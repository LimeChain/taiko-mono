package proposer

import (
	"bytes"
	"context"
	"encoding/hex"

	"fmt"
	"math/big"
	"strconv"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"

	"github.com/ethereum/go-ethereum/rlp"
	"github.com/urfave/cli/v2"
	"golang.org/x/sync/errgroup"

	"github.com/taikoxyz/taiko-mono/packages/taiko-client/bindings"
	"github.com/taikoxyz/taiko-mono/packages/taiko-client/bindings/encoding"
	"github.com/taikoxyz/taiko-mono/packages/taiko-client/types"

	"github.com/taikoxyz/taiko-mono/packages/taiko-client/internal/metrics"
	"github.com/taikoxyz/taiko-mono/packages/taiko-client/internal/utils"
	"github.com/taikoxyz/taiko-mono/packages/taiko-client/pkg/rpc"
	selector "github.com/taikoxyz/taiko-mono/packages/taiko-client/proposer/prover_selector"
	builder "github.com/taikoxyz/taiko-mono/packages/taiko-client/proposer/transaction_builder"
)

// EligibleSlot represents a slot that a proposer
// may be eligible to participate in.
type EligibleSlot struct {
	Slot       uint64
	IsPrimary  bool
	IsFallback bool
}

// Proposer is responsible for proposing new transactions
// from the L2 execution engine at regular intervals.
type Proposer struct {
	// configurations
	*Config

	// RPC clients
	RPC *rpc.Client

	// State variables for proposer duties
	proposerDutiesSlot uint64
	eligibleSlots      []*EligibleSlot

	l1HeadSlotTimer *time.Timer

	validatorPublicKeyHex string
	validatorAddress      common.Address

	// Private keys and account addresses
	proposerAddress common.Address

	tiers    []*rpc.TierProviderTierWithID
	tierFees []encoding.TierFee

	// Prover selector
	proverSelector selector.ProverSelector

	// Transaction builder
	txBuilder builder.ProposeBlockTransactionBuilder

	// Protocol configurations
	protocolConfigs *bindings.TaikoDataConfig

	lastProposedAt time.Time

	txmgr *builder.SimpleTxManager

	ctx context.Context
	wg  sync.WaitGroup
}

// InitFromCli New initializes the given proposer instance based on the command line flags.
func (p *Proposer) InitFromCli(ctx context.Context, c *cli.Context) error {
	cfg, err := NewConfigFromCliContext(c)
	if err != nil {
		return err
	}

	return p.InitFromConfig(ctx, cfg)
}

