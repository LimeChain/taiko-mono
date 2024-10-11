package mock

import (
	"context"

	"github.com/prysmaticlabs/prysm/v4/beacon-chain/rpc/eth/blob"

	"github.com/taikoxyz/taiko-mono/packages/taiko-client/types"
)

type BeaconClient struct{}

func NewBeaconClient() *BeaconClient {
	return &BeaconClient{}
}

func (c *BeaconClient) GetNextProposerDuties(
	_ context.Context,
	_ uint64,
	_ uint64) ([]*types.ProposerDuty, error) {
	return []*types.ProposerDuty{}, nil
}

func (c *BeaconClient) GetBlobs(_ context.Context, _ uint64) ([]*blob.Sidecar, error) {
	return []*blob.Sidecar{}, nil
}

func (c *BeaconClient) GetL1HeadSlot() uint64 {
	return 0
}

func (c *BeaconClient) GetTimestampBySlot(_ uint64) uint64 {
	return 0
}

func (c *BeaconClient) GetGenesisTimestamp() uint64 {
	return 0
}

func (c *BeaconClient) GetSecondsPerSlot() uint64 {
	return 0
}

func (c *BeaconClient) GetSlotsPerEpoch() uint64 {
	return 0
}
