// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ERC1155Handler

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

// ERC1155HandlerMetaData contains all meta data concerning the ERC1155Handler contract.
var ERC1155HandlerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridgeAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"_bridgeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_burnList\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_contractWhitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_lockMintUnlockList\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"_resourceIDToTokenContractAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_tokenContractAddressToResourceID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setBurnable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setLockMintUnlockable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setResource\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"depositer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"metaData\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"executeProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561001057600080fd5b506040516200138738038062001387833981016040819052610031916100c1565b6001600160601b0319606082901b166080526100536301ffc9a760e01b610069565b610063630271189760e51b610069565b50610126565b6001600160e01b0319808216141561009c5760405162461bcd60e51b8152600401610093906100ef565b60405180910390fd5b6001600160e01b0319166000908152600560205260409020805460ff19166001179055565b6000602082840312156100d2578081fd5b81516001600160a01b03811681146100e8578182fd5b9392505050565b6020808252601c908201527f4552433136353a20696e76616c696420696e7465726661636520696400000000604082015260600190565b60805160601c61123e62000149600039806102fe5280610529525061123e6000f3fe608060405234801561001057600080fd5b50600436106100f45760003560e01c80637f79bea811610097578063bc197c8111610066578063bc197c81146101f1578063c8ba6c8714610211578063e248cff214610231578063f23a6e6114610244576100f4565b80637f79bea814610198578063b07e54bb146101ab578063b1f06463146101cb578063b8fa3736146101de576100f4565b80630968f264116100d35780630968f2641461014a5780630a6d55d81461015d578063318c136e1461017d5780636a70d08114610185576100f4565b8062d8e7ee146100f957806301ffc9a71461010e57806307b7ed9914610137575b600080fd5b61010c61010736600461097f565b610257565b005b61012161011c366004610d8d565b61026b565b60405161012e9190610f5d565b60405180910390f35b61010c61014536600461097f565b61028a565b61010c610158366004610db5565b61029b565b61017061016b366004610ca2565b6102e1565b60405161012e9190610e56565b6101706102fc565b61012161019336600461097f565b610320565b6101216101a636600461097f565b610335565b6101be6101b9366004610ce9565b61034a565b60405161012e9190610f86565b6101216101d936600461097f565b610400565b61010c6101ec366004610cba565b610415565b6102046101ff366004610a88565b61042b565b60405161012e9190610f71565b61022461021f36600461097f565b61043c565b60405161012e9190610f68565b61010c61023f366004610d43565b61044e565b610204610252366004610b32565b61050d565b61025f61051e565b61026881610568565b50565b6001600160e01b03191660009081526005602052604090205460ff1690565b61029261051e565b610268816105fd565b6102a361051e565b6000806060806060858060200190518101906102bf91906109a2565b9398509196509450925090506102d9853086868686610692565b505050505050565b6000602081905290815260409020546001600160a01b031681565b7f000000000000000000000000000000000000000000000000000000000000000081565b60036020526000908152604090205460ff1681565b60026020526000908152604090205460ff1681565b606061035461051e565b60608061036384860186610b99565b60008981526020819052604090205491935091506001600160a01b0316806103a65760405162461bcd60e51b815260040161039d906110ce565b60405180910390fd5b6001600160a01b03811660009081526003602052604090205460ff16156103d8576103d381888585610703565b6103f5565b6103f5818830868660405180602001604052806000815250610692565b505050949350505050565b60046020526000908152604090205460ff1681565b61041d61051e565b6104278282610763565b5050565b63bc197c8160e01b95945050505050565b60016020526000908152604090205481565b61045661051e565b606080808061046785870187610bfa565b60208083015160008d81528083526040808220546001600160a01b031680835260029094529020549599509397509195509350909160ff166104bb5760405162461bcd60e51b815260040161039d90611110565b6001600160a01b03811660009081526003602052604090205460ff16156104f1576104ec818360601c8888876107b0565b610502565b61050281308460601c898988610692565b505050505050505050565b63f23a6e6160e01b95945050505050565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146105665760405162461bcd60e51b815260040161039d90610ff6565b565b6001600160a01b03811660009081526002602052604090205460ff166105a05760405162461bcd60e51b815260040161039d9061108a565b6001600160a01b03811660009081526003602052604090205460ff16156105d95760405162461bcd60e51b815260040161039d9061102d565b6001600160a01b03166000908152600460205260409020805460ff19166001179055565b6001600160a01b03811660009081526002602052604090205460ff166106355760405162461bcd60e51b815260040161039d9061108a565b6001600160a01b03811660009081526004602052604090205460ff161561066e5760405162461bcd60e51b815260040161039d90610f99565b6001600160a01b03166000908152600360205260409020805460ff19166001179055565b604051631759616b60e11b815286906001600160a01b03821690632eb2c2d6906106c89089908990899089908990600401610e6a565b600060405180830381600087803b1580156106e257600080fd5b505af11580156106f6573d6000803e3d6000fd5b5050505050505050505050565b604051631ac8311560e21b815284906001600160a01b03821690636b20c4549061073590879087908790600401610ec8565b600060405180830381600087803b15801561074f57600080fd5b505af1158015610502573d6000803e3d6000fd5b60008281526020818152604080832080546001600160a01b039095166001600160a01b0319909516851790559282526001808252838320949094556002905220805460ff19169091179055565b604051630fbfeffd60e11b815285906001600160a01b03821690631f7fdffa906107e4908890889088908890600401610f08565b600060405180830381600087803b1580156107fe57600080fd5b505af1158015610812573d6000803e3d6000fd5b50505050505050505050565b600082601f83011261082e578081fd5b813561084161083c8261117f565b611158565b81815291506020808301908481018184028601820187101561086257600080fd5b60005b8481101561088157813584529282019290820190600101610865565b505050505092915050565b600082601f83011261089c578081fd5b81516108aa61083c8261117f565b8181529150602080830190848101818402860182018710156108cb57600080fd5b60005b84811015610881578151845292820192908201906001016108ce565b60008083601f8401126108fb578182fd5b50813567ffffffffffffffff811115610912578182fd5b60208301915083602082850101111561092a57600080fd5b9250929050565b600082601f830112610941578081fd5b813561094f61083c8261119f565b915080825283602082850101111561096657600080fd5b8060208401602084013760009082016020015292915050565b600060208284031215610990578081fd5b813561099b816111f3565b9392505050565b600080600080600060a086880312156109b9578081fd5b85516109c4816111f3565b60208701519095506109d5816111f3565b604087015190945067ffffffffffffffff808211156109f2578283fd5b6109fe89838a0161088c565b94506060880151915080821115610a13578283fd5b610a1f89838a0161088c565b93506080880151915080821115610a34578283fd5b508601601f81018813610a45578182fd5b8051610a5361083c8261119f565b818152896020838501011115610a67578384fd5b610a788260208301602086016111c3565b8093505050509295509295909350565b600080600080600060a08688031215610a9f578081fd5b8535610aaa816111f3565b94506020860135610aba816111f3565b9350604086013567ffffffffffffffff80821115610ad6578283fd5b610ae289838a0161081e565b94506060880135915080821115610af7578283fd5b610b0389838a0161081e565b93506080880135915080821115610b18578283fd5b50610b2588828901610931565b9150509295509295909350565b600080600080600060a08688031215610b49578081fd5b8535610b54816111f3565b94506020860135610b64816111f3565b93506040860135925060608601359150608086013567ffffffffffffffff811115610b8d578182fd5b610b2588828901610931565b60008060408385031215610bab578182fd5b823567ffffffffffffffff80821115610bc2578384fd5b610bce8683870161081e565b93506020850135915080821115610be3578283fd5b50610bf08582860161081e565b9150509250929050565b60008060008060808587031215610c0f578384fd5b843567ffffffffffffffff80821115610c26578586fd5b610c328883890161081e565b95506020870135915080821115610c47578485fd5b610c538883890161081e565b94506040870135915080821115610c68578384fd5b610c7488838901610931565b93506060870135915080821115610c89578283fd5b50610c9687828801610931565b91505092959194509250565b600060208284031215610cb3578081fd5b5035919050565b60008060408385031215610ccc578182fd5b823591506020830135610cde816111f3565b809150509250929050565b60008060008060608587031215610cfe578182fd5b843593506020850135610d10816111f3565b9250604085013567ffffffffffffffff811115610d2b578283fd5b610d37878288016108ea565b95989497509550505050565b600080600060408486031215610d57578081fd5b83359250602084013567ffffffffffffffff811115610d74578182fd5b610d80868287016108ea565b9497909650939450505050565b600060208284031215610d9e578081fd5b81356001600160e01b03198116811461099b578182fd5b600060208284031215610dc6578081fd5b813567ffffffffffffffff811115610ddc578182fd5b610de884828501610931565b949350505050565b6000815180845260208085019450808401835b83811015610e1f57815187529582019590820190600101610e03565b509495945050505050565b60008151808452610e428160208601602086016111c3565b601f01601f19169290920160200192915050565b6001600160a01b0391909116815260200190565b6001600160a01b0386811682528516602082015260a060408201819052600090610e9690830186610df0565b8281036060840152610ea88186610df0565b90508281036080840152610ebc8185610e2a565b98975050505050505050565b6001600160a01b0384168152606060208201819052600090610eec90830185610df0565b8281036040840152610efe8185610df0565b9695505050505050565b6001600160a01b0385168152608060208201819052600090610f2c90830186610df0565b8281036040840152610f3e8186610df0565b90508281036060840152610f528185610e2a565b979650505050505050565b901515815260200190565b90815260200190565b6001600160e01b031991909116815260200190565b60006020825261099b6020830184610e2a565b6020808252603b908201527f70726f766964656420636f6e74726163742063616e6e6f74206265206c6f636b60408201527f4d696e74556e6c6f636b61626c6520616e64206275726e61626c650000000000606082015260800190565b6020808252601e908201527f73656e646572206d7573742062652062726964676520636f6e74726163740000604082015260600190565b6020808252603b908201527f70726f766964656420636f6e74726163742063616e6e6f74206265206275726e60408201527f61626c6520616e64206c6f636b4d696e74556e6c6f636b61626c650000000000606082015260800190565b60208082526024908201527f70726f766964656420636f6e7472616374206973206e6f742077686974656c696040820152631cdd195960e21b606082015260800190565b60208082526022908201527f70726f7669646564207265736f75726365494420646f6573206e6f74206578696040820152611cdd60f21b606082015260800190565b60208082526028908201527f70726f766964656420746f6b656e41646472657373206973206e6f74207768696040820152671d195b1a5cdd195960c21b606082015260800190565b60405181810167ffffffffffffffff8111828210171561117757600080fd5b604052919050565b600067ffffffffffffffff821115611195578081fd5b5060209081020190565b600067ffffffffffffffff8211156111b5578081fd5b50601f01601f191660200190565b60005b838110156111de5781810151838201526020016111c6565b838111156111ed576000848401525b50505050565b6001600160a01b038116811461026857600080fdfea2646970667358221220f28172d000063d6514b8b4a07523d74c01a38368329c8a11c033b86c0bfc9bc264736f6c634300060c0033",
}