// InitFromConfig initializes the proposer instance based on the given configurations.
func (p *Proposer) InitFromConfig(ctx context.Context, cfg *Config) (err error) {
	p.proposerAddress = crypto.PubkeyToAddress(cfg.L1ProposerPrivKey.PublicKey)
	p.ctx = ctx
	p.Config = cfg
	p.lastProposedAt = time.Now()

	// RPC clients
	if p.RPC, err = rpc.NewClient(p.ctx, cfg.ClientConfig); err != nil {
		return fmt.Errorf("initialize rpc clients error: %w", err)
	}

	if _, ok := p.RPC.L1MevBoost.(*rpc.MevBoostClient); ok {
		if p.validatorPublicKeyHex, err = p.RPC.L1MevBoost.GetValidatorPubKeyHex(); err != nil {
			return fmt.Errorf("failed to get validator pubkey hex error: %w", err)
		}

		p.validatorAddress = GetAddressFromBlsPublikKeyHex(p.validatorPublicKeyHex[2:])
	}

	// Protocol configs
	protocolConfigs, err := p.RPC.TaikoL1.GetConfig(&bind.CallOpts{Context: ctx})
	if err != nil {
		return fmt.Errorf("failed to get protocol configs: %w", err)
	}
	p.protocolConfigs = &protocolConfigs

	log.Info("Protocol configs", "configs", p.protocolConfigs)

	if p.tiers, err = p.RPC.GetTiers(ctx); err != nil {
		return err
	}
	if err := p.initTierFees(); err != nil {
		return err
	}

	if p.txmgr, err = builder.NewSimpleTxManager(
		"proposer",
		log.Root(),
		&metrics.TxMgrMetrics,
		*cfg.TxmgrConfigs,
	); err != nil {
		return err
	}

	if p.proverSelector, err = selector.NewETHFeeEOASelector(
		&protocolConfigs,
		p.RPC,
		p.proposerAddress,
		cfg.TaikoL1Address,
		cfg.ProverSetAddress,
		p.tierFees,
		cfg.TierFeePriceBump,
		cfg.ProverEndpoints,
		cfg.MaxTierFeePriceBumps,
	); err != nil {
		return err
	}

	if cfg.BlobAllowed {
		p.txBuilder = builder.NewBlobTransactionBuilder(
			p.RPC,
			p.L1ProposerPrivKey,
			p.proverSelector,
			p.Config.L1BlockBuilderTip,
			cfg.TaikoL1Address,
			cfg.ProverSetAddress,
			cfg.L2SuggestedFeeRecipient,
			cfg.ProposeBlockTxGasLimit,
			cfg.ExtraData,
		)
	} else {
		p.txBuilder = builder.NewCalldataTransactionBuilder(
			p.RPC,
			p.L1ProposerPrivKey,
			p.proverSelector,
			p.Config.L1BlockBuilderTip,
			cfg.L2SuggestedFeeRecipient,
			cfg.TaikoL1Address,
			cfg.ProverSetAddress,
			cfg.ProposeBlockTxGasLimit,
			cfg.ExtraData,
		)
	}

	// TODO: Uncomment once BLS signatures are supported (after EIP-2573) in solidity
	// pubKeyBytes, err := hex.DecodeString(p.validatorPublicKeyHex[2:])
	// if err != nil {
	// 	return fmt.Errorf("failed to decode validator public key: %w", err)
	// }

	// TODO: Remove once BLS signatures are supported (after EIP-2573) in solidity
	pubKeyBytes := make([]byte, 48)
	pubKeyBytes[31] = 0x01

	// Verify if the validator is eligible to propose
	isEligible, err := p.RPC.SequencerRegistry.IsEligible(&bind.CallOpts{Context: ctx}, pubKeyBytes)
	if err != nil {
		return fmt.Errorf("failed to get is eligible error: %w", err)
	}

	if !isEligible {
		return fmt.Errorf("the validator is not eligible signer, pubkey: %s", p.validatorPublicKeyHex)
	}

	// Wait until L2 execution engine is synced at first.
	if err := p.RPC.WaitTillL2ExecutionEngineSynced(ctx); err != nil {
		return fmt.Errorf("failed to wait until L2 execution engine synced: %w", err)
	}

	log.Info("Proposer initialized successfully")

	return nil
}

// Start starts the proposer's main loop.
func (p *Proposer) Start() error {
	go p.eventLoop()
	return nil
}

// Close closes the proposer instance.
func (p *Proposer) Close(_ context.Context) {
	p.wg.Wait()
}

