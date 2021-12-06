// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package HandlerHelpers

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

// HandlerHelpersMetaData contains all meta data concerning the HandlerHelpers contract.
var HandlerHelpersMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridgeAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"_bridgeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_burnList\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_contractWhitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_lockMintUnlockList\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"_resourceIDToTokenContractAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_tokenContractAddressToResourceID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setResource\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setBurnable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setLockMintUnlockable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561001057600080fd5b5060405161073a38038061073a8339818101604052602081101561003357600080fd5b5051606081901b6001600160601b0319166080526001600160a01b03166106ce61006c60003980610303528061039752506106ce6000f3fe608060405234801561001057600080fd5b506004361061009d5760003560e01c80636a70d081116100665780636a70d081146101d75780637f79bea814610211578063b1f0646314610237578063b8fa37361461025d578063c8ba6c87146102895761009d565b8062d8e7ee146100a257806307b7ed99146100ca5780630968f264146100f05780630a6d55d814610196578063318c136e146101cf575b600080fd5b6100c8600480360360208110156100b857600080fd5b50356001600160a01b03166102c1565b005b6100c8600480360360208110156100e057600080fd5b50356001600160a01b03166102d5565b6100c86004803603602081101561010657600080fd5b81019060208101813564010000000081111561012157600080fd5b82018360208201111561013357600080fd5b8035906020019184600183028401116401000000008311171561015557600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506102d2945050505050565b6101b3600480360360208110156101ac57600080fd5b50356102e6565b604080516001600160a01b039092168252519081900360200190f35b6101b3610301565b6101fd600480360360208110156101ed57600080fd5b50356001600160a01b0316610325565b604080519115158252519081900360200190f35b6101fd6004803603602081101561022757600080fd5b50356001600160a01b031661033a565b6101fd6004803603602081101561024d57600080fd5b50356001600160a01b031661034f565b6100c86004803603604081101561027357600080fd5b50803590602001356001600160a01b0316610364565b6102af6004803603602081101561029f57600080fd5b50356001600160a01b031661037a565b60408051918252519081900360200190f35b6102c961038c565b6102d28161040b565b50565b6102dd61038c565b6102d2816104de565b6000602081905290815260409020546001600160a01b031681565b7f000000000000000000000000000000000000000000000000000000000000000081565b60036020526000908152604090205460ff1681565b60026020526000908152604090205460ff1681565b60046020526000908152604090205460ff1681565b61036c61038c565b61037682826105b1565b5050565b60016020526000908152604090205481565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610409576040805162461bcd60e51b815260206004820152601e60248201527f73656e646572206d7573742062652062726964676520636f6e74726163740000604482015290519081900360640190fd5b565b6001600160a01b03811660009081526002602052604090205460ff166104625760405162461bcd60e51b81526004018080602001828103825260248152602001806106756024913960400191505060405180910390fd5b6001600160a01b03811660009081526003602052604090205460ff16156104ba5760405162461bcd60e51b815260040180806020018281038252603b81526020018061063a603b913960400191505060405180910390fd5b6001600160a01b03166000908152600460205260409020805460ff19166001179055565b6001600160a01b03811660009081526002602052604090205460ff166105355760405162461bcd60e51b81526004018080602001828103825260248152602001806106756024913960400191505060405180910390fd5b6001600160a01b03811660009081526004602052604090205460ff161561058d5760405162461bcd60e51b815260040180806020018281038252603b8152602001806105ff603b913960400191505060405180910390fd5b6001600160a01b03166000908152600360205260409020805460ff19166001179055565b60008281526020818152604080832080546001600160a01b039095166001600160a01b0319909516851790559282526001808252838320949094556002905220805460ff1916909117905556fe70726f766964656420636f6e74726163742063616e6e6f74206265206c6f636b4d696e74556e6c6f636b61626c6520616e64206275726e61626c6570726f766964656420636f6e74726163742063616e6e6f74206265206275726e61626c6520616e64206c6f636b4d696e74556e6c6f636b61626c6570726f766964656420636f6e7472616374206973206e6f742077686974656c6973746564a2646970667358221220f154db92306de8a7ff32f5e7bb0be8f2e1ad83e0a86d2e28a584a9355215eb3964736f6c634300060c0033",
}

