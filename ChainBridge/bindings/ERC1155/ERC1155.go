// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ERC1155

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

// ERC1155MetaData contains all meta data concerning the ERC1155 contract.
var ERC1155MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uri_\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferSingle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"URI\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"uri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeBatchTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040516200182e3803806200182e833981810160405260208110156200003757600080fd5b81019080805160405193929190846401000000008211156200005857600080fd5b9083019060208201858111156200006e57600080fd5b82516401000000008111828201881017156200008957600080fd5b82525081516020918201929091019080838360005b83811015620000b85781810151838201526020016200009e565b50505050905090810190601f168015620000e65780820380516001836020036101000a031916815260200191505b50604052506200010191506301ffc9a760e01b905062000137565b6200010c81620001bc565b6200011e636cdb3d1360e11b62000137565b620001306303a24d0760e21b62000137565b5062000271565b6001600160e01b0319808216141562000197576040805162461bcd60e51b815260206004820152601c60248201527f4552433136353a20696e76616c696420696e7465726661636520696400000000604482015290519081900360640190fd5b6001600160e01b0319166000908152602081905260409020805460ff19166001179055565b8051620001d1906003906020840190620001d5565b5050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200021857805160ff191683800117855562000248565b8280016001018555821562000248579182015b82811115620002485782518255916020019190600101906200022b565b50620002569291506200025a565b5090565b5b808211156200025657600081556001016200025b565b6115ad80620002816000396000f3fe608060405234801561001057600080fd5b50600436106100875760003560e01c80634e1273f41161005b5780634e1273f41461035a578063a22cb465146104cd578063e985e9c5146104fb578063f242432a1461052957610087565b8062fdd58e1461008c57806301ffc9a7146100ca5780630e89341c146101055780632eb2c2d614610197575b600080fd5b6100b8600480360360408110156100a257600080fd5b506001600160a01b0381351690602001356105f2565b60408051918252519081900360200190f35b6100f1600480360360208110156100e057600080fd5b50356001600160e01b031916610661565b604080519115158252519081900360200190f35b6101226004803603602081101561011b57600080fd5b5035610680565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561015c578181015183820152602001610144565b50505050905090810190601f1680156101895780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b610358600480360360a08110156101ad57600080fd5b6001600160a01b038235811692602081013590911691810190606081016040820135600160201b8111156101e057600080fd5b8201836020820111156101f257600080fd5b803590602001918460208302840111600160201b8311171561021357600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561026257600080fd5b82018360208201111561027457600080fd5b803590602001918460208302840111600160201b8311171561029557600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156102e457600080fd5b8201836020820111156102f657600080fd5b803590602001918460018302840111600160201b8311171561031757600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610718945050505050565b005b61047d6004803603604081101561037057600080fd5b810190602081018135600160201b81111561038a57600080fd5b82018360208201111561039c57600080fd5b803590602001918460208302840111600160201b831117156103bd57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561040c57600080fd5b82018360208201111561041e57600080fd5b803590602001918460208302840111600160201b8311171561043f57600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550610a16945050505050565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156104b95781810151838201526020016104a1565b505050509050019250505060405180910390f35b610358600480360360408110156104e357600080fd5b506001600160a01b0381351690602001351515610b02565b6100f16004803603604081101561051157600080fd5b506001600160a01b0381358116916020013516610bf1565b610358600480360360a081101561053f57600080fd5b6001600160a01b03823581169260208101359091169160408201359160608101359181019060a081016080820135600160201b81111561057e57600080fd5b82018360208201111561059057600080fd5b803590602001918460018302840111600160201b831117156105b157600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610c1f945050505050565b60006001600160a01b0383166106395760405162461bcd60e51b815260040180806020018281038252602b815260200180611429602b913960400191505060405180910390fd5b5060009081526001602090815260408083206001600160a01b03949094168352929052205490565b6001600160e01b03191660009081526020819052604090205460ff1690565b60038054604080516020601f600260001961010060018816150201909516949094049384018190048102820181019092528281526060939092909183018282801561070c5780601f106106e15761010080835404028352916020019161070c565b820191906000526020600020905b8154815290600101906020018083116106ef57829003601f168201915b50505050509050919050565b81518351146107585760405162461bcd60e51b81526004018080602001828103825260288152602001806115506028913960400191505060405180910390fd5b6001600160a01b03841661079d5760405162461bcd60e51b815260040180806020018281038252602581526020018061147d6025913960400191505060405180910390fd5b6107a5610dea565b6001600160a01b0316856001600160a01b031614806107d057506107d0856107cb610dea565b610bf1565b61080b5760405162461bcd60e51b81526004018080602001828103825260328152602001806114a26032913960400191505060405180910390fd5b6000610815610dea565b9050610825818787878787610a0e565b60005b845181101561092657600085828151811061083f57fe5b60200260200101519050600085838151811061085757fe5b602002602001015190506108c4816040518060600160405280602a81526020016114d4602a91396001600086815260200190815260200160002060008d6001600160a01b03166001600160a01b0316815260200190815260200160002054610def9092919063ffffffff16565b60008381526001602090815260408083206001600160a01b038e811685529252808320939093558a16815220546108fb9082610e86565b60009283526001602081815260408086206001600160a01b038d168752909152909320555001610828565b50846001600160a01b0316866001600160a01b0316826001600160a01b03167f4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb8787604051808060200180602001838103835285818151815260200191508051906020019060200280838360005b838110156109ac578181015183820152602001610994565b50505050905001838103825284818151815260200191508051906020019060200280838360005b838110156109eb5781810151838201526020016109d3565b5050505090500194505050505060405180910390a4610a0e818787878787610ee7565b505050505050565b60608151835114610a585760405162461bcd60e51b81526004018080602001828103825260298152602001806115276029913960400191505060405180910390fd5b6060835167ffffffffffffffff81118015610a7257600080fd5b50604051908082528060200260200182016040528015610a9c578160200160208202803683370190505b50905060005b8451811015610afa57610adb858281518110610aba57fe5b6020026020010151858381518110610ace57fe5b60200260200101516105f2565b828281518110610ae757fe5b6020908102919091010152600101610aa2565b509392505050565b816001600160a01b0316610b14610dea565b6001600160a01b03161415610b5a5760405162461bcd60e51b81526004018080602001828103825260298152602001806114fe6029913960400191505060405180910390fd5b8060026000610b67610dea565b6001600160a01b03908116825260208083019390935260409182016000908120918716808252919093529120805460ff191692151592909217909155610bab610dea565b6001600160a01b03167f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c318360405180821515815260200191505060405180910390a35050565b6001600160a01b03918216600090815260026020908152604080832093909416825291909152205460ff1690565b6001600160a01b038416610c645760405162461bcd60e51b815260040180806020018281038252602581526020018061147d6025913960400191505060405180910390fd5b610c6c610dea565b6001600160a01b0316856001600160a01b03161480610c925750610c92856107cb610dea565b610ccd5760405162461bcd60e51b81526004018080602001828103825260298152602001806114546029913960400191505060405180910390fd5b6000610cd7610dea565b9050610cf7818787610ce888611166565b610cf188611166565b87610a0e565b610d3e836040518060600160405280602a81526020016114d4602a913960008781526001602090815260408083206001600160a01b038d1684529091529020549190610def565b60008581526001602090815260408083206001600160a01b038b81168552925280832093909355871681522054610d759084610e86565b60008581526001602090815260408083206001600160a01b03808b168086529184529382902094909455805188815291820187905280518a8416938616927fc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f6292908290030190a4610a0e8187878787876111aa565b335b90565b60008184841115610e7e5760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015610e43578181015183820152602001610e2b565b50505050905090810190601f168015610e705780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505050900390565b600082820183811015610ee0576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b610ef9846001600160a01b031661131b565b15610a0e57836001600160a01b031663bc197c8187878686866040518663ffffffff1660e01b815260040180866001600160a01b03168152602001856001600160a01b03168152602001806020018060200180602001848103845287818151815260200191508051906020019060200280838360005b83811015610f87578181015183820152602001610f6f565b50505050905001848103835286818151815260200191508051906020019060200280838360005b83811015610fc6578181015183820152602001610fae565b50505050905001848103825285818151815260200191508051906020019080838360005b83811015611002578181015183820152602001610fea565b50505050905090810190601f16801561102f5780820380516001836020036101000a031916815260200191505b5098505050505050505050602060405180830381600087803b15801561105457600080fd5b505af192505050801561107957506040513d602081101561107457600080fd5b505160015b61110e57611085611327565b8061109057506110d7565b60405162461bcd60e51b8152602060048201818152835160248401528351849391928392604401919085019080838360008315610e43578181015183820152602001610e2b565b60405162461bcd60e51b81526004018080602001828103825260348152602001806113cd6034913960400191505060405180910390fd5b6001600160e01b0319811663bc197c8160e01b1461115d5760405162461bcd60e51b81526004018080602001828103825260288152602001806114016028913960400191505060405180910390fd5b50505050505050565b60408051600180825281830190925260609182919060208083019080368337019050509050828160008151811061119957fe5b602090810291909101015292915050565b6111bc846001600160a01b031661131b565b15610a0e57836001600160a01b031663f23a6e6187878686866040518663ffffffff1660e01b815260040180866001600160a01b03168152602001856001600160a01b0316815260200184815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561124b578181015183820152602001611233565b50505050905090810190601f1680156112785780820380516001836020036101000a031916815260200191505b509650505050505050602060405180830381600087803b15801561129b57600080fd5b505af19250505080156112c057506040513d60208110156112bb57600080fd5b505160015b6112cc57611085611327565b6001600160e01b0319811663f23a6e6160e01b1461115d5760405162461bcd60e51b81526004018080602001828103825260288152602001806114016028913960400191505060405180910390fd5b3b151590565b60e01c90565b600060443d101561133757610dec565b600481823e6308c379a061134b8251611321565b1461135557610dec565b6040513d600319016004823e80513d67ffffffffffffffff81602484011181841117156113855750505050610dec565b8284019250825191508082111561139f5750505050610dec565b503d830160208284010111156113b757505050610dec565b601f01601f191681016020016040529150509056fe455243313135353a207472616e7366657220746f206e6f6e2045524331313535526563656976657220696d706c656d656e746572455243313135353a204552433131353552656365697665722072656a656374656420746f6b656e73455243313135353a2062616c616e636520717565727920666f7220746865207a65726f2061646472657373455243313135353a2063616c6c6572206973206e6f74206f776e6572206e6f7220617070726f766564455243313135353a207472616e7366657220746f20746865207a65726f2061646472657373455243313135353a207472616e736665722063616c6c6572206973206e6f74206f776e6572206e6f7220617070726f766564455243313135353a20696e73756666696369656e742062616c616e636520666f72207472616e73666572455243313135353a2073657474696e6720617070726f76616c2073746174757320666f722073656c66455243313135353a206163636f756e747320616e6420696473206c656e677468206d69736d61746368455243313135353a2069647320616e6420616d6f756e7473206c656e677468206d69736d61746368a26469706673582212209f68f1f423f8203ed19a867b53896e31477b10d7b0719535eeed1b1c4e6c7f8464736f6c634300060c0033",
}

