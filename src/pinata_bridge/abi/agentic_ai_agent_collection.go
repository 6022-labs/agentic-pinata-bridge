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

// AgenticAIAgentCollectionMetaData contains all meta data concerning the AgenticAIAgentCollection contract.
var AgenticAIAgentCollectionMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agentDescriptor\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721IncorrectOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721InsufficientApproval\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC721InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721NonexistentToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"InvalidName\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidShortString\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"expectedSigner\",\"type\":\"address\"}],\"name\":\"InvalidSigner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"NotSigner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"SignerUnchanged\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"}],\"name\":\"StringTooLong\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"UsedName\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nonce\",\"type\":\"bytes32\"}],\"name\":\"UsedNonce\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Minted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"SignerAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"SignerRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldSigner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newSigner\",\"type\":\"address\"}],\"name\":\"TokenSignerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SIGNER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"addSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"creatorOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"imageOf\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nonce\",\"type\":\"bytes32\"}],\"name\":\"isNonceUsed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"agentAddresses\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"role\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"image\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"nonce\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"nameOf\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"removeSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"roleOf\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"signerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"tokenIdOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agentDescriptor\",\"type\":\"address\"}],\"name\":\"updateAgentDescriptor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"newSigner\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"nonce\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"updateTokenSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"walletsOf\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x610160604052600160095534801561001657600080fd5b5060405161544e38038061544e83398181016040528101906100389190610528565b6040518060400160405280601881526020017f4167656e74696341494167656e74436f6c6c656374696f6e00000000000000008152506040518060400160405280600181526020017f31000000000000000000000000000000000000000000000000000000000000008152506040518060400160405280601881526020017f4167656e74696341494167656e74436f6c6c656374696f6e00000000000000008152506040518060400160405280600781526020017f41494167656e7400000000000000000000000000000000000000000000000000815250816000908161011f91906107a5565b50806001908161012f91906107a5565b50505061014660068361023760201b90919060201c565b610120818152505061016260078261023760201b90919060201c565b6101408181525050818051906020012060e08181525050808051906020012061010081815250504660a0818152505061019f61028760201b60201c565b608081815250503073ffffffffffffffffffffffffffffffffffffffff1660c08173ffffffffffffffffffffffffffffffffffffffff1681525050505080601260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506102306000801b336102e260201b60201c565b5050610a3f565b600060208351101561025957610252836103e060201b60201c565b9050610281565b826102698361044860201b60201c565b600001908161027891906107a5565b5060ff60001b90505b92915050565b60007f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f60e0516101005146306040516020016102c79594939291906108ae565b60405160208183030381529060405280519060200120905090565b60006102f4838361045260201b60201c565b6103d55760016008600085815260200190815260200160002060000160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055506103726104bd60201b60201c565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a4600190506103da565b600090505b92915050565b600080829050601f8151111561042d57826040517f305a27a90000000000000000000000000000000000000000000000000000000081526004016104249190610986565b60405180910390fd5b805181610439906109d8565b60001c1760001b915050919050565b6000819050919050565b60006008600084815260200190815260200160002060000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b600033905090565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006104f5826104ca565b9050919050565b610505816104ea565b811461051057600080fd5b50565b600081519050610522816104fc565b92915050565b60006020828403121561053e5761053d6104c5565b5b600061054c84828501610513565b91505092915050565b600081519050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806105d657607f821691505b6020821081036105e9576105e861058f565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026106517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610614565b61065b8683610614565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b60006106a261069d61069884610673565b61067d565b610673565b9050919050565b6000819050919050565b6106bc83610687565b6106d06106c8826106a9565b848454610621565b825550505050565b600090565b6106e56106d8565b6106f08184846106b3565b505050565b5b81811015610714576107096000826106dd565b6001810190506106f6565b5050565b601f8211156107595761072a816105ef565b61073384610604565b81016020851015610742578190505b61075661074e85610604565b8301826106f5565b50505b505050565b600082821c905092915050565b600061077c6000198460080261075e565b1980831691505092915050565b6000610795838361076b565b9150826002028217905092915050565b6107ae82610555565b67ffffffffffffffff8111156107c7576107c6610560565b5b6107d182546105be565b6107dc828285610718565b600060209050601f83116001811461080f57600084156107fd578287015190505b6108078582610789565b86555061086f565b601f19841661081d866105ef565b60005b8281101561084557848901518255600182019150602085019450602081019050610820565b86831015610862578489015161085e601f89168261076b565b8355505b6001600288020188555050505b505050505050565b6000819050919050565b61088a81610877565b82525050565b61089981610673565b82525050565b6108a8816104ea565b82525050565b600060a0820190506108c36000830188610881565b6108d06020830187610881565b6108dd6040830186610881565b6108ea6060830185610890565b6108f7608083018461089f565b9695505050505050565b600082825260208201905092915050565b60005b83811015610930578082015181840152602081019050610915565b60008484015250505050565b6000601f19601f8301169050919050565b600061095882610555565b6109628185610901565b9350610972818560208601610912565b61097b8161093c565b840191505092915050565b600060208201905081810360008301526109a0818461094d565b905092915050565b600081519050919050565b6000819050602082019050919050565b60006109cf8251610877565b80915050919050565b60006109e3826109a8565b826109ed846109b3565b90506109f8816109c3565b92506020821015610a3857610a337fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff83602003600802610614565b831692505b5050919050565b60805160a05160c05160e0516101005161012051610140516149b5610a9960003960006121b60152600061217b015260006132490152600061322801526000612e0601526000612e5c01526000612e8501526149b56000f3fe608060405234801561001057600080fd5b50600436106102115760003560e01c806375794a3c11610125578063a398d819116100ad578063d06075621161007c578063d060756214610688578063d547741f146106b8578063e985e9c5146106d4578063eb12d61e14610704578063eeadae2a1461072057610211565b8063a398d819146105dc578063b88d4fde1461060c578063b90665e514610628578063c87b56dd1461065857610211565b806391d14854116100f457806391d148541461053657806395d89b4114610566578063a1ebf35d14610584578063a217fddf146105a2578063a22cb465146105c057610211565b806375794a3c146104a857806383648671146104c657806384b0196e146104f65780638c2c7ec51461051a57610211565b80632f2ff15d116101a8578063589a174311610177578063589a1743146103cc5780636352211e146103fc5780636457f5031461042c57806370a082311461044857806372b1d8cf1461047857610211565b80632f2ff15d1461034857806336568abe1461036457806342842e0e146103805780635161fdf51461039c57610211565b8063095ea7b3116101e4578063095ea7b3146102c45780630e316ab7146102e057806323b872dd146102fc578063248a9ca31461031857610211565b806301ffc9a714610216578063051a26641461024657806306fdde0314610276578063081812fc14610294575b600080fd5b610230600480360381019061022b91906134a0565b61073c565b60405161023d91906134e8565b60405180910390f35b610260600480360381019061025b9190613539565b61075e565b60405161026d91906135f6565b60405180910390f35b61027e61080d565b60405161028b91906135f6565b60405180910390f35b6102ae60048036038101906102a99190613539565b61089f565b6040516102bb9190613659565b60405180910390f35b6102de60048036038101906102d991906136a0565b6108bb565b005b6102fa60048036038101906102f591906136e0565b6108d1565b005b6103166004803603810190610311919061370d565b610944565b005b610332600480360381019061032d9190613796565b610a46565b60405161033f91906137d2565b60405180910390f35b610362600480360381019061035d91906137ed565b610a66565b005b61037e600480360381019061037991906137ed565b610a88565b005b61039a6004803603810190610395919061370d565b610b03565b005b6103b660048036038101906103b19190613539565b610b23565b6040516103c39190613659565b60405180910390f35b6103e660048036038101906103e19190613539565b610b6a565b6040516103f39190613659565b60405180910390f35b61041660048036038101906104119190613539565b610bb1565b6040516104239190613659565b60405180910390f35b610446600480360381019061044191906136e0565b610bc3565b005b610462600480360381019061045d91906136e0565b610c15565b60405161046f919061383c565b60405180910390f35b610492600480360381019061048d9190613539565b610ccf565b60405161049f91906135f6565b60405180910390f35b6104b0610d7e565b6040516104bd919061383c565b60405180910390f35b6104e060048036038101906104db9190613539565b610d88565b6040516104ed9190613915565b60405180910390f35b6104fe610e33565b6040516105119796959493929190613a30565b60405180910390f35b610534600480360381019061052f9190613be9565b610edd565b005b610550600480360381019061054b91906137ed565b611216565b60405161055d91906134e8565b60405180910390f35b61056e611281565b60405161057b91906135f6565b60405180910390f35b61058c611313565b60405161059991906137d2565b60405180910390f35b6105aa611337565b6040516105b791906137d2565b60405180910390f35b6105da60048036038101906105d59190613c98565b61133e565b005b6105f660048036038101906105f19190613d79565b611354565b604051610603919061383c565b60405180910390f35b61062660048036038101906106219190613dc2565b61137c565b005b610642600480360381019061063d9190613796565b6113a1565b60405161064f91906134e8565b60405180910390f35b610672600480360381019061066d9190613539565b6113cb565b60405161067f91906135f6565b60405180910390f35b6106a2600480360381019061069d9190613539565b611481565b6040516106af91906135f6565b60405180910390f35b6106d260048036038101906106cd91906137ed565b611530565b005b6106ee60048036038101906106e99190613e45565b611552565b6040516106fb91906134e8565b60405180910390f35b61071e600480360381019061071991906136e0565b6115e6565b005b61073a60048036038101906107359190613f4d565b611659565b005b600061074782611ac0565b80610757575061075682611ba2565b5b9050919050565b606061076982611c1c565b50600e60008381526020019081526020016000208054610788906140be565b80601f01602080910402602001604051908101604052809291908181526020018280546107b4906140be565b80156108015780601f106107d657610100808354040283529160200191610801565b820191906000526020600020905b8154815290600101906020018083116107e457829003601f168201915b50505050509050919050565b60606000805461081c906140be565b80601f0160208091040260200160405190810160405280929190818152602001828054610848906140be565b80156108955780601f1061086a57610100808354040283529160200191610895565b820191906000526020600020905b81548152906001019060200180831161087857829003601f168201915b5050505050905090565b60006108aa82611c80565b506108b482611d08565b9050919050565b6108cd82826108c8611d45565b611d4d565b5050565b6000801b6108de81611d5f565b6109087fe2f4eaae4a9751e85a3e4a7b9587827a877f29914755229b07a7b2da98285f7083611d73565b507f3525e22824a8a7df2c9a6029941c824cf95b6447f1e13d5128fd3826d35afe8b826040516109389190613659565b60405180910390a15050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036109b65760006040517f64a0ae920000000000000000000000000000000000000000000000000000000081526004016109ad9190613659565b60405180910390fd5b60006109ca83836109c5611d45565b611e66565b90508373ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610a40578382826040517f64283d7b000000000000000000000000000000000000000000000000000000008152600401610a37939291906140ef565b60405180910390fd5b50505050565b600060086000838152602001908152602001600020600101549050919050565b610a6f82610a46565b610a7881611d5f565b610a828383612080565b50505050565b610a90611d45565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610af4576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610afe8282611d73565b505050565b610b1e8383836040518060200160405280600081525061137c565b505050565b6000610b2e82611c1c565b50600b600083815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b6000610b7582611c1c565b50600c600083815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b6000610bbc82611c80565b9050919050565b6000801b610bd081611d5f565b81601260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b60008073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610c885760006040517f89c62b64000000000000000000000000000000000000000000000000000000008152600401610c7f9190613659565b60405180910390fd5b600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b6060610cda82611c1c565b50600a60008381526020019081526020016000208054610cf9906140be565b80601f0160208091040260200160405190810160405280929190818152602001828054610d25906140be565b8015610d725780601f10610d4757610100808354040283529160200191610d72565b820191906000526020600020905b815481529060010190602001808311610d5557829003601f168201915b50505050509050919050565b6000600954905090565b6060610d9382611c1c565b50600d6000838152602001908152602001600020805480602002602001604051908101604052809291908181526020018280548015610e2757602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311610ddd575b50505050509050919050565b600060608060008060006060610e47612172565b610e4f6121ad565b46306000801b600067ffffffffffffffff811115610e7057610e6f613abe565b5b604051908082528060200260200182016040528015610e9e5781602001602082028036833780820191505090505b507f0f00000000000000000000000000000000000000000000000000000000000000959493929190965096509650965096509650965090919293949596565b6011600083815260200190815260200160002060009054906101000a900460ff1615610f4057816040517f5c146343000000000000000000000000000000000000000000000000000000008152600401610f3791906137d2565b60405180910390fd5b610f48611d45565b73ffffffffffffffffffffffffffffffffffffffff16610f6785610bb1565b73ffffffffffffffffffffffffffffffffffffffff1614610fc857610f8a611d45565b846040517f5c3c03c3000000000000000000000000000000000000000000000000000000008152600401610fbf929190614126565b60405180910390fd5b610ff27fe2f4eaae4a9751e85a3e4a7b9587827a877f29914755229b07a7b2da98285f7084611216565b61103357826040517f2a6edb2b00000000000000000000000000000000000000000000000000000000815260040161102a9190613659565b60405180910390fd5b6000600b600086815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508373ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036110db57806040517fa13983290000000000000000000000000000000000000000000000000000000081526004016110d29190613659565b60405180910390fd5b60006110e9868686866121e8565b90506111157fe2f4eaae4a9751e85a3e4a7b9587827a877f29914755229b07a7b2da98285f7082611216565b61115657806040517f2a6edb2b00000000000000000000000000000000000000000000000000000000815260040161114d9190613659565b60405180910390fd5b84600b600088815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060016011600086815260200190815260200160002060006101000a81548160ff021916908315150217905550857fe51e6ee1b7a3c7c5d0e801f5306ab38325b6cf7d09e0dda2cf4e9bb59c00fcba838760405161120692919061414f565b60405180910390a2505050505050565b60006008600084815260200190815260200160002060000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b606060018054611290906140be565b80601f01602080910402602001604051908101604052809291908181526020018280546112bc906140be565b80156113095780601f106112de57610100808354040283529160200191611309565b820191906000526020600020905b8154815290600101906020018083116112ec57829003601f168201915b5050505050905090565b7fe2f4eaae4a9751e85a3e4a7b9587827a877f29914755229b07a7b2da98285f7081565b6000801b81565b611350611349611d45565b838361225d565b5050565b600060108260405161136691906141b4565b9081526020016040518091039020549050919050565b611387848484610944565b61139b611392611d45565b858585856123cc565b50505050565b60006011600083815260200190815260200160002060009054906101000a900460ff169050919050565b60606113d682611c1c565b50601260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663b25457d330846040518363ffffffff1660e01b815260040161143492919061422a565b600060405180830381865afa158015611451573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525081019061147a91906142c3565b9050919050565b606061148c82611c1c565b50600f600083815260200190815260200160002080546114ab906140be565b80601f01602080910402602001604051908101604052809291908181526020018280546114d7906140be565b80156115245780601f106114f957610100808354040283529160200191611524565b820191906000526020600020905b81548152906001019060200180831161150757829003601f168201915b50505050509050919050565b61153982610a46565b61154281611d5f565b61154c8383611d73565b50505050565b6000600560008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b6000801b6115f381611d5f565b61161d7fe2f4eaae4a9751e85a3e4a7b9587827a877f29914755229b07a7b2da98285f7083612080565b507f47d1c22a25bb3a5d4e481b9b1e6944c2eade3181a0a20b495ed61d35b5323f248260405161164d9190613659565b60405180910390a15050565b6116628461257d565b6116a357836040517f7f19f48d00000000000000000000000000000000000000000000000000000000815260040161169a91906135f6565b60405180910390fd5b6011600083815260200190815260200160002060009054906101000a900460ff161561170657816040517f5c1463430000000000000000000000000000000000000000000000000000000081526004016116fd91906137d2565b60405180910390fd5b600060108560405161171891906141b4565b9081526020016040518091039020541461176957836040517f36cd42c300000000000000000000000000000000000000000000000000000000815260040161176091906135f6565b60405180910390fd5b600061177b89898989898989896126e3565b90506117a77fe2f4eaae4a9751e85a3e4a7b9587827a877f29914755229b07a7b2da98285f7082611216565b6117e857806040517f2a6edb2b0000000000000000000000000000000000000000000000000000000081526004016117df9190613659565b60405180910390fd5b8673ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161461185a5780876040517f7ba5ffb500000000000000000000000000000000000000000000000000000000815260040161185192919061414f565b60405180910390fd5b60016011600085815260200190815260200160002060006101000a81548160ff02191690831515021790555085600a6000600954815260200190815260200160002090816118a891906144ae565b5084600e6000600954815260200190815260200160002090816118cb91906144ae565b5086600b6000600954815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550611928611d45565b600c6000600954815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555087600d6000600954815260200190815260200160002090805190602001906119a492919061338d565b5083600f6000600954815260200190815260200160002090816119c791906144ae565b506009546010866040516119db91906141b4565b9081526020016040518091039020819055506119f989600954612814565b8673ffffffffffffffffffffffffffffffffffffffff168973ffffffffffffffffffffffffffffffffffffffff167f9d228d69b5fdb8d273a2336f8fb8612d039631024ea9bf09c424a9503aa078f0600954604051611a58919061383c565b60405180910390a36009547fe51e6ee1b7a3c7c5d0e801f5306ab38325b6cf7d09e0dda2cf4e9bb59c00fcba600089604051611a9592919061414f565b60405180910390a260096000815480929190611ab0906145af565b9190505550505050505050505050565b60007f80ac58cd000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19161480611b8b57507f5b5e139f000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916145b80611b9b5750611b9a82612832565b5b9050919050565b60007f7965db0b000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19161480611c155750611c1482611ac0565b5b9050919050565b600080821480611c2e57506009548210155b15611c7057816040517f7e273289000000000000000000000000000000000000000000000000000000008152600401611c67919061383c565b60405180910390fd5b611c798261289c565b9050919050565b600080611c8c8361289c565b9050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603611cff57826040517f7e273289000000000000000000000000000000000000000000000000000000008152600401611cf6919061383c565b60405180910390fd5b80915050919050565b60006004600083815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b600033905090565b611d5a83838360016128d9565b505050565b611d7081611d6b611d45565b612a9e565b50565b6000611d7f8383611216565b15611e5b5760006008600085815260200190815260200160002060000160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550611df8611d45565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16847ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b60405160405180910390a460019050611e60565b600090505b92915050565b600080611e728461289c565b9050600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614611eb457611eb3818486612aef565b5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614611f4557611ef66000856000806128d9565b6001600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825403925050819055505b600073ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff1614611fc8576001600360008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825401925050819055505b846002600086815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550838573ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60405160405180910390a4809150509392505050565b600061208c8383611216565b6121675760016008600085815260200190815260200160002060000160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550612104611d45565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a46001905061216c565b600090505b92915050565b60606121a860067f0000000000000000000000000000000000000000000000000000000000000000612bb390919063ffffffff16565b905090565b60606121e360077f0000000000000000000000000000000000000000000000000000000000000000612bb390919063ffffffff16565b905090565b6000807f26b808267d3374b647158bad9a412514863c285b07dc6f424271401dc84e000786868660405160200161222294939291906145f7565b604051602081830303815290604052805190602001209050600061224582612c63565b90506122518185612c7d565b92505050949350505050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036122ce57816040517f5b08ba180000000000000000000000000000000000000000000000000000000081526004016122c59190613659565b60405180910390fd5b80600560008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31836040516123bf91906134e8565b60405180910390a3505050565b60008373ffffffffffffffffffffffffffffffffffffffff163b1115612576578273ffffffffffffffffffffffffffffffffffffffff1663150b7a02868685856040518563ffffffff1660e01b815260040161242b9493929190614691565b6020604051808303816000875af192505050801561246757506040513d601f19601f8201168201806040525081019061246491906146f2565b60015b6124eb573d8060008114612497576040519150601f19603f3d011682016040523d82523d6000602084013e61249c565b606091505b5060008151036124e357836040517f64a0ae920000000000000000000000000000000000000000000000000000000081526004016124da9190613659565b60405180910390fd5b805181602001fd5b63150b7a0260e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19161461257457836040517f64a0ae9200000000000000000000000000000000000000000000000000000000815260040161256b9190613659565b60405180910390fd5b505b5050505050565b60008082905060008151036125965760009150506126de565b60005b81518110156126d75760008282815181106125b7576125b661471f565b5b602001015160f81c60f81b9050606160f81b817effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916101580156126205750607a60f81b817effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191611155b1580156126865750603060f81b817effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916101580156126845750603960f81b817effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191611155b155b80156126b85750602d60f81b817effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614155b156126c957600093505050506126de565b508080600101915050612599565b5060019150505b919050565b600080886040516020016126f791906147de565b60405160208183030381529060405280519060200120905060008660405160200161272291906141b4565b60405160208183030381529060405280519060200120905060008660405160200161274d91906141b4565b60405160208183030381529060405280519060200120905060008960405160200161277891906141b4565b60405160208183030381529060405280519060200120905060007f53b22f5b00b43a987f3c0816840e3890870cf7b6460aceed6df717b4b0a3f0938e868e8588888e6040516020016127d19897969594939291906147f5565b60405160208183030381529060405280519060200120905060006127f482612c63565b90506128008189612c7d565b965050505050505098975050505050505050565b61282e828260405180602001604052806000815250612ca9565b5050565b60007f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b60006002600083815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b80806129125750600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614155b15612a4657600061292284611c80565b9050600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415801561298d57508273ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614155b80156129a0575061299e8184611552565b155b156129e257826040517fa9fbf51f0000000000000000000000000000000000000000000000000000000081526004016129d99190613659565b60405180910390fd5b8115612a4457838573ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45b505b836004600085815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050505050565b612aa88282611216565b612aeb5780826040517fe2517d3f000000000000000000000000000000000000000000000000000000008152600401612ae2929190614873565b60405180910390fd5b5050565b612afa838383612ccd565b612bae57600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603612b6f57806040517f7e273289000000000000000000000000000000000000000000000000000000008152600401612b66919061383c565b60405180910390fd5b81816040517f177e802f000000000000000000000000000000000000000000000000000000008152600401612ba5929190614126565b60405180910390fd5b505050565b606060ff60001b8314612bd057612bc983612d8e565b9050612c5d565b818054612bdc906140be565b80601f0160208091040260200160405190810160405280929190818152602001828054612c08906140be565b8015612c555780601f10612c2a57610100808354040283529160200191612c55565b820191906000526020600020905b815481529060010190602001808311612c3857829003601f168201915b505050505090505b92915050565b6000612c76612c70612e02565b83612eb9565b9050919050565b600080600080612c8d8686612efa565b925092509250612c9d8282612f56565b82935050505092915050565b612cb383836130ba565b612cc8612cbe611d45565b60008585856123cc565b505050565b60008073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614158015612d8557508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff161480612d465750612d458484611552565b5b80612d8457508273ffffffffffffffffffffffffffffffffffffffff16612d6c83611d08565b73ffffffffffffffffffffffffffffffffffffffff16145b5b90509392505050565b60606000612d9b836131b3565b90506000602067ffffffffffffffff811115612dba57612db9613abe565b5b6040519080825280601f01601f191660200182016040528015612dec5781602001600182028036833780820191505090505b5090508181528360208201528092505050919050565b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff16148015612e7e57507f000000000000000000000000000000000000000000000000000000000000000046145b15612eab577f00000000000000000000000000000000000000000000000000000000000000009050612eb6565b612eb3613203565b90505b90565b60006040517f190100000000000000000000000000000000000000000000000000000000000081528360028201528260228201526042812091505092915050565b60008060006041845103612f3f5760008060006020870151925060408701519150606087015160001a9050612f3188828585613299565b955095509550505050612f4f565b60006002855160001b9250925092505b9250925092565b60006003811115612f6a57612f6961489c565b5b826003811115612f7d57612f7c61489c565b5b03156130b65760016003811115612f9757612f9661489c565b5b826003811115612faa57612fa961489c565b5b03612fe1576040517ff645eedf00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60026003811115612ff557612ff461489c565b5b8260038111156130085761300761489c565b5b0361304d578060001c6040517ffce698f7000000000000000000000000000000000000000000000000000000008152600401613044919061383c565b60405180910390fd5b6003808111156130605761305f61489c565b5b8260038111156130735761307261489c565b5b036130b557806040517fd78bce0c0000000000000000000000000000000000000000000000000000000081526004016130ac91906137d2565b60405180910390fd5b5b5050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff160361312c5760006040517f64a0ae920000000000000000000000000000000000000000000000000000000081526004016131239190613659565b60405180910390fd5b600061313a83836000611e66565b9050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16146131ae5760006040517f73c6ac6e0000000000000000000000000000000000000000000000000000000081526004016131a59190613659565b60405180910390fd5b505050565b60008060ff8360001c169050601f8111156131fa576040517fb3512b0c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80915050919050565b60007f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f7f00000000000000000000000000000000000000000000000000000000000000007f0000000000000000000000000000000000000000000000000000000000000000463060405160200161327e9594939291906148cb565b60405160208183030381529060405280519060200120905090565b60008060007f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08460001c11156132d9576000600385925092509250613383565b6000600188888888604051600081526020016040526040516132fe949392919061493a565b6020604051602081039080840390855afa158015613320573d6000803e3d6000fd5b505050602060405103519050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff160361337457600060016000801b93509350935050613383565b8060008060001b935093509350505b9450945094915050565b828054828255906000526020600020908101928215613406579160200282015b828111156134055782518260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550916020019190600101906133ad565b5b5090506134139190613417565b5090565b5b80821115613430576000816000905550600101613418565b5090565b6000604051905090565b600080fd5b600080fd5b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b61347d81613448565b811461348857600080fd5b50565b60008135905061349a81613474565b92915050565b6000602082840312156134b6576134b561343e565b5b60006134c48482850161348b565b91505092915050565b60008115159050919050565b6134e2816134cd565b82525050565b60006020820190506134fd60008301846134d9565b92915050565b6000819050919050565b61351681613503565b811461352157600080fd5b50565b6000813590506135338161350d565b92915050565b60006020828403121561354f5761354e61343e565b5b600061355d84828501613524565b91505092915050565b600081519050919050565b600082825260208201905092915050565b60005b838110156135a0578082015181840152602081019050613585565b60008484015250505050565b6000601f19601f8301169050919050565b60006135c882613566565b6135d28185613571565b93506135e2818560208601613582565b6135eb816135ac565b840191505092915050565b6000602082019050818103600083015261361081846135bd565b905092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061364382613618565b9050919050565b61365381613638565b82525050565b600060208201905061366e600083018461364a565b92915050565b61367d81613638565b811461368857600080fd5b50565b60008135905061369a81613674565b92915050565b600080604083850312156136b7576136b661343e565b5b60006136c58582860161368b565b92505060206136d685828601613524565b9150509250929050565b6000602082840312156136f6576136f561343e565b5b60006137048482850161368b565b91505092915050565b6000806000606084860312156137265761372561343e565b5b60006137348682870161368b565b93505060206137458682870161368b565b925050604061375686828701613524565b9150509250925092565b6000819050919050565b61377381613760565b811461377e57600080fd5b50565b6000813590506137908161376a565b92915050565b6000602082840312156137ac576137ab61343e565b5b60006137ba84828501613781565b91505092915050565b6137cc81613760565b82525050565b60006020820190506137e760008301846137c3565b92915050565b600080604083850312156138045761380361343e565b5b600061381285828601613781565b92505060206138238582860161368b565b9150509250929050565b61383681613503565b82525050565b6000602082019050613851600083018461382d565b92915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b61388c81613638565b82525050565b600061389e8383613883565b60208301905092915050565b6000602082019050919050565b60006138c282613857565b6138cc8185613862565b93506138d783613873565b8060005b838110156139085781516138ef8882613892565b97506138fa836138aa565b9250506001810190506138db565b5085935050505092915050565b6000602082019050818103600083015261392f81846138b7565b905092915050565b60007fff0000000000000000000000000000000000000000000000000000000000000082169050919050565b61396c81613937565b82525050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b6139a781613503565b82525050565b60006139b9838361399e565b60208301905092915050565b6000602082019050919050565b60006139dd82613972565b6139e7818561397d565b93506139f28361398e565b8060005b83811015613a23578151613a0a88826139ad565b9750613a15836139c5565b9250506001810190506139f6565b5085935050505092915050565b600060e082019050613a45600083018a613963565b8181036020830152613a5781896135bd565b90508181036040830152613a6b81886135bd565b9050613a7a606083018761382d565b613a87608083018661364a565b613a9460a08301856137c3565b81810360c0830152613aa681846139d2565b905098975050505050505050565b600080fd5b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b613af6826135ac565b810181811067ffffffffffffffff82111715613b1557613b14613abe565b5b80604052505050565b6000613b28613434565b9050613b348282613aed565b919050565b600067ffffffffffffffff821115613b5457613b53613abe565b5b613b5d826135ac565b9050602081019050919050565b82818337600083830152505050565b6000613b8c613b8784613b39565b613b1e565b905082815260208101848484011115613ba857613ba7613ab9565b5b613bb3848285613b6a565b509392505050565b600082601f830112613bd057613bcf613ab4565b5b8135613be0848260208601613b79565b91505092915050565b60008060008060808587031215613c0357613c0261343e565b5b6000613c1187828801613524565b9450506020613c228782880161368b565b9350506040613c3387828801613781565b925050606085013567ffffffffffffffff811115613c5457613c53613443565b5b613c6087828801613bbb565b91505092959194509250565b613c75816134cd565b8114613c8057600080fd5b50565b600081359050613c9281613c6c565b92915050565b60008060408385031215613caf57613cae61343e565b5b6000613cbd8582860161368b565b9250506020613cce85828601613c83565b9150509250929050565b600067ffffffffffffffff821115613cf357613cf2613abe565b5b613cfc826135ac565b9050602081019050919050565b6000613d1c613d1784613cd8565b613b1e565b905082815260208101848484011115613d3857613d37613ab9565b5b613d43848285613b6a565b509392505050565b600082601f830112613d6057613d5f613ab4565b5b8135613d70848260208601613d09565b91505092915050565b600060208284031215613d8f57613d8e61343e565b5b600082013567ffffffffffffffff811115613dad57613dac613443565b5b613db984828501613d4b565b91505092915050565b60008060008060808587031215613ddc57613ddb61343e565b5b6000613dea8782880161368b565b9450506020613dfb8782880161368b565b9350506040613e0c87828801613524565b925050606085013567ffffffffffffffff811115613e2d57613e2c613443565b5b613e3987828801613bbb565b91505092959194509250565b60008060408385031215613e5c57613e5b61343e565b5b6000613e6a8582860161368b565b9250506020613e7b8582860161368b565b9150509250929050565b600067ffffffffffffffff821115613ea057613e9f613abe565b5b602082029050602081019050919050565b600080fd5b6000613ec9613ec484613e85565b613b1e565b90508083825260208201905060208402830185811115613eec57613eeb613eb1565b5b835b81811015613f155780613f01888261368b565b845260208401935050602081019050613eee565b5050509392505050565b600082601f830112613f3457613f33613ab4565b5b8135613f44848260208601613eb6565b91505092915050565b600080600080600080600080610100898b031215613f6e57613f6d61343e565b5b6000613f7c8b828c0161368b565b985050602089013567ffffffffffffffff811115613f9d57613f9c613443565b5b613fa98b828c01613f1f565b9750506040613fba8b828c0161368b565b965050606089013567ffffffffffffffff811115613fdb57613fda613443565b5b613fe78b828c01613d4b565b955050608089013567ffffffffffffffff81111561400857614007613443565b5b6140148b828c01613d4b565b94505060a089013567ffffffffffffffff81111561403557614034613443565b5b6140418b828c01613d4b565b93505060c06140528b828c01613781565b92505060e089013567ffffffffffffffff81111561407357614072613443565b5b61407f8b828c01613bbb565b9150509295985092959890939650565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806140d657607f821691505b6020821081036140e9576140e861408f565b5b50919050565b6000606082019050614104600083018661364a565b614111602083018561382d565b61411e604083018461364a565b949350505050565b600060408201905061413b600083018561364a565b614148602083018461382d565b9392505050565b6000604082019050614164600083018561364a565b614171602083018461364a565b9392505050565b600081905092915050565b600061418e82613566565b6141988185614178565b93506141a8818560208601613582565b80840191505092915050565b60006141c08284614183565b915081905092915050565b6000819050919050565b60006141f06141eb6141e684613618565b6141cb565b613618565b9050919050565b6000614202826141d5565b9050919050565b6000614214826141f7565b9050919050565b61422481614209565b82525050565b600060408201905061423f600083018561421b565b61424c602083018461382d565b9392505050565b600061426661426184613cd8565b613b1e565b90508281526020810184848401111561428257614281613ab9565b5b61428d848285613582565b509392505050565b600082601f8301126142aa576142a9613ab4565b5b81516142ba848260208601614253565b91505092915050565b6000602082840312156142d9576142d861343e565b5b600082015167ffffffffffffffff8111156142f7576142f6613443565b5b61430384828501614295565b91505092915050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b60006008830261436e7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82614331565b6143788683614331565b95508019841693508086168417925050509392505050565b60006143ab6143a66143a184613503565b6141cb565b613503565b9050919050565b6000819050919050565b6143c583614390565b6143d96143d1826143b2565b84845461433e565b825550505050565b600090565b6143ee6143e1565b6143f98184846143bc565b505050565b5b8181101561441d576144126000826143e6565b6001810190506143ff565b5050565b601f821115614462576144338161430c565b61443c84614321565b8101602085101561444b578190505b61445f61445785614321565b8301826143fe565b50505b505050565b600082821c905092915050565b600061448560001984600802614467565b1980831691505092915050565b600061449e8383614474565b9150826002028217905092915050565b6144b782613566565b67ffffffffffffffff8111156144d0576144cf613abe565b5b6144da82546140be565b6144e5828285614421565b600060209050601f8311600181146145185760008415614506578287015190505b6145108582614492565b865550614578565b601f1984166145268661430c565b60005b8281101561454e57848901518255600182019150602085019450602081019050614529565b8683101561456b5784890151614567601f891682614474565b8355505b6001600288020188555050505b505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006145ba82613503565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036145ec576145eb614580565b5b600182019050919050565b600060808201905061460c60008301876137c3565b614619602083018661382d565b614626604083018561364a565b61463360608301846137c3565b95945050505050565b600081519050919050565b600082825260208201905092915050565b60006146638261463c565b61466d8185614647565b935061467d818560208601613582565b614686816135ac565b840191505092915050565b60006080820190506146a6600083018761364a565b6146b3602083018661364a565b6146c0604083018561382d565b81810360608301526146d28184614658565b905095945050505050565b6000815190506146ec81613474565b92915050565b6000602082840312156147085761470761343e565b5b6000614716848285016146dd565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600081905092915050565b61476281613638565b82525050565b60006147748383614759565b60208301905092915050565b600061478b82613857565b614795818561474e565b93506147a083613873565b8060005b838110156147d15781516147b88882614768565b97506147c3836138aa565b9250506001810190506147a4565b5085935050505092915050565b60006147ea8284614780565b915081905092915050565b60006101008201905061480b600083018b6137c3565b614818602083018a61364a565b61482560408301896137c3565b614832606083018861364a565b61483f60808301876137c3565b61484c60a08301866137c3565b61485960c08301856137c3565b61486660e08301846137c3565b9998505050505050505050565b6000604082019050614888600083018561364a565b61489560208301846137c3565b9392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b600060a0820190506148e060008301886137c3565b6148ed60208301876137c3565b6148fa60408301866137c3565b614907606083018561382d565b614914608083018461364a565b9695505050505050565b600060ff82169050919050565b6149348161491e565b82525050565b600060808201905061494f60008301876137c3565b61495c602083018661492b565b61496960408301856137c3565b61497660608301846137c3565b9594505050505056fea2646970667358221220918be4a71c411c6fdd7c4b072ba49689d4747e5aaf4d1e4b24fcd47dd974aa0f64736f6c634300081c0033",
}