// HandlerHelpersABI is the input ABI used to generate the binding from.
// Deprecated: Use HandlerHelpersMetaData.ABI instead.
var HandlerHelpersABI = HandlerHelpersMetaData.ABI

// HandlerHelpersBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use HandlerHelpersMetaData.Bin instead.
var HandlerHelpersBin = HandlerHelpersMetaData.Bin

// DeployHandlerHelpers deploys a new Ethereum contract, binding an instance of HandlerHelpers to it.
func DeployHandlerHelpers(auth *bind.TransactOpts, backend bind.ContractBackend, bridgeAddress common.Address) (common.Address, *types.Transaction, *HandlerHelpers, error) {
	parsed, err := HandlerHelpersMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(HandlerHelpersBin), backend, bridgeAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &HandlerHelpers{HandlerHelpersCaller: HandlerHelpersCaller{contract: contract}, HandlerHelpersTransactor: HandlerHelpersTransactor{contract: contract}, HandlerHelpersFilterer: HandlerHelpersFilterer{contract: contract}}, nil
}

// HandlerHelpers is an auto generated Go binding around an Ethereum contract.
type HandlerHelpers struct {
	HandlerHelpersCaller     // Read-only binding to the contract
	HandlerHelpersTransactor // Write-only binding to the contract
	HandlerHelpersFilterer   // Log filterer for contract events
}

