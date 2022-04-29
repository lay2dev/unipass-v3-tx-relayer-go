// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package asset

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

// AssetMetaData contains all meta data concerning the Asset contract.
var AssetMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"E_Call_Failed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"E_Call_Success\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"E_Transfer_Native\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"E_Transfer_Token\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"assetNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chainID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"entryContract\",\"outputs\":[{\"internalType\":\"contractEntry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"sigKeyIndex\",\"type\":\"uint256\"}],\"name\":\"execute\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_entryAddr\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_registerEmail\",\"type\":\"bytes32\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registerEmail\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"sigKeyIndex\",\"type\":\"uint256\"}],\"name\":\"transferNativeToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"sigKeyIndex\",\"type\":\"uint256\"}],\"name\":\"transferToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// AssetABI is the input ABI used to generate the binding from.
// Deprecated: Use AssetMetaData.ABI instead.
var AssetABI = AssetMetaData.ABI

// Asset is an auto generated Go binding around an Ethereum contract.
type Asset struct {
	AssetCaller     // Read-only binding to the contract
	AssetTransactor // Write-only binding to the contract
	AssetFilterer   // Log filterer for contract events
}

// AssetCaller is an auto generated read-only Go binding around an Ethereum contract.
type AssetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AssetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AssetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AssetSession struct {
	Contract     *Asset            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AssetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AssetCallerSession struct {
	Contract *AssetCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AssetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AssetTransactorSession struct {
	Contract     *AssetTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AssetRaw is an auto generated low-level Go binding around an Ethereum contract.
type AssetRaw struct {
	Contract *Asset // Generic contract binding to access the raw methods on
}

// AssetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AssetCallerRaw struct {
	Contract *AssetCaller // Generic read-only contract binding to access the raw methods on
}

// AssetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AssetTransactorRaw struct {
	Contract *AssetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAsset creates a new instance of Asset, bound to a specific deployed contract.
func NewAsset(address common.Address, backend bind.ContractBackend) (*Asset, error) {
	contract, err := bindAsset(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Asset{AssetCaller: AssetCaller{contract: contract}, AssetTransactor: AssetTransactor{contract: contract}, AssetFilterer: AssetFilterer{contract: contract}}, nil
}

// NewAssetCaller creates a new read-only instance of Asset, bound to a specific deployed contract.
func NewAssetCaller(address common.Address, caller bind.ContractCaller) (*AssetCaller, error) {
	contract, err := bindAsset(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AssetCaller{contract: contract}, nil
}

// NewAssetTransactor creates a new write-only instance of Asset, bound to a specific deployed contract.
func NewAssetTransactor(address common.Address, transactor bind.ContractTransactor) (*AssetTransactor, error) {
	contract, err := bindAsset(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AssetTransactor{contract: contract}, nil
}

// NewAssetFilterer creates a new log filterer instance of Asset, bound to a specific deployed contract.
func NewAssetFilterer(address common.Address, filterer bind.ContractFilterer) (*AssetFilterer, error) {
	contract, err := bindAsset(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AssetFilterer{contract: contract}, nil
}

// bindAsset binds a generic wrapper to an already deployed contract.
func bindAsset(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AssetABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Asset *AssetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Asset.Contract.AssetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Asset *AssetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Asset.Contract.AssetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Asset *AssetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Asset.Contract.AssetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Asset *AssetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Asset.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Asset *AssetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Asset.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Asset *AssetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Asset.Contract.contract.Transact(opts, method, params...)
}

// AssetNonce is a free data retrieval call binding the contract method 0x7ae9925a.
//
// Solidity: function assetNonce() view returns(uint256)
func (_Asset *AssetCaller) AssetNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Asset.contract.Call(opts, &out, "assetNonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AssetNonce is a free data retrieval call binding the contract method 0x7ae9925a.
//
// Solidity: function assetNonce() view returns(uint256)
func (_Asset *AssetSession) AssetNonce() (*big.Int, error) {
	return _Asset.Contract.AssetNonce(&_Asset.CallOpts)
}

// AssetNonce is a free data retrieval call binding the contract method 0x7ae9925a.
//
// Solidity: function assetNonce() view returns(uint256)
func (_Asset *AssetCallerSession) AssetNonce() (*big.Int, error) {
	return _Asset.Contract.AssetNonce(&_Asset.CallOpts)
}

// ChainID is a free data retrieval call binding the contract method 0xadc879e9.
//
// Solidity: function chainID() view returns(uint256)
func (_Asset *AssetCaller) ChainID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Asset.contract.Call(opts, &out, "chainID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChainID is a free data retrieval call binding the contract method 0xadc879e9.
//
// Solidity: function chainID() view returns(uint256)
func (_Asset *AssetSession) ChainID() (*big.Int, error) {
	return _Asset.Contract.ChainID(&_Asset.CallOpts)
}

// ChainID is a free data retrieval call binding the contract method 0xadc879e9.
//
// Solidity: function chainID() view returns(uint256)
func (_Asset *AssetCallerSession) ChainID() (*big.Int, error) {
	return _Asset.Contract.ChainID(&_Asset.CallOpts)
}

// EntryContract is a free data retrieval call binding the contract method 0xfb4803f2.
//
// Solidity: function entryContract() view returns(address)
func (_Asset *AssetCaller) EntryContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Asset.contract.Call(opts, &out, "entryContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EntryContract is a free data retrieval call binding the contract method 0xfb4803f2.
//
// Solidity: function entryContract() view returns(address)
func (_Asset *AssetSession) EntryContract() (common.Address, error) {
	return _Asset.Contract.EntryContract(&_Asset.CallOpts)
}

// EntryContract is a free data retrieval call binding the contract method 0xfb4803f2.
//
// Solidity: function entryContract() view returns(address)
func (_Asset *AssetCallerSession) EntryContract() (common.Address, error) {
	return _Asset.Contract.EntryContract(&_Asset.CallOpts)
}

// RegisterEmail is a free data retrieval call binding the contract method 0x517741fb.
//
// Solidity: function registerEmail() view returns(bytes32)
func (_Asset *AssetCaller) RegisterEmail(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Asset.contract.Call(opts, &out, "registerEmail")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RegisterEmail is a free data retrieval call binding the contract method 0x517741fb.
//
// Solidity: function registerEmail() view returns(bytes32)
func (_Asset *AssetSession) RegisterEmail() ([32]byte, error) {
	return _Asset.Contract.RegisterEmail(&_Asset.CallOpts)
}

// RegisterEmail is a free data retrieval call binding the contract method 0x517741fb.
//
// Solidity: function registerEmail() view returns(bytes32)
func (_Asset *AssetCallerSession) RegisterEmail() ([32]byte, error) {
	return _Asset.Contract.RegisterEmail(&_Asset.CallOpts)
}

// Execute is a paid mutator transaction binding the contract method 0x910d7ea8.
//
// Solidity: function execute(uint256 nonce, address to, uint256 value, bytes callData, address feeToken, uint256 feeAmount, bytes sig, uint256 sigKeyIndex) returns(bool, bytes)
func (_Asset *AssetTransactor) Execute(opts *bind.TransactOpts, nonce *big.Int, to common.Address, value *big.Int, callData []byte, feeToken common.Address, feeAmount *big.Int, sig []byte, sigKeyIndex *big.Int) (*types.Transaction, error) {
	return _Asset.contract.Transact(opts, "execute", nonce, to, value, callData, feeToken, feeAmount, sig, sigKeyIndex)
}

// Execute is a paid mutator transaction binding the contract method 0x910d7ea8.
//
// Solidity: function execute(uint256 nonce, address to, uint256 value, bytes callData, address feeToken, uint256 feeAmount, bytes sig, uint256 sigKeyIndex) returns(bool, bytes)
func (_Asset *AssetSession) Execute(nonce *big.Int, to common.Address, value *big.Int, callData []byte, feeToken common.Address, feeAmount *big.Int, sig []byte, sigKeyIndex *big.Int) (*types.Transaction, error) {
	return _Asset.Contract.Execute(&_Asset.TransactOpts, nonce, to, value, callData, feeToken, feeAmount, sig, sigKeyIndex)
}

// Execute is a paid mutator transaction binding the contract method 0x910d7ea8.
//
// Solidity: function execute(uint256 nonce, address to, uint256 value, bytes callData, address feeToken, uint256 feeAmount, bytes sig, uint256 sigKeyIndex) returns(bool, bytes)
func (_Asset *AssetTransactorSession) Execute(nonce *big.Int, to common.Address, value *big.Int, callData []byte, feeToken common.Address, feeAmount *big.Int, sig []byte, sigKeyIndex *big.Int) (*types.Transaction, error) {
	return _Asset.Contract.Execute(&_Asset.TransactOpts, nonce, to, value, callData, feeToken, feeAmount, sig, sigKeyIndex)
}

// Init is a paid mutator transaction binding the contract method 0x2cc0b254.
//
// Solidity: function init(address _entryAddr, bytes32 _registerEmail) returns()
func (_Asset *AssetTransactor) Init(opts *bind.TransactOpts, _entryAddr common.Address, _registerEmail [32]byte) (*types.Transaction, error) {
	return _Asset.contract.Transact(opts, "init", _entryAddr, _registerEmail)
}

// Init is a paid mutator transaction binding the contract method 0x2cc0b254.
//
// Solidity: function init(address _entryAddr, bytes32 _registerEmail) returns()
func (_Asset *AssetSession) Init(_entryAddr common.Address, _registerEmail [32]byte) (*types.Transaction, error) {
	return _Asset.Contract.Init(&_Asset.TransactOpts, _entryAddr, _registerEmail)
}

// Init is a paid mutator transaction binding the contract method 0x2cc0b254.
//
// Solidity: function init(address _entryAddr, bytes32 _registerEmail) returns()
func (_Asset *AssetTransactorSession) Init(_entryAddr common.Address, _registerEmail [32]byte) (*types.Transaction, error) {
	return _Asset.Contract.Init(&_Asset.TransactOpts, _entryAddr, _registerEmail)
}

// TransferNativeToken is a paid mutator transaction binding the contract method 0x3cff2f01.
//
// Solidity: function transferNativeToken(uint256 nonce, address to, uint256 amount, address feeToken, uint256 feeAmount, bytes sig, uint256 sigKeyIndex) returns()
func (_Asset *AssetTransactor) TransferNativeToken(opts *bind.TransactOpts, nonce *big.Int, to common.Address, amount *big.Int, feeToken common.Address, feeAmount *big.Int, sig []byte, sigKeyIndex *big.Int) (*types.Transaction, error) {
	return _Asset.contract.Transact(opts, "transferNativeToken", nonce, to, amount, feeToken, feeAmount, sig, sigKeyIndex)
}

// TransferNativeToken is a paid mutator transaction binding the contract method 0x3cff2f01.
//
// Solidity: function transferNativeToken(uint256 nonce, address to, uint256 amount, address feeToken, uint256 feeAmount, bytes sig, uint256 sigKeyIndex) returns()
func (_Asset *AssetSession) TransferNativeToken(nonce *big.Int, to common.Address, amount *big.Int, feeToken common.Address, feeAmount *big.Int, sig []byte, sigKeyIndex *big.Int) (*types.Transaction, error) {
	return _Asset.Contract.TransferNativeToken(&_Asset.TransactOpts, nonce, to, amount, feeToken, feeAmount, sig, sigKeyIndex)
}

// TransferNativeToken is a paid mutator transaction binding the contract method 0x3cff2f01.
//
// Solidity: function transferNativeToken(uint256 nonce, address to, uint256 amount, address feeToken, uint256 feeAmount, bytes sig, uint256 sigKeyIndex) returns()
func (_Asset *AssetTransactorSession) TransferNativeToken(nonce *big.Int, to common.Address, amount *big.Int, feeToken common.Address, feeAmount *big.Int, sig []byte, sigKeyIndex *big.Int) (*types.Transaction, error) {
	return _Asset.Contract.TransferNativeToken(&_Asset.TransactOpts, nonce, to, amount, feeToken, feeAmount, sig, sigKeyIndex)
}

// TransferToken is a paid mutator transaction binding the contract method 0x6cfd64a5.
//
// Solidity: function transferToken(uint256 nonce, address token, address to, uint256 amount, address feeToken, uint256 feeAmount, bytes sig, uint256 sigKeyIndex) returns()
func (_Asset *AssetTransactor) TransferToken(opts *bind.TransactOpts, nonce *big.Int, token common.Address, to common.Address, amount *big.Int, feeToken common.Address, feeAmount *big.Int, sig []byte, sigKeyIndex *big.Int) (*types.Transaction, error) {
	return _Asset.contract.Transact(opts, "transferToken", nonce, token, to, amount, feeToken, feeAmount, sig, sigKeyIndex)
}

// TransferToken is a paid mutator transaction binding the contract method 0x6cfd64a5.
//
// Solidity: function transferToken(uint256 nonce, address token, address to, uint256 amount, address feeToken, uint256 feeAmount, bytes sig, uint256 sigKeyIndex) returns()
func (_Asset *AssetSession) TransferToken(nonce *big.Int, token common.Address, to common.Address, amount *big.Int, feeToken common.Address, feeAmount *big.Int, sig []byte, sigKeyIndex *big.Int) (*types.Transaction, error) {
	return _Asset.Contract.TransferToken(&_Asset.TransactOpts, nonce, token, to, amount, feeToken, feeAmount, sig, sigKeyIndex)
}

// TransferToken is a paid mutator transaction binding the contract method 0x6cfd64a5.
//
// Solidity: function transferToken(uint256 nonce, address token, address to, uint256 amount, address feeToken, uint256 feeAmount, bytes sig, uint256 sigKeyIndex) returns()
func (_Asset *AssetTransactorSession) TransferToken(nonce *big.Int, token common.Address, to common.Address, amount *big.Int, feeToken common.Address, feeAmount *big.Int, sig []byte, sigKeyIndex *big.Int) (*types.Transaction, error) {
	return _Asset.Contract.TransferToken(&_Asset.TransactOpts, nonce, token, to, amount, feeToken, feeAmount, sig, sigKeyIndex)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Asset *AssetTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Asset.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Asset *AssetSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Asset.Contract.Fallback(&_Asset.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Asset *AssetTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Asset.Contract.Fallback(&_Asset.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Asset *AssetTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Asset.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Asset *AssetSession) Receive() (*types.Transaction, error) {
	return _Asset.Contract.Receive(&_Asset.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Asset *AssetTransactorSession) Receive() (*types.Transaction, error) {
	return _Asset.Contract.Receive(&_Asset.TransactOpts)
}

// AssetECallFailedIterator is returned from FilterECallFailed and is used to iterate over the raw logs and unpacked data for ECallFailed events raised by the Asset contract.
type AssetECallFailedIterator struct {
	Event *AssetECallFailed // Event containing the contract specifics and raw log

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
func (it *AssetECallFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetECallFailed)
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
		it.Event = new(AssetECallFailed)
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
func (it *AssetECallFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetECallFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetECallFailed represents a ECallFailed event raised by the Asset contract.
type AssetECallFailed struct {
	Arg0 [32]byte
	Arg1 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterECallFailed is a free log retrieval operation binding the contract event 0xe58e1506e3b37e68a149230d64d9b29a5234ca9df751c07243bf296a5d046687.
//
// Solidity: event E_Call_Failed(bytes32 arg0, address arg1)
func (_Asset *AssetFilterer) FilterECallFailed(opts *bind.FilterOpts) (*AssetECallFailedIterator, error) {

	logs, sub, err := _Asset.contract.FilterLogs(opts, "E_Call_Failed")
	if err != nil {
		return nil, err
	}
	return &AssetECallFailedIterator{contract: _Asset.contract, event: "E_Call_Failed", logs: logs, sub: sub}, nil
}

// WatchECallFailed is a free log subscription operation binding the contract event 0xe58e1506e3b37e68a149230d64d9b29a5234ca9df751c07243bf296a5d046687.
//
// Solidity: event E_Call_Failed(bytes32 arg0, address arg1)
func (_Asset *AssetFilterer) WatchECallFailed(opts *bind.WatchOpts, sink chan<- *AssetECallFailed) (event.Subscription, error) {

	logs, sub, err := _Asset.contract.WatchLogs(opts, "E_Call_Failed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetECallFailed)
				if err := _Asset.contract.UnpackLog(event, "E_Call_Failed", log); err != nil {
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

// ParseECallFailed is a log parse operation binding the contract event 0xe58e1506e3b37e68a149230d64d9b29a5234ca9df751c07243bf296a5d046687.
//
// Solidity: event E_Call_Failed(bytes32 arg0, address arg1)
func (_Asset *AssetFilterer) ParseECallFailed(log types.Log) (*AssetECallFailed, error) {
	event := new(AssetECallFailed)
	if err := _Asset.contract.UnpackLog(event, "E_Call_Failed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssetECallSuccessIterator is returned from FilterECallSuccess and is used to iterate over the raw logs and unpacked data for ECallSuccess events raised by the Asset contract.
type AssetECallSuccessIterator struct {
	Event *AssetECallSuccess // Event containing the contract specifics and raw log

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
func (it *AssetECallSuccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetECallSuccess)
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
		it.Event = new(AssetECallSuccess)
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
func (it *AssetECallSuccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetECallSuccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetECallSuccess represents a ECallSuccess event raised by the Asset contract.
type AssetECallSuccess struct {
	Arg0 [32]byte
	Arg1 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterECallSuccess is a free log retrieval operation binding the contract event 0xe47c82e669f0997954d8f2f085371a24b7c10674062e5620daa5aca48a2b1cc8.
//
// Solidity: event E_Call_Success(bytes32 arg0, address arg1)
func (_Asset *AssetFilterer) FilterECallSuccess(opts *bind.FilterOpts) (*AssetECallSuccessIterator, error) {

	logs, sub, err := _Asset.contract.FilterLogs(opts, "E_Call_Success")
	if err != nil {
		return nil, err
	}
	return &AssetECallSuccessIterator{contract: _Asset.contract, event: "E_Call_Success", logs: logs, sub: sub}, nil
}

// WatchECallSuccess is a free log subscription operation binding the contract event 0xe47c82e669f0997954d8f2f085371a24b7c10674062e5620daa5aca48a2b1cc8.
//
// Solidity: event E_Call_Success(bytes32 arg0, address arg1)
func (_Asset *AssetFilterer) WatchECallSuccess(opts *bind.WatchOpts, sink chan<- *AssetECallSuccess) (event.Subscription, error) {

	logs, sub, err := _Asset.contract.WatchLogs(opts, "E_Call_Success")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetECallSuccess)
				if err := _Asset.contract.UnpackLog(event, "E_Call_Success", log); err != nil {
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

// ParseECallSuccess is a log parse operation binding the contract event 0xe47c82e669f0997954d8f2f085371a24b7c10674062e5620daa5aca48a2b1cc8.
//
// Solidity: event E_Call_Success(bytes32 arg0, address arg1)
func (_Asset *AssetFilterer) ParseECallSuccess(log types.Log) (*AssetECallSuccess, error) {
	event := new(AssetECallSuccess)
	if err := _Asset.contract.UnpackLog(event, "E_Call_Success", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssetETransferNativeIterator is returned from FilterETransferNative and is used to iterate over the raw logs and unpacked data for ETransferNative events raised by the Asset contract.
type AssetETransferNativeIterator struct {
	Event *AssetETransferNative // Event containing the contract specifics and raw log

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
func (it *AssetETransferNativeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetETransferNative)
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
		it.Event = new(AssetETransferNative)
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
func (it *AssetETransferNativeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetETransferNativeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetETransferNative represents a ETransferNative event raised by the Asset contract.
type AssetETransferNative struct {
	Arg0 [32]byte
	Arg1 common.Address
	Arg2 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterETransferNative is a free log retrieval operation binding the contract event 0x25a1016e9f61fd45630e7222f6b92e1314493b84ab59c1e4ad4aee7be93c76a0.
//
// Solidity: event E_Transfer_Native(bytes32 arg0, address arg1, uint256 arg2)
func (_Asset *AssetFilterer) FilterETransferNative(opts *bind.FilterOpts) (*AssetETransferNativeIterator, error) {

	logs, sub, err := _Asset.contract.FilterLogs(opts, "E_Transfer_Native")
	if err != nil {
		return nil, err
	}
	return &AssetETransferNativeIterator{contract: _Asset.contract, event: "E_Transfer_Native", logs: logs, sub: sub}, nil
}

// WatchETransferNative is a free log subscription operation binding the contract event 0x25a1016e9f61fd45630e7222f6b92e1314493b84ab59c1e4ad4aee7be93c76a0.
//
// Solidity: event E_Transfer_Native(bytes32 arg0, address arg1, uint256 arg2)
func (_Asset *AssetFilterer) WatchETransferNative(opts *bind.WatchOpts, sink chan<- *AssetETransferNative) (event.Subscription, error) {

	logs, sub, err := _Asset.contract.WatchLogs(opts, "E_Transfer_Native")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetETransferNative)
				if err := _Asset.contract.UnpackLog(event, "E_Transfer_Native", log); err != nil {
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

// ParseETransferNative is a log parse operation binding the contract event 0x25a1016e9f61fd45630e7222f6b92e1314493b84ab59c1e4ad4aee7be93c76a0.
//
// Solidity: event E_Transfer_Native(bytes32 arg0, address arg1, uint256 arg2)
func (_Asset *AssetFilterer) ParseETransferNative(log types.Log) (*AssetETransferNative, error) {
	event := new(AssetETransferNative)
	if err := _Asset.contract.UnpackLog(event, "E_Transfer_Native", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssetETransferTokenIterator is returned from FilterETransferToken and is used to iterate over the raw logs and unpacked data for ETransferToken events raised by the Asset contract.
type AssetETransferTokenIterator struct {
	Event *AssetETransferToken // Event containing the contract specifics and raw log

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
func (it *AssetETransferTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetETransferToken)
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
		it.Event = new(AssetETransferToken)
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
func (it *AssetETransferTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetETransferTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetETransferToken represents a ETransferToken event raised by the Asset contract.
type AssetETransferToken struct {
	Arg0 [32]byte
	Arg1 common.Address
	Arg2 common.Address
	Arg3 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterETransferToken is a free log retrieval operation binding the contract event 0x085dd04177f537dbaac184ad498c023fa70a8975916c0af128247bf50f57804d.
//
// Solidity: event E_Transfer_Token(bytes32 arg0, address arg1, address arg2, uint256 arg3)
func (_Asset *AssetFilterer) FilterETransferToken(opts *bind.FilterOpts) (*AssetETransferTokenIterator, error) {

	logs, sub, err := _Asset.contract.FilterLogs(opts, "E_Transfer_Token")
	if err != nil {
		return nil, err
	}
	return &AssetETransferTokenIterator{contract: _Asset.contract, event: "E_Transfer_Token", logs: logs, sub: sub}, nil
}

// WatchETransferToken is a free log subscription operation binding the contract event 0x085dd04177f537dbaac184ad498c023fa70a8975916c0af128247bf50f57804d.
//
// Solidity: event E_Transfer_Token(bytes32 arg0, address arg1, address arg2, uint256 arg3)
func (_Asset *AssetFilterer) WatchETransferToken(opts *bind.WatchOpts, sink chan<- *AssetETransferToken) (event.Subscription, error) {

	logs, sub, err := _Asset.contract.WatchLogs(opts, "E_Transfer_Token")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetETransferToken)
				if err := _Asset.contract.UnpackLog(event, "E_Transfer_Token", log); err != nil {
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

// ParseETransferToken is a log parse operation binding the contract event 0x085dd04177f537dbaac184ad498c023fa70a8975916c0af128247bf50f57804d.
//
// Solidity: event E_Transfer_Token(bytes32 arg0, address arg1, address arg2, uint256 arg3)
func (_Asset *AssetFilterer) ParseETransferToken(log types.Log) (*AssetETransferToken, error) {
	event := new(AssetETransferToken)
	if err := _Asset.contract.UnpackLog(event, "E_Transfer_Token", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

