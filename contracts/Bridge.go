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

// BridgeMetaData contains all meta data concerning the Bridge contract.
var BridgeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_mimc\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_controller\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"commitment\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"leafIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"CommitmentAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"LeafAdded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"cur\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_commitment\",\"type\":\"uint256\"}],\"name\":\"addCommitment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"commitments\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"controller\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"j\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"k\",\"type\":\"uint256\"}],\"name\":\"getLeaf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getMerkleProof\",\"outputs\":[{\"internalType\":\"uint256[8]\",\"name\":\"\",\"type\":\"uint256[8]\"},{\"internalType\":\"uint256[8]\",\"name\":\"\",\"type\":\"uint256[8]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNextCommitmentIdx\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"cur\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"input\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sk\",\"type\":\"uint256\"}],\"name\":\"getPoseidon\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRoot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"leaf\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depth\",\"type\":\"uint256\"}],\"name\":\"getUniqueLeaf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"com\",\"type\":\"uint256\"}],\"name\":\"insert\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_root\",\"type\":\"uint256\"}],\"name\":\"isKnownRoot\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"no_leaves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"roots\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"serials\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tree_depth\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateTree\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"a\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"b\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"c\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"input\",\"type\":\"uint256[2]\"}],\"name\":\"verify\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"a\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"b\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"c\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"input\",\"type\":\"uint256[2]\"}],\"name\":\"verifyProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"r\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// BridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use BridgeMetaData.ABI instead.
var BridgeABI = BridgeMetaData.ABI

// Bridge is an auto generated Go binding around an Ethereum contract.
type Bridge struct {
	BridgeCaller     // Read-only binding to the contract
	BridgeTransactor // Write-only binding to the contract
	BridgeFilterer   // Log filterer for contract events
}

// BridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeSession struct {
	Contract     *Bridge           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeCallerSession struct {
	Contract *BridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeTransactorSession struct {
	Contract     *BridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeRaw struct {
	Contract *Bridge // Generic contract binding to access the raw methods on
}

// BridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeCallerRaw struct {
	Contract *BridgeCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeTransactorRaw struct {
	Contract *BridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridge creates a new instance of Bridge, bound to a specific deployed contract.
func NewBridge(address common.Address, backend bind.ContractBackend) (*Bridge, error) {
	contract, err := bindBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bridge{BridgeCaller: BridgeCaller{contract: contract}, BridgeTransactor: BridgeTransactor{contract: contract}, BridgeFilterer: BridgeFilterer{contract: contract}}, nil
}

// NewBridgeCaller creates a new read-only instance of Bridge, bound to a specific deployed contract.
func NewBridgeCaller(address common.Address, caller bind.ContractCaller) (*BridgeCaller, error) {
	contract, err := bindBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeCaller{contract: contract}, nil
}

// NewBridgeTransactor creates a new write-only instance of Bridge, bound to a specific deployed contract.
func NewBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeTransactor, error) {
	contract, err := bindBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeTransactor{contract: contract}, nil
}

// NewBridgeFilterer creates a new log filterer instance of Bridge, bound to a specific deployed contract.
func NewBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeFilterer, error) {
	contract, err := bindBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeFilterer{contract: contract}, nil
}

// bindBridge binds a generic wrapper to an already deployed contract.
func bindBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BridgeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridge *BridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bridge.Contract.BridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridge *BridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.Contract.BridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridge *BridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridge.Contract.BridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridge *BridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridge *BridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridge *BridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridge.Contract.contract.Transact(opts, method, params...)
}

// MT is a free data retrieval call binding the contract method 0x4a1f11a7.
//
// Solidity: function MT() view returns(uint256 cur)
func (_Bridge *BridgeCaller) MT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "MT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MT is a free data retrieval call binding the contract method 0x4a1f11a7.
//
// Solidity: function MT() view returns(uint256 cur)
func (_Bridge *BridgeSession) MT() (*big.Int, error) {
	return _Bridge.Contract.MT(&_Bridge.CallOpts)
}

// MT is a free data retrieval call binding the contract method 0x4a1f11a7.
//
// Solidity: function MT() view returns(uint256 cur)
func (_Bridge *BridgeCallerSession) MT() (*big.Int, error) {
	return _Bridge.Contract.MT(&_Bridge.CallOpts)
}

// Commitments is a free data retrieval call binding the contract method 0x49ce8997.
//
// Solidity: function commitments(uint256 ) view returns(bool)
func (_Bridge *BridgeCaller) Commitments(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "commitments", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Commitments is a free data retrieval call binding the contract method 0x49ce8997.
//
// Solidity: function commitments(uint256 ) view returns(bool)
func (_Bridge *BridgeSession) Commitments(arg0 *big.Int) (bool, error) {
	return _Bridge.Contract.Commitments(&_Bridge.CallOpts, arg0)
}

// Commitments is a free data retrieval call binding the contract method 0x49ce8997.
//
// Solidity: function commitments(uint256 ) view returns(bool)
func (_Bridge *BridgeCallerSession) Commitments(arg0 *big.Int) (bool, error) {
	return _Bridge.Contract.Commitments(&_Bridge.CallOpts, arg0)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_Bridge *BridgeCaller) Controller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "controller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_Bridge *BridgeSession) Controller() (common.Address, error) {
	return _Bridge.Contract.Controller(&_Bridge.CallOpts)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_Bridge *BridgeCallerSession) Controller() (common.Address, error) {
	return _Bridge.Contract.Controller(&_Bridge.CallOpts)
}

// GetLeaf is a free data retrieval call binding the contract method 0x2fb2798d.
//
// Solidity: function getLeaf(uint256 j, uint256 k) view returns(uint256 root)
func (_Bridge *BridgeCaller) GetLeaf(opts *bind.CallOpts, j *big.Int, k *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "getLeaf", j, k)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLeaf is a free data retrieval call binding the contract method 0x2fb2798d.
//
// Solidity: function getLeaf(uint256 j, uint256 k) view returns(uint256 root)
func (_Bridge *BridgeSession) GetLeaf(j *big.Int, k *big.Int) (*big.Int, error) {
	return _Bridge.Contract.GetLeaf(&_Bridge.CallOpts, j, k)
}

// GetLeaf is a free data retrieval call binding the contract method 0x2fb2798d.
//
// Solidity: function getLeaf(uint256 j, uint256 k) view returns(uint256 root)
func (_Bridge *BridgeCallerSession) GetLeaf(j *big.Int, k *big.Int) (*big.Int, error) {
	return _Bridge.Contract.GetLeaf(&_Bridge.CallOpts, j, k)
}

// GetMerkleProof is a free data retrieval call binding the contract method 0xc41dd397.
//
// Solidity: function getMerkleProof(uint256 index) view returns(uint256[8], uint256[8])
func (_Bridge *BridgeCaller) GetMerkleProof(opts *bind.CallOpts, index *big.Int) ([8]*big.Int, [8]*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "getMerkleProof", index)

	if err != nil {
		return *new([8]*big.Int), *new([8]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([8]*big.Int)).(*[8]*big.Int)
	out1 := *abi.ConvertType(out[1], new([8]*big.Int)).(*[8]*big.Int)

	return out0, out1, err

}

// GetMerkleProof is a free data retrieval call binding the contract method 0xc41dd397.
//
// Solidity: function getMerkleProof(uint256 index) view returns(uint256[8], uint256[8])
func (_Bridge *BridgeSession) GetMerkleProof(index *big.Int) ([8]*big.Int, [8]*big.Int, error) {
	return _Bridge.Contract.GetMerkleProof(&_Bridge.CallOpts, index)
}

// GetMerkleProof is a free data retrieval call binding the contract method 0xc41dd397.
//
// Solidity: function getMerkleProof(uint256 index) view returns(uint256[8], uint256[8])
func (_Bridge *BridgeCallerSession) GetMerkleProof(index *big.Int) ([8]*big.Int, [8]*big.Int, error) {
	return _Bridge.Contract.GetMerkleProof(&_Bridge.CallOpts, index)
}

// GetNextCommitmentIdx is a free data retrieval call binding the contract method 0xf9a541d6.
//
// Solidity: function getNextCommitmentIdx() view returns(uint256 cur)
func (_Bridge *BridgeCaller) GetNextCommitmentIdx(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "getNextCommitmentIdx")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextCommitmentIdx is a free data retrieval call binding the contract method 0xf9a541d6.
//
// Solidity: function getNextCommitmentIdx() view returns(uint256 cur)
func (_Bridge *BridgeSession) GetNextCommitmentIdx() (*big.Int, error) {
	return _Bridge.Contract.GetNextCommitmentIdx(&_Bridge.CallOpts)
}

