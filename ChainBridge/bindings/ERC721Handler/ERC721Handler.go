// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ERC721Handler

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

// ERC721HandlerMetaData contains all meta data concerning the ERC721Handler contract.
var ERC721HandlerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridgeAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"_bridgeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_burnList\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_contractWhitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_lockMintUnlockList\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"_resourceIDToTokenContractAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_tokenContractAddressToResourceID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setBurnable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setLockMintUnlockable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setResource\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"depositer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"metaData\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"executeProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561001057600080fd5b5060405161114238038061114283398101604081905261002f91610044565b60601b6001600160601b031916608052610072565b600060208284031215610055578081fd5b81516001600160a01b038116811461006b578182fd5b9392505050565b60805160601c6110ae6100946000398061024c528061067552506110ae6000f3fe608060405234801561001057600080fd5b50600436106100b35760003560e01c80637f79bea8116100715780637f79bea814610144578063b07e54bb14610157578063b1f0646314610177578063b8fa37361461018a578063c8ba6c871461019d578063e248cff2146101bd576100b3565b8062d8e7ee146100b857806307b7ed99146100cd5780630968f264146100e05780630a6d55d8146100f3578063318c136e1461011c5780636a70d08114610124575b600080fd5b6100cb6100c6366004610adf565b6101d0565b005b6100cb6100db366004610adf565b6101e4565b6100cb6100ee366004610c48565b6101f5565b610106610101366004610b5d565b61022f565b6040516101139190610dac565b60405180910390f35b61010661024a565b610137610132366004610adf565b61026e565b6040516101139190610e0b565b610137610152366004610adf565b610283565b61016a610165366004610ba4565b610298565b6040516101139190610e34565b610137610185366004610adf565b6103ec565b6100cb610198366004610b75565b610401565b6101b06101ab366004610adf565b610417565b6040516101139190610e16565b6100cb6101cb366004610bfe565b610429565b6101d861066a565b6101e1816106b4565b50565b6101ec61066a565b6101e181610749565b6101fd61066a565b6000806000838060200190518101906102169190610afb565b91945092509050610229833084846107de565b50505050565b6000602081905290815260409020546001600160a01b031681565b7f000000000000000000000000000000000000000000000000000000000000000081565b60036020526000908152604090205460ff1681565b60026020526000908152604090205460ff1681565b60606102a261066a565b60006102b083850185610b5d565b600087815260208181526040808320546001600160a01b03168084526002909252909120549192509060ff166103015760405162461bcd60e51b81526004016102f890610f7c565b60405180910390fd5b61031b6001600160a01b038216635b5e139f60e01b610849565b156103a65760405163c87b56dd60e01b815281906001600160a01b0382169063c87b56dd9061034e908690600401610e16565b60006040518083038186803b15801561036657600080fd5b505afa15801561037a573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526103a29190810190610cc1565b9350505b6001600160a01b03811660009081526003602052604090205460ff16156103d6576103d1818361086c565b6103e2565b6103e2818730856107de565b5050949350505050565b60046020526000908152604090205460ff1681565b61040961066a565b61041382826108d1565b5050565b60016020526000908152604090205481565b61043161066a565b600080606081808261044587890189610d43565b90965094506040808601935061045e908490898b61100f565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509296506104a49250899150859050818b61100f565b8101906104b19190610b5d565b91506104c760208484018101908501898b61100f565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201829052506020898101518f83528282526040808420546001600160a01b0316808552600290935290922054959650909490935060ff16915061054990505760405162461bcd60e51b81526004016102f890610f7c565b6001600160a01b03811660009081526003602052604090205460ff161561057e57610579818360601c8a8661091e565b61065d565b6001600160a01b03811660009081526004602052604090205460ff161561064e576040516370a0823160e01b815281906000906001600160a01b038316906370a08231906105d0903090600401610dac565b60206040518083038186803b1580156105e857600080fd5b505afa1580156105fc573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106209190610d2b565b1161063957610634828460601c8b8761091e565b610648565b61064882308560601c8c6107de565b5061065d565b61065d81308460601c8b6107de565b5050505050505050505050565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146106b25760405162461bcd60e51b81526004016102f890610ea4565b565b6001600160a01b03811660009081526002602052604090205460ff166106ec5760405162461bcd60e51b81526004016102f890610f38565b6001600160a01b03811660009081526003602052604090205460ff16156107255760405162461bcd60e51b81526004016102f890610edb565b6001600160a01b03166000908152600460205260409020805460ff19166001179055565b6001600160a01b03811660009081526002602052604090205460ff166107815760405162461bcd60e51b81526004016102f890610f38565b6001600160a01b03811660009081526004602052604090205460ff16156107ba5760405162461bcd60e51b81526004016102f890610e47565b6001600160a01b03166000908152600360205260409020805460ff19166001179055565b6040516323b872dd60e01b815284906001600160a01b038216906323b872dd9061081090879087908790600401610dc0565b600060405180830381600087803b15801561082a57600080fd5b505af115801561083e573d6000803e3d6000fd5b505050505050505050565b600061085483610950565b801561086557506108658383610984565b9392505050565b604051630852cd8d60e31b815282906001600160a01b038216906342966c689061089a908590600401610e16565b600060405180830381600087803b1580156108b457600080fd5b505af11580156108c8573d6000803e3d6000fd5b50505050505050565b60008281526020818152604080832080546001600160a01b039095166001600160a01b0319909516851790559282526001808252838320949094556002905220805460ff19169091179055565b6040516334ff261960e21b815284906001600160a01b0382169063d3fc98649061081090879087908790600401610de4565b6000610963826301ffc9a760e01b610984565b801561097e575061097c826001600160e01b0319610984565b155b92915050565b600080600061099385856109aa565b915091508180156109a15750805b95945050505050565b60008060606301ffc9a760e01b846040516024016109c89190610e1f565b604051602081830303815290604052906001600160e01b0319166020820180516001600160e01b038381831617835250505050905060006060866001600160a01b031661753084604051610a1c9190610d90565b6000604051808303818686fa925050503d8060008114610a58576040519150601f19603f3d011682016040523d82523d6000602084013e610a5d565b606091505b5091509150602081511015610a7b5760008094509450505050610a98565b8181806020019051810190610a909190610b3d565b945094505050505b9250929050565b60008083601f840112610ab0578182fd5b50813567ffffffffffffffff811115610ac7578182fd5b602083019150836020828501011115610a9857600080fd5b600060208284031215610af0578081fd5b813561086581611063565b600080600060608486031215610b0f578182fd5b8351610b1a81611063565b6020850151909350610b2b81611063565b80925050604084015190509250925092565b600060208284031215610b4e578081fd5b81518015158114610865578182fd5b600060208284031215610b6e578081fd5b5035919050565b60008060408385031215610b87578182fd5b823591506020830135610b9981611063565b809150509250929050565b60008060008060608587031215610bb9578081fd5b843593506020850135610bcb81611063565b9250604085013567ffffffffffffffff811115610be6578182fd5b610bf287828801610a9f565b95989497509550505050565b600080600060408486031215610c12578283fd5b83359250602084013567ffffffffffffffff811115610c2f578283fd5b610c3b86828701610a9f565b9497909650939450505050565b600060208284031215610c59578081fd5b813567ffffffffffffffff811115610c6f578182fd5b8201601f81018413610c7f578182fd5b8035610c92610c8d82610feb565b610fc4565b818152856020838501011115610ca6578384fd5b81602084016020830137908101602001929092525092915050565b600060208284031215610cd2578081fd5b815167ffffffffffffffff811115610ce8578182fd5b8201601f81018413610cf8578182fd5b8051610d06610c8d82610feb565b818152856020838501011115610d1a578384fd5b6109a1826020830160208601611037565b600060208284031215610d3c578081fd5b5051919050565b60008060408385031215610d55578182fd5b50508035926020909101359150565b60008151808452610d7c816020860160208601611037565b601f01601f19169290920160200192915050565b60008251610da2818460208701611037565b9190910192915050565b6001600160a01b0391909116815260200190565b6001600160a01b039384168152919092166020820152604081019190915260600190565b600060018060a01b0385168252836020830152606060408301526109a16060830184610d64565b901515815260200190565b90815260200190565b6001600160e01b031991909116815260200190565b6000602082526108656020830184610d64565b6020808252603b908201527f70726f766964656420636f6e74726163742063616e6e6f74206265206c6f636b60408201527f4d696e74556e6c6f636b61626c6520616e64206275726e61626c650000000000606082015260800190565b6020808252601e908201527f73656e646572206d7573742062652062726964676520636f6e74726163740000604082015260600190565b6020808252603b908201527f70726f766964656420636f6e74726163742063616e6e6f74206265206275726e60408201527f61626c6520616e64206c6f636b4d696e74556e6c6f636b61626c650000000000606082015260800190565b60208082526024908201527f70726f766964656420636f6e7472616374206973206e6f742077686974656c696040820152631cdd195960e21b606082015260800190565b60208082526028908201527f70726f766964656420746f6b656e41646472657373206973206e6f74207768696040820152671d195b1a5cdd195960c21b606082015260800190565b60405181810167ffffffffffffffff81118282101715610fe357600080fd5b604052919050565b600067ffffffffffffffff821115611001578081fd5b50601f01601f191660200190565b6000808585111561101e578182fd5b8386111561102a578182fd5b5050820193919092039150565b60005b8381101561105257818101518382015260200161103a565b838111156102295750506000910152565b6001600160a01b03811681146101e157600080fdfea26469706673582212207627722f91eeb76f9f949bab9f2e720571b79b8d19f03f4dd4da14857313f44464736f6c634300060c0033",
}

