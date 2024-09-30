package builder

import (
	"context"

	"github.com/taikoxyz/taiko-mono/packages/taiko-client/bindings/encoding"
)

// ProposeBlockTransactionBuilder is an interface for building a TaikoL1.proposeBlock transaction.
type ProposeBlockTransactionBuilder interface {
	Build(
		ctx context.Context,
		tierFees []encoding.TierFee,
		includeParentMetaHash bool,
		txListBytes []byte,
	) (*TxCandidate, error)
}
