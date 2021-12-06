// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ERC20Handler

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

// ERC20HandlerMetaData contains all meta data concerning the ERC20Handler contract.
var ERC20HandlerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridgeAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"_bridgeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_burnList\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_contractWhitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_lockMintUnlockList\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"_resourceIDToTokenContractAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_tokenContractAddressToResourceID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setBurnable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setLockMintUnlockable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setResource\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"depositer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"executeProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561001057600080fd5b50604051610e67380380610e6783398101604081905261002f91610044565b60601b6001600160601b031916608052610072565b600060208284031215610055578081fd5b81516001600160a01b038116811461006b578182fd5b9392505050565b60805160601c610dd36100946000398061024b528061048c5250610dd36000f3fe608060405234801561001057600080fd5b50600436106100b35760003560e01c80637f79bea8116100715780637f79bea814610144578063b07e54bb14610157578063b1f0646314610177578063b8fa37361461018a578063c8ba6c871461019d578063e248cff2146101bd576100b3565b8062d8e7ee146100b857806307b7ed99146100cd5780630968f264146100e05780630a6d55d8146100f3578063318c136e1461011c5780636a70d08114610124575b600080fd5b6100cb6100c636600461086a565b6101d0565b005b6100cb6100db36600461086a565b6101e4565b6100cb6100ee3660046109da565b6101f5565b6101066101013660046108ef565b61022e565b6040516101139190610ab2565b60405180910390f35b610106610249565b61013761013236600461086a565b61026d565b6040516101139190610b03565b61013761015236600461086a565b610282565b61016a610165366004610936565b610297565b6040516101139190610b17565b61013761018536600461086a565b610347565b6100cb610198366004610907565b61035c565b6101b06101ab36600461086a565b610372565b6040516101139190610b0e565b6100cb6101cb366004610990565b610384565b6101d8610481565b6101e1816104cb565b50565b6101ec610481565b6101e181610560565b6101fd610481565b600080600083806020019051810190610216919061088d565b919450925090506102288383836105f5565b50505050565b6000602081905290815260409020546001600160a01b031681565b7f000000000000000000000000000000000000000000000000000000000000000081565b60036020526000908152604090205460ff1681565b60026020526000908152604090205460ff1681565b60606102a1610481565b60006102af838501856108ef565b600087815260208181526040808320546001600160a01b03168084526002909252909120549192509060ff166103005760405162461bcd60e51b81526004016102f790610ce0565b60405180910390fd5b6001600160a01b03811660009081526003602052604090205460ff16156103315761032c818784610601565b61033d565b61033d8187308561065f565b5050949350505050565b60046020526000908152604090205460ff1681565b610364610481565b61036e8282610673565b5050565b60016020526000908152604090205481565b61038c610481565b600080606061039d84860186610a75565b90935091506103b26040808401908688610d28565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201829052508a81526020818152604080832054828801516001600160a01b039091168085526002909352922054959650909490935060ff16915061043590505760405162461bcd60e51b81526004016102f790610ce0565b6001600160a01b03811660009081526003602052604090205460ff161561046957610464818360601c876106c0565b610477565b610477818360601c876105f5565b5050505050505050565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146104c95760405162461bcd60e51b81526004016102f790610bdc565b565b6001600160a01b03811660009081526002602052604090205460ff166105035760405162461bcd60e51b81526004016102f790610c70565b6001600160a01b03811660009081526003602052604090205460ff161561053c5760405162461bcd60e51b81526004016102f790610c13565b6001600160a01b03166000908152600460205260409020805460ff19166001179055565b6001600160a01b03811660009081526002602052604090205460ff166105985760405162461bcd60e51b81526004016102f790610c70565b6001600160a01b03811660009081526004602052604090205460ff16156105d15760405162461bcd60e51b81526004016102f790610b4a565b6001600160a01b03166000908152600360205260409020805460ff19166001179055565b826102288184846106f0565b60405163079cc67960e41b815283906001600160a01b038216906379cc6790906106319086908690600401610aea565b600060405180830381600087803b15801561064b57600080fd5b505af1158015610477573d6000803e3d6000fd5b8361066c8185858561074b565b5050505050565b60008281526020818152604080832080546001600160a01b039095166001600160a01b0319909516851790559282526001808252838320949094556002905220805460ff19169091179055565b6040516340c10f1960e01b815283906001600160a01b038216906340c10f19906106319086908690600401610aea565b6107468363a9059cbb60e01b848460405160240161070f929190610aea565b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b03199093169290921790915261076c565b505050565b610228846323b872dd60e01b85858560405160240161070f93929190610ac6565b60006060836001600160a01b0316836040516107889190610a96565b6000604051808303816000865af19150503d80600081146107c5576040519150601f19603f3d011682016040523d82523d6000602084013e6107ca565b606091505b5091509150816107ec5760405162461bcd60e51b81526004016102f790610cb4565b805115610228578080602001905181019061080791906108cf565b6102285760405162461bcd60e51b81526004016102f790610ba7565b60008083601f840112610834578182fd5b50813567ffffffffffffffff81111561084b578182fd5b60208301915083602082850101111561086357600080fd5b9250929050565b60006020828403121561087b578081fd5b813561088681610d88565b9392505050565b6000806000606084860312156108a1578182fd5b83516108ac81610d88565b60208501519093506108bd81610d88565b80925050604084015190509250925092565b6000602082840312156108e0578081fd5b81518015158114610886578182fd5b600060208284031215610900578081fd5b5035919050565b60008060408385031215610919578182fd5b82359150602083013561092b81610d88565b809150509250929050565b6000806000806060858703121561094b578081fd5b84359350602085013561095d81610d88565b9250604085013567ffffffffffffffff811115610978578182fd5b61098487828801610823565b95989497509550505050565b6000806000604084860312156109a4578283fd5b83359250602084013567ffffffffffffffff8111156109c1578283fd5b6109cd86828701610823565b9497909650939450505050565b6000602082840312156109eb578081fd5b813567ffffffffffffffff80821115610a02578283fd5b818401915084601f830112610a15578283fd5b813581811115610a23578384fd5b604051601f8201601f191681016020018381118282101715610a43578586fd5b604052818152838201602001871015610a5a578485fd5b610a6b826020830160208701610d50565b9695505050505050565b60008060408385031215610a87578182fd5b50508035926020909101359150565b60008251610aa8818460208701610d5c565b9190910192915050565b6001600160a01b0391909116815260200190565b6001600160a01b039384168152919092166020820152604081019190915260600190565b6001600160a01b03929092168252602082015260400190565b901515815260200190565b90815260200190565b6000602082528251806020840152610b36816040850160208701610d5c565b601f01601f19169190910160400192915050565b6020808252603b908201527f70726f766964656420636f6e74726163742063616e6e6f74206265206c6f636b60408201527f4d696e74556e6c6f636b61626c6520616e64206275726e61626c650000000000606082015260800190565b6020808252818101527f45524332303a206f7065726174696f6e20646964206e6f742073756363656564604082015260600190565b6020808252601e908201527f73656e646572206d7573742062652062726964676520636f6e74726163740000604082015260600190565b6020808252603b908201527f70726f766964656420636f6e74726163742063616e6e6f74206265206275726e60408201527f61626c6520616e64206c6f636b4d696e74556e6c6f636b61626c650000000000606082015260800190565b60208082526024908201527f70726f766964656420636f6e7472616374206973206e6f742077686974656c696040820152631cdd195960e21b606082015260800190565b602080825260129082015271115490cc8c0e8818d85b1b0819985a5b195960721b604082015260600190565b60208082526028908201527f70726f766964656420746f6b656e41646472657373206973206e6f74207768696040820152671d195b1a5cdd195960c21b606082015260800190565b60008085851115610d37578182fd5b83861115610d43578182fd5b5050820193919092039150565b82818337506000910152565b60005b83811015610d77578181015183820152602001610d5f565b838111156102285750506000910152565b6001600160a01b03811681146101e157600080fdfea264697066735822122077178213f349092dd75670f09197530292c45c46ecbaf516353bcbb2521eb6d564736f6c634300060c0033",
}

