// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"math/big"
	"strings"
	"errors"

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





	// AintGatewayMetaData contains all meta data concerning the AintGateway contract.
	// var AintGatewayMetaData = &bind.MetaData{
	// 	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_withdrawCount\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_withdraws\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"addWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"cancelWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"to\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getWithdrawFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"setWithdrawFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"withdrawExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
		
	// }
	// AintGatewayABI is the input ABI used to generate the binding from.
	// Deprecated: Use AintGatewayMetaData.ABI instead.
	var AintGatewayABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_withdrawCount\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_withdraws\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"addWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"cancelWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"to\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getWithdrawFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"setWithdrawFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"withdrawExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

	

	

	// AintGateway is an auto generated Go binding around an Ethereum contract.
	type AintGateway struct {
	  AintGatewayCaller     // Read-only binding to the contract
	  AintGatewayTransactor // Write-only binding to the contract
	  AintGatewayFilterer   // Log filterer for contract events
	}

	// AintGatewayCaller is an auto generated read-only Go binding around an Ethereum contract.
	type AintGatewayCaller struct {
	  contract *bind.BoundContract // Generic contract wrapper for the low level calls
	}

	// AintGatewayTransactor is an auto generated write-only Go binding around an Ethereum contract.
	type AintGatewayTransactor struct {
	  contract *bind.BoundContract // Generic contract wrapper for the low level calls
	}

	// AintGatewayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
	type AintGatewayFilterer struct {
	  contract *bind.BoundContract // Generic contract wrapper for the low level calls
	}

	// AintGatewaySession is an auto generated Go binding around an Ethereum contract,
	// with pre-set call and transact options.
	type AintGatewaySession struct {
	  Contract     *AintGateway        // Generic contract binding to set the session for
	  CallOpts     bind.CallOpts     // Call options to use throughout this session
	  TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
	}

	// AintGatewayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
	// with pre-set call options.
	type AintGatewayCallerSession struct {
	  Contract *AintGatewayCaller // Generic contract caller binding to set the session for
	  CallOpts bind.CallOpts    // Call options to use throughout this session
	}

	// AintGatewayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
	// with pre-set transact options.
	type AintGatewayTransactorSession struct {
	  Contract     *AintGatewayTransactor // Generic contract transactor binding to set the session for
	  TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
	}

	// AintGatewayRaw is an auto generated low-level Go binding around an Ethereum contract.
	type AintGatewayRaw struct {
	  Contract *AintGateway // Generic contract binding to access the raw methods on
	}

	// AintGatewayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
	type AintGatewayCallerRaw struct {
		Contract *AintGatewayCaller // Generic read-only contract binding to access the raw methods on
	}

	// AintGatewayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
	type AintGatewayTransactorRaw struct {
		Contract *AintGatewayTransactor // Generic write-only contract binding to access the raw methods on
	}

	// NewAintGateway creates a new instance of AintGateway, bound to a specific deployed contract.
	func NewAintGateway(address common.Address, backend bind.ContractBackend) (*AintGateway, error) {
	  contract, err := bindAintGateway(address, backend, backend, backend)
	  if err != nil {
	    return nil, err
	  }
	  return &AintGateway{ AintGatewayCaller: AintGatewayCaller{contract: contract}, AintGatewayTransactor: AintGatewayTransactor{contract: contract}, AintGatewayFilterer: AintGatewayFilterer{contract: contract} }, nil
	}

	// NewAintGatewayCaller creates a new read-only instance of AintGateway, bound to a specific deployed contract.
	func NewAintGatewayCaller(address common.Address, caller bind.ContractCaller) (*AintGatewayCaller, error) {
	  contract, err := bindAintGateway(address, caller, nil, nil)
	  if err != nil {
	    return nil, err
	  }
	  return &AintGatewayCaller{contract: contract}, nil
	}

	// NewAintGatewayTransactor creates a new write-only instance of AintGateway, bound to a specific deployed contract.
	func NewAintGatewayTransactor(address common.Address, transactor bind.ContractTransactor) (*AintGatewayTransactor, error) {
	  contract, err := bindAintGateway(address, nil, transactor, nil)
	  if err != nil {
	    return nil, err
	  }
	  return &AintGatewayTransactor{contract: contract}, nil
	}

	// NewAintGatewayFilterer creates a new log filterer instance of AintGateway, bound to a specific deployed contract.
 	func NewAintGatewayFilterer(address common.Address, filterer bind.ContractFilterer) (*AintGatewayFilterer, error) {
 	  contract, err := bindAintGateway(address, nil, nil, filterer)
 	  if err != nil {
 	    return nil, err
 	  }
 	  return &AintGatewayFilterer{contract: contract}, nil
 	}

	// bindAintGateway binds a generic wrapper to an already deployed contract.
	func bindAintGateway(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	  parsed, err := AintGatewayMetaData.GetAbi()
	  if err != nil {
	    return nil, err
	  }
	  return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
	}

	// Call invokes the (constant) contract method with params as input values and
	// sets the output to result. The result type might be a single field for simple
	// returns, a slice of interfaces for anonymous returns and a struct for named
	// returns.
	func (_AintGateway *AintGatewayRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
		return _AintGateway.Contract.AintGatewayCaller.contract.Call(opts, result, method, params...)
	}

	// Transfer initiates a plain transaction to move funds to the contract, calling
	// its default method if one is available.
	func (_AintGateway *AintGatewayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
		return _AintGateway.Contract.AintGatewayTransactor.contract.Transfer(opts)
	}

	// Transact invokes the (paid) contract method with params as input values.
	func (_AintGateway *AintGatewayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
		return _AintGateway.Contract.AintGatewayTransactor.contract.Transact(opts, method, params...)
	}

	// Call invokes the (constant) contract method with params as input values and
	// sets the output to result. The result type might be a single field for simple
	// returns, a slice of interfaces for anonymous returns and a struct for named
	// returns.
	func (_AintGateway *AintGatewayCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
		return _AintGateway.Contract.contract.Call(opts, result, method, params...)
	}

	// Transfer initiates a plain transaction to move funds to the contract, calling
	// its default method if one is available.
	func (_AintGateway *AintGatewayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
		return _AintGateway.Contract.contract.Transfer(opts)
	}

	// Transact invokes the (paid) contract method with params as input values.
	func (_AintGateway *AintGatewayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
		return _AintGateway.Contract.contract.Transact(opts, method, params...)
	}

	
		// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
		//
		// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
		func (_AintGateway *AintGatewayCaller) DEFAULTADMINROLE(opts *bind.CallOpts ) ([32]byte, error) {
			var out []interface{}
			err := _AintGateway.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE" )
			
			if err != nil {
				return *new([32]byte),  err
			}
			
			out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

			return out0,  err
			
		}

		// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
		//
		// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
		func (_AintGateway *AintGatewaySession) DEFAULTADMINROLE() ( [32]byte,  error) {
		  return _AintGateway.Contract.DEFAULTADMINROLE(&_AintGateway.CallOpts )
		}

		// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
		//
		// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
		func (_AintGateway *AintGatewayCallerSession) DEFAULTADMINROLE() ( [32]byte,  error) {
		  return _AintGateway.Contract.DEFAULTADMINROLE(&_AintGateway.CallOpts )
		}
	
		// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
		//
		// Solidity: function MINTER_ROLE() view returns(bytes32)
		func (_AintGateway *AintGatewayCaller) MINTERROLE(opts *bind.CallOpts ) ([32]byte, error) {
			var out []interface{}
			err := _AintGateway.contract.Call(opts, &out, "MINTER_ROLE" )
			
			if err != nil {
				return *new([32]byte),  err
			}
			
			out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

			return out0,  err
			
		}

		// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
		//
		// Solidity: function MINTER_ROLE() view returns(bytes32)
		func (_AintGateway *AintGatewaySession) MINTERROLE() ( [32]byte,  error) {
		  return _AintGateway.Contract.MINTERROLE(&_AintGateway.CallOpts )
		}

		// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
		//
		// Solidity: function MINTER_ROLE() view returns(bytes32)
		func (_AintGateway *AintGatewayCallerSession) MINTERROLE() ( [32]byte,  error) {
		  return _AintGateway.Contract.MINTERROLE(&_AintGateway.CallOpts )
		}
	
		// WithdrawCount is a free data retrieval call binding the contract method 0xfad72d12.
		//
		// Solidity: function _withdrawCount(address ) view returns(uint8)
		func (_AintGateway *AintGatewayCaller) WithdrawCount(opts *bind.CallOpts , arg0 common.Address ) (uint8, error) {
			var out []interface{}
			err := _AintGateway.contract.Call(opts, &out, "_withdrawCount" , arg0)
			
			if err != nil {
				return *new(uint8),  err
			}
			
			out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

			return out0,  err
			
		}

		// WithdrawCount is a free data retrieval call binding the contract method 0xfad72d12.
		//
		// Solidity: function _withdrawCount(address ) view returns(uint8)
		func (_AintGateway *AintGatewaySession) WithdrawCount( arg0 common.Address ) ( uint8,  error) {
		  return _AintGateway.Contract.WithdrawCount(&_AintGateway.CallOpts , arg0)
		}

		// WithdrawCount is a free data retrieval call binding the contract method 0xfad72d12.
		//
		// Solidity: function _withdrawCount(address ) view returns(uint8)
		func (_AintGateway *AintGatewayCallerSession) WithdrawCount( arg0 common.Address ) ( uint8,  error) {
		  return _AintGateway.Contract.WithdrawCount(&_AintGateway.CallOpts , arg0)
		}
	
		// Withdraws is a free data retrieval call binding the contract method 0x2590983b.
		//
		// Solidity: function _withdraws(address ) view returns(uint256)
		func (_AintGateway *AintGatewayCaller) Withdraws(opts *bind.CallOpts , arg0 common.Address ) (*big.Int, error) {
			var out []interface{}
			err := _AintGateway.contract.Call(opts, &out, "_withdraws" , arg0)
			
			if err != nil {
				return *new(*big.Int),  err
			}
			
			out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

			return out0,  err
			
		}

		// Withdraws is a free data retrieval call binding the contract method 0x2590983b.
		//
		// Solidity: function _withdraws(address ) view returns(uint256)
		func (_AintGateway *AintGatewaySession) Withdraws( arg0 common.Address ) ( *big.Int,  error) {
		  return _AintGateway.Contract.Withdraws(&_AintGateway.CallOpts , arg0)
		}

		// Withdraws is a free data retrieval call binding the contract method 0x2590983b.
		//
		// Solidity: function _withdraws(address ) view returns(uint256)
		func (_AintGateway *AintGatewayCallerSession) Withdraws( arg0 common.Address ) ( *big.Int,  error) {
		  return _AintGateway.Contract.Withdraws(&_AintGateway.CallOpts , arg0)
		}
	
		// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
		//
		// Solidity: function allowance(address owner, address spender) view returns(uint256)
		func (_AintGateway *AintGatewayCaller) Allowance(opts *bind.CallOpts , owner common.Address , spender common.Address ) (*big.Int, error) {
			var out []interface{}
			err := _AintGateway.contract.Call(opts, &out, "allowance" , owner, spender)
			
			if err != nil {
				return *new(*big.Int),  err
			}
			
			out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

			return out0,  err
			
		}

		// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
		//
		// Solidity: function allowance(address owner, address spender) view returns(uint256)
		func (_AintGateway *AintGatewaySession) Allowance( owner common.Address , spender common.Address ) ( *big.Int,  error) {
		  return _AintGateway.Contract.Allowance(&_AintGateway.CallOpts , owner, spender)
		}

		// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
		//
		// Solidity: function allowance(address owner, address spender) view returns(uint256)
		func (_AintGateway *AintGatewayCallerSession) Allowance( owner common.Address , spender common.Address ) ( *big.Int,  error) {
		  return _AintGateway.Contract.Allowance(&_AintGateway.CallOpts , owner, spender)
		}
	
		// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
		//
		// Solidity: function balanceOf(address account) view returns(uint256)
		func (_AintGateway *AintGatewayCaller) BalanceOf(opts *bind.CallOpts , account common.Address ) (*big.Int, error) {
			var out []interface{}
			err := _AintGateway.contract.Call(opts, &out, "balanceOf" , account)
			
			if err != nil {
				return *new(*big.Int),  err
			}
			
			out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

			return out0,  err
			
		}

		// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
		//
		// Solidity: function balanceOf(address account) view returns(uint256)
		func (_AintGateway *AintGatewaySession) BalanceOf( account common.Address ) ( *big.Int,  error) {
		  return _AintGateway.Contract.BalanceOf(&_AintGateway.CallOpts , account)
		}

		// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
		//
		// Solidity: function balanceOf(address account) view returns(uint256)
		func (_AintGateway *AintGatewayCallerSession) BalanceOf( account common.Address ) ( *big.Int,  error) {
		  return _AintGateway.Contract.BalanceOf(&_AintGateway.CallOpts , account)
		}
	
		// Decimals is a free data retrieval call binding the contract method 0x313ce567.
		//
		// Solidity: function decimals() view returns(uint8)
		func (_AintGateway *AintGatewayCaller) Decimals(opts *bind.CallOpts ) (uint8, error) {
			var out []interface{}
			err := _AintGateway.contract.Call(opts, &out, "decimals" )
			
			if err != nil {
				return *new(uint8),  err
			}
			
			out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

			return out0,  err
			
		}

		// Decimals is a free data retrieval call binding the contract method 0x313ce567.
		//
		// Solidity: function decimals() view returns(uint8)
		func (_AintGateway *AintGatewaySession) Decimals() ( uint8,  error) {
		  return _AintGateway.Contract.Decimals(&_AintGateway.CallOpts )
		}

		// Decimals is a free data retrieval call binding the contract method 0x313ce567.
		//
		// Solidity: function decimals() view returns(uint8)
		func (_AintGateway *AintGatewayCallerSession) Decimals() ( uint8,  error) {
		  return _AintGateway.Contract.Decimals(&_AintGateway.CallOpts )
		}
	
		// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
		//
		// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
		func (_AintGateway *AintGatewayCaller) GetRoleAdmin(opts *bind.CallOpts , role [32]byte ) ([32]byte, error) {
			var out []interface{}
			err := _AintGateway.contract.Call(opts, &out, "getRoleAdmin" , role)
			
			if err != nil {
				return *new([32]byte),  err
			}
			
			out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

			return out0,  err
			
		}

		// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
		//
		// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
		func (_AintGateway *AintGatewaySession) GetRoleAdmin( role [32]byte ) ( [32]byte,  error) {
		  return _AintGateway.Contract.GetRoleAdmin(&_AintGateway.CallOpts , role)
		}

		// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
		//
		// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
		func (_AintGateway *AintGatewayCallerSession) GetRoleAdmin( role [32]byte ) ( [32]byte,  error) {
		  return _AintGateway.Contract.GetRoleAdmin(&_AintGateway.CallOpts , role)
		}
	
		// GetWithdrawFee is a free data retrieval call binding the contract method 0x1540aa89.
		//
		// Solidity: function getWithdrawFee() view returns(uint256)
		func (_AintGateway *AintGatewayCaller) GetWithdrawFee(opts *bind.CallOpts ) (*big.Int, error) {
			var out []interface{}
			err := _AintGateway.contract.Call(opts, &out, "getWithdrawFee" )
			
			if err != nil {
				return *new(*big.Int),  err
			}
			
			out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

			return out0,  err
			
		}

		// GetWithdrawFee is a free data retrieval call binding the contract method 0x1540aa89.
		//
		// Solidity: function getWithdrawFee() view returns(uint256)
		func (_AintGateway *AintGatewaySession) GetWithdrawFee() ( *big.Int,  error) {
		  return _AintGateway.Contract.GetWithdrawFee(&_AintGateway.CallOpts )
		}

		// GetWithdrawFee is a free data retrieval call binding the contract method 0x1540aa89.
		//
		// Solidity: function getWithdrawFee() view returns(uint256)
		func (_AintGateway *AintGatewayCallerSession) GetWithdrawFee() ( *big.Int,  error) {
		  return _AintGateway.Contract.GetWithdrawFee(&_AintGateway.CallOpts )
		}
	
		// HasRole is a free data retrieval call binding the contract method 0x91d14854.
		//
		// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
		func (_AintGateway *AintGatewayCaller) HasRole(opts *bind.CallOpts , role [32]byte , account common.Address ) (bool, error) {
			var out []interface{}
			err := _AintGateway.contract.Call(opts, &out, "hasRole" , role, account)
			
			if err != nil {
				return *new(bool),  err
			}
			
			out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

			return out0,  err
			
		}

		// HasRole is a free data retrieval call binding the contract method 0x91d14854.
		//
		// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
		func (_AintGateway *AintGatewaySession) HasRole( role [32]byte , account common.Address ) ( bool,  error) {
		  return _AintGateway.Contract.HasRole(&_AintGateway.CallOpts , role, account)
		}

		// HasRole is a free data retrieval call binding the contract method 0x91d14854.
		//
		// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
		func (_AintGateway *AintGatewayCallerSession) HasRole( role [32]byte , account common.Address ) ( bool,  error) {
		  return _AintGateway.Contract.HasRole(&_AintGateway.CallOpts , role, account)
		}
	
		// Name is a free data retrieval call binding the contract method 0x06fdde03.
		//
		// Solidity: function name() view returns(string)
		func (_AintGateway *AintGatewayCaller) Name(opts *bind.CallOpts ) (string, error) {
			var out []interface{}
			err := _AintGateway.contract.Call(opts, &out, "name" )
			
			if err != nil {
				return *new(string),  err
			}
			
			out0 := *abi.ConvertType(out[0], new(string)).(*string)

			return out0,  err
			
		}

		// Name is a free data retrieval call binding the contract method 0x06fdde03.
		//
		// Solidity: function name() view returns(string)
		func (_AintGateway *AintGatewaySession) Name() ( string,  error) {
		  return _AintGateway.Contract.Name(&_AintGateway.CallOpts )
		}

		// Name is a free data retrieval call binding the contract method 0x06fdde03.
		//
		// Solidity: function name() view returns(string)
		func (_AintGateway *AintGatewayCallerSession) Name() ( string,  error) {
		  return _AintGateway.Contract.Name(&_AintGateway.CallOpts )
		}
	
		// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
		//
		// Solidity: function owner() view returns(address)
		func (_AintGateway *AintGatewayCaller) Owner(opts *bind.CallOpts ) (common.Address, error) {
			var out []interface{}
			err := _AintGateway.contract.Call(opts, &out, "owner" )
			
			if err != nil {
				return *new(common.Address),  err
			}
			
			out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

			return out0,  err
			
		}

		// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
		//
		// Solidity: function owner() view returns(address)
		func (_AintGateway *AintGatewaySession) Owner() ( common.Address,  error) {
		  return _AintGateway.Contract.Owner(&_AintGateway.CallOpts )
		}

		// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
		//
		// Solidity: function owner() view returns(address)
		func (_AintGateway *AintGatewayCallerSession) Owner() ( common.Address,  error) {
		  return _AintGateway.Contract.Owner(&_AintGateway.CallOpts )
		}
	
		// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
		//
		// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
		func (_AintGateway *AintGatewayCaller) SupportsInterface(opts *bind.CallOpts , interfaceId [4]byte ) (bool, error) {
			var out []interface{}
			err := _AintGateway.contract.Call(opts, &out, "supportsInterface" , interfaceId)
			
			if err != nil {
				return *new(bool),  err
			}
			
			out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

			return out0,  err
			
		}

		// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
		//
		// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
		func (_AintGateway *AintGatewaySession) SupportsInterface( interfaceId [4]byte ) ( bool,  error) {
		  return _AintGateway.Contract.SupportsInterface(&_AintGateway.CallOpts , interfaceId)
		}

		// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
		//
		// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
		func (_AintGateway *AintGatewayCallerSession) SupportsInterface( interfaceId [4]byte ) ( bool,  error) {
		  return _AintGateway.Contract.SupportsInterface(&_AintGateway.CallOpts , interfaceId)
		}
	
		// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
		//
		// Solidity: function symbol() view returns(string)
		func (_AintGateway *AintGatewayCaller) Symbol(opts *bind.CallOpts ) (string, error) {
			var out []interface{}
			err := _AintGateway.contract.Call(opts, &out, "symbol" )
			
			if err != nil {
				return *new(string),  err
			}
			
			out0 := *abi.ConvertType(out[0], new(string)).(*string)

			return out0,  err
			
		}

		// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
		//
		// Solidity: function symbol() view returns(string)
		func (_AintGateway *AintGatewaySession) Symbol() ( string,  error) {
		  return _AintGateway.Contract.Symbol(&_AintGateway.CallOpts )
		}

		// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
		//
		// Solidity: function symbol() view returns(string)
		func (_AintGateway *AintGatewayCallerSession) Symbol() ( string,  error) {
		  return _AintGateway.Contract.Symbol(&_AintGateway.CallOpts )
		}
	
		// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
		//
		// Solidity: function totalSupply() view returns(uint256)
		func (_AintGateway *AintGatewayCaller) TotalSupply(opts *bind.CallOpts ) (*big.Int, error) {
			var out []interface{}
			err := _AintGateway.contract.Call(opts, &out, "totalSupply" )
			
			if err != nil {
				return *new(*big.Int),  err
			}
			
			out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

			return out0,  err
			
		}

		// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
		//
		// Solidity: function totalSupply() view returns(uint256)
		func (_AintGateway *AintGatewaySession) TotalSupply() ( *big.Int,  error) {
		  return _AintGateway.Contract.TotalSupply(&_AintGateway.CallOpts )
		}

		// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
		//
		// Solidity: function totalSupply() view returns(uint256)
		func (_AintGateway *AintGatewayCallerSession) TotalSupply() ( *big.Int,  error) {
		  return _AintGateway.Contract.TotalSupply(&_AintGateway.CallOpts )
		}
	
		// WithdrawExists is a free data retrieval call binding the contract method 0x90cde51e.
		//
		// Solidity: function withdrawExists(address addr) view returns(bool)
		func (_AintGateway *AintGatewayCaller) WithdrawExists(opts *bind.CallOpts , addr common.Address ) (bool, error) {
			var out []interface{}
			err := _AintGateway.contract.Call(opts, &out, "withdrawExists" , addr)
			
			if err != nil {
				return *new(bool),  err
			}
			
			out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

			return out0,  err
			
		}

		// WithdrawExists is a free data retrieval call binding the contract method 0x90cde51e.
		//
		// Solidity: function withdrawExists(address addr) view returns(bool)
		func (_AintGateway *AintGatewaySession) WithdrawExists( addr common.Address ) ( bool,  error) {
		  return _AintGateway.Contract.WithdrawExists(&_AintGateway.CallOpts , addr)
		}

		// WithdrawExists is a free data retrieval call binding the contract method 0x90cde51e.
		//
		// Solidity: function withdrawExists(address addr) view returns(bool)
		func (_AintGateway *AintGatewayCallerSession) WithdrawExists( addr common.Address ) ( bool,  error) {
		  return _AintGateway.Contract.WithdrawExists(&_AintGateway.CallOpts , addr)
		}
	

	
		// AddWithdraw is a paid mutator transaction binding the contract method 0xf389a512.
		//
		// Solidity: function addWithdraw(address addr, uint256 amount) returns()
		func (_AintGateway *AintGatewayTransactor) AddWithdraw(opts *bind.TransactOpts , addr common.Address , amount *big.Int ) (*types.Transaction, error) {
			return _AintGateway.contract.Transact(opts, "addWithdraw" , addr, amount)
		}

		// AddWithdraw is a paid mutator transaction binding the contract method 0xf389a512.
		//
		// Solidity: function addWithdraw(address addr, uint256 amount) returns()
		func (_AintGateway *AintGatewaySession) AddWithdraw( addr common.Address , amount *big.Int ) (*types.Transaction, error) {
		  return _AintGateway.Contract.AddWithdraw(&_AintGateway.TransactOpts , addr, amount)
		}

		// AddWithdraw is a paid mutator transaction binding the contract method 0xf389a512.
		//
		// Solidity: function addWithdraw(address addr, uint256 amount) returns()
		func (_AintGateway *AintGatewayTransactorSession) AddWithdraw( addr common.Address , amount *big.Int ) (*types.Transaction, error) {
		  return _AintGateway.Contract.AddWithdraw(&_AintGateway.TransactOpts , addr, amount)
		}
	
		// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
		//
		// Solidity: function approve(address spender, uint256 value) returns(bool)
		func (_AintGateway *AintGatewayTransactor) Approve(opts *bind.TransactOpts , spender common.Address , value *big.Int ) (*types.Transaction, error) {
			return _AintGateway.contract.Transact(opts, "approve" , spender, value)
		}

		// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
		//
		// Solidity: function approve(address spender, uint256 value) returns(bool)
		func (_AintGateway *AintGatewaySession) Approve( spender common.Address , value *big.Int ) (*types.Transaction, error) {
		  return _AintGateway.Contract.Approve(&_AintGateway.TransactOpts , spender, value)
		}

		// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
		//
		// Solidity: function approve(address spender, uint256 value) returns(bool)
		func (_AintGateway *AintGatewayTransactorSession) Approve( spender common.Address , value *big.Int ) (*types.Transaction, error) {
		  return _AintGateway.Contract.Approve(&_AintGateway.TransactOpts , spender, value)
		}
	
		// Burn is a paid mutator transaction binding the contract method 0x42966c68.
		//
		// Solidity: function burn(uint256 value) returns()
		func (_AintGateway *AintGatewayTransactor) Burn(opts *bind.TransactOpts , value *big.Int ) (*types.Transaction, error) {
			return _AintGateway.contract.Transact(opts, "burn" , value)
		}

		// Burn is a paid mutator transaction binding the contract method 0x42966c68.
		//
		// Solidity: function burn(uint256 value) returns()
		func (_AintGateway *AintGatewaySession) Burn( value *big.Int ) (*types.Transaction, error) {
		  return _AintGateway.Contract.Burn(&_AintGateway.TransactOpts , value)
		}

		// Burn is a paid mutator transaction binding the contract method 0x42966c68.
		//
		// Solidity: function burn(uint256 value) returns()
		func (_AintGateway *AintGatewayTransactorSession) Burn( value *big.Int ) (*types.Transaction, error) {
		  return _AintGateway.Contract.Burn(&_AintGateway.TransactOpts , value)
		}
	
		// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
		//
		// Solidity: function burnFrom(address account, uint256 value) returns()
		func (_AintGateway *AintGatewayTransactor) BurnFrom(opts *bind.TransactOpts , account common.Address , value *big.Int ) (*types.Transaction, error) {
			return _AintGateway.contract.Transact(opts, "burnFrom" , account, value)
		}

		// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
		//
		// Solidity: function burnFrom(address account, uint256 value) returns()
		func (_AintGateway *AintGatewaySession) BurnFrom( account common.Address , value *big.Int ) (*types.Transaction, error) {
		  return _AintGateway.Contract.BurnFrom(&_AintGateway.TransactOpts , account, value)
		}

		// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
		//
		// Solidity: function burnFrom(address account, uint256 value) returns()
		func (_AintGateway *AintGatewayTransactorSession) BurnFrom( account common.Address , value *big.Int ) (*types.Transaction, error) {
		  return _AintGateway.Contract.BurnFrom(&_AintGateway.TransactOpts , account, value)
		}
	
		// CancelWithdraw is a paid mutator transaction binding the contract method 0xe9919629.
		//
		// Solidity: function cancelWithdraw(address addr) returns()
		func (_AintGateway *AintGatewayTransactor) CancelWithdraw(opts *bind.TransactOpts , addr common.Address ) (*types.Transaction, error) {
			return _AintGateway.contract.Transact(opts, "cancelWithdraw" , addr)
		}

		// CancelWithdraw is a paid mutator transaction binding the contract method 0xe9919629.
		//
		// Solidity: function cancelWithdraw(address addr) returns()
		func (_AintGateway *AintGatewaySession) CancelWithdraw( addr common.Address ) (*types.Transaction, error) {
		  return _AintGateway.Contract.CancelWithdraw(&_AintGateway.TransactOpts , addr)
		}

		// CancelWithdraw is a paid mutator transaction binding the contract method 0xe9919629.
		//
		// Solidity: function cancelWithdraw(address addr) returns()
		func (_AintGateway *AintGatewayTransactorSession) CancelWithdraw( addr common.Address ) (*types.Transaction, error) {
		  return _AintGateway.Contract.CancelWithdraw(&_AintGateway.TransactOpts , addr)
		}
	
		// Deposit is a paid mutator transaction binding the contract method 0x8e27d719.
		//
		// Solidity: function deposit(string to, uint256 amount) returns(string)
		func (_AintGateway *AintGatewayTransactor) Deposit(opts *bind.TransactOpts , to string , amount *big.Int ) (*types.Transaction, error) {
			return _AintGateway.contract.Transact(opts, "deposit" , to, amount)
		}

		// Deposit is a paid mutator transaction binding the contract method 0x8e27d719.
		//
		// Solidity: function deposit(string to, uint256 amount) returns(string)
		func (_AintGateway *AintGatewaySession) Deposit( to string , amount *big.Int ) (*types.Transaction, error) {
		  return _AintGateway.Contract.Deposit(&_AintGateway.TransactOpts , to, amount)
		}

		// Deposit is a paid mutator transaction binding the contract method 0x8e27d719.
		//
		// Solidity: function deposit(string to, uint256 amount) returns(string)
		func (_AintGateway *AintGatewayTransactorSession) Deposit( to string , amount *big.Int ) (*types.Transaction, error) {
		  return _AintGateway.Contract.Deposit(&_AintGateway.TransactOpts , to, amount)
		}
	
		// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
		//
		// Solidity: function grantRole(bytes32 role, address account) returns()
		func (_AintGateway *AintGatewayTransactor) GrantRole(opts *bind.TransactOpts , role [32]byte , account common.Address ) (*types.Transaction, error) {
			return _AintGateway.contract.Transact(opts, "grantRole" , role, account)
		}

		// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
		//
		// Solidity: function grantRole(bytes32 role, address account) returns()
		func (_AintGateway *AintGatewaySession) GrantRole( role [32]byte , account common.Address ) (*types.Transaction, error) {
		  return _AintGateway.Contract.GrantRole(&_AintGateway.TransactOpts , role, account)
		}

		// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
		//
		// Solidity: function grantRole(bytes32 role, address account) returns()
		func (_AintGateway *AintGatewayTransactorSession) GrantRole( role [32]byte , account common.Address ) (*types.Transaction, error) {
		  return _AintGateway.Contract.GrantRole(&_AintGateway.TransactOpts , role, account)
		}
	
		// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
		//
		// Solidity: function renounceOwnership() returns()
		func (_AintGateway *AintGatewayTransactor) RenounceOwnership(opts *bind.TransactOpts ) (*types.Transaction, error) {
			return _AintGateway.contract.Transact(opts, "renounceOwnership" )
		}

		// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
		//
		// Solidity: function renounceOwnership() returns()
		func (_AintGateway *AintGatewaySession) RenounceOwnership() (*types.Transaction, error) {
		  return _AintGateway.Contract.RenounceOwnership(&_AintGateway.TransactOpts )
		}

		// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
		//
		// Solidity: function renounceOwnership() returns()
		func (_AintGateway *AintGatewayTransactorSession) RenounceOwnership() (*types.Transaction, error) {
		  return _AintGateway.Contract.RenounceOwnership(&_AintGateway.TransactOpts )
		}
	
		// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
		//
		// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
		func (_AintGateway *AintGatewayTransactor) RenounceRole(opts *bind.TransactOpts , role [32]byte , callerConfirmation common.Address ) (*types.Transaction, error) {
			return _AintGateway.contract.Transact(opts, "renounceRole" , role, callerConfirmation)
		}

		// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
		//
		// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
		func (_AintGateway *AintGatewaySession) RenounceRole( role [32]byte , callerConfirmation common.Address ) (*types.Transaction, error) {
		  return _AintGateway.Contract.RenounceRole(&_AintGateway.TransactOpts , role, callerConfirmation)
		}

		// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
		//
		// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
		func (_AintGateway *AintGatewayTransactorSession) RenounceRole( role [32]byte , callerConfirmation common.Address ) (*types.Transaction, error) {
		  return _AintGateway.Contract.RenounceRole(&_AintGateway.TransactOpts , role, callerConfirmation)
		}
	
		// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
		//
		// Solidity: function revokeRole(bytes32 role, address account) returns()
		func (_AintGateway *AintGatewayTransactor) RevokeRole(opts *bind.TransactOpts , role [32]byte , account common.Address ) (*types.Transaction, error) {
			return _AintGateway.contract.Transact(opts, "revokeRole" , role, account)
		}

		// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
		//
		// Solidity: function revokeRole(bytes32 role, address account) returns()
		func (_AintGateway *AintGatewaySession) RevokeRole( role [32]byte , account common.Address ) (*types.Transaction, error) {
		  return _AintGateway.Contract.RevokeRole(&_AintGateway.TransactOpts , role, account)
		}

		// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
		//
		// Solidity: function revokeRole(bytes32 role, address account) returns()
		func (_AintGateway *AintGatewayTransactorSession) RevokeRole( role [32]byte , account common.Address ) (*types.Transaction, error) {
		  return _AintGateway.Contract.RevokeRole(&_AintGateway.TransactOpts , role, account)
		}
	
		// SetWithdrawFee is a paid mutator transaction binding the contract method 0xb6ac642a.
		//
		// Solidity: function setWithdrawFee(uint256 amount) returns()
		func (_AintGateway *AintGatewayTransactor) SetWithdrawFee(opts *bind.TransactOpts , amount *big.Int ) (*types.Transaction, error) {
			return _AintGateway.contract.Transact(opts, "setWithdrawFee" , amount)
		}

		// SetWithdrawFee is a paid mutator transaction binding the contract method 0xb6ac642a.
		//
		// Solidity: function setWithdrawFee(uint256 amount) returns()
		func (_AintGateway *AintGatewaySession) SetWithdrawFee( amount *big.Int ) (*types.Transaction, error) {
		  return _AintGateway.Contract.SetWithdrawFee(&_AintGateway.TransactOpts , amount)
		}

		// SetWithdrawFee is a paid mutator transaction binding the contract method 0xb6ac642a.
		//
		// Solidity: function setWithdrawFee(uint256 amount) returns()
		func (_AintGateway *AintGatewayTransactorSession) SetWithdrawFee( amount *big.Int ) (*types.Transaction, error) {
		  return _AintGateway.Contract.SetWithdrawFee(&_AintGateway.TransactOpts , amount)
		}
	
		// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
		//
		// Solidity: function transfer(address to, uint256 value) returns(bool)
		func (_AintGateway *AintGatewayTransactor) Transfer(opts *bind.TransactOpts , to common.Address , value *big.Int ) (*types.Transaction, error) {
			return _AintGateway.contract.Transact(opts, "transfer" , to, value)
		}

		// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
		//
		// Solidity: function transfer(address to, uint256 value) returns(bool)
		func (_AintGateway *AintGatewaySession) Transfer( to common.Address , value *big.Int ) (*types.Transaction, error) {
		  return _AintGateway.Contract.Transfer(&_AintGateway.TransactOpts , to, value)
		}

		// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
		//
		// Solidity: function transfer(address to, uint256 value) returns(bool)
		func (_AintGateway *AintGatewayTransactorSession) Transfer( to common.Address , value *big.Int ) (*types.Transaction, error) {
		  return _AintGateway.Contract.Transfer(&_AintGateway.TransactOpts , to, value)
		}
	
		// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
		//
		// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
		func (_AintGateway *AintGatewayTransactor) TransferFrom(opts *bind.TransactOpts , from common.Address , to common.Address , value *big.Int ) (*types.Transaction, error) {
			return _AintGateway.contract.Transact(opts, "transferFrom" , from, to, value)
		}

		// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
		//
		// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
		func (_AintGateway *AintGatewaySession) TransferFrom( from common.Address , to common.Address , value *big.Int ) (*types.Transaction, error) {
		  return _AintGateway.Contract.TransferFrom(&_AintGateway.TransactOpts , from, to, value)
		}

		// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
		//
		// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
		func (_AintGateway *AintGatewayTransactorSession) TransferFrom( from common.Address , to common.Address , value *big.Int ) (*types.Transaction, error) {
		  return _AintGateway.Contract.TransferFrom(&_AintGateway.TransactOpts , from, to, value)
		}
	
		// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
		//
		// Solidity: function transferOwnership(address newOwner) returns()
		func (_AintGateway *AintGatewayTransactor) TransferOwnership(opts *bind.TransactOpts , newOwner common.Address ) (*types.Transaction, error) {
			return _AintGateway.contract.Transact(opts, "transferOwnership" , newOwner)
		}

		// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
		//
		// Solidity: function transferOwnership(address newOwner) returns()
		func (_AintGateway *AintGatewaySession) TransferOwnership( newOwner common.Address ) (*types.Transaction, error) {
		  return _AintGateway.Contract.TransferOwnership(&_AintGateway.TransactOpts , newOwner)
		}

		// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
		//
		// Solidity: function transferOwnership(address newOwner) returns()
		func (_AintGateway *AintGatewayTransactorSession) TransferOwnership( newOwner common.Address ) (*types.Transaction, error) {
		  return _AintGateway.Contract.TransferOwnership(&_AintGateway.TransactOpts , newOwner)
		}
	
		// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
		//
		// Solidity: function withdraw() payable returns()
		func (_AintGateway *AintGatewayTransactor) Withdraw(opts *bind.TransactOpts ) (*types.Transaction, error) {
			return _AintGateway.contract.Transact(opts, "withdraw" )
		}

		// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
		//
		// Solidity: function withdraw() payable returns()
		func (_AintGateway *AintGatewaySession) Withdraw() (*types.Transaction, error) {
		  return _AintGateway.Contract.Withdraw(&_AintGateway.TransactOpts )
		}

		// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
		//
		// Solidity: function withdraw() payable returns()
		func (_AintGateway *AintGatewayTransactorSession) Withdraw() (*types.Transaction, error) {
		  return _AintGateway.Contract.Withdraw(&_AintGateway.TransactOpts )
		}
	

	

	

	
		// AintGatewayApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the AintGateway contract.
		type AintGatewayApprovalIterator struct {
			Event *AintGatewayApproval // Event containing the contract specifics and raw log

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
		func (it *AintGatewayApprovalIterator) Next() bool {
			// If the iterator failed, stop iterating
			if (it.fail != nil) {
				return false
			}
			// If the iterator completed, deliver directly whatever's available
			if (it.done) {
				select {
				case log := <-it.logs:
					it.Event = new(AintGatewayApproval)
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
				it.Event = new(AintGatewayApproval)
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
		func (it *AintGatewayApprovalIterator) Error() error {
			return it.fail
		}
		// Close terminates the iteration process, releasing any pending underlying
		// resources.
		func (it *AintGatewayApprovalIterator) Close() error {
			it.sub.Unsubscribe()
			return nil
		}

		// AintGatewayApproval represents a Approval event raised by the AintGateway contract.
		type AintGatewayApproval struct { 
			Owner common.Address; 
			Spender common.Address; 
			Value *big.Int; 
			Raw types.Log // Blockchain specific contextual infos
		}

		// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
		//
		// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
 		func (_AintGateway *AintGatewayFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*AintGatewayApprovalIterator, error) {
			
			var ownerRule []interface{}
			for _, ownerItem := range owner {
				ownerRule = append(ownerRule, ownerItem)
			}
			var spenderRule []interface{}
			for _, spenderItem := range spender {
				spenderRule = append(spenderRule, spenderItem)
			}
			

			logs, sub, err := _AintGateway.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
			if err != nil {
				return nil, err
			}
			return &AintGatewayApprovalIterator{contract: _AintGateway.contract, event: "Approval", logs: logs, sub: sub}, nil
 		}

		// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
		//
		// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
		func (_AintGateway *AintGatewayFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *AintGatewayApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {
			
			var ownerRule []interface{}
			for _, ownerItem := range owner {
				ownerRule = append(ownerRule, ownerItem)
			}
			var spenderRule []interface{}
			for _, spenderItem := range spender {
				spenderRule = append(spenderRule, spenderItem)
			}
			

			logs, sub, err := _AintGateway.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
			if err != nil {
				return nil, err
			}
			return event.NewSubscription(func(quit <-chan struct{}) error {
				defer sub.Unsubscribe()
				for {
					select {
					case log := <-logs:
						// New log arrived, parse the event and forward to the user
						event := new(AintGatewayApproval)
						if err := _AintGateway.contract.UnpackLog(event, "Approval", log); err != nil {
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
		// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
		func (_AintGateway *AintGatewayFilterer) ParseApproval(log types.Log) (*AintGatewayApproval, error) {
			event := new(AintGatewayApproval)
			if err := _AintGateway.contract.UnpackLog(event, "Approval", log); err != nil {
				return nil, err
			}
			event.Raw = log
			return event, nil
		}

 	
		// AintGatewayOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AintGateway contract.
		type AintGatewayOwnershipTransferredIterator struct {
			Event *AintGatewayOwnershipTransferred // Event containing the contract specifics and raw log

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
		func (it *AintGatewayOwnershipTransferredIterator) Next() bool {
			// If the iterator failed, stop iterating
			if (it.fail != nil) {
				return false
			}
			// If the iterator completed, deliver directly whatever's available
			if (it.done) {
				select {
				case log := <-it.logs:
					it.Event = new(AintGatewayOwnershipTransferred)
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
				it.Event = new(AintGatewayOwnershipTransferred)
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
		func (it *AintGatewayOwnershipTransferredIterator) Error() error {
			return it.fail
		}
		// Close terminates the iteration process, releasing any pending underlying
		// resources.
		func (it *AintGatewayOwnershipTransferredIterator) Close() error {
			it.sub.Unsubscribe()
			return nil
		}

		// AintGatewayOwnershipTransferred represents a OwnershipTransferred event raised by the AintGateway contract.
		type AintGatewayOwnershipTransferred struct { 
			PreviousOwner common.Address; 
			NewOwner common.Address; 
			Raw types.Log // Blockchain specific contextual infos
		}

		// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
		//
		// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
 		func (_AintGateway *AintGatewayFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AintGatewayOwnershipTransferredIterator, error) {
			
			var previousOwnerRule []interface{}
			for _, previousOwnerItem := range previousOwner {
				previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
			}
			var newOwnerRule []interface{}
			for _, newOwnerItem := range newOwner {
				newOwnerRule = append(newOwnerRule, newOwnerItem)
			}

			logs, sub, err := _AintGateway.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
			if err != nil {
				return nil, err
			}
			return &AintGatewayOwnershipTransferredIterator{contract: _AintGateway.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
 		}

		// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
		//
		// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
		func (_AintGateway *AintGatewayFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AintGatewayOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {
			
			var previousOwnerRule []interface{}
			for _, previousOwnerItem := range previousOwner {
				previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
			}
			var newOwnerRule []interface{}
			for _, newOwnerItem := range newOwner {
				newOwnerRule = append(newOwnerRule, newOwnerItem)
			}

			logs, sub, err := _AintGateway.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
			if err != nil {
				return nil, err
			}
			return event.NewSubscription(func(quit <-chan struct{}) error {
				defer sub.Unsubscribe()
				for {
					select {
					case log := <-logs:
						// New log arrived, parse the event and forward to the user
						event := new(AintGatewayOwnershipTransferred)
						if err := _AintGateway.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

		// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
		//
		// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
		func (_AintGateway *AintGatewayFilterer) ParseOwnershipTransferred(log types.Log) (*AintGatewayOwnershipTransferred, error) {
			event := new(AintGatewayOwnershipTransferred)
			if err := _AintGateway.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
				return nil, err
			}
			event.Raw = log
			return event, nil
		}

 	
		// AintGatewayRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the AintGateway contract.
		type AintGatewayRoleAdminChangedIterator struct {
			Event *AintGatewayRoleAdminChanged // Event containing the contract specifics and raw log

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
		func (it *AintGatewayRoleAdminChangedIterator) Next() bool {
			// If the iterator failed, stop iterating
			if (it.fail != nil) {
				return false
			}
			// If the iterator completed, deliver directly whatever's available
			if (it.done) {
				select {
				case log := <-it.logs:
					it.Event = new(AintGatewayRoleAdminChanged)
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
				it.Event = new(AintGatewayRoleAdminChanged)
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
		func (it *AintGatewayRoleAdminChangedIterator) Error() error {
			return it.fail
		}
		// Close terminates the iteration process, releasing any pending underlying
		// resources.
		func (it *AintGatewayRoleAdminChangedIterator) Close() error {
			it.sub.Unsubscribe()
			return nil
		}

		// AintGatewayRoleAdminChanged represents a RoleAdminChanged event raised by the AintGateway contract.
		type AintGatewayRoleAdminChanged struct { 
			Role [32]byte; 
			PreviousAdminRole [32]byte; 
			NewAdminRole [32]byte; 
			Raw types.Log // Blockchain specific contextual infos
		}

		// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
		//
		// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
 		func (_AintGateway *AintGatewayFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*AintGatewayRoleAdminChangedIterator, error) {
			
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

			logs, sub, err := _AintGateway.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
			if err != nil {
				return nil, err
			}
			return &AintGatewayRoleAdminChangedIterator{contract: _AintGateway.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
 		}

		// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
		//
		// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
		func (_AintGateway *AintGatewayFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *AintGatewayRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {
			
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

			logs, sub, err := _AintGateway.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
			if err != nil {
				return nil, err
			}
			return event.NewSubscription(func(quit <-chan struct{}) error {
				defer sub.Unsubscribe()
				for {
					select {
					case log := <-logs:
						// New log arrived, parse the event and forward to the user
						event := new(AintGatewayRoleAdminChanged)
						if err := _AintGateway.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
		func (_AintGateway *AintGatewayFilterer) ParseRoleAdminChanged(log types.Log) (*AintGatewayRoleAdminChanged, error) {
			event := new(AintGatewayRoleAdminChanged)
			if err := _AintGateway.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
				return nil, err
			}
			event.Raw = log
			return event, nil
		}

 	
		// AintGatewayRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the AintGateway contract.
		type AintGatewayRoleGrantedIterator struct {
			Event *AintGatewayRoleGranted // Event containing the contract specifics and raw log

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
		func (it *AintGatewayRoleGrantedIterator) Next() bool {
			// If the iterator failed, stop iterating
			if (it.fail != nil) {
				return false
			}
			// If the iterator completed, deliver directly whatever's available
			if (it.done) {
				select {
				case log := <-it.logs:
					it.Event = new(AintGatewayRoleGranted)
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
				it.Event = new(AintGatewayRoleGranted)
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
		func (it *AintGatewayRoleGrantedIterator) Error() error {
			return it.fail
		}
		// Close terminates the iteration process, releasing any pending underlying
		// resources.
		func (it *AintGatewayRoleGrantedIterator) Close() error {
			it.sub.Unsubscribe()
			return nil
		}

		// AintGatewayRoleGranted represents a RoleGranted event raised by the AintGateway contract.
		type AintGatewayRoleGranted struct { 
			Role [32]byte; 
			Account common.Address; 
			Sender common.Address; 
			Raw types.Log // Blockchain specific contextual infos
		}

		// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
		//
		// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
 		func (_AintGateway *AintGatewayFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AintGatewayRoleGrantedIterator, error) {
			
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

			logs, sub, err := _AintGateway.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
			if err != nil {
				return nil, err
			}
			return &AintGatewayRoleGrantedIterator{contract: _AintGateway.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
 		}

		// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
		//
		// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
		func (_AintGateway *AintGatewayFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *AintGatewayRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {
			
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

			logs, sub, err := _AintGateway.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
			if err != nil {
				return nil, err
			}
			return event.NewSubscription(func(quit <-chan struct{}) error {
				defer sub.Unsubscribe()
				for {
					select {
					case log := <-logs:
						// New log arrived, parse the event and forward to the user
						event := new(AintGatewayRoleGranted)
						if err := _AintGateway.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
		func (_AintGateway *AintGatewayFilterer) ParseRoleGranted(log types.Log) (*AintGatewayRoleGranted, error) {
			event := new(AintGatewayRoleGranted)
			if err := _AintGateway.contract.UnpackLog(event, "RoleGranted", log); err != nil {
				return nil, err
			}
			event.Raw = log
			return event, nil
		}

 	
		// AintGatewayRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the AintGateway contract.
		type AintGatewayRoleRevokedIterator struct {
			Event *AintGatewayRoleRevoked // Event containing the contract specifics and raw log

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
		func (it *AintGatewayRoleRevokedIterator) Next() bool {
			// If the iterator failed, stop iterating
			if (it.fail != nil) {
				return false
			}
			// If the iterator completed, deliver directly whatever's available
			if (it.done) {
				select {
				case log := <-it.logs:
					it.Event = new(AintGatewayRoleRevoked)
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
				it.Event = new(AintGatewayRoleRevoked)
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
		func (it *AintGatewayRoleRevokedIterator) Error() error {
			return it.fail
		}
		// Close terminates the iteration process, releasing any pending underlying
		// resources.
		func (it *AintGatewayRoleRevokedIterator) Close() error {
			it.sub.Unsubscribe()
			return nil
		}

		// AintGatewayRoleRevoked represents a RoleRevoked event raised by the AintGateway contract.
		type AintGatewayRoleRevoked struct { 
			Role [32]byte; 
			Account common.Address; 
			Sender common.Address; 
			Raw types.Log // Blockchain specific contextual infos
		}

		// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
		//
		// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
 		func (_AintGateway *AintGatewayFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AintGatewayRoleRevokedIterator, error) {
			
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

			logs, sub, err := _AintGateway.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
			if err != nil {
				return nil, err
			}
			return &AintGatewayRoleRevokedIterator{contract: _AintGateway.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
 		}

		// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
		//
		// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
		func (_AintGateway *AintGatewayFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *AintGatewayRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {
			
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

			logs, sub, err := _AintGateway.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
			if err != nil {
				return nil, err
			}
			return event.NewSubscription(func(quit <-chan struct{}) error {
				defer sub.Unsubscribe()
				for {
					select {
					case log := <-logs:
						// New log arrived, parse the event and forward to the user
						event := new(AintGatewayRoleRevoked)
						if err := _AintGateway.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
		func (_AintGateway *AintGatewayFilterer) ParseRoleRevoked(log types.Log) (*AintGatewayRoleRevoked, error) {
			event := new(AintGatewayRoleRevoked)
			if err := _AintGateway.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
				return nil, err
			}
			event.Raw = log
			return event, nil
		}

 	
		// AintGatewayTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the AintGateway contract.
		type AintGatewayTransferIterator struct {
			Event *AintGatewayTransfer // Event containing the contract specifics and raw log

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
		func (it *AintGatewayTransferIterator) Next() bool {
			// If the iterator failed, stop iterating
			if (it.fail != nil) {
				return false
			}
			// If the iterator completed, deliver directly whatever's available
			if (it.done) {
				select {
				case log := <-it.logs:
					it.Event = new(AintGatewayTransfer)
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
				it.Event = new(AintGatewayTransfer)
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
		func (it *AintGatewayTransferIterator) Error() error {
			return it.fail
		}
		// Close terminates the iteration process, releasing any pending underlying
		// resources.
		func (it *AintGatewayTransferIterator) Close() error {
			it.sub.Unsubscribe()
			return nil
		}

		// AintGatewayTransfer represents a Transfer event raised by the AintGateway contract.
		type AintGatewayTransfer struct { 
			From common.Address; 
			To common.Address; 
			Value *big.Int; 
			Raw types.Log // Blockchain specific contextual infos
		}

		// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
		//
		// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
 		func (_AintGateway *AintGatewayFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*AintGatewayTransferIterator, error) {
			
			var fromRule []interface{}
			for _, fromItem := range from {
				fromRule = append(fromRule, fromItem)
			}
			var toRule []interface{}
			for _, toItem := range to {
				toRule = append(toRule, toItem)
			}
			

			logs, sub, err := _AintGateway.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
			if err != nil {
				return nil, err
			}
			return &AintGatewayTransferIterator{contract: _AintGateway.contract, event: "Transfer", logs: logs, sub: sub}, nil
 		}

		// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
		//
		// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
		func (_AintGateway *AintGatewayFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *AintGatewayTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {
			
			var fromRule []interface{}
			for _, fromItem := range from {
				fromRule = append(fromRule, fromItem)
			}
			var toRule []interface{}
			for _, toItem := range to {
				toRule = append(toRule, toItem)
			}
			

			logs, sub, err := _AintGateway.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
			if err != nil {
				return nil, err
			}
			return event.NewSubscription(func(quit <-chan struct{}) error {
				defer sub.Unsubscribe()
				for {
					select {
					case log := <-logs:
						// New log arrived, parse the event and forward to the user
						event := new(AintGatewayTransfer)
						if err := _AintGateway.contract.UnpackLog(event, "Transfer", log); err != nil {
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
		// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
		func (_AintGateway *AintGatewayFilterer) ParseTransfer(log types.Log) (*AintGatewayTransfer, error) {
			event := new(AintGatewayTransfer)
			if err := _AintGateway.contract.UnpackLog(event, "Transfer", log); err != nil {
				return nil, err
			}
			event.Raw = log
			return event, nil
		}