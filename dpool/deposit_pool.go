// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dpool

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

// DepositPoolABI is the input ABI used to generate the binding from.
const DepositPoolABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_authWithdraw\",\"type\":\"bytes\"},{\"name\":\"_signature\",\"type\":\"bytes\"},{\"name\":\"_channelId\",\"type\":\"uint256\"}],\"name\":\"authorizedWithdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_tokenContract\",\"type\":\"address\"},{\"name\":\"_tokenType\",\"type\":\"uint256\"}],\"name\":\"withdrawERCToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"},{\"name\":\"_tokenContract\",\"type\":\"address\"}],\"name\":\"getRemainingBalance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_receipient\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_receipient\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_tokenContract\",\"type\":\"address\"},{\"name\":\"_tokenType\",\"type\":\"uint256\"}],\"name\":\"depositERCToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// DepositPoolBin is the compiled bytecode used for deploying new contracts.
const DepositPoolBin = `0x608060405234801561001057600080fd5b5061157a806100206000396000f30060806040526004361061005e5763ffffffff60e060020a60003504166306fa40f381146100635780632e1a7d4d146100fe578063783202581461011657806381ae85411461013d578063f340fa0114610176578063fcc3e9911461018a575b600080fd5b34801561006f57600080fd5b506040805160206004803580820135601f81018490048402850184019095528484526100fc94369492936024939284019190819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a99988101979196509182019450925082915084018382808284375094975050933594506101b89350505050565b005b34801561010a57600080fd5b506100fc600435610606565b34801561012257600080fd5b506100fc600435600160a060020a036024351660443561069f565b34801561014957600080fd5b50610164600160a060020a0360043581169060243516610788565b60408051918252519081900360200190f35b6100fc600160a060020a03600435166107b1565b34801561019657600080fd5b506100fc600160a060020a03600435811690602435906044351660643561081d565b6101c061149d565b6101c86114bf565b6000806000806000806000806101dd8c6108f3565b99506101e88d610918565b98508c6040518082805190602001908083835b6020831061021a5780518252601f1990920191602091820191016101fb565b51815160001960209485036101000a01908116901991909116179052604080519490920184900384207f19457468657265756d205369676e6564204d6573736167653a0a3332000000008552601c8501819052825194859003603c019094206000858152600190925291902054929c509a505060ff1615915061029e905057600080fd5b895151895151146102ae57600080fd5b60c089015115806102c3575060018960c00151145b15156102ce57600080fd5b600088815260016020819052604091829020805460ff191682179055908a015160a08b015160c08c01519198509650945092505b8951518310156105f757602089015180518490811061031d57fe5b9060200190602002015191508160001415610337576105ec565b8951805160019189918690811061034a57fe5b906020019060200201518c602001518681518110151561036657fe5b906020019060200201518d604001518781518110151561038257fe5b60209081029091018101516040805160008082528185018084529790975260ff9095168582015260608501939093526080840152905160a0808401949293601f19830193908390039091019190865af11580156103e3573d6000803e3d6000fd5b50505060206040510351905080600160a060020a031689600001518481518110151561040b57fe5b60209081029091010151600160a060020a03161461042857600080fd5b600160a060020a038082166000908152602081815260408083209389168352929052205482111561045857600080fd5b600160a060020a038082166000908152602081815260408083209389168352929052205461048c908363ffffffff61093516565b600160a060020a03828116600090815260208181526040808320938a16835292905220558315156105385785600160a060020a0316636e553f65838d846040518463ffffffff1660e060020a0281526004018083815260200182600160a060020a0316600160a060020a03168152602001925050506000604051808303818588803b15801561051a57600080fd5b505af115801561052e573d6000803e3d6000fd5b50505050506105f7565b60018414156105ea5761055b600160a060020a038616878463ffffffff61094716565b604080517f235edd9a000000000000000000000000000000000000000000000000000000008152600481018d9052600160a060020a0383811660248301526044820185905291519188169163235edd9a9160648082019260009290919082900301818387803b1580156105cd57600080fd5b505af11580156105e1573d6000803e3d6000fd5b505050506105f7565bfe5b600190920191610302565b50505050505050505050505050565b3360009081526020818152604080832083805290915290205481111561062b57600080fd5b33600090815260208181526040808320838052909152902054610654908263ffffffff61093516565b33600081815260208181526040808320838052909152808220939093559151909183156108fc02918491818181858888f1935050505015801561069b573d6000803e3d6000fd5b5050565b600160a060020a03821615156106b457600080fd5b6106c682600160a060020a03166109e1565b15156106d157600080fd5b600181146106de57600080fd5b33600090815260208181526040808320600160a060020a038616845290915290205483111561070c57600080fd5b33600090815260208181526040808320600160a060020a038616845290915290205461073e908463ffffffff61093516565b33600090815260208181526040808320600160a060020a038716845290915290205560018114156105ea57610783600160a060020a038316338563ffffffff6109e916565b505050565b600160a060020a0391821660009081526020818152604080832093909416825291909152205490565b600160a060020a03811615156107c657600080fd5b600160a060020a0381166000908152602081815260408083208380529091529020546107f8903463ffffffff610a4c16565b600160a060020a03909116600090815260208181526040808320838052909152902055565b600160a060020a038416151561083257600080fd5b600160a060020a038216151561084757600080fd5b61085982600160a060020a03166109e1565b151561086457600080fd5b6001811461087157600080fd5b600160a060020a03808516600090815260208181526040808320938616835292905220546108a5908463ffffffff610a4c16565b600160a060020a038086166000908152602081815260408083209387168352929052205560018114156105ea576108ed600160a060020a03831633308663ffffffff610a5f16565b50505050565b6108fb61149d565b61090361149d565b6109106020848551610b0a565b509392505050565b6109206114bf565b6109286114bf565b6109106020848551610cc9565b60008282111561094157fe5b50900390565b82600160a060020a031663095ea7b383836040518363ffffffff1660e060020a0281526004018083600160a060020a0316600160a060020a0316815260200182815260200192505050602060405180830381600087803b1580156109aa57600080fd5b505af11580156109be573d6000803e3d6000fd5b505050506040513d60208110156109d457600080fd5b5051151561078357600080fd5b6000903b1190565b82600160a060020a031663a9059cbb83836040518363ffffffff1660e060020a0281526004018083600160a060020a0316600160a060020a0316815260200182815260200192505050602060405180830381600087803b1580156109aa57600080fd5b81810182811015610a5957fe5b92915050565b604080517f23b872dd000000000000000000000000000000000000000000000000000000008152600160a060020a0385811660048301528481166024830152604482018490529151918616916323b872dd916064808201926020929091908290030181600087803b158015610ad357600080fd5b505af1158015610ae7573d6000803e3d6000fd5b505050506040513d6020811015610afd57600080fd5b505115156108ed57600080fd5b610b1261149d565b6000610b1c61149d565b610b2461150f565b60008080895b8881018b1015610bab57610b3e8b8b610f3b565b9c8d019c919550935091508360011415610b6e57610b658b8b610b5f610f78565b88610f86565b8b019a50610ba6565b8360021415610b8a57610b658b8b610b84610f78565b8861100c565b836003141561005e57610b658b8b610ba0610f78565b88611075565b610b2a565b9950898460016020020151604051908082528060200260200182016040528015610bdf578160200160208202803883390190505b508652604080860151815181815260208281028201019092528015610c0e578160200160208202803883390190505b5060208701528460036020020151604051908082528060200260200182016040528015610c45578160200160208202803883390190505b5060408701525b8881018b1015610cb857610c608b8b610f3b565b9c8d019c919550935091508360011415610c8957610c808b8b8888610f86565b8b019a50610cb3565b8360021415610c9e57610c808b8b888861100c565b836003141561005e57610c808b8b8888611075565b610c4c565b509399969850959650505050505050565b610cd16114bf565b6000610cdb6114bf565b610ce361152e565b60008080895b8881018b1015610dbe57610cfd8b8b610f3b565b9c8d019c919550935091508360011415610d2d57610d248b8b610d1e6110de565b886110e6565b8b019a50610db9565b8360021415610d4957610d248b8b610d436110de565b88611153565b8360031415610d5e57610d248b8b8888611162565b8360041415610d7a57610d248b8b610d746110de565b886111b3565b8360051415610d8f57610d248b8b888861121c565b8360061415610da457610d248b8b8888611265565b836007141561005e57610d248b8b88886112b6565b610ce9565b9950898460016020020151604051908082528060200260200182016040528015610df2578160200160208202803883390190505b508652604080860151815181815260208281028201019092528015610e21578160200160208202803883390190505b5060208701528460046020020151604051908082528060200260200182016040528015610e58578160200160208202803883390190505b5060608701525b8881018b1015610cb857610e738b8b610f3b565b9c8d019c919550935091508360011415610e9c57610e938b8b88886110e6565b8b019a50610f36565b8360021415610eb157610e938b8b8888611153565b8360031415610ecd57610e938b8b610ec76110de565b88611162565b8360041415610ee257610e938b8b88886111b3565b8360051415610efe57610e938b8b610ef86110de565b8861121c565b8360061415610f1a57610e938b8b610f146110de565b88611265565b836007141561005e57610e938b8b610f306110de565b886112b6565b610e5f565b6000806000806000806000610f5089896112ff565b93509350836007166005811115610f6357fe5b91506008840499919850919650945050505050565b610f8061149d565b50600090565b6000806000610f958787611344565b91509150610fa285611363565b15610fbc57600184815b6020020180519091019052611002565b8451602085015181518492918103908110610fd357fe5b60ff929092166020928302919091018201528401516000101561100257600184815b6020020180519190910390525b9695505050505050565b600080600061101b8787611367565b9150915061102885611363565b15611037576001846002610fac565b602085015160408501518151849291810390811061105157fe5b60209081029091010152600084600260200201511115611002576001846002610ff5565b60008060006110848787611367565b9150915061109185611363565b156110a0576001846003610fac565b60408501516060850151815184929181039081106110ba57fe5b60209081029091010152600084600360200201511115611002576001846003610ff5565b610f806114bf565b60008060006110f58787611381565b9150915061110285611363565b156111105760018481610fac565b845160208501518151849291810390811061112757fe5b600160a060020a0392909216602092830291909101820152840151600010156110025760018481610ff5565b600080600061101b8787611393565b60008060006111718787611381565b9150915061117e85611363565b1561118d576001846003610fac565b600160a060020a0382166040860152606084015160001015611002576001846003610ff5565b60008060006111c28787611393565b915091506111cf85611363565b156111de576001846004610fac565b60608501516080850151815184929181039081106111f857fe5b60209081029091010152600084600460200201511115611002576001846004610ff5565b600080600061122b8787611393565b9150915061123885611363565b15611247576001846005610fac565b6080850182905260a084015160001015611002576001846005610ff5565b60008060006112748787611381565b9150915061128185611363565b15611290576001846006610fac565b600160a060020a03821660a086015260c084015160001015611002576001846006610ff5565b60008060006112c58787611393565b915091506112d285611363565b156112e1576001846007610fac565b60c0850182905260e084015160001015611002576001846007610ff5565b908101906000808080805b865160001a90508160070260020a81607f1602831792506001820191506001870196506080811660801461130a5750909590945092505050565b600080600080611356600187876113a1565b9097909650945050505050565b1590565b600080611376602085856113eb565b915091509250929050565b60008060008061135660148787611430565b600080600080611356602087875b60008060008060006113b387876112ff565b91509150600387019650858701519250600282036020036101000a838115156113d857fe5b0498600190920197509095505050505050565b60008060008060006113fd87876112ff565b909350915082820160ff60038a01161461141657600080fd5b505050600393909201830151949390920160ff1692915050565b60008060008060008061144388886112ff565b909350915082820160ff60038b01161461145c57600080fd5b6003880197508688015193508860ff1690505b6020811015611487576101008404935060010161146f565b50919760039790970160ff169695505050505050565b6060604051908101604052806060815260200160608152602001606081525090565b60e06040519081016040528060608152602001606081526020016000600160a060020a0316815260200160608152602001600081526020016000600160a060020a03168152602001600081525090565b6080604051908101604052806004906020820280388339509192915050565b6101006040519081016040528060089060208202803883395091929150505600a165627a7a72305820a73e43793061d6700873efc76f5b82ce48b761884840a2ff9fbb82d52659d5390029`

// DeployDepositPool deploys a new Ethereum contract, binding an instance of DepositPool to it.
func DeployDepositPool(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DepositPool, error) {
	parsed, err := abi.JSON(strings.NewReader(DepositPoolABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DepositPoolBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DepositPool{DepositPoolCaller: DepositPoolCaller{contract: contract}, DepositPoolTransactor: DepositPoolTransactor{contract: contract}, DepositPoolFilterer: DepositPoolFilterer{contract: contract}}, nil
}

// DepositPool is an auto generated Go binding around an Ethereum contract.
type DepositPool struct {
	DepositPoolCaller     // Read-only binding to the contract
	DepositPoolTransactor // Write-only binding to the contract
	DepositPoolFilterer   // Log filterer for contract events
}

// DepositPoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type DepositPoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DepositPoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DepositPoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DepositPoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DepositPoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DepositPoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DepositPoolSession struct {
	Contract     *DepositPool      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DepositPoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DepositPoolCallerSession struct {
	Contract *DepositPoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// DepositPoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DepositPoolTransactorSession struct {
	Contract     *DepositPoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// DepositPoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type DepositPoolRaw struct {
	Contract *DepositPool // Generic contract binding to access the raw methods on
}

// DepositPoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DepositPoolCallerRaw struct {
	Contract *DepositPoolCaller // Generic read-only contract binding to access the raw methods on
}

// DepositPoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DepositPoolTransactorRaw struct {
	Contract *DepositPoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDepositPool creates a new instance of DepositPool, bound to a specific deployed contract.
func NewDepositPool(address common.Address, backend bind.ContractBackend) (*DepositPool, error) {
	contract, err := bindDepositPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DepositPool{DepositPoolCaller: DepositPoolCaller{contract: contract}, DepositPoolTransactor: DepositPoolTransactor{contract: contract}, DepositPoolFilterer: DepositPoolFilterer{contract: contract}}, nil
}

// NewDepositPoolCaller creates a new read-only instance of DepositPool, bound to a specific deployed contract.
func NewDepositPoolCaller(address common.Address, caller bind.ContractCaller) (*DepositPoolCaller, error) {
	contract, err := bindDepositPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DepositPoolCaller{contract: contract}, nil
}

// NewDepositPoolTransactor creates a new write-only instance of DepositPool, bound to a specific deployed contract.
func NewDepositPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*DepositPoolTransactor, error) {
	contract, err := bindDepositPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DepositPoolTransactor{contract: contract}, nil
}

// NewDepositPoolFilterer creates a new log filterer instance of DepositPool, bound to a specific deployed contract.
func NewDepositPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*DepositPoolFilterer, error) {
	contract, err := bindDepositPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DepositPoolFilterer{contract: contract}, nil
}