func (p *Proposer) eventLoop() {
	p.wg.Add(1)

	defer func() {
		p.l1HeadSlotTimer.Stop()
		p.wg.Done()
	}()

	for {
		p.updateL1HeadSlotTicker()

		select {
		case <-p.ctx.Done():
			return
		case <-p.l1HeadSlotTimer.C:
			l1HeadSlot := p.RPC.L1Beacon.GetL1HeadSlot()

			log.Info("L1 head slot", "slot", l1HeadSlot)

			// Sync L1 proposer duties
			updated, err := p.syncL1ProposerDuties(l1HeadSlot)
			if err != nil {
				log.Error("Sync L1 proposer duties operation error", "error", err)
				continue
			}

			if updated {
				log.Info("L1 proposer duties updated successfully")

				l1Slots := make([]uint64, 0)
				for _, slot := range p.eligibleSlots {
					if slot.IsPrimary || slot.IsFallback {
						l1Slots = append(l1Slots, slot.Slot)
					}
				}

				// Update L2 config and slots sends the eligible slots
				// along with L1 params to the L2 execution engine.
				_, err := p.RPC.UpdateL2ConfigAndSlots(
					p.ctx,
					p.RPC.L1Beacon.GetGenesisTimestamp(),
					l1Slots,
					p.protocolConfigs.BlockMaxGasLimit,
					rpc.BlockMaxTxListBytes,
					p.proposerAddress,
					p.LocalAddresses,
					p.MaxProposedTxListsPerEpoch,
				)
				if err != nil {
					log.Error("Update slots and config params", "error", err)
					continue
				}

				log.Debug("Updated L2 config and slots successfully", "L1 eligible slots", l1Slots)
			}

			log.Debug(
				"Checking if the proposer is eligible for the slots",
				"slot",
				l1HeadSlot+1,
				"slot",
				l1HeadSlot+2,
			)

			nonce := p.txmgr.GetNonce()
			if nonce != nil {
				txmgrNonce := *nonce
				netNonce, err := p.RPC.L1.PendingNonceAt(p.ctx, p.proposerAddress)
				if err != nil {
					log.Error("Failed to get proposer nonce", "error", err)
					break
				}

				// Reset the nonce if the proposer nonce is ahead of the network nonce.
				if AbsInt(int64(txmgrNonce)-int64(netNonce)) > 1 {
					log.Warn("Proposer nonce is ahead of the network nonce", "network nonce", netNonce, "txmng nonce", txmgrNonce)
					p.txmgr.ResetNonce()
				}
			}

			eligibleSlot, err := p.isEligibleForL1Slot(l1HeadSlot + 1)
			if err != nil {
				log.Error("Failed to check if the proposer is eligible for the L1 slot", "error", err)
				continue
			}
			// If the proposer is a fallback for the slot, sends a propose operation to the taiko L11 contract.
			if eligibleSlot.IsFallback {
				metrics.ProposerProposeEpochCounter.Add(1)
				log.Debug("Proposer is fallback for the L1 slot", "slot", l1HeadSlot+1)

				// Attempt a propose operation
				if err := p.ProposeOp(p.ctx); err != nil {
					log.Error("Propose operation error", "error", err)
					continue
				}
			}

			eligibleSlot, err = p.isEligibleForL1Slot(l1HeadSlot + 2)
			if err != nil {
				log.Error("Failed to check if the proposer is eligible for the L1 slot", "error", err)
				continue
			}
			// If the proposer is a primary for the slot, propose the block to the mev-boost.
			if eligibleSlot.IsPrimary {
				p.preconfDelay(l1HeadSlot)

				metrics.ProposerProposeEpochCounter.Add(1)
				log.Debug("Proposer is primary for the L1 slot", "slot", l1HeadSlot+2)

				// Attempt a preconf operation
				if err := p.PreconfOp(p.ctx); err != nil {
					log.Error("Preconf operation error", "error", err)
					continue
				}
			}
		}
	}
}

// preconfDelay waits for a specified delay before performing a pre-confirmation operation.
func (p *Proposer) preconfDelay(l1HeadSlot uint64) {
	currentSlotTS := p.RPC.L1Beacon.GetTimestampBySlot(l1HeadSlot)
	secsInSlot := time.Now().UTC().Unix() - int64(currentSlotTS)
	durationSec := int64(p.PreconfDelay) - secsInSlot
	if durationSec > 0 {
		time.Sleep(time.Duration(durationSec) * time.Second)
	}
}

// syncL1ProposerDuties checks if the L1 proposer duties need to be updated based on the current slot.
// If required, it updates the eligible slots for the proposer.
func (p *Proposer) syncL1ProposerDuties(l1Slot uint64) (updated bool, err error) {
	updated = false

	if p.shouldUpdateDuties(l1Slot) {
		log.Info(
			"Start synching L1 proposer duties",
			"headSlot", l1Slot,
		)

		updated, err = p.updateProposerDuties(p.ctx)
		if err != nil {
			return false, fmt.Errorf("failed to update proposer duties: %w", err)
		}
	}

	p.proposerDutiesSlot = l1Slot

	return updated, nil
}

// shouldUpdateDuties checks if the proposer duties should be updated based on the current slot.
func (p *Proposer) shouldUpdateDuties(headSlot uint64) bool {
	if p.proposerDutiesSlot == 0 {
		return true
	}
	lastProposerDutyDistance := headSlot - p.proposerDutiesSlot
	return headSlot%p.ProposerDutiesUpdateFreq == 0 ||
		lastProposerDutyDistance >= p.ProposerDutiesUpdateFreq
}

