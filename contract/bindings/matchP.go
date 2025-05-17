// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bingdings

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

// MatchPgameInfo is an auto generated low-level Go binding around an user-defined struct.
type MatchPgameInfo struct {
	Id             *big.Int
	Name           string
	Players        []common.Address
	StartTime      *big.Int
	EndTime        *big.Int
	IsSettled      *big.Int
	Player0Balance *big.Int
	Player1Balance *big.Int
}

// BingdingsMetaData contains all meta data concerning the Bingdings contract.
var BingdingsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GameAlreadyOver\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"gameId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"builder\",\"type\":\"address\"}],\"name\":\"CreateGame\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"DoNothing\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gameId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"player\",\"type\":\"address\"}],\"name\":\"JoinGame\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gameId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winnerVoter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_time\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_bonus\",\"type\":\"uint256\"}],\"name\":\"Settle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gameId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_time\",\"type\":\"uint256\"}],\"name\":\"Settled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gameId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"player\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Vote\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Aplayers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Avoters\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MatchProtocol\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_nonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_totalScore\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"addWhiteList\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_stake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalWinning\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalLosing\",\"type\":\"uint256\"}],\"name\":\"calculateBonusTwoDecimals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"bonus\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_endTime\",\"type\":\"uint256\"}],\"name\":\"createGame\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"endAppend\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeDecimals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_gameId\",\"type\":\"uint256\"}],\"name\":\"forceGameOver\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"gameInfoList\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"isSettled\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"player0Balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"player1Balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"games\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"exist\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"isSettled\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"averageScore\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index_slot\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllGames\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"players\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"isSettled\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"player0Balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"player1Balance\",\"type\":\"uint256\"}],\"internalType\":\"structMatchP.gameInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGameCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"index0\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"protocol\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"st\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"et\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isGain\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isPlayer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_gameId\",\"type\":\"uint256\"}],\"name\":\"isVoting\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_gameId\",\"type\":\"uint256\"}],\"name\":\"joinGame\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"playerStakeBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"setToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_gameId\",\"type\":\"uint256\"}],\"name\":\"settlement\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_staker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startAppend\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_gameId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_player\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"vote\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"voteBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"voterGameStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"gameId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"player\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"stakeAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"whiteList\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// BingdingsABI is the input ABI used to generate the binding from.
// Deprecated: Use BingdingsMetaData.ABI instead.
var BingdingsABI = BingdingsMetaData.ABI

// Bingdings is an auto generated Go binding around an Ethereum contract.
type Bingdings struct {
	BingdingsCaller     // Read-only binding to the contract
	BingdingsTransactor // Write-only binding to the contract
	BingdingsFilterer   // Log filterer for contract events
}