// bindDepositPool binds a generic wrapper to an already deployed contract.
func bindDepositPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DepositPoolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DepositPool *DepositPoolRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DepositPool.Contract.DepositPoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DepositPool *DepositPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DepositPool.Contract.DepositPoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DepositPool *DepositPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DepositPool.Contract.DepositPoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DepositPool *DepositPoolCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DepositPool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DepositPool *DepositPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DepositPool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DepositPool *DepositPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DepositPool.Contract.contract.Transact(opts, method, params...)
}

// GetRemainingBalance is a free data retrieval call binding the contract method 0x81ae8541.
//
// Solidity: function getRemainingBalance(_addr address, _tokenContract address) constant returns(uint256)
func (_DepositPool *DepositPoolCaller) GetRemainingBalance(opts *bind.CallOpts, _addr common.Address, _tokenContract common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DepositPool.contract.Call(opts, out, "getRemainingBalance", _addr, _tokenContract)
	return *ret0, err
}

// GetRemainingBalance is a free data retrieval call binding the contract method 0x81ae8541.
//
// Solidity: function getRemainingBalance(_addr address, _tokenContract address) constant returns(uint256)
func (_DepositPool *DepositPoolSession) GetRemainingBalance(_addr common.Address, _tokenContract common.Address) (*big.Int, error) {
	return _DepositPool.Contract.GetRemainingBalance(&_DepositPool.CallOpts, _addr, _tokenContract)
}