// GetNextCommitmentIdx is a free data retrieval call binding the contract method 0xf9a541d6.
//
// Solidity: function getNextCommitmentIdx() view returns(uint256 cur)
func (_Bridge *BridgeCallerSession) GetNextCommitmentIdx() (*big.Int, error) {
	return _Bridge.Contract.GetNextCommitmentIdx(&_Bridge.CallOpts)
}

// GetPoseidon is a free data retrieval call binding the contract method 0xf662db24.
//
// Solidity: function getPoseidon(uint256 input, uint256 sk) view returns(uint256)
func (_Bridge *BridgeCaller) GetPoseidon(opts *bind.CallOpts, input *big.Int, sk *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "getPoseidon", input, sk)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPoseidon is a free data retrieval call binding the contract method 0xf662db24.
//
// Solidity: function getPoseidon(uint256 input, uint256 sk) view returns(uint256)
func (_Bridge *BridgeSession) GetPoseidon(input *big.Int, sk *big.Int) (*big.Int, error) {
	return _Bridge.Contract.GetPoseidon(&_Bridge.CallOpts, input, sk)
}

// GetPoseidon is a free data retrieval call binding the contract method 0xf662db24.
//
// Solidity: function getPoseidon(uint256 input, uint256 sk) view returns(uint256)
func (_Bridge *BridgeCallerSession) GetPoseidon(input *big.Int, sk *big.Int) (*big.Int, error) {
	return _Bridge.Contract.GetPoseidon(&_Bridge.CallOpts, input, sk)
}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() view returns(uint256 root)
func (_Bridge *BridgeCaller) GetRoot(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "getRoot")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() view returns(uint256 root)
func (_Bridge *BridgeSession) GetRoot() (*big.Int, error) {
	return _Bridge.Contract.GetRoot(&_Bridge.CallOpts)
}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() view returns(uint256 root)
func (_Bridge *BridgeCallerSession) GetRoot() (*big.Int, error) {
	return _Bridge.Contract.GetRoot(&_Bridge.CallOpts)
}

// GetUniqueLeaf is a free data retrieval call binding the contract method 0x4d556bed.
//
// Solidity: function getUniqueLeaf(uint256 leaf, uint256 depth) view returns(uint256)
func (_Bridge *BridgeCaller) GetUniqueLeaf(opts *bind.CallOpts, leaf *big.Int, depth *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "getUniqueLeaf", leaf, depth)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUniqueLeaf is a free data retrieval call binding the contract method 0x4d556bed.
//
// Solidity: function getUniqueLeaf(uint256 leaf, uint256 depth) view returns(uint256)
func (_Bridge *BridgeSession) GetUniqueLeaf(leaf *big.Int, depth *big.Int) (*big.Int, error) {
	return _Bridge.Contract.GetUniqueLeaf(&_Bridge.CallOpts, leaf, depth)
}

// GetUniqueLeaf is a free data retrieval call binding the contract method 0x4d556bed.
//
// Solidity: function getUniqueLeaf(uint256 leaf, uint256 depth) view returns(uint256)
func (_Bridge *BridgeCallerSession) GetUniqueLeaf(leaf *big.Int, depth *big.Int) (*big.Int, error) {
	return _Bridge.Contract.GetUniqueLeaf(&_Bridge.CallOpts, leaf, depth)
}

// IsKnownRoot is a free data retrieval call binding the contract method 0xa6232a93.
//
// Solidity: function isKnownRoot(uint256 _root) view returns(bool)
func (_Bridge *BridgeCaller) IsKnownRoot(opts *bind.CallOpts, _root *big.Int) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "isKnownRoot", _root)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsKnownRoot is a free data retrieval call binding the contract method 0xa6232a93.
//
// Solidity: function isKnownRoot(uint256 _root) view returns(bool)
func (_Bridge *BridgeSession) IsKnownRoot(_root *big.Int) (bool, error) {
	return _Bridge.Contract.IsKnownRoot(&_Bridge.CallOpts, _root)
}

// IsKnownRoot is a free data retrieval call binding the contract method 0xa6232a93.
//
// Solidity: function isKnownRoot(uint256 _root) view returns(bool)
func (_Bridge *BridgeCallerSession) IsKnownRoot(_root *big.Int) (bool, error) {
	return _Bridge.Contract.IsKnownRoot(&_Bridge.CallOpts, _root)
}