// BingdingsCaller is an auto generated read-only Go binding around an Ethereum contract.
type BingdingsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BingdingsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BingdingsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BingdingsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BingdingsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BingdingsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BingdingsSession struct {
	Contract     *Bingdings        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BingdingsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BingdingsCallerSession struct {
	Contract *BingdingsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// BingdingsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BingdingsTransactorSession struct {
	Contract     *BingdingsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// BingdingsRaw is an auto generated low-level Go binding around an Ethereum contract.
type BingdingsRaw struct {
	Contract *Bingdings // Generic contract binding to access the raw methods on
}

// BingdingsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BingdingsCallerRaw struct {
	Contract *BingdingsCaller // Generic read-only contract binding to access the raw methods on
}

// BingdingsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BingdingsTransactorRaw struct {
	Contract *BingdingsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBingdings creates a new instance of Bingdings, bound to a specific deployed contract.
func NewBingdings(address common.Address, backend bind.ContractBackend) (*Bingdings, error) {
	contract, err := bindBingdings(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bingdings{BingdingsCaller: BingdingsCaller{contract: contract}, BingdingsTransactor: BingdingsTransactor{contract: contract}, BingdingsFilterer: BingdingsFilterer{contract: contract}}, nil
}

// NewBingdingsCaller creates a new read-only instance of Bingdings, bound to a specific deployed contract.
func NewBingdingsCaller(address common.Address, caller bind.ContractCaller) (*BingdingsCaller, error) {
	contract, err := bindBingdings(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BingdingsCaller{contract: contract}, nil
}

// NewBingdingsTransactor creates a new write-only instance of Bingdings, bound to a specific deployed contract.
func NewBingdingsTransactor(address common.Address, transactor bind.ContractTransactor) (*BingdingsTransactor, error) {
	contract, err := bindBingdings(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BingdingsTransactor{contract: contract}, nil
}

// NewBingdingsFilterer creates a new log filterer instance of Bingdings, bound to a specific deployed contract.
func NewBingdingsFilterer(address common.Address, filterer bind.ContractFilterer) (*BingdingsFilterer, error) {
	contract, err := bindBingdings(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BingdingsFilterer{contract: contract}, nil
}

// bindBingdings binds a generic wrapper to an already deployed contract.
func bindBingdings(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BingdingsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bingdings *BingdingsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bingdings.Contract.BingdingsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bingdings *BingdingsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bingdings.Contract.BingdingsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bingdings *BingdingsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bingdings.Contract.BingdingsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bingdings *BingdingsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bingdings.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bingdings *BingdingsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bingdings.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bingdings *BingdingsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bingdings.Contract.contract.Transact(opts, method, params...)
}

// Aplayers is a free data retrieval call binding the contract method 0x3c59d2ff.
//
// Solidity: function Aplayers(uint256 , uint256 ) view returns(address)
func (_Bingdings *BingdingsCaller) Aplayers(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Bingdings.contract.Call(opts, &out, "Aplayers", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Aplayers is a free data retrieval call binding the contract method 0x3c59d2ff.
//
// Solidity: function Aplayers(uint256 , uint256 ) view returns(address)
func (_Bingdings *BingdingsSession) Aplayers(arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	return _Bingdings.Contract.Aplayers(&_Bingdings.CallOpts, arg0, arg1)
}

// Aplayers is a free data retrieval call binding the contract method 0x3c59d2ff.
//
// Solidity: function Aplayers(uint256 , uint256 ) view returns(address)
func (_Bingdings *BingdingsCallerSession) Aplayers(arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	return _Bingdings.Contract.Aplayers(&_Bingdings.CallOpts, arg0, arg1)
}

// Avoters is a free data retrieval call binding the contract method 0x880c4589.
//
// Solidity: function Avoters(uint256 , uint256 ) view returns(address)
func (_Bingdings *BingdingsCaller) Avoters(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Bingdings.contract.Call(opts, &out, "Avoters", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Avoters is a free data retrieval call binding the contract method 0x880c4589.
//
// Solidity: function Avoters(uint256 , uint256 ) view returns(address)
func (_Bingdings *BingdingsSession) Avoters(arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	return _Bingdings.Contract.Avoters(&_Bingdings.CallOpts, arg0, arg1)
}

// Avoters is a free data retrieval call binding the contract method 0x880c4589.
//
// Solidity: function Avoters(uint256 , uint256 ) view returns(address)
func (_Bingdings *BingdingsCallerSession) Avoters(arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	return _Bingdings.Contract.Avoters(&_Bingdings.CallOpts, arg0, arg1)
}

// MatchProtocol is a free data retrieval call binding the contract method 0x372ca54a.
//
// Solidity: function MatchProtocol() view returns(address)
func (_Bingdings *BingdingsCaller) MatchProtocol(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bingdings.contract.Call(opts, &out, "MatchProtocol")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MatchProtocol is a free data retrieval call binding the contract method 0x372ca54a.
//
// Solidity: function MatchProtocol() view returns(address)
func (_Bingdings *BingdingsSession) MatchProtocol() (common.Address, error) {
	return _Bingdings.Contract.MatchProtocol(&_Bingdings.CallOpts)
}

// MatchProtocol is a free data retrieval call binding the contract method 0x372ca54a.
//
// Solidity: function MatchProtocol() view returns(address)
func (_Bingdings *BingdingsCallerSession) MatchProtocol() (common.Address, error) {
	return _Bingdings.Contract.MatchProtocol(&_Bingdings.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Bingdings *BingdingsCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Bingdings.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Bingdings *BingdingsSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Bingdings.Contract.UPGRADEINTERFACEVERSION(&_Bingdings.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Bingdings *BingdingsCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Bingdings.Contract.UPGRADEINTERFACEVERSION(&_Bingdings.CallOpts)
}

// Nonce is a free data retrieval call binding the contract method 0x5d1222aa.
//
// Solidity: function _nonce() view returns(uint256)
func (_Bingdings *BingdingsCaller) Nonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bingdings.contract.Call(opts, &out, "_nonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonce is a free data retrieval call binding the contract method 0x5d1222aa.
//
// Solidity: function _nonce() view returns(uint256)
func (_Bingdings *BingdingsSession) Nonce() (*big.Int, error) {
	return _Bingdings.Contract.Nonce(&_Bingdings.CallOpts)
}

// Nonce is a free data retrieval call binding the contract method 0x5d1222aa.
//
// Solidity: function _nonce() view returns(uint256)
func (_Bingdings *BingdingsCallerSession) Nonce() (*big.Int, error) {
	return _Bingdings.Contract.Nonce(&_Bingdings.CallOpts)
}

// TotalScore is a free data retrieval call binding the contract method 0xe44c78e0.
//
// Solidity: function _totalScore() view returns(uint256)
func (_Bingdings *BingdingsCaller) TotalScore(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bingdings.contract.Call(opts, &out, "_totalScore")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalScore is a free data retrieval call binding the contract method 0xe44c78e0.
//
// Solidity: function _totalScore() view returns(uint256)
func (_Bingdings *BingdingsSession) TotalScore() (*big.Int, error) {
	return _Bingdings.Contract.TotalScore(&_Bingdings.CallOpts)
}

// TotalScore is a free data retrieval call binding the contract method 0xe44c78e0.
//
// Solidity: function _totalScore() view returns(uint256)
func (_Bingdings *BingdingsCallerSession) TotalScore() (*big.Int, error) {
	return _Bingdings.Contract.TotalScore(&_Bingdings.CallOpts)
}

// CalculateBonusTwoDecimals is a free data retrieval call binding the contract method 0x7c5d2ed9.
//
// Solidity: function calculateBonusTwoDecimals(uint256 _stake, uint256 totalWinning, uint256 totalLosing) pure returns(uint256 bonus)
func (_Bingdings *BingdingsCaller) CalculateBonusTwoDecimals(opts *bind.CallOpts, _stake *big.Int, totalWinning *big.Int, totalLosing *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Bingdings.contract.Call(opts, &out, "calculateBonusTwoDecimals", _stake, totalWinning, totalLosing)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateBonusTwoDecimals is a free data retrieval call binding the contract method 0x7c5d2ed9.
//
// Solidity: function calculateBonusTwoDecimals(uint256 _stake, uint256 totalWinning, uint256 totalLosing) pure returns(uint256 bonus)
func (_Bingdings *BingdingsSession) CalculateBonusTwoDecimals(_stake *big.Int, totalWinning *big.Int, totalLosing *big.Int) (*big.Int, error) {
	return _Bingdings.Contract.CalculateBonusTwoDecimals(&_Bingdings.CallOpts, _stake, totalWinning, totalLosing)
}

// CalculateBonusTwoDecimals is a free data retrieval call binding the contract method 0x7c5d2ed9.
//
// Solidity: function calculateBonusTwoDecimals(uint256 _stake, uint256 totalWinning, uint256 totalLosing) pure returns(uint256 bonus)
func (_Bingdings *BingdingsCallerSession) CalculateBonusTwoDecimals(_stake *big.Int, totalWinning *big.Int, totalLosing *big.Int) (*big.Int, error) {
	return _Bingdings.Contract.CalculateBonusTwoDecimals(&_Bingdings.CallOpts, _stake, totalWinning, totalLosing)
}

// EndAppend is a free data retrieval call binding the contract method 0xe5402bc4.
//
// Solidity: function endAppend() view returns(uint256)
func (_Bingdings *BingdingsCaller) EndAppend(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bingdings.contract.Call(opts, &out, "endAppend")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EndAppend is a free data retrieval call binding the contract method 0xe5402bc4.
//
// Solidity: function endAppend() view returns(uint256)
func (_Bingdings *BingdingsSession) EndAppend() (*big.Int, error) {
	return _Bingdings.Contract.EndAppend(&_Bingdings.CallOpts)
}

// EndAppend is a free data retrieval call binding the contract method 0xe5402bc4.
//
// Solidity: function endAppend() view returns(uint256)
func (_Bingdings *BingdingsCallerSession) EndAppend() (*big.Int, error) {
	return _Bingdings.Contract.EndAppend(&_Bingdings.CallOpts)
}

// FeeDecimals is a free data retrieval call binding the contract method 0xcc0f1786.
//
// Solidity: function feeDecimals() view returns(uint256)
func (_Bingdings *BingdingsCaller) FeeDecimals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bingdings.contract.Call(opts, &out, "feeDecimals")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeDecimals is a free data retrieval call binding the contract method 0xcc0f1786.
//
// Solidity: function feeDecimals() view returns(uint256)
func (_Bingdings *BingdingsSession) FeeDecimals() (*big.Int, error) {
	return _Bingdings.Contract.FeeDecimals(&_Bingdings.CallOpts)
}

// FeeDecimals is a free data retrieval call binding the contract method 0xcc0f1786.
//
// Solidity: function feeDecimals() view returns(uint256)
func (_Bingdings *BingdingsCallerSession) FeeDecimals() (*big.Int, error) {
	return _Bingdings.Contract.FeeDecimals(&_Bingdings.CallOpts)
}

// FeeRate is a free data retrieval call binding the contract method 0x978bbdb9.
//
// Solidity: function feeRate() view returns(uint256)
func (_Bingdings *BingdingsCaller) FeeRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bingdings.contract.Call(opts, &out, "feeRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeRate is a free data retrieval call binding the contract method 0x978bbdb9.
//
// Solidity: function feeRate() view returns(uint256)
func (_Bingdings *BingdingsSession) FeeRate() (*big.Int, error) {
	return _Bingdings.Contract.FeeRate(&_Bingdings.CallOpts)
}

// FeeRate is a free data retrieval call binding the contract method 0x978bbdb9.
//
// Solidity: function feeRate() view returns(uint256)
func (_Bingdings *BingdingsCallerSession) FeeRate() (*big.Int, error) {
	return _Bingdings.Contract.FeeRate(&_Bingdings.CallOpts)
}

// GameInfoList is a free data retrieval call binding the contract method 0xd979a56d.
//
// Solidity: function gameInfoList(uint256 ) view returns(uint256 id, string name, uint256 startTime, uint256 endTime, uint256 isSettled, uint256 player0Balance, uint256 player1Balance)
func (_Bingdings *BingdingsCaller) GameInfoList(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id             *big.Int
	Name           string
	StartTime      *big.Int
	EndTime        *big.Int
	IsSettled      *big.Int
	Player0Balance *big.Int
	Player1Balance *big.Int
}, error) {
	var out []interface{}
	err := _Bingdings.contract.Call(opts, &out, "gameInfoList", arg0)

	outstruct := new(struct {
		Id             *big.Int
		Name           string
		StartTime      *big.Int
		EndTime        *big.Int
		IsSettled      *big.Int
		Player0Balance *big.Int
		Player1Balance *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.StartTime = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.EndTime = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.IsSettled = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Player0Balance = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Player1Balance = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GameInfoList is a free data retrieval call binding the contract method 0xd979a56d.
//
// Solidity: function gameInfoList(uint256 ) view returns(uint256 id, string name, uint256 startTime, uint256 endTime, uint256 isSettled, uint256 player0Balance, uint256 player1Balance)
func (_Bingdings *BingdingsSession) GameInfoList(arg0 *big.Int) (struct {
	Id             *big.Int
	Name           string
	StartTime      *big.Int
	EndTime        *big.Int
	IsSettled      *big.Int
	Player0Balance *big.Int
	Player1Balance *big.Int
}, error) {
	return _Bingdings.Contract.GameInfoList(&_Bingdings.CallOpts, arg0)
}

// GameInfoList is a free data retrieval call binding the contract method 0xd979a56d.
//
// Solidity: function gameInfoList(uint256 ) view returns(uint256 id, string name, uint256 startTime, uint256 endTime, uint256 isSettled, uint256 player0Balance, uint256 player1Balance)
func (_Bingdings *BingdingsCallerSession) GameInfoList(arg0 *big.Int) (struct {
	Id             *big.Int
	Name           string
	StartTime      *big.Int
	EndTime        *big.Int
	IsSettled      *big.Int
	Player0Balance *big.Int
	Player1Balance *big.Int
}, error) {
	return _Bingdings.Contract.GameInfoList(&_Bingdings.CallOpts, arg0)
}

// Games is a free data retrieval call binding the contract method 0x117a5b90.
//
// Solidity: function games(uint256 ) view returns(string name, uint256 id, uint256 startTime, uint256 endTime, uint8 exist, uint8 isSettled, uint256 averageScore, uint256 index_slot)
func (_Bingdings *BingdingsCaller) Games(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Name         string
	Id           *big.Int
	StartTime    *big.Int
	EndTime      *big.Int
	Exist        uint8
	IsSettled    uint8
	AverageScore *big.Int
	IndexSlot    *big.Int
}, error) {
	var out []interface{}
	err := _Bingdings.contract.Call(opts, &out, "games", arg0)

	outstruct := new(struct {
		Name         string
		Id           *big.Int
		StartTime    *big.Int
		EndTime      *big.Int
		Exist        uint8
		IsSettled    uint8
		AverageScore *big.Int
		IndexSlot    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Id = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.StartTime = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.EndTime = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Exist = *abi.ConvertType(out[4], new(uint8)).(*uint8)
	outstruct.IsSettled = *abi.ConvertType(out[5], new(uint8)).(*uint8)
	outstruct.AverageScore = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.IndexSlot = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Games is a free data retrieval call binding the contract method 0x117a5b90.
//
// Solidity: function games(uint256 ) view returns(string name, uint256 id, uint256 startTime, uint256 endTime, uint8 exist, uint8 isSettled, uint256 averageScore, uint256 index_slot)
func (_Bingdings *BingdingsSession) Games(arg0 *big.Int) (struct {
	Name         string
	Id           *big.Int
	StartTime    *big.Int
	EndTime      *big.Int
	Exist        uint8
	IsSettled    uint8
	AverageScore *big.Int
	IndexSlot    *big.Int
}, error) {
	return _Bingdings.Contract.Games(&_Bingdings.CallOpts, arg0)
}

// Games is a free data retrieval call binding the contract method 0x117a5b90.
//
// Solidity: function games(uint256 ) view returns(string name, uint256 id, uint256 startTime, uint256 endTime, uint8 exist, uint8 isSettled, uint256 averageScore, uint256 index_slot)
func (_Bingdings *BingdingsCallerSession) Games(arg0 *big.Int) (struct {
	Name         string
	Id           *big.Int
	StartTime    *big.Int
	EndTime      *big.Int
	Exist        uint8
	IsSettled    uint8
	AverageScore *big.Int
	IndexSlot    *big.Int
}, error) {
	return _Bingdings.Contract.Games(&_Bingdings.CallOpts, arg0)
}

// GetAllGames is a free data retrieval call binding the contract method 0xdb1c45f9.
//
// Solidity: function getAllGames() view returns((uint256,string,address[],uint256,uint256,uint256,uint256,uint256)[])
func (_Bingdings *BingdingsCaller) GetAllGames(opts *bind.CallOpts) ([]MatchPgameInfo, error) {
	var out []interface{}
	err := _Bingdings.contract.Call(opts, &out, "getAllGames")

	if err != nil {
		return *new([]MatchPgameInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]MatchPgameInfo)).(*[]MatchPgameInfo)

	return out0, err

}

// GetAllGames is a free data retrieval call binding the contract method 0xdb1c45f9.
//
// Solidity: function getAllGames() view returns((uint256,string,address[],uint256,uint256,uint256,uint256,uint256)[])
func (_Bingdings *BingdingsSession) GetAllGames() ([]MatchPgameInfo, error) {
	return _Bingdings.Contract.GetAllGames(&_Bingdings.CallOpts)
}

// GetAllGames is a free data retrieval call binding the contract method 0xdb1c45f9.
//
// Solidity: function getAllGames() view returns((uint256,string,address[],uint256,uint256,uint256,uint256,uint256)[])
func (_Bingdings *BingdingsCallerSession) GetAllGames() ([]MatchPgameInfo, error) {
	return _Bingdings.Contract.GetAllGames(&_Bingdings.CallOpts)
}

// GetGameCount is a free data retrieval call binding the contract method 0xe87c0ee6.
//
// Solidity: function getGameCount() view returns(uint256)
func (_Bingdings *BingdingsCaller) GetGameCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bingdings.contract.Call(opts, &out, "getGameCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetGameCount is a free data retrieval call binding the contract method 0xe87c0ee6.
//
// Solidity: function getGameCount() view returns(uint256)
func (_Bingdings *BingdingsSession) GetGameCount() (*big.Int, error) {
	return _Bingdings.Contract.GetGameCount(&_Bingdings.CallOpts)
}

// GetGameCount is a free data retrieval call binding the contract method 0xe87c0ee6.
//
// Solidity: function getGameCount() view returns(uint256)
func (_Bingdings *BingdingsCallerSession) GetGameCount() (*big.Int, error) {
	return _Bingdings.Contract.GetGameCount(&_Bingdings.CallOpts)
}

// Index0 is a free data retrieval call binding the contract method 0x32c0defd.
//
// Solidity: function index0() view returns(uint256)
func (_Bingdings *BingdingsCaller) Index0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bingdings.contract.Call(opts, &out, "index0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Index0 is a free data retrieval call binding the contract method 0x32c0defd.
//
// Solidity: function index0() view returns(uint256)
func (_Bingdings *BingdingsSession) Index0() (*big.Int, error) {
	return _Bingdings.Contract.Index0(&_Bingdings.CallOpts)
}

// Index0 is a free data retrieval call binding the contract method 0x32c0defd.
//
// Solidity: function index0() view returns(uint256)
func (_Bingdings *BingdingsCallerSession) Index0() (*big.Int, error) {
	return _Bingdings.Contract.Index0(&_Bingdings.CallOpts)
}

// IsGain is a free data retrieval call binding the contract method 0x9f438a14.
//
// Solidity: function isGain(address ) view returns(bool)
func (_Bingdings *BingdingsCaller) IsGain(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Bingdings.contract.Call(opts, &out, "isGain", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsGain is a free data retrieval call binding the contract method 0x9f438a14.
//
// Solidity: function isGain(address ) view returns(bool)
func (_Bingdings *BingdingsSession) IsGain(arg0 common.Address) (bool, error) {
	return _Bingdings.Contract.IsGain(&_Bingdings.CallOpts, arg0)
}

// IsGain is a free data retrieval call binding the contract method 0x9f438a14.
//
// Solidity: function isGain(address ) view returns(bool)
func (_Bingdings *BingdingsCallerSession) IsGain(arg0 common.Address) (bool, error) {
	return _Bingdings.Contract.IsGain(&_Bingdings.CallOpts, arg0)
}

// IsPlayer is a free data retrieval call binding the contract method 0xe6d3538d.
//
// Solidity: function isPlayer(uint256 , address ) view returns(bool)
func (_Bingdings *BingdingsCaller) IsPlayer(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _Bingdings.contract.Call(opts, &out, "isPlayer", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPlayer is a free data retrieval call binding the contract method 0xe6d3538d.
//
// Solidity: function isPlayer(uint256 , address ) view returns(bool)
func (_Bingdings *BingdingsSession) IsPlayer(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _Bingdings.Contract.IsPlayer(&_Bingdings.CallOpts, arg0, arg1)
}

// IsPlayer is a free data retrieval call binding the contract method 0xe6d3538d.
//
// Solidity: function isPlayer(uint256 , address ) view returns(bool)
func (_Bingdings *BingdingsCallerSession) IsPlayer(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _Bingdings.Contract.IsPlayer(&_Bingdings.CallOpts, arg0, arg1)
}

// IsVoting is a free data retrieval call binding the contract method 0x97f72300.
//
// Solidity: function isVoting(uint256 _gameId) view returns(bool)
func (_Bingdings *BingdingsCaller) IsVoting(opts *bind.CallOpts, _gameId *big.Int) (bool, error) {
	var out []interface{}
	err := _Bingdings.contract.Call(opts, &out, "isVoting", _gameId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsVoting is a free data retrieval call binding the contract method 0x97f72300.
//
// Solidity: function isVoting(uint256 _gameId) view returns(bool)
func (_Bingdings *BingdingsSession) IsVoting(_gameId *big.Int) (bool, error) {
	return _Bingdings.Contract.IsVoting(&_Bingdings.CallOpts, _gameId)
}

// IsVoting is a free data retrieval call binding the contract method 0x97f72300.
//
// Solidity: function isVoting(uint256 _gameId) view returns(bool)
func (_Bingdings *BingdingsCallerSession) IsVoting(_gameId *big.Int) (bool, error) {
	return _Bingdings.Contract.IsVoting(&_Bingdings.CallOpts, _gameId)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bingdings *BingdingsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bingdings.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bingdings *BingdingsSession) Owner() (common.Address, error) {
	return _Bingdings.Contract.Owner(&_Bingdings.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bingdings *BingdingsCallerSession) Owner() (common.Address, error) {
	return _Bingdings.Contract.Owner(&_Bingdings.CallOpts)
}

// PlayerStakeBalance is a free data retrieval call binding the contract method 0xf131c153.
//
// Solidity: function playerStakeBalance(uint256 , address ) view returns(uint256)
func (_Bingdings *BingdingsCaller) PlayerStakeBalance(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bingdings.contract.Call(opts, &out, "playerStakeBalance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PlayerStakeBalance is a free data retrieval call binding the contract method 0xf131c153.
//
// Solidity: function playerStakeBalance(uint256 , address ) view returns(uint256)
func (_Bingdings *BingdingsSession) PlayerStakeBalance(arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	return _Bingdings.Contract.PlayerStakeBalance(&_Bingdings.CallOpts, arg0, arg1)
}

// PlayerStakeBalance is a free data retrieval call binding the contract method 0xf131c153.
//
// Solidity: function playerStakeBalance(uint256 , address ) view returns(uint256)
func (_Bingdings *BingdingsCallerSession) PlayerStakeBalance(arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	return _Bingdings.Contract.PlayerStakeBalance(&_Bingdings.CallOpts, arg0, arg1)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Bingdings *BingdingsCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Bingdings.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Bingdings *BingdingsSession) ProxiableUUID() ([32]byte, error) {
	return _Bingdings.Contract.ProxiableUUID(&_Bingdings.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Bingdings *BingdingsCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Bingdings.Contract.ProxiableUUID(&_Bingdings.CallOpts)
}

// StartAppend is a free data retrieval call binding the contract method 0xd768930a.
//
// Solidity: function startAppend() view returns(uint256)
func (_Bingdings *BingdingsCaller) StartAppend(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bingdings.contract.Call(opts, &out, "startAppend")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartAppend is a free data retrieval call binding the contract method 0xd768930a.
//
// Solidity: function startAppend() view returns(uint256)
func (_Bingdings *BingdingsSession) StartAppend() (*big.Int, error) {
	return _Bingdings.Contract.StartAppend(&_Bingdings.CallOpts)
}

// StartAppend is a free data retrieval call binding the contract method 0xd768930a.
//
// Solidity: function startAppend() view returns(uint256)
func (_Bingdings *BingdingsCallerSession) StartAppend() (*big.Int, error) {
	return _Bingdings.Contract.StartAppend(&_Bingdings.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Bingdings *BingdingsCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bingdings.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Bingdings *BingdingsSession) Token() (common.Address, error) {
	return _Bingdings.Contract.Token(&_Bingdings.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Bingdings *BingdingsCallerSession) Token() (common.Address, error) {
	return _Bingdings.Contract.Token(&_Bingdings.CallOpts)
}

// VoteBalance is a free data retrieval call binding the contract method 0x81269a18.
//
// Solidity: function voteBalance(address ) view returns(uint256)
func (_Bingdings *BingdingsCaller) VoteBalance(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bingdings.contract.Call(opts, &out, "voteBalance", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VoteBalance is a free data retrieval call binding the contract method 0x81269a18.
//
// Solidity: function voteBalance(address ) view returns(uint256)
func (_Bingdings *BingdingsSession) VoteBalance(arg0 common.Address) (*big.Int, error) {
	return _Bingdings.Contract.VoteBalance(&_Bingdings.CallOpts, arg0)
}

// VoteBalance is a free data retrieval call binding the contract method 0x81269a18.
//
// Solidity: function voteBalance(address ) view returns(uint256)
func (_Bingdings *BingdingsCallerSession) VoteBalance(arg0 common.Address) (*big.Int, error) {
	return _Bingdings.Contract.VoteBalance(&_Bingdings.CallOpts, arg0)
}

// VoterGameStake is a free data retrieval call binding the contract method 0xdb324918.
//
// Solidity: function voterGameStake(address , uint256 ) view returns(uint256 gameId, address player, uint256 stakeAmount)
func (_Bingdings *BingdingsCaller) VoterGameStake(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	GameId      *big.Int
	Player      common.Address
	StakeAmount *big.Int
}, error) {
	var out []interface{}
	err := _Bingdings.contract.Call(opts, &out, "voterGameStake", arg0, arg1)

	outstruct := new(struct {
		GameId      *big.Int
		Player      common.Address
		StakeAmount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.GameId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Player = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.StakeAmount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// VoterGameStake is a free data retrieval call binding the contract method 0xdb324918.
//
// Solidity: function voterGameStake(address , uint256 ) view returns(uint256 gameId, address player, uint256 stakeAmount)
func (_Bingdings *BingdingsSession) VoterGameStake(arg0 common.Address, arg1 *big.Int) (struct {
	GameId      *big.Int
	Player      common.Address
	StakeAmount *big.Int
}, error) {
	return _Bingdings.Contract.VoterGameStake(&_Bingdings.CallOpts, arg0, arg1)
}

// VoterGameStake is a free data retrieval call binding the contract method 0xdb324918.
//
// Solidity: function voterGameStake(address , uint256 ) view returns(uint256 gameId, address player, uint256 stakeAmount)
func (_Bingdings *BingdingsCallerSession) VoterGameStake(arg0 common.Address, arg1 *big.Int) (struct {
	GameId      *big.Int
	Player      common.Address
	StakeAmount *big.Int
}, error) {
	return _Bingdings.Contract.VoterGameStake(&_Bingdings.CallOpts, arg0, arg1)
}

// WhiteList is a free data retrieval call binding the contract method 0x372c12b1.
//
// Solidity: function whiteList(address ) view returns(bool)
func (_Bingdings *BingdingsCaller) WhiteList(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Bingdings.contract.Call(opts, &out, "whiteList", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WhiteList is a free data retrieval call binding the contract method 0x372c12b1.
//
// Solidity: function whiteList(address ) view returns(bool)
func (_Bingdings *BingdingsSession) WhiteList(arg0 common.Address) (bool, error) {
	return _Bingdings.Contract.WhiteList(&_Bingdings.CallOpts, arg0)
}

// WhiteList is a free data retrieval call binding the contract method 0x372c12b1.
//
// Solidity: function whiteList(address ) view returns(bool)
func (_Bingdings *BingdingsCallerSession) WhiteList(arg0 common.Address) (bool, error) {
	return _Bingdings.Contract.WhiteList(&_Bingdings.CallOpts, arg0)
}

// AddWhiteList is a paid mutator transaction binding the contract method 0xe7cd4a04.
//
// Solidity: function addWhiteList(address addr) returns(bool)
func (_Bingdings *BingdingsTransactor) AddWhiteList(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Bingdings.contract.Transact(opts, "addWhiteList", addr)
}

// AddWhiteList is a paid mutator transaction binding the contract method 0xe7cd4a04.
//
// Solidity: function addWhiteList(address addr) returns(bool)
func (_Bingdings *BingdingsSession) AddWhiteList(addr common.Address) (*types.Transaction, error) {
	return _Bingdings.Contract.AddWhiteList(&_Bingdings.TransactOpts, addr)
}

// AddWhiteList is a paid mutator transaction binding the contract method 0xe7cd4a04.
//
// Solidity: function addWhiteList(address addr) returns(bool)
func (_Bingdings *BingdingsTransactorSession) AddWhiteList(addr common.Address) (*types.Transaction, error) {
	return _Bingdings.Contract.AddWhiteList(&_Bingdings.TransactOpts, addr)
}

// CreateGame is a paid mutator transaction binding the contract method 0x0b6d715b.
//
// Solidity: function createGame(string _name, uint256 _startTime, uint256 _endTime) returns(uint256)
func (_Bingdings *BingdingsTransactor) CreateGame(opts *bind.TransactOpts, _name string, _startTime *big.Int, _endTime *big.Int) (*types.Transaction, error) {
	return _Bingdings.contract.Transact(opts, "createGame", _name, _startTime, _endTime)
}

// CreateGame is a paid mutator transaction binding the contract method 0x0b6d715b.
//
// Solidity: function createGame(string _name, uint256 _startTime, uint256 _endTime) returns(uint256)
func (_Bingdings *BingdingsSession) CreateGame(_name string, _startTime *big.Int, _endTime *big.Int) (*types.Transaction, error) {
	return _Bingdings.Contract.CreateGame(&_Bingdings.TransactOpts, _name, _startTime, _endTime)
}

// CreateGame is a paid mutator transaction binding the contract method 0x0b6d715b.
//
// Solidity: function createGame(string _name, uint256 _startTime, uint256 _endTime) returns(uint256)
func (_Bingdings *BingdingsTransactorSession) CreateGame(_name string, _startTime *big.Int, _endTime *big.Int) (*types.Transaction, error) {
	return _Bingdings.Contract.CreateGame(&_Bingdings.TransactOpts, _name, _startTime, _endTime)
}

// ForceGameOver is a paid mutator transaction binding the contract method 0xb9b11e0e.
//
// Solidity: function forceGameOver(uint256 _gameId) returns(bool)
func (_Bingdings *BingdingsTransactor) ForceGameOver(opts *bind.TransactOpts, _gameId *big.Int) (*types.Transaction, error) {
	return _Bingdings.contract.Transact(opts, "forceGameOver", _gameId)
}

// ForceGameOver is a paid mutator transaction binding the contract method 0xb9b11e0e.
//
// Solidity: function forceGameOver(uint256 _gameId) returns(bool)
func (_Bingdings *BingdingsSession) ForceGameOver(_gameId *big.Int) (*types.Transaction, error) {
	return _Bingdings.Contract.ForceGameOver(&_Bingdings.TransactOpts, _gameId)
}

// ForceGameOver is a paid mutator transaction binding the contract method 0xb9b11e0e.
//
// Solidity: function forceGameOver(uint256 _gameId) returns(bool)
func (_Bingdings *BingdingsTransactorSession) ForceGameOver(_gameId *big.Int) (*types.Transaction, error) {
	return _Bingdings.Contract.ForceGameOver(&_Bingdings.TransactOpts, _gameId)
}

// GetToken is a paid mutator transaction binding the contract method 0x21df0da7.
//
// Solidity: function getToken() returns(bool)
func (_Bingdings *BingdingsTransactor) GetToken(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bingdings.contract.Transact(opts, "getToken")
}

// GetToken is a paid mutator transaction binding the contract method 0x21df0da7.
//
// Solidity: function getToken() returns(bool)
func (_Bingdings *BingdingsSession) GetToken() (*types.Transaction, error) {
	return _Bingdings.Contract.GetToken(&_Bingdings.TransactOpts)
}

// GetToken is a paid mutator transaction binding the contract method 0x21df0da7.
//
// Solidity: function getToken() returns(bool)
func (_Bingdings *BingdingsTransactorSession) GetToken() (*types.Transaction, error) {
	return _Bingdings.Contract.GetToken(&_Bingdings.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x7a1ac61e.
//
// Solidity: function initialize(address protocol, uint256 st, uint256 et) returns()
func (_Bingdings *BingdingsTransactor) Initialize(opts *bind.TransactOpts, protocol common.Address, st *big.Int, et *big.Int) (*types.Transaction, error) {
	return _Bingdings.contract.Transact(opts, "initialize", protocol, st, et)
}

// Initialize is a paid mutator transaction binding the contract method 0x7a1ac61e.
//
// Solidity: function initialize(address protocol, uint256 st, uint256 et) returns()
func (_Bingdings *BingdingsSession) Initialize(protocol common.Address, st *big.Int, et *big.Int) (*types.Transaction, error) {
	return _Bingdings.Contract.Initialize(&_Bingdings.TransactOpts, protocol, st, et)
}

// Initialize is a paid mutator transaction binding the contract method 0x7a1ac61e.
//
// Solidity: function initialize(address protocol, uint256 st, uint256 et) returns()
func (_Bingdings *BingdingsTransactorSession) Initialize(protocol common.Address, st *big.Int, et *big.Int) (*types.Transaction, error) {
	return _Bingdings.Contract.Initialize(&_Bingdings.TransactOpts, protocol, st, et)
}

// JoinGame is a paid mutator transaction binding the contract method 0xefaa55a0.
//
// Solidity: function joinGame(uint256 _gameId) returns(bool)
func (_Bingdings *BingdingsTransactor) JoinGame(opts *bind.TransactOpts, _gameId *big.Int) (*types.Transaction, error) {
	return _Bingdings.contract.Transact(opts, "joinGame", _gameId)
}

// JoinGame is a paid mutator transaction binding the contract method 0xefaa55a0.
//
// Solidity: function joinGame(uint256 _gameId) returns(bool)
func (_Bingdings *BingdingsSession) JoinGame(_gameId *big.Int) (*types.Transaction, error) {
	return _Bingdings.Contract.JoinGame(&_Bingdings.TransactOpts, _gameId)
}

// JoinGame is a paid mutator transaction binding the contract method 0xefaa55a0.
//
// Solidity: function joinGame(uint256 _gameId) returns(bool)
func (_Bingdings *BingdingsTransactorSession) JoinGame(_gameId *big.Int) (*types.Transaction, error) {
	return _Bingdings.Contract.JoinGame(&_Bingdings.TransactOpts, _gameId)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bingdings *BingdingsTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bingdings.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bingdings *BingdingsSession) RenounceOwnership() (*types.Transaction, error) {
	return _Bingdings.Contract.RenounceOwnership(&_Bingdings.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bingdings *BingdingsTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Bingdings.Contract.RenounceOwnership(&_Bingdings.TransactOpts)
}

// SetToken is a paid mutator transaction binding the contract method 0x144fa6d7.
//
// Solidity: function setToken(address _token) returns()
func (_Bingdings *BingdingsTransactor) SetToken(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _Bingdings.contract.Transact(opts, "setToken", _token)
}

// SetToken is a paid mutator transaction binding the contract method 0x144fa6d7.
//
// Solidity: function setToken(address _token) returns()
func (_Bingdings *BingdingsSession) SetToken(_token common.Address) (*types.Transaction, error) {
	return _Bingdings.Contract.SetToken(&_Bingdings.TransactOpts, _token)
}

// SetToken is a paid mutator transaction binding the contract method 0x144fa6d7.
//
// Solidity: function setToken(address _token) returns()
func (_Bingdings *BingdingsTransactorSession) SetToken(_token common.Address) (*types.Transaction, error) {
	return _Bingdings.Contract.SetToken(&_Bingdings.TransactOpts, _token)
}

// Settlement is a paid mutator transaction binding the contract method 0x60ad2391.
//
// Solidity: function settlement(uint256 _gameId) returns(bool)
func (_Bingdings *BingdingsTransactor) Settlement(opts *bind.TransactOpts, _gameId *big.Int) (*types.Transaction, error) {
	return _Bingdings.contract.Transact(opts, "settlement", _gameId)
}

// Settlement is a paid mutator transaction binding the contract method 0x60ad2391.
//
// Solidity: function settlement(uint256 _gameId) returns(bool)
func (_Bingdings *BingdingsSession) Settlement(_gameId *big.Int) (*types.Transaction, error) {
	return _Bingdings.Contract.Settlement(&_Bingdings.TransactOpts, _gameId)
}

// Settlement is a paid mutator transaction binding the contract method 0x60ad2391.
//
// Solidity: function settlement(uint256 _gameId) returns(bool)
func (_Bingdings *BingdingsTransactorSession) Settlement(_gameId *big.Int) (*types.Transaction, error) {
	return _Bingdings.Contract.Settlement(&_Bingdings.TransactOpts, _gameId)
}

// Stake is a paid mutator transaction binding the contract method 0xadc9772e.
//
// Solidity: function stake(address _staker, uint256 _amount) returns(bool)
func (_Bingdings *BingdingsTransactor) Stake(opts *bind.TransactOpts, _staker common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Bingdings.contract.Transact(opts, "stake", _staker, _amount)
}

// Stake is a paid mutator transaction binding the contract method 0xadc9772e.
//
// Solidity: function stake(address _staker, uint256 _amount) returns(bool)
func (_Bingdings *BingdingsSession) Stake(_staker common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Bingdings.Contract.Stake(&_Bingdings.TransactOpts, _staker, _amount)
}

// Stake is a paid mutator transaction binding the contract method 0xadc9772e.
//
// Solidity: function stake(address _staker, uint256 _amount) returns(bool)
func (_Bingdings *BingdingsTransactorSession) Stake(_staker common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Bingdings.Contract.Stake(&_Bingdings.TransactOpts, _staker, _amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bingdings *BingdingsTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Bingdings.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bingdings *BingdingsSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Bingdings.Contract.TransferOwnership(&_Bingdings.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bingdings *BingdingsTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Bingdings.Contract.TransferOwnership(&_Bingdings.TransactOpts, newOwner)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Bingdings *BingdingsTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Bingdings.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Bingdings *BingdingsSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Bingdings.Contract.UpgradeToAndCall(&_Bingdings.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Bingdings *BingdingsTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Bingdings.Contract.UpgradeToAndCall(&_Bingdings.TransactOpts, newImplementation, data)
}

// Vote is a paid mutator transaction binding the contract method 0x4de8737c.
//
// Solidity: function vote(uint256 _gameId, address _player, uint256 amount) returns(bool)
func (_Bingdings *BingdingsTransactor) Vote(opts *bind.TransactOpts, _gameId *big.Int, _player common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bingdings.contract.Transact(opts, "vote", _gameId, _player, amount)
}

// Vote is a paid mutator transaction binding the contract method 0x4de8737c.
//
// Solidity: function vote(uint256 _gameId, address _player, uint256 amount) returns(bool)
func (_Bingdings *BingdingsSession) Vote(_gameId *big.Int, _player common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bingdings.Contract.Vote(&_Bingdings.TransactOpts, _gameId, _player, amount)
}

// Vote is a paid mutator transaction binding the contract method 0x4de8737c.
//
// Solidity: function vote(uint256 _gameId, address _player, uint256 amount) returns(bool)
func (_Bingdings *BingdingsTransactorSession) Vote(_gameId *big.Int, _player common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bingdings.Contract.Vote(&_Bingdings.TransactOpts, _gameId, _player, amount)
}

// BingdingsCreateGameIterator is returned from FilterCreateGame and is used to iterate over the raw logs and unpacked data for CreateGame events raised by the Bingdings contract.
type BingdingsCreateGameIterator struct {
	Event *BingdingsCreateGame // Event containing the contract specifics and raw log

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
func (it *BingdingsCreateGameIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BingdingsCreateGame)
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
		it.Event = new(BingdingsCreateGame)
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
func (it *BingdingsCreateGameIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BingdingsCreateGameIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BingdingsCreateGame represents a CreateGame event raised by the Bingdings contract.
type BingdingsCreateGame struct {
	GameId    *big.Int
	Name      common.Hash
	StartTime *big.Int
	EndTime   *big.Int
	Builder   common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCreateGame is a free log retrieval operation binding the contract event 0x1512862c5902d1ca8ca4ec660f29eb674347a39d327974dc2f50b428e204879e.
//
// Solidity: event CreateGame(uint256 indexed gameId, string indexed _name, uint256 startTime, uint256 endTime, address builder)
func (_Bingdings *BingdingsFilterer) FilterCreateGame(opts *bind.FilterOpts, gameId []*big.Int, _name []string) (*BingdingsCreateGameIterator, error) {

	var gameIdRule []interface{}
	for _, gameIdItem := range gameId {
		gameIdRule = append(gameIdRule, gameIdItem)
	}
	var _nameRule []interface{}
	for _, _nameItem := range _name {
		_nameRule = append(_nameRule, _nameItem)
	}

	logs, sub, err := _Bingdings.contract.FilterLogs(opts, "CreateGame", gameIdRule, _nameRule)
	if err != nil {
		return nil, err
	}
	return &BingdingsCreateGameIterator{contract: _Bingdings.contract, event: "CreateGame", logs: logs, sub: sub}, nil
}

// WatchCreateGame is a free log subscription operation binding the contract event 0x1512862c5902d1ca8ca4ec660f29eb674347a39d327974dc2f50b428e204879e.
//
// Solidity: event CreateGame(uint256 indexed gameId, string indexed _name, uint256 startTime, uint256 endTime, address builder)
func (_Bingdings *BingdingsFilterer) WatchCreateGame(opts *bind.WatchOpts, sink chan<- *BingdingsCreateGame, gameId []*big.Int, _name []string) (event.Subscription, error) {

	var gameIdRule []interface{}
	for _, gameIdItem := range gameId {
		gameIdRule = append(gameIdRule, gameIdItem)
	}
	var _nameRule []interface{}
	for _, _nameItem := range _name {
		_nameRule = append(_nameRule, _nameItem)
	}

	logs, sub, err := _Bingdings.contract.WatchLogs(opts, "CreateGame", gameIdRule, _nameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BingdingsCreateGame)
				if err := _Bingdings.contract.UnpackLog(event, "CreateGame", log); err != nil {
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

// ParseCreateGame is a log parse operation binding the contract event 0x1512862c5902d1ca8ca4ec660f29eb674347a39d327974dc2f50b428e204879e.
//
// Solidity: event CreateGame(uint256 indexed gameId, string indexed _name, uint256 startTime, uint256 endTime, address builder)
func (_Bingdings *BingdingsFilterer) ParseCreateGame(log types.Log) (*BingdingsCreateGame, error) {
	event := new(BingdingsCreateGame)
	if err := _Bingdings.contract.UnpackLog(event, "CreateGame", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BingdingsDoNothingIterator is returned from FilterDoNothing and is used to iterate over the raw logs and unpacked data for DoNothing events raised by the Bingdings contract.
type BingdingsDoNothingIterator struct {
	Event *BingdingsDoNothing // Event containing the contract specifics and raw log

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
func (it *BingdingsDoNothingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BingdingsDoNothing)
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
		it.Event = new(BingdingsDoNothing)
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
func (it *BingdingsDoNothingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BingdingsDoNothingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BingdingsDoNothing represents a DoNothing event raised by the Bingdings contract.
type BingdingsDoNothing struct {
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDoNothing is a free log retrieval operation binding the contract event 0x81731ee8979ae4adf353227f954a9f4732a1fbfa4be38b31aff6619a5cdbe4a0.
//
// Solidity: event DoNothing(address indexed sender)
func (_Bingdings *BingdingsFilterer) FilterDoNothing(opts *bind.FilterOpts, sender []common.Address) (*BingdingsDoNothingIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Bingdings.contract.FilterLogs(opts, "DoNothing", senderRule)
	if err != nil {
		return nil, err
	}
	return &BingdingsDoNothingIterator{contract: _Bingdings.contract, event: "DoNothing", logs: logs, sub: sub}, nil
}

// WatchDoNothing is a free log subscription operation binding the contract event 0x81731ee8979ae4adf353227f954a9f4732a1fbfa4be38b31aff6619a5cdbe4a0.
//
// Solidity: event DoNothing(address indexed sender)
func (_Bingdings *BingdingsFilterer) WatchDoNothing(opts *bind.WatchOpts, sink chan<- *BingdingsDoNothing, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Bingdings.contract.WatchLogs(opts, "DoNothing", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BingdingsDoNothing)
				if err := _Bingdings.contract.UnpackLog(event, "DoNothing", log); err != nil {
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

// ParseDoNothing is a log parse operation binding the contract event 0x81731ee8979ae4adf353227f954a9f4732a1fbfa4be38b31aff6619a5cdbe4a0.
//
// Solidity: event DoNothing(address indexed sender)
func (_Bingdings *BingdingsFilterer) ParseDoNothing(log types.Log) (*BingdingsDoNothing, error) {
	event := new(BingdingsDoNothing)
	if err := _Bingdings.contract.UnpackLog(event, "DoNothing", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BingdingsInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Bingdings contract.
type BingdingsInitializedIterator struct {
	Event *BingdingsInitialized // Event containing the contract specifics and raw log

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
func (it *BingdingsInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BingdingsInitialized)
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
		it.Event = new(BingdingsInitialized)
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
func (it *BingdingsInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BingdingsInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BingdingsInitialized represents a Initialized event raised by the Bingdings contract.
type BingdingsInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Bingdings *BingdingsFilterer) FilterInitialized(opts *bind.FilterOpts) (*BingdingsInitializedIterator, error) {

	logs, sub, err := _Bingdings.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &BingdingsInitializedIterator{contract: _Bingdings.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Bingdings *BingdingsFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *BingdingsInitialized) (event.Subscription, error) {

	logs, sub, err := _Bingdings.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BingdingsInitialized)
				if err := _Bingdings.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Bingdings *BingdingsFilterer) ParseInitialized(log types.Log) (*BingdingsInitialized, error) {
	event := new(BingdingsInitialized)
	if err := _Bingdings.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BingdingsJoinGameIterator is returned from FilterJoinGame and is used to iterate over the raw logs and unpacked data for JoinGame events raised by the Bingdings contract.
type BingdingsJoinGameIterator struct {
	Event *BingdingsJoinGame // Event containing the contract specifics and raw log

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
func (it *BingdingsJoinGameIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BingdingsJoinGame)
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
		it.Event = new(BingdingsJoinGame)
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
func (it *BingdingsJoinGameIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BingdingsJoinGameIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BingdingsJoinGame represents a JoinGame event raised by the Bingdings contract.
type BingdingsJoinGame struct {
	GameId *big.Int
	Player common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterJoinGame is a free log retrieval operation binding the contract event 0x94a23cd53bf5ba1d2903ed9a1cd6418ae0d681ab5bbae27df124f57f70a9b627.
//
// Solidity: event JoinGame(uint256 gameId, address player)
func (_Bingdings *BingdingsFilterer) FilterJoinGame(opts *bind.FilterOpts) (*BingdingsJoinGameIterator, error) {

	logs, sub, err := _Bingdings.contract.FilterLogs(opts, "JoinGame")
	if err != nil {
		return nil, err
	}
	return &BingdingsJoinGameIterator{contract: _Bingdings.contract, event: "JoinGame", logs: logs, sub: sub}, nil
}

// WatchJoinGame is a free log subscription operation binding the contract event 0x94a23cd53bf5ba1d2903ed9a1cd6418ae0d681ab5bbae27df124f57f70a9b627.
//
// Solidity: event JoinGame(uint256 gameId, address player)
func (_Bingdings *BingdingsFilterer) WatchJoinGame(opts *bind.WatchOpts, sink chan<- *BingdingsJoinGame) (event.Subscription, error) {

	logs, sub, err := _Bingdings.contract.WatchLogs(opts, "JoinGame")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BingdingsJoinGame)
				if err := _Bingdings.contract.UnpackLog(event, "JoinGame", log); err != nil {
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

// ParseJoinGame is a log parse operation binding the contract event 0x94a23cd53bf5ba1d2903ed9a1cd6418ae0d681ab5bbae27df124f57f70a9b627.
//
// Solidity: event JoinGame(uint256 gameId, address player)
func (_Bingdings *BingdingsFilterer) ParseJoinGame(log types.Log) (*BingdingsJoinGame, error) {
	event := new(BingdingsJoinGame)
	if err := _Bingdings.contract.UnpackLog(event, "JoinGame", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BingdingsOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Bingdings contract.
type BingdingsOwnershipTransferredIterator struct {
	Event *BingdingsOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BingdingsOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BingdingsOwnershipTransferred)
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
		it.Event = new(BingdingsOwnershipTransferred)
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
func (it *BingdingsOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BingdingsOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BingdingsOwnershipTransferred represents a OwnershipTransferred event raised by the Bingdings contract.
type BingdingsOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Bingdings *BingdingsFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BingdingsOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Bingdings.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BingdingsOwnershipTransferredIterator{contract: _Bingdings.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Bingdings *BingdingsFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BingdingsOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Bingdings.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BingdingsOwnershipTransferred)
				if err := _Bingdings.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Bingdings *BingdingsFilterer) ParseOwnershipTransferred(log types.Log) (*BingdingsOwnershipTransferred, error) {
	event := new(BingdingsOwnershipTransferred)
	if err := _Bingdings.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BingdingsSettleIterator is returned from FilterSettle and is used to iterate over the raw logs and unpacked data for Settle events raised by the Bingdings contract.
type BingdingsSettleIterator struct {
	Event *BingdingsSettle // Event containing the contract specifics and raw log

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
func (it *BingdingsSettleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BingdingsSettle)
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
		it.Event = new(BingdingsSettle)
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
func (it *BingdingsSettleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BingdingsSettleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BingdingsSettle represents a Settle event raised by the Bingdings contract.
type BingdingsSettle struct {
	GameId      *big.Int
	WinnerVoter common.Address
	Time        *big.Int
	Bonus       *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSettle is a free log retrieval operation binding the contract event 0xefaa342df122305c6fde76e0644963c2d3c478219dadfaace396521c1902d302.
//
// Solidity: event Settle(uint256 gameId, address winnerVoter, uint256 _time, uint256 _bonus)
func (_Bingdings *BingdingsFilterer) FilterSettle(opts *bind.FilterOpts) (*BingdingsSettleIterator, error) {

	logs, sub, err := _Bingdings.contract.FilterLogs(opts, "Settle")
	if err != nil {
		return nil, err
	}
	return &BingdingsSettleIterator{contract: _Bingdings.contract, event: "Settle", logs: logs, sub: sub}, nil
}

// WatchSettle is a free log subscription operation binding the contract event 0xefaa342df122305c6fde76e0644963c2d3c478219dadfaace396521c1902d302.
//
// Solidity: event Settle(uint256 gameId, address winnerVoter, uint256 _time, uint256 _bonus)
func (_Bingdings *BingdingsFilterer) WatchSettle(opts *bind.WatchOpts, sink chan<- *BingdingsSettle) (event.Subscription, error) {

	logs, sub, err := _Bingdings.contract.WatchLogs(opts, "Settle")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BingdingsSettle)
				if err := _Bingdings.contract.UnpackLog(event, "Settle", log); err != nil {
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

// ParseSettle is a log parse operation binding the contract event 0xefaa342df122305c6fde76e0644963c2d3c478219dadfaace396521c1902d302.
//
// Solidity: event Settle(uint256 gameId, address winnerVoter, uint256 _time, uint256 _bonus)
func (_Bingdings *BingdingsFilterer) ParseSettle(log types.Log) (*BingdingsSettle, error) {
	event := new(BingdingsSettle)
	if err := _Bingdings.contract.UnpackLog(event, "Settle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BingdingsSettledIterator is returned from FilterSettled and is used to iterate over the raw logs and unpacked data for Settled events raised by the Bingdings contract.
type BingdingsSettledIterator struct {
	Event *BingdingsSettled // Event containing the contract specifics and raw log

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
func (it *BingdingsSettledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BingdingsSettled)
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
		it.Event = new(BingdingsSettled)
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
func (it *BingdingsSettledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BingdingsSettledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BingdingsSettled represents a Settled event raised by the Bingdings contract.
type BingdingsSettled struct {
	GameId *big.Int
	Winner common.Address
	Time   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSettled is a free log retrieval operation binding the contract event 0x2c35d68fdf40b18e913bb877373b4a4fc67810e2546dc5c9f9208eb8494057cb.
//
// Solidity: event Settled(uint256 gameId, address winner, uint256 _time)
func (_Bingdings *BingdingsFilterer) FilterSettled(opts *bind.FilterOpts) (*BingdingsSettledIterator, error) {

	logs, sub, err := _Bingdings.contract.FilterLogs(opts, "Settled")
	if err != nil {
		return nil, err
	}
	return &BingdingsSettledIterator{contract: _Bingdings.contract, event: "Settled", logs: logs, sub: sub}, nil
}

// WatchSettled is a free log subscription operation binding the contract event 0x2c35d68fdf40b18e913bb877373b4a4fc67810e2546dc5c9f9208eb8494057cb.
//
// Solidity: event Settled(uint256 gameId, address winner, uint256 _time)
func (_Bingdings *BingdingsFilterer) WatchSettled(opts *bind.WatchOpts, sink chan<- *BingdingsSettled) (event.Subscription, error) {

	logs, sub, err := _Bingdings.contract.WatchLogs(opts, "Settled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BingdingsSettled)
				if err := _Bingdings.contract.UnpackLog(event, "Settled", log); err != nil {
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

// ParseSettled is a log parse operation binding the contract event 0x2c35d68fdf40b18e913bb877373b4a4fc67810e2546dc5c9f9208eb8494057cb.
//
// Solidity: event Settled(uint256 gameId, address winner, uint256 _time)
func (_Bingdings *BingdingsFilterer) ParseSettled(log types.Log) (*BingdingsSettled, error) {
	event := new(BingdingsSettled)
	if err := _Bingdings.contract.UnpackLog(event, "Settled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BingdingsUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Bingdings contract.
type BingdingsUpgradedIterator struct {
	Event *BingdingsUpgraded // Event containing the contract specifics and raw log

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
func (it *BingdingsUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BingdingsUpgraded)
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
		it.Event = new(BingdingsUpgraded)
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
func (it *BingdingsUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BingdingsUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BingdingsUpgraded represents a Upgraded event raised by the Bingdings contract.
type BingdingsUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Bingdings *BingdingsFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*BingdingsUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Bingdings.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &BingdingsUpgradedIterator{contract: _Bingdings.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Bingdings *BingdingsFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *BingdingsUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Bingdings.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BingdingsUpgraded)
				if err := _Bingdings.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Bingdings *BingdingsFilterer) ParseUpgraded(log types.Log) (*BingdingsUpgraded, error) {
	event := new(BingdingsUpgraded)
	if err := _Bingdings.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BingdingsVoteIterator is returned from FilterVote and is used to iterate over the raw logs and unpacked data for Vote events raised by the Bingdings contract.
type BingdingsVoteIterator struct {
	Event *BingdingsVote // Event containing the contract specifics and raw log

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
func (it *BingdingsVoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BingdingsVote)
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
		it.Event = new(BingdingsVote)
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
func (it *BingdingsVoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BingdingsVoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BingdingsVote represents a Vote event raised by the Bingdings contract.
type BingdingsVote struct {
	GameId *big.Int
	Voter  common.Address
	Player common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterVote is a free log retrieval operation binding the contract event 0xd3e566fa4d4d8b17c007ba24b9dbb382c9185a03c9ef59477ee3c6132ca0a306.
//
// Solidity: event Vote(uint256 gameId, address voter, address player, uint256 amount)
func (_Bingdings *BingdingsFilterer) FilterVote(opts *bind.FilterOpts) (*BingdingsVoteIterator, error) {

	logs, sub, err := _Bingdings.contract.FilterLogs(opts, "Vote")
	if err != nil {
		return nil, err
	}
	return &BingdingsVoteIterator{contract: _Bingdings.contract, event: "Vote", logs: logs, sub: sub}, nil
}

// WatchVote is a free log subscription operation binding the contract event 0xd3e566fa4d4d8b17c007ba24b9dbb382c9185a03c9ef59477ee3c6132ca0a306.
//
// Solidity: event Vote(uint256 gameId, address voter, address player, uint256 amount)
func (_Bingdings *BingdingsFilterer) WatchVote(opts *bind.WatchOpts, sink chan<- *BingdingsVote) (event.Subscription, error) {

	logs, sub, err := _Bingdings.contract.WatchLogs(opts, "Vote")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BingdingsVote)
				if err := _Bingdings.contract.UnpackLog(event, "Vote", log); err != nil {
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

// ParseVote is a log parse operation binding the contract event 0xd3e566fa4d4d8b17c007ba24b9dbb382c9185a03c9ef59477ee3c6132ca0a306.
//
// Solidity: event Vote(uint256 gameId, address voter, address player, uint256 amount)
func (_Bingdings *BingdingsFilterer) ParseVote(log types.Log) (*BingdingsVote, error) {
	event := new(BingdingsVote)
	if err := _Bingdings.contract.UnpackLog(event, "Vote", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