func (p *Proposer) getBlockBySlot(slot uint64) (uint64, error) {
	timestamp := p.RPC.L1Beacon.GetTimestampBySlot(slot)

	// Start searching from the latest block
	latestBlock, err := p.RPC.L1.BlockByNumber(context.Background(), nil)
	if err != nil {
		return 0, fmt.Errorf("failed to fetch the latest block: %w", err)
	}

	var low, high uint64 = 0, latestBlock.NumberU64()
	for low <= high {
		mid := (low + high) / 2
		block, err := p.RPC.L1.BlockByNumber(context.Background(), big.NewInt(int64(mid)))
		if err != nil {
			return 0, fmt.Errorf("failed to fetch block number %d: %w", mid, err)
		}

		if block.Time() < timestamp {
			low = mid + 1
		} else if block.Time() > timestamp {
			high = mid - 1
		} else {
			return block.NumberU64(), nil
		}
	}

	return low, nil
}

func (p *Proposer) updateProposerDuties(ctx context.Context) (bool, error) {
	proposerDuties, err := p.RPC.L1Beacon.GetNextProposerDuties(
		ctx,
		p.RPC.L1Beacon.GetL1HeadSlot(),
		p.MaxProposerDutiesSlots,
	)
	if err != nil {
		return false, fmt.Errorf("failed to get next proposer duties: %w", err)
	}

	if len(proposerDuties) == 0 {
		log.Warn("Empty proposer duties")
		return false, nil
	}

	var wg sync.WaitGroup
	errors := make(chan error, len(proposerDuties))
	var mu sync.Mutex

	height, err := p.getBlockBySlot(p.RPC.L1Beacon.GetL1HeadSlot())
	if err != nil {
		return false, fmt.Errorf("failed to get block by slot: %w", err)
	}

	var eligibleSlots []*EligibleSlot

	for _, duty := range proposerDuties {
		wg.Add(1)
		go func(duty *types.ProposerDuty) {
			defer wg.Done()
			eligibleSlot, err := p.handleDuty(
				ctx,
				duty,
				p.RPC.L1Beacon.GetL1HeadSlot(),
				height,
			)
			if err != nil {
				errors <- err
				return
			}

			mu.Lock()
			if eligibleSlot != nil {
				eligibleSlots = append(eligibleSlots, eligibleSlot)
			}
			mu.Unlock()
		}(duty)
	}

	wg.Wait()
	close(errors)

	for err := range errors {
		if err != nil {
			return false, err
		}
	}

	p.eligibleSlots = eligibleSlots

	return true, nil
}

