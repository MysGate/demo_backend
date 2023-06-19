// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// CrossControllerOrder is an auto generated low-level Go binding around an user-defined struct.
type CrossControllerOrder struct {
	OrderId     *big.Int
	SrcChainId  *big.Int
	SrcAddress  common.Address
	SrcToken    common.Address
	SrcAmount   *big.Int
	DestChainId *big.Int
	DestAddress common.Address
	DestToken   common.Address
	Porter      common.Address
}

// CrossControllerProof is an auto generated low-level Go binding around an user-defined struct.
type CrossControllerProof struct {
	A [2]*big.Int
	B [2][2]*big.Int
	C [2]*big.Int
}

// CrossControllerReceipt is an auto generated low-level Go binding around an user-defined struct.
type CrossControllerReceipt struct {
	ProofHash  [32]byte
	DestTxHash [32]byte
}

// CrossMetaData contains all meta data concerning the Cross contract.
var CrossMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"proofHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"destTxHash\",\"type\":\"bytes32\"}],\"indexed\":false,\"internalType\":\"structCrossController.Receipt\",\"name\":\"receipt\",\"type\":\"tuple\"}],\"name\":\"CommitReceipt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"proofHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"destTxHash\",\"type\":\"bytes32\"}],\"indexed\":false,\"internalType\":\"structCrossController.Receipt\",\"name\":\"receipt\",\"type\":\"tuple\"}],\"name\":\"CommitReceiptWithZk\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"srcAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"porter\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structCrossController.Order\",\"name\":\"order\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"srcTokenDecimals\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"crossAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"paidAmount\",\"type\":\"uint256\"}],\"name\":\"CrossFrom\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"srcAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"porter\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structCrossController.Order\",\"name\":\"order\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fixedFeeAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"floatFeeAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"crossAmount\",\"type\":\"uint256\"}],\"name\":\"CrossTo\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_enable\",\"type\":\"bool\"}],\"name\":\"EnabledZkVerifier\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"floatFee\",\"type\":\"uint256\"}],\"name\":\"SettedFloatFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"}],\"name\":\"SettedPorterFactory\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_zkVerifier\",\"type\":\"address\"}],\"name\":\"SettedZkVerifier\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_commitment\",\"type\":\"uint256\"}],\"name\":\"addCommitment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"proofHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"destTxHash\",\"type\":\"bytes32\"}],\"internalType\":\"structCrossController.Receipt\",\"name\":\"receipt\",\"type\":\"tuple\"}],\"name\":\"commitReceipt\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"a\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"b\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"c\",\"type\":\"uint256[2]\"}],\"internalType\":\"structCrossController.Proof\",\"name\":\"proof\",\"type\":\"tuple\"},{\"internalType\":\"uint256[2]\",\"name\":\"input\",\"type\":\"uint256[2]\"},{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"destTxHash\",\"type\":\"bytes32\"}],\"name\":\"commitReceiptWithZK\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"srcAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"porter\",\"type\":\"address\"}],\"internalType\":\"structCrossController.Order\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"srcTokenDecimals\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"crossAmount\",\"type\":\"uint256\"}],\"name\":\"crossFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"srcAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"porter\",\"type\":\"address\"}],\"internalType\":\"structCrossController.Order\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"crossTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentChainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_enable\",\"type\":\"bool\"}],\"name\":\"enableZkVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"floatFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"orderHashes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"orders\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"srcAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"porter\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"paidOrders\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"pendingOrders\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"porterFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"receipts\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"proofHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"destTxHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_floatFee\",\"type\":\"uint256\"}],\"name\":\"setFloatFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"}],\"name\":\"setPorterFactory\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_verifier\",\"type\":\"address\"}],\"name\":\"setZkVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"zkVerifier\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// CrossABI is the input ABI used to generate the binding from.
// Deprecated: Use CrossMetaData.ABI instead.
var CrossABI = CrossMetaData.ABI

// Cross is an auto generated Go binding around an Ethereum contract.
type Cross struct {
	CrossCaller     // Read-only binding to the contract
	CrossTransactor // Write-only binding to the contract
	CrossFilterer   // Log filterer for contract events
}

// CrossCaller is an auto generated read-only Go binding around an Ethereum contract.
type CrossCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CrossTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CrossFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CrossSession struct {
	Contract     *Cross            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CrossCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CrossCallerSession struct {
	Contract *CrossCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// CrossTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CrossTransactorSession struct {
	Contract     *CrossTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CrossRaw is an auto generated low-level Go binding around an Ethereum contract.
type CrossRaw struct {
	Contract *Cross // Generic contract binding to access the raw methods on
}

// CrossCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CrossCallerRaw struct {
	Contract *CrossCaller // Generic read-only contract binding to access the raw methods on
}

// CrossTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CrossTransactorRaw struct {
	Contract *CrossTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCross creates a new instance of Cross, bound to a specific deployed contract.
func NewCross(address common.Address, backend bind.ContractBackend) (*Cross, error) {
	contract, err := bindCross(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Cross{CrossCaller: CrossCaller{contract: contract}, CrossTransactor: CrossTransactor{contract: contract}, CrossFilterer: CrossFilterer{contract: contract}}, nil
}

// NewCrossCaller creates a new read-only instance of Cross, bound to a specific deployed contract.
func NewCrossCaller(address common.Address, caller bind.ContractCaller) (*CrossCaller, error) {
	contract, err := bindCross(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CrossCaller{contract: contract}, nil
}

// NewCrossTransactor creates a new write-only instance of Cross, bound to a specific deployed contract.
func NewCrossTransactor(address common.Address, transactor bind.ContractTransactor) (*CrossTransactor, error) {
	contract, err := bindCross(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CrossTransactor{contract: contract}, nil
}

// NewCrossFilterer creates a new log filterer instance of Cross, bound to a specific deployed contract.
func NewCrossFilterer(address common.Address, filterer bind.ContractFilterer) (*CrossFilterer, error) {
	contract, err := bindCross(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CrossFilterer{contract: contract}, nil
}

// bindCross binds a generic wrapper to an already deployed contract.
func bindCross(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CrossMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Cross *CrossRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Cross.Contract.CrossCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Cross *CrossRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cross.Contract.CrossTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Cross *CrossRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Cross.Contract.CrossTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Cross *CrossCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Cross.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Cross *CrossTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cross.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Cross *CrossTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Cross.Contract.contract.Transact(opts, method, params...)
}

// CurrentChainId is a free data retrieval call binding the contract method 0x6cbadbfa.
//
// Solidity: function currentChainId() view returns(uint256)
func (_Cross *CrossCaller) CurrentChainId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cross.contract.Call(opts, &out, "currentChainId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentChainId is a free data retrieval call binding the contract method 0x6cbadbfa.
//
// Solidity: function currentChainId() view returns(uint256)
func (_Cross *CrossSession) CurrentChainId() (*big.Int, error) {
	return _Cross.Contract.CurrentChainId(&_Cross.CallOpts)
}

// CurrentChainId is a free data retrieval call binding the contract method 0x6cbadbfa.
//
// Solidity: function currentChainId() view returns(uint256)
func (_Cross *CrossCallerSession) CurrentChainId() (*big.Int, error) {
	return _Cross.Contract.CurrentChainId(&_Cross.CallOpts)
}

// Enable is a free data retrieval call binding the contract method 0xa3907d71.
//
// Solidity: function enable() view returns(bool)
func (_Cross *CrossCaller) Enable(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Cross.contract.Call(opts, &out, "enable")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Enable is a free data retrieval call binding the contract method 0xa3907d71.
//
// Solidity: function enable() view returns(bool)
func (_Cross *CrossSession) Enable() (bool, error) {
	return _Cross.Contract.Enable(&_Cross.CallOpts)
}

// Enable is a free data retrieval call binding the contract method 0xa3907d71.
//
// Solidity: function enable() view returns(bool)
func (_Cross *CrossCallerSession) Enable() (bool, error) {
	return _Cross.Contract.Enable(&_Cross.CallOpts)
}

// FloatFee is a free data retrieval call binding the contract method 0x7effecea.
//
// Solidity: function floatFee() view returns(uint256)
func (_Cross *CrossCaller) FloatFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cross.contract.Call(opts, &out, "floatFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FloatFee is a free data retrieval call binding the contract method 0x7effecea.
//
// Solidity: function floatFee() view returns(uint256)
func (_Cross *CrossSession) FloatFee() (*big.Int, error) {
	return _Cross.Contract.FloatFee(&_Cross.CallOpts)
}

// FloatFee is a free data retrieval call binding the contract method 0x7effecea.
//
// Solidity: function floatFee() view returns(uint256)
func (_Cross *CrossCallerSession) FloatFee() (*big.Int, error) {
	return _Cross.Contract.FloatFee(&_Cross.CallOpts)
}

// OrderHashes is a free data retrieval call binding the contract method 0xdf69558c.
//
// Solidity: function orderHashes(uint256 ) view returns(bytes32)
func (_Cross *CrossCaller) OrderHashes(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Cross.contract.Call(opts, &out, "orderHashes", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OrderHashes is a free data retrieval call binding the contract method 0xdf69558c.
//
// Solidity: function orderHashes(uint256 ) view returns(bytes32)
func (_Cross *CrossSession) OrderHashes(arg0 *big.Int) ([32]byte, error) {
	return _Cross.Contract.OrderHashes(&_Cross.CallOpts, arg0)
}

// OrderHashes is a free data retrieval call binding the contract method 0xdf69558c.
//
// Solidity: function orderHashes(uint256 ) view returns(bytes32)
func (_Cross *CrossCallerSession) OrderHashes(arg0 *big.Int) ([32]byte, error) {
	return _Cross.Contract.OrderHashes(&_Cross.CallOpts, arg0)
}

// Orders is a free data retrieval call binding the contract method 0x9c3f1e90.
//
// Solidity: function orders(bytes32 ) view returns(uint256 orderId, uint256 srcChainId, address srcAddress, address srcToken, uint256 srcAmount, uint256 destChainId, address destAddress, address destToken, address porter)
func (_Cross *CrossCaller) Orders(opts *bind.CallOpts, arg0 [32]byte) (struct {
	OrderId     *big.Int
	SrcChainId  *big.Int
	SrcAddress  common.Address
	SrcToken    common.Address
	SrcAmount   *big.Int
	DestChainId *big.Int
	DestAddress common.Address
	DestToken   common.Address
	Porter      common.Address
}, error) {
	var out []interface{}
	err := _Cross.contract.Call(opts, &out, "orders", arg0)

	outstruct := new(struct {
		OrderId     *big.Int
		SrcChainId  *big.Int
		SrcAddress  common.Address
		SrcToken    common.Address
		SrcAmount   *big.Int
		DestChainId *big.Int
		DestAddress common.Address
		DestToken   common.Address
		Porter      common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.OrderId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.SrcChainId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.SrcAddress = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.SrcToken = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.SrcAmount = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.DestChainId = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.DestAddress = *abi.ConvertType(out[6], new(common.Address)).(*common.Address)
	outstruct.DestToken = *abi.ConvertType(out[7], new(common.Address)).(*common.Address)
	outstruct.Porter = *abi.ConvertType(out[8], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Orders is a free data retrieval call binding the contract method 0x9c3f1e90.
//
// Solidity: function orders(bytes32 ) view returns(uint256 orderId, uint256 srcChainId, address srcAddress, address srcToken, uint256 srcAmount, uint256 destChainId, address destAddress, address destToken, address porter)
func (_Cross *CrossSession) Orders(arg0 [32]byte) (struct {
	OrderId     *big.Int
	SrcChainId  *big.Int
	SrcAddress  common.Address
	SrcToken    common.Address
	SrcAmount   *big.Int
	DestChainId *big.Int
	DestAddress common.Address
	DestToken   common.Address
	Porter      common.Address
}, error) {
	return _Cross.Contract.Orders(&_Cross.CallOpts, arg0)
}

// Orders is a free data retrieval call binding the contract method 0x9c3f1e90.
//
// Solidity: function orders(bytes32 ) view returns(uint256 orderId, uint256 srcChainId, address srcAddress, address srcToken, uint256 srcAmount, uint256 destChainId, address destAddress, address destToken, address porter)
func (_Cross *CrossCallerSession) Orders(arg0 [32]byte) (struct {
	OrderId     *big.Int
	SrcChainId  *big.Int
	SrcAddress  common.Address
	SrcToken    common.Address
	SrcAmount   *big.Int
	DestChainId *big.Int
	DestAddress common.Address
	DestToken   common.Address
	Porter      common.Address
}, error) {
	return _Cross.Contract.Orders(&_Cross.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Cross *CrossCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Cross.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Cross *CrossSession) Owner() (common.Address, error) {
	return _Cross.Contract.Owner(&_Cross.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Cross *CrossCallerSession) Owner() (common.Address, error) {
	return _Cross.Contract.Owner(&_Cross.CallOpts)
}

// PaidOrders is a free data retrieval call binding the contract method 0x39c51fea.
//
// Solidity: function paidOrders(bytes32 ) view returns(bool)
func (_Cross *CrossCaller) PaidOrders(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Cross.contract.Call(opts, &out, "paidOrders", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PaidOrders is a free data retrieval call binding the contract method 0x39c51fea.
//
// Solidity: function paidOrders(bytes32 ) view returns(bool)
func (_Cross *CrossSession) PaidOrders(arg0 [32]byte) (bool, error) {
	return _Cross.Contract.PaidOrders(&_Cross.CallOpts, arg0)
}

// PaidOrders is a free data retrieval call binding the contract method 0x39c51fea.
//
// Solidity: function paidOrders(bytes32 ) view returns(bool)
func (_Cross *CrossCallerSession) PaidOrders(arg0 [32]byte) (bool, error) {
	return _Cross.Contract.PaidOrders(&_Cross.CallOpts, arg0)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Cross *CrossCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Cross.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Cross *CrossSession) Paused() (bool, error) {
	return _Cross.Contract.Paused(&_Cross.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Cross *CrossCallerSession) Paused() (bool, error) {
	return _Cross.Contract.Paused(&_Cross.CallOpts)
}

// PendingOrders is a free data retrieval call binding the contract method 0x7b55766f.
//
// Solidity: function pendingOrders(bytes32 ) view returns(bool)
func (_Cross *CrossCaller) PendingOrders(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Cross.contract.Call(opts, &out, "pendingOrders", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PendingOrders is a free data retrieval call binding the contract method 0x7b55766f.
//
// Solidity: function pendingOrders(bytes32 ) view returns(bool)
func (_Cross *CrossSession) PendingOrders(arg0 [32]byte) (bool, error) {
	return _Cross.Contract.PendingOrders(&_Cross.CallOpts, arg0)
}

// PendingOrders is a free data retrieval call binding the contract method 0x7b55766f.
//
// Solidity: function pendingOrders(bytes32 ) view returns(bool)
func (_Cross *CrossCallerSession) PendingOrders(arg0 [32]byte) (bool, error) {
	return _Cross.Contract.PendingOrders(&_Cross.CallOpts, arg0)
}

// PorterFactory is a free data retrieval call binding the contract method 0x2f803f81.
//
// Solidity: function porterFactory() view returns(address)
func (_Cross *CrossCaller) PorterFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Cross.contract.Call(opts, &out, "porterFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PorterFactory is a free data retrieval call binding the contract method 0x2f803f81.
//
// Solidity: function porterFactory() view returns(address)
func (_Cross *CrossSession) PorterFactory() (common.Address, error) {
	return _Cross.Contract.PorterFactory(&_Cross.CallOpts)
}

// PorterFactory is a free data retrieval call binding the contract method 0x2f803f81.
//
// Solidity: function porterFactory() view returns(address)
func (_Cross *CrossCallerSession) PorterFactory() (common.Address, error) {
	return _Cross.Contract.PorterFactory(&_Cross.CallOpts)
}

// Receipts is a free data retrieval call binding the contract method 0xef6cf04d.
//
// Solidity: function receipts(bytes32 ) view returns(bytes32 proofHash, bytes32 destTxHash)
func (_Cross *CrossCaller) Receipts(opts *bind.CallOpts, arg0 [32]byte) (struct {
	ProofHash  [32]byte
	DestTxHash [32]byte
}, error) {
	var out []interface{}
	err := _Cross.contract.Call(opts, &out, "receipts", arg0)

	outstruct := new(struct {
		ProofHash  [32]byte
		DestTxHash [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ProofHash = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.DestTxHash = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// Receipts is a free data retrieval call binding the contract method 0xef6cf04d.
//
// Solidity: function receipts(bytes32 ) view returns(bytes32 proofHash, bytes32 destTxHash)
func (_Cross *CrossSession) Receipts(arg0 [32]byte) (struct {
	ProofHash  [32]byte
	DestTxHash [32]byte
}, error) {
	return _Cross.Contract.Receipts(&_Cross.CallOpts, arg0)
}

// Receipts is a free data retrieval call binding the contract method 0xef6cf04d.
//
// Solidity: function receipts(bytes32 ) view returns(bytes32 proofHash, bytes32 destTxHash)
func (_Cross *CrossCallerSession) Receipts(arg0 [32]byte) (struct {
	ProofHash  [32]byte
	DestTxHash [32]byte
}, error) {
	return _Cross.Contract.Receipts(&_Cross.CallOpts, arg0)
}

// ZkVerifier is a free data retrieval call binding the contract method 0xd6df096d.
//
// Solidity: function zkVerifier() view returns(address)
func (_Cross *CrossCaller) ZkVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Cross.contract.Call(opts, &out, "zkVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ZkVerifier is a free data retrieval call binding the contract method 0xd6df096d.
//
// Solidity: function zkVerifier() view returns(address)
func (_Cross *CrossSession) ZkVerifier() (common.Address, error) {
	return _Cross.Contract.ZkVerifier(&_Cross.CallOpts)
}

// ZkVerifier is a free data retrieval call binding the contract method 0xd6df096d.
//
// Solidity: function zkVerifier() view returns(address)
func (_Cross *CrossCallerSession) ZkVerifier() (common.Address, error) {
	return _Cross.Contract.ZkVerifier(&_Cross.CallOpts)
}

// AddCommitment is a paid mutator transaction binding the contract method 0x797a6890.
//
// Solidity: function addCommitment(uint256 _commitment) returns()
func (_Cross *CrossTransactor) AddCommitment(opts *bind.TransactOpts, _commitment *big.Int) (*types.Transaction, error) {
	return _Cross.contract.Transact(opts, "addCommitment", _commitment)
}

// AddCommitment is a paid mutator transaction binding the contract method 0x797a6890.
//
// Solidity: function addCommitment(uint256 _commitment) returns()
func (_Cross *CrossSession) AddCommitment(_commitment *big.Int) (*types.Transaction, error) {
	return _Cross.Contract.AddCommitment(&_Cross.TransactOpts, _commitment)
}

// AddCommitment is a paid mutator transaction binding the contract method 0x797a6890.
//
// Solidity: function addCommitment(uint256 _commitment) returns()
func (_Cross *CrossTransactorSession) AddCommitment(_commitment *big.Int) (*types.Transaction, error) {
	return _Cross.Contract.AddCommitment(&_Cross.TransactOpts, _commitment)
}

// CommitReceipt is a paid mutator transaction binding the contract method 0x2d75219c.
//
// Solidity: function commitReceipt(bytes32 orderHash, (bytes32,bytes32) receipt) returns()
func (_Cross *CrossTransactor) CommitReceipt(opts *bind.TransactOpts, orderHash [32]byte, receipt CrossControllerReceipt) (*types.Transaction, error) {
	return _Cross.contract.Transact(opts, "commitReceipt", orderHash, receipt)
}

// CommitReceipt is a paid mutator transaction binding the contract method 0x2d75219c.
//
// Solidity: function commitReceipt(bytes32 orderHash, (bytes32,bytes32) receipt) returns()
func (_Cross *CrossSession) CommitReceipt(orderHash [32]byte, receipt CrossControllerReceipt) (*types.Transaction, error) {
	return _Cross.Contract.CommitReceipt(&_Cross.TransactOpts, orderHash, receipt)
}

// CommitReceipt is a paid mutator transaction binding the contract method 0x2d75219c.
//
// Solidity: function commitReceipt(bytes32 orderHash, (bytes32,bytes32) receipt) returns()
func (_Cross *CrossTransactorSession) CommitReceipt(orderHash [32]byte, receipt CrossControllerReceipt) (*types.Transaction, error) {
	return _Cross.Contract.CommitReceipt(&_Cross.TransactOpts, orderHash, receipt)
}

// CommitReceiptWithZK is a paid mutator transaction binding the contract method 0x102c955b.
//
// Solidity: function commitReceiptWithZK((uint256[2],uint256[2][2],uint256[2]) proof, uint256[2] input, bytes32 orderHash, bytes32 destTxHash) returns()
func (_Cross *CrossTransactor) CommitReceiptWithZK(opts *bind.TransactOpts, proof CrossControllerProof, input [2]*big.Int, orderHash [32]byte, destTxHash [32]byte) (*types.Transaction, error) {
	return _Cross.contract.Transact(opts, "commitReceiptWithZK", proof, input, orderHash, destTxHash)
}

// CommitReceiptWithZK is a paid mutator transaction binding the contract method 0x102c955b.
//
// Solidity: function commitReceiptWithZK((uint256[2],uint256[2][2],uint256[2]) proof, uint256[2] input, bytes32 orderHash, bytes32 destTxHash) returns()
func (_Cross *CrossSession) CommitReceiptWithZK(proof CrossControllerProof, input [2]*big.Int, orderHash [32]byte, destTxHash [32]byte) (*types.Transaction, error) {
	return _Cross.Contract.CommitReceiptWithZK(&_Cross.TransactOpts, proof, input, orderHash, destTxHash)
}

// CommitReceiptWithZK is a paid mutator transaction binding the contract method 0x102c955b.
//
// Solidity: function commitReceiptWithZK((uint256[2],uint256[2][2],uint256[2]) proof, uint256[2] input, bytes32 orderHash, bytes32 destTxHash) returns()
func (_Cross *CrossTransactorSession) CommitReceiptWithZK(proof CrossControllerProof, input [2]*big.Int, orderHash [32]byte, destTxHash [32]byte) (*types.Transaction, error) {
	return _Cross.Contract.CommitReceiptWithZK(&_Cross.TransactOpts, proof, input, orderHash, destTxHash)
}

// CrossFrom is a paid mutator transaction binding the contract method 0x88dd09f6.
//
// Solidity: function crossFrom((uint256,uint256,address,address,uint256,uint256,address,address,address) order, uint8 srcTokenDecimals, uint256 crossAmount) returns()
func (_Cross *CrossTransactor) CrossFrom(opts *bind.TransactOpts, order CrossControllerOrder, srcTokenDecimals uint8, crossAmount *big.Int) (*types.Transaction, error) {
	return _Cross.contract.Transact(opts, "crossFrom", order, srcTokenDecimals, crossAmount)
}

// CrossFrom is a paid mutator transaction binding the contract method 0x88dd09f6.
//
// Solidity: function crossFrom((uint256,uint256,address,address,uint256,uint256,address,address,address) order, uint8 srcTokenDecimals, uint256 crossAmount) returns()
func (_Cross *CrossSession) CrossFrom(order CrossControllerOrder, srcTokenDecimals uint8, crossAmount *big.Int) (*types.Transaction, error) {
	return _Cross.Contract.CrossFrom(&_Cross.TransactOpts, order, srcTokenDecimals, crossAmount)
}

// CrossFrom is a paid mutator transaction binding the contract method 0x88dd09f6.
//
// Solidity: function crossFrom((uint256,uint256,address,address,uint256,uint256,address,address,address) order, uint8 srcTokenDecimals, uint256 crossAmount) returns()
func (_Cross *CrossTransactorSession) CrossFrom(order CrossControllerOrder, srcTokenDecimals uint8, crossAmount *big.Int) (*types.Transaction, error) {
	return _Cross.Contract.CrossFrom(&_Cross.TransactOpts, order, srcTokenDecimals, crossAmount)
}

// CrossTo is a paid mutator transaction binding the contract method 0xdef66322.
//
// Solidity: function crossTo((uint256,uint256,address,address,uint256,uint256,address,address,address) order) returns()
func (_Cross *CrossTransactor) CrossTo(opts *bind.TransactOpts, order CrossControllerOrder) (*types.Transaction, error) {
	return _Cross.contract.Transact(opts, "crossTo", order)
}

// CrossTo is a paid mutator transaction binding the contract method 0xdef66322.
//
// Solidity: function crossTo((uint256,uint256,address,address,uint256,uint256,address,address,address) order) returns()
func (_Cross *CrossSession) CrossTo(order CrossControllerOrder) (*types.Transaction, error) {
	return _Cross.Contract.CrossTo(&_Cross.TransactOpts, order)
}

// CrossTo is a paid mutator transaction binding the contract method 0xdef66322.
//
// Solidity: function crossTo((uint256,uint256,address,address,uint256,uint256,address,address,address) order) returns()
func (_Cross *CrossTransactorSession) CrossTo(order CrossControllerOrder) (*types.Transaction, error) {
	return _Cross.Contract.CrossTo(&_Cross.TransactOpts, order)
}

// EnableZkVerifier is a paid mutator transaction binding the contract method 0x4056e986.
//
// Solidity: function enableZkVerifier(bool _enable) returns()
func (_Cross *CrossTransactor) EnableZkVerifier(opts *bind.TransactOpts, _enable bool) (*types.Transaction, error) {
	return _Cross.contract.Transact(opts, "enableZkVerifier", _enable)
}

// EnableZkVerifier is a paid mutator transaction binding the contract method 0x4056e986.
//
// Solidity: function enableZkVerifier(bool _enable) returns()
func (_Cross *CrossSession) EnableZkVerifier(_enable bool) (*types.Transaction, error) {
	return _Cross.Contract.EnableZkVerifier(&_Cross.TransactOpts, _enable)
}

// EnableZkVerifier is a paid mutator transaction binding the contract method 0x4056e986.
//
// Solidity: function enableZkVerifier(bool _enable) returns()
func (_Cross *CrossTransactorSession) EnableZkVerifier(_enable bool) (*types.Transaction, error) {
	return _Cross.Contract.EnableZkVerifier(&_Cross.TransactOpts, _enable)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _owner) returns()
func (_Cross *CrossTransactor) Initialize(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _Cross.contract.Transact(opts, "initialize", _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _owner) returns()
func (_Cross *CrossSession) Initialize(_owner common.Address) (*types.Transaction, error) {
	return _Cross.Contract.Initialize(&_Cross.TransactOpts, _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _owner) returns()
func (_Cross *CrossTransactorSession) Initialize(_owner common.Address) (*types.Transaction, error) {
	return _Cross.Contract.Initialize(&_Cross.TransactOpts, _owner)
}

// SetFloatFee is a paid mutator transaction binding the contract method 0xcb23c2a7.
//
// Solidity: function setFloatFee(uint256 _floatFee) returns()
func (_Cross *CrossTransactor) SetFloatFee(opts *bind.TransactOpts, _floatFee *big.Int) (*types.Transaction, error) {
	return _Cross.contract.Transact(opts, "setFloatFee", _floatFee)
}

// SetFloatFee is a paid mutator transaction binding the contract method 0xcb23c2a7.
//
// Solidity: function setFloatFee(uint256 _floatFee) returns()
func (_Cross *CrossSession) SetFloatFee(_floatFee *big.Int) (*types.Transaction, error) {
	return _Cross.Contract.SetFloatFee(&_Cross.TransactOpts, _floatFee)
}

// SetFloatFee is a paid mutator transaction binding the contract method 0xcb23c2a7.
//
// Solidity: function setFloatFee(uint256 _floatFee) returns()
func (_Cross *CrossTransactorSession) SetFloatFee(_floatFee *big.Int) (*types.Transaction, error) {
	return _Cross.Contract.SetFloatFee(&_Cross.TransactOpts, _floatFee)
}

// SetPorterFactory is a paid mutator transaction binding the contract method 0x5e553013.
//
// Solidity: function setPorterFactory(address _factory) returns()
func (_Cross *CrossTransactor) SetPorterFactory(opts *bind.TransactOpts, _factory common.Address) (*types.Transaction, error) {
	return _Cross.contract.Transact(opts, "setPorterFactory", _factory)
}

// SetPorterFactory is a paid mutator transaction binding the contract method 0x5e553013.
//
// Solidity: function setPorterFactory(address _factory) returns()
func (_Cross *CrossSession) SetPorterFactory(_factory common.Address) (*types.Transaction, error) {
	return _Cross.Contract.SetPorterFactory(&_Cross.TransactOpts, _factory)
}

// SetPorterFactory is a paid mutator transaction binding the contract method 0x5e553013.
//
// Solidity: function setPorterFactory(address _factory) returns()
func (_Cross *CrossTransactorSession) SetPorterFactory(_factory common.Address) (*types.Transaction, error) {
	return _Cross.Contract.SetPorterFactory(&_Cross.TransactOpts, _factory)
}

// SetZkVerifier is a paid mutator transaction binding the contract method 0x96913528.
//
// Solidity: function setZkVerifier(address _verifier) returns()
func (_Cross *CrossTransactor) SetZkVerifier(opts *bind.TransactOpts, _verifier common.Address) (*types.Transaction, error) {
	return _Cross.contract.Transact(opts, "setZkVerifier", _verifier)
}

// SetZkVerifier is a paid mutator transaction binding the contract method 0x96913528.
//
// Solidity: function setZkVerifier(address _verifier) returns()
func (_Cross *CrossSession) SetZkVerifier(_verifier common.Address) (*types.Transaction, error) {
	return _Cross.Contract.SetZkVerifier(&_Cross.TransactOpts, _verifier)
}

// SetZkVerifier is a paid mutator transaction binding the contract method 0x96913528.
//
// Solidity: function setZkVerifier(address _verifier) returns()
func (_Cross *CrossTransactorSession) SetZkVerifier(_verifier common.Address) (*types.Transaction, error) {
	return _Cross.Contract.SetZkVerifier(&_Cross.TransactOpts, _verifier)
}

// CrossCommitReceiptIterator is returned from FilterCommitReceipt and is used to iterate over the raw logs and unpacked data for CommitReceipt events raised by the Cross contract.
type CrossCommitReceiptIterator struct {
	Event *CrossCommitReceipt // Event containing the contract specifics and raw log

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
func (it *CrossCommitReceiptIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossCommitReceipt)
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
		it.Event = new(CrossCommitReceipt)
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
func (it *CrossCommitReceiptIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossCommitReceiptIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossCommitReceipt represents a CommitReceipt event raised by the Cross contract.
type CrossCommitReceipt struct {
	Validator common.Address
	OrderHash [32]byte
	Receipt   CrossControllerReceipt
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCommitReceipt is a free log retrieval operation binding the contract event 0x581db44feed8ab7f2b0e591fd633c1326a4ba3ea20a5c346ab38fd1f42208e81.
//
// Solidity: event CommitReceipt(address indexed validator, bytes32 indexed orderHash, (bytes32,bytes32) receipt)
func (_Cross *CrossFilterer) FilterCommitReceipt(opts *bind.FilterOpts, validator []common.Address, orderHash [][32]byte) (*CrossCommitReceiptIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}

	logs, sub, err := _Cross.contract.FilterLogs(opts, "CommitReceipt", validatorRule, orderHashRule)
	if err != nil {
		return nil, err
	}
	return &CrossCommitReceiptIterator{contract: _Cross.contract, event: "CommitReceipt", logs: logs, sub: sub}, nil
}

// WatchCommitReceipt is a free log subscription operation binding the contract event 0x581db44feed8ab7f2b0e591fd633c1326a4ba3ea20a5c346ab38fd1f42208e81.
//
// Solidity: event CommitReceipt(address indexed validator, bytes32 indexed orderHash, (bytes32,bytes32) receipt)
func (_Cross *CrossFilterer) WatchCommitReceipt(opts *bind.WatchOpts, sink chan<- *CrossCommitReceipt, validator []common.Address, orderHash [][32]byte) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}

	logs, sub, err := _Cross.contract.WatchLogs(opts, "CommitReceipt", validatorRule, orderHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossCommitReceipt)
				if err := _Cross.contract.UnpackLog(event, "CommitReceipt", log); err != nil {
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

// ParseCommitReceipt is a log parse operation binding the contract event 0x581db44feed8ab7f2b0e591fd633c1326a4ba3ea20a5c346ab38fd1f42208e81.
//
// Solidity: event CommitReceipt(address indexed validator, bytes32 indexed orderHash, (bytes32,bytes32) receipt)
func (_Cross *CrossFilterer) ParseCommitReceipt(log types.Log) (*CrossCommitReceipt, error) {
	event := new(CrossCommitReceipt)
	if err := _Cross.contract.UnpackLog(event, "CommitReceipt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossCommitReceiptWithZkIterator is returned from FilterCommitReceiptWithZk and is used to iterate over the raw logs and unpacked data for CommitReceiptWithZk events raised by the Cross contract.
type CrossCommitReceiptWithZkIterator struct {
	Event *CrossCommitReceiptWithZk // Event containing the contract specifics and raw log

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
func (it *CrossCommitReceiptWithZkIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossCommitReceiptWithZk)
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
		it.Event = new(CrossCommitReceiptWithZk)
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
func (it *CrossCommitReceiptWithZkIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossCommitReceiptWithZkIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossCommitReceiptWithZk represents a CommitReceiptWithZk event raised by the Cross contract.
type CrossCommitReceiptWithZk struct {
	Validator common.Address
	OrderHash [32]byte
	Receipt   CrossControllerReceipt
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCommitReceiptWithZk is a free log retrieval operation binding the contract event 0xd58cb8984ca093fbb83c2c99048edb7c88d0c472ba426de692d9bf80ca7438f5.
//
// Solidity: event CommitReceiptWithZk(address indexed validator, bytes32 indexed orderHash, (bytes32,bytes32) receipt)
func (_Cross *CrossFilterer) FilterCommitReceiptWithZk(opts *bind.FilterOpts, validator []common.Address, orderHash [][32]byte) (*CrossCommitReceiptWithZkIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}

	logs, sub, err := _Cross.contract.FilterLogs(opts, "CommitReceiptWithZk", validatorRule, orderHashRule)
	if err != nil {
		return nil, err
	}
	return &CrossCommitReceiptWithZkIterator{contract: _Cross.contract, event: "CommitReceiptWithZk", logs: logs, sub: sub}, nil
}

// WatchCommitReceiptWithZk is a free log subscription operation binding the contract event 0xd58cb8984ca093fbb83c2c99048edb7c88d0c472ba426de692d9bf80ca7438f5.
//
// Solidity: event CommitReceiptWithZk(address indexed validator, bytes32 indexed orderHash, (bytes32,bytes32) receipt)
func (_Cross *CrossFilterer) WatchCommitReceiptWithZk(opts *bind.WatchOpts, sink chan<- *CrossCommitReceiptWithZk, validator []common.Address, orderHash [][32]byte) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}

	logs, sub, err := _Cross.contract.WatchLogs(opts, "CommitReceiptWithZk", validatorRule, orderHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossCommitReceiptWithZk)
				if err := _Cross.contract.UnpackLog(event, "CommitReceiptWithZk", log); err != nil {
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

// ParseCommitReceiptWithZk is a log parse operation binding the contract event 0xd58cb8984ca093fbb83c2c99048edb7c88d0c472ba426de692d9bf80ca7438f5.
//
// Solidity: event CommitReceiptWithZk(address indexed validator, bytes32 indexed orderHash, (bytes32,bytes32) receipt)
func (_Cross *CrossFilterer) ParseCommitReceiptWithZk(log types.Log) (*CrossCommitReceiptWithZk, error) {
	event := new(CrossCommitReceiptWithZk)
	if err := _Cross.contract.UnpackLog(event, "CommitReceiptWithZk", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossCrossFromIterator is returned from FilterCrossFrom and is used to iterate over the raw logs and unpacked data for CrossFrom events raised by the Cross contract.
type CrossCrossFromIterator struct {
	Event *CrossCrossFrom // Event containing the contract specifics and raw log

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
func (it *CrossCrossFromIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossCrossFrom)
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
		it.Event = new(CrossCrossFrom)
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
func (it *CrossCrossFromIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossCrossFromIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossCrossFrom represents a CrossFrom event raised by the Cross contract.
type CrossCrossFrom struct {
	Validator        common.Address
	Order            CrossControllerOrder
	SrcTokenDecimals uint8
	CrossAmount      *big.Int
	PaidAmount       *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterCrossFrom is a free log retrieval operation binding the contract event 0x104f0c1d6ebbba9acf834bd5f27d78481d562d83159d076b974d16bca9c66c21.
//
// Solidity: event CrossFrom(address indexed validator, (uint256,uint256,address,address,uint256,uint256,address,address,address) order, uint8 srcTokenDecimals, uint256 crossAmount, uint256 paidAmount)
func (_Cross *CrossFilterer) FilterCrossFrom(opts *bind.FilterOpts, validator []common.Address) (*CrossCrossFromIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Cross.contract.FilterLogs(opts, "CrossFrom", validatorRule)
	if err != nil {
		return nil, err
	}
	return &CrossCrossFromIterator{contract: _Cross.contract, event: "CrossFrom", logs: logs, sub: sub}, nil
}

// WatchCrossFrom is a free log subscription operation binding the contract event 0x104f0c1d6ebbba9acf834bd5f27d78481d562d83159d076b974d16bca9c66c21.
//
// Solidity: event CrossFrom(address indexed validator, (uint256,uint256,address,address,uint256,uint256,address,address,address) order, uint8 srcTokenDecimals, uint256 crossAmount, uint256 paidAmount)
func (_Cross *CrossFilterer) WatchCrossFrom(opts *bind.WatchOpts, sink chan<- *CrossCrossFrom, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Cross.contract.WatchLogs(opts, "CrossFrom", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossCrossFrom)
				if err := _Cross.contract.UnpackLog(event, "CrossFrom", log); err != nil {
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

// ParseCrossFrom is a log parse operation binding the contract event 0x104f0c1d6ebbba9acf834bd5f27d78481d562d83159d076b974d16bca9c66c21.
//
// Solidity: event CrossFrom(address indexed validator, (uint256,uint256,address,address,uint256,uint256,address,address,address) order, uint8 srcTokenDecimals, uint256 crossAmount, uint256 paidAmount)
func (_Cross *CrossFilterer) ParseCrossFrom(log types.Log) (*CrossCrossFrom, error) {
	event := new(CrossCrossFrom)
	if err := _Cross.contract.UnpackLog(event, "CrossFrom", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossCrossToIterator is returned from FilterCrossTo and is used to iterate over the raw logs and unpacked data for CrossTo events raised by the Cross contract.
type CrossCrossToIterator struct {
	Event *CrossCrossTo // Event containing the contract specifics and raw log

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
func (it *CrossCrossToIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossCrossTo)
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
		it.Event = new(CrossCrossTo)
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
func (it *CrossCrossToIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossCrossToIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossCrossTo represents a CrossTo event raised by the Cross contract.
type CrossCrossTo struct {
	Account        common.Address
	Order          CrossControllerOrder
	FixedFeeAmount *big.Int
	FloatFeeAmount *big.Int
	CrossAmount    *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterCrossTo is a free log retrieval operation binding the contract event 0xeb354ff2ff6b3d6392f3c14565a5e0c60fc642b456cd2538e94968fbc54467e8.
//
// Solidity: event CrossTo(address indexed account, (uint256,uint256,address,address,uint256,uint256,address,address,address) order, uint256 fixedFeeAmount, uint256 floatFeeAmount, uint256 crossAmount)
func (_Cross *CrossFilterer) FilterCrossTo(opts *bind.FilterOpts, account []common.Address) (*CrossCrossToIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Cross.contract.FilterLogs(opts, "CrossTo", accountRule)
	if err != nil {
		return nil, err
	}
	return &CrossCrossToIterator{contract: _Cross.contract, event: "CrossTo", logs: logs, sub: sub}, nil
}

// WatchCrossTo is a free log subscription operation binding the contract event 0xeb354ff2ff6b3d6392f3c14565a5e0c60fc642b456cd2538e94968fbc54467e8.
//
// Solidity: event CrossTo(address indexed account, (uint256,uint256,address,address,uint256,uint256,address,address,address) order, uint256 fixedFeeAmount, uint256 floatFeeAmount, uint256 crossAmount)
func (_Cross *CrossFilterer) WatchCrossTo(opts *bind.WatchOpts, sink chan<- *CrossCrossTo, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Cross.contract.WatchLogs(opts, "CrossTo", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossCrossTo)
				if err := _Cross.contract.UnpackLog(event, "CrossTo", log); err != nil {
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

// ParseCrossTo is a log parse operation binding the contract event 0xeb354ff2ff6b3d6392f3c14565a5e0c60fc642b456cd2538e94968fbc54467e8.
//
// Solidity: event CrossTo(address indexed account, (uint256,uint256,address,address,uint256,uint256,address,address,address) order, uint256 fixedFeeAmount, uint256 floatFeeAmount, uint256 crossAmount)
func (_Cross *CrossFilterer) ParseCrossTo(log types.Log) (*CrossCrossTo, error) {
	event := new(CrossCrossTo)
	if err := _Cross.contract.UnpackLog(event, "CrossTo", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossEnabledZkVerifierIterator is returned from FilterEnabledZkVerifier and is used to iterate over the raw logs and unpacked data for EnabledZkVerifier events raised by the Cross contract.
type CrossEnabledZkVerifierIterator struct {
	Event *CrossEnabledZkVerifier // Event containing the contract specifics and raw log

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
func (it *CrossEnabledZkVerifierIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossEnabledZkVerifier)
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
		it.Event = new(CrossEnabledZkVerifier)
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
func (it *CrossEnabledZkVerifierIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossEnabledZkVerifierIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossEnabledZkVerifier represents a EnabledZkVerifier event raised by the Cross contract.
type CrossEnabledZkVerifier struct {
	Enable bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEnabledZkVerifier is a free log retrieval operation binding the contract event 0xd274b64d953f0ab17c3752f6943a135156a325cc228f1ec0a710a428d7460042.
//
// Solidity: event EnabledZkVerifier(bool _enable)
func (_Cross *CrossFilterer) FilterEnabledZkVerifier(opts *bind.FilterOpts) (*CrossEnabledZkVerifierIterator, error) {

	logs, sub, err := _Cross.contract.FilterLogs(opts, "EnabledZkVerifier")
	if err != nil {
		return nil, err
	}
	return &CrossEnabledZkVerifierIterator{contract: _Cross.contract, event: "EnabledZkVerifier", logs: logs, sub: sub}, nil
}

// WatchEnabledZkVerifier is a free log subscription operation binding the contract event 0xd274b64d953f0ab17c3752f6943a135156a325cc228f1ec0a710a428d7460042.
//
// Solidity: event EnabledZkVerifier(bool _enable)
func (_Cross *CrossFilterer) WatchEnabledZkVerifier(opts *bind.WatchOpts, sink chan<- *CrossEnabledZkVerifier) (event.Subscription, error) {

	logs, sub, err := _Cross.contract.WatchLogs(opts, "EnabledZkVerifier")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossEnabledZkVerifier)
				if err := _Cross.contract.UnpackLog(event, "EnabledZkVerifier", log); err != nil {
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

// ParseEnabledZkVerifier is a log parse operation binding the contract event 0xd274b64d953f0ab17c3752f6943a135156a325cc228f1ec0a710a428d7460042.
//
// Solidity: event EnabledZkVerifier(bool _enable)
func (_Cross *CrossFilterer) ParseEnabledZkVerifier(log types.Log) (*CrossEnabledZkVerifier, error) {
	event := new(CrossEnabledZkVerifier)
	if err := _Cross.contract.UnpackLog(event, "EnabledZkVerifier", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Cross contract.
type CrossInitializedIterator struct {
	Event *CrossInitialized // Event containing the contract specifics and raw log

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
func (it *CrossInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossInitialized)
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
		it.Event = new(CrossInitialized)
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
func (it *CrossInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossInitialized represents a Initialized event raised by the Cross contract.
type CrossInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Cross *CrossFilterer) FilterInitialized(opts *bind.FilterOpts) (*CrossInitializedIterator, error) {

	logs, sub, err := _Cross.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &CrossInitializedIterator{contract: _Cross.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Cross *CrossFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *CrossInitialized) (event.Subscription, error) {

	logs, sub, err := _Cross.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossInitialized)
				if err := _Cross.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Cross *CrossFilterer) ParseInitialized(log types.Log) (*CrossInitialized, error) {
	event := new(CrossInitialized)
	if err := _Cross.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Cross contract.
type CrossPausedIterator struct {
	Event *CrossPaused // Event containing the contract specifics and raw log

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
func (it *CrossPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossPaused)
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
		it.Event = new(CrossPaused)
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
func (it *CrossPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossPaused represents a Paused event raised by the Cross contract.
type CrossPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Cross *CrossFilterer) FilterPaused(opts *bind.FilterOpts) (*CrossPausedIterator, error) {

	logs, sub, err := _Cross.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &CrossPausedIterator{contract: _Cross.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Cross *CrossFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *CrossPaused) (event.Subscription, error) {

	logs, sub, err := _Cross.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossPaused)
				if err := _Cross.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Cross *CrossFilterer) ParsePaused(log types.Log) (*CrossPaused, error) {
	event := new(CrossPaused)
	if err := _Cross.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossSettedFloatFeeIterator is returned from FilterSettedFloatFee and is used to iterate over the raw logs and unpacked data for SettedFloatFee events raised by the Cross contract.
type CrossSettedFloatFeeIterator struct {
	Event *CrossSettedFloatFee // Event containing the contract specifics and raw log

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
func (it *CrossSettedFloatFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossSettedFloatFee)
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
		it.Event = new(CrossSettedFloatFee)
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
func (it *CrossSettedFloatFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossSettedFloatFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossSettedFloatFee represents a SettedFloatFee event raised by the Cross contract.
type CrossSettedFloatFee struct {
	FloatFee *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSettedFloatFee is a free log retrieval operation binding the contract event 0x43d82450da77837bc3bed7812c648e023bf57da520dfce930ae9b975a6e6955d.
//
// Solidity: event SettedFloatFee(uint256 floatFee)
func (_Cross *CrossFilterer) FilterSettedFloatFee(opts *bind.FilterOpts) (*CrossSettedFloatFeeIterator, error) {

	logs, sub, err := _Cross.contract.FilterLogs(opts, "SettedFloatFee")
	if err != nil {
		return nil, err
	}
	return &CrossSettedFloatFeeIterator{contract: _Cross.contract, event: "SettedFloatFee", logs: logs, sub: sub}, nil
}

// WatchSettedFloatFee is a free log subscription operation binding the contract event 0x43d82450da77837bc3bed7812c648e023bf57da520dfce930ae9b975a6e6955d.
//
// Solidity: event SettedFloatFee(uint256 floatFee)
func (_Cross *CrossFilterer) WatchSettedFloatFee(opts *bind.WatchOpts, sink chan<- *CrossSettedFloatFee) (event.Subscription, error) {

	logs, sub, err := _Cross.contract.WatchLogs(opts, "SettedFloatFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossSettedFloatFee)
				if err := _Cross.contract.UnpackLog(event, "SettedFloatFee", log); err != nil {
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

// ParseSettedFloatFee is a log parse operation binding the contract event 0x43d82450da77837bc3bed7812c648e023bf57da520dfce930ae9b975a6e6955d.
//
// Solidity: event SettedFloatFee(uint256 floatFee)
func (_Cross *CrossFilterer) ParseSettedFloatFee(log types.Log) (*CrossSettedFloatFee, error) {
	event := new(CrossSettedFloatFee)
	if err := _Cross.contract.UnpackLog(event, "SettedFloatFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossSettedPorterFactoryIterator is returned from FilterSettedPorterFactory and is used to iterate over the raw logs and unpacked data for SettedPorterFactory events raised by the Cross contract.
type CrossSettedPorterFactoryIterator struct {
	Event *CrossSettedPorterFactory // Event containing the contract specifics and raw log

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
func (it *CrossSettedPorterFactoryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossSettedPorterFactory)
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
		it.Event = new(CrossSettedPorterFactory)
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
func (it *CrossSettedPorterFactoryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossSettedPorterFactoryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossSettedPorterFactory represents a SettedPorterFactory event raised by the Cross contract.
type CrossSettedPorterFactory struct {
	Factory common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSettedPorterFactory is a free log retrieval operation binding the contract event 0x1e213eda8bdfbd30f63a9301525e1aebd4f3845d7c251a0c524a014b6f2c41bb.
//
// Solidity: event SettedPorterFactory(address _factory)
func (_Cross *CrossFilterer) FilterSettedPorterFactory(opts *bind.FilterOpts) (*CrossSettedPorterFactoryIterator, error) {

	logs, sub, err := _Cross.contract.FilterLogs(opts, "SettedPorterFactory")
	if err != nil {
		return nil, err
	}
	return &CrossSettedPorterFactoryIterator{contract: _Cross.contract, event: "SettedPorterFactory", logs: logs, sub: sub}, nil
}

// WatchSettedPorterFactory is a free log subscription operation binding the contract event 0x1e213eda8bdfbd30f63a9301525e1aebd4f3845d7c251a0c524a014b6f2c41bb.
//
// Solidity: event SettedPorterFactory(address _factory)
func (_Cross *CrossFilterer) WatchSettedPorterFactory(opts *bind.WatchOpts, sink chan<- *CrossSettedPorterFactory) (event.Subscription, error) {

	logs, sub, err := _Cross.contract.WatchLogs(opts, "SettedPorterFactory")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossSettedPorterFactory)
				if err := _Cross.contract.UnpackLog(event, "SettedPorterFactory", log); err != nil {
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

// ParseSettedPorterFactory is a log parse operation binding the contract event 0x1e213eda8bdfbd30f63a9301525e1aebd4f3845d7c251a0c524a014b6f2c41bb.
//
// Solidity: event SettedPorterFactory(address _factory)
func (_Cross *CrossFilterer) ParseSettedPorterFactory(log types.Log) (*CrossSettedPorterFactory, error) {
	event := new(CrossSettedPorterFactory)
	if err := _Cross.contract.UnpackLog(event, "SettedPorterFactory", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossSettedZkVerifierIterator is returned from FilterSettedZkVerifier and is used to iterate over the raw logs and unpacked data for SettedZkVerifier events raised by the Cross contract.
type CrossSettedZkVerifierIterator struct {
	Event *CrossSettedZkVerifier // Event containing the contract specifics and raw log

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
func (it *CrossSettedZkVerifierIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossSettedZkVerifier)
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
		it.Event = new(CrossSettedZkVerifier)
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
func (it *CrossSettedZkVerifierIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossSettedZkVerifierIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossSettedZkVerifier represents a SettedZkVerifier event raised by the Cross contract.
type CrossSettedZkVerifier struct {
	ZkVerifier common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSettedZkVerifier is a free log retrieval operation binding the contract event 0x7a2f3aaefa2e60a6eec0505b8a895c8e7e96cd14358fd18c9d7d917f1552ef5a.
//
// Solidity: event SettedZkVerifier(address _zkVerifier)
func (_Cross *CrossFilterer) FilterSettedZkVerifier(opts *bind.FilterOpts) (*CrossSettedZkVerifierIterator, error) {

	logs, sub, err := _Cross.contract.FilterLogs(opts, "SettedZkVerifier")
	if err != nil {
		return nil, err
	}
	return &CrossSettedZkVerifierIterator{contract: _Cross.contract, event: "SettedZkVerifier", logs: logs, sub: sub}, nil
}

// WatchSettedZkVerifier is a free log subscription operation binding the contract event 0x7a2f3aaefa2e60a6eec0505b8a895c8e7e96cd14358fd18c9d7d917f1552ef5a.
//
// Solidity: event SettedZkVerifier(address _zkVerifier)
func (_Cross *CrossFilterer) WatchSettedZkVerifier(opts *bind.WatchOpts, sink chan<- *CrossSettedZkVerifier) (event.Subscription, error) {

	logs, sub, err := _Cross.contract.WatchLogs(opts, "SettedZkVerifier")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossSettedZkVerifier)
				if err := _Cross.contract.UnpackLog(event, "SettedZkVerifier", log); err != nil {
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

// ParseSettedZkVerifier is a log parse operation binding the contract event 0x7a2f3aaefa2e60a6eec0505b8a895c8e7e96cd14358fd18c9d7d917f1552ef5a.
//
// Solidity: event SettedZkVerifier(address _zkVerifier)
func (_Cross *CrossFilterer) ParseSettedZkVerifier(log types.Log) (*CrossSettedZkVerifier, error) {
	event := new(CrossSettedZkVerifier)
	if err := _Cross.contract.UnpackLog(event, "SettedZkVerifier", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Cross contract.
type CrossUnpausedIterator struct {
	Event *CrossUnpaused // Event containing the contract specifics and raw log

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
func (it *CrossUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossUnpaused)
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
		it.Event = new(CrossUnpaused)
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
func (it *CrossUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossUnpaused represents a Unpaused event raised by the Cross contract.
type CrossUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Cross *CrossFilterer) FilterUnpaused(opts *bind.FilterOpts) (*CrossUnpausedIterator, error) {

	logs, sub, err := _Cross.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &CrossUnpausedIterator{contract: _Cross.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Cross *CrossFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *CrossUnpaused) (event.Subscription, error) {

	logs, sub, err := _Cross.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossUnpaused)
				if err := _Cross.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Cross *CrossFilterer) ParseUnpaused(log types.Log) (*CrossUnpaused, error) {
	event := new(CrossUnpaused)
	if err := _Cross.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
