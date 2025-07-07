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

// AddOrUpdateImageProposal is an auto generated low-level Go binding around an user-defined struct.
type AddOrUpdateImageProposal struct {
	Id      *big.Int
	TokenId *big.Int
	Image   KeyValue
}

// AgentInformation is an auto generated low-level Go binding around an user-defined struct.
type AgentInformation struct {
	Name    string
	Role    string
	Signer  common.Address
	Mission string
	Creator common.Address
	Wallets []common.Address
	Images  []KeyValue
}

// KeyValue is an auto generated low-level Go binding around an user-defined struct.
type KeyValue struct {
	Key   string
	Value string
}

// MintProposal is an auto generated low-level Go binding around an user-defined struct.
type MintProposal struct {
	Id      *big.Int
	Name    string
	Role    string
	Owner   common.Address
	Signer  common.Address
	Mission string
	Creator common.Address
	Wallets []common.Address
	Images  []KeyValue
}

// AgentCollectionV1MetaData contains all meta data concerning the AgentCollectionV1 contract.
var AgentCollectionV1MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721IncorrectOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721InsufficientApproval\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC721InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721NonexistentToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"InvalidName\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"expectedSigner\",\"type\":\"address\"}],\"name\":\"InvalidSigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MissingDefaultImage\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoModerators\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"NotModerator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"NotSigner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"SignerUnchanged\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"UsedName\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nonce\",\"type\":\"bytes32\"}],\"name\":\"UsedNonce\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"AgentImageProposalCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"AgentImageProposalRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"oldValue\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newValue\",\"type\":\"string\"}],\"name\":\"AgentImageUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"oldMission\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newMission\",\"type\":\"string\"}],\"name\":\"AgentMissionUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldSigner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newSigner\",\"type\":\"address\"}],\"name\":\"AgentSignerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldCollectionsDescriptor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newCollectionsDescriptor\",\"type\":\"address\"}],\"name\":\"CollectionsDescriptorUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"oldDescription\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newDescription\",\"type\":\"string\"}],\"name\":\"DescriptionUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"MintProposalCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"MintProposalRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Minted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MODERATOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"internalType\":\"structKeyValue\",\"name\":\"image\",\"type\":\"tuple\"}],\"name\":\"addOrUpdateImage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"addOrUpdateImageFromProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalIndex\",\"type\":\"uint256\"}],\"name\":\"addOrUpdateImageProposal\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"internalType\":\"structKeyValue\",\"name\":\"image\",\"type\":\"tuple\"}],\"internalType\":\"structAddOrUpdateImageProposal\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"addOrUpdateImageProposalsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"collectionsDescriptor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"collectionsManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"internalType\":\"structKeyValue\",\"name\":\"image\",\"type\":\"tuple\"}],\"name\":\"createAddOrUpdateImageProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"agentAddresses\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"role\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"mission\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"internalType\":\"structKeyValue[]\",\"name\":\"images\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"nonce\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"createMintProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"creatorOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"description\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"imagesOf\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"internalType\":\"structKeyValue[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"informationOf\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"role\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"mission\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"wallets\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"internalType\":\"structKeyValue[]\",\"name\":\"images\",\"type\":\"tuple[]\"}],\"internalType\":\"structAgentInformation\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"initialDescription\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"moderator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"initialCollectionsManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"initialCollectionsDescriptor\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isModerator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nonce\",\"type\":\"bytes32\"}],\"name\":\"isNonceUsed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"agentAddresses\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"role\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"mission\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"internalType\":\"structKeyValue[]\",\"name\":\"images\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"nonce\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"mintFromProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalIndex\",\"type\":\"uint256\"}],\"name\":\"mintProposal\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"role\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"mission\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"wallets\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"internalType\":\"structKeyValue[]\",\"name\":\"images\",\"type\":\"tuple[]\"}],\"internalType\":\"structMintProposal\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintProposalsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"missionOf\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"moderatorCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"nameOf\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"refuseAddOrUpdateImageProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"refuseMintProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"removeImage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"roleOf\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"signerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"tokenIdOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"newMission\",\"type\":\"string\"}],\"name\":\"updateAgentMission\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"newSigner\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"nonce\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"updateAgentSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newCollectionsDescriptor\",\"type\":\"address\"}],\"name\":\"updateCollectionsDescriptor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newDescription\",\"type\":\"string\"}],\"name\":\"updateDescription\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"walletsOf\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080806040523460d2577ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005460ff8160401c1660c1576002600160401b03196001600160401b03821601605c575b60405161588190816100d88239f35b6001600160401b0319166001600160401b039081177ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005581527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a13880604d565b63f92ee8a960e01b60005260046000fd5b600080fdfe608080604052600436101561001357600080fd5b60003560e01c90816301ffc9a71461297e5750806302d5146314612786578063051a26641461275057806306fdde03146126a0578063081812fc14612678578063095ea7b3146125865780630ee7e60b146125685780631164e719146124345780631278fc8b146123e757806323b872dd146123d0578063248a9ca3146123aa5780632f2ff15d1461237957806334b481e71461235b57806336568abe1461231557806337c11b2a146122ec578063419f5aa01461226b57806342842e0e1461223b57806350cffb68146120b05780635161fdf51461206f578063589a17431461202e5780636352211e14611ffe57806367c1e20b14611fda5780636d21b48414611d2557806370a0823114611cd35780637284e41614611cb757806372b1d8cf14611c7e57806375794a3c14611c60578063797669c914611c37578063798931a314611bc65780638364867114611b7957806384b0196e14611a4e57806391d14854146119f457806395d89b411461191f578063a217fddf14611903578063a22cb46514611867578063a398d8191461180b578063ab1fdaef146117e2578063af90754514611795578063b31fcdbf14611771578063b88d4fde14611716578063b90665e5146116e5578063baf3dfd31461126b578063bb046abc1461121e578063c088c9a7146110f2578063c87b56dd14610fe5578063d0c0fcdf14610f5d578063d10321bd146106fb578063d547741f146106c5578063e6f1d66a146105df578063e735b48a1461049b578063e8c43e7c14610396578063e985e9c514610348578063f53957ad1461032a578063fa6f3936146102da5763fede22631461027c57600080fd5b346102d55760203660031901126102d557610295613a86565b60043560005260076020526102d36040600020546102ce6102be6102b8836131e3565b5061371a565b6040602082015191015190614c81565b613c70565b005b600080fd5b346102d55760203660031901126102d5576102f3612bde565b6001600160a01b031660009081526000805160206155ac833981519152602090815260409182902054915160ff9092161515825290f35b346102d55760003660031901126102d5576020600454604051908152f35b346102d55760403660031901126102d557610361612bde565b61037261036c612bf4565b91613803565b9060018060a01b0316600052602052602060ff604060002054166040519015158152f35b346102d5576103a436612c0a565b906103ae81613a0b565b506103b881613a41565b6001600160a01b031633141580610474575b61045c5760005260026020526103e860066040600020019182614d04565b5080546103f481613006565b916104026040519384612a7f565b818352602083019060005260206000206000915b83831061043e5761042685614de6565b1561042d57005b634d41f25760e11b60005260046000fd5b6002602060019261044e856134f0565b815201920192019190610416565b635c3c03c360e01b6000523360045260245260446000fd5b503360009081526000805160206155ac833981519152602052604090205460ff16156103ca565b346102d55760203660031901126102d5576004356001600160401b0381116102d5576104cb903690600401612abd565b6104d3613a86565b6104db612d6d565b815191906001600160401b0383116105c957610501836104fc600854612d33565b61329d565b6020601f841160011461055457836000805160206155cc8339815191529461053192600091610549575b50613311565b6008555b610544604051928392836134cb565b0390a1005b90508401518661052b565b601f198416906008600052806000209160005b8181106105b157509185916000805160206155cc8339815191529660019410610598575b5050811b01600855610535565b85015160001960f88460031b161c19169055858061058b565b91926020600181928689015181550194019201610567565b634e487b7160e01b600052604160045260246000fd5b346102d55760203660031901126102d5576105f8613a86565b600435600052600560205260406000205461061b6106158261383c565b50613958565b90602082019161063b602084518160405193828580945193849201612b96565b810160018152030190205461069a57606081015160e082015160808301516040840151955160a085015161010086015160c0909601516102d398610695976001600160a01b0392831697909693959193918316921661422f565b6148db565b82516040516336cd42c360e01b8152602060048201529081906106c1906024830190612bb9565b0390fd5b346102d55760403660031901126102d5576102d36004356106e4612bf4565b906106f66106f1826136f9565b613ad4565b613b4b565b346102d55760e03660031901126102d5576004356001600160401b0381116102d55761072b903690600401612abd565b6024356001600160401b0381116102d55761074a903690600401612abd565b6044356001600160401b0381116102d557610769903690600401612abd565b6064356001600160a01b03811681036102d557608435906001600160a01b03821682036102d55760a4356001600160a01b03811693908490036102d55760c4356001600160a01b038116908190036102d55760008051602061578c83398151915254604081901c60ff161595906001600160401b03811680159081610f55575b6001149081610f4b575b159081610f42575b50610f31576001600160401b0319811660011760008051602061578c8339815191525586610f08575b50600b80546001600160a01b0319908116929092179055600c80549091169190911790556001600a558051906001600160401b0382116105c95761086d826104fc600854612d33565b602090601f8311600114610e9757826108a7936108bb96959361089993600092610e8c575b5050613311565b6008555b6001600355614eb3565b506108b36009546134bc565b600955614f2f565b506108c461538f565b6108cc61538f565b82516001600160401b0381116105c9576108f46000805160206155ec83398151915254612d33565b601f8111610e42575b506020601f8211600114610dc1578190610921939495600092610b36575050613311565b6000805160206155ec833981519152555b81516001600160401b0381116105c95761095a60008051602061580c83398151915254612d33565b601f8111610d77575b50602092601f8211600114610cf85761098792938291600092610ced575050613311565b60008051602061580c833981519152555b6040908151916109a88184612a7f565b600f83526e20b3b2b73a21b7b63632b1ba34b7b760891b60208401528051916109d18284612a7f565b60018352603160f81b60208401526109e761538f565b6109ef61538f565b83516001600160401b0381116105c957610a1760008051602061562c83398151915254612d33565b601f8111610ca3575b50602094601f8211600114610c2257610a46929394958291600092610c17575050613311565b60008051602061562c833981519152555b82516001600160401b0381116105c957610a7f60008051602061564c83398151915254612d33565b601f8111610bc2575b506020601f8211600114610b41578190610aac939495600092610b36575050613311565b60008051602061564c833981519152555b600060008051602061566c8339815191525560006000805160206157ac83398151915255610ae961538f565b610aef57005b60008051602061578c833981519152805460ff60401b1916905551600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a1005b015190508580610892565b601f1982169060008051602061564c833981519152600052806000209160005b818110610baa57509583600195969710610b91575b505050811b0160008051602061564c83398151915255610abd565b015160001960f88460031b161c19169055848080610b76565b9192602060018192868b015181550194019201610b61565b60008051602061564c833981519152600052610c07906000805160206157ec833981519152601f840160051c81019160208510610c0d575b601f0160051c0190613286565b84610a88565b9091508190610bfa565b015190508680610892565b601f1982169560008051602061562c833981519152600052806000209160005b888110610c8b57508360019596979810610c72575b505050811b0160008051602061562c83398151915255610a57565b015160001960f88460031b161c19169055858080610c57565b91926020600181928685015181550194019201610c42565b60008051602061562c833981519152600052610ce7906000805160206156cc833981519152601f840160051c81019160208510610c0d57601f0160051c0190613286565b85610a20565b015190508480610892565b601f1982169360008051602061580c833981519152600052806000209160005b868110610d5f5750836001959610610d46575b505050811b0160008051602061580c83398151915255610998565b015160001960f88460031b161c19169055838080610d2b565b91926020600181928685015181550194019201610d18565b60008051602061580c833981519152600052610dbb9060008051602061560c833981519152601f840160051c81019160208510610c0d57601f0160051c0190613286565b83610963565b601f198216906000805160206155ec833981519152600052806000209160005b818110610e2a57509583600195969710610e11575b505050811b016000805160206155ec83398151915255610932565b015160001960f88460031b161c19169055848080610df6565b9192602060018192868b015181550194019201610de1565b6000805160206155ec833981519152600052610e869060008051602061582c833981519152601f840160051c81019160208510610c0d57601f0160051c0190613286565b846108fd565b015190508a80610892565b90601f198316916008600052816000209260005b818110610ef057509260019285926108a7966108bb99989610610ed7575b505050811b0160085561089d565b015160001960f88460031b161c19169055898080610ec9565b92936020600181928786015181550195019301610eab565b6001600160481b0319166001600160401b011760008051602061578c8339815191525588610824565b63f92ee8a960e01b60005260046000fd5b9050158a6107fb565b303b1591506107f3565b8891506107e9565b346102d557610f6b36612b64565b610f73613dae565b610f7c82613a0b565b50610f8682613a41565b6001600160a01b031633141580610fbe575b610fa5576102d391614c81565b50635c3c03c360e01b6000523360045260245260446000fd5b503360009081526000805160206155ac833981519152602052604090205460ff1615610f98565b346102d55760203660031901126102d55760043561100281613a0b565b50600c5460405163b25457d360e01b81523060048201526024810192909252600090829060449082906001600160a01b03165afa9081156110e65760009161105f575b6040516020808252819061105b90820185612bb9565b0390f35b3d8083833e61106e8183612a7f565b8101906020818303126110de578051906001600160401b0382116110e2570181601f820112156110de578051926110a484612aa2565b926110b26040519485612a7f565b848452602085840101116110db575061105b926110d59160208085019101612b96565b90611045565b80fd5b8280fd5b8380fd5b6040513d6000823e3d90fd5b346102d55760203660031901126102d557606061010060405161111481612a63565b60008152826020820152826040820152600083820152600060808201528260a0820152600060c08201528260e0820152015261105b61115761061560043561383c565b60405191829160208352805160208401526101006112096111df6111a561118f602086015161012060408a0152610140890190612bb9565b6040860151888203601f190160608a0152612bb9565b60608501516001600160a01b039081166080898101919091528601511660a088810191909152850151878203601f190160c0890152612bb9565b60c08401516001600160a01b031660e087810191909152840151868203601f190184880152612cf6565b910151838203601f1901610120850152612c63565b346102d55760203660031901126102d55760043561123b81613a0b565b50600052600260205261105b6112576003604060002001612f64565b604051918291602083526020830190612bb9565b346102d55761129661127c3661301d565b81838587898b8d9b9f97989d8f6112916139f0565b613dfb565b60005260006020526040600020600160ff1982541617905560035493604051946112bf86612a63565b85526020850192835260408501978852606085019660018060a01b031696878152608086019160018060a01b03169889835260a0870193845260c087019433865260e088019687526101008801988952600454600160401b8110156105c95780600161132e920160045561383c565b9890986116cf5751885551805160018901916001600160401b0382116105c9576113628261135c8554612d33565b856132d7565b602090601f83116001146116685761138492916000918361165d575050613311565b90555b51805160028801916001600160401b0382116105c9576113ab8261135c8554612d33565b602090601f83116001146115f6576113cd9291600091836115eb575050613311565b90555b516003860180546001600160a01b03199081166001600160a01b03938416179091559151600487018054909316911617905551805160058501916001600160401b0382116105c9576114268261135c8554612d33565b602090601f831160011461158457611448929160009183610e8c575050613311565b90555b516006830180546001600160a01b0319166001600160a01b03929092169190911790555180519060078301906001600160401b0383116105c957602090611492848461385b565b0190600052602060002060005b838110611567575050505060080190519060208251926114bf84846138de565b019060005260206000206000915b8383106115495785856004549160001983019283116115335761152e926003546000526005602052604060002055600354917ff11434e9253f4ff089cbc840cb7a5db1fd5881d5fe71d35d6e611d2a12aa49486020604051858152a36134bc565b600355005b634e487b7160e01b600052601160045260246000fd5b600260208261155b6001945186613324565b019201920191906114cd565b82516001600160a01b03168183015560209092019160010161149f565b90601f1983169184600052816000209260005b8181106115d357509084600195949392106115ba575b505050811b01905561144b565b015160001960f88460031b161c191690558980806115ad565b92936020600181928786015181550195019301611597565b015190508d80610892565b90601f1983169184600052816000209260005b818110611645575090846001959493921061162c575b505050811b0190556113d0565b015160001960f88460031b161c191690558c808061161f565b92936020600181928786015181550195019301611609565b015190508e80610892565b90601f1983169184600052816000209260005b8181106116b7575090846001959493921061169e575b505050811b019055611387565b015160001960f88460031b161c191690558d8080611691565b9293602060018192878601518155019501930161167b565b634e487b7160e01b600052600060045260246000fd5b346102d55760203660031901126102d5576004356000526000602052602060ff604060002054166040519015158152f35b346102d55760803660031901126102d55761172f612bde565b611737612bf4565b6064359190604435906001600160401b0384116102d55761175f6102d3943690600401612abd565b9261176b838383613570565b33614b44565b346102d55760203660031901126102d55761178a613a86565b6102d36004356148db565b346102d5576102d36117c36117a93661301d565b81838587898b8e9f97989e9d809a9e6112919e9c9e613dae565b60005260006020526040600020600160ff19825416179055339661422f565b346102d55760003660031901126102d557600c546040516001600160a01b039091168152602090f35b346102d55760203660031901126102d5576004356001600160401b0381116102d557611854602061184181933690600401612abd565b8160405193828580945193849201612b96565b8101600181520301902054604051908152f35b346102d55760403660031901126102d557611880612bde565b602435908115158092036102d5576001600160a01b03169081156118ee576118a733613803565b82600052602052604060002060ff1981541660ff83161790556040519081527f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c3160203392a3005b50630b61174360e31b60005260045260246000fd5b346102d55760003660031901126102d557602060405160008152f35b346102d55760003660031901126102d557604051600060008051602061580c8339815191525461194e81612d33565b80845290600181169081156119d05750600114611976575b61105b8361125781850382612a7f565b60008051602061580c833981519152600090815260008051602061560c833981519152939250905b8082106119b657509091508101602001611257611966565b91926001816020925483858801015201910190929161199e565b60ff191660208086019190915291151560051b840190910191506112579050611966565b346102d55760403660031901126102d557611a0d612bf4565b60043560005260008051602061572c83398151915260205260406000209060018060a01b0316600052602052602060ff604060002054166040519015158152f35b346102d55760003660031901126102d55760008051602061566c833981519152541580611b62575b15611b2557611ac7611a86612e26565b611a8e612ec5565b6020611ad560405192611aa18385612a7f565b600084526000368137604051958695600f60f81b875260e08588015260e0870190612bb9565b908582036040870152612bb9565b466060850152306080850152600060a085015283810360c085015281808451928381520193019160005b828110611b0e57505050500390f35b835185528695509381019392810192600101611aff565b60405162461bcd60e51b81526020600482015260156024820152741152540dcc4c8e88155b9a5b9a5d1a585b1a5e9959605a1b6044820152606490fd5b506000805160206157ac8339815191525415611a76565b346102d55760203660031901126102d557600435611b9681613a0b565b50600052600260205261105b611bb26005604060002001613744565b604051918291602083526020830190612cf6565b346102d55760203660031901126102d5577fdaff73e0662ecb070f01270cb2aa647158f64dccd9df9b2f36a4de72d2e18d1e611c00612bde565b611c08613a86565b600c80546001600160a01b038381166001600160a01b03198316179092556040519283926105449216836137b0565b346102d55760003660031901126102d55760206040516000805160206156ec8339815191528152f35b346102d55760003660031901126102d5576020600a54604051908152f35b346102d55760203660031901126102d557600435611c9b81613a0b565b50600052600260205261105b6112576001604060002001612f64565b346102d55760003660031901126102d55761105b611257612d6d565b346102d55760203660031901126102d557611cec612bde565b6001600160a01b03811615611d0f57611d066020916137ca565b54604051908152f35b6322718ad960e21b600052600060045260246000fd5b346102d55760803660031901126102d557600435611d41612bf4565b60443591906064356001600160401b0381116102d557611d65903690600401612abd565b83600052600060205260ff60406000205416611fc557611d8483613a41565b336001600160a01b0390911603611fac57600b54604051637df73e2760e01b81526001600160a01b038481166004830181905292169591906020816024818a5afa9081156110e657600091611f8d575b5015611f7957600085815260026020819052604090912001546001600160a01b031692818414611f64576020611e85611e7c6024938660408051611e1781612a12565b8c815288878201520152611e77604051858101907ff2c6ac72efb8079ea7681f2985e6e973fadb946f39ff4a45b95c0a6110423feb82528c604082015288606082015289608082015260808152611e6f60a082612a7f565b5190206150e5565b6153bd565b909291926153f9565b604051637df73e2760e01b81526001600160a01b0390911660048201819052989092839182905afa9081156110e657600091611f35575b5015611f2057846000805160206157cc833981519152959650600052600260205260026040600020019060018060a01b031982541617905560005260006020526040600020600160ff19825416179055611f1b604051928392836137b0565b0390a2005b85632a6edb2b60e01b60005260045260246000fd5b611f57915060203d602011611f5d575b611f4f8183612a7f565b810190613798565b87611ebc565b503d611f45565b8363a139832960e01b60005260045260246000fd5b632a6edb2b60e01b60005260045260246000fd5b611fa6915060203d602011611f5d57611f4f8183612a7f565b87611dd4565b82635c3c03c360e01b6000523360045260245260446000fd5b83635c14634360e01b60005260045260246000fd5b346102d55760203660031901126102d557611ff3613a86565b6102d3600435613c70565b346102d55760203660031901126102d557602061201c600435613a41565b6040516001600160a01b039091168152f35b346102d55760203660031901126102d55760043561204b81613a0b565b506000526002602052602060018060a01b0360046040600020015416604051908152f35b346102d55760203660031901126102d55760043561208c81613a0b565b506000526002602052602060018060a01b0360026040600020015416604051908152f35b346102d55760203660031901126102d557600435606060c06040516120d481612a48565b828152826020820152600060408201528280820152600060808201528260a0820152015261210181613a0b565b50600052600260205260e0602061105b60406000206122286040519161212683612a48565b61212f81612f64565b835261213d60018201612f64565b83860190815260028201546001600160a01b0316604085019081529091906122069061216b60038401612f64565b6060870190815260048401546001600160a01b031660808801908152946121e4906121d16121ae60066121a060058a01613744565b9860a08d01998a5201613517565b9960c081019a8b526040519d8d8f9e928f938452519201526101008d0190612bb9565b90518b8203601f190160408d0152612bb9565b91516001600160a01b031660608a015251888203601f190160808a0152612bb9565b91516001600160a01b031660a087015251858203601f190160c0870152612cf6565b9051838203601f190160e0850152612c63565b346102d5576102d361224c36612cbc565b906040519261225c602085612a7f565b6000845261176b838383613570565b346102d55760203660031901126102d55760405161228881612a12565b600081526000602082015260408051916122a183612a2d565b6060835260606020840152015260406122be6102b86004356131e3565b61105b8251928392602084528051602085015260208101518285015201516060808401526080830190612c39565b346102d55760003660031901126102d557600b546040516001600160a01b039091168152602090f35b346102d55760403660031901126102d55761232e612bf4565b336001600160a01b0382160361234a576102d390600435613b4b565b63334bd91960e11b60005260046000fd5b346102d55760003660031901126102d5576020600654604051908152f35b346102d55760403660031901126102d5576102d3600435612398612bf4565b906123a56106f1826136f9565b613b1e565b346102d55760203660031901126102d55760206123c86004356136f9565b604051908152f35b346102d5576102d36123e136612cbc565b91613570565b346102d55760203660031901126102d55760043561240481613a0b565b50600052600260205261105b6124206006604060002001613517565b604051918291602083526020830190612c63565b346102d55761244236612c0a565b9061244c81613a0b565b5061245681613a41565b336001600160a01b039091160361045c5780600052600260205260036040600020019161248283612f64565b8151936001600160401b0385116105c9576124a7856124a18354612d33565b836132d7565b6020601f86116001146124f457856000805160206156ac83398151915295966124d7926000916124e95750613311565b90555b611f1b604051928392836134cb565b90508501518861052b565b601f1986169082600052806000209160005b81811061255057509187916000805160206156ac833981519152979860019410612537575b5050811b0190556124da565b86015160001960f88460031b161c19169055878061252b565b9192602060018192868a015181550194019201612506565b346102d55760003660031901126102d5576020600954604051908152f35b346102d55760403660031901126102d55761259f612bde565b6024356125ab81613a41565b33151580612665575b80612641575b61262c5781906001600160a01b0384811691167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925600080a4600090815260008051602061576c8339815191526020526040902080546001600160a01b0319166001600160a01b03909216919091179055005b63a9fbf51f60e01b6000523360045260246000fd5b5061264b81613803565b336000908152602091909152604090205460ff16156125ba565b506001600160a01b0381163314156125b4565b346102d55760203660031901126102d557602061201c60043561269a81613a41565b50613a5e565b346102d55760003660031901126102d55760405160006000805160206155ec833981519152546126cf81612d33565b80845290600181169081156119d057506001146126f65761105b8361125781850382612a7f565b6000805160206155ec833981519152600090815260008051602061582c833981519152939250905b80821061273657509091508101602001611257611966565b91926001816020925483858801015201910190929161271e565b346102d55760203660031901126102d55760043561276d81613a0b565b50600052600260205261105b6112576040600020612f64565b346102d55761279436612b64565b9061279d6139f0565b6127a681613a0b565b506127b081613a41565b6001600160a01b031633141580612957575b61045c5760009060005b60065481101561295057838260016127e3846131e3565b50015414806128e0575b6127fa57506001016127cc565b9061281392935061280c6002916131e3565b5001613324565b60015b1561281d57005b6003546040519061282d82612a12565b815260208101908282526040810190848252600654600160401b8110156105c95780600161285e92016006556131e3565b9290926116cf5761287d93600292518455516001840155519101613324565b6006546000198101929083116115335761152e92600354600052600760205260406000205551906128b06003549261349c565b907f8b06557af55313466fb2c7be6bf1db02af881cb0308214410a604185c79af52a6020604051858152a36134bc565b5060026129066129146128f2856131e3565b506040519283916020830195869101613218565b03601f198101835282612a7f565b5190208151604051612947602082816129368183019687815193849201612b96565b81010301601f198101835282612a7f565b519020146127ed565b5090612816565b503360009081526000805160206155ac833981519152602052604090205460ff16156127c2565b346102d55760203660031901126102d5576004359063ffffffff60e01b82168092036102d5576020916380ac58cd60e01b81148015612a02575b80156129f2575b809181156129d0575b505015158152f35b637965db0b60e01b14915081156129ea575b5083806129c8565b9050836129e2565b506301ffc9a760e01b81146129bf565b50635b5e139f60e01b81146129b8565b606081019081106001600160401b038211176105c957604052565b604081019081106001600160401b038211176105c957604052565b60e081019081106001600160401b038211176105c957604052565b61012081019081106001600160401b038211176105c957604052565b601f909101601f19168101906001600160401b038211908210176105c957604052565b6001600160401b0381116105c957601f01601f191660200190565b81601f820112156102d557602081359101612ad782612aa2565b92612ae56040519485612a7f565b828452828201116102d55781600092602092838601378301015290565b91906040838203126102d55760405190612b1b82612a2d565b909283919080356001600160401b0381116102d55782612b3c918301612abd565b83526020810135916001600160401b0383116102d557602092612b5f9201612abd565b910152565b9060406003198301126102d55760043591602435906001600160401b0382116102d557612b9391600401612b02565b90565b60005b838110612ba95750506000910152565b8181015183820152602001612b99565b90602091612bd281518092818552858086019101612b96565b601f01601f1916010190565b600435906001600160a01b03821682036102d557565b602435906001600160a01b03821682036102d557565b9060406003198301126102d55760043591602435906001600160401b0382116102d557612b9391600401612abd565b612b93916020612c528351604084526040840190612bb9565b920151906020818403910152612bb9565b9080602083519182815201916020808360051b8301019401926000915b838310612c8f57505050505090565b9091929394602080612cad600193601f198682030187528951612c39565b97019301930191939290612c80565b60609060031901126102d5576004356001600160a01b03811681036102d557906024356001600160a01b03811681036102d5579060443590565b906020808351928381520192019060005b818110612d145750505090565b82516001600160a01b0316845260209384019390920191600101612d07565b90600182811c92168015612d63575b6020831014612d4d57565b634e487b7160e01b600052602260045260246000fd5b91607f1691612d42565b6040519060008260085491612d8183612d33565b8083529260018116908115612e075750600114612da7575b612da592500383612a7f565b565b506008600090815290917ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee35b818310612deb575050906020612da592820101612d99565b6020919350806001915483858901015201910190918492612dd3565b60209250612da594915060ff191682840152151560051b820101612d99565b6040519060008260008051602061562c8339815191525491612e4783612d33565b8083529260018116908115612e075750600114612e6a57612da592500383612a7f565b5060008051602061562c833981519152600090815290916000805160206156cc8339815191525b818310612ea9575050906020612da592820101612d99565b6020919350806001915483858901015201910190918492612e91565b6040519060008260008051602061564c8339815191525491612ee683612d33565b8083529260018116908115612e075750600114612f0957612da592500383612a7f565b5060008051602061564c833981519152600090815290916000805160206157ec8339815191525b818310612f48575050906020612da592820101612d99565b6020919350806001915483858901015201910190918492612f30565b9060405191826000825492612f7884612d33565b8084529360018116908115612fe45750600114612f9d575b50612da592500383612a7f565b90506000929192526020600020906000915b818310612fc8575050906020612da59282010138612f90565b6020919350806001915483858901015201910190918492612faf565b905060209250612da594915060ff191682840152151560051b82010138612f90565b6001600160401b0381116105c95760051b60200190565b6101206003198201126102d5576004356001600160a01b03811681036102d557916024356001600160401b0381116102d557826023820112156102d557806004013561306881613006565b916130766040519384612a7f565b8183526024602084019260051b820101908582116102d557602401915b8183106131c3575091939150506044356001600160a01b03811681036102d557916064356001600160401b0381116102d557826130d291600401612abd565b916084356001600160401b0381116102d557816130f191600401612abd565b9160a4356001600160401b0381116102d5578261311091600401612abd565b9160c4356001600160401b0381116102d557816023820112156102d557806004013561313b81613006565b916131496040519384612a7f565b8183526024602084019260051b820101918483116102d55760248201905b83821061319557505050509160e43591610104359060018060401b0382116102d557612b9391600401612abd565b81356001600160401b0381116102d5576020916131b88860248594880101612b02565b815201910190613167565b82356001600160a01b03811681036102d557815260209283019201613093565b60065481101561320257600660005260206000209060021b0190600090565b634e487b7160e01b600052603260045260246000fd5b6000929181549161322883612d33565b9260018116908115613273575060011461324157505050565b909192935060005260206000206000905b83821061325f5750500190565b600181602092548486015201910190613252565b60ff191683525050811515909102019150565b818110613291575050565b60008155600101613286565b90601f82116132aa575050565b612da59160086000526020600020906020601f840160051c83019310610c0d57601f0160051c0190613286565b9190601f81116132e657505050565b612da5926000526020600020906020601f840160051c83019310610c0d57601f0160051c0190613286565b8160011b916000199060031b1c19161790565b81518051909391906001600160401b0381116105c95761334e816133488454612d33565b846132d7565b6020601f82116001146134305781600193926133779260209697986000926133c3575050613311565b81555b0192015191825160018060401b0381116105c95761339c816133488454612d33565b6020601f82116001146133ce5781906133bf9394956000926133c3575050613311565b9055565b015190503880610892565b601f1982169083600052806000209160005b818110613418575095836001959697106133ff575b505050811b019055565b015160001960f88460031b161c191690553880806133f5565b9192602060018192868b0151815501940192016133e0565b601f1982169083600052806000209160005b818110613484575091839160209697986001969587951061346b575b505050811b01815561337a565b015160001960f88460031b161c1916905538808061345e565b9192602060018192868c015181550194019201613442565b6134b490602060405192828480945193849201612b96565b810103902090565b60001981146115335760010190565b90916134e2612b9393604084526040840190612bb9565b916020818403910152612bb9565b906040516134fd81612a2d565b6020612b5f6001839561350f81612f64565b855201612f64565b90815461352381613006565b926135316040519485612a7f565b818452602084019060005260206000206000915b8383106135525750505050565b60026020600192613562856134f0565b815201920192019190613545565b909291906001600160a01b03841680156136e35760009461359084614e60565b953315158061364a575b5091849160008051602061570c8339815191526135c997989460018060a01b0386169889938461362b576137ca565b8054600101905584815260008051602061574c8339815191526020526040812080546001600160a01b0319168517905580a4506001600160a01b031680830361361157505050565b6364283d7b60e01b60005260045260245260445260646000fd5b61363487614e88565b61363d886137ca565b80546000190190556137ca565b80613693575b1561365b573861359a565b84906001600160a01b03881661367d57602491637e27328960e01b8252600452fd5b60449163177e802f60e01b825233600452602452fd5b506001600160a01b038716331480156136c6575b806136505750336001600160a01b036136bf87613a5e565b1614613650565b506136d087613803565b338252602052604081205460ff166136a7565b633250574960e11b600052600060045260246000fd5b60005260008051602061572c83398151915260205260016040600020015490565b9060405161372781612a12565b6040612b5f600283958054855260018101546020860152016134f0565b906040519182815491828252602082019060005260206000209260005b818110613776575050612da592500383612a7f565b84546001600160a01b0316835260019485019487945060209093019201613761565b908160209103126102d5575180151581036102d55790565b6001600160a01b0391821681529116602082015260400190565b6001600160a01b031660009081527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab00793036020526040902090565b6001600160a01b031660009081527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab00793056020526040902090565b6004548110156132025760046000526009602060002091020190600090565b90600160401b81116105c95781549080835581811061387957505050565b612da59260005260206000209182019101613286565b6138998154612d33565b90816138a3575050565b81601f600093116001146138b5575055565b818352602083206138d191601f0160051c810190600101613286565b8082528160208120915555565b600160401b82116105c9578054908281558183106138fb57505050565b6001600160ff1b0382168203611533576001600160ff1b03831683036115335760005260206000209060011b81019160011b015b81811061393a575050565b8061394660029261388f565b6139526001820161388f565b0161392f565b9060405161396581612a63565b610100612b5f600883958054855261397f60018201612f64565b602086015261399060028201612f64565b604086015260038101546001600160a01b03908116606087015260048201541660808601526139c160058201612f64565b60a086015260068101546001600160a01b031660c08601526139e560078201613744565b60e086015201613517565b600954156139fa57565b63cf8aa8cb60e01b60005260046000fd5b80158015613a35575b613a2157612b9390614e60565b637e27328960e01b60005260045260246000fd5b50600a54811015613a14565b613a4a81614e60565b906001600160a01b03821615613a21575090565b600090815260008051602061576c83398151915260205260409020546001600160a01b031690565b3360009081526000805160206155ac833981519152602052604090205460ff1615613aad57565b63e2517d3f60e01b600052336004526000805160206156ec83398151915260245260446000fd5b600081815260008051602061572c8339815191526020908152604080832033845290915290205460ff1615613b065750565b63e2517d3f60e01b6000523360045260245260446000fd5b90612b93916000805160206156ec8339815191528103614fb557613b436009546134bc565b600955614fb5565b906000805160206156ec8339815191528214613b6b575b612b9391615045565b60095491821561153357600019909201600955613b62565b919091828114613c4757613b978354612d33565b6001600160401b0381116105c957613bb3816133488454612d33565b600093601f8211600114613be2576133bf9293948291600092613bd7575050613311565b015490503880610892565b845260208085208386529085209094601f198316815b818110613c2f57509583600195969710613c1657505050811b019055565b015460001960f88460031b161c191690553880806133f5565b9192600180602092868b015481550194019201613bf8565b509050565b818103613c57575050565b60018083613c68612da59585613b83565b019101613b83565b80600052600760205260406000205490613c8c6102b8836131e3565b90600654600019810190811161153357808403613d49575b506006548015613d335760001901613cbb816131e3565b6116cf576003816000613ce2935560006001820155613cdc6002820161388f565b0161388f565b6006556000526007602052600060408120557f1b684b8ada21a0a6786e403a19ac9c0de5efdb63365e9744c92d21521a5c545f6020613d296040828501519401515161349c565b93604051908152a3565b634e487b7160e01b600052603160045260246000fd5b613d52906131e3565b50613d5c846131e3565b6116cf57818103613d8a575b5050613d73836131e3565b505460005260076020528260406000205538613ca4565b60028083613da79454845560018101546001850155019101613c4c565b3880613d68565b600954151580613dd4575b613dbf57565b63198dd43760e21b6000523360045260246000fd5b503360009081526000805160206155ac833981519152602052604090205460ff1615613db9565b989497929596939093613e0d82615167565b1561420d5780600052600060205260ff604060002054166141f9576040516020818451613e3d8183858901612b96565b81016001815203019020546141d757613e5588614de6565b1561042d57604051996101008b016001600160401b0381118c8210176105c95760405260018060a01b03168a528460208b015260408a019260018060a01b03169485845260608b0197885260808b0192835260a08b01998a5260c08b0198895260e08b0191825260405180602081019283602082519192019060005b8181106141b55750505003601f1981018252613eed9082612a7f565b5190209151604051806020810192838151602081930191613f0d92612b96565b810103808252613f209060200182612a7f565b5190209651604051806020810192838151602081930191613f4092612b96565b810103808252613f539060200182612a7f565b5190209851604051806020810192838151602081930191613f7392612b96565b810103808252613f869060200182612a7f565b519020998851519a613f978c613006565b9b6040519c613fa6908e612a7f565b808d52613fb290613006565b60208d019990601f1901368b3760005b8b51805182101561404457908e61403d828f602061403481613ff385613fea8160019c614dd2565b51519551614dd2565b510151604051928391816140108185019788815193849201612b96565b830161402482518093858085019101612b96565b010103601f198101835282612a7f565b51902092614dd2565b5201613fc2565b505093979b91959950939799919599604051906020820180935190919060005b81811061419f57505050926141179a98959282614094611e779a9794611e7c9d9a9703601f198101835282612a7f565b5190209360018060a01b039051169660018060a01b039051169451946040519660208801987f57547536a15daff6624f9ef1f9fcc4d254a4a97e839c6e82a9b8a397cb6d118e8a5260408901526060880152608087015260a086015260c085015260e08401526101008301526101208201526101208152611e6f61014082612a7f565b600b54604051637df73e2760e01b81526001600160a01b039283166004820181905292909160209183916024918391165afa9081156110e657600091614180575b5015611f7957818103614169575050565b637ba5ffb560e01b60005260045260245260446000fd5b614199915060203d602011611f5d57611f4f8183612a7f565b38614158565b8251845260209384019390920191600101614064565b82516001600160a01b0316845285945060209384019390920191600101613ed1565b6040516336cd42c360e01b815260206004820152806106c16024820185612bb9565b635c14634360e01b60005260045260246000fd5b604051637f19f48d60e01b815260206004820152806106c16024820185612bb9565b97969091929593946040519361424485612a48565b86855260208501978852604085019060018060a01b03169788825260608601928352608086019360018060a01b0316845260a0860194855260c08601968752600a5460005260026020526040600020955180519060018060401b0382116105c9576142b9826142b38a54612d33565b8a6132d7565b602090601f83116001146147b5576142db9291600091836133c3575050613311565b86555b51805160018701916001600160401b0382116105c9576143028261135c8554612d33565b602090601f831160011461474e576143249291600091836133c3575050613311565b90555b516002850180546001600160a01b0319166001600160a01b039290921691909117905551805160038501916001600160401b0382116105c95761436e8261135c8554612d33565b602090601f83116001146146e7576143909291600091836133c3575050613311565b90555b516004830180546001600160a01b0319166001600160a01b03929092169190911790555180519060058301906001600160401b0383116105c9576020906143da848461385b565b0190600052602060002060005b8381106146ca5750505050600601905190602082519261440784846138de565b019060005260206000206000915b8383106146ac575050505060209061443f600a549283928160405193828580945193849201612b96565b81016001815203019020556020906040519061445b8383612a7f565b600082526001600160a01b0386169586156136e35761447982614e60565b6001600160a01b038116801515919084908a908461468d575b61449b866137ca565b80546001019055600083815260008051602061574c8339815191528a526040812080546001600160a01b0319168417905560008051602061570c8339815191529080a450614677573b614546575b5050926000805160206157cc8339815191526040614541949584600a549586957f9d228d69b5fdb8d273a2336f8fb8612d039631024ea9bf09c424a9503aa078f0848651898152a382519160008352820152a26134bc565b600a55565b9161458491819594936040519384928392630a85bd0160e11b8452336004850152600060248501526044840152608060648401526084830190612bb9565b03816000895af18091600091614635575b50906145ef575050503d6000146145e8573d6145b081612aa2565b906145be6040519283612a7f565b81523d60008383013e5b805191826145e55783633250574960e11b60005260045260246000fd5b01fd5b60606145c8565b919290916001600160e01b03191663757a42ff60e11b01614620576000805160206157cc83398151915260406144e9565b83633250574960e11b60005260045260246000fd5b8581813d8311614670575b61464a8183612a7f565b8101031261466c5751906001600160e01b0319821682036110db575038614595565b5080fd5b503d614640565b6339e3563760e11b600052600060045260246000fd5b61469683614e88565b61469f846137ca565b8054600019019055614492565b60026020826146be6001945186613324565b01920192019190614415565b82516001600160a01b0316818301556020909201916001016143e7565b90601f1983169184600052816000209260005b818110614736575090846001959493921061471d575b505050811b019055614393565b015160001960f88460031b161c19169055388080614710565b929360206001819287860151815501950193016146fa565b90601f1983169184600052816000209260005b81811061479d5750908460019594939210614784575b505050811b019055614327565b015160001960f88460031b161c19169055388080614777565b92936020600181928786015181550195019301614761565b90601f1983169189600052816000209260005b81811061480457509084600195949392106147eb575b505050811b0186556142de565b015160001960f88460031b161c191690553880806147de565b929360206001819287860151815501950193016147c8565b818114614885578154916001600160401b0383116105c95761483e838361385b565b60005260206000209060005260206000208154916000925b848410614864575050505050565b600191820180546001600160a01b0390921684860155939091019290614856565b5050565b8181146148855781549161489d83836138de565b60005260206000209060005260206000206000915b8383106148bf5750505050565b600280826148cf60019486613c4c565b019201920191906148b2565b806000526005602052604060002054906148f76106158361383c565b90600454600019810190811161153357808403614a48575b506004548015613d3357600019016149268161383c565b6116cf578060006008925561493d6001820161388f565b6149496002820161388f565b60006003820155600060048201556149636005820161388f565b600060068201556007810180546000825580614a2c575b50500180549060008155816149e6575b50506004556000526005602052600060408120557fc3cee524930441f1267d79dc72f161b7625e84f8052db5f175e091c0e86bda3c602060018060a01b0360608401511692608060018060a01b039101511693604051908152a3565b6001600160ff1b03821682036115335760005260206000209060011b8101905b8181101561498a5780614a1a60029261388f565b614a266001820161388f565b01614a06565b614a4191600052602060002090810190613286565b388061497a565b614a519061383c565b50614a5b8461383c565b6116cf57818103614a89575b5050614a728361383c565b50546000526005602052826040600020553861490f565b60088083614b3d94548455614aa46001820160018601613b83565b614ab46002820160028601613b83565b60038181015490850180546001600160a01b039283166001600160a01b0319918216179091556004808401549087018054919093169116179055614afe6005808301908601613b83565b60068181015490850180546001600160a01b0319166001600160a01b0392909216919091179055614b35600780830190860161481c565b019101614889565b3880614a67565b823b614b52575b5050505050565b604051630a85bd0160e11b81526001600160a01b039182166004820152918116602483015260448201939093526080606482015291169160209082908190614b9e906084830190612bb9565b03816000865af18091600091614c3e575b5090614c0857503d15614c01573d614bc681612aa2565b90614bd46040519283612a7f565b81523d6000602083013e5b80519081614bfc5782633250574960e11b60005260045260246000fd5b602001fd5b6060614bdf565b6001600160e01b03191663757a42ff60e11b01614c2a57503880808080614b4b565b633250574960e11b60005260045260246000fd5b6020813d602011614c79575b81614c5760209383612a7f565b8101031261466c5751906001600160e01b0319821682036110db575038614baf565b3d9150614c4a565b8060005260026020527f102c9e6462b7d4e559e582578ea590c7f08c5d0551bd6c2d047768a4b1377e656006604060002001614cd3614cc98551926020870193845191615200565b945191519161349c565b93614ce3604051928392836134cb565b0390a3565b80548210156132025760005260206000209060011b0190600090565b60005b815480821015614dc957614d1b8284614ce8565b50604051614d3181612906602082018095613218565b5190206040516020810190614d53602082816129368b87815193849201612b96565b51902014614d645750600101614d07565b91925090600019810190811161153357614d81614d899184614ce8565b509183614ce8565b9190916116cf57614d9991613c4c565b80548015613d33576000190190614db08282614ce8565b6116cf57600181613cdc614dc39361388f565b55600190565b50505050600090565b80518210156132025760209160051b010190565b60005b8151811015614e5957614dfc8183614dd2565b5151604051614e1b602082816129368183019687815193849201612b96565b519020604051602081019066191959985d5b1d60ca1b825260078152614e42602782612a7f565b51902014614e5257600101614de9565b5050600190565b5050600090565b600090815260008051602061574c83398151915260205260409020546001600160a01b031690565b600090815260008051602061576c8339815191526020526040902080546001600160a01b0319169055565b6001600160a01b038116600090815260008051602061568c833981519152602052604090205460ff16614f29576001600160a01b0316600081815260008051602061568c83398151915260205260408120805460ff1916600117905533919060008051602061558c8339815191528180a4600190565b50600090565b6001600160a01b03811660009081526000805160206155ac833981519152602052604090205460ff16614f29576001600160a01b031660008181526000805160206155ac83398151915260205260408120805460ff191660011790553391906000805160206156ec8339815191529060008051602061558c8339815191529080a4600190565b600081815260008051602061572c833981519152602090815260408083206001600160a01b038616845290915290205460ff16614e5957600081815260008051602061572c833981519152602090815260408083206001600160a01b0395909516808452949091528120805460ff191660011790553392919060008051602061558c8339815191529080a4600190565b600081815260008051602061572c833981519152602090815260408083206001600160a01b038616845290915290205460ff1615614e5957600081815260008051602061572c833981519152602090815260408083206001600160a01b0395909516808452949091528120805460ff19169055339291907ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9080a4600190565b6042906150f0615500565b6150f8615558565b6040519060208201927f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f8452604083015260608201524660808201523060a082015260a0815261514960c082612a7f565b519020906040519161190160f01b8352600283015260228201522090565b805115614f295760005b8151811015614e5257818101602001516000906001600160f81b031916606160f81b811015806151f2575b1590816151cd575b816151be575b506151b85750600101615171565b91505090565b602d60f81b14159050386151aa565b9050600360fc1b811015806151e4575b15906151a4565b50603960f81b8111156151dd565b50603d60f91b81111561519c565b9290919260005b81548110156153365761521a8183614ce8565b5060405161523081612906602082018095613218565b5190206040516020810190615252602082816129368b87815193849201612b96565b5190201461526257600101615207565b8192935080615281600161527a615287948296614ce8565b5001612f64565b93614ce8565b5084519101906001600160401b0381116105c9576152a9816133488454612d33565b602094601f82116001146152d3576152ce9293949582916000926133c3575050613311565b905590565b601f1982169583600052806000209160005b88811061531e57508360019596979810615305575b505050811b01905590565b015160001960f88460031b161c191690553880806152fa565b919260206001819286850151815501940192016152e5565b5090926040519261534684612a2d565b835260208301528054600160401b8110156105c95761536a91600182018155614ce8565b9190916116cf5761537a91613324565b604051615388602082612a7f565b6000815290565b60ff60008051602061578c8339815191525460401c16156153ac57565b631afcd79f60e31b60005260046000fd5b81519190604183036153ee576153e792506020820151906060604084015193015160001a90615481565b9192909190565b505060009160029190565b919091600481101561546b578061540f57509050565b60006001820361542a5763f645eedf60e01b60005260046000fd5b5060028103615448578263fce698f760e01b60005260045260246000fd5b9091600360009214615458575050565b6335e2f38360e21b825260045260249150fd5b634e487b7160e01b600052602160045260246000fd5b91906fa2a8918ca85bafe22016d0b997e4df60600160ff1b0384116154f4579160209360809260ff60009560405194855216868401526040830152606082015282805260015afa156110e6576000516001600160a01b038116156154e85790600090600090565b50600090600190600090565b50505060009160039190565b615508612e26565b8051908115615518576020012090565b505060008051602061566c8339815191525480156155335790565b507fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a47090565b615560612ec5565b8051908115615570576020012090565b50506000805160206157ac833981519152548015615533579056fe2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d1a864d4eac0891250c5be130dc5ed49253048fe7a091bb5fcd3f026d0ce30376e21432e1fe2b572d5803dd7316b7a854952317b42017f920a616ec70cdb8a5c180bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab0079300f4bad0a69248f59680a4f2b3000328cec71a413447c96781cfe5996daa8c456ea16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d102a16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d103a16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d100b7db2dd08fcb62d0c9e08c51941cae53c267786a0b75803fb7960902fc8ef97d0631fae4c84380e8028a67e87eb9ca8b60d5b14feafecaabaedf3aeb32a81cdd42ad5d3e1f2e6e70edcf6d991b8a3023d3fca8047a131592f9edb9fd9b89d57d71f3d55856e4058ed06ee057d79ada615f65cdf5f9ee88181b914225088f834fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b62680080bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930280bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab0079304f0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00a16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d101c0e455cb361d77f8df0f9d49668a41d5a79c14f930e2a8ede12b9735f02b37c65f9ce34815f8e11431c7bb75a8e6886a91478f7ffc1dbb0a98dc240fddd76b7580bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930137c58c799b6609234b945e882912ee9ad34948a1dfaa20a97485e1a7752bbf81a2646970667358221220be16687307dc57b55b761cd336cbea9f0a33d54098c402322fbda4361b7a3f4464736f6c634300081c0033",
}