// ERC1155ABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC1155MetaData.ABI instead.
var ERC1155ABI = ERC1155MetaData.ABI

// ERC1155Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ERC1155MetaData.Bin instead.
var ERC1155Bin = ERC1155MetaData.Bin

// DeployERC1155 deploys a new Ethereum contract, binding an instance of ERC1155 to it.
func DeployERC1155(auth *bind.TransactOpts, backend bind.ContractBackend, uri_ string) (common.Address, *types.Transaction, *ERC1155, error) {
	parsed, err := ERC1155MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ERC1155Bin), backend, uri_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC1155{ERC1155Caller: ERC1155Caller{contract: contract}, ERC1155Transactor: ERC1155Transactor{contract: contract}, ERC1155Filterer: ERC1155Filterer{contract: contract}}, nil
}

// ERC1155 is an auto generated Go binding around an Ethereum contract.
type ERC1155 struct {
	ERC1155Caller     // Read-only binding to the contract
	ERC1155Transactor // Write-only binding to the contract
	ERC1155Filterer   // Log filterer for contract events
}

// ERC1155Caller is an auto generated read-only Go binding around an Ethereum contract.
type ERC1155Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1155Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC1155Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1155Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC1155Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1155Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC1155Session struct {
	Contract     *ERC1155          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC1155CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC1155CallerSession struct {
	Contract *ERC1155Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ERC1155TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC1155TransactorSession struct {
	Contract     *ERC1155Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ERC1155Raw is an auto generated low-level Go binding around an Ethereum contract.
type ERC1155Raw struct {
	Contract *ERC1155 // Generic contract binding to access the raw methods on
}

// ERC1155CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC1155CallerRaw struct {
	Contract *ERC1155Caller // Generic read-only contract binding to access the raw methods on
}

// ERC1155TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC1155TransactorRaw struct {
	Contract *ERC1155Transactor // Generic write-only contract binding to access the raw methods on
}

// NewERC1155 creates a new instance of ERC1155, bound to a specific deployed contract.
func NewERC1155(address common.Address, backend bind.ContractBackend) (*ERC1155, error) {
	contract, err := bindERC1155(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC1155{ERC1155Caller: ERC1155Caller{contract: contract}, ERC1155Transactor: ERC1155Transactor{contract: contract}, ERC1155Filterer: ERC1155Filterer{contract: contract}}, nil
}

// NewERC1155Caller creates a new read-only instance of ERC1155, bound to a specific deployed contract.
func NewERC1155Caller(address common.Address, caller bind.ContractCaller) (*ERC1155Caller, error) {
	contract, err := bindERC1155(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC1155Caller{contract: contract}, nil
}

// NewERC1155Transactor creates a new write-only instance of ERC1155, bound to a specific deployed contract.
func NewERC1155Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC1155Transactor, error) {
	contract, err := bindERC1155(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC1155Transactor{contract: contract}, nil
}

// NewERC1155Filterer creates a new log filterer instance of ERC1155, bound to a specific deployed contract.
func NewERC1155Filterer(address common.Address, filterer bind.ContractFilterer) (*ERC1155Filterer, error) {
	contract, err := bindERC1155(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC1155Filterer{contract: contract}, nil
}

// bindERC1155 binds a generic wrapper to an already deployed contract.
func bindERC1155(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC1155ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC1155 *ERC1155Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC1155.Contract.ERC1155Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC1155 *ERC1155Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC1155.Contract.ERC1155Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC1155 *ERC1155Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC1155.Contract.ERC1155Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC1155 *ERC1155CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC1155.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC1155 *ERC1155TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC1155.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC1155 *ERC1155TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC1155.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_ERC1155 *ERC1155Caller) BalanceOf(opts *bind.CallOpts, account common.Address, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ERC1155.contract.Call(opts, &out, "balanceOf", account, id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_ERC1155 *ERC1155Session) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _ERC1155.Contract.BalanceOf(&_ERC1155.CallOpts, account, id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_ERC1155 *ERC1155CallerSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _ERC1155.Contract.BalanceOf(&_ERC1155.CallOpts, account, id)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_ERC1155 *ERC1155Caller) BalanceOfBatch(opts *bind.CallOpts, accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _ERC1155.contract.Call(opts, &out, "balanceOfBatch", accounts, ids)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_ERC1155 *ERC1155Session) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _ERC1155.Contract.BalanceOfBatch(&_ERC1155.CallOpts, accounts, ids)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_ERC1155 *ERC1155CallerSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _ERC1155.Contract.BalanceOfBatch(&_ERC1155.CallOpts, accounts, ids)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_ERC1155 *ERC1155Caller) IsApprovedForAll(opts *bind.CallOpts, account common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _ERC1155.contract.Call(opts, &out, "isApprovedForAll", account, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_ERC1155 *ERC1155Session) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _ERC1155.Contract.IsApprovedForAll(&_ERC1155.CallOpts, account, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_ERC1155 *ERC1155CallerSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _ERC1155.Contract.IsApprovedForAll(&_ERC1155.CallOpts, account, operator)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC1155 *ERC1155Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ERC1155.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC1155 *ERC1155Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC1155.Contract.SupportsInterface(&_ERC1155.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC1155 *ERC1155CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC1155.Contract.SupportsInterface(&_ERC1155.CallOpts, interfaceId)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_ERC1155 *ERC1155Caller) Uri(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _ERC1155.contract.Call(opts, &out, "uri", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_ERC1155 *ERC1155Session) Uri(arg0 *big.Int) (string, error) {
	return _ERC1155.Contract.Uri(&_ERC1155.CallOpts, arg0)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_ERC1155 *ERC1155CallerSession) Uri(arg0 *big.Int) (string, error) {
	return _ERC1155.Contract.Uri(&_ERC1155.CallOpts, arg0)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_ERC1155 *ERC1155Transactor) SafeBatchTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _ERC1155.contract.Transact(opts, "safeBatchTransferFrom", from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_ERC1155 *ERC1155Session) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _ERC1155.Contract.SafeBatchTransferFrom(&_ERC1155.TransactOpts, from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_ERC1155 *ERC1155TransactorSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _ERC1155.Contract.SafeBatchTransferFrom(&_ERC1155.TransactOpts, from, to, ids, amounts, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_ERC1155 *ERC1155Transactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC1155.contract.Transact(opts, "safeTransferFrom", from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_ERC1155 *ERC1155Session) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC1155.Contract.SafeTransferFrom(&_ERC1155.TransactOpts, from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_ERC1155 *ERC1155TransactorSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC1155.Contract.SafeTransferFrom(&_ERC1155.TransactOpts, from, to, id, amount, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ERC1155 *ERC1155Transactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _ERC1155.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ERC1155 *ERC1155Session) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ERC1155.Contract.SetApprovalForAll(&_ERC1155.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ERC1155 *ERC1155TransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ERC1155.Contract.SetApprovalForAll(&_ERC1155.TransactOpts, operator, approved)
}

// ERC1155ApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the ERC1155 contract.
type ERC1155ApprovalForAllIterator struct {
	Event *ERC1155ApprovalForAll // Event containing the contract specifics and raw log

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
func (it *ERC1155ApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155ApprovalForAll)
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
		it.Event = new(ERC1155ApprovalForAll)
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
func (it *ERC1155ApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155ApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155ApprovalForAll represents a ApprovalForAll event raised by the ERC1155 contract.
type ERC1155ApprovalForAll struct {
	Account  common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_ERC1155 *ERC1155Filterer) FilterApprovalForAll(opts *bind.FilterOpts, account []common.Address, operator []common.Address) (*ERC1155ApprovalForAllIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ERC1155.contract.FilterLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ERC1155ApprovalForAllIterator{contract: _ERC1155.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_ERC1155 *ERC1155Filterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ERC1155ApprovalForAll, account []common.Address, operator []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ERC1155.contract.WatchLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155ApprovalForAll)
				if err := _ERC1155.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_ERC1155 *ERC1155Filterer) ParseApprovalForAll(log types.Log) (*ERC1155ApprovalForAll, error) {
	event := new(ERC1155ApprovalForAll)
	if err := _ERC1155.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1155TransferBatchIterator is returned from FilterTransferBatch and is used to iterate over the raw logs and unpacked data for TransferBatch events raised by the ERC1155 contract.
type ERC1155TransferBatchIterator struct {
	Event *ERC1155TransferBatch // Event containing the contract specifics and raw log

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
func (it *ERC1155TransferBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155TransferBatch)
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
		it.Event = new(ERC1155TransferBatch)
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
func (it *ERC1155TransferBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155TransferBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155TransferBatch represents a TransferBatch event raised by the ERC1155 contract.
type ERC1155TransferBatch struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Ids      []*big.Int
	Values   []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferBatch is a free log retrieval operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_ERC1155 *ERC1155Filterer) FilterTransferBatch(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*ERC1155TransferBatchIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC1155.contract.FilterLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC1155TransferBatchIterator{contract: _ERC1155.contract, event: "TransferBatch", logs: logs, sub: sub}, nil
}

// WatchTransferBatch is a free log subscription operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_ERC1155 *ERC1155Filterer) WatchTransferBatch(opts *bind.WatchOpts, sink chan<- *ERC1155TransferBatch, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC1155.contract.WatchLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155TransferBatch)
				if err := _ERC1155.contract.UnpackLog(event, "TransferBatch", log); err != nil {
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

// ParseTransferBatch is a log parse operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_ERC1155 *ERC1155Filterer) ParseTransferBatch(log types.Log) (*ERC1155TransferBatch, error) {
	event := new(ERC1155TransferBatch)
	if err := _ERC1155.contract.UnpackLog(event, "TransferBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1155TransferSingleIterator is returned from FilterTransferSingle and is used to iterate over the raw logs and unpacked data for TransferSingle events raised by the ERC1155 contract.
type ERC1155TransferSingleIterator struct {
	Event *ERC1155TransferSingle // Event containing the contract specifics and raw log

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
func (it *ERC1155TransferSingleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155TransferSingle)
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
		it.Event = new(ERC1155TransferSingle)
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
func (it *ERC1155TransferSingleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155TransferSingleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155TransferSingle represents a TransferSingle event raised by the ERC1155 contract.
type ERC1155TransferSingle struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Id       *big.Int
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferSingle is a free log retrieval operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_ERC1155 *ERC1155Filterer) FilterTransferSingle(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*ERC1155TransferSingleIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC1155.contract.FilterLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC1155TransferSingleIterator{contract: _ERC1155.contract, event: "TransferSingle", logs: logs, sub: sub}, nil
}

// WatchTransferSingle is a free log subscription operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_ERC1155 *ERC1155Filterer) WatchTransferSingle(opts *bind.WatchOpts, sink chan<- *ERC1155TransferSingle, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC1155.contract.WatchLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155TransferSingle)
				if err := _ERC1155.contract.UnpackLog(event, "TransferSingle", log); err != nil {
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

// ParseTransferSingle is a log parse operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_ERC1155 *ERC1155Filterer) ParseTransferSingle(log types.Log) (*ERC1155TransferSingle, error) {
	event := new(ERC1155TransferSingle)
	if err := _ERC1155.contract.UnpackLog(event, "TransferSingle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1155URIIterator is returned from FilterURI and is used to iterate over the raw logs and unpacked data for URI events raised by the ERC1155 contract.
type ERC1155URIIterator struct {
	Event *ERC1155URI // Event containing the contract specifics and raw log

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
func (it *ERC1155URIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155URI)
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
		it.Event = new(ERC1155URI)
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
func (it *ERC1155URIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155URIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155URI represents a URI event raised by the ERC1155 contract.
type ERC1155URI struct {
	Value string
	Id    *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterURI is a free log retrieval operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_ERC1155 *ERC1155Filterer) FilterURI(opts *bind.FilterOpts, id []*big.Int) (*ERC1155URIIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _ERC1155.contract.FilterLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return &ERC1155URIIterator{contract: _ERC1155.contract, event: "URI", logs: logs, sub: sub}, nil
}

// WatchURI is a free log subscription operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_ERC1155 *ERC1155Filterer) WatchURI(opts *bind.WatchOpts, sink chan<- *ERC1155URI, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _ERC1155.contract.WatchLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155URI)
				if err := _ERC1155.contract.UnpackLog(event, "URI", log); err != nil {
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

// ParseURI is a log parse operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_ERC1155 *ERC1155Filterer) ParseURI(log types.Log) (*ERC1155URI, error) {
	event := new(ERC1155URI)
	if err := _ERC1155.contract.UnpackLog(event, "URI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
