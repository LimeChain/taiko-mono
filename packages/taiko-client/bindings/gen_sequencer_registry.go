// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ISequencerRegistrySequencer is an auto generated low-level Go binding around an user-defined struct.
type ISequencerRegistrySequencer struct {
	Pubkey            []byte
	Metadata          []byte
	Signer            common.Address
	ActivationBlock   *big.Int
	DeactivationBlock *big.Int
}

// ISequencerRegistryValidatorProof is an auto generated low-level Go binding around an user-defined struct.
type ISequencerRegistryValidatorProof struct {
	CurrentEpoch    uint64
	ActivationEpoch uint64
	ExitEpoch       uint64
	ValidatorIndex  *big.Int
	Slashed         bool
	ProofSlot       *big.Int
	SszProof        []byte
}

// SequencerRegistryMetaData contains all meta data concerning the SequencerRegistry contract.
var SequencerRegistryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"ACTIVATION_TIMEOUT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEACTIVATION_PERIOD\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PROTOCOL_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"activate\",\"inputs\":[{\"name\":\"pubkey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"validatorProof\",\"type\":\"tuple\",\"internalType\":\"structISequencerRegistry.ValidatorProof\",\"components\":[{\"name\":\"currentEpoch\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"activationEpoch\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"exitEpoch\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"validatorIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"slashed\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"proofSlot\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sszProof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"activationTimeout\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"addressManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allValidators\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"changeRegistration\",\"inputs\":[{\"name\":\"signer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"metadata\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"authHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deactivate\",\"inputs\":[{\"name\":\"authHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deactivationPeriod\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"eligibleCountAt\",\"inputs\":[{\"name\":\"blockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"fallbackSigner\",\"inputs\":[{\"name\":\"_blockNum\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"forceDeactivate\",\"inputs\":[{\"name\":\"pubkey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"validatorProof\",\"type\":\"tuple\",\"internalType\":\"structISequencerRegistry.ValidatorProof\",\"components\":[{\"name\":\"currentEpoch\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"activationEpoch\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"exitEpoch\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"validatorIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"slashed\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"proofSlot\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sszProof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"impl\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"inNonReentrant\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"init\",\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isEligible\",\"inputs\":[{\"name\":\"pubkey\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isEligibleSigner\",\"inputs\":[{\"name\":\"signer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRegistered\",\"inputs\":[{\"name\":\"signer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lastUnpausedAt\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nonces\",\"inputs\":[{\"name\":\"pubkey\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingOwner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"protocolVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"register\",\"inputs\":[{\"name\":\"signer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"metadata\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"authHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"validatorProof\",\"type\":\"tuple\",\"internalType\":\"structISequencerRegistry.ValidatorProof\",\"components\":[{\"name\":\"currentEpoch\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"activationEpoch\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"exitEpoch\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"validatorIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"slashed\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"proofSlot\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sszProof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"resolve\",\"inputs\":[{\"name\":\"_chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"_name\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_allowZeroAddress\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"addresspayable\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"resolve\",\"inputs\":[{\"name\":\"_name\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_allowZeroAddress\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"addresspayable\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"sequencerByIndex\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"signer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"metadata\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"pubkey\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"sequencersToPubkeyHash\",\"inputs\":[{\"name\":\"sequencer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"pubkeyHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"statusOf\",\"inputs\":[{\"name\":\"pubkey\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"data\",\"type\":\"tuple\",\"internalType\":\"structISequencerRegistry.Sequencer\",\"components\":[{\"name\":\"pubkey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"metadata\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"activationBlock\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deactivationBlock\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeTo\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"validators\",\"inputs\":[{\"name\":\"pubkeyHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"pubkey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"metadata\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"activationBlock\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deactivationBlock\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"AdminChanged\",\"inputs\":[{\"name\":\"previousAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BeaconUpgraded\",\"inputs\":[{\"name\":\"beacon\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferStarted\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SequencerActivated\",\"inputs\":[{\"name\":\"signer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SequencerChanged\",\"inputs\":[{\"name\":\"oldSigner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newSigner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"pubkey\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SequencerDeactivated\",\"inputs\":[{\"name\":\"signer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SequencerRegistered\",\"inputs\":[{\"name\":\"signer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"pubkey\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"FUNC_NOT_IMPLEMENTED\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"INVALID_PAUSE_STATUS\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"REENTRANT_CALL\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RESOLVER_DENIED\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RESOLVER_INVALID_MANAGER\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RESOLVER_UNEXPECTED_CHAINID\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RESOLVER_ZERO_ADDR\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"name\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"SR_BLOCK_TOO_LOW\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SR_INDEX_OUT_OF_BOUNDS\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SR_INVALID_ADDRESS\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SR_INVALID_AUTH_HASH\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SR_INVALID_AUTH_SIGNATURE\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SR_INVALID_PROOF\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SR_NO_ELIGIBLE_SEQUENCERS\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SR_SIGNER_REGISTERED\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SR_VALIDATOR_DEACTIVATED\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SR_VALIDATOR_NOT_REGISTERED\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SR_VALIDATOR_REGISTERED\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZERO_ADDRESS\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZERO_VALUE\",\"inputs\":[]}]",
}

// SequencerRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use SequencerRegistryMetaData.ABI instead.
var SequencerRegistryABI = SequencerRegistryMetaData.ABI

// SequencerRegistry is an auto generated Go binding around an Ethereum contract.
type SequencerRegistry struct {
	SequencerRegistryCaller     // Read-only binding to the contract
	SequencerRegistryTransactor // Write-only binding to the contract
	SequencerRegistryFilterer   // Log filterer for contract events
}

// SequencerRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type SequencerRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SequencerRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SequencerRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SequencerRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SequencerRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SequencerRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SequencerRegistrySession struct {
	Contract     *SequencerRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// SequencerRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SequencerRegistryCallerSession struct {
	Contract *SequencerRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// SequencerRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SequencerRegistryTransactorSession struct {
	Contract     *SequencerRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// SequencerRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type SequencerRegistryRaw struct {
	Contract *SequencerRegistry // Generic contract binding to access the raw methods on
}

// SequencerRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SequencerRegistryCallerRaw struct {
	Contract *SequencerRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// SequencerRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SequencerRegistryTransactorRaw struct {
	Contract *SequencerRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSequencerRegistry creates a new instance of SequencerRegistry, bound to a specific deployed contract.
func NewSequencerRegistry(address common.Address, backend bind.ContractBackend) (*SequencerRegistry, error) {
	contract, err := bindSequencerRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SequencerRegistry{SequencerRegistryCaller: SequencerRegistryCaller{contract: contract}, SequencerRegistryTransactor: SequencerRegistryTransactor{contract: contract}, SequencerRegistryFilterer: SequencerRegistryFilterer{contract: contract}}, nil
}

// NewSequencerRegistryCaller creates a new read-only instance of SequencerRegistry, bound to a specific deployed contract.
func NewSequencerRegistryCaller(address common.Address, caller bind.ContractCaller) (*SequencerRegistryCaller, error) {
	contract, err := bindSequencerRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SequencerRegistryCaller{contract: contract}, nil
}

// NewSequencerRegistryTransactor creates a new write-only instance of SequencerRegistry, bound to a specific deployed contract.
func NewSequencerRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*SequencerRegistryTransactor, error) {
	contract, err := bindSequencerRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SequencerRegistryTransactor{contract: contract}, nil
}

// NewSequencerRegistryFilterer creates a new log filterer instance of SequencerRegistry, bound to a specific deployed contract.
func NewSequencerRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*SequencerRegistryFilterer, error) {
	contract, err := bindSequencerRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SequencerRegistryFilterer{contract: contract}, nil
}

// bindSequencerRegistry binds a generic wrapper to an already deployed contract.
func bindSequencerRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SequencerRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SequencerRegistry *SequencerRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SequencerRegistry.Contract.SequencerRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SequencerRegistry *SequencerRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SequencerRegistry.Contract.SequencerRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SequencerRegistry *SequencerRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SequencerRegistry.Contract.SequencerRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SequencerRegistry *SequencerRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SequencerRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SequencerRegistry *SequencerRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SequencerRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SequencerRegistry *SequencerRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SequencerRegistry.Contract.contract.Transact(opts, method, params...)
}