// AgentCollectionV1ABI is the input ABI used to generate the binding from.
// Deprecated: Use AgentCollectionV1MetaData.ABI instead.
var AgentCollectionV1ABI = AgentCollectionV1MetaData.ABI

// AgentCollectionV1Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AgentCollectionV1MetaData.Bin instead.
var AgentCollectionV1Bin = AgentCollectionV1MetaData.Bin

// DeployAgentCollectionV1 deploys a new Ethereum contract, binding an instance of AgentCollectionV1 to it.
func DeployAgentCollectionV1(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AgentCollectionV1, error) {
	parsed, err := AgentCollectionV1MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AgentCollectionV1Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AgentCollectionV1{AgentCollectionV1Caller: AgentCollectionV1Caller{contract: contract}, AgentCollectionV1Transactor: AgentCollectionV1Transactor{contract: contract}, AgentCollectionV1Filterer: AgentCollectionV1Filterer{contract: contract}}, nil
}

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

// AddOrUpdateImageProposal is a free data retrieval call binding the contract method 0x419f5aa0.
//
// Solidity: function addOrUpdateImageProposal(uint256 proposalIndex) view returns((uint256,uint256,(string,string)))
func (_AgentCollectionV1 *AgentCollectionV1Caller) AddOrUpdateImageProposal(opts *bind.CallOpts, proposalIndex *big.Int) (AddOrUpdateImageProposal, error) {
	var out []interface{}
	err := _AgentCollectionV1.contract.Call(opts, &out, "addOrUpdateImageProposal", proposalIndex)

	if err != nil {
		return *new(AddOrUpdateImageProposal), err
	}

	out0 := *abi.ConvertType(out[0], new(AddOrUpdateImageProposal)).(*AddOrUpdateImageProposal)

	return out0, err

}