// handleDuty processes a proposer duty for a specific L1 slot and checks if the current proposer is
// eligible to propose a block in that slot, either as the primary proposer or as a fallback.
func (p *Proposer) handleDuty(
	ctx context.Context,
	duty *types.ProposerDuty,
	l1HeadSlot uint64,
	l1Height uint64,
) (*EligibleSlot, error) {
	// Convert the duty slot from string to integer
	dutySlot, err := strconv.Atoi(duty.Slot)
	if err != nil {
		return nil, fmt.Errorf("failed to convert L1 duty slot to integer: %w", err)
	}

	// Create a new EligibleSlot with the duty slot
	eligibleSlot := &EligibleSlot{
		Slot:       uint64(dutySlot),
		IsPrimary:  false,
		IsFallback: false,
	}

	// Check if the duty is assigned to this proposer based on the public key
	if duty.PubKey != p.validatorPublicKeyHex {
		// Decode the duty's public key
		pubKeyBytes, err := hex.DecodeString(duty.PubKey[2:])
		if err != nil {
			return nil, fmt.Errorf("failed to decode duty public key: %w", err)
		}

		// Check if another validator is eligible to propose
		isAnotherValEligible, err := p.RPC.SequencerRegistry.IsEligible(&bind.CallOpts{Context: ctx}, pubKeyBytes)
		if err != nil {
			return nil, fmt.Errorf("failed to get is eligible: %w", err)
		}

		// If another validator is eligible, skip this duty
		if isAnotherValEligible {
			return eligibleSlot, nil
		}

		// Calculate the block number corresponding to the fallback signer for this slot
		blockNum := l1Height + uint64(dutySlot) - l1HeadSlot

		// Check if the fallback signer for the calculated block number matches the proposer's address
		fallbackSigner, err := p.RPC.SequencerRegistry.FallbackSigner(
			&bind.CallOpts{Context: ctx},
			big.NewInt(int64(blockNum)),
		)

		if err != nil {
			return nil, fmt.Errorf("failed to get FallbackSigner: %w", err)
		}

		// If the fallback signer matches this proposer's address, mark it as fallback
		if fallbackSigner == p.validatorAddress {
			eligibleSlot.IsFallback = true
		}
	} else {
		eligibleSlot.IsPrimary = true
	}

	// Check if the dutySlot is in the same epoch as the (dutySlot - 2).
	// This is a temp fix for the commit-boost error "can only set constraints for current epoch".
	slotsPerEpoch := p.RPC.L1Beacon.GetSlotsPerEpoch()
	prevL1HeadSlotEpoch := (uint64(dutySlot) - 2) / slotsPerEpoch
	l1SlotEpoch := uint64(dutySlot) / slotsPerEpoch
	// If the duty slot is not in the same epoch, switch the role to fallback
	if prevL1HeadSlotEpoch != l1SlotEpoch && eligibleSlot.IsPrimary {
		eligibleSlot.IsPrimary = false
		eligibleSlot.IsFallback = true
	}

	return eligibleSlot, nil
}

// isEligibleForL1Slot checks if the proposer is eligible to propose a block for the given L1 slot.
// It looks for the slot in the list of eligible slots and returns the corresponding eligibility status.
func (p *Proposer) isEligibleForL1Slot(slot uint64) (*EligibleSlot, error) {
	// Ensure there are eligible slots available
	if len(p.eligibleSlots) == 0 {
		return nil, fmt.Errorf("no eligible slots found")
	}

	// Loop through the eligible slots to find a matching slot
	for _, eligibleSlot := range p.eligibleSlots {
		if slot == eligibleSlot.Slot {
			return eligibleSlot, nil
		}
	}

	// If no matching slot is found, return an error
	return nil, fmt.Errorf("eligible slot not found")
}

// fetchTxListToPropose fetches prebuilt list of transactions from L2 execution engine.
func (p *Proposer) fetchTxListToPropose(l1Slot uint64) ([]ethTypes.Transactions, error) {
	preBuiltTxLists, err := p.RPC.FetchTxList(p.ctx, l1Slot)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch transaction pool content: %w", err)
	}

	txLists := []ethTypes.Transactions{}
	for _, prebuiltTxList := range preBuiltTxLists {
		txLists = append(txLists, prebuiltTxList.TxList)
	}
	if len(txLists) == 0 {
		log.Info(
			"Pool content is empty, proposing an empty block",
			"lastProposedAt", p.lastProposedAt,
		)
		txLists = append(txLists, ethTypes.Transactions{})
	}

	// If LocalAddressesOnly is set, filter the transactions by the local addresses.
	if p.LocalAddressesOnly {
		var (
			localTxsLists []ethTypes.Transactions
			signer        = ethTypes.LatestSignerForChainID(p.RPC.L2.ChainID)
		)
		for _, txs := range txLists {
			var filtered ethTypes.Transactions
			for _, tx := range txs {
				sender, err := ethTypes.Sender(signer, tx)
				if err != nil {
					return nil, err
				}

				for _, localAddress := range p.LocalAddresses {
					if sender == localAddress {
						filtered = append(filtered, tx)
					}
				}
			}

			if filtered.Len() != 0 {
				localTxsLists = append(localTxsLists, filtered)
			}
		}
		txLists = localTxsLists
	}

	log.Info("Transaction lists", "size", len(txLists))

	return txLists, nil
}

