// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

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
	_ = abi.ConvertType
)

// AgentCollectionsManagerMetaData contains all meta data concerning the AgentCollectionsManager contract.
var AgentCollectionsManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"collectionId\",\"type\":\"uint256\"}],\"name\":\"CollectionDoesNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"CollectionNameAlreadyExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"name\":\"CollectionSymbolAlreadyExists\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"collectionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"collectionAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"name\":\"CollectionCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CREATOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SIGNER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newCollection\",\"type\":\"address\"}],\"name\":\"addCollection\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"collectionIdByName\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"collectionIdBySymbol\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"collections\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"isSigner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"moderatorOf\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextCollectionId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234602357600180556013336028565b50604051610bf290816100b68239f35b600080fd5b6001600160a01b0381166000908152600080516020610ca8833981519152602052604090205460ff1660af576001600160a01b03166000818152600080516020610ca883398151915260205260408120805460ff191660011790553391907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d8180a4600190565b5060009056fe608080604052600436101561001357600080fd5b60003560e01c90816301ffc9a7146108085750806308fa3c65146107e1578063103b970e146107a7578063248a9ca3146107815780632f2ff15d1461075057806336568abe1461070a5780634b2d81481461050b5780637df73e27146104a95780638aeda25a1461048057806391d1485414610433578063a174e77a146101b5578063a1ebf35d1461017a578063a217fddf1461015e578063d547741f14610126578063e77d6f7c146101085763fdbda0ec146100cf57600080fd5b34610103576020366003190112610103576004356000526002602052602060018060a01b0360406000205416604051908152f35b600080fd5b34610103576000366003190112610103576020600154604051908152f35b346101035760403660031901126101035761015c60043561014561093a565b9061015761015282610966565b610a54565b610b1a565b005b3461010357600036600319011261010357602060405160008152f35b346101035760003660031901126101035760206040517fe2f4eaae4a9751e85a3e4a7b9587827a877f29914755229b07a7b2da98285f708152f35b34610103576020366003190112610103576101ce610950565b3360009081527ffb6cfbf0a6e77dfd309859bf4bad2c7342bd2b734e42c01cbc85d3e6dd74f95e602052604090205460ff161561040c576040516306fdde0360e01b81526001600160a01b039190911690600081600481855afa9081156103e7576000916103f3575b506040516395d89b4160e01b815291600083600481845afa9283156103e7576000936103c4575b5060405182519060208181860193610277818387610917565b81016003815203019020546103a257604051908451916020818188019461029f818388610917565b810160048152030190205461037c579261037794926020610334817fe1344d16944f31a1bd59b121889e600e71fcc0716ae3fe38175b7ef2db1678b39561036f986001546000526002835260406000208660018060a01b0319825416179055610318600154958692604051809381928c51928391610917565b8101600381520301902055604051809381928a51928391610917565b81016004815203019020556001549485946103616040519485948552606060208601526060850190610a2f565b908382036040850152610a2f565b0390a2610992565b600155005b604051632456bbdf60e11b8152602060048201528061039e6024820188610a2f565b0390fd5b604051630a3942ab60e41b8152602060048201528061039e6024820186610a2f565b6103e091933d8091833e6103d8818361085b565b8101906109ce565b918361025e565b6040513d6000823e3d90fd5b610406913d8091833e6103d8818361085b565b82610237565b63e2517d3f60e01b60005233600452600080516020610b9d83398151915260245260446000fd5b346101035760403660031901126101035761044c61093a565b600435600052600060205260406000209060018060a01b0316600052602052602060ff604060002054166040519015158152f35b34610103576000366003190112610103576020604051600080516020610b9d8339815191528152f35b34610103576020366003190112610103576104c2610950565b6001600160a01b031660009081527f059f08e7d7ba1c82eddc57afae67f80df851baf38a099607a779825038c3ce5b602090815260409182902054915160ff9092161515825290f35b3461010357602036600319011261010357610524610950565b6001805460009290916001600160a01b03909116905b828110610687575061056461054e846109b7565b9361055c604051958661085b565b8085526109b7565b602084019290601f1901368437600160005b8282106105cb5784866040519182916020830190602084525180915260408301919060005b8181106105a9575050500390f35b82516001600160a01b031684528594506020938401939092019160010161059b565b81600052600260205260018060a01b0360406000205416604051637d379c9b60e11b8152856004820152602081602481855afa9081156103e757600091610659575b5061061e575b509060010190610576565b8651821015610643578161063d91602060019460051b8a010152610992565b90610613565b634e487b7160e01b600052603260045260246000fd5b61067a915060203d8111610680575b610672818361085b565b81019061097a565b8861060d565b503d610668565b8060005260026020526024602060018060a01b036040600020541660405192838092637d379c9b60e11b82528760048301525afa9081156103e7576000916106ec575b506106d8575b60010161053a565b926106e4600191610992565b9390506106d0565b610704915060203d811161068057610672818361085b565b856106ca565b346101035760403660031901126101035761072361093a565b336001600160a01b0382160361073f5761015c90600435610b1a565b63334bd91960e11b60005260046000fd5b346101035760403660031901126101035761015c60043561076f61093a565b9061077c61015282610966565b610a8f565b3461010357602036600319011261010357602061079f600435610966565b604051908152f35b346101035760206107ce816107bb366108af565b8160405193828580945193849201610917565b8101600481520301902054604051908152f35b346101035760206107f5816107bb366108af565b8101600381520301902054604051908152f35b34610103576020366003190112610103576004359063ffffffff60e01b821680920361010357602091637965db0b60e01b811490811561084a575b5015158152f35b6301ffc9a760e01b14905083610843565b601f909101601f19168101906001600160401b0382119082101761087e57604052565b634e487b7160e01b600052604160045260246000fd5b6001600160401b03811161087e57601f01601f191660200190565b6020600319820112610103576004356001600160401b0381116101035781602382011215610103578060040135906108e682610894565b926108f4604051948561085b565b828452602483830101116101035781600092602460209301838601378301015290565b60005b83811061092a5750506000910152565b818101518382015260200161091a565b602435906001600160a01b038216820361010357565b600435906001600160a01b038216820361010357565b600052600060205260016040600020015490565b90816020910312610103575180151581036101035790565b60001981146109a15760010190565b634e487b7160e01b600052601160045260246000fd5b6001600160401b03811161087e5760051b60200190565b602081830312610103578051906001600160401b038211610103570181601f82011215610103578051610a0081610894565b92610a0e604051948561085b565b8184526020828401011161010357610a2c9160208085019101610917565b90565b90602091610a4881518092818552858086019101610917565b601f01601f1916010190565b60008181526020818152604080832033845290915290205460ff1615610a775750565b63e2517d3f60e01b6000523360045260245260446000fd5b6000818152602081815260408083206001600160a01b038616845290915290205460ff16610b13576000818152602081815260408083206001600160a01b0395909516808452949091528120805460ff19166001179055339291907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d9080a4600190565b5050600090565b6000818152602081815260408083206001600160a01b038616845290915290205460ff1615610b13576000818152602081815260408083206001600160a01b0395909516808452949091528120805460ff19169055339291907ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9080a460019056fe828634d95e775031b9ff576b159a8509d3053581a8c9c4d7d86899e0afcd882fa26469706673582212200cc687c403e76d7e1d19b422da7807752a229c7b9eac9b54d55c96050892e70764736f6c634300081c0033ad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5",
}

// AgentCollectionsManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use AgentCollectionsManagerMetaData.ABI instead.
var AgentCollectionsManagerABI = AgentCollectionsManagerMetaData.ABI

// AgentCollectionsManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AgentCollectionsManagerMetaData.Bin instead.
var AgentCollectionsManagerBin = AgentCollectionsManagerMetaData.Bin

// DeployAgentCollectionsManager deploys a new Ethereum contract, binding an instance of AgentCollectionsManager to it.
func DeployAgentCollectionsManager(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AgentCollectionsManager, error) {
	parsed, err := AgentCollectionsManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AgentCollectionsManagerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AgentCollectionsManager{AgentCollectionsManagerCaller: AgentCollectionsManagerCaller{contract: contract}, AgentCollectionsManagerTransactor: AgentCollectionsManagerTransactor{contract: contract}, AgentCollectionsManagerFilterer: AgentCollectionsManagerFilterer{contract: contract}}, nil
}

// AgentCollectionsManager is an auto generated Go binding around an Ethereum contract.
type AgentCollectionsManager struct {
	AgentCollectionsManagerCaller     // Read-only binding to the contract
	AgentCollectionsManagerTransactor // Write-only binding to the contract
	AgentCollectionsManagerFilterer   // Log filterer for contract events
}

// AgentCollectionsManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type AgentCollectionsManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgentCollectionsManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AgentCollectionsManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgentCollectionsManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AgentCollectionsManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgentCollectionsManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AgentCollectionsManagerSession struct {
	Contract     *AgentCollectionsManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// AgentCollectionsManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AgentCollectionsManagerCallerSession struct {
	Contract *AgentCollectionsManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// AgentCollectionsManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AgentCollectionsManagerTransactorSession struct {
	Contract     *AgentCollectionsManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// AgentCollectionsManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type AgentCollectionsManagerRaw struct {
	Contract *AgentCollectionsManager // Generic contract binding to access the raw methods on
}

// AgentCollectionsManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AgentCollectionsManagerCallerRaw struct {
	Contract *AgentCollectionsManagerCaller // Generic read-only contract binding to access the raw methods on
}

// AgentCollectionsManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AgentCollectionsManagerTransactorRaw struct {
	Contract *AgentCollectionsManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAgentCollectionsManager creates a new instance of AgentCollectionsManager, bound to a specific deployed contract.
func NewAgentCollectionsManager(address common.Address, backend bind.ContractBackend) (*AgentCollectionsManager, error) {
	contract, err := bindAgentCollectionsManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AgentCollectionsManager{AgentCollectionsManagerCaller: AgentCollectionsManagerCaller{contract: contract}, AgentCollectionsManagerTransactor: AgentCollectionsManagerTransactor{contract: contract}, AgentCollectionsManagerFilterer: AgentCollectionsManagerFilterer{contract: contract}}, nil
}

// NewAgentCollectionsManagerCaller creates a new read-only instance of AgentCollectionsManager, bound to a specific deployed contract.
func NewAgentCollectionsManagerCaller(address common.Address, caller bind.ContractCaller) (*AgentCollectionsManagerCaller, error) {
	contract, err := bindAgentCollectionsManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AgentCollectionsManagerCaller{contract: contract}, nil
}

// NewAgentCollectionsManagerTransactor creates a new write-only instance of AgentCollectionsManager, bound to a specific deployed contract.
func NewAgentCollectionsManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*AgentCollectionsManagerTransactor, error) {
	contract, err := bindAgentCollectionsManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AgentCollectionsManagerTransactor{contract: contract}, nil
}

// NewAgentCollectionsManagerFilterer creates a new log filterer instance of AgentCollectionsManager, bound to a specific deployed contract.
func NewAgentCollectionsManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*AgentCollectionsManagerFilterer, error) {
	contract, err := bindAgentCollectionsManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AgentCollectionsManagerFilterer{contract: contract}, nil
}

// bindAgentCollectionsManager binds a generic wrapper to an already deployed contract.
func bindAgentCollectionsManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AgentCollectionsManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AgentCollectionsManager *AgentCollectionsManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AgentCollectionsManager.Contract.AgentCollectionsManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AgentCollectionsManager *AgentCollectionsManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AgentCollectionsManager.Contract.AgentCollectionsManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AgentCollectionsManager *AgentCollectionsManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AgentCollectionsManager.Contract.AgentCollectionsManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AgentCollectionsManager *AgentCollectionsManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AgentCollectionsManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AgentCollectionsManager *AgentCollectionsManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AgentCollectionsManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AgentCollectionsManager *AgentCollectionsManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AgentCollectionsManager.Contract.contract.Transact(opts, method, params...)
}

