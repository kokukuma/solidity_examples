// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package coin

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

// CoinMetaData contains all meta data concerning the Coin contract.
var CoinMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Sent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"get\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"send\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506106d8806100606000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c8063075461721461005c57806327e235e31461007a57806340c10f19146100aa578063c2bc2efc146100c6578063d0679d34146100f6575b600080fd5b610064610112565b604051610071919061041c565b60405180910390f35b610094600480360381019061008f9190610468565b610136565b6040516100a191906104ae565b60405180910390f35b6100c460048036038101906100bf91906104f5565b61014e565b005b6100e060048036038101906100db9190610468565b610225565b6040516100ed91906104ae565b60405180910390f35b610110600480360381019061010b91906104f5565b61026e565b005b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60016020528060005260406000206000915090505481565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146101a657600080fd5b789f4f2726179a224501d762422c946590d9100000000000000081106101cb57600080fd5b80600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461021a9190610564565b925050819055505050565b6000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548111156102f0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102e790610617565b60405180910390fd5b80600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461033f9190610637565b9250508190555080600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546103959190610564565b925050819055507f3990db2d31862302a685e8086b5755072a6e2b5b780af1ee81ece35ee3cd33453383836040516103cf9392919061066b565b60405180910390a15050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610406826103db565b9050919050565b610416816103fb565b82525050565b6000602082019050610431600083018461040d565b92915050565b600080fd5b610445816103fb565b811461045057600080fd5b50565b6000813590506104628161043c565b92915050565b60006020828403121561047e5761047d610437565b5b600061048c84828501610453565b91505092915050565b6000819050919050565b6104a881610495565b82525050565b60006020820190506104c3600083018461049f565b92915050565b6104d281610495565b81146104dd57600080fd5b50565b6000813590506104ef816104c9565b92915050565b6000806040838503121561050c5761050b610437565b5b600061051a85828601610453565b925050602061052b858286016104e0565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061056f82610495565b915061057a83610495565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff038211156105af576105ae610535565b5b828201905092915050565b600082825260208201905092915050565b7f496e73756666696369656e742062616c616e63652e0000000000000000000000600082015250565b60006106016015836105ba565b915061060c826105cb565b602082019050919050565b60006020820190508181036000830152610630816105f4565b9050919050565b600061064282610495565b915061064d83610495565b9250828210156106605761065f610535565b5b828203905092915050565b6000606082019050610680600083018661040d565b61068d602083018561040d565b61069a604083018461049f565b94935050505056fea264697066735822122040e1777544ba4b958fb769f6d193d1385712dc4ecbca0902b767734e2ffd139b64736f6c634300080d0033",
}

// CoinABI is the input ABI used to generate the binding from.
// Deprecated: Use CoinMetaData.ABI instead.
var CoinABI = CoinMetaData.ABI

// CoinBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CoinMetaData.Bin instead.
var CoinBin = CoinMetaData.Bin

// DeployCoin deploys a new Ethereum contract, binding an instance of Coin to it.
func DeployCoin(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Coin, error) {
	parsed, err := CoinMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CoinBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Coin{CoinCaller: CoinCaller{contract: contract}, CoinTransactor: CoinTransactor{contract: contract}, CoinFilterer: CoinFilterer{contract: contract}}, nil
}

// Coin is an auto generated Go binding around an Ethereum contract.
type Coin struct {
	CoinCaller     // Read-only binding to the contract
	CoinTransactor // Write-only binding to the contract
	CoinFilterer   // Log filterer for contract events
}