// updateL1HeadSlotTicker updates or resets the timer that controls when the next L1 head slot is processed.
// It calculates the time until the next L1 slot is expected and sets a timer for the next action.
func (p *Proposer) updateL1HeadSlotTicker() {
	// Stop the current timer if it's already running, to avoid conflicts.
	if p.l1HeadSlotTimer != nil {
		p.l1HeadSlotTimer.Stop()
	}

	// Get the timestamp of the next L1 slot by querying the beacon chain.
	nextSlotTS := p.RPC.L1Beacon.GetTimestampBySlot(p.RPC.L1Beacon.GetL1HeadSlot() + 1)

	// Calculate the number of seconds remaining until the next slot occurs.
	durationSec := int64(nextSlotTS) - time.Now().UTC().Unix()

	// Get the number of seconds in a slot from the beacon chain configuration.
	secPerSlot := p.RPC.L1Beacon.GetSecondsPerSlot()

	// If the calculated duration is negative or zero, adjust the duration to the next slot.
	// This can happen if the system is slightly behind or if the next slot has already started.
	if durationSec <= 0 {
		durationSec = int64(secPerSlot) - AbsInt(durationSec)
	}

	// Log the time remaining until the next L1 head slot for debugging purposes.
	log.Debug(
		"Setting L1 head slot timer",
		"next L1 head slot",
		p.RPC.L1Beacon.GetL1HeadSlot()+1,
		"time to slot",
		durationSec,
	)

	// Create a new timer for the next L1 slot, based on the calculated duration.
	// This timer will trigger when the next slot starts, allowing the proposer to take the necessary actions.
	duration := time.Duration(durationSec) * time.Second
	p.l1HeadSlotTimer = time.NewTimer(duration)
}