// ERC20HandlerABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC20HandlerMetaData.ABI instead.
var ERC20HandlerABI = ERC20HandlerMetaData.ABI

// ERC20HandlerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ERC20HandlerMetaData.Bin instead.
var ERC20HandlerBin = ERC20HandlerMetaData.Bin

// DeployERC20Handler deploys a new Ethereum contract, binding an instance of ERC20Handler to it.
func DeployERC20Handler(auth *bind.TransactOpts, backend bind.ContractBackend, bridgeAddress common.Address) (common.Address, *types.Transaction, *ERC20Handler, error) {
	parsed, err := ERC20HandlerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ERC20HandlerBin), backend, bridgeAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20Handler{ERC20HandlerCaller: ERC20HandlerCaller{contract: contract}, ERC20HandlerTransactor: ERC20HandlerTransactor{contract: contract}, ERC20HandlerFilterer: ERC20HandlerFilterer{contract: contract}}, nil
}

// ERC20Handler is an auto generated Go binding around an Ethereum contract.
type ERC20Handler struct {
	ERC20HandlerCaller     // Read-only binding to the contract
	ERC20HandlerTransactor // Write-only binding to the contract
	ERC20HandlerFilterer   // Log filterer for contract events
}

// ERC20HandlerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20HandlerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20HandlerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20HandlerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20HandlerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20HandlerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20HandlerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20HandlerSession struct {
	Contract     *ERC20Handler     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20HandlerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20HandlerCallerSession struct {
	Contract *ERC20HandlerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ERC20HandlerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20HandlerTransactorSession struct {
	Contract     *ERC20HandlerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ERC20HandlerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20HandlerRaw struct {
	Contract *ERC20Handler // Generic contract binding to access the raw methods on
}

// ERC20HandlerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20HandlerCallerRaw struct {
	Contract *ERC20HandlerCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20HandlerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20HandlerTransactorRaw struct {
	Contract *ERC20HandlerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20Handler creates a new instance of ERC20Handler, bound to a specific deployed contract.
func NewERC20Handler(address common.Address, backend bind.ContractBackend) (*ERC20Handler, error) {
	contract, err := bindERC20Handler(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20Handler{ERC20HandlerCaller: ERC20HandlerCaller{contract: contract}, ERC20HandlerTransactor: ERC20HandlerTransactor{contract: contract}, ERC20HandlerFilterer: ERC20HandlerFilterer{contract: contract}}, nil
}

// NewERC20HandlerCaller creates a new read-only instance of ERC20Handler, bound to a specific deployed contract.
func NewERC20HandlerCaller(address common.Address, caller bind.ContractCaller) (*ERC20HandlerCaller, error) {
	contract, err := bindERC20Handler(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20HandlerCaller{contract: contract}, nil
}

// NewERC20HandlerTransactor creates a new write-only instance of ERC20Handler, bound to a specific deployed contract.
func NewERC20HandlerTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20HandlerTransactor, error) {
	contract, err := bindERC20Handler(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20HandlerTransactor{contract: contract}, nil
}

// NewERC20HandlerFilterer creates a new log filterer instance of ERC20Handler, bound to a specific deployed contract.
func NewERC20HandlerFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20HandlerFilterer, error) {
	contract, err := bindERC20Handler(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20HandlerFilterer{contract: contract}, nil
}

// bindERC20Handler binds a generic wrapper to an already deployed contract.
func bindERC20Handler(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20HandlerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Handler *ERC20HandlerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20Handler.Contract.ERC20HandlerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Handler *ERC20HandlerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Handler.Contract.ERC20HandlerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Handler *ERC20HandlerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Handler.Contract.ERC20HandlerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Handler *ERC20HandlerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20Handler.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Handler *ERC20HandlerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Handler.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Handler *ERC20HandlerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Handler.Contract.contract.Transact(opts, method, params...)
}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_ERC20Handler *ERC20HandlerCaller) BridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20Handler.contract.Call(opts, &out, "_bridgeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_ERC20Handler *ERC20HandlerSession) BridgeAddress() (common.Address, error) {
	return _ERC20Handler.Contract.BridgeAddress(&_ERC20Handler.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_ERC20Handler *ERC20HandlerCallerSession) BridgeAddress() (common.Address, error) {
	return _ERC20Handler.Contract.BridgeAddress(&_ERC20Handler.CallOpts)
}

// BurnList is a free data retrieval call binding the contract method 0x6a70d081.
//
// Solidity: function _burnList(address ) view returns(bool)
func (_ERC20Handler *ERC20HandlerCaller) BurnList(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _ERC20Handler.contract.Call(opts, &out, "_burnList", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BurnList is a free data retrieval call binding the contract method 0x6a70d081.
//
// Solidity: function _burnList(address ) view returns(bool)
func (_ERC20Handler *ERC20HandlerSession) BurnList(arg0 common.Address) (bool, error) {
	return _ERC20Handler.Contract.BurnList(&_ERC20Handler.CallOpts, arg0)
}

// BurnList is a free data retrieval call binding the contract method 0x6a70d081.
//
// Solidity: function _burnList(address ) view returns(bool)
func (_ERC20Handler *ERC20HandlerCallerSession) BurnList(arg0 common.Address) (bool, error) {
	return _ERC20Handler.Contract.BurnList(&_ERC20Handler.CallOpts, arg0)
}

// ContractWhitelist is a free data retrieval call binding the contract method 0x7f79bea8.
//
// Solidity: function _contractWhitelist(address ) view returns(bool)
func (_ERC20Handler *ERC20HandlerCaller) ContractWhitelist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _ERC20Handler.contract.Call(opts, &out, "_contractWhitelist", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ContractWhitelist is a free data retrieval call binding the contract method 0x7f79bea8.
//
// Solidity: function _contractWhitelist(address ) view returns(bool)
func (_ERC20Handler *ERC20HandlerSession) ContractWhitelist(arg0 common.Address) (bool, error) {
	return _ERC20Handler.Contract.ContractWhitelist(&_ERC20Handler.CallOpts, arg0)
}

// ContractWhitelist is a free data retrieval call binding the contract method 0x7f79bea8.
//
// Solidity: function _contractWhitelist(address ) view returns(bool)
func (_ERC20Handler *ERC20HandlerCallerSession) ContractWhitelist(arg0 common.Address) (bool, error) {
	return _ERC20Handler.Contract.ContractWhitelist(&_ERC20Handler.CallOpts, arg0)
}

// LockMintUnlockList is a free data retrieval call binding the contract method 0xb1f06463.
//
// Solidity: function _lockMintUnlockList(address ) view returns(bool)
func (_ERC20Handler *ERC20HandlerCaller) LockMintUnlockList(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _ERC20Handler.contract.Call(opts, &out, "_lockMintUnlockList", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// LockMintUnlockList is a free data retrieval call binding the contract method 0xb1f06463.
//
// Solidity: function _lockMintUnlockList(address ) view returns(bool)
func (_ERC20Handler *ERC20HandlerSession) LockMintUnlockList(arg0 common.Address) (bool, error) {
	return _ERC20Handler.Contract.LockMintUnlockList(&_ERC20Handler.CallOpts, arg0)
}

// LockMintUnlockList is a free data retrieval call binding the contract method 0xb1f06463.
//
// Solidity: function _lockMintUnlockList(address ) view returns(bool)
func (_ERC20Handler *ERC20HandlerCallerSession) LockMintUnlockList(arg0 common.Address) (bool, error) {
	return _ERC20Handler.Contract.LockMintUnlockList(&_ERC20Handler.CallOpts, arg0)
}

// ResourceIDToTokenContractAddress is a free data retrieval call binding the contract method 0x0a6d55d8.
//
// Solidity: function _resourceIDToTokenContractAddress(bytes32 ) view returns(address)
func (_ERC20Handler *ERC20HandlerCaller) ResourceIDToTokenContractAddress(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _ERC20Handler.contract.Call(opts, &out, "_resourceIDToTokenContractAddress", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ResourceIDToTokenContractAddress is a free data retrieval call binding the contract method 0x0a6d55d8.
//
// Solidity: function _resourceIDToTokenContractAddress(bytes32 ) view returns(address)
func (_ERC20Handler *ERC20HandlerSession) ResourceIDToTokenContractAddress(arg0 [32]byte) (common.Address, error) {
	return _ERC20Handler.Contract.ResourceIDToTokenContractAddress(&_ERC20Handler.CallOpts, arg0)
}

// ResourceIDToTokenContractAddress is a free data retrieval call binding the contract method 0x0a6d55d8.
//
// Solidity: function _resourceIDToTokenContractAddress(bytes32 ) view returns(address)
func (_ERC20Handler *ERC20HandlerCallerSession) ResourceIDToTokenContractAddress(arg0 [32]byte) (common.Address, error) {
	return _ERC20Handler.Contract.ResourceIDToTokenContractAddress(&_ERC20Handler.CallOpts, arg0)
}

// TokenContractAddressToResourceID is a free data retrieval call binding the contract method 0xc8ba6c87.
//
// Solidity: function _tokenContractAddressToResourceID(address ) view returns(bytes32)
func (_ERC20Handler *ERC20HandlerCaller) TokenContractAddressToResourceID(opts *bind.CallOpts, arg0 common.Address) ([32]byte, error) {
	var out []interface{}
	err := _ERC20Handler.contract.Call(opts, &out, "_tokenContractAddressToResourceID", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TokenContractAddressToResourceID is a free data retrieval call binding the contract method 0xc8ba6c87.
//
// Solidity: function _tokenContractAddressToResourceID(address ) view returns(bytes32)
func (_ERC20Handler *ERC20HandlerSession) TokenContractAddressToResourceID(arg0 common.Address) ([32]byte, error) {
	return _ERC20Handler.Contract.TokenContractAddressToResourceID(&_ERC20Handler.CallOpts, arg0)
}

// TokenContractAddressToResourceID is a free data retrieval call binding the contract method 0xc8ba6c87.
//
// Solidity: function _tokenContractAddressToResourceID(address ) view returns(bytes32)
func (_ERC20Handler *ERC20HandlerCallerSession) TokenContractAddressToResourceID(arg0 common.Address) ([32]byte, error) {
	return _ERC20Handler.Contract.TokenContractAddressToResourceID(&_ERC20Handler.CallOpts, arg0)
}

// Deposit is a paid mutator transaction binding the contract method 0xb07e54bb.
//
// Solidity: function deposit(bytes32 resourceID, address depositer, bytes data) returns(bytes)
func (_ERC20Handler *ERC20HandlerTransactor) Deposit(opts *bind.TransactOpts, resourceID [32]byte, depositer common.Address, data []byte) (*types.Transaction, error) {
	return _ERC20Handler.contract.Transact(opts, "deposit", resourceID, depositer, data)
}

// Deposit is a paid mutator transaction binding the contract method 0xb07e54bb.
//
// Solidity: function deposit(bytes32 resourceID, address depositer, bytes data) returns(bytes)
func (_ERC20Handler *ERC20HandlerSession) Deposit(resourceID [32]byte, depositer common.Address, data []byte) (*types.Transaction, error) {
	return _ERC20Handler.Contract.Deposit(&_ERC20Handler.TransactOpts, resourceID, depositer, data)
}

// Deposit is a paid mutator transaction binding the contract method 0xb07e54bb.
//
// Solidity: function deposit(bytes32 resourceID, address depositer, bytes data) returns(bytes)
func (_ERC20Handler *ERC20HandlerTransactorSession) Deposit(resourceID [32]byte, depositer common.Address, data []byte) (*types.Transaction, error) {
	return _ERC20Handler.Contract.Deposit(&_ERC20Handler.TransactOpts, resourceID, depositer, data)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0xe248cff2.
//
// Solidity: function executeProposal(bytes32 resourceID, bytes data) returns()
func (_ERC20Handler *ERC20HandlerTransactor) ExecuteProposal(opts *bind.TransactOpts, resourceID [32]byte, data []byte) (*types.Transaction, error) {
	return _ERC20Handler.contract.Transact(opts, "executeProposal", resourceID, data)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0xe248cff2.
//
// Solidity: function executeProposal(bytes32 resourceID, bytes data) returns()
func (_ERC20Handler *ERC20HandlerSession) ExecuteProposal(resourceID [32]byte, data []byte) (*types.Transaction, error) {
	return _ERC20Handler.Contract.ExecuteProposal(&_ERC20Handler.TransactOpts, resourceID, data)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0xe248cff2.
//
// Solidity: function executeProposal(bytes32 resourceID, bytes data) returns()
func (_ERC20Handler *ERC20HandlerTransactorSession) ExecuteProposal(resourceID [32]byte, data []byte) (*types.Transaction, error) {
	return _ERC20Handler.Contract.ExecuteProposal(&_ERC20Handler.TransactOpts, resourceID, data)
}

// SetBurnable is a paid mutator transaction binding the contract method 0x07b7ed99.
//
// Solidity: function setBurnable(address contractAddress) returns()
func (_ERC20Handler *ERC20HandlerTransactor) SetBurnable(opts *bind.TransactOpts, contractAddress common.Address) (*types.Transaction, error) {
	return _ERC20Handler.contract.Transact(opts, "setBurnable", contractAddress)
}

// SetBurnable is a paid mutator transaction binding the contract method 0x07b7ed99.
//
// Solidity: function setBurnable(address contractAddress) returns()
func (_ERC20Handler *ERC20HandlerSession) SetBurnable(contractAddress common.Address) (*types.Transaction, error) {
	return _ERC20Handler.Contract.SetBurnable(&_ERC20Handler.TransactOpts, contractAddress)
}

// SetBurnable is a paid mutator transaction binding the contract method 0x07b7ed99.
//
// Solidity: function setBurnable(address contractAddress) returns()
func (_ERC20Handler *ERC20HandlerTransactorSession) SetBurnable(contractAddress common.Address) (*types.Transaction, error) {
	return _ERC20Handler.Contract.SetBurnable(&_ERC20Handler.TransactOpts, contractAddress)
}

// SetLockMintUnlockable is a paid mutator transaction binding the contract method 0x00d8e7ee.
//
// Solidity: function setLockMintUnlockable(address contractAddress) returns()
func (_ERC20Handler *ERC20HandlerTransactor) SetLockMintUnlockable(opts *bind.TransactOpts, contractAddress common.Address) (*types.Transaction, error) {
	return _ERC20Handler.contract.Transact(opts, "setLockMintUnlockable", contractAddress)
}

// SetLockMintUnlockable is a paid mutator transaction binding the contract method 0x00d8e7ee.
//
// Solidity: function setLockMintUnlockable(address contractAddress) returns()
func (_ERC20Handler *ERC20HandlerSession) SetLockMintUnlockable(contractAddress common.Address) (*types.Transaction, error) {
	return _ERC20Handler.Contract.SetLockMintUnlockable(&_ERC20Handler.TransactOpts, contractAddress)
}

// SetLockMintUnlockable is a paid mutator transaction binding the contract method 0x00d8e7ee.
//
// Solidity: function setLockMintUnlockable(address contractAddress) returns()
func (_ERC20Handler *ERC20HandlerTransactorSession) SetLockMintUnlockable(contractAddress common.Address) (*types.Transaction, error) {
	return _ERC20Handler.Contract.SetLockMintUnlockable(&_ERC20Handler.TransactOpts, contractAddress)
}

// SetResource is a paid mutator transaction binding the contract method 0xb8fa3736.
//
// Solidity: function setResource(bytes32 resourceID, address contractAddress) returns()
func (_ERC20Handler *ERC20HandlerTransactor) SetResource(opts *bind.TransactOpts, resourceID [32]byte, contractAddress common.Address) (*types.Transaction, error) {
	return _ERC20Handler.contract.Transact(opts, "setResource", resourceID, contractAddress)
}

// SetResource is a paid mutator transaction binding the contract method 0xb8fa3736.
//
// Solidity: function setResource(bytes32 resourceID, address contractAddress) returns()
func (_ERC20Handler *ERC20HandlerSession) SetResource(resourceID [32]byte, contractAddress common.Address) (*types.Transaction, error) {
	return _ERC20Handler.Contract.SetResource(&_ERC20Handler.TransactOpts, resourceID, contractAddress)
}

// SetResource is a paid mutator transaction binding the contract method 0xb8fa3736.
//
// Solidity: function setResource(bytes32 resourceID, address contractAddress) returns()
func (_ERC20Handler *ERC20HandlerTransactorSession) SetResource(resourceID [32]byte, contractAddress common.Address) (*types.Transaction, error) {
	return _ERC20Handler.Contract.SetResource(&_ERC20Handler.TransactOpts, resourceID, contractAddress)
}

// Withdraw is a paid mutator transaction binding the contract method 0x0968f264.
//
// Solidity: function withdraw(bytes data) returns()
func (_ERC20Handler *ERC20HandlerTransactor) Withdraw(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _ERC20Handler.contract.Transact(opts, "withdraw", data)
}

// Withdraw is a paid mutator transaction binding the contract method 0x0968f264.
//
// Solidity: function withdraw(bytes data) returns()
func (_ERC20Handler *ERC20HandlerSession) Withdraw(data []byte) (*types.Transaction, error) {
	return _ERC20Handler.Contract.Withdraw(&_ERC20Handler.TransactOpts, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0x0968f264.
//
// Solidity: function withdraw(bytes data) returns()
func (_ERC20Handler *ERC20HandlerTransactorSession) Withdraw(data []byte) (*types.Transaction, error) {
	return _ERC20Handler.Contract.Withdraw(&_ERC20Handler.TransactOpts, data)
}