// NoLeaves is a free data retrieval call binding the contract method 0xa03515fd.
//
// Solidity: function no_leaves() view returns(uint256)
func (_Bridge *BridgeCaller) NoLeaves(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "no_leaves")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NoLeaves is a free data retrieval call binding the contract method 0xa03515fd.
//
// Solidity: function no_leaves() view returns(uint256)
func (_Bridge *BridgeSession) NoLeaves() (*big.Int, error) {
	return _Bridge.Contract.NoLeaves(&_Bridge.CallOpts)
}

// NoLeaves is a free data retrieval call binding the contract method 0xa03515fd.
//
// Solidity: function no_leaves() view returns(uint256)
func (_Bridge *BridgeCallerSession) NoLeaves() (*big.Int, error) {
	return _Bridge.Contract.NoLeaves(&_Bridge.CallOpts)
}

// Roots is a free data retrieval call binding the contract method 0xc2b40ae4.
//
// Solidity: function roots(uint256 ) view returns(bool)
func (_Bridge *BridgeCaller) Roots(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "roots", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Roots is a free data retrieval call binding the contract method 0xc2b40ae4.
//
// Solidity: function roots(uint256 ) view returns(bool)
func (_Bridge *BridgeSession) Roots(arg0 *big.Int) (bool, error) {
	return _Bridge.Contract.Roots(&_Bridge.CallOpts, arg0)
}

// Roots is a free data retrieval call binding the contract method 0xc2b40ae4.
//
// Solidity: function roots(uint256 ) view returns(bool)
func (_Bridge *BridgeCallerSession) Roots(arg0 *big.Int) (bool, error) {
	return _Bridge.Contract.Roots(&_Bridge.CallOpts, arg0)
}

// Serials is a free data retrieval call binding the contract method 0x98cd96af.
//
// Solidity: function serials(uint256 ) view returns(bool)
func (_Bridge *BridgeCaller) Serials(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "serials", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Serials is a free data retrieval call binding the contract method 0x98cd96af.
//
// Solidity: function serials(uint256 ) view returns(bool)
func (_Bridge *BridgeSession) Serials(arg0 *big.Int) (bool, error) {
	return _Bridge.Contract.Serials(&_Bridge.CallOpts, arg0)
}

// Serials is a free data retrieval call binding the contract method 0x98cd96af.
//
// Solidity: function serials(uint256 ) view returns(bool)
func (_Bridge *BridgeCallerSession) Serials(arg0 *big.Int) (bool, error) {
	return _Bridge.Contract.Serials(&_Bridge.CallOpts, arg0)
}

// TreeDepth is a free data retrieval call binding the contract method 0xe08dfff8.
//
// Solidity: function tree_depth() view returns(uint256)
func (_Bridge *BridgeCaller) TreeDepth(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "tree_depth")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TreeDepth is a free data retrieval call binding the contract method 0xe08dfff8.
//
// Solidity: function tree_depth() view returns(uint256)
func (_Bridge *BridgeSession) TreeDepth() (*big.Int, error) {
	return _Bridge.Contract.TreeDepth(&_Bridge.CallOpts)
}

// TreeDepth is a free data retrieval call binding the contract method 0xe08dfff8.
//
// Solidity: function tree_depth() view returns(uint256)
func (_Bridge *BridgeCallerSession) TreeDepth() (*big.Int, error) {
	return _Bridge.Contract.TreeDepth(&_Bridge.CallOpts)
}

// Verify is a free data retrieval call binding the contract method 0xdff66416.
//
// Solidity: function verify(uint256[2] a, uint256[2][2] b, uint256[2] c, uint256[2] input) view returns(bool)
func (_Bridge *BridgeCaller) Verify(opts *bind.CallOpts, a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, input [2]*big.Int) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "verify", a, b, c, input)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Verify is a free data retrieval call binding the contract method 0xdff66416.
//
// Solidity: function verify(uint256[2] a, uint256[2][2] b, uint256[2] c, uint256[2] input) view returns(bool)
func (_Bridge *BridgeSession) Verify(a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, input [2]*big.Int) (bool, error) {
	return _Bridge.Contract.Verify(&_Bridge.CallOpts, a, b, c, input)
}

// Verify is a free data retrieval call binding the contract method 0xdff66416.
//
// Solidity: function verify(uint256[2] a, uint256[2][2] b, uint256[2] c, uint256[2] input) view returns(bool)
func (_Bridge *BridgeCallerSession) Verify(a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, input [2]*big.Int) (bool, error) {
	return _Bridge.Contract.Verify(&_Bridge.CallOpts, a, b, c, input)
}

// VerifyProof is a free data retrieval call binding the contract method 0xf5c9d69e.
//
// Solidity: function verifyProof(uint256[2] a, uint256[2][2] b, uint256[2] c, uint256[2] input) view returns(bool r)
func (_Bridge *BridgeCaller) VerifyProof(opts *bind.CallOpts, a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, input [2]*big.Int) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "verifyProof", a, b, c, input)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyProof is a free data retrieval call binding the contract method 0xf5c9d69e.
//
// Solidity: function verifyProof(uint256[2] a, uint256[2][2] b, uint256[2] c, uint256[2] input) view returns(bool r)
func (_Bridge *BridgeSession) VerifyProof(a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, input [2]*big.Int) (bool, error) {
	return _Bridge.Contract.VerifyProof(&_Bridge.CallOpts, a, b, c, input)
}