// ProposeOp performs a proposing operation, fetching transactions
// from L2 execution engine's tx pool, splitting them by proposing constraints,
// and then proposing them to TaikoL1 contract.
func (p *Proposer) ProposeOp(ctx context.Context) error {
	log.Info(
		"Start fetching L2 execution engine's transaction pool content",
		"lastProposedAt", p.lastProposedAt,
	)

	txLists, err := p.fetchTxListToPropose(p.RPC.L1Beacon.GetL1HeadSlot() + 1)
	if err != nil {
		return err
	}

	// If the pool content is empty, return.
	if len(txLists) == 0 {
		return nil
	}

	log.Warn("Tx list content", "txs", txLists)

	g, gCtx := errgroup.WithContext(ctx)
	// Propose all L2 transactions lists.
	for _, txsList := range txLists[:utils.Min(p.MaxProposedTxListsPerEpoch, uint64(len(txLists)))] {
		txs := txsList
		g.Go(func() error {
			txListBytes, err := rlp.EncodeToBytes(txs)
			if err != nil {
				return fmt.Errorf("failed to encode transactions: %w", err)
			}

			if err := p.ProposeTxList(gCtx, txListBytes, uint(txs.Len())); err != nil {
				return err
			}

			p.lastProposedAt = time.Now()
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		return err
	}

	return nil
}

// ProposeTxList proposes the given transactions list to TaikoL1 smart contract.
func (p *Proposer) ProposeTxList(
	ctx context.Context,
	txListBytes []byte,
	txNum uint,
) error {
	compressedTxListBytes, err := utils.Compress(txListBytes)
	if err != nil {
		return err
	}

	txCandidate, err := p.txBuilder.Build(
		ctx,
		p.tierFees,
		p.IncludeParentMetaHash,
		compressedTxListBytes,
	)
	if err != nil {
		log.Warn("Failed to build TaikoL1.proposeBlock transaction", "error", encoding.TryParsingCustomError(err))
		return err
	}

	err = p.txmgr.Send(ctx, *txCandidate)
	if err != nil {
		log.Warn("Failed to send TaikoL1.proposeBlock transaction", "error", encoding.TryParsingCustomError(err))
		return err
	}

	log.Info("📝 Propose transactions succeeded", "txs", txNum)

	metrics.ProposerProposedTxListsCounter.Add(1)
	metrics.ProposerProposedTxsCounter.Add(float64(txNum))

	return nil
}

// PreconfOp performs a preconfirmation operation, fetching transactions and
// proposing them to the mev-boost.
func (p *Proposer) PreconfOp(ctx context.Context) error {
	log.Info(
		"Start fetching L2 execution engine's transaction pool content",
		"lastProposedAt", p.lastProposedAt,
	)

	txLists, err := p.fetchTxListToPropose(p.RPC.L1Beacon.GetL1HeadSlot() + 2)
	if err != nil {
		return err
	}

	// If the pool content is empty, return.
	if len(txLists) == 0 {
		return nil
	}

	log.Warn("Tx list content", "txs", txLists)

	g, gCtx := errgroup.WithContext(ctx)
	// Propose all L2 transactions lists.
	for _, txsList := range txLists[:utils.Min(p.MaxProposedTxListsPerEpoch, uint64(len(txLists)))] {
		txs := txsList
		g.Go(func() error {
			txListBytes, err := rlp.EncodeToBytes(txs)
			if err != nil {
				return fmt.Errorf("failed to encode transactions: %w", err)
			}

			if err := p.PreconfTxList(gCtx, txListBytes, uint(txs.Len())); err != nil {
				return err
			}

			p.lastProposedAt = time.Now()
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		return err
	}

	return nil
}

// PreconfTxList proposes the given transactions list to mev-boost.
func (p *Proposer) PreconfTxList(
	ctx context.Context,
	txListBytes []byte,
	txNum uint,
) error {
	compressedTxListBytes, err := utils.Compress(txListBytes)
	if err != nil {
		return err
	}

	txCandidate, err := p.txBuilder.Build(
		ctx,
		p.tierFees,
		p.IncludeParentMetaHash,
		compressedTxListBytes,
	)
	if err != nil {
		log.Warn("Failed to build TaikoL1.proposeBlock transaction", "error", encoding.TryParsingCustomError(err))
		return err
	}

	tx, err := p.txmgr.CraftTx(ctx, *txCandidate)
	if err != nil {
		log.Warn("Failed to craft TaikoL1.preconfBlock transaction", "error", encoding.TryParsingCustomError(err))
		return err
	}
	newTx, err := p.txmgr.IncreaseGasPrice(ctx, tx)
	if err != nil {
		log.Error("unable to increase gas", "err", err)
		return err
	}

	log.Debug("Setting validator mev boost constraints", "slot", p.RPC.L1Beacon.GetL1HeadSlot()+2)

	if err = p.RPC.L1MevBoost.SetConstraints(p.RPC.L1Beacon.GetL1HeadSlot()+2, newTx); err != nil {
		p.txmgr.DecNonce()
		return fmt.Errorf("failed to set validator mev boost constraints: %w", err)
	}

	log.Info("📝 Preconf tx sent to the mev-boost", "txs", txNum)

	metrics.ProposerProposedTxListsCounter.Add(1)
	metrics.ProposerProposedTxsCounter.Add(float64(txNum))

	return nil
}

// Name returns the application name.
func (p *Proposer) Name() string {
	return "proposer"
}

// initTierFees initializes the proving fees for every proof tier configured in the protocol for the proposer.
func (p *Proposer) initTierFees() error {
	for _, tier := range p.tiers {
		log.Info(
			"Protocol tier",
			"id", tier.ID,
			"name", string(bytes.TrimRight(tier.VerifierName[:], "\x00")),
			"validityBond", utils.WeiToEther(tier.ValidityBond),
			"contestBond", utils.WeiToEther(tier.ContestBond),
			"provingWindow", tier.ProvingWindow,
			"cooldownWindow", tier.CooldownWindow,
		)

		switch tier.ID {
		case encoding.TierOptimisticID:
			p.tierFees = append(p.tierFees, encoding.TierFee{Tier: tier.ID, Fee: p.OptimisticTierFee})
		case encoding.TierSgxID:
			p.tierFees = append(p.tierFees, encoding.TierFee{Tier: tier.ID, Fee: p.SgxTierFee})
		case encoding.TierGuardianMinorityID:
			p.tierFees = append(p.tierFees, encoding.TierFee{Tier: tier.ID, Fee: common.Big0})
		case encoding.TierGuardianMajorityID:
			// Guardian prover should not charge any fee.
			p.tierFees = append(p.tierFees, encoding.TierFee{Tier: tier.ID, Fee: common.Big0})
		default:
			return fmt.Errorf("unknown tier: %d", tier.ID)
		}
	}

	return nil
}