// ERC1155HandlerABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC1155HandlerMetaData.ABI instead.
var ERC1155HandlerABI = ERC1155HandlerMetaData.ABI

// ERC1155HandlerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ERC1155HandlerMetaData.Bin instead.
var ERC1155HandlerBin = ERC1155HandlerMetaData.Bin

// DeployERC1155Handler deploys a new Ethereum contract, binding an instance of ERC1155Handler to it.
func DeployERC1155Handler(auth *bind.TransactOpts, backend bind.ContractBackend, bridgeAddress common.Address) (common.Address, *types.Transaction, *ERC1155Handler, error) {
	parsed, err := ERC1155HandlerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ERC1155HandlerBin), backend, bridgeAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC1155Handler{ERC1155HandlerCaller: ERC1155HandlerCaller{contract: contract}, ERC1155HandlerTransactor: ERC1155HandlerTransactor{contract: contract}, ERC1155HandlerFilterer: ERC1155HandlerFilterer{contract: contract}}, nil
}

// ERC1155Handler is an auto generated Go binding around an Ethereum contract.
type ERC1155Handler struct {
	ERC1155HandlerCaller     // Read-only binding to the contract
	ERC1155HandlerTransactor // Write-only binding to the contract
	ERC1155HandlerFilterer   // Log filterer for contract events
}