// CREATORROLE is a free data retrieval call binding the contract method 0x8aeda25a.
//
// Solidity: function CREATOR_ROLE() view returns(bytes32)
func (_AgentCollectionsManager *AgentCollectionsManagerCaller) CREATORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AgentCollectionsManager.contract.Call(opts, &out, "CREATOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CREATORROLE is a free data retrieval call binding the contract method 0x8aeda25a.
//
// Solidity: function CREATOR_ROLE() view returns(bytes32)
func (_AgentCollectionsManager *AgentCollectionsManagerSession) CREATORROLE() ([32]byte, error) {
	return _AgentCollectionsManager.Contract.CREATORROLE(&_AgentCollectionsManager.CallOpts)
}

// CREATORROLE is a free data retrieval call binding the contract method 0x8aeda25a.
//
// Solidity: function CREATOR_ROLE() view returns(bytes32)
func (_AgentCollectionsManager *AgentCollectionsManagerCallerSession) CREATORROLE() ([32]byte, error) {
	return _AgentCollectionsManager.Contract.CREATORROLE(&_AgentCollectionsManager.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AgentCollectionsManager *AgentCollectionsManagerCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AgentCollectionsManager.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AgentCollectionsManager *AgentCollectionsManagerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _AgentCollectionsManager.Contract.DEFAULTADMINROLE(&_AgentCollectionsManager.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AgentCollectionsManager *AgentCollectionsManagerCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _AgentCollectionsManager.Contract.DEFAULTADMINROLE(&_AgentCollectionsManager.CallOpts)
}

// SIGNERROLE is a free data retrieval call binding the contract method 0xa1ebf35d.
//
// Solidity: function SIGNER_ROLE() view returns(bytes32)
func (_AgentCollectionsManager *AgentCollectionsManagerCaller) SIGNERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AgentCollectionsManager.contract.Call(opts, &out, "SIGNER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SIGNERROLE is a free data retrieval call binding the contract method 0xa1ebf35d.
//
// Solidity: function SIGNER_ROLE() view returns(bytes32)
func (_AgentCollectionsManager *AgentCollectionsManagerSession) SIGNERROLE() ([32]byte, error) {
	return _AgentCollectionsManager.Contract.SIGNERROLE(&_AgentCollectionsManager.CallOpts)
}

// SIGNERROLE is a free data retrieval call binding the contract method 0xa1ebf35d.
//
// Solidity: function SIGNER_ROLE() view returns(bytes32)
func (_AgentCollectionsManager *AgentCollectionsManagerCallerSession) SIGNERROLE() ([32]byte, error) {
	return _AgentCollectionsManager.Contract.SIGNERROLE(&_AgentCollectionsManager.CallOpts)
}

// CollectionIdByName is a free data retrieval call binding the contract method 0x08fa3c65.
//
// Solidity: function collectionIdByName(string ) view returns(uint256)
func (_AgentCollectionsManager *AgentCollectionsManagerCaller) CollectionIdByName(opts *bind.CallOpts, arg0 string) (*big.Int, error) {
	var out []interface{}
	err := _AgentCollectionsManager.contract.Call(opts, &out, "collectionIdByName", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CollectionIdByName is a free data retrieval call binding the contract method 0x08fa3c65.
//
// Solidity: function collectionIdByName(string ) view returns(uint256)
func (_AgentCollectionsManager *AgentCollectionsManagerSession) CollectionIdByName(arg0 string) (*big.Int, error) {
	return _AgentCollectionsManager.Contract.CollectionIdByName(&_AgentCollectionsManager.CallOpts, arg0)
}

// CollectionIdByName is a free data retrieval call binding the contract method 0x08fa3c65.
//
// Solidity: function collectionIdByName(string ) view returns(uint256)
func (_AgentCollectionsManager *AgentCollectionsManagerCallerSession) CollectionIdByName(arg0 string) (*big.Int, error) {
	return _AgentCollectionsManager.Contract.CollectionIdByName(&_AgentCollectionsManager.CallOpts, arg0)
}

// CollectionIdBySymbol is a free data retrieval call binding the contract method 0x103b970e.
//
// Solidity: function collectionIdBySymbol(string ) view returns(uint256)
func (_AgentCollectionsManager *AgentCollectionsManagerCaller) CollectionIdBySymbol(opts *bind.CallOpts, arg0 string) (*big.Int, error) {
	var out []interface{}
	err := _AgentCollectionsManager.contract.Call(opts, &out, "collectionIdBySymbol", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CollectionIdBySymbol is a free data retrieval call binding the contract method 0x103b970e.
//
// Solidity: function collectionIdBySymbol(string ) view returns(uint256)
func (_AgentCollectionsManager *AgentCollectionsManagerSession) CollectionIdBySymbol(arg0 string) (*big.Int, error) {
	return _AgentCollectionsManager.Contract.CollectionIdBySymbol(&_AgentCollectionsManager.CallOpts, arg0)
}

// CollectionIdBySymbol is a free data retrieval call binding the contract method 0x103b970e.
//
// Solidity: function collectionIdBySymbol(string ) view returns(uint256)
func (_AgentCollectionsManager *AgentCollectionsManagerCallerSession) CollectionIdBySymbol(arg0 string) (*big.Int, error) {
	return _AgentCollectionsManager.Contract.CollectionIdBySymbol(&_AgentCollectionsManager.CallOpts, arg0)
}

// Collections is a free data retrieval call binding the contract method 0xfdbda0ec.
//
// Solidity: function collections(uint256 ) view returns(address)
func (_AgentCollectionsManager *AgentCollectionsManagerCaller) Collections(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _AgentCollectionsManager.contract.Call(opts, &out, "collections", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Collections is a free data retrieval call binding the contract method 0xfdbda0ec.
//
// Solidity: function collections(uint256 ) view returns(address)
func (_AgentCollectionsManager *AgentCollectionsManagerSession) Collections(arg0 *big.Int) (common.Address, error) {
	return _AgentCollectionsManager.Contract.Collections(&_AgentCollectionsManager.CallOpts, arg0)
}

// Collections is a free data retrieval call binding the contract method 0xfdbda0ec.
//
// Solidity: function collections(uint256 ) view returns(address)
func (_AgentCollectionsManager *AgentCollectionsManagerCallerSession) Collections(arg0 *big.Int) (common.Address, error) {
	return _AgentCollectionsManager.Contract.Collections(&_AgentCollectionsManager.CallOpts, arg0)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AgentCollectionsManager *AgentCollectionsManagerCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _AgentCollectionsManager.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AgentCollectionsManager *AgentCollectionsManagerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _AgentCollectionsManager.Contract.GetRoleAdmin(&_AgentCollectionsManager.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AgentCollectionsManager *AgentCollectionsManagerCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _AgentCollectionsManager.Contract.GetRoleAdmin(&_AgentCollectionsManager.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AgentCollectionsManager *AgentCollectionsManagerCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _AgentCollectionsManager.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AgentCollectionsManager *AgentCollectionsManagerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _AgentCollectionsManager.Contract.HasRole(&_AgentCollectionsManager.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AgentCollectionsManager *AgentCollectionsManagerCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _AgentCollectionsManager.Contract.HasRole(&_AgentCollectionsManager.CallOpts, role, account)
}

// IsSigner is a free data retrieval call binding the contract method 0x7df73e27.
//
// Solidity: function isSigner(address signer) view returns(bool)
func (_AgentCollectionsManager *AgentCollectionsManagerCaller) IsSigner(opts *bind.CallOpts, signer common.Address) (bool, error) {
	var out []interface{}
	err := _AgentCollectionsManager.contract.Call(opts, &out, "isSigner", signer)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSigner is a free data retrieval call binding the contract method 0x7df73e27.
//
// Solidity: function isSigner(address signer) view returns(bool)
func (_AgentCollectionsManager *AgentCollectionsManagerSession) IsSigner(signer common.Address) (bool, error) {
	return _AgentCollectionsManager.Contract.IsSigner(&_AgentCollectionsManager.CallOpts, signer)
}

// IsSigner is a free data retrieval call binding the contract method 0x7df73e27.
//
// Solidity: function isSigner(address signer) view returns(bool)
func (_AgentCollectionsManager *AgentCollectionsManagerCallerSession) IsSigner(signer common.Address) (bool, error) {
	return _AgentCollectionsManager.Contract.IsSigner(&_AgentCollectionsManager.CallOpts, signer)
}

// ModeratorOf is a free data retrieval call binding the contract method 0x4b2d8148.
//
// Solidity: function moderatorOf(address account) view returns(address[])
func (_AgentCollectionsManager *AgentCollectionsManagerCaller) ModeratorOf(opts *bind.CallOpts, account common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _AgentCollectionsManager.contract.Call(opts, &out, "moderatorOf", account)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ModeratorOf is a free data retrieval call binding the contract method 0x4b2d8148.
//
// Solidity: function moderatorOf(address account) view returns(address[])
func (_AgentCollectionsManager *AgentCollectionsManagerSession) ModeratorOf(account common.Address) ([]common.Address, error) {
	return _AgentCollectionsManager.Contract.ModeratorOf(&_AgentCollectionsManager.CallOpts, account)
}

// ModeratorOf is a free data retrieval call binding the contract method 0x4b2d8148.
//
// Solidity: function moderatorOf(address account) view returns(address[])
func (_AgentCollectionsManager *AgentCollectionsManagerCallerSession) ModeratorOf(account common.Address) ([]common.Address, error) {
	return _AgentCollectionsManager.Contract.ModeratorOf(&_AgentCollectionsManager.CallOpts, account)
}

// NextCollectionId is a free data retrieval call binding the contract method 0xe77d6f7c.
//
// Solidity: function nextCollectionId() view returns(uint256)
func (_AgentCollectionsManager *AgentCollectionsManagerCaller) NextCollectionId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AgentCollectionsManager.contract.Call(opts, &out, "nextCollectionId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextCollectionId is a free data retrieval call binding the contract method 0xe77d6f7c.
//
// Solidity: function nextCollectionId() view returns(uint256)
func (_AgentCollectionsManager *AgentCollectionsManagerSession) NextCollectionId() (*big.Int, error) {
	return _AgentCollectionsManager.Contract.NextCollectionId(&_AgentCollectionsManager.CallOpts)
}

// NextCollectionId is a free data retrieval call binding the contract method 0xe77d6f7c.
//
// Solidity: function nextCollectionId() view returns(uint256)
func (_AgentCollectionsManager *AgentCollectionsManagerCallerSession) NextCollectionId() (*big.Int, error) {
	return _AgentCollectionsManager.Contract.NextCollectionId(&_AgentCollectionsManager.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AgentCollectionsManager *AgentCollectionsManagerCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _AgentCollectionsManager.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AgentCollectionsManager *AgentCollectionsManagerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _AgentCollectionsManager.Contract.SupportsInterface(&_AgentCollectionsManager.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AgentCollectionsManager *AgentCollectionsManagerCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _AgentCollectionsManager.Contract.SupportsInterface(&_AgentCollectionsManager.CallOpts, interfaceId)
}

// AddCollection is a paid mutator transaction binding the contract method 0xa174e77a.
//
// Solidity: function addCollection(address newCollection) returns()
func (_AgentCollectionsManager *AgentCollectionsManagerTransactor) AddCollection(opts *bind.TransactOpts, newCollection common.Address) (*types.Transaction, error) {
	return _AgentCollectionsManager.contract.Transact(opts, "addCollection", newCollection)
}

// AddCollection is a paid mutator transaction binding the contract method 0xa174e77a.
//
// Solidity: function addCollection(address newCollection) returns()
func (_AgentCollectionsManager *AgentCollectionsManagerSession) AddCollection(newCollection common.Address) (*types.Transaction, error) {
	return _AgentCollectionsManager.Contract.AddCollection(&_AgentCollectionsManager.TransactOpts, newCollection)
}

// AddCollection is a paid mutator transaction binding the contract method 0xa174e77a.
//
// Solidity: function addCollection(address newCollection) returns()
func (_AgentCollectionsManager *AgentCollectionsManagerTransactorSession) AddCollection(newCollection common.Address) (*types.Transaction, error) {
	return _AgentCollectionsManager.Contract.AddCollection(&_AgentCollectionsManager.TransactOpts, newCollection)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AgentCollectionsManager *AgentCollectionsManagerTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AgentCollectionsManager.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AgentCollectionsManager *AgentCollectionsManagerSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AgentCollectionsManager.Contract.GrantRole(&_AgentCollectionsManager.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AgentCollectionsManager *AgentCollectionsManagerTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AgentCollectionsManager.Contract.GrantRole(&_AgentCollectionsManager.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_AgentCollectionsManager *AgentCollectionsManagerTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _AgentCollectionsManager.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_AgentCollectionsManager *AgentCollectionsManagerSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _AgentCollectionsManager.Contract.RenounceRole(&_AgentCollectionsManager.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_AgentCollectionsManager *AgentCollectionsManagerTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _AgentCollectionsManager.Contract.RenounceRole(&_AgentCollectionsManager.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AgentCollectionsManager *AgentCollectionsManagerTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AgentCollectionsManager.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AgentCollectionsManager *AgentCollectionsManagerSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AgentCollectionsManager.Contract.RevokeRole(&_AgentCollectionsManager.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AgentCollectionsManager *AgentCollectionsManagerTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AgentCollectionsManager.Contract.RevokeRole(&_AgentCollectionsManager.TransactOpts, role, account)
}

// AgentCollectionsManagerCollectionCreatedIterator is returned from FilterCollectionCreated and is used to iterate over the raw logs and unpacked data for CollectionCreated events raised by the AgentCollectionsManager contract.
type AgentCollectionsManagerCollectionCreatedIterator struct {
	Event *AgentCollectionsManagerCollectionCreated // Event containing the contract specifics and raw log

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
func (it *AgentCollectionsManagerCollectionCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentCollectionsManagerCollectionCreated)
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
		it.Event = new(AgentCollectionsManagerCollectionCreated)
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
func (it *AgentCollectionsManagerCollectionCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentCollectionsManagerCollectionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentCollectionsManagerCollectionCreated represents a CollectionCreated event raised by the AgentCollectionsManager contract.
type AgentCollectionsManagerCollectionCreated struct {
	CollectionId      *big.Int
	CollectionAddress common.Address
	Name              string
	Symbol            string
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterCollectionCreated is a free log retrieval operation binding the contract event 0xe1344d16944f31a1bd59b121889e600e71fcc0716ae3fe38175b7ef2db1678b3.
//
// Solidity: event CollectionCreated(uint256 indexed collectionId, address collectionAddress, string name, string symbol)
func (_AgentCollectionsManager *AgentCollectionsManagerFilterer) FilterCollectionCreated(opts *bind.FilterOpts, collectionId []*big.Int) (*AgentCollectionsManagerCollectionCreatedIterator, error) {

	var collectionIdRule []interface{}
	for _, collectionIdItem := range collectionId {
		collectionIdRule = append(collectionIdRule, collectionIdItem)
	}

	logs, sub, err := _AgentCollectionsManager.contract.FilterLogs(opts, "CollectionCreated", collectionIdRule)
	if err != nil {
		return nil, err
	}
	return &AgentCollectionsManagerCollectionCreatedIterator{contract: _AgentCollectionsManager.contract, event: "CollectionCreated", logs: logs, sub: sub}, nil
}

// WatchCollectionCreated is a free log subscription operation binding the contract event 0xe1344d16944f31a1bd59b121889e600e71fcc0716ae3fe38175b7ef2db1678b3.
//
// Solidity: event CollectionCreated(uint256 indexed collectionId, address collectionAddress, string name, string symbol)
func (_AgentCollectionsManager *AgentCollectionsManagerFilterer) WatchCollectionCreated(opts *bind.WatchOpts, sink chan<- *AgentCollectionsManagerCollectionCreated, collectionId []*big.Int) (event.Subscription, error) {

	var collectionIdRule []interface{}
	for _, collectionIdItem := range collectionId {
		collectionIdRule = append(collectionIdRule, collectionIdItem)
	}

	logs, sub, err := _AgentCollectionsManager.contract.WatchLogs(opts, "CollectionCreated", collectionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentCollectionsManagerCollectionCreated)
				if err := _AgentCollectionsManager.contract.UnpackLog(event, "CollectionCreated", log); err != nil {
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

// ParseCollectionCreated is a log parse operation binding the contract event 0xe1344d16944f31a1bd59b121889e600e71fcc0716ae3fe38175b7ef2db1678b3.
//
// Solidity: event CollectionCreated(uint256 indexed collectionId, address collectionAddress, string name, string symbol)
func (_AgentCollectionsManager *AgentCollectionsManagerFilterer) ParseCollectionCreated(log types.Log) (*AgentCollectionsManagerCollectionCreated, error) {
	event := new(AgentCollectionsManagerCollectionCreated)
	if err := _AgentCollectionsManager.contract.UnpackLog(event, "CollectionCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentCollectionsManagerRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the AgentCollectionsManager contract.
type AgentCollectionsManagerRoleAdminChangedIterator struct {
	Event *AgentCollectionsManagerRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *AgentCollectionsManagerRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentCollectionsManagerRoleAdminChanged)
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
		it.Event = new(AgentCollectionsManagerRoleAdminChanged)
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
func (it *AgentCollectionsManagerRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentCollectionsManagerRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentCollectionsManagerRoleAdminChanged represents a RoleAdminChanged event raised by the AgentCollectionsManager contract.
type AgentCollectionsManagerRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_AgentCollectionsManager *AgentCollectionsManagerFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*AgentCollectionsManagerRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _AgentCollectionsManager.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &AgentCollectionsManagerRoleAdminChangedIterator{contract: _AgentCollectionsManager.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_AgentCollectionsManager *AgentCollectionsManagerFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *AgentCollectionsManagerRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _AgentCollectionsManager.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentCollectionsManagerRoleAdminChanged)
				if err := _AgentCollectionsManager.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_AgentCollectionsManager *AgentCollectionsManagerFilterer) ParseRoleAdminChanged(log types.Log) (*AgentCollectionsManagerRoleAdminChanged, error) {
	event := new(AgentCollectionsManagerRoleAdminChanged)
	if err := _AgentCollectionsManager.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentCollectionsManagerRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the AgentCollectionsManager contract.
type AgentCollectionsManagerRoleGrantedIterator struct {
	Event *AgentCollectionsManagerRoleGranted // Event containing the contract specifics and raw log

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
func (it *AgentCollectionsManagerRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentCollectionsManagerRoleGranted)
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
		it.Event = new(AgentCollectionsManagerRoleGranted)
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
func (it *AgentCollectionsManagerRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentCollectionsManagerRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentCollectionsManagerRoleGranted represents a RoleGranted event raised by the AgentCollectionsManager contract.
type AgentCollectionsManagerRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AgentCollectionsManager *AgentCollectionsManagerFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AgentCollectionsManagerRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AgentCollectionsManager.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AgentCollectionsManagerRoleGrantedIterator{contract: _AgentCollectionsManager.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AgentCollectionsManager *AgentCollectionsManagerFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *AgentCollectionsManagerRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AgentCollectionsManager.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentCollectionsManagerRoleGranted)
				if err := _AgentCollectionsManager.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AgentCollectionsManager *AgentCollectionsManagerFilterer) ParseRoleGranted(log types.Log) (*AgentCollectionsManagerRoleGranted, error) {
	event := new(AgentCollectionsManagerRoleGranted)
	if err := _AgentCollectionsManager.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentCollectionsManagerRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the AgentCollectionsManager contract.
type AgentCollectionsManagerRoleRevokedIterator struct {
	Event *AgentCollectionsManagerRoleRevoked // Event containing the contract specifics and raw log

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
func (it *AgentCollectionsManagerRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentCollectionsManagerRoleRevoked)
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
		it.Event = new(AgentCollectionsManagerRoleRevoked)
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
func (it *AgentCollectionsManagerRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentCollectionsManagerRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentCollectionsManagerRoleRevoked represents a RoleRevoked event raised by the AgentCollectionsManager contract.
type AgentCollectionsManagerRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AgentCollectionsManager *AgentCollectionsManagerFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AgentCollectionsManagerRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AgentCollectionsManager.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AgentCollectionsManagerRoleRevokedIterator{contract: _AgentCollectionsManager.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AgentCollectionsManager *AgentCollectionsManagerFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *AgentCollectionsManagerRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AgentCollectionsManager.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentCollectionsManagerRoleRevoked)
				if err := _AgentCollectionsManager.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AgentCollectionsManager *AgentCollectionsManagerFilterer) ParseRoleRevoked(log types.Log) (*AgentCollectionsManagerRoleRevoked, error) {
	event := new(AgentCollectionsManagerRoleRevoked)
	if err := _AgentCollectionsManager.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
