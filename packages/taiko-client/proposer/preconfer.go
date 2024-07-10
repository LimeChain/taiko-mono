package proposer

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"time"

	"golang.org/x/sync/errgroup"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/beacon/engine"
	"github.com/ethereum/go-ethereum/common"
	consensus "github.com/ethereum/go-ethereum/consensus/taiko"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/taikoxyz/taiko-mono/packages/taiko-client/bindings/encoding"
	anchorTxConstructor "github.com/taikoxyz/taiko-mono/packages/taiko-client/driver/anchor_tx_constructor"
)

// StartL2Preconfirmations rebuilds preconfirmation blocks with the latest pool txs.
func (p *Proposer) StartL2Preconfirmations() {
	for {
		time.Sleep(5 * time.Second)
		log.Info("Fetching mempool txs")

		filterPoolContent := time.Now().Before(p.lastProposedAt.Add(p.MinProposingInternal))
		poolTxs, err := p.fetchPoolContent(filterPoolContent)
		if err != nil {
			log.Error("Failed fetching pool content", "error", err)
			continue
		}

		if len(poolTxs) == 0 {
			log.Debug("Skipping preconfirmation block rebuild, no pool txs")
			continue
		}

		log.Debug("Current pool tx list", "count", len(poolTxs))
		log.Debug("Current first pool tx list", "txs", len(poolTxs[0]))

		preconfirmedTxs, err := p.fetchPreconfirmedTxs()
		if err != nil {
			log.Error("Failed fetching preconfirmed block transactions", "error", err)
			continue
		}
		log.Debug("Preconfirmed tx lists", "length", len(preconfirmedTxs))

		// TODO(limechain): handle multiple tx lists

		var txs types.Transactions
		if len(preconfirmedTxs) > 0 {
			txs = append(txs, preconfirmedTxs[0]...)
		}

		txs = append(txs, poolTxs[0]...)
		log.Debug("Txs for processing", "length", len(txs))

		g, gCtx := errgroup.WithContext(p.ctx)

		g.Go(func() error {
			if err := p.BuildPreconfBlock(gCtx, txs); err != nil {
				log.Error("Failed to build preconfirmation block", "error", err)
				return err
			}
			return nil
		})
	}
}

