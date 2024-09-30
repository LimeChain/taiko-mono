package rpc

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	"github.com/taikoxyz/taiko-mono/packages/taiko-client/pkg/mock"
)

func newTestClient(t *testing.T) *Client {
	client, err := NewClient(context.Background(), &ClientConfig{
		L1Endpoint:               os.Getenv("L1_NODE_WS_ENDPOINT"),
		L2Endpoint:               os.Getenv("L2_EXECUTION_ENGINE_WS_ENDPOINT"),
		TaikoL1Address:           common.HexToAddress(os.Getenv("TAIKO_L1_ADDRESS")),
		TaikoL2Address:           common.HexToAddress(os.Getenv("TAIKO_L2_ADDRESS")),
		SequencerRegistryAddress: common.HexToAddress(os.Getenv("SEQUENCER_REGISTRY")),
		TaikoTokenAddress:        common.HexToAddress(os.Getenv("TAIKO_TOKEN_ADDRESS")),
		ProverSetAddress:         common.HexToAddress(os.Getenv("PROVER_SET_ADDRESS")),
		L2EngineEndpoint:         os.Getenv("L2_EXECUTION_ENGINE_AUTH_ENDPOINT"),
		JwtSecret:                os.Getenv("JWT_SECRET"),
	})
	client.L1MevBoost = mock.NewMevBoostClient()
	client.L1Beacon = mock.NewBeaconClient()

	require.Nil(t, err)
	require.NotNil(t, client)

	return client
}

func newTestClientWithTimeout(t *testing.T) *Client {
	client, err := NewClient(context.Background(), &ClientConfig{
		L1Endpoint:               os.Getenv("L1_NODE_WS_ENDPOINT"),
		L2Endpoint:               os.Getenv("L2_EXECUTION_ENGINE_WS_ENDPOINT"),
		TaikoL1Address:           common.HexToAddress(os.Getenv("TAIKO_L1_ADDRESS")),
		TaikoL2Address:           common.HexToAddress(os.Getenv("TAIKO_L2_ADDRESS")),
		SequencerRegistryAddress: common.HexToAddress(os.Getenv("SEQUENCER_REGISTRY")),
		TaikoTokenAddress:        common.HexToAddress(os.Getenv("TAIKO_TOKEN_ADDRESS")),
		ProverSetAddress:         common.HexToAddress(os.Getenv("PROVER_SET_ADDRESS")),
		L2EngineEndpoint:         os.Getenv("L2_EXECUTION_ENGINE_AUTH_ENDPOINT"),
		JwtSecret:                os.Getenv("JWT_SECRET"),
		Timeout:                  5 * time.Second,
	})
	client.L1MevBoost = mock.NewMevBoostClient()
	client.L1Beacon = mock.NewBeaconClient()

	require.Nil(t, err)
	require.NotNil(t, client)

	return client
}
