// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package resolver

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

// VirtContractResolverABI is the input ABI used to generate the binding from.
const VirtContractResolverABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"virt\",\"type\":\"bytes32\"}],\"name\":\"resolve\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_code\",\"type\":\"bytes\"},{\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"deploy\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"virtAddr\",\"type\":\"bytes32\"}],\"name\":\"Deploy\",\"type\":\"event\"}]"

// VirtContractResolverBin is the compiled bytecode used for deploying new contracts.
const VirtContractResolverBin = `0x608060405234801561001057600080fd5b5061026d806100206000396000f30060806040526004361061004b5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416635c23bdf581146100505780639c4ae2d014610091575b600080fd5b34801561005c57600080fd5b506100686004356100c9565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b34801561009d57600080fd5b506100b5602460048035828101929101359035610122565b604080519115158252519081900360200190f35b60008181526020819052604081205473ffffffffffffffffffffffffffffffffffffffff1615156100f957600080fd5b5060009081526020819052604090205473ffffffffffffffffffffffffffffffffffffffff1690565b6000806060600086868660405180848480828437820191505082815260200193505050506040518091039020925086868080601f016020809104026020016040519081016040528093929190818152602001838380828437505050600087815260208190526040902054939550505073ffffffffffffffffffffffffffffffffffffffff9091161590506101b557600080fd5b8151602083016000f060008481526020818152604091829020805473ffffffffffffffffffffffffffffffffffffffff191673ffffffffffffffffffffffffffffffffffffffff8516179055815186815291519293507f149208daa30a9306858cc9c171c3510e0e50ab5d59ed2027a37a728430dd02e492918290030190a150600196955050505050505600a165627a7a72305820af7f128f068c34fd3a89220fef76f7eddb0ea9868225cb2e7021d3f22f9769820029`