// BuildPreconfBlock builds virtual block that provides pre-confirmation receipts for the contained TXs.
func (p *Proposer) BuildPreconfBlock(ctx context.Context, txList types.Transactions) error {
	var (
		l2Head   *types.Header
		pbCursor *types.PreconfBlockCursor
		err      error
	)
	l2Head, err = p.rpc.L2.HeaderByNumber(ctx, nil)
	if err != nil {
		return err
	}
	log.Info("preconfer: l2 head", "number", l2Head.Number, "hash", l2Head.Hash())

	// Mark any pre-confirmed transactions that are already executed to be
	// skipped in the next pre-confirmation block rebuild.
	if pbCursor, err = p.rpc.L2.GetPreconfBlockCursor(ctx); err != nil {
		return fmt.Errorf("failed to get preconfirmation block cursor: %s", err)
	}
	if pbCursor != nil && (l2Head.Number.Uint64() > pbCursor.Number) {
		skip := true
		if err = p.rpc.L2.UpdatePreconfBlockCursor(ctx, &pbCursor.Hash, new(big.Int).SetUint64(pbCursor.Number), new(big.Int).SetUint64(pbCursor.ProposedTxCount), &skip); err != nil {
			return fmt.Errorf("failed to update preconfirmation block cursor: %s", err)
		}
		return nil
	}

	// parameters of the TaikoL2.anchor transaction
	l1Origin, err := p.rpc.L2.HeadL1Origin(ctx)
	if err != nil {
		return err
	}
	l1Height := l1Origin.L1BlockHeight
	l1Hash := l1Origin.L1BlockHash

	baseFeeInfo, err := p.rpc.TaikoL2.GetBasefee(
		&bind.CallOpts{BlockNumber: l2Head.Number, Context: ctx},
		l1Height.Uint64(),
		uint32(l2Head.GasUsed),
	)
	if err != nil {
		return fmt.Errorf("preconfer: failed to get L2 baseFee: %w", encoding.TryParsingCustomError(err))
	}

	anchorConstructor, err := anchorTxConstructor.New(p.rpc)
	if err != nil {
		return err
	}
	anchorTx, err := anchorConstructor.AssembleAnchorTx(
		ctx,
		l1Height,
		l1Hash,
		new(big.Int).Add(l2Head.Number, common.Big1),
		baseFeeInfo.Basefee,
		l2Head.GasUsed,
	)
	if err != nil {
		return fmt.Errorf("preconfer: failed to create TaikoL2.anchor transaction: %w", err)
	}
	log.Info("Anchor tx", "hash", anchorTx.Hash().String())

	// Insert the anchor transaction at the head of the transactions list
	txList = append([]*types.Transaction{anchorTx}, txList...)
	txListBytes, err := rlp.EncodeToBytes(txList)
	if err != nil {
		return fmt.Errorf("preconfer: failed to encode transactions: %w", err)
	}

	fc := &engine.ForkchoiceStateV1{HeadBlockHash: l2Head.Hash()}

	timestamp := uint64(time.Now().Unix())
	coinbase := common.Address{}

	attributes := &engine.PayloadAttributes{
		Timestamp:             timestamp,
		Random:                common.Hash{},
		SuggestedFeeRecipient: coinbase,
		Withdrawals:           make(types.Withdrawals, 0),
		BlockMetadata: &engine.BlockMetadata{
			HighestBlockID: new(big.Int).Add(l2Head.Number, common.Big1),
			Beneficiary:    coinbase,
			GasLimit:       uint64(21000) + consensus.AnchorGasLimit,
			Timestamp:      timestamp,
			TxList:         txListBytes,
			MixHash:        common.Hash{},
			ExtraData:      []byte{},
		},
		BaseFeePerGas: baseFeeInfo.Basefee,
		L1Origin: &rawdb.L1Origin{
			BlockID:       new(big.Int).Add(l2Head.Number, common.Big1),
			L2BlockHash:   common.Hash{}, // Will be set by taiko-geth.
			L1BlockHeight: l1Height,
			L1BlockHash:   l1Hash,
		},
		PreconfBlock: true,
	}

	// Start building payload
	log.Debug("preconfer: start building payload")
	fcRes, err := p.rpc.L2Engine.ForkchoiceUpdate(ctx, fc, attributes)
	if err != nil {
		log.Debug("preconfer: failed to update fork choice")
		return fmt.Errorf("preconfer: failed to update fork choice: %w", err)
	}
	if fcRes.PayloadStatus.Status != engine.VALID {
		log.Debug("Preconfer: unexpected ForkchoiceUpdate response status")
		return fmt.Errorf("preconfer: unexpected ForkchoiceUpdate response status: %s", fcRes.PayloadStatus.Status)
	}
	if fcRes.PayloadID == nil {
		log.Debug("Preconfer: empty payload ID")
		return errors.New("preconfer: empty payload ID")
	}

	// Get the built payload
	log.Debug("Preconfer: get built payload")
	payload, err := p.rpc.L2Engine.GetPayload(ctx, fcRes.PayloadID)
	if err != nil {
		log.Debug("Preconfer: failed to get payload")
		return fmt.Errorf("preconfer: failed to get payload: %w", err)
	}

	// Execute the payload
	log.Debug("preconfer: execute the payload", "block hash", payload.BlockHash.String(), "txs", len(payload.Transactions))
	execStatus, err := p.rpc.L2Engine.NewPayload(ctx, payload)
	if err != nil {
		log.Debug("preconfer: failed to create a new payload")
		return fmt.Errorf("preconfer: failed to create a new payload: %w", err)
	}
	if execStatus.Status != engine.VALID {
		return fmt.Errorf("preconfer: unexpected NewPayload response status: %s", execStatus.Status)
	}

	// fc := &engine.ForkchoiceStateV1{
	// 	HeadBlockHash:      payload.BlockHash,
	// 	SafeBlockHash:      payload.BlockHash,
	// 	FinalizedBlockHash: payload.BlockHash,
	// }

	// // Update the fork choice
	// fcRes, err := pr.rpc.L2Engine.ForkchoiceUpdate(ctx, fc, nil)
	// if err != nil {
	// 	return nil, err
	// }
	// if fcRes.PayloadStatus.Status != engine.VALID {
	// 	return nil, fmt.Errorf("Preconfer: unexpected ForkchoiceUpdate response status: %s", fcRes.PayloadStatus.Status)
	// }

	// TODO: move this as part of the payload execution
	err = p.rpc.L2.UpdatePreconfBlockCursor(ctx, &payload.BlockHash, new(big.Int).SetUint64(payload.Number), new(big.Int), new(bool))
	if err != nil {
		return fmt.Errorf("preconfer: failed to update preconfirmation block: %w", err)
	}

	log.Debug("preconfer: payload", "hash", payload.BlockHash.String(), "txs", len(payload.Transactions))

	return nil
}