// ERC1155HandlerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC1155HandlerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1155HandlerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC1155HandlerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1155HandlerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC1155HandlerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1155HandlerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC1155HandlerSession struct {
	Contract     *ERC1155Handler   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC1155HandlerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC1155HandlerCallerSession struct {
	Contract *ERC1155HandlerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ERC1155HandlerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC1155HandlerTransactorSession struct {
	Contract     *ERC1155HandlerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ERC1155HandlerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC1155HandlerRaw struct {
	Contract *ERC1155Handler // Generic contract binding to access the raw methods on
}

// ERC1155HandlerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC1155HandlerCallerRaw struct {
	Contract *ERC1155HandlerCaller // Generic read-only contract binding to access the raw methods on
}

// ERC1155HandlerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC1155HandlerTransactorRaw struct {
	Contract *ERC1155HandlerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC1155Handler creates a new instance of ERC1155Handler, bound to a specific deployed contract.
func NewERC1155Handler(address common.Address, backend bind.ContractBackend) (*ERC1155Handler, error) {
	contract, err := bindERC1155Handler(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC1155Handler{ERC1155HandlerCaller: ERC1155HandlerCaller{contract: contract}, ERC1155HandlerTransactor: ERC1155HandlerTransactor{contract: contract}, ERC1155HandlerFilterer: ERC1155HandlerFilterer{contract: contract}}, nil
}

// NewERC1155HandlerCaller creates a new read-only instance of ERC1155Handler, bound to a specific deployed contract.
func NewERC1155HandlerCaller(address common.Address, caller bind.ContractCaller) (*ERC1155HandlerCaller, error) {
	contract, err := bindERC1155Handler(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC1155HandlerCaller{contract: contract}, nil
}

// NewERC1155HandlerTransactor creates a new write-only instance of ERC1155Handler, bound to a specific deployed contract.
func NewERC1155HandlerTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC1155HandlerTransactor, error) {
	contract, err := bindERC1155Handler(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC1155HandlerTransactor{contract: contract}, nil
}

// NewERC1155HandlerFilterer creates a new log filterer instance of ERC1155Handler, bound to a specific deployed contract.
func NewERC1155HandlerFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC1155HandlerFilterer, error) {
	contract, err := bindERC1155Handler(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC1155HandlerFilterer{contract: contract}, nil
}

// bindERC1155Handler binds a generic wrapper to an already deployed contract.
func bindERC1155Handler(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC1155HandlerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC1155Handler *ERC1155HandlerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC1155Handler.Contract.ERC1155HandlerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC1155Handler *ERC1155HandlerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC1155Handler.Contract.ERC1155HandlerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC1155Handler *ERC1155HandlerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC1155Handler.Contract.ERC1155HandlerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC1155Handler *ERC1155HandlerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC1155Handler.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC1155Handler *ERC1155HandlerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC1155Handler.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC1155Handler *ERC1155HandlerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC1155Handler.Contract.contract.Transact(opts, method, params...)
}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_ERC1155Handler *ERC1155HandlerCaller) BridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC1155Handler.contract.Call(opts, &out, "_bridgeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_ERC1155Handler *ERC1155HandlerSession) BridgeAddress() (common.Address, error) {
	return _ERC1155Handler.Contract.BridgeAddress(&_ERC1155Handler.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_ERC1155Handler *ERC1155HandlerCallerSession) BridgeAddress() (common.Address, error) {
	return _ERC1155Handler.Contract.BridgeAddress(&_ERC1155Handler.CallOpts)
}

// BurnList is a free data retrieval call binding the contract method 0x6a70d081.
//
// Solidity: function _burnList(address ) view returns(bool)
func (_ERC1155Handler *ERC1155HandlerCaller) BurnList(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _ERC1155Handler.contract.Call(opts, &out, "_burnList", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BurnList is a free data retrieval call binding the contract method 0x6a70d081.
//
// Solidity: function _burnList(address ) view returns(bool)
func (_ERC1155Handler *ERC1155HandlerSession) BurnList(arg0 common.Address) (bool, error) {
	return _ERC1155Handler.Contract.BurnList(&_ERC1155Handler.CallOpts, arg0)
}

// BurnList is a free data retrieval call binding the contract method 0x6a70d081.
//
// Solidity: function _burnList(address ) view returns(bool)
func (_ERC1155Handler *ERC1155HandlerCallerSession) BurnList(arg0 common.Address) (bool, error) {
	return _ERC1155Handler.Contract.BurnList(&_ERC1155Handler.CallOpts, arg0)
}

// ContractWhitelist is a free data retrieval call binding the contract method 0x7f79bea8.
//
// Solidity: function _contractWhitelist(address ) view returns(bool)
func (_ERC1155Handler *ERC1155HandlerCaller) ContractWhitelist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _ERC1155Handler.contract.Call(opts, &out, "_contractWhitelist", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ContractWhitelist is a free data retrieval call binding the contract method 0x7f79bea8.
//
// Solidity: function _contractWhitelist(address ) view returns(bool)
func (_ERC1155Handler *ERC1155HandlerSession) ContractWhitelist(arg0 common.Address) (bool, error) {
	return _ERC1155Handler.Contract.ContractWhitelist(&_ERC1155Handler.CallOpts, arg0)
}

// ContractWhitelist is a free data retrieval call binding the contract method 0x7f79bea8.
//
// Solidity: function _contractWhitelist(address ) view returns(bool)
func (_ERC1155Handler *ERC1155HandlerCallerSession) ContractWhitelist(arg0 common.Address) (bool, error) {
	return _ERC1155Handler.Contract.ContractWhitelist(&_ERC1155Handler.CallOpts, arg0)
}

// LockMintUnlockList is a free data retrieval call binding the contract method 0xb1f06463.
//
// Solidity: function _lockMintUnlockList(address ) view returns(bool)
func (_ERC1155Handler *ERC1155HandlerCaller) LockMintUnlockList(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _ERC1155Handler.contract.Call(opts, &out, "_lockMintUnlockList", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// LockMintUnlockList is a free data retrieval call binding the contract method 0xb1f06463.
//
// Solidity: function _lockMintUnlockList(address ) view returns(bool)
func (_ERC1155Handler *ERC1155HandlerSession) LockMintUnlockList(arg0 common.Address) (bool, error) {
	return _ERC1155Handler.Contract.LockMintUnlockList(&_ERC1155Handler.CallOpts, arg0)
}

// LockMintUnlockList is a free data retrieval call binding the contract method 0xb1f06463.
//
// Solidity: function _lockMintUnlockList(address ) view returns(bool)
func (_ERC1155Handler *ERC1155HandlerCallerSession) LockMintUnlockList(arg0 common.Address) (bool, error) {
	return _ERC1155Handler.Contract.LockMintUnlockList(&_ERC1155Handler.CallOpts, arg0)
}

// ResourceIDToTokenContractAddress is a free data retrieval call binding the contract method 0x0a6d55d8.
//
// Solidity: function _resourceIDToTokenContractAddress(bytes32 ) view returns(address)
func (_ERC1155Handler *ERC1155HandlerCaller) ResourceIDToTokenContractAddress(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _ERC1155Handler.contract.Call(opts, &out, "_resourceIDToTokenContractAddress", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ResourceIDToTokenContractAddress is a free data retrieval call binding the contract method 0x0a6d55d8.
//
// Solidity: function _resourceIDToTokenContractAddress(bytes32 ) view returns(address)
func (_ERC1155Handler *ERC1155HandlerSession) ResourceIDToTokenContractAddress(arg0 [32]byte) (common.Address, error) {
	return _ERC1155Handler.Contract.ResourceIDToTokenContractAddress(&_ERC1155Handler.CallOpts, arg0)
}

// ResourceIDToTokenContractAddress is a free data retrieval call binding the contract method 0x0a6d55d8.
//
// Solidity: function _resourceIDToTokenContractAddress(bytes32 ) view returns(address)
func (_ERC1155Handler *ERC1155HandlerCallerSession) ResourceIDToTokenContractAddress(arg0 [32]byte) (common.Address, error) {
	return _ERC1155Handler.Contract.ResourceIDToTokenContractAddress(&_ERC1155Handler.CallOpts, arg0)
}

// TokenContractAddressToResourceID is a free data retrieval call binding the contract method 0xc8ba6c87.
//
// Solidity: function _tokenContractAddressToResourceID(address ) view returns(bytes32)
func (_ERC1155Handler *ERC1155HandlerCaller) TokenContractAddressToResourceID(opts *bind.CallOpts, arg0 common.Address) ([32]byte, error) {
	var out []interface{}
	err := _ERC1155Handler.contract.Call(opts, &out, "_tokenContractAddressToResourceID", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TokenContractAddressToResourceID is a free data retrieval call binding the contract method 0xc8ba6c87.
//
// Solidity: function _tokenContractAddressToResourceID(address ) view returns(bytes32)
func (_ERC1155Handler *ERC1155HandlerSession) TokenContractAddressToResourceID(arg0 common.Address) ([32]byte, error) {
	return _ERC1155Handler.Contract.TokenContractAddressToResourceID(&_ERC1155Handler.CallOpts, arg0)
}

// TokenContractAddressToResourceID is a free data retrieval call binding the contract method 0xc8ba6c87.
//
// Solidity: function _tokenContractAddressToResourceID(address ) view returns(bytes32)
func (_ERC1155Handler *ERC1155HandlerCallerSession) TokenContractAddressToResourceID(arg0 common.Address) ([32]byte, error) {
	return _ERC1155Handler.Contract.TokenContractAddressToResourceID(&_ERC1155Handler.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC1155Handler *ERC1155HandlerCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ERC1155Handler.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC1155Handler *ERC1155HandlerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC1155Handler.Contract.SupportsInterface(&_ERC1155Handler.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC1155Handler *ERC1155HandlerCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC1155Handler.Contract.SupportsInterface(&_ERC1155Handler.CallOpts, interfaceId)
}

// Deposit is a paid mutator transaction binding the contract method 0xb07e54bb.
//
// Solidity: function deposit(bytes32 resourceID, address depositer, bytes data) returns(bytes metaData)
func (_ERC1155Handler *ERC1155HandlerTransactor) Deposit(opts *bind.TransactOpts, resourceID [32]byte, depositer common.Address, data []byte) (*types.Transaction, error) {
	return _ERC1155Handler.contract.Transact(opts, "deposit", resourceID, depositer, data)
}

// Deposit is a paid mutator transaction binding the contract method 0xb07e54bb.
//
// Solidity: function deposit(bytes32 resourceID, address depositer, bytes data) returns(bytes metaData)
func (_ERC1155Handler *ERC1155HandlerSession) Deposit(resourceID [32]byte, depositer common.Address, data []byte) (*types.Transaction, error) {
	return _ERC1155Handler.Contract.Deposit(&_ERC1155Handler.TransactOpts, resourceID, depositer, data)
}

// Deposit is a paid mutator transaction binding the contract method 0xb07e54bb.
//
// Solidity: function deposit(bytes32 resourceID, address depositer, bytes data) returns(bytes metaData)
func (_ERC1155Handler *ERC1155HandlerTransactorSession) Deposit(resourceID [32]byte, depositer common.Address, data []byte) (*types.Transaction, error) {
	return _ERC1155Handler.Contract.Deposit(&_ERC1155Handler.TransactOpts, resourceID, depositer, data)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0xe248cff2.
//
// Solidity: function executeProposal(bytes32 resourceID, bytes data) returns()
func (_ERC1155Handler *ERC1155HandlerTransactor) ExecuteProposal(opts *bind.TransactOpts, resourceID [32]byte, data []byte) (*types.Transaction, error) {
	return _ERC1155Handler.contract.Transact(opts, "executeProposal", resourceID, data)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0xe248cff2.
//
// Solidity: function executeProposal(bytes32 resourceID, bytes data) returns()
func (_ERC1155Handler *ERC1155HandlerSession) ExecuteProposal(resourceID [32]byte, data []byte) (*types.Transaction, error) {
	return _ERC1155Handler.Contract.ExecuteProposal(&_ERC1155Handler.TransactOpts, resourceID, data)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0xe248cff2.
//
// Solidity: function executeProposal(bytes32 resourceID, bytes data) returns()
func (_ERC1155Handler *ERC1155HandlerTransactorSession) ExecuteProposal(resourceID [32]byte, data []byte) (*types.Transaction, error) {
	return _ERC1155Handler.Contract.ExecuteProposal(&_ERC1155Handler.TransactOpts, resourceID, data)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_ERC1155Handler *ERC1155HandlerTransactor) OnERC1155BatchReceived(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _ERC1155Handler.contract.Transact(opts, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_ERC1155Handler *ERC1155HandlerSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _ERC1155Handler.Contract.OnERC1155BatchReceived(&_ERC1155Handler.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_ERC1155Handler *ERC1155HandlerTransactorSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _ERC1155Handler.Contract.OnERC1155BatchReceived(&_ERC1155Handler.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_ERC1155Handler *ERC1155HandlerTransactor) OnERC1155Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _ERC1155Handler.contract.Transact(opts, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_ERC1155Handler *ERC1155HandlerSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _ERC1155Handler.Contract.OnERC1155Received(&_ERC1155Handler.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_ERC1155Handler *ERC1155HandlerTransactorSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _ERC1155Handler.Contract.OnERC1155Received(&_ERC1155Handler.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// SetBurnable is a paid mutator transaction binding the contract method 0x07b7ed99.
//
// Solidity: function setBurnable(address contractAddress) returns()
func (_ERC1155Handler *ERC1155HandlerTransactor) SetBurnable(opts *bind.TransactOpts, contractAddress common.Address) (*types.Transaction, error) {
	return _ERC1155Handler.contract.Transact(opts, "setBurnable", contractAddress)
}

// SetBurnable is a paid mutator transaction binding the contract method 0x07b7ed99.
//
// Solidity: function setBurnable(address contractAddress) returns()
func (_ERC1155Handler *ERC1155HandlerSession) SetBurnable(contractAddress common.Address) (*types.Transaction, error) {
	return _ERC1155Handler.Contract.SetBurnable(&_ERC1155Handler.TransactOpts, contractAddress)
}

// SetBurnable is a paid mutator transaction binding the contract method 0x07b7ed99.
//
// Solidity: function setBurnable(address contractAddress) returns()
func (_ERC1155Handler *ERC1155HandlerTransactorSession) SetBurnable(contractAddress common.Address) (*types.Transaction, error) {
	return _ERC1155Handler.Contract.SetBurnable(&_ERC1155Handler.TransactOpts, contractAddress)
}

// SetLockMintUnlockable is a paid mutator transaction binding the contract method 0x00d8e7ee.
//
// Solidity: function setLockMintUnlockable(address contractAddress) returns()
func (_ERC1155Handler *ERC1155HandlerTransactor) SetLockMintUnlockable(opts *bind.TransactOpts, contractAddress common.Address) (*types.Transaction, error) {
	return _ERC1155Handler.contract.Transact(opts, "setLockMintUnlockable", contractAddress)
}

// SetLockMintUnlockable is a paid mutator transaction binding the contract method 0x00d8e7ee.
//
// Solidity: function setLockMintUnlockable(address contractAddress) returns()
func (_ERC1155Handler *ERC1155HandlerSession) SetLockMintUnlockable(contractAddress common.Address) (*types.Transaction, error) {
	return _ERC1155Handler.Contract.SetLockMintUnlockable(&_ERC1155Handler.TransactOpts, contractAddress)
}

// SetLockMintUnlockable is a paid mutator transaction binding the contract method 0x00d8e7ee.
//
// Solidity: function setLockMintUnlockable(address contractAddress) returns()
func (_ERC1155Handler *ERC1155HandlerTransactorSession) SetLockMintUnlockable(contractAddress common.Address) (*types.Transaction, error) {
	return _ERC1155Handler.Contract.SetLockMintUnlockable(&_ERC1155Handler.TransactOpts, contractAddress)
}

// SetResource is a paid mutator transaction binding the contract method 0xb8fa3736.
//
// Solidity: function setResource(bytes32 resourceID, address contractAddress) returns()
func (_ERC1155Handler *ERC1155HandlerTransactor) SetResource(opts *bind.TransactOpts, resourceID [32]byte, contractAddress common.Address) (*types.Transaction, error) {
	return _ERC1155Handler.contract.Transact(opts, "setResource", resourceID, contractAddress)
}

// SetResource is a paid mutator transaction binding the contract method 0xb8fa3736.
//
// Solidity: function setResource(bytes32 resourceID, address contractAddress) returns()
func (_ERC1155Handler *ERC1155HandlerSession) SetResource(resourceID [32]byte, contractAddress common.Address) (*types.Transaction, error) {
	return _ERC1155Handler.Contract.SetResource(&_ERC1155Handler.TransactOpts, resourceID, contractAddress)
}

// SetResource is a paid mutator transaction binding the contract method 0xb8fa3736.
//
// Solidity: function setResource(bytes32 resourceID, address contractAddress) returns()
func (_ERC1155Handler *ERC1155HandlerTransactorSession) SetResource(resourceID [32]byte, contractAddress common.Address) (*types.Transaction, error) {
	return _ERC1155Handler.Contract.SetResource(&_ERC1155Handler.TransactOpts, resourceID, contractAddress)
}

// Withdraw is a paid mutator transaction binding the contract method 0x0968f264.
//
// Solidity: function withdraw(bytes data) returns()
func (_ERC1155Handler *ERC1155HandlerTransactor) Withdraw(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _ERC1155Handler.contract.Transact(opts, "withdraw", data)
}

// Withdraw is a paid mutator transaction binding the contract method 0x0968f264.
//
// Solidity: function withdraw(bytes data) returns()
func (_ERC1155Handler *ERC1155HandlerSession) Withdraw(data []byte) (*types.Transaction, error) {
	return _ERC1155Handler.Contract.Withdraw(&_ERC1155Handler.TransactOpts, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0x0968f264.
//
// Solidity: function withdraw(bytes data) returns()
func (_ERC1155Handler *ERC1155HandlerTransactorSession) Withdraw(data []byte) (*types.Transaction, error) {
	return _ERC1155Handler.Contract.Withdraw(&_ERC1155Handler.TransactOpts, data)
}
