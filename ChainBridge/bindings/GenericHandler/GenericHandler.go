// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package GenericHandler

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

// GenericHandlerMetaData contains all meta data concerning the GenericHandler contract.
var GenericHandlerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridgeAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"_bridgeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_contractAddressToDepositFunctionDepositerOffset\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_contractAddressToDepositFunctionSignature\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_contractAddressToExecuteFunctionSignature\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_contractAddressToResourceID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_contractWhitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"_resourceIDToContractAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"depositFunctionSig\",\"type\":\"bytes4\"},{\"internalType\":\"uint256\",\"name\":\"depositFunctionDepositerOffset\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"executeFunctionSig\",\"type\":\"bytes4\"}],\"name\":\"setResource\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"depositer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"executeProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561001057600080fd5b50604051610b80380380610b8083398101604081905261002f91610044565b60601b6001600160601b031916608052610072565b600060208284031215610055578081fd5b81516001600160a01b038116811461006b578182fd5b9392505050565b60805160601c610aec610094600039806101a452806105de5250610aec6000f3fe608060405234801561001057600080fd5b506004361061009e5760003560e01c8063c54c2a1111610066578063c54c2a1114610141578063cb62446314610154578063de319d9914610167578063e248cff21461017c578063ec97d3b41461018f5761009e565b8063318c136e146100a35780637f79bea8146100c1578063a5c3a985146100e1578063aa50800b14610101578063b07e54bb14610121575b600080fd5b6100ab6101a2565b6040516100b8919061089f565b60405180910390f35b6100d46100cf36600461070c565b6101c6565b6040516100b891906108b3565b6100f46100ef36600461070c565b6101db565b6040516100b891906108c7565b61011461010f36600461070c565b6101f0565b6040516100b891906108be565b61013461012f3660046107af565b610202565b6040516100b891906108dc565b6100ab61014f36600461072e565b6103f5565b6100f461016236600461070c565b610410565b61017a610175366004610746565b610425565b005b61017a61018a366004610808565b610441565b61011461019d36600461070c565b6105c1565b7f000000000000000000000000000000000000000000000000000000000000000081565b60056020526000908152604090205460ff1681565b60046020526000908152604090205460e01b81565b60036020526000908152604090205481565b606061020c6105d3565b6000606061021c8486018661072e565b915061022e6020808401908688610a45565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201829052508b815260208181526040808320546001600160a01b031680845260039092529091205494955093925050811590506102ce57828101602001516001600160a01b038916606082901c146102cc5760405162461bcd60e51b81526004016102c39061098c565b60405180910390fd5b505b6001600160a01b03821660009081526005602052604090205460ff166103065760405162461bcd60e51b81526004016102c3906109c3565b6001600160a01b03821660009081526002602052604090205460e01b6001600160e01b03198116156103e75760608185604051602001610347929190610852565b604051602081830303815290604052905060006060856001600160a01b0316836040516103749190610883565b6000604051808303816000865af19150503d80600081146103b1576040519150601f19603f3d011682016040523d82523d6000602084013e6103b6565b606091505b5091509150816103d85760405162461bcd60e51b81526004016102c390610a0e565b97506103ed9650505050505050565b50505050505b949350505050565b6000602081905290815260409020546001600160a01b031681565b60026020526000908152604090205460e01b81565b61042d6105d3565b61043a858585858561061d565b5050505050565b6104496105d3565b600060606104598385018561072e565b915061046b6020808401908587610a45565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092018290525089815260208181526040808320546001600160a01b03168084526005909252909120549495509360ff1692506104e69150505760405162461bcd60e51b81526004016102c3906109c3565b6001600160a01b03811660009081526004602052604090205460e01b6001600160e01b03198116156105b85760608184604051602001610527929190610852565b60405160208183030381529060405290506000836001600160a01b0316826040516105529190610883565b6000604051808303816000865af19150503d806000811461058f576040519150601f19603f3d011682016040523d82523d6000602084013e610594565b606091505b50509050806105b55760405162461bcd60e51b81526004016102c390610946565b50505b50505050505050565b60016020526000908152604090205481565b336001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461061b5760405162461bcd60e51b81526004016102c39061090f565b565b60008581526020818152604080832080546001600160a01b039098166001600160a01b03199098168817905595825260018082528683209790975560028152858220805460e096871c63ffffffff199182161790915560038252868320949094556004815285822080549390951c92909316919091179092556005905220805460ff19169091179055565b80356001600160a01b03811681146106bf57600080fd5b92915050565b60008083601f8401126106d6578182fd5b50813567ffffffffffffffff8111156106ed578182fd5b60208301915083602082850101111561070557600080fd5b9250929050565b60006020828403121561071d578081fd5b61072783836106a8565b9392505050565b60006020828403121561073f578081fd5b5035919050565b600080600080600060a0868803121561075d578081fd5b8535945060208601356001600160a01b038116811461077a578182fd5b9350604086013561078a81610a9d565b92506060860135915060808601356107a181610a9d565b809150509295509295909350565b600080600080606085870312156107c4578384fd5b843593506107d586602087016106a8565b9250604085013567ffffffffffffffff8111156107f0578283fd5b6107fc878288016106c5565b95989497509550505050565b60008060006040848603121561081c578283fd5b83359250602084013567ffffffffffffffff811115610839578283fd5b610845868287016106c5565b9497909650939450505050565b6001600160e01b0319831681528151600090610875816004850160208701610a6d565b919091016004019392505050565b60008251610895818460208701610a6d565b9190910192915050565b6001600160a01b0391909116815260200190565b901515815260200190565b90815260200190565b6001600160e01b031991909116815260200190565b60006020825282518060208401526108fb816040850160208701610a6d565b601f01601f19169190910160400192915050565b6020808252601e908201527f73656e646572206d7573742062652062726964676520636f6e74726163740000604082015260600190565b60208082526026908201527f64656c656761746563616c6c20746f20636f6e7472616374416464726573732060408201526519985a5b195960d21b606082015260800190565b6020808252601f908201527f696e636f7272656374206465706f736974657220696e20746865206461746100604082015260600190565b6020808252602b908201527f70726f766964656420636f6e747261637441646472657373206973206e6f742060408201526a1dda1a5d195b1a5cdd195960aa1b606082015260800190565b6020808252601e908201527f63616c6c20746f20636f6e747261637441646472657373206661696c65640000604082015260600190565b60008085851115610a54578182fd5b83861115610a60578182fd5b5050820193919092039150565b60005b83811015610a88578181015183820152602001610a70565b83811115610a97576000848401525b50505050565b6001600160e01b031981168114610ab357600080fd5b5056fea26469706673582212206629f1b20a59266ba81cd21a0a9659a1ccffc5bda3d96cf6d3743ab5f1eab7cf64736f6c634300060c0033",
}

