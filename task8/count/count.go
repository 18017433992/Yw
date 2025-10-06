// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package count

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

// CountMetaData contains all meta data concerning the Count contract.
var CountMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_version\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"CountNumber\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CountNmb\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f5ffd5b506040516107cb3803806107cb83398181016040528101906100319190610194565b806001908161004091906103eb565b50506104ba565b5f604051905090565b5f5ffd5b5f5ffd5b5f5ffd5b5f5ffd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6100a682610060565b810181811067ffffffffffffffff821117156100c5576100c4610070565b5b80604052505050565b5f6100d7610047565b90506100e3828261009d565b919050565b5f67ffffffffffffffff82111561010257610101610070565b5b61010b82610060565b9050602081019050919050565b8281835e5f83830152505050565b5f610138610133846100e8565b6100ce565b9050828152602081018484840111156101545761015361005c565b5b61015f848285610118565b509392505050565b5f82601f83011261017b5761017a610058565b5b815161018b848260208601610126565b91505092915050565b5f602082840312156101a9576101a8610050565b5b5f82015167ffffffffffffffff8111156101c6576101c5610054565b5b6101d284828501610167565b91505092915050565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061022957607f821691505b60208210810361023c5761023b6101e5565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f6008830261029e7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610263565b6102a88683610263565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f6102ec6102e76102e2846102c0565b6102c9565b6102c0565b9050919050565b5f819050919050565b610305836102d2565b610319610311826102f3565b84845461026f565b825550505050565b5f5f905090565b610330610321565b61033b8184846102fc565b505050565b5b8181101561035e576103535f82610328565b600181019050610341565b5050565b601f8211156103a35761037481610242565b61037d84610254565b8101602085101561038c578190505b6103a061039885610254565b830182610340565b50505b505050565b5f82821c905092915050565b5f6103c35f19846008026103a8565b1980831691505092915050565b5f6103db83836103b4565b9150826002028217905092915050565b6103f4826101db565b67ffffffffffffffff81111561040d5761040c610070565b5b6104178254610212565b610422828285610362565b5f60209050601f831160018114610453575f8415610441578287015190505b61044b85826103d0565b8655506104b2565b601f19841661046186610242565b5f5b8281101561048857848901518255600182019150602085019450602081019050610463565b868310156104a557848901516104a1601f8916826103b4565b8355505b6001600288020188555050505b505050505050565b610304806104c75f395ff3fe608060405234801561000f575f5ffd5b5060043610610034575f3560e01c806354fd4d5014610038578063e211945014610056575b5f5ffd5b610040610060565b60405161004d91906101ac565b60405180910390f35b61005e6100ec565b005b6001805461006d906101f9565b80601f0160208091040260200160405190810160405280929190818152602001828054610099906101f9565b80156100e45780601f106100bb576101008083540402835291602001916100e4565b820191905f5260205f20905b8154815290600101906020018083116100c757829003601f168201915b505050505081565b5f5f8154809291906100fd9061025f565b91905055507feca011e810d3353e4c1e9308544e45836b238ae8a7330a7e53ca8b4e03810df95f5460405161013291906102b5565b60405180910390a1565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f61017e8261013c565b6101888185610146565b9350610198818560208601610156565b6101a181610164565b840191505092915050565b5f6020820190508181035f8301526101c48184610174565b905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061021057607f821691505b602082108103610223576102226101cc565b5b50919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f819050919050565b5f61026982610256565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361029b5761029a610229565b5b600182019050919050565b6102af81610256565b82525050565b5f6020820190506102c85f8301846102a6565b9291505056fea2646970667358221220858f96fefe3ee6aa9428f16c9c414b9cb0c90330e6c7fe3012d525b46794787e64736f6c634300081e0033",
}

// CountABI is the input ABI used to generate the binding from.
// Deprecated: Use CountMetaData.ABI instead.
var CountABI = CountMetaData.ABI

// CountBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CountMetaData.Bin instead.
var CountBin = CountMetaData.Bin

// DeployCount deploys a new Ethereum contract, binding an instance of Count to it.
func DeployCount(auth *bind.TransactOpts, backend bind.ContractBackend, _version string) (common.Address, *types.Transaction, *Count, error) {
	parsed, err := CountMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CountBin), backend, _version)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Count{CountCaller: CountCaller{contract: contract}, CountTransactor: CountTransactor{contract: contract}, CountFilterer: CountFilterer{contract: contract}}, nil
}

// Count is an auto generated Go binding around an Ethereum contract.
type Count struct {
	CountCaller     // Read-only binding to the contract
	CountTransactor // Write-only binding to the contract
	CountFilterer   // Log filterer for contract events
}

