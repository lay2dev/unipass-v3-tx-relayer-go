// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package entry

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

// EntryPendingState is an auto generated low-level Go binding around an user-defined struct.
type EntryPendingState struct {
	Key       UnipassUtilsPubKey
	Timestamp *big.Int
	ResetKeys bool
}

// EntryRecoveryEmail is an auto generated low-level Go binding around an user-defined struct.
type EntryRecoveryEmail struct {
	Threshold *big.Int
	Emails    [][32]byte
}

// UnipassUtilsPubKey is an auto generated low-level Go binding around an user-defined struct.
type UnipassUtilsPubKey struct {
	KeyType *big.Int
	Key     []byte
}

// EntryMetaData contains all meta data concerning the Entry contract.
var EntryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"E_AddLocalKey\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"E_CancelRecovery\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"E_CompleteRecovery\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"E_DelLocalKey\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"E_QuickAddLocalKey\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"E_QuickRegister\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"E_Register\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"E_SetAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"timeLockSwitch\",\"type\":\"bool\"}],\"name\":\"E_SetTimeLockSwitch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"E_StartRecovery\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"name\":\"E_UpdateQuickLogin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"E_UpdateRecoveryEmail\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"registerEmail\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"username\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newKeyType\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"newKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"newKeySig\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"sigKeyIndex\",\"type\":\"uint256\"}],\"name\":\"addLocalKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"registerEmail\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"username\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"sigKeyIndex\",\"type\":\"uint256\"}],\"name\":\"cancelRecovery\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chainID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"registerEmail\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"keyType\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"}],\"name\":\"checkUserKey\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"registerEmail\",\"type\":\"bytes32\"}],\"name\":\"completeRecovery\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"registerEmail\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"username\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delKeyIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"sigKeyIndex\",\"type\":\"uint256\"}],\"name\":\"delLocalKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"registerEmail\",\"type\":\"bytes32\"}],\"name\":\"getUserKeys\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"keyType\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"}],\"internalType\":\"structUnipassUtils.PubKey[]\",\"name\":\"keys\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"registerEmail\",\"type\":\"bytes32\"}],\"name\":\"getUserName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"registerEmail\",\"type\":\"bytes32\"}],\"name\":\"getUserNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"registerEmail\",\"type\":\"bytes32\"}],\"name\":\"getUserPendingState\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"keyType\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"}],\"internalType\":\"structUnipassUtils.PubKey\",\"name\":\"key\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"resetKeys\",\"type\":\"bool\"}],\"internalType\":\"structEntry.PendingState\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_logic\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_chainID\",\"type\":\"uint8\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"registerEmail\",\"type\":\"bytes32\"}],\"name\":\"isUseEmailRegister\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"registerEmail\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"sigKeyType\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sigKey\",\"type\":\"bytes\"}],\"name\":\"isUserKey\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"logic\",\"outputs\":[{\"internalType\":\"contractLogic\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"registerEmail\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"username\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newKeyType\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"newKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"newKeySig\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"emailHeaders\",\"type\":\"bytes[]\"}],\"name\":\"quickAddLocalKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"registerEmail\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"originUsername\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"keyType\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"source\",\"type\":\"string\"}],\"name\":\"quickRegister\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"registerEmail\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"originUsername\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"emailHeader\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"keyType\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"source\",\"type\":\"string\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newAdmin\",\"type\":\"address\"}],\"name\":\"setAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_switch\",\"type\":\"bool\"}],\"name\":\"setTimeLockSwitch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"registerEmail\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"username\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"resetKeys\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"newKeyType\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"newKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"newKeySig\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"emailHeaders\",\"type\":\"bytes[]\"}],\"name\":\"startRecovery\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timeLockSwitch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalUsers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"registerEmail\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"username\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"quickLogin\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"sigKeyIndex\",\"type\":\"uint256\"}],\"name\":\"updateQuickLogin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"registerEmail\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"username\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"emails\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"sigKeyIndex\",\"type\":\"uint256\"}],\"name\":\"updateRecoveryEmail\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"users\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"registerEmail\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"originUsername\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isUseEmailRegister\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"quickLogin\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"emails\",\"type\":\"bytes32[]\"}],\"internalType\":\"structEntry.RecoveryEmail\",\"name\":\"recoveryEmails\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"assetContractAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isRecovering\",\"type\":\"bool\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"keyType\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"}],\"internalType\":\"structUnipassUtils.PubKey\",\"name\":\"key\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"resetKeys\",\"type\":\"bool\"}],\"internalType\":\"structEntry.PendingState\",\"name\":\"pendingState\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"source\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"registerEmail\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"sigKeyIndex\",\"type\":\"uint256\"}],\"name\":\"verify\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"registerEmail\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"sigKeyType\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sigKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"verifyUserSig\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// EntryABI is the input ABI used to generate the binding from.
// Deprecated: Use EntryMetaData.ABI instead.
var EntryABI = EntryMetaData.ABI

// Entry is an auto generated Go binding around an Ethereum contract.
type Entry struct {
	EntryCaller     // Read-only binding to the contract
	EntryTransactor // Write-only binding to the contract
	EntryFilterer   // Log filterer for contract events
}

