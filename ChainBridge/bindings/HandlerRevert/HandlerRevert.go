// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package HandlerRevert

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

// HandlerRevertMetaData contains all meta data concerning the HandlerRevert contract.
var HandlerRevertMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridgeAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"_bridgeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_burnList\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_contractWhitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_lockMintUnlockList\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"_resourceIDToTokenContractAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_tokenContractAddressToResourceID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setBurnable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setLockMintUnlockable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setResource\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"executeProposal\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"virtualIncreaseBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561001057600080fd5b5060405161083b38038061083b8339818101604052602081101561003357600080fd5b5051606081901b6001600160601b0319166080526001600160a01b03166107cf61006c600039806103ad528061049852506107cf6000f3fe608060405234801561001057600080fd5b50600436106100b35760003560e01c80637f79bea8116100715780637f79bea814610227578063b1f064631461024d578063b8fa373614610273578063c238eea11461029f578063c8ba6c87146102bc578063e248cff2146102f4576100b3565b8062d8e7ee146100b857806307b7ed99146100e05780630968f264146101065780630a6d55d8146101ac578063318c136e146101e55780636a70d081146101ed575b600080fd5b6100de600480360360208110156100ce57600080fd5b50356001600160a01b031661036b565b005b6100de600480360360208110156100f657600080fd5b50356001600160a01b031661037f565b6100de6004803603602081101561011c57600080fd5b81019060208101813564010000000081111561013757600080fd5b82018360208201111561014957600080fd5b8035906020019184600183028401116401000000008311171561016b57600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061037c945050505050565b6101c9600480360360208110156101c257600080fd5b5035610390565b604080516001600160a01b039092168252519081900360200190f35b6101c96103ab565b6102136004803603602081101561020357600080fd5b50356001600160a01b03166103cf565b604080519115158252519081900360200190f35b6102136004803603602081101561023d57600080fd5b50356001600160a01b03166103e4565b6102136004803603602081101561026357600080fd5b50356001600160a01b03166103f9565b6100de6004803603604081101561028957600080fd5b50803590602001356001600160a01b031661040e565b6100de600480360360208110156102b557600080fd5b5035610424565b6102e2600480360360208110156102d257600080fd5b50356001600160a01b0316610429565b60408051918252519081900360200190f35b6100de6004803603604081101561030a57600080fd5b8135919081019060408101602082013564010000000081111561032c57600080fd5b82018360208201111561033e57600080fd5b8035906020019184600183028401116401000000008311171561036057600080fd5b50909250905061043b565b61037361048d565b61037c8161050c565b50565b61038761048d565b61037c816105df565b6000602081905290815260409020546001600160a01b031681565b7f000000000000000000000000000000000000000000000000000000000000000081565b60036020526000908152604090205460ff1681565b60026020526000908152604090205460ff1681565b60046020526000908152604090205460ff1681565b61041661048d565b61042082826106b2565b5050565b600555565b60016020526000908152604090205481565b600554610488576040805162461bcd60e51b815260206004820152601660248201527514dbdb595d1a1a5b99c8189859081a185c1c195b995960521b604482015290519081900360640190fd5b505050565b336001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461050a576040805162461bcd60e51b815260206004820152601e60248201527f73656e646572206d7573742062652062726964676520636f6e74726163740000604482015290519081900360640190fd5b565b6001600160a01b03811660009081526002602052604090205460ff166105635760405162461bcd60e51b81526004018080602001828103825260248152602001806107766024913960400191505060405180910390fd5b6001600160a01b03811660009081526003602052604090205460ff16156105bb5760405162461bcd60e51b815260040180806020018281038252603b81526020018061073b603b913960400191505060405180910390fd5b6001600160a01b03166000908152600460205260409020805460ff19166001179055565b6001600160a01b03811660009081526002602052604090205460ff166106365760405162461bcd60e51b81526004018080602001828103825260248152602001806107766024913960400191505060405180910390fd5b6001600160a01b03811660009081526004602052604090205460ff161561068e5760405162461bcd60e51b815260040180806020018281038252603b815260200180610700603b913960400191505060405180910390fd5b6001600160a01b03166000908152600360205260409020805460ff19166001179055565b60008281526020818152604080832080546001600160a01b039095166001600160a01b0319909516851790559282526001808252838320949094556002905220805460ff1916909117905556fe70726f766964656420636f6e74726163742063616e6e6f74206265206c6f636b4d696e74556e6c6f636b61626c6520616e64206275726e61626c6570726f766964656420636f6e74726163742063616e6e6f74206265206275726e61626c6520616e64206c6f636b4d696e74556e6c6f636b61626c6570726f766964656420636f6e7472616374206973206e6f742077686974656c6973746564a2646970667358221220fd71e9508bf8d8be1f55867a1a161f43e9ae6e9697813996af36282118db9fe064736f6c634300060c0033",
}

