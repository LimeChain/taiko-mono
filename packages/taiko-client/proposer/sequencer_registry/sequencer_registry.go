package main

import (
	"bytes"
	"context"
	"log"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/joho/godotenv"

	"github.com/taikoxyz/taiko-mono/packages/taiko-client/bindings"
	"github.com/taikoxyz/taiko-mono/packages/taiko-client/pkg/rpc"
)

const (
	defaultTimeout = 1 * time.Minute
)

type ISequencerRegistryValidatorProof struct {
	CurrentEpoch    uint64
	ActivationEpoch uint64
	ExitEpoch       uint64
	ValidatorIndex  *big.Int
	Slashed         bool
	ProofSlot       *big.Int
	SszProof        []byte
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file", err)
	}
}

func encodePacked(input ...[]byte) []byte {
	return bytes.Join(input, nil)
}

func register(ctx context.Context, client *rpc.EthClient, auth *bind.TransactOpts, chainId *big.Int) error {
	registryAddress := common.HexToAddress(os.Getenv("SEQUENCER_REGISTRY"))

	registry, err := bindings.NewSequencerRegistry(registryAddress, client)
	if err != nil {
		log.Fatalf("Failed to instantiate a SequencerRegistry contract: %v", err)
	}

	validatorAddress := common.HexToAddress(os.Getenv("VALIDATOR_ADDRESS"))
	fnSelector := [4]byte{0xe6, 0xe9, 0x11, 0x57}
	packed := encodePacked(
		[]byte{0x1},
		registryAddress.Bytes(),
		math.U256Bytes(chainId),
		math.U256Bytes(big.NewInt(0)), fnSelector[:], validatorAddress.Bytes(),
		[]byte{},
	)

	authHash := crypto.Keccak256(packed)

	if len(authHash) != 32 {
		log.Fatalf("authHash length is not 32 bytes: got %d", len(authHash))
	}

	var authHashArray [32]byte
	copy(authHashArray[:], authHash)

	tx, err := registry.Register(
		auth,
		validatorAddress,
		[]byte{},
		authHashArray,
		[]byte{},
		bindings.ISequencerRegistryValidatorProof{
			CurrentEpoch:    0,
			ActivationEpoch: 0,
			ExitEpoch:       0,
			ValidatorIndex:  big.NewInt(0),
			Slashed:         false,
			ProofSlot:       big.NewInt(0),
			SszProof:        []byte{},
		},
	)

	if err != nil {
		log.Fatalf("failed to register sequencer in the registry", "error", err)
		return err
	}

	receipt, err := bind.WaitMined(ctx, client, tx)
	if err != nil {
		log.Fatalf("transaction mining error", "error", err)
		return err
	}
	if receipt.Status != 1 {
		log.Fatalf("transaction failed", "error", err)
		return err
	}

	log.Printf("Sequencer registered successfully")
	return nil
}

func activate(ctx context.Context, client *rpc.EthClient, auth *bind.TransactOpts) error {
	registryAddress := common.HexToAddress(os.Getenv("TAIKOL1"))

	registry, err := bindings.NewTaikoL1Client(registryAddress, client)
	if err != nil {
		log.Fatalf("Failed to instantiate a TaikoL1 contract: %v", err)
	}

	auth.Value = big.NewInt(1e18) // Staking 1 ETH

	pubkey := make([]byte, 48)
	pubkey[31] = 0x01

	tx, err := registry.StakeSequencer(
		auth,
		pubkey, // publicKey
		bindings.ISequencerRegistryValidatorProof{
			CurrentEpoch:    0,
			ActivationEpoch: 0,
			ExitEpoch:       0,
			ValidatorIndex:  big.NewInt(0),
			Slashed:         false,
			ProofSlot:       big.NewInt(0),
			SszProof:        []byte{},
		},
	)

	if err != nil {
		log.Fatalf("failed to stake sequencer", "error", err)
		return err
	}

	receipt, err := bind.WaitMined(ctx, client, tx)
	if err != nil {
		log.Fatalf("transaction mining error", "error", err)
		return err
	}
	if receipt.Status != 1 {
		log.Fatalf("transaction failed", "error", err)
		return err
	}

	log.Printf("Sequencer activated successfully")
	return nil
}

func main() {
	var client *rpc.EthClient

	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s [register|activate]", os.Args[0])
		os.Exit(1)
	}

	loadEnv()

	ctx := context.Background()

	privateKey := os.Getenv("PRIVATE_KEY")
	rpcURL := os.Getenv("RPC_URL")
	chainId := os.Getenv("CHAIN_ID")

	l1ProposerPrivKey, err := crypto.ToECDSA(common.FromHex(privateKey))
	if err != nil {
		log.Fatalf("invalid L1 proposer private key: %w", err)
		os.Exit(1)
	}

	chainIdInt, ok := new(big.Int).SetString(chainId, 10)
	if !ok {
		log.Fatalf("Invalid CHAIN_ID")
		os.Exit(1)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(l1ProposerPrivKey, chainIdInt) // Adjust the chain ID as needed
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
		os.Exit(1)
	}

	if client, err = rpc.NewEthClient(ctx, rpcURL, defaultTimeout); err != nil {
		log.Fatalf("Failed to connect to L1 endpoint, retrying", "endpoint", rpcURL, "err", err)
		os.Exit(1)
	}

	switch os.Args[1] {
	case "register":
		err := register(ctx, client, auth, chainIdInt)
		if err != nil {
			log.Fatalf("Failed to register sequencer: %v", err)
			os.Exit(1)
		}
	case "activate":
		err := activate(ctx, client, auth)
		if err != nil {
			log.Fatalf("Failed to register sequencer: %v", err)
			os.Exit(1)
		}
	default:
		log.Fatalf("Unknown action: %s", os.Args[1])
	}
}
