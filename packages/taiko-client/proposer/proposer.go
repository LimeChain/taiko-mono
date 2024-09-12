package proposer

import (
	"bytes"
	"context"

	"encoding/hex"
	"fmt"
	"math/big"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/log"

	"github.com/ethereum/go-ethereum/rlp"
	"github.com/urfave/cli/v2"
	"golang.org/x/sync/errgroup"

	"github.com/taikoxyz/taiko-mono/packages/taiko-client/bindings"
	"github.com/taikoxyz/taiko-mono/packages/taiko-client/bindings/encoding"

	"github.com/taikoxyz/taiko-mono/packages/taiko-client/internal/metrics"
	"github.com/taikoxyz/taiko-mono/packages/taiko-client/internal/utils"
	"github.com/taikoxyz/taiko-mono/packages/taiko-client/pkg/rpc"
	selector "github.com/taikoxyz/taiko-mono/packages/taiko-client/proposer/prover_selector"
	builder "github.com/taikoxyz/taiko-mono/packages/taiko-client/proposer/transaction_builder"
)

// Proposer keep proposing new transactions from L2 execution engine at a fixed interval.
type Proposer struct {
	// configurations
	*Config

	// RPC clients
	rpc *rpc.Client

	l1HeadSlot     atomic.Value
	l1HeadSlotFeed event.Feed

	proposerDutiesSlot atomic.Value
	proposerDuties     []*rpc.ProposerDuty

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
	if p.rpc, err = rpc.NewClient(p.ctx, cfg.ClientConfig); err != nil {
		return fmt.Errorf("initialize rpc clients error: %w", err)
	}

	if p.validatorPublicKeyHex, err = p.rpc.L1MevBoost.GetValidatorPubKeyHex(); err != nil {
		return fmt.Errorf("failed to get validator pubkey hex error: %w", err)
	}

	p.validatorAddress = GetAddressFromBlsPublikKeyHex(p.validatorPublicKeyHex[2:])

	// Protocol configs
	protocolConfigs, err := p.rpc.TaikoL1.GetConfig(&bind.CallOpts{Context: ctx})
	if err != nil {
		return fmt.Errorf("failed to get protocol configs: %w", err)
	}
	p.protocolConfigs = &protocolConfigs

	log.Info("Protocol configs", "configs", p.protocolConfigs)

	if p.tiers, err = p.rpc.GetTiers(ctx); err != nil {
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
		p.rpc,
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
			p.rpc,
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
			p.rpc,
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

	isEligible, err := p.rpc.SequencerRegistry.IsEligible(&bind.CallOpts{Context: ctx}, pubKeyBytes)
	if err != nil {
		return fmt.Errorf("failed to get is eligible error: %w", err)
	}

	if !isEligible {
		return fmt.Errorf("the validator is not eligible signer, pubkey: %s", p.validatorPublicKeyHex)
	}

	// Wait until L2 execution engine is synced at first.
	if err := p.rpc.WaitTillL2ExecutionEngineSynced(ctx); err != nil {
		return fmt.Errorf("failed to wait until L2 execution engine synced: %w", err)
	}

	p.updateL1HeadSlot(p.ctx)

	return nil
}

// Start starts the proposer's main loop.
func (p *Proposer) Start() error {
	go p.syncL1()
	go p.proposeBlock()
	go p.eventLoop()
	return nil
}

// Close closes the proposer instance.
func (p *Proposer) Close(_ context.Context) {
	p.wg.Wait()
}

func (p *Proposer) syncL1() {
	p.wg.Add(1)

	var (
		l1HeadSlotFeedCh1  = make(chan uint64, 10)
		l1HeadSlotFeedSub1 = p.l1HeadSlotFeed.Subscribe(l1HeadSlotFeedCh1)
	)

	defer func() {
		l1HeadSlotFeedSub1.Unsubscribe()
		p.wg.Done()
	}()

	for {
		select {
		case <-p.ctx.Done():
			return
		case l1HeadSlot := <-l1HeadSlotFeedCh1:
			updated, err := p.syncL1ProposerDuties(p.ctx, l1HeadSlot)
			if err != nil {
				log.Error("Sync L1 proposer duties operation error", "error", err)
				continue
			}

			if updated {
				log.Info("L1 proposer duties updated successfully")

				_, err := p.rpc.UpdateL2ConfigAndSlots(
					p.ctx,
					p.rpc.L1Beacon.GetGenesisTimestamp(),
					p.getL1ValidatorAssignedSlots(),
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
			}
		}
	}
}

func (p *Proposer) proposeBlock() {
	p.wg.Add(1)

	var (
		l1HeadSlotFeedCh2  = make(chan uint64, 10)
		l1HeadSlotFeedSub2 = p.l1HeadSlotFeed.Subscribe(l1HeadSlotFeedCh2)
	)

	defer func() {
		l1HeadSlotFeedSub2.Unsubscribe()
		p.wg.Done()
	}()

	for {
		select {
		case <-p.ctx.Done():
			return
		case <-l1HeadSlotFeedCh2:
			time.Sleep(p.ProposeDelay * time.Second)

			metrics.ProposerProposeEpochCounter.Add(1)

			// Attempt a proposing operation
			if err := p.ProposeOp(p.ctx); err != nil {
				log.Error("Proposing operation error", "error", err)
				continue
			}
		}
	}
}

// eventLoop starts the main loop of Taiko proposer.
func (p *Proposer) eventLoop() {
	p.wg.Add(1)

	defer p.wg.Done()

	for {
		select {
		case <-p.ctx.Done():
			return
		default:
			time.Sleep(500 * time.Millisecond)

			updated, l1HeadSlot := p.updateL1HeadSlot(p.ctx)
			if updated {
				p.l1HeadSlotFeed.Send(l1HeadSlot)
			}
		}
	}
}

// updateL1HeadSlot updates the L1 head slot concurrent safely.
func (p *Proposer) updateL1HeadSlot(ctx context.Context) (bool, uint64) {
	l1HeadSlot := p.rpc.L1Beacon.GetL1HeadSlot()

	prevL1HeadSlot := p.l1HeadSlot.Swap(l1HeadSlot)

	if prevL1HeadSlot == l1HeadSlot {
		return false, l1HeadSlot
	}

	log.Info("L1 head slot updated", "headSlot", l1HeadSlot)

	return true, l1HeadSlot
}

func (p *Proposer) syncL1ProposerDuties(ctx context.Context, l1HeadSlot uint64) (updated bool, err error) {
	updated = false

	if p.shouldUpdateDuties(l1HeadSlot) {
		log.Info(
			"Start synching L1 proposer duties",
			"headSlot", l1HeadSlot,
		)

		updated, err = p.updateProposerDuties(p.ctx)
		if err != nil {
			return false, fmt.Errorf("failed to update proposer duties: %w", err)
		}
	}

	p.proposerDutiesSlot.Store(l1HeadSlot)

	return updated, nil
}

func (p *Proposer) shouldUpdateDuties(headSlot uint64) bool {
	proposerDutiesSlot := p.proposerDutiesSlot.Load()
	if proposerDutiesSlot == nil {
		return true
	}
	lastProposerDutyDistance := headSlot - proposerDutiesSlot.(uint64)
	return headSlot%p.ProposerDutiesUpdateFreq == 0 ||
		lastProposerDutyDistance >= p.ProposerDutiesUpdateFreq
}

func (p *Proposer) getL1ValidatorAssignedSlots() []uint64 {
	assignedSlots := make([]uint64, 0)
	for _, duty := range p.proposerDuties {
		slot, err := strconv.Atoi(duty.Slot)
		if err != nil {
			log.Error("Failed to convert slot to integer", "slot", duty.Slot, "error", err)
			continue
		}
		if duty.PubKey == p.validatorPublicKeyHex {
			assignedSlots = append(assignedSlots, uint64(slot))
		}
	}

	return assignedSlots
}

func (p *Proposer) getBlockBySlot(slot uint64) (uint64, error) {
	timestamp := p.rpc.L1Beacon.GetTimestampBySlot(slot)

	// Start searching from the latest block
	latestBlock, err := p.rpc.L1.BlockByNumber(context.Background(), nil)
	if err != nil {
		return 0, fmt.Errorf("failed to fetch the latest block: %w", err)
	}

	var low, high uint64 = 0, latestBlock.NumberU64()
	for low <= high {
		mid := (low + high) / 2
		block, err := p.rpc.L1.BlockByNumber(context.Background(), big.NewInt(int64(mid)))
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
	l1HeadSlot := p.l1HeadSlot.Load().(uint64)

	proposerDuties, err := p.rpc.L1Beacon.GetNextProposerDuties(ctx, l1HeadSlot, p.MaxProposerDutiesSlots)
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

	height, err := p.getBlockBySlot(l1HeadSlot)

	if err != nil {
		return false, fmt.Errorf("failed to get block by slot: %w", err)
	}

	for i, duty := range proposerDuties {
		wg.Add(1)
		go func(duty *rpc.ProposerDuty, index int) {
			defer wg.Done()
			updatedDuty, err := p.handleDuty(ctx, duty, index, l1HeadSlot, height)
			if err != nil {
				errors <- err
				return
			}

			mu.Lock()
			proposerDuties[index] = updatedDuty
			mu.Unlock()
		}(duty, i)
	}

	wg.Wait()
	close(errors)

	for err := range errors {
		if err != nil {
			return false, err
		}
	}

	p.proposerDuties = proposerDuties

	return true, nil
}

func (p *Proposer) handleDuty(ctx context.Context, duty *rpc.ProposerDuty, index int, headSlot uint64, height uint64) (*rpc.ProposerDuty, error) {
	if duty.PubKey != p.validatorPublicKeyHex {
		pubKeyBytes, err := hex.DecodeString(duty.PubKey[2:])
		if err != nil {
			return duty, fmt.Errorf("failed to decode duty public key: %w", err)
		}

		isEligible, err := p.rpc.SequencerRegistry.IsEligible(&bind.CallOpts{Context: ctx}, pubKeyBytes)
		if err != nil {
			return duty, fmt.Errorf("failed to get is eligible: %w", err)
		}

		if !isEligible {
			dutySlot, err := strconv.Atoi(duty.Slot)
			if err != nil {
				return duty, fmt.Errorf("failed to convert duty slot to integer: %w", err)
			}

			blockNum := height + uint64(dutySlot) - headSlot

			fallbackSigner, err := p.rpc.SequencerRegistry.FallbackSigner(&bind.CallOpts{Context: ctx}, big.NewInt(int64(blockNum)))
			if err != nil {
				return duty, fmt.Errorf("failed to get FallbackSigner: %w", err)
			}

			if fallbackSigner == p.validatorAddress {
				duty.PubKey = p.validatorPublicKeyHex
			}
		}
	}

	return duty, nil
}

// fetchTxListToPropose fetches prebuilt list of transactions from L2 execution engine.
func (p *Proposer) fetchTxListToPropose() ([]types.Transactions, error) {
	preBuiltTxLists, err := p.rpc.FetchTxList(p.ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch transaction pool content: %w", err)
	}

	txLists := []types.Transactions{}
	for i, prebuiltTxList := range preBuiltTxLists {
		// Filter the pool content if the filterPoolContent flag is set.
		if prebuiltTxList.EstimatedGasUsed < p.MinGasUsed && prebuiltTxList.BytesLength < p.MinTxListBytes {
			log.Info(
				"Pool content skipped",
				"index", i,
				"estimatedGasUsed", prebuiltTxList.EstimatedGasUsed,
				"minGasUsed", p.MinGasUsed,
				"bytesLength", prebuiltTxList.BytesLength,
				"minBytesLength", p.MinTxListBytes,
			)
			break
		}
		txLists = append(txLists, prebuiltTxList.TxList)
	}
	// If the pool content is empty and the checkPoolContent flag is not set, return an empty list.
	if len(txLists) == 0 {
		log.Info(
			"Pool content is empty, proposing an empty block",
			"lastProposedAt", p.lastProposedAt,
		)
		txLists = append(txLists, types.Transactions{})
	}

	// If LocalAddressesOnly is set, filter the transactions by the local addresses.
	if p.LocalAddressesOnly {
		var (
			localTxsLists []types.Transactions
			signer        = types.LatestSignerForChainID(p.rpc.L2.ChainID)
		)
		for _, txs := range txLists {
			var filtered types.Transactions
			for _, tx := range txs {
				sender, err := types.Sender(signer, tx)
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

func (p *Proposer) isEligibleValidatorForNextSlot() bool {
	l1HeadSlot := p.l1HeadSlot.Load().(uint64)
	nextSlot := strconv.FormatUint(l1HeadSlot+1, 10)
	var nextDuty *rpc.ProposerDuty

	for _, duty := range p.proposerDuties {
		if duty.Slot == nextSlot {
			nextDuty = duty
			break
		}
	}

	if nextDuty == nil {
		log.Warn("Empty next proposer duty")
		return false
	}

	return nextDuty.PubKey == p.validatorPublicKeyHex
}

// ProposeOp performs a proposing operation, fetching transactions
// from L2 execution engine's tx pool, splitting them by proposing constraints,
// and then proposing them to TaikoL1 contract.
func (p *Proposer) ProposeOp(ctx context.Context) error {
	log.Info(
		"Start fetching L2 execution engine's transaction pool content",
		"lastProposedAt", p.lastProposedAt,
	)

	txLists, err := p.fetchTxListToPropose()
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
	for _, txs := range txLists[:utils.Min(p.MaxProposedTxListsPerEpoch, uint64(len(txLists)))] {
		nonce, err := p.rpc.L1.PendingNonceAt(ctx, p.proposerAddress)
		if err != nil {
			log.Error("Failed to get proposer nonce", "error", err)
			break
		}

		log.Info("Proposer current pending nonce", "nonce", nonce)

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

		if err := p.rpc.WaitL1NewPendingTransaction(ctx, p.proposerAddress, nonce); err != nil {
			log.Error("Failed to wait for new pending transaction", "error", err)
		}
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

	if p.isEligibleValidatorForNextSlot() {
		// If the validator is eligible for the next slot, set the constraints directly.
		tx, err := p.txmgr.CraftTx(ctx, *txCandidate)
		if err != nil {
			log.Warn("Failed to craft TaikoL1.proposeBlock transaction", "error", encoding.TryParsingCustomError(err))
			return err
		}

		l1HeadSlot := p.l1HeadSlot.Load().(uint64)
		nextSlot := l1HeadSlot + 1

		if err = p.rpc.L1MevBoost.SetConstraints(nextSlot, tx); err != nil {
			return fmt.Errorf("failed to set validator mev boost constraints: %w", err)
		}

		// Since we are not sending the transaction to the network, but instead to the MEV-Boost , we need to reset the nonce.
		p.txmgr.ResetNonce()
	} else {
		receipt, err := p.txmgr.Send(ctx, *txCandidate)
		if err != nil {
			log.Warn("Failed to send TaikoL1.proposeBlock transaction", "error", encoding.TryParsingCustomError(err))
			return err
		}

		if receipt.Status != types.ReceiptStatusSuccessful {
			return fmt.Errorf("failed to propose block: %s", receipt.TxHash.Hex())
		}
	}

	log.Info("📝 Propose transactions succeeded", "txs", txNum)

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
