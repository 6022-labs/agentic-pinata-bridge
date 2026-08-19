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

// AddOrUpdateAgentImageProposal is an auto generated low-level Go binding around an user-defined struct.
type AddOrUpdateAgentImageProposal struct {
	Id      *big.Int
	TokenId *big.Int
	Image   KeyValue
}

// AgentAddress is an auto generated low-level Go binding around an user-defined struct.
type AgentAddress struct {
	AddressType string
	Value       string
}

// AgentInformation is an auto generated low-level Go binding around an user-defined struct.
type AgentInformation struct {
	Name       string
	Creator    common.Address
	Images     []KeyValue
	Attributes []KeyValue
	CloneOf    NullableUint256
	Addresses  []AgentAddress
}

// KeyValue is an auto generated low-level Go binding around an user-defined struct.
type KeyValue struct {
	Key   string
	Value string
}

// MintProposal is an auto generated low-level Go binding around an user-defined struct.
type MintProposal struct {
	Id         *big.Int
	Name       string
	Owner      common.Address
	Creator    common.Address
	Images     []KeyValue
	Attributes []KeyValue
	CloneOf    NullableUint256
	Addresses  []AgentAddress
}

// NullableUint256 is an auto generated low-level Go binding around an user-defined struct.
type NullableUint256 struct {
	IsSet bool
	Value *big.Int
}

// AgentCollectionV1MetaData contains all meta data concerning the AgentCollectionV1 contract.
var AgentCollectionV1MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"addressType\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"AgentAddressNotFound\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"DuplicateKey\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721IncorrectOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721InsufficientApproval\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC721InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721NonexistentToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"addressType\",\"type\":\"string\"}],\"name\":\"EmptyAddressValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"addressType\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"InvalidAddressType\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cloneOf\",\"type\":\"uint256\"}],\"name\":\"InvalidCloneOf\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"InvalidCollectionName\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"InvalidEvmAddressValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"InvalidName\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MissingDefaultImage\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MissingEvmAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"MissingRequiredKey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoModerators\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"NotModerator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"ProposalNotFound\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"StringsInsufficientHexLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"addressType\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"UsedAgentAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"UsedName\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"parameter\",\"type\":\"string\"}],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"addressType\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"AgentAddressAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"addressType\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"AgentAddressRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"oldValue\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newValue\",\"type\":\"string\"}],\"name\":\"AgentAttributeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"AgentImageProposalCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"AgentImageProposalRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"AgentImageRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"oldValue\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newValue\",\"type\":\"string\"}],\"name\":\"AgentImageUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldCollectionsDescriptor\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newCollectionsDescriptor\",\"type\":\"address\"}],\"name\":\"CollectionsDescriptorUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"oldDescription\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newDescription\",\"type\":\"string\"}],\"name\":\"DescriptionUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"MintProposalCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"MintProposalRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Minted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MODERATOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"addressType\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"addAgentAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"addOrUpdateAgentAttribute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"internalType\":\"structKeyValue\",\"name\":\"image\",\"type\":\"tuple\"}],\"name\":\"addOrUpdateAgentImage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"addOrUpdateAgentImageFromProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalIndex\",\"type\":\"uint256\"}],\"name\":\"addOrUpdateAgentImageProposal\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"internalType\":\"structKeyValue\",\"name\":\"image\",\"type\":\"tuple\"}],\"internalType\":\"structAddOrUpdateAgentImageProposal\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"addOrUpdateAgentImageProposalsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"addressOfByIndex\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"addressType\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"internalType\":\"structAgentAddress\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"addressType\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"addressToTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"addressesCountOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"attributeOfByIndex\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"internalType\":\"structKeyValue\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"attributeOfByKey\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"attributesCountOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"collectionsDescriptor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"internalType\":\"structKeyValue\",\"name\":\"image\",\"type\":\"tuple\"}],\"name\":\"createAddOrUpdateAgentImageProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"agentName\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"addressType\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"internalType\":\"structAgentAddress[]\",\"name\":\"addresses\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"internalType\":\"structKeyValue[]\",\"name\":\"images\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"internalType\":\"structKeyValue[]\",\"name\":\"attributes\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isSet\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structNullableUint256\",\"name\":\"cloneOf\",\"type\":\"tuple\"}],\"name\":\"createMintProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"creatorOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"description\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"imageOfByIndex\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"internalType\":\"structKeyValue\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"imagesCountOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"informationOf\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"internalType\":\"structKeyValue[]\",\"name\":\"images\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"internalType\":\"structKeyValue[]\",\"name\":\"attributes\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isSet\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structNullableUint256\",\"name\":\"cloneOf\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"addressType\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"internalType\":\"structAgentAddress[]\",\"name\":\"addresses\",\"type\":\"tuple[]\"}],\"internalType\":\"structAgentInformation\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"initialDescription\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"moderator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"initialCollectionsDescriptor\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isModerator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"agentName\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"addressType\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"internalType\":\"structAgentAddress[]\",\"name\":\"addresses\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"internalType\":\"structKeyValue[]\",\"name\":\"images\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"internalType\":\"structKeyValue[]\",\"name\":\"attributes\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isSet\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structNullableUint256\",\"name\":\"cloneOf\",\"type\":\"tuple\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"mintFromProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalIndex\",\"type\":\"uint256\"}],\"name\":\"mintProposal\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"internalType\":\"structKeyValue[]\",\"name\":\"images\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"internalType\":\"structKeyValue[]\",\"name\":\"attributes\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isSet\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structNullableUint256\",\"name\":\"cloneOf\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"addressType\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"internalType\":\"structAgentAddress[]\",\"name\":\"addresses\",\"type\":\"tuple[]\"}],\"internalType\":\"structMintProposal\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintProposalsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"moderatorCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"nameOf\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"agentName\",\"type\":\"string\"}],\"name\":\"nameToTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"refuseAddOrUpdateAgentImageProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"refuseMintProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"addressType\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"removeAgentAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"removeAgentImage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newCollectionsDescriptor\",\"type\":\"address\"}],\"name\":\"updateCollectionsDescriptor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newDescription\",\"type\":\"string\"}],\"name\":\"updateDescription\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// AgentCollectionV1ABI is the input ABI used to generate the binding from.
// Deprecated: Use AgentCollectionV1MetaData.ABI instead.
var AgentCollectionV1ABI = AgentCollectionV1MetaData.ABI

// AgentCollectionV1 is an auto generated Go binding around an Ethereum contract.
type AgentCollectionV1 struct {
	AgentCollectionV1Caller     // Read-only binding to the contract
	AgentCollectionV1Transactor // Write-only binding to the contract
	AgentCollectionV1Filterer   // Log filterer for contract events
}

// AgentCollectionV1Caller is an auto generated read-only Go binding around an Ethereum contract.
type AgentCollectionV1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgentCollectionV1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type AgentCollectionV1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgentCollectionV1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AgentCollectionV1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgentCollectionV1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AgentCollectionV1Session struct {
	Contract     *AgentCollectionV1 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// AgentCollectionV1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AgentCollectionV1CallerSession struct {
	Contract *AgentCollectionV1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// AgentCollectionV1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AgentCollectionV1TransactorSession struct {
	Contract     *AgentCollectionV1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// AgentCollectionV1Raw is an auto generated low-level Go binding around an Ethereum contract.
type AgentCollectionV1Raw struct {
	Contract *AgentCollectionV1 // Generic contract binding to access the raw methods on
}

// AgentCollectionV1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AgentCollectionV1CallerRaw struct {
	Contract *AgentCollectionV1Caller // Generic read-only contract binding to access the raw methods on
}

// AgentCollectionV1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AgentCollectionV1TransactorRaw struct {
	Contract *AgentCollectionV1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewAgentCollectionV1 creates a new instance of AgentCollectionV1, bound to a specific deployed contract.
func NewAgentCollectionV1(address common.Address, backend bind.ContractBackend) (*AgentCollectionV1, error) {
	contract, err := bindAgentCollectionV1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AgentCollectionV1{AgentCollectionV1Caller: AgentCollectionV1Caller{contract: contract}, AgentCollectionV1Transactor: AgentCollectionV1Transactor{contract: contract}, AgentCollectionV1Filterer: AgentCollectionV1Filterer{contract: contract}}, nil
}

// NewAgentCollectionV1Caller creates a new read-only instance of AgentCollectionV1, bound to a specific deployed contract.
func NewAgentCollectionV1Caller(address common.Address, caller bind.ContractCaller) (*AgentCollectionV1Caller, error) {
	contract, err := bindAgentCollectionV1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AgentCollectionV1Caller{contract: contract}, nil
}

// NewAgentCollectionV1Transactor creates a new write-only instance of AgentCollectionV1, bound to a specific deployed contract.
func NewAgentCollectionV1Transactor(address common.Address, transactor bind.ContractTransactor) (*AgentCollectionV1Transactor, error) {
	contract, err := bindAgentCollectionV1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AgentCollectionV1Transactor{contract: contract}, nil
}

// NewAgentCollectionV1Filterer creates a new log filterer instance of AgentCollectionV1, bound to a specific deployed contract.
func NewAgentCollectionV1Filterer(address common.Address, filterer bind.ContractFilterer) (*AgentCollectionV1Filterer, error) {
	contract, err := bindAgentCollectionV1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AgentCollectionV1Filterer{contract: contract}, nil
}