// DeployVirtContractResolver deploys a new Ethereum contract, binding an instance of VirtContractResolver to it.
func DeployVirtContractResolver(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *VirtContractResolver, error) {
	parsed, err := abi.JSON(strings.NewReader(VirtContractResolverABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(VirtContractResolverBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VirtContractResolver{VirtContractResolverCaller: VirtContractResolverCaller{contract: contract}, VirtContractResolverTransactor: VirtContractResolverTransactor{contract: contract}, VirtContractResolverFilterer: VirtContractResolverFilterer{contract: contract}}, nil
}

// VirtContractResolver is an auto generated Go binding around an Ethereum contract.
type VirtContractResolver struct {
	VirtContractResolverCaller     // Read-only binding to the contract
	VirtContractResolverTransactor // Write-only binding to the contract
	VirtContractResolverFilterer   // Log filterer for contract events
}

// VirtContractResolverCaller is an auto generated read-only Go binding around an Ethereum contract.
type VirtContractResolverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VirtContractResolverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VirtContractResolverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VirtContractResolverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VirtContractResolverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VirtContractResolverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VirtContractResolverSession struct {
	Contract     *VirtContractResolver // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// VirtContractResolverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VirtContractResolverCallerSession struct {
	Contract *VirtContractResolverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// VirtContractResolverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VirtContractResolverTransactorSession struct {
	Contract     *VirtContractResolverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// VirtContractResolverRaw is an auto generated low-level Go binding around an Ethereum contract.
type VirtContractResolverRaw struct {
	Contract *VirtContractResolver // Generic contract binding to access the raw methods on
}

// VirtContractResolverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VirtContractResolverCallerRaw struct {
	Contract *VirtContractResolverCaller // Generic read-only contract binding to access the raw methods on
}

// VirtContractResolverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VirtContractResolverTransactorRaw struct {
	Contract *VirtContractResolverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVirtContractResolver creates a new instance of VirtContractResolver, bound to a specific deployed contract.
func NewVirtContractResolver(address common.Address, backend bind.ContractBackend) (*VirtContractResolver, error) {
	contract, err := bindVirtContractResolver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VirtContractResolver{VirtContractResolverCaller: VirtContractResolverCaller{contract: contract}, VirtContractResolverTransactor: VirtContractResolverTransactor{contract: contract}, VirtContractResolverFilterer: VirtContractResolverFilterer{contract: contract}}, nil
}

// NewVirtContractResolverCaller creates a new read-only instance of VirtContractResolver, bound to a specific deployed contract.
func NewVirtContractResolverCaller(address common.Address, caller bind.ContractCaller) (*VirtContractResolverCaller, error) {
	contract, err := bindVirtContractResolver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VirtContractResolverCaller{contract: contract}, nil
}

// NewVirtContractResolverTransactor creates a new write-only instance of VirtContractResolver, bound to a specific deployed contract.
func NewVirtContractResolverTransactor(address common.Address, transactor bind.ContractTransactor) (*VirtContractResolverTransactor, error) {
	contract, err := bindVirtContractResolver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VirtContractResolverTransactor{contract: contract}, nil
}

// NewVirtContractResolverFilterer creates a new log filterer instance of VirtContractResolver, bound to a specific deployed contract.
func NewVirtContractResolverFilterer(address common.Address, filterer bind.ContractFilterer) (*VirtContractResolverFilterer, error) {
	contract, err := bindVirtContractResolver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VirtContractResolverFilterer{contract: contract}, nil
}

// bindVirtContractResolver binds a generic wrapper to an already deployed contract.
func bindVirtContractResolver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VirtContractResolverABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VirtContractResolver *VirtContractResolverRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _VirtContractResolver.Contract.VirtContractResolverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VirtContractResolver *VirtContractResolverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VirtContractResolver.Contract.VirtContractResolverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VirtContractResolver *VirtContractResolverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VirtContractResolver.Contract.VirtContractResolverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VirtContractResolver *VirtContractResolverCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _VirtContractResolver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VirtContractResolver *VirtContractResolverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VirtContractResolver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VirtContractResolver *VirtContractResolverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VirtContractResolver.Contract.contract.Transact(opts, method, params...)
}

// Resolve is a free data retrieval call binding the contract method 0x5c23bdf5.
//
// Solidity: function resolve(virt bytes32) constant returns(address)
func (_VirtContractResolver *VirtContractResolverCaller) Resolve(opts *bind.CallOpts, virt [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _VirtContractResolver.contract.Call(opts, out, "resolve", virt)
	return *ret0, err
}

// Resolve is a free data retrieval call binding the contract method 0x5c23bdf5.
//
// Solidity: function resolve(virt bytes32) constant returns(address)
func (_VirtContractResolver *VirtContractResolverSession) Resolve(virt [32]byte) (common.Address, error) {
	return _VirtContractResolver.Contract.Resolve(&_VirtContractResolver.CallOpts, virt)
}

// Resolve is a free data retrieval call binding the contract method 0x5c23bdf5.
//
// Solidity: function resolve(virt bytes32) constant returns(address)
func (_VirtContractResolver *VirtContractResolverCallerSession) Resolve(virt [32]byte) (common.Address, error) {
	return _VirtContractResolver.Contract.Resolve(&_VirtContractResolver.CallOpts, virt)
}

// Deploy is a paid mutator transaction binding the contract method 0x9c4ae2d0.
//
// Solidity: function deploy(_code bytes, _nonce uint256) returns(bool)
func (_VirtContractResolver *VirtContractResolverTransactor) Deploy(opts *bind.TransactOpts, _code []byte, _nonce *big.Int) (*types.Transaction, error) {
	return _VirtContractResolver.contract.Transact(opts, "deploy", _code, _nonce)
}

// Deploy is a paid mutator transaction binding the contract method 0x9c4ae2d0.
//
// Solidity: function deploy(_code bytes, _nonce uint256) returns(bool)
func (_VirtContractResolver *VirtContractResolverSession) Deploy(_code []byte, _nonce *big.Int) (*types.Transaction, error) {
	return _VirtContractResolver.Contract.Deploy(&_VirtContractResolver.TransactOpts, _code, _nonce)
}

// Deploy is a paid mutator transaction binding the contract method 0x9c4ae2d0.
//
// Solidity: function deploy(_code bytes, _nonce uint256) returns(bool)
func (_VirtContractResolver *VirtContractResolverTransactorSession) Deploy(_code []byte, _nonce *big.Int) (*types.Transaction, error) {
	return _VirtContractResolver.Contract.Deploy(&_VirtContractResolver.TransactOpts, _code, _nonce)
}

// VirtContractResolverDeployIterator is returned from FilterDeploy and is used to iterate over the raw logs and unpacked data for Deploy events raised by the VirtContractResolver contract.
type VirtContractResolverDeployIterator struct {
	Event *VirtContractResolverDeploy // Event containing the contract specifics and raw log

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
func (it *VirtContractResolverDeployIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VirtContractResolverDeploy)
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
		it.Event = new(VirtContractResolverDeploy)
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
func (it *VirtContractResolverDeployIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VirtContractResolverDeployIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VirtContractResolverDeploy represents a Deploy event raised by the VirtContractResolver contract.
type VirtContractResolverDeploy struct {
	VirtAddr [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDeploy is a free log retrieval operation binding the contract event 0x149208daa30a9306858cc9c171c3510e0e50ab5d59ed2027a37a728430dd02e4.
//
// Solidity: e Deploy(virtAddr bytes32)
func (_VirtContractResolver *VirtContractResolverFilterer) FilterDeploy(opts *bind.FilterOpts) (*VirtContractResolverDeployIterator, error) {

	logs, sub, err := _VirtContractResolver.contract.FilterLogs(opts, "Deploy")
	if err != nil {
		return nil, err
	}
	return &VirtContractResolverDeployIterator{contract: _VirtContractResolver.contract, event: "Deploy", logs: logs, sub: sub}, nil
}

// WatchDeploy is a free log subscription operation binding the contract event 0x149208daa30a9306858cc9c171c3510e0e50ab5d59ed2027a37a728430dd02e4.
//
// Solidity: e Deploy(virtAddr bytes32)
func (_VirtContractResolver *VirtContractResolverFilterer) WatchDeploy(opts *bind.WatchOpts, sink chan<- *VirtContractResolverDeploy) (event.Subscription, error) {

	logs, sub, err := _VirtContractResolver.contract.WatchLogs(opts, "Deploy")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VirtContractResolverDeploy)
				if err := _VirtContractResolver.contract.UnpackLog(event, "Deploy", log); err != nil {
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