// VerifyProof is a free data retrieval call binding the contract method 0xf5c9d69e.
//
// Solidity: function verifyProof(uint256[2] a, uint256[2][2] b, uint256[2] c, uint256[2] input) view returns(bool r)
func (_Bridge *BridgeCallerSession) VerifyProof(a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, input [2]*big.Int) (bool, error) {
	return _Bridge.Contract.VerifyProof(&_Bridge.CallOpts, a, b, c, input)
}

// AddCommitment is a paid mutator transaction binding the contract method 0x797a6890.
//
// Solidity: function addCommitment(uint256 _commitment) returns()
func (_Bridge *BridgeTransactor) AddCommitment(opts *bind.TransactOpts, _commitment *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "addCommitment", _commitment)
}

// AddCommitment is a paid mutator transaction binding the contract method 0x797a6890.
//
// Solidity: function addCommitment(uint256 _commitment) returns()
func (_Bridge *BridgeSession) AddCommitment(_commitment *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.AddCommitment(&_Bridge.TransactOpts, _commitment)
}

// AddCommitment is a paid mutator transaction binding the contract method 0x797a6890.
//
// Solidity: function addCommitment(uint256 _commitment) returns()
func (_Bridge *BridgeTransactorSession) AddCommitment(_commitment *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.AddCommitment(&_Bridge.TransactOpts, _commitment)
}

// Insert is a paid mutator transaction binding the contract method 0x90b5561d.
//
// Solidity: function insert(uint256 com) returns(uint256)
func (_Bridge *BridgeTransactor) Insert(opts *bind.TransactOpts, com *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "insert", com)
}

// Insert is a paid mutator transaction binding the contract method 0x90b5561d.
//
// Solidity: function insert(uint256 com) returns(uint256)
func (_Bridge *BridgeSession) Insert(com *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.Insert(&_Bridge.TransactOpts, com)
}

// Insert is a paid mutator transaction binding the contract method 0x90b5561d.
//
// Solidity: function insert(uint256 com) returns(uint256)
func (_Bridge *BridgeTransactorSession) Insert(com *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.Insert(&_Bridge.TransactOpts, com)
}

// UpdateTree is a paid mutator transaction binding the contract method 0xee427b1d.
//
// Solidity: function updateTree() returns(uint256 root)
func (_Bridge *BridgeTransactor) UpdateTree(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "updateTree")
}

// UpdateTree is a paid mutator transaction binding the contract method 0xee427b1d.
//
// Solidity: function updateTree() returns(uint256 root)
func (_Bridge *BridgeSession) UpdateTree() (*types.Transaction, error) {
	return _Bridge.Contract.UpdateTree(&_Bridge.TransactOpts)
}

// UpdateTree is a paid mutator transaction binding the contract method 0xee427b1d.
//
// Solidity: function updateTree() returns(uint256 root)
func (_Bridge *BridgeTransactorSession) UpdateTree() (*types.Transaction, error) {
	return _Bridge.Contract.UpdateTree(&_Bridge.TransactOpts)
}

