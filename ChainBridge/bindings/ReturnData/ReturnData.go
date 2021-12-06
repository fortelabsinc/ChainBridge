// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ReturnData

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

// ReturnDataMetaData contains all meta data concerning the ReturnData contract.
var ReturnDataMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"argument\",\"type\":\"string\"}],\"name\":\"returnData\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"response\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061011c806100206000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c8063e5f92e8614602d575b600080fd5b60cd60048036036020811015604157600080fd5b810190602081018135640100000000811115605b57600080fd5b820183602082011115606c57600080fd5b80359060200191846001830284011164010000000083111715608d57600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955060df945050505050565b60408051918252519081900360200190f35b602001519056fea26469706673582212205521168a97bf324910513a0f9c4eac18776628c5b50313d52f04f308ba0b0c2064736f6c634300060c0033",
}

// ReturnDataABI is the input ABI used to generate the binding from.
// Deprecated: Use ReturnDataMetaData.ABI instead.
var ReturnDataABI = ReturnDataMetaData.ABI

// ReturnDataBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ReturnDataMetaData.Bin instead.
var ReturnDataBin = ReturnDataMetaData.Bin

// DeployReturnData deploys a new Ethereum contract, binding an instance of ReturnData to it.
func DeployReturnData(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ReturnData, error) {
	parsed, err := ReturnDataMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ReturnDataBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ReturnData{ReturnDataCaller: ReturnDataCaller{contract: contract}, ReturnDataTransactor: ReturnDataTransactor{contract: contract}, ReturnDataFilterer: ReturnDataFilterer{contract: contract}}, nil
}

// ReturnData is an auto generated Go binding around an Ethereum contract.
type ReturnData struct {
	ReturnDataCaller     // Read-only binding to the contract
	ReturnDataTransactor // Write-only binding to the contract
	ReturnDataFilterer   // Log filterer for contract events
}

// ReturnDataCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReturnDataCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReturnDataTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReturnDataTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReturnDataFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReturnDataFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReturnDataSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReturnDataSession struct {
	Contract     *ReturnData       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ReturnDataCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReturnDataCallerSession struct {
	Contract *ReturnDataCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ReturnDataTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReturnDataTransactorSession struct {
	Contract     *ReturnDataTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ReturnDataRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReturnDataRaw struct {
	Contract *ReturnData // Generic contract binding to access the raw methods on
}

// ReturnDataCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReturnDataCallerRaw struct {
	Contract *ReturnDataCaller // Generic read-only contract binding to access the raw methods on
}

// ReturnDataTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReturnDataTransactorRaw struct {
	Contract *ReturnDataTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReturnData creates a new instance of ReturnData, bound to a specific deployed contract.
func NewReturnData(address common.Address, backend bind.ContractBackend) (*ReturnData, error) {
	contract, err := bindReturnData(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ReturnData{ReturnDataCaller: ReturnDataCaller{contract: contract}, ReturnDataTransactor: ReturnDataTransactor{contract: contract}, ReturnDataFilterer: ReturnDataFilterer{contract: contract}}, nil
}

// NewReturnDataCaller creates a new read-only instance of ReturnData, bound to a specific deployed contract.
func NewReturnDataCaller(address common.Address, caller bind.ContractCaller) (*ReturnDataCaller, error) {
	contract, err := bindReturnData(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReturnDataCaller{contract: contract}, nil
}

// NewReturnDataTransactor creates a new write-only instance of ReturnData, bound to a specific deployed contract.
func NewReturnDataTransactor(address common.Address, transactor bind.ContractTransactor) (*ReturnDataTransactor, error) {
	contract, err := bindReturnData(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReturnDataTransactor{contract: contract}, nil
}

// NewReturnDataFilterer creates a new log filterer instance of ReturnData, bound to a specific deployed contract.
func NewReturnDataFilterer(address common.Address, filterer bind.ContractFilterer) (*ReturnDataFilterer, error) {
	contract, err := bindReturnData(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReturnDataFilterer{contract: contract}, nil
}

// bindReturnData binds a generic wrapper to an already deployed contract.
func bindReturnData(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ReturnDataABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReturnData *ReturnDataRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReturnData.Contract.ReturnDataCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReturnData *ReturnDataRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReturnData.Contract.ReturnDataTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReturnData *ReturnDataRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReturnData.Contract.ReturnDataTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReturnData *ReturnDataCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReturnData.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReturnData *ReturnDataTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReturnData.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReturnData *ReturnDataTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReturnData.Contract.contract.Transact(opts, method, params...)
}

// ReturnData is a free data retrieval call binding the contract method 0xe5f92e86.
//
// Solidity: function returnData(string argument) pure returns(bytes32 response)
func (_ReturnData *ReturnDataCaller) ReturnData(opts *bind.CallOpts, argument string) ([32]byte, error) {
	var out []interface{}
	err := _ReturnData.contract.Call(opts, &out, "returnData", argument)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ReturnData is a free data retrieval call binding the contract method 0xe5f92e86.
//
// Solidity: function returnData(string argument) pure returns(bytes32 response)
func (_ReturnData *ReturnDataSession) ReturnData(argument string) ([32]byte, error) {
	return _ReturnData.Contract.ReturnData(&_ReturnData.CallOpts, argument)
}

// ReturnData is a free data retrieval call binding the contract method 0xe5f92e86.
//
// Solidity: function returnData(string argument) pure returns(bytes32 response)
func (_ReturnData *ReturnDataCallerSession) ReturnData(argument string) ([32]byte, error) {
	return _ReturnData.Contract.ReturnData(&_ReturnData.CallOpts, argument)
}