// CoinCaller is an auto generated read-only Go binding around an Ethereum contract.
type CoinCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CoinTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CoinTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CoinFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CoinFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CoinSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CoinSession struct {
	Contract     *Coin             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CoinCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CoinCallerSession struct {
	Contract *CoinCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// CoinTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CoinTransactorSession struct {
	Contract     *CoinTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CoinRaw is an auto generated low-level Go binding around an Ethereum contract.
type CoinRaw struct {
	Contract *Coin // Generic contract binding to access the raw methods on
}

// CoinCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CoinCallerRaw struct {
	Contract *CoinCaller // Generic read-only contract binding to access the raw methods on
}

// CoinTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CoinTransactorRaw struct {
	Contract *CoinTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCoin creates a new instance of Coin, bound to a specific deployed contract.
func NewCoin(address common.Address, backend bind.ContractBackend) (*Coin, error) {
	contract, err := bindCoin(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Coin{CoinCaller: CoinCaller{contract: contract}, CoinTransactor: CoinTransactor{contract: contract}, CoinFilterer: CoinFilterer{contract: contract}}, nil
}

// NewCoinCaller creates a new read-only instance of Coin, bound to a specific deployed contract.
func NewCoinCaller(address common.Address, caller bind.ContractCaller) (*CoinCaller, error) {
	contract, err := bindCoin(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CoinCaller{contract: contract}, nil
}

// NewCoinTransactor creates a new write-only instance of Coin, bound to a specific deployed contract.
func NewCoinTransactor(address common.Address, transactor bind.ContractTransactor) (*CoinTransactor, error) {
	contract, err := bindCoin(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CoinTransactor{contract: contract}, nil
}

// NewCoinFilterer creates a new log filterer instance of Coin, bound to a specific deployed contract.
func NewCoinFilterer(address common.Address, filterer bind.ContractFilterer) (*CoinFilterer, error) {
	contract, err := bindCoin(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CoinFilterer{contract: contract}, nil
}

// bindCoin binds a generic wrapper to an already deployed contract.
func bindCoin(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CoinABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Coin *CoinRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Coin.Contract.CoinCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Coin *CoinRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Coin.Contract.CoinTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Coin *CoinRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Coin.Contract.CoinTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Coin *CoinCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Coin.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Coin *CoinTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Coin.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Coin *CoinTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Coin.Contract.contract.Transact(opts, method, params...)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_Coin *CoinCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Coin.contract.Call(opts, &out, "balances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_Coin *CoinSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _Coin.Contract.Balances(&_Coin.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_Coin *CoinCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _Coin.Contract.Balances(&_Coin.CallOpts, arg0)
}

// Get is a free data retrieval call binding the contract method 0xc2bc2efc.
//
// Solidity: function get(address receiver) view returns(uint256)
func (_Coin *CoinCaller) Get(opts *bind.CallOpts, receiver common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Coin.contract.Call(opts, &out, "get", receiver)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Get is a free data retrieval call binding the contract method 0xc2bc2efc.
//
// Solidity: function get(address receiver) view returns(uint256)
func (_Coin *CoinSession) Get(receiver common.Address) (*big.Int, error) {
	return _Coin.Contract.Get(&_Coin.CallOpts, receiver)
}

// Get is a free data retrieval call binding the contract method 0xc2bc2efc.
//
// Solidity: function get(address receiver) view returns(uint256)
func (_Coin *CoinCallerSession) Get(receiver common.Address) (*big.Int, error) {
	return _Coin.Contract.Get(&_Coin.CallOpts, receiver)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_Coin *CoinCaller) Minter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Coin.contract.Call(opts, &out, "minter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_Coin *CoinSession) Minter() (common.Address, error) {
	return _Coin.Contract.Minter(&_Coin.CallOpts)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_Coin *CoinCallerSession) Minter() (common.Address, error) {
	return _Coin.Contract.Minter(&_Coin.CallOpts)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address receiver, uint256 amount) returns()
func (_Coin *CoinTransactor) Mint(opts *bind.TransactOpts, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Coin.contract.Transact(opts, "mint", receiver, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address receiver, uint256 amount) returns()
func (_Coin *CoinSession) Mint(receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Coin.Contract.Mint(&_Coin.TransactOpts, receiver, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address receiver, uint256 amount) returns()
func (_Coin *CoinTransactorSession) Mint(receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Coin.Contract.Mint(&_Coin.TransactOpts, receiver, amount)
}

// Send is a paid mutator transaction binding the contract method 0xd0679d34.
//
// Solidity: function send(address receiver, uint256 amount) returns()
func (_Coin *CoinTransactor) Send(opts *bind.TransactOpts, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Coin.contract.Transact(opts, "send", receiver, amount)
}

// Send is a paid mutator transaction binding the contract method 0xd0679d34.
//
// Solidity: function send(address receiver, uint256 amount) returns()
func (_Coin *CoinSession) Send(receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Coin.Contract.Send(&_Coin.TransactOpts, receiver, amount)
}

// Send is a paid mutator transaction binding the contract method 0xd0679d34.
//
// Solidity: function send(address receiver, uint256 amount) returns()
func (_Coin *CoinTransactorSession) Send(receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Coin.Contract.Send(&_Coin.TransactOpts, receiver, amount)
}

// CoinSentIterator is returned from FilterSent and is used to iterate over the raw logs and unpacked data for Sent events raised by the Coin contract.
type CoinSentIterator struct {
	Event *CoinSent // Event containing the contract specifics and raw log

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
func (it *CoinSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoinSent)
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
		it.Event = new(CoinSent)
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
func (it *CoinSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoinSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoinSent represents a Sent event raised by the Coin contract.
type CoinSent struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSent is a free log retrieval operation binding the contract event 0x3990db2d31862302a685e8086b5755072a6e2b5b780af1ee81ece35ee3cd3345.
//
// Solidity: event Sent(address from, address to, uint256 amount)
func (_Coin *CoinFilterer) FilterSent(opts *bind.FilterOpts) (*CoinSentIterator, error) {

	logs, sub, err := _Coin.contract.FilterLogs(opts, "Sent")
	if err != nil {
		return nil, err
	}
	return &CoinSentIterator{contract: _Coin.contract, event: "Sent", logs: logs, sub: sub}, nil
}

// WatchSent is a free log subscription operation binding the contract event 0x3990db2d31862302a685e8086b5755072a6e2b5b780af1ee81ece35ee3cd3345.
//
// Solidity: event Sent(address from, address to, uint256 amount)
func (_Coin *CoinFilterer) WatchSent(opts *bind.WatchOpts, sink chan<- *CoinSent) (event.Subscription, error) {

	logs, sub, err := _Coin.contract.WatchLogs(opts, "Sent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoinSent)
				if err := _Coin.contract.UnpackLog(event, "Sent", log); err != nil {
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

// ParseSent is a log parse operation binding the contract event 0x3990db2d31862302a685e8086b5755072a6e2b5b780af1ee81ece35ee3cd3345.
//
// Solidity: event Sent(address from, address to, uint256 amount)
func (_Coin *CoinFilterer) ParseSent(log types.Log) (*CoinSent, error) {
	event := new(CoinSent)
	if err := _Coin.contract.UnpackLog(event, "Sent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