// ERC721HandlerABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC721HandlerMetaData.ABI instead.
var ERC721HandlerABI = ERC721HandlerMetaData.ABI

// ERC721HandlerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ERC721HandlerMetaData.Bin instead.
var ERC721HandlerBin = ERC721HandlerMetaData.Bin

// DeployERC721Handler deploys a new Ethereum contract, binding an instance of ERC721Handler to it.
func DeployERC721Handler(auth *bind.TransactOpts, backend bind.ContractBackend, bridgeAddress common.Address) (common.Address, *types.Transaction, *ERC721Handler, error) {
	parsed, err := ERC721HandlerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ERC721HandlerBin), backend, bridgeAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC721Handler{ERC721HandlerCaller: ERC721HandlerCaller{contract: contract}, ERC721HandlerTransactor: ERC721HandlerTransactor{contract: contract}, ERC721HandlerFilterer: ERC721HandlerFilterer{contract: contract}}, nil
}

// ERC721Handler is an auto generated Go binding around an Ethereum contract.
type ERC721Handler struct {
	ERC721HandlerCaller     // Read-only binding to the contract
	ERC721HandlerTransactor // Write-only binding to the contract
	ERC721HandlerFilterer   // Log filterer for contract events
}

// ERC721HandlerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC721HandlerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721HandlerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC721HandlerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721HandlerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC721HandlerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721HandlerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC721HandlerSession struct {
	Contract     *ERC721Handler    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC721HandlerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC721HandlerCallerSession struct {
	Contract *ERC721HandlerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ERC721HandlerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC721HandlerTransactorSession struct {
	Contract     *ERC721HandlerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ERC721HandlerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC721HandlerRaw struct {
	Contract *ERC721Handler // Generic contract binding to access the raw methods on
}

// ERC721HandlerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC721HandlerCallerRaw struct {
	Contract *ERC721HandlerCaller // Generic read-only contract binding to access the raw methods on
}

// ERC721HandlerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC721HandlerTransactorRaw struct {
	Contract *ERC721HandlerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC721Handler creates a new instance of ERC721Handler, bound to a specific deployed contract.
func NewERC721Handler(address common.Address, backend bind.ContractBackend) (*ERC721Handler, error) {
	contract, err := bindERC721Handler(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC721Handler{ERC721HandlerCaller: ERC721HandlerCaller{contract: contract}, ERC721HandlerTransactor: ERC721HandlerTransactor{contract: contract}, ERC721HandlerFilterer: ERC721HandlerFilterer{contract: contract}}, nil
}

// NewERC721HandlerCaller creates a new read-only instance of ERC721Handler, bound to a specific deployed contract.
func NewERC721HandlerCaller(address common.Address, caller bind.ContractCaller) (*ERC721HandlerCaller, error) {
	contract, err := bindERC721Handler(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721HandlerCaller{contract: contract}, nil
}

// NewERC721HandlerTransactor creates a new write-only instance of ERC721Handler, bound to a specific deployed contract.
func NewERC721HandlerTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC721HandlerTransactor, error) {
	contract, err := bindERC721Handler(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721HandlerTransactor{contract: contract}, nil
}

// NewERC721HandlerFilterer creates a new log filterer instance of ERC721Handler, bound to a specific deployed contract.
func NewERC721HandlerFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC721HandlerFilterer, error) {
	contract, err := bindERC721Handler(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC721HandlerFilterer{contract: contract}, nil
}

// bindERC721Handler binds a generic wrapper to an already deployed contract.
func bindERC721Handler(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC721HandlerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721Handler *ERC721HandlerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC721Handler.Contract.ERC721HandlerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721Handler *ERC721HandlerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721Handler.Contract.ERC721HandlerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721Handler *ERC721HandlerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721Handler.Contract.ERC721HandlerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721Handler *ERC721HandlerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC721Handler.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721Handler *ERC721HandlerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721Handler.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721Handler *ERC721HandlerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721Handler.Contract.contract.Transact(opts, method, params...)
}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_ERC721Handler *ERC721HandlerCaller) BridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC721Handler.contract.Call(opts, &out, "_bridgeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_ERC721Handler *ERC721HandlerSession) BridgeAddress() (common.Address, error) {
	return _ERC721Handler.Contract.BridgeAddress(&_ERC721Handler.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_ERC721Handler *ERC721HandlerCallerSession) BridgeAddress() (common.Address, error) {
	return _ERC721Handler.Contract.BridgeAddress(&_ERC721Handler.CallOpts)
}

// BurnList is a free data retrieval call binding the contract method 0x6a70d081.
//
// Solidity: function _burnList(address ) view returns(bool)
func (_ERC721Handler *ERC721HandlerCaller) BurnList(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _ERC721Handler.contract.Call(opts, &out, "_burnList", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BurnList is a free data retrieval call binding the contract method 0x6a70d081.
//
// Solidity: function _burnList(address ) view returns(bool)
func (_ERC721Handler *ERC721HandlerSession) BurnList(arg0 common.Address) (bool, error) {
	return _ERC721Handler.Contract.BurnList(&_ERC721Handler.CallOpts, arg0)
}

// BurnList is a free data retrieval call binding the contract method 0x6a70d081.
//
// Solidity: function _burnList(address ) view returns(bool)
func (_ERC721Handler *ERC721HandlerCallerSession) BurnList(arg0 common.Address) (bool, error) {
	return _ERC721Handler.Contract.BurnList(&_ERC721Handler.CallOpts, arg0)
}

// ContractWhitelist is a free data retrieval call binding the contract method 0x7f79bea8.
//
// Solidity: function _contractWhitelist(address ) view returns(bool)
func (_ERC721Handler *ERC721HandlerCaller) ContractWhitelist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _ERC721Handler.contract.Call(opts, &out, "_contractWhitelist", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ContractWhitelist is a free data retrieval call binding the contract method 0x7f79bea8.
//
// Solidity: function _contractWhitelist(address ) view returns(bool)
func (_ERC721Handler *ERC721HandlerSession) ContractWhitelist(arg0 common.Address) (bool, error) {
	return _ERC721Handler.Contract.ContractWhitelist(&_ERC721Handler.CallOpts, arg0)
}

// ContractWhitelist is a free data retrieval call binding the contract method 0x7f79bea8.
//
// Solidity: function _contractWhitelist(address ) view returns(bool)
func (_ERC721Handler *ERC721HandlerCallerSession) ContractWhitelist(arg0 common.Address) (bool, error) {
	return _ERC721Handler.Contract.ContractWhitelist(&_ERC721Handler.CallOpts, arg0)
}

// LockMintUnlockList is a free data retrieval call binding the contract method 0xb1f06463.
//
// Solidity: function _lockMintUnlockList(address ) view returns(bool)
func (_ERC721Handler *ERC721HandlerCaller) LockMintUnlockList(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _ERC721Handler.contract.Call(opts, &out, "_lockMintUnlockList", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// LockMintUnlockList is a free data retrieval call binding the contract method 0xb1f06463.
//
// Solidity: function _lockMintUnlockList(address ) view returns(bool)
func (_ERC721Handler *ERC721HandlerSession) LockMintUnlockList(arg0 common.Address) (bool, error) {
	return _ERC721Handler.Contract.LockMintUnlockList(&_ERC721Handler.CallOpts, arg0)
}

// LockMintUnlockList is a free data retrieval call binding the contract method 0xb1f06463.
//
// Solidity: function _lockMintUnlockList(address ) view returns(bool)
func (_ERC721Handler *ERC721HandlerCallerSession) LockMintUnlockList(arg0 common.Address) (bool, error) {
	return _ERC721Handler.Contract.LockMintUnlockList(&_ERC721Handler.CallOpts, arg0)
}

// ResourceIDToTokenContractAddress is a free data retrieval call binding the contract method 0x0a6d55d8.
//
// Solidity: function _resourceIDToTokenContractAddress(bytes32 ) view returns(address)
func (_ERC721Handler *ERC721HandlerCaller) ResourceIDToTokenContractAddress(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _ERC721Handler.contract.Call(opts, &out, "_resourceIDToTokenContractAddress", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ResourceIDToTokenContractAddress is a free data retrieval call binding the contract method 0x0a6d55d8.
//
// Solidity: function _resourceIDToTokenContractAddress(bytes32 ) view returns(address)
func (_ERC721Handler *ERC721HandlerSession) ResourceIDToTokenContractAddress(arg0 [32]byte) (common.Address, error) {
	return _ERC721Handler.Contract.ResourceIDToTokenContractAddress(&_ERC721Handler.CallOpts, arg0)
}

// ResourceIDToTokenContractAddress is a free data retrieval call binding the contract method 0x0a6d55d8.
//
// Solidity: function _resourceIDToTokenContractAddress(bytes32 ) view returns(address)
func (_ERC721Handler *ERC721HandlerCallerSession) ResourceIDToTokenContractAddress(arg0 [32]byte) (common.Address, error) {
	return _ERC721Handler.Contract.ResourceIDToTokenContractAddress(&_ERC721Handler.CallOpts, arg0)
}

// TokenContractAddressToResourceID is a free data retrieval call binding the contract method 0xc8ba6c87.
//
// Solidity: function _tokenContractAddressToResourceID(address ) view returns(bytes32)
func (_ERC721Handler *ERC721HandlerCaller) TokenContractAddressToResourceID(opts *bind.CallOpts, arg0 common.Address) ([32]byte, error) {
	var out []interface{}
	err := _ERC721Handler.contract.Call(opts, &out, "_tokenContractAddressToResourceID", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TokenContractAddressToResourceID is a free data retrieval call binding the contract method 0xc8ba6c87.
//
// Solidity: function _tokenContractAddressToResourceID(address ) view returns(bytes32)
func (_ERC721Handler *ERC721HandlerSession) TokenContractAddressToResourceID(arg0 common.Address) ([32]byte, error) {
	return _ERC721Handler.Contract.TokenContractAddressToResourceID(&_ERC721Handler.CallOpts, arg0)
}

// TokenContractAddressToResourceID is a free data retrieval call binding the contract method 0xc8ba6c87.
//
// Solidity: function _tokenContractAddressToResourceID(address ) view returns(bytes32)
func (_ERC721Handler *ERC721HandlerCallerSession) TokenContractAddressToResourceID(arg0 common.Address) ([32]byte, error) {
	return _ERC721Handler.Contract.TokenContractAddressToResourceID(&_ERC721Handler.CallOpts, arg0)
}

// Deposit is a paid mutator transaction binding the contract method 0xb07e54bb.
//
// Solidity: function deposit(bytes32 resourceID, address depositer, bytes data) returns(bytes metaData)
func (_ERC721Handler *ERC721HandlerTransactor) Deposit(opts *bind.TransactOpts, resourceID [32]byte, depositer common.Address, data []byte) (*types.Transaction, error) {
	return _ERC721Handler.contract.Transact(opts, "deposit", resourceID, depositer, data)
}

// Deposit is a paid mutator transaction binding the contract method 0xb07e54bb.
//
// Solidity: function deposit(bytes32 resourceID, address depositer, bytes data) returns(bytes metaData)
func (_ERC721Handler *ERC721HandlerSession) Deposit(resourceID [32]byte, depositer common.Address, data []byte) (*types.Transaction, error) {
	return _ERC721Handler.Contract.Deposit(&_ERC721Handler.TransactOpts, resourceID, depositer, data)
}

// Deposit is a paid mutator transaction binding the contract method 0xb07e54bb.
//
// Solidity: function deposit(bytes32 resourceID, address depositer, bytes data) returns(bytes metaData)
func (_ERC721Handler *ERC721HandlerTransactorSession) Deposit(resourceID [32]byte, depositer common.Address, data []byte) (*types.Transaction, error) {
	return _ERC721Handler.Contract.Deposit(&_ERC721Handler.TransactOpts, resourceID, depositer, data)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0xe248cff2.
//
// Solidity: function executeProposal(bytes32 resourceID, bytes data) returns()
func (_ERC721Handler *ERC721HandlerTransactor) ExecuteProposal(opts *bind.TransactOpts, resourceID [32]byte, data []byte) (*types.Transaction, error) {
	return _ERC721Handler.contract.Transact(opts, "executeProposal", resourceID, data)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0xe248cff2.
//
// Solidity: function executeProposal(bytes32 resourceID, bytes data) returns()
func (_ERC721Handler *ERC721HandlerSession) ExecuteProposal(resourceID [32]byte, data []byte) (*types.Transaction, error) {
	return _ERC721Handler.Contract.ExecuteProposal(&_ERC721Handler.TransactOpts, resourceID, data)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0xe248cff2.
//
// Solidity: function executeProposal(bytes32 resourceID, bytes data) returns()
func (_ERC721Handler *ERC721HandlerTransactorSession) ExecuteProposal(resourceID [32]byte, data []byte) (*types.Transaction, error) {
	return _ERC721Handler.Contract.ExecuteProposal(&_ERC721Handler.TransactOpts, resourceID, data)
}

// SetBurnable is a paid mutator transaction binding the contract method 0x07b7ed99.
//
// Solidity: function setBurnable(address contractAddress) returns()
func (_ERC721Handler *ERC721HandlerTransactor) SetBurnable(opts *bind.TransactOpts, contractAddress common.Address) (*types.Transaction, error) {
	return _ERC721Handler.contract.Transact(opts, "setBurnable", contractAddress)
}

// SetBurnable is a paid mutator transaction binding the contract method 0x07b7ed99.
//
// Solidity: function setBurnable(address contractAddress) returns()
func (_ERC721Handler *ERC721HandlerSession) SetBurnable(contractAddress common.Address) (*types.Transaction, error) {
	return _ERC721Handler.Contract.SetBurnable(&_ERC721Handler.TransactOpts, contractAddress)
}

// SetBurnable is a paid mutator transaction binding the contract method 0x07b7ed99.
//
// Solidity: function setBurnable(address contractAddress) returns()
func (_ERC721Handler *ERC721HandlerTransactorSession) SetBurnable(contractAddress common.Address) (*types.Transaction, error) {
	return _ERC721Handler.Contract.SetBurnable(&_ERC721Handler.TransactOpts, contractAddress)
}

// SetLockMintUnlockable is a paid mutator transaction binding the contract method 0x00d8e7ee.
//
// Solidity: function setLockMintUnlockable(address contractAddress) returns()
func (_ERC721Handler *ERC721HandlerTransactor) SetLockMintUnlockable(opts *bind.TransactOpts, contractAddress common.Address) (*types.Transaction, error) {
	return _ERC721Handler.contract.Transact(opts, "setLockMintUnlockable", contractAddress)
}

// SetLockMintUnlockable is a paid mutator transaction binding the contract method 0x00d8e7ee.
//
// Solidity: function setLockMintUnlockable(address contractAddress) returns()
func (_ERC721Handler *ERC721HandlerSession) SetLockMintUnlockable(contractAddress common.Address) (*types.Transaction, error) {
	return _ERC721Handler.Contract.SetLockMintUnlockable(&_ERC721Handler.TransactOpts, contractAddress)
}

// SetLockMintUnlockable is a paid mutator transaction binding the contract method 0x00d8e7ee.
//
// Solidity: function setLockMintUnlockable(address contractAddress) returns()
func (_ERC721Handler *ERC721HandlerTransactorSession) SetLockMintUnlockable(contractAddress common.Address) (*types.Transaction, error) {
	return _ERC721Handler.Contract.SetLockMintUnlockable(&_ERC721Handler.TransactOpts, contractAddress)
}

// SetResource is a paid mutator transaction binding the contract method 0xb8fa3736.
//
// Solidity: function setResource(bytes32 resourceID, address contractAddress) returns()
func (_ERC721Handler *ERC721HandlerTransactor) SetResource(opts *bind.TransactOpts, resourceID [32]byte, contractAddress common.Address) (*types.Transaction, error) {
	return _ERC721Handler.contract.Transact(opts, "setResource", resourceID, contractAddress)
}

// SetResource is a paid mutator transaction binding the contract method 0xb8fa3736.
//
// Solidity: function setResource(bytes32 resourceID, address contractAddress) returns()
func (_ERC721Handler *ERC721HandlerSession) SetResource(resourceID [32]byte, contractAddress common.Address) (*types.Transaction, error) {
	return _ERC721Handler.Contract.SetResource(&_ERC721Handler.TransactOpts, resourceID, contractAddress)
}

// SetResource is a paid mutator transaction binding the contract method 0xb8fa3736.
//
// Solidity: function setResource(bytes32 resourceID, address contractAddress) returns()
func (_ERC721Handler *ERC721HandlerTransactorSession) SetResource(resourceID [32]byte, contractAddress common.Address) (*types.Transaction, error) {
	return _ERC721Handler.Contract.SetResource(&_ERC721Handler.TransactOpts, resourceID, contractAddress)
}

// Withdraw is a paid mutator transaction binding the contract method 0x0968f264.
//
// Solidity: function withdraw(bytes data) returns()
func (_ERC721Handler *ERC721HandlerTransactor) Withdraw(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _ERC721Handler.contract.Transact(opts, "withdraw", data)
}

// Withdraw is a paid mutator transaction binding the contract method 0x0968f264.
//
// Solidity: function withdraw(bytes data) returns()
func (_ERC721Handler *ERC721HandlerSession) Withdraw(data []byte) (*types.Transaction, error) {
	return _ERC721Handler.Contract.Withdraw(&_ERC721Handler.TransactOpts, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0x0968f264.
//
// Solidity: function withdraw(bytes data) returns()
func (_ERC721Handler *ERC721HandlerTransactorSession) Withdraw(data []byte) (*types.Transaction, error) {
	return _ERC721Handler.Contract.Withdraw(&_ERC721Handler.TransactOpts, data)
}
