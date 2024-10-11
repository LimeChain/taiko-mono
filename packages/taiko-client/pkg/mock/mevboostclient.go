package mock

import (
	"github.com/ethereum/go-ethereum/core/types"
)

type MevBoostClient struct{}

func NewMevBoostClient() *MevBoostClient {
	return &MevBoostClient{}
}

func (c *MevBoostClient) SetConstraints(_ uint64, _ *types.Transaction) error {
	return nil
}

func (c *MevBoostClient) GetValidatorPubKeyHex() (string, error) {
	return "", nil
}