// AgenticAIAgentCollectionABI is the input ABI used to generate the binding from.
// Deprecated: Use AgenticAIAgentCollectionMetaData.ABI instead.
var AgenticAIAgentCollectionABI = AgenticAIAgentCollectionMetaData.ABI

// AgenticAIAgentCollectionBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AgenticAIAgentCollectionMetaData.Bin instead.
var AgenticAIAgentCollectionBin = AgenticAIAgentCollectionMetaData.Bin

// DeployAgenticAIAgentCollection deploys a new Ethereum contract, binding an instance of AgenticAIAgentCollection to it.
func DeployAgenticAIAgentCollection(auth *bind.TransactOpts, backend bind.ContractBackend, agentDescriptor common.Address) (common.Address, *types.Transaction, *AgenticAIAgentCollection, error) {
	parsed, err := AgenticAIAgentCollectionMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AgenticAIAgentCollectionBin), backend, agentDescriptor)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AgenticAIAgentCollection{AgenticAIAgentCollectionCaller: AgenticAIAgentCollectionCaller{contract: contract}, AgenticAIAgentCollectionTransactor: AgenticAIAgentCollectionTransactor{contract: contract}, AgenticAIAgentCollectionFilterer: AgenticAIAgentCollectionFilterer{contract: contract}}, nil
}