// ACTIVATIONTIMEOUT is a free data retrieval call binding the contract method 0x8f3b2275.
//
// Solidity: function ACTIVATION_TIMEOUT() view returns(uint8)
func (_SequencerRegistry *SequencerRegistryCaller) ACTIVATIONTIMEOUT(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _SequencerRegistry.contract.Call(opts, &out, "ACTIVATION_TIMEOUT")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ACTIVATIONTIMEOUT is a free data retrieval call binding the contract method 0x8f3b2275.
//
// Solidity: function ACTIVATION_TIMEOUT() view returns(uint8)
func (_SequencerRegistry *SequencerRegistrySession) ACTIVATIONTIMEOUT() (uint8, error) {
	return _SequencerRegistry.Contract.ACTIVATIONTIMEOUT(&_SequencerRegistry.CallOpts)
}

// ACTIVATIONTIMEOUT is a free data retrieval call binding the contract method 0x8f3b2275.
//
// Solidity: function ACTIVATION_TIMEOUT() view returns(uint8)
func (_SequencerRegistry *SequencerRegistryCallerSession) ACTIVATIONTIMEOUT() (uint8, error) {
	return _SequencerRegistry.Contract.ACTIVATIONTIMEOUT(&_SequencerRegistry.CallOpts)
}

// DEACTIVATIONPERIOD is a free data retrieval call binding the contract method 0x7c20672e.
//
// Solidity: function DEACTIVATION_PERIOD() view returns(uint8)
func (_SequencerRegistry *SequencerRegistryCaller) DEACTIVATIONPERIOD(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _SequencerRegistry.contract.Call(opts, &out, "DEACTIVATION_PERIOD")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// DEACTIVATIONPERIOD is a free data retrieval call binding the contract method 0x7c20672e.
//
// Solidity: function DEACTIVATION_PERIOD() view returns(uint8)
func (_SequencerRegistry *SequencerRegistrySession) DEACTIVATIONPERIOD() (uint8, error) {
	return _SequencerRegistry.Contract.DEACTIVATIONPERIOD(&_SequencerRegistry.CallOpts)
}

// DEACTIVATIONPERIOD is a free data retrieval call binding the contract method 0x7c20672e.
//
// Solidity: function DEACTIVATION_PERIOD() view returns(uint8)
func (_SequencerRegistry *SequencerRegistryCallerSession) DEACTIVATIONPERIOD() (uint8, error) {
	return _SequencerRegistry.Contract.DEACTIVATIONPERIOD(&_SequencerRegistry.CallOpts)
}

// PROTOCOLVERSION is a free data retrieval call binding the contract method 0xaa3aa460.
//
// Solidity: function PROTOCOL_VERSION() view returns(uint8)
func (_SequencerRegistry *SequencerRegistryCaller) PROTOCOLVERSION(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _SequencerRegistry.contract.Call(opts, &out, "PROTOCOL_VERSION")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// PROTOCOLVERSION is a free data retrieval call binding the contract method 0xaa3aa460.
//
// Solidity: function PROTOCOL_VERSION() view returns(uint8)
func (_SequencerRegistry *SequencerRegistrySession) PROTOCOLVERSION() (uint8, error) {
	return _SequencerRegistry.Contract.PROTOCOLVERSION(&_SequencerRegistry.CallOpts)
}

// PROTOCOLVERSION is a free data retrieval call binding the contract method 0xaa3aa460.
//
// Solidity: function PROTOCOL_VERSION() view returns(uint8)
func (_SequencerRegistry *SequencerRegistryCallerSession) PROTOCOLVERSION() (uint8, error) {
	return _SequencerRegistry.Contract.PROTOCOLVERSION(&_SequencerRegistry.CallOpts)
}

// ActivationTimeout is a free data retrieval call binding the contract method 0x30bd382e.
//
// Solidity: function activationTimeout() pure returns(uint8)
func (_SequencerRegistry *SequencerRegistryCaller) ActivationTimeout(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _SequencerRegistry.contract.Call(opts, &out, "activationTimeout")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ActivationTimeout is a free data retrieval call binding the contract method 0x30bd382e.
//
// Solidity: function activationTimeout() pure returns(uint8)
func (_SequencerRegistry *SequencerRegistrySession) ActivationTimeout() (uint8, error) {
	return _SequencerRegistry.Contract.ActivationTimeout(&_SequencerRegistry.CallOpts)
}

// ActivationTimeout is a free data retrieval call binding the contract method 0x30bd382e.
//
// Solidity: function activationTimeout() pure returns(uint8)
func (_SequencerRegistry *SequencerRegistryCallerSession) ActivationTimeout() (uint8, error) {
	return _SequencerRegistry.Contract.ActivationTimeout(&_SequencerRegistry.CallOpts)
}

// AddressManager is a free data retrieval call binding the contract method 0x3ab76e9f.
//
// Solidity: function addressManager() view returns(address)
func (_SequencerRegistry *SequencerRegistryCaller) AddressManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SequencerRegistry.contract.Call(opts, &out, "addressManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressManager is a free data retrieval call binding the contract method 0x3ab76e9f.
//
// Solidity: function addressManager() view returns(address)
func (_SequencerRegistry *SequencerRegistrySession) AddressManager() (common.Address, error) {
	return _SequencerRegistry.Contract.AddressManager(&_SequencerRegistry.CallOpts)
}

// AddressManager is a free data retrieval call binding the contract method 0x3ab76e9f.
//
// Solidity: function addressManager() view returns(address)
func (_SequencerRegistry *SequencerRegistryCallerSession) AddressManager() (common.Address, error) {
	return _SequencerRegistry.Contract.AddressManager(&_SequencerRegistry.CallOpts)
}

// AllValidators is a free data retrieval call binding the contract method 0xbcecf81b.
//
// Solidity: function allValidators(uint256 ) view returns(bytes32)
func (_SequencerRegistry *SequencerRegistryCaller) AllValidators(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _SequencerRegistry.contract.Call(opts, &out, "allValidators", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AllValidators is a free data retrieval call binding the contract method 0xbcecf81b.
//
// Solidity: function allValidators(uint256 ) view returns(bytes32)
func (_SequencerRegistry *SequencerRegistrySession) AllValidators(arg0 *big.Int) ([32]byte, error) {
	return _SequencerRegistry.Contract.AllValidators(&_SequencerRegistry.CallOpts, arg0)
}

// AllValidators is a free data retrieval call binding the contract method 0xbcecf81b.
//
// Solidity: function allValidators(uint256 ) view returns(bytes32)
func (_SequencerRegistry *SequencerRegistryCallerSession) AllValidators(arg0 *big.Int) ([32]byte, error) {
	return _SequencerRegistry.Contract.AllValidators(&_SequencerRegistry.CallOpts, arg0)
}

// DeactivationPeriod is a free data retrieval call binding the contract method 0x887fba65.
//
// Solidity: function deactivationPeriod() pure returns(uint8)
func (_SequencerRegistry *SequencerRegistryCaller) DeactivationPeriod(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _SequencerRegistry.contract.Call(opts, &out, "deactivationPeriod")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// DeactivationPeriod is a free data retrieval call binding the contract method 0x887fba65.
//
// Solidity: function deactivationPeriod() pure returns(uint8)
func (_SequencerRegistry *SequencerRegistrySession) DeactivationPeriod() (uint8, error) {
	return _SequencerRegistry.Contract.DeactivationPeriod(&_SequencerRegistry.CallOpts)
}

// DeactivationPeriod is a free data retrieval call binding the contract method 0x887fba65.
//
// Solidity: function deactivationPeriod() pure returns(uint8)
func (_SequencerRegistry *SequencerRegistryCallerSession) DeactivationPeriod() (uint8, error) {
	return _SequencerRegistry.Contract.DeactivationPeriod(&_SequencerRegistry.CallOpts)
}

// EligibleCountAt is a free data retrieval call binding the contract method 0xec53c45c.
//
// Solidity: function eligibleCountAt(uint256 blockNumber) view returns(uint256)
func (_SequencerRegistry *SequencerRegistryCaller) EligibleCountAt(opts *bind.CallOpts, blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SequencerRegistry.contract.Call(opts, &out, "eligibleCountAt", blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EligibleCountAt is a free data retrieval call binding the contract method 0xec53c45c.
//
// Solidity: function eligibleCountAt(uint256 blockNumber) view returns(uint256)
func (_SequencerRegistry *SequencerRegistrySession) EligibleCountAt(blockNumber *big.Int) (*big.Int, error) {
	return _SequencerRegistry.Contract.EligibleCountAt(&_SequencerRegistry.CallOpts, blockNumber)
}

// EligibleCountAt is a free data retrieval call binding the contract method 0xec53c45c.
//
// Solidity: function eligibleCountAt(uint256 blockNumber) view returns(uint256)
func (_SequencerRegistry *SequencerRegistryCallerSession) EligibleCountAt(blockNumber *big.Int) (*big.Int, error) {
	return _SequencerRegistry.Contract.EligibleCountAt(&_SequencerRegistry.CallOpts, blockNumber)
}

// FallbackSigner is a free data retrieval call binding the contract method 0x5cf858d0.
//
// Solidity: function fallbackSigner(uint256 _blockNum) view returns(address)
func (_SequencerRegistry *SequencerRegistryCaller) FallbackSigner(opts *bind.CallOpts, _blockNum *big.Int) (common.Address, error) {
	var out []interface{}
	err := _SequencerRegistry.contract.Call(opts, &out, "fallbackSigner", _blockNum)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FallbackSigner is a free data retrieval call binding the contract method 0x5cf858d0.
//
// Solidity: function fallbackSigner(uint256 _blockNum) view returns(address)
func (_SequencerRegistry *SequencerRegistrySession) FallbackSigner(_blockNum *big.Int) (common.Address, error) {
	return _SequencerRegistry.Contract.FallbackSigner(&_SequencerRegistry.CallOpts, _blockNum)
}

// FallbackSigner is a free data retrieval call binding the contract method 0x5cf858d0.
//
// Solidity: function fallbackSigner(uint256 _blockNum) view returns(address)
func (_SequencerRegistry *SequencerRegistryCallerSession) FallbackSigner(_blockNum *big.Int) (common.Address, error) {
	return _SequencerRegistry.Contract.FallbackSigner(&_SequencerRegistry.CallOpts, _blockNum)
}

// Impl is a free data retrieval call binding the contract method 0x8abf6077.
//
// Solidity: function impl() view returns(address)
func (_SequencerRegistry *SequencerRegistryCaller) Impl(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SequencerRegistry.contract.Call(opts, &out, "impl")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Impl is a free data retrieval call binding the contract method 0x8abf6077.
//
// Solidity: function impl() view returns(address)
func (_SequencerRegistry *SequencerRegistrySession) Impl() (common.Address, error) {
	return _SequencerRegistry.Contract.Impl(&_SequencerRegistry.CallOpts)
}

// Impl is a free data retrieval call binding the contract method 0x8abf6077.
//
// Solidity: function impl() view returns(address)
func (_SequencerRegistry *SequencerRegistryCallerSession) Impl() (common.Address, error) {
	return _SequencerRegistry.Contract.Impl(&_SequencerRegistry.CallOpts)
}

// InNonReentrant is a free data retrieval call binding the contract method 0x3075db56.
//
// Solidity: function inNonReentrant() view returns(bool)
func (_SequencerRegistry *SequencerRegistryCaller) InNonReentrant(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _SequencerRegistry.contract.Call(opts, &out, "inNonReentrant")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// InNonReentrant is a free data retrieval call binding the contract method 0x3075db56.
//
// Solidity: function inNonReentrant() view returns(bool)
func (_SequencerRegistry *SequencerRegistrySession) InNonReentrant() (bool, error) {
	return _SequencerRegistry.Contract.InNonReentrant(&_SequencerRegistry.CallOpts)
}

// InNonReentrant is a free data retrieval call binding the contract method 0x3075db56.
//
// Solidity: function inNonReentrant() view returns(bool)
func (_SequencerRegistry *SequencerRegistryCallerSession) InNonReentrant() (bool, error) {
	return _SequencerRegistry.Contract.InNonReentrant(&_SequencerRegistry.CallOpts)
}

// IsEligible is a free data retrieval call binding the contract method 0xdc80b827.
//
// Solidity: function isEligible(bytes pubkey) view returns(bool)
func (_SequencerRegistry *SequencerRegistryCaller) IsEligible(opts *bind.CallOpts, pubkey []byte) (bool, error) {
	var out []interface{}
	err := _SequencerRegistry.contract.Call(opts, &out, "isEligible", pubkey)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsEligible is a free data retrieval call binding the contract method 0xdc80b827.
//
// Solidity: function isEligible(bytes pubkey) view returns(bool)
func (_SequencerRegistry *SequencerRegistrySession) IsEligible(pubkey []byte) (bool, error) {
	return _SequencerRegistry.Contract.IsEligible(&_SequencerRegistry.CallOpts, pubkey)
}

// IsEligible is a free data retrieval call binding the contract method 0xdc80b827.
//
// Solidity: function isEligible(bytes pubkey) view returns(bool)
func (_SequencerRegistry *SequencerRegistryCallerSession) IsEligible(pubkey []byte) (bool, error) {
	return _SequencerRegistry.Contract.IsEligible(&_SequencerRegistry.CallOpts, pubkey)
}

// IsEligibleSigner is a free data retrieval call binding the contract method 0x1872ae94.
//
// Solidity: function isEligibleSigner(address signer) view returns(bool)
func (_SequencerRegistry *SequencerRegistryCaller) IsEligibleSigner(opts *bind.CallOpts, signer common.Address) (bool, error) {
	var out []interface{}
	err := _SequencerRegistry.contract.Call(opts, &out, "isEligibleSigner", signer)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsEligibleSigner is a free data retrieval call binding the contract method 0x1872ae94.
//
// Solidity: function isEligibleSigner(address signer) view returns(bool)
func (_SequencerRegistry *SequencerRegistrySession) IsEligibleSigner(signer common.Address) (bool, error) {
	return _SequencerRegistry.Contract.IsEligibleSigner(&_SequencerRegistry.CallOpts, signer)
}

// IsEligibleSigner is a free data retrieval call binding the contract method 0x1872ae94.
//
// Solidity: function isEligibleSigner(address signer) view returns(bool)
func (_SequencerRegistry *SequencerRegistryCallerSession) IsEligibleSigner(signer common.Address) (bool, error) {
	return _SequencerRegistry.Contract.IsEligibleSigner(&_SequencerRegistry.CallOpts, signer)
}

// IsRegistered is a free data retrieval call binding the contract method 0xc3c5a547.
//
// Solidity: function isRegistered(address signer) view returns(bool)
func (_SequencerRegistry *SequencerRegistryCaller) IsRegistered(opts *bind.CallOpts, signer common.Address) (bool, error) {
	var out []interface{}
	err := _SequencerRegistry.contract.Call(opts, &out, "isRegistered", signer)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRegistered is a free data retrieval call binding the contract method 0xc3c5a547.
//
// Solidity: function isRegistered(address signer) view returns(bool)
func (_SequencerRegistry *SequencerRegistrySession) IsRegistered(signer common.Address) (bool, error) {
	return _SequencerRegistry.Contract.IsRegistered(&_SequencerRegistry.CallOpts, signer)
}

// IsRegistered is a free data retrieval call binding the contract method 0xc3c5a547.
//
// Solidity: function isRegistered(address signer) view returns(bool)
func (_SequencerRegistry *SequencerRegistryCallerSession) IsRegistered(signer common.Address) (bool, error) {
	return _SequencerRegistry.Contract.IsRegistered(&_SequencerRegistry.CallOpts, signer)
}

// LastUnpausedAt is a free data retrieval call binding the contract method 0xe07baba6.
//
// Solidity: function lastUnpausedAt() view returns(uint64)
func (_SequencerRegistry *SequencerRegistryCaller) LastUnpausedAt(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _SequencerRegistry.contract.Call(opts, &out, "lastUnpausedAt")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastUnpausedAt is a free data retrieval call binding the contract method 0xe07baba6.
//
// Solidity: function lastUnpausedAt() view returns(uint64)
func (_SequencerRegistry *SequencerRegistrySession) LastUnpausedAt() (uint64, error) {
	return _SequencerRegistry.Contract.LastUnpausedAt(&_SequencerRegistry.CallOpts)
}

// LastUnpausedAt is a free data retrieval call binding the contract method 0xe07baba6.
//
// Solidity: function lastUnpausedAt() view returns(uint64)
func (_SequencerRegistry *SequencerRegistryCallerSession) LastUnpausedAt() (uint64, error) {
	return _SequencerRegistry.Contract.LastUnpausedAt(&_SequencerRegistry.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x4690642b.
//
// Solidity: function nonces(bytes pubkey) view returns(uint256)
func (_SequencerRegistry *SequencerRegistryCaller) Nonces(opts *bind.CallOpts, pubkey []byte) (*big.Int, error) {
	var out []interface{}
	err := _SequencerRegistry.contract.Call(opts, &out, "nonces", pubkey)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x4690642b.
//
// Solidity: function nonces(bytes pubkey) view returns(uint256)
func (_SequencerRegistry *SequencerRegistrySession) Nonces(pubkey []byte) (*big.Int, error) {
	return _SequencerRegistry.Contract.Nonces(&_SequencerRegistry.CallOpts, pubkey)
}

// Nonces is a free data retrieval call binding the contract method 0x4690642b.
//
// Solidity: function nonces(bytes pubkey) view returns(uint256)
func (_SequencerRegistry *SequencerRegistryCallerSession) Nonces(pubkey []byte) (*big.Int, error) {
	return _SequencerRegistry.Contract.Nonces(&_SequencerRegistry.CallOpts, pubkey)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SequencerRegistry *SequencerRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SequencerRegistry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SequencerRegistry *SequencerRegistrySession) Owner() (common.Address, error) {
	return _SequencerRegistry.Contract.Owner(&_SequencerRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SequencerRegistry *SequencerRegistryCallerSession) Owner() (common.Address, error) {
	return _SequencerRegistry.Contract.Owner(&_SequencerRegistry.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_SequencerRegistry *SequencerRegistryCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _SequencerRegistry.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_SequencerRegistry *SequencerRegistrySession) Paused() (bool, error) {
	return _SequencerRegistry.Contract.Paused(&_SequencerRegistry.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_SequencerRegistry *SequencerRegistryCallerSession) Paused() (bool, error) {
	return _SequencerRegistry.Contract.Paused(&_SequencerRegistry.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_SequencerRegistry *SequencerRegistryCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SequencerRegistry.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_SequencerRegistry *SequencerRegistrySession) PendingOwner() (common.Address, error) {
	return _SequencerRegistry.Contract.PendingOwner(&_SequencerRegistry.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_SequencerRegistry *SequencerRegistryCallerSession) PendingOwner() (common.Address, error) {
	return _SequencerRegistry.Contract.PendingOwner(&_SequencerRegistry.CallOpts)
}

// ProtocolVersion is a free data retrieval call binding the contract method 0x2ae9c600.
//
// Solidity: function protocolVersion() pure returns(uint8)
func (_SequencerRegistry *SequencerRegistryCaller) ProtocolVersion(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _SequencerRegistry.contract.Call(opts, &out, "protocolVersion")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ProtocolVersion is a free data retrieval call binding the contract method 0x2ae9c600.
//
// Solidity: function protocolVersion() pure returns(uint8)
func (_SequencerRegistry *SequencerRegistrySession) ProtocolVersion() (uint8, error) {
	return _SequencerRegistry.Contract.ProtocolVersion(&_SequencerRegistry.CallOpts)
}

// ProtocolVersion is a free data retrieval call binding the contract method 0x2ae9c600.
//
// Solidity: function protocolVersion() pure returns(uint8)
func (_SequencerRegistry *SequencerRegistryCallerSession) ProtocolVersion() (uint8, error) {
	return _SequencerRegistry.Contract.ProtocolVersion(&_SequencerRegistry.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_SequencerRegistry *SequencerRegistryCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SequencerRegistry.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_SequencerRegistry *SequencerRegistrySession) ProxiableUUID() ([32]byte, error) {
	return _SequencerRegistry.Contract.ProxiableUUID(&_SequencerRegistry.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_SequencerRegistry *SequencerRegistryCallerSession) ProxiableUUID() ([32]byte, error) {
	return _SequencerRegistry.Contract.ProxiableUUID(&_SequencerRegistry.CallOpts)
}

// Resolve is a free data retrieval call binding the contract method 0x3eb6b8cf.
//
// Solidity: function resolve(uint64 _chainId, bytes32 _name, bool _allowZeroAddress) view returns(address)
func (_SequencerRegistry *SequencerRegistryCaller) Resolve(opts *bind.CallOpts, _chainId uint64, _name [32]byte, _allowZeroAddress bool) (common.Address, error) {
	var out []interface{}
	err := _SequencerRegistry.contract.Call(opts, &out, "resolve", _chainId, _name, _allowZeroAddress)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Resolve is a free data retrieval call binding the contract method 0x3eb6b8cf.
//
// Solidity: function resolve(uint64 _chainId, bytes32 _name, bool _allowZeroAddress) view returns(address)
func (_SequencerRegistry *SequencerRegistrySession) Resolve(_chainId uint64, _name [32]byte, _allowZeroAddress bool) (common.Address, error) {
	return _SequencerRegistry.Contract.Resolve(&_SequencerRegistry.CallOpts, _chainId, _name, _allowZeroAddress)
}

// Resolve is a free data retrieval call binding the contract method 0x3eb6b8cf.
//
// Solidity: function resolve(uint64 _chainId, bytes32 _name, bool _allowZeroAddress) view returns(address)
func (_SequencerRegistry *SequencerRegistryCallerSession) Resolve(_chainId uint64, _name [32]byte, _allowZeroAddress bool) (common.Address, error) {
	return _SequencerRegistry.Contract.Resolve(&_SequencerRegistry.CallOpts, _chainId, _name, _allowZeroAddress)
}

// Resolve0 is a free data retrieval call binding the contract method 0xa86f9d9e.
//
// Solidity: function resolve(bytes32 _name, bool _allowZeroAddress) view returns(address)
func (_SequencerRegistry *SequencerRegistryCaller) Resolve0(opts *bind.CallOpts, _name [32]byte, _allowZeroAddress bool) (common.Address, error) {
	var out []interface{}
	err := _SequencerRegistry.contract.Call(opts, &out, "resolve0", _name, _allowZeroAddress)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Resolve0 is a free data retrieval call binding the contract method 0xa86f9d9e.
//
// Solidity: function resolve(bytes32 _name, bool _allowZeroAddress) view returns(address)
func (_SequencerRegistry *SequencerRegistrySession) Resolve0(_name [32]byte, _allowZeroAddress bool) (common.Address, error) {
	return _SequencerRegistry.Contract.Resolve0(&_SequencerRegistry.CallOpts, _name, _allowZeroAddress)
}

// Resolve0 is a free data retrieval call binding the contract method 0xa86f9d9e.
//
// Solidity: function resolve(bytes32 _name, bool _allowZeroAddress) view returns(address)
func (_SequencerRegistry *SequencerRegistryCallerSession) Resolve0(_name [32]byte, _allowZeroAddress bool) (common.Address, error) {
	return _SequencerRegistry.Contract.Resolve0(&_SequencerRegistry.CallOpts, _name, _allowZeroAddress)
}

// SequencerByIndex is a free data retrieval call binding the contract method 0xf7c82658.
//
// Solidity: function sequencerByIndex(uint256 index) view returns(address signer, bytes metadata, bytes pubkey)
func (_SequencerRegistry *SequencerRegistryCaller) SequencerByIndex(opts *bind.CallOpts, index *big.Int) (struct {
	Signer   common.Address
	Metadata []byte
	Pubkey   []byte
}, error) {
	var out []interface{}
	err := _SequencerRegistry.contract.Call(opts, &out, "sequencerByIndex", index)

	outstruct := new(struct {
		Signer   common.Address
		Metadata []byte
		Pubkey   []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Signer = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Metadata = *abi.ConvertType(out[1], new([]byte)).(*[]byte)
	outstruct.Pubkey = *abi.ConvertType(out[2], new([]byte)).(*[]byte)

	return *outstruct, err

}

// SequencerByIndex is a free data retrieval call binding the contract method 0xf7c82658.
//
// Solidity: function sequencerByIndex(uint256 index) view returns(address signer, bytes metadata, bytes pubkey)
func (_SequencerRegistry *SequencerRegistrySession) SequencerByIndex(index *big.Int) (struct {
	Signer   common.Address
	Metadata []byte
	Pubkey   []byte
}, error) {
	return _SequencerRegistry.Contract.SequencerByIndex(&_SequencerRegistry.CallOpts, index)
}

// SequencerByIndex is a free data retrieval call binding the contract method 0xf7c82658.
//
// Solidity: function sequencerByIndex(uint256 index) view returns(address signer, bytes metadata, bytes pubkey)
func (_SequencerRegistry *SequencerRegistryCallerSession) SequencerByIndex(index *big.Int) (struct {
	Signer   common.Address
	Metadata []byte
	Pubkey   []byte
}, error) {
	return _SequencerRegistry.Contract.SequencerByIndex(&_SequencerRegistry.CallOpts, index)
}

// SequencersToPubkeyHash is a free data retrieval call binding the contract method 0xed22f0db.
//
// Solidity: function sequencersToPubkeyHash(address sequencer) view returns(bytes32 pubkeyHash)
func (_SequencerRegistry *SequencerRegistryCaller) SequencersToPubkeyHash(opts *bind.CallOpts, sequencer common.Address) ([32]byte, error) {
	var out []interface{}
	err := _SequencerRegistry.contract.Call(opts, &out, "sequencersToPubkeyHash", sequencer)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SequencersToPubkeyHash is a free data retrieval call binding the contract method 0xed22f0db.
//
// Solidity: function sequencersToPubkeyHash(address sequencer) view returns(bytes32 pubkeyHash)
func (_SequencerRegistry *SequencerRegistrySession) SequencersToPubkeyHash(sequencer common.Address) ([32]byte, error) {
	return _SequencerRegistry.Contract.SequencersToPubkeyHash(&_SequencerRegistry.CallOpts, sequencer)
}

// SequencersToPubkeyHash is a free data retrieval call binding the contract method 0xed22f0db.
//
// Solidity: function sequencersToPubkeyHash(address sequencer) view returns(bytes32 pubkeyHash)
func (_SequencerRegistry *SequencerRegistryCallerSession) SequencersToPubkeyHash(sequencer common.Address) ([32]byte, error) {
	return _SequencerRegistry.Contract.SequencersToPubkeyHash(&_SequencerRegistry.CallOpts, sequencer)
}

// StatusOf is a free data retrieval call binding the contract method 0x1df398ec.
//
// Solidity: function statusOf(bytes pubkey) view returns((bytes,bytes,address,uint256,uint256) data)
func (_SequencerRegistry *SequencerRegistryCaller) StatusOf(opts *bind.CallOpts, pubkey []byte) (ISequencerRegistrySequencer, error) {
	var out []interface{}
	err := _SequencerRegistry.contract.Call(opts, &out, "statusOf", pubkey)

	if err != nil {
		return *new(ISequencerRegistrySequencer), err
	}

	out0 := *abi.ConvertType(out[0], new(ISequencerRegistrySequencer)).(*ISequencerRegistrySequencer)

	return out0, err

}

// StatusOf is a free data retrieval call binding the contract method 0x1df398ec.
//
// Solidity: function statusOf(bytes pubkey) view returns((bytes,bytes,address,uint256,uint256) data)
func (_SequencerRegistry *SequencerRegistrySession) StatusOf(pubkey []byte) (ISequencerRegistrySequencer, error) {
	return _SequencerRegistry.Contract.StatusOf(&_SequencerRegistry.CallOpts, pubkey)
}

// StatusOf is a free data retrieval call binding the contract method 0x1df398ec.
//
// Solidity: function statusOf(bytes pubkey) view returns((bytes,bytes,address,uint256,uint256) data)
func (_SequencerRegistry *SequencerRegistryCallerSession) StatusOf(pubkey []byte) (ISequencerRegistrySequencer, error) {
	return _SequencerRegistry.Contract.StatusOf(&_SequencerRegistry.CallOpts, pubkey)
}

// Validators is a free data retrieval call binding the contract method 0x9bdafcb3.
//
// Solidity: function validators(bytes32 pubkeyHash) view returns(bytes pubkey, bytes metadata, address signer, uint256 activationBlock, uint256 deactivationBlock)
func (_SequencerRegistry *SequencerRegistryCaller) Validators(opts *bind.CallOpts, pubkeyHash [32]byte) (struct {
	Pubkey            []byte
	Metadata          []byte
	Signer            common.Address
	ActivationBlock   *big.Int
	DeactivationBlock *big.Int
}, error) {
	var out []interface{}
	err := _SequencerRegistry.contract.Call(opts, &out, "validators", pubkeyHash)

	outstruct := new(struct {
		Pubkey            []byte
		Metadata          []byte
		Signer            common.Address
		ActivationBlock   *big.Int
		DeactivationBlock *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Pubkey = *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	outstruct.Metadata = *abi.ConvertType(out[1], new([]byte)).(*[]byte)
	outstruct.Signer = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.ActivationBlock = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.DeactivationBlock = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Validators is a free data retrieval call binding the contract method 0x9bdafcb3.
//
// Solidity: function validators(bytes32 pubkeyHash) view returns(bytes pubkey, bytes metadata, address signer, uint256 activationBlock, uint256 deactivationBlock)
func (_SequencerRegistry *SequencerRegistrySession) Validators(pubkeyHash [32]byte) (struct {
	Pubkey            []byte
	Metadata          []byte
	Signer            common.Address
	ActivationBlock   *big.Int
	DeactivationBlock *big.Int
}, error) {
	return _SequencerRegistry.Contract.Validators(&_SequencerRegistry.CallOpts, pubkeyHash)
}

// Validators is a free data retrieval call binding the contract method 0x9bdafcb3.
//
// Solidity: function validators(bytes32 pubkeyHash) view returns(bytes pubkey, bytes metadata, address signer, uint256 activationBlock, uint256 deactivationBlock)
func (_SequencerRegistry *SequencerRegistryCallerSession) Validators(pubkeyHash [32]byte) (struct {
	Pubkey            []byte
	Metadata          []byte
	Signer            common.Address
	ActivationBlock   *big.Int
	DeactivationBlock *big.Int
}, error) {
	return _SequencerRegistry.Contract.Validators(&_SequencerRegistry.CallOpts, pubkeyHash)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_SequencerRegistry *SequencerRegistryTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SequencerRegistry.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_SequencerRegistry *SequencerRegistrySession) AcceptOwnership() (*types.Transaction, error) {
	return _SequencerRegistry.Contract.AcceptOwnership(&_SequencerRegistry.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_SequencerRegistry *SequencerRegistryTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _SequencerRegistry.Contract.AcceptOwnership(&_SequencerRegistry.TransactOpts)
}

// Activate is a paid mutator transaction binding the contract method 0xb9a7c230.
//
// Solidity: function activate(bytes pubkey, (uint64,uint64,uint64,uint256,bool,uint256,bytes) validatorProof) returns()
func (_SequencerRegistry *SequencerRegistryTransactor) Activate(opts *bind.TransactOpts, pubkey []byte, validatorProof ISequencerRegistryValidatorProof) (*types.Transaction, error) {
	return _SequencerRegistry.contract.Transact(opts, "activate", pubkey, validatorProof)
}

// Activate is a paid mutator transaction binding the contract method 0xb9a7c230.
//
// Solidity: function activate(bytes pubkey, (uint64,uint64,uint64,uint256,bool,uint256,bytes) validatorProof) returns()
func (_SequencerRegistry *SequencerRegistrySession) Activate(pubkey []byte, validatorProof ISequencerRegistryValidatorProof) (*types.Transaction, error) {
	return _SequencerRegistry.Contract.Activate(&_SequencerRegistry.TransactOpts, pubkey, validatorProof)
}

// Activate is a paid mutator transaction binding the contract method 0xb9a7c230.
//
// Solidity: function activate(bytes pubkey, (uint64,uint64,uint64,uint256,bool,uint256,bytes) validatorProof) returns()
func (_SequencerRegistry *SequencerRegistryTransactorSession) Activate(pubkey []byte, validatorProof ISequencerRegistryValidatorProof) (*types.Transaction, error) {
	return _SequencerRegistry.Contract.Activate(&_SequencerRegistry.TransactOpts, pubkey, validatorProof)
}

// ChangeRegistration is a paid mutator transaction binding the contract method 0x6869e219.
//
// Solidity: function changeRegistration(address signer, bytes metadata, bytes32 authHash, bytes signature) returns()
func (_SequencerRegistry *SequencerRegistryTransactor) ChangeRegistration(opts *bind.TransactOpts, signer common.Address, metadata []byte, authHash [32]byte, signature []byte) (*types.Transaction, error) {
	return _SequencerRegistry.contract.Transact(opts, "changeRegistration", signer, metadata, authHash, signature)
}

// ChangeRegistration is a paid mutator transaction binding the contract method 0x6869e219.
//
// Solidity: function changeRegistration(address signer, bytes metadata, bytes32 authHash, bytes signature) returns()
func (_SequencerRegistry *SequencerRegistrySession) ChangeRegistration(signer common.Address, metadata []byte, authHash [32]byte, signature []byte) (*types.Transaction, error) {
	return _SequencerRegistry.Contract.ChangeRegistration(&_SequencerRegistry.TransactOpts, signer, metadata, authHash, signature)
}

// ChangeRegistration is a paid mutator transaction binding the contract method 0x6869e219.
//
// Solidity: function changeRegistration(address signer, bytes metadata, bytes32 authHash, bytes signature) returns()
func (_SequencerRegistry *SequencerRegistryTransactorSession) ChangeRegistration(signer common.Address, metadata []byte, authHash [32]byte, signature []byte) (*types.Transaction, error) {
	return _SequencerRegistry.Contract.ChangeRegistration(&_SequencerRegistry.TransactOpts, signer, metadata, authHash, signature)
}

// Deactivate is a paid mutator transaction binding the contract method 0x29eda8e6.
//
// Solidity: function deactivate(bytes32 authHash, bytes signature) returns()
func (_SequencerRegistry *SequencerRegistryTransactor) Deactivate(opts *bind.TransactOpts, authHash [32]byte, signature []byte) (*types.Transaction, error) {
	return _SequencerRegistry.contract.Transact(opts, "deactivate", authHash, signature)
}

// Deactivate is a paid mutator transaction binding the contract method 0x29eda8e6.
//
// Solidity: function deactivate(bytes32 authHash, bytes signature) returns()
func (_SequencerRegistry *SequencerRegistrySession) Deactivate(authHash [32]byte, signature []byte) (*types.Transaction, error) {
	return _SequencerRegistry.Contract.Deactivate(&_SequencerRegistry.TransactOpts, authHash, signature)
}

// Deactivate is a paid mutator transaction binding the contract method 0x29eda8e6.
//
// Solidity: function deactivate(bytes32 authHash, bytes signature) returns()
func (_SequencerRegistry *SequencerRegistryTransactorSession) Deactivate(authHash [32]byte, signature []byte) (*types.Transaction, error) {
	return _SequencerRegistry.Contract.Deactivate(&_SequencerRegistry.TransactOpts, authHash, signature)
}

// ForceDeactivate is a paid mutator transaction binding the contract method 0x5bc95351.
//
// Solidity: function forceDeactivate(bytes pubkey, (uint64,uint64,uint64,uint256,bool,uint256,bytes) validatorProof) returns()
func (_SequencerRegistry *SequencerRegistryTransactor) ForceDeactivate(opts *bind.TransactOpts, pubkey []byte, validatorProof ISequencerRegistryValidatorProof) (*types.Transaction, error) {
	return _SequencerRegistry.contract.Transact(opts, "forceDeactivate", pubkey, validatorProof)
}

// ForceDeactivate is a paid mutator transaction binding the contract method 0x5bc95351.
//
// Solidity: function forceDeactivate(bytes pubkey, (uint64,uint64,uint64,uint256,bool,uint256,bytes) validatorProof) returns()
func (_SequencerRegistry *SequencerRegistrySession) ForceDeactivate(pubkey []byte, validatorProof ISequencerRegistryValidatorProof) (*types.Transaction, error) {
	return _SequencerRegistry.Contract.ForceDeactivate(&_SequencerRegistry.TransactOpts, pubkey, validatorProof)
}

// ForceDeactivate is a paid mutator transaction binding the contract method 0x5bc95351.
//
// Solidity: function forceDeactivate(bytes pubkey, (uint64,uint64,uint64,uint256,bool,uint256,bytes) validatorProof) returns()
func (_SequencerRegistry *SequencerRegistryTransactorSession) ForceDeactivate(pubkey []byte, validatorProof ISequencerRegistryValidatorProof) (*types.Transaction, error) {
	return _SequencerRegistry.Contract.ForceDeactivate(&_SequencerRegistry.TransactOpts, pubkey, validatorProof)
}

// Init is a paid mutator transaction binding the contract method 0x19ab453c.
//
// Solidity: function init(address _owner) returns()
func (_SequencerRegistry *SequencerRegistryTransactor) Init(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _SequencerRegistry.contract.Transact(opts, "init", _owner)
}

// Init is a paid mutator transaction binding the contract method 0x19ab453c.
//
// Solidity: function init(address _owner) returns()
func (_SequencerRegistry *SequencerRegistrySession) Init(_owner common.Address) (*types.Transaction, error) {
	return _SequencerRegistry.Contract.Init(&_SequencerRegistry.TransactOpts, _owner)
}

// Init is a paid mutator transaction binding the contract method 0x19ab453c.
//
// Solidity: function init(address _owner) returns()
func (_SequencerRegistry *SequencerRegistryTransactorSession) Init(_owner common.Address) (*types.Transaction, error) {
	return _SequencerRegistry.Contract.Init(&_SequencerRegistry.TransactOpts, _owner)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_SequencerRegistry *SequencerRegistryTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SequencerRegistry.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_SequencerRegistry *SequencerRegistrySession) Pause() (*types.Transaction, error) {
	return _SequencerRegistry.Contract.Pause(&_SequencerRegistry.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_SequencerRegistry *SequencerRegistryTransactorSession) Pause() (*types.Transaction, error) {
	return _SequencerRegistry.Contract.Pause(&_SequencerRegistry.TransactOpts)
}

// Register is a paid mutator transaction binding the contract method 0xe6e91157.
//
// Solidity: function register(address signer, bytes metadata, bytes32 authHash, bytes signature, (uint64,uint64,uint64,uint256,bool,uint256,bytes) validatorProof) returns()
func (_SequencerRegistry *SequencerRegistryTransactor) Register(opts *bind.TransactOpts, signer common.Address, metadata []byte, authHash [32]byte, signature []byte, validatorProof ISequencerRegistryValidatorProof) (*types.Transaction, error) {
	return _SequencerRegistry.contract.Transact(opts, "register", signer, metadata, authHash, signature, validatorProof)
}

// Register is a paid mutator transaction binding the contract method 0xe6e91157.
//
// Solidity: function register(address signer, bytes metadata, bytes32 authHash, bytes signature, (uint64,uint64,uint64,uint256,bool,uint256,bytes) validatorProof) returns()
func (_SequencerRegistry *SequencerRegistrySession) Register(signer common.Address, metadata []byte, authHash [32]byte, signature []byte, validatorProof ISequencerRegistryValidatorProof) (*types.Transaction, error) {
	return _SequencerRegistry.Contract.Register(&_SequencerRegistry.TransactOpts, signer, metadata, authHash, signature, validatorProof)
}

// Register is a paid mutator transaction binding the contract method 0xe6e91157.
//
// Solidity: function register(address signer, bytes metadata, bytes32 authHash, bytes signature, (uint64,uint64,uint64,uint256,bool,uint256,bytes) validatorProof) returns()
func (_SequencerRegistry *SequencerRegistryTransactorSession) Register(signer common.Address, metadata []byte, authHash [32]byte, signature []byte, validatorProof ISequencerRegistryValidatorProof) (*types.Transaction, error) {
	return _SequencerRegistry.Contract.Register(&_SequencerRegistry.TransactOpts, signer, metadata, authHash, signature, validatorProof)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SequencerRegistry *SequencerRegistryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SequencerRegistry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SequencerRegistry *SequencerRegistrySession) RenounceOwnership() (*types.Transaction, error) {
	return _SequencerRegistry.Contract.RenounceOwnership(&_SequencerRegistry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SequencerRegistry *SequencerRegistryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _SequencerRegistry.Contract.RenounceOwnership(&_SequencerRegistry.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SequencerRegistry *SequencerRegistryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SequencerRegistry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SequencerRegistry *SequencerRegistrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SequencerRegistry.Contract.TransferOwnership(&_SequencerRegistry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SequencerRegistry *SequencerRegistryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SequencerRegistry.Contract.TransferOwnership(&_SequencerRegistry.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_SequencerRegistry *SequencerRegistryTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SequencerRegistry.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_SequencerRegistry *SequencerRegistrySession) Unpause() (*types.Transaction, error) {
	return _SequencerRegistry.Contract.Unpause(&_SequencerRegistry.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_SequencerRegistry *SequencerRegistryTransactorSession) Unpause() (*types.Transaction, error) {
	return _SequencerRegistry.Contract.Unpause(&_SequencerRegistry.TransactOpts)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_SequencerRegistry *SequencerRegistryTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _SequencerRegistry.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_SequencerRegistry *SequencerRegistrySession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _SequencerRegistry.Contract.UpgradeTo(&_SequencerRegistry.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_SequencerRegistry *SequencerRegistryTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _SequencerRegistry.Contract.UpgradeTo(&_SequencerRegistry.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_SequencerRegistry *SequencerRegistryTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _SequencerRegistry.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_SequencerRegistry *SequencerRegistrySession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _SequencerRegistry.Contract.UpgradeToAndCall(&_SequencerRegistry.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_SequencerRegistry *SequencerRegistryTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _SequencerRegistry.Contract.UpgradeToAndCall(&_SequencerRegistry.TransactOpts, newImplementation, data)
}

// SequencerRegistryAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the SequencerRegistry contract.
type SequencerRegistryAdminChangedIterator struct {
	Event *SequencerRegistryAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SequencerRegistryAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SequencerRegistryAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SequencerRegistryAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SequencerRegistryAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SequencerRegistryAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SequencerRegistryAdminChanged represents a AdminChanged event raised by the SequencerRegistry contract.
type SequencerRegistryAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_SequencerRegistry *SequencerRegistryFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*SequencerRegistryAdminChangedIterator, error) {

	logs, sub, err := _SequencerRegistry.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &SequencerRegistryAdminChangedIterator{contract: _SequencerRegistry.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_SequencerRegistry *SequencerRegistryFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *SequencerRegistryAdminChanged) (event.Subscription, error) {

	logs, sub, err := _SequencerRegistry.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SequencerRegistryAdminChanged)
				if err := _SequencerRegistry.contract.UnpackLog(event, "AdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_SequencerRegistry *SequencerRegistryFilterer) ParseAdminChanged(log types.Log) (*SequencerRegistryAdminChanged, error) {
	event := new(SequencerRegistryAdminChanged)
	if err := _SequencerRegistry.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SequencerRegistryBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the SequencerRegistry contract.
type SequencerRegistryBeaconUpgradedIterator struct {
	Event *SequencerRegistryBeaconUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SequencerRegistryBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SequencerRegistryBeaconUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SequencerRegistryBeaconUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SequencerRegistryBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SequencerRegistryBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SequencerRegistryBeaconUpgraded represents a BeaconUpgraded event raised by the SequencerRegistry contract.
type SequencerRegistryBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_SequencerRegistry *SequencerRegistryFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*SequencerRegistryBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _SequencerRegistry.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &SequencerRegistryBeaconUpgradedIterator{contract: _SequencerRegistry.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_SequencerRegistry *SequencerRegistryFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *SequencerRegistryBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _SequencerRegistry.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SequencerRegistryBeaconUpgraded)
				if err := _SequencerRegistry.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_SequencerRegistry *SequencerRegistryFilterer) ParseBeaconUpgraded(log types.Log) (*SequencerRegistryBeaconUpgraded, error) {
	event := new(SequencerRegistryBeaconUpgraded)
	if err := _SequencerRegistry.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SequencerRegistryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the SequencerRegistry contract.
type SequencerRegistryInitializedIterator struct {
	Event *SequencerRegistryInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SequencerRegistryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SequencerRegistryInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SequencerRegistryInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SequencerRegistryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SequencerRegistryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SequencerRegistryInitialized represents a Initialized event raised by the SequencerRegistry contract.
type SequencerRegistryInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_SequencerRegistry *SequencerRegistryFilterer) FilterInitialized(opts *bind.FilterOpts) (*SequencerRegistryInitializedIterator, error) {

	logs, sub, err := _SequencerRegistry.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &SequencerRegistryInitializedIterator{contract: _SequencerRegistry.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_SequencerRegistry *SequencerRegistryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *SequencerRegistryInitialized) (event.Subscription, error) {

	logs, sub, err := _SequencerRegistry.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SequencerRegistryInitialized)
				if err := _SequencerRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_SequencerRegistry *SequencerRegistryFilterer) ParseInitialized(log types.Log) (*SequencerRegistryInitialized, error) {
	event := new(SequencerRegistryInitialized)
	if err := _SequencerRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SequencerRegistryOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the SequencerRegistry contract.
type SequencerRegistryOwnershipTransferStartedIterator struct {
	Event *SequencerRegistryOwnershipTransferStarted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SequencerRegistryOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SequencerRegistryOwnershipTransferStarted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SequencerRegistryOwnershipTransferStarted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SequencerRegistryOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SequencerRegistryOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SequencerRegistryOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the SequencerRegistry contract.
type SequencerRegistryOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_SequencerRegistry *SequencerRegistryFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SequencerRegistryOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SequencerRegistry.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SequencerRegistryOwnershipTransferStartedIterator{contract: _SequencerRegistry.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_SequencerRegistry *SequencerRegistryFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *SequencerRegistryOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SequencerRegistry.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SequencerRegistryOwnershipTransferStarted)
				if err := _SequencerRegistry.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferStarted is a log parse operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_SequencerRegistry *SequencerRegistryFilterer) ParseOwnershipTransferStarted(log types.Log) (*SequencerRegistryOwnershipTransferStarted, error) {
	event := new(SequencerRegistryOwnershipTransferStarted)
	if err := _SequencerRegistry.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SequencerRegistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SequencerRegistry contract.
type SequencerRegistryOwnershipTransferredIterator struct {
	Event *SequencerRegistryOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SequencerRegistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SequencerRegistryOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SequencerRegistryOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SequencerRegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SequencerRegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SequencerRegistryOwnershipTransferred represents a OwnershipTransferred event raised by the SequencerRegistry contract.
type SequencerRegistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SequencerRegistry *SequencerRegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SequencerRegistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SequencerRegistry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SequencerRegistryOwnershipTransferredIterator{contract: _SequencerRegistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SequencerRegistry *SequencerRegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SequencerRegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SequencerRegistry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SequencerRegistryOwnershipTransferred)
				if err := _SequencerRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SequencerRegistry *SequencerRegistryFilterer) ParseOwnershipTransferred(log types.Log) (*SequencerRegistryOwnershipTransferred, error) {
	event := new(SequencerRegistryOwnershipTransferred)
	if err := _SequencerRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SequencerRegistryPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the SequencerRegistry contract.
type SequencerRegistryPausedIterator struct {
	Event *SequencerRegistryPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SequencerRegistryPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SequencerRegistryPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SequencerRegistryPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SequencerRegistryPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SequencerRegistryPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SequencerRegistryPaused represents a Paused event raised by the SequencerRegistry contract.
type SequencerRegistryPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_SequencerRegistry *SequencerRegistryFilterer) FilterPaused(opts *bind.FilterOpts) (*SequencerRegistryPausedIterator, error) {

	logs, sub, err := _SequencerRegistry.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &SequencerRegistryPausedIterator{contract: _SequencerRegistry.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_SequencerRegistry *SequencerRegistryFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *SequencerRegistryPaused) (event.Subscription, error) {

	logs, sub, err := _SequencerRegistry.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SequencerRegistryPaused)
				if err := _SequencerRegistry.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_SequencerRegistry *SequencerRegistryFilterer) ParsePaused(log types.Log) (*SequencerRegistryPaused, error) {
	event := new(SequencerRegistryPaused)
	if err := _SequencerRegistry.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SequencerRegistrySequencerActivatedIterator is returned from FilterSequencerActivated and is used to iterate over the raw logs and unpacked data for SequencerActivated events raised by the SequencerRegistry contract.
type SequencerRegistrySequencerActivatedIterator struct {
	Event *SequencerRegistrySequencerActivated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SequencerRegistrySequencerActivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SequencerRegistrySequencerActivated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SequencerRegistrySequencerActivated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SequencerRegistrySequencerActivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SequencerRegistrySequencerActivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SequencerRegistrySequencerActivated represents a SequencerActivated event raised by the SequencerRegistry contract.
type SequencerRegistrySequencerActivated struct {
	Signer common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSequencerActivated is a free log retrieval operation binding the contract event 0x6630fca1bb566a362b39a495527169e1d413ba6e4dcca7e28e61126bbf596475.
//
// Solidity: event SequencerActivated(address indexed signer)
func (_SequencerRegistry *SequencerRegistryFilterer) FilterSequencerActivated(opts *bind.FilterOpts, signer []common.Address) (*SequencerRegistrySequencerActivatedIterator, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _SequencerRegistry.contract.FilterLogs(opts, "SequencerActivated", signerRule)
	if err != nil {
		return nil, err
	}
	return &SequencerRegistrySequencerActivatedIterator{contract: _SequencerRegistry.contract, event: "SequencerActivated", logs: logs, sub: sub}, nil
}

// WatchSequencerActivated is a free log subscription operation binding the contract event 0x6630fca1bb566a362b39a495527169e1d413ba6e4dcca7e28e61126bbf596475.
//
// Solidity: event SequencerActivated(address indexed signer)
func (_SequencerRegistry *SequencerRegistryFilterer) WatchSequencerActivated(opts *bind.WatchOpts, sink chan<- *SequencerRegistrySequencerActivated, signer []common.Address) (event.Subscription, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _SequencerRegistry.contract.WatchLogs(opts, "SequencerActivated", signerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SequencerRegistrySequencerActivated)
				if err := _SequencerRegistry.contract.UnpackLog(event, "SequencerActivated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSequencerActivated is a log parse operation binding the contract event 0x6630fca1bb566a362b39a495527169e1d413ba6e4dcca7e28e61126bbf596475.
//
// Solidity: event SequencerActivated(address indexed signer)
func (_SequencerRegistry *SequencerRegistryFilterer) ParseSequencerActivated(log types.Log) (*SequencerRegistrySequencerActivated, error) {
	event := new(SequencerRegistrySequencerActivated)
	if err := _SequencerRegistry.contract.UnpackLog(event, "SequencerActivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SequencerRegistrySequencerChangedIterator is returned from FilterSequencerChanged and is used to iterate over the raw logs and unpacked data for SequencerChanged events raised by the SequencerRegistry contract.
type SequencerRegistrySequencerChangedIterator struct {
	Event *SequencerRegistrySequencerChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SequencerRegistrySequencerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SequencerRegistrySequencerChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SequencerRegistrySequencerChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SequencerRegistrySequencerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SequencerRegistrySequencerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SequencerRegistrySequencerChanged represents a SequencerChanged event raised by the SequencerRegistry contract.
type SequencerRegistrySequencerChanged struct {
	OldSigner common.Address
	NewSigner common.Address
	Pubkey    []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSequencerChanged is a free log retrieval operation binding the contract event 0xc83f03300ac02196ac22ef96e9666a948b88b2c4772c4c3c3baa2f84a0eb5c9b.
//
// Solidity: event SequencerChanged(address indexed oldSigner, address indexed newSigner, bytes pubkey)
func (_SequencerRegistry *SequencerRegistryFilterer) FilterSequencerChanged(opts *bind.FilterOpts, oldSigner []common.Address, newSigner []common.Address) (*SequencerRegistrySequencerChangedIterator, error) {

	var oldSignerRule []interface{}
	for _, oldSignerItem := range oldSigner {
		oldSignerRule = append(oldSignerRule, oldSignerItem)
	}
	var newSignerRule []interface{}
	for _, newSignerItem := range newSigner {
		newSignerRule = append(newSignerRule, newSignerItem)
	}

	logs, sub, err := _SequencerRegistry.contract.FilterLogs(opts, "SequencerChanged", oldSignerRule, newSignerRule)
	if err != nil {
		return nil, err
	}
	return &SequencerRegistrySequencerChangedIterator{contract: _SequencerRegistry.contract, event: "SequencerChanged", logs: logs, sub: sub}, nil
}

// WatchSequencerChanged is a free log subscription operation binding the contract event 0xc83f03300ac02196ac22ef96e9666a948b88b2c4772c4c3c3baa2f84a0eb5c9b.
//
// Solidity: event SequencerChanged(address indexed oldSigner, address indexed newSigner, bytes pubkey)
func (_SequencerRegistry *SequencerRegistryFilterer) WatchSequencerChanged(opts *bind.WatchOpts, sink chan<- *SequencerRegistrySequencerChanged, oldSigner []common.Address, newSigner []common.Address) (event.Subscription, error) {

	var oldSignerRule []interface{}
	for _, oldSignerItem := range oldSigner {
		oldSignerRule = append(oldSignerRule, oldSignerItem)
	}
	var newSignerRule []interface{}
	for _, newSignerItem := range newSigner {
		newSignerRule = append(newSignerRule, newSignerItem)
	}

	logs, sub, err := _SequencerRegistry.contract.WatchLogs(opts, "SequencerChanged", oldSignerRule, newSignerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SequencerRegistrySequencerChanged)
				if err := _SequencerRegistry.contract.UnpackLog(event, "SequencerChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSequencerChanged is a log parse operation binding the contract event 0xc83f03300ac02196ac22ef96e9666a948b88b2c4772c4c3c3baa2f84a0eb5c9b.
//
// Solidity: event SequencerChanged(address indexed oldSigner, address indexed newSigner, bytes pubkey)
func (_SequencerRegistry *SequencerRegistryFilterer) ParseSequencerChanged(log types.Log) (*SequencerRegistrySequencerChanged, error) {
	event := new(SequencerRegistrySequencerChanged)
	if err := _SequencerRegistry.contract.UnpackLog(event, "SequencerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SequencerRegistrySequencerDeactivatedIterator is returned from FilterSequencerDeactivated and is used to iterate over the raw logs and unpacked data for SequencerDeactivated events raised by the SequencerRegistry contract.
type SequencerRegistrySequencerDeactivatedIterator struct {
	Event *SequencerRegistrySequencerDeactivated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SequencerRegistrySequencerDeactivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SequencerRegistrySequencerDeactivated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SequencerRegistrySequencerDeactivated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SequencerRegistrySequencerDeactivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SequencerRegistrySequencerDeactivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SequencerRegistrySequencerDeactivated represents a SequencerDeactivated event raised by the SequencerRegistry contract.
type SequencerRegistrySequencerDeactivated struct {
	Signer common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSequencerDeactivated is a free log retrieval operation binding the contract event 0xe511138637c40bb6949783dee2df37335f8bad4815aba500f4f429adb64f1ade.
//
// Solidity: event SequencerDeactivated(address indexed signer)
func (_SequencerRegistry *SequencerRegistryFilterer) FilterSequencerDeactivated(opts *bind.FilterOpts, signer []common.Address) (*SequencerRegistrySequencerDeactivatedIterator, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _SequencerRegistry.contract.FilterLogs(opts, "SequencerDeactivated", signerRule)
	if err != nil {
		return nil, err
	}
	return &SequencerRegistrySequencerDeactivatedIterator{contract: _SequencerRegistry.contract, event: "SequencerDeactivated", logs: logs, sub: sub}, nil
}

// WatchSequencerDeactivated is a free log subscription operation binding the contract event 0xe511138637c40bb6949783dee2df37335f8bad4815aba500f4f429adb64f1ade.
//
// Solidity: event SequencerDeactivated(address indexed signer)
func (_SequencerRegistry *SequencerRegistryFilterer) WatchSequencerDeactivated(opts *bind.WatchOpts, sink chan<- *SequencerRegistrySequencerDeactivated, signer []common.Address) (event.Subscription, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _SequencerRegistry.contract.WatchLogs(opts, "SequencerDeactivated", signerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SequencerRegistrySequencerDeactivated)
				if err := _SequencerRegistry.contract.UnpackLog(event, "SequencerDeactivated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSequencerDeactivated is a log parse operation binding the contract event 0xe511138637c40bb6949783dee2df37335f8bad4815aba500f4f429adb64f1ade.
//
// Solidity: event SequencerDeactivated(address indexed signer)
func (_SequencerRegistry *SequencerRegistryFilterer) ParseSequencerDeactivated(log types.Log) (*SequencerRegistrySequencerDeactivated, error) {
	event := new(SequencerRegistrySequencerDeactivated)
	if err := _SequencerRegistry.contract.UnpackLog(event, "SequencerDeactivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SequencerRegistrySequencerRegisteredIterator is returned from FilterSequencerRegistered and is used to iterate over the raw logs and unpacked data for SequencerRegistered events raised by the SequencerRegistry contract.
type SequencerRegistrySequencerRegisteredIterator struct {
	Event *SequencerRegistrySequencerRegistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SequencerRegistrySequencerRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SequencerRegistrySequencerRegistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SequencerRegistrySequencerRegistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SequencerRegistrySequencerRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SequencerRegistrySequencerRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SequencerRegistrySequencerRegistered represents a SequencerRegistered event raised by the SequencerRegistry contract.
type SequencerRegistrySequencerRegistered struct {
	Signer common.Address
	Pubkey []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSequencerRegistered is a free log retrieval operation binding the contract event 0x9a8d0d88dd19eadc11fc14c4044e8f378ff8d10056bc126e2c3283339617b98d.
//
// Solidity: event SequencerRegistered(address indexed signer, bytes pubkey)
func (_SequencerRegistry *SequencerRegistryFilterer) FilterSequencerRegistered(opts *bind.FilterOpts, signer []common.Address) (*SequencerRegistrySequencerRegisteredIterator, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _SequencerRegistry.contract.FilterLogs(opts, "SequencerRegistered", signerRule)
	if err != nil {
		return nil, err
	}
	return &SequencerRegistrySequencerRegisteredIterator{contract: _SequencerRegistry.contract, event: "SequencerRegistered", logs: logs, sub: sub}, nil
}

// WatchSequencerRegistered is a free log subscription operation binding the contract event 0x9a8d0d88dd19eadc11fc14c4044e8f378ff8d10056bc126e2c3283339617b98d.
//
// Solidity: event SequencerRegistered(address indexed signer, bytes pubkey)
func (_SequencerRegistry *SequencerRegistryFilterer) WatchSequencerRegistered(opts *bind.WatchOpts, sink chan<- *SequencerRegistrySequencerRegistered, signer []common.Address) (event.Subscription, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _SequencerRegistry.contract.WatchLogs(opts, "SequencerRegistered", signerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SequencerRegistrySequencerRegistered)
				if err := _SequencerRegistry.contract.UnpackLog(event, "SequencerRegistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSequencerRegistered is a log parse operation binding the contract event 0x9a8d0d88dd19eadc11fc14c4044e8f378ff8d10056bc126e2c3283339617b98d.
//
// Solidity: event SequencerRegistered(address indexed signer, bytes pubkey)
func (_SequencerRegistry *SequencerRegistryFilterer) ParseSequencerRegistered(log types.Log) (*SequencerRegistrySequencerRegistered, error) {
	event := new(SequencerRegistrySequencerRegistered)
	if err := _SequencerRegistry.contract.UnpackLog(event, "SequencerRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SequencerRegistryUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the SequencerRegistry contract.
type SequencerRegistryUnpausedIterator struct {
	Event *SequencerRegistryUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SequencerRegistryUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SequencerRegistryUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SequencerRegistryUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SequencerRegistryUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SequencerRegistryUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SequencerRegistryUnpaused represents a Unpaused event raised by the SequencerRegistry contract.
type SequencerRegistryUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_SequencerRegistry *SequencerRegistryFilterer) FilterUnpaused(opts *bind.FilterOpts) (*SequencerRegistryUnpausedIterator, error) {

	logs, sub, err := _SequencerRegistry.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &SequencerRegistryUnpausedIterator{contract: _SequencerRegistry.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_SequencerRegistry *SequencerRegistryFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *SequencerRegistryUnpaused) (event.Subscription, error) {

	logs, sub, err := _SequencerRegistry.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SequencerRegistryUnpaused)
				if err := _SequencerRegistry.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_SequencerRegistry *SequencerRegistryFilterer) ParseUnpaused(log types.Log) (*SequencerRegistryUnpaused, error) {
	event := new(SequencerRegistryUnpaused)
	if err := _SequencerRegistry.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SequencerRegistryUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the SequencerRegistry contract.
type SequencerRegistryUpgradedIterator struct {
	Event *SequencerRegistryUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SequencerRegistryUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SequencerRegistryUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SequencerRegistryUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SequencerRegistryUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SequencerRegistryUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SequencerRegistryUpgraded represents a Upgraded event raised by the SequencerRegistry contract.
type SequencerRegistryUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_SequencerRegistry *SequencerRegistryFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*SequencerRegistryUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _SequencerRegistry.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &SequencerRegistryUpgradedIterator{contract: _SequencerRegistry.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_SequencerRegistry *SequencerRegistryFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *SequencerRegistryUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _SequencerRegistry.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SequencerRegistryUpgraded)
				if err := _SequencerRegistry.contract.UnpackLog(event, "Upgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_SequencerRegistry *SequencerRegistryFilterer) ParseUpgraded(log types.Log) (*SequencerRegistryUpgraded, error) {
	event := new(SequencerRegistryUpgraded)
	if err := _SequencerRegistry.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
