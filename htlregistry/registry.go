// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package htlregistry

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// BooleanCondABI is the input ABI used to generate the binding from.
const BooleanCondABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"query\",\"type\":\"bytes\"},{\"name\":\"timeout\",\"type\":\"uint256\"}],\"name\":\"isFinalized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"query\",\"type\":\"bytes\"}],\"name\":\"isSatisfied\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// BooleanCondBin is the compiled bytecode used for deploying new contracts.
const BooleanCondBin = `0x`

// DeployBooleanCond deploys a new Ethereum contract, binding an instance of BooleanCond to it.
func DeployBooleanCond(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BooleanCond, error) {
	parsed, err := abi.JSON(strings.NewReader(BooleanCondABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BooleanCondBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BooleanCond{BooleanCondCaller: BooleanCondCaller{contract: contract}, BooleanCondTransactor: BooleanCondTransactor{contract: contract}, BooleanCondFilterer: BooleanCondFilterer{contract: contract}}, nil
}

// BooleanCond is an auto generated Go binding around an Ethereum contract.
type BooleanCond struct {
	BooleanCondCaller     // Read-only binding to the contract
	BooleanCondTransactor // Write-only binding to the contract
	BooleanCondFilterer   // Log filterer for contract events
}

// BooleanCondCaller is an auto generated read-only Go binding around an Ethereum contract.
type BooleanCondCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BooleanCondTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BooleanCondTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BooleanCondFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BooleanCondFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BooleanCondSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BooleanCondSession struct {
	Contract     *BooleanCond      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BooleanCondCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BooleanCondCallerSession struct {
	Contract *BooleanCondCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// BooleanCondTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BooleanCondTransactorSession struct {
	Contract     *BooleanCondTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// BooleanCondRaw is an auto generated low-level Go binding around an Ethereum contract.
type BooleanCondRaw struct {
	Contract *BooleanCond // Generic contract binding to access the raw methods on
}

// BooleanCondCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BooleanCondCallerRaw struct {
	Contract *BooleanCondCaller // Generic read-only contract binding to access the raw methods on
}

// BooleanCondTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BooleanCondTransactorRaw struct {
	Contract *BooleanCondTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBooleanCond creates a new instance of BooleanCond, bound to a specific deployed contract.
func NewBooleanCond(address common.Address, backend bind.ContractBackend) (*BooleanCond, error) {
	contract, err := bindBooleanCond(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BooleanCond{BooleanCondCaller: BooleanCondCaller{contract: contract}, BooleanCondTransactor: BooleanCondTransactor{contract: contract}, BooleanCondFilterer: BooleanCondFilterer{contract: contract}}, nil
}

// NewBooleanCondCaller creates a new read-only instance of BooleanCond, bound to a specific deployed contract.
func NewBooleanCondCaller(address common.Address, caller bind.ContractCaller) (*BooleanCondCaller, error) {
	contract, err := bindBooleanCond(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BooleanCondCaller{contract: contract}, nil
}

// NewBooleanCondTransactor creates a new write-only instance of BooleanCond, bound to a specific deployed contract.
func NewBooleanCondTransactor(address common.Address, transactor bind.ContractTransactor) (*BooleanCondTransactor, error) {
	contract, err := bindBooleanCond(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BooleanCondTransactor{contract: contract}, nil
}

// NewBooleanCondFilterer creates a new log filterer instance of BooleanCond, bound to a specific deployed contract.
func NewBooleanCondFilterer(address common.Address, filterer bind.ContractFilterer) (*BooleanCondFilterer, error) {
	contract, err := bindBooleanCond(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BooleanCondFilterer{contract: contract}, nil
}

// bindBooleanCond binds a generic wrapper to an already deployed contract.
func bindBooleanCond(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BooleanCondABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BooleanCond *BooleanCondRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BooleanCond.Contract.BooleanCondCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BooleanCond *BooleanCondRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BooleanCond.Contract.BooleanCondTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BooleanCond *BooleanCondRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BooleanCond.Contract.BooleanCondTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BooleanCond *BooleanCondCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BooleanCond.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BooleanCond *BooleanCondTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BooleanCond.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BooleanCond *BooleanCondTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BooleanCond.Contract.contract.Transact(opts, method, params...)
}

// IsFinalized is a free data retrieval call binding the contract method 0xd2121c2b.
//
// Solidity: function isFinalized(query bytes, timeout uint256) constant returns(bool)
func (_BooleanCond *BooleanCondCaller) IsFinalized(opts *bind.CallOpts, query []byte, timeout *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BooleanCond.contract.Call(opts, out, "isFinalized", query, timeout)
	return *ret0, err
}

// IsFinalized is a free data retrieval call binding the contract method 0xd2121c2b.
//
// Solidity: function isFinalized(query bytes, timeout uint256) constant returns(bool)
func (_BooleanCond *BooleanCondSession) IsFinalized(query []byte, timeout *big.Int) (bool, error) {
	return _BooleanCond.Contract.IsFinalized(&_BooleanCond.CallOpts, query, timeout)
}

// IsFinalized is a free data retrieval call binding the contract method 0xd2121c2b.
//
// Solidity: function isFinalized(query bytes, timeout uint256) constant returns(bool)
func (_BooleanCond *BooleanCondCallerSession) IsFinalized(query []byte, timeout *big.Int) (bool, error) {
	return _BooleanCond.Contract.IsFinalized(&_BooleanCond.CallOpts, query, timeout)
}

// IsSatisfied is a free data retrieval call binding the contract method 0xd6f12ca3.
//
// Solidity: function isSatisfied(query bytes) constant returns(bool)
func (_BooleanCond *BooleanCondCaller) IsSatisfied(opts *bind.CallOpts, query []byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BooleanCond.contract.Call(opts, out, "isSatisfied", query)
	return *ret0, err
}

// IsSatisfied is a free data retrieval call binding the contract method 0xd6f12ca3.
//
// Solidity: function isSatisfied(query bytes) constant returns(bool)
func (_BooleanCond *BooleanCondSession) IsSatisfied(query []byte) (bool, error) {
	return _BooleanCond.Contract.IsSatisfied(&_BooleanCond.CallOpts, query)
}

// IsSatisfied is a free data retrieval call binding the contract method 0xd6f12ca3.
//
// Solidity: function isSatisfied(query bytes) constant returns(bool)
func (_BooleanCond *BooleanCondCallerSession) IsSatisfied(query []byte) (bool, error) {
	return _BooleanCond.Contract.IsSatisfied(&_BooleanCond.CallOpts, query)
}

// HTLRegistryABI is the input ABI used to generate the binding from.
const HTLRegistryABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"query\",\"type\":\"bytes\"},{\"name\":\"timeout\",\"type\":\"uint256\"}],\"name\":\"isFinalized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"query\",\"type\":\"bytes\"}],\"name\":\"isSatisfied\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"secret\",\"type\":\"bytes\"}],\"name\":\"resolve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"secret\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"secretHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"SecretRegistry\",\"type\":\"event\"}]"

// HTLRegistryBin is the compiled bytecode used for deploying new contracts.
const HTLRegistryBin = `0x608060405234801561001057600080fd5b506102c9806100206000396000f3006080604052600436106100565763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663d2121c2b811461005b578063d6f12ca314610093578063e4056186146100b3575b600080fd5b34801561006757600080fd5b5061007f6024600480358281019291013590356100d5565b604080519115158252519081900360200190f35b34801561009f57600080fd5b5061007f6004803560248101910135610166565b3480156100bf57600080fd5b506100d360048035602481019101356101c8565b005b600080602084146100e557600080fd5b61011e85858080601f0160208091040260200160405190810160405280939291908181526020018383808284375061027d945050505050565b6000818152602081905260409020549091501580159061014b575060008181526020819052604090205483115b15610159576001915061015e565b600091505b509392505050565b6000806020831461017657600080fd5b6101af84848080601f0160208091040260200160405190810160405280939291908181526020018383808284375061027d945050505050565b6000908152602081905260409020541515949350505050565b6000828260405180838380828437604080519190930181900390206000818152602081905292909220549195505015925061020591505057600080fd5b6000818152602081815260409182902043908190558251918201849052918101829052606080825281018490527ff07454d4b7438e6b3d0e7273b625029cc4d7566ca87f4a2174062d55d4033915918591859185918060808101868680828437604051920182900397509095505050505050a1505050565b600081516000141561029157506000610298565b5060208101515b9190505600a165627a7a7230582083f2be93ed6f92f20a574ad3f54f13e0b3c90adc12cec690464a60b3d30627d60029`

// DeployHTLRegistry deploys a new Ethereum contract, binding an instance of HTLRegistry to it.
func DeployHTLRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *HTLRegistry, error) {
	parsed, err := abi.JSON(strings.NewReader(HTLRegistryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(HTLRegistryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &HTLRegistry{HTLRegistryCaller: HTLRegistryCaller{contract: contract}, HTLRegistryTransactor: HTLRegistryTransactor{contract: contract}, HTLRegistryFilterer: HTLRegistryFilterer{contract: contract}}, nil
}

// HTLRegistry is an auto generated Go binding around an Ethereum contract.
type HTLRegistry struct {
	HTLRegistryCaller     // Read-only binding to the contract
	HTLRegistryTransactor // Write-only binding to the contract
	HTLRegistryFilterer   // Log filterer for contract events
}

// HTLRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type HTLRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HTLRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HTLRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HTLRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HTLRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HTLRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HTLRegistrySession struct {
	Contract     *HTLRegistry      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HTLRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HTLRegistryCallerSession struct {
	Contract *HTLRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// HTLRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HTLRegistryTransactorSession struct {
	Contract     *HTLRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// HTLRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type HTLRegistryRaw struct {
	Contract *HTLRegistry // Generic contract binding to access the raw methods on
}

// HTLRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HTLRegistryCallerRaw struct {
	Contract *HTLRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// HTLRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HTLRegistryTransactorRaw struct {
	Contract *HTLRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHTLRegistry creates a new instance of HTLRegistry, bound to a specific deployed contract.
func NewHTLRegistry(address common.Address, backend bind.ContractBackend) (*HTLRegistry, error) {
	contract, err := bindHTLRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HTLRegistry{HTLRegistryCaller: HTLRegistryCaller{contract: contract}, HTLRegistryTransactor: HTLRegistryTransactor{contract: contract}, HTLRegistryFilterer: HTLRegistryFilterer{contract: contract}}, nil
}

// NewHTLRegistryCaller creates a new read-only instance of HTLRegistry, bound to a specific deployed contract.
func NewHTLRegistryCaller(address common.Address, caller bind.ContractCaller) (*HTLRegistryCaller, error) {
	contract, err := bindHTLRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HTLRegistryCaller{contract: contract}, nil
}

// NewHTLRegistryTransactor creates a new write-only instance of HTLRegistry, bound to a specific deployed contract.
func NewHTLRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*HTLRegistryTransactor, error) {
	contract, err := bindHTLRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HTLRegistryTransactor{contract: contract}, nil
}

// NewHTLRegistryFilterer creates a new log filterer instance of HTLRegistry, bound to a specific deployed contract.
func NewHTLRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*HTLRegistryFilterer, error) {
	contract, err := bindHTLRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HTLRegistryFilterer{contract: contract}, nil
}

// bindHTLRegistry binds a generic wrapper to an already deployed contract.
func bindHTLRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HTLRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HTLRegistry *HTLRegistryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HTLRegistry.Contract.HTLRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HTLRegistry *HTLRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HTLRegistry.Contract.HTLRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HTLRegistry *HTLRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HTLRegistry.Contract.HTLRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HTLRegistry *HTLRegistryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HTLRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HTLRegistry *HTLRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HTLRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HTLRegistry *HTLRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HTLRegistry.Contract.contract.Transact(opts, method, params...)
}

// IsFinalized is a free data retrieval call binding the contract method 0xd2121c2b.
//
// Solidity: function isFinalized(query bytes, timeout uint256) constant returns(bool)
func (_HTLRegistry *HTLRegistryCaller) IsFinalized(opts *bind.CallOpts, query []byte, timeout *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _HTLRegistry.contract.Call(opts, out, "isFinalized", query, timeout)
	return *ret0, err
}

// IsFinalized is a free data retrieval call binding the contract method 0xd2121c2b.
//
// Solidity: function isFinalized(query bytes, timeout uint256) constant returns(bool)
func (_HTLRegistry *HTLRegistrySession) IsFinalized(query []byte, timeout *big.Int) (bool, error) {
	return _HTLRegistry.Contract.IsFinalized(&_HTLRegistry.CallOpts, query, timeout)
}

// IsFinalized is a free data retrieval call binding the contract method 0xd2121c2b.
//
// Solidity: function isFinalized(query bytes, timeout uint256) constant returns(bool)
func (_HTLRegistry *HTLRegistryCallerSession) IsFinalized(query []byte, timeout *big.Int) (bool, error) {
	return _HTLRegistry.Contract.IsFinalized(&_HTLRegistry.CallOpts, query, timeout)
}

// IsSatisfied is a free data retrieval call binding the contract method 0xd6f12ca3.
//
// Solidity: function isSatisfied(query bytes) constant returns(bool)
func (_HTLRegistry *HTLRegistryCaller) IsSatisfied(opts *bind.CallOpts, query []byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _HTLRegistry.contract.Call(opts, out, "isSatisfied", query)
	return *ret0, err
}

// IsSatisfied is a free data retrieval call binding the contract method 0xd6f12ca3.
//
// Solidity: function isSatisfied(query bytes) constant returns(bool)
func (_HTLRegistry *HTLRegistrySession) IsSatisfied(query []byte) (bool, error) {
	return _HTLRegistry.Contract.IsSatisfied(&_HTLRegistry.CallOpts, query)
}

// IsSatisfied is a free data retrieval call binding the contract method 0xd6f12ca3.
//
// Solidity: function isSatisfied(query bytes) constant returns(bool)
func (_HTLRegistry *HTLRegistryCallerSession) IsSatisfied(query []byte) (bool, error) {
	return _HTLRegistry.Contract.IsSatisfied(&_HTLRegistry.CallOpts, query)
}

// Resolve is a paid mutator transaction binding the contract method 0xe4056186.
//
// Solidity: function resolve(secret bytes) returns()
func (_HTLRegistry *HTLRegistryTransactor) Resolve(opts *bind.TransactOpts, secret []byte) (*types.Transaction, error) {
	return _HTLRegistry.contract.Transact(opts, "resolve", secret)
}

// Resolve is a paid mutator transaction binding the contract method 0xe4056186.
//
// Solidity: function resolve(secret bytes) returns()
func (_HTLRegistry *HTLRegistrySession) Resolve(secret []byte) (*types.Transaction, error) {
	return _HTLRegistry.Contract.Resolve(&_HTLRegistry.TransactOpts, secret)
}

// Resolve is a paid mutator transaction binding the contract method 0xe4056186.
//
// Solidity: function resolve(secret bytes) returns()
func (_HTLRegistry *HTLRegistryTransactorSession) Resolve(secret []byte) (*types.Transaction, error) {
	return _HTLRegistry.Contract.Resolve(&_HTLRegistry.TransactOpts, secret)
}

// HTLRegistrySecretRegistryIterator is returned from FilterSecretRegistry and is used to iterate over the raw logs and unpacked data for SecretRegistry events raised by the HTLRegistry contract.
type HTLRegistrySecretRegistryIterator struct {
	Event *HTLRegistrySecretRegistry // Event containing the contract specifics and raw log

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
func (it *HTLRegistrySecretRegistryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HTLRegistrySecretRegistry)
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
		it.Event = new(HTLRegistrySecretRegistry)
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
func (it *HTLRegistrySecretRegistryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HTLRegistrySecretRegistryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HTLRegistrySecretRegistry represents a SecretRegistry event raised by the HTLRegistry contract.
type HTLRegistrySecretRegistry struct {
	Secret     []byte
	SecretHash [32]byte
	Time       *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSecretRegistry is a free log retrieval operation binding the contract event 0xf07454d4b7438e6b3d0e7273b625029cc4d7566ca87f4a2174062d55d4033915.
//
// Solidity: e SecretRegistry(secret bytes, secretHash bytes32, time uint256)
func (_HTLRegistry *HTLRegistryFilterer) FilterSecretRegistry(opts *bind.FilterOpts) (*HTLRegistrySecretRegistryIterator, error) {

	logs, sub, err := _HTLRegistry.contract.FilterLogs(opts, "SecretRegistry")
	if err != nil {
		return nil, err
	}
	return &HTLRegistrySecretRegistryIterator{contract: _HTLRegistry.contract, event: "SecretRegistry", logs: logs, sub: sub}, nil
}

// WatchSecretRegistry is a free log subscription operation binding the contract event 0xf07454d4b7438e6b3d0e7273b625029cc4d7566ca87f4a2174062d55d4033915.
//
// Solidity: e SecretRegistry(secret bytes, secretHash bytes32, time uint256)
func (_HTLRegistry *HTLRegistryFilterer) WatchSecretRegistry(opts *bind.WatchOpts, sink chan<- *HTLRegistrySecretRegistry) (event.Subscription, error) {

	logs, sub, err := _HTLRegistry.contract.WatchLogs(opts, "SecretRegistry")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HTLRegistrySecretRegistry)
				if err := _HTLRegistry.contract.UnpackLog(event, "SecretRegistry", log); err != nil {
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
