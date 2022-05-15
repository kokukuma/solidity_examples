// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package sample

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
)

// SampleMetaData contains all meta data concerning the Sample contract.
var SampleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"get\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"}],\"name\":\"set\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506101e6806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c806360fe47b11461003b5780636d4ce63c14610057575b600080fd5b610055600480360381019061005091906100d4565b610075565b005b61005f610090565b60405161006c9190610110565b60405180910390f35b80600080828254610086919061015a565b9250508190555050565b60008054905090565b600080fd5b6000819050919050565b6100b18161009e565b81146100bc57600080fd5b50565b6000813590506100ce816100a8565b92915050565b6000602082840312156100ea576100e9610099565b5b60006100f8848285016100bf565b91505092915050565b61010a8161009e565b82525050565b60006020820190506101256000830184610101565b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006101658261009e565b91506101708361009e565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff038211156101a5576101a461012b565b5b82820190509291505056fea2646970667358221220e3e3188d6dca6e176a12cce5f5f1347aed716df0ebcc1edc22ba5d807965efc664736f6c634300080d0033",
}

// SampleABI is the input ABI used to generate the binding from.
// Deprecated: Use SampleMetaData.ABI instead.
var SampleABI = SampleMetaData.ABI

// SampleBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SampleMetaData.Bin instead.
var SampleBin = SampleMetaData.Bin

// DeploySample deploys a new Ethereum contract, binding an instance of Sample to it.
func DeploySample(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Sample, error) {
	parsed, err := SampleMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SampleBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Sample{SampleCaller: SampleCaller{contract: contract}, SampleTransactor: SampleTransactor{contract: contract}, SampleFilterer: SampleFilterer{contract: contract}}, nil
}

// Sample is an auto generated Go binding around an Ethereum contract.
type Sample struct {
	SampleCaller     // Read-only binding to the contract
	SampleTransactor // Write-only binding to the contract
	SampleFilterer   // Log filterer for contract events
}

// SampleCaller is an auto generated read-only Go binding around an Ethereum contract.
type SampleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SampleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SampleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SampleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SampleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SampleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SampleSession struct {
	Contract     *Sample           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SampleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SampleCallerSession struct {
	Contract *SampleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SampleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SampleTransactorSession struct {
	Contract     *SampleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SampleRaw is an auto generated low-level Go binding around an Ethereum contract.
type SampleRaw struct {
	Contract *Sample // Generic contract binding to access the raw methods on
}

// SampleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SampleCallerRaw struct {
	Contract *SampleCaller // Generic read-only contract binding to access the raw methods on
}

// SampleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SampleTransactorRaw struct {
	Contract *SampleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSample creates a new instance of Sample, bound to a specific deployed contract.
func NewSample(address common.Address, backend bind.ContractBackend) (*Sample, error) {
	contract, err := bindSample(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Sample{SampleCaller: SampleCaller{contract: contract}, SampleTransactor: SampleTransactor{contract: contract}, SampleFilterer: SampleFilterer{contract: contract}}, nil
}

// NewSampleCaller creates a new read-only instance of Sample, bound to a specific deployed contract.
func NewSampleCaller(address common.Address, caller bind.ContractCaller) (*SampleCaller, error) {
	contract, err := bindSample(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SampleCaller{contract: contract}, nil
}

// NewSampleTransactor creates a new write-only instance of Sample, bound to a specific deployed contract.
func NewSampleTransactor(address common.Address, transactor bind.ContractTransactor) (*SampleTransactor, error) {
	contract, err := bindSample(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SampleTransactor{contract: contract}, nil
}

// NewSampleFilterer creates a new log filterer instance of Sample, bound to a specific deployed contract.
func NewSampleFilterer(address common.Address, filterer bind.ContractFilterer) (*SampleFilterer, error) {
	contract, err := bindSample(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SampleFilterer{contract: contract}, nil
}

// bindSample binds a generic wrapper to an already deployed contract.
func bindSample(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SampleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sample *SampleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sample.Contract.SampleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sample *SampleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sample.Contract.SampleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sample *SampleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sample.Contract.SampleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sample *SampleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sample.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sample *SampleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sample.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sample *SampleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sample.Contract.contract.Transact(opts, method, params...)
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns(uint256)
func (_Sample *SampleCaller) Get(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Sample.contract.Call(opts, &out, "get")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns(uint256)
func (_Sample *SampleSession) Get() (*big.Int, error) {
	return _Sample.Contract.Get(&_Sample.CallOpts)
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns(uint256)
func (_Sample *SampleCallerSession) Get() (*big.Int, error) {
	return _Sample.Contract.Get(&_Sample.CallOpts)
}

// Set is a paid mutator transaction binding the contract method 0x60fe47b1.
//
// Solidity: function set(uint256 x) returns()
func (_Sample *SampleTransactor) Set(opts *bind.TransactOpts, x *big.Int) (*types.Transaction, error) {
	return _Sample.contract.Transact(opts, "set", x)
}

// Set is a paid mutator transaction binding the contract method 0x60fe47b1.
//
// Solidity: function set(uint256 x) returns()
func (_Sample *SampleSession) Set(x *big.Int) (*types.Transaction, error) {
	return _Sample.Contract.Set(&_Sample.TransactOpts, x)
}

// Set is a paid mutator transaction binding the contract method 0x60fe47b1.
//
// Solidity: function set(uint256 x) returns()
func (_Sample *SampleTransactorSession) Set(x *big.Int) (*types.Transaction, error) {
	return _Sample.Contract.Set(&_Sample.TransactOpts, x)
}