// HandlerHelpersCaller is an auto generated read-only Go binding around an Ethereum contract.
type HandlerHelpersCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HandlerHelpersTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HandlerHelpersTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HandlerHelpersFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HandlerHelpersFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HandlerHelpersSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HandlerHelpersSession struct {
	Contract     *HandlerHelpers   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HandlerHelpersCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HandlerHelpersCallerSession struct {
	Contract *HandlerHelpersCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// HandlerHelpersTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HandlerHelpersTransactorSession struct {
	Contract     *HandlerHelpersTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// HandlerHelpersRaw is an auto generated low-level Go binding around an Ethereum contract.
type HandlerHelpersRaw struct {
	Contract *HandlerHelpers // Generic contract binding to access the raw methods on
}

// HandlerHelpersCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HandlerHelpersCallerRaw struct {
	Contract *HandlerHelpersCaller // Generic read-only contract binding to access the raw methods on
}

// HandlerHelpersTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HandlerHelpersTransactorRaw struct {
	Contract *HandlerHelpersTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHandlerHelpers creates a new instance of HandlerHelpers, bound to a specific deployed contract.
func NewHandlerHelpers(address common.Address, backend bind.ContractBackend) (*HandlerHelpers, error) {
	contract, err := bindHandlerHelpers(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HandlerHelpers{HandlerHelpersCaller: HandlerHelpersCaller{contract: contract}, HandlerHelpersTransactor: HandlerHelpersTransactor{contract: contract}, HandlerHelpersFilterer: HandlerHelpersFilterer{contract: contract}}, nil
}

// NewHandlerHelpersCaller creates a new read-only instance of HandlerHelpers, bound to a specific deployed contract.
func NewHandlerHelpersCaller(address common.Address, caller bind.ContractCaller) (*HandlerHelpersCaller, error) {
	contract, err := bindHandlerHelpers(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HandlerHelpersCaller{contract: contract}, nil
}

// NewHandlerHelpersTransactor creates a new write-only instance of HandlerHelpers, bound to a specific deployed contract.
func NewHandlerHelpersTransactor(address common.Address, transactor bind.ContractTransactor) (*HandlerHelpersTransactor, error) {
	contract, err := bindHandlerHelpers(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HandlerHelpersTransactor{contract: contract}, nil
}

// NewHandlerHelpersFilterer creates a new log filterer instance of HandlerHelpers, bound to a specific deployed contract.
func NewHandlerHelpersFilterer(address common.Address, filterer bind.ContractFilterer) (*HandlerHelpersFilterer, error) {
	contract, err := bindHandlerHelpers(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HandlerHelpersFilterer{contract: contract}, nil
}

// bindHandlerHelpers binds a generic wrapper to an already deployed contract.
func bindHandlerHelpers(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HandlerHelpersABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HandlerHelpers *HandlerHelpersRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HandlerHelpers.Contract.HandlerHelpersCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HandlerHelpers *HandlerHelpersRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HandlerHelpers.Contract.HandlerHelpersTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HandlerHelpers *HandlerHelpersRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HandlerHelpers.Contract.HandlerHelpersTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HandlerHelpers *HandlerHelpersCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HandlerHelpers.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HandlerHelpers *HandlerHelpersTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HandlerHelpers.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HandlerHelpers *HandlerHelpersTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HandlerHelpers.Contract.contract.Transact(opts, method, params...)
}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_HandlerHelpers *HandlerHelpersCaller) BridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HandlerHelpers.contract.Call(opts, &out, "_bridgeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_HandlerHelpers *HandlerHelpersSession) BridgeAddress() (common.Address, error) {
	return _HandlerHelpers.Contract.BridgeAddress(&_HandlerHelpers.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_HandlerHelpers *HandlerHelpersCallerSession) BridgeAddress() (common.Address, error) {
	return _HandlerHelpers.Contract.BridgeAddress(&_HandlerHelpers.CallOpts)
}

// BurnList is a free data retrieval call binding the contract method 0x6a70d081.
//
// Solidity: function _burnList(address ) view returns(bool)
func (_HandlerHelpers *HandlerHelpersCaller) BurnList(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _HandlerHelpers.contract.Call(opts, &out, "_burnList", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BurnList is a free data retrieval call binding the contract method 0x6a70d081.
//
// Solidity: function _burnList(address ) view returns(bool)
func (_HandlerHelpers *HandlerHelpersSession) BurnList(arg0 common.Address) (bool, error) {
	return _HandlerHelpers.Contract.BurnList(&_HandlerHelpers.CallOpts, arg0)
}

// BurnList is a free data retrieval call binding the contract method 0x6a70d081.
//
// Solidity: function _burnList(address ) view returns(bool)
func (_HandlerHelpers *HandlerHelpersCallerSession) BurnList(arg0 common.Address) (bool, error) {
	return _HandlerHelpers.Contract.BurnList(&_HandlerHelpers.CallOpts, arg0)
}

// ContractWhitelist is a free data retrieval call binding the contract method 0x7f79bea8.
//
// Solidity: function _contractWhitelist(address ) view returns(bool)
func (_HandlerHelpers *HandlerHelpersCaller) ContractWhitelist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _HandlerHelpers.contract.Call(opts, &out, "_contractWhitelist", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ContractWhitelist is a free data retrieval call binding the contract method 0x7f79bea8.
//
// Solidity: function _contractWhitelist(address ) view returns(bool)
func (_HandlerHelpers *HandlerHelpersSession) ContractWhitelist(arg0 common.Address) (bool, error) {
	return _HandlerHelpers.Contract.ContractWhitelist(&_HandlerHelpers.CallOpts, arg0)
}

// ContractWhitelist is a free data retrieval call binding the contract method 0x7f79bea8.
//
// Solidity: function _contractWhitelist(address ) view returns(bool)
func (_HandlerHelpers *HandlerHelpersCallerSession) ContractWhitelist(arg0 common.Address) (bool, error) {
	return _HandlerHelpers.Contract.ContractWhitelist(&_HandlerHelpers.CallOpts, arg0)
}

// LockMintUnlockList is a free data retrieval call binding the contract method 0xb1f06463.
//
// Solidity: function _lockMintUnlockList(address ) view returns(bool)
func (_HandlerHelpers *HandlerHelpersCaller) LockMintUnlockList(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _HandlerHelpers.contract.Call(opts, &out, "_lockMintUnlockList", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// LockMintUnlockList is a free data retrieval call binding the contract method 0xb1f06463.
//
// Solidity: function _lockMintUnlockList(address ) view returns(bool)
func (_HandlerHelpers *HandlerHelpersSession) LockMintUnlockList(arg0 common.Address) (bool, error) {
	return _HandlerHelpers.Contract.LockMintUnlockList(&_HandlerHelpers.CallOpts, arg0)
}

// LockMintUnlockList is a free data retrieval call binding the contract method 0xb1f06463.
//
// Solidity: function _lockMintUnlockList(address ) view returns(bool)
func (_HandlerHelpers *HandlerHelpersCallerSession) LockMintUnlockList(arg0 common.Address) (bool, error) {
	return _HandlerHelpers.Contract.LockMintUnlockList(&_HandlerHelpers.CallOpts, arg0)
}

// ResourceIDToTokenContractAddress is a free data retrieval call binding the contract method 0x0a6d55d8.
//
// Solidity: function _resourceIDToTokenContractAddress(bytes32 ) view returns(address)
func (_HandlerHelpers *HandlerHelpersCaller) ResourceIDToTokenContractAddress(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _HandlerHelpers.contract.Call(opts, &out, "_resourceIDToTokenContractAddress", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ResourceIDToTokenContractAddress is a free data retrieval call binding the contract method 0x0a6d55d8.
//
// Solidity: function _resourceIDToTokenContractAddress(bytes32 ) view returns(address)
func (_HandlerHelpers *HandlerHelpersSession) ResourceIDToTokenContractAddress(arg0 [32]byte) (common.Address, error) {
	return _HandlerHelpers.Contract.ResourceIDToTokenContractAddress(&_HandlerHelpers.CallOpts, arg0)
}

// ResourceIDToTokenContractAddress is a free data retrieval call binding the contract method 0x0a6d55d8.
//
// Solidity: function _resourceIDToTokenContractAddress(bytes32 ) view returns(address)
func (_HandlerHelpers *HandlerHelpersCallerSession) ResourceIDToTokenContractAddress(arg0 [32]byte) (common.Address, error) {
	return _HandlerHelpers.Contract.ResourceIDToTokenContractAddress(&_HandlerHelpers.CallOpts, arg0)
}

// TokenContractAddressToResourceID is a free data retrieval call binding the contract method 0xc8ba6c87.
//
// Solidity: function _tokenContractAddressToResourceID(address ) view returns(bytes32)
func (_HandlerHelpers *HandlerHelpersCaller) TokenContractAddressToResourceID(opts *bind.CallOpts, arg0 common.Address) ([32]byte, error) {
	var out []interface{}
	err := _HandlerHelpers.contract.Call(opts, &out, "_tokenContractAddressToResourceID", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TokenContractAddressToResourceID is a free data retrieval call binding the contract method 0xc8ba6c87.
//
// Solidity: function _tokenContractAddressToResourceID(address ) view returns(bytes32)
func (_HandlerHelpers *HandlerHelpersSession) TokenContractAddressToResourceID(arg0 common.Address) ([32]byte, error) {
	return _HandlerHelpers.Contract.TokenContractAddressToResourceID(&_HandlerHelpers.CallOpts, arg0)
}

// TokenContractAddressToResourceID is a free data retrieval call binding the contract method 0xc8ba6c87.
//
// Solidity: function _tokenContractAddressToResourceID(address ) view returns(bytes32)
func (_HandlerHelpers *HandlerHelpersCallerSession) TokenContractAddressToResourceID(arg0 common.Address) ([32]byte, error) {
	return _HandlerHelpers.Contract.TokenContractAddressToResourceID(&_HandlerHelpers.CallOpts, arg0)
}

// SetBurnable is a paid mutator transaction binding the contract method 0x07b7ed99.
//
// Solidity: function setBurnable(address contractAddress) returns()
func (_HandlerHelpers *HandlerHelpersTransactor) SetBurnable(opts *bind.TransactOpts, contractAddress common.Address) (*types.Transaction, error) {
	return _HandlerHelpers.contract.Transact(opts, "setBurnable", contractAddress)
}

// SetBurnable is a paid mutator transaction binding the contract method 0x07b7ed99.
//
// Solidity: function setBurnable(address contractAddress) returns()
func (_HandlerHelpers *HandlerHelpersSession) SetBurnable(contractAddress common.Address) (*types.Transaction, error) {
	return _HandlerHelpers.Contract.SetBurnable(&_HandlerHelpers.TransactOpts, contractAddress)
}

// SetBurnable is a paid mutator transaction binding the contract method 0x07b7ed99.
//
// Solidity: function setBurnable(address contractAddress) returns()
func (_HandlerHelpers *HandlerHelpersTransactorSession) SetBurnable(contractAddress common.Address) (*types.Transaction, error) {
	return _HandlerHelpers.Contract.SetBurnable(&_HandlerHelpers.TransactOpts, contractAddress)
}

// SetLockMintUnlockable is a paid mutator transaction binding the contract method 0x00d8e7ee.
//
// Solidity: function setLockMintUnlockable(address contractAddress) returns()
func (_HandlerHelpers *HandlerHelpersTransactor) SetLockMintUnlockable(opts *bind.TransactOpts, contractAddress common.Address) (*types.Transaction, error) {
	return _HandlerHelpers.contract.Transact(opts, "setLockMintUnlockable", contractAddress)
}

// SetLockMintUnlockable is a paid mutator transaction binding the contract method 0x00d8e7ee.
//
// Solidity: function setLockMintUnlockable(address contractAddress) returns()
func (_HandlerHelpers *HandlerHelpersSession) SetLockMintUnlockable(contractAddress common.Address) (*types.Transaction, error) {
	return _HandlerHelpers.Contract.SetLockMintUnlockable(&_HandlerHelpers.TransactOpts, contractAddress)
}

// SetLockMintUnlockable is a paid mutator transaction binding the contract method 0x00d8e7ee.
//
// Solidity: function setLockMintUnlockable(address contractAddress) returns()
func (_HandlerHelpers *HandlerHelpersTransactorSession) SetLockMintUnlockable(contractAddress common.Address) (*types.Transaction, error) {
	return _HandlerHelpers.Contract.SetLockMintUnlockable(&_HandlerHelpers.TransactOpts, contractAddress)
}

// SetResource is a paid mutator transaction binding the contract method 0xb8fa3736.
//
// Solidity: function setResource(bytes32 resourceID, address contractAddress) returns()
func (_HandlerHelpers *HandlerHelpersTransactor) SetResource(opts *bind.TransactOpts, resourceID [32]byte, contractAddress common.Address) (*types.Transaction, error) {
	return _HandlerHelpers.contract.Transact(opts, "setResource", resourceID, contractAddress)
}

// SetResource is a paid mutator transaction binding the contract method 0xb8fa3736.
//
// Solidity: function setResource(bytes32 resourceID, address contractAddress) returns()
func (_HandlerHelpers *HandlerHelpersSession) SetResource(resourceID [32]byte, contractAddress common.Address) (*types.Transaction, error) {
	return _HandlerHelpers.Contract.SetResource(&_HandlerHelpers.TransactOpts, resourceID, contractAddress)
}

// SetResource is a paid mutator transaction binding the contract method 0xb8fa3736.
//
// Solidity: function setResource(bytes32 resourceID, address contractAddress) returns()
func (_HandlerHelpers *HandlerHelpersTransactorSession) SetResource(resourceID [32]byte, contractAddress common.Address) (*types.Transaction, error) {
	return _HandlerHelpers.Contract.SetResource(&_HandlerHelpers.TransactOpts, resourceID, contractAddress)
}

// Withdraw is a paid mutator transaction binding the contract method 0x0968f264.
//
// Solidity: function withdraw(bytes data) returns()
func (_HandlerHelpers *HandlerHelpersTransactor) Withdraw(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _HandlerHelpers.contract.Transact(opts, "withdraw", data)
}

// Withdraw is a paid mutator transaction binding the contract method 0x0968f264.
//
// Solidity: function withdraw(bytes data) returns()
func (_HandlerHelpers *HandlerHelpersSession) Withdraw(data []byte) (*types.Transaction, error) {
	return _HandlerHelpers.Contract.Withdraw(&_HandlerHelpers.TransactOpts, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0x0968f264.
//
// Solidity: function withdraw(bytes data) returns()
func (_HandlerHelpers *HandlerHelpersTransactorSession) Withdraw(data []byte) (*types.Transaction, error) {
	return _HandlerHelpers.Contract.Withdraw(&_HandlerHelpers.TransactOpts, data)
}
