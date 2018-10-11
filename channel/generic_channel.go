// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package channel

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

// AddressUtilsABI is the input ABI used to generate the binding from.
const AddressUtilsABI = "[]"

// AddressUtilsBin is the compiled bytecode used for deploying new contracts.
const AddressUtilsBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a723058206a1900e97a39a7183ad2cf65e18c0d59340f0c607810126302be20212e3a5e520029`

// DeployAddressUtils deploys a new Ethereum contract, binding an instance of AddressUtils to it.
func DeployAddressUtils(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AddressUtils, error) {
	parsed, err := abi.JSON(strings.NewReader(AddressUtilsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AddressUtilsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AddressUtils{AddressUtilsCaller: AddressUtilsCaller{contract: contract}, AddressUtilsTransactor: AddressUtilsTransactor{contract: contract}, AddressUtilsFilterer: AddressUtilsFilterer{contract: contract}}, nil
}

// AddressUtils is an auto generated Go binding around an Ethereum contract.
type AddressUtils struct {
	AddressUtilsCaller     // Read-only binding to the contract
	AddressUtilsTransactor // Write-only binding to the contract
	AddressUtilsFilterer   // Log filterer for contract events
}

// AddressUtilsCaller is an auto generated read-only Go binding around an Ethereum contract.
type AddressUtilsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressUtilsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AddressUtilsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressUtilsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AddressUtilsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressUtilsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AddressUtilsSession struct {
	Contract     *AddressUtils     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AddressUtilsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AddressUtilsCallerSession struct {
	Contract *AddressUtilsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// AddressUtilsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AddressUtilsTransactorSession struct {
	Contract     *AddressUtilsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// AddressUtilsRaw is an auto generated low-level Go binding around an Ethereum contract.
type AddressUtilsRaw struct {
	Contract *AddressUtils // Generic contract binding to access the raw methods on
}

// AddressUtilsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AddressUtilsCallerRaw struct {
	Contract *AddressUtilsCaller // Generic read-only contract binding to access the raw methods on
}

// AddressUtilsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AddressUtilsTransactorRaw struct {
	Contract *AddressUtilsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAddressUtils creates a new instance of AddressUtils, bound to a specific deployed contract.
func NewAddressUtils(address common.Address, backend bind.ContractBackend) (*AddressUtils, error) {
	contract, err := bindAddressUtils(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AddressUtils{AddressUtilsCaller: AddressUtilsCaller{contract: contract}, AddressUtilsTransactor: AddressUtilsTransactor{contract: contract}, AddressUtilsFilterer: AddressUtilsFilterer{contract: contract}}, nil
}

// NewAddressUtilsCaller creates a new read-only instance of AddressUtils, bound to a specific deployed contract.
func NewAddressUtilsCaller(address common.Address, caller bind.ContractCaller) (*AddressUtilsCaller, error) {
	contract, err := bindAddressUtils(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AddressUtilsCaller{contract: contract}, nil
}

// NewAddressUtilsTransactor creates a new write-only instance of AddressUtils, bound to a specific deployed contract.
func NewAddressUtilsTransactor(address common.Address, transactor bind.ContractTransactor) (*AddressUtilsTransactor, error) {
	contract, err := bindAddressUtils(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AddressUtilsTransactor{contract: contract}, nil
}

// NewAddressUtilsFilterer creates a new log filterer instance of AddressUtils, bound to a specific deployed contract.
func NewAddressUtilsFilterer(address common.Address, filterer bind.ContractFilterer) (*AddressUtilsFilterer, error) {
	contract, err := bindAddressUtils(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AddressUtilsFilterer{contract: contract}, nil
}

// bindAddressUtils binds a generic wrapper to an already deployed contract.
func bindAddressUtils(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AddressUtilsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AddressUtils *AddressUtilsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AddressUtils.Contract.AddressUtilsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AddressUtils *AddressUtilsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressUtils.Contract.AddressUtilsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AddressUtils *AddressUtilsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AddressUtils.Contract.AddressUtilsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AddressUtils *AddressUtilsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AddressUtils.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AddressUtils *AddressUtilsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressUtils.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AddressUtils *AddressUtilsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AddressUtils.Contract.contract.Transact(opts, method, params...)
}

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

// DepositPoolInterfaceABI is the input ABI used to generate the binding from.
const DepositPoolInterfaceABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_authWithdraw\",\"type\":\"bytes\"},{\"name\":\"_signature\",\"type\":\"bytes\"},{\"name\":\"_channelId\",\"type\":\"uint256\"}],\"name\":\"authorizedWithdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// DepositPoolInterfaceBin is the compiled bytecode used for deploying new contracts.
const DepositPoolInterfaceBin = `0x`

// DeployDepositPoolInterface deploys a new Ethereum contract, binding an instance of DepositPoolInterface to it.
func DeployDepositPoolInterface(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DepositPoolInterface, error) {
	parsed, err := abi.JSON(strings.NewReader(DepositPoolInterfaceABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DepositPoolInterfaceBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DepositPoolInterface{DepositPoolInterfaceCaller: DepositPoolInterfaceCaller{contract: contract}, DepositPoolInterfaceTransactor: DepositPoolInterfaceTransactor{contract: contract}, DepositPoolInterfaceFilterer: DepositPoolInterfaceFilterer{contract: contract}}, nil
}

// DepositPoolInterface is an auto generated Go binding around an Ethereum contract.
type DepositPoolInterface struct {
	DepositPoolInterfaceCaller     // Read-only binding to the contract
	DepositPoolInterfaceTransactor // Write-only binding to the contract
	DepositPoolInterfaceFilterer   // Log filterer for contract events
}

// DepositPoolInterfaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type DepositPoolInterfaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DepositPoolInterfaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DepositPoolInterfaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DepositPoolInterfaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DepositPoolInterfaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DepositPoolInterfaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DepositPoolInterfaceSession struct {
	Contract     *DepositPoolInterface // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// DepositPoolInterfaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DepositPoolInterfaceCallerSession struct {
	Contract *DepositPoolInterfaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// DepositPoolInterfaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DepositPoolInterfaceTransactorSession struct {
	Contract     *DepositPoolInterfaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// DepositPoolInterfaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type DepositPoolInterfaceRaw struct {
	Contract *DepositPoolInterface // Generic contract binding to access the raw methods on
}

// DepositPoolInterfaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DepositPoolInterfaceCallerRaw struct {
	Contract *DepositPoolInterfaceCaller // Generic read-only contract binding to access the raw methods on
}

// DepositPoolInterfaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DepositPoolInterfaceTransactorRaw struct {
	Contract *DepositPoolInterfaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDepositPoolInterface creates a new instance of DepositPoolInterface, bound to a specific deployed contract.
func NewDepositPoolInterface(address common.Address, backend bind.ContractBackend) (*DepositPoolInterface, error) {
	contract, err := bindDepositPoolInterface(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DepositPoolInterface{DepositPoolInterfaceCaller: DepositPoolInterfaceCaller{contract: contract}, DepositPoolInterfaceTransactor: DepositPoolInterfaceTransactor{contract: contract}, DepositPoolInterfaceFilterer: DepositPoolInterfaceFilterer{contract: contract}}, nil
}

// NewDepositPoolInterfaceCaller creates a new read-only instance of DepositPoolInterface, bound to a specific deployed contract.
func NewDepositPoolInterfaceCaller(address common.Address, caller bind.ContractCaller) (*DepositPoolInterfaceCaller, error) {
	contract, err := bindDepositPoolInterface(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DepositPoolInterfaceCaller{contract: contract}, nil
}

// NewDepositPoolInterfaceTransactor creates a new write-only instance of DepositPoolInterface, bound to a specific deployed contract.
func NewDepositPoolInterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*DepositPoolInterfaceTransactor, error) {
	contract, err := bindDepositPoolInterface(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DepositPoolInterfaceTransactor{contract: contract}, nil
}

// NewDepositPoolInterfaceFilterer creates a new log filterer instance of DepositPoolInterface, bound to a specific deployed contract.
func NewDepositPoolInterfaceFilterer(address common.Address, filterer bind.ContractFilterer) (*DepositPoolInterfaceFilterer, error) {
	contract, err := bindDepositPoolInterface(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DepositPoolInterfaceFilterer{contract: contract}, nil
}

// bindDepositPoolInterface binds a generic wrapper to an already deployed contract.
func bindDepositPoolInterface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DepositPoolInterfaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DepositPoolInterface *DepositPoolInterfaceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DepositPoolInterface.Contract.DepositPoolInterfaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DepositPoolInterface *DepositPoolInterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DepositPoolInterface.Contract.DepositPoolInterfaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DepositPoolInterface *DepositPoolInterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DepositPoolInterface.Contract.DepositPoolInterfaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DepositPoolInterface *DepositPoolInterfaceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DepositPoolInterface.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DepositPoolInterface *DepositPoolInterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DepositPoolInterface.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DepositPoolInterface *DepositPoolInterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DepositPoolInterface.Contract.contract.Transact(opts, method, params...)
}

// AuthorizedWithdraw is a paid mutator transaction binding the contract method 0x06fa40f3.
//
// Solidity: function authorizedWithdraw(_authWithdraw bytes, _signature bytes, _channelId uint256) returns()
func (_DepositPoolInterface *DepositPoolInterfaceTransactor) AuthorizedWithdraw(opts *bind.TransactOpts, _authWithdraw []byte, _signature []byte, _channelId *big.Int) (*types.Transaction, error) {
	return _DepositPoolInterface.contract.Transact(opts, "authorizedWithdraw", _authWithdraw, _signature, _channelId)
}

// AuthorizedWithdraw is a paid mutator transaction binding the contract method 0x06fa40f3.
//
// Solidity: function authorizedWithdraw(_authWithdraw bytes, _signature bytes, _channelId uint256) returns()
func (_DepositPoolInterface *DepositPoolInterfaceSession) AuthorizedWithdraw(_authWithdraw []byte, _signature []byte, _channelId *big.Int) (*types.Transaction, error) {
	return _DepositPoolInterface.Contract.AuthorizedWithdraw(&_DepositPoolInterface.TransactOpts, _authWithdraw, _signature, _channelId)
}

// AuthorizedWithdraw is a paid mutator transaction binding the contract method 0x06fa40f3.
//
// Solidity: function authorizedWithdraw(_authWithdraw bytes, _signature bytes, _channelId uint256) returns()
func (_DepositPoolInterface *DepositPoolInterfaceTransactorSession) AuthorizedWithdraw(_authWithdraw []byte, _signature []byte, _channelId *big.Int) (*types.Transaction, error) {
	return _DepositPoolInterface.Contract.AuthorizedWithdraw(&_DepositPoolInterface.TransactOpts, _authWithdraw, _signature, _channelId)
}

// ERC20ABI is the input ABI used to generate the binding from.
const ERC20ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"who\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// ERC20Bin is the compiled bytecode used for deploying new contracts.
const ERC20Bin = `0x`

// DeployERC20 deploys a new Ethereum contract, binding an instance of ERC20 to it.
func DeployERC20(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC20, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC20Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20{ERC20Caller: ERC20Caller{contract: contract}, ERC20Transactor: ERC20Transactor{contract: contract}, ERC20Filterer: ERC20Filterer{contract: contract}}, nil
}

// ERC20 is an auto generated Go binding around an Ethereum contract.
type ERC20 struct {
	ERC20Caller     // Read-only binding to the contract
	ERC20Transactor // Write-only binding to the contract
	ERC20Filterer   // Log filterer for contract events
}

// ERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20Session struct {
	Contract     *ERC20            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20CallerSession struct {
	Contract *ERC20Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20TransactorSession struct {
	Contract     *ERC20Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20Raw struct {
	Contract *ERC20 // Generic contract binding to access the raw methods on
}

// ERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20CallerRaw struct {
	Contract *ERC20Caller // Generic read-only contract binding to access the raw methods on
}

// ERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20TransactorRaw struct {
	Contract *ERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20 creates a new instance of ERC20, bound to a specific deployed contract.
func NewERC20(address common.Address, backend bind.ContractBackend) (*ERC20, error) {
	contract, err := bindERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20{ERC20Caller: ERC20Caller{contract: contract}, ERC20Transactor: ERC20Transactor{contract: contract}, ERC20Filterer: ERC20Filterer{contract: contract}}, nil
}

// NewERC20Caller creates a new read-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Caller(address common.Address, caller bind.ContractCaller) (*ERC20Caller, error) {
	contract, err := bindERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20Caller{contract: contract}, nil
}

// NewERC20Transactor creates a new write-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC20Transactor, error) {
	contract, err := bindERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20Transactor{contract: contract}, nil
}

// NewERC20Filterer creates a new log filterer instance of ERC20, bound to a specific deployed contract.
func NewERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*ERC20Filterer, error) {
	contract, err := bindERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20Filterer{contract: contract}, nil
}

// bindERC20 binds a generic wrapper to an already deployed contract.
func bindERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20 *ERC20Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20.Contract.ERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20 *ERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.Contract.ERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20 *ERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20.Contract.ERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20 *ERC20CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20 *ERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20 *ERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_ERC20 *ERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_ERC20 *ERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_ERC20 *ERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_ERC20 *ERC20Caller) BalanceOf(opts *bind.CallOpts, who common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20.contract.Call(opts, out, "balanceOf", who)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_ERC20 *ERC20Session) BalanceOf(who common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, who)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_ERC20 *ERC20CallerSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, who)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20 *ERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20 *ERC20Session) TotalSupply() (*big.Int, error) {
	return _ERC20.Contract.TotalSupply(&_ERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20 *ERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20.Contract.TotalSupply(&_ERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_ERC20 *ERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_ERC20 *ERC20Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_ERC20 *ERC20TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, spender, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_ERC20 *ERC20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_ERC20 *ERC20Session) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_ERC20 *ERC20TransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_ERC20 *ERC20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_ERC20 *ERC20Session) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_ERC20 *ERC20TransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, from, to, value)
}

// ERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC20 contract.
type ERC20ApprovalIterator struct {
	Event *ERC20Approval // Event containing the contract specifics and raw log

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
func (it *ERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20Approval)
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
		it.Event = new(ERC20Approval)
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
func (it *ERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20Approval represents a Approval event raised by the ERC20 contract.
type ERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_ERC20 *ERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20ApprovalIterator{contract: _ERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_ERC20 *ERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20Approval)
				if err := _ERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20 contract.
type ERC20TransferIterator struct {
	Event *ERC20Transfer // Event containing the contract specifics and raw log

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
func (it *ERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20Transfer)
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
		it.Event = new(ERC20Transfer)
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
func (it *ERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20Transfer represents a Transfer event raised by the ERC20 contract.
type ERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_ERC20 *ERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TransferIterator{contract: _ERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_ERC20 *ERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20Transfer)
				if err := _ERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ERC20BasicABI is the input ABI used to generate the binding from.
const ERC20BasicABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"who\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// ERC20BasicBin is the compiled bytecode used for deploying new contracts.
const ERC20BasicBin = `0x`

// DeployERC20Basic deploys a new Ethereum contract, binding an instance of ERC20Basic to it.
func DeployERC20Basic(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC20Basic, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20BasicABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC20BasicBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20Basic{ERC20BasicCaller: ERC20BasicCaller{contract: contract}, ERC20BasicTransactor: ERC20BasicTransactor{contract: contract}, ERC20BasicFilterer: ERC20BasicFilterer{contract: contract}}, nil
}

// ERC20Basic is an auto generated Go binding around an Ethereum contract.
type ERC20Basic struct {
	ERC20BasicCaller     // Read-only binding to the contract
	ERC20BasicTransactor // Write-only binding to the contract
	ERC20BasicFilterer   // Log filterer for contract events
}

// ERC20BasicCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20BasicCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BasicTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20BasicTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BasicFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20BasicFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BasicSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20BasicSession struct {
	Contract     *ERC20Basic       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20BasicCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20BasicCallerSession struct {
	Contract *ERC20BasicCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ERC20BasicTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20BasicTransactorSession struct {
	Contract     *ERC20BasicTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ERC20BasicRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20BasicRaw struct {
	Contract *ERC20Basic // Generic contract binding to access the raw methods on
}

// ERC20BasicCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20BasicCallerRaw struct {
	Contract *ERC20BasicCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20BasicTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20BasicTransactorRaw struct {
	Contract *ERC20BasicTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20Basic creates a new instance of ERC20Basic, bound to a specific deployed contract.
func NewERC20Basic(address common.Address, backend bind.ContractBackend) (*ERC20Basic, error) {
	contract, err := bindERC20Basic(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20Basic{ERC20BasicCaller: ERC20BasicCaller{contract: contract}, ERC20BasicTransactor: ERC20BasicTransactor{contract: contract}, ERC20BasicFilterer: ERC20BasicFilterer{contract: contract}}, nil
}

// NewERC20BasicCaller creates a new read-only instance of ERC20Basic, bound to a specific deployed contract.
func NewERC20BasicCaller(address common.Address, caller bind.ContractCaller) (*ERC20BasicCaller, error) {
	contract, err := bindERC20Basic(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20BasicCaller{contract: contract}, nil
}

// NewERC20BasicTransactor creates a new write-only instance of ERC20Basic, bound to a specific deployed contract.
func NewERC20BasicTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20BasicTransactor, error) {
	contract, err := bindERC20Basic(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20BasicTransactor{contract: contract}, nil
}

// NewERC20BasicFilterer creates a new log filterer instance of ERC20Basic, bound to a specific deployed contract.
func NewERC20BasicFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20BasicFilterer, error) {
	contract, err := bindERC20Basic(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20BasicFilterer{contract: contract}, nil
}

// bindERC20Basic binds a generic wrapper to an already deployed contract.
func bindERC20Basic(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20BasicABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Basic *ERC20BasicRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20Basic.Contract.ERC20BasicCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Basic *ERC20BasicRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Basic.Contract.ERC20BasicTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Basic *ERC20BasicRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Basic.Contract.ERC20BasicTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Basic *ERC20BasicCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20Basic.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Basic *ERC20BasicTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Basic.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Basic *ERC20BasicTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Basic.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_ERC20Basic *ERC20BasicCaller) BalanceOf(opts *bind.CallOpts, who common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20Basic.contract.Call(opts, out, "balanceOf", who)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_ERC20Basic *ERC20BasicSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _ERC20Basic.Contract.BalanceOf(&_ERC20Basic.CallOpts, who)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_ERC20Basic *ERC20BasicCallerSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _ERC20Basic.Contract.BalanceOf(&_ERC20Basic.CallOpts, who)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20Basic *ERC20BasicCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20Basic.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20Basic *ERC20BasicSession) TotalSupply() (*big.Int, error) {
	return _ERC20Basic.Contract.TotalSupply(&_ERC20Basic.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20Basic *ERC20BasicCallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20Basic.Contract.TotalSupply(&_ERC20Basic.CallOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_ERC20Basic *ERC20BasicTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20Basic.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_ERC20Basic *ERC20BasicSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20Basic.Contract.Transfer(&_ERC20Basic.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_ERC20Basic *ERC20BasicTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20Basic.Contract.Transfer(&_ERC20Basic.TransactOpts, to, value)
}

// ERC20BasicTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20Basic contract.
type ERC20BasicTransferIterator struct {
	Event *ERC20BasicTransfer // Event containing the contract specifics and raw log

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
func (it *ERC20BasicTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20BasicTransfer)
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
		it.Event = new(ERC20BasicTransfer)
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
func (it *ERC20BasicTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20BasicTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20BasicTransfer represents a Transfer event raised by the ERC20Basic contract.
type ERC20BasicTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_ERC20Basic *ERC20BasicFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20BasicTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20Basic.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20BasicTransferIterator{contract: _ERC20Basic.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_ERC20Basic *ERC20BasicFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20BasicTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20Basic.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20BasicTransfer)
				if err := _ERC20Basic.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// GenericCondABI is the input ABI used to generate the binding from.
const GenericCondABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"query\",\"type\":\"bytes\"}],\"name\":\"getStateUpdate\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"query\",\"type\":\"bytes\"},{\"name\":\"timeout\",\"type\":\"uint256\"}],\"name\":\"isFinalized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// GenericCondBin is the compiled bytecode used for deploying new contracts.
const GenericCondBin = `0x`

// DeployGenericCond deploys a new Ethereum contract, binding an instance of GenericCond to it.
func DeployGenericCond(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GenericCond, error) {
	parsed, err := abi.JSON(strings.NewReader(GenericCondABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(GenericCondBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GenericCond{GenericCondCaller: GenericCondCaller{contract: contract}, GenericCondTransactor: GenericCondTransactor{contract: contract}, GenericCondFilterer: GenericCondFilterer{contract: contract}}, nil
}

// GenericCond is an auto generated Go binding around an Ethereum contract.
type GenericCond struct {
	GenericCondCaller     // Read-only binding to the contract
	GenericCondTransactor // Write-only binding to the contract
	GenericCondFilterer   // Log filterer for contract events
}

// GenericCondCaller is an auto generated read-only Go binding around an Ethereum contract.
type GenericCondCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GenericCondTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GenericCondTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GenericCondFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GenericCondFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GenericCondSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GenericCondSession struct {
	Contract     *GenericCond      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GenericCondCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GenericCondCallerSession struct {
	Contract *GenericCondCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// GenericCondTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GenericCondTransactorSession struct {
	Contract     *GenericCondTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// GenericCondRaw is an auto generated low-level Go binding around an Ethereum contract.
type GenericCondRaw struct {
	Contract *GenericCond // Generic contract binding to access the raw methods on
}

// GenericCondCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GenericCondCallerRaw struct {
	Contract *GenericCondCaller // Generic read-only contract binding to access the raw methods on
}

// GenericCondTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GenericCondTransactorRaw struct {
	Contract *GenericCondTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGenericCond creates a new instance of GenericCond, bound to a specific deployed contract.
func NewGenericCond(address common.Address, backend bind.ContractBackend) (*GenericCond, error) {
	contract, err := bindGenericCond(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GenericCond{GenericCondCaller: GenericCondCaller{contract: contract}, GenericCondTransactor: GenericCondTransactor{contract: contract}, GenericCondFilterer: GenericCondFilterer{contract: contract}}, nil
}

// NewGenericCondCaller creates a new read-only instance of GenericCond, bound to a specific deployed contract.
func NewGenericCondCaller(address common.Address, caller bind.ContractCaller) (*GenericCondCaller, error) {
	contract, err := bindGenericCond(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GenericCondCaller{contract: contract}, nil
}

// NewGenericCondTransactor creates a new write-only instance of GenericCond, bound to a specific deployed contract.
func NewGenericCondTransactor(address common.Address, transactor bind.ContractTransactor) (*GenericCondTransactor, error) {
	contract, err := bindGenericCond(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GenericCondTransactor{contract: contract}, nil
}

// NewGenericCondFilterer creates a new log filterer instance of GenericCond, bound to a specific deployed contract.
func NewGenericCondFilterer(address common.Address, filterer bind.ContractFilterer) (*GenericCondFilterer, error) {
	contract, err := bindGenericCond(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GenericCondFilterer{contract: contract}, nil
}

// bindGenericCond binds a generic wrapper to an already deployed contract.
func bindGenericCond(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GenericCondABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GenericCond *GenericCondRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GenericCond.Contract.GenericCondCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GenericCond *GenericCondRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GenericCond.Contract.GenericCondTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GenericCond *GenericCondRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GenericCond.Contract.GenericCondTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GenericCond *GenericCondCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GenericCond.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GenericCond *GenericCondTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GenericCond.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GenericCond *GenericCondTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GenericCond.Contract.contract.Transact(opts, method, params...)
}

// GetStateUpdate is a free data retrieval call binding the contract method 0x195a4bc6.
//
// Solidity: function getStateUpdate(query bytes) constant returns(bytes)
func (_GenericCond *GenericCondCaller) GetStateUpdate(opts *bind.CallOpts, query []byte) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _GenericCond.contract.Call(opts, out, "getStateUpdate", query)
	return *ret0, err
}

// GetStateUpdate is a free data retrieval call binding the contract method 0x195a4bc6.
//
// Solidity: function getStateUpdate(query bytes) constant returns(bytes)
func (_GenericCond *GenericCondSession) GetStateUpdate(query []byte) ([]byte, error) {
	return _GenericCond.Contract.GetStateUpdate(&_GenericCond.CallOpts, query)
}

// GetStateUpdate is a free data retrieval call binding the contract method 0x195a4bc6.
//
// Solidity: function getStateUpdate(query bytes) constant returns(bytes)
func (_GenericCond *GenericCondCallerSession) GetStateUpdate(query []byte) ([]byte, error) {
	return _GenericCond.Contract.GetStateUpdate(&_GenericCond.CallOpts, query)
}

// IsFinalized is a free data retrieval call binding the contract method 0xd2121c2b.
//
// Solidity: function isFinalized(query bytes, timeout uint256) constant returns(bool)
func (_GenericCond *GenericCondCaller) IsFinalized(opts *bind.CallOpts, query []byte, timeout *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _GenericCond.contract.Call(opts, out, "isFinalized", query, timeout)
	return *ret0, err
}

// IsFinalized is a free data retrieval call binding the contract method 0xd2121c2b.
//
// Solidity: function isFinalized(query bytes, timeout uint256) constant returns(bool)
func (_GenericCond *GenericCondSession) IsFinalized(query []byte, timeout *big.Int) (bool, error) {
	return _GenericCond.Contract.IsFinalized(&_GenericCond.CallOpts, query, timeout)
}

// IsFinalized is a free data retrieval call binding the contract method 0xd2121c2b.
//
// Solidity: function isFinalized(query bytes, timeout uint256) constant returns(bool)
func (_GenericCond *GenericCondCallerSession) IsFinalized(query []byte, timeout *big.Int) (bool, error) {
	return _GenericCond.Contract.IsFinalized(&_GenericCond.CallOpts, query, timeout)
}

// GenericConditionalChannelABI is the input ABI used to generate the binding from.
const GenericConditionalChannelABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"resolver\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_channelId\",\"type\":\"uint256\"},{\"name\":\"_receipient\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"depositERCToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_channelId\",\"type\":\"uint256\"}],\"name\":\"getDepositMap\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"},{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"depositPool\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_channelId\",\"type\":\"uint256\"}],\"name\":\"confirmSettle\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_peers\",\"type\":\"address[]\"},{\"name\":\"_withdrawalTimeout\",\"type\":\"uint256[]\"},{\"name\":\"_settleTimeoutIncrement\",\"type\":\"uint256\"},{\"name\":\"_tokenContract\",\"type\":\"address\"},{\"name\":\"_tokenType\",\"type\":\"uint256\"}],\"name\":\"openChannel\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_channelId\",\"type\":\"uint256\"}],\"name\":\"getChannelSettleTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_channelId\",\"type\":\"uint256\"},{\"name\":\"_stateProof\",\"type\":\"bytes\"},{\"name\":\"_signature\",\"type\":\"bytes\"},{\"name\":\"_signatureOfSignature\",\"type\":\"bytes\"}],\"name\":\"cooperativeSettle\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_channelId\",\"type\":\"uint256\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"intendWithdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_channelId\",\"type\":\"uint256\"},{\"name\":\"_stateProof\",\"type\":\"bytes\"},{\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"intendSettleStateProof\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_channelId\",\"type\":\"uint256\"},{\"name\":\"_receipient\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_channelId\",\"type\":\"uint256\"}],\"name\":\"getTokenType\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"channelLength\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"chainId\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_channelId\",\"type\":\"uint256\"}],\"name\":\"getTokenContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_channelId\",\"type\":\"uint256\"},{\"name\":\"_withdrawId\",\"type\":\"uint256\"}],\"name\":\"confirmWithdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_channelId\",\"type\":\"uint256\"},{\"name\":\"_withdrawId\",\"type\":\"uint256\"},{\"name\":\"dispute\",\"type\":\"bytes\"}],\"name\":\"disputeWithdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_channelId\",\"type\":\"uint256\"},{\"name\":\"_cooperativeWithdrawProof\",\"type\":\"bytes\"},{\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"cooperativeWithdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_channelId\",\"type\":\"uint256\"}],\"name\":\"getChannelStatus\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_channelId\",\"type\":\"uint256\"},{\"name\":\"_proof\",\"type\":\"bytes32[]\"},{\"name\":\"_conditionGroup\",\"type\":\"bytes\"}],\"name\":\"resolveConditionalStateTransition\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_channelId\",\"type\":\"uint256\"},{\"name\":\"_peer\",\"type\":\"address\"}],\"name\":\"getDepositAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_authWithdraw\",\"type\":\"bytes\"},{\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"authOpenChannel\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_chainId\",\"type\":\"uint256\"},{\"name\":\"_virtResolver\",\"type\":\"address\"},{\"name\":\"_depositPool\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"channelId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"peers\",\"type\":\"address[]\"},{\"indexed\":false,\"name\":\"uintTokenType\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"tokenContract\",\"type\":\"address\"}],\"name\":\"OpenChannel\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"channelId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"peers\",\"type\":\"address[]\"},{\"indexed\":false,\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"channelId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"stateProofNonce\",\"type\":\"uint256\"}],\"name\":\"IntendSettle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"channelId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"condGroupHash\",\"type\":\"bytes32\"}],\"name\":\"ResolveCondGroup\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"channelId\",\"type\":\"uint256\"}],\"name\":\"ConfirmSettle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"channelId\",\"type\":\"uint256\"}],\"name\":\"ConfirmSettleFail\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"channelId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"withdrawalAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"balance\",\"type\":\"uint256\"}],\"name\":\"CooperativeWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"channelId\",\"type\":\"uint256\"}],\"name\":\"CooperativeSettle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"channelId\",\"type\":\"uint256\"}],\"name\":\"CooperativeSettleFail\",\"type\":\"event\"}]"

// GenericConditionalChannelBin is the compiled bytecode used for deploying new contracts.
const GenericConditionalChannelBin = `0x608060405234801561001057600080fd5b5060405160608062005d2d833981016040908152815160208301519190920151600092909255600160025560038054600160a060020a03928316600160a060020a03199182161790915560048054929093169116179055615cb680620000776000396000f3006080604052600436106101275763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166304f3bcec811461012c578063235edd9a1461015d578063298e7c69146101865780632a343709146102375780632d9df2411461024c578063417bac3f1461026457806342e556ce1461030757806358ffdbde1461033157806362771a771461040b5780636d80e4c2146104265780636e553f65146104c257806376aed2e1146104d95780638e7607b4146105155780639a8a05921461052a578063a07e51991461053f578063aa7c221414610557578063ab7a0a8014610572578063b194b0b7146105d2578063c3edbc041461066e578063d4c8464214610696578063de75a5c01461072e578063f53a728014610752575b600080fd5b34801561013857600080fd5b506101416107dc565b60408051600160a060020a039092168252519081900360200190f35b34801561016957600080fd5b50610184600435600160a060020a03602435166044356107eb565b005b34801561019257600080fd5b5061019e600435610a60565b604051808060200180602001838103835285818151815260200191508051906020019060200280838360005b838110156101e25781810151838201526020016101ca565b50505050905001838103825284818151815260200191508051906020019060200280838360005b83811015610221578181015183820152602001610209565b5050505090500194505050505060405180910390f35b34801561024357600080fd5b50610141610b83565b34801561025857600080fd5b50610184600435610b92565b34801561027057600080fd5b506040805160206004803580820135838102808601850190965280855261018495369593946024949385019291829185019084908082843750506040805187358901803560208181028481018201909552818452989b9a99890198929750908201955093508392508501908490808284375094975050843595505050506020820135600160a060020a031691604001359050610e7d565b34801561031357600080fd5b5061031f600435611157565b60408051918252519081900360200190f35b34801561033d57600080fd5b5060408051602060046024803582810135601f810185900485028601850190965285855261018495833595369560449491939091019190819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a99988101979196509182019450925082915084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a9998810197919650918201945092508291508401838280828437509497506111699650505050505050565b34801561041757600080fd5b50610184600435602435611610565b34801561043257600080fd5b5060408051602060046024803582810135601f810185900485028601850190965285855261018495833595369560449491939091019190819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a99988101979196509182019450925082915084018382808284375094975061170b9650505050505050565b610184600435600160a060020a03602435166119c5565b3480156104e557600080fd5b506104f1600435611bef565b6040518082600181111561050157fe5b60ff16815260200191505060405180910390f35b34801561052157600080fd5b5061031f611c0e565b34801561053657600080fd5b5061031f611c14565b34801561054b57600080fd5b50610141600435611c1a565b34801561056357600080fd5b50610184600435602435611c38565b34801561057e57600080fd5b50604080516020600460443581810135601f8101849004840285018401909552848452610184948235946024803595369594606494920191908190840183828082843750949750611dd49650505050505050565b3480156105de57600080fd5b5060408051602060046024803582810135601f810185900485028601850190965285855261018495833595369560449491939091019190819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a999881019791965091820194509250829150840183828082843750949750611e299650505050505050565b34801561067a57600080fd5b50610686600435612103565b6040518082600381111561050157fe5b3480156106a257600080fd5b506040805160206004602480358281013584810280870186019097528086526101849684359636966044959194909101929182918501908490808284375050604080516020601f89358b018035918201839004830284018301909452808352979a99988101979196509182019450925082915084018382808284375094975061211b9650505050505050565b34801561073a57600080fd5b5061031f600435600160a060020a03602435166122ad565b6040805160206004803580820135601f810184900484028501840190955284845261018494369492936024939284019190819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a9998810197919650918201945092508291508401838280828437509497506122d99650505050505050565b600354600160a060020a031681565b6000838152600160208190526040822060609183918791600282015460ff16600381111561081557fe5b1480610832575060028082015460ff16600381111561083057fe5b145b151561083d57600080fd5b60008881526001602081905260409091209550601586015460a060020a900460ff16600181111561086a57fe5b1461087457600080fd5b600160a060020a0387166000908152600b8601602052604090205460ff16151561089d57600080fd5b600160a060020a0387166000908152600c860160205260409020546108c8908763ffffffff61279d16565b600160a060020a0388166000908152600c8701602090815260409182902092909255600a870154815181815281840281019093019091528015610915578160200160208202803883390190505b50935061092288886127b0565b925085848481518110151561093357fe5b90602001906020020181815250507fe33829dd3495a41ba02f88de5eadec2635e25192c73dee8c65625206aa61e09a8886600a018660405180848152602001806020018060200183810383528581815481526020019150805480156109c157602002820191906000526020600020905b8154600160a060020a031681526001909101906020018083116109a3575b50508381038252845181528451602091820191808701910280838360005b838110156109f75781810151838201526020016109df565b505050509050019550505050505060405180910390a16001601586015460a060020a900460ff166001811115610a2957fe5b1415610a54576015850154610a4f90600160a060020a031633308963ffffffff61284516565b610a56565bfe5b5050505050505050565b6000818152600160209081526040808320600a81015482518181528185028101909401909252606093849391928492908015610aa6578160200160208202803883390190505b509150600090505b600a830154811015610b165782600c01600084600a0183815481101515610ad157fe5b6000918252602080832090910154600160a060020a031683528201929092526040019020548251839083908110610b0457fe5b60209081029091010152600101610aae565b82600a018281805480602002602001604051908101604052809291908181526020018280548015610b7057602002820191906000526020600020905b8154600160a060020a03168152600190910190602001808311610b52575b5050505050915094509450505050915091565b600454600160a060020a031681565b6000818152600160205260408120908060028084015460ff166003811115610bb657fe5b14610bc057600080fd5b82544311610bcd57600080fd5b610bd6836128f0565b1515610c14576040805185815290517f5a68b21ebf14d22f6a89bd2f92d1c899f2cfff1c821cdfbebd2af882389a287b9181900360200190a1610e77565b610c25838460050160000154612ce0565b1515610c63576040805185815290517f5a68b21ebf14d22f6a89bd2f92d1c899f2cfff1c821cdfbebd2af882389a287b9181900360200190a1610e77565b60028301805460ff19166003179055601583015460009060ff60a060020a909104166001811115610c9057fe5b1415610d7a576040805185815290517f05bfa6014fe9e35763ca08066a4f5a0e562a787ba9d77d27732eb4aa0b66b2369181900360200190a1600091505b600a830154821015610d7557600a8301805483908110610cea57fe5b6000918252602082200154600a85018054600160a060020a03909216926108fc92600d8801929087908110610d1b57fe5b6000918252602080832090910154600160a060020a03168352820192909252604090810182205490518115909302929091818181858888f19350505050158015610d69573d6000803e3d6000fd5b50600190910190610cce565b610e77565b6001601584015460a060020a900460ff166001811115610d9657fe5b1415610a54576040805185815290517f05bfa6014fe9e35763ca08066a4f5a0e562a787ba9d77d27732eb4aa0b66b2369181900360200190a150506015810154600090600160a060020a03165b600a830154821015610d7557610e6c83600a0183815481101515610e0357fe5b6000918252602082200154600a86018054600160a060020a0390921692600d88019290919087908110610e3257fe5b6000918252602080832090910154600160a060020a039081168452908301939093526040909101902054908416919063ffffffff612dd816565b600190910190610de3565b50505050565b845160009081901515610e8f57600080fd5b8551875114610e9d57600080fd5b600280546000908152600160208190526040909120808201889055918201805460ff191690911790559150610ed28385612e90565b60158301805474ff0000000000000000000000000000000000000000191660a060020a836001811115610f0157fe5b02179055505060158101805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03851617905560005b865181101561107b5781600b0160008883815181101515610f5157fe5b6020908102909101810151600160a060020a031682528101919091526040016000205460ff1615610f8157600080fd5b600182600b0160008984815181101515610f9757fe5b602090810291909101810151600160a060020a03168252810191909152604001600020805460ff19169115159190911790558651600a830190889083908110610fdc57fe5b602090810291909101810151825460018101845560009384529190922001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03909216919091179055855186908290811061103257fe5b90602001906020020151826010016000898481518110151561105057fe5b6020908102909101810151600160a060020a0316825281019190915260400160002055600101610f34565b7f9d51eb7f3cbaa052d1e423f1630d320f56f424e7c1fd7f4621654ccfd0d41b2f600254888460150160149054906101000a900460ff1660018111156110bd57fe5b601586015460408051858152908101839052600160a060020a039091166060820181905260806020808401828152865192850192909252855192939260a0840191878101910280838360005b83811015611121578181015183820152602001611109565b505050509050019550505050505060405180910390a160025461114b90600163ffffffff61279d16565b60025550505050505050565b60009081526001602052604090205490565b60008060006111766158f9565b61117e61592c565b61118661592c565b60008a8152600160208190526040822082918d9190600282015460ff1660038111156111ae57fe5b14806111cb575060028082015460ff1660038111156111c957fe5b145b15156111d657600080fd5b8c6040518082805190602001908083835b602083106112065780518252601f1990920191602091820191016111e7565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051809103902099508b6040518082805190602001908083835b602083106112655780518252601f199092019160209182019101611246565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390209850600160008f815260200190815260200160002097506112b28d612ef8565b80516000908152600e8a01602052604090205490975060ff16156112d557600080fd5b60608701518e146112e557600080fd5b60058801548751116112f657600080fd5b60408701511561130557600080fd5b6080870151431161131557600080fd5b61131e8c612f1d565b95506113298b612f1d565b9450611336888b88612f3a565b151561134157600080fd5b61134c888a87612f3a565b151561135757600080fd5b611361888861314a565b151561139f57604080518f815290517f9af0573e2d3394e13e9c62615248cd22581e1862e87a84ddb544247904fd9cb29181900360200190a1611600565b6113ad888860000151612ce0565b15156113eb57604080518f815290517f9af0573e2d3394e13e9c62615248cd22581e1862e87a84ddb544247904fd9cb29181900360200190a1611600565b60028801805460ff19166003179055601588015460009060ff60a060020a90910416600181111561141857fe5b141561150257604080518f815290517f6230d4127b5014fc7595da99cf59ce68bac5d9276e21712c85f91c06a96c722f9181900360200190a1600093505b600a8801548410156114fd57600a880180548590811061147257fe5b6000918252602082200154600a8a018054600160a060020a03909216926108fc92600d8d019290899081106114a357fe5b6000918252602080832090910154600160a060020a03168352820192909252604090810182205490518115909302929091818181858888f193505050501580156114f1573d6000803e3d6000fd5b50600190930192611456565b611600565b6001601589015460a060020a900460ff16600181111561151e57fe5b1415610a5457604080518f815290517f6230d4127b5014fc7595da99cf59ce68bac5d9276e21712c85f91c06a96c722f9181900360200190a1601588015460009450600160a060020a031692505b600a8801548410156114fd576115f588600a018581548110151561158c57fe5b6000918252602082200154600a8b018054600160a060020a0390921692600d8d0192909190899081106115bb57fe5b6000918252602080832090910154600160a060020a039081168452908301939093526040909101902054908616919063ffffffff612dd816565b60019093019261156c565b5050505050505050505050505050565b600082815260016020819052604082208291859190600282015460ff16600381111561163857fe5b1480611655575060028082015460ff16600381111561165357fe5b145b151561166057600080fd5b6000868152600160209081526040808320338452600b81019092529091205490945060ff16151561169057600080fd5b5050336000908152601083016020908152604080832054600f9095018252808320815160608101835296875243909501868301908152908601838152855460018082018855968552929093209551600390920290950190815593519284019290925550516002909101805460ff191691151591909117905550565b6000806117166158f9565b61171e61592c565b60008781526001602081905260409091208891600282015460ff16600381111561174457fe5b1480611761575060028082015460ff16600381111561175f57fe5b145b151561176c57600080fd5b876040518082805190602001908083835b6020831061179c5780518252601f19909201916020918201910161177d565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390209550600160008a815260200190815260200160002094506117e988612ef8565b80516000908152600e8701602052604090205490945060ff161561180c57600080fd5b6060840151891461181c57600080fd5b600585015484511161182d57600080fd5b61183687612f1d565b9250611843858785612f3a565b151561184e57600080fd5b6118578561335b565b61186085613461565b83516005860190815560208086015180518793926118859260068b019291019061594e565b50604082015160028281019190915560608301516003830155608090920151600490910155858101805460ff1916909117905560098501544311156118d2576001850154430185556118e0565b600185015460098601540185555b60068501805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815261197a938993919290918301828280156119705780601f1061194557610100808354040283529160200191611970565b820191906000526020600020905b81548152906001019060200180831161195357829003601f168201915b50505050506134c4565b6005850154604080518b8152602081019290925280517f5adbec4d8dfaaf515f95e278796bcb261b13c8df87a959188fd5f5fb05bccd029281900390910190a1505050505050505050565b6000828152600160208190526040822060609183918691600282015460ff1660038111156119ef57fe5b1480611a0c575060028082015460ff166003811115611a0a57fe5b145b1515611a1757600080fd5b6000878152600160209081526040808320600160a060020a038a168452600b81019092529091205490955060ff161515611a5057600080fd5b6000601586015460a060020a900460ff166001811115611a6c57fe5b14611a7657600080fd5b600160a060020a0386166000908152600c86016020526040902054611aa1903463ffffffff61279d16565b600160a060020a0387166000908152600c8701602090815260409182902092909255600a870154815181815281840281019093019091528015611aee578160200160208202803883390190505b509350611afb87876127b0565b9250348484815181101515611b0c57fe5b90602001906020020181815250507fe33829dd3495a41ba02f88de5eadec2635e25192c73dee8c65625206aa61e09a8786600a01866040518084815260200180602001806020018381038352858181548152602001915080548015611b9a57602002820191906000526020600020905b8154600160a060020a03168152600190910190602001808311611b7c575b50508381038252845181528451602091820191808701910280838360005b83811015611bd0578181015183820152602001611bb8565b505050509050019550505050505060405180910390a150505050505050565b60009081526001602052604090206015015460a060020a900460ff1690565b60025481565b60005481565b600090815260016020526040902060150154600160a060020a031690565b600082815260016020819052604082208291859190600282015460ff166003811115611c6057fe5b1480611c7d575060028082015460ff166003811115611c7b57fe5b145b1515611c8857600080fd5b6000868152600160209081526040808320338452600f810190925290912080549195509086908110611cb657fe5b906000526020600020906003020192508260010154600014158015611ce05750600283015460ff16155b1515611ceb57600080fd5b60018301544311611cfb57600080fd5b8254336000908152600c86016020526040902054611d1e9163ffffffff61358316565b336000908152600c86016020526040812091909155601585015460a060020a900460ff166001811115611d4d57fe5b1415611d86578254604051339180156108fc02916000818181858888f19350505050158015611d80573d6000803e3d6000fd5b50611dcc565b6001601585015460a060020a900460ff166001811115611da257fe5b1415610a545782546015850154611dcc91600160a060020a0390911690339063ffffffff612dd816565b505050505050565b60008381526001602081905260409091208491600282015460ff166003811115611dfa57fe5b1480611e17575060028082015460ff166003811115611e1557fe5b145b1515610a5457600080fd5b5050505050565b600080611e346159cc565b6000611e3e61592c565b600088815260016020819052604082208a91600282015460ff166003811115611e6357fe5b1480611e80575060028082015460ff166003811115611e7e57fe5b145b1515611e8b57600080fd5b896040518082805190602001908083835b60208310611ebb5780518252601f199092019160209182019101611e9c565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390209750600160008c81526020019081526020016000209650611f088a613595565b8051600090815260148901602052604090205490965060ff1615611f2b57600080fd5b60208601518b14611f3b57600080fd5b6060860151600160a060020a0381166000908152600b8901602052604090205490955060ff161515611f6c57600080fd5b600160a060020a0385163314611f8157600080fd5b611f8a89612f1d565b9350611f97878986612f3a565b1515611fa257600080fd5b855160009081526014880160209081526040808320805460ff1916600117905580890151600160a060020a0389168452600c8b0190925290912054909350611ff0908463ffffffff61358316565b600160a060020a0386166000818152600c8a01602090815260409182902084905581518f815290810187905280820192909252606082019290925290517f0375fe0230defc24bed394834804158bdbd5e92e4c3a6dc88f4db5439080f7b89181900360800190a16000601588015460a060020a900460ff16600181111561207357fe5b14156120b557604051600160a060020a0386169084156108fc029085906000818181858888f193505050501580156120af573d6000803e3d6000fd5b506120f6565b6001601588015460a060020a900460ff1660018111156120d157fe5b1415610a545760158701546120f690600160a060020a0316868563ffffffff612dd816565b5050505050505050505050565b60009081526001602052604090206002015460ff1690565b6000806121266159fe565b60008681526001602081905260409091208791600282015460ff16600381111561214c57fe5b1480612169575060028082015460ff16600381111561216757fe5b145b151561217457600080fd5b60008881526001602052604090208054909550431061219257600080fd5b856040518082805190602001908083835b602083106121c25780518252601f1990920191602091820191016121a3565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390209350612203878660050160020154866135b2565b5061220e8585613701565b61221786613755565b92508460010154438660000154031015612235576001850154430185555b6001836020015163ffffffff161415612292576122528584613772565b604080518981526020810186905281517fda7220c7681df4c0b2cb90420d72d108cc12db4cbed30a5c4f5cb59f04d3d49e929181900390910190a1610a56565b602083015163ffffffff161515610a545761225285846139da565b6000918252600160209081526040808420600160a060020a03939093168452600c909201905290205490565b6122e1615a2d565b6000806122ec61592c565b60606122f787613c98565b805180519196503391600090811061230b57fe5b60209081029091010151600160a060020a03161461232857600080fd5b60c0850151158061233d575060018560c00151145b151561234857600080fd5b60c0850151151561238057348560200151600081518110151561236757fe5b602090810290910101511461237b57600080fd5b61238b565b341561238b57600080fd5b6123ac8560000151866060015187608001518860a001518960c00151610e7d565b6001600060016002540381526020019081526020016000209350866040518082805190602001908083835b602083106123f65780518252601f1990920191602091820191016123d7565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040518091039020925061242e86612f1d565b915061243b848484612f3a565b151561244657600080fd5b60208501518051600090811061245857fe5b9060200190602002015184600c0160008760000151600081518110151561247b57fe5b90602001906020020151600160a060020a0316600160a060020a031681526020019081526020016000208190555083600a01805490506040519080825280602002602001820160405280156124da578160200160208202803883390190505b509050846020015160008151811015156124f057fe5b9060200190602002015181600081518110151561250957fe5b90602001906020020181815250507fe33829dd3495a41ba02f88de5eadec2635e25192c73dee8c65625206aa61e09a60016002540385600a0183604051808481526020018060200180602001838103835285818154815260200191508054801561259c57602002820191906000526020600020905b8154600160a060020a0316815260019091019060200180831161257e575b50508381038252845181528451602091820191808701910280838360005b838110156125d25781810151838201526020016125ba565b505050509050019550505050505060405180910390a1600480546002546040517f06fa40f30000000000000000000000000000000000000000000000000000000081526000199091016044820181905260609382019384528a5160648301528a51600160a060020a03909316936306fa40f3938c938c93929182916024810191608490910190602088019080838360005b8381101561267b578181015183820152602001612663565b50505050905090810190601f1680156126a85780820380516001836020036101000a031916815260200191505b50838103825285518152855160209182019187019080838360005b838110156126db5781810151838201526020016126c3565b50505050905090810190601f1680156127085780820380516001836020036101000a031916815260200191505b5095505050505050600060405180830381600087803b15801561272a57600080fd5b505af115801561273e573d6000803e3d6000fd5b506001925061274b915050565b8560c0015114156127945761279433308760200151600081518110151561276e57fe5b6020908102909101015160a0890151600160a060020a031692919063ffffffff61284516565b50505050505050565b818101828110156127aa57fe5b92915050565b6000828152600160209081526040808320600160a060020a0385168452600b8101909252822054829060ff1615156127e757600080fd5b5060005b600a820154811015610a545783600160a060020a031682600a018281548110151561281257fe5b600091825260209091200154600160a060020a031614156128355780925061283d565b6001016127eb565b505092915050565b604080517f23b872dd000000000000000000000000000000000000000000000000000000008152600160a060020a0385811660048301528481166024830152604482018490529151918616916323b872dd916064808201926020929091908290030181600087803b1580156128b957600080fd5b505af11580156128cd573d6000803e3d6000fd5b505050506040513d60208110156128e357600080fd5b50511515610e7757600080fd5b600080805b600a8401548210156129885783600c01600085600a018481548110151561291857fe5b6000918252602080832090910154600160a060020a03168352820192909252604001812054600a860180549192600d8801929091908690811061295757fe5b6000918252602080832090910154600160a060020a03168352820192909252604001902055600191909101906128f5565b600091505b600a840154821015612ab9575060005b600a840154811015612aae57612a6a84601101600086600a01858154811015156129c357fe5b6000918252602080832090910154600160a060020a031683528201929092526040018120600a87018054919291859081106129fa57fe5b6000918252602080832090910154600160a060020a03168352820192909252604001812054600a870180549192600d89019290919086908110612a3957fe5b6000918252602080832090910154600160a060020a031683528201929092526040019020549063ffffffff61279d16565b84600d01600086600a0184815481101515612a8157fe5b6000918252602080832090910154600160a060020a0316835282019290925260400190205560010161299d565b60019091019061298d565b600091505b600a840154821015612cd4575060005b600a840154811015612cc95783601101600085600a0184815481101515612af157fe5b6000918252602080832090910154600160a060020a031683528201929092526040018120600a8601805491929184908110612b2857fe5b6000918252602080832090910154600160a060020a03168352820192909252604001812054600a860180549192600d88019290919086908110612b6757fe5b6000918252602080832090910154600160a060020a031683528201929092526040019020541015612bc45760058401546000908152600e850160205260409020805460ff19166001179055612bbb84613cb5565b60009250612cd9565b612c8584601101600086600a0185815481101515612bde57fe5b6000918252602080832090910154600160a060020a031683528201929092526040018120600a8701805491929185908110612c1557fe5b6000918252602080832090910154600160a060020a03168352820192909252604001812054600a870180549192600d89019290919087908110612c5457fe5b6000918252602080832090910154600160a060020a031683528201929092526040019020549063ffffffff61358316565b84600d01600086600a0185815481101515612c9c57fe5b6000918252602080832090910154600160a060020a03168352820192909252604001902055600101612ace565b600190910190612abe565b600192505b5050919050565b60008080805b600a860154811015612d9657612d3e86600d01600088600a0184815481101515612d0c57fe5b6000918252602080832090910154600160a060020a03168352820192909252604001902054849063ffffffff61279d16565b9250612d8c86600c01600088600a0184815481101515612d5a57fe5b6000918252602080832090910154600160a060020a03168352820192909252604001902054839063ffffffff61279d16565b9150600101612ce6565b828214612dca576000858152600e870160205260409020805460ff19166001179055612dc186613cb5565b60009350612dcf565b600193505b50505092915050565b82600160a060020a031663a9059cbb83836040518363ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018083600160a060020a0316600160a060020a0316815260200182815260200192505050602060405180830381600087803b158015612e5457600080fd5b505af1158015612e68573d6000803e3d6000fd5b505050506040513d6020811015612e7e57600080fd5b50511515612e8b57600080fd5b505050565b6000821515612eb557600160a060020a03821615612ead57600080fd5b5060006127aa565b6001831415610a5457600160a060020a0382161515612ed357600080fd5b612ee582600160a060020a0316613d12565b1515612ef057600080fd5b5060016127aa565b612f006158f9565b612f086158f9565b612f156020848551613d1a565b509392505050565b612f2561592c565b612f2d61592c565b612f156020848551613e9e565b805151600a84015460009182918291829114612f5557600080fd5b612f5e8761404c565b604080517f19457468657265756d205369676e6564204d6573736167653a0a3332000000008152601c8101889052905190819003603c0190209250600091505b8451518210156131235784518051600191859185908110612fbb57fe5b90602001906020020151876020015185815181101515612fd757fe5b90602001906020020151886040015186815181101515612ff357fe5b60209081029091018101516040805160008082528185018084529790975260ff9095168582015260608501939093526080840152905160a0808401949293601f19830193908390039091019190865af1158015613054573d6000803e3d6000fd5b505060408051601f190151600160a060020a0381166000908152600b8b01602052919091205490925060ff16905080156130a95750600160a060020a038116600090815260128801602052604090205460ff16155b1561310f576003870180546001808201835560009283526020808420909201805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a038616908117909155835260128a019091526040909120805460ff19169091179055613118565b60009350613140565b600190910190612f9e565b600a87015460038801541461313b5760009350613140565b600193505b5050509392505050565b6000613154615a7d565b600061315e615a8f565b61316b85602001516140b4565b9250600091505b600a8601548210156132055785600c01600087600a018481548110151561319557fe5b6000918252602080832090910154600160a060020a03168352820192909252604001812054600a880180549192600d8a0192909190869081106131d457fe5b6000918252602080832090910154600160a060020a0316835282019290925260400190205560019190910190613172565b600091505b82515182101561328c57825180518390811061322257fe5b602090810290910181015160408082015182840151600160a060020a03166000908152600d8b019094529220549092506132619163ffffffff61279d16565b602080830151600160a060020a03166000908152600d8901909152604090205560019091019061320a565b600091505b825151821015612dca5782518051839081106132a957fe5b60209081029091018101516040808201518251600160a060020a03166000908152600d8b0190945292205490925010156133035784516000908152600e870160205260409020805460ff19166001179055612dc186613cb5565b6040808201518251600160a060020a03166000908152600d890160205291909120546133349163ffffffff61358316565b8151600160a060020a03166000908152600d88016020526040902055600190910190613291565b6000805b600a830154821015613401575060005b600a8301548110156133f65782601101600084600a018481548110151561339257fe5b6000918252602080832090910154600160a060020a031683528201929092526040018120600a85018054919291849081106133c957fe5b6000918252602080832090910154600160a060020a0316835282019290925260400181205560010161336f565b60019091019061335f565b600091505b600a830154821015612e8b5782600f01600084600a018481548110151561342957fe5b6000918252602080832090910154600160a060020a03168352820192909252604001812061345691615aaf565b600190910190613406565b60005b60048201548110156134b257816013016000836004018381548110151561348757fe5b600091825260208083209091015483528201929092526040019020805460ff19169055600101613464565b6134c0600483016000615ad3565b5050565b6134cc615a7d565b60006134d6615a8f565b6134df846140b4565b9250600091505b825151821015611e225782518051839081106134fe57fe5b60209081029091018101516040808201518251600160a060020a03908116600090815260118b018652838120858701519092168152945292205490925061354a9163ffffffff61279d16565b8151600160a060020a039081166000908152601188016020908152604080832082870151909416835292905220556001909101906134e6565b60008282111561358f57fe5b50900390565b61359d6159cc565b6135a56159cc565b612f1560208485516140d1565b60008181805b86518210156136f45786828151811015156135cf57fe5b6020908102909101015190508083101561366857604080516020808201869052818301849052825180830384018152606090920192839052815191929182918401908083835b602083106136345780518252601f199092019160209182019101613615565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051809103902092506136e9565b604080516020808201849052818301869052825180830384018152606090920192839052815191929182918401908083835b602083106136b95780518252601f19909201916020918201910161369a565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051809103902092505b6001909101906135b8565b5050929092149392505050565b600081815260138301602052604090205460ff161561371f57600080fd5b600081815260138301602090815260408220805460ff191660019081179091556004909401805494850181558252902090910155565b61375d6159fe565b6137656159fe565b612f156020848551614213565b600061377c615af1565b600080600093505b8451518410156139cc57845180518590811061379c57fe5b9060200190602002015192506137b1836143e0565b60a0840151602080860151604080517fd2121c2b00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff8316602482015260048101918252845160448201528451959750879650600160a060020a0387169563d2121c2b95948392606401919086019080838360005b8381101561384357818101518382015260200161382b565b50505050905090810190601f1680156138705780820380516001836020036101000a031916815260200191505b509350505050602060405180830381600087803b15801561389057600080fd5b505af11580156138a4573d6000803e3d6000fd5b505050506040513d60208110156138ba57600080fd5b505180156139b6575060c08301516040517fd6f12ca3000000000000000000000000000000000000000000000000000000008152602060048201818152835160248401528351600160a060020a0386169463d6f12ca394909383926044909201919085019080838360005b8381101561393d578181015183820152602001613925565b50505050905090810190601f16801561396a5780820380516001836020036101000a031916815260200191505b5092505050602060405180830381600087803b15801561398957600080fd5b505af115801561399d573d6000803e3d6000fd5b505050506040513d60208110156139b357600080fd5b50515b15156139c157600080fd5b600190930192613784565b611dcc8686606001516134c4565b60006139e4615af1565b6000806060600094505b855151851015612794578551805186908110613a0657fe5b906020019060200201519350613a1b846143e0565b60a0850151602080870151604080517fd2121c2b00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff8316602482015260048101918252845160448201528451959850889750600160a060020a0388169563d2121c2b95948392606401919086019080838360005b83811015613aad578181015183820152602001613a95565b50505050905090810190601f168015613ada5780820380516001836020036101000a031916815260200191505b509350505050602060405180830381600087803b158015613afa57600080fd5b505af1158015613b0e573d6000803e3d6000fd5b505050506040513d6020811015613b2457600080fd5b50511515613b3157600080fd5b60c08401516040517f195a4bc6000000000000000000000000000000000000000000000000000000008152602060048201818152835160248401528351600160a060020a0387169463195a4bc694909383926044909201919085019080838360005b83811015613bab578181015183820152602001613b93565b50505050905090810190601f168015613bd85780820380516001836020036101000a031916815260200191505b5092505050600060405180830381600087803b158015613bf757600080fd5b505af1158015613c0b573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526020811015613c3457600080fd5b810190808051640100000000811115613c4c57600080fd5b82016020810184811115613c5f57600080fd5b8151640100000000811182820187101715613c7957600080fd5b50509291905050509050613c8d87826134c4565b6001909401936139ee565b613ca0615a2d565b613ca8615a2d565b612f1560208485516144cd565b60006005820181815590613ccc6006840182615b2f565b506000600282018190556003820181905560049091018190558155613cf08161404c565b613cf98161335b565b613d0281613461565b600201805460ff19166001179055565b6000903b1190565b613d226158f9565b6000613d2c6158f9565b613d34615b73565b60008080895b8881018b1015613dd057613d4e8b8b614739565b9c8d019c919550935091508360011415613d7757613d6e8b8b8888614776565b8b019a50613dcb565b8360021415613d8c57613d6e8b8b88886147d8565b8360031415613da157613d6e8b8b8888614826565b8360041415613db657613d6e8b8b8888614873565b836005141561012757613d6e8b8b88886148c0565b613d3a565b809a505b8881018b1015613e8d57613de88b8b614739565b9c8d019c919550935091508360011415613e1857613e0f8b8b613e0961490d565b88614776565b8b019a50613e88565b8360021415613e3457613e0f8b8b613e2e61490d565b886147d8565b8360031415613e5057613e0f8b8b613e4a61490d565b88614826565b8360041415613e6c57613e0f8b8b613e6661490d565b88614873565b836005141561012757613e0f8b8b613e8261490d565b886148c0565b613dd4565b509399969850959650505050505050565b613ea661592c565b6000613eb061592c565b613eb8615b92565b60008080895b8881018b1015613f3f57613ed28b8b614739565b9c8d019c919550935091508360011415613f0257613ef98b8b613ef361491b565b88614923565b8b019a50613f3a565b8360021415613f1e57613ef98b8b613f1861491b565b8861498f565b836003141561012757613ef98b8b613f3461491b565b886149f8565b613ebe565b9950898460016020020151604051908082528060200260200182016040528015613f73578160200160208202803883390190505b508652604080860151815181815260208281028201019092528015613fa2578160200160208202803883390190505b5060208701528460036020020151604051908082528060200260200182016040528015613fd9578160200160208202803883390190505b5060408701525b8881018b1015613e8d57613ff48b8b614739565b9c8d019c91955093509150836001141561401d576140148b8b8888614923565b8b019a50614047565b8360021415614032576140148b8b888861498f565b8360031415610127576140148b8b88886149f8565b613fe0565b60005b60038201548110156140a657816012016000836003018381548110151561407257fe5b6000918252602080832090910154600160a060020a031683528201929092526040019020805460ff1916905560010161404f565b6134c0600383016000615ad3565b6140bc615a7d565b6140c4615a7d565b612f156020848551614a61565b6140d96159cc565b60006140e36159cc565b6140eb615bb1565b60008080895b8881018b1015614172576141058b8b614739565b9c8d019c91955093509150836001141561412e576141258b8b8888614b4a565b8b019a5061416d565b8360021415614143576141258b8b8888614b8d565b8360031415614158576141258b8b8888614b9c565b8360041415610127576141258b8b8888614bab565b6140f1565b809a505b8881018b1015613e8d5761418a8b8b614739565b9c8d019c9195509350915083600114156141ba576141b18b8b6141ab614bfc565b88614b4a565b8b019a5061420e565b83600214156141d6576141b18b8b6141d0614bfc565b88614b8d565b83600314156141f2576141b18b8b6141ec614bfc565b88614b9c565b8360041415610127576141b18b8b614208614bfc565b88614bab565b614176565b61421b6159fe565b60006142256159fe565b61422d615bb1565b60008080895b8881018b10156142c2576142478b8b614739565b9c8d019c9195509350915083600114156142775761426e8b8b614268614c04565b88614c0c565b8b019a506142bd565b836002141561428c5761426e8b8b8888614c68565b83600314156142a85761426e8b8b6142a2614c04565b88614ca9565b83600414156101275761426e8b8b8888614d09565b614233565b995089846001602002015160405190808252806020026020018201604052801561430657816020015b6142f3615af1565b8152602001906001900390816142eb5790505b508652846003602002015160405190808252806020026020018201604052801561434a57816020015b614337615bd0565b81526020019060019003908161432f5790505b5060408701525b8881018b1015613e8d576143658b8b614739565b9c8d019c91955093509150836001141561438e576143858b8b8888614c0c565b8b019a506143db565b83600214156143aa576143858b8b6143a4614c04565b88614c68565b83600314156143bf576143858b8b8888614ca9565b8360041415610127576143858b8b6143d5614c04565b88614d09565b614351565b600043826020015167ffffffffffffffff161015156143fe57600080fd5b608082015163ffffffff1615156144ae576003546060830151604080517f5c23bdf5000000000000000000000000000000000000000000000000000000008152600481019290925251600160a060020a0390921691635c23bdf5916024808201926020929091908290030181600087803b15801561447b57600080fd5b505af115801561448f573d6000803e3d6000fd5b505050506040513d60208110156144a557600080fd5b505190506144c8565b6001826080015163ffffffff161415610a54575060608101515b919050565b6144d5615a2d565b60006144df615a2d565b6144e7615be7565b60008080895b8881018b10156145c2576145018b8b614739565b9c8d019c919550935091508360011415614531576145288b8b614522614d46565b88614d4e565b8b019a506145bd565b836002141561454d576145288b8b614547614d46565b88614dbb565b8360031415614562576145288b8b8888614dca565b836004141561457e576145288b8b614578614d46565b88614e1b565b8360051415614593576145288b8b88886148c0565b83600614156145a8576145288b8b8888614e84565b8360071415610127576145288b8b8888614ed9565b6144ed565b99508984600160200201516040519080825280602002602001820160405280156145f6578160200160208202803883390190505b508652604080860151815181815260208281028201019092528015614625578160200160208202803883390190505b506020870152846004602002015160405190808252806020026020018201604052801561465c578160200160208202803883390190505b5060608701525b8881018b1015613e8d576146778b8b614739565b9c8d019c9195509350915083600114156146a0576146978b8b8888614d4e565b8b019a50614734565b83600214156146b5576146978b8b8888614dbb565b83600314156146d1576146978b8b6146cb614d46565b88614dca565b83600414156146e6576146978b8b8888614e1b565b83600514156146fc576146978b8b613e82614d46565b8360061415614718576146978b8b614712614d46565b88614e84565b8360071415610127576146978b8b61472e614d46565b88614ed9565b614663565b600080600080600080600061474e8989614f26565b9350935083600716600581111561476157fe5b91506008840499919850919650945050505050565b60008060006147858787614f6b565b9150915061479285614f8a565b156147ac57600184815b60200201805190910190526147ce565b8185526020840151600010156147ce57600184815b6020020180519190910390525b9695505050505050565b6000606060006147e88787614f8e565b915091506147f585614f8a565b1561480457600184600261479c565b6020850182905260008460025b602002015111156147ce5760018460026147c1565b60008060006148358787614fa7565b9150915061484285614f8a565b1561485157600184600361479c565b6040850182905260008460035b602002015111156147ce5760018460036147c1565b60008060006148828787614f6b565b9150915061488f85614f8a565b1561489e57600184600461479c565b6060850182905260008460045b602002015111156147ce5760018460046147c1565b60008060006148cf8787614f6b565b915091506148dc85614f8a565b156148eb57600184600561479c565b6080850182905260008460055b602002015111156147ce5760018460056147c1565b6149156158f9565b50600090565b61491561592c565b60008060006149328787614fb6565b9150915061493f85614f8a565b1561494d576001848161479c565b845160208501518151849291810390811061496457fe5b60ff90921660209283029091019091015260008460015b602002015111156147ce57600184816147c1565b600080600061499e8787614fa7565b915091506149ab85614f8a565b156149ba57600184600261479c565b60208501516040850151815184929181039081106149d457fe5b602090810290910101526000846002602002015111156147ce5760018460026147c1565b6000806000614a078787614fa7565b91509150614a1485614f8a565b15614a2357600184600361479c565b6040850151606085015181518492918103908110614a3d57fe5b602090810290910101526000846003602002015111156147ce5760018460036147c1565b614a69615a7d565b6000614a73615a7d565b614a7b615c07565b60008080895b8881018b1015614ac557614a958b8b614739565b9c8d019c91955093509150836001141561012757614abc8b8b614ab6614fc8565b88614fd0565b8b019a50614a81565b9950898460016020020151604051908082528060200260200182016040528015614b0957816020015b614af6615a8f565b815260200190600190039081614aee5790505b5086525b8881018b1015613e8d57614b218b8b614739565b9c8d019c91955093509150836001141561012757614b418b8b8888614fd0565b8b019a50614b0d565b6000806000614b598787614f6b565b91509150614b6685614f8a565b15614b74576001848161479c565b8185526020840151600010156147ce57600184816147c1565b60008060006147e88787614f6b565b60008060006148358787614f6b565b6000806000614bba878761503b565b91509150614bc785614f8a565b15614bd657600184600461479c565b600160a060020a03821660608601526080840151600010156147ce5760018460046147c1565b6149156159cc565b6149156159fe565b6000614c16615af1565b6000614c22878761504d565b91509150614c2f85614f8a565b15614c3d576001848161479c565b8451602085015181518492918103908110614c5457fe5b60209081029091010152600084600161497b565b6000806000614c77878761508f565b91509150614c8485614f8a565b15614c9357600184600261479c565b63ffffffff821660208601526000846002614811565b6000614cb3615bd0565b6000614cbf878761509f565b91509150614ccc85614f8a565b15614cdb57600184600361479c565b6040850151606085015181518492918103908110614cf557fe5b60209081029091010152600084600361485e565b600060606000614d198787614f8e565b91509150614d2685614f8a565b15614d3557600184600461479c565b6060850182905260008460046148ab565b614915615a2d565b6000806000614d5d878761503b565b91509150614d6a85614f8a565b15614d78576001848161479c565b8451602085015181518492918103908110614d8f57fe5b600160a060020a0392909216602092830291909101820152840151600010156147ce57600184816147c1565b600080600061499e8787614f6b565b6000806000614dd9878761503b565b91509150614de685614f8a565b15614df557600184600361479c565b600160a060020a03821660408601526060840151600010156147ce5760018460036147c1565b6000806000614e2a8787614f6b565b91509150614e3785614f8a565b15614e4657600184600461479c565b6060850151608085015181518492918103908110614e6057fe5b602090810290910101526000846004602002015111156147ce5760018460046147c1565b6000806000614e93878761503b565b91509150614ea085614f8a565b15614eaf57600184600661479c565b600160a060020a03821660a086015260008460065b602002015111156147ce5760018460066147c1565b6000806000614ee88787614f6b565b91509150614ef585614f8a565b15614f0457600184600761479c565b60c0850182905260008460075b602002015111156147ce5760018460076147c1565b908101906000808080805b865160001a90508160070260020a81607f16028317925060018201915060018701965060808116608014614f315750909590945092505050565b600080600080614f7d602087876150d2565b9097909650945050505050565b1590565b60606000614f9c848461511c565b915091509250929050565b600080614f9c6020858561519f565b600080600080614f7d600187876150d2565b614915615a7d565b6000614fda615a8f565b6000614fe687876151e4565b91509150614ff385614f8a565b15615001576001848161479c565b845160208501518151849291810390811061501857fe5b602090810290910101526000846001602002015111156147ce57600184816147c1565b600080600080614f7d60148787615217565b615055615af1565b6000806000615062615af1565b61506c8787614f26565b925092508187019650615080878785615284565b50979190920195509350505050565b600080600080614f7d8686614f26565b6150a7615bd0565b60008060006150b4615bd0565b6150be8787614f26565b925092508187019650615080878785615459565b60008060008060006150e48787614f26565b91509150600387019650858701519250600282036020036101000a8381151561510957fe5b0498600190920197509095505050505050565b60606000806000606061512f8787614f26565b92509250826040519080825280601f01601f191660200182016040528015615161578160200160208202803883390190505b5096860182019690506020810160005b84811461518f57885160001a82536001988901989182019101615171565b5090979190920195509350505050565b60008060008060006151b18787614f26565b909350915082820160ff60038a0116146151ca57600080fd5b505050600393909201830151949390920160ff1692915050565b6151ec615a8f565b60008060006151f9615a8f565b6152038787614f26565b925092508187019650615080878785615533565b60008060008060008061522a8888614f26565b909350915082820160ff60038b01161461524357600080fd5b6003880197508688015193508860ff1690505b602081101561526e5761010084049350600101615256565b50919760039790970160ff169695505050505050565b61528c615af1565b6000615296615af1565b61529e615be7565b60008080895b8881018b1015615364576152b88b8b614739565b9c8d019c9195509350915083600114156152e1576152d88b8b8888615644565b8b019a5061535f565b83600214156152f6576152d88b8b8888615685565b836003141561530b576152d88b8b88886156ca565b8360041415615320576152d88b8b888861570b565b8360051415615335576152d88b8b888861571a565b836006141561534a576152d88b8b888861575b565b8360071415610127576152d88b8b8888615798565b6152a4565b809a505b8881018b1015613e8d5761537c8b8b614739565b9c8d019c9195509350915083600114156153ac576153a38b8b61539d6157d5565b88615644565b8b019a50615454565b83600214156153c8576153a38b8b6153c26157d5565b88615685565b83600314156153e4576153a38b8b6153de6157d5565b886156ca565b8360041415615400576153a38b8b6153fa6157d5565b8861570b565b836005141561541c576153a38b8b6154166157d5565b8861571a565b8360061415615438576153a38b8b6154326157d5565b8861575b565b8360071415610127576153a38b8b61544e6157d5565b88615798565b615368565b615461615bd0565b600061546b615bd0565b615473615c22565b60008080895b8881018b10156154d05761548d8b8b614739565b9c8d019c9195509350915083600114156154b6576154ad8b8b88886157dd565b8b019a506154cb565b8360021415610127576154ad8b8b8888614b8d565b615479565b809a505b8881018b1015613e8d576154e88b8b614739565b9c8d019c9195509350915083600114156155185761550f8b8b615509615829565b886157dd565b8b019a5061552e565b83600214156101275761550f8b8b6141d0615829565b6154d4565b61553b615a8f565b6000615545615a8f565b61554d615b92565b60008080895b8881018b10156155bf576155678b8b614739565b9c8d019c919550935091508360011415615590576155878b8b8888615831565b8b019a506155ba565b83600214156155a5576155878b8b8888615871565b8360031415610127576155878b8b88886158b5565b615553565b809a505b8881018b1015613e8d576155d78b8b614739565b9c8d019c919550935091508360011415615607576155fe8b8b6155f86158f1565b88615831565b8b019a5061563f565b8360021415615623576155fe8b8b61561d6158f1565b88615871565b8360031415610127576155fe8b8b6156396158f1565b886158b5565b6155c3565b6000806000615653878761508f565b9150915061566085614f8a565b1561566e576001848161479c565b67ffffffffffffffff82168552600084600161497b565b6000806000615694878761508f565b915091506156a185614f8a565b156156b057600184600261479c565b67ffffffffffffffff821660208601526000846002614811565b60008060006156d9878761508f565b915091506156e685614f8a565b156156f557600184600361479c565b63ffffffff82166040860152600084600361485e565b6000806000614d198787614fa7565b6000806000615729878761508f565b9150915061573685614f8a565b1561574557600184600561479c565b63ffffffff8216608086015260008460056148f8565b60006060600061576b8787614f8e565b9150915061577885614f8a565b1561578757600184600661479c565b60a085018290526000846006614ec4565b6000606060006157a88787614f8e565b915091506157b585614f8a565b156157c457600184600761479c565b60c085018290526000846007614f11565b614915615af1565b60008060006157ec878761503b565b915091506157f985614f8a565b15615807576001848161479c565b600160a060020a03821685526020840151600010156147ce57600184816147c1565b614915615bd0565b6000806000615840878761503b565b9150915061584d85614f8a565b1561585b576001848161479c565b600160a060020a0382168552600084600161497b565b6000806000615880878761503b565b9150915061588d85614f8a565b1561589c57600184600261479c565b600160a060020a03821660208601526000846002614811565b60008060006158c48787614f6b565b915091506158d185614f8a565b156158e057600184600361479c565b60408501829052600084600361485e565b614915615a8f565b60a06040519081016040528060008152602001606081526020016000801916815260200160008152602001600081525090565b6060604051908101604052806060815260200160608152602001606081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061598f57805160ff19168380011785556159bc565b828001600101855582156159bc579182015b828111156159bc5782518255916020019190600101906159a1565b506159c8929150615c41565b5090565b6080604051908101604052806000815260200160008152602001600081526020016000600160a060020a031681525090565b60806040519081016040528060608152602001600063ffffffff16815260200160608152602001606081525090565b60e06040519081016040528060608152602001606081526020016000600160a060020a0316815260200160608152602001600081526020016000600160a060020a03168152602001600081525090565b60408051602081019091526060815290565b604080516060810182526000808252602082018190529181019190915290565b5080546000825560030290600052602060002090810190615ad09190615c5e565b50565b5080546000825590600052602060002090810190615ad09190615c41565b6040805160e0810182526000808252602082018190529181018290526060808201839052608082019290925260a0810182905260c081019190915290565b50805460018160011615610100020316600290046000825580601f10615b555750615ad0565b601f016020900490600052602060002090810190615ad09190615c41565b60c0604051908101604052806006906020820280388339509192915050565b6080604051908101604052806004906020820280388339509192915050565b60a0604051908101604052806005906020820280388339509192915050565b604080518082019091526000808252602082015290565b610100604051908101604052806008906020820280388339509192915050565b60408051808201825290600290829080388339509192915050565b6060604051908101604052806003906020820280388339509192915050565b615c5b91905b808211156159c85760008155600101615c47565b90565b615c5b91905b808211156159c8576000808255600182015560028101805460ff19169055600301615c645600a165627a7a723058206b1e4157bcc841017ebefdab2ac91f37460c6b29f7108c7b76bcfdcaec3c301c0029`

// DeployGenericConditionalChannel deploys a new Ethereum contract, binding an instance of GenericConditionalChannel to it.
func DeployGenericConditionalChannel(auth *bind.TransactOpts, backend bind.ContractBackend, _chainId *big.Int, _virtResolver common.Address, _depositPool common.Address) (common.Address, *types.Transaction, *GenericConditionalChannel, error) {
	parsed, err := abi.JSON(strings.NewReader(GenericConditionalChannelABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(GenericConditionalChannelBin), backend, _chainId, _virtResolver, _depositPool)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GenericConditionalChannel{GenericConditionalChannelCaller: GenericConditionalChannelCaller{contract: contract}, GenericConditionalChannelTransactor: GenericConditionalChannelTransactor{contract: contract}, GenericConditionalChannelFilterer: GenericConditionalChannelFilterer{contract: contract}}, nil
}

// GenericConditionalChannel is an auto generated Go binding around an Ethereum contract.
type GenericConditionalChannel struct {
	GenericConditionalChannelCaller     // Read-only binding to the contract
	GenericConditionalChannelTransactor // Write-only binding to the contract
	GenericConditionalChannelFilterer   // Log filterer for contract events
}

// GenericConditionalChannelCaller is an auto generated read-only Go binding around an Ethereum contract.
type GenericConditionalChannelCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GenericConditionalChannelTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GenericConditionalChannelTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GenericConditionalChannelFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GenericConditionalChannelFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GenericConditionalChannelSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GenericConditionalChannelSession struct {
	Contract     *GenericConditionalChannel // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// GenericConditionalChannelCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GenericConditionalChannelCallerSession struct {
	Contract *GenericConditionalChannelCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// GenericConditionalChannelTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GenericConditionalChannelTransactorSession struct {
	Contract     *GenericConditionalChannelTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// GenericConditionalChannelRaw is an auto generated low-level Go binding around an Ethereum contract.
type GenericConditionalChannelRaw struct {
	Contract *GenericConditionalChannel // Generic contract binding to access the raw methods on
}

// GenericConditionalChannelCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GenericConditionalChannelCallerRaw struct {
	Contract *GenericConditionalChannelCaller // Generic read-only contract binding to access the raw methods on
}

// GenericConditionalChannelTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GenericConditionalChannelTransactorRaw struct {
	Contract *GenericConditionalChannelTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGenericConditionalChannel creates a new instance of GenericConditionalChannel, bound to a specific deployed contract.
func NewGenericConditionalChannel(address common.Address, backend bind.ContractBackend) (*GenericConditionalChannel, error) {
	contract, err := bindGenericConditionalChannel(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GenericConditionalChannel{GenericConditionalChannelCaller: GenericConditionalChannelCaller{contract: contract}, GenericConditionalChannelTransactor: GenericConditionalChannelTransactor{contract: contract}, GenericConditionalChannelFilterer: GenericConditionalChannelFilterer{contract: contract}}, nil
}

// NewGenericConditionalChannelCaller creates a new read-only instance of GenericConditionalChannel, bound to a specific deployed contract.
func NewGenericConditionalChannelCaller(address common.Address, caller bind.ContractCaller) (*GenericConditionalChannelCaller, error) {
	contract, err := bindGenericConditionalChannel(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GenericConditionalChannelCaller{contract: contract}, nil
}

// NewGenericConditionalChannelTransactor creates a new write-only instance of GenericConditionalChannel, bound to a specific deployed contract.
func NewGenericConditionalChannelTransactor(address common.Address, transactor bind.ContractTransactor) (*GenericConditionalChannelTransactor, error) {
	contract, err := bindGenericConditionalChannel(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GenericConditionalChannelTransactor{contract: contract}, nil
}

// NewGenericConditionalChannelFilterer creates a new log filterer instance of GenericConditionalChannel, bound to a specific deployed contract.
func NewGenericConditionalChannelFilterer(address common.Address, filterer bind.ContractFilterer) (*GenericConditionalChannelFilterer, error) {
	contract, err := bindGenericConditionalChannel(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GenericConditionalChannelFilterer{contract: contract}, nil
}

// bindGenericConditionalChannel binds a generic wrapper to an already deployed contract.
func bindGenericConditionalChannel(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GenericConditionalChannelABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GenericConditionalChannel *GenericConditionalChannelRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GenericConditionalChannel.Contract.GenericConditionalChannelCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GenericConditionalChannel *GenericConditionalChannelRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GenericConditionalChannel.Contract.GenericConditionalChannelTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GenericConditionalChannel *GenericConditionalChannelRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GenericConditionalChannel.Contract.GenericConditionalChannelTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GenericConditionalChannel *GenericConditionalChannelCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GenericConditionalChannel.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GenericConditionalChannel *GenericConditionalChannelTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GenericConditionalChannel.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GenericConditionalChannel *GenericConditionalChannelTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GenericConditionalChannel.Contract.contract.Transact(opts, method, params...)
}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() constant returns(uint256)
func (_GenericConditionalChannel *GenericConditionalChannelCaller) ChainId(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _GenericConditionalChannel.contract.Call(opts, out, "chainId")
	return *ret0, err
}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() constant returns(uint256)
func (_GenericConditionalChannel *GenericConditionalChannelSession) ChainId() (*big.Int, error) {
	return _GenericConditionalChannel.Contract.ChainId(&_GenericConditionalChannel.CallOpts)
}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() constant returns(uint256)
func (_GenericConditionalChannel *GenericConditionalChannelCallerSession) ChainId() (*big.Int, error) {
	return _GenericConditionalChannel.Contract.ChainId(&_GenericConditionalChannel.CallOpts)
}

// ChannelLength is a free data retrieval call binding the contract method 0x8e7607b4.
//
// Solidity: function channelLength() constant returns(uint256)
func (_GenericConditionalChannel *GenericConditionalChannelCaller) ChannelLength(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _GenericConditionalChannel.contract.Call(opts, out, "channelLength")
	return *ret0, err
}

// ChannelLength is a free data retrieval call binding the contract method 0x8e7607b4.
//
// Solidity: function channelLength() constant returns(uint256)
func (_GenericConditionalChannel *GenericConditionalChannelSession) ChannelLength() (*big.Int, error) {
	return _GenericConditionalChannel.Contract.ChannelLength(&_GenericConditionalChannel.CallOpts)
}

// ChannelLength is a free data retrieval call binding the contract method 0x8e7607b4.
//
// Solidity: function channelLength() constant returns(uint256)
func (_GenericConditionalChannel *GenericConditionalChannelCallerSession) ChannelLength() (*big.Int, error) {
	return _GenericConditionalChannel.Contract.ChannelLength(&_GenericConditionalChannel.CallOpts)
}

// DepositPool is a free data retrieval call binding the contract method 0x2a343709.
//
// Solidity: function depositPool() constant returns(address)
func (_GenericConditionalChannel *GenericConditionalChannelCaller) DepositPool(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _GenericConditionalChannel.contract.Call(opts, out, "depositPool")
	return *ret0, err
}

// DepositPool is a free data retrieval call binding the contract method 0x2a343709.
//
// Solidity: function depositPool() constant returns(address)
func (_GenericConditionalChannel *GenericConditionalChannelSession) DepositPool() (common.Address, error) {
	return _GenericConditionalChannel.Contract.DepositPool(&_GenericConditionalChannel.CallOpts)
}

// DepositPool is a free data retrieval call binding the contract method 0x2a343709.
//
// Solidity: function depositPool() constant returns(address)
func (_GenericConditionalChannel *GenericConditionalChannelCallerSession) DepositPool() (common.Address, error) {
	return _GenericConditionalChannel.Contract.DepositPool(&_GenericConditionalChannel.CallOpts)
}

// GetChannelSettleTime is a free data retrieval call binding the contract method 0x42e556ce.
//
// Solidity: function getChannelSettleTime(_channelId uint256) constant returns(uint256)
func (_GenericConditionalChannel *GenericConditionalChannelCaller) GetChannelSettleTime(opts *bind.CallOpts, _channelId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _GenericConditionalChannel.contract.Call(opts, out, "getChannelSettleTime", _channelId)
	return *ret0, err
}

// GetChannelSettleTime is a free data retrieval call binding the contract method 0x42e556ce.
//
// Solidity: function getChannelSettleTime(_channelId uint256) constant returns(uint256)
func (_GenericConditionalChannel *GenericConditionalChannelSession) GetChannelSettleTime(_channelId *big.Int) (*big.Int, error) {
	return _GenericConditionalChannel.Contract.GetChannelSettleTime(&_GenericConditionalChannel.CallOpts, _channelId)
}

// GetChannelSettleTime is a free data retrieval call binding the contract method 0x42e556ce.
//
// Solidity: function getChannelSettleTime(_channelId uint256) constant returns(uint256)
func (_GenericConditionalChannel *GenericConditionalChannelCallerSession) GetChannelSettleTime(_channelId *big.Int) (*big.Int, error) {
	return _GenericConditionalChannel.Contract.GetChannelSettleTime(&_GenericConditionalChannel.CallOpts, _channelId)
}

// GetChannelStatus is a free data retrieval call binding the contract method 0xc3edbc04.
//
// Solidity: function getChannelStatus(_channelId uint256) constant returns(uint8)
func (_GenericConditionalChannel *GenericConditionalChannelCaller) GetChannelStatus(opts *bind.CallOpts, _channelId *big.Int) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _GenericConditionalChannel.contract.Call(opts, out, "getChannelStatus", _channelId)
	return *ret0, err
}

// GetChannelStatus is a free data retrieval call binding the contract method 0xc3edbc04.
//
// Solidity: function getChannelStatus(_channelId uint256) constant returns(uint8)
func (_GenericConditionalChannel *GenericConditionalChannelSession) GetChannelStatus(_channelId *big.Int) (uint8, error) {
	return _GenericConditionalChannel.Contract.GetChannelStatus(&_GenericConditionalChannel.CallOpts, _channelId)
}

// GetChannelStatus is a free data retrieval call binding the contract method 0xc3edbc04.
//
// Solidity: function getChannelStatus(_channelId uint256) constant returns(uint8)
func (_GenericConditionalChannel *GenericConditionalChannelCallerSession) GetChannelStatus(_channelId *big.Int) (uint8, error) {
	return _GenericConditionalChannel.Contract.GetChannelStatus(&_GenericConditionalChannel.CallOpts, _channelId)
}

// GetDepositAmount is a free data retrieval call binding the contract method 0xde75a5c0.
//
// Solidity: function getDepositAmount(_channelId uint256, _peer address) constant returns(uint256)
func (_GenericConditionalChannel *GenericConditionalChannelCaller) GetDepositAmount(opts *bind.CallOpts, _channelId *big.Int, _peer common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _GenericConditionalChannel.contract.Call(opts, out, "getDepositAmount", _channelId, _peer)
	return *ret0, err
}

// GetDepositAmount is a free data retrieval call binding the contract method 0xde75a5c0.
//
// Solidity: function getDepositAmount(_channelId uint256, _peer address) constant returns(uint256)
func (_GenericConditionalChannel *GenericConditionalChannelSession) GetDepositAmount(_channelId *big.Int, _peer common.Address) (*big.Int, error) {
	return _GenericConditionalChannel.Contract.GetDepositAmount(&_GenericConditionalChannel.CallOpts, _channelId, _peer)
}

// GetDepositAmount is a free data retrieval call binding the contract method 0xde75a5c0.
//
// Solidity: function getDepositAmount(_channelId uint256, _peer address) constant returns(uint256)
func (_GenericConditionalChannel *GenericConditionalChannelCallerSession) GetDepositAmount(_channelId *big.Int, _peer common.Address) (*big.Int, error) {
	return _GenericConditionalChannel.Contract.GetDepositAmount(&_GenericConditionalChannel.CallOpts, _channelId, _peer)
}

// GetDepositMap is a free data retrieval call binding the contract method 0x298e7c69.
//
// Solidity: function getDepositMap(_channelId uint256) constant returns(address[], uint256[])
func (_GenericConditionalChannel *GenericConditionalChannelCaller) GetDepositMap(opts *bind.CallOpts, _channelId *big.Int) ([]common.Address, []*big.Int, error) {
	var (
		ret0 = new([]common.Address)
		ret1 = new([]*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _GenericConditionalChannel.contract.Call(opts, out, "getDepositMap", _channelId)
	return *ret0, *ret1, err
}

// GetDepositMap is a free data retrieval call binding the contract method 0x298e7c69.
//
// Solidity: function getDepositMap(_channelId uint256) constant returns(address[], uint256[])
func (_GenericConditionalChannel *GenericConditionalChannelSession) GetDepositMap(_channelId *big.Int) ([]common.Address, []*big.Int, error) {
	return _GenericConditionalChannel.Contract.GetDepositMap(&_GenericConditionalChannel.CallOpts, _channelId)
}

// GetDepositMap is a free data retrieval call binding the contract method 0x298e7c69.
//
// Solidity: function getDepositMap(_channelId uint256) constant returns(address[], uint256[])
func (_GenericConditionalChannel *GenericConditionalChannelCallerSession) GetDepositMap(_channelId *big.Int) ([]common.Address, []*big.Int, error) {
	return _GenericConditionalChannel.Contract.GetDepositMap(&_GenericConditionalChannel.CallOpts, _channelId)
}

// GetTokenContract is a free data retrieval call binding the contract method 0xa07e5199.
//
// Solidity: function getTokenContract(_channelId uint256) constant returns(address)
func (_GenericConditionalChannel *GenericConditionalChannelCaller) GetTokenContract(opts *bind.CallOpts, _channelId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _GenericConditionalChannel.contract.Call(opts, out, "getTokenContract", _channelId)
	return *ret0, err
}

// GetTokenContract is a free data retrieval call binding the contract method 0xa07e5199.
//
// Solidity: function getTokenContract(_channelId uint256) constant returns(address)
func (_GenericConditionalChannel *GenericConditionalChannelSession) GetTokenContract(_channelId *big.Int) (common.Address, error) {
	return _GenericConditionalChannel.Contract.GetTokenContract(&_GenericConditionalChannel.CallOpts, _channelId)
}

// GetTokenContract is a free data retrieval call binding the contract method 0xa07e5199.
//
// Solidity: function getTokenContract(_channelId uint256) constant returns(address)
func (_GenericConditionalChannel *GenericConditionalChannelCallerSession) GetTokenContract(_channelId *big.Int) (common.Address, error) {
	return _GenericConditionalChannel.Contract.GetTokenContract(&_GenericConditionalChannel.CallOpts, _channelId)
}

// GetTokenType is a free data retrieval call binding the contract method 0x76aed2e1.
//
// Solidity: function getTokenType(_channelId uint256) constant returns(uint8)
func (_GenericConditionalChannel *GenericConditionalChannelCaller) GetTokenType(opts *bind.CallOpts, _channelId *big.Int) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _GenericConditionalChannel.contract.Call(opts, out, "getTokenType", _channelId)
	return *ret0, err
}

// GetTokenType is a free data retrieval call binding the contract method 0x76aed2e1.
//
// Solidity: function getTokenType(_channelId uint256) constant returns(uint8)
func (_GenericConditionalChannel *GenericConditionalChannelSession) GetTokenType(_channelId *big.Int) (uint8, error) {
	return _GenericConditionalChannel.Contract.GetTokenType(&_GenericConditionalChannel.CallOpts, _channelId)
}

// GetTokenType is a free data retrieval call binding the contract method 0x76aed2e1.
//
// Solidity: function getTokenType(_channelId uint256) constant returns(uint8)
func (_GenericConditionalChannel *GenericConditionalChannelCallerSession) GetTokenType(_channelId *big.Int) (uint8, error) {
	return _GenericConditionalChannel.Contract.GetTokenType(&_GenericConditionalChannel.CallOpts, _channelId)
}

// Resolver is a free data retrieval call binding the contract method 0x04f3bcec.
//
// Solidity: function resolver() constant returns(address)
func (_GenericConditionalChannel *GenericConditionalChannelCaller) Resolver(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _GenericConditionalChannel.contract.Call(opts, out, "resolver")
	return *ret0, err
}

// Resolver is a free data retrieval call binding the contract method 0x04f3bcec.
//
// Solidity: function resolver() constant returns(address)
func (_GenericConditionalChannel *GenericConditionalChannelSession) Resolver() (common.Address, error) {
	return _GenericConditionalChannel.Contract.Resolver(&_GenericConditionalChannel.CallOpts)
}

// Resolver is a free data retrieval call binding the contract method 0x04f3bcec.
//
// Solidity: function resolver() constant returns(address)
func (_GenericConditionalChannel *GenericConditionalChannelCallerSession) Resolver() (common.Address, error) {
	return _GenericConditionalChannel.Contract.Resolver(&_GenericConditionalChannel.CallOpts)
}

// AuthOpenChannel is a paid mutator transaction binding the contract method 0xf53a7280.
//
// Solidity: function authOpenChannel(_authWithdraw bytes, _signature bytes) returns()
func (_GenericConditionalChannel *GenericConditionalChannelTransactor) AuthOpenChannel(opts *bind.TransactOpts, _authWithdraw []byte, _signature []byte) (*types.Transaction, error) {
	return _GenericConditionalChannel.contract.Transact(opts, "authOpenChannel", _authWithdraw, _signature)
}

// AuthOpenChannel is a paid mutator transaction binding the contract method 0xf53a7280.
//
// Solidity: function authOpenChannel(_authWithdraw bytes, _signature bytes) returns()
func (_GenericConditionalChannel *GenericConditionalChannelSession) AuthOpenChannel(_authWithdraw []byte, _signature []byte) (*types.Transaction, error) {
	return _GenericConditionalChannel.Contract.AuthOpenChannel(&_GenericConditionalChannel.TransactOpts, _authWithdraw, _signature)
}

// AuthOpenChannel is a paid mutator transaction binding the contract method 0xf53a7280.
//
// Solidity: function authOpenChannel(_authWithdraw bytes, _signature bytes) returns()
func (_GenericConditionalChannel *GenericConditionalChannelTransactorSession) AuthOpenChannel(_authWithdraw []byte, _signature []byte) (*types.Transaction, error) {
	return _GenericConditionalChannel.Contract.AuthOpenChannel(&_GenericConditionalChannel.TransactOpts, _authWithdraw, _signature)
}

// ConfirmSettle is a paid mutator transaction binding the contract method 0x2d9df241.
//
// Solidity: function confirmSettle(_channelId uint256) returns()
func (_GenericConditionalChannel *GenericConditionalChannelTransactor) ConfirmSettle(opts *bind.TransactOpts, _channelId *big.Int) (*types.Transaction, error) {
	return _GenericConditionalChannel.contract.Transact(opts, "confirmSettle", _channelId)
}

// ConfirmSettle is a paid mutator transaction binding the contract method 0x2d9df241.
//
// Solidity: function confirmSettle(_channelId uint256) returns()
func (_GenericConditionalChannel *GenericConditionalChannelSession) ConfirmSettle(_channelId *big.Int) (*types.Transaction, error) {
	return _GenericConditionalChannel.Contract.ConfirmSettle(&_GenericConditionalChannel.TransactOpts, _channelId)
}

// ConfirmSettle is a paid mutator transaction binding the contract method 0x2d9df241.
//
// Solidity: function confirmSettle(_channelId uint256) returns()
func (_GenericConditionalChannel *GenericConditionalChannelTransactorSession) ConfirmSettle(_channelId *big.Int) (*types.Transaction, error) {
	return _GenericConditionalChannel.Contract.ConfirmSettle(&_GenericConditionalChannel.TransactOpts, _channelId)
}

// ConfirmWithdraw is a paid mutator transaction binding the contract method 0xaa7c2214.
//
// Solidity: function confirmWithdraw(_channelId uint256, _withdrawId uint256) returns()
func (_GenericConditionalChannel *GenericConditionalChannelTransactor) ConfirmWithdraw(opts *bind.TransactOpts, _channelId *big.Int, _withdrawId *big.Int) (*types.Transaction, error) {
	return _GenericConditionalChannel.contract.Transact(opts, "confirmWithdraw", _channelId, _withdrawId)
}

// ConfirmWithdraw is a paid mutator transaction binding the contract method 0xaa7c2214.
//
// Solidity: function confirmWithdraw(_channelId uint256, _withdrawId uint256) returns()
func (_GenericConditionalChannel *GenericConditionalChannelSession) ConfirmWithdraw(_channelId *big.Int, _withdrawId *big.Int) (*types.Transaction, error) {
	return _GenericConditionalChannel.Contract.ConfirmWithdraw(&_GenericConditionalChannel.TransactOpts, _channelId, _withdrawId)
}

// ConfirmWithdraw is a paid mutator transaction binding the contract method 0xaa7c2214.
//
// Solidity: function confirmWithdraw(_channelId uint256, _withdrawId uint256) returns()
func (_GenericConditionalChannel *GenericConditionalChannelTransactorSession) ConfirmWithdraw(_channelId *big.Int, _withdrawId *big.Int) (*types.Transaction, error) {
	return _GenericConditionalChannel.Contract.ConfirmWithdraw(&_GenericConditionalChannel.TransactOpts, _channelId, _withdrawId)
}

// CooperativeSettle is a paid mutator transaction binding the contract method 0x58ffdbde.
//
// Solidity: function cooperativeSettle(_channelId uint256, _stateProof bytes, _signature bytes, _signatureOfSignature bytes) returns()
func (_GenericConditionalChannel *GenericConditionalChannelTransactor) CooperativeSettle(opts *bind.TransactOpts, _channelId *big.Int, _stateProof []byte, _signature []byte, _signatureOfSignature []byte) (*types.Transaction, error) {
	return _GenericConditionalChannel.contract.Transact(opts, "cooperativeSettle", _channelId, _stateProof, _signature, _signatureOfSignature)
}

// CooperativeSettle is a paid mutator transaction binding the contract method 0x58ffdbde.
//
// Solidity: function cooperativeSettle(_channelId uint256, _stateProof bytes, _signature bytes, _signatureOfSignature bytes) returns()
func (_GenericConditionalChannel *GenericConditionalChannelSession) CooperativeSettle(_channelId *big.Int, _stateProof []byte, _signature []byte, _signatureOfSignature []byte) (*types.Transaction, error) {
	return _GenericConditionalChannel.Contract.CooperativeSettle(&_GenericConditionalChannel.TransactOpts, _channelId, _stateProof, _signature, _signatureOfSignature)
}

// CooperativeSettle is a paid mutator transaction binding the contract method 0x58ffdbde.
//
// Solidity: function cooperativeSettle(_channelId uint256, _stateProof bytes, _signature bytes, _signatureOfSignature bytes) returns()
func (_GenericConditionalChannel *GenericConditionalChannelTransactorSession) CooperativeSettle(_channelId *big.Int, _stateProof []byte, _signature []byte, _signatureOfSignature []byte) (*types.Transaction, error) {
	return _GenericConditionalChannel.Contract.CooperativeSettle(&_GenericConditionalChannel.TransactOpts, _channelId, _stateProof, _signature, _signatureOfSignature)
}

// CooperativeWithdraw is a paid mutator transaction binding the contract method 0xb194b0b7.
//
// Solidity: function cooperativeWithdraw(_channelId uint256, _cooperativeWithdrawProof bytes, _signature bytes) returns()
func (_GenericConditionalChannel *GenericConditionalChannelTransactor) CooperativeWithdraw(opts *bind.TransactOpts, _channelId *big.Int, _cooperativeWithdrawProof []byte, _signature []byte) (*types.Transaction, error) {
	return _GenericConditionalChannel.contract.Transact(opts, "cooperativeWithdraw", _channelId, _cooperativeWithdrawProof, _signature)
}

// CooperativeWithdraw is a paid mutator transaction binding the contract method 0xb194b0b7.
//
// Solidity: function cooperativeWithdraw(_channelId uint256, _cooperativeWithdrawProof bytes, _signature bytes) returns()
func (_GenericConditionalChannel *GenericConditionalChannelSession) CooperativeWithdraw(_channelId *big.Int, _cooperativeWithdrawProof []byte, _signature []byte) (*types.Transaction, error) {
	return _GenericConditionalChannel.Contract.CooperativeWithdraw(&_GenericConditionalChannel.TransactOpts, _channelId, _cooperativeWithdrawProof, _signature)
}

// CooperativeWithdraw is a paid mutator transaction binding the contract method 0xb194b0b7.
//
// Solidity: function cooperativeWithdraw(_channelId uint256, _cooperativeWithdrawProof bytes, _signature bytes) returns()
func (_GenericConditionalChannel *GenericConditionalChannelTransactorSession) CooperativeWithdraw(_channelId *big.Int, _cooperativeWithdrawProof []byte, _signature []byte) (*types.Transaction, error) {
	return _GenericConditionalChannel.Contract.CooperativeWithdraw(&_GenericConditionalChannel.TransactOpts, _channelId, _cooperativeWithdrawProof, _signature)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(_channelId uint256, _receipient address) returns()
func (_GenericConditionalChannel *GenericConditionalChannelTransactor) Deposit(opts *bind.TransactOpts, _channelId *big.Int, _receipient common.Address) (*types.Transaction, error) {
	return _GenericConditionalChannel.contract.Transact(opts, "deposit", _channelId, _receipient)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(_channelId uint256, _receipient address) returns()
func (_GenericConditionalChannel *GenericConditionalChannelSession) Deposit(_channelId *big.Int, _receipient common.Address) (*types.Transaction, error) {
	return _GenericConditionalChannel.Contract.Deposit(&_GenericConditionalChannel.TransactOpts, _channelId, _receipient)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(_channelId uint256, _receipient address) returns()
func (_GenericConditionalChannel *GenericConditionalChannelTransactorSession) Deposit(_channelId *big.Int, _receipient common.Address) (*types.Transaction, error) {
	return _GenericConditionalChannel.Contract.Deposit(&_GenericConditionalChannel.TransactOpts, _channelId, _receipient)
}

// DepositERCToken is a paid mutator transaction binding the contract method 0x235edd9a.
//
// Solidity: function depositERCToken(_channelId uint256, _receipient address, _amount uint256) returns()
func (_GenericConditionalChannel *GenericConditionalChannelTransactor) DepositERCToken(opts *bind.TransactOpts, _channelId *big.Int, _receipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GenericConditionalChannel.contract.Transact(opts, "depositERCToken", _channelId, _receipient, _amount)
}

// DepositERCToken is a paid mutator transaction binding the contract method 0x235edd9a.
//
// Solidity: function depositERCToken(_channelId uint256, _receipient address, _amount uint256) returns()
func (_GenericConditionalChannel *GenericConditionalChannelSession) DepositERCToken(_channelId *big.Int, _receipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GenericConditionalChannel.Contract.DepositERCToken(&_GenericConditionalChannel.TransactOpts, _channelId, _receipient, _amount)
}

// DepositERCToken is a paid mutator transaction binding the contract method 0x235edd9a.
//
// Solidity: function depositERCToken(_channelId uint256, _receipient address, _amount uint256) returns()
func (_GenericConditionalChannel *GenericConditionalChannelTransactorSession) DepositERCToken(_channelId *big.Int, _receipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GenericConditionalChannel.Contract.DepositERCToken(&_GenericConditionalChannel.TransactOpts, _channelId, _receipient, _amount)
}

// DisputeWithdraw is a paid mutator transaction binding the contract method 0xab7a0a80.
//
// Solidity: function disputeWithdraw(_channelId uint256, _withdrawId uint256, dispute bytes) returns()
func (_GenericConditionalChannel *GenericConditionalChannelTransactor) DisputeWithdraw(opts *bind.TransactOpts, _channelId *big.Int, _withdrawId *big.Int, dispute []byte) (*types.Transaction, error) {
	return _GenericConditionalChannel.contract.Transact(opts, "disputeWithdraw", _channelId, _withdrawId, dispute)
}

// DisputeWithdraw is a paid mutator transaction binding the contract method 0xab7a0a80.
//
// Solidity: function disputeWithdraw(_channelId uint256, _withdrawId uint256, dispute bytes) returns()
func (_GenericConditionalChannel *GenericConditionalChannelSession) DisputeWithdraw(_channelId *big.Int, _withdrawId *big.Int, dispute []byte) (*types.Transaction, error) {
	return _GenericConditionalChannel.Contract.DisputeWithdraw(&_GenericConditionalChannel.TransactOpts, _channelId, _withdrawId, dispute)
}

// DisputeWithdraw is a paid mutator transaction binding the contract method 0xab7a0a80.
//
// Solidity: function disputeWithdraw(_channelId uint256, _withdrawId uint256, dispute bytes) returns()
func (_GenericConditionalChannel *GenericConditionalChannelTransactorSession) DisputeWithdraw(_channelId *big.Int, _withdrawId *big.Int, dispute []byte) (*types.Transaction, error) {
	return _GenericConditionalChannel.Contract.DisputeWithdraw(&_GenericConditionalChannel.TransactOpts, _channelId, _withdrawId, dispute)
}

// IntendSettleStateProof is a paid mutator transaction binding the contract method 0x6d80e4c2.
//
// Solidity: function intendSettleStateProof(_channelId uint256, _stateProof bytes, _signature bytes) returns()
func (_GenericConditionalChannel *GenericConditionalChannelTransactor) IntendSettleStateProof(opts *bind.TransactOpts, _channelId *big.Int, _stateProof []byte, _signature []byte) (*types.Transaction, error) {
	return _GenericConditionalChannel.contract.Transact(opts, "intendSettleStateProof", _channelId, _stateProof, _signature)
}

// IntendSettleStateProof is a paid mutator transaction binding the contract method 0x6d80e4c2.
//
// Solidity: function intendSettleStateProof(_channelId uint256, _stateProof bytes, _signature bytes) returns()
func (_GenericConditionalChannel *GenericConditionalChannelSession) IntendSettleStateProof(_channelId *big.Int, _stateProof []byte, _signature []byte) (*types.Transaction, error) {
	return _GenericConditionalChannel.Contract.IntendSettleStateProof(&_GenericConditionalChannel.TransactOpts, _channelId, _stateProof, _signature)
}

// IntendSettleStateProof is a paid mutator transaction binding the contract method 0x6d80e4c2.
//
// Solidity: function intendSettleStateProof(_channelId uint256, _stateProof bytes, _signature bytes) returns()
func (_GenericConditionalChannel *GenericConditionalChannelTransactorSession) IntendSettleStateProof(_channelId *big.Int, _stateProof []byte, _signature []byte) (*types.Transaction, error) {
	return _GenericConditionalChannel.Contract.IntendSettleStateProof(&_GenericConditionalChannel.TransactOpts, _channelId, _stateProof, _signature)
}

// IntendWithdraw is a paid mutator transaction binding the contract method 0x62771a77.
//
// Solidity: function intendWithdraw(_channelId uint256, _amount uint256) returns()
func (_GenericConditionalChannel *GenericConditionalChannelTransactor) IntendWithdraw(opts *bind.TransactOpts, _channelId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _GenericConditionalChannel.contract.Transact(opts, "intendWithdraw", _channelId, _amount)
}

// IntendWithdraw is a paid mutator transaction binding the contract method 0x62771a77.
//
// Solidity: function intendWithdraw(_channelId uint256, _amount uint256) returns()
func (_GenericConditionalChannel *GenericConditionalChannelSession) IntendWithdraw(_channelId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _GenericConditionalChannel.Contract.IntendWithdraw(&_GenericConditionalChannel.TransactOpts, _channelId, _amount)
}

// IntendWithdraw is a paid mutator transaction binding the contract method 0x62771a77.
//
// Solidity: function intendWithdraw(_channelId uint256, _amount uint256) returns()
func (_GenericConditionalChannel *GenericConditionalChannelTransactorSession) IntendWithdraw(_channelId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _GenericConditionalChannel.Contract.IntendWithdraw(&_GenericConditionalChannel.TransactOpts, _channelId, _amount)
}

// OpenChannel is a paid mutator transaction binding the contract method 0x417bac3f.
//
// Solidity: function openChannel(_peers address[], _withdrawalTimeout uint256[], _settleTimeoutIncrement uint256, _tokenContract address, _tokenType uint256) returns()
func (_GenericConditionalChannel *GenericConditionalChannelTransactor) OpenChannel(opts *bind.TransactOpts, _peers []common.Address, _withdrawalTimeout []*big.Int, _settleTimeoutIncrement *big.Int, _tokenContract common.Address, _tokenType *big.Int) (*types.Transaction, error) {
	return _GenericConditionalChannel.contract.Transact(opts, "openChannel", _peers, _withdrawalTimeout, _settleTimeoutIncrement, _tokenContract, _tokenType)
}

// OpenChannel is a paid mutator transaction binding the contract method 0x417bac3f.
//
// Solidity: function openChannel(_peers address[], _withdrawalTimeout uint256[], _settleTimeoutIncrement uint256, _tokenContract address, _tokenType uint256) returns()
func (_GenericConditionalChannel *GenericConditionalChannelSession) OpenChannel(_peers []common.Address, _withdrawalTimeout []*big.Int, _settleTimeoutIncrement *big.Int, _tokenContract common.Address, _tokenType *big.Int) (*types.Transaction, error) {
	return _GenericConditionalChannel.Contract.OpenChannel(&_GenericConditionalChannel.TransactOpts, _peers, _withdrawalTimeout, _settleTimeoutIncrement, _tokenContract, _tokenType)
}

// OpenChannel is a paid mutator transaction binding the contract method 0x417bac3f.
//
// Solidity: function openChannel(_peers address[], _withdrawalTimeout uint256[], _settleTimeoutIncrement uint256, _tokenContract address, _tokenType uint256) returns()
func (_GenericConditionalChannel *GenericConditionalChannelTransactorSession) OpenChannel(_peers []common.Address, _withdrawalTimeout []*big.Int, _settleTimeoutIncrement *big.Int, _tokenContract common.Address, _tokenType *big.Int) (*types.Transaction, error) {
	return _GenericConditionalChannel.Contract.OpenChannel(&_GenericConditionalChannel.TransactOpts, _peers, _withdrawalTimeout, _settleTimeoutIncrement, _tokenContract, _tokenType)
}

// ResolveConditionalStateTransition is a paid mutator transaction binding the contract method 0xd4c84642.
//
// Solidity: function resolveConditionalStateTransition(_channelId uint256, _proof bytes32[], _conditionGroup bytes) returns()
func (_GenericConditionalChannel *GenericConditionalChannelTransactor) ResolveConditionalStateTransition(opts *bind.TransactOpts, _channelId *big.Int, _proof [][32]byte, _conditionGroup []byte) (*types.Transaction, error) {
	return _GenericConditionalChannel.contract.Transact(opts, "resolveConditionalStateTransition", _channelId, _proof, _conditionGroup)
}

// ResolveConditionalStateTransition is a paid mutator transaction binding the contract method 0xd4c84642.
//
// Solidity: function resolveConditionalStateTransition(_channelId uint256, _proof bytes32[], _conditionGroup bytes) returns()
func (_GenericConditionalChannel *GenericConditionalChannelSession) ResolveConditionalStateTransition(_channelId *big.Int, _proof [][32]byte, _conditionGroup []byte) (*types.Transaction, error) {
	return _GenericConditionalChannel.Contract.ResolveConditionalStateTransition(&_GenericConditionalChannel.TransactOpts, _channelId, _proof, _conditionGroup)
}

// ResolveConditionalStateTransition is a paid mutator transaction binding the contract method 0xd4c84642.
//
// Solidity: function resolveConditionalStateTransition(_channelId uint256, _proof bytes32[], _conditionGroup bytes) returns()
func (_GenericConditionalChannel *GenericConditionalChannelTransactorSession) ResolveConditionalStateTransition(_channelId *big.Int, _proof [][32]byte, _conditionGroup []byte) (*types.Transaction, error) {
	return _GenericConditionalChannel.Contract.ResolveConditionalStateTransition(&_GenericConditionalChannel.TransactOpts, _channelId, _proof, _conditionGroup)
}

// GenericConditionalChannelConfirmSettleIterator is returned from FilterConfirmSettle and is used to iterate over the raw logs and unpacked data for ConfirmSettle events raised by the GenericConditionalChannel contract.
type GenericConditionalChannelConfirmSettleIterator struct {
	Event *GenericConditionalChannelConfirmSettle // Event containing the contract specifics and raw log

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
func (it *GenericConditionalChannelConfirmSettleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GenericConditionalChannelConfirmSettle)
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
		it.Event = new(GenericConditionalChannelConfirmSettle)
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
func (it *GenericConditionalChannelConfirmSettleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GenericConditionalChannelConfirmSettleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GenericConditionalChannelConfirmSettle represents a ConfirmSettle event raised by the GenericConditionalChannel contract.
type GenericConditionalChannelConfirmSettle struct {
	ChannelId *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterConfirmSettle is a free log retrieval operation binding the contract event 0x05bfa6014fe9e35763ca08066a4f5a0e562a787ba9d77d27732eb4aa0b66b236.
//
// Solidity: e ConfirmSettle(channelId uint256)
func (_GenericConditionalChannel *GenericConditionalChannelFilterer) FilterConfirmSettle(opts *bind.FilterOpts) (*GenericConditionalChannelConfirmSettleIterator, error) {

	logs, sub, err := _GenericConditionalChannel.contract.FilterLogs(opts, "ConfirmSettle")
	if err != nil {
		return nil, err
	}
	return &GenericConditionalChannelConfirmSettleIterator{contract: _GenericConditionalChannel.contract, event: "ConfirmSettle", logs: logs, sub: sub}, nil
}

// WatchConfirmSettle is a free log subscription operation binding the contract event 0x05bfa6014fe9e35763ca08066a4f5a0e562a787ba9d77d27732eb4aa0b66b236.
//
// Solidity: e ConfirmSettle(channelId uint256)
func (_GenericConditionalChannel *GenericConditionalChannelFilterer) WatchConfirmSettle(opts *bind.WatchOpts, sink chan<- *GenericConditionalChannelConfirmSettle) (event.Subscription, error) {

	logs, sub, err := _GenericConditionalChannel.contract.WatchLogs(opts, "ConfirmSettle")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GenericConditionalChannelConfirmSettle)
				if err := _GenericConditionalChannel.contract.UnpackLog(event, "ConfirmSettle", log); err != nil {
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

// GenericConditionalChannelConfirmSettleFailIterator is returned from FilterConfirmSettleFail and is used to iterate over the raw logs and unpacked data for ConfirmSettleFail events raised by the GenericConditionalChannel contract.
type GenericConditionalChannelConfirmSettleFailIterator struct {
	Event *GenericConditionalChannelConfirmSettleFail // Event containing the contract specifics and raw log

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
func (it *GenericConditionalChannelConfirmSettleFailIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GenericConditionalChannelConfirmSettleFail)
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
		it.Event = new(GenericConditionalChannelConfirmSettleFail)
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
func (it *GenericConditionalChannelConfirmSettleFailIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GenericConditionalChannelConfirmSettleFailIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GenericConditionalChannelConfirmSettleFail represents a ConfirmSettleFail event raised by the GenericConditionalChannel contract.
type GenericConditionalChannelConfirmSettleFail struct {
	ChannelId *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterConfirmSettleFail is a free log retrieval operation binding the contract event 0x5a68b21ebf14d22f6a89bd2f92d1c899f2cfff1c821cdfbebd2af882389a287b.
//
// Solidity: e ConfirmSettleFail(channelId uint256)
func (_GenericConditionalChannel *GenericConditionalChannelFilterer) FilterConfirmSettleFail(opts *bind.FilterOpts) (*GenericConditionalChannelConfirmSettleFailIterator, error) {

	logs, sub, err := _GenericConditionalChannel.contract.FilterLogs(opts, "ConfirmSettleFail")
	if err != nil {
		return nil, err
	}
	return &GenericConditionalChannelConfirmSettleFailIterator{contract: _GenericConditionalChannel.contract, event: "ConfirmSettleFail", logs: logs, sub: sub}, nil
}

// WatchConfirmSettleFail is a free log subscription operation binding the contract event 0x5a68b21ebf14d22f6a89bd2f92d1c899f2cfff1c821cdfbebd2af882389a287b.
//
// Solidity: e ConfirmSettleFail(channelId uint256)
func (_GenericConditionalChannel *GenericConditionalChannelFilterer) WatchConfirmSettleFail(opts *bind.WatchOpts, sink chan<- *GenericConditionalChannelConfirmSettleFail) (event.Subscription, error) {

	logs, sub, err := _GenericConditionalChannel.contract.WatchLogs(opts, "ConfirmSettleFail")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GenericConditionalChannelConfirmSettleFail)
				if err := _GenericConditionalChannel.contract.UnpackLog(event, "ConfirmSettleFail", log); err != nil {
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

// GenericConditionalChannelCooperativeSettleIterator is returned from FilterCooperativeSettle and is used to iterate over the raw logs and unpacked data for CooperativeSettle events raised by the GenericConditionalChannel contract.
type GenericConditionalChannelCooperativeSettleIterator struct {
	Event *GenericConditionalChannelCooperativeSettle // Event containing the contract specifics and raw log

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
func (it *GenericConditionalChannelCooperativeSettleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GenericConditionalChannelCooperativeSettle)
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
		it.Event = new(GenericConditionalChannelCooperativeSettle)
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
func (it *GenericConditionalChannelCooperativeSettleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GenericConditionalChannelCooperativeSettleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GenericConditionalChannelCooperativeSettle represents a CooperativeSettle event raised by the GenericConditionalChannel contract.
type GenericConditionalChannelCooperativeSettle struct {
	ChannelId *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCooperativeSettle is a free log retrieval operation binding the contract event 0x6230d4127b5014fc7595da99cf59ce68bac5d9276e21712c85f91c06a96c722f.
//
// Solidity: e CooperativeSettle(channelId uint256)
func (_GenericConditionalChannel *GenericConditionalChannelFilterer) FilterCooperativeSettle(opts *bind.FilterOpts) (*GenericConditionalChannelCooperativeSettleIterator, error) {

	logs, sub, err := _GenericConditionalChannel.contract.FilterLogs(opts, "CooperativeSettle")
	if err != nil {
		return nil, err
	}
	return &GenericConditionalChannelCooperativeSettleIterator{contract: _GenericConditionalChannel.contract, event: "CooperativeSettle", logs: logs, sub: sub}, nil
}

// WatchCooperativeSettle is a free log subscription operation binding the contract event 0x6230d4127b5014fc7595da99cf59ce68bac5d9276e21712c85f91c06a96c722f.
//
// Solidity: e CooperativeSettle(channelId uint256)
func (_GenericConditionalChannel *GenericConditionalChannelFilterer) WatchCooperativeSettle(opts *bind.WatchOpts, sink chan<- *GenericConditionalChannelCooperativeSettle) (event.Subscription, error) {

	logs, sub, err := _GenericConditionalChannel.contract.WatchLogs(opts, "CooperativeSettle")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GenericConditionalChannelCooperativeSettle)
				if err := _GenericConditionalChannel.contract.UnpackLog(event, "CooperativeSettle", log); err != nil {
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

// GenericConditionalChannelCooperativeSettleFailIterator is returned from FilterCooperativeSettleFail and is used to iterate over the raw logs and unpacked data for CooperativeSettleFail events raised by the GenericConditionalChannel contract.
type GenericConditionalChannelCooperativeSettleFailIterator struct {
	Event *GenericConditionalChannelCooperativeSettleFail // Event containing the contract specifics and raw log

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
func (it *GenericConditionalChannelCooperativeSettleFailIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GenericConditionalChannelCooperativeSettleFail)
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
		it.Event = new(GenericConditionalChannelCooperativeSettleFail)
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
func (it *GenericConditionalChannelCooperativeSettleFailIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GenericConditionalChannelCooperativeSettleFailIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GenericConditionalChannelCooperativeSettleFail represents a CooperativeSettleFail event raised by the GenericConditionalChannel contract.
type GenericConditionalChannelCooperativeSettleFail struct {
	ChannelId *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCooperativeSettleFail is a free log retrieval operation binding the contract event 0x9af0573e2d3394e13e9c62615248cd22581e1862e87a84ddb544247904fd9cb2.
//
// Solidity: e CooperativeSettleFail(channelId uint256)
func (_GenericConditionalChannel *GenericConditionalChannelFilterer) FilterCooperativeSettleFail(opts *bind.FilterOpts) (*GenericConditionalChannelCooperativeSettleFailIterator, error) {

	logs, sub, err := _GenericConditionalChannel.contract.FilterLogs(opts, "CooperativeSettleFail")
	if err != nil {
		return nil, err
	}
	return &GenericConditionalChannelCooperativeSettleFailIterator{contract: _GenericConditionalChannel.contract, event: "CooperativeSettleFail", logs: logs, sub: sub}, nil
}

// WatchCooperativeSettleFail is a free log subscription operation binding the contract event 0x9af0573e2d3394e13e9c62615248cd22581e1862e87a84ddb544247904fd9cb2.
//
// Solidity: e CooperativeSettleFail(channelId uint256)
func (_GenericConditionalChannel *GenericConditionalChannelFilterer) WatchCooperativeSettleFail(opts *bind.WatchOpts, sink chan<- *GenericConditionalChannelCooperativeSettleFail) (event.Subscription, error) {

	logs, sub, err := _GenericConditionalChannel.contract.WatchLogs(opts, "CooperativeSettleFail")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GenericConditionalChannelCooperativeSettleFail)
				if err := _GenericConditionalChannel.contract.UnpackLog(event, "CooperativeSettleFail", log); err != nil {
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

// GenericConditionalChannelCooperativeWithdrawIterator is returned from FilterCooperativeWithdraw and is used to iterate over the raw logs and unpacked data for CooperativeWithdraw events raised by the GenericConditionalChannel contract.
type GenericConditionalChannelCooperativeWithdrawIterator struct {
	Event *GenericConditionalChannelCooperativeWithdraw // Event containing the contract specifics and raw log

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
func (it *GenericConditionalChannelCooperativeWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GenericConditionalChannelCooperativeWithdraw)
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
		it.Event = new(GenericConditionalChannelCooperativeWithdraw)
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
func (it *GenericConditionalChannelCooperativeWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GenericConditionalChannelCooperativeWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GenericConditionalChannelCooperativeWithdraw represents a CooperativeWithdraw event raised by the GenericConditionalChannel contract.
type GenericConditionalChannelCooperativeWithdraw struct {
	ChannelId        *big.Int
	WithdrawalAmount *big.Int
	Receiver         common.Address
	Balance          *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterCooperativeWithdraw is a free log retrieval operation binding the contract event 0x0375fe0230defc24bed394834804158bdbd5e92e4c3a6dc88f4db5439080f7b8.
//
// Solidity: e CooperativeWithdraw(channelId uint256, withdrawalAmount uint256, receiver address, balance uint256)
func (_GenericConditionalChannel *GenericConditionalChannelFilterer) FilterCooperativeWithdraw(opts *bind.FilterOpts) (*GenericConditionalChannelCooperativeWithdrawIterator, error) {

	logs, sub, err := _GenericConditionalChannel.contract.FilterLogs(opts, "CooperativeWithdraw")
	if err != nil {
		return nil, err
	}
	return &GenericConditionalChannelCooperativeWithdrawIterator{contract: _GenericConditionalChannel.contract, event: "CooperativeWithdraw", logs: logs, sub: sub}, nil
}

// WatchCooperativeWithdraw is a free log subscription operation binding the contract event 0x0375fe0230defc24bed394834804158bdbd5e92e4c3a6dc88f4db5439080f7b8.
//
// Solidity: e CooperativeWithdraw(channelId uint256, withdrawalAmount uint256, receiver address, balance uint256)
func (_GenericConditionalChannel *GenericConditionalChannelFilterer) WatchCooperativeWithdraw(opts *bind.WatchOpts, sink chan<- *GenericConditionalChannelCooperativeWithdraw) (event.Subscription, error) {

	logs, sub, err := _GenericConditionalChannel.contract.WatchLogs(opts, "CooperativeWithdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GenericConditionalChannelCooperativeWithdraw)
				if err := _GenericConditionalChannel.contract.UnpackLog(event, "CooperativeWithdraw", log); err != nil {
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

// GenericConditionalChannelDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the GenericConditionalChannel contract.
type GenericConditionalChannelDepositIterator struct {
	Event *GenericConditionalChannelDeposit // Event containing the contract specifics and raw log

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
func (it *GenericConditionalChannelDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GenericConditionalChannelDeposit)
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
		it.Event = new(GenericConditionalChannelDeposit)
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
func (it *GenericConditionalChannelDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GenericConditionalChannelDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GenericConditionalChannelDeposit represents a Deposit event raised by the GenericConditionalChannel contract.
type GenericConditionalChannelDeposit struct {
	ChannelId *big.Int
	Peers     []common.Address
	Amounts   []*big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe33829dd3495a41ba02f88de5eadec2635e25192c73dee8c65625206aa61e09a.
//
// Solidity: e Deposit(channelId uint256, peers address[], amounts uint256[])
func (_GenericConditionalChannel *GenericConditionalChannelFilterer) FilterDeposit(opts *bind.FilterOpts) (*GenericConditionalChannelDepositIterator, error) {

	logs, sub, err := _GenericConditionalChannel.contract.FilterLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return &GenericConditionalChannelDepositIterator{contract: _GenericConditionalChannel.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe33829dd3495a41ba02f88de5eadec2635e25192c73dee8c65625206aa61e09a.
//
// Solidity: e Deposit(channelId uint256, peers address[], amounts uint256[])
func (_GenericConditionalChannel *GenericConditionalChannelFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *GenericConditionalChannelDeposit) (event.Subscription, error) {

	logs, sub, err := _GenericConditionalChannel.contract.WatchLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GenericConditionalChannelDeposit)
				if err := _GenericConditionalChannel.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// GenericConditionalChannelIntendSettleIterator is returned from FilterIntendSettle and is used to iterate over the raw logs and unpacked data for IntendSettle events raised by the GenericConditionalChannel contract.
type GenericConditionalChannelIntendSettleIterator struct {
	Event *GenericConditionalChannelIntendSettle // Event containing the contract specifics and raw log

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
func (it *GenericConditionalChannelIntendSettleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GenericConditionalChannelIntendSettle)
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
		it.Event = new(GenericConditionalChannelIntendSettle)
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
func (it *GenericConditionalChannelIntendSettleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GenericConditionalChannelIntendSettleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GenericConditionalChannelIntendSettle represents a IntendSettle event raised by the GenericConditionalChannel contract.
type GenericConditionalChannelIntendSettle struct {
	ChannelId       *big.Int
	StateProofNonce *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterIntendSettle is a free log retrieval operation binding the contract event 0x5adbec4d8dfaaf515f95e278796bcb261b13c8df87a959188fd5f5fb05bccd02.
//
// Solidity: e IntendSettle(channelId uint256, stateProofNonce uint256)
func (_GenericConditionalChannel *GenericConditionalChannelFilterer) FilterIntendSettle(opts *bind.FilterOpts) (*GenericConditionalChannelIntendSettleIterator, error) {

	logs, sub, err := _GenericConditionalChannel.contract.FilterLogs(opts, "IntendSettle")
	if err != nil {
		return nil, err
	}
	return &GenericConditionalChannelIntendSettleIterator{contract: _GenericConditionalChannel.contract, event: "IntendSettle", logs: logs, sub: sub}, nil
}

// WatchIntendSettle is a free log subscription operation binding the contract event 0x5adbec4d8dfaaf515f95e278796bcb261b13c8df87a959188fd5f5fb05bccd02.
//
// Solidity: e IntendSettle(channelId uint256, stateProofNonce uint256)
func (_GenericConditionalChannel *GenericConditionalChannelFilterer) WatchIntendSettle(opts *bind.WatchOpts, sink chan<- *GenericConditionalChannelIntendSettle) (event.Subscription, error) {

	logs, sub, err := _GenericConditionalChannel.contract.WatchLogs(opts, "IntendSettle")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GenericConditionalChannelIntendSettle)
				if err := _GenericConditionalChannel.contract.UnpackLog(event, "IntendSettle", log); err != nil {
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

// GenericConditionalChannelOpenChannelIterator is returned from FilterOpenChannel and is used to iterate over the raw logs and unpacked data for OpenChannel events raised by the GenericConditionalChannel contract.
type GenericConditionalChannelOpenChannelIterator struct {
	Event *GenericConditionalChannelOpenChannel // Event containing the contract specifics and raw log

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
func (it *GenericConditionalChannelOpenChannelIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GenericConditionalChannelOpenChannel)
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
		it.Event = new(GenericConditionalChannelOpenChannel)
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
func (it *GenericConditionalChannelOpenChannelIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GenericConditionalChannelOpenChannelIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GenericConditionalChannelOpenChannel represents a OpenChannel event raised by the GenericConditionalChannel contract.
type GenericConditionalChannelOpenChannel struct {
	ChannelId     *big.Int
	Peers         []common.Address
	UintTokenType *big.Int
	TokenContract common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOpenChannel is a free log retrieval operation binding the contract event 0x9d51eb7f3cbaa052d1e423f1630d320f56f424e7c1fd7f4621654ccfd0d41b2f.
//
// Solidity: e OpenChannel(channelId uint256, peers address[], uintTokenType uint256, tokenContract address)
func (_GenericConditionalChannel *GenericConditionalChannelFilterer) FilterOpenChannel(opts *bind.FilterOpts) (*GenericConditionalChannelOpenChannelIterator, error) {

	logs, sub, err := _GenericConditionalChannel.contract.FilterLogs(opts, "OpenChannel")
	if err != nil {
		return nil, err
	}
	return &GenericConditionalChannelOpenChannelIterator{contract: _GenericConditionalChannel.contract, event: "OpenChannel", logs: logs, sub: sub}, nil
}

// WatchOpenChannel is a free log subscription operation binding the contract event 0x9d51eb7f3cbaa052d1e423f1630d320f56f424e7c1fd7f4621654ccfd0d41b2f.
//
// Solidity: e OpenChannel(channelId uint256, peers address[], uintTokenType uint256, tokenContract address)
func (_GenericConditionalChannel *GenericConditionalChannelFilterer) WatchOpenChannel(opts *bind.WatchOpts, sink chan<- *GenericConditionalChannelOpenChannel) (event.Subscription, error) {

	logs, sub, err := _GenericConditionalChannel.contract.WatchLogs(opts, "OpenChannel")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GenericConditionalChannelOpenChannel)
				if err := _GenericConditionalChannel.contract.UnpackLog(event, "OpenChannel", log); err != nil {
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

// GenericConditionalChannelResolveCondGroupIterator is returned from FilterResolveCondGroup and is used to iterate over the raw logs and unpacked data for ResolveCondGroup events raised by the GenericConditionalChannel contract.
type GenericConditionalChannelResolveCondGroupIterator struct {
	Event *GenericConditionalChannelResolveCondGroup // Event containing the contract specifics and raw log

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
func (it *GenericConditionalChannelResolveCondGroupIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GenericConditionalChannelResolveCondGroup)
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
		it.Event = new(GenericConditionalChannelResolveCondGroup)
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
func (it *GenericConditionalChannelResolveCondGroupIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GenericConditionalChannelResolveCondGroupIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GenericConditionalChannelResolveCondGroup represents a ResolveCondGroup event raised by the GenericConditionalChannel contract.
type GenericConditionalChannelResolveCondGroup struct {
	ChannelId     *big.Int
	CondGroupHash [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterResolveCondGroup is a free log retrieval operation binding the contract event 0xda7220c7681df4c0b2cb90420d72d108cc12db4cbed30a5c4f5cb59f04d3d49e.
//
// Solidity: e ResolveCondGroup(channelId uint256, condGroupHash bytes32)
func (_GenericConditionalChannel *GenericConditionalChannelFilterer) FilterResolveCondGroup(opts *bind.FilterOpts) (*GenericConditionalChannelResolveCondGroupIterator, error) {

	logs, sub, err := _GenericConditionalChannel.contract.FilterLogs(opts, "ResolveCondGroup")
	if err != nil {
		return nil, err
	}
	return &GenericConditionalChannelResolveCondGroupIterator{contract: _GenericConditionalChannel.contract, event: "ResolveCondGroup", logs: logs, sub: sub}, nil
}

// WatchResolveCondGroup is a free log subscription operation binding the contract event 0xda7220c7681df4c0b2cb90420d72d108cc12db4cbed30a5c4f5cb59f04d3d49e.
//
// Solidity: e ResolveCondGroup(channelId uint256, condGroupHash bytes32)
func (_GenericConditionalChannel *GenericConditionalChannelFilterer) WatchResolveCondGroup(opts *bind.WatchOpts, sink chan<- *GenericConditionalChannelResolveCondGroup) (event.Subscription, error) {

	logs, sub, err := _GenericConditionalChannel.contract.WatchLogs(opts, "ResolveCondGroup")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GenericConditionalChannelResolveCondGroup)
				if err := _GenericConditionalChannel.contract.UnpackLog(event, "ResolveCondGroup", log); err != nil {
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

// MerkleProofABI is the input ABI used to generate the binding from.
const MerkleProofABI = "[]"

// MerkleProofBin is the compiled bytecode used for deploying new contracts.
const MerkleProofBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a72305820c044a07fb8a6c613098fb94b5f50a125e457b76c5dde02b28e166e9dd9c465190029`

// DeployMerkleProof deploys a new Ethereum contract, binding an instance of MerkleProof to it.
func DeployMerkleProof(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MerkleProof, error) {
	parsed, err := abi.JSON(strings.NewReader(MerkleProofABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MerkleProofBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MerkleProof{MerkleProofCaller: MerkleProofCaller{contract: contract}, MerkleProofTransactor: MerkleProofTransactor{contract: contract}, MerkleProofFilterer: MerkleProofFilterer{contract: contract}}, nil
}

// MerkleProof is an auto generated Go binding around an Ethereum contract.
type MerkleProof struct {
	MerkleProofCaller     // Read-only binding to the contract
	MerkleProofTransactor // Write-only binding to the contract
	MerkleProofFilterer   // Log filterer for contract events
}

// MerkleProofCaller is an auto generated read-only Go binding around an Ethereum contract.
type MerkleProofCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleProofTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MerkleProofTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleProofFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MerkleProofFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleProofSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MerkleProofSession struct {
	Contract     *MerkleProof      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MerkleProofCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MerkleProofCallerSession struct {
	Contract *MerkleProofCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// MerkleProofTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MerkleProofTransactorSession struct {
	Contract     *MerkleProofTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// MerkleProofRaw is an auto generated low-level Go binding around an Ethereum contract.
type MerkleProofRaw struct {
	Contract *MerkleProof // Generic contract binding to access the raw methods on
}

// MerkleProofCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MerkleProofCallerRaw struct {
	Contract *MerkleProofCaller // Generic read-only contract binding to access the raw methods on
}

// MerkleProofTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MerkleProofTransactorRaw struct {
	Contract *MerkleProofTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMerkleProof creates a new instance of MerkleProof, bound to a specific deployed contract.
func NewMerkleProof(address common.Address, backend bind.ContractBackend) (*MerkleProof, error) {
	contract, err := bindMerkleProof(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MerkleProof{MerkleProofCaller: MerkleProofCaller{contract: contract}, MerkleProofTransactor: MerkleProofTransactor{contract: contract}, MerkleProofFilterer: MerkleProofFilterer{contract: contract}}, nil
}

// NewMerkleProofCaller creates a new read-only instance of MerkleProof, bound to a specific deployed contract.
func NewMerkleProofCaller(address common.Address, caller bind.ContractCaller) (*MerkleProofCaller, error) {
	contract, err := bindMerkleProof(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MerkleProofCaller{contract: contract}, nil
}

// NewMerkleProofTransactor creates a new write-only instance of MerkleProof, bound to a specific deployed contract.
func NewMerkleProofTransactor(address common.Address, transactor bind.ContractTransactor) (*MerkleProofTransactor, error) {
	contract, err := bindMerkleProof(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MerkleProofTransactor{contract: contract}, nil
}

// NewMerkleProofFilterer creates a new log filterer instance of MerkleProof, bound to a specific deployed contract.
func NewMerkleProofFilterer(address common.Address, filterer bind.ContractFilterer) (*MerkleProofFilterer, error) {
	contract, err := bindMerkleProof(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MerkleProofFilterer{contract: contract}, nil
}

// bindMerkleProof binds a generic wrapper to an already deployed contract.
func bindMerkleProof(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MerkleProofABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MerkleProof *MerkleProofRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MerkleProof.Contract.MerkleProofCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MerkleProof *MerkleProofRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MerkleProof.Contract.MerkleProofTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MerkleProof *MerkleProofRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MerkleProof.Contract.MerkleProofTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MerkleProof *MerkleProofCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MerkleProof.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MerkleProof *MerkleProofTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MerkleProof.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MerkleProof *MerkleProofTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MerkleProof.Contract.contract.Transact(opts, method, params...)
}

// SafeERC20ABI is the input ABI used to generate the binding from.
const SafeERC20ABI = "[]"

// SafeERC20Bin is the compiled bytecode used for deploying new contracts.
const SafeERC20Bin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a72305820e6837ce769829bfe677748ffc0da0ed9838bb6b503c441960de30ab46c080fa70029`

// DeploySafeERC20 deploys a new Ethereum contract, binding an instance of SafeERC20 to it.
func DeploySafeERC20(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeERC20, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeERC20ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeERC20Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeERC20{SafeERC20Caller: SafeERC20Caller{contract: contract}, SafeERC20Transactor: SafeERC20Transactor{contract: contract}, SafeERC20Filterer: SafeERC20Filterer{contract: contract}}, nil
}

// SafeERC20 is an auto generated Go binding around an Ethereum contract.
type SafeERC20 struct {
	SafeERC20Caller     // Read-only binding to the contract
	SafeERC20Transactor // Write-only binding to the contract
	SafeERC20Filterer   // Log filterer for contract events
}

// SafeERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type SafeERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeERC20Session struct {
	Contract     *SafeERC20        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeERC20CallerSession struct {
	Contract *SafeERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// SafeERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeERC20TransactorSession struct {
	Contract     *SafeERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SafeERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type SafeERC20Raw struct {
	Contract *SafeERC20 // Generic contract binding to access the raw methods on
}

// SafeERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeERC20CallerRaw struct {
	Contract *SafeERC20Caller // Generic read-only contract binding to access the raw methods on
}

// SafeERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeERC20TransactorRaw struct {
	Contract *SafeERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeERC20 creates a new instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20(address common.Address, backend bind.ContractBackend) (*SafeERC20, error) {
	contract, err := bindSafeERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeERC20{SafeERC20Caller: SafeERC20Caller{contract: contract}, SafeERC20Transactor: SafeERC20Transactor{contract: contract}, SafeERC20Filterer: SafeERC20Filterer{contract: contract}}, nil
}

// NewSafeERC20Caller creates a new read-only instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20Caller(address common.Address, caller bind.ContractCaller) (*SafeERC20Caller, error) {
	contract, err := bindSafeERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeERC20Caller{contract: contract}, nil
}

// NewSafeERC20Transactor creates a new write-only instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*SafeERC20Transactor, error) {
	contract, err := bindSafeERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeERC20Transactor{contract: contract}, nil
}

// NewSafeERC20Filterer creates a new log filterer instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*SafeERC20Filterer, error) {
	contract, err := bindSafeERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeERC20Filterer{contract: contract}, nil
}

// bindSafeERC20 binds a generic wrapper to an already deployed contract.
func bindSafeERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeERC20 *SafeERC20Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeERC20.Contract.SafeERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeERC20 *SafeERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeERC20.Contract.SafeERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeERC20 *SafeERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeERC20.Contract.SafeERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeERC20 *SafeERC20CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeERC20 *SafeERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeERC20 *SafeERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeERC20.Contract.contract.Transact(opts, method, params...)
}

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
const SafeMathBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a7230582078cb3175d55f2a85d43aa57aeb4f714a413e157c9c71472c41dafc1ac270732a0029`

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

// VirtualChannelResolverInterfaceABI is the input ABI used to generate the binding from.
const VirtualChannelResolverInterfaceABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_virtualAddr\",\"type\":\"bytes32\"}],\"name\":\"resolve\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"code\",\"type\":\"bytes\"},{\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"deploy\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// VirtualChannelResolverInterfaceBin is the compiled bytecode used for deploying new contracts.
const VirtualChannelResolverInterfaceBin = `0x`

// DeployVirtualChannelResolverInterface deploys a new Ethereum contract, binding an instance of VirtualChannelResolverInterface to it.
func DeployVirtualChannelResolverInterface(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *VirtualChannelResolverInterface, error) {
	parsed, err := abi.JSON(strings.NewReader(VirtualChannelResolverInterfaceABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(VirtualChannelResolverInterfaceBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VirtualChannelResolverInterface{VirtualChannelResolverInterfaceCaller: VirtualChannelResolverInterfaceCaller{contract: contract}, VirtualChannelResolverInterfaceTransactor: VirtualChannelResolverInterfaceTransactor{contract: contract}, VirtualChannelResolverInterfaceFilterer: VirtualChannelResolverInterfaceFilterer{contract: contract}}, nil
}

// VirtualChannelResolverInterface is an auto generated Go binding around an Ethereum contract.
type VirtualChannelResolverInterface struct {
	VirtualChannelResolverInterfaceCaller     // Read-only binding to the contract
	VirtualChannelResolverInterfaceTransactor // Write-only binding to the contract
	VirtualChannelResolverInterfaceFilterer   // Log filterer for contract events
}

// VirtualChannelResolverInterfaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type VirtualChannelResolverInterfaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VirtualChannelResolverInterfaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VirtualChannelResolverInterfaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VirtualChannelResolverInterfaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VirtualChannelResolverInterfaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VirtualChannelResolverInterfaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VirtualChannelResolverInterfaceSession struct {
	Contract     *VirtualChannelResolverInterface // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                    // Call options to use throughout this session
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// VirtualChannelResolverInterfaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VirtualChannelResolverInterfaceCallerSession struct {
	Contract *VirtualChannelResolverInterfaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                          // Call options to use throughout this session
}

// VirtualChannelResolverInterfaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VirtualChannelResolverInterfaceTransactorSession struct {
	Contract     *VirtualChannelResolverInterfaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                          // Transaction auth options to use throughout this session
}

// VirtualChannelResolverInterfaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type VirtualChannelResolverInterfaceRaw struct {
	Contract *VirtualChannelResolverInterface // Generic contract binding to access the raw methods on
}

// VirtualChannelResolverInterfaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VirtualChannelResolverInterfaceCallerRaw struct {
	Contract *VirtualChannelResolverInterfaceCaller // Generic read-only contract binding to access the raw methods on
}

// VirtualChannelResolverInterfaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VirtualChannelResolverInterfaceTransactorRaw struct {
	Contract *VirtualChannelResolverInterfaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVirtualChannelResolverInterface creates a new instance of VirtualChannelResolverInterface, bound to a specific deployed contract.
func NewVirtualChannelResolverInterface(address common.Address, backend bind.ContractBackend) (*VirtualChannelResolverInterface, error) {
	contract, err := bindVirtualChannelResolverInterface(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VirtualChannelResolverInterface{VirtualChannelResolverInterfaceCaller: VirtualChannelResolverInterfaceCaller{contract: contract}, VirtualChannelResolverInterfaceTransactor: VirtualChannelResolverInterfaceTransactor{contract: contract}, VirtualChannelResolverInterfaceFilterer: VirtualChannelResolverInterfaceFilterer{contract: contract}}, nil
}

// NewVirtualChannelResolverInterfaceCaller creates a new read-only instance of VirtualChannelResolverInterface, bound to a specific deployed contract.
func NewVirtualChannelResolverInterfaceCaller(address common.Address, caller bind.ContractCaller) (*VirtualChannelResolverInterfaceCaller, error) {
	contract, err := bindVirtualChannelResolverInterface(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VirtualChannelResolverInterfaceCaller{contract: contract}, nil
}

// NewVirtualChannelResolverInterfaceTransactor creates a new write-only instance of VirtualChannelResolverInterface, bound to a specific deployed contract.
func NewVirtualChannelResolverInterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*VirtualChannelResolverInterfaceTransactor, error) {
	contract, err := bindVirtualChannelResolverInterface(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VirtualChannelResolverInterfaceTransactor{contract: contract}, nil
}

// NewVirtualChannelResolverInterfaceFilterer creates a new log filterer instance of VirtualChannelResolverInterface, bound to a specific deployed contract.
func NewVirtualChannelResolverInterfaceFilterer(address common.Address, filterer bind.ContractFilterer) (*VirtualChannelResolverInterfaceFilterer, error) {
	contract, err := bindVirtualChannelResolverInterface(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VirtualChannelResolverInterfaceFilterer{contract: contract}, nil
}

// bindVirtualChannelResolverInterface binds a generic wrapper to an already deployed contract.
func bindVirtualChannelResolverInterface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VirtualChannelResolverInterfaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VirtualChannelResolverInterface *VirtualChannelResolverInterfaceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _VirtualChannelResolverInterface.Contract.VirtualChannelResolverInterfaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VirtualChannelResolverInterface *VirtualChannelResolverInterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VirtualChannelResolverInterface.Contract.VirtualChannelResolverInterfaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VirtualChannelResolverInterface *VirtualChannelResolverInterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VirtualChannelResolverInterface.Contract.VirtualChannelResolverInterfaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VirtualChannelResolverInterface *VirtualChannelResolverInterfaceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _VirtualChannelResolverInterface.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VirtualChannelResolverInterface *VirtualChannelResolverInterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VirtualChannelResolverInterface.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VirtualChannelResolverInterface *VirtualChannelResolverInterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VirtualChannelResolverInterface.Contract.contract.Transact(opts, method, params...)
}

// Resolve is a free data retrieval call binding the contract method 0x5c23bdf5.
//
// Solidity: function resolve(_virtualAddr bytes32) constant returns(address)
func (_VirtualChannelResolverInterface *VirtualChannelResolverInterfaceCaller) Resolve(opts *bind.CallOpts, _virtualAddr [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _VirtualChannelResolverInterface.contract.Call(opts, out, "resolve", _virtualAddr)
	return *ret0, err
}

// Resolve is a free data retrieval call binding the contract method 0x5c23bdf5.
//
// Solidity: function resolve(_virtualAddr bytes32) constant returns(address)
func (_VirtualChannelResolverInterface *VirtualChannelResolverInterfaceSession) Resolve(_virtualAddr [32]byte) (common.Address, error) {
	return _VirtualChannelResolverInterface.Contract.Resolve(&_VirtualChannelResolverInterface.CallOpts, _virtualAddr)
}

// Resolve is a free data retrieval call binding the contract method 0x5c23bdf5.
//
// Solidity: function resolve(_virtualAddr bytes32) constant returns(address)
func (_VirtualChannelResolverInterface *VirtualChannelResolverInterfaceCallerSession) Resolve(_virtualAddr [32]byte) (common.Address, error) {
	return _VirtualChannelResolverInterface.Contract.Resolve(&_VirtualChannelResolverInterface.CallOpts, _virtualAddr)
}

// Deploy is a paid mutator transaction binding the contract method 0x9c4ae2d0.
//
// Solidity: function deploy(code bytes, nonce uint256) returns(bool)
func (_VirtualChannelResolverInterface *VirtualChannelResolverInterfaceTransactor) Deploy(opts *bind.TransactOpts, code []byte, nonce *big.Int) (*types.Transaction, error) {
	return _VirtualChannelResolverInterface.contract.Transact(opts, "deploy", code, nonce)
}

// Deploy is a paid mutator transaction binding the contract method 0x9c4ae2d0.
//
// Solidity: function deploy(code bytes, nonce uint256) returns(bool)
func (_VirtualChannelResolverInterface *VirtualChannelResolverInterfaceSession) Deploy(code []byte, nonce *big.Int) (*types.Transaction, error) {
	return _VirtualChannelResolverInterface.Contract.Deploy(&_VirtualChannelResolverInterface.TransactOpts, code, nonce)
}

// Deploy is a paid mutator transaction binding the contract method 0x9c4ae2d0.
//
// Solidity: function deploy(code bytes, nonce uint256) returns(bool)
func (_VirtualChannelResolverInterface *VirtualChannelResolverInterfaceTransactorSession) Deploy(code []byte, nonce *big.Int) (*types.Transaction, error) {
	return _VirtualChannelResolverInterface.Contract.Deploy(&_VirtualChannelResolverInterface.TransactOpts, code, nonce)
}

// PbABI is the input ABI used to generate the binding from.
const PbABI = "[]"

// PbBin is the compiled bytecode used for deploying new contracts.
const PbBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a723058205c99cdb50a45b7392e5bccb5f87ae374e89e756811996c5099c32aae4b1d15270029`

// DeployPb deploys a new Ethereum contract, binding an instance of Pb to it.
func DeployPb(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Pb, error) {
	parsed, err := abi.JSON(strings.NewReader(PbABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PbBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Pb{PbCaller: PbCaller{contract: contract}, PbTransactor: PbTransactor{contract: contract}, PbFilterer: PbFilterer{contract: contract}}, nil
}

// Pb is an auto generated Go binding around an Ethereum contract.
type Pb struct {
	PbCaller     // Read-only binding to the contract
	PbTransactor // Write-only binding to the contract
	PbFilterer   // Log filterer for contract events
}

// PbCaller is an auto generated read-only Go binding around an Ethereum contract.
type PbCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PbTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PbTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PbFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PbFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PbSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PbSession struct {
	Contract     *Pb               // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PbCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PbCallerSession struct {
	Contract *PbCaller     // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PbTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PbTransactorSession struct {
	Contract     *PbTransactor     // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PbRaw is an auto generated low-level Go binding around an Ethereum contract.
type PbRaw struct {
	Contract *Pb // Generic contract binding to access the raw methods on
}

// PbCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PbCallerRaw struct {
	Contract *PbCaller // Generic read-only contract binding to access the raw methods on
}

// PbTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PbTransactorRaw struct {
	Contract *PbTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPb creates a new instance of Pb, bound to a specific deployed contract.
func NewPb(address common.Address, backend bind.ContractBackend) (*Pb, error) {
	contract, err := bindPb(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pb{PbCaller: PbCaller{contract: contract}, PbTransactor: PbTransactor{contract: contract}, PbFilterer: PbFilterer{contract: contract}}, nil
}

// NewPbCaller creates a new read-only instance of Pb, bound to a specific deployed contract.
func NewPbCaller(address common.Address, caller bind.ContractCaller) (*PbCaller, error) {
	contract, err := bindPb(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PbCaller{contract: contract}, nil
}

// NewPbTransactor creates a new write-only instance of Pb, bound to a specific deployed contract.
func NewPbTransactor(address common.Address, transactor bind.ContractTransactor) (*PbTransactor, error) {
	contract, err := bindPb(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PbTransactor{contract: contract}, nil
}

// NewPbFilterer creates a new log filterer instance of Pb, bound to a specific deployed contract.
func NewPbFilterer(address common.Address, filterer bind.ContractFilterer) (*PbFilterer, error) {
	contract, err := bindPb(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PbFilterer{contract: contract}, nil
}

// bindPb binds a generic wrapper to an already deployed contract.
func bindPb(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PbABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pb *PbRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Pb.Contract.PbCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pb *PbRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pb.Contract.PbTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pb *PbRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pb.Contract.PbTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pb *PbCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Pb.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pb *PbTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pb.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pb *PbTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pb.Contract.contract.Transact(opts, method, params...)
}

// PbRpcAuthorizedWithdrawABI is the input ABI used to generate the binding from.
const PbRpcAuthorizedWithdrawABI = "[]"

// PbRpcAuthorizedWithdrawBin is the compiled bytecode used for deploying new contracts.
const PbRpcAuthorizedWithdrawBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a723058206c7cfee6e802dad3f3c27eb422785db4312b5c014442e1923d73da7e4f654ad30029`

// DeployPbRpcAuthorizedWithdraw deploys a new Ethereum contract, binding an instance of PbRpcAuthorizedWithdraw to it.
func DeployPbRpcAuthorizedWithdraw(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PbRpcAuthorizedWithdraw, error) {
	parsed, err := abi.JSON(strings.NewReader(PbRpcAuthorizedWithdrawABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PbRpcAuthorizedWithdrawBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PbRpcAuthorizedWithdraw{PbRpcAuthorizedWithdrawCaller: PbRpcAuthorizedWithdrawCaller{contract: contract}, PbRpcAuthorizedWithdrawTransactor: PbRpcAuthorizedWithdrawTransactor{contract: contract}, PbRpcAuthorizedWithdrawFilterer: PbRpcAuthorizedWithdrawFilterer{contract: contract}}, nil
}

// PbRpcAuthorizedWithdraw is an auto generated Go binding around an Ethereum contract.
type PbRpcAuthorizedWithdraw struct {
	PbRpcAuthorizedWithdrawCaller     // Read-only binding to the contract
	PbRpcAuthorizedWithdrawTransactor // Write-only binding to the contract
	PbRpcAuthorizedWithdrawFilterer   // Log filterer for contract events
}

// PbRpcAuthorizedWithdrawCaller is an auto generated read-only Go binding around an Ethereum contract.
type PbRpcAuthorizedWithdrawCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PbRpcAuthorizedWithdrawTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PbRpcAuthorizedWithdrawTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PbRpcAuthorizedWithdrawFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PbRpcAuthorizedWithdrawFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PbRpcAuthorizedWithdrawSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PbRpcAuthorizedWithdrawSession struct {
	Contract     *PbRpcAuthorizedWithdraw // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// PbRpcAuthorizedWithdrawCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PbRpcAuthorizedWithdrawCallerSession struct {
	Contract *PbRpcAuthorizedWithdrawCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// PbRpcAuthorizedWithdrawTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PbRpcAuthorizedWithdrawTransactorSession struct {
	Contract     *PbRpcAuthorizedWithdrawTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// PbRpcAuthorizedWithdrawRaw is an auto generated low-level Go binding around an Ethereum contract.
type PbRpcAuthorizedWithdrawRaw struct {
	Contract *PbRpcAuthorizedWithdraw // Generic contract binding to access the raw methods on
}

// PbRpcAuthorizedWithdrawCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PbRpcAuthorizedWithdrawCallerRaw struct {
	Contract *PbRpcAuthorizedWithdrawCaller // Generic read-only contract binding to access the raw methods on
}

// PbRpcAuthorizedWithdrawTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PbRpcAuthorizedWithdrawTransactorRaw struct {
	Contract *PbRpcAuthorizedWithdrawTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPbRpcAuthorizedWithdraw creates a new instance of PbRpcAuthorizedWithdraw, bound to a specific deployed contract.
func NewPbRpcAuthorizedWithdraw(address common.Address, backend bind.ContractBackend) (*PbRpcAuthorizedWithdraw, error) {
	contract, err := bindPbRpcAuthorizedWithdraw(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PbRpcAuthorizedWithdraw{PbRpcAuthorizedWithdrawCaller: PbRpcAuthorizedWithdrawCaller{contract: contract}, PbRpcAuthorizedWithdrawTransactor: PbRpcAuthorizedWithdrawTransactor{contract: contract}, PbRpcAuthorizedWithdrawFilterer: PbRpcAuthorizedWithdrawFilterer{contract: contract}}, nil
}

// NewPbRpcAuthorizedWithdrawCaller creates a new read-only instance of PbRpcAuthorizedWithdraw, bound to a specific deployed contract.
func NewPbRpcAuthorizedWithdrawCaller(address common.Address, caller bind.ContractCaller) (*PbRpcAuthorizedWithdrawCaller, error) {
	contract, err := bindPbRpcAuthorizedWithdraw(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PbRpcAuthorizedWithdrawCaller{contract: contract}, nil
}

// NewPbRpcAuthorizedWithdrawTransactor creates a new write-only instance of PbRpcAuthorizedWithdraw, bound to a specific deployed contract.
func NewPbRpcAuthorizedWithdrawTransactor(address common.Address, transactor bind.ContractTransactor) (*PbRpcAuthorizedWithdrawTransactor, error) {
	contract, err := bindPbRpcAuthorizedWithdraw(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PbRpcAuthorizedWithdrawTransactor{contract: contract}, nil
}

// NewPbRpcAuthorizedWithdrawFilterer creates a new log filterer instance of PbRpcAuthorizedWithdraw, bound to a specific deployed contract.
func NewPbRpcAuthorizedWithdrawFilterer(address common.Address, filterer bind.ContractFilterer) (*PbRpcAuthorizedWithdrawFilterer, error) {
	contract, err := bindPbRpcAuthorizedWithdraw(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PbRpcAuthorizedWithdrawFilterer{contract: contract}, nil
}

// bindPbRpcAuthorizedWithdraw binds a generic wrapper to an already deployed contract.
func bindPbRpcAuthorizedWithdraw(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PbRpcAuthorizedWithdrawABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PbRpcAuthorizedWithdraw *PbRpcAuthorizedWithdrawRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PbRpcAuthorizedWithdraw.Contract.PbRpcAuthorizedWithdrawCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PbRpcAuthorizedWithdraw *PbRpcAuthorizedWithdrawRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PbRpcAuthorizedWithdraw.Contract.PbRpcAuthorizedWithdrawTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PbRpcAuthorizedWithdraw *PbRpcAuthorizedWithdrawRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PbRpcAuthorizedWithdraw.Contract.PbRpcAuthorizedWithdrawTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PbRpcAuthorizedWithdraw *PbRpcAuthorizedWithdrawCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PbRpcAuthorizedWithdraw.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PbRpcAuthorizedWithdraw *PbRpcAuthorizedWithdrawTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PbRpcAuthorizedWithdraw.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PbRpcAuthorizedWithdraw *PbRpcAuthorizedWithdrawTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PbRpcAuthorizedWithdraw.Contract.contract.Transact(opts, method, params...)
}

// PbRpcConditionABI is the input ABI used to generate the binding from.
const PbRpcConditionABI = "[]"

// PbRpcConditionBin is the compiled bytecode used for deploying new contracts.
const PbRpcConditionBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a72305820fc9b20ae5d48d04290643b9adebf50052d9cb6aae2e412af12d21fa442f597ff0029`

// DeployPbRpcCondition deploys a new Ethereum contract, binding an instance of PbRpcCondition to it.
func DeployPbRpcCondition(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PbRpcCondition, error) {
	parsed, err := abi.JSON(strings.NewReader(PbRpcConditionABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PbRpcConditionBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PbRpcCondition{PbRpcConditionCaller: PbRpcConditionCaller{contract: contract}, PbRpcConditionTransactor: PbRpcConditionTransactor{contract: contract}, PbRpcConditionFilterer: PbRpcConditionFilterer{contract: contract}}, nil
}

// PbRpcCondition is an auto generated Go binding around an Ethereum contract.
type PbRpcCondition struct {
	PbRpcConditionCaller     // Read-only binding to the contract
	PbRpcConditionTransactor // Write-only binding to the contract
	PbRpcConditionFilterer   // Log filterer for contract events
}

// PbRpcConditionCaller is an auto generated read-only Go binding around an Ethereum contract.
type PbRpcConditionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PbRpcConditionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PbRpcConditionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PbRpcConditionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PbRpcConditionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PbRpcConditionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PbRpcConditionSession struct {
	Contract     *PbRpcCondition   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PbRpcConditionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PbRpcConditionCallerSession struct {
	Contract *PbRpcConditionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// PbRpcConditionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PbRpcConditionTransactorSession struct {
	Contract     *PbRpcConditionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// PbRpcConditionRaw is an auto generated low-level Go binding around an Ethereum contract.
type PbRpcConditionRaw struct {
	Contract *PbRpcCondition // Generic contract binding to access the raw methods on
}

// PbRpcConditionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PbRpcConditionCallerRaw struct {
	Contract *PbRpcConditionCaller // Generic read-only contract binding to access the raw methods on
}

// PbRpcConditionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PbRpcConditionTransactorRaw struct {
	Contract *PbRpcConditionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPbRpcCondition creates a new instance of PbRpcCondition, bound to a specific deployed contract.
func NewPbRpcCondition(address common.Address, backend bind.ContractBackend) (*PbRpcCondition, error) {
	contract, err := bindPbRpcCondition(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PbRpcCondition{PbRpcConditionCaller: PbRpcConditionCaller{contract: contract}, PbRpcConditionTransactor: PbRpcConditionTransactor{contract: contract}, PbRpcConditionFilterer: PbRpcConditionFilterer{contract: contract}}, nil
}

// NewPbRpcConditionCaller creates a new read-only instance of PbRpcCondition, bound to a specific deployed contract.
func NewPbRpcConditionCaller(address common.Address, caller bind.ContractCaller) (*PbRpcConditionCaller, error) {
	contract, err := bindPbRpcCondition(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PbRpcConditionCaller{contract: contract}, nil
}

// NewPbRpcConditionTransactor creates a new write-only instance of PbRpcCondition, bound to a specific deployed contract.
func NewPbRpcConditionTransactor(address common.Address, transactor bind.ContractTransactor) (*PbRpcConditionTransactor, error) {
	contract, err := bindPbRpcCondition(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PbRpcConditionTransactor{contract: contract}, nil
}

// NewPbRpcConditionFilterer creates a new log filterer instance of PbRpcCondition, bound to a specific deployed contract.
func NewPbRpcConditionFilterer(address common.Address, filterer bind.ContractFilterer) (*PbRpcConditionFilterer, error) {
	contract, err := bindPbRpcCondition(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PbRpcConditionFilterer{contract: contract}, nil
}

// bindPbRpcCondition binds a generic wrapper to an already deployed contract.
func bindPbRpcCondition(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PbRpcConditionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PbRpcCondition *PbRpcConditionRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PbRpcCondition.Contract.PbRpcConditionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PbRpcCondition *PbRpcConditionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PbRpcCondition.Contract.PbRpcConditionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PbRpcCondition *PbRpcConditionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PbRpcCondition.Contract.PbRpcConditionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PbRpcCondition *PbRpcConditionCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PbRpcCondition.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PbRpcCondition *PbRpcConditionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PbRpcCondition.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PbRpcCondition *PbRpcConditionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PbRpcCondition.Contract.contract.Transact(opts, method, params...)
}

// PbRpcConditionGroupABI is the input ABI used to generate the binding from.
const PbRpcConditionGroupABI = "[]"

// PbRpcConditionGroupBin is the compiled bytecode used for deploying new contracts.
const PbRpcConditionGroupBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a723058206e4637ea677c315662ceb9ce82a6e4ba8b834f119e1b2c1c4c3e397eb28eba240029`

// DeployPbRpcConditionGroup deploys a new Ethereum contract, binding an instance of PbRpcConditionGroup to it.
func DeployPbRpcConditionGroup(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PbRpcConditionGroup, error) {
	parsed, err := abi.JSON(strings.NewReader(PbRpcConditionGroupABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PbRpcConditionGroupBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PbRpcConditionGroup{PbRpcConditionGroupCaller: PbRpcConditionGroupCaller{contract: contract}, PbRpcConditionGroupTransactor: PbRpcConditionGroupTransactor{contract: contract}, PbRpcConditionGroupFilterer: PbRpcConditionGroupFilterer{contract: contract}}, nil
}

// PbRpcConditionGroup is an auto generated Go binding around an Ethereum contract.
type PbRpcConditionGroup struct {
	PbRpcConditionGroupCaller     // Read-only binding to the contract
	PbRpcConditionGroupTransactor // Write-only binding to the contract
	PbRpcConditionGroupFilterer   // Log filterer for contract events
}

// PbRpcConditionGroupCaller is an auto generated read-only Go binding around an Ethereum contract.
type PbRpcConditionGroupCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PbRpcConditionGroupTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PbRpcConditionGroupTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PbRpcConditionGroupFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PbRpcConditionGroupFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PbRpcConditionGroupSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PbRpcConditionGroupSession struct {
	Contract     *PbRpcConditionGroup // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// PbRpcConditionGroupCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PbRpcConditionGroupCallerSession struct {
	Contract *PbRpcConditionGroupCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// PbRpcConditionGroupTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PbRpcConditionGroupTransactorSession struct {
	Contract     *PbRpcConditionGroupTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// PbRpcConditionGroupRaw is an auto generated low-level Go binding around an Ethereum contract.
type PbRpcConditionGroupRaw struct {
	Contract *PbRpcConditionGroup // Generic contract binding to access the raw methods on
}

// PbRpcConditionGroupCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PbRpcConditionGroupCallerRaw struct {
	Contract *PbRpcConditionGroupCaller // Generic read-only contract binding to access the raw methods on
}

// PbRpcConditionGroupTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PbRpcConditionGroupTransactorRaw struct {
	Contract *PbRpcConditionGroupTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPbRpcConditionGroup creates a new instance of PbRpcConditionGroup, bound to a specific deployed contract.
func NewPbRpcConditionGroup(address common.Address, backend bind.ContractBackend) (*PbRpcConditionGroup, error) {
	contract, err := bindPbRpcConditionGroup(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PbRpcConditionGroup{PbRpcConditionGroupCaller: PbRpcConditionGroupCaller{contract: contract}, PbRpcConditionGroupTransactor: PbRpcConditionGroupTransactor{contract: contract}, PbRpcConditionGroupFilterer: PbRpcConditionGroupFilterer{contract: contract}}, nil
}

// NewPbRpcConditionGroupCaller creates a new read-only instance of PbRpcConditionGroup, bound to a specific deployed contract.
func NewPbRpcConditionGroupCaller(address common.Address, caller bind.ContractCaller) (*PbRpcConditionGroupCaller, error) {
	contract, err := bindPbRpcConditionGroup(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PbRpcConditionGroupCaller{contract: contract}, nil
}

// NewPbRpcConditionGroupTransactor creates a new write-only instance of PbRpcConditionGroup, bound to a specific deployed contract.
func NewPbRpcConditionGroupTransactor(address common.Address, transactor bind.ContractTransactor) (*PbRpcConditionGroupTransactor, error) {
	contract, err := bindPbRpcConditionGroup(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PbRpcConditionGroupTransactor{contract: contract}, nil
}

// NewPbRpcConditionGroupFilterer creates a new log filterer instance of PbRpcConditionGroup, bound to a specific deployed contract.
func NewPbRpcConditionGroupFilterer(address common.Address, filterer bind.ContractFilterer) (*PbRpcConditionGroupFilterer, error) {
	contract, err := bindPbRpcConditionGroup(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PbRpcConditionGroupFilterer{contract: contract}, nil
}

// bindPbRpcConditionGroup binds a generic wrapper to an already deployed contract.
func bindPbRpcConditionGroup(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PbRpcConditionGroupABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PbRpcConditionGroup *PbRpcConditionGroupRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PbRpcConditionGroup.Contract.PbRpcConditionGroupCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PbRpcConditionGroup *PbRpcConditionGroupRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PbRpcConditionGroup.Contract.PbRpcConditionGroupTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PbRpcConditionGroup *PbRpcConditionGroupRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PbRpcConditionGroup.Contract.PbRpcConditionGroupTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PbRpcConditionGroup *PbRpcConditionGroupCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PbRpcConditionGroup.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PbRpcConditionGroup *PbRpcConditionGroupTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PbRpcConditionGroup.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PbRpcConditionGroup *PbRpcConditionGroupTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PbRpcConditionGroup.Contract.contract.Transact(opts, method, params...)
}

// PbRpcCooperativeWithdrawProofABI is the input ABI used to generate the binding from.
const PbRpcCooperativeWithdrawProofABI = "[]"

// PbRpcCooperativeWithdrawProofBin is the compiled bytecode used for deploying new contracts.
const PbRpcCooperativeWithdrawProofBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a723058206c4a5b3ad604a6d8d25d1c3faed26777ab33b0c3f35bb35da402eede12bc8c050029`

// DeployPbRpcCooperativeWithdrawProof deploys a new Ethereum contract, binding an instance of PbRpcCooperativeWithdrawProof to it.
func DeployPbRpcCooperativeWithdrawProof(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PbRpcCooperativeWithdrawProof, error) {
	parsed, err := abi.JSON(strings.NewReader(PbRpcCooperativeWithdrawProofABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PbRpcCooperativeWithdrawProofBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PbRpcCooperativeWithdrawProof{PbRpcCooperativeWithdrawProofCaller: PbRpcCooperativeWithdrawProofCaller{contract: contract}, PbRpcCooperativeWithdrawProofTransactor: PbRpcCooperativeWithdrawProofTransactor{contract: contract}, PbRpcCooperativeWithdrawProofFilterer: PbRpcCooperativeWithdrawProofFilterer{contract: contract}}, nil
}

// PbRpcCooperativeWithdrawProof is an auto generated Go binding around an Ethereum contract.
type PbRpcCooperativeWithdrawProof struct {
	PbRpcCooperativeWithdrawProofCaller     // Read-only binding to the contract
	PbRpcCooperativeWithdrawProofTransactor // Write-only binding to the contract
	PbRpcCooperativeWithdrawProofFilterer   // Log filterer for contract events
}

// PbRpcCooperativeWithdrawProofCaller is an auto generated read-only Go binding around an Ethereum contract.
type PbRpcCooperativeWithdrawProofCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PbRpcCooperativeWithdrawProofTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PbRpcCooperativeWithdrawProofTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PbRpcCooperativeWithdrawProofFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PbRpcCooperativeWithdrawProofFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PbRpcCooperativeWithdrawProofSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PbRpcCooperativeWithdrawProofSession struct {
	Contract     *PbRpcCooperativeWithdrawProof // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                  // Call options to use throughout this session
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// PbRpcCooperativeWithdrawProofCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PbRpcCooperativeWithdrawProofCallerSession struct {
	Contract *PbRpcCooperativeWithdrawProofCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                        // Call options to use throughout this session
}

// PbRpcCooperativeWithdrawProofTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PbRpcCooperativeWithdrawProofTransactorSession struct {
	Contract     *PbRpcCooperativeWithdrawProofTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                        // Transaction auth options to use throughout this session
}

// PbRpcCooperativeWithdrawProofRaw is an auto generated low-level Go binding around an Ethereum contract.
type PbRpcCooperativeWithdrawProofRaw struct {
	Contract *PbRpcCooperativeWithdrawProof // Generic contract binding to access the raw methods on
}

// PbRpcCooperativeWithdrawProofCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PbRpcCooperativeWithdrawProofCallerRaw struct {
	Contract *PbRpcCooperativeWithdrawProofCaller // Generic read-only contract binding to access the raw methods on
}

// PbRpcCooperativeWithdrawProofTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PbRpcCooperativeWithdrawProofTransactorRaw struct {
	Contract *PbRpcCooperativeWithdrawProofTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPbRpcCooperativeWithdrawProof creates a new instance of PbRpcCooperativeWithdrawProof, bound to a specific deployed contract.
func NewPbRpcCooperativeWithdrawProof(address common.Address, backend bind.ContractBackend) (*PbRpcCooperativeWithdrawProof, error) {
	contract, err := bindPbRpcCooperativeWithdrawProof(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PbRpcCooperativeWithdrawProof{PbRpcCooperativeWithdrawProofCaller: PbRpcCooperativeWithdrawProofCaller{contract: contract}, PbRpcCooperativeWithdrawProofTransactor: PbRpcCooperativeWithdrawProofTransactor{contract: contract}, PbRpcCooperativeWithdrawProofFilterer: PbRpcCooperativeWithdrawProofFilterer{contract: contract}}, nil
}

// NewPbRpcCooperativeWithdrawProofCaller creates a new read-only instance of PbRpcCooperativeWithdrawProof, bound to a specific deployed contract.
func NewPbRpcCooperativeWithdrawProofCaller(address common.Address, caller bind.ContractCaller) (*PbRpcCooperativeWithdrawProofCaller, error) {
	contract, err := bindPbRpcCooperativeWithdrawProof(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PbRpcCooperativeWithdrawProofCaller{contract: contract}, nil
}

// NewPbRpcCooperativeWithdrawProofTransactor creates a new write-only instance of PbRpcCooperativeWithdrawProof, bound to a specific deployed contract.
func NewPbRpcCooperativeWithdrawProofTransactor(address common.Address, transactor bind.ContractTransactor) (*PbRpcCooperativeWithdrawProofTransactor, error) {
	contract, err := bindPbRpcCooperativeWithdrawProof(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PbRpcCooperativeWithdrawProofTransactor{contract: contract}, nil
}

// NewPbRpcCooperativeWithdrawProofFilterer creates a new log filterer instance of PbRpcCooperativeWithdrawProof, bound to a specific deployed contract.
func NewPbRpcCooperativeWithdrawProofFilterer(address common.Address, filterer bind.ContractFilterer) (*PbRpcCooperativeWithdrawProofFilterer, error) {
	contract, err := bindPbRpcCooperativeWithdrawProof(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PbRpcCooperativeWithdrawProofFilterer{contract: contract}, nil
}

// bindPbRpcCooperativeWithdrawProof binds a generic wrapper to an already deployed contract.
func bindPbRpcCooperativeWithdrawProof(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PbRpcCooperativeWithdrawProofABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PbRpcCooperativeWithdrawProof *PbRpcCooperativeWithdrawProofRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PbRpcCooperativeWithdrawProof.Contract.PbRpcCooperativeWithdrawProofCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PbRpcCooperativeWithdrawProof *PbRpcCooperativeWithdrawProofRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PbRpcCooperativeWithdrawProof.Contract.PbRpcCooperativeWithdrawProofTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PbRpcCooperativeWithdrawProof *PbRpcCooperativeWithdrawProofRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PbRpcCooperativeWithdrawProof.Contract.PbRpcCooperativeWithdrawProofTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PbRpcCooperativeWithdrawProof *PbRpcCooperativeWithdrawProofCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PbRpcCooperativeWithdrawProof.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PbRpcCooperativeWithdrawProof *PbRpcCooperativeWithdrawProofTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PbRpcCooperativeWithdrawProof.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PbRpcCooperativeWithdrawProof *PbRpcCooperativeWithdrawProofTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PbRpcCooperativeWithdrawProof.Contract.contract.Transact(opts, method, params...)
}

// PbRpcMultiSignatureABI is the input ABI used to generate the binding from.
const PbRpcMultiSignatureABI = "[]"

// PbRpcMultiSignatureBin is the compiled bytecode used for deploying new contracts.
const PbRpcMultiSignatureBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a723058200a0a9f8cc52ffecd7ee4676730e1d8ca601a0ddab387d699d2d8b954ad6fb0430029`

// DeployPbRpcMultiSignature deploys a new Ethereum contract, binding an instance of PbRpcMultiSignature to it.
func DeployPbRpcMultiSignature(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PbRpcMultiSignature, error) {
	parsed, err := abi.JSON(strings.NewReader(PbRpcMultiSignatureABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PbRpcMultiSignatureBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PbRpcMultiSignature{PbRpcMultiSignatureCaller: PbRpcMultiSignatureCaller{contract: contract}, PbRpcMultiSignatureTransactor: PbRpcMultiSignatureTransactor{contract: contract}, PbRpcMultiSignatureFilterer: PbRpcMultiSignatureFilterer{contract: contract}}, nil
}

// PbRpcMultiSignature is an auto generated Go binding around an Ethereum contract.
type PbRpcMultiSignature struct {
	PbRpcMultiSignatureCaller     // Read-only binding to the contract
	PbRpcMultiSignatureTransactor // Write-only binding to the contract
	PbRpcMultiSignatureFilterer   // Log filterer for contract events
}

// PbRpcMultiSignatureCaller is an auto generated read-only Go binding around an Ethereum contract.
type PbRpcMultiSignatureCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PbRpcMultiSignatureTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PbRpcMultiSignatureTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PbRpcMultiSignatureFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PbRpcMultiSignatureFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PbRpcMultiSignatureSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PbRpcMultiSignatureSession struct {
	Contract     *PbRpcMultiSignature // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// PbRpcMultiSignatureCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PbRpcMultiSignatureCallerSession struct {
	Contract *PbRpcMultiSignatureCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// PbRpcMultiSignatureTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PbRpcMultiSignatureTransactorSession struct {
	Contract     *PbRpcMultiSignatureTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// PbRpcMultiSignatureRaw is an auto generated low-level Go binding around an Ethereum contract.
type PbRpcMultiSignatureRaw struct {
	Contract *PbRpcMultiSignature // Generic contract binding to access the raw methods on
}

// PbRpcMultiSignatureCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PbRpcMultiSignatureCallerRaw struct {
	Contract *PbRpcMultiSignatureCaller // Generic read-only contract binding to access the raw methods on
}

// PbRpcMultiSignatureTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PbRpcMultiSignatureTransactorRaw struct {
	Contract *PbRpcMultiSignatureTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPbRpcMultiSignature creates a new instance of PbRpcMultiSignature, bound to a specific deployed contract.
func NewPbRpcMultiSignature(address common.Address, backend bind.ContractBackend) (*PbRpcMultiSignature, error) {
	contract, err := bindPbRpcMultiSignature(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PbRpcMultiSignature{PbRpcMultiSignatureCaller: PbRpcMultiSignatureCaller{contract: contract}, PbRpcMultiSignatureTransactor: PbRpcMultiSignatureTransactor{contract: contract}, PbRpcMultiSignatureFilterer: PbRpcMultiSignatureFilterer{contract: contract}}, nil
}

// NewPbRpcMultiSignatureCaller creates a new read-only instance of PbRpcMultiSignature, bound to a specific deployed contract.
func NewPbRpcMultiSignatureCaller(address common.Address, caller bind.ContractCaller) (*PbRpcMultiSignatureCaller, error) {
	contract, err := bindPbRpcMultiSignature(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PbRpcMultiSignatureCaller{contract: contract}, nil
}

// NewPbRpcMultiSignatureTransactor creates a new write-only instance of PbRpcMultiSignature, bound to a specific deployed contract.
func NewPbRpcMultiSignatureTransactor(address common.Address, transactor bind.ContractTransactor) (*PbRpcMultiSignatureTransactor, error) {
	contract, err := bindPbRpcMultiSignature(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PbRpcMultiSignatureTransactor{contract: contract}, nil
}

// NewPbRpcMultiSignatureFilterer creates a new log filterer instance of PbRpcMultiSignature, bound to a specific deployed contract.
func NewPbRpcMultiSignatureFilterer(address common.Address, filterer bind.ContractFilterer) (*PbRpcMultiSignatureFilterer, error) {
	contract, err := bindPbRpcMultiSignature(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PbRpcMultiSignatureFilterer{contract: contract}, nil
}

// bindPbRpcMultiSignature binds a generic wrapper to an already deployed contract.
func bindPbRpcMultiSignature(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PbRpcMultiSignatureABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PbRpcMultiSignature *PbRpcMultiSignatureRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PbRpcMultiSignature.Contract.PbRpcMultiSignatureCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PbRpcMultiSignature *PbRpcMultiSignatureRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PbRpcMultiSignature.Contract.PbRpcMultiSignatureTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PbRpcMultiSignature *PbRpcMultiSignatureRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PbRpcMultiSignature.Contract.PbRpcMultiSignatureTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PbRpcMultiSignature *PbRpcMultiSignatureCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PbRpcMultiSignature.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PbRpcMultiSignature *PbRpcMultiSignatureTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PbRpcMultiSignature.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PbRpcMultiSignature *PbRpcMultiSignatureTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PbRpcMultiSignature.Contract.contract.Transact(opts, method, params...)
}

// PbRpcPaymentBooleanAndResolveLogicABI is the input ABI used to generate the binding from.
const PbRpcPaymentBooleanAndResolveLogicABI = "[]"

// PbRpcPaymentBooleanAndResolveLogicBin is the compiled bytecode used for deploying new contracts.
const PbRpcPaymentBooleanAndResolveLogicBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a723058206f330da8bcbd93928206b437247af03dfa5a78244a85e9063f9232a0168a0b8a0029`

// DeployPbRpcPaymentBooleanAndResolveLogic deploys a new Ethereum contract, binding an instance of PbRpcPaymentBooleanAndResolveLogic to it.
func DeployPbRpcPaymentBooleanAndResolveLogic(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PbRpcPaymentBooleanAndResolveLogic, error) {
	parsed, err := abi.JSON(strings.NewReader(PbRpcPaymentBooleanAndResolveLogicABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PbRpcPaymentBooleanAndResolveLogicBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PbRpcPaymentBooleanAndResolveLogic{PbRpcPaymentBooleanAndResolveLogicCaller: PbRpcPaymentBooleanAndResolveLogicCaller{contract: contract}, PbRpcPaymentBooleanAndResolveLogicTransactor: PbRpcPaymentBooleanAndResolveLogicTransactor{contract: contract}, PbRpcPaymentBooleanAndResolveLogicFilterer: PbRpcPaymentBooleanAndResolveLogicFilterer{contract: contract}}, nil
}

// PbRpcPaymentBooleanAndResolveLogic is an auto generated Go binding around an Ethereum contract.
type PbRpcPaymentBooleanAndResolveLogic struct {
	PbRpcPaymentBooleanAndResolveLogicCaller     // Read-only binding to the contract
	PbRpcPaymentBooleanAndResolveLogicTransactor // Write-only binding to the contract
	PbRpcPaymentBooleanAndResolveLogicFilterer   // Log filterer for contract events
}

// PbRpcPaymentBooleanAndResolveLogicCaller is an auto generated read-only Go binding around an Ethereum contract.
type PbRpcPaymentBooleanAndResolveLogicCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PbRpcPaymentBooleanAndResolveLogicTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PbRpcPaymentBooleanAndResolveLogicTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PbRpcPaymentBooleanAndResolveLogicFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PbRpcPaymentBooleanAndResolveLogicFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PbRpcPaymentBooleanAndResolveLogicSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PbRpcPaymentBooleanAndResolveLogicSession struct {
	Contract     *PbRpcPaymentBooleanAndResolveLogic // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                       // Call options to use throughout this session
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// PbRpcPaymentBooleanAndResolveLogicCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PbRpcPaymentBooleanAndResolveLogicCallerSession struct {
	Contract *PbRpcPaymentBooleanAndResolveLogicCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                             // Call options to use throughout this session
}

// PbRpcPaymentBooleanAndResolveLogicTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PbRpcPaymentBooleanAndResolveLogicTransactorSession struct {
	Contract     *PbRpcPaymentBooleanAndResolveLogicTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                             // Transaction auth options to use throughout this session
}

// PbRpcPaymentBooleanAndResolveLogicRaw is an auto generated low-level Go binding around an Ethereum contract.
type PbRpcPaymentBooleanAndResolveLogicRaw struct {
	Contract *PbRpcPaymentBooleanAndResolveLogic // Generic contract binding to access the raw methods on
}

// PbRpcPaymentBooleanAndResolveLogicCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PbRpcPaymentBooleanAndResolveLogicCallerRaw struct {
	Contract *PbRpcPaymentBooleanAndResolveLogicCaller // Generic read-only contract binding to access the raw methods on
}

// PbRpcPaymentBooleanAndResolveLogicTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PbRpcPaymentBooleanAndResolveLogicTransactorRaw struct {
	Contract *PbRpcPaymentBooleanAndResolveLogicTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPbRpcPaymentBooleanAndResolveLogic creates a new instance of PbRpcPaymentBooleanAndResolveLogic, bound to a specific deployed contract.
func NewPbRpcPaymentBooleanAndResolveLogic(address common.Address, backend bind.ContractBackend) (*PbRpcPaymentBooleanAndResolveLogic, error) {
	contract, err := bindPbRpcPaymentBooleanAndResolveLogic(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PbRpcPaymentBooleanAndResolveLogic{PbRpcPaymentBooleanAndResolveLogicCaller: PbRpcPaymentBooleanAndResolveLogicCaller{contract: contract}, PbRpcPaymentBooleanAndResolveLogicTransactor: PbRpcPaymentBooleanAndResolveLogicTransactor{contract: contract}, PbRpcPaymentBooleanAndResolveLogicFilterer: PbRpcPaymentBooleanAndResolveLogicFilterer{contract: contract}}, nil
}

// NewPbRpcPaymentBooleanAndResolveLogicCaller creates a new read-only instance of PbRpcPaymentBooleanAndResolveLogic, bound to a specific deployed contract.
func NewPbRpcPaymentBooleanAndResolveLogicCaller(address common.Address, caller bind.ContractCaller) (*PbRpcPaymentBooleanAndResolveLogicCaller, error) {
	contract, err := bindPbRpcPaymentBooleanAndResolveLogic(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PbRpcPaymentBooleanAndResolveLogicCaller{contract: contract}, nil
}

// NewPbRpcPaymentBooleanAndResolveLogicTransactor creates a new write-only instance of PbRpcPaymentBooleanAndResolveLogic, bound to a specific deployed contract.
func NewPbRpcPaymentBooleanAndResolveLogicTransactor(address common.Address, transactor bind.ContractTransactor) (*PbRpcPaymentBooleanAndResolveLogicTransactor, error) {
	contract, err := bindPbRpcPaymentBooleanAndResolveLogic(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PbRpcPaymentBooleanAndResolveLogicTransactor{contract: contract}, nil
}

// NewPbRpcPaymentBooleanAndResolveLogicFilterer creates a new log filterer instance of PbRpcPaymentBooleanAndResolveLogic, bound to a specific deployed contract.
func NewPbRpcPaymentBooleanAndResolveLogicFilterer(address common.Address, filterer bind.ContractFilterer) (*PbRpcPaymentBooleanAndResolveLogicFilterer, error) {
	contract, err := bindPbRpcPaymentBooleanAndResolveLogic(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PbRpcPaymentBooleanAndResolveLogicFilterer{contract: contract}, nil
}

// bindPbRpcPaymentBooleanAndResolveLogic binds a generic wrapper to an already deployed contract.
func bindPbRpcPaymentBooleanAndResolveLogic(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PbRpcPaymentBooleanAndResolveLogicABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PbRpcPaymentBooleanAndResolveLogic *PbRpcPaymentBooleanAndResolveLogicRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PbRpcPaymentBooleanAndResolveLogic.Contract.PbRpcPaymentBooleanAndResolveLogicCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PbRpcPaymentBooleanAndResolveLogic *PbRpcPaymentBooleanAndResolveLogicRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PbRpcPaymentBooleanAndResolveLogic.Contract.PbRpcPaymentBooleanAndResolveLogicTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PbRpcPaymentBooleanAndResolveLogic *PbRpcPaymentBooleanAndResolveLogicRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PbRpcPaymentBooleanAndResolveLogic.Contract.PbRpcPaymentBooleanAndResolveLogicTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PbRpcPaymentBooleanAndResolveLogic *PbRpcPaymentBooleanAndResolveLogicCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PbRpcPaymentBooleanAndResolveLogic.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PbRpcPaymentBooleanAndResolveLogic *PbRpcPaymentBooleanAndResolveLogicTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PbRpcPaymentBooleanAndResolveLogic.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PbRpcPaymentBooleanAndResolveLogic *PbRpcPaymentBooleanAndResolveLogicTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PbRpcPaymentBooleanAndResolveLogic.Contract.contract.Transact(opts, method, params...)
}

// PbRpcPaymentChannelStateABI is the input ABI used to generate the binding from.
const PbRpcPaymentChannelStateABI = "[]"

// PbRpcPaymentChannelStateBin is the compiled bytecode used for deploying new contracts.
const PbRpcPaymentChannelStateBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a723058200d6bc7f4017144b607bac1e133270386bb0ee73838e6b31b21654decb526282a0029`

// DeployPbRpcPaymentChannelState deploys a new Ethereum contract, binding an instance of PbRpcPaymentChannelState to it.
func DeployPbRpcPaymentChannelState(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PbRpcPaymentChannelState, error) {
	parsed, err := abi.JSON(strings.NewReader(PbRpcPaymentChannelStateABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PbRpcPaymentChannelStateBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PbRpcPaymentChannelState{PbRpcPaymentChannelStateCaller: PbRpcPaymentChannelStateCaller{contract: contract}, PbRpcPaymentChannelStateTransactor: PbRpcPaymentChannelStateTransactor{contract: contract}, PbRpcPaymentChannelStateFilterer: PbRpcPaymentChannelStateFilterer{contract: contract}}, nil
}

// PbRpcPaymentChannelState is an auto generated Go binding around an Ethereum contract.
type PbRpcPaymentChannelState struct {
	PbRpcPaymentChannelStateCaller     // Read-only binding to the contract
	PbRpcPaymentChannelStateTransactor // Write-only binding to the contract
	PbRpcPaymentChannelStateFilterer   // Log filterer for contract events
}

// PbRpcPaymentChannelStateCaller is an auto generated read-only Go binding around an Ethereum contract.
type PbRpcPaymentChannelStateCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PbRpcPaymentChannelStateTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PbRpcPaymentChannelStateTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PbRpcPaymentChannelStateFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PbRpcPaymentChannelStateFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PbRpcPaymentChannelStateSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PbRpcPaymentChannelStateSession struct {
	Contract     *PbRpcPaymentChannelState // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// PbRpcPaymentChannelStateCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PbRpcPaymentChannelStateCallerSession struct {
	Contract *PbRpcPaymentChannelStateCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// PbRpcPaymentChannelStateTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PbRpcPaymentChannelStateTransactorSession struct {
	Contract     *PbRpcPaymentChannelStateTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// PbRpcPaymentChannelStateRaw is an auto generated low-level Go binding around an Ethereum contract.
type PbRpcPaymentChannelStateRaw struct {
	Contract *PbRpcPaymentChannelState // Generic contract binding to access the raw methods on
}

// PbRpcPaymentChannelStateCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PbRpcPaymentChannelStateCallerRaw struct {
	Contract *PbRpcPaymentChannelStateCaller // Generic read-only contract binding to access the raw methods on
}

// PbRpcPaymentChannelStateTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PbRpcPaymentChannelStateTransactorRaw struct {
	Contract *PbRpcPaymentChannelStateTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPbRpcPaymentChannelState creates a new instance of PbRpcPaymentChannelState, bound to a specific deployed contract.
func NewPbRpcPaymentChannelState(address common.Address, backend bind.ContractBackend) (*PbRpcPaymentChannelState, error) {
	contract, err := bindPbRpcPaymentChannelState(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PbRpcPaymentChannelState{PbRpcPaymentChannelStateCaller: PbRpcPaymentChannelStateCaller{contract: contract}, PbRpcPaymentChannelStateTransactor: PbRpcPaymentChannelStateTransactor{contract: contract}, PbRpcPaymentChannelStateFilterer: PbRpcPaymentChannelStateFilterer{contract: contract}}, nil
}

// NewPbRpcPaymentChannelStateCaller creates a new read-only instance of PbRpcPaymentChannelState, bound to a specific deployed contract.
func NewPbRpcPaymentChannelStateCaller(address common.Address, caller bind.ContractCaller) (*PbRpcPaymentChannelStateCaller, error) {
	contract, err := bindPbRpcPaymentChannelState(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PbRpcPaymentChannelStateCaller{contract: contract}, nil
}

// NewPbRpcPaymentChannelStateTransactor creates a new write-only instance of PbRpcPaymentChannelState, bound to a specific deployed contract.
func NewPbRpcPaymentChannelStateTransactor(address common.Address, transactor bind.ContractTransactor) (*PbRpcPaymentChannelStateTransactor, error) {
	contract, err := bindPbRpcPaymentChannelState(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PbRpcPaymentChannelStateTransactor{contract: contract}, nil
}

// NewPbRpcPaymentChannelStateFilterer creates a new log filterer instance of PbRpcPaymentChannelState, bound to a specific deployed contract.
func NewPbRpcPaymentChannelStateFilterer(address common.Address, filterer bind.ContractFilterer) (*PbRpcPaymentChannelStateFilterer, error) {
	contract, err := bindPbRpcPaymentChannelState(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PbRpcPaymentChannelStateFilterer{contract: contract}, nil
}

// bindPbRpcPaymentChannelState binds a generic wrapper to an already deployed contract.
func bindPbRpcPaymentChannelState(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PbRpcPaymentChannelStateABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PbRpcPaymentChannelState *PbRpcPaymentChannelStateRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PbRpcPaymentChannelState.Contract.PbRpcPaymentChannelStateCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PbRpcPaymentChannelState *PbRpcPaymentChannelStateRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PbRpcPaymentChannelState.Contract.PbRpcPaymentChannelStateTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PbRpcPaymentChannelState *PbRpcPaymentChannelStateRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PbRpcPaymentChannelState.Contract.PbRpcPaymentChannelStateTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PbRpcPaymentChannelState *PbRpcPaymentChannelStateCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PbRpcPaymentChannelState.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PbRpcPaymentChannelState *PbRpcPaymentChannelStateTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PbRpcPaymentChannelState.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PbRpcPaymentChannelState *PbRpcPaymentChannelStateTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PbRpcPaymentChannelState.Contract.contract.Transact(opts, method, params...)
}

// PbRpcStateDepositMapEntryABI is the input ABI used to generate the binding from.
const PbRpcStateDepositMapEntryABI = "[]"

// PbRpcStateDepositMapEntryBin is the compiled bytecode used for deploying new contracts.
const PbRpcStateDepositMapEntryBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a723058206cd43dec5e3ee180a77ddca2af33fcdec764eb53fd954d08056aff27aba8489a0029`

// DeployPbRpcStateDepositMapEntry deploys a new Ethereum contract, binding an instance of PbRpcStateDepositMapEntry to it.
func DeployPbRpcStateDepositMapEntry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PbRpcStateDepositMapEntry, error) {
	parsed, err := abi.JSON(strings.NewReader(PbRpcStateDepositMapEntryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PbRpcStateDepositMapEntryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PbRpcStateDepositMapEntry{PbRpcStateDepositMapEntryCaller: PbRpcStateDepositMapEntryCaller{contract: contract}, PbRpcStateDepositMapEntryTransactor: PbRpcStateDepositMapEntryTransactor{contract: contract}, PbRpcStateDepositMapEntryFilterer: PbRpcStateDepositMapEntryFilterer{contract: contract}}, nil
}

// PbRpcStateDepositMapEntry is an auto generated Go binding around an Ethereum contract.
type PbRpcStateDepositMapEntry struct {
	PbRpcStateDepositMapEntryCaller     // Read-only binding to the contract
	PbRpcStateDepositMapEntryTransactor // Write-only binding to the contract
	PbRpcStateDepositMapEntryFilterer   // Log filterer for contract events
}

// PbRpcStateDepositMapEntryCaller is an auto generated read-only Go binding around an Ethereum contract.
type PbRpcStateDepositMapEntryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PbRpcStateDepositMapEntryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PbRpcStateDepositMapEntryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PbRpcStateDepositMapEntryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PbRpcStateDepositMapEntryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PbRpcStateDepositMapEntrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PbRpcStateDepositMapEntrySession struct {
	Contract     *PbRpcStateDepositMapEntry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// PbRpcStateDepositMapEntryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PbRpcStateDepositMapEntryCallerSession struct {
	Contract *PbRpcStateDepositMapEntryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// PbRpcStateDepositMapEntryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PbRpcStateDepositMapEntryTransactorSession struct {
	Contract     *PbRpcStateDepositMapEntryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// PbRpcStateDepositMapEntryRaw is an auto generated low-level Go binding around an Ethereum contract.
type PbRpcStateDepositMapEntryRaw struct {
	Contract *PbRpcStateDepositMapEntry // Generic contract binding to access the raw methods on
}

// PbRpcStateDepositMapEntryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PbRpcStateDepositMapEntryCallerRaw struct {
	Contract *PbRpcStateDepositMapEntryCaller // Generic read-only contract binding to access the raw methods on
}

// PbRpcStateDepositMapEntryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PbRpcStateDepositMapEntryTransactorRaw struct {
	Contract *PbRpcStateDepositMapEntryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPbRpcStateDepositMapEntry creates a new instance of PbRpcStateDepositMapEntry, bound to a specific deployed contract.
func NewPbRpcStateDepositMapEntry(address common.Address, backend bind.ContractBackend) (*PbRpcStateDepositMapEntry, error) {
	contract, err := bindPbRpcStateDepositMapEntry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PbRpcStateDepositMapEntry{PbRpcStateDepositMapEntryCaller: PbRpcStateDepositMapEntryCaller{contract: contract}, PbRpcStateDepositMapEntryTransactor: PbRpcStateDepositMapEntryTransactor{contract: contract}, PbRpcStateDepositMapEntryFilterer: PbRpcStateDepositMapEntryFilterer{contract: contract}}, nil
}

// NewPbRpcStateDepositMapEntryCaller creates a new read-only instance of PbRpcStateDepositMapEntry, bound to a specific deployed contract.
func NewPbRpcStateDepositMapEntryCaller(address common.Address, caller bind.ContractCaller) (*PbRpcStateDepositMapEntryCaller, error) {
	contract, err := bindPbRpcStateDepositMapEntry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PbRpcStateDepositMapEntryCaller{contract: contract}, nil
}

// NewPbRpcStateDepositMapEntryTransactor creates a new write-only instance of PbRpcStateDepositMapEntry, bound to a specific deployed contract.
func NewPbRpcStateDepositMapEntryTransactor(address common.Address, transactor bind.ContractTransactor) (*PbRpcStateDepositMapEntryTransactor, error) {
	contract, err := bindPbRpcStateDepositMapEntry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PbRpcStateDepositMapEntryTransactor{contract: contract}, nil
}

// NewPbRpcStateDepositMapEntryFilterer creates a new log filterer instance of PbRpcStateDepositMapEntry, bound to a specific deployed contract.
func NewPbRpcStateDepositMapEntryFilterer(address common.Address, filterer bind.ContractFilterer) (*PbRpcStateDepositMapEntryFilterer, error) {
	contract, err := bindPbRpcStateDepositMapEntry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PbRpcStateDepositMapEntryFilterer{contract: contract}, nil
}

// bindPbRpcStateDepositMapEntry binds a generic wrapper to an already deployed contract.
func bindPbRpcStateDepositMapEntry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PbRpcStateDepositMapEntryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PbRpcStateDepositMapEntry *PbRpcStateDepositMapEntryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PbRpcStateDepositMapEntry.Contract.PbRpcStateDepositMapEntryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PbRpcStateDepositMapEntry *PbRpcStateDepositMapEntryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PbRpcStateDepositMapEntry.Contract.PbRpcStateDepositMapEntryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PbRpcStateDepositMapEntry *PbRpcStateDepositMapEntryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PbRpcStateDepositMapEntry.Contract.PbRpcStateDepositMapEntryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PbRpcStateDepositMapEntry *PbRpcStateDepositMapEntryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PbRpcStateDepositMapEntry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PbRpcStateDepositMapEntry *PbRpcStateDepositMapEntryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PbRpcStateDepositMapEntry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PbRpcStateDepositMapEntry *PbRpcStateDepositMapEntryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PbRpcStateDepositMapEntry.Contract.contract.Transact(opts, method, params...)
}

// PbRpcStateProofABI is the input ABI used to generate the binding from.
const PbRpcStateProofABI = "[]"

// PbRpcStateProofBin is the compiled bytecode used for deploying new contracts.
const PbRpcStateProofBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a723058208f7a6cdc8c7accf32f8abd8d3254eeeef85cf50494554aa9abedaf40f2d3db220029`

// DeployPbRpcStateProof deploys a new Ethereum contract, binding an instance of PbRpcStateProof to it.
func DeployPbRpcStateProof(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PbRpcStateProof, error) {
	parsed, err := abi.JSON(strings.NewReader(PbRpcStateProofABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PbRpcStateProofBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PbRpcStateProof{PbRpcStateProofCaller: PbRpcStateProofCaller{contract: contract}, PbRpcStateProofTransactor: PbRpcStateProofTransactor{contract: contract}, PbRpcStateProofFilterer: PbRpcStateProofFilterer{contract: contract}}, nil
}

// PbRpcStateProof is an auto generated Go binding around an Ethereum contract.
type PbRpcStateProof struct {
	PbRpcStateProofCaller     // Read-only binding to the contract
	PbRpcStateProofTransactor // Write-only binding to the contract
	PbRpcStateProofFilterer   // Log filterer for contract events
}

// PbRpcStateProofCaller is an auto generated read-only Go binding around an Ethereum contract.
type PbRpcStateProofCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PbRpcStateProofTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PbRpcStateProofTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PbRpcStateProofFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PbRpcStateProofFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PbRpcStateProofSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PbRpcStateProofSession struct {
	Contract     *PbRpcStateProof  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PbRpcStateProofCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PbRpcStateProofCallerSession struct {
	Contract *PbRpcStateProofCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// PbRpcStateProofTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PbRpcStateProofTransactorSession struct {
	Contract     *PbRpcStateProofTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// PbRpcStateProofRaw is an auto generated low-level Go binding around an Ethereum contract.
type PbRpcStateProofRaw struct {
	Contract *PbRpcStateProof // Generic contract binding to access the raw methods on
}

// PbRpcStateProofCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PbRpcStateProofCallerRaw struct {
	Contract *PbRpcStateProofCaller // Generic read-only contract binding to access the raw methods on
}

// PbRpcStateProofTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PbRpcStateProofTransactorRaw struct {
	Contract *PbRpcStateProofTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPbRpcStateProof creates a new instance of PbRpcStateProof, bound to a specific deployed contract.
func NewPbRpcStateProof(address common.Address, backend bind.ContractBackend) (*PbRpcStateProof, error) {
	contract, err := bindPbRpcStateProof(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PbRpcStateProof{PbRpcStateProofCaller: PbRpcStateProofCaller{contract: contract}, PbRpcStateProofTransactor: PbRpcStateProofTransactor{contract: contract}, PbRpcStateProofFilterer: PbRpcStateProofFilterer{contract: contract}}, nil
}

// NewPbRpcStateProofCaller creates a new read-only instance of PbRpcStateProof, bound to a specific deployed contract.
func NewPbRpcStateProofCaller(address common.Address, caller bind.ContractCaller) (*PbRpcStateProofCaller, error) {
	contract, err := bindPbRpcStateProof(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PbRpcStateProofCaller{contract: contract}, nil
}

// NewPbRpcStateProofTransactor creates a new write-only instance of PbRpcStateProof, bound to a specific deployed contract.
func NewPbRpcStateProofTransactor(address common.Address, transactor bind.ContractTransactor) (*PbRpcStateProofTransactor, error) {
	contract, err := bindPbRpcStateProof(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PbRpcStateProofTransactor{contract: contract}, nil
}

// NewPbRpcStateProofFilterer creates a new log filterer instance of PbRpcStateProof, bound to a specific deployed contract.
func NewPbRpcStateProofFilterer(address common.Address, filterer bind.ContractFilterer) (*PbRpcStateProofFilterer, error) {
	contract, err := bindPbRpcStateProof(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PbRpcStateProofFilterer{contract: contract}, nil
}

// bindPbRpcStateProof binds a generic wrapper to an already deployed contract.
func bindPbRpcStateProof(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PbRpcStateProofABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PbRpcStateProof *PbRpcStateProofRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PbRpcStateProof.Contract.PbRpcStateProofCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PbRpcStateProof *PbRpcStateProofRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PbRpcStateProof.Contract.PbRpcStateProofTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PbRpcStateProof *PbRpcStateProofRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PbRpcStateProof.Contract.PbRpcStateProofTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PbRpcStateProof *PbRpcStateProofCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PbRpcStateProof.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PbRpcStateProof *PbRpcStateProofTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PbRpcStateProof.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PbRpcStateProof *PbRpcStateProofTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PbRpcStateProof.Contract.contract.Transact(opts, method, params...)
}

// PbRpcTransferMapEntryABI is the input ABI used to generate the binding from.
const PbRpcTransferMapEntryABI = "[]"

// PbRpcTransferMapEntryBin is the compiled bytecode used for deploying new contracts.
const PbRpcTransferMapEntryBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a723058206d4e04570e3333872151eaf1e6448f9d5db67d64977a29c98735ecdab295fc620029`

// DeployPbRpcTransferMapEntry deploys a new Ethereum contract, binding an instance of PbRpcTransferMapEntry to it.
func DeployPbRpcTransferMapEntry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PbRpcTransferMapEntry, error) {
	parsed, err := abi.JSON(strings.NewReader(PbRpcTransferMapEntryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PbRpcTransferMapEntryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PbRpcTransferMapEntry{PbRpcTransferMapEntryCaller: PbRpcTransferMapEntryCaller{contract: contract}, PbRpcTransferMapEntryTransactor: PbRpcTransferMapEntryTransactor{contract: contract}, PbRpcTransferMapEntryFilterer: PbRpcTransferMapEntryFilterer{contract: contract}}, nil
}

// PbRpcTransferMapEntry is an auto generated Go binding around an Ethereum contract.
type PbRpcTransferMapEntry struct {
	PbRpcTransferMapEntryCaller     // Read-only binding to the contract
	PbRpcTransferMapEntryTransactor // Write-only binding to the contract
	PbRpcTransferMapEntryFilterer   // Log filterer for contract events
}

// PbRpcTransferMapEntryCaller is an auto generated read-only Go binding around an Ethereum contract.
type PbRpcTransferMapEntryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PbRpcTransferMapEntryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PbRpcTransferMapEntryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PbRpcTransferMapEntryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PbRpcTransferMapEntryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PbRpcTransferMapEntrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PbRpcTransferMapEntrySession struct {
	Contract     *PbRpcTransferMapEntry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// PbRpcTransferMapEntryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PbRpcTransferMapEntryCallerSession struct {
	Contract *PbRpcTransferMapEntryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// PbRpcTransferMapEntryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PbRpcTransferMapEntryTransactorSession struct {
	Contract     *PbRpcTransferMapEntryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// PbRpcTransferMapEntryRaw is an auto generated low-level Go binding around an Ethereum contract.
type PbRpcTransferMapEntryRaw struct {
	Contract *PbRpcTransferMapEntry // Generic contract binding to access the raw methods on
}

// PbRpcTransferMapEntryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PbRpcTransferMapEntryCallerRaw struct {
	Contract *PbRpcTransferMapEntryCaller // Generic read-only contract binding to access the raw methods on
}

// PbRpcTransferMapEntryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PbRpcTransferMapEntryTransactorRaw struct {
	Contract *PbRpcTransferMapEntryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPbRpcTransferMapEntry creates a new instance of PbRpcTransferMapEntry, bound to a specific deployed contract.
func NewPbRpcTransferMapEntry(address common.Address, backend bind.ContractBackend) (*PbRpcTransferMapEntry, error) {
	contract, err := bindPbRpcTransferMapEntry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PbRpcTransferMapEntry{PbRpcTransferMapEntryCaller: PbRpcTransferMapEntryCaller{contract: contract}, PbRpcTransferMapEntryTransactor: PbRpcTransferMapEntryTransactor{contract: contract}, PbRpcTransferMapEntryFilterer: PbRpcTransferMapEntryFilterer{contract: contract}}, nil
}

// NewPbRpcTransferMapEntryCaller creates a new read-only instance of PbRpcTransferMapEntry, bound to a specific deployed contract.
func NewPbRpcTransferMapEntryCaller(address common.Address, caller bind.ContractCaller) (*PbRpcTransferMapEntryCaller, error) {
	contract, err := bindPbRpcTransferMapEntry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PbRpcTransferMapEntryCaller{contract: contract}, nil
}

// NewPbRpcTransferMapEntryTransactor creates a new write-only instance of PbRpcTransferMapEntry, bound to a specific deployed contract.
func NewPbRpcTransferMapEntryTransactor(address common.Address, transactor bind.ContractTransactor) (*PbRpcTransferMapEntryTransactor, error) {
	contract, err := bindPbRpcTransferMapEntry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PbRpcTransferMapEntryTransactor{contract: contract}, nil
}

// NewPbRpcTransferMapEntryFilterer creates a new log filterer instance of PbRpcTransferMapEntry, bound to a specific deployed contract.
func NewPbRpcTransferMapEntryFilterer(address common.Address, filterer bind.ContractFilterer) (*PbRpcTransferMapEntryFilterer, error) {
	contract, err := bindPbRpcTransferMapEntry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PbRpcTransferMapEntryFilterer{contract: contract}, nil
}

// bindPbRpcTransferMapEntry binds a generic wrapper to an already deployed contract.
func bindPbRpcTransferMapEntry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PbRpcTransferMapEntryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PbRpcTransferMapEntry *PbRpcTransferMapEntryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PbRpcTransferMapEntry.Contract.PbRpcTransferMapEntryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PbRpcTransferMapEntry *PbRpcTransferMapEntryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PbRpcTransferMapEntry.Contract.PbRpcTransferMapEntryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PbRpcTransferMapEntry *PbRpcTransferMapEntryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PbRpcTransferMapEntry.Contract.PbRpcTransferMapEntryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PbRpcTransferMapEntry *PbRpcTransferMapEntryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PbRpcTransferMapEntry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PbRpcTransferMapEntry *PbRpcTransferMapEntryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PbRpcTransferMapEntry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PbRpcTransferMapEntry *PbRpcTransferMapEntryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PbRpcTransferMapEntry.Contract.contract.Transact(opts, method, params...)
}