// HandlerRevertABI is the input ABI used to generate the binding from.
// Deprecated: Use HandlerRevertMetaData.ABI instead.
var HandlerRevertABI = HandlerRevertMetaData.ABI

// HandlerRevertBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use HandlerRevertMetaData.Bin instead.
var HandlerRevertBin = HandlerRevertMetaData.Bin

// DeployHandlerRevert deploys a new Ethereum contract, binding an instance of HandlerRevert to it.
func DeployHandlerRevert(auth *bind.TransactOpts, backend bind.ContractBackend, bridgeAddress common.Address) (common.Address, *types.Transaction, *HandlerRevert, error) {
	parsed, err := HandlerRevertMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(HandlerRevertBin), backend, bridgeAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &HandlerRevert{HandlerRevertCaller: HandlerRevertCaller{contract: contract}, HandlerRevertTransactor: HandlerRevertTransactor{contract: contract}, HandlerRevertFilterer: HandlerRevertFilterer{contract: contract}}, nil
}

// HandlerRevert is an auto generated Go binding around an Ethereum contract.
type HandlerRevert struct {
	HandlerRevertCaller     // Read-only binding to the contract
	HandlerRevertTransactor // Write-only binding to the contract
	HandlerRevertFilterer   // Log filterer for contract events
}