// BridgeCommitmentAddedIterator is returned from FilterCommitmentAdded and is used to iterate over the raw logs and unpacked data for CommitmentAdded events raised by the Bridge contract.
type BridgeCommitmentAddedIterator struct {
	Event *BridgeCommitmentAdded // Event containing the contract specifics and raw log

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
func (it *BridgeCommitmentAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeCommitmentAdded)
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
		it.Event = new(BridgeCommitmentAdded)
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
func (it *BridgeCommitmentAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeCommitmentAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeCommitmentAdded represents a CommitmentAdded event raised by the Bridge contract.
type BridgeCommitmentAdded struct {
	Commitment *big.Int
	LeafIndex  *big.Int
	Timestamp  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCommitmentAdded is a free log retrieval operation binding the contract event 0x65ba9a9f4b62e3d394afcc0ed4b32089839f4dd83f4cfda7e495e09f2cc5657f.
//
// Solidity: event CommitmentAdded(uint256 indexed commitment, uint256 leafIndex, uint256 timestamp)
func (_Bridge *BridgeFilterer) FilterCommitmentAdded(opts *bind.FilterOpts, commitment []*big.Int) (*BridgeCommitmentAddedIterator, error) {

	var commitmentRule []interface{}
	for _, commitmentItem := range commitment {
		commitmentRule = append(commitmentRule, commitmentItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "CommitmentAdded", commitmentRule)
	if err != nil {
		return nil, err
	}
	return &BridgeCommitmentAddedIterator{contract: _Bridge.contract, event: "CommitmentAdded", logs: logs, sub: sub}, nil
}

// WatchCommitmentAdded is a free log subscription operation binding the contract event 0x65ba9a9f4b62e3d394afcc0ed4b32089839f4dd83f4cfda7e495e09f2cc5657f.
//
// Solidity: event CommitmentAdded(uint256 indexed commitment, uint256 leafIndex, uint256 timestamp)
func (_Bridge *BridgeFilterer) WatchCommitmentAdded(opts *bind.WatchOpts, sink chan<- *BridgeCommitmentAdded, commitment []*big.Int) (event.Subscription, error) {

	var commitmentRule []interface{}
	for _, commitmentItem := range commitment {
		commitmentRule = append(commitmentRule, commitmentItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "CommitmentAdded", commitmentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeCommitmentAdded)
				if err := _Bridge.contract.UnpackLog(event, "CommitmentAdded", log); err != nil {
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

// ParseCommitmentAdded is a log parse operation binding the contract event 0x65ba9a9f4b62e3d394afcc0ed4b32089839f4dd83f4cfda7e495e09f2cc5657f.
//
// Solidity: event CommitmentAdded(uint256 indexed commitment, uint256 leafIndex, uint256 timestamp)
func (_Bridge *BridgeFilterer) ParseCommitmentAdded(log types.Log) (*BridgeCommitmentAdded, error) {
	event := new(BridgeCommitmentAdded)
	if err := _Bridge.contract.UnpackLog(event, "CommitmentAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeLeafAddedIterator is returned from FilterLeafAdded and is used to iterate over the raw logs and unpacked data for LeafAdded events raised by the Bridge contract.
type BridgeLeafAddedIterator struct {
	Event *BridgeLeafAdded // Event containing the contract specifics and raw log

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
func (it *BridgeLeafAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeLeafAdded)
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
		it.Event = new(BridgeLeafAdded)
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
func (it *BridgeLeafAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeLeafAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeLeafAdded represents a LeafAdded event raised by the Bridge contract.
type BridgeLeafAdded struct {
	Index *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterLeafAdded is a free log retrieval operation binding the contract event 0xee14d40b05aa42dee41bd28416dcb3ca632a33e9f6ea33a4bca7ffcb95938858.
//
// Solidity: event LeafAdded(uint256 index)
func (_Bridge *BridgeFilterer) FilterLeafAdded(opts *bind.FilterOpts) (*BridgeLeafAddedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "LeafAdded")
	if err != nil {
		return nil, err
	}
	return &BridgeLeafAddedIterator{contract: _Bridge.contract, event: "LeafAdded", logs: logs, sub: sub}, nil
}

// WatchLeafAdded is a free log subscription operation binding the contract event 0xee14d40b05aa42dee41bd28416dcb3ca632a33e9f6ea33a4bca7ffcb95938858.
//
// Solidity: event LeafAdded(uint256 index)
func (_Bridge *BridgeFilterer) WatchLeafAdded(opts *bind.WatchOpts, sink chan<- *BridgeLeafAdded) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "LeafAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeLeafAdded)
				if err := _Bridge.contract.UnpackLog(event, "LeafAdded", log); err != nil {
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

// ParseLeafAdded is a log parse operation binding the contract event 0xee14d40b05aa42dee41bd28416dcb3ca632a33e9f6ea33a4bca7ffcb95938858.
//
// Solidity: event LeafAdded(uint256 index)
func (_Bridge *BridgeFilterer) ParseLeafAdded(log types.Log) (*BridgeLeafAdded, error) {
	event := new(BridgeLeafAdded)
	if err := _Bridge.contract.UnpackLog(event, "LeafAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