// GetRemainingBalance is a free data retrieval call binding the contract method 0x81ae8541.
//
// Solidity: function getRemainingBalance(_addr address, _tokenContract address) constant returns(uint256)
func (_DepositPool *DepositPoolCallerSession) GetRemainingBalance(_addr common.Address, _tokenContract common.Address) (*big.Int, error) {
	return _DepositPool.Contract.GetRemainingBalance(&_DepositPool.CallOpts, _addr, _tokenContract)
}

// AuthorizedWithdraw is a paid mutator transaction binding the contract method 0x06fa40f3.
//
// Solidity: function authorizedWithdraw(_authWithdraw bytes, _signature bytes, _channelId uint256) returns()
func (_DepositPool *DepositPoolTransactor) AuthorizedWithdraw(opts *bind.TransactOpts, _authWithdraw []byte, _signature []byte, _channelId *big.Int) (*types.Transaction, error) {
	return _DepositPool.contract.Transact(opts, "authorizedWithdraw", _authWithdraw, _signature, _channelId)
}

// AuthorizedWithdraw is a paid mutator transaction binding the contract method 0x06fa40f3.
//
// Solidity: function authorizedWithdraw(_authWithdraw bytes, _signature bytes, _channelId uint256) returns()
func (_DepositPool *DepositPoolSession) AuthorizedWithdraw(_authWithdraw []byte, _signature []byte, _channelId *big.Int) (*types.Transaction, error) {
	return _DepositPool.Contract.AuthorizedWithdraw(&_DepositPool.TransactOpts, _authWithdraw, _signature, _channelId)
}