// bindAgentCollectionV1 binds a generic wrapper to an already deployed contract.
func bindAgentCollectionV1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AgentCollectionV1MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AgentCollectionV1 *AgentCollectionV1Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AgentCollectionV1.Contract.AgentCollectionV1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AgentCollectionV1 *AgentCollectionV1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AgentCollectionV1.Contract.AgentCollectionV1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AgentCollectionV1 *AgentCollectionV1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AgentCollectionV1.Contract.AgentCollectionV1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AgentCollectionV1 *AgentCollectionV1CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AgentCollectionV1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AgentCollectionV1 *AgentCollectionV1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AgentCollectionV1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AgentCollectionV1 *AgentCollectionV1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AgentCollectionV1.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AgentCollectionV1 *AgentCollectionV1Caller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AgentCollectionV1.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AgentCollectionV1 *AgentCollectionV1Session) DEFAULTADMINROLE() ([32]byte, error) {
	return _AgentCollectionV1.Contract.DEFAULTADMINROLE(&_AgentCollectionV1.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AgentCollectionV1 *AgentCollectionV1CallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _AgentCollectionV1.Contract.DEFAULTADMINROLE(&_AgentCollectionV1.CallOpts)
}

// MODERATORROLE is a free data retrieval call binding the contract method 0x797669c9.
//
// Solidity: function MODERATOR_ROLE() view returns(bytes32)
func (_AgentCollectionV1 *AgentCollectionV1Caller) MODERATORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AgentCollectionV1.contract.Call(opts, &out, "MODERATOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MODERATORROLE is a free data retrieval call binding the contract method 0x797669c9.
//
// Solidity: function MODERATOR_ROLE() view returns(bytes32)
func (_AgentCollectionV1 *AgentCollectionV1Session) MODERATORROLE() ([32]byte, error) {
	return _AgentCollectionV1.Contract.MODERATORROLE(&_AgentCollectionV1.CallOpts)
}

// MODERATORROLE is a free data retrieval call binding the contract method 0x797669c9.
//
// Solidity: function MODERATOR_ROLE() view returns(bytes32)
func (_AgentCollectionV1 *AgentCollectionV1CallerSession) MODERATORROLE() ([32]byte, error) {
	return _AgentCollectionV1.Contract.MODERATORROLE(&_AgentCollectionV1.CallOpts)
}

// AddOrUpdateAgentImageProposal is a free data retrieval call binding the contract method 0x202ddfa9.
//
// Solidity: function addOrUpdateAgentImageProposal(uint256 proposalIndex) view returns((uint256,uint256,(string,string)))
func (_AgentCollectionV1 *AgentCollectionV1Caller) AddOrUpdateAgentImageProposal(opts *bind.CallOpts, proposalIndex *big.Int) (AddOrUpdateAgentImageProposal, error) {
	var out []interface{}
	err := _AgentCollectionV1.contract.Call(opts, &out, "addOrUpdateAgentImageProposal", proposalIndex)

	if err != nil {
		return *new(AddOrUpdateAgentImageProposal), err
	}

	out0 := *abi.ConvertType(out[0], new(AddOrUpdateAgentImageProposal)).(*AddOrUpdateAgentImageProposal)

	return out0, err

}

// AddOrUpdateAgentImageProposal is a free data retrieval call binding the contract method 0x202ddfa9.
//
// Solidity: function addOrUpdateAgentImageProposal(uint256 proposalIndex) view returns((uint256,uint256,(string,string)))
func (_AgentCollectionV1 *AgentCollectionV1Session) AddOrUpdateAgentImageProposal(proposalIndex *big.Int) (AddOrUpdateAgentImageProposal, error) {
	return _AgentCollectionV1.Contract.AddOrUpdateAgentImageProposal(&_AgentCollectionV1.CallOpts, proposalIndex)
}

// AddOrUpdateAgentImageProposal is a free data retrieval call binding the contract method 0x202ddfa9.
//
// Solidity: function addOrUpdateAgentImageProposal(uint256 proposalIndex) view returns((uint256,uint256,(string,string)))
func (_AgentCollectionV1 *AgentCollectionV1CallerSession) AddOrUpdateAgentImageProposal(proposalIndex *big.Int) (AddOrUpdateAgentImageProposal, error) {
	return _AgentCollectionV1.Contract.AddOrUpdateAgentImageProposal(&_AgentCollectionV1.CallOpts, proposalIndex)
}

// AddOrUpdateAgentImageProposalsLength is a free data retrieval call binding the contract method 0xc563848c.
//
// Solidity: function addOrUpdateAgentImageProposalsLength() view returns(uint256)
func (_AgentCollectionV1 *AgentCollectionV1Caller) AddOrUpdateAgentImageProposalsLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AgentCollectionV1.contract.Call(opts, &out, "addOrUpdateAgentImageProposalsLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AddOrUpdateAgentImageProposalsLength is a free data retrieval call binding the contract method 0xc563848c.
//
// Solidity: function addOrUpdateAgentImageProposalsLength() view returns(uint256)
func (_AgentCollectionV1 *AgentCollectionV1Session) AddOrUpdateAgentImageProposalsLength() (*big.Int, error) {
	return _AgentCollectionV1.Contract.AddOrUpdateAgentImageProposalsLength(&_AgentCollectionV1.CallOpts)
}

// AddOrUpdateAgentImageProposalsLength is a free data retrieval call binding the contract method 0xc563848c.
//
// Solidity: function addOrUpdateAgentImageProposalsLength() view returns(uint256)
func (_AgentCollectionV1 *AgentCollectionV1CallerSession) AddOrUpdateAgentImageProposalsLength() (*big.Int, error) {
	return _AgentCollectionV1.Contract.AddOrUpdateAgentImageProposalsLength(&_AgentCollectionV1.CallOpts)
}

// AddressOfByIndex is a free data retrieval call binding the contract method 0xc5e7f33b.
//
// Solidity: function addressOfByIndex(uint256 tokenId, uint256 index) view returns((string,string))
func (_AgentCollectionV1 *AgentCollectionV1Caller) AddressOfByIndex(opts *bind.CallOpts, tokenId *big.Int, index *big.Int) (AgentAddress, error) {
	var out []interface{}
	err := _AgentCollectionV1.contract.Call(opts, &out, "addressOfByIndex", tokenId, index)

	if err != nil {
		return *new(AgentAddress), err
	}

	out0 := *abi.ConvertType(out[0], new(AgentAddress)).(*AgentAddress)

	return out0, err

}

// AddressOfByIndex is a free data retrieval call binding the contract method 0xc5e7f33b.
//
// Solidity: function addressOfByIndex(uint256 tokenId, uint256 index) view returns((string,string))
func (_AgentCollectionV1 *AgentCollectionV1Session) AddressOfByIndex(tokenId *big.Int, index *big.Int) (AgentAddress, error) {
	return _AgentCollectionV1.Contract.AddressOfByIndex(&_AgentCollectionV1.CallOpts, tokenId, index)
}

// AddressOfByIndex is a free data retrieval call binding the contract method 0xc5e7f33b.
//
// Solidity: function addressOfByIndex(uint256 tokenId, uint256 index) view returns((string,string))
func (_AgentCollectionV1 *AgentCollectionV1CallerSession) AddressOfByIndex(tokenId *big.Int, index *big.Int) (AgentAddress, error) {
	return _AgentCollectionV1.Contract.AddressOfByIndex(&_AgentCollectionV1.CallOpts, tokenId, index)
}

// AddressToTokenId is a free data retrieval call binding the contract method 0x2f71351a.
//
// Solidity: function addressToTokenId(string addressType, string value) view returns(uint256)
func (_AgentCollectionV1 *AgentCollectionV1Caller) AddressToTokenId(opts *bind.CallOpts, addressType string, value string) (*big.Int, error) {
	var out []interface{}
	err := _AgentCollectionV1.contract.Call(opts, &out, "addressToTokenId", addressType, value)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AddressToTokenId is a free data retrieval call binding the contract method 0x2f71351a.
//
// Solidity: function addressToTokenId(string addressType, string value) view returns(uint256)
func (_AgentCollectionV1 *AgentCollectionV1Session) AddressToTokenId(addressType string, value string) (*big.Int, error) {
	return _AgentCollectionV1.Contract.AddressToTokenId(&_AgentCollectionV1.CallOpts, addressType, value)
}

// AddressToTokenId is a free data retrieval call binding the contract method 0x2f71351a.
//
// Solidity: function addressToTokenId(string addressType, string value) view returns(uint256)
func (_AgentCollectionV1 *AgentCollectionV1CallerSession) AddressToTokenId(addressType string, value string) (*big.Int, error) {
	return _AgentCollectionV1.Contract.AddressToTokenId(&_AgentCollectionV1.CallOpts, addressType, value)
}

// AddressesCountOf is a free data retrieval call binding the contract method 0x4affab02.
//
// Solidity: function addressesCountOf(uint256 tokenId) view returns(uint256)
func (_AgentCollectionV1 *AgentCollectionV1Caller) AddressesCountOf(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AgentCollectionV1.contract.Call(opts, &out, "addressesCountOf", tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AddressesCountOf is a free data retrieval call binding the contract method 0x4affab02.
//
// Solidity: function addressesCountOf(uint256 tokenId) view returns(uint256)
func (_AgentCollectionV1 *AgentCollectionV1Session) AddressesCountOf(tokenId *big.Int) (*big.Int, error) {
	return _AgentCollectionV1.Contract.AddressesCountOf(&_AgentCollectionV1.CallOpts, tokenId)
}

// AddressesCountOf is a free data retrieval call binding the contract method 0x4affab02.
//
// Solidity: function addressesCountOf(uint256 tokenId) view returns(uint256)
func (_AgentCollectionV1 *AgentCollectionV1CallerSession) AddressesCountOf(tokenId *big.Int) (*big.Int, error) {
	return _AgentCollectionV1.Contract.AddressesCountOf(&_AgentCollectionV1.CallOpts, tokenId)
}

// AttributeOfByIndex is a free data retrieval call binding the contract method 0x643bf08d.
//
// Solidity: function attributeOfByIndex(uint256 tokenId, uint256 index) view returns((string,string))
func (_AgentCollectionV1 *AgentCollectionV1Caller) AttributeOfByIndex(opts *bind.CallOpts, tokenId *big.Int, index *big.Int) (KeyValue, error) {
	var out []interface{}
	err := _AgentCollectionV1.contract.Call(opts, &out, "attributeOfByIndex", tokenId, index)

	if err != nil {
		return *new(KeyValue), err
	}

	out0 := *abi.ConvertType(out[0], new(KeyValue)).(*KeyValue)

	return out0, err

}

// AttributeOfByIndex is a free data retrieval call binding the contract method 0x643bf08d.
//
// Solidity: function attributeOfByIndex(uint256 tokenId, uint256 index) view returns((string,string))
func (_AgentCollectionV1 *AgentCollectionV1Session) AttributeOfByIndex(tokenId *big.Int, index *big.Int) (KeyValue, error) {
	return _AgentCollectionV1.Contract.AttributeOfByIndex(&_AgentCollectionV1.CallOpts, tokenId, index)
}

// AttributeOfByIndex is a free data retrieval call binding the contract method 0x643bf08d.
//
// Solidity: function attributeOfByIndex(uint256 tokenId, uint256 index) view returns((string,string))
func (_AgentCollectionV1 *AgentCollectionV1CallerSession) AttributeOfByIndex(tokenId *big.Int, index *big.Int) (KeyValue, error) {
	return _AgentCollectionV1.Contract.AttributeOfByIndex(&_AgentCollectionV1.CallOpts, tokenId, index)
}

// AttributeOfByKey is a free data retrieval call binding the contract method 0x154df1dd.
//
// Solidity: function attributeOfByKey(uint256 tokenId, string key) view returns(string)
func (_AgentCollectionV1 *AgentCollectionV1Caller) AttributeOfByKey(opts *bind.CallOpts, tokenId *big.Int, key string) (string, error) {
	var out []interface{}
	err := _AgentCollectionV1.contract.Call(opts, &out, "attributeOfByKey", tokenId, key)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// AttributeOfByKey is a free data retrieval call binding the contract method 0x154df1dd.
//
// Solidity: function attributeOfByKey(uint256 tokenId, string key) view returns(string)
func (_AgentCollectionV1 *AgentCollectionV1Session) AttributeOfByKey(tokenId *big.Int, key string) (string, error) {
	return _AgentCollectionV1.Contract.AttributeOfByKey(&_AgentCollectionV1.CallOpts, tokenId, key)
}

// AttributeOfByKey is a free data retrieval call binding the contract method 0x154df1dd.
//
// Solidity: function attributeOfByKey(uint256 tokenId, string key) view returns(string)
func (_AgentCollectionV1 *AgentCollectionV1CallerSession) AttributeOfByKey(tokenId *big.Int, key string) (string, error) {
	return _AgentCollectionV1.Contract.AttributeOfByKey(&_AgentCollectionV1.CallOpts, tokenId, key)
}

// AttributesCountOf is a free data retrieval call binding the contract method 0x1609670e.
//
// Solidity: function attributesCountOf(uint256 tokenId) view returns(uint256)
func (_AgentCollectionV1 *AgentCollectionV1Caller) AttributesCountOf(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AgentCollectionV1.contract.Call(opts, &out, "attributesCountOf", tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AttributesCountOf is a free data retrieval call binding the contract method 0x1609670e.
//
// Solidity: function attributesCountOf(uint256 tokenId) view returns(uint256)
func (_AgentCollectionV1 *AgentCollectionV1Session) AttributesCountOf(tokenId *big.Int) (*big.Int, error) {
	return _AgentCollectionV1.Contract.AttributesCountOf(&_AgentCollectionV1.CallOpts, tokenId)
}

// AttributesCountOf is a free data retrieval call binding the contract method 0x1609670e.
//
// Solidity: function attributesCountOf(uint256 tokenId) view returns(uint256)
func (_AgentCollectionV1 *AgentCollectionV1CallerSession) AttributesCountOf(tokenId *big.Int) (*big.Int, error) {
	return _AgentCollectionV1.Contract.AttributesCountOf(&_AgentCollectionV1.CallOpts, tokenId)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_AgentCollectionV1 *AgentCollectionV1Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AgentCollectionV1.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_AgentCollectionV1 *AgentCollectionV1Session) BalanceOf(owner common.Address) (*big.Int, error) {
	return _AgentCollectionV1.Contract.BalanceOf(&_AgentCollectionV1.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_AgentCollectionV1 *AgentCollectionV1CallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _AgentCollectionV1.Contract.BalanceOf(&_AgentCollectionV1.CallOpts, owner)
}

// CollectionsDescriptor is a free data retrieval call binding the contract method 0xab1fdaef.
//
// Solidity: function collectionsDescriptor() view returns(address)
func (_AgentCollectionV1 *AgentCollectionV1Caller) CollectionsDescriptor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AgentCollectionV1.contract.Call(opts, &out, "collectionsDescriptor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CollectionsDescriptor is a free data retrieval call binding the contract method 0xab1fdaef.
//
// Solidity: function collectionsDescriptor() view returns(address)
func (_AgentCollectionV1 *AgentCollectionV1Session) CollectionsDescriptor() (common.Address, error) {
	return _AgentCollectionV1.Contract.CollectionsDescriptor(&_AgentCollectionV1.CallOpts)
}

// CollectionsDescriptor is a free data retrieval call binding the contract method 0xab1fdaef.
//
// Solidity: function collectionsDescriptor() view returns(address)
func (_AgentCollectionV1 *AgentCollectionV1CallerSession) CollectionsDescriptor() (common.Address, error) {
	return _AgentCollectionV1.Contract.CollectionsDescriptor(&_AgentCollectionV1.CallOpts)
}

// CreatorOf is a free data retrieval call binding the contract method 0x589a1743.
//
// Solidity: function creatorOf(uint256 tokenId) view returns(address)
func (_AgentCollectionV1 *AgentCollectionV1Caller) CreatorOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _AgentCollectionV1.contract.Call(opts, &out, "creatorOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CreatorOf is a free data retrieval call binding the contract method 0x589a1743.
//
// Solidity: function creatorOf(uint256 tokenId) view returns(address)
func (_AgentCollectionV1 *AgentCollectionV1Session) CreatorOf(tokenId *big.Int) (common.Address, error) {
	return _AgentCollectionV1.Contract.CreatorOf(&_AgentCollectionV1.CallOpts, tokenId)
}

// CreatorOf is a free data retrieval call binding the contract method 0x589a1743.
//
// Solidity: function creatorOf(uint256 tokenId) view returns(address)
func (_AgentCollectionV1 *AgentCollectionV1CallerSession) CreatorOf(tokenId *big.Int) (common.Address, error) {
	return _AgentCollectionV1.Contract.CreatorOf(&_AgentCollectionV1.CallOpts, tokenId)
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_AgentCollectionV1 *AgentCollectionV1Caller) Description(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AgentCollectionV1.contract.Call(opts, &out, "description")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_AgentCollectionV1 *AgentCollectionV1Session) Description() (string, error) {
	return _AgentCollectionV1.Contract.Description(&_AgentCollectionV1.CallOpts)
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_AgentCollectionV1 *AgentCollectionV1CallerSession) Description() (string, error) {
	return _AgentCollectionV1.Contract.Description(&_AgentCollectionV1.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_AgentCollectionV1 *AgentCollectionV1Caller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _AgentCollectionV1.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_AgentCollectionV1 *AgentCollectionV1Session) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _AgentCollectionV1.Contract.GetApproved(&_AgentCollectionV1.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_AgentCollectionV1 *AgentCollectionV1CallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _AgentCollectionV1.Contract.GetApproved(&_AgentCollectionV1.CallOpts, tokenId)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AgentCollectionV1 *AgentCollectionV1Caller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _AgentCollectionV1.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AgentCollectionV1 *AgentCollectionV1Session) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _AgentCollectionV1.Contract.GetRoleAdmin(&_AgentCollectionV1.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AgentCollectionV1 *AgentCollectionV1CallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _AgentCollectionV1.Contract.GetRoleAdmin(&_AgentCollectionV1.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AgentCollectionV1 *AgentCollectionV1Caller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _AgentCollectionV1.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AgentCollectionV1 *AgentCollectionV1Session) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _AgentCollectionV1.Contract.HasRole(&_AgentCollectionV1.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AgentCollectionV1 *AgentCollectionV1CallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _AgentCollectionV1.Contract.HasRole(&_AgentCollectionV1.CallOpts, role, account)
}

// ImageOfByIndex is a free data retrieval call binding the contract method 0x9c53ba99.
//
// Solidity: function imageOfByIndex(uint256 tokenId, uint256 index) view returns((string,string))
func (_AgentCollectionV1 *AgentCollectionV1Caller) ImageOfByIndex(opts *bind.CallOpts, tokenId *big.Int, index *big.Int) (KeyValue, error) {
	var out []interface{}
	err := _AgentCollectionV1.contract.Call(opts, &out, "imageOfByIndex", tokenId, index)

	if err != nil {
		return *new(KeyValue), err
	}

	out0 := *abi.ConvertType(out[0], new(KeyValue)).(*KeyValue)

	return out0, err

}

// ImageOfByIndex is a free data retrieval call binding the contract method 0x9c53ba99.
//
// Solidity: function imageOfByIndex(uint256 tokenId, uint256 index) view returns((string,string))
func (_AgentCollectionV1 *AgentCollectionV1Session) ImageOfByIndex(tokenId *big.Int, index *big.Int) (KeyValue, error) {
	return _AgentCollectionV1.Contract.ImageOfByIndex(&_AgentCollectionV1.CallOpts, tokenId, index)
}

// ImageOfByIndex is a free data retrieval call binding the contract method 0x9c53ba99.
//
// Solidity: function imageOfByIndex(uint256 tokenId, uint256 index) view returns((string,string))
func (_AgentCollectionV1 *AgentCollectionV1CallerSession) ImageOfByIndex(tokenId *big.Int, index *big.Int) (KeyValue, error) {
	return _AgentCollectionV1.Contract.ImageOfByIndex(&_AgentCollectionV1.CallOpts, tokenId, index)
}

// ImagesCountOf is a free data retrieval call binding the contract method 0xcdd649d6.
//
// Solidity: function imagesCountOf(uint256 tokenId) view returns(uint256)
func (_AgentCollectionV1 *AgentCollectionV1Caller) ImagesCountOf(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AgentCollectionV1.contract.Call(opts, &out, "imagesCountOf", tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ImagesCountOf is a free data retrieval call binding the contract method 0xcdd649d6.
//
// Solidity: function imagesCountOf(uint256 tokenId) view returns(uint256)
func (_AgentCollectionV1 *AgentCollectionV1Session) ImagesCountOf(tokenId *big.Int) (*big.Int, error) {
	return _AgentCollectionV1.Contract.ImagesCountOf(&_AgentCollectionV1.CallOpts, tokenId)
}

// ImagesCountOf is a free data retrieval call binding the contract method 0xcdd649d6.
//
// Solidity: function imagesCountOf(uint256 tokenId) view returns(uint256)
func (_AgentCollectionV1 *AgentCollectionV1CallerSession) ImagesCountOf(tokenId *big.Int) (*big.Int, error) {
	return _AgentCollectionV1.Contract.ImagesCountOf(&_AgentCollectionV1.CallOpts, tokenId)
}

// InformationOf is a free data retrieval call binding the contract method 0x50cffb68.
//
// Solidity: function informationOf(uint256 tokenId) view returns((string,address,(string,string)[],(string,string)[],(bool,uint256),(string,string)[]))
func (_AgentCollectionV1 *AgentCollectionV1Caller) InformationOf(opts *bind.CallOpts, tokenId *big.Int) (AgentInformation, error) {
	var out []interface{}
	err := _AgentCollectionV1.contract.Call(opts, &out, "informationOf", tokenId)

	if err != nil {
		return *new(AgentInformation), err
	}

	out0 := *abi.ConvertType(out[0], new(AgentInformation)).(*AgentInformation)

	return out0, err

}

// InformationOf is a free data retrieval call binding the contract method 0x50cffb68.
//
// Solidity: function informationOf(uint256 tokenId) view returns((string,address,(string,string)[],(string,string)[],(bool,uint256),(string,string)[]))
func (_AgentCollectionV1 *AgentCollectionV1Session) InformationOf(tokenId *big.Int) (AgentInformation, error) {
	return _AgentCollectionV1.Contract.InformationOf(&_AgentCollectionV1.CallOpts, tokenId)
}

// InformationOf is a free data retrieval call binding the contract method 0x50cffb68.
//
// Solidity: function informationOf(uint256 tokenId) view returns((string,address,(string,string)[],(string,string)[],(bool,uint256),(string,string)[]))
func (_AgentCollectionV1 *AgentCollectionV1CallerSession) InformationOf(tokenId *big.Int) (AgentInformation, error) {
	return _AgentCollectionV1.Contract.InformationOf(&_AgentCollectionV1.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_AgentCollectionV1 *AgentCollectionV1Caller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _AgentCollectionV1.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_AgentCollectionV1 *AgentCollectionV1Session) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _AgentCollectionV1.Contract.IsApprovedForAll(&_AgentCollectionV1.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_AgentCollectionV1 *AgentCollectionV1CallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _AgentCollectionV1.Contract.IsApprovedForAll(&_AgentCollectionV1.CallOpts, owner, operator)
}

// IsModerator is a free data retrieval call binding the contract method 0xfa6f3936.
//
// Solidity: function isModerator(address account) view returns(bool)
func (_AgentCollectionV1 *AgentCollectionV1Caller) IsModerator(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _AgentCollectionV1.contract.Call(opts, &out, "isModerator", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsModerator is a free data retrieval call binding the contract method 0xfa6f3936.
//
// Solidity: function isModerator(address account) view returns(bool)
func (_AgentCollectionV1 *AgentCollectionV1Session) IsModerator(account common.Address) (bool, error) {
	return _AgentCollectionV1.Contract.IsModerator(&_AgentCollectionV1.CallOpts, account)
}

// IsModerator is a free data retrieval call binding the contract method 0xfa6f3936.
//
// Solidity: function isModerator(address account) view returns(bool)
func (_AgentCollectionV1 *AgentCollectionV1CallerSession) IsModerator(account common.Address) (bool, error) {
	return _AgentCollectionV1.Contract.IsModerator(&_AgentCollectionV1.CallOpts, account)
}

// MintProposal is a free data retrieval call binding the contract method 0xc088c9a7.
//
// Solidity: function mintProposal(uint256 proposalIndex) view returns((uint256,string,address,address,(string,string)[],(string,string)[],(bool,uint256),(string,string)[]))
func (_AgentCollectionV1 *AgentCollectionV1Caller) MintProposal(opts *bind.CallOpts, proposalIndex *big.Int) (MintProposal, error) {
	var out []interface{}
	err := _AgentCollectionV1.contract.Call(opts, &out, "mintProposal", proposalIndex)

	if err != nil {
		return *new(MintProposal), err
	}

	out0 := *abi.ConvertType(out[0], new(MintProposal)).(*MintProposal)

	return out0, err

}

// MintProposal is a free data retrieval call binding the contract method 0xc088c9a7.
//
// Solidity: function mintProposal(uint256 proposalIndex) view returns((uint256,string,address,address,(string,string)[],(string,string)[],(bool,uint256),(string,string)[]))
func (_AgentCollectionV1 *AgentCollectionV1Session) MintProposal(proposalIndex *big.Int) (MintProposal, error) {
	return _AgentCollectionV1.Contract.MintProposal(&_AgentCollectionV1.CallOpts, proposalIndex)
}

// MintProposal is a free data retrieval call binding the contract method 0xc088c9a7.
//
// Solidity: function mintProposal(uint256 proposalIndex) view returns((uint256,string,address,address,(string,string)[],(string,string)[],(bool,uint256),(string,string)[]))
func (_AgentCollectionV1 *AgentCollectionV1CallerSession) MintProposal(proposalIndex *big.Int) (MintProposal, error) {
	return _AgentCollectionV1.Contract.MintProposal(&_AgentCollectionV1.CallOpts, proposalIndex)
}

// MintProposalsLength is a free data retrieval call binding the contract method 0xf53957ad.
//
// Solidity: function mintProposalsLength() view returns(uint256)
func (_AgentCollectionV1 *AgentCollectionV1Caller) MintProposalsLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AgentCollectionV1.contract.Call(opts, &out, "mintProposalsLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintProposalsLength is a free data retrieval call binding the contract method 0xf53957ad.
//
// Solidity: function mintProposalsLength() view returns(uint256)
func (_AgentCollectionV1 *AgentCollectionV1Session) MintProposalsLength() (*big.Int, error) {
	return _AgentCollectionV1.Contract.MintProposalsLength(&_AgentCollectionV1.CallOpts)
}

// MintProposalsLength is a free data retrieval call binding the contract method 0xf53957ad.
//
// Solidity: function mintProposalsLength() view returns(uint256)
func (_AgentCollectionV1 *AgentCollectionV1CallerSession) MintProposalsLength() (*big.Int, error) {
	return _AgentCollectionV1.Contract.MintProposalsLength(&_AgentCollectionV1.CallOpts)
}

// ModeratorCount is a free data retrieval call binding the contract method 0x0ee7e60b.
//
// Solidity: function moderatorCount() view returns(uint256)
func (_AgentCollectionV1 *AgentCollectionV1Caller) ModeratorCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AgentCollectionV1.contract.Call(opts, &out, "moderatorCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ModeratorCount is a free data retrieval call binding the contract method 0x0ee7e60b.
//
// Solidity: function moderatorCount() view returns(uint256)
func (_AgentCollectionV1 *AgentCollectionV1Session) ModeratorCount() (*big.Int, error) {
	return _AgentCollectionV1.Contract.ModeratorCount(&_AgentCollectionV1.CallOpts)
}

// ModeratorCount is a free data retrieval call binding the contract method 0x0ee7e60b.
//
// Solidity: function moderatorCount() view returns(uint256)
func (_AgentCollectionV1 *AgentCollectionV1CallerSession) ModeratorCount() (*big.Int, error) {
	return _AgentCollectionV1.Contract.ModeratorCount(&_AgentCollectionV1.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_AgentCollectionV1 *AgentCollectionV1Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AgentCollectionV1.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_AgentCollectionV1 *AgentCollectionV1Session) Name() (string, error) {
	return _AgentCollectionV1.Contract.Name(&_AgentCollectionV1.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_AgentCollectionV1 *AgentCollectionV1CallerSession) Name() (string, error) {
	return _AgentCollectionV1.Contract.Name(&_AgentCollectionV1.CallOpts)
}

// NameOf is a free data retrieval call binding the contract method 0x051a2664.
//
// Solidity: function nameOf(uint256 tokenId) view returns(string)
func (_AgentCollectionV1 *AgentCollectionV1Caller) NameOf(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _AgentCollectionV1.contract.Call(opts, &out, "nameOf", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// NameOf is a free data retrieval call binding the contract method 0x051a2664.
//
// Solidity: function nameOf(uint256 tokenId) view returns(string)
func (_AgentCollectionV1 *AgentCollectionV1Session) NameOf(tokenId *big.Int) (string, error) {
	return _AgentCollectionV1.Contract.NameOf(&_AgentCollectionV1.CallOpts, tokenId)
}

// NameOf is a free data retrieval call binding the contract method 0x051a2664.
//
// Solidity: function nameOf(uint256 tokenId) view returns(string)
func (_AgentCollectionV1 *AgentCollectionV1CallerSession) NameOf(tokenId *big.Int) (string, error) {
	return _AgentCollectionV1.Contract.NameOf(&_AgentCollectionV1.CallOpts, tokenId)
}

// NameToTokenId is a free data retrieval call binding the contract method 0xdd001254.
//
// Solidity: function nameToTokenId(string agentName) view returns(uint256)
func (_AgentCollectionV1 *AgentCollectionV1Caller) NameToTokenId(opts *bind.CallOpts, agentName string) (*big.Int, error) {
	var out []interface{}
	err := _AgentCollectionV1.contract.Call(opts, &out, "nameToTokenId", agentName)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NameToTokenId is a free data retrieval call binding the contract method 0xdd001254.
//
// Solidity: function nameToTokenId(string agentName) view returns(uint256)
func (_AgentCollectionV1 *AgentCollectionV1Session) NameToTokenId(agentName string) (*big.Int, error) {
	return _AgentCollectionV1.Contract.NameToTokenId(&_AgentCollectionV1.CallOpts, agentName)
}

// NameToTokenId is a free data retrieval call binding the contract method 0xdd001254.
//
// Solidity: function nameToTokenId(string agentName) view returns(uint256)
func (_AgentCollectionV1 *AgentCollectionV1CallerSession) NameToTokenId(agentName string) (*big.Int, error) {
	return _AgentCollectionV1.Contract.NameToTokenId(&_AgentCollectionV1.CallOpts, agentName)
}

// NextTokenId is a free data retrieval call binding the contract method 0x75794a3c.
//
// Solidity: function nextTokenId() view returns(uint256)
func (_AgentCollectionV1 *AgentCollectionV1Caller) NextTokenId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AgentCollectionV1.contract.Call(opts, &out, "nextTokenId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextTokenId is a free data retrieval call binding the contract method 0x75794a3c.
//
// Solidity: function nextTokenId() view returns(uint256)
func (_AgentCollectionV1 *AgentCollectionV1Session) NextTokenId() (*big.Int, error) {
	return _AgentCollectionV1.Contract.NextTokenId(&_AgentCollectionV1.CallOpts)
}

// NextTokenId is a free data retrieval call binding the contract method 0x75794a3c.
//
// Solidity: function nextTokenId() view returns(uint256)
func (_AgentCollectionV1 *AgentCollectionV1CallerSession) NextTokenId() (*big.Int, error) {
	return _AgentCollectionV1.Contract.NextTokenId(&_AgentCollectionV1.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_AgentCollectionV1 *AgentCollectionV1Caller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _AgentCollectionV1.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_AgentCollectionV1 *AgentCollectionV1Session) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _AgentCollectionV1.Contract.OwnerOf(&_AgentCollectionV1.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_AgentCollectionV1 *AgentCollectionV1CallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _AgentCollectionV1.Contract.OwnerOf(&_AgentCollectionV1.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AgentCollectionV1 *AgentCollectionV1Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _AgentCollectionV1.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AgentCollectionV1 *AgentCollectionV1Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _AgentCollectionV1.Contract.SupportsInterface(&_AgentCollectionV1.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AgentCollectionV1 *AgentCollectionV1CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _AgentCollectionV1.Contract.SupportsInterface(&_AgentCollectionV1.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_AgentCollectionV1 *AgentCollectionV1Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AgentCollectionV1.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_AgentCollectionV1 *AgentCollectionV1Session) Symbol() (string, error) {
	return _AgentCollectionV1.Contract.Symbol(&_AgentCollectionV1.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_AgentCollectionV1 *AgentCollectionV1CallerSession) Symbol() (string, error) {
	return _AgentCollectionV1.Contract.Symbol(&_AgentCollectionV1.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_AgentCollectionV1 *AgentCollectionV1Caller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _AgentCollectionV1.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_AgentCollectionV1 *AgentCollectionV1Session) TokenURI(tokenId *big.Int) (string, error) {
	return _AgentCollectionV1.Contract.TokenURI(&_AgentCollectionV1.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_AgentCollectionV1 *AgentCollectionV1CallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _AgentCollectionV1.Contract.TokenURI(&_AgentCollectionV1.CallOpts, tokenId)
}

// AddAgentAddress is a paid mutator transaction binding the contract method 0x84f62845.
//
// Solidity: function addAgentAddress(uint256 tokenId, string addressType, string value) returns()
func (_AgentCollectionV1 *AgentCollectionV1Transactor) AddAgentAddress(opts *bind.TransactOpts, tokenId *big.Int, addressType string, value string) (*types.Transaction, error) {
	return _AgentCollectionV1.contract.Transact(opts, "addAgentAddress", tokenId, addressType, value)
}

// AddAgentAddress is a paid mutator transaction binding the contract method 0x84f62845.
//
// Solidity: function addAgentAddress(uint256 tokenId, string addressType, string value) returns()
func (_AgentCollectionV1 *AgentCollectionV1Session) AddAgentAddress(tokenId *big.Int, addressType string, value string) (*types.Transaction, error) {
	return _AgentCollectionV1.Contract.AddAgentAddress(&_AgentCollectionV1.TransactOpts, tokenId, addressType, value)
}

// AddAgentAddress is a paid mutator transaction binding the contract method 0x84f62845.
//
// Solidity: function addAgentAddress(uint256 tokenId, string addressType, string value) returns()
func (_AgentCollectionV1 *AgentCollectionV1TransactorSession) AddAgentAddress(tokenId *big.Int, addressType string, value string) (*types.Transaction, error) {
	return _AgentCollectionV1.Contract.AddAgentAddress(&_AgentCollectionV1.TransactOpts, tokenId, addressType, value)
}

// AddOrUpdateAgentAttribute is a paid mutator transaction binding the contract method 0xd769ad0f.
//
// Solidity: function addOrUpdateAgentAttribute(uint256 tokenId, string key, string value) returns()
func (_AgentCollectionV1 *AgentCollectionV1Transactor) AddOrUpdateAgentAttribute(opts *bind.TransactOpts, tokenId *big.Int, key string, value string) (*types.Transaction, error) {
	return _AgentCollectionV1.contract.Transact(opts, "addOrUpdateAgentAttribute", tokenId, key, value)
}

// AddOrUpdateAgentAttribute is a paid mutator transaction binding the contract method 0xd769ad0f.
//
// Solidity: function addOrUpdateAgentAttribute(uint256 tokenId, string key, string value) returns()
func (_AgentCollectionV1 *AgentCollectionV1Session) AddOrUpdateAgentAttribute(tokenId *big.Int, key string, value string) (*types.Transaction, error) {
	return _AgentCollectionV1.Contract.AddOrUpdateAgentAttribute(&_AgentCollectionV1.TransactOpts, tokenId, key, value)
}

// AddOrUpdateAgentAttribute is a paid mutator transaction binding the contract method 0xd769ad0f.
//
// Solidity: function addOrUpdateAgentAttribute(uint256 tokenId, string key, string value) returns()
func (_AgentCollectionV1 *AgentCollectionV1TransactorSession) AddOrUpdateAgentAttribute(tokenId *big.Int, key string, value string) (*types.Transaction, error) {
	return _AgentCollectionV1.Contract.AddOrUpdateAgentAttribute(&_AgentCollectionV1.TransactOpts, tokenId, key, value)
}

// AddOrUpdateAgentImage is a paid mutator transaction binding the contract method 0x86abf25b.
//
// Solidity: function addOrUpdateAgentImage(uint256 tokenId, (string,string) image) returns()
func (_AgentCollectionV1 *AgentCollectionV1Transactor) AddOrUpdateAgentImage(opts *bind.TransactOpts, tokenId *big.Int, image KeyValue) (*types.Transaction, error) {
	return _AgentCollectionV1.contract.Transact(opts, "addOrUpdateAgentImage", tokenId, image)
}

// AddOrUpdateAgentImage is a paid mutator transaction binding the contract method 0x86abf25b.
//
// Solidity: function addOrUpdateAgentImage(uint256 tokenId, (string,string) image) returns()
func (_AgentCollectionV1 *AgentCollectionV1Session) AddOrUpdateAgentImage(tokenId *big.Int, image KeyValue) (*types.Transaction, error) {
	return _AgentCollectionV1.Contract.AddOrUpdateAgentImage(&_AgentCollectionV1.TransactOpts, tokenId, image)
}

// AddOrUpdateAgentImage is a paid mutator transaction binding the contract method 0x86abf25b.
//
// Solidity: function addOrUpdateAgentImage(uint256 tokenId, (string,string) image) returns()
func (_AgentCollectionV1 *AgentCollectionV1TransactorSession) AddOrUpdateAgentImage(tokenId *big.Int, image KeyValue) (*types.Transaction, error) {
	return _AgentCollectionV1.Contract.AddOrUpdateAgentImage(&_AgentCollectionV1.TransactOpts, tokenId, image)
}

// AddOrUpdateAgentImageFromProposal is a paid mutator transaction binding the contract method 0xdc5d05f2.
//
// Solidity: function addOrUpdateAgentImageFromProposal(uint256 proposalId) returns()
func (_AgentCollectionV1 *AgentCollectionV1Transactor) AddOrUpdateAgentImageFromProposal(opts *bind.TransactOpts, proposalId *big.Int) (*types.Transaction, error) {
	return _AgentCollectionV1.contract.Transact(opts, "addOrUpdateAgentImageFromProposal", proposalId)
}

// AddOrUpdateAgentImageFromProposal is a paid mutator transaction binding the contract method 0xdc5d05f2.
//
// Solidity: function addOrUpdateAgentImageFromProposal(uint256 proposalId) returns()
func (_AgentCollectionV1 *AgentCollectionV1Session) AddOrUpdateAgentImageFromProposal(proposalId *big.Int) (*types.Transaction, error) {
	return _AgentCollectionV1.Contract.AddOrUpdateAgentImageFromProposal(&_AgentCollectionV1.TransactOpts, proposalId)
}

// AddOrUpdateAgentImageFromProposal is a paid mutator transaction binding the contract method 0xdc5d05f2.
//
// Solidity: function addOrUpdateAgentImageFromProposal(uint256 proposalId) returns()
func (_AgentCollectionV1 *AgentCollectionV1TransactorSession) AddOrUpdateAgentImageFromProposal(proposalId *big.Int) (*types.Transaction, error) {
	return _AgentCollectionV1.Contract.AddOrUpdateAgentImageFromProposal(&_AgentCollectionV1.TransactOpts, proposalId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_AgentCollectionV1 *AgentCollectionV1Transactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _AgentCollectionV1.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_AgentCollectionV1 *AgentCollectionV1Session) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _AgentCollectionV1.Contract.Approve(&_AgentCollectionV1.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_AgentCollectionV1 *AgentCollectionV1TransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _AgentCollectionV1.Contract.Approve(&_AgentCollectionV1.TransactOpts, to, tokenId)
}

// CreateAddOrUpdateAgentImageProposal is a paid mutator transaction binding the contract method 0x1b2b4d36.
//
// Solidity: function createAddOrUpdateAgentImageProposal(uint256 tokenId, (string,string) image) returns()
func (_AgentCollectionV1 *AgentCollectionV1Transactor) CreateAddOrUpdateAgentImageProposal(opts *bind.TransactOpts, tokenId *big.Int, image KeyValue) (*types.Transaction, error) {
	return _AgentCollectionV1.contract.Transact(opts, "createAddOrUpdateAgentImageProposal", tokenId, image)
}

// CreateAddOrUpdateAgentImageProposal is a paid mutator transaction binding the contract method 0x1b2b4d36.
//
// Solidity: function createAddOrUpdateAgentImageProposal(uint256 tokenId, (string,string) image) returns()
func (_AgentCollectionV1 *AgentCollectionV1Session) CreateAddOrUpdateAgentImageProposal(tokenId *big.Int, image KeyValue) (*types.Transaction, error) {
	return _AgentCollectionV1.Contract.CreateAddOrUpdateAgentImageProposal(&_AgentCollectionV1.TransactOpts, tokenId, image)
}

// CreateAddOrUpdateAgentImageProposal is a paid mutator transaction binding the contract method 0x1b2b4d36.
//
// Solidity: function createAddOrUpdateAgentImageProposal(uint256 tokenId, (string,string) image) returns()
func (_AgentCollectionV1 *AgentCollectionV1TransactorSession) CreateAddOrUpdateAgentImageProposal(tokenId *big.Int, image KeyValue) (*types.Transaction, error) {
	return _AgentCollectionV1.Contract.CreateAddOrUpdateAgentImageProposal(&_AgentCollectionV1.TransactOpts, tokenId, image)
}

// CreateMintProposal is a paid mutator transaction binding the contract method 0xa801287f.
//
// Solidity: function createMintProposal(address to, string agentName, (string,string)[] addresses, (string,string)[] images, (string,string)[] attributes, (bool,uint256) cloneOf) returns()
func (_AgentCollectionV1 *AgentCollectionV1Transactor) CreateMintProposal(opts *bind.TransactOpts, to common.Address, agentName string, addresses []AgentAddress, images []KeyValue, attributes []KeyValue, cloneOf NullableUint256) (*types.Transaction, error) {
	return _AgentCollectionV1.contract.Transact(opts, "createMintProposal", to, agentName, addresses, images, attributes, cloneOf)
}

// CreateMintProposal is a paid mutator transaction binding the contract method 0xa801287f.
//
// Solidity: function createMintProposal(address to, string agentName, (string,string)[] addresses, (string,string)[] images, (string,string)[] attributes, (bool,uint256) cloneOf) returns()
func (_AgentCollectionV1 *AgentCollectionV1Session) CreateMintProposal(to common.Address, agentName string, addresses []AgentAddress, images []KeyValue, attributes []KeyValue, cloneOf NullableUint256) (*types.Transaction, error) {
	return _AgentCollectionV1.Contract.CreateMintProposal(&_AgentCollectionV1.TransactOpts, to, agentName, addresses, images, attributes, cloneOf)
}

// CreateMintProposal is a paid mutator transaction binding the contract method 0xa801287f.
//
// Solidity: function createMintProposal(address to, string agentName, (string,string)[] addresses, (string,string)[] images, (string,string)[] attributes, (bool,uint256) cloneOf) returns()
func (_AgentCollectionV1 *AgentCollectionV1TransactorSession) CreateMintProposal(to common.Address, agentName string, addresses []AgentAddress, images []KeyValue, attributes []KeyValue, cloneOf NullableUint256) (*types.Transaction, error) {
	return _AgentCollectionV1.Contract.CreateMintProposal(&_AgentCollectionV1.TransactOpts, to, agentName, addresses, images, attributes, cloneOf)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AgentCollectionV1 *AgentCollectionV1Transactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AgentCollectionV1.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AgentCollectionV1 *AgentCollectionV1Session) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AgentCollectionV1.Contract.GrantRole(&_AgentCollectionV1.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AgentCollectionV1 *AgentCollectionV1TransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AgentCollectionV1.Contract.GrantRole(&_AgentCollectionV1.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x9630c8ac.
//
// Solidity: function initialize(string name_, string symbol_, string initialDescription, address admin, address moderator, address initialCollectionsDescriptor) returns()
func (_AgentCollectionV1 *AgentCollectionV1Transactor) Initialize(opts *bind.TransactOpts, name_ string, symbol_ string, initialDescription string, admin common.Address, moderator common.Address, initialCollectionsDescriptor common.Address) (*types.Transaction, error) {
	return _AgentCollectionV1.contract.Transact(opts, "initialize", name_, symbol_, initialDescription, admin, moderator, initialCollectionsDescriptor)
}

// Initialize is a paid mutator transaction binding the contract method 0x9630c8ac.
//
// Solidity: function initialize(string name_, string symbol_, string initialDescription, address admin, address moderator, address initialCollectionsDescriptor) returns()
func (_AgentCollectionV1 *AgentCollectionV1Session) Initialize(name_ string, symbol_ string, initialDescription string, admin common.Address, moderator common.Address, initialCollectionsDescriptor common.Address) (*types.Transaction, error) {
	return _AgentCollectionV1.Contract.Initialize(&_AgentCollectionV1.TransactOpts, name_, symbol_, initialDescription, admin, moderator, initialCollectionsDescriptor)
}

// Initialize is a paid mutator transaction binding the contract method 0x9630c8ac.
//
// Solidity: function initialize(string name_, string symbol_, string initialDescription, address admin, address moderator, address initialCollectionsDescriptor) returns()
func (_AgentCollectionV1 *AgentCollectionV1TransactorSession) Initialize(name_ string, symbol_ string, initialDescription string, admin common.Address, moderator common.Address, initialCollectionsDescriptor common.Address) (*types.Transaction, error) {
	return _AgentCollectionV1.Contract.Initialize(&_AgentCollectionV1.TransactOpts, name_, symbol_, initialDescription, admin, moderator, initialCollectionsDescriptor)
}

// Mint is a paid mutator transaction binding the contract method 0xe7fc82d9.
//
// Solidity: function mint(address to, string agentName, (string,string)[] addresses, (string,string)[] images, (string,string)[] attributes, (bool,uint256) cloneOf) returns()
func (_AgentCollectionV1 *AgentCollectionV1Transactor) Mint(opts *bind.TransactOpts, to common.Address, agentName string, addresses []AgentAddress, images []KeyValue, attributes []KeyValue, cloneOf NullableUint256) (*types.Transaction, error) {
	return _AgentCollectionV1.contract.Transact(opts, "mint", to, agentName, addresses, images, attributes, cloneOf)
}

// Mint is a paid mutator transaction binding the contract method 0xe7fc82d9.
//
// Solidity: function mint(address to, string agentName, (string,string)[] addresses, (string,string)[] images, (string,string)[] attributes, (bool,uint256) cloneOf) returns()
func (_AgentCollectionV1 *AgentCollectionV1Session) Mint(to common.Address, agentName string, addresses []AgentAddress, images []KeyValue, attributes []KeyValue, cloneOf NullableUint256) (*types.Transaction, error) {
	return _AgentCollectionV1.Contract.Mint(&_AgentCollectionV1.TransactOpts, to, agentName, addresses, images, attributes, cloneOf)
}

// Mint is a paid mutator transaction binding the contract method 0xe7fc82d9.
//
// Solidity: function mint(address to, string agentName, (string,string)[] addresses, (string,string)[] images, (string,string)[] attributes, (bool,uint256) cloneOf) returns()
func (_AgentCollectionV1 *AgentCollectionV1TransactorSession) Mint(to common.Address, agentName string, addresses []AgentAddress, images []KeyValue, attributes []KeyValue, cloneOf NullableUint256) (*types.Transaction, error) {
	return _AgentCollectionV1.Contract.Mint(&_AgentCollectionV1.TransactOpts, to, agentName, addresses, images, attributes, cloneOf)
}

// MintFromProposal is a paid mutator transaction binding the contract method 0xe6f1d66a.
//
// Solidity: function mintFromProposal(uint256 proposalId) returns()
func (_AgentCollectionV1 *AgentCollectionV1Transactor) MintFromProposal(opts *bind.TransactOpts, proposalId *big.Int) (*types.Transaction, error) {
	return _AgentCollectionV1.contract.Transact(opts, "mintFromProposal", proposalId)
}

// MintFromProposal is a paid mutator transaction binding the contract method 0xe6f1d66a.
//
// Solidity: function mintFromProposal(uint256 proposalId) returns()
func (_AgentCollectionV1 *AgentCollectionV1Session) MintFromProposal(proposalId *big.Int) (*types.Transaction, error) {
	return _AgentCollectionV1.Contract.MintFromProposal(&_AgentCollectionV1.TransactOpts, proposalId)
}

// MintFromProposal is a paid mutator transaction binding the contract method 0xe6f1d66a.
//
// Solidity: function mintFromProposal(uint256 proposalId) returns()
func (_AgentCollectionV1 *AgentCollectionV1TransactorSession) MintFromProposal(proposalId *big.Int) (*types.Transaction, error) {
	return _AgentCollectionV1.Contract.MintFromProposal(&_AgentCollectionV1.TransactOpts, proposalId)
}

// RefuseAddOrUpdateAgentImageProposal is a paid mutator transaction binding the contract method 0xdb8076e5.
//
// Solidity: function refuseAddOrUpdateAgentImageProposal(uint256 proposalId) returns()
func (_AgentCollectionV1 *AgentCollectionV1Transactor) RefuseAddOrUpdateAgentImageProposal(opts *bind.TransactOpts, proposalId *big.Int) (*types.Transaction, error) {
	return _AgentCollectionV1.contract.Transact(opts, "refuseAddOrUpdateAgentImageProposal", proposalId)
}

// RefuseAddOrUpdateAgentImageProposal is a paid mutator transaction binding the contract method 0xdb8076e5.
//
// Solidity: function refuseAddOrUpdateAgentImageProposal(uint256 proposalId) returns()
func (_AgentCollectionV1 *AgentCollectionV1Session) RefuseAddOrUpdateAgentImageProposal(proposalId *big.Int) (*types.Transaction, error) {
	return _AgentCollectionV1.Contract.RefuseAddOrUpdateAgentImageProposal(&_AgentCollectionV1.TransactOpts, proposalId)
}

// RefuseAddOrUpdateAgentImageProposal is a paid mutator transaction binding the contract method 0xdb8076e5.
//
// Solidity: function refuseAddOrUpdateAgentImageProposal(uint256 proposalId) returns()
func (_AgentCollectionV1 *AgentCollectionV1TransactorSession) RefuseAddOrUpdateAgentImageProposal(proposalId *big.Int) (*types.Transaction, error) {
	return _AgentCollectionV1.Contract.RefuseAddOrUpdateAgentImageProposal(&_AgentCollectionV1.TransactOpts, proposalId)
}

// RefuseMintProposal is a paid mutator transaction binding the contract method 0xb31fcdbf.
//
// Solidity: function refuseMintProposal(uint256 proposalId) returns()
func (_AgentCollectionV1 *AgentCollectionV1Transactor) RefuseMintProposal(opts *bind.TransactOpts, proposalId *big.Int) (*types.Transaction, error) {
	return _AgentCollectionV1.contract.Transact(opts, "refuseMintProposal", proposalId)
}

// RefuseMintProposal is a paid mutator transaction binding the contract method 0xb31fcdbf.
//
// Solidity: function refuseMintProposal(uint256 proposalId) returns()
func (_AgentCollectionV1 *AgentCollectionV1Session) RefuseMintProposal(proposalId *big.Int) (*types.Transaction, error) {
	return _AgentCollectionV1.Contract.RefuseMintProposal(&_AgentCollectionV1.TransactOpts, proposalId)
}

// RefuseMintProposal is a paid mutator transaction binding the contract method 0xb31fcdbf.
//
// Solidity: function refuseMintProposal(uint256 proposalId) returns()
func (_AgentCollectionV1 *AgentCollectionV1TransactorSession) RefuseMintProposal(proposalId *big.Int) (*types.Transaction, error) {
	return _AgentCollectionV1.Contract.RefuseMintProposal(&_AgentCollectionV1.TransactOpts, proposalId)
}

// RemoveAgentAddress is a paid mutator transaction binding the contract method 0xc643edfe.
//
// Solidity: function removeAgentAddress(uint256 tokenId, string addressType, string value) returns()
func (_AgentCollectionV1 *AgentCollectionV1Transactor) RemoveAgentAddress(opts *bind.TransactOpts, tokenId *big.Int, addressType string, value string) (*types.Transaction, error) {
	return _AgentCollectionV1.contract.Transact(opts, "removeAgentAddress", tokenId, addressType, value)
}

// RemoveAgentAddress is a paid mutator transaction binding the contract method 0xc643edfe.
//
// Solidity: function removeAgentAddress(uint256 tokenId, string addressType, string value) returns()
func (_AgentCollectionV1 *AgentCollectionV1Session) RemoveAgentAddress(tokenId *big.Int, addressType string, value string) (*types.Transaction, error) {
	return _AgentCollectionV1.Contract.RemoveAgentAddress(&_AgentCollectionV1.TransactOpts, tokenId, addressType, value)
}

// RemoveAgentAddress is a paid mutator transaction binding the contract method 0xc643edfe.
//
// Solidity: function removeAgentAddress(uint256 tokenId, string addressType, string value) returns()
func (_AgentCollectionV1 *AgentCollectionV1TransactorSession) RemoveAgentAddress(tokenId *big.Int, addressType string, value string) (*types.Transaction, error) {
	return _AgentCollectionV1.Contract.RemoveAgentAddress(&_AgentCollectionV1.TransactOpts, tokenId, addressType, value)
}

// RemoveAgentImage is a paid mutator transaction binding the contract method 0x05393a28.
//
// Solidity: function removeAgentImage(uint256 tokenId, string key) returns()
func (_AgentCollectionV1 *AgentCollectionV1Transactor) RemoveAgentImage(opts *bind.TransactOpts, tokenId *big.Int, key string) (*types.Transaction, error) {
	return _AgentCollectionV1.contract.Transact(opts, "removeAgentImage", tokenId, key)
}

// RemoveAgentImage is a paid mutator transaction binding the contract method 0x05393a28.
//
// Solidity: function removeAgentImage(uint256 tokenId, string key) returns()
func (_AgentCollectionV1 *AgentCollectionV1Session) RemoveAgentImage(tokenId *big.Int, key string) (*types.Transaction, error) {
	return _AgentCollectionV1.Contract.RemoveAgentImage(&_AgentCollectionV1.TransactOpts, tokenId, key)
}

// RemoveAgentImage is a paid mutator transaction binding the contract method 0x05393a28.
//
// Solidity: function removeAgentImage(uint256 tokenId, string key) returns()
func (_AgentCollectionV1 *AgentCollectionV1TransactorSession) RemoveAgentImage(tokenId *big.Int, key string) (*types.Transaction, error) {
	return _AgentCollectionV1.Contract.RemoveAgentImage(&_AgentCollectionV1.TransactOpts, tokenId, key)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_AgentCollectionV1 *AgentCollectionV1Transactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _AgentCollectionV1.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_AgentCollectionV1 *AgentCollectionV1Session) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _AgentCollectionV1.Contract.RenounceRole(&_AgentCollectionV1.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_AgentCollectionV1 *AgentCollectionV1TransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _AgentCollectionV1.Contract.RenounceRole(&_AgentCollectionV1.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AgentCollectionV1 *AgentCollectionV1Transactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AgentCollectionV1.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AgentCollectionV1 *AgentCollectionV1Session) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AgentCollectionV1.Contract.RevokeRole(&_AgentCollectionV1.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AgentCollectionV1 *AgentCollectionV1TransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AgentCollectionV1.Contract.RevokeRole(&_AgentCollectionV1.TransactOpts, role, account)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_AgentCollectionV1 *AgentCollectionV1Transactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _AgentCollectionV1.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_AgentCollectionV1 *AgentCollectionV1Session) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _AgentCollectionV1.Contract.SafeTransferFrom(&_AgentCollectionV1.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_AgentCollectionV1 *AgentCollectionV1TransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _AgentCollectionV1.Contract.SafeTransferFrom(&_AgentCollectionV1.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_AgentCollectionV1 *AgentCollectionV1Transactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _AgentCollectionV1.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_AgentCollectionV1 *AgentCollectionV1Session) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _AgentCollectionV1.Contract.SafeTransferFrom0(&_AgentCollectionV1.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_AgentCollectionV1 *AgentCollectionV1TransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _AgentCollectionV1.Contract.SafeTransferFrom0(&_AgentCollectionV1.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_AgentCollectionV1 *AgentCollectionV1Transactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _AgentCollectionV1.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_AgentCollectionV1 *AgentCollectionV1Session) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _AgentCollectionV1.Contract.SetApprovalForAll(&_AgentCollectionV1.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_AgentCollectionV1 *AgentCollectionV1TransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _AgentCollectionV1.Contract.SetApprovalForAll(&_AgentCollectionV1.TransactOpts, operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_AgentCollectionV1 *AgentCollectionV1Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _AgentCollectionV1.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_AgentCollectionV1 *AgentCollectionV1Session) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _AgentCollectionV1.Contract.TransferFrom(&_AgentCollectionV1.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_AgentCollectionV1 *AgentCollectionV1TransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _AgentCollectionV1.Contract.TransferFrom(&_AgentCollectionV1.TransactOpts, from, to, tokenId)
}

// UpdateCollectionsDescriptor is a paid mutator transaction binding the contract method 0x798931a3.
//
// Solidity: function updateCollectionsDescriptor(address newCollectionsDescriptor) returns()
func (_AgentCollectionV1 *AgentCollectionV1Transactor) UpdateCollectionsDescriptor(opts *bind.TransactOpts, newCollectionsDescriptor common.Address) (*types.Transaction, error) {
	return _AgentCollectionV1.contract.Transact(opts, "updateCollectionsDescriptor", newCollectionsDescriptor)
}

// UpdateCollectionsDescriptor is a paid mutator transaction binding the contract method 0x798931a3.
//
// Solidity: function updateCollectionsDescriptor(address newCollectionsDescriptor) returns()
func (_AgentCollectionV1 *AgentCollectionV1Session) UpdateCollectionsDescriptor(newCollectionsDescriptor common.Address) (*types.Transaction, error) {
	return _AgentCollectionV1.Contract.UpdateCollectionsDescriptor(&_AgentCollectionV1.TransactOpts, newCollectionsDescriptor)
}

// UpdateCollectionsDescriptor is a paid mutator transaction binding the contract method 0x798931a3.
//
// Solidity: function updateCollectionsDescriptor(address newCollectionsDescriptor) returns()
func (_AgentCollectionV1 *AgentCollectionV1TransactorSession) UpdateCollectionsDescriptor(newCollectionsDescriptor common.Address) (*types.Transaction, error) {
	return _AgentCollectionV1.Contract.UpdateCollectionsDescriptor(&_AgentCollectionV1.TransactOpts, newCollectionsDescriptor)
}

// UpdateDescription is a paid mutator transaction binding the contract method 0xe735b48a.
//
// Solidity: function updateDescription(string newDescription) returns()
func (_AgentCollectionV1 *AgentCollectionV1Transactor) UpdateDescription(opts *bind.TransactOpts, newDescription string) (*types.Transaction, error) {
	return _AgentCollectionV1.contract.Transact(opts, "updateDescription", newDescription)
}

// UpdateDescription is a paid mutator transaction binding the contract method 0xe735b48a.
//
// Solidity: function updateDescription(string newDescription) returns()
func (_AgentCollectionV1 *AgentCollectionV1Session) UpdateDescription(newDescription string) (*types.Transaction, error) {
	return _AgentCollectionV1.Contract.UpdateDescription(&_AgentCollectionV1.TransactOpts, newDescription)
}

// UpdateDescription is a paid mutator transaction binding the contract method 0xe735b48a.
//
// Solidity: function updateDescription(string newDescription) returns()
func (_AgentCollectionV1 *AgentCollectionV1TransactorSession) UpdateDescription(newDescription string) (*types.Transaction, error) {
	return _AgentCollectionV1.Contract.UpdateDescription(&_AgentCollectionV1.TransactOpts, newDescription)
}

// AgentCollectionV1AgentAddressAddedIterator is returned from FilterAgentAddressAdded and is used to iterate over the raw logs and unpacked data for AgentAddressAdded events raised by the AgentCollectionV1 contract.
type AgentCollectionV1AgentAddressAddedIterator struct {
	Event *AgentCollectionV1AgentAddressAdded // Event containing the contract specifics and raw log

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
func (it *AgentCollectionV1AgentAddressAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentCollectionV1AgentAddressAdded)
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
		it.Event = new(AgentCollectionV1AgentAddressAdded)
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
func (it *AgentCollectionV1AgentAddressAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentCollectionV1AgentAddressAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentCollectionV1AgentAddressAdded represents a AgentAddressAdded event raised by the AgentCollectionV1 contract.
type AgentCollectionV1AgentAddressAdded struct {
	TokenId     *big.Int
	AddressType string
	Value       string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAgentAddressAdded is a free log retrieval operation binding the contract event 0x2f6a4339fb3581220ac44c101401de6fadea0c48ccffdb219e7cc2daf98b62b3.
//
// Solidity: event AgentAddressAdded(uint256 indexed tokenId, string addressType, string value)
func (_AgentCollectionV1 *AgentCollectionV1Filterer) FilterAgentAddressAdded(opts *bind.FilterOpts, tokenId []*big.Int) (*AgentCollectionV1AgentAddressAddedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _AgentCollectionV1.contract.FilterLogs(opts, "AgentAddressAdded", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &AgentCollectionV1AgentAddressAddedIterator{contract: _AgentCollectionV1.contract, event: "AgentAddressAdded", logs: logs, sub: sub}, nil
}

// WatchAgentAddressAdded is a free log subscription operation binding the contract event 0x2f6a4339fb3581220ac44c101401de6fadea0c48ccffdb219e7cc2daf98b62b3.
//
// Solidity: event AgentAddressAdded(uint256 indexed tokenId, string addressType, string value)
func (_AgentCollectionV1 *AgentCollectionV1Filterer) WatchAgentAddressAdded(opts *bind.WatchOpts, sink chan<- *AgentCollectionV1AgentAddressAdded, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _AgentCollectionV1.contract.WatchLogs(opts, "AgentAddressAdded", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentCollectionV1AgentAddressAdded)
				if err := _AgentCollectionV1.contract.UnpackLog(event, "AgentAddressAdded", log); err != nil {
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

// ParseAgentAddressAdded is a log parse operation binding the contract event 0x2f6a4339fb3581220ac44c101401de6fadea0c48ccffdb219e7cc2daf98b62b3.
//
// Solidity: event AgentAddressAdded(uint256 indexed tokenId, string addressType, string value)
func (_AgentCollectionV1 *AgentCollectionV1Filterer) ParseAgentAddressAdded(log types.Log) (*AgentCollectionV1AgentAddressAdded, error) {
	event := new(AgentCollectionV1AgentAddressAdded)
	if err := _AgentCollectionV1.contract.UnpackLog(event, "AgentAddressAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentCollectionV1AgentAddressRemovedIterator is returned from FilterAgentAddressRemoved and is used to iterate over the raw logs and unpacked data for AgentAddressRemoved events raised by the AgentCollectionV1 contract.
type AgentCollectionV1AgentAddressRemovedIterator struct {
	Event *AgentCollectionV1AgentAddressRemoved // Event containing the contract specifics and raw log

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
func (it *AgentCollectionV1AgentAddressRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentCollectionV1AgentAddressRemoved)
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
		it.Event = new(AgentCollectionV1AgentAddressRemoved)
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
func (it *AgentCollectionV1AgentAddressRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentCollectionV1AgentAddressRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentCollectionV1AgentAddressRemoved represents a AgentAddressRemoved event raised by the AgentCollectionV1 contract.
type AgentCollectionV1AgentAddressRemoved struct {
	TokenId     *big.Int
	AddressType string
	Value       string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAgentAddressRemoved is a free log retrieval operation binding the contract event 0xcc6b0af363cb8a9d8ab09f9b188640368afc3a8860b6608c464e1cd86f04a6e3.
//
// Solidity: event AgentAddressRemoved(uint256 indexed tokenId, string addressType, string value)
func (_AgentCollectionV1 *AgentCollectionV1Filterer) FilterAgentAddressRemoved(opts *bind.FilterOpts, tokenId []*big.Int) (*AgentCollectionV1AgentAddressRemovedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _AgentCollectionV1.contract.FilterLogs(opts, "AgentAddressRemoved", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &AgentCollectionV1AgentAddressRemovedIterator{contract: _AgentCollectionV1.contract, event: "AgentAddressRemoved", logs: logs, sub: sub}, nil
}

// WatchAgentAddressRemoved is a free log subscription operation binding the contract event 0xcc6b0af363cb8a9d8ab09f9b188640368afc3a8860b6608c464e1cd86f04a6e3.
//
// Solidity: event AgentAddressRemoved(uint256 indexed tokenId, string addressType, string value)
func (_AgentCollectionV1 *AgentCollectionV1Filterer) WatchAgentAddressRemoved(opts *bind.WatchOpts, sink chan<- *AgentCollectionV1AgentAddressRemoved, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _AgentCollectionV1.contract.WatchLogs(opts, "AgentAddressRemoved", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentCollectionV1AgentAddressRemoved)
				if err := _AgentCollectionV1.contract.UnpackLog(event, "AgentAddressRemoved", log); err != nil {
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

// ParseAgentAddressRemoved is a log parse operation binding the contract event 0xcc6b0af363cb8a9d8ab09f9b188640368afc3a8860b6608c464e1cd86f04a6e3.
//
// Solidity: event AgentAddressRemoved(uint256 indexed tokenId, string addressType, string value)
func (_AgentCollectionV1 *AgentCollectionV1Filterer) ParseAgentAddressRemoved(log types.Log) (*AgentCollectionV1AgentAddressRemoved, error) {
	event := new(AgentCollectionV1AgentAddressRemoved)
	if err := _AgentCollectionV1.contract.UnpackLog(event, "AgentAddressRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentCollectionV1AgentAttributeUpdatedIterator is returned from FilterAgentAttributeUpdated and is used to iterate over the raw logs and unpacked data for AgentAttributeUpdated events raised by the AgentCollectionV1 contract.
type AgentCollectionV1AgentAttributeUpdatedIterator struct {
	Event *AgentCollectionV1AgentAttributeUpdated // Event containing the contract specifics and raw log

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
func (it *AgentCollectionV1AgentAttributeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentCollectionV1AgentAttributeUpdated)
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
		it.Event = new(AgentCollectionV1AgentAttributeUpdated)
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
func (it *AgentCollectionV1AgentAttributeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentCollectionV1AgentAttributeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentCollectionV1AgentAttributeUpdated represents a AgentAttributeUpdated event raised by the AgentCollectionV1 contract.
type AgentCollectionV1AgentAttributeUpdated struct {
	TokenId  *big.Int
	Key      common.Hash
	OldValue string
	NewValue string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAgentAttributeUpdated is a free log retrieval operation binding the contract event 0x351c3779bcc4c2572503322d51afd15a65216f375c126e5fe6c0c47949c64bb0.
//
// Solidity: event AgentAttributeUpdated(uint256 indexed tokenId, string indexed key, string oldValue, string newValue)
func (_AgentCollectionV1 *AgentCollectionV1Filterer) FilterAgentAttributeUpdated(opts *bind.FilterOpts, tokenId []*big.Int, key []string) (*AgentCollectionV1AgentAttributeUpdatedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _AgentCollectionV1.contract.FilterLogs(opts, "AgentAttributeUpdated", tokenIdRule, keyRule)
	if err != nil {
		return nil, err
	}
	return &AgentCollectionV1AgentAttributeUpdatedIterator{contract: _AgentCollectionV1.contract, event: "AgentAttributeUpdated", logs: logs, sub: sub}, nil
}

// WatchAgentAttributeUpdated is a free log subscription operation binding the contract event 0x351c3779bcc4c2572503322d51afd15a65216f375c126e5fe6c0c47949c64bb0.
//
// Solidity: event AgentAttributeUpdated(uint256 indexed tokenId, string indexed key, string oldValue, string newValue)
func (_AgentCollectionV1 *AgentCollectionV1Filterer) WatchAgentAttributeUpdated(opts *bind.WatchOpts, sink chan<- *AgentCollectionV1AgentAttributeUpdated, tokenId []*big.Int, key []string) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _AgentCollectionV1.contract.WatchLogs(opts, "AgentAttributeUpdated", tokenIdRule, keyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentCollectionV1AgentAttributeUpdated)
				if err := _AgentCollectionV1.contract.UnpackLog(event, "AgentAttributeUpdated", log); err != nil {
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

// ParseAgentAttributeUpdated is a log parse operation binding the contract event 0x351c3779bcc4c2572503322d51afd15a65216f375c126e5fe6c0c47949c64bb0.
//
// Solidity: event AgentAttributeUpdated(uint256 indexed tokenId, string indexed key, string oldValue, string newValue)
func (_AgentCollectionV1 *AgentCollectionV1Filterer) ParseAgentAttributeUpdated(log types.Log) (*AgentCollectionV1AgentAttributeUpdated, error) {
	event := new(AgentCollectionV1AgentAttributeUpdated)
	if err := _AgentCollectionV1.contract.UnpackLog(event, "AgentAttributeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentCollectionV1AgentImageProposalCreatedIterator is returned from FilterAgentImageProposalCreated and is used to iterate over the raw logs and unpacked data for AgentImageProposalCreated events raised by the AgentCollectionV1 contract.
type AgentCollectionV1AgentImageProposalCreatedIterator struct {
	Event *AgentCollectionV1AgentImageProposalCreated // Event containing the contract specifics and raw log

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
func (it *AgentCollectionV1AgentImageProposalCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentCollectionV1AgentImageProposalCreated)
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
		it.Event = new(AgentCollectionV1AgentImageProposalCreated)
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
func (it *AgentCollectionV1AgentImageProposalCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentCollectionV1AgentImageProposalCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentCollectionV1AgentImageProposalCreated represents a AgentImageProposalCreated event raised by the AgentCollectionV1 contract.
type AgentCollectionV1AgentImageProposalCreated struct {
	TokenId    *big.Int
	Key        common.Hash
	ProposalId *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAgentImageProposalCreated is a free log retrieval operation binding the contract event 0x8b06557af55313466fb2c7be6bf1db02af881cb0308214410a604185c79af52a.
//
// Solidity: event AgentImageProposalCreated(uint256 indexed tokenId, string indexed key, uint256 proposalId)
func (_AgentCollectionV1 *AgentCollectionV1Filterer) FilterAgentImageProposalCreated(opts *bind.FilterOpts, tokenId []*big.Int, key []string) (*AgentCollectionV1AgentImageProposalCreatedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _AgentCollectionV1.contract.FilterLogs(opts, "AgentImageProposalCreated", tokenIdRule, keyRule)
	if err != nil {
		return nil, err
	}
	return &AgentCollectionV1AgentImageProposalCreatedIterator{contract: _AgentCollectionV1.contract, event: "AgentImageProposalCreated", logs: logs, sub: sub}, nil
}

// WatchAgentImageProposalCreated is a free log subscription operation binding the contract event 0x8b06557af55313466fb2c7be6bf1db02af881cb0308214410a604185c79af52a.
//
// Solidity: event AgentImageProposalCreated(uint256 indexed tokenId, string indexed key, uint256 proposalId)
func (_AgentCollectionV1 *AgentCollectionV1Filterer) WatchAgentImageProposalCreated(opts *bind.WatchOpts, sink chan<- *AgentCollectionV1AgentImageProposalCreated, tokenId []*big.Int, key []string) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _AgentCollectionV1.contract.WatchLogs(opts, "AgentImageProposalCreated", tokenIdRule, keyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentCollectionV1AgentImageProposalCreated)
				if err := _AgentCollectionV1.contract.UnpackLog(event, "AgentImageProposalCreated", log); err != nil {
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

// ParseAgentImageProposalCreated is a log parse operation binding the contract event 0x8b06557af55313466fb2c7be6bf1db02af881cb0308214410a604185c79af52a.
//
// Solidity: event AgentImageProposalCreated(uint256 indexed tokenId, string indexed key, uint256 proposalId)
func (_AgentCollectionV1 *AgentCollectionV1Filterer) ParseAgentImageProposalCreated(log types.Log) (*AgentCollectionV1AgentImageProposalCreated, error) {
	event := new(AgentCollectionV1AgentImageProposalCreated)
	if err := _AgentCollectionV1.contract.UnpackLog(event, "AgentImageProposalCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentCollectionV1AgentImageProposalRemovedIterator is returned from FilterAgentImageProposalRemoved and is used to iterate over the raw logs and unpacked data for AgentImageProposalRemoved events raised by the AgentCollectionV1 contract.
type AgentCollectionV1AgentImageProposalRemovedIterator struct {
	Event *AgentCollectionV1AgentImageProposalRemoved // Event containing the contract specifics and raw log

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
func (it *AgentCollectionV1AgentImageProposalRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentCollectionV1AgentImageProposalRemoved)
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
		it.Event = new(AgentCollectionV1AgentImageProposalRemoved)
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
func (it *AgentCollectionV1AgentImageProposalRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentCollectionV1AgentImageProposalRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentCollectionV1AgentImageProposalRemoved represents a AgentImageProposalRemoved event raised by the AgentCollectionV1 contract.
type AgentCollectionV1AgentImageProposalRemoved struct {
	TokenId    *big.Int
	Key        common.Hash
	ProposalId *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAgentImageProposalRemoved is a free log retrieval operation binding the contract event 0x1b684b8ada21a0a6786e403a19ac9c0de5efdb63365e9744c92d21521a5c545f.
//
// Solidity: event AgentImageProposalRemoved(uint256 indexed tokenId, string indexed key, uint256 proposalId)
func (_AgentCollectionV1 *AgentCollectionV1Filterer) FilterAgentImageProposalRemoved(opts *bind.FilterOpts, tokenId []*big.Int, key []string) (*AgentCollectionV1AgentImageProposalRemovedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _AgentCollectionV1.contract.FilterLogs(opts, "AgentImageProposalRemoved", tokenIdRule, keyRule)
	if err != nil {
		return nil, err
	}
	return &AgentCollectionV1AgentImageProposalRemovedIterator{contract: _AgentCollectionV1.contract, event: "AgentImageProposalRemoved", logs: logs, sub: sub}, nil
}

// WatchAgentImageProposalRemoved is a free log subscription operation binding the contract event 0x1b684b8ada21a0a6786e403a19ac9c0de5efdb63365e9744c92d21521a5c545f.
//
// Solidity: event AgentImageProposalRemoved(uint256 indexed tokenId, string indexed key, uint256 proposalId)
func (_AgentCollectionV1 *AgentCollectionV1Filterer) WatchAgentImageProposalRemoved(opts *bind.WatchOpts, sink chan<- *AgentCollectionV1AgentImageProposalRemoved, tokenId []*big.Int, key []string) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _AgentCollectionV1.contract.WatchLogs(opts, "AgentImageProposalRemoved", tokenIdRule, keyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentCollectionV1AgentImageProposalRemoved)
				if err := _AgentCollectionV1.contract.UnpackLog(event, "AgentImageProposalRemoved", log); err != nil {
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

// ParseAgentImageProposalRemoved is a log parse operation binding the contract event 0x1b684b8ada21a0a6786e403a19ac9c0de5efdb63365e9744c92d21521a5c545f.
//
// Solidity: event AgentImageProposalRemoved(uint256 indexed tokenId, string indexed key, uint256 proposalId)
func (_AgentCollectionV1 *AgentCollectionV1Filterer) ParseAgentImageProposalRemoved(log types.Log) (*AgentCollectionV1AgentImageProposalRemoved, error) {
	event := new(AgentCollectionV1AgentImageProposalRemoved)
	if err := _AgentCollectionV1.contract.UnpackLog(event, "AgentImageProposalRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentCollectionV1AgentImageRemovedIterator is returned from FilterAgentImageRemoved and is used to iterate over the raw logs and unpacked data for AgentImageRemoved events raised by the AgentCollectionV1 contract.
type AgentCollectionV1AgentImageRemovedIterator struct {
	Event *AgentCollectionV1AgentImageRemoved // Event containing the contract specifics and raw log

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
func (it *AgentCollectionV1AgentImageRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentCollectionV1AgentImageRemoved)
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
		it.Event = new(AgentCollectionV1AgentImageRemoved)
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
func (it *AgentCollectionV1AgentImageRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentCollectionV1AgentImageRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentCollectionV1AgentImageRemoved represents a AgentImageRemoved event raised by the AgentCollectionV1 contract.
type AgentCollectionV1AgentImageRemoved struct {
	TokenId *big.Int
	Key     common.Hash
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAgentImageRemoved is a free log retrieval operation binding the contract event 0x6fd39f310cfda0ebdeced7ea4054b710c3e04122836c659b0f85b2a143ae4ff0.
//
// Solidity: event AgentImageRemoved(uint256 indexed tokenId, string indexed key)
func (_AgentCollectionV1 *AgentCollectionV1Filterer) FilterAgentImageRemoved(opts *bind.FilterOpts, tokenId []*big.Int, key []string) (*AgentCollectionV1AgentImageRemovedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _AgentCollectionV1.contract.FilterLogs(opts, "AgentImageRemoved", tokenIdRule, keyRule)
	if err != nil {
		return nil, err
	}
	return &AgentCollectionV1AgentImageRemovedIterator{contract: _AgentCollectionV1.contract, event: "AgentImageRemoved", logs: logs, sub: sub}, nil
}

// WatchAgentImageRemoved is a free log subscription operation binding the contract event 0x6fd39f310cfda0ebdeced7ea4054b710c3e04122836c659b0f85b2a143ae4ff0.
//
// Solidity: event AgentImageRemoved(uint256 indexed tokenId, string indexed key)
func (_AgentCollectionV1 *AgentCollectionV1Filterer) WatchAgentImageRemoved(opts *bind.WatchOpts, sink chan<- *AgentCollectionV1AgentImageRemoved, tokenId []*big.Int, key []string) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _AgentCollectionV1.contract.WatchLogs(opts, "AgentImageRemoved", tokenIdRule, keyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentCollectionV1AgentImageRemoved)
				if err := _AgentCollectionV1.contract.UnpackLog(event, "AgentImageRemoved", log); err != nil {
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

// ParseAgentImageRemoved is a log parse operation binding the contract event 0x6fd39f310cfda0ebdeced7ea4054b710c3e04122836c659b0f85b2a143ae4ff0.
//
// Solidity: event AgentImageRemoved(uint256 indexed tokenId, string indexed key)
func (_AgentCollectionV1 *AgentCollectionV1Filterer) ParseAgentImageRemoved(log types.Log) (*AgentCollectionV1AgentImageRemoved, error) {
	event := new(AgentCollectionV1AgentImageRemoved)
	if err := _AgentCollectionV1.contract.UnpackLog(event, "AgentImageRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentCollectionV1AgentImageUpdatedIterator is returned from FilterAgentImageUpdated and is used to iterate over the raw logs and unpacked data for AgentImageUpdated events raised by the AgentCollectionV1 contract.
type AgentCollectionV1AgentImageUpdatedIterator struct {
	Event *AgentCollectionV1AgentImageUpdated // Event containing the contract specifics and raw log

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
func (it *AgentCollectionV1AgentImageUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentCollectionV1AgentImageUpdated)
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
		it.Event = new(AgentCollectionV1AgentImageUpdated)
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
func (it *AgentCollectionV1AgentImageUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentCollectionV1AgentImageUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentCollectionV1AgentImageUpdated represents a AgentImageUpdated event raised by the AgentCollectionV1 contract.
type AgentCollectionV1AgentImageUpdated struct {
	TokenId  *big.Int
	Key      common.Hash
	OldValue string
	NewValue string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAgentImageUpdated is a free log retrieval operation binding the contract event 0x102c9e6462b7d4e559e582578ea590c7f08c5d0551bd6c2d047768a4b1377e65.
//
// Solidity: event AgentImageUpdated(uint256 indexed tokenId, string indexed key, string oldValue, string newValue)
func (_AgentCollectionV1 *AgentCollectionV1Filterer) FilterAgentImageUpdated(opts *bind.FilterOpts, tokenId []*big.Int, key []string) (*AgentCollectionV1AgentImageUpdatedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _AgentCollectionV1.contract.FilterLogs(opts, "AgentImageUpdated", tokenIdRule, keyRule)
	if err != nil {
		return nil, err
	}
	return &AgentCollectionV1AgentImageUpdatedIterator{contract: _AgentCollectionV1.contract, event: "AgentImageUpdated", logs: logs, sub: sub}, nil
}

// WatchAgentImageUpdated is a free log subscription operation binding the contract event 0x102c9e6462b7d4e559e582578ea590c7f08c5d0551bd6c2d047768a4b1377e65.
//
// Solidity: event AgentImageUpdated(uint256 indexed tokenId, string indexed key, string oldValue, string newValue)
func (_AgentCollectionV1 *AgentCollectionV1Filterer) WatchAgentImageUpdated(opts *bind.WatchOpts, sink chan<- *AgentCollectionV1AgentImageUpdated, tokenId []*big.Int, key []string) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _AgentCollectionV1.contract.WatchLogs(opts, "AgentImageUpdated", tokenIdRule, keyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentCollectionV1AgentImageUpdated)
				if err := _AgentCollectionV1.contract.UnpackLog(event, "AgentImageUpdated", log); err != nil {
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

// ParseAgentImageUpdated is a log parse operation binding the contract event 0x102c9e6462b7d4e559e582578ea590c7f08c5d0551bd6c2d047768a4b1377e65.
//
// Solidity: event AgentImageUpdated(uint256 indexed tokenId, string indexed key, string oldValue, string newValue)
func (_AgentCollectionV1 *AgentCollectionV1Filterer) ParseAgentImageUpdated(log types.Log) (*AgentCollectionV1AgentImageUpdated, error) {
	event := new(AgentCollectionV1AgentImageUpdated)
	if err := _AgentCollectionV1.contract.UnpackLog(event, "AgentImageUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentCollectionV1ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the AgentCollectionV1 contract.
type AgentCollectionV1ApprovalIterator struct {
	Event *AgentCollectionV1Approval // Event containing the contract specifics and raw log

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
func (it *AgentCollectionV1ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentCollectionV1Approval)
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
		it.Event = new(AgentCollectionV1Approval)
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
func (it *AgentCollectionV1ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentCollectionV1ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentCollectionV1Approval represents a Approval event raised by the AgentCollectionV1 contract.
type AgentCollectionV1Approval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_AgentCollectionV1 *AgentCollectionV1Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*AgentCollectionV1ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _AgentCollectionV1.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &AgentCollectionV1ApprovalIterator{contract: _AgentCollectionV1.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_AgentCollectionV1 *AgentCollectionV1Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *AgentCollectionV1Approval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _AgentCollectionV1.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentCollectionV1Approval)
				if err := _AgentCollectionV1.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_AgentCollectionV1 *AgentCollectionV1Filterer) ParseApproval(log types.Log) (*AgentCollectionV1Approval, error) {
	event := new(AgentCollectionV1Approval)
	if err := _AgentCollectionV1.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentCollectionV1ApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the AgentCollectionV1 contract.
type AgentCollectionV1ApprovalForAllIterator struct {
	Event *AgentCollectionV1ApprovalForAll // Event containing the contract specifics and raw log

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
func (it *AgentCollectionV1ApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentCollectionV1ApprovalForAll)
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
		it.Event = new(AgentCollectionV1ApprovalForAll)
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
func (it *AgentCollectionV1ApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentCollectionV1ApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentCollectionV1ApprovalForAll represents a ApprovalForAll event raised by the AgentCollectionV1 contract.
type AgentCollectionV1ApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_AgentCollectionV1 *AgentCollectionV1Filterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*AgentCollectionV1ApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _AgentCollectionV1.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &AgentCollectionV1ApprovalForAllIterator{contract: _AgentCollectionV1.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_AgentCollectionV1 *AgentCollectionV1Filterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *AgentCollectionV1ApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _AgentCollectionV1.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentCollectionV1ApprovalForAll)
				if err := _AgentCollectionV1.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_AgentCollectionV1 *AgentCollectionV1Filterer) ParseApprovalForAll(log types.Log) (*AgentCollectionV1ApprovalForAll, error) {
	event := new(AgentCollectionV1ApprovalForAll)
	if err := _AgentCollectionV1.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentCollectionV1CollectionsDescriptorUpdatedIterator is returned from FilterCollectionsDescriptorUpdated and is used to iterate over the raw logs and unpacked data for CollectionsDescriptorUpdated events raised by the AgentCollectionV1 contract.
type AgentCollectionV1CollectionsDescriptorUpdatedIterator struct {
	Event *AgentCollectionV1CollectionsDescriptorUpdated // Event containing the contract specifics and raw log

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
func (it *AgentCollectionV1CollectionsDescriptorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentCollectionV1CollectionsDescriptorUpdated)
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
		it.Event = new(AgentCollectionV1CollectionsDescriptorUpdated)
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
func (it *AgentCollectionV1CollectionsDescriptorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentCollectionV1CollectionsDescriptorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentCollectionV1CollectionsDescriptorUpdated represents a CollectionsDescriptorUpdated event raised by the AgentCollectionV1 contract.
type AgentCollectionV1CollectionsDescriptorUpdated struct {
	OldCollectionsDescriptor common.Address
	NewCollectionsDescriptor common.Address
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterCollectionsDescriptorUpdated is a free log retrieval operation binding the contract event 0xdaff73e0662ecb070f01270cb2aa647158f64dccd9df9b2f36a4de72d2e18d1e.
//
// Solidity: event CollectionsDescriptorUpdated(address indexed oldCollectionsDescriptor, address indexed newCollectionsDescriptor)
func (_AgentCollectionV1 *AgentCollectionV1Filterer) FilterCollectionsDescriptorUpdated(opts *bind.FilterOpts, oldCollectionsDescriptor []common.Address, newCollectionsDescriptor []common.Address) (*AgentCollectionV1CollectionsDescriptorUpdatedIterator, error) {

	var oldCollectionsDescriptorRule []interface{}
	for _, oldCollectionsDescriptorItem := range oldCollectionsDescriptor {
		oldCollectionsDescriptorRule = append(oldCollectionsDescriptorRule, oldCollectionsDescriptorItem)
	}
	var newCollectionsDescriptorRule []interface{}
	for _, newCollectionsDescriptorItem := range newCollectionsDescriptor {
		newCollectionsDescriptorRule = append(newCollectionsDescriptorRule, newCollectionsDescriptorItem)
	}

	logs, sub, err := _AgentCollectionV1.contract.FilterLogs(opts, "CollectionsDescriptorUpdated", oldCollectionsDescriptorRule, newCollectionsDescriptorRule)
	if err != nil {
		return nil, err
	}
	return &AgentCollectionV1CollectionsDescriptorUpdatedIterator{contract: _AgentCollectionV1.contract, event: "CollectionsDescriptorUpdated", logs: logs, sub: sub}, nil
}

// WatchCollectionsDescriptorUpdated is a free log subscription operation binding the contract event 0xdaff73e0662ecb070f01270cb2aa647158f64dccd9df9b2f36a4de72d2e18d1e.
//
// Solidity: event CollectionsDescriptorUpdated(address indexed oldCollectionsDescriptor, address indexed newCollectionsDescriptor)
func (_AgentCollectionV1 *AgentCollectionV1Filterer) WatchCollectionsDescriptorUpdated(opts *bind.WatchOpts, sink chan<- *AgentCollectionV1CollectionsDescriptorUpdated, oldCollectionsDescriptor []common.Address, newCollectionsDescriptor []common.Address) (event.Subscription, error) {

	var oldCollectionsDescriptorRule []interface{}
	for _, oldCollectionsDescriptorItem := range oldCollectionsDescriptor {
		oldCollectionsDescriptorRule = append(oldCollectionsDescriptorRule, oldCollectionsDescriptorItem)
	}
	var newCollectionsDescriptorRule []interface{}
	for _, newCollectionsDescriptorItem := range newCollectionsDescriptor {
		newCollectionsDescriptorRule = append(newCollectionsDescriptorRule, newCollectionsDescriptorItem)
	}

	logs, sub, err := _AgentCollectionV1.contract.WatchLogs(opts, "CollectionsDescriptorUpdated", oldCollectionsDescriptorRule, newCollectionsDescriptorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentCollectionV1CollectionsDescriptorUpdated)
				if err := _AgentCollectionV1.contract.UnpackLog(event, "CollectionsDescriptorUpdated", log); err != nil {
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

// ParseCollectionsDescriptorUpdated is a log parse operation binding the contract event 0xdaff73e0662ecb070f01270cb2aa647158f64dccd9df9b2f36a4de72d2e18d1e.
//
// Solidity: event CollectionsDescriptorUpdated(address indexed oldCollectionsDescriptor, address indexed newCollectionsDescriptor)
func (_AgentCollectionV1 *AgentCollectionV1Filterer) ParseCollectionsDescriptorUpdated(log types.Log) (*AgentCollectionV1CollectionsDescriptorUpdated, error) {
	event := new(AgentCollectionV1CollectionsDescriptorUpdated)
	if err := _AgentCollectionV1.contract.UnpackLog(event, "CollectionsDescriptorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentCollectionV1DescriptionUpdatedIterator is returned from FilterDescriptionUpdated and is used to iterate over the raw logs and unpacked data for DescriptionUpdated events raised by the AgentCollectionV1 contract.
type AgentCollectionV1DescriptionUpdatedIterator struct {
	Event *AgentCollectionV1DescriptionUpdated // Event containing the contract specifics and raw log

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
func (it *AgentCollectionV1DescriptionUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentCollectionV1DescriptionUpdated)
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
		it.Event = new(AgentCollectionV1DescriptionUpdated)
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
func (it *AgentCollectionV1DescriptionUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentCollectionV1DescriptionUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentCollectionV1DescriptionUpdated represents a DescriptionUpdated event raised by the AgentCollectionV1 contract.
type AgentCollectionV1DescriptionUpdated struct {
	OldDescription string
	NewDescription string
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDescriptionUpdated is a free log retrieval operation binding the contract event 0xe21432e1fe2b572d5803dd7316b7a854952317b42017f920a616ec70cdb8a5c1.
//
// Solidity: event DescriptionUpdated(string oldDescription, string newDescription)
func (_AgentCollectionV1 *AgentCollectionV1Filterer) FilterDescriptionUpdated(opts *bind.FilterOpts) (*AgentCollectionV1DescriptionUpdatedIterator, error) {

	logs, sub, err := _AgentCollectionV1.contract.FilterLogs(opts, "DescriptionUpdated")
	if err != nil {
		return nil, err
	}
	return &AgentCollectionV1DescriptionUpdatedIterator{contract: _AgentCollectionV1.contract, event: "DescriptionUpdated", logs: logs, sub: sub}, nil
}

// WatchDescriptionUpdated is a free log subscription operation binding the contract event 0xe21432e1fe2b572d5803dd7316b7a854952317b42017f920a616ec70cdb8a5c1.
//
// Solidity: event DescriptionUpdated(string oldDescription, string newDescription)
func (_AgentCollectionV1 *AgentCollectionV1Filterer) WatchDescriptionUpdated(opts *bind.WatchOpts, sink chan<- *AgentCollectionV1DescriptionUpdated) (event.Subscription, error) {

	logs, sub, err := _AgentCollectionV1.contract.WatchLogs(opts, "DescriptionUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentCollectionV1DescriptionUpdated)
				if err := _AgentCollectionV1.contract.UnpackLog(event, "DescriptionUpdated", log); err != nil {
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

// ParseDescriptionUpdated is a log parse operation binding the contract event 0xe21432e1fe2b572d5803dd7316b7a854952317b42017f920a616ec70cdb8a5c1.
//
// Solidity: event DescriptionUpdated(string oldDescription, string newDescription)
func (_AgentCollectionV1 *AgentCollectionV1Filterer) ParseDescriptionUpdated(log types.Log) (*AgentCollectionV1DescriptionUpdated, error) {
	event := new(AgentCollectionV1DescriptionUpdated)
	if err := _AgentCollectionV1.contract.UnpackLog(event, "DescriptionUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentCollectionV1InitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the AgentCollectionV1 contract.
type AgentCollectionV1InitializedIterator struct {
	Event *AgentCollectionV1Initialized // Event containing the contract specifics and raw log

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
func (it *AgentCollectionV1InitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentCollectionV1Initialized)
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
		it.Event = new(AgentCollectionV1Initialized)
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
func (it *AgentCollectionV1InitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentCollectionV1InitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentCollectionV1Initialized represents a Initialized event raised by the AgentCollectionV1 contract.
type AgentCollectionV1Initialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_AgentCollectionV1 *AgentCollectionV1Filterer) FilterInitialized(opts *bind.FilterOpts) (*AgentCollectionV1InitializedIterator, error) {

	logs, sub, err := _AgentCollectionV1.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &AgentCollectionV1InitializedIterator{contract: _AgentCollectionV1.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_AgentCollectionV1 *AgentCollectionV1Filterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *AgentCollectionV1Initialized) (event.Subscription, error) {

	logs, sub, err := _AgentCollectionV1.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentCollectionV1Initialized)
				if err := _AgentCollectionV1.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_AgentCollectionV1 *AgentCollectionV1Filterer) ParseInitialized(log types.Log) (*AgentCollectionV1Initialized, error) {
	event := new(AgentCollectionV1Initialized)
	if err := _AgentCollectionV1.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentCollectionV1MintProposalCreatedIterator is returned from FilterMintProposalCreated and is used to iterate over the raw logs and unpacked data for MintProposalCreated events raised by the AgentCollectionV1 contract.
type AgentCollectionV1MintProposalCreatedIterator struct {
	Event *AgentCollectionV1MintProposalCreated // Event containing the contract specifics and raw log

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
func (it *AgentCollectionV1MintProposalCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentCollectionV1MintProposalCreated)
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
		it.Event = new(AgentCollectionV1MintProposalCreated)
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
func (it *AgentCollectionV1MintProposalCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentCollectionV1MintProposalCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentCollectionV1MintProposalCreated represents a MintProposalCreated event raised by the AgentCollectionV1 contract.
type AgentCollectionV1MintProposalCreated struct {
	From       common.Address
	ProposalId *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMintProposalCreated is a free log retrieval operation binding the contract event 0x0116c43c83acf6cd8597ce110e50b139e8c85a1544440913be14f9043a244361.
//
// Solidity: event MintProposalCreated(address indexed from, uint256 proposalId)
func (_AgentCollectionV1 *AgentCollectionV1Filterer) FilterMintProposalCreated(opts *bind.FilterOpts, from []common.Address) (*AgentCollectionV1MintProposalCreatedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _AgentCollectionV1.contract.FilterLogs(opts, "MintProposalCreated", fromRule)
	if err != nil {
		return nil, err
	}
	return &AgentCollectionV1MintProposalCreatedIterator{contract: _AgentCollectionV1.contract, event: "MintProposalCreated", logs: logs, sub: sub}, nil
}

// WatchMintProposalCreated is a free log subscription operation binding the contract event 0x0116c43c83acf6cd8597ce110e50b139e8c85a1544440913be14f9043a244361.
//
// Solidity: event MintProposalCreated(address indexed from, uint256 proposalId)
func (_AgentCollectionV1 *AgentCollectionV1Filterer) WatchMintProposalCreated(opts *bind.WatchOpts, sink chan<- *AgentCollectionV1MintProposalCreated, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _AgentCollectionV1.contract.WatchLogs(opts, "MintProposalCreated", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentCollectionV1MintProposalCreated)
				if err := _AgentCollectionV1.contract.UnpackLog(event, "MintProposalCreated", log); err != nil {
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

// ParseMintProposalCreated is a log parse operation binding the contract event 0x0116c43c83acf6cd8597ce110e50b139e8c85a1544440913be14f9043a244361.
//
// Solidity: event MintProposalCreated(address indexed from, uint256 proposalId)
func (_AgentCollectionV1 *AgentCollectionV1Filterer) ParseMintProposalCreated(log types.Log) (*AgentCollectionV1MintProposalCreated, error) {
	event := new(AgentCollectionV1MintProposalCreated)
	if err := _AgentCollectionV1.contract.UnpackLog(event, "MintProposalCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentCollectionV1MintProposalRemovedIterator is returned from FilterMintProposalRemoved and is used to iterate over the raw logs and unpacked data for MintProposalRemoved events raised by the AgentCollectionV1 contract.
type AgentCollectionV1MintProposalRemovedIterator struct {
	Event *AgentCollectionV1MintProposalRemoved // Event containing the contract specifics and raw log

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
func (it *AgentCollectionV1MintProposalRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentCollectionV1MintProposalRemoved)
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
		it.Event = new(AgentCollectionV1MintProposalRemoved)
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
func (it *AgentCollectionV1MintProposalRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentCollectionV1MintProposalRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentCollectionV1MintProposalRemoved represents a MintProposalRemoved event raised by the AgentCollectionV1 contract.
type AgentCollectionV1MintProposalRemoved struct {
	Owner      common.Address
	ProposalId *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMintProposalRemoved is a free log retrieval operation binding the contract event 0x632209d0b827b0b37a988c9af91bfd7e849809b379e1efa66dc47082953b8700.
//
// Solidity: event MintProposalRemoved(address indexed owner, uint256 proposalId)
func (_AgentCollectionV1 *AgentCollectionV1Filterer) FilterMintProposalRemoved(opts *bind.FilterOpts, owner []common.Address) (*AgentCollectionV1MintProposalRemovedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _AgentCollectionV1.contract.FilterLogs(opts, "MintProposalRemoved", ownerRule)
	if err != nil {
		return nil, err
	}
	return &AgentCollectionV1MintProposalRemovedIterator{contract: _AgentCollectionV1.contract, event: "MintProposalRemoved", logs: logs, sub: sub}, nil
}

// WatchMintProposalRemoved is a free log subscription operation binding the contract event 0x632209d0b827b0b37a988c9af91bfd7e849809b379e1efa66dc47082953b8700.
//
// Solidity: event MintProposalRemoved(address indexed owner, uint256 proposalId)
func (_AgentCollectionV1 *AgentCollectionV1Filterer) WatchMintProposalRemoved(opts *bind.WatchOpts, sink chan<- *AgentCollectionV1MintProposalRemoved, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _AgentCollectionV1.contract.WatchLogs(opts, "MintProposalRemoved", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentCollectionV1MintProposalRemoved)
				if err := _AgentCollectionV1.contract.UnpackLog(event, "MintProposalRemoved", log); err != nil {
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

// ParseMintProposalRemoved is a log parse operation binding the contract event 0x632209d0b827b0b37a988c9af91bfd7e849809b379e1efa66dc47082953b8700.
//
// Solidity: event MintProposalRemoved(address indexed owner, uint256 proposalId)
func (_AgentCollectionV1 *AgentCollectionV1Filterer) ParseMintProposalRemoved(log types.Log) (*AgentCollectionV1MintProposalRemoved, error) {
	event := new(AgentCollectionV1MintProposalRemoved)
	if err := _AgentCollectionV1.contract.UnpackLog(event, "MintProposalRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentCollectionV1MintedIterator is returned from FilterMinted and is used to iterate over the raw logs and unpacked data for Minted events raised by the AgentCollectionV1 contract.
type AgentCollectionV1MintedIterator struct {
	Event *AgentCollectionV1Minted // Event containing the contract specifics and raw log

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
func (it *AgentCollectionV1MintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentCollectionV1Minted)
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
		it.Event = new(AgentCollectionV1Minted)
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
func (it *AgentCollectionV1MintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentCollectionV1MintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentCollectionV1Minted represents a Minted event raised by the AgentCollectionV1 contract.
type AgentCollectionV1Minted struct {
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMinted is a free log retrieval operation binding the contract event 0x30385c845b448a36257a6a1716e6ad2e1bc2cbe333cde1e69fe849ad6511adfe.
//
// Solidity: event Minted(address indexed to, uint256 tokenId)
func (_AgentCollectionV1 *AgentCollectionV1Filterer) FilterMinted(opts *bind.FilterOpts, to []common.Address) (*AgentCollectionV1MintedIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AgentCollectionV1.contract.FilterLogs(opts, "Minted", toRule)
	if err != nil {
		return nil, err
	}
	return &AgentCollectionV1MintedIterator{contract: _AgentCollectionV1.contract, event: "Minted", logs: logs, sub: sub}, nil
}

// WatchMinted is a free log subscription operation binding the contract event 0x30385c845b448a36257a6a1716e6ad2e1bc2cbe333cde1e69fe849ad6511adfe.
//
// Solidity: event Minted(address indexed to, uint256 tokenId)
func (_AgentCollectionV1 *AgentCollectionV1Filterer) WatchMinted(opts *bind.WatchOpts, sink chan<- *AgentCollectionV1Minted, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AgentCollectionV1.contract.WatchLogs(opts, "Minted", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentCollectionV1Minted)
				if err := _AgentCollectionV1.contract.UnpackLog(event, "Minted", log); err != nil {
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

// ParseMinted is a log parse operation binding the contract event 0x30385c845b448a36257a6a1716e6ad2e1bc2cbe333cde1e69fe849ad6511adfe.
//
// Solidity: event Minted(address indexed to, uint256 tokenId)
func (_AgentCollectionV1 *AgentCollectionV1Filterer) ParseMinted(log types.Log) (*AgentCollectionV1Minted, error) {
	event := new(AgentCollectionV1Minted)
	if err := _AgentCollectionV1.contract.UnpackLog(event, "Minted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentCollectionV1RoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the AgentCollectionV1 contract.
type AgentCollectionV1RoleAdminChangedIterator struct {
	Event *AgentCollectionV1RoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *AgentCollectionV1RoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentCollectionV1RoleAdminChanged)
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
		it.Event = new(AgentCollectionV1RoleAdminChanged)
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
func (it *AgentCollectionV1RoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentCollectionV1RoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentCollectionV1RoleAdminChanged represents a RoleAdminChanged event raised by the AgentCollectionV1 contract.
type AgentCollectionV1RoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_AgentCollectionV1 *AgentCollectionV1Filterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*AgentCollectionV1RoleAdminChangedIterator, error) {

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

	logs, sub, err := _AgentCollectionV1.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &AgentCollectionV1RoleAdminChangedIterator{contract: _AgentCollectionV1.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_AgentCollectionV1 *AgentCollectionV1Filterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *AgentCollectionV1RoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _AgentCollectionV1.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentCollectionV1RoleAdminChanged)
				if err := _AgentCollectionV1.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_AgentCollectionV1 *AgentCollectionV1Filterer) ParseRoleAdminChanged(log types.Log) (*AgentCollectionV1RoleAdminChanged, error) {
	event := new(AgentCollectionV1RoleAdminChanged)
	if err := _AgentCollectionV1.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentCollectionV1RoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the AgentCollectionV1 contract.
type AgentCollectionV1RoleGrantedIterator struct {
	Event *AgentCollectionV1RoleGranted // Event containing the contract specifics and raw log

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
func (it *AgentCollectionV1RoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentCollectionV1RoleGranted)
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
		it.Event = new(AgentCollectionV1RoleGranted)
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
func (it *AgentCollectionV1RoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentCollectionV1RoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentCollectionV1RoleGranted represents a RoleGranted event raised by the AgentCollectionV1 contract.
type AgentCollectionV1RoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AgentCollectionV1 *AgentCollectionV1Filterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AgentCollectionV1RoleGrantedIterator, error) {

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

	logs, sub, err := _AgentCollectionV1.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AgentCollectionV1RoleGrantedIterator{contract: _AgentCollectionV1.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AgentCollectionV1 *AgentCollectionV1Filterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *AgentCollectionV1RoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _AgentCollectionV1.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentCollectionV1RoleGranted)
				if err := _AgentCollectionV1.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_AgentCollectionV1 *AgentCollectionV1Filterer) ParseRoleGranted(log types.Log) (*AgentCollectionV1RoleGranted, error) {
	event := new(AgentCollectionV1RoleGranted)
	if err := _AgentCollectionV1.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentCollectionV1RoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the AgentCollectionV1 contract.
type AgentCollectionV1RoleRevokedIterator struct {
	Event *AgentCollectionV1RoleRevoked // Event containing the contract specifics and raw log

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
func (it *AgentCollectionV1RoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentCollectionV1RoleRevoked)
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
		it.Event = new(AgentCollectionV1RoleRevoked)
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
func (it *AgentCollectionV1RoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentCollectionV1RoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentCollectionV1RoleRevoked represents a RoleRevoked event raised by the AgentCollectionV1 contract.
type AgentCollectionV1RoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AgentCollectionV1 *AgentCollectionV1Filterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AgentCollectionV1RoleRevokedIterator, error) {

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

	logs, sub, err := _AgentCollectionV1.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AgentCollectionV1RoleRevokedIterator{contract: _AgentCollectionV1.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AgentCollectionV1 *AgentCollectionV1Filterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *AgentCollectionV1RoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _AgentCollectionV1.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentCollectionV1RoleRevoked)
				if err := _AgentCollectionV1.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_AgentCollectionV1 *AgentCollectionV1Filterer) ParseRoleRevoked(log types.Log) (*AgentCollectionV1RoleRevoked, error) {
	event := new(AgentCollectionV1RoleRevoked)
	if err := _AgentCollectionV1.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentCollectionV1TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the AgentCollectionV1 contract.
type AgentCollectionV1TransferIterator struct {
	Event *AgentCollectionV1Transfer // Event containing the contract specifics and raw log

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
func (it *AgentCollectionV1TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentCollectionV1Transfer)
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
		it.Event = new(AgentCollectionV1Transfer)
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
func (it *AgentCollectionV1TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentCollectionV1TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentCollectionV1Transfer represents a Transfer event raised by the AgentCollectionV1 contract.
type AgentCollectionV1Transfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_AgentCollectionV1 *AgentCollectionV1Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*AgentCollectionV1TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _AgentCollectionV1.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &AgentCollectionV1TransferIterator{contract: _AgentCollectionV1.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_AgentCollectionV1 *AgentCollectionV1Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *AgentCollectionV1Transfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _AgentCollectionV1.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentCollectionV1Transfer)
				if err := _AgentCollectionV1.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_AgentCollectionV1 *AgentCollectionV1Filterer) ParseTransfer(log types.Log) (*AgentCollectionV1Transfer, error) {
	event := new(AgentCollectionV1Transfer)
	if err := _AgentCollectionV1.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// --- RawLog() methods added by go generate ---
func (e AgentCollectionV1AgentAddressAdded) RawLog() types.Log            { return e.Raw }
func (e AgentCollectionV1AgentAddressRemoved) RawLog() types.Log          { return e.Raw }
func (e AgentCollectionV1AgentAttributeUpdated) RawLog() types.Log        { return e.Raw }
func (e AgentCollectionV1AgentImageProposalCreated) RawLog() types.Log    { return e.Raw }
func (e AgentCollectionV1AgentImageProposalRemoved) RawLog() types.Log    { return e.Raw }
func (e AgentCollectionV1AgentImageRemoved) RawLog() types.Log            { return e.Raw }
func (e AgentCollectionV1AgentImageUpdated) RawLog() types.Log            { return e.Raw }
func (e AgentCollectionV1Approval) RawLog() types.Log                     { return e.Raw }
func (e AgentCollectionV1ApprovalForAll) RawLog() types.Log               { return e.Raw }
func (e AgentCollectionV1CollectionsDescriptorUpdated) RawLog() types.Log { return e.Raw }
func (e AgentCollectionV1DescriptionUpdated) RawLog() types.Log           { return e.Raw }
func (e AgentCollectionV1Initialized) RawLog() types.Log                  { return e.Raw }
func (e AgentCollectionV1MintProposalCreated) RawLog() types.Log          { return e.Raw }
func (e AgentCollectionV1MintProposalRemoved) RawLog() types.Log          { return e.Raw }
func (e AgentCollectionV1Minted) RawLog() types.Log                       { return e.Raw }
func (e AgentCollectionV1RoleAdminChanged) RawLog() types.Log             { return e.Raw }
func (e AgentCollectionV1RoleGranted) RawLog() types.Log                  { return e.Raw }
func (e AgentCollectionV1RoleRevoked) RawLog() types.Log                  { return e.Raw }
func (e AgentCollectionV1Transfer) RawLog() types.Log                     { return e.Raw }