// HandlerRevertCaller is an auto generated read-only Go binding around an Ethereum contract.
type HandlerRevertCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HandlerRevertTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HandlerRevertTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HandlerRevertFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HandlerRevertFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HandlerRevertSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HandlerRevertSession struct {
	Contract     *HandlerRevert    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HandlerRevertCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HandlerRevertCallerSession struct {
	Contract *HandlerRevertCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// HandlerRevertTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HandlerRevertTransactorSession struct {
	Contract     *HandlerRevertTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// HandlerRevertRaw is an auto generated low-level Go binding around an Ethereum contract.
type HandlerRevertRaw struct {
	Contract *HandlerRevert // Generic contract binding to access the raw methods on
}

// HandlerRevertCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HandlerRevertCallerRaw struct {
	Contract *HandlerRevertCaller // Generic read-only contract binding to access the raw methods on
}

// HandlerRevertTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HandlerRevertTransactorRaw struct {
	Contract *HandlerRevertTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHandlerRevert creates a new instance of HandlerRevert, bound to a specific deployed contract.
func NewHandlerRevert(address common.Address, backend bind.ContractBackend) (*HandlerRevert, error) {
	contract, err := bindHandlerRevert(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HandlerRevert{HandlerRevertCaller: HandlerRevertCaller{contract: contract}, HandlerRevertTransactor: HandlerRevertTransactor{contract: contract}, HandlerRevertFilterer: HandlerRevertFilterer{contract: contract}}, nil
}

// NewHandlerRevertCaller creates a new read-only instance of HandlerRevert, bound to a specific deployed contract.
func NewHandlerRevertCaller(address common.Address, caller bind.ContractCaller) (*HandlerRevertCaller, error) {
	contract, err := bindHandlerRevert(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HandlerRevertCaller{contract: contract}, nil
}

// NewHandlerRevertTransactor creates a new write-only instance of HandlerRevert, bound to a specific deployed contract.
func NewHandlerRevertTransactor(address common.Address, transactor bind.ContractTransactor) (*HandlerRevertTransactor, error) {
	contract, err := bindHandlerRevert(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HandlerRevertTransactor{contract: contract}, nil
}

// NewHandlerRevertFilterer creates a new log filterer instance of HandlerRevert, bound to a specific deployed contract.
func NewHandlerRevertFilterer(address common.Address, filterer bind.ContractFilterer) (*HandlerRevertFilterer, error) {
	contract, err := bindHandlerRevert(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HandlerRevertFilterer{contract: contract}, nil
}

// bindHandlerRevert binds a generic wrapper to an already deployed contract.
func bindHandlerRevert(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HandlerRevertABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HandlerRevert *HandlerRevertRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HandlerRevert.Contract.HandlerRevertCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HandlerRevert *HandlerRevertRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HandlerRevert.Contract.HandlerRevertTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HandlerRevert *HandlerRevertRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HandlerRevert.Contract.HandlerRevertTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HandlerRevert *HandlerRevertCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HandlerRevert.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HandlerRevert *HandlerRevertTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HandlerRevert.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HandlerRevert *HandlerRevertTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HandlerRevert.Contract.contract.Transact(opts, method, params...)
}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_HandlerRevert *HandlerRevertCaller) BridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HandlerRevert.contract.Call(opts, &out, "_bridgeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_HandlerRevert *HandlerRevertSession) BridgeAddress() (common.Address, error) {
	return _HandlerRevert.Contract.BridgeAddress(&_HandlerRevert.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_HandlerRevert *HandlerRevertCallerSession) BridgeAddress() (common.Address, error) {
	return _HandlerRevert.Contract.BridgeAddress(&_HandlerRevert.CallOpts)
}

// BurnList is a free data retrieval call binding the contract method 0x6a70d081.
//
// Solidity: function _burnList(address ) view returns(bool)
func (_HandlerRevert *HandlerRevertCaller) BurnList(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _HandlerRevert.contract.Call(opts, &out, "_burnList", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BurnList is a free data retrieval call binding the contract method 0x6a70d081.
//
// Solidity: function _burnList(address ) view returns(bool)
func (_HandlerRevert *HandlerRevertSession) BurnList(arg0 common.Address) (bool, error) {
	return _HandlerRevert.Contract.BurnList(&_HandlerRevert.CallOpts, arg0)
}

// BurnList is a free data retrieval call binding the contract method 0x6a70d081.
//
// Solidity: function _burnList(address ) view returns(bool)
func (_HandlerRevert *HandlerRevertCallerSession) BurnList(arg0 common.Address) (bool, error) {
	return _HandlerRevert.Contract.BurnList(&_HandlerRevert.CallOpts, arg0)
}

// ContractWhitelist is a free data retrieval call binding the contract method 0x7f79bea8.
//
// Solidity: function _contractWhitelist(address ) view returns(bool)
func (_HandlerRevert *HandlerRevertCaller) ContractWhitelist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _HandlerRevert.contract.Call(opts, &out, "_contractWhitelist", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ContractWhitelist is a free data retrieval call binding the contract method 0x7f79bea8.
//
// Solidity: function _contractWhitelist(address ) view returns(bool)
func (_HandlerRevert *HandlerRevertSession) ContractWhitelist(arg0 common.Address) (bool, error) {
	return _HandlerRevert.Contract.ContractWhitelist(&_HandlerRevert.CallOpts, arg0)
}

// ContractWhitelist is a free data retrieval call binding the contract method 0x7f79bea8.
//
// Solidity: function _contractWhitelist(address ) view returns(bool)
func (_HandlerRevert *HandlerRevertCallerSession) ContractWhitelist(arg0 common.Address) (bool, error) {
	return _HandlerRevert.Contract.ContractWhitelist(&_HandlerRevert.CallOpts, arg0)
}

// LockMintUnlockList is a free data retrieval call binding the contract method 0xb1f06463.
//
// Solidity: function _lockMintUnlockList(address ) view returns(bool)
func (_HandlerRevert *HandlerRevertCaller) LockMintUnlockList(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _HandlerRevert.contract.Call(opts, &out, "_lockMintUnlockList", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// LockMintUnlockList is a free data retrieval call binding the contract method 0xb1f06463.
//
// Solidity: function _lockMintUnlockList(address ) view returns(bool)
func (_HandlerRevert *HandlerRevertSession) LockMintUnlockList(arg0 common.Address) (bool, error) {
	return _HandlerRevert.Contract.LockMintUnlockList(&_HandlerRevert.CallOpts, arg0)
}

// LockMintUnlockList is a free data retrieval call binding the contract method 0xb1f06463.
//
// Solidity: function _lockMintUnlockList(address ) view returns(bool)
func (_HandlerRevert *HandlerRevertCallerSession) LockMintUnlockList(arg0 common.Address) (bool, error) {
	return _HandlerRevert.Contract.LockMintUnlockList(&_HandlerRevert.CallOpts, arg0)
}

// ResourceIDToTokenContractAddress is a free data retrieval call binding the contract method 0x0a6d55d8.
//
// Solidity: function _resourceIDToTokenContractAddress(bytes32 ) view returns(address)
func (_HandlerRevert *HandlerRevertCaller) ResourceIDToTokenContractAddress(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _HandlerRevert.contract.Call(opts, &out, "_resourceIDToTokenContractAddress", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ResourceIDToTokenContractAddress is a free data retrieval call binding the contract method 0x0a6d55d8.
//
// Solidity: function _resourceIDToTokenContractAddress(bytes32 ) view returns(address)
func (_HandlerRevert *HandlerRevertSession) ResourceIDToTokenContractAddress(arg0 [32]byte) (common.Address, error) {
	return _HandlerRevert.Contract.ResourceIDToTokenContractAddress(&_HandlerRevert.CallOpts, arg0)
}

// ResourceIDToTokenContractAddress is a free data retrieval call binding the contract method 0x0a6d55d8.
//
// Solidity: function _resourceIDToTokenContractAddress(bytes32 ) view returns(address)
func (_HandlerRevert *HandlerRevertCallerSession) ResourceIDToTokenContractAddress(arg0 [32]byte) (common.Address, error) {
	return _HandlerRevert.Contract.ResourceIDToTokenContractAddress(&_HandlerRevert.CallOpts, arg0)
}

// TokenContractAddressToResourceID is a free data retrieval call binding the contract method 0xc8ba6c87.
//
// Solidity: function _tokenContractAddressToResourceID(address ) view returns(bytes32)
func (_HandlerRevert *HandlerRevertCaller) TokenContractAddressToResourceID(opts *bind.CallOpts, arg0 common.Address) ([32]byte, error) {
	var out []interface{}
	err := _HandlerRevert.contract.Call(opts, &out, "_tokenContractAddressToResourceID", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TokenContractAddressToResourceID is a free data retrieval call binding the contract method 0xc8ba6c87.
//
// Solidity: function _tokenContractAddressToResourceID(address ) view returns(bytes32)
func (_HandlerRevert *HandlerRevertSession) TokenContractAddressToResourceID(arg0 common.Address) ([32]byte, error) {
	return _HandlerRevert.Contract.TokenContractAddressToResourceID(&_HandlerRevert.CallOpts, arg0)
}

// TokenContractAddressToResourceID is a free data retrieval call binding the contract method 0xc8ba6c87.
//
// Solidity: function _tokenContractAddressToResourceID(address ) view returns(bytes32)
func (_HandlerRevert *HandlerRevertCallerSession) TokenContractAddressToResourceID(arg0 common.Address) ([32]byte, error) {
	return _HandlerRevert.Contract.TokenContractAddressToResourceID(&_HandlerRevert.CallOpts, arg0)
}

// ExecuteProposal is a free data retrieval call binding the contract method 0xe248cff2.
//
// Solidity: function executeProposal(bytes32 , bytes ) view returns()
func (_HandlerRevert *HandlerRevertCaller) ExecuteProposal(opts *bind.CallOpts, arg0 [32]byte, arg1 []byte) error {
	var out []interface{}
	err := _HandlerRevert.contract.Call(opts, &out, "executeProposal", arg0, arg1)

	if err != nil {
		return err
	}

	return err

}

// ExecuteProposal is a free data retrieval call binding the contract method 0xe248cff2.
//
// Solidity: function executeProposal(bytes32 , bytes ) view returns()
func (_HandlerRevert *HandlerRevertSession) ExecuteProposal(arg0 [32]byte, arg1 []byte) error {
	return _HandlerRevert.Contract.ExecuteProposal(&_HandlerRevert.CallOpts, arg0, arg1)
}

// ExecuteProposal is a free data retrieval call binding the contract method 0xe248cff2.
//
// Solidity: function executeProposal(bytes32 , bytes ) view returns()
func (_HandlerRevert *HandlerRevertCallerSession) ExecuteProposal(arg0 [32]byte, arg1 []byte) error {
	return _HandlerRevert.Contract.ExecuteProposal(&_HandlerRevert.CallOpts, arg0, arg1)
}

// SetBurnable is a paid mutator transaction binding the contract method 0x07b7ed99.
//
// Solidity: function setBurnable(address contractAddress) returns()
func (_HandlerRevert *HandlerRevertTransactor) SetBurnable(opts *bind.TransactOpts, contractAddress common.Address) (*types.Transaction, error) {
	return _HandlerRevert.contract.Transact(opts, "setBurnable", contractAddress)
}

// SetBurnable is a paid mutator transaction binding the contract method 0x07b7ed99.
//
// Solidity: function setBurnable(address contractAddress) returns()
func (_HandlerRevert *HandlerRevertSession) SetBurnable(contractAddress common.Address) (*types.Transaction, error) {
	return _HandlerRevert.Contract.SetBurnable(&_HandlerRevert.TransactOpts, contractAddress)
}

// SetBurnable is a paid mutator transaction binding the contract method 0x07b7ed99.
//
// Solidity: function setBurnable(address contractAddress) returns()
func (_HandlerRevert *HandlerRevertTransactorSession) SetBurnable(contractAddress common.Address) (*types.Transaction, error) {
	return _HandlerRevert.Contract.SetBurnable(&_HandlerRevert.TransactOpts, contractAddress)
}

// SetLockMintUnlockable is a paid mutator transaction binding the contract method 0x00d8e7ee.
//
// Solidity: function setLockMintUnlockable(address contractAddress) returns()
func (_HandlerRevert *HandlerRevertTransactor) SetLockMintUnlockable(opts *bind.TransactOpts, contractAddress common.Address) (*types.Transaction, error) {
	return _HandlerRevert.contract.Transact(opts, "setLockMintUnlockable", contractAddress)
}

// SetLockMintUnlockable is a paid mutator transaction binding the contract method 0x00d8e7ee.
//
// Solidity: function setLockMintUnlockable(address contractAddress) returns()
func (_HandlerRevert *HandlerRevertSession) SetLockMintUnlockable(contractAddress common.Address) (*types.Transaction, error) {
	return _HandlerRevert.Contract.SetLockMintUnlockable(&_HandlerRevert.TransactOpts, contractAddress)
}

// SetLockMintUnlockable is a paid mutator transaction binding the contract method 0x00d8e7ee.
//
// Solidity: function setLockMintUnlockable(address contractAddress) returns()
func (_HandlerRevert *HandlerRevertTransactorSession) SetLockMintUnlockable(contractAddress common.Address) (*types.Transaction, error) {
	return _HandlerRevert.Contract.SetLockMintUnlockable(&_HandlerRevert.TransactOpts, contractAddress)
}

// SetResource is a paid mutator transaction binding the contract method 0xb8fa3736.
//
// Solidity: function setResource(bytes32 resourceID, address contractAddress) returns()
func (_HandlerRevert *HandlerRevertTransactor) SetResource(opts *bind.TransactOpts, resourceID [32]byte, contractAddress common.Address) (*types.Transaction, error) {
	return _HandlerRevert.contract.Transact(opts, "setResource", resourceID, contractAddress)
}

// SetResource is a paid mutator transaction binding the contract method 0xb8fa3736.
//
// Solidity: function setResource(bytes32 resourceID, address contractAddress) returns()
func (_HandlerRevert *HandlerRevertSession) SetResource(resourceID [32]byte, contractAddress common.Address) (*types.Transaction, error) {
	return _HandlerRevert.Contract.SetResource(&_HandlerRevert.TransactOpts, resourceID, contractAddress)
}

// SetResource is a paid mutator transaction binding the contract method 0xb8fa3736.
//
// Solidity: function setResource(bytes32 resourceID, address contractAddress) returns()
func (_HandlerRevert *HandlerRevertTransactorSession) SetResource(resourceID [32]byte, contractAddress common.Address) (*types.Transaction, error) {
	return _HandlerRevert.Contract.SetResource(&_HandlerRevert.TransactOpts, resourceID, contractAddress)
}

// VirtualIncreaseBalance is a paid mutator transaction binding the contract method 0xc238eea1.
//
// Solidity: function virtualIncreaseBalance(uint256 amount) returns()
func (_HandlerRevert *HandlerRevertTransactor) VirtualIncreaseBalance(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _HandlerRevert.contract.Transact(opts, "virtualIncreaseBalance", amount)
}

// VirtualIncreaseBalance is a paid mutator transaction binding the contract method 0xc238eea1.
//
// Solidity: function virtualIncreaseBalance(uint256 amount) returns()
func (_HandlerRevert *HandlerRevertSession) VirtualIncreaseBalance(amount *big.Int) (*types.Transaction, error) {
	return _HandlerRevert.Contract.VirtualIncreaseBalance(&_HandlerRevert.TransactOpts, amount)
}

// VirtualIncreaseBalance is a paid mutator transaction binding the contract method 0xc238eea1.
//
// Solidity: function virtualIncreaseBalance(uint256 amount) returns()
func (_HandlerRevert *HandlerRevertTransactorSession) VirtualIncreaseBalance(amount *big.Int) (*types.Transaction, error) {
	return _HandlerRevert.Contract.VirtualIncreaseBalance(&_HandlerRevert.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x0968f264.
//
// Solidity: function withdraw(bytes data) returns()
func (_HandlerRevert *HandlerRevertTransactor) Withdraw(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _HandlerRevert.contract.Transact(opts, "withdraw", data)
}

// Withdraw is a paid mutator transaction binding the contract method 0x0968f264.
//
// Solidity: function withdraw(bytes data) returns()
func (_HandlerRevert *HandlerRevertSession) Withdraw(data []byte) (*types.Transaction, error) {
	return _HandlerRevert.Contract.Withdraw(&_HandlerRevert.TransactOpts, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0x0968f264.
//
// Solidity: function withdraw(bytes data) returns()
func (_HandlerRevert *HandlerRevertTransactorSession) Withdraw(data []byte) (*types.Transaction, error) {
	return _HandlerRevert.Contract.Withdraw(&_HandlerRevert.TransactOpts, data)
}