// AddOrUpdateImageProposal is a free data retrieval call binding the contract method 0x419f5aa0.
//
// Solidity: function addOrUpdateImageProposal(uint256 proposalIndex) view returns((uint256,uint256,(string,string)))
func (_AgentCollectionV1 *AgentCollectionV1Session) AddOrUpdateImageProposal(proposalIndex *big.Int) (AddOrUpdateImageProposal, error) {
	return _AgentCollectionV1.Contract.AddOrUpdateImageProposal(&_AgentCollectionV1.CallOpts, proposalIndex)
}

// AddOrUpdateImageProposal is a free data retrieval call binding the contract method 0x419f5aa0.
//
// Solidity: function addOrUpdateImageProposal(uint256 proposalIndex) view returns((uint256,uint256,(string,string)))
func (_AgentCollectionV1 *AgentCollectionV1CallerSession) AddOrUpdateImageProposal(proposalIndex *big.Int) (AddOrUpdateImageProposal, error) {
	return _AgentCollectionV1.Contract.AddOrUpdateImageProposal(&_AgentCollectionV1.CallOpts, proposalIndex)
}

// AddOrUpdateImageProposalsLength is a free data retrieval call binding the contract method 0x34b481e7.
//
// Solidity: function addOrUpdateImageProposalsLength() view returns(uint256)
func (_AgentCollectionV1 *AgentCollectionV1Caller) AddOrUpdateImageProposalsLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AgentCollectionV1.contract.Call(opts, &out, "addOrUpdateImageProposalsLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AddOrUpdateImageProposalsLength is a free data retrieval call binding the contract method 0x34b481e7.
//
// Solidity: function addOrUpdateImageProposalsLength() view returns(uint256)
func (_AgentCollectionV1 *AgentCollectionV1Session) AddOrUpdateImageProposalsLength() (*big.Int, error) {
	return _AgentCollectionV1.Contract.AddOrUpdateImageProposalsLength(&_AgentCollectionV1.CallOpts)
}

// AddOrUpdateImageProposalsLength is a free data retrieval call binding the contract method 0x34b481e7.
//
// Solidity: function addOrUpdateImageProposalsLength() view returns(uint256)
func (_AgentCollectionV1 *AgentCollectionV1CallerSession) AddOrUpdateImageProposalsLength() (*big.Int, error) {
	return _AgentCollectionV1.Contract.AddOrUpdateImageProposalsLength(&_AgentCollectionV1.CallOpts)
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

// CollectionsManager is a free data retrieval call binding the contract method 0x37c11b2a.
//
// Solidity: function collectionsManager() view returns(address)
func (_AgentCollectionV1 *AgentCollectionV1Caller) CollectionsManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AgentCollectionV1.contract.Call(opts, &out, "collectionsManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CollectionsManager is a free data retrieval call binding the contract method 0x37c11b2a.
//
// Solidity: function collectionsManager() view returns(address)
func (_AgentCollectionV1 *AgentCollectionV1Session) CollectionsManager() (common.Address, error) {
	return _AgentCollectionV1.Contract.CollectionsManager(&_AgentCollectionV1.CallOpts)
}

// CollectionsManager is a free data retrieval call binding the contract method 0x37c11b2a.
//
// Solidity: function collectionsManager() view returns(address)
func (_AgentCollectionV1 *AgentCollectionV1CallerSession) CollectionsManager() (common.Address, error) {
	return _AgentCollectionV1.Contract.CollectionsManager(&_AgentCollectionV1.CallOpts)
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

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_AgentCollectionV1 *AgentCollectionV1Caller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _AgentCollectionV1.contract.Call(opts, &out, "eip712Domain")

	outstruct := new(struct {
		Fields            [1]byte
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
		Salt              [32]byte
		Extensions        []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_AgentCollectionV1 *AgentCollectionV1Session) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _AgentCollectionV1.Contract.Eip712Domain(&_AgentCollectionV1.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_AgentCollectionV1 *AgentCollectionV1CallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _AgentCollectionV1.Contract.Eip712Domain(&_AgentCollectionV1.CallOpts)
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

// ImagesOf is a free data retrieval call binding the contract method 0x1278fc8b.
//
// Solidity: function imagesOf(uint256 tokenId) view returns((string,string)[])
func (_AgentCollectionV1 *AgentCollectionV1Caller) ImagesOf(opts *bind.CallOpts, tokenId *big.Int) ([]KeyValue, error) {
	var out []interface{}
	err := _AgentCollectionV1.contract.Call(opts, &out, "imagesOf", tokenId)

	if err != nil {
		return *new([]KeyValue), err
	}

	out0 := *abi.ConvertType(out[0], new([]KeyValue)).(*[]KeyValue)

	return out0, err

}

// ImagesOf is a free data retrieval call binding the contract method 0x1278fc8b.
//
// Solidity: function imagesOf(uint256 tokenId) view returns((string,string)[])
func (_AgentCollectionV1 *AgentCollectionV1Session) ImagesOf(tokenId *big.Int) ([]KeyValue, error) {
	return _AgentCollectionV1.Contract.ImagesOf(&_AgentCollectionV1.CallOpts, tokenId)
}

// ImagesOf is a free data retrieval call binding the contract method 0x1278fc8b.
//
// Solidity: function imagesOf(uint256 tokenId) view returns((string,string)[])
func (_AgentCollectionV1 *AgentCollectionV1CallerSession) ImagesOf(tokenId *big.Int) ([]KeyValue, error) {
	return _AgentCollectionV1.Contract.ImagesOf(&_AgentCollectionV1.CallOpts, tokenId)
}

// InformationOf is a free data retrieval call binding the contract method 0x50cffb68.
//
// Solidity: function informationOf(uint256 tokenId) view returns((string,string,address,string,address,address[],(string,string)[]))
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
// Solidity: function informationOf(uint256 tokenId) view returns((string,string,address,string,address,address[],(string,string)[]))
func (_AgentCollectionV1 *AgentCollectionV1Session) InformationOf(tokenId *big.Int) (AgentInformation, error) {
	return _AgentCollectionV1.Contract.InformationOf(&_AgentCollectionV1.CallOpts, tokenId)
}

// InformationOf is a free data retrieval call binding the contract method 0x50cffb68.
//
// Solidity: function informationOf(uint256 tokenId) view returns((string,string,address,string,address,address[],(string,string)[]))
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

// IsNonceUsed is a free data retrieval call binding the contract method 0xb90665e5.
//
// Solidity: function isNonceUsed(bytes32 nonce) view returns(bool)
func (_AgentCollectionV1 *AgentCollectionV1Caller) IsNonceUsed(opts *bind.CallOpts, nonce [32]byte) (bool, error) {
	var out []interface{}
	err := _AgentCollectionV1.contract.Call(opts, &out, "isNonceUsed", nonce)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsNonceUsed is a free data retrieval call binding the contract method 0xb90665e5.
//
// Solidity: function isNonceUsed(bytes32 nonce) view returns(bool)
func (_AgentCollectionV1 *AgentCollectionV1Session) IsNonceUsed(nonce [32]byte) (bool, error) {
	return _AgentCollectionV1.Contract.IsNonceUsed(&_AgentCollectionV1.CallOpts, nonce)
}

// IsNonceUsed is a free data retrieval call binding the contract method 0xb90665e5.
//
// Solidity: function isNonceUsed(bytes32 nonce) view returns(bool)
func (_AgentCollectionV1 *AgentCollectionV1CallerSession) IsNonceUsed(nonce [32]byte) (bool, error) {
	return _AgentCollectionV1.Contract.IsNonceUsed(&_AgentCollectionV1.CallOpts, nonce)
}

// MintProposal is a free data retrieval call binding the contract method 0xc088c9a7.
//
// Solidity: function mintProposal(uint256 proposalIndex) view returns((uint256,string,string,address,address,string,address,address[],(string,string)[]))
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
// Solidity: function mintProposal(uint256 proposalIndex) view returns((uint256,string,string,address,address,string,address,address[],(string,string)[]))
func (_AgentCollectionV1 *AgentCollectionV1Session) MintProposal(proposalIndex *big.Int) (MintProposal, error) {
	return _AgentCollectionV1.Contract.MintProposal(&_AgentCollectionV1.CallOpts, proposalIndex)
}

// MintProposal is a free data retrieval call binding the contract method 0xc088c9a7.
//
// Solidity: function mintProposal(uint256 proposalIndex) view returns((uint256,string,string,address,address,string,address,address[],(string,string)[]))
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

// MissionOf is a free data retrieval call binding the contract method 0xbb046abc.
//
// Solidity: function missionOf(uint256 tokenId) view returns(string)
func (_AgentCollectionV1 *AgentCollectionV1Caller) MissionOf(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _AgentCollectionV1.contract.Call(opts, &out, "missionOf", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// MissionOf is a free data retrieval call binding the contract method 0xbb046abc.
//
// Solidity: function missionOf(uint256 tokenId) view returns(string)
func (_AgentCollectionV1 *AgentCollectionV1Session) MissionOf(tokenId *big.Int) (string, error) {
	return _AgentCollectionV1.Contract.MissionOf(&_AgentCollectionV1.CallOpts, tokenId)
}

// MissionOf is a free data retrieval call binding the contract method 0xbb046abc.
//
// Solidity: function missionOf(uint256 tokenId) view returns(string)
func (_AgentCollectionV1 *AgentCollectionV1CallerSession) MissionOf(tokenId *big.Int) (string, error) {
	return _AgentCollectionV1.Contract.MissionOf(&_AgentCollectionV1.CallOpts, tokenId)
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

// RoleOf is a free data retrieval call binding the contract method 0x72b1d8cf.
//
// Solidity: function roleOf(uint256 tokenId) view returns(string)
func (_AgentCollectionV1 *AgentCollectionV1Caller) RoleOf(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _AgentCollectionV1.contract.Call(opts, &out, "roleOf", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// RoleOf is a free data retrieval call binding the contract method 0x72b1d8cf.
//
// Solidity: function roleOf(uint256 tokenId) view returns(string)
func (_AgentCollectionV1 *AgentCollectionV1Session) RoleOf(tokenId *big.Int) (string, error) {
	return _AgentCollectionV1.Contract.RoleOf(&_AgentCollectionV1.CallOpts, tokenId)
}

// RoleOf is a free data retrieval call binding the contract method 0x72b1d8cf.
//
// Solidity: function roleOf(uint256 tokenId) view returns(string)
func (_AgentCollectionV1 *AgentCollectionV1CallerSession) RoleOf(tokenId *big.Int) (string, error) {
	return _AgentCollectionV1.Contract.RoleOf(&_AgentCollectionV1.CallOpts, tokenId)
}

// SignerOf is a free data retrieval call binding the contract method 0x5161fdf5.
//
// Solidity: function signerOf(uint256 tokenId) view returns(address)
func (_AgentCollectionV1 *AgentCollectionV1Caller) SignerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _AgentCollectionV1.contract.Call(opts, &out, "signerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SignerOf is a free data retrieval call binding the contract method 0x5161fdf5.
//
// Solidity: function signerOf(uint256 tokenId) view returns(address)
func (_AgentCollectionV1 *AgentCollectionV1Session) SignerOf(tokenId *big.Int) (common.Address, error) {
	return _AgentCollectionV1.Contract.SignerOf(&_AgentCollectionV1.CallOpts, tokenId)
}

// SignerOf is a free data retrieval call binding the contract method 0x5161fdf5.
//
// Solidity: function signerOf(uint256 tokenId) view returns(address)
func (_AgentCollectionV1 *AgentCollectionV1CallerSession) SignerOf(tokenId *big.Int) (common.Address, error) {
	return _AgentCollectionV1.Contract.SignerOf(&_AgentCollectionV1.CallOpts, tokenId)
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

// TokenIdOf is a free data retrieval call binding the contract method 0xa398d819.
//
// Solidity: function tokenIdOf(string name) view returns(uint256)
func (_AgentCollectionV1 *AgentCollectionV1Caller) TokenIdOf(opts *bind.CallOpts, name string) (*big.Int, error) {
	var out []interface{}
	err := _AgentCollectionV1.contract.Call(opts, &out, "tokenIdOf", name)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenIdOf is a free data retrieval call binding the contract method 0xa398d819.
//
// Solidity: function tokenIdOf(string name) view returns(uint256)
func (_AgentCollectionV1 *AgentCollectionV1Session) TokenIdOf(name string) (*big.Int, error) {
	return _AgentCollectionV1.Contract.TokenIdOf(&_AgentCollectionV1.CallOpts, name)
}

// TokenIdOf is a free data retrieval call binding the contract method 0xa398d819.
//
// Solidity: function tokenIdOf(string name) view returns(uint256)
func (_AgentCollectionV1 *AgentCollectionV1CallerSession) TokenIdOf(name string) (*big.Int, error) {
	return _AgentCollectionV1.Contract.TokenIdOf(&_AgentCollectionV1.CallOpts, name)
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

// WalletsOf is a free data retrieval call binding the contract method 0x83648671.
//
// Solidity: function walletsOf(uint256 tokenId) view returns(address[])
func (_AgentCollectionV1 *AgentCollectionV1Caller) WalletsOf(opts *bind.CallOpts, tokenId *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _AgentCollectionV1.contract.Call(opts, &out, "walletsOf", tokenId)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// WalletsOf is a free data retrieval call binding the contract method 0x83648671.
//
// Solidity: function walletsOf(uint256 tokenId) view returns(address[])
func (_AgentCollectionV1 *AgentCollectionV1Session) WalletsOf(tokenId *big.Int) ([]common.Address, error) {
	return _AgentCollectionV1.Contract.WalletsOf(&_AgentCollectionV1.CallOpts, tokenId)
}

// WalletsOf is a free data retrieval call binding the contract method 0x83648671.
//
// Solidity: function walletsOf(uint256 tokenId) view returns(address[])
func (_AgentCollectionV1 *AgentCollectionV1CallerSession) WalletsOf(tokenId *big.Int) ([]common.Address, error) {
	return _AgentCollectionV1.Contract.WalletsOf(&_AgentCollectionV1.CallOpts, tokenId)
}

// AddOrUpdateImage is a paid mutator transaction binding the contract method 0xd0c0fcdf.
//
// Solidity: function addOrUpdateImage(uint256 tokenId, (string,string) image) returns()
func (_AgentCollectionV1 *AgentCollectionV1Transactor) AddOrUpdateImage(opts *bind.TransactOpts, tokenId *big.Int, image KeyValue) (*types.Transaction, error) {
	return _AgentCollectionV1.contract.Transact(opts, "addOrUpdateImage", tokenId, image)
}

// AddOrUpdateImage is a paid mutator transaction binding the contract method 0xd0c0fcdf.
//
// Solidity: function addOrUpdateImage(uint256 tokenId, (string,string) image) returns()
func (_AgentCollectionV1 *AgentCollectionV1Session) AddOrUpdateImage(tokenId *big.Int, image KeyValue) (*types.Transaction, error) {
	return _AgentCollectionV1.Contract.AddOrUpdateImage(&_AgentCollectionV1.TransactOpts, tokenId, image)
}

// AddOrUpdateImage is a paid mutator transaction binding the contract method 0xd0c0fcdf.
//
// Solidity: function addOrUpdateImage(uint256 tokenId, (string,string) image) returns()
func (_AgentCollectionV1 *AgentCollectionV1TransactorSession) AddOrUpdateImage(tokenId *big.Int, image KeyValue) (*types.Transaction, error) {
	return _AgentCollectionV1.Contract.AddOrUpdateImage(&_AgentCollectionV1.TransactOpts, tokenId, image)
}

// AddOrUpdateImageFromProposal is a paid mutator transaction binding the contract method 0xfede2263.
//
// Solidity: function addOrUpdateImageFromProposal(uint256 proposalId) returns()
func (_AgentCollectionV1 *AgentCollectionV1Transactor) AddOrUpdateImageFromProposal(opts *bind.TransactOpts, proposalId *big.Int) (*types.Transaction, error) {
	return _AgentCollectionV1.contract.Transact(opts, "addOrUpdateImageFromProposal", proposalId)
}

// AddOrUpdateImageFromProposal is a paid mutator transaction binding the contract method 0xfede2263.
//
// Solidity: function addOrUpdateImageFromProposal(uint256 proposalId) returns()
func (_AgentCollectionV1 *AgentCollectionV1Session) AddOrUpdateImageFromProposal(proposalId *big.Int) (*types.Transaction, error) {
	return _AgentCollectionV1.Contract.AddOrUpdateImageFromProposal(&_AgentCollectionV1.TransactOpts, proposalId)
}

// AddOrUpdateImageFromProposal is a paid mutator transaction binding the contract method 0xfede2263.
//
// Solidity: function addOrUpdateImageFromProposal(uint256 proposalId) returns()
func (_AgentCollectionV1 *AgentCollectionV1TransactorSession) AddOrUpdateImageFromProposal(proposalId *big.Int) (*types.Transaction, error) {
	return _AgentCollectionV1.Contract.AddOrUpdateImageFromProposal(&_AgentCollectionV1.TransactOpts, proposalId)
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

// CreateAddOrUpdateImageProposal is a paid mutator transaction binding the contract method 0x02d51463.
//
// Solidity: function createAddOrUpdateImageProposal(uint256 tokenId, (string,string) image) returns()
func (_AgentCollectionV1 *AgentCollectionV1Transactor) CreateAddOrUpdateImageProposal(opts *bind.TransactOpts, tokenId *big.Int, image KeyValue) (*types.Transaction, error) {
	return _AgentCollectionV1.contract.Transact(opts, "createAddOrUpdateImageProposal", tokenId, image)
}

// CreateAddOrUpdateImageProposal is a paid mutator transaction binding the contract method 0x02d51463.
//
// Solidity: function createAddOrUpdateImageProposal(uint256 tokenId, (string,string) image) returns()
func (_AgentCollectionV1 *AgentCollectionV1Session) CreateAddOrUpdateImageProposal(tokenId *big.Int, image KeyValue) (*types.Transaction, error) {
	return _AgentCollectionV1.Contract.CreateAddOrUpdateImageProposal(&_AgentCollectionV1.TransactOpts, tokenId, image)
}

// CreateAddOrUpdateImageProposal is a paid mutator transaction binding the contract method 0x02d51463.
//
// Solidity: function createAddOrUpdateImageProposal(uint256 tokenId, (string,string) image) returns()
func (_AgentCollectionV1 *AgentCollectionV1TransactorSession) CreateAddOrUpdateImageProposal(tokenId *big.Int, image KeyValue) (*types.Transaction, error) {
	return _AgentCollectionV1.Contract.CreateAddOrUpdateImageProposal(&_AgentCollectionV1.TransactOpts, tokenId, image)
}

// CreateMintProposal is a paid mutator transaction binding the contract method 0xbaf3dfd3.
//
// Solidity: function createMintProposal(address to, address[] agentAddresses, address signer, string role, string name, string mission, (string,string)[] images, bytes32 nonce, bytes signature) returns()
func (_AgentCollectionV1 *AgentCollectionV1Transactor) CreateMintProposal(opts *bind.TransactOpts, to common.Address, agentAddresses []common.Address, signer common.Address, role string, name string, mission string, images []KeyValue, nonce [32]byte, signature []byte) (*types.Transaction, error) {
	return _AgentCollectionV1.contract.Transact(opts, "createMintProposal", to, agentAddresses, signer, role, name, mission, images, nonce, signature)
}

// CreateMintProposal is a paid mutator transaction binding the contract method 0xbaf3dfd3.
//
// Solidity: function createMintProposal(address to, address[] agentAddresses, address signer, string role, string name, string mission, (string,string)[] images, bytes32 nonce, bytes signature) returns()
func (_AgentCollectionV1 *AgentCollectionV1Session) CreateMintProposal(to common.Address, agentAddresses []common.Address, signer common.Address, role string, name string, mission string, images []KeyValue, nonce [32]byte, signature []byte) (*types.Transaction, error) {
	return _AgentCollectionV1.Contract.CreateMintProposal(&_AgentCollectionV1.TransactOpts, to, agentAddresses, signer, role, name, mission, images, nonce, signature)
}

// CreateMintProposal is a paid mutator transaction binding the contract method 0xbaf3dfd3.
//
// Solidity: function createMintProposal(address to, address[] agentAddresses, address signer, string role, string name, string mission, (string,string)[] images, bytes32 nonce, bytes signature) returns()
func (_AgentCollectionV1 *AgentCollectionV1TransactorSession) CreateMintProposal(to common.Address, agentAddresses []common.Address, signer common.Address, role string, name string, mission string, images []KeyValue, nonce [32]byte, signature []byte) (*types.Transaction, error) {
	return _AgentCollectionV1.Contract.CreateMintProposal(&_AgentCollectionV1.TransactOpts, to, agentAddresses, signer, role, name, mission, images, nonce, signature)
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

// Initialize is a paid mutator transaction binding the contract method 0xd10321bd.
//
// Solidity: function initialize(string name, string symbol, string initialDescription, address admin, address moderator, address initialCollectionsManager, address initialCollectionsDescriptor) returns()
func (_AgentCollectionV1 *AgentCollectionV1Transactor) Initialize(opts *bind.TransactOpts, name string, symbol string, initialDescription string, admin common.Address, moderator common.Address, initialCollectionsManager common.Address, initialCollectionsDescriptor common.Address) (*types.Transaction, error) {
	return _AgentCollectionV1.contract.Transact(opts, "initialize", name, symbol, initialDescription, admin, moderator, initialCollectionsManager, initialCollectionsDescriptor)
}

// Initialize is a paid mutator transaction binding the contract method 0xd10321bd.
//
// Solidity: function initialize(string name, string symbol, string initialDescription, address admin, address moderator, address initialCollectionsManager, address initialCollectionsDescriptor) returns()
func (_AgentCollectionV1 *AgentCollectionV1Session) Initialize(name string, symbol string, initialDescription string, admin common.Address, moderator common.Address, initialCollectionsManager common.Address, initialCollectionsDescriptor common.Address) (*types.Transaction, error) {
	return _AgentCollectionV1.Contract.Initialize(&_AgentCollectionV1.TransactOpts, name, symbol, initialDescription, admin, moderator, initialCollectionsManager, initialCollectionsDescriptor)
}

// Initialize is a paid mutator transaction binding the contract method 0xd10321bd.
//
// Solidity: function initialize(string name, string symbol, string initialDescription, address admin, address moderator, address initialCollectionsManager, address initialCollectionsDescriptor) returns()
func (_AgentCollectionV1 *AgentCollectionV1TransactorSession) Initialize(name string, symbol string, initialDescription string, admin common.Address, moderator common.Address, initialCollectionsManager common.Address, initialCollectionsDescriptor common.Address) (*types.Transaction, error) {
	return _AgentCollectionV1.Contract.Initialize(&_AgentCollectionV1.TransactOpts, name, symbol, initialDescription, admin, moderator, initialCollectionsManager, initialCollectionsDescriptor)
}

// Mint is a paid mutator transaction binding the contract method 0xaf907545.
//
// Solidity: function mint(address to, address[] agentAddresses, address signer, string role, string name, string mission, (string,string)[] images, bytes32 nonce, bytes signature) returns()
func (_AgentCollectionV1 *AgentCollectionV1Transactor) Mint(opts *bind.TransactOpts, to common.Address, agentAddresses []common.Address, signer common.Address, role string, name string, mission string, images []KeyValue, nonce [32]byte, signature []byte) (*types.Transaction, error) {
	return _AgentCollectionV1.contract.Transact(opts, "mint", to, agentAddresses, signer, role, name, mission, images, nonce, signature)
}

// Mint is a paid mutator transaction binding the contract method 0xaf907545.
//
// Solidity: function mint(address to, address[] agentAddresses, address signer, string role, string name, string mission, (string,string)[] images, bytes32 nonce, bytes signature) returns()
func (_AgentCollectionV1 *AgentCollectionV1Session) Mint(to common.Address, agentAddresses []common.Address, signer common.Address, role string, name string, mission string, images []KeyValue, nonce [32]byte, signature []byte) (*types.Transaction, error) {
	return _AgentCollectionV1.Contract.Mint(&_AgentCollectionV1.TransactOpts, to, agentAddresses, signer, role, name, mission, images, nonce, signature)
}

// Mint is a paid mutator transaction binding the contract method 0xaf907545.
//
// Solidity: function mint(address to, address[] agentAddresses, address signer, string role, string name, string mission, (string,string)[] images, bytes32 nonce, bytes signature) returns()
func (_AgentCollectionV1 *AgentCollectionV1TransactorSession) Mint(to common.Address, agentAddresses []common.Address, signer common.Address, role string, name string, mission string, images []KeyValue, nonce [32]byte, signature []byte) (*types.Transaction, error) {
	return _AgentCollectionV1.Contract.Mint(&_AgentCollectionV1.TransactOpts, to, agentAddresses, signer, role, name, mission, images, nonce, signature)
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

// RefuseAddOrUpdateImageProposal is a paid mutator transaction binding the contract method 0x67c1e20b.
//
// Solidity: function refuseAddOrUpdateImageProposal(uint256 proposalId) returns()
func (_AgentCollectionV1 *AgentCollectionV1Transactor) RefuseAddOrUpdateImageProposal(opts *bind.TransactOpts, proposalId *big.Int) (*types.Transaction, error) {
	return _AgentCollectionV1.contract.Transact(opts, "refuseAddOrUpdateImageProposal", proposalId)
}

// RefuseAddOrUpdateImageProposal is a paid mutator transaction binding the contract method 0x67c1e20b.
//
// Solidity: function refuseAddOrUpdateImageProposal(uint256 proposalId) returns()
func (_AgentCollectionV1 *AgentCollectionV1Session) RefuseAddOrUpdateImageProposal(proposalId *big.Int) (*types.Transaction, error) {
	return _AgentCollectionV1.Contract.RefuseAddOrUpdateImageProposal(&_AgentCollectionV1.TransactOpts, proposalId)
}

// RefuseAddOrUpdateImageProposal is a paid mutator transaction binding the contract method 0x67c1e20b.
//
// Solidity: function refuseAddOrUpdateImageProposal(uint256 proposalId) returns()
func (_AgentCollectionV1 *AgentCollectionV1TransactorSession) RefuseAddOrUpdateImageProposal(proposalId *big.Int) (*types.Transaction, error) {
	return _AgentCollectionV1.Contract.RefuseAddOrUpdateImageProposal(&_AgentCollectionV1.TransactOpts, proposalId)
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

// RemoveImage is a paid mutator transaction binding the contract method 0xe8c43e7c.
//
// Solidity: function removeImage(uint256 tokenId, string key) returns()
func (_AgentCollectionV1 *AgentCollectionV1Transactor) RemoveImage(opts *bind.TransactOpts, tokenId *big.Int, key string) (*types.Transaction, error) {
	return _AgentCollectionV1.contract.Transact(opts, "removeImage", tokenId, key)
}

// RemoveImage is a paid mutator transaction binding the contract method 0xe8c43e7c.
//
// Solidity: function removeImage(uint256 tokenId, string key) returns()
func (_AgentCollectionV1 *AgentCollectionV1Session) RemoveImage(tokenId *big.Int, key string) (*types.Transaction, error) {
	return _AgentCollectionV1.Contract.RemoveImage(&_AgentCollectionV1.TransactOpts, tokenId, key)
}

// RemoveImage is a paid mutator transaction binding the contract method 0xe8c43e7c.
//
// Solidity: function removeImage(uint256 tokenId, string key) returns()
func (_AgentCollectionV1 *AgentCollectionV1TransactorSession) RemoveImage(tokenId *big.Int, key string) (*types.Transaction, error) {
	return _AgentCollectionV1.Contract.RemoveImage(&_AgentCollectionV1.TransactOpts, tokenId, key)
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

// UpdateAgentMission is a paid mutator transaction binding the contract method 0x1164e719.
//
// Solidity: function updateAgentMission(uint256 tokenId, string newMission) returns()
func (_AgentCollectionV1 *AgentCollectionV1Transactor) UpdateAgentMission(opts *bind.TransactOpts, tokenId *big.Int, newMission string) (*types.Transaction, error) {
	return _AgentCollectionV1.contract.Transact(opts, "updateAgentMission", tokenId, newMission)
}

// UpdateAgentMission is a paid mutator transaction binding the contract method 0x1164e719.
//
// Solidity: function updateAgentMission(uint256 tokenId, string newMission) returns()
func (_AgentCollectionV1 *AgentCollectionV1Session) UpdateAgentMission(tokenId *big.Int, newMission string) (*types.Transaction, error) {
	return _AgentCollectionV1.Contract.UpdateAgentMission(&_AgentCollectionV1.TransactOpts, tokenId, newMission)
}

// UpdateAgentMission is a paid mutator transaction binding the contract method 0x1164e719.
//
// Solidity: function updateAgentMission(uint256 tokenId, string newMission) returns()
func (_AgentCollectionV1 *AgentCollectionV1TransactorSession) UpdateAgentMission(tokenId *big.Int, newMission string) (*types.Transaction, error) {
	return _AgentCollectionV1.Contract.UpdateAgentMission(&_AgentCollectionV1.TransactOpts, tokenId, newMission)
}

// UpdateAgentSigner is a paid mutator transaction binding the contract method 0x6d21b484.
//
// Solidity: function updateAgentSigner(uint256 tokenId, address newSigner, bytes32 nonce, bytes signature) returns()
func (_AgentCollectionV1 *AgentCollectionV1Transactor) UpdateAgentSigner(opts *bind.TransactOpts, tokenId *big.Int, newSigner common.Address, nonce [32]byte, signature []byte) (*types.Transaction, error) {
	return _AgentCollectionV1.contract.Transact(opts, "updateAgentSigner", tokenId, newSigner, nonce, signature)
}

// UpdateAgentSigner is a paid mutator transaction binding the contract method 0x6d21b484.
//
// Solidity: function updateAgentSigner(uint256 tokenId, address newSigner, bytes32 nonce, bytes signature) returns()
func (_AgentCollectionV1 *AgentCollectionV1Session) UpdateAgentSigner(tokenId *big.Int, newSigner common.Address, nonce [32]byte, signature []byte) (*types.Transaction, error) {
	return _AgentCollectionV1.Contract.UpdateAgentSigner(&_AgentCollectionV1.TransactOpts, tokenId, newSigner, nonce, signature)
}

// UpdateAgentSigner is a paid mutator transaction binding the contract method 0x6d21b484.
//
// Solidity: function updateAgentSigner(uint256 tokenId, address newSigner, bytes32 nonce, bytes signature) returns()
func (_AgentCollectionV1 *AgentCollectionV1TransactorSession) UpdateAgentSigner(tokenId *big.Int, newSigner common.Address, nonce [32]byte, signature []byte) (*types.Transaction, error) {
	return _AgentCollectionV1.Contract.UpdateAgentSigner(&_AgentCollectionV1.TransactOpts, tokenId, newSigner, nonce, signature)
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

// AgentCollectionV1AgentMissionUpdatedIterator is returned from FilterAgentMissionUpdated and is used to iterate over the raw logs and unpacked data for AgentMissionUpdated events raised by the AgentCollectionV1 contract.
type AgentCollectionV1AgentMissionUpdatedIterator struct {
	Event *AgentCollectionV1AgentMissionUpdated // Event containing the contract specifics and raw log

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
func (it *AgentCollectionV1AgentMissionUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentCollectionV1AgentMissionUpdated)
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
		it.Event = new(AgentCollectionV1AgentMissionUpdated)
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
func (it *AgentCollectionV1AgentMissionUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentCollectionV1AgentMissionUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentCollectionV1AgentMissionUpdated represents a AgentMissionUpdated event raised by the AgentCollectionV1 contract.
type AgentCollectionV1AgentMissionUpdated struct {
	TokenId    *big.Int
	OldMission string
	NewMission string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAgentMissionUpdated is a free log retrieval operation binding the contract event 0x0631fae4c84380e8028a67e87eb9ca8b60d5b14feafecaabaedf3aeb32a81cdd.
//
// Solidity: event AgentMissionUpdated(uint256 indexed tokenId, string oldMission, string newMission)
func (_AgentCollectionV1 *AgentCollectionV1Filterer) FilterAgentMissionUpdated(opts *bind.FilterOpts, tokenId []*big.Int) (*AgentCollectionV1AgentMissionUpdatedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _AgentCollectionV1.contract.FilterLogs(opts, "AgentMissionUpdated", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &AgentCollectionV1AgentMissionUpdatedIterator{contract: _AgentCollectionV1.contract, event: "AgentMissionUpdated", logs: logs, sub: sub}, nil
}

// WatchAgentMissionUpdated is a free log subscription operation binding the contract event 0x0631fae4c84380e8028a67e87eb9ca8b60d5b14feafecaabaedf3aeb32a81cdd.
//
// Solidity: event AgentMissionUpdated(uint256 indexed tokenId, string oldMission, string newMission)
func (_AgentCollectionV1 *AgentCollectionV1Filterer) WatchAgentMissionUpdated(opts *bind.WatchOpts, sink chan<- *AgentCollectionV1AgentMissionUpdated, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _AgentCollectionV1.contract.WatchLogs(opts, "AgentMissionUpdated", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentCollectionV1AgentMissionUpdated)
				if err := _AgentCollectionV1.contract.UnpackLog(event, "AgentMissionUpdated", log); err != nil {
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

// ParseAgentMissionUpdated is a log parse operation binding the contract event 0x0631fae4c84380e8028a67e87eb9ca8b60d5b14feafecaabaedf3aeb32a81cdd.
//
// Solidity: event AgentMissionUpdated(uint256 indexed tokenId, string oldMission, string newMission)
func (_AgentCollectionV1 *AgentCollectionV1Filterer) ParseAgentMissionUpdated(log types.Log) (*AgentCollectionV1AgentMissionUpdated, error) {
	event := new(AgentCollectionV1AgentMissionUpdated)
	if err := _AgentCollectionV1.contract.UnpackLog(event, "AgentMissionUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentCollectionV1AgentSignerUpdatedIterator is returned from FilterAgentSignerUpdated and is used to iterate over the raw logs and unpacked data for AgentSignerUpdated events raised by the AgentCollectionV1 contract.
type AgentCollectionV1AgentSignerUpdatedIterator struct {
	Event *AgentCollectionV1AgentSignerUpdated // Event containing the contract specifics and raw log

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
func (it *AgentCollectionV1AgentSignerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentCollectionV1AgentSignerUpdated)
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
		it.Event = new(AgentCollectionV1AgentSignerUpdated)
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
func (it *AgentCollectionV1AgentSignerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentCollectionV1AgentSignerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentCollectionV1AgentSignerUpdated represents a AgentSignerUpdated event raised by the AgentCollectionV1 contract.
type AgentCollectionV1AgentSignerUpdated struct {
	TokenId   *big.Int
	OldSigner common.Address
	NewSigner common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAgentSignerUpdated is a free log retrieval operation binding the contract event 0xc0e455cb361d77f8df0f9d49668a41d5a79c14f930e2a8ede12b9735f02b37c6.
//
// Solidity: event AgentSignerUpdated(uint256 indexed tokenId, address oldSigner, address newSigner)
func (_AgentCollectionV1 *AgentCollectionV1Filterer) FilterAgentSignerUpdated(opts *bind.FilterOpts, tokenId []*big.Int) (*AgentCollectionV1AgentSignerUpdatedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _AgentCollectionV1.contract.FilterLogs(opts, "AgentSignerUpdated", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &AgentCollectionV1AgentSignerUpdatedIterator{contract: _AgentCollectionV1.contract, event: "AgentSignerUpdated", logs: logs, sub: sub}, nil
}

// WatchAgentSignerUpdated is a free log subscription operation binding the contract event 0xc0e455cb361d77f8df0f9d49668a41d5a79c14f930e2a8ede12b9735f02b37c6.
//
// Solidity: event AgentSignerUpdated(uint256 indexed tokenId, address oldSigner, address newSigner)
func (_AgentCollectionV1 *AgentCollectionV1Filterer) WatchAgentSignerUpdated(opts *bind.WatchOpts, sink chan<- *AgentCollectionV1AgentSignerUpdated, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _AgentCollectionV1.contract.WatchLogs(opts, "AgentSignerUpdated", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentCollectionV1AgentSignerUpdated)
				if err := _AgentCollectionV1.contract.UnpackLog(event, "AgentSignerUpdated", log); err != nil {
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

// ParseAgentSignerUpdated is a log parse operation binding the contract event 0xc0e455cb361d77f8df0f9d49668a41d5a79c14f930e2a8ede12b9735f02b37c6.
//
// Solidity: event AgentSignerUpdated(uint256 indexed tokenId, address oldSigner, address newSigner)
func (_AgentCollectionV1 *AgentCollectionV1Filterer) ParseAgentSignerUpdated(log types.Log) (*AgentCollectionV1AgentSignerUpdated, error) {
	event := new(AgentCollectionV1AgentSignerUpdated)
	if err := _AgentCollectionV1.contract.UnpackLog(event, "AgentSignerUpdated", log); err != nil {
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
// Solidity: event CollectionsDescriptorUpdated(address oldCollectionsDescriptor, address newCollectionsDescriptor)
func (_AgentCollectionV1 *AgentCollectionV1Filterer) FilterCollectionsDescriptorUpdated(opts *bind.FilterOpts) (*AgentCollectionV1CollectionsDescriptorUpdatedIterator, error) {

	logs, sub, err := _AgentCollectionV1.contract.FilterLogs(opts, "CollectionsDescriptorUpdated")
	if err != nil {
		return nil, err
	}
	return &AgentCollectionV1CollectionsDescriptorUpdatedIterator{contract: _AgentCollectionV1.contract, event: "CollectionsDescriptorUpdated", logs: logs, sub: sub}, nil
}

// WatchCollectionsDescriptorUpdated is a free log subscription operation binding the contract event 0xdaff73e0662ecb070f01270cb2aa647158f64dccd9df9b2f36a4de72d2e18d1e.
//
// Solidity: event CollectionsDescriptorUpdated(address oldCollectionsDescriptor, address newCollectionsDescriptor)
func (_AgentCollectionV1 *AgentCollectionV1Filterer) WatchCollectionsDescriptorUpdated(opts *bind.WatchOpts, sink chan<- *AgentCollectionV1CollectionsDescriptorUpdated) (event.Subscription, error) {

	logs, sub, err := _AgentCollectionV1.contract.WatchLogs(opts, "CollectionsDescriptorUpdated")
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
// Solidity: event CollectionsDescriptorUpdated(address oldCollectionsDescriptor, address newCollectionsDescriptor)
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

// AgentCollectionV1EIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the AgentCollectionV1 contract.
type AgentCollectionV1EIP712DomainChangedIterator struct {
	Event *AgentCollectionV1EIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *AgentCollectionV1EIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentCollectionV1EIP712DomainChanged)
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
		it.Event = new(AgentCollectionV1EIP712DomainChanged)
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
func (it *AgentCollectionV1EIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentCollectionV1EIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentCollectionV1EIP712DomainChanged represents a EIP712DomainChanged event raised by the AgentCollectionV1 contract.
type AgentCollectionV1EIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_AgentCollectionV1 *AgentCollectionV1Filterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*AgentCollectionV1EIP712DomainChangedIterator, error) {

	logs, sub, err := _AgentCollectionV1.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &AgentCollectionV1EIP712DomainChangedIterator{contract: _AgentCollectionV1.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_AgentCollectionV1 *AgentCollectionV1Filterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *AgentCollectionV1EIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _AgentCollectionV1.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentCollectionV1EIP712DomainChanged)
				if err := _AgentCollectionV1.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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

// ParseEIP712DomainChanged is a log parse operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_AgentCollectionV1 *AgentCollectionV1Filterer) ParseEIP712DomainChanged(log types.Log) (*AgentCollectionV1EIP712DomainChanged, error) {
	event := new(AgentCollectionV1EIP712DomainChanged)
	if err := _AgentCollectionV1.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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
	Signer     common.Address
	ProposalId *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMintProposalCreated is a free log retrieval operation binding the contract event 0xf11434e9253f4ff089cbc840cb7a5db1fd5881d5fe71d35d6e611d2a12aa4948.
//
// Solidity: event MintProposalCreated(address indexed from, address indexed signer, uint256 proposalId)
func (_AgentCollectionV1 *AgentCollectionV1Filterer) FilterMintProposalCreated(opts *bind.FilterOpts, from []common.Address, signer []common.Address) (*AgentCollectionV1MintProposalCreatedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _AgentCollectionV1.contract.FilterLogs(opts, "MintProposalCreated", fromRule, signerRule)
	if err != nil {
		return nil, err
	}
	return &AgentCollectionV1MintProposalCreatedIterator{contract: _AgentCollectionV1.contract, event: "MintProposalCreated", logs: logs, sub: sub}, nil
}

// WatchMintProposalCreated is a free log subscription operation binding the contract event 0xf11434e9253f4ff089cbc840cb7a5db1fd5881d5fe71d35d6e611d2a12aa4948.
//
// Solidity: event MintProposalCreated(address indexed from, address indexed signer, uint256 proposalId)
func (_AgentCollectionV1 *AgentCollectionV1Filterer) WatchMintProposalCreated(opts *bind.WatchOpts, sink chan<- *AgentCollectionV1MintProposalCreated, from []common.Address, signer []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _AgentCollectionV1.contract.WatchLogs(opts, "MintProposalCreated", fromRule, signerRule)
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

// ParseMintProposalCreated is a log parse operation binding the contract event 0xf11434e9253f4ff089cbc840cb7a5db1fd5881d5fe71d35d6e611d2a12aa4948.
//
// Solidity: event MintProposalCreated(address indexed from, address indexed signer, uint256 proposalId)
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
	Signer     common.Address
	ProposalId *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMintProposalRemoved is a free log retrieval operation binding the contract event 0xc3cee524930441f1267d79dc72f161b7625e84f8052db5f175e091c0e86bda3c.
//
// Solidity: event MintProposalRemoved(address indexed owner, address indexed signer, uint256 proposalId)
func (_AgentCollectionV1 *AgentCollectionV1Filterer) FilterMintProposalRemoved(opts *bind.FilterOpts, owner []common.Address, signer []common.Address) (*AgentCollectionV1MintProposalRemovedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _AgentCollectionV1.contract.FilterLogs(opts, "MintProposalRemoved", ownerRule, signerRule)
	if err != nil {
		return nil, err
	}
	return &AgentCollectionV1MintProposalRemovedIterator{contract: _AgentCollectionV1.contract, event: "MintProposalRemoved", logs: logs, sub: sub}, nil
}

// WatchMintProposalRemoved is a free log subscription operation binding the contract event 0xc3cee524930441f1267d79dc72f161b7625e84f8052db5f175e091c0e86bda3c.
//
// Solidity: event MintProposalRemoved(address indexed owner, address indexed signer, uint256 proposalId)
func (_AgentCollectionV1 *AgentCollectionV1Filterer) WatchMintProposalRemoved(opts *bind.WatchOpts, sink chan<- *AgentCollectionV1MintProposalRemoved, owner []common.Address, signer []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _AgentCollectionV1.contract.WatchLogs(opts, "MintProposalRemoved", ownerRule, signerRule)
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

// ParseMintProposalRemoved is a log parse operation binding the contract event 0xc3cee524930441f1267d79dc72f161b7625e84f8052db5f175e091c0e86bda3c.
//
// Solidity: event MintProposalRemoved(address indexed owner, address indexed signer, uint256 proposalId)
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
	Signer  common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMinted is a free log retrieval operation binding the contract event 0x9d228d69b5fdb8d273a2336f8fb8612d039631024ea9bf09c424a9503aa078f0.
//
// Solidity: event Minted(address indexed to, address indexed signer, uint256 tokenId)
func (_AgentCollectionV1 *AgentCollectionV1Filterer) FilterMinted(opts *bind.FilterOpts, to []common.Address, signer []common.Address) (*AgentCollectionV1MintedIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _AgentCollectionV1.contract.FilterLogs(opts, "Minted", toRule, signerRule)
	if err != nil {
		return nil, err
	}
	return &AgentCollectionV1MintedIterator{contract: _AgentCollectionV1.contract, event: "Minted", logs: logs, sub: sub}, nil
}

// WatchMinted is a free log subscription operation binding the contract event 0x9d228d69b5fdb8d273a2336f8fb8612d039631024ea9bf09c424a9503aa078f0.
//
// Solidity: event Minted(address indexed to, address indexed signer, uint256 tokenId)
func (_AgentCollectionV1 *AgentCollectionV1Filterer) WatchMinted(opts *bind.WatchOpts, sink chan<- *AgentCollectionV1Minted, to []common.Address, signer []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _AgentCollectionV1.contract.WatchLogs(opts, "Minted", toRule, signerRule)
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

// ParseMinted is a log parse operation binding the contract event 0x9d228d69b5fdb8d273a2336f8fb8612d039631024ea9bf09c424a9503aa078f0.
//
// Solidity: event Minted(address indexed to, address indexed signer, uint256 tokenId)
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