// EntryCaller is an auto generated read-only Go binding around an Ethereum contract.
type EntryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EntryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EntryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EntryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EntryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EntrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EntrySession struct {
	Contract     *Entry            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EntryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EntryCallerSession struct {
	Contract *EntryCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// EntryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EntryTransactorSession struct {
	Contract     *EntryTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EntryRaw is an auto generated low-level Go binding around an Ethereum contract.
type EntryRaw struct {
	Contract *Entry // Generic contract binding to access the raw methods on
}

// EntryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EntryCallerRaw struct {
	Contract *EntryCaller // Generic read-only contract binding to access the raw methods on
}

// EntryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EntryTransactorRaw struct {
	Contract *EntryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEntry creates a new instance of Entry, bound to a specific deployed contract.
func NewEntry(address common.Address, backend bind.ContractBackend) (*Entry, error) {
	contract, err := bindEntry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Entry{EntryCaller: EntryCaller{contract: contract}, EntryTransactor: EntryTransactor{contract: contract}, EntryFilterer: EntryFilterer{contract: contract}}, nil
}

// NewEntryCaller creates a new read-only instance of Entry, bound to a specific deployed contract.
func NewEntryCaller(address common.Address, caller bind.ContractCaller) (*EntryCaller, error) {
	contract, err := bindEntry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EntryCaller{contract: contract}, nil
}

// NewEntryTransactor creates a new write-only instance of Entry, bound to a specific deployed contract.
func NewEntryTransactor(address common.Address, transactor bind.ContractTransactor) (*EntryTransactor, error) {
	contract, err := bindEntry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EntryTransactor{contract: contract}, nil
}

// NewEntryFilterer creates a new log filterer instance of Entry, bound to a specific deployed contract.
func NewEntryFilterer(address common.Address, filterer bind.ContractFilterer) (*EntryFilterer, error) {
	contract, err := bindEntry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EntryFilterer{contract: contract}, nil
}

// bindEntry binds a generic wrapper to an already deployed contract.
func bindEntry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EntryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Entry *EntryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Entry.Contract.EntryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Entry *EntryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Entry.Contract.EntryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Entry *EntryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Entry.Contract.EntryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Entry *EntryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Entry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Entry *EntryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Entry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Entry *EntryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Entry.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Entry *EntryCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Entry.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Entry *EntrySession) Admin() (common.Address, error) {
	return _Entry.Contract.Admin(&_Entry.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Entry *EntryCallerSession) Admin() (common.Address, error) {
	return _Entry.Contract.Admin(&_Entry.CallOpts)
}

// ChainID is a free data retrieval call binding the contract method 0xadc879e9.
//
// Solidity: function chainID() view returns(uint8)
func (_Entry *EntryCaller) ChainID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Entry.contract.Call(opts, &out, "chainID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ChainID is a free data retrieval call binding the contract method 0xadc879e9.
//
// Solidity: function chainID() view returns(uint8)
func (_Entry *EntrySession) ChainID() (uint8, error) {
	return _Entry.Contract.ChainID(&_Entry.CallOpts)
}

// ChainID is a free data retrieval call binding the contract method 0xadc879e9.
//
// Solidity: function chainID() view returns(uint8)
func (_Entry *EntryCallerSession) ChainID() (uint8, error) {
	return _Entry.Contract.ChainID(&_Entry.CallOpts)
}

// CheckUserKey is a free data retrieval call binding the contract method 0xa82285ac.
//
// Solidity: function checkUserKey(bytes32 registerEmail, uint256 keyType, bytes key) view returns(bool, uint256)
func (_Entry *EntryCaller) CheckUserKey(opts *bind.CallOpts, registerEmail [32]byte, keyType *big.Int, key []byte) (bool, *big.Int, error) {
	var out []interface{}
	err := _Entry.contract.Call(opts, &out, "checkUserKey", registerEmail, keyType, key)

	if err != nil {
		return *new(bool), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// CheckUserKey is a free data retrieval call binding the contract method 0xa82285ac.
//
// Solidity: function checkUserKey(bytes32 registerEmail, uint256 keyType, bytes key) view returns(bool, uint256)
func (_Entry *EntrySession) CheckUserKey(registerEmail [32]byte, keyType *big.Int, key []byte) (bool, *big.Int, error) {
	return _Entry.Contract.CheckUserKey(&_Entry.CallOpts, registerEmail, keyType, key)
}

// CheckUserKey is a free data retrieval call binding the contract method 0xa82285ac.
//
// Solidity: function checkUserKey(bytes32 registerEmail, uint256 keyType, bytes key) view returns(bool, uint256)
func (_Entry *EntryCallerSession) CheckUserKey(registerEmail [32]byte, keyType *big.Int, key []byte) (bool, *big.Int, error) {
	return _Entry.Contract.CheckUserKey(&_Entry.CallOpts, registerEmail, keyType, key)
}

// GetUserKeys is a free data retrieval call binding the contract method 0x64d9dcc1.
//
// Solidity: function getUserKeys(bytes32 registerEmail) view returns((uint256,bytes)[] keys)
func (_Entry *EntryCaller) GetUserKeys(opts *bind.CallOpts, registerEmail [32]byte) ([]UnipassUtilsPubKey, error) {
	var out []interface{}
	err := _Entry.contract.Call(opts, &out, "getUserKeys", registerEmail)

	if err != nil {
		return *new([]UnipassUtilsPubKey), err
	}

	out0 := *abi.ConvertType(out[0], new([]UnipassUtilsPubKey)).(*[]UnipassUtilsPubKey)

	return out0, err

}

// GetUserKeys is a free data retrieval call binding the contract method 0x64d9dcc1.
//
// Solidity: function getUserKeys(bytes32 registerEmail) view returns((uint256,bytes)[] keys)
func (_Entry *EntrySession) GetUserKeys(registerEmail [32]byte) ([]UnipassUtilsPubKey, error) {
	return _Entry.Contract.GetUserKeys(&_Entry.CallOpts, registerEmail)
}

// GetUserKeys is a free data retrieval call binding the contract method 0x64d9dcc1.
//
// Solidity: function getUserKeys(bytes32 registerEmail) view returns((uint256,bytes)[] keys)
func (_Entry *EntryCallerSession) GetUserKeys(registerEmail [32]byte) ([]UnipassUtilsPubKey, error) {
	return _Entry.Contract.GetUserKeys(&_Entry.CallOpts, registerEmail)
}

// GetUserName is a free data retrieval call binding the contract method 0x21a11805.
//
// Solidity: function getUserName(bytes32 registerEmail) view returns(string)
func (_Entry *EntryCaller) GetUserName(opts *bind.CallOpts, registerEmail [32]byte) (string, error) {
	var out []interface{}
	err := _Entry.contract.Call(opts, &out, "getUserName", registerEmail)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetUserName is a free data retrieval call binding the contract method 0x21a11805.
//
// Solidity: function getUserName(bytes32 registerEmail) view returns(string)
func (_Entry *EntrySession) GetUserName(registerEmail [32]byte) (string, error) {
	return _Entry.Contract.GetUserName(&_Entry.CallOpts, registerEmail)
}

// GetUserName is a free data retrieval call binding the contract method 0x21a11805.
//
// Solidity: function getUserName(bytes32 registerEmail) view returns(string)
func (_Entry *EntryCallerSession) GetUserName(registerEmail [32]byte) (string, error) {
	return _Entry.Contract.GetUserName(&_Entry.CallOpts, registerEmail)
}

// GetUserNonce is a free data retrieval call binding the contract method 0xb8601cf3.
//
// Solidity: function getUserNonce(bytes32 registerEmail) view returns(uint256)
func (_Entry *EntryCaller) GetUserNonce(opts *bind.CallOpts, registerEmail [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Entry.contract.Call(opts, &out, "getUserNonce", registerEmail)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserNonce is a free data retrieval call binding the contract method 0xb8601cf3.
//
// Solidity: function getUserNonce(bytes32 registerEmail) view returns(uint256)
func (_Entry *EntrySession) GetUserNonce(registerEmail [32]byte) (*big.Int, error) {
	return _Entry.Contract.GetUserNonce(&_Entry.CallOpts, registerEmail)
}

// GetUserNonce is a free data retrieval call binding the contract method 0xb8601cf3.
//
// Solidity: function getUserNonce(bytes32 registerEmail) view returns(uint256)
func (_Entry *EntryCallerSession) GetUserNonce(registerEmail [32]byte) (*big.Int, error) {
	return _Entry.Contract.GetUserNonce(&_Entry.CallOpts, registerEmail)
}

// GetUserPendingState is a free data retrieval call binding the contract method 0x02f89f51.
//
// Solidity: function getUserPendingState(bytes32 registerEmail) view returns(((uint256,bytes),uint256,bool))
func (_Entry *EntryCaller) GetUserPendingState(opts *bind.CallOpts, registerEmail [32]byte) (EntryPendingState, error) {
	var out []interface{}
	err := _Entry.contract.Call(opts, &out, "getUserPendingState", registerEmail)

	if err != nil {
		return *new(EntryPendingState), err
	}

	out0 := *abi.ConvertType(out[0], new(EntryPendingState)).(*EntryPendingState)

	return out0, err

}

// GetUserPendingState is a free data retrieval call binding the contract method 0x02f89f51.
//
// Solidity: function getUserPendingState(bytes32 registerEmail) view returns(((uint256,bytes),uint256,bool))
func (_Entry *EntrySession) GetUserPendingState(registerEmail [32]byte) (EntryPendingState, error) {
	return _Entry.Contract.GetUserPendingState(&_Entry.CallOpts, registerEmail)
}

// GetUserPendingState is a free data retrieval call binding the contract method 0x02f89f51.
//
// Solidity: function getUserPendingState(bytes32 registerEmail) view returns(((uint256,bytes),uint256,bool))
func (_Entry *EntryCallerSession) GetUserPendingState(registerEmail [32]byte) (EntryPendingState, error) {
	return _Entry.Contract.GetUserPendingState(&_Entry.CallOpts, registerEmail)
}

// IsUseEmailRegister is a free data retrieval call binding the contract method 0x88bf26e7.
//
// Solidity: function isUseEmailRegister(bytes32 registerEmail) view returns(bool)
func (_Entry *EntryCaller) IsUseEmailRegister(opts *bind.CallOpts, registerEmail [32]byte) (bool, error) {
	var out []interface{}
	err := _Entry.contract.Call(opts, &out, "isUseEmailRegister", registerEmail)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsUseEmailRegister is a free data retrieval call binding the contract method 0x88bf26e7.
//
// Solidity: function isUseEmailRegister(bytes32 registerEmail) view returns(bool)
func (_Entry *EntrySession) IsUseEmailRegister(registerEmail [32]byte) (bool, error) {
	return _Entry.Contract.IsUseEmailRegister(&_Entry.CallOpts, registerEmail)
}

// IsUseEmailRegister is a free data retrieval call binding the contract method 0x88bf26e7.
//
// Solidity: function isUseEmailRegister(bytes32 registerEmail) view returns(bool)
func (_Entry *EntryCallerSession) IsUseEmailRegister(registerEmail [32]byte) (bool, error) {
	return _Entry.Contract.IsUseEmailRegister(&_Entry.CallOpts, registerEmail)
}

// IsUserKey is a free data retrieval call binding the contract method 0x7231b188.
//
// Solidity: function isUserKey(bytes32 registerEmail, uint256 sigKeyType, bytes sigKey) view returns(bool)
func (_Entry *EntryCaller) IsUserKey(opts *bind.CallOpts, registerEmail [32]byte, sigKeyType *big.Int, sigKey []byte) (bool, error) {
	var out []interface{}
	err := _Entry.contract.Call(opts, &out, "isUserKey", registerEmail, sigKeyType, sigKey)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsUserKey is a free data retrieval call binding the contract method 0x7231b188.
//
// Solidity: function isUserKey(bytes32 registerEmail, uint256 sigKeyType, bytes sigKey) view returns(bool)
func (_Entry *EntrySession) IsUserKey(registerEmail [32]byte, sigKeyType *big.Int, sigKey []byte) (bool, error) {
	return _Entry.Contract.IsUserKey(&_Entry.CallOpts, registerEmail, sigKeyType, sigKey)
}

// IsUserKey is a free data retrieval call binding the contract method 0x7231b188.
//
// Solidity: function isUserKey(bytes32 registerEmail, uint256 sigKeyType, bytes sigKey) view returns(bool)
func (_Entry *EntryCallerSession) IsUserKey(registerEmail [32]byte, sigKeyType *big.Int, sigKey []byte) (bool, error) {
	return _Entry.Contract.IsUserKey(&_Entry.CallOpts, registerEmail, sigKeyType, sigKey)
}

// Logic is a free data retrieval call binding the contract method 0xd7dfa0dd.
//
// Solidity: function logic() view returns(address)
func (_Entry *EntryCaller) Logic(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Entry.contract.Call(opts, &out, "logic")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Logic is a free data retrieval call binding the contract method 0xd7dfa0dd.
//
// Solidity: function logic() view returns(address)
func (_Entry *EntrySession) Logic() (common.Address, error) {
	return _Entry.Contract.Logic(&_Entry.CallOpts)
}

// Logic is a free data retrieval call binding the contract method 0xd7dfa0dd.
//
// Solidity: function logic() view returns(address)
func (_Entry *EntryCallerSession) Logic() (common.Address, error) {
	return _Entry.Contract.Logic(&_Entry.CallOpts)
}

// TimeLockSwitch is a free data retrieval call binding the contract method 0xb7da1992.
//
// Solidity: function timeLockSwitch() view returns(bool)
func (_Entry *EntryCaller) TimeLockSwitch(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Entry.contract.Call(opts, &out, "timeLockSwitch")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TimeLockSwitch is a free data retrieval call binding the contract method 0xb7da1992.
//
// Solidity: function timeLockSwitch() view returns(bool)
func (_Entry *EntrySession) TimeLockSwitch() (bool, error) {
	return _Entry.Contract.TimeLockSwitch(&_Entry.CallOpts)
}

// TimeLockSwitch is a free data retrieval call binding the contract method 0xb7da1992.
//
// Solidity: function timeLockSwitch() view returns(bool)
func (_Entry *EntryCallerSession) TimeLockSwitch() (bool, error) {
	return _Entry.Contract.TimeLockSwitch(&_Entry.CallOpts)
}

// TotalUsers is a free data retrieval call binding the contract method 0xbff1f9e1.
//
// Solidity: function totalUsers() view returns(uint256)
func (_Entry *EntryCaller) TotalUsers(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Entry.contract.Call(opts, &out, "totalUsers")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalUsers is a free data retrieval call binding the contract method 0xbff1f9e1.
//
// Solidity: function totalUsers() view returns(uint256)
func (_Entry *EntrySession) TotalUsers() (*big.Int, error) {
	return _Entry.Contract.TotalUsers(&_Entry.CallOpts)
}

// TotalUsers is a free data retrieval call binding the contract method 0xbff1f9e1.
//
// Solidity: function totalUsers() view returns(uint256)
func (_Entry *EntryCallerSession) TotalUsers() (*big.Int, error) {
	return _Entry.Contract.TotalUsers(&_Entry.CallOpts)
}

// Users is a free data retrieval call binding the contract method 0xcea6ab98.
//
// Solidity: function users(bytes32 ) view returns(bytes32 registerEmail, string originUsername, bool isUseEmailRegister, uint256 nonce, bool quickLogin, (uint256,bytes32[]) recoveryEmails, address assetContractAddress, bool isRecovering, ((uint256,bytes),uint256,bool) pendingState, bytes source)
func (_Entry *EntryCaller) Users(opts *bind.CallOpts, arg0 [32]byte) (struct {
	RegisterEmail        [32]byte
	OriginUsername       string
	IsUseEmailRegister   bool
	Nonce                *big.Int
	QuickLogin           bool
	RecoveryEmails       EntryRecoveryEmail
	AssetContractAddress common.Address
	IsRecovering         bool
	PendingState         EntryPendingState
	Source               []byte
}, error) {
	var out []interface{}
	err := _Entry.contract.Call(opts, &out, "users", arg0)

	outstruct := new(struct {
		RegisterEmail        [32]byte
		OriginUsername       string
		IsUseEmailRegister   bool
		Nonce                *big.Int
		QuickLogin           bool
		RecoveryEmails       EntryRecoveryEmail
		AssetContractAddress common.Address
		IsRecovering         bool
		PendingState         EntryPendingState
		Source               []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RegisterEmail = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.OriginUsername = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.IsUseEmailRegister = *abi.ConvertType(out[2], new(bool)).(*bool)
	outstruct.Nonce = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.QuickLogin = *abi.ConvertType(out[4], new(bool)).(*bool)
	outstruct.RecoveryEmails = *abi.ConvertType(out[5], new(EntryRecoveryEmail)).(*EntryRecoveryEmail)
	outstruct.AssetContractAddress = *abi.ConvertType(out[6], new(common.Address)).(*common.Address)
	outstruct.IsRecovering = *abi.ConvertType(out[7], new(bool)).(*bool)
	outstruct.PendingState = *abi.ConvertType(out[8], new(EntryPendingState)).(*EntryPendingState)
	outstruct.Source = *abi.ConvertType(out[9], new([]byte)).(*[]byte)

	return *outstruct, err

}

// Users is a free data retrieval call binding the contract method 0xcea6ab98.
//
// Solidity: function users(bytes32 ) view returns(bytes32 registerEmail, string originUsername, bool isUseEmailRegister, uint256 nonce, bool quickLogin, (uint256,bytes32[]) recoveryEmails, address assetContractAddress, bool isRecovering, ((uint256,bytes),uint256,bool) pendingState, bytes source)
func (_Entry *EntrySession) Users(arg0 [32]byte) (struct {
	RegisterEmail        [32]byte
	OriginUsername       string
	IsUseEmailRegister   bool
	Nonce                *big.Int
	QuickLogin           bool
	RecoveryEmails       EntryRecoveryEmail
	AssetContractAddress common.Address
	IsRecovering         bool
	PendingState         EntryPendingState
	Source               []byte
}, error) {
	return _Entry.Contract.Users(&_Entry.CallOpts, arg0)
}

// Users is a free data retrieval call binding the contract method 0xcea6ab98.
//
// Solidity: function users(bytes32 ) view returns(bytes32 registerEmail, string originUsername, bool isUseEmailRegister, uint256 nonce, bool quickLogin, (uint256,bytes32[]) recoveryEmails, address assetContractAddress, bool isRecovering, ((uint256,bytes),uint256,bool) pendingState, bytes source)
func (_Entry *EntryCallerSession) Users(arg0 [32]byte) (struct {
	RegisterEmail        [32]byte
	OriginUsername       string
	IsUseEmailRegister   bool
	Nonce                *big.Int
	QuickLogin           bool
	RecoveryEmails       EntryRecoveryEmail
	AssetContractAddress common.Address
	IsRecovering         bool
	PendingState         EntryPendingState
	Source               []byte
}, error) {
	return _Entry.Contract.Users(&_Entry.CallOpts, arg0)
}

// Verify is a free data retrieval call binding the contract method 0x8cfdd6f3.
//
// Solidity: function verify(bytes32 registerEmail, bytes32 hash, bytes sig, uint256 sigKeyIndex) view returns(bool)
func (_Entry *EntryCaller) Verify(opts *bind.CallOpts, registerEmail [32]byte, hash [32]byte, sig []byte, sigKeyIndex *big.Int) (bool, error) {
	var out []interface{}
	err := _Entry.contract.Call(opts, &out, "verify", registerEmail, hash, sig, sigKeyIndex)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Verify is a free data retrieval call binding the contract method 0x8cfdd6f3.
//
// Solidity: function verify(bytes32 registerEmail, bytes32 hash, bytes sig, uint256 sigKeyIndex) view returns(bool)
func (_Entry *EntrySession) Verify(registerEmail [32]byte, hash [32]byte, sig []byte, sigKeyIndex *big.Int) (bool, error) {
	return _Entry.Contract.Verify(&_Entry.CallOpts, registerEmail, hash, sig, sigKeyIndex)
}

// Verify is a free data retrieval call binding the contract method 0x8cfdd6f3.
//
// Solidity: function verify(bytes32 registerEmail, bytes32 hash, bytes sig, uint256 sigKeyIndex) view returns(bool)
func (_Entry *EntryCallerSession) Verify(registerEmail [32]byte, hash [32]byte, sig []byte, sigKeyIndex *big.Int) (bool, error) {
	return _Entry.Contract.Verify(&_Entry.CallOpts, registerEmail, hash, sig, sigKeyIndex)
}

// VerifyUserSig is a free data retrieval call binding the contract method 0x878be590.
//
// Solidity: function verifyUserSig(bytes32 registerEmail, bytes message, uint256 sigKeyType, bytes sigKey, bytes sig) view returns(bool)
func (_Entry *EntryCaller) VerifyUserSig(opts *bind.CallOpts, registerEmail [32]byte, message []byte, sigKeyType *big.Int, sigKey []byte, sig []byte) (bool, error) {
	var out []interface{}
	err := _Entry.contract.Call(opts, &out, "verifyUserSig", registerEmail, message, sigKeyType, sigKey, sig)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyUserSig is a free data retrieval call binding the contract method 0x878be590.
//
// Solidity: function verifyUserSig(bytes32 registerEmail, bytes message, uint256 sigKeyType, bytes sigKey, bytes sig) view returns(bool)
func (_Entry *EntrySession) VerifyUserSig(registerEmail [32]byte, message []byte, sigKeyType *big.Int, sigKey []byte, sig []byte) (bool, error) {
	return _Entry.Contract.VerifyUserSig(&_Entry.CallOpts, registerEmail, message, sigKeyType, sigKey, sig)
}

// VerifyUserSig is a free data retrieval call binding the contract method 0x878be590.
//
// Solidity: function verifyUserSig(bytes32 registerEmail, bytes message, uint256 sigKeyType, bytes sigKey, bytes sig) view returns(bool)
func (_Entry *EntryCallerSession) VerifyUserSig(registerEmail [32]byte, message []byte, sigKeyType *big.Int, sigKey []byte, sig []byte) (bool, error) {
	return _Entry.Contract.VerifyUserSig(&_Entry.CallOpts, registerEmail, message, sigKeyType, sigKey, sig)
}

// AddLocalKey is a paid mutator transaction binding the contract method 0xd14f8e8a.
//
// Solidity: function addLocalKey(bytes32 registerEmail, bytes32 username, uint256 nonce, uint256 newKeyType, bytes newKey, bytes newKeySig, bytes sig, uint256 sigKeyIndex) returns()
func (_Entry *EntryTransactor) AddLocalKey(opts *bind.TransactOpts, registerEmail [32]byte, username [32]byte, nonce *big.Int, newKeyType *big.Int, newKey []byte, newKeySig []byte, sig []byte, sigKeyIndex *big.Int) (*types.Transaction, error) {
	return _Entry.contract.Transact(opts, "addLocalKey", registerEmail, username, nonce, newKeyType, newKey, newKeySig, sig, sigKeyIndex)
}

// AddLocalKey is a paid mutator transaction binding the contract method 0xd14f8e8a.
//
// Solidity: function addLocalKey(bytes32 registerEmail, bytes32 username, uint256 nonce, uint256 newKeyType, bytes newKey, bytes newKeySig, bytes sig, uint256 sigKeyIndex) returns()
func (_Entry *EntrySession) AddLocalKey(registerEmail [32]byte, username [32]byte, nonce *big.Int, newKeyType *big.Int, newKey []byte, newKeySig []byte, sig []byte, sigKeyIndex *big.Int) (*types.Transaction, error) {
	return _Entry.Contract.AddLocalKey(&_Entry.TransactOpts, registerEmail, username, nonce, newKeyType, newKey, newKeySig, sig, sigKeyIndex)
}

// AddLocalKey is a paid mutator transaction binding the contract method 0xd14f8e8a.
//
// Solidity: function addLocalKey(bytes32 registerEmail, bytes32 username, uint256 nonce, uint256 newKeyType, bytes newKey, bytes newKeySig, bytes sig, uint256 sigKeyIndex) returns()
func (_Entry *EntryTransactorSession) AddLocalKey(registerEmail [32]byte, username [32]byte, nonce *big.Int, newKeyType *big.Int, newKey []byte, newKeySig []byte, sig []byte, sigKeyIndex *big.Int) (*types.Transaction, error) {
	return _Entry.Contract.AddLocalKey(&_Entry.TransactOpts, registerEmail, username, nonce, newKeyType, newKey, newKeySig, sig, sigKeyIndex)
}

// CancelRecovery is a paid mutator transaction binding the contract method 0x5b67034d.
//
// Solidity: function cancelRecovery(bytes32 registerEmail, bytes32 username, uint256 nonce, bytes sig, uint256 sigKeyIndex) returns()
func (_Entry *EntryTransactor) CancelRecovery(opts *bind.TransactOpts, registerEmail [32]byte, username [32]byte, nonce *big.Int, sig []byte, sigKeyIndex *big.Int) (*types.Transaction, error) {
	return _Entry.contract.Transact(opts, "cancelRecovery", registerEmail, username, nonce, sig, sigKeyIndex)
}

// CancelRecovery is a paid mutator transaction binding the contract method 0x5b67034d.
//
// Solidity: function cancelRecovery(bytes32 registerEmail, bytes32 username, uint256 nonce, bytes sig, uint256 sigKeyIndex) returns()
func (_Entry *EntrySession) CancelRecovery(registerEmail [32]byte, username [32]byte, nonce *big.Int, sig []byte, sigKeyIndex *big.Int) (*types.Transaction, error) {
	return _Entry.Contract.CancelRecovery(&_Entry.TransactOpts, registerEmail, username, nonce, sig, sigKeyIndex)
}

// CancelRecovery is a paid mutator transaction binding the contract method 0x5b67034d.
//
// Solidity: function cancelRecovery(bytes32 registerEmail, bytes32 username, uint256 nonce, bytes sig, uint256 sigKeyIndex) returns()
func (_Entry *EntryTransactorSession) CancelRecovery(registerEmail [32]byte, username [32]byte, nonce *big.Int, sig []byte, sigKeyIndex *big.Int) (*types.Transaction, error) {
	return _Entry.Contract.CancelRecovery(&_Entry.TransactOpts, registerEmail, username, nonce, sig, sigKeyIndex)
}

// CompleteRecovery is a paid mutator transaction binding the contract method 0x1272c596.
//
// Solidity: function completeRecovery(bytes32 registerEmail) returns()
func (_Entry *EntryTransactor) CompleteRecovery(opts *bind.TransactOpts, registerEmail [32]byte) (*types.Transaction, error) {
	return _Entry.contract.Transact(opts, "completeRecovery", registerEmail)
}

// CompleteRecovery is a paid mutator transaction binding the contract method 0x1272c596.
//
// Solidity: function completeRecovery(bytes32 registerEmail) returns()
func (_Entry *EntrySession) CompleteRecovery(registerEmail [32]byte) (*types.Transaction, error) {
	return _Entry.Contract.CompleteRecovery(&_Entry.TransactOpts, registerEmail)
}

// CompleteRecovery is a paid mutator transaction binding the contract method 0x1272c596.
//
// Solidity: function completeRecovery(bytes32 registerEmail) returns()
func (_Entry *EntryTransactorSession) CompleteRecovery(registerEmail [32]byte) (*types.Transaction, error) {
	return _Entry.Contract.CompleteRecovery(&_Entry.TransactOpts, registerEmail)
}

// DelLocalKey is a paid mutator transaction binding the contract method 0xbaf508fe.
//
// Solidity: function delLocalKey(bytes32 registerEmail, bytes32 username, uint256 nonce, uint256 delKeyIndex, bytes sig, uint256 sigKeyIndex) returns()
func (_Entry *EntryTransactor) DelLocalKey(opts *bind.TransactOpts, registerEmail [32]byte, username [32]byte, nonce *big.Int, delKeyIndex *big.Int, sig []byte, sigKeyIndex *big.Int) (*types.Transaction, error) {
	return _Entry.contract.Transact(opts, "delLocalKey", registerEmail, username, nonce, delKeyIndex, sig, sigKeyIndex)
}

// DelLocalKey is a paid mutator transaction binding the contract method 0xbaf508fe.
//
// Solidity: function delLocalKey(bytes32 registerEmail, bytes32 username, uint256 nonce, uint256 delKeyIndex, bytes sig, uint256 sigKeyIndex) returns()
func (_Entry *EntrySession) DelLocalKey(registerEmail [32]byte, username [32]byte, nonce *big.Int, delKeyIndex *big.Int, sig []byte, sigKeyIndex *big.Int) (*types.Transaction, error) {
	return _Entry.Contract.DelLocalKey(&_Entry.TransactOpts, registerEmail, username, nonce, delKeyIndex, sig, sigKeyIndex)
}

// DelLocalKey is a paid mutator transaction binding the contract method 0xbaf508fe.
//
// Solidity: function delLocalKey(bytes32 registerEmail, bytes32 username, uint256 nonce, uint256 delKeyIndex, bytes sig, uint256 sigKeyIndex) returns()
func (_Entry *EntryTransactorSession) DelLocalKey(registerEmail [32]byte, username [32]byte, nonce *big.Int, delKeyIndex *big.Int, sig []byte, sigKeyIndex *big.Int) (*types.Transaction, error) {
	return _Entry.Contract.DelLocalKey(&_Entry.TransactOpts, registerEmail, username, nonce, delKeyIndex, sig, sigKeyIndex)
}

// Init is a paid mutator transaction binding the contract method 0xf8ebd43b.
//
// Solidity: function init(address _logic, address _admin, uint8 _chainID) returns()
func (_Entry *EntryTransactor) Init(opts *bind.TransactOpts, _logic common.Address, _admin common.Address, _chainID uint8) (*types.Transaction, error) {
	return _Entry.contract.Transact(opts, "init", _logic, _admin, _chainID)
}

// Init is a paid mutator transaction binding the contract method 0xf8ebd43b.
//
// Solidity: function init(address _logic, address _admin, uint8 _chainID) returns()
func (_Entry *EntrySession) Init(_logic common.Address, _admin common.Address, _chainID uint8) (*types.Transaction, error) {
	return _Entry.Contract.Init(&_Entry.TransactOpts, _logic, _admin, _chainID)
}

// Init is a paid mutator transaction binding the contract method 0xf8ebd43b.
//
// Solidity: function init(address _logic, address _admin, uint8 _chainID) returns()
func (_Entry *EntryTransactorSession) Init(_logic common.Address, _admin common.Address, _chainID uint8) (*types.Transaction, error) {
	return _Entry.Contract.Init(&_Entry.TransactOpts, _logic, _admin, _chainID)
}

// QuickAddLocalKey is a paid mutator transaction binding the contract method 0xcf133f08.
//
// Solidity: function quickAddLocalKey(bytes32 registerEmail, bytes32 username, uint256 nonce, uint256 newKeyType, bytes newKey, bytes newKeySig, bytes[] emailHeaders) returns()
func (_Entry *EntryTransactor) QuickAddLocalKey(opts *bind.TransactOpts, registerEmail [32]byte, username [32]byte, nonce *big.Int, newKeyType *big.Int, newKey []byte, newKeySig []byte, emailHeaders [][]byte) (*types.Transaction, error) {
	return _Entry.contract.Transact(opts, "quickAddLocalKey", registerEmail, username, nonce, newKeyType, newKey, newKeySig, emailHeaders)
}

// QuickAddLocalKey is a paid mutator transaction binding the contract method 0xcf133f08.
//
// Solidity: function quickAddLocalKey(bytes32 registerEmail, bytes32 username, uint256 nonce, uint256 newKeyType, bytes newKey, bytes newKeySig, bytes[] emailHeaders) returns()
func (_Entry *EntrySession) QuickAddLocalKey(registerEmail [32]byte, username [32]byte, nonce *big.Int, newKeyType *big.Int, newKey []byte, newKeySig []byte, emailHeaders [][]byte) (*types.Transaction, error) {
	return _Entry.Contract.QuickAddLocalKey(&_Entry.TransactOpts, registerEmail, username, nonce, newKeyType, newKey, newKeySig, emailHeaders)
}

// QuickAddLocalKey is a paid mutator transaction binding the contract method 0xcf133f08.
//
// Solidity: function quickAddLocalKey(bytes32 registerEmail, bytes32 username, uint256 nonce, uint256 newKeyType, bytes newKey, bytes newKeySig, bytes[] emailHeaders) returns()
func (_Entry *EntryTransactorSession) QuickAddLocalKey(registerEmail [32]byte, username [32]byte, nonce *big.Int, newKeyType *big.Int, newKey []byte, newKeySig []byte, emailHeaders [][]byte) (*types.Transaction, error) {
	return _Entry.Contract.QuickAddLocalKey(&_Entry.TransactOpts, registerEmail, username, nonce, newKeyType, newKey, newKeySig, emailHeaders)
}

// QuickRegister is a paid mutator transaction binding the contract method 0xf0a08ba9.
//
// Solidity: function quickRegister(bytes32 registerEmail, string originUsername, uint256 keyType, bytes key, bytes sig, string source) returns()
func (_Entry *EntryTransactor) QuickRegister(opts *bind.TransactOpts, registerEmail [32]byte, originUsername string, keyType *big.Int, key []byte, sig []byte, source string) (*types.Transaction, error) {
	return _Entry.contract.Transact(opts, "quickRegister", registerEmail, originUsername, keyType, key, sig, source)
}

// QuickRegister is a paid mutator transaction binding the contract method 0xf0a08ba9.
//
// Solidity: function quickRegister(bytes32 registerEmail, string originUsername, uint256 keyType, bytes key, bytes sig, string source) returns()
func (_Entry *EntrySession) QuickRegister(registerEmail [32]byte, originUsername string, keyType *big.Int, key []byte, sig []byte, source string) (*types.Transaction, error) {
	return _Entry.Contract.QuickRegister(&_Entry.TransactOpts, registerEmail, originUsername, keyType, key, sig, source)
}

// QuickRegister is a paid mutator transaction binding the contract method 0xf0a08ba9.
//
// Solidity: function quickRegister(bytes32 registerEmail, string originUsername, uint256 keyType, bytes key, bytes sig, string source) returns()
func (_Entry *EntryTransactorSession) QuickRegister(registerEmail [32]byte, originUsername string, keyType *big.Int, key []byte, sig []byte, source string) (*types.Transaction, error) {
	return _Entry.Contract.QuickRegister(&_Entry.TransactOpts, registerEmail, originUsername, keyType, key, sig, source)
}

// Register is a paid mutator transaction binding the contract method 0xb865d7bf.
//
// Solidity: function register(bytes32 registerEmail, string originUsername, bytes emailHeader, uint256 keyType, bytes key, bytes sig, string source) returns()
func (_Entry *EntryTransactor) Register(opts *bind.TransactOpts, registerEmail [32]byte, originUsername string, emailHeader []byte, keyType *big.Int, key []byte, sig []byte, source string) (*types.Transaction, error) {
	return _Entry.contract.Transact(opts, "register", registerEmail, originUsername, emailHeader, keyType, key, sig, source)
}

// Register is a paid mutator transaction binding the contract method 0xb865d7bf.
//
// Solidity: function register(bytes32 registerEmail, string originUsername, bytes emailHeader, uint256 keyType, bytes key, bytes sig, string source) returns()
func (_Entry *EntrySession) Register(registerEmail [32]byte, originUsername string, emailHeader []byte, keyType *big.Int, key []byte, sig []byte, source string) (*types.Transaction, error) {
	return _Entry.Contract.Register(&_Entry.TransactOpts, registerEmail, originUsername, emailHeader, keyType, key, sig, source)
}

// Register is a paid mutator transaction binding the contract method 0xb865d7bf.
//
// Solidity: function register(bytes32 registerEmail, string originUsername, bytes emailHeader, uint256 keyType, bytes key, bytes sig, string source) returns()
func (_Entry *EntryTransactorSession) Register(registerEmail [32]byte, originUsername string, emailHeader []byte, keyType *big.Int, key []byte, sig []byte, source string) (*types.Transaction, error) {
	return _Entry.Contract.Register(&_Entry.TransactOpts, registerEmail, originUsername, emailHeader, keyType, key, sig, source)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address _newAdmin) returns()
func (_Entry *EntryTransactor) SetAdmin(opts *bind.TransactOpts, _newAdmin common.Address) (*types.Transaction, error) {
	return _Entry.contract.Transact(opts, "setAdmin", _newAdmin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address _newAdmin) returns()
func (_Entry *EntrySession) SetAdmin(_newAdmin common.Address) (*types.Transaction, error) {
	return _Entry.Contract.SetAdmin(&_Entry.TransactOpts, _newAdmin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address _newAdmin) returns()
func (_Entry *EntryTransactorSession) SetAdmin(_newAdmin common.Address) (*types.Transaction, error) {
	return _Entry.Contract.SetAdmin(&_Entry.TransactOpts, _newAdmin)
}

// SetTimeLockSwitch is a paid mutator transaction binding the contract method 0x5ee24823.
//
// Solidity: function setTimeLockSwitch(bool _switch) returns()
func (_Entry *EntryTransactor) SetTimeLockSwitch(opts *bind.TransactOpts, _switch bool) (*types.Transaction, error) {
	return _Entry.contract.Transact(opts, "setTimeLockSwitch", _switch)
}

// SetTimeLockSwitch is a paid mutator transaction binding the contract method 0x5ee24823.
//
// Solidity: function setTimeLockSwitch(bool _switch) returns()
func (_Entry *EntrySession) SetTimeLockSwitch(_switch bool) (*types.Transaction, error) {
	return _Entry.Contract.SetTimeLockSwitch(&_Entry.TransactOpts, _switch)
}

// SetTimeLockSwitch is a paid mutator transaction binding the contract method 0x5ee24823.
//
// Solidity: function setTimeLockSwitch(bool _switch) returns()
func (_Entry *EntryTransactorSession) SetTimeLockSwitch(_switch bool) (*types.Transaction, error) {
	return _Entry.Contract.SetTimeLockSwitch(&_Entry.TransactOpts, _switch)
}

// StartRecovery is a paid mutator transaction binding the contract method 0xb2744923.
//
// Solidity: function startRecovery(bytes32 registerEmail, bytes32 username, uint256 nonce, bool resetKeys, uint256 newKeyType, bytes newKey, bytes newKeySig, bytes[] emailHeaders) returns()
func (_Entry *EntryTransactor) StartRecovery(opts *bind.TransactOpts, registerEmail [32]byte, username [32]byte, nonce *big.Int, resetKeys bool, newKeyType *big.Int, newKey []byte, newKeySig []byte, emailHeaders [][]byte) (*types.Transaction, error) {
	return _Entry.contract.Transact(opts, "startRecovery", registerEmail, username, nonce, resetKeys, newKeyType, newKey, newKeySig, emailHeaders)
}

// StartRecovery is a paid mutator transaction binding the contract method 0xb2744923.
//
// Solidity: function startRecovery(bytes32 registerEmail, bytes32 username, uint256 nonce, bool resetKeys, uint256 newKeyType, bytes newKey, bytes newKeySig, bytes[] emailHeaders) returns()
func (_Entry *EntrySession) StartRecovery(registerEmail [32]byte, username [32]byte, nonce *big.Int, resetKeys bool, newKeyType *big.Int, newKey []byte, newKeySig []byte, emailHeaders [][]byte) (*types.Transaction, error) {
	return _Entry.Contract.StartRecovery(&_Entry.TransactOpts, registerEmail, username, nonce, resetKeys, newKeyType, newKey, newKeySig, emailHeaders)
}

// StartRecovery is a paid mutator transaction binding the contract method 0xb2744923.
//
// Solidity: function startRecovery(bytes32 registerEmail, bytes32 username, uint256 nonce, bool resetKeys, uint256 newKeyType, bytes newKey, bytes newKeySig, bytes[] emailHeaders) returns()
func (_Entry *EntryTransactorSession) StartRecovery(registerEmail [32]byte, username [32]byte, nonce *big.Int, resetKeys bool, newKeyType *big.Int, newKey []byte, newKeySig []byte, emailHeaders [][]byte) (*types.Transaction, error) {
	return _Entry.Contract.StartRecovery(&_Entry.TransactOpts, registerEmail, username, nonce, resetKeys, newKeyType, newKey, newKeySig, emailHeaders)
}

// UpdateQuickLogin is a paid mutator transaction binding the contract method 0x38a1cfcd.
//
// Solidity: function updateQuickLogin(bytes32 registerEmail, bytes32 username, uint256 nonce, bool quickLogin, bytes sig, uint256 sigKeyIndex) returns()
func (_Entry *EntryTransactor) UpdateQuickLogin(opts *bind.TransactOpts, registerEmail [32]byte, username [32]byte, nonce *big.Int, quickLogin bool, sig []byte, sigKeyIndex *big.Int) (*types.Transaction, error) {
	return _Entry.contract.Transact(opts, "updateQuickLogin", registerEmail, username, nonce, quickLogin, sig, sigKeyIndex)
}

// UpdateQuickLogin is a paid mutator transaction binding the contract method 0x38a1cfcd.
//
// Solidity: function updateQuickLogin(bytes32 registerEmail, bytes32 username, uint256 nonce, bool quickLogin, bytes sig, uint256 sigKeyIndex) returns()
func (_Entry *EntrySession) UpdateQuickLogin(registerEmail [32]byte, username [32]byte, nonce *big.Int, quickLogin bool, sig []byte, sigKeyIndex *big.Int) (*types.Transaction, error) {
	return _Entry.Contract.UpdateQuickLogin(&_Entry.TransactOpts, registerEmail, username, nonce, quickLogin, sig, sigKeyIndex)
}

// UpdateQuickLogin is a paid mutator transaction binding the contract method 0x38a1cfcd.
//
// Solidity: function updateQuickLogin(bytes32 registerEmail, bytes32 username, uint256 nonce, bool quickLogin, bytes sig, uint256 sigKeyIndex) returns()
func (_Entry *EntryTransactorSession) UpdateQuickLogin(registerEmail [32]byte, username [32]byte, nonce *big.Int, quickLogin bool, sig []byte, sigKeyIndex *big.Int) (*types.Transaction, error) {
	return _Entry.Contract.UpdateQuickLogin(&_Entry.TransactOpts, registerEmail, username, nonce, quickLogin, sig, sigKeyIndex)
}

// UpdateRecoveryEmail is a paid mutator transaction binding the contract method 0xfd8aa3c9.
//
// Solidity: function updateRecoveryEmail(bytes32 registerEmail, bytes32 username, uint256 nonce, bytes32[] emails, uint256 threshold, bytes sig, uint256 sigKeyIndex) returns()
func (_Entry *EntryTransactor) UpdateRecoveryEmail(opts *bind.TransactOpts, registerEmail [32]byte, username [32]byte, nonce *big.Int, emails [][32]byte, threshold *big.Int, sig []byte, sigKeyIndex *big.Int) (*types.Transaction, error) {
	return _Entry.contract.Transact(opts, "updateRecoveryEmail", registerEmail, username, nonce, emails, threshold, sig, sigKeyIndex)
}

// UpdateRecoveryEmail is a paid mutator transaction binding the contract method 0xfd8aa3c9.
//
// Solidity: function updateRecoveryEmail(bytes32 registerEmail, bytes32 username, uint256 nonce, bytes32[] emails, uint256 threshold, bytes sig, uint256 sigKeyIndex) returns()
func (_Entry *EntrySession) UpdateRecoveryEmail(registerEmail [32]byte, username [32]byte, nonce *big.Int, emails [][32]byte, threshold *big.Int, sig []byte, sigKeyIndex *big.Int) (*types.Transaction, error) {
	return _Entry.Contract.UpdateRecoveryEmail(&_Entry.TransactOpts, registerEmail, username, nonce, emails, threshold, sig, sigKeyIndex)
}

// UpdateRecoveryEmail is a paid mutator transaction binding the contract method 0xfd8aa3c9.
//
// Solidity: function updateRecoveryEmail(bytes32 registerEmail, bytes32 username, uint256 nonce, bytes32[] emails, uint256 threshold, bytes sig, uint256 sigKeyIndex) returns()
func (_Entry *EntryTransactorSession) UpdateRecoveryEmail(registerEmail [32]byte, username [32]byte, nonce *big.Int, emails [][32]byte, threshold *big.Int, sig []byte, sigKeyIndex *big.Int) (*types.Transaction, error) {
	return _Entry.Contract.UpdateRecoveryEmail(&_Entry.TransactOpts, registerEmail, username, nonce, emails, threshold, sig, sigKeyIndex)
}

// EntryEAddLocalKeyIterator is returned from FilterEAddLocalKey and is used to iterate over the raw logs and unpacked data for EAddLocalKey events raised by the Entry contract.
type EntryEAddLocalKeyIterator struct {
	Event *EntryEAddLocalKey // Event containing the contract specifics and raw log

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
func (it *EntryEAddLocalKeyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EntryEAddLocalKey)
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
		it.Event = new(EntryEAddLocalKey)
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
func (it *EntryEAddLocalKeyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EntryEAddLocalKeyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EntryEAddLocalKey represents a EAddLocalKey event raised by the Entry contract.
type EntryEAddLocalKey struct {
	Arg0 [32]byte
	Arg1 *big.Int
	Arg2 []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterEAddLocalKey is a free log retrieval operation binding the contract event 0xba01e23b19ec4f972ccd3df77ca29e838669270d72915e62fde2bf0c8858c587.
//
// Solidity: event E_AddLocalKey(bytes32 arg0, uint256 arg1, bytes arg2)
func (_Entry *EntryFilterer) FilterEAddLocalKey(opts *bind.FilterOpts) (*EntryEAddLocalKeyIterator, error) {

	logs, sub, err := _Entry.contract.FilterLogs(opts, "E_AddLocalKey")
	if err != nil {
		return nil, err
	}
	return &EntryEAddLocalKeyIterator{contract: _Entry.contract, event: "E_AddLocalKey", logs: logs, sub: sub}, nil
}

// WatchEAddLocalKey is a free log subscription operation binding the contract event 0xba01e23b19ec4f972ccd3df77ca29e838669270d72915e62fde2bf0c8858c587.
//
// Solidity: event E_AddLocalKey(bytes32 arg0, uint256 arg1, bytes arg2)
func (_Entry *EntryFilterer) WatchEAddLocalKey(opts *bind.WatchOpts, sink chan<- *EntryEAddLocalKey) (event.Subscription, error) {

	logs, sub, err := _Entry.contract.WatchLogs(opts, "E_AddLocalKey")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EntryEAddLocalKey)
				if err := _Entry.contract.UnpackLog(event, "E_AddLocalKey", log); err != nil {
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

// ParseEAddLocalKey is a log parse operation binding the contract event 0xba01e23b19ec4f972ccd3df77ca29e838669270d72915e62fde2bf0c8858c587.
//
// Solidity: event E_AddLocalKey(bytes32 arg0, uint256 arg1, bytes arg2)
func (_Entry *EntryFilterer) ParseEAddLocalKey(log types.Log) (*EntryEAddLocalKey, error) {
	event := new(EntryEAddLocalKey)
	if err := _Entry.contract.UnpackLog(event, "E_AddLocalKey", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EntryECancelRecoveryIterator is returned from FilterECancelRecovery and is used to iterate over the raw logs and unpacked data for ECancelRecovery events raised by the Entry contract.
type EntryECancelRecoveryIterator struct {
	Event *EntryECancelRecovery // Event containing the contract specifics and raw log

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
func (it *EntryECancelRecoveryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EntryECancelRecovery)
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
		it.Event = new(EntryECancelRecovery)
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
func (it *EntryECancelRecoveryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EntryECancelRecoveryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EntryECancelRecovery represents a ECancelRecovery event raised by the Entry contract.
type EntryECancelRecovery struct {
	Arg0 [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterECancelRecovery is a free log retrieval operation binding the contract event 0x4b1423cdf92637eaa20ad2ef2a8221d879fcc6d5760d373d27169a42c2676756.
//
// Solidity: event E_CancelRecovery(bytes32 arg0)
func (_Entry *EntryFilterer) FilterECancelRecovery(opts *bind.FilterOpts) (*EntryECancelRecoveryIterator, error) {

	logs, sub, err := _Entry.contract.FilterLogs(opts, "E_CancelRecovery")
	if err != nil {
		return nil, err
	}
	return &EntryECancelRecoveryIterator{contract: _Entry.contract, event: "E_CancelRecovery", logs: logs, sub: sub}, nil
}

// WatchECancelRecovery is a free log subscription operation binding the contract event 0x4b1423cdf92637eaa20ad2ef2a8221d879fcc6d5760d373d27169a42c2676756.
//
// Solidity: event E_CancelRecovery(bytes32 arg0)
func (_Entry *EntryFilterer) WatchECancelRecovery(opts *bind.WatchOpts, sink chan<- *EntryECancelRecovery) (event.Subscription, error) {

	logs, sub, err := _Entry.contract.WatchLogs(opts, "E_CancelRecovery")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EntryECancelRecovery)
				if err := _Entry.contract.UnpackLog(event, "E_CancelRecovery", log); err != nil {
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

// ParseECancelRecovery is a log parse operation binding the contract event 0x4b1423cdf92637eaa20ad2ef2a8221d879fcc6d5760d373d27169a42c2676756.
//
// Solidity: event E_CancelRecovery(bytes32 arg0)
func (_Entry *EntryFilterer) ParseECancelRecovery(log types.Log) (*EntryECancelRecovery, error) {
	event := new(EntryECancelRecovery)
	if err := _Entry.contract.UnpackLog(event, "E_CancelRecovery", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EntryECompleteRecoveryIterator is returned from FilterECompleteRecovery and is used to iterate over the raw logs and unpacked data for ECompleteRecovery events raised by the Entry contract.
type EntryECompleteRecoveryIterator struct {
	Event *EntryECompleteRecovery // Event containing the contract specifics and raw log

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
func (it *EntryECompleteRecoveryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EntryECompleteRecovery)
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
		it.Event = new(EntryECompleteRecovery)
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
func (it *EntryECompleteRecoveryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EntryECompleteRecoveryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EntryECompleteRecovery represents a ECompleteRecovery event raised by the Entry contract.
type EntryECompleteRecovery struct {
	Arg0 [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterECompleteRecovery is a free log retrieval operation binding the contract event 0xd79e6a66d3f7988e922fc879bad67c849349032f464bd95c5738a5b7318c593a.
//
// Solidity: event E_CompleteRecovery(bytes32 arg0)
func (_Entry *EntryFilterer) FilterECompleteRecovery(opts *bind.FilterOpts) (*EntryECompleteRecoveryIterator, error) {

	logs, sub, err := _Entry.contract.FilterLogs(opts, "E_CompleteRecovery")
	if err != nil {
		return nil, err
	}
	return &EntryECompleteRecoveryIterator{contract: _Entry.contract, event: "E_CompleteRecovery", logs: logs, sub: sub}, nil
}

// WatchECompleteRecovery is a free log subscription operation binding the contract event 0xd79e6a66d3f7988e922fc879bad67c849349032f464bd95c5738a5b7318c593a.
//
// Solidity: event E_CompleteRecovery(bytes32 arg0)
func (_Entry *EntryFilterer) WatchECompleteRecovery(opts *bind.WatchOpts, sink chan<- *EntryECompleteRecovery) (event.Subscription, error) {

	logs, sub, err := _Entry.contract.WatchLogs(opts, "E_CompleteRecovery")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EntryECompleteRecovery)
				if err := _Entry.contract.UnpackLog(event, "E_CompleteRecovery", log); err != nil {
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

// ParseECompleteRecovery is a log parse operation binding the contract event 0xd79e6a66d3f7988e922fc879bad67c849349032f464bd95c5738a5b7318c593a.
//
// Solidity: event E_CompleteRecovery(bytes32 arg0)
func (_Entry *EntryFilterer) ParseECompleteRecovery(log types.Log) (*EntryECompleteRecovery, error) {
	event := new(EntryECompleteRecovery)
	if err := _Entry.contract.UnpackLog(event, "E_CompleteRecovery", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EntryEDelLocalKeyIterator is returned from FilterEDelLocalKey and is used to iterate over the raw logs and unpacked data for EDelLocalKey events raised by the Entry contract.
type EntryEDelLocalKeyIterator struct {
	Event *EntryEDelLocalKey // Event containing the contract specifics and raw log

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
func (it *EntryEDelLocalKeyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EntryEDelLocalKey)
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
		it.Event = new(EntryEDelLocalKey)
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
func (it *EntryEDelLocalKeyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EntryEDelLocalKeyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EntryEDelLocalKey represents a EDelLocalKey event raised by the Entry contract.
type EntryEDelLocalKey struct {
	Arg0 [32]byte
	Arg1 *big.Int
	Arg2 []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterEDelLocalKey is a free log retrieval operation binding the contract event 0xf68263e2bbfc3f94b732e3bed7d28a3f38bcf4b96144c0bb3df6bde48decf6ee.
//
// Solidity: event E_DelLocalKey(bytes32 arg0, uint256 arg1, bytes arg2)
func (_Entry *EntryFilterer) FilterEDelLocalKey(opts *bind.FilterOpts) (*EntryEDelLocalKeyIterator, error) {

	logs, sub, err := _Entry.contract.FilterLogs(opts, "E_DelLocalKey")
	if err != nil {
		return nil, err
	}
	return &EntryEDelLocalKeyIterator{contract: _Entry.contract, event: "E_DelLocalKey", logs: logs, sub: sub}, nil
}

// WatchEDelLocalKey is a free log subscription operation binding the contract event 0xf68263e2bbfc3f94b732e3bed7d28a3f38bcf4b96144c0bb3df6bde48decf6ee.
//
// Solidity: event E_DelLocalKey(bytes32 arg0, uint256 arg1, bytes arg2)
func (_Entry *EntryFilterer) WatchEDelLocalKey(opts *bind.WatchOpts, sink chan<- *EntryEDelLocalKey) (event.Subscription, error) {

	logs, sub, err := _Entry.contract.WatchLogs(opts, "E_DelLocalKey")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EntryEDelLocalKey)
				if err := _Entry.contract.UnpackLog(event, "E_DelLocalKey", log); err != nil {
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

// ParseEDelLocalKey is a log parse operation binding the contract event 0xf68263e2bbfc3f94b732e3bed7d28a3f38bcf4b96144c0bb3df6bde48decf6ee.
//
// Solidity: event E_DelLocalKey(bytes32 arg0, uint256 arg1, bytes arg2)
func (_Entry *EntryFilterer) ParseEDelLocalKey(log types.Log) (*EntryEDelLocalKey, error) {
	event := new(EntryEDelLocalKey)
	if err := _Entry.contract.UnpackLog(event, "E_DelLocalKey", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EntryEQuickAddLocalKeyIterator is returned from FilterEQuickAddLocalKey and is used to iterate over the raw logs and unpacked data for EQuickAddLocalKey events raised by the Entry contract.
type EntryEQuickAddLocalKeyIterator struct {
	Event *EntryEQuickAddLocalKey // Event containing the contract specifics and raw log

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
func (it *EntryEQuickAddLocalKeyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EntryEQuickAddLocalKey)
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
		it.Event = new(EntryEQuickAddLocalKey)
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
func (it *EntryEQuickAddLocalKeyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EntryEQuickAddLocalKeyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EntryEQuickAddLocalKey represents a EQuickAddLocalKey event raised by the Entry contract.
type EntryEQuickAddLocalKey struct {
	Arg0 [32]byte
	Arg1 *big.Int
	Arg2 []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterEQuickAddLocalKey is a free log retrieval operation binding the contract event 0x292d4a3525ab6e4806162015f244867dd0373fd03eba76874d1b46b749060079.
//
// Solidity: event E_QuickAddLocalKey(bytes32 arg0, uint256 arg1, bytes arg2)
func (_Entry *EntryFilterer) FilterEQuickAddLocalKey(opts *bind.FilterOpts) (*EntryEQuickAddLocalKeyIterator, error) {

	logs, sub, err := _Entry.contract.FilterLogs(opts, "E_QuickAddLocalKey")
	if err != nil {
		return nil, err
	}
	return &EntryEQuickAddLocalKeyIterator{contract: _Entry.contract, event: "E_QuickAddLocalKey", logs: logs, sub: sub}, nil
}

// WatchEQuickAddLocalKey is a free log subscription operation binding the contract event 0x292d4a3525ab6e4806162015f244867dd0373fd03eba76874d1b46b749060079.
//
// Solidity: event E_QuickAddLocalKey(bytes32 arg0, uint256 arg1, bytes arg2)
func (_Entry *EntryFilterer) WatchEQuickAddLocalKey(opts *bind.WatchOpts, sink chan<- *EntryEQuickAddLocalKey) (event.Subscription, error) {

	logs, sub, err := _Entry.contract.WatchLogs(opts, "E_QuickAddLocalKey")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EntryEQuickAddLocalKey)
				if err := _Entry.contract.UnpackLog(event, "E_QuickAddLocalKey", log); err != nil {
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

// ParseEQuickAddLocalKey is a log parse operation binding the contract event 0x292d4a3525ab6e4806162015f244867dd0373fd03eba76874d1b46b749060079.
//
// Solidity: event E_QuickAddLocalKey(bytes32 arg0, uint256 arg1, bytes arg2)
func (_Entry *EntryFilterer) ParseEQuickAddLocalKey(log types.Log) (*EntryEQuickAddLocalKey, error) {
	event := new(EntryEQuickAddLocalKey)
	if err := _Entry.contract.UnpackLog(event, "E_QuickAddLocalKey", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EntryEQuickRegisterIterator is returned from FilterEQuickRegister and is used to iterate over the raw logs and unpacked data for EQuickRegister events raised by the Entry contract.
type EntryEQuickRegisterIterator struct {
	Event *EntryEQuickRegister // Event containing the contract specifics and raw log

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
func (it *EntryEQuickRegisterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EntryEQuickRegister)
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
		it.Event = new(EntryEQuickRegister)
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
func (it *EntryEQuickRegisterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EntryEQuickRegisterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EntryEQuickRegister represents a EQuickRegister event raised by the Entry contract.
type EntryEQuickRegister struct {
	Arg0 [32]byte
	Arg1 string
	Arg2 string
	Arg3 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterEQuickRegister is a free log retrieval operation binding the contract event 0x517adabacc69023b2ae60a6475142ead892f0791ad5541623bade97d3350c2db.
//
// Solidity: event E_QuickRegister(bytes32 arg0, string arg1, string arg2, address arg3)
func (_Entry *EntryFilterer) FilterEQuickRegister(opts *bind.FilterOpts) (*EntryEQuickRegisterIterator, error) {

	logs, sub, err := _Entry.contract.FilterLogs(opts, "E_QuickRegister")
	if err != nil {
		return nil, err
	}
	return &EntryEQuickRegisterIterator{contract: _Entry.contract, event: "E_QuickRegister", logs: logs, sub: sub}, nil
}

// WatchEQuickRegister is a free log subscription operation binding the contract event 0x517adabacc69023b2ae60a6475142ead892f0791ad5541623bade97d3350c2db.
//
// Solidity: event E_QuickRegister(bytes32 arg0, string arg1, string arg2, address arg3)
func (_Entry *EntryFilterer) WatchEQuickRegister(opts *bind.WatchOpts, sink chan<- *EntryEQuickRegister) (event.Subscription, error) {

	logs, sub, err := _Entry.contract.WatchLogs(opts, "E_QuickRegister")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EntryEQuickRegister)
				if err := _Entry.contract.UnpackLog(event, "E_QuickRegister", log); err != nil {
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

// ParseEQuickRegister is a log parse operation binding the contract event 0x517adabacc69023b2ae60a6475142ead892f0791ad5541623bade97d3350c2db.
//
// Solidity: event E_QuickRegister(bytes32 arg0, string arg1, string arg2, address arg3)
func (_Entry *EntryFilterer) ParseEQuickRegister(log types.Log) (*EntryEQuickRegister, error) {
	event := new(EntryEQuickRegister)
	if err := _Entry.contract.UnpackLog(event, "E_QuickRegister", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EntryERegisterIterator is returned from FilterERegister and is used to iterate over the raw logs and unpacked data for ERegister events raised by the Entry contract.
type EntryERegisterIterator struct {
	Event *EntryERegister // Event containing the contract specifics and raw log

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
func (it *EntryERegisterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EntryERegister)
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
		it.Event = new(EntryERegister)
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
func (it *EntryERegisterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EntryERegisterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EntryERegister represents a ERegister event raised by the Entry contract.
type EntryERegister struct {
	Arg0 [32]byte
	Arg1 string
	Arg2 string
	Arg3 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterERegister is a free log retrieval operation binding the contract event 0x6a00be676778e9c189742c5801bc41a11ebdf6ea78444d02974c250b3a747736.
//
// Solidity: event E_Register(bytes32 arg0, string arg1, string arg2, address arg3)
func (_Entry *EntryFilterer) FilterERegister(opts *bind.FilterOpts) (*EntryERegisterIterator, error) {

	logs, sub, err := _Entry.contract.FilterLogs(opts, "E_Register")
	if err != nil {
		return nil, err
	}
	return &EntryERegisterIterator{contract: _Entry.contract, event: "E_Register", logs: logs, sub: sub}, nil
}

// WatchERegister is a free log subscription operation binding the contract event 0x6a00be676778e9c189742c5801bc41a11ebdf6ea78444d02974c250b3a747736.
//
// Solidity: event E_Register(bytes32 arg0, string arg1, string arg2, address arg3)
func (_Entry *EntryFilterer) WatchERegister(opts *bind.WatchOpts, sink chan<- *EntryERegister) (event.Subscription, error) {

	logs, sub, err := _Entry.contract.WatchLogs(opts, "E_Register")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EntryERegister)
				if err := _Entry.contract.UnpackLog(event, "E_Register", log); err != nil {
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

// ParseERegister is a log parse operation binding the contract event 0x6a00be676778e9c189742c5801bc41a11ebdf6ea78444d02974c250b3a747736.
//
// Solidity: event E_Register(bytes32 arg0, string arg1, string arg2, address arg3)
func (_Entry *EntryFilterer) ParseERegister(log types.Log) (*EntryERegister, error) {
	event := new(EntryERegister)
	if err := _Entry.contract.UnpackLog(event, "E_Register", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EntryESetAdminIterator is returned from FilterESetAdmin and is used to iterate over the raw logs and unpacked data for ESetAdmin events raised by the Entry contract.
type EntryESetAdminIterator struct {
	Event *EntryESetAdmin // Event containing the contract specifics and raw log

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
func (it *EntryESetAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EntryESetAdmin)
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
		it.Event = new(EntryESetAdmin)
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
func (it *EntryESetAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EntryESetAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EntryESetAdmin represents a ESetAdmin event raised by the Entry contract.
type EntryESetAdmin struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterESetAdmin is a free log retrieval operation binding the contract event 0x55a6de4609f40ad7bdd9ad01a493f1d45ffeca88a938c76910318fb4ff5b3f44.
//
// Solidity: event E_SetAdmin(address arg0)
func (_Entry *EntryFilterer) FilterESetAdmin(opts *bind.FilterOpts) (*EntryESetAdminIterator, error) {

	logs, sub, err := _Entry.contract.FilterLogs(opts, "E_SetAdmin")
	if err != nil {
		return nil, err
	}
	return &EntryESetAdminIterator{contract: _Entry.contract, event: "E_SetAdmin", logs: logs, sub: sub}, nil
}

// WatchESetAdmin is a free log subscription operation binding the contract event 0x55a6de4609f40ad7bdd9ad01a493f1d45ffeca88a938c76910318fb4ff5b3f44.
//
// Solidity: event E_SetAdmin(address arg0)
func (_Entry *EntryFilterer) WatchESetAdmin(opts *bind.WatchOpts, sink chan<- *EntryESetAdmin) (event.Subscription, error) {

	logs, sub, err := _Entry.contract.WatchLogs(opts, "E_SetAdmin")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EntryESetAdmin)
				if err := _Entry.contract.UnpackLog(event, "E_SetAdmin", log); err != nil {
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

// ParseESetAdmin is a log parse operation binding the contract event 0x55a6de4609f40ad7bdd9ad01a493f1d45ffeca88a938c76910318fb4ff5b3f44.
//
// Solidity: event E_SetAdmin(address arg0)
func (_Entry *EntryFilterer) ParseESetAdmin(log types.Log) (*EntryESetAdmin, error) {
	event := new(EntryESetAdmin)
	if err := _Entry.contract.UnpackLog(event, "E_SetAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EntryESetTimeLockSwitchIterator is returned from FilterESetTimeLockSwitch and is used to iterate over the raw logs and unpacked data for ESetTimeLockSwitch events raised by the Entry contract.
type EntryESetTimeLockSwitchIterator struct {
	Event *EntryESetTimeLockSwitch // Event containing the contract specifics and raw log

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
func (it *EntryESetTimeLockSwitchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EntryESetTimeLockSwitch)
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
		it.Event = new(EntryESetTimeLockSwitch)
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
func (it *EntryESetTimeLockSwitchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EntryESetTimeLockSwitchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EntryESetTimeLockSwitch represents a ESetTimeLockSwitch event raised by the Entry contract.
type EntryESetTimeLockSwitch struct {
	TimeLockSwitch bool
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterESetTimeLockSwitch is a free log retrieval operation binding the contract event 0xd1e3419646ae3203b9f48668f59b37e95b663c6da65e28a49ecbe78ec15c7cf3.
//
// Solidity: event E_SetTimeLockSwitch(bool timeLockSwitch)
func (_Entry *EntryFilterer) FilterESetTimeLockSwitch(opts *bind.FilterOpts) (*EntryESetTimeLockSwitchIterator, error) {

	logs, sub, err := _Entry.contract.FilterLogs(opts, "E_SetTimeLockSwitch")
	if err != nil {
		return nil, err
	}
	return &EntryESetTimeLockSwitchIterator{contract: _Entry.contract, event: "E_SetTimeLockSwitch", logs: logs, sub: sub}, nil
}

// WatchESetTimeLockSwitch is a free log subscription operation binding the contract event 0xd1e3419646ae3203b9f48668f59b37e95b663c6da65e28a49ecbe78ec15c7cf3.
//
// Solidity: event E_SetTimeLockSwitch(bool timeLockSwitch)
func (_Entry *EntryFilterer) WatchESetTimeLockSwitch(opts *bind.WatchOpts, sink chan<- *EntryESetTimeLockSwitch) (event.Subscription, error) {

	logs, sub, err := _Entry.contract.WatchLogs(opts, "E_SetTimeLockSwitch")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EntryESetTimeLockSwitch)
				if err := _Entry.contract.UnpackLog(event, "E_SetTimeLockSwitch", log); err != nil {
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

// ParseESetTimeLockSwitch is a log parse operation binding the contract event 0xd1e3419646ae3203b9f48668f59b37e95b663c6da65e28a49ecbe78ec15c7cf3.
//
// Solidity: event E_SetTimeLockSwitch(bool timeLockSwitch)
func (_Entry *EntryFilterer) ParseESetTimeLockSwitch(log types.Log) (*EntryESetTimeLockSwitch, error) {
	event := new(EntryESetTimeLockSwitch)
	if err := _Entry.contract.UnpackLog(event, "E_SetTimeLockSwitch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EntryEStartRecoveryIterator is returned from FilterEStartRecovery and is used to iterate over the raw logs and unpacked data for EStartRecovery events raised by the Entry contract.
type EntryEStartRecoveryIterator struct {
	Event *EntryEStartRecovery // Event containing the contract specifics and raw log

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
func (it *EntryEStartRecoveryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EntryEStartRecovery)
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
		it.Event = new(EntryEStartRecovery)
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
func (it *EntryEStartRecoveryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EntryEStartRecoveryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EntryEStartRecovery represents a EStartRecovery event raised by the Entry contract.
type EntryEStartRecovery struct {
	Arg0 [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterEStartRecovery is a free log retrieval operation binding the contract event 0x5ee7951198402251425eb336794576897218caab80e12b7cbabd198afcb8040e.
//
// Solidity: event E_StartRecovery(bytes32 arg0)
func (_Entry *EntryFilterer) FilterEStartRecovery(opts *bind.FilterOpts) (*EntryEStartRecoveryIterator, error) {

	logs, sub, err := _Entry.contract.FilterLogs(opts, "E_StartRecovery")
	if err != nil {
		return nil, err
	}
	return &EntryEStartRecoveryIterator{contract: _Entry.contract, event: "E_StartRecovery", logs: logs, sub: sub}, nil
}

// WatchEStartRecovery is a free log subscription operation binding the contract event 0x5ee7951198402251425eb336794576897218caab80e12b7cbabd198afcb8040e.
//
// Solidity: event E_StartRecovery(bytes32 arg0)
func (_Entry *EntryFilterer) WatchEStartRecovery(opts *bind.WatchOpts, sink chan<- *EntryEStartRecovery) (event.Subscription, error) {

	logs, sub, err := _Entry.contract.WatchLogs(opts, "E_StartRecovery")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EntryEStartRecovery)
				if err := _Entry.contract.UnpackLog(event, "E_StartRecovery", log); err != nil {
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

// ParseEStartRecovery is a log parse operation binding the contract event 0x5ee7951198402251425eb336794576897218caab80e12b7cbabd198afcb8040e.
//
// Solidity: event E_StartRecovery(bytes32 arg0)
func (_Entry *EntryFilterer) ParseEStartRecovery(log types.Log) (*EntryEStartRecovery, error) {
	event := new(EntryEStartRecovery)
	if err := _Entry.contract.UnpackLog(event, "E_StartRecovery", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EntryEUpdateQuickLoginIterator is returned from FilterEUpdateQuickLogin and is used to iterate over the raw logs and unpacked data for EUpdateQuickLogin events raised by the Entry contract.
type EntryEUpdateQuickLoginIterator struct {
	Event *EntryEUpdateQuickLogin // Event containing the contract specifics and raw log

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
func (it *EntryEUpdateQuickLoginIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EntryEUpdateQuickLogin)
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
		it.Event = new(EntryEUpdateQuickLogin)
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
func (it *EntryEUpdateQuickLoginIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EntryEUpdateQuickLoginIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EntryEUpdateQuickLogin represents a EUpdateQuickLogin event raised by the Entry contract.
type EntryEUpdateQuickLogin struct {
	Arg0 [32]byte
	Arg1 bool
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterEUpdateQuickLogin is a free log retrieval operation binding the contract event 0x5dd674453b4973bed65ef808aa323f191e249ea7af997b9edaa0d525bb6a260a.
//
// Solidity: event E_UpdateQuickLogin(bytes32 arg0, bool arg1)
func (_Entry *EntryFilterer) FilterEUpdateQuickLogin(opts *bind.FilterOpts) (*EntryEUpdateQuickLoginIterator, error) {

	logs, sub, err := _Entry.contract.FilterLogs(opts, "E_UpdateQuickLogin")
	if err != nil {
		return nil, err
	}
	return &EntryEUpdateQuickLoginIterator{contract: _Entry.contract, event: "E_UpdateQuickLogin", logs: logs, sub: sub}, nil
}

// WatchEUpdateQuickLogin is a free log subscription operation binding the contract event 0x5dd674453b4973bed65ef808aa323f191e249ea7af997b9edaa0d525bb6a260a.
//
// Solidity: event E_UpdateQuickLogin(bytes32 arg0, bool arg1)
func (_Entry *EntryFilterer) WatchEUpdateQuickLogin(opts *bind.WatchOpts, sink chan<- *EntryEUpdateQuickLogin) (event.Subscription, error) {

	logs, sub, err := _Entry.contract.WatchLogs(opts, "E_UpdateQuickLogin")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EntryEUpdateQuickLogin)
				if err := _Entry.contract.UnpackLog(event, "E_UpdateQuickLogin", log); err != nil {
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

// ParseEUpdateQuickLogin is a log parse operation binding the contract event 0x5dd674453b4973bed65ef808aa323f191e249ea7af997b9edaa0d525bb6a260a.
//
// Solidity: event E_UpdateQuickLogin(bytes32 arg0, bool arg1)
func (_Entry *EntryFilterer) ParseEUpdateQuickLogin(log types.Log) (*EntryEUpdateQuickLogin, error) {
	event := new(EntryEUpdateQuickLogin)
	if err := _Entry.contract.UnpackLog(event, "E_UpdateQuickLogin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EntryEUpdateRecoveryEmailIterator is returned from FilterEUpdateRecoveryEmail and is used to iterate over the raw logs and unpacked data for EUpdateRecoveryEmail events raised by the Entry contract.
type EntryEUpdateRecoveryEmailIterator struct {
	Event *EntryEUpdateRecoveryEmail // Event containing the contract specifics and raw log

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
func (it *EntryEUpdateRecoveryEmailIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EntryEUpdateRecoveryEmail)
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
		it.Event = new(EntryEUpdateRecoveryEmail)
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
func (it *EntryEUpdateRecoveryEmailIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EntryEUpdateRecoveryEmailIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EntryEUpdateRecoveryEmail represents a EUpdateRecoveryEmail event raised by the Entry contract.
type EntryEUpdateRecoveryEmail struct {
	Arg0 [32]byte
	Arg1 [][32]byte
	Arg2 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterEUpdateRecoveryEmail is a free log retrieval operation binding the contract event 0x9151ae2a7fd84fa316b5ff0f1479ded5e5f7cff9de7661dcf1e08f868e3282dc.
//
// Solidity: event E_UpdateRecoveryEmail(bytes32 arg0, bytes32[] arg1, uint256 arg2)
func (_Entry *EntryFilterer) FilterEUpdateRecoveryEmail(opts *bind.FilterOpts) (*EntryEUpdateRecoveryEmailIterator, error) {

	logs, sub, err := _Entry.contract.FilterLogs(opts, "E_UpdateRecoveryEmail")
	if err != nil {
		return nil, err
	}
	return &EntryEUpdateRecoveryEmailIterator{contract: _Entry.contract, event: "E_UpdateRecoveryEmail", logs: logs, sub: sub}, nil
}

// WatchEUpdateRecoveryEmail is a free log subscription operation binding the contract event 0x9151ae2a7fd84fa316b5ff0f1479ded5e5f7cff9de7661dcf1e08f868e3282dc.
//
// Solidity: event E_UpdateRecoveryEmail(bytes32 arg0, bytes32[] arg1, uint256 arg2)
func (_Entry *EntryFilterer) WatchEUpdateRecoveryEmail(opts *bind.WatchOpts, sink chan<- *EntryEUpdateRecoveryEmail) (event.Subscription, error) {

	logs, sub, err := _Entry.contract.WatchLogs(opts, "E_UpdateRecoveryEmail")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EntryEUpdateRecoveryEmail)
				if err := _Entry.contract.UnpackLog(event, "E_UpdateRecoveryEmail", log); err != nil {
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

// ParseEUpdateRecoveryEmail is a log parse operation binding the contract event 0x9151ae2a7fd84fa316b5ff0f1479ded5e5f7cff9de7661dcf1e08f868e3282dc.
//
// Solidity: event E_UpdateRecoveryEmail(bytes32 arg0, bytes32[] arg1, uint256 arg2)
func (_Entry *EntryFilterer) ParseEUpdateRecoveryEmail(log types.Log) (*EntryEUpdateRecoveryEmail, error) {
	event := new(EntryEUpdateRecoveryEmail)
	if err := _Entry.contract.UnpackLog(event, "E_UpdateRecoveryEmail", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