// GenericHandlerABI is the input ABI used to generate the binding from.
// Deprecated: Use GenericHandlerMetaData.ABI instead.
var GenericHandlerABI = GenericHandlerMetaData.ABI

// GenericHandlerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GenericHandlerMetaData.Bin instead.
var GenericHandlerBin = GenericHandlerMetaData.Bin

// DeployGenericHandler deploys a new Ethereum contract, binding an instance of GenericHandler to it.
func DeployGenericHandler(auth *bind.TransactOpts, backend bind.ContractBackend, bridgeAddress common.Address) (common.Address, *types.Transaction, *GenericHandler, error) {
	parsed, err := GenericHandlerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GenericHandlerBin), backend, bridgeAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GenericHandler{GenericHandlerCaller: GenericHandlerCaller{contract: contract}, GenericHandlerTransactor: GenericHandlerTransactor{contract: contract}, GenericHandlerFilterer: GenericHandlerFilterer{contract: contract}}, nil
}

// GenericHandler is an auto generated Go binding around an Ethereum contract.
type GenericHandler struct {
	GenericHandlerCaller     // Read-only binding to the contract
	GenericHandlerTransactor // Write-only binding to the contract
	GenericHandlerFilterer   // Log filterer for contract events
}

// GenericHandlerCaller is an auto generated read-only Go binding around an Ethereum contract.
type GenericHandlerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GenericHandlerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GenericHandlerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GenericHandlerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GenericHandlerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GenericHandlerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GenericHandlerSession struct {
	Contract     *GenericHandler   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GenericHandlerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GenericHandlerCallerSession struct {
	Contract *GenericHandlerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// GenericHandlerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GenericHandlerTransactorSession struct {
	Contract     *GenericHandlerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// GenericHandlerRaw is an auto generated low-level Go binding around an Ethereum contract.
type GenericHandlerRaw struct {
	Contract *GenericHandler // Generic contract binding to access the raw methods on
}

// GenericHandlerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GenericHandlerCallerRaw struct {
	Contract *GenericHandlerCaller // Generic read-only contract binding to access the raw methods on
}

// GenericHandlerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GenericHandlerTransactorRaw struct {
	Contract *GenericHandlerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGenericHandler creates a new instance of GenericHandler, bound to a specific deployed contract.
func NewGenericHandler(address common.Address, backend bind.ContractBackend) (*GenericHandler, error) {
	contract, err := bindGenericHandler(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GenericHandler{GenericHandlerCaller: GenericHandlerCaller{contract: contract}, GenericHandlerTransactor: GenericHandlerTransactor{contract: contract}, GenericHandlerFilterer: GenericHandlerFilterer{contract: contract}}, nil
}

// NewGenericHandlerCaller creates a new read-only instance of GenericHandler, bound to a specific deployed contract.
func NewGenericHandlerCaller(address common.Address, caller bind.ContractCaller) (*GenericHandlerCaller, error) {
	contract, err := bindGenericHandler(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GenericHandlerCaller{contract: contract}, nil
}

// NewGenericHandlerTransactor creates a new write-only instance of GenericHandler, bound to a specific deployed contract.
func NewGenericHandlerTransactor(address common.Address, transactor bind.ContractTransactor) (*GenericHandlerTransactor, error) {
	contract, err := bindGenericHandler(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GenericHandlerTransactor{contract: contract}, nil
}

// NewGenericHandlerFilterer creates a new log filterer instance of GenericHandler, bound to a specific deployed contract.
func NewGenericHandlerFilterer(address common.Address, filterer bind.ContractFilterer) (*GenericHandlerFilterer, error) {
	contract, err := bindGenericHandler(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GenericHandlerFilterer{contract: contract}, nil
}

// bindGenericHandler binds a generic wrapper to an already deployed contract.
func bindGenericHandler(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GenericHandlerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GenericHandler *GenericHandlerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GenericHandler.Contract.GenericHandlerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GenericHandler *GenericHandlerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GenericHandler.Contract.GenericHandlerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GenericHandler *GenericHandlerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GenericHandler.Contract.GenericHandlerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GenericHandler *GenericHandlerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GenericHandler.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GenericHandler *GenericHandlerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GenericHandler.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GenericHandler *GenericHandlerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GenericHandler.Contract.contract.Transact(opts, method, params...)
}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_GenericHandler *GenericHandlerCaller) BridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GenericHandler.contract.Call(opts, &out, "_bridgeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_GenericHandler *GenericHandlerSession) BridgeAddress() (common.Address, error) {
	return _GenericHandler.Contract.BridgeAddress(&_GenericHandler.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_GenericHandler *GenericHandlerCallerSession) BridgeAddress() (common.Address, error) {
	return _GenericHandler.Contract.BridgeAddress(&_GenericHandler.CallOpts)
}

// ContractAddressToDepositFunctionDepositerOffset is a free data retrieval call binding the contract method 0xaa50800b.
//
// Solidity: function _contractAddressToDepositFunctionDepositerOffset(address ) view returns(uint256)
func (_GenericHandler *GenericHandlerCaller) ContractAddressToDepositFunctionDepositerOffset(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GenericHandler.contract.Call(opts, &out, "_contractAddressToDepositFunctionDepositerOffset", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ContractAddressToDepositFunctionDepositerOffset is a free data retrieval call binding the contract method 0xaa50800b.
//
// Solidity: function _contractAddressToDepositFunctionDepositerOffset(address ) view returns(uint256)
func (_GenericHandler *GenericHandlerSession) ContractAddressToDepositFunctionDepositerOffset(arg0 common.Address) (*big.Int, error) {
	return _GenericHandler.Contract.ContractAddressToDepositFunctionDepositerOffset(&_GenericHandler.CallOpts, arg0)
}

// ContractAddressToDepositFunctionDepositerOffset is a free data retrieval call binding the contract method 0xaa50800b.
//
// Solidity: function _contractAddressToDepositFunctionDepositerOffset(address ) view returns(uint256)
func (_GenericHandler *GenericHandlerCallerSession) ContractAddressToDepositFunctionDepositerOffset(arg0 common.Address) (*big.Int, error) {
	return _GenericHandler.Contract.ContractAddressToDepositFunctionDepositerOffset(&_GenericHandler.CallOpts, arg0)
}

// ContractAddressToDepositFunctionSignature is a free data retrieval call binding the contract method 0xcb624463.
//
// Solidity: function _contractAddressToDepositFunctionSignature(address ) view returns(bytes4)
func (_GenericHandler *GenericHandlerCaller) ContractAddressToDepositFunctionSignature(opts *bind.CallOpts, arg0 common.Address) ([4]byte, error) {
	var out []interface{}
	err := _GenericHandler.contract.Call(opts, &out, "_contractAddressToDepositFunctionSignature", arg0)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// ContractAddressToDepositFunctionSignature is a free data retrieval call binding the contract method 0xcb624463.
//
// Solidity: function _contractAddressToDepositFunctionSignature(address ) view returns(bytes4)
func (_GenericHandler *GenericHandlerSession) ContractAddressToDepositFunctionSignature(arg0 common.Address) ([4]byte, error) {
	return _GenericHandler.Contract.ContractAddressToDepositFunctionSignature(&_GenericHandler.CallOpts, arg0)
}

// ContractAddressToDepositFunctionSignature is a free data retrieval call binding the contract method 0xcb624463.
//
// Solidity: function _contractAddressToDepositFunctionSignature(address ) view returns(bytes4)
func (_GenericHandler *GenericHandlerCallerSession) ContractAddressToDepositFunctionSignature(arg0 common.Address) ([4]byte, error) {
	return _GenericHandler.Contract.ContractAddressToDepositFunctionSignature(&_GenericHandler.CallOpts, arg0)
}

// ContractAddressToExecuteFunctionSignature is a free data retrieval call binding the contract method 0xa5c3a985.
//
// Solidity: function _contractAddressToExecuteFunctionSignature(address ) view returns(bytes4)
func (_GenericHandler *GenericHandlerCaller) ContractAddressToExecuteFunctionSignature(opts *bind.CallOpts, arg0 common.Address) ([4]byte, error) {
	var out []interface{}
	err := _GenericHandler.contract.Call(opts, &out, "_contractAddressToExecuteFunctionSignature", arg0)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// ContractAddressToExecuteFunctionSignature is a free data retrieval call binding the contract method 0xa5c3a985.
//
// Solidity: function _contractAddressToExecuteFunctionSignature(address ) view returns(bytes4)
func (_GenericHandler *GenericHandlerSession) ContractAddressToExecuteFunctionSignature(arg0 common.Address) ([4]byte, error) {
	return _GenericHandler.Contract.ContractAddressToExecuteFunctionSignature(&_GenericHandler.CallOpts, arg0)
}

// ContractAddressToExecuteFunctionSignature is a free data retrieval call binding the contract method 0xa5c3a985.
//
// Solidity: function _contractAddressToExecuteFunctionSignature(address ) view returns(bytes4)
func (_GenericHandler *GenericHandlerCallerSession) ContractAddressToExecuteFunctionSignature(arg0 common.Address) ([4]byte, error) {
	return _GenericHandler.Contract.ContractAddressToExecuteFunctionSignature(&_GenericHandler.CallOpts, arg0)
}

// ContractAddressToResourceID is a free data retrieval call binding the contract method 0xec97d3b4.
//
// Solidity: function _contractAddressToResourceID(address ) view returns(bytes32)
func (_GenericHandler *GenericHandlerCaller) ContractAddressToResourceID(opts *bind.CallOpts, arg0 common.Address) ([32]byte, error) {
	var out []interface{}
	err := _GenericHandler.contract.Call(opts, &out, "_contractAddressToResourceID", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ContractAddressToResourceID is a free data retrieval call binding the contract method 0xec97d3b4.
//
// Solidity: function _contractAddressToResourceID(address ) view returns(bytes32)
func (_GenericHandler *GenericHandlerSession) ContractAddressToResourceID(arg0 common.Address) ([32]byte, error) {
	return _GenericHandler.Contract.ContractAddressToResourceID(&_GenericHandler.CallOpts, arg0)
}

// ContractAddressToResourceID is a free data retrieval call binding the contract method 0xec97d3b4.
//
// Solidity: function _contractAddressToResourceID(address ) view returns(bytes32)
func (_GenericHandler *GenericHandlerCallerSession) ContractAddressToResourceID(arg0 common.Address) ([32]byte, error) {
	return _GenericHandler.Contract.ContractAddressToResourceID(&_GenericHandler.CallOpts, arg0)
}

// ContractWhitelist is a free data retrieval call binding the contract method 0x7f79bea8.
//
// Solidity: function _contractWhitelist(address ) view returns(bool)
func (_GenericHandler *GenericHandlerCaller) ContractWhitelist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _GenericHandler.contract.Call(opts, &out, "_contractWhitelist", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ContractWhitelist is a free data retrieval call binding the contract method 0x7f79bea8.
//
// Solidity: function _contractWhitelist(address ) view returns(bool)
func (_GenericHandler *GenericHandlerSession) ContractWhitelist(arg0 common.Address) (bool, error) {
	return _GenericHandler.Contract.ContractWhitelist(&_GenericHandler.CallOpts, arg0)
}

// ContractWhitelist is a free data retrieval call binding the contract method 0x7f79bea8.
//
// Solidity: function _contractWhitelist(address ) view returns(bool)
func (_GenericHandler *GenericHandlerCallerSession) ContractWhitelist(arg0 common.Address) (bool, error) {
	return _GenericHandler.Contract.ContractWhitelist(&_GenericHandler.CallOpts, arg0)
}

// ResourceIDToContractAddress is a free data retrieval call binding the contract method 0xc54c2a11.
//
// Solidity: function _resourceIDToContractAddress(bytes32 ) view returns(address)
func (_GenericHandler *GenericHandlerCaller) ResourceIDToContractAddress(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _GenericHandler.contract.Call(opts, &out, "_resourceIDToContractAddress", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ResourceIDToContractAddress is a free data retrieval call binding the contract method 0xc54c2a11.
//
// Solidity: function _resourceIDToContractAddress(bytes32 ) view returns(address)
func (_GenericHandler *GenericHandlerSession) ResourceIDToContractAddress(arg0 [32]byte) (common.Address, error) {
	return _GenericHandler.Contract.ResourceIDToContractAddress(&_GenericHandler.CallOpts, arg0)
}

// ResourceIDToContractAddress is a free data retrieval call binding the contract method 0xc54c2a11.
//
// Solidity: function _resourceIDToContractAddress(bytes32 ) view returns(address)
func (_GenericHandler *GenericHandlerCallerSession) ResourceIDToContractAddress(arg0 [32]byte) (common.Address, error) {
	return _GenericHandler.Contract.ResourceIDToContractAddress(&_GenericHandler.CallOpts, arg0)
}

// Deposit is a paid mutator transaction binding the contract method 0xb07e54bb.
//
// Solidity: function deposit(bytes32 resourceID, address depositer, bytes data) returns(bytes)
func (_GenericHandler *GenericHandlerTransactor) Deposit(opts *bind.TransactOpts, resourceID [32]byte, depositer common.Address, data []byte) (*types.Transaction, error) {
	return _GenericHandler.contract.Transact(opts, "deposit", resourceID, depositer, data)
}

// Deposit is a paid mutator transaction binding the contract method 0xb07e54bb.
//
// Solidity: function deposit(bytes32 resourceID, address depositer, bytes data) returns(bytes)
func (_GenericHandler *GenericHandlerSession) Deposit(resourceID [32]byte, depositer common.Address, data []byte) (*types.Transaction, error) {
	return _GenericHandler.Contract.Deposit(&_GenericHandler.TransactOpts, resourceID, depositer, data)
}

// Deposit is a paid mutator transaction binding the contract method 0xb07e54bb.
//
// Solidity: function deposit(bytes32 resourceID, address depositer, bytes data) returns(bytes)
func (_GenericHandler *GenericHandlerTransactorSession) Deposit(resourceID [32]byte, depositer common.Address, data []byte) (*types.Transaction, error) {
	return _GenericHandler.Contract.Deposit(&_GenericHandler.TransactOpts, resourceID, depositer, data)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0xe248cff2.
//
// Solidity: function executeProposal(bytes32 resourceID, bytes data) returns()
func (_GenericHandler *GenericHandlerTransactor) ExecuteProposal(opts *bind.TransactOpts, resourceID [32]byte, data []byte) (*types.Transaction, error) {
	return _GenericHandler.contract.Transact(opts, "executeProposal", resourceID, data)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0xe248cff2.
//
// Solidity: function executeProposal(bytes32 resourceID, bytes data) returns()
func (_GenericHandler *GenericHandlerSession) ExecuteProposal(resourceID [32]byte, data []byte) (*types.Transaction, error) {
	return _GenericHandler.Contract.ExecuteProposal(&_GenericHandler.TransactOpts, resourceID, data)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0xe248cff2.
//
// Solidity: function executeProposal(bytes32 resourceID, bytes data) returns()
func (_GenericHandler *GenericHandlerTransactorSession) ExecuteProposal(resourceID [32]byte, data []byte) (*types.Transaction, error) {
	return _GenericHandler.Contract.ExecuteProposal(&_GenericHandler.TransactOpts, resourceID, data)
}

// SetResource is a paid mutator transaction binding the contract method 0xde319d99.
//
// Solidity: function setResource(bytes32 resourceID, address contractAddress, bytes4 depositFunctionSig, uint256 depositFunctionDepositerOffset, bytes4 executeFunctionSig) returns()
func (_GenericHandler *GenericHandlerTransactor) SetResource(opts *bind.TransactOpts, resourceID [32]byte, contractAddress common.Address, depositFunctionSig [4]byte, depositFunctionDepositerOffset *big.Int, executeFunctionSig [4]byte) (*types.Transaction, error) {
	return _GenericHandler.contract.Transact(opts, "setResource", resourceID, contractAddress, depositFunctionSig, depositFunctionDepositerOffset, executeFunctionSig)
}

// SetResource is a paid mutator transaction binding the contract method 0xde319d99.
//
// Solidity: function setResource(bytes32 resourceID, address contractAddress, bytes4 depositFunctionSig, uint256 depositFunctionDepositerOffset, bytes4 executeFunctionSig) returns()
func (_GenericHandler *GenericHandlerSession) SetResource(resourceID [32]byte, contractAddress common.Address, depositFunctionSig [4]byte, depositFunctionDepositerOffset *big.Int, executeFunctionSig [4]byte) (*types.Transaction, error) {
	return _GenericHandler.Contract.SetResource(&_GenericHandler.TransactOpts, resourceID, contractAddress, depositFunctionSig, depositFunctionDepositerOffset, executeFunctionSig)
}

// SetResource is a paid mutator transaction binding the contract method 0xde319d99.
//
// Solidity: function setResource(bytes32 resourceID, address contractAddress, bytes4 depositFunctionSig, uint256 depositFunctionDepositerOffset, bytes4 executeFunctionSig) returns()
func (_GenericHandler *GenericHandlerTransactorSession) SetResource(resourceID [32]byte, contractAddress common.Address, depositFunctionSig [4]byte, depositFunctionDepositerOffset *big.Int, executeFunctionSig [4]byte) (*types.Transaction, error) {
	return _GenericHandler.Contract.SetResource(&_GenericHandler.TransactOpts, resourceID, contractAddress, depositFunctionSig, depositFunctionDepositerOffset, executeFunctionSig)
}