// AgenticAIAgentCollection is an auto generated Go binding around an Ethereum contract.
type AgenticAIAgentCollection struct {
	AgenticAIAgentCollectionCaller     // Read-only binding to the contract
	AgenticAIAgentCollectionTransactor // Write-only binding to the contract
	AgenticAIAgentCollectionFilterer   // Log filterer for contract events
}

// AgenticAIAgentCollectionCaller is an auto generated read-only Go binding around an Ethereum contract.
type AgenticAIAgentCollectionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgenticAIAgentCollectionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AgenticAIAgentCollectionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgenticAIAgentCollectionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AgenticAIAgentCollectionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgenticAIAgentCollectionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AgenticAIAgentCollectionSession struct {
	Contract     *AgenticAIAgentCollection // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// AgenticAIAgentCollectionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AgenticAIAgentCollectionCallerSession struct {
	Contract *AgenticAIAgentCollectionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// AgenticAIAgentCollectionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AgenticAIAgentCollectionTransactorSession struct {
	Contract     *AgenticAIAgentCollectionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// AgenticAIAgentCollectionRaw is an auto generated low-level Go binding around an Ethereum contract.
type AgenticAIAgentCollectionRaw struct {
	Contract *AgenticAIAgentCollection // Generic contract binding to access the raw methods on
}

// AgenticAIAgentCollectionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AgenticAIAgentCollectionCallerRaw struct {
	Contract *AgenticAIAgentCollectionCaller // Generic read-only contract binding to access the raw methods on
}

// AgenticAIAgentCollectionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AgenticAIAgentCollectionTransactorRaw struct {
	Contract *AgenticAIAgentCollectionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAgenticAIAgentCollection creates a new instance of AgenticAIAgentCollection, bound to a specific deployed contract.
func NewAgenticAIAgentCollection(address common.Address, backend bind.ContractBackend) (*AgenticAIAgentCollection, error) {
	contract, err := bindAgenticAIAgentCollection(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AgenticAIAgentCollection{AgenticAIAgentCollectionCaller: AgenticAIAgentCollectionCaller{contract: contract}, AgenticAIAgentCollectionTransactor: AgenticAIAgentCollectionTransactor{contract: contract}, AgenticAIAgentCollectionFilterer: AgenticAIAgentCollectionFilterer{contract: contract}}, nil
}

// NewAgenticAIAgentCollectionCaller creates a new read-only instance of AgenticAIAgentCollection, bound to a specific deployed contract.
func NewAgenticAIAgentCollectionCaller(address common.Address, caller bind.ContractCaller) (*AgenticAIAgentCollectionCaller, error) {
	contract, err := bindAgenticAIAgentCollection(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AgenticAIAgentCollectionCaller{contract: contract}, nil
}

// NewAgenticAIAgentCollectionTransactor creates a new write-only instance of AgenticAIAgentCollection, bound to a specific deployed contract.
func NewAgenticAIAgentCollectionTransactor(address common.Address, transactor bind.ContractTransactor) (*AgenticAIAgentCollectionTransactor, error) {
	contract, err := bindAgenticAIAgentCollection(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AgenticAIAgentCollectionTransactor{contract: contract}, nil
}

// NewAgenticAIAgentCollectionFilterer creates a new log filterer instance of AgenticAIAgentCollection, bound to a specific deployed contract.
func NewAgenticAIAgentCollectionFilterer(address common.Address, filterer bind.ContractFilterer) (*AgenticAIAgentCollectionFilterer, error) {
	contract, err := bindAgenticAIAgentCollection(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AgenticAIAgentCollectionFilterer{contract: contract}, nil
}

// bindAgenticAIAgentCollection binds a generic wrapper to an already deployed contract.
func bindAgenticAIAgentCollection(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AgenticAIAgentCollectionMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AgenticAIAgentCollection.Contract.AgenticAIAgentCollectionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AgenticAIAgentCollection.Contract.AgenticAIAgentCollectionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AgenticAIAgentCollection.Contract.AgenticAIAgentCollectionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AgenticAIAgentCollection.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AgenticAIAgentCollection.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AgenticAIAgentCollection.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AgenticAIAgentCollection.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _AgenticAIAgentCollection.Contract.DEFAULTADMINROLE(&_AgenticAIAgentCollection.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _AgenticAIAgentCollection.Contract.DEFAULTADMINROLE(&_AgenticAIAgentCollection.CallOpts)
}

// SIGNERROLE is a free data retrieval call binding the contract method 0xa1ebf35d.
//
// Solidity: function SIGNER_ROLE() view returns(bytes32)
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionCaller) SIGNERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AgenticAIAgentCollection.contract.Call(opts, &out, "SIGNER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SIGNERROLE is a free data retrieval call binding the contract method 0xa1ebf35d.
//
// Solidity: function SIGNER_ROLE() view returns(bytes32)
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionSession) SIGNERROLE() ([32]byte, error) {
	return _AgenticAIAgentCollection.Contract.SIGNERROLE(&_AgenticAIAgentCollection.CallOpts)
}

// SIGNERROLE is a free data retrieval call binding the contract method 0xa1ebf35d.
//
// Solidity: function SIGNER_ROLE() view returns(bytes32)
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionCallerSession) SIGNERROLE() ([32]byte, error) {
	return _AgenticAIAgentCollection.Contract.SIGNERROLE(&_AgenticAIAgentCollection.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AgenticAIAgentCollection.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _AgenticAIAgentCollection.Contract.BalanceOf(&_AgenticAIAgentCollection.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _AgenticAIAgentCollection.Contract.BalanceOf(&_AgenticAIAgentCollection.CallOpts, owner)
}

// CreatorOf is a free data retrieval call binding the contract method 0x589a1743.
//
// Solidity: function creatorOf(uint256 tokenId) view returns(address)
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionCaller) CreatorOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _AgenticAIAgentCollection.contract.Call(opts, &out, "creatorOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CreatorOf is a free data retrieval call binding the contract method 0x589a1743.
//
// Solidity: function creatorOf(uint256 tokenId) view returns(address)
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionSession) CreatorOf(tokenId *big.Int) (common.Address, error) {
	return _AgenticAIAgentCollection.Contract.CreatorOf(&_AgenticAIAgentCollection.CallOpts, tokenId)
}

// CreatorOf is a free data retrieval call binding the contract method 0x589a1743.
//
// Solidity: function creatorOf(uint256 tokenId) view returns(address)
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionCallerSession) CreatorOf(tokenId *big.Int) (common.Address, error) {
	return _AgenticAIAgentCollection.Contract.CreatorOf(&_AgenticAIAgentCollection.CallOpts, tokenId)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _AgenticAIAgentCollection.contract.Call(opts, &out, "eip712Domain")

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
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _AgenticAIAgentCollection.Contract.Eip712Domain(&_AgenticAIAgentCollection.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _AgenticAIAgentCollection.Contract.Eip712Domain(&_AgenticAIAgentCollection.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _AgenticAIAgentCollection.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _AgenticAIAgentCollection.Contract.GetApproved(&_AgenticAIAgentCollection.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _AgenticAIAgentCollection.Contract.GetApproved(&_AgenticAIAgentCollection.CallOpts, tokenId)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _AgenticAIAgentCollection.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _AgenticAIAgentCollection.Contract.GetRoleAdmin(&_AgenticAIAgentCollection.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _AgenticAIAgentCollection.Contract.GetRoleAdmin(&_AgenticAIAgentCollection.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _AgenticAIAgentCollection.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _AgenticAIAgentCollection.Contract.HasRole(&_AgenticAIAgentCollection.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _AgenticAIAgentCollection.Contract.HasRole(&_AgenticAIAgentCollection.CallOpts, role, account)
}

// ImageOf is a free data retrieval call binding the contract method 0xd0607562.
//
// Solidity: function imageOf(uint256 tokenId) view returns(string)
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionCaller) ImageOf(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _AgenticAIAgentCollection.contract.Call(opts, &out, "imageOf", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ImageOf is a free data retrieval call binding the contract method 0xd0607562.
//
// Solidity: function imageOf(uint256 tokenId) view returns(string)
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionSession) ImageOf(tokenId *big.Int) (string, error) {
	return _AgenticAIAgentCollection.Contract.ImageOf(&_AgenticAIAgentCollection.CallOpts, tokenId)
}

// ImageOf is a free data retrieval call binding the contract method 0xd0607562.
//
// Solidity: function imageOf(uint256 tokenId) view returns(string)
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionCallerSession) ImageOf(tokenId *big.Int) (string, error) {
	return _AgenticAIAgentCollection.Contract.ImageOf(&_AgenticAIAgentCollection.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _AgenticAIAgentCollection.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _AgenticAIAgentCollection.Contract.IsApprovedForAll(&_AgenticAIAgentCollection.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _AgenticAIAgentCollection.Contract.IsApprovedForAll(&_AgenticAIAgentCollection.CallOpts, owner, operator)
}

// IsNonceUsed is a free data retrieval call binding the contract method 0xb90665e5.
//
// Solidity: function isNonceUsed(bytes32 nonce) view returns(bool)
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionCaller) IsNonceUsed(opts *bind.CallOpts, nonce [32]byte) (bool, error) {
	var out []interface{}
	err := _AgenticAIAgentCollection.contract.Call(opts, &out, "isNonceUsed", nonce)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsNonceUsed is a free data retrieval call binding the contract method 0xb90665e5.
//
// Solidity: function isNonceUsed(bytes32 nonce) view returns(bool)
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionSession) IsNonceUsed(nonce [32]byte) (bool, error) {
	return _AgenticAIAgentCollection.Contract.IsNonceUsed(&_AgenticAIAgentCollection.CallOpts, nonce)
}

// IsNonceUsed is a free data retrieval call binding the contract method 0xb90665e5.
//
// Solidity: function isNonceUsed(bytes32 nonce) view returns(bool)
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionCallerSession) IsNonceUsed(nonce [32]byte) (bool, error) {
	return _AgenticAIAgentCollection.Contract.IsNonceUsed(&_AgenticAIAgentCollection.CallOpts, nonce)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AgenticAIAgentCollection.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionSession) Name() (string, error) {
	return _AgenticAIAgentCollection.Contract.Name(&_AgenticAIAgentCollection.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionCallerSession) Name() (string, error) {
	return _AgenticAIAgentCollection.Contract.Name(&_AgenticAIAgentCollection.CallOpts)
}

// NameOf is a free data retrieval call binding the contract method 0x051a2664.
//
// Solidity: function nameOf(uint256 tokenId) view returns(string)
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionCaller) NameOf(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _AgenticAIAgentCollection.contract.Call(opts, &out, "nameOf", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// NameOf is a free data retrieval call binding the contract method 0x051a2664.
//
// Solidity: function nameOf(uint256 tokenId) view returns(string)
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionSession) NameOf(tokenId *big.Int) (string, error) {
	return _AgenticAIAgentCollection.Contract.NameOf(&_AgenticAIAgentCollection.CallOpts, tokenId)
}

// NameOf is a free data retrieval call binding the contract method 0x051a2664.
//
// Solidity: function nameOf(uint256 tokenId) view returns(string)
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionCallerSession) NameOf(tokenId *big.Int) (string, error) {
	return _AgenticAIAgentCollection.Contract.NameOf(&_AgenticAIAgentCollection.CallOpts, tokenId)
}

// NextTokenId is a free data retrieval call binding the contract method 0x75794a3c.
//
// Solidity: function nextTokenId() view returns(uint256)
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionCaller) NextTokenId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AgenticAIAgentCollection.contract.Call(opts, &out, "nextTokenId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextTokenId is a free data retrieval call binding the contract method 0x75794a3c.
//
// Solidity: function nextTokenId() view returns(uint256)
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionSession) NextTokenId() (*big.Int, error) {
	return _AgenticAIAgentCollection.Contract.NextTokenId(&_AgenticAIAgentCollection.CallOpts)
}

// NextTokenId is a free data retrieval call binding the contract method 0x75794a3c.
//
// Solidity: function nextTokenId() view returns(uint256)
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionCallerSession) NextTokenId() (*big.Int, error) {
	return _AgenticAIAgentCollection.Contract.NextTokenId(&_AgenticAIAgentCollection.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _AgenticAIAgentCollection.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _AgenticAIAgentCollection.Contract.OwnerOf(&_AgenticAIAgentCollection.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _AgenticAIAgentCollection.Contract.OwnerOf(&_AgenticAIAgentCollection.CallOpts, tokenId)
}

// RoleOf is a free data retrieval call binding the contract method 0x72b1d8cf.
//
// Solidity: function roleOf(uint256 tokenId) view returns(string)
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionCaller) RoleOf(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _AgenticAIAgentCollection.contract.Call(opts, &out, "roleOf", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// RoleOf is a free data retrieval call binding the contract method 0x72b1d8cf.
//
// Solidity: function roleOf(uint256 tokenId) view returns(string)
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionSession) RoleOf(tokenId *big.Int) (string, error) {
	return _AgenticAIAgentCollection.Contract.RoleOf(&_AgenticAIAgentCollection.CallOpts, tokenId)
}

// RoleOf is a free data retrieval call binding the contract method 0x72b1d8cf.
//
// Solidity: function roleOf(uint256 tokenId) view returns(string)
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionCallerSession) RoleOf(tokenId *big.Int) (string, error) {
	return _AgenticAIAgentCollection.Contract.RoleOf(&_AgenticAIAgentCollection.CallOpts, tokenId)
}

// SignerOf is a free data retrieval call binding the contract method 0x5161fdf5.
//
// Solidity: function signerOf(uint256 tokenId) view returns(address)
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionCaller) SignerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _AgenticAIAgentCollection.contract.Call(opts, &out, "signerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SignerOf is a free data retrieval call binding the contract method 0x5161fdf5.
//
// Solidity: function signerOf(uint256 tokenId) view returns(address)
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionSession) SignerOf(tokenId *big.Int) (common.Address, error) {
	return _AgenticAIAgentCollection.Contract.SignerOf(&_AgenticAIAgentCollection.CallOpts, tokenId)
}

// SignerOf is a free data retrieval call binding the contract method 0x5161fdf5.
//
// Solidity: function signerOf(uint256 tokenId) view returns(address)
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionCallerSession) SignerOf(tokenId *big.Int) (common.Address, error) {
	return _AgenticAIAgentCollection.Contract.SignerOf(&_AgenticAIAgentCollection.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _AgenticAIAgentCollection.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _AgenticAIAgentCollection.Contract.SupportsInterface(&_AgenticAIAgentCollection.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _AgenticAIAgentCollection.Contract.SupportsInterface(&_AgenticAIAgentCollection.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AgenticAIAgentCollection.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionSession) Symbol() (string, error) {
	return _AgenticAIAgentCollection.Contract.Symbol(&_AgenticAIAgentCollection.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionCallerSession) Symbol() (string, error) {
	return _AgenticAIAgentCollection.Contract.Symbol(&_AgenticAIAgentCollection.CallOpts)
}

// TokenIdOf is a free data retrieval call binding the contract method 0xa398d819.
//
// Solidity: function tokenIdOf(string name) view returns(uint256)
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionCaller) TokenIdOf(opts *bind.CallOpts, name string) (*big.Int, error) {
	var out []interface{}
	err := _AgenticAIAgentCollection.contract.Call(opts, &out, "tokenIdOf", name)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenIdOf is a free data retrieval call binding the contract method 0xa398d819.
//
// Solidity: function tokenIdOf(string name) view returns(uint256)
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionSession) TokenIdOf(name string) (*big.Int, error) {
	return _AgenticAIAgentCollection.Contract.TokenIdOf(&_AgenticAIAgentCollection.CallOpts, name)
}

// TokenIdOf is a free data retrieval call binding the contract method 0xa398d819.
//
// Solidity: function tokenIdOf(string name) view returns(uint256)
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionCallerSession) TokenIdOf(name string) (*big.Int, error) {
	return _AgenticAIAgentCollection.Contract.TokenIdOf(&_AgenticAIAgentCollection.CallOpts, name)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _AgenticAIAgentCollection.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionSession) TokenURI(tokenId *big.Int) (string, error) {
	return _AgenticAIAgentCollection.Contract.TokenURI(&_AgenticAIAgentCollection.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _AgenticAIAgentCollection.Contract.TokenURI(&_AgenticAIAgentCollection.CallOpts, tokenId)
}

// WalletsOf is a free data retrieval call binding the contract method 0x83648671.
//
// Solidity: function walletsOf(uint256 tokenId) view returns(address[])
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionCaller) WalletsOf(opts *bind.CallOpts, tokenId *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _AgenticAIAgentCollection.contract.Call(opts, &out, "walletsOf", tokenId)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// WalletsOf is a free data retrieval call binding the contract method 0x83648671.
//
// Solidity: function walletsOf(uint256 tokenId) view returns(address[])
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionSession) WalletsOf(tokenId *big.Int) ([]common.Address, error) {
	return _AgenticAIAgentCollection.Contract.WalletsOf(&_AgenticAIAgentCollection.CallOpts, tokenId)
}

// WalletsOf is a free data retrieval call binding the contract method 0x83648671.
//
// Solidity: function walletsOf(uint256 tokenId) view returns(address[])
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionCallerSession) WalletsOf(tokenId *big.Int) ([]common.Address, error) {
	return _AgenticAIAgentCollection.Contract.WalletsOf(&_AgenticAIAgentCollection.CallOpts, tokenId)
}

// AddSigner is a paid mutator transaction binding the contract method 0xeb12d61e.
//
// Solidity: function addSigner(address signer) returns()
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionTransactor) AddSigner(opts *bind.TransactOpts, signer common.Address) (*types.Transaction, error) {
	return _AgenticAIAgentCollection.contract.Transact(opts, "addSigner", signer)
}

// AddSigner is a paid mutator transaction binding the contract method 0xeb12d61e.
//
// Solidity: function addSigner(address signer) returns()
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionSession) AddSigner(signer common.Address) (*types.Transaction, error) {
	return _AgenticAIAgentCollection.Contract.AddSigner(&_AgenticAIAgentCollection.TransactOpts, signer)
}

// AddSigner is a paid mutator transaction binding the contract method 0xeb12d61e.
//
// Solidity: function addSigner(address signer) returns()
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionTransactorSession) AddSigner(signer common.Address) (*types.Transaction, error) {
	return _AgenticAIAgentCollection.Contract.AddSigner(&_AgenticAIAgentCollection.TransactOpts, signer)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _AgenticAIAgentCollection.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _AgenticAIAgentCollection.Contract.Approve(&_AgenticAIAgentCollection.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _AgenticAIAgentCollection.Contract.Approve(&_AgenticAIAgentCollection.TransactOpts, to, tokenId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AgenticAIAgentCollection.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AgenticAIAgentCollection.Contract.GrantRole(&_AgenticAIAgentCollection.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AgenticAIAgentCollection.Contract.GrantRole(&_AgenticAIAgentCollection.TransactOpts, role, account)
}

// Mint is a paid mutator transaction binding the contract method 0xeeadae2a.
//
// Solidity: function mint(address to, address[] agentAddresses, address signer, string role, string name, string image, bytes32 nonce, bytes signature) returns()
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionTransactor) Mint(opts *bind.TransactOpts, to common.Address, agentAddresses []common.Address, signer common.Address, role string, name string, image string, nonce [32]byte, signature []byte) (*types.Transaction, error) {
	return _AgenticAIAgentCollection.contract.Transact(opts, "mint", to, agentAddresses, signer, role, name, image, nonce, signature)
}

// Mint is a paid mutator transaction binding the contract method 0xeeadae2a.
//
// Solidity: function mint(address to, address[] agentAddresses, address signer, string role, string name, string image, bytes32 nonce, bytes signature) returns()
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionSession) Mint(to common.Address, agentAddresses []common.Address, signer common.Address, role string, name string, image string, nonce [32]byte, signature []byte) (*types.Transaction, error) {
	return _AgenticAIAgentCollection.Contract.Mint(&_AgenticAIAgentCollection.TransactOpts, to, agentAddresses, signer, role, name, image, nonce, signature)
}

// Mint is a paid mutator transaction binding the contract method 0xeeadae2a.
//
// Solidity: function mint(address to, address[] agentAddresses, address signer, string role, string name, string image, bytes32 nonce, bytes signature) returns()
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionTransactorSession) Mint(to common.Address, agentAddresses []common.Address, signer common.Address, role string, name string, image string, nonce [32]byte, signature []byte) (*types.Transaction, error) {
	return _AgenticAIAgentCollection.Contract.Mint(&_AgenticAIAgentCollection.TransactOpts, to, agentAddresses, signer, role, name, image, nonce, signature)
}

// RemoveSigner is a paid mutator transaction binding the contract method 0x0e316ab7.
//
// Solidity: function removeSigner(address signer) returns()
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionTransactor) RemoveSigner(opts *bind.TransactOpts, signer common.Address) (*types.Transaction, error) {
	return _AgenticAIAgentCollection.contract.Transact(opts, "removeSigner", signer)
}

// RemoveSigner is a paid mutator transaction binding the contract method 0x0e316ab7.
//
// Solidity: function removeSigner(address signer) returns()
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionSession) RemoveSigner(signer common.Address) (*types.Transaction, error) {
	return _AgenticAIAgentCollection.Contract.RemoveSigner(&_AgenticAIAgentCollection.TransactOpts, signer)
}

// RemoveSigner is a paid mutator transaction binding the contract method 0x0e316ab7.
//
// Solidity: function removeSigner(address signer) returns()
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionTransactorSession) RemoveSigner(signer common.Address) (*types.Transaction, error) {
	return _AgenticAIAgentCollection.Contract.RemoveSigner(&_AgenticAIAgentCollection.TransactOpts, signer)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _AgenticAIAgentCollection.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _AgenticAIAgentCollection.Contract.RenounceRole(&_AgenticAIAgentCollection.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _AgenticAIAgentCollection.Contract.RenounceRole(&_AgenticAIAgentCollection.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AgenticAIAgentCollection.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AgenticAIAgentCollection.Contract.RevokeRole(&_AgenticAIAgentCollection.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AgenticAIAgentCollection.Contract.RevokeRole(&_AgenticAIAgentCollection.TransactOpts, role, account)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _AgenticAIAgentCollection.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _AgenticAIAgentCollection.Contract.SafeTransferFrom(&_AgenticAIAgentCollection.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _AgenticAIAgentCollection.Contract.SafeTransferFrom(&_AgenticAIAgentCollection.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _AgenticAIAgentCollection.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _AgenticAIAgentCollection.Contract.SafeTransferFrom0(&_AgenticAIAgentCollection.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _AgenticAIAgentCollection.Contract.SafeTransferFrom0(&_AgenticAIAgentCollection.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _AgenticAIAgentCollection.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _AgenticAIAgentCollection.Contract.SetApprovalForAll(&_AgenticAIAgentCollection.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _AgenticAIAgentCollection.Contract.SetApprovalForAll(&_AgenticAIAgentCollection.TransactOpts, operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _AgenticAIAgentCollection.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _AgenticAIAgentCollection.Contract.TransferFrom(&_AgenticAIAgentCollection.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _AgenticAIAgentCollection.Contract.TransferFrom(&_AgenticAIAgentCollection.TransactOpts, from, to, tokenId)
}

// UpdateAgentDescriptor is a paid mutator transaction binding the contract method 0x6457f503.
//
// Solidity: function updateAgentDescriptor(address agentDescriptor) returns()
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionTransactor) UpdateAgentDescriptor(opts *bind.TransactOpts, agentDescriptor common.Address) (*types.Transaction, error) {
	return _AgenticAIAgentCollection.contract.Transact(opts, "updateAgentDescriptor", agentDescriptor)
}

// UpdateAgentDescriptor is a paid mutator transaction binding the contract method 0x6457f503.
//
// Solidity: function updateAgentDescriptor(address agentDescriptor) returns()
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionSession) UpdateAgentDescriptor(agentDescriptor common.Address) (*types.Transaction, error) {
	return _AgenticAIAgentCollection.Contract.UpdateAgentDescriptor(&_AgenticAIAgentCollection.TransactOpts, agentDescriptor)
}

// UpdateAgentDescriptor is a paid mutator transaction binding the contract method 0x6457f503.
//
// Solidity: function updateAgentDescriptor(address agentDescriptor) returns()
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionTransactorSession) UpdateAgentDescriptor(agentDescriptor common.Address) (*types.Transaction, error) {
	return _AgenticAIAgentCollection.Contract.UpdateAgentDescriptor(&_AgenticAIAgentCollection.TransactOpts, agentDescriptor)
}

// UpdateTokenSigner is a paid mutator transaction binding the contract method 0x8c2c7ec5.
//
// Solidity: function updateTokenSigner(uint256 tokenId, address newSigner, bytes32 nonce, bytes signature) returns()
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionTransactor) UpdateTokenSigner(opts *bind.TransactOpts, tokenId *big.Int, newSigner common.Address, nonce [32]byte, signature []byte) (*types.Transaction, error) {
	return _AgenticAIAgentCollection.contract.Transact(opts, "updateTokenSigner", tokenId, newSigner, nonce, signature)
}

// UpdateTokenSigner is a paid mutator transaction binding the contract method 0x8c2c7ec5.
//
// Solidity: function updateTokenSigner(uint256 tokenId, address newSigner, bytes32 nonce, bytes signature) returns()
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionSession) UpdateTokenSigner(tokenId *big.Int, newSigner common.Address, nonce [32]byte, signature []byte) (*types.Transaction, error) {
	return _AgenticAIAgentCollection.Contract.UpdateTokenSigner(&_AgenticAIAgentCollection.TransactOpts, tokenId, newSigner, nonce, signature)
}

// UpdateTokenSigner is a paid mutator transaction binding the contract method 0x8c2c7ec5.
//
// Solidity: function updateTokenSigner(uint256 tokenId, address newSigner, bytes32 nonce, bytes signature) returns()
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionTransactorSession) UpdateTokenSigner(tokenId *big.Int, newSigner common.Address, nonce [32]byte, signature []byte) (*types.Transaction, error) {
	return _AgenticAIAgentCollection.Contract.UpdateTokenSigner(&_AgenticAIAgentCollection.TransactOpts, tokenId, newSigner, nonce, signature)
}

// AgenticAIAgentCollectionApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the AgenticAIAgentCollection contract.
type AgenticAIAgentCollectionApprovalIterator struct {
	Event *AgenticAIAgentCollectionApproval // Event containing the contract specifics and raw log

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
func (it *AgenticAIAgentCollectionApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgenticAIAgentCollectionApproval)
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
		it.Event = new(AgenticAIAgentCollectionApproval)
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
func (it *AgenticAIAgentCollectionApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgenticAIAgentCollectionApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgenticAIAgentCollectionApproval represents a Approval event raised by the AgenticAIAgentCollection contract.
type AgenticAIAgentCollectionApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*AgenticAIAgentCollectionApprovalIterator, error) {

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

	logs, sub, err := _AgenticAIAgentCollection.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &AgenticAIAgentCollectionApprovalIterator{contract: _AgenticAIAgentCollection.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *AgenticAIAgentCollectionApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _AgenticAIAgentCollection.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgenticAIAgentCollectionApproval)
				if err := _AgenticAIAgentCollection.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionFilterer) ParseApproval(log types.Log) (*AgenticAIAgentCollectionApproval, error) {
	event := new(AgenticAIAgentCollectionApproval)
	if err := _AgenticAIAgentCollection.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgenticAIAgentCollectionApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the AgenticAIAgentCollection contract.
type AgenticAIAgentCollectionApprovalForAllIterator struct {
	Event *AgenticAIAgentCollectionApprovalForAll // Event containing the contract specifics and raw log

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
func (it *AgenticAIAgentCollectionApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgenticAIAgentCollectionApprovalForAll)
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
		it.Event = new(AgenticAIAgentCollectionApprovalForAll)
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
func (it *AgenticAIAgentCollectionApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgenticAIAgentCollectionApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgenticAIAgentCollectionApprovalForAll represents a ApprovalForAll event raised by the AgenticAIAgentCollection contract.
type AgenticAIAgentCollectionApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*AgenticAIAgentCollectionApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _AgenticAIAgentCollection.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &AgenticAIAgentCollectionApprovalForAllIterator{contract: _AgenticAIAgentCollection.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *AgenticAIAgentCollectionApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _AgenticAIAgentCollection.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgenticAIAgentCollectionApprovalForAll)
				if err := _AgenticAIAgentCollection.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionFilterer) ParseApprovalForAll(log types.Log) (*AgenticAIAgentCollectionApprovalForAll, error) {
	event := new(AgenticAIAgentCollectionApprovalForAll)
	if err := _AgenticAIAgentCollection.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgenticAIAgentCollectionEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the AgenticAIAgentCollection contract.
type AgenticAIAgentCollectionEIP712DomainChangedIterator struct {
	Event *AgenticAIAgentCollectionEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *AgenticAIAgentCollectionEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgenticAIAgentCollectionEIP712DomainChanged)
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
		it.Event = new(AgenticAIAgentCollectionEIP712DomainChanged)
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
func (it *AgenticAIAgentCollectionEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgenticAIAgentCollectionEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgenticAIAgentCollectionEIP712DomainChanged represents a EIP712DomainChanged event raised by the AgenticAIAgentCollection contract.
type AgenticAIAgentCollectionEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*AgenticAIAgentCollectionEIP712DomainChangedIterator, error) {

	logs, sub, err := _AgenticAIAgentCollection.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &AgenticAIAgentCollectionEIP712DomainChangedIterator{contract: _AgenticAIAgentCollection.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *AgenticAIAgentCollectionEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _AgenticAIAgentCollection.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgenticAIAgentCollectionEIP712DomainChanged)
				if err := _AgenticAIAgentCollection.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionFilterer) ParseEIP712DomainChanged(log types.Log) (*AgenticAIAgentCollectionEIP712DomainChanged, error) {
	event := new(AgenticAIAgentCollectionEIP712DomainChanged)
	if err := _AgenticAIAgentCollection.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgenticAIAgentCollectionMintedIterator is returned from FilterMinted and is used to iterate over the raw logs and unpacked data for Minted events raised by the AgenticAIAgentCollection contract.
type AgenticAIAgentCollectionMintedIterator struct {
	Event *AgenticAIAgentCollectionMinted // Event containing the contract specifics and raw log

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
func (it *AgenticAIAgentCollectionMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgenticAIAgentCollectionMinted)
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
		it.Event = new(AgenticAIAgentCollectionMinted)
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
func (it *AgenticAIAgentCollectionMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgenticAIAgentCollectionMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgenticAIAgentCollectionMinted represents a Minted event raised by the AgenticAIAgentCollection contract.
type AgenticAIAgentCollectionMinted struct {
	To      common.Address
	Signer  common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMinted is a free log retrieval operation binding the contract event 0x9d228d69b5fdb8d273a2336f8fb8612d039631024ea9bf09c424a9503aa078f0.
//
// Solidity: event Minted(address indexed to, address indexed signer, uint256 tokenId)
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionFilterer) FilterMinted(opts *bind.FilterOpts, to []common.Address, signer []common.Address) (*AgenticAIAgentCollectionMintedIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _AgenticAIAgentCollection.contract.FilterLogs(opts, "Minted", toRule, signerRule)
	if err != nil {
		return nil, err
	}
	return &AgenticAIAgentCollectionMintedIterator{contract: _AgenticAIAgentCollection.contract, event: "Minted", logs: logs, sub: sub}, nil
}

// WatchMinted is a free log subscription operation binding the contract event 0x9d228d69b5fdb8d273a2336f8fb8612d039631024ea9bf09c424a9503aa078f0.
//
// Solidity: event Minted(address indexed to, address indexed signer, uint256 tokenId)
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionFilterer) WatchMinted(opts *bind.WatchOpts, sink chan<- *AgenticAIAgentCollectionMinted, to []common.Address, signer []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _AgenticAIAgentCollection.contract.WatchLogs(opts, "Minted", toRule, signerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgenticAIAgentCollectionMinted)
				if err := _AgenticAIAgentCollection.contract.UnpackLog(event, "Minted", log); err != nil {
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
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionFilterer) ParseMinted(log types.Log) (*AgenticAIAgentCollectionMinted, error) {
	event := new(AgenticAIAgentCollectionMinted)
	if err := _AgenticAIAgentCollection.contract.UnpackLog(event, "Minted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgenticAIAgentCollectionRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the AgenticAIAgentCollection contract.
type AgenticAIAgentCollectionRoleAdminChangedIterator struct {
	Event *AgenticAIAgentCollectionRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *AgenticAIAgentCollectionRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgenticAIAgentCollectionRoleAdminChanged)
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
		it.Event = new(AgenticAIAgentCollectionRoleAdminChanged)
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
func (it *AgenticAIAgentCollectionRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgenticAIAgentCollectionRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgenticAIAgentCollectionRoleAdminChanged represents a RoleAdminChanged event raised by the AgenticAIAgentCollection contract.
type AgenticAIAgentCollectionRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*AgenticAIAgentCollectionRoleAdminChangedIterator, error) {

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

	logs, sub, err := _AgenticAIAgentCollection.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &AgenticAIAgentCollectionRoleAdminChangedIterator{contract: _AgenticAIAgentCollection.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *AgenticAIAgentCollectionRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _AgenticAIAgentCollection.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgenticAIAgentCollectionRoleAdminChanged)
				if err := _AgenticAIAgentCollection.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionFilterer) ParseRoleAdminChanged(log types.Log) (*AgenticAIAgentCollectionRoleAdminChanged, error) {
	event := new(AgenticAIAgentCollectionRoleAdminChanged)
	if err := _AgenticAIAgentCollection.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgenticAIAgentCollectionRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the AgenticAIAgentCollection contract.
type AgenticAIAgentCollectionRoleGrantedIterator struct {
	Event *AgenticAIAgentCollectionRoleGranted // Event containing the contract specifics and raw log

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
func (it *AgenticAIAgentCollectionRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgenticAIAgentCollectionRoleGranted)
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
		it.Event = new(AgenticAIAgentCollectionRoleGranted)
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
func (it *AgenticAIAgentCollectionRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgenticAIAgentCollectionRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgenticAIAgentCollectionRoleGranted represents a RoleGranted event raised by the AgenticAIAgentCollection contract.
type AgenticAIAgentCollectionRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AgenticAIAgentCollectionRoleGrantedIterator, error) {

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

	logs, sub, err := _AgenticAIAgentCollection.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AgenticAIAgentCollectionRoleGrantedIterator{contract: _AgenticAIAgentCollection.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *AgenticAIAgentCollectionRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _AgenticAIAgentCollection.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgenticAIAgentCollectionRoleGranted)
				if err := _AgenticAIAgentCollection.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionFilterer) ParseRoleGranted(log types.Log) (*AgenticAIAgentCollectionRoleGranted, error) {
	event := new(AgenticAIAgentCollectionRoleGranted)
	if err := _AgenticAIAgentCollection.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgenticAIAgentCollectionRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the AgenticAIAgentCollection contract.
type AgenticAIAgentCollectionRoleRevokedIterator struct {
	Event *AgenticAIAgentCollectionRoleRevoked // Event containing the contract specifics and raw log

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
func (it *AgenticAIAgentCollectionRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgenticAIAgentCollectionRoleRevoked)
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
		it.Event = new(AgenticAIAgentCollectionRoleRevoked)
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
func (it *AgenticAIAgentCollectionRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgenticAIAgentCollectionRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgenticAIAgentCollectionRoleRevoked represents a RoleRevoked event raised by the AgenticAIAgentCollection contract.
type AgenticAIAgentCollectionRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AgenticAIAgentCollectionRoleRevokedIterator, error) {

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

	logs, sub, err := _AgenticAIAgentCollection.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AgenticAIAgentCollectionRoleRevokedIterator{contract: _AgenticAIAgentCollection.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *AgenticAIAgentCollectionRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _AgenticAIAgentCollection.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgenticAIAgentCollectionRoleRevoked)
				if err := _AgenticAIAgentCollection.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionFilterer) ParseRoleRevoked(log types.Log) (*AgenticAIAgentCollectionRoleRevoked, error) {
	event := new(AgenticAIAgentCollectionRoleRevoked)
	if err := _AgenticAIAgentCollection.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgenticAIAgentCollectionSignerAddedIterator is returned from FilterSignerAdded and is used to iterate over the raw logs and unpacked data for SignerAdded events raised by the AgenticAIAgentCollection contract.
type AgenticAIAgentCollectionSignerAddedIterator struct {
	Event *AgenticAIAgentCollectionSignerAdded // Event containing the contract specifics and raw log

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
func (it *AgenticAIAgentCollectionSignerAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgenticAIAgentCollectionSignerAdded)
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
		it.Event = new(AgenticAIAgentCollectionSignerAdded)
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
func (it *AgenticAIAgentCollectionSignerAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgenticAIAgentCollectionSignerAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgenticAIAgentCollectionSignerAdded represents a SignerAdded event raised by the AgenticAIAgentCollection contract.
type AgenticAIAgentCollectionSignerAdded struct {
	Signer common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSignerAdded is a free log retrieval operation binding the contract event 0x47d1c22a25bb3a5d4e481b9b1e6944c2eade3181a0a20b495ed61d35b5323f24.
//
// Solidity: event SignerAdded(address signer)
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionFilterer) FilterSignerAdded(opts *bind.FilterOpts) (*AgenticAIAgentCollectionSignerAddedIterator, error) {

	logs, sub, err := _AgenticAIAgentCollection.contract.FilterLogs(opts, "SignerAdded")
	if err != nil {
		return nil, err
	}
	return &AgenticAIAgentCollectionSignerAddedIterator{contract: _AgenticAIAgentCollection.contract, event: "SignerAdded", logs: logs, sub: sub}, nil
}

// WatchSignerAdded is a free log subscription operation binding the contract event 0x47d1c22a25bb3a5d4e481b9b1e6944c2eade3181a0a20b495ed61d35b5323f24.
//
// Solidity: event SignerAdded(address signer)
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionFilterer) WatchSignerAdded(opts *bind.WatchOpts, sink chan<- *AgenticAIAgentCollectionSignerAdded) (event.Subscription, error) {

	logs, sub, err := _AgenticAIAgentCollection.contract.WatchLogs(opts, "SignerAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgenticAIAgentCollectionSignerAdded)
				if err := _AgenticAIAgentCollection.contract.UnpackLog(event, "SignerAdded", log); err != nil {
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

// ParseSignerAdded is a log parse operation binding the contract event 0x47d1c22a25bb3a5d4e481b9b1e6944c2eade3181a0a20b495ed61d35b5323f24.
//
// Solidity: event SignerAdded(address signer)
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionFilterer) ParseSignerAdded(log types.Log) (*AgenticAIAgentCollectionSignerAdded, error) {
	event := new(AgenticAIAgentCollectionSignerAdded)
	if err := _AgenticAIAgentCollection.contract.UnpackLog(event, "SignerAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgenticAIAgentCollectionSignerRemovedIterator is returned from FilterSignerRemoved and is used to iterate over the raw logs and unpacked data for SignerRemoved events raised by the AgenticAIAgentCollection contract.
type AgenticAIAgentCollectionSignerRemovedIterator struct {
	Event *AgenticAIAgentCollectionSignerRemoved // Event containing the contract specifics and raw log

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
func (it *AgenticAIAgentCollectionSignerRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgenticAIAgentCollectionSignerRemoved)
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
		it.Event = new(AgenticAIAgentCollectionSignerRemoved)
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
func (it *AgenticAIAgentCollectionSignerRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgenticAIAgentCollectionSignerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgenticAIAgentCollectionSignerRemoved represents a SignerRemoved event raised by the AgenticAIAgentCollection contract.
type AgenticAIAgentCollectionSignerRemoved struct {
	Signer common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSignerRemoved is a free log retrieval operation binding the contract event 0x3525e22824a8a7df2c9a6029941c824cf95b6447f1e13d5128fd3826d35afe8b.
//
// Solidity: event SignerRemoved(address signer)
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionFilterer) FilterSignerRemoved(opts *bind.FilterOpts) (*AgenticAIAgentCollectionSignerRemovedIterator, error) {

	logs, sub, err := _AgenticAIAgentCollection.contract.FilterLogs(opts, "SignerRemoved")
	if err != nil {
		return nil, err
	}
	return &AgenticAIAgentCollectionSignerRemovedIterator{contract: _AgenticAIAgentCollection.contract, event: "SignerRemoved", logs: logs, sub: sub}, nil
}

// WatchSignerRemoved is a free log subscription operation binding the contract event 0x3525e22824a8a7df2c9a6029941c824cf95b6447f1e13d5128fd3826d35afe8b.
//
// Solidity: event SignerRemoved(address signer)
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionFilterer) WatchSignerRemoved(opts *bind.WatchOpts, sink chan<- *AgenticAIAgentCollectionSignerRemoved) (event.Subscription, error) {

	logs, sub, err := _AgenticAIAgentCollection.contract.WatchLogs(opts, "SignerRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgenticAIAgentCollectionSignerRemoved)
				if err := _AgenticAIAgentCollection.contract.UnpackLog(event, "SignerRemoved", log); err != nil {
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

// ParseSignerRemoved is a log parse operation binding the contract event 0x3525e22824a8a7df2c9a6029941c824cf95b6447f1e13d5128fd3826d35afe8b.
//
// Solidity: event SignerRemoved(address signer)
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionFilterer) ParseSignerRemoved(log types.Log) (*AgenticAIAgentCollectionSignerRemoved, error) {
	event := new(AgenticAIAgentCollectionSignerRemoved)
	if err := _AgenticAIAgentCollection.contract.UnpackLog(event, "SignerRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgenticAIAgentCollectionTokenSignerChangedIterator is returned from FilterTokenSignerChanged and is used to iterate over the raw logs and unpacked data for TokenSignerChanged events raised by the AgenticAIAgentCollection contract.
type AgenticAIAgentCollectionTokenSignerChangedIterator struct {
	Event *AgenticAIAgentCollectionTokenSignerChanged // Event containing the contract specifics and raw log

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
func (it *AgenticAIAgentCollectionTokenSignerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgenticAIAgentCollectionTokenSignerChanged)
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
		it.Event = new(AgenticAIAgentCollectionTokenSignerChanged)
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
func (it *AgenticAIAgentCollectionTokenSignerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgenticAIAgentCollectionTokenSignerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgenticAIAgentCollectionTokenSignerChanged represents a TokenSignerChanged event raised by the AgenticAIAgentCollection contract.
type AgenticAIAgentCollectionTokenSignerChanged struct {
	TokenId   *big.Int
	OldSigner common.Address
	NewSigner common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTokenSignerChanged is a free log retrieval operation binding the contract event 0xe51e6ee1b7a3c7c5d0e801f5306ab38325b6cf7d09e0dda2cf4e9bb59c00fcba.
//
// Solidity: event TokenSignerChanged(uint256 indexed tokenId, address oldSigner, address newSigner)
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionFilterer) FilterTokenSignerChanged(opts *bind.FilterOpts, tokenId []*big.Int) (*AgenticAIAgentCollectionTokenSignerChangedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _AgenticAIAgentCollection.contract.FilterLogs(opts, "TokenSignerChanged", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &AgenticAIAgentCollectionTokenSignerChangedIterator{contract: _AgenticAIAgentCollection.contract, event: "TokenSignerChanged", logs: logs, sub: sub}, nil
}

// WatchTokenSignerChanged is a free log subscription operation binding the contract event 0xe51e6ee1b7a3c7c5d0e801f5306ab38325b6cf7d09e0dda2cf4e9bb59c00fcba.
//
// Solidity: event TokenSignerChanged(uint256 indexed tokenId, address oldSigner, address newSigner)
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionFilterer) WatchTokenSignerChanged(opts *bind.WatchOpts, sink chan<- *AgenticAIAgentCollectionTokenSignerChanged, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _AgenticAIAgentCollection.contract.WatchLogs(opts, "TokenSignerChanged", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgenticAIAgentCollectionTokenSignerChanged)
				if err := _AgenticAIAgentCollection.contract.UnpackLog(event, "TokenSignerChanged", log); err != nil {
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

// ParseTokenSignerChanged is a log parse operation binding the contract event 0xe51e6ee1b7a3c7c5d0e801f5306ab38325b6cf7d09e0dda2cf4e9bb59c00fcba.
//
// Solidity: event TokenSignerChanged(uint256 indexed tokenId, address oldSigner, address newSigner)
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionFilterer) ParseTokenSignerChanged(log types.Log) (*AgenticAIAgentCollectionTokenSignerChanged, error) {
	event := new(AgenticAIAgentCollectionTokenSignerChanged)
	if err := _AgenticAIAgentCollection.contract.UnpackLog(event, "TokenSignerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgenticAIAgentCollectionTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the AgenticAIAgentCollection contract.
type AgenticAIAgentCollectionTransferIterator struct {
	Event *AgenticAIAgentCollectionTransfer // Event containing the contract specifics and raw log

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
func (it *AgenticAIAgentCollectionTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgenticAIAgentCollectionTransfer)
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
		it.Event = new(AgenticAIAgentCollectionTransfer)
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
func (it *AgenticAIAgentCollectionTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgenticAIAgentCollectionTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgenticAIAgentCollectionTransfer represents a Transfer event raised by the AgenticAIAgentCollection contract.
type AgenticAIAgentCollectionTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*AgenticAIAgentCollectionTransferIterator, error) {

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

	logs, sub, err := _AgenticAIAgentCollection.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &AgenticAIAgentCollectionTransferIterator{contract: _AgenticAIAgentCollection.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *AgenticAIAgentCollectionTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _AgenticAIAgentCollection.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgenticAIAgentCollectionTransfer)
				if err := _AgenticAIAgentCollection.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_AgenticAIAgentCollection *AgenticAIAgentCollectionFilterer) ParseTransfer(log types.Log) (*AgenticAIAgentCollectionTransfer, error) {
	event := new(AgenticAIAgentCollectionTransfer)
	if err := _AgenticAIAgentCollection.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