// AuthorizedWithdraw is a paid mutator transaction binding the contract method 0x06fa40f3.
//
// Solidity: function authorizedWithdraw(_authWithdraw bytes, _signature bytes, _channelId uint256) returns()
func (_DepositPool *DepositPoolTransactorSession) AuthorizedWithdraw(_authWithdraw []byte, _signature []byte, _channelId *big.Int) (*types.Transaction, error) {
	return _DepositPool.Contract.AuthorizedWithdraw(&_DepositPool.TransactOpts, _authWithdraw, _signature, _channelId)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(_receipient address) returns()
func (_DepositPool *DepositPoolTransactor) Deposit(opts *bind.TransactOpts, _receipient common.Address) (*types.Transaction, error) {
	return _DepositPool.contract.Transact(opts, "deposit", _receipient)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(_receipient address) returns()
func (_DepositPool *DepositPoolSession) Deposit(_receipient common.Address) (*types.Transaction, error) {
	return _DepositPool.Contract.Deposit(&_DepositPool.TransactOpts, _receipient)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(_receipient address) returns()
func (_DepositPool *DepositPoolTransactorSession) Deposit(_receipient common.Address) (*types.Transaction, error) {
	return _DepositPool.Contract.Deposit(&_DepositPool.TransactOpts, _receipient)
}

// DepositERCToken is a paid mutator transaction binding the contract method 0xfcc3e991.
//
// Solidity: function depositERCToken(_receipient address, _amount uint256, _tokenContract address, _tokenType uint256) returns()
func (_DepositPool *DepositPoolTransactor) DepositERCToken(opts *bind.TransactOpts, _receipient common.Address, _amount *big.Int, _tokenContract common.Address, _tokenType *big.Int) (*types.Transaction, error) {
	return _DepositPool.contract.Transact(opts, "depositERCToken", _receipient, _amount, _tokenContract, _tokenType)
}

// DepositERCToken is a paid mutator transaction binding the contract method 0xfcc3e991.
//
// Solidity: function depositERCToken(_receipient address, _amount uint256, _tokenContract address, _tokenType uint256) returns()
func (_DepositPool *DepositPoolSession) DepositERCToken(_receipient common.Address, _amount *big.Int, _tokenContract common.Address, _tokenType *big.Int) (*types.Transaction, error) {
	return _DepositPool.Contract.DepositERCToken(&_DepositPool.TransactOpts, _receipient, _amount, _tokenContract, _tokenType)
}

// DepositERCToken is a paid mutator transaction binding the contract method 0xfcc3e991.
//
// Solidity: function depositERCToken(_receipient address, _amount uint256, _tokenContract address, _tokenType uint256) returns()
func (_DepositPool *DepositPoolTransactorSession) DepositERCToken(_receipient common.Address, _amount *big.Int, _tokenContract common.Address, _tokenType *big.Int) (*types.Transaction, error) {
	return _DepositPool.Contract.DepositERCToken(&_DepositPool.TransactOpts, _receipient, _amount, _tokenContract, _tokenType)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(_amount uint256) returns()
func (_DepositPool *DepositPoolTransactor) Withdraw(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _DepositPool.contract.Transact(opts, "withdraw", _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(_amount uint256) returns()
func (_DepositPool *DepositPoolSession) Withdraw(_amount *big.Int) (*types.Transaction, error) {
	return _DepositPool.Contract.Withdraw(&_DepositPool.TransactOpts, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(_amount uint256) returns()
func (_DepositPool *DepositPoolTransactorSession) Withdraw(_amount *big.Int) (*types.Transaction, error) {
	return _DepositPool.Contract.Withdraw(&_DepositPool.TransactOpts, _amount)
}

// WithdrawERCToken is a paid mutator transaction binding the contract method 0x78320258.
//
// Solidity: function withdrawERCToken(_amount uint256, _tokenContract address, _tokenType uint256) returns()
func (_DepositPool *DepositPoolTransactor) WithdrawERCToken(opts *bind.TransactOpts, _amount *big.Int, _tokenContract common.Address, _tokenType *big.Int) (*types.Transaction, error) {
	return _DepositPool.contract.Transact(opts, "withdrawERCToken", _amount, _tokenContract, _tokenType)
}

// WithdrawERCToken is a paid mutator transaction binding the contract method 0x78320258.
//
// Solidity: function withdrawERCToken(_amount uint256, _tokenContract address, _tokenType uint256) returns()
func (_DepositPool *DepositPoolSession) WithdrawERCToken(_amount *big.Int, _tokenContract common.Address, _tokenType *big.Int) (*types.Transaction, error) {
	return _DepositPool.Contract.WithdrawERCToken(&_DepositPool.TransactOpts, _amount, _tokenContract, _tokenType)
}

// WithdrawERCToken is a paid mutator transaction binding the contract method 0x78320258.
//
// Solidity: function withdrawERCToken(_amount uint256, _tokenContract address, _tokenType uint256) returns()
func (_DepositPool *DepositPoolTransactorSession) WithdrawERCToken(_amount *big.Int, _tokenContract common.Address, _tokenType *big.Int) (*types.Transaction, error) {
	return _DepositPool.Contract.WithdrawERCToken(&_DepositPool.TransactOpts, _amount, _tokenContract, _tokenType)
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

// GenericChannelInterfaceABI is the input ABI used to generate the binding from.
const GenericChannelInterfaceABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_channelId\",\"type\":\"uint256\"},{\"name\":\"_receipient\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"depositERCToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_channelId\",\"type\":\"uint256\"},{\"name\":\"_receipient\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// GenericChannelInterfaceBin is the compiled bytecode used for deploying new contracts.
const GenericChannelInterfaceBin = `0x`

// DeployGenericChannelInterface deploys a new Ethereum contract, binding an instance of GenericChannelInterface to it.
func DeployGenericChannelInterface(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GenericChannelInterface, error) {
	parsed, err := abi.JSON(strings.NewReader(GenericChannelInterfaceABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(GenericChannelInterfaceBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GenericChannelInterface{GenericChannelInterfaceCaller: GenericChannelInterfaceCaller{contract: contract}, GenericChannelInterfaceTransactor: GenericChannelInterfaceTransactor{contract: contract}, GenericChannelInterfaceFilterer: GenericChannelInterfaceFilterer{contract: contract}}, nil
}

// GenericChannelInterface is an auto generated Go binding around an Ethereum contract.
type GenericChannelInterface struct {
	GenericChannelInterfaceCaller     // Read-only binding to the contract
	GenericChannelInterfaceTransactor // Write-only binding to the contract
	GenericChannelInterfaceFilterer   // Log filterer for contract events
}

// GenericChannelInterfaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type GenericChannelInterfaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GenericChannelInterfaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GenericChannelInterfaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GenericChannelInterfaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GenericChannelInterfaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GenericChannelInterfaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GenericChannelInterfaceSession struct {
	Contract     *GenericChannelInterface // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// GenericChannelInterfaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GenericChannelInterfaceCallerSession struct {
	Contract *GenericChannelInterfaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// GenericChannelInterfaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GenericChannelInterfaceTransactorSession struct {
	Contract     *GenericChannelInterfaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// GenericChannelInterfaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type GenericChannelInterfaceRaw struct {
	Contract *GenericChannelInterface // Generic contract binding to access the raw methods on
}

// GenericChannelInterfaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GenericChannelInterfaceCallerRaw struct {
	Contract *GenericChannelInterfaceCaller // Generic read-only contract binding to access the raw methods on
}

// GenericChannelInterfaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GenericChannelInterfaceTransactorRaw struct {
	Contract *GenericChannelInterfaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGenericChannelInterface creates a new instance of GenericChannelInterface, bound to a specific deployed contract.
func NewGenericChannelInterface(address common.Address, backend bind.ContractBackend) (*GenericChannelInterface, error) {
	contract, err := bindGenericChannelInterface(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GenericChannelInterface{GenericChannelInterfaceCaller: GenericChannelInterfaceCaller{contract: contract}, GenericChannelInterfaceTransactor: GenericChannelInterfaceTransactor{contract: contract}, GenericChannelInterfaceFilterer: GenericChannelInterfaceFilterer{contract: contract}}, nil
}

// NewGenericChannelInterfaceCaller creates a new read-only instance of GenericChannelInterface, bound to a specific deployed contract.
func NewGenericChannelInterfaceCaller(address common.Address, caller bind.ContractCaller) (*GenericChannelInterfaceCaller, error) {
	contract, err := bindGenericChannelInterface(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GenericChannelInterfaceCaller{contract: contract}, nil
}

// NewGenericChannelInterfaceTransactor creates a new write-only instance of GenericChannelInterface, bound to a specific deployed contract.
func NewGenericChannelInterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*GenericChannelInterfaceTransactor, error) {
	contract, err := bindGenericChannelInterface(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GenericChannelInterfaceTransactor{contract: contract}, nil
}

// NewGenericChannelInterfaceFilterer creates a new log filterer instance of GenericChannelInterface, bound to a specific deployed contract.
func NewGenericChannelInterfaceFilterer(address common.Address, filterer bind.ContractFilterer) (*GenericChannelInterfaceFilterer, error) {
	contract, err := bindGenericChannelInterface(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GenericChannelInterfaceFilterer{contract: contract}, nil
}

// bindGenericChannelInterface binds a generic wrapper to an already deployed contract.
func bindGenericChannelInterface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GenericChannelInterfaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GenericChannelInterface *GenericChannelInterfaceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GenericChannelInterface.Contract.GenericChannelInterfaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GenericChannelInterface *GenericChannelInterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GenericChannelInterface.Contract.GenericChannelInterfaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GenericChannelInterface *GenericChannelInterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GenericChannelInterface.Contract.GenericChannelInterfaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GenericChannelInterface *GenericChannelInterfaceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GenericChannelInterface.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GenericChannelInterface *GenericChannelInterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GenericChannelInterface.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GenericChannelInterface *GenericChannelInterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GenericChannelInterface.Contract.contract.Transact(opts, method, params...)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(_channelId uint256, _receipient address) returns()
func (_GenericChannelInterface *GenericChannelInterfaceTransactor) Deposit(opts *bind.TransactOpts, _channelId *big.Int, _receipient common.Address) (*types.Transaction, error) {
	return _GenericChannelInterface.contract.Transact(opts, "deposit", _channelId, _receipient)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(_channelId uint256, _receipient address) returns()
func (_GenericChannelInterface *GenericChannelInterfaceSession) Deposit(_channelId *big.Int, _receipient common.Address) (*types.Transaction, error) {
	return _GenericChannelInterface.Contract.Deposit(&_GenericChannelInterface.TransactOpts, _channelId, _receipient)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(_channelId uint256, _receipient address) returns()
func (_GenericChannelInterface *GenericChannelInterfaceTransactorSession) Deposit(_channelId *big.Int, _receipient common.Address) (*types.Transaction, error) {
	return _GenericChannelInterface.Contract.Deposit(&_GenericChannelInterface.TransactOpts, _channelId, _receipient)
}

// DepositERCToken is a paid mutator transaction binding the contract method 0x235edd9a.
//
// Solidity: function depositERCToken(_channelId uint256, _receipient address, _amount uint256) returns()
func (_GenericChannelInterface *GenericChannelInterfaceTransactor) DepositERCToken(opts *bind.TransactOpts, _channelId *big.Int, _receipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GenericChannelInterface.contract.Transact(opts, "depositERCToken", _channelId, _receipient, _amount)
}

// DepositERCToken is a paid mutator transaction binding the contract method 0x235edd9a.
//
// Solidity: function depositERCToken(_channelId uint256, _receipient address, _amount uint256) returns()
func (_GenericChannelInterface *GenericChannelInterfaceSession) DepositERCToken(_channelId *big.Int, _receipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GenericChannelInterface.Contract.DepositERCToken(&_GenericChannelInterface.TransactOpts, _channelId, _receipient, _amount)
}

// DepositERCToken is a paid mutator transaction binding the contract method 0x235edd9a.
//
// Solidity: function depositERCToken(_channelId uint256, _receipient address, _amount uint256) returns()
func (_GenericChannelInterface *GenericChannelInterfaceTransactorSession) DepositERCToken(_channelId *big.Int, _receipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GenericChannelInterface.Contract.DepositERCToken(&_GenericChannelInterface.TransactOpts, _channelId, _receipient, _amount)
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