// CountCaller is an auto generated read-only Go binding around an Ethereum contract.
type CountCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CountTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CountTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CountFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CountFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CountSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CountSession struct {
	Contract     *Count            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CountCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CountCallerSession struct {
	Contract *CountCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// CountTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CountTransactorSession struct {
	Contract     *CountTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CountRaw is an auto generated low-level Go binding around an Ethereum contract.
type CountRaw struct {
	Contract *Count // Generic contract binding to access the raw methods on
}

// CountCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CountCallerRaw struct {
	Contract *CountCaller // Generic read-only contract binding to access the raw methods on
}

// CountTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CountTransactorRaw struct {
	Contract *CountTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCount creates a new instance of Count, bound to a specific deployed contract.
func NewCount(address common.Address, backend bind.ContractBackend) (*Count, error) {
	contract, err := bindCount(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Count{CountCaller: CountCaller{contract: contract}, CountTransactor: CountTransactor{contract: contract}, CountFilterer: CountFilterer{contract: contract}}, nil
}

// NewCountCaller creates a new read-only instance of Count, bound to a specific deployed contract.
func NewCountCaller(address common.Address, caller bind.ContractCaller) (*CountCaller, error) {
	contract, err := bindCount(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CountCaller{contract: contract}, nil
}

// NewCountTransactor creates a new write-only instance of Count, bound to a specific deployed contract.
func NewCountTransactor(address common.Address, transactor bind.ContractTransactor) (*CountTransactor, error) {
	contract, err := bindCount(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CountTransactor{contract: contract}, nil
}

// NewCountFilterer creates a new log filterer instance of Count, bound to a specific deployed contract.
func NewCountFilterer(address common.Address, filterer bind.ContractFilterer) (*CountFilterer, error) {
	contract, err := bindCount(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CountFilterer{contract: contract}, nil
}

// bindCount binds a generic wrapper to an already deployed contract.
func bindCount(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CountMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Count *CountRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Count.Contract.CountCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Count *CountRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Count.Contract.CountTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Count *CountRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Count.Contract.CountTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Count *CountCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Count.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Count *CountTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Count.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Count *CountTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Count.Contract.contract.Transact(opts, method, params...)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Count *CountCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Count.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Count *CountSession) Version() (string, error) {
	return _Count.Contract.Version(&_Count.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Count *CountCallerSession) Version() (string, error) {
	return _Count.Contract.Version(&_Count.CallOpts)
}

// CountNmb is a paid mutator transaction binding the contract method 0xe2119450.
//
// Solidity: function CountNmb() returns()
func (_Count *CountTransactor) CountNmb(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Count.contract.Transact(opts, "CountNmb")
}

// CountNmb is a paid mutator transaction binding the contract method 0xe2119450.
//
// Solidity: function CountNmb() returns()
func (_Count *CountSession) CountNmb() (*types.Transaction, error) {
	return _Count.Contract.CountNmb(&_Count.TransactOpts)
}

// CountNmb is a paid mutator transaction binding the contract method 0xe2119450.
//
// Solidity: function CountNmb() returns()
func (_Count *CountTransactorSession) CountNmb() (*types.Transaction, error) {
	return _Count.Contract.CountNmb(&_Count.TransactOpts)
}

// CountCountNumberIterator is returned from FilterCountNumber and is used to iterate over the raw logs and unpacked data for CountNumber events raised by the Count contract.
type CountCountNumberIterator struct {
	Event *CountCountNumber // Event containing the contract specifics and raw log

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
func (it *CountCountNumberIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CountCountNumber)
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
		it.Event = new(CountCountNumber)
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
func (it *CountCountNumberIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CountCountNumberIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CountCountNumber represents a CountNumber event raised by the Count contract.
type CountCountNumber struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterCountNumber is a free log retrieval operation binding the contract event 0xeca011e810d3353e4c1e9308544e45836b238ae8a7330a7e53ca8b4e03810df9.
//
// Solidity: event CountNumber(uint256 id)
func (_Count *CountFilterer) FilterCountNumber(opts *bind.FilterOpts) (*CountCountNumberIterator, error) {

	logs, sub, err := _Count.contract.FilterLogs(opts, "CountNumber")
	if err != nil {
		return nil, err
	}
	return &CountCountNumberIterator{contract: _Count.contract, event: "CountNumber", logs: logs, sub: sub}, nil
}

// WatchCountNumber is a free log subscription operation binding the contract event 0xeca011e810d3353e4c1e9308544e45836b238ae8a7330a7e53ca8b4e03810df9.
//
// Solidity: event CountNumber(uint256 id)
func (_Count *CountFilterer) WatchCountNumber(opts *bind.WatchOpts, sink chan<- *CountCountNumber) (event.Subscription, error) {

	logs, sub, err := _Count.contract.WatchLogs(opts, "CountNumber")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CountCountNumber)
				if err := _Count.contract.UnpackLog(event, "CountNumber", log); err != nil {
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

// ParseCountNumber is a log parse operation binding the contract event 0xeca011e810d3353e4c1e9308544e45836b238ae8a7330a7e53ca8b4e03810df9.
//
// Solidity: event CountNumber(uint256 id)
func (_Count *CountFilterer) ParseCountNumber(log types.Log) (*CountCountNumber, error) {
	event := new(CountCountNumber)
	if err := _Count.contract.UnpackLog(event, "CountNumber", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
