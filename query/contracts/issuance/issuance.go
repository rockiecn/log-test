// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package issuance

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ControlPermissionABI is the input ABI used to generate the binding from.
const ControlPermissionABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"role\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"SetRole\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"auth\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_role\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_role\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"bytes[5]\",\"name\":\"signs\",\"type\":\"bytes[5]\"}],\"name\":\"setRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ControlPermissionFuncSigs maps the 4-byte function signature to its string representation.
var ControlPermissionFuncSigs = map[string]string{
	"de9375f2": "auth()",
	"9e97b8f6": "hasRole(uint8,address)",
	"cfbfe5f3": "setRole(uint8,address,bytes[5])",
}

// ControlPermissionBin is the compiled bytecode used for deploying new contracts.
var ControlPermissionBin = "0x608060405234801561001057600080fd5b5060405161062438038061062483398101604081905261002f91610054565b600180546001600160a01b0319166001600160a01b0392909216919091179055610084565b60006020828403121561006657600080fd5b81516001600160a01b038116811461007d57600080fd5b9392505050565b610591806100936000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80639e97b8f614610046578063cfbfe5f31461006e578063de9375f214610083575b600080fd5b6100596100543660046102f5565b6100ae565b60405190151581526020015b60405180910390f35b61008161007c366004610398565b610152565b005b600154610096906001600160a01b031681565b6040516001600160a01b039091168152602001610065565b60ff82166000908152602081905260408120546001600160a01b038084169116036100db5750600161014c565b6000805b60088160ff16101561014557806002188516915060008260ff16118015610123575060ff82166000908152602081905260409020546001600160a01b038581169116145b156101335760019250505061014c565b8061013d816104a7565b9150506100df565b5060009150505b92915050565b813b600081900361019e5760405162461bcd60e51b81526020600482015260126024820152713732b2b21031b7b73a3930b1ba1030b2323960711b604482015260640160405180910390fd5b6040516bffffffffffffffffffffffff1930606090811b8216602084015266736574526f6c6560c81b603484015285901b16603b8201526001600160f81b031960f886901b16604f82015260009060500160408051601f1981840301815290829052805160209091012060015463a96bba9d60e01b83529092506001600160a01b03169063a96bba9d9061023890849087906004016104d4565b600060405180830381600087803b15801561025257600080fd5b505af1158015610266573d6000803e3d6000fd5b5050505060ff851660008181526020819052604080822080546001600160a01b0319166001600160a01b03891690811790915590519092917f5c22095cd1123887402a709a9698bc2961c12cb22eeb16e5d76480931f63ede091a35050505050565b803560ff811681146102d957600080fd5b919050565b80356001600160a01b03811681146102d957600080fd5b6000806040838503121561030857600080fd5b610311836102c8565b915061031f602084016102de565b90509250929050565b634e487b7160e01b600052604160045260246000fd5b60405160a0810167ffffffffffffffff8111828210171561036157610361610328565b60405290565b604051601f8201601f1916810167ffffffffffffffff8111828210171561039057610390610328565b604052919050565b6000806000606084860312156103ad57600080fd5b6103b6846102c8565b925060206103c58186016102de565b9250604085013567ffffffffffffffff808211156103e257600080fd5b8187019150601f88818401126103f757600080fd5b6103ff61033e565b8060a085018b81111561041157600080fd5b855b818110156104955780358681111561042b5760008081fd5b87018581018e1361043c5760008081fd5b80358781111561044e5761044e610328565b61045f818801601f19168b01610367565b8181528f8b8385010111156104745760008081fd5b818b84018c83013760009181018b0191909152855250928701928701610413565b50508096505050505050509250925092565b600060ff821660ff81036104cb57634e487b7160e01b600052601160045260246000fd5b60010192915050565b8281526040602080830182905260009160e08401919084018584805b600581101561054d57878603603f1901845282518051808852835b8181101561052657828101880151898201890152870161050b565b508781018701849052601f01601f19169096018501955092840192918401916001016104f0565b50939897505050505050505056fea2646970667358221220e03d75989d35e65fb3d12332962ed33b8ad19816f663923f7dd2412ba21f486264736f6c63430008100033"

// DeployControlPermission deploys a new Ethereum contract, binding an instance of ControlPermission to it.
func DeployControlPermission(auth *bind.TransactOpts, backend bind.ContractBackend, _a common.Address) (common.Address, *types.Transaction, *ControlPermission, error) {
	parsed, err := abi.JSON(strings.NewReader(ControlPermissionABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ControlPermissionBin), backend, _a)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ControlPermission{ControlPermissionCaller: ControlPermissionCaller{contract: contract}, ControlPermissionTransactor: ControlPermissionTransactor{contract: contract}, ControlPermissionFilterer: ControlPermissionFilterer{contract: contract}}, nil
}

// ControlPermission is an auto generated Go binding around an Ethereum contract.
type ControlPermission struct {
	ControlPermissionCaller     // Read-only binding to the contract
	ControlPermissionTransactor // Write-only binding to the contract
	ControlPermissionFilterer   // Log filterer for contract events
}

// ControlPermissionCaller is an auto generated read-only Go binding around an Ethereum contract.
type ControlPermissionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ControlPermissionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ControlPermissionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ControlPermissionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ControlPermissionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ControlPermissionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ControlPermissionSession struct {
	Contract     *ControlPermission // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ControlPermissionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ControlPermissionCallerSession struct {
	Contract *ControlPermissionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ControlPermissionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ControlPermissionTransactorSession struct {
	Contract     *ControlPermissionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ControlPermissionRaw is an auto generated low-level Go binding around an Ethereum contract.
type ControlPermissionRaw struct {
	Contract *ControlPermission // Generic contract binding to access the raw methods on
}

// ControlPermissionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ControlPermissionCallerRaw struct {
	Contract *ControlPermissionCaller // Generic read-only contract binding to access the raw methods on
}

// ControlPermissionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ControlPermissionTransactorRaw struct {
	Contract *ControlPermissionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewControlPermission creates a new instance of ControlPermission, bound to a specific deployed contract.
func NewControlPermission(address common.Address, backend bind.ContractBackend) (*ControlPermission, error) {
	contract, err := bindControlPermission(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ControlPermission{ControlPermissionCaller: ControlPermissionCaller{contract: contract}, ControlPermissionTransactor: ControlPermissionTransactor{contract: contract}, ControlPermissionFilterer: ControlPermissionFilterer{contract: contract}}, nil
}

// NewControlPermissionCaller creates a new read-only instance of ControlPermission, bound to a specific deployed contract.
func NewControlPermissionCaller(address common.Address, caller bind.ContractCaller) (*ControlPermissionCaller, error) {
	contract, err := bindControlPermission(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ControlPermissionCaller{contract: contract}, nil
}

// NewControlPermissionTransactor creates a new write-only instance of ControlPermission, bound to a specific deployed contract.
func NewControlPermissionTransactor(address common.Address, transactor bind.ContractTransactor) (*ControlPermissionTransactor, error) {
	contract, err := bindControlPermission(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ControlPermissionTransactor{contract: contract}, nil
}

// NewControlPermissionFilterer creates a new log filterer instance of ControlPermission, bound to a specific deployed contract.
func NewControlPermissionFilterer(address common.Address, filterer bind.ContractFilterer) (*ControlPermissionFilterer, error) {
	contract, err := bindControlPermission(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ControlPermissionFilterer{contract: contract}, nil
}

// bindControlPermission binds a generic wrapper to an already deployed contract.
func bindControlPermission(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ControlPermissionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ControlPermission *ControlPermissionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ControlPermission.Contract.ControlPermissionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ControlPermission *ControlPermissionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ControlPermission.Contract.ControlPermissionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ControlPermission *ControlPermissionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ControlPermission.Contract.ControlPermissionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ControlPermission *ControlPermissionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ControlPermission.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ControlPermission *ControlPermissionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ControlPermission.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ControlPermission *ControlPermissionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ControlPermission.Contract.contract.Transact(opts, method, params...)
}

// Auth is a free data retrieval call binding the contract method 0xde9375f2.
//
// Solidity: function auth() view returns(address)
func (_ControlPermission *ControlPermissionCaller) Auth(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ControlPermission.contract.Call(opts, &out, "auth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Auth is a free data retrieval call binding the contract method 0xde9375f2.
//
// Solidity: function auth() view returns(address)
func (_ControlPermission *ControlPermissionSession) Auth() (common.Address, error) {
	return _ControlPermission.Contract.Auth(&_ControlPermission.CallOpts)
}

// Auth is a free data retrieval call binding the contract method 0xde9375f2.
//
// Solidity: function auth() view returns(address)
func (_ControlPermission *ControlPermissionCallerSession) Auth() (common.Address, error) {
	return _ControlPermission.Contract.Auth(&_ControlPermission.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x9e97b8f6.
//
// Solidity: function hasRole(uint8 _role, address _a) view returns(bool)
func (_ControlPermission *ControlPermissionCaller) HasRole(opts *bind.CallOpts, _role uint8, _a common.Address) (bool, error) {
	var out []interface{}
	err := _ControlPermission.contract.Call(opts, &out, "hasRole", _role, _a)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x9e97b8f6.
//
// Solidity: function hasRole(uint8 _role, address _a) view returns(bool)
func (_ControlPermission *ControlPermissionSession) HasRole(_role uint8, _a common.Address) (bool, error) {
	return _ControlPermission.Contract.HasRole(&_ControlPermission.CallOpts, _role, _a)
}

// HasRole is a free data retrieval call binding the contract method 0x9e97b8f6.
//
// Solidity: function hasRole(uint8 _role, address _a) view returns(bool)
func (_ControlPermission *ControlPermissionCallerSession) HasRole(_role uint8, _a common.Address) (bool, error) {
	return _ControlPermission.Contract.HasRole(&_ControlPermission.CallOpts, _role, _a)
}

// SetRole is a paid mutator transaction binding the contract method 0xcfbfe5f3.
//
// Solidity: function setRole(uint8 _role, address _a, bytes[5] signs) returns()
func (_ControlPermission *ControlPermissionTransactor) SetRole(opts *bind.TransactOpts, _role uint8, _a common.Address, signs [5][]byte) (*types.Transaction, error) {
	return _ControlPermission.contract.Transact(opts, "setRole", _role, _a, signs)
}

// SetRole is a paid mutator transaction binding the contract method 0xcfbfe5f3.
//
// Solidity: function setRole(uint8 _role, address _a, bytes[5] signs) returns()
func (_ControlPermission *ControlPermissionSession) SetRole(_role uint8, _a common.Address, signs [5][]byte) (*types.Transaction, error) {
	return _ControlPermission.Contract.SetRole(&_ControlPermission.TransactOpts, _role, _a, signs)
}

// SetRole is a paid mutator transaction binding the contract method 0xcfbfe5f3.
//
// Solidity: function setRole(uint8 _role, address _a, bytes[5] signs) returns()
func (_ControlPermission *ControlPermissionTransactorSession) SetRole(_role uint8, _a common.Address, signs [5][]byte) (*types.Transaction, error) {
	return _ControlPermission.Contract.SetRole(&_ControlPermission.TransactOpts, _role, _a, signs)
}

// ControlPermissionSetRoleIterator is returned from FilterSetRole and is used to iterate over the raw logs and unpacked data for SetRole events raised by the ControlPermission contract.
type ControlPermissionSetRoleIterator struct {
	Event *ControlPermissionSetRole // Event containing the contract specifics and raw log

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
func (it *ControlPermissionSetRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControlPermissionSetRole)
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
		it.Event = new(ControlPermissionSetRole)
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
func (it *ControlPermissionSetRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControlPermissionSetRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControlPermissionSetRole represents a SetRole event raised by the ControlPermission contract.
type ControlPermissionSetRole struct {
	Role uint8
	A    common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSetRole is a free log retrieval operation binding the contract event 0x5c22095cd1123887402a709a9698bc2961c12cb22eeb16e5d76480931f63ede0.
//
// Solidity: event SetRole(uint8 indexed role, address indexed a)
func (_ControlPermission *ControlPermissionFilterer) FilterSetRole(opts *bind.FilterOpts, role []uint8, a []common.Address) (*ControlPermissionSetRoleIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var aRule []interface{}
	for _, aItem := range a {
		aRule = append(aRule, aItem)
	}

	logs, sub, err := _ControlPermission.contract.FilterLogs(opts, "SetRole", roleRule, aRule)
	if err != nil {
		return nil, err
	}
	return &ControlPermissionSetRoleIterator{contract: _ControlPermission.contract, event: "SetRole", logs: logs, sub: sub}, nil
}

// WatchSetRole is a free log subscription operation binding the contract event 0x5c22095cd1123887402a709a9698bc2961c12cb22eeb16e5d76480931f63ede0.
//
// Solidity: event SetRole(uint8 indexed role, address indexed a)
func (_ControlPermission *ControlPermissionFilterer) WatchSetRole(opts *bind.WatchOpts, sink chan<- *ControlPermissionSetRole, role []uint8, a []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var aRule []interface{}
	for _, aItem := range a {
		aRule = append(aRule, aItem)
	}

	logs, sub, err := _ControlPermission.contract.WatchLogs(opts, "SetRole", roleRule, aRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControlPermissionSetRole)
				if err := _ControlPermission.contract.UnpackLog(event, "SetRole", log); err != nil {
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

// ParseSetRole is a log parse operation binding the contract event 0x5c22095cd1123887402a709a9698bc2961c12cb22eeb16e5d76480931f63ede0.
//
// Solidity: event SetRole(uint8 indexed role, address indexed a)
func (_ControlPermission *ControlPermissionFilterer) ParseSetRole(log types.Log) (*ControlPermissionSetRole, error) {
	event := new(ControlPermissionSetRole)
	if err := _ControlPermission.contract.UnpackLog(event, "SetRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAuthABI is the input ABI used to generate the binding from.
const IAuthABI = "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_h\",\"type\":\"bytes32\"},{\"internalType\":\"bytes[5]\",\"name\":\"signs\",\"type\":\"bytes[5]\"}],\"name\":\"perm\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IAuthFuncSigs maps the 4-byte function signature to its string representation.
var IAuthFuncSigs = map[string]string{
	"a96bba9d": "perm(bytes32,bytes[5])",
}

// IAuth is an auto generated Go binding around an Ethereum contract.
type IAuth struct {
	IAuthCaller     // Read-only binding to the contract
	IAuthTransactor // Write-only binding to the contract
	IAuthFilterer   // Log filterer for contract events
}

// IAuthCaller is an auto generated read-only Go binding around an Ethereum contract.
type IAuthCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAuthTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IAuthTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAuthFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IAuthFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAuthSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IAuthSession struct {
	Contract     *IAuth            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IAuthCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IAuthCallerSession struct {
	Contract *IAuthCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IAuthTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IAuthTransactorSession struct {
	Contract     *IAuthTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IAuthRaw is an auto generated low-level Go binding around an Ethereum contract.
type IAuthRaw struct {
	Contract *IAuth // Generic contract binding to access the raw methods on
}

// IAuthCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IAuthCallerRaw struct {
	Contract *IAuthCaller // Generic read-only contract binding to access the raw methods on
}

// IAuthTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IAuthTransactorRaw struct {
	Contract *IAuthTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIAuth creates a new instance of IAuth, bound to a specific deployed contract.
func NewIAuth(address common.Address, backend bind.ContractBackend) (*IAuth, error) {
	contract, err := bindIAuth(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IAuth{IAuthCaller: IAuthCaller{contract: contract}, IAuthTransactor: IAuthTransactor{contract: contract}, IAuthFilterer: IAuthFilterer{contract: contract}}, nil
}

// NewIAuthCaller creates a new read-only instance of IAuth, bound to a specific deployed contract.
func NewIAuthCaller(address common.Address, caller bind.ContractCaller) (*IAuthCaller, error) {
	contract, err := bindIAuth(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IAuthCaller{contract: contract}, nil
}

// NewIAuthTransactor creates a new write-only instance of IAuth, bound to a specific deployed contract.
func NewIAuthTransactor(address common.Address, transactor bind.ContractTransactor) (*IAuthTransactor, error) {
	contract, err := bindIAuth(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IAuthTransactor{contract: contract}, nil
}

// NewIAuthFilterer creates a new log filterer instance of IAuth, bound to a specific deployed contract.
func NewIAuthFilterer(address common.Address, filterer bind.ContractFilterer) (*IAuthFilterer, error) {
	contract, err := bindIAuth(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IAuthFilterer{contract: contract}, nil
}

// bindIAuth binds a generic wrapper to an already deployed contract.
func bindIAuth(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IAuthABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAuth *IAuthRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAuth.Contract.IAuthCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAuth *IAuthRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAuth.Contract.IAuthTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAuth *IAuthRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAuth.Contract.IAuthTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAuth *IAuthCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAuth.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAuth *IAuthTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAuth.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAuth *IAuthTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAuth.Contract.contract.Transact(opts, method, params...)
}

// Perm is a paid mutator transaction binding the contract method 0xa96bba9d.
//
// Solidity: function perm(bytes32 _h, bytes[5] signs) returns()
func (_IAuth *IAuthTransactor) Perm(opts *bind.TransactOpts, _h [32]byte, signs [5][]byte) (*types.Transaction, error) {
	return _IAuth.contract.Transact(opts, "perm", _h, signs)
}

// Perm is a paid mutator transaction binding the contract method 0xa96bba9d.
//
// Solidity: function perm(bytes32 _h, bytes[5] signs) returns()
func (_IAuth *IAuthSession) Perm(_h [32]byte, signs [5][]byte) (*types.Transaction, error) {
	return _IAuth.Contract.Perm(&_IAuth.TransactOpts, _h, signs)
}

// Perm is a paid mutator transaction binding the contract method 0xa96bba9d.
//
// Solidity: function perm(bytes32 _h, bytes[5] signs) returns()
func (_IAuth *IAuthTransactorSession) Perm(_h [32]byte, signs [5][]byte) (*types.Transaction, error) {
	return _IAuth.Contract.Perm(&_IAuth.TransactOpts, _h, signs)
}

// IControlPermissionABI is the input ABI used to generate the binding from.
const IControlPermissionABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"role\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"SetRole\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_role\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_role\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"bytes[5]\",\"name\":\"signs\",\"type\":\"bytes[5]\"}],\"name\":\"setRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IControlPermissionFuncSigs maps the 4-byte function signature to its string representation.
var IControlPermissionFuncSigs = map[string]string{
	"9e97b8f6": "hasRole(uint8,address)",
	"cfbfe5f3": "setRole(uint8,address,bytes[5])",
}

// IControlPermission is an auto generated Go binding around an Ethereum contract.
type IControlPermission struct {
	IControlPermissionCaller     // Read-only binding to the contract
	IControlPermissionTransactor // Write-only binding to the contract
	IControlPermissionFilterer   // Log filterer for contract events
}

// IControlPermissionCaller is an auto generated read-only Go binding around an Ethereum contract.
type IControlPermissionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IControlPermissionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IControlPermissionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IControlPermissionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IControlPermissionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IControlPermissionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IControlPermissionSession struct {
	Contract     *IControlPermission // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IControlPermissionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IControlPermissionCallerSession struct {
	Contract *IControlPermissionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// IControlPermissionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IControlPermissionTransactorSession struct {
	Contract     *IControlPermissionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// IControlPermissionRaw is an auto generated low-level Go binding around an Ethereum contract.
type IControlPermissionRaw struct {
	Contract *IControlPermission // Generic contract binding to access the raw methods on
}

// IControlPermissionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IControlPermissionCallerRaw struct {
	Contract *IControlPermissionCaller // Generic read-only contract binding to access the raw methods on
}

// IControlPermissionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IControlPermissionTransactorRaw struct {
	Contract *IControlPermissionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIControlPermission creates a new instance of IControlPermission, bound to a specific deployed contract.
func NewIControlPermission(address common.Address, backend bind.ContractBackend) (*IControlPermission, error) {
	contract, err := bindIControlPermission(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IControlPermission{IControlPermissionCaller: IControlPermissionCaller{contract: contract}, IControlPermissionTransactor: IControlPermissionTransactor{contract: contract}, IControlPermissionFilterer: IControlPermissionFilterer{contract: contract}}, nil
}

// NewIControlPermissionCaller creates a new read-only instance of IControlPermission, bound to a specific deployed contract.
func NewIControlPermissionCaller(address common.Address, caller bind.ContractCaller) (*IControlPermissionCaller, error) {
	contract, err := bindIControlPermission(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IControlPermissionCaller{contract: contract}, nil
}

// NewIControlPermissionTransactor creates a new write-only instance of IControlPermission, bound to a specific deployed contract.
func NewIControlPermissionTransactor(address common.Address, transactor bind.ContractTransactor) (*IControlPermissionTransactor, error) {
	contract, err := bindIControlPermission(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IControlPermissionTransactor{contract: contract}, nil
}

// NewIControlPermissionFilterer creates a new log filterer instance of IControlPermission, bound to a specific deployed contract.
func NewIControlPermissionFilterer(address common.Address, filterer bind.ContractFilterer) (*IControlPermissionFilterer, error) {
	contract, err := bindIControlPermission(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IControlPermissionFilterer{contract: contract}, nil
}

// bindIControlPermission binds a generic wrapper to an already deployed contract.
func bindIControlPermission(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IControlPermissionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IControlPermission *IControlPermissionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IControlPermission.Contract.IControlPermissionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IControlPermission *IControlPermissionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IControlPermission.Contract.IControlPermissionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IControlPermission *IControlPermissionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IControlPermission.Contract.IControlPermissionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IControlPermission *IControlPermissionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IControlPermission.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IControlPermission *IControlPermissionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IControlPermission.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IControlPermission *IControlPermissionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IControlPermission.Contract.contract.Transact(opts, method, params...)
}

// HasRole is a free data retrieval call binding the contract method 0x9e97b8f6.
//
// Solidity: function hasRole(uint8 _role, address _a) view returns(bool)
func (_IControlPermission *IControlPermissionCaller) HasRole(opts *bind.CallOpts, _role uint8, _a common.Address) (bool, error) {
	var out []interface{}
	err := _IControlPermission.contract.Call(opts, &out, "hasRole", _role, _a)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x9e97b8f6.
//
// Solidity: function hasRole(uint8 _role, address _a) view returns(bool)
func (_IControlPermission *IControlPermissionSession) HasRole(_role uint8, _a common.Address) (bool, error) {
	return _IControlPermission.Contract.HasRole(&_IControlPermission.CallOpts, _role, _a)
}

// HasRole is a free data retrieval call binding the contract method 0x9e97b8f6.
//
// Solidity: function hasRole(uint8 _role, address _a) view returns(bool)
func (_IControlPermission *IControlPermissionCallerSession) HasRole(_role uint8, _a common.Address) (bool, error) {
	return _IControlPermission.Contract.HasRole(&_IControlPermission.CallOpts, _role, _a)
}

// SetRole is a paid mutator transaction binding the contract method 0xcfbfe5f3.
//
// Solidity: function setRole(uint8 _role, address _a, bytes[5] signs) returns()
func (_IControlPermission *IControlPermissionTransactor) SetRole(opts *bind.TransactOpts, _role uint8, _a common.Address, signs [5][]byte) (*types.Transaction, error) {
	return _IControlPermission.contract.Transact(opts, "setRole", _role, _a, signs)
}

// SetRole is a paid mutator transaction binding the contract method 0xcfbfe5f3.
//
// Solidity: function setRole(uint8 _role, address _a, bytes[5] signs) returns()
func (_IControlPermission *IControlPermissionSession) SetRole(_role uint8, _a common.Address, signs [5][]byte) (*types.Transaction, error) {
	return _IControlPermission.Contract.SetRole(&_IControlPermission.TransactOpts, _role, _a, signs)
}

// SetRole is a paid mutator transaction binding the contract method 0xcfbfe5f3.
//
// Solidity: function setRole(uint8 _role, address _a, bytes[5] signs) returns()
func (_IControlPermission *IControlPermissionTransactorSession) SetRole(_role uint8, _a common.Address, signs [5][]byte) (*types.Transaction, error) {
	return _IControlPermission.Contract.SetRole(&_IControlPermission.TransactOpts, _role, _a, signs)
}

// IControlPermissionSetRoleIterator is returned from FilterSetRole and is used to iterate over the raw logs and unpacked data for SetRole events raised by the IControlPermission contract.
type IControlPermissionSetRoleIterator struct {
	Event *IControlPermissionSetRole // Event containing the contract specifics and raw log

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
func (it *IControlPermissionSetRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IControlPermissionSetRole)
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
		it.Event = new(IControlPermissionSetRole)
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
func (it *IControlPermissionSetRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IControlPermissionSetRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IControlPermissionSetRole represents a SetRole event raised by the IControlPermission contract.
type IControlPermissionSetRole struct {
	Role uint8
	A    common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSetRole is a free log retrieval operation binding the contract event 0x5c22095cd1123887402a709a9698bc2961c12cb22eeb16e5d76480931f63ede0.
//
// Solidity: event SetRole(uint8 indexed role, address indexed a)
func (_IControlPermission *IControlPermissionFilterer) FilterSetRole(opts *bind.FilterOpts, role []uint8, a []common.Address) (*IControlPermissionSetRoleIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var aRule []interface{}
	for _, aItem := range a {
		aRule = append(aRule, aItem)
	}

	logs, sub, err := _IControlPermission.contract.FilterLogs(opts, "SetRole", roleRule, aRule)
	if err != nil {
		return nil, err
	}
	return &IControlPermissionSetRoleIterator{contract: _IControlPermission.contract, event: "SetRole", logs: logs, sub: sub}, nil
}

// WatchSetRole is a free log subscription operation binding the contract event 0x5c22095cd1123887402a709a9698bc2961c12cb22eeb16e5d76480931f63ede0.
//
// Solidity: event SetRole(uint8 indexed role, address indexed a)
func (_IControlPermission *IControlPermissionFilterer) WatchSetRole(opts *bind.WatchOpts, sink chan<- *IControlPermissionSetRole, role []uint8, a []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var aRule []interface{}
	for _, aItem := range a {
		aRule = append(aRule, aItem)
	}

	logs, sub, err := _IControlPermission.contract.WatchLogs(opts, "SetRole", roleRule, aRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IControlPermissionSetRole)
				if err := _IControlPermission.contract.UnpackLog(event, "SetRole", log); err != nil {
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

// ParseSetRole is a log parse operation binding the contract event 0x5c22095cd1123887402a709a9698bc2961c12cb22eeb16e5d76480931f63ede0.
//
// Solidity: event SetRole(uint8 indexed role, address indexed a)
func (_IControlPermission *IControlPermissionFilterer) ParseSetRole(log types.Log) (*IControlPermissionSetRole, error) {
	event := new(IControlPermissionSetRole)
	if err := _IControlPermission.contract.UnpackLog(event, "SetRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IIssuanceABI is the input ABI used to generate the binding from.
const IIssuanceABI = "[{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_start\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_end\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_size\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"_sPrice\",\"type\":\"uint256\"}],\"name\":\"issu\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IIssuanceFuncSigs maps the 4-byte function signature to its string representation.
var IIssuanceFuncSigs = map[string]string{
	"13df3828": "issu(uint64,uint64,uint64,uint256)",
}

// IIssuance is an auto generated Go binding around an Ethereum contract.
type IIssuance struct {
	IIssuanceCaller     // Read-only binding to the contract
	IIssuanceTransactor // Write-only binding to the contract
	IIssuanceFilterer   // Log filterer for contract events
}

// IIssuanceCaller is an auto generated read-only Go binding around an Ethereum contract.
type IIssuanceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IIssuanceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IIssuanceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IIssuanceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IIssuanceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IIssuanceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IIssuanceSession struct {
	Contract     *IIssuance        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IIssuanceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IIssuanceCallerSession struct {
	Contract *IIssuanceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// IIssuanceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IIssuanceTransactorSession struct {
	Contract     *IIssuanceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IIssuanceRaw is an auto generated low-level Go binding around an Ethereum contract.
type IIssuanceRaw struct {
	Contract *IIssuance // Generic contract binding to access the raw methods on
}

// IIssuanceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IIssuanceCallerRaw struct {
	Contract *IIssuanceCaller // Generic read-only contract binding to access the raw methods on
}

// IIssuanceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IIssuanceTransactorRaw struct {
	Contract *IIssuanceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIIssuance creates a new instance of IIssuance, bound to a specific deployed contract.
func NewIIssuance(address common.Address, backend bind.ContractBackend) (*IIssuance, error) {
	contract, err := bindIIssuance(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IIssuance{IIssuanceCaller: IIssuanceCaller{contract: contract}, IIssuanceTransactor: IIssuanceTransactor{contract: contract}, IIssuanceFilterer: IIssuanceFilterer{contract: contract}}, nil
}

// NewIIssuanceCaller creates a new read-only instance of IIssuance, bound to a specific deployed contract.
func NewIIssuanceCaller(address common.Address, caller bind.ContractCaller) (*IIssuanceCaller, error) {
	contract, err := bindIIssuance(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IIssuanceCaller{contract: contract}, nil
}

// NewIIssuanceTransactor creates a new write-only instance of IIssuance, bound to a specific deployed contract.
func NewIIssuanceTransactor(address common.Address, transactor bind.ContractTransactor) (*IIssuanceTransactor, error) {
	contract, err := bindIIssuance(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IIssuanceTransactor{contract: contract}, nil
}

// NewIIssuanceFilterer creates a new log filterer instance of IIssuance, bound to a specific deployed contract.
func NewIIssuanceFilterer(address common.Address, filterer bind.ContractFilterer) (*IIssuanceFilterer, error) {
	contract, err := bindIIssuance(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IIssuanceFilterer{contract: contract}, nil
}

// bindIIssuance binds a generic wrapper to an already deployed contract.
func bindIIssuance(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IIssuanceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IIssuance *IIssuanceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IIssuance.Contract.IIssuanceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IIssuance *IIssuanceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IIssuance.Contract.IIssuanceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IIssuance *IIssuanceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IIssuance.Contract.IIssuanceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IIssuance *IIssuanceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IIssuance.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IIssuance *IIssuanceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IIssuance.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IIssuance *IIssuanceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IIssuance.Contract.contract.Transact(opts, method, params...)
}

// Issu is a paid mutator transaction binding the contract method 0x13df3828.
//
// Solidity: function issu(uint64 _start, uint64 _end, uint64 _size, uint256 _sPrice) returns(uint256)
func (_IIssuance *IIssuanceTransactor) Issu(opts *bind.TransactOpts, _start uint64, _end uint64, _size uint64, _sPrice *big.Int) (*types.Transaction, error) {
	return _IIssuance.contract.Transact(opts, "issu", _start, _end, _size, _sPrice)
}

// Issu is a paid mutator transaction binding the contract method 0x13df3828.
//
// Solidity: function issu(uint64 _start, uint64 _end, uint64 _size, uint256 _sPrice) returns(uint256)
func (_IIssuance *IIssuanceSession) Issu(_start uint64, _end uint64, _size uint64, _sPrice *big.Int) (*types.Transaction, error) {
	return _IIssuance.Contract.Issu(&_IIssuance.TransactOpts, _start, _end, _size, _sPrice)
}

// Issu is a paid mutator transaction binding the contract method 0x13df3828.
//
// Solidity: function issu(uint64 _start, uint64 _end, uint64 _size, uint256 _sPrice) returns(uint256)
func (_IIssuance *IIssuanceTransactorSession) Issu(_start uint64, _end uint64, _size uint64, _sPrice *big.Int) (*types.Transaction, error) {
	return _IIssuance.Contract.Issu(&_IIssuance.TransactOpts, _start, _end, _size, _sPrice)
}

// IssuanceABI is the input ABI used to generate the binding from.
const IssuanceABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"role\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"SetRole\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"auth\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"expire\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_role\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_start\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_end\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_size\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"_sPrice\",\"type\":\"uint256\"}],\"name\":\"issu\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"last\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardResid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_role\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"bytes[5]\",\"name\":\"signs\",\"type\":\"bytes[5]\"}],\"name\":\"setRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"size\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"spaceTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sprice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stageRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stageReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stageRewardResid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stageSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"targetReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IssuanceFuncSigs maps the 4-byte function signature to its string representation.
var IssuanceFuncSigs = map[string]string{
	"de9375f2": "auth()",
	"22228560": "expire(uint64)",
	"9e97b8f6": "hasRole(uint8,address)",
	"13df3828": "issu(uint64,uint64,uint64,uint256)",
	"47799da8": "last()",
	"48ac60aa": "rewardResid()",
	"cfbfe5f3": "setRole(uint8,address,bytes[5])",
	"949d225d": "size()",
	"43a2755c": "spaceTime()",
	"8937e7f0": "sprice()",
	"e72b5250": "stageRatio()",
	"e12e4b21": "stageReward()",
	"bd6ba39a": "stageRewardResid()",
	"fb752b8e": "stageSize()",
	"eda7461a": "targetReward()",
	"54fd4d50": "version()",
}

// IssuanceBin is the compiled bytecode used for deploying new contracts.
var IssuanceBin = "0x60806040526001805461ffff60a01b1916600160a11b1790556af8277896582678ac0000006002819055600355612710600481905565010000000000600581905561004a91906100ca565b610056906103606100ca565b60065560065460075534801561006b57600080fd5b50604051610eae380380610eae83398101604081905261008a916100f7565b600180546001600160a01b0319166001600160a01b0392909216919091179055600880546001600160401b031916426001600160401b0316179055610127565b60008160001904831182151516156100f257634e487b7160e01b600052601160045260246000fd5b500290565b60006020828403121561010957600080fd5b81516001600160a01b038116811461012057600080fd5b9392505050565b610d78806101366000396000f3fe608060405234801561001057600080fd5b50600436106101005760003560e01c80639e97b8f611610097578063e12e4b2111610066578063e12e4b211461023f578063e72b525014610248578063eda7461a14610251578063fb752b8e1461025a57600080fd5b80639e97b8f6146101d3578063bd6ba39a146101f6578063cfbfe5f3146101ff578063de9375f21461021457600080fd5b806348ac60aa116100d357806348ac60aa1461017f57806354fd4d50146101885780638937e7f0146101b0578063949d225d146101b957600080fd5b806313df382814610105578063222285601461012b57806343a2755c1461014b57806347799da814610154575b600080fd5b61011861011336600461090e565b610263565b6040519081526020015b60405180910390f35b610118610139366004610959565b600b6020526000908152604090205481565b610118600a5481565b600854610167906001600160401b031681565b6040516001600160401b039091168152602001610122565b61011860035481565b60015461019d90600160a01b900461ffff1681565b60405161ffff9091168152602001610122565b61011860095481565b60085461016790600160401b90046001600160401b031681565b6101e66101e13660046109a3565b610637565b6040519015158152602001610122565b61011860075481565b61021261020d366004610a44565b6106db565b005b600154610227906001600160a01b031681565b6040516001600160a01b039091168152602001610122565b61011860065481565b61011860045481565b61011860025481565b61011860055481565b600060026102718133610852565b6001600160401b0385166000908152600b602052604081208054859290610299908490610b68565b909155505060085442906102b9906001600160401b031662015180610b7b565b6001600160401b0316816001600160401b031611156102ee576008546102eb906001600160401b031662015180610b7b565b90505b600854600090610307906001600160401b031683610ba2565b6001600160401b031660095461031d9190610bc2565b905061032c6201518083610bf7565b6008546001600160401b039182169161034a91620151809116610bf7565b6001600160401b031610156103d75760006103686201518084610bf7565b6103759062015180610c1d565b6001600160401b0381166000908152600b60205260408120546009805493945090928392906103a5908490610c4c565b909155506103b590508285610ba2565b6103c8906001600160401b031682610bc2565b6103d29084610c4c565b925050505b6008805467ffffffffffffffff19166001600160401b038416179055600454612710906104049083610bc2565b61040e9190610c5f565b905060075481111561041f57506007545b60035481111561042e57506003545b80600760008282546104409190610c4c565b9250508190555080600360008282546104599190610c4c565b9091555061046990508888610ba2565b6001600160401b0316866001600160401b03166104869190610bc2565b600a60008282546104979190610b68565b9250508190555084600960008282546104b09190610b68565b90915550506008805487919081906104d9908490600160401b90046001600160401b0316610b7b565b92506101000a8154816001600160401b0302191690836001600160401b031602179055506005546283d600600a546105119190610c5f565b1061062c57670de0b6b3a7640000600554106105455760026005600082825461053a9190610bc2565b9091555061055e9050565b6020600560008282546105589190610bc2565b90915550505b60085460009061058090600160401b90046001600160401b0316612710610c1d565b6001600160401b0316600a546105969190610c5f565b9050600a600454600f6105a99190610bc2565b6105b39190610c5f565b600481905560055482916105c691610bc2565b6105d09190610bc2565b60068190556003541161062457806005546105eb9190610bc2565b6105f6906002610bc2565b6003546106039190610c5f565b6004819055600554829161061691610bc2565b6106209190610bc2565b6006555b506006546007555b979650505050505050565b60ff82166000908152602081905260408120546001600160a01b03808416911603610664575060016106d5565b6000805b60088160ff1610156106ce57806002188516915060008260ff161180156106ac575060ff82166000908152602081905260409020546001600160a01b038581169116145b156106bc576001925050506106d5565b806106c681610c73565b915050610668565b5060009150505b92915050565b813b60008190036107285760405162461bcd60e51b81526020600482015260126024820152713732b2b21031b7b73a3930b1ba1030b2323960711b60448201526064015b60405180910390fd5b6040516bffffffffffffffffffffffff1930606090811b8216602084015266736574526f6c6560c81b603484015285901b16603b8201526001600160f81b031960f886901b16604f82015260009060500160408051601f1981840301815290829052805160209091012060015463a96bba9d60e01b83529092506001600160a01b03169063a96bba9d906107c29084908790600401610cd8565b600060405180830381600087803b1580156107dc57600080fd5b505af11580156107f0573d6000803e3d6000fd5b5050505060ff851660008181526020819052604080822080546001600160a01b0319166001600160a01b03891690811790915590519092917f5c22095cd1123887402a709a9698bc2961c12cb22eeb16e5d76480931f63ede091a35050505050565b61085c8282610637565b6108ee5760405171024b739ba30b731b29d1030b1b1b7bab73a160751b60208201526bffffffffffffffffffffffff19606083901b1660328201527001034b99036b4b9b9b4b733903937b6329607d1b60468201526001600160f81b031960f884901b16605782015260580160408051601f198184030181529082905262461bcd60e51b825261071f91600401610d2f565b5050565b80356001600160401b038116811461090957600080fd5b919050565b6000806000806080858703121561092457600080fd5b61092d856108f2565b935061093b602086016108f2565b9250610949604086016108f2565b9396929550929360600135925050565b60006020828403121561096b57600080fd5b610974826108f2565b9392505050565b803560ff8116811461090957600080fd5b80356001600160a01b038116811461090957600080fd5b600080604083850312156109b657600080fd5b6109bf8361097b565b91506109cd6020840161098c565b90509250929050565b634e487b7160e01b600052604160045260246000fd5b60405160a081016001600160401b0381118282101715610a0e57610a0e6109d6565b60405290565b604051601f8201601f191681016001600160401b0381118282101715610a3c57610a3c6109d6565b604052919050565b600080600060608486031215610a5957600080fd5b610a628461097b565b92506020610a7181860161098c565b925060408501356001600160401b0380821115610a8d57600080fd5b8187019150601f8881840112610aa257600080fd5b610aaa6109ec565b8060a085018b811115610abc57600080fd5b855b81811015610b4057803586811115610ad65760008081fd5b87018581018e13610ae75760008081fd5b803587811115610af957610af96109d6565b610b0a818801601f19168b01610a14565b8181528f8b838501011115610b1f5760008081fd5b818b84018c83013760009181018b0191909152855250928701928701610abe565b50508096505050505050509250925092565b634e487b7160e01b600052601160045260246000fd5b808201808211156106d5576106d5610b52565b6001600160401b03818116838216019080821115610b9b57610b9b610b52565b5092915050565b6001600160401b03828116828216039080821115610b9b57610b9b610b52565b6000816000190483118215151615610bdc57610bdc610b52565b500290565b634e487b7160e01b600052601260045260246000fd5b60006001600160401b0380841680610c1157610c11610be1565b92169190910492915050565b60006001600160401b0380831681851681830481118215151615610c4357610c43610b52565b02949350505050565b818103818111156106d5576106d5610b52565b600082610c6e57610c6e610be1565b500490565b600060ff821660ff8103610c8957610c89610b52565b60010192915050565b6000815180845260005b81811015610cb857602081850181015186830182015201610c9c565b506000602082860101526020601f19601f83011685010191505092915050565b8281526040602080830182905260009160e084019190840185845b6005811015610d2257603f19878603018352610d10858351610c92565b94509183019190830190600101610cf3565b5092979650505050505050565b6020815260006109746020830184610c9256fea2646970667358221220493043a0dfdc7716b7e8d584d525353d369a34c84d40c9579b1a9ed0eb63d29d64736f6c63430008100033"

// DeployIssuance deploys a new Ethereum contract, binding an instance of Issuance to it.
func DeployIssuance(auth *bind.TransactOpts, backend bind.ContractBackend, _a common.Address) (common.Address, *types.Transaction, *Issuance, error) {
	parsed, err := abi.JSON(strings.NewReader(IssuanceABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IssuanceBin), backend, _a)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Issuance{IssuanceCaller: IssuanceCaller{contract: contract}, IssuanceTransactor: IssuanceTransactor{contract: contract}, IssuanceFilterer: IssuanceFilterer{contract: contract}}, nil
}

// Issuance is an auto generated Go binding around an Ethereum contract.
type Issuance struct {
	IssuanceCaller     // Read-only binding to the contract
	IssuanceTransactor // Write-only binding to the contract
	IssuanceFilterer   // Log filterer for contract events
}

// IssuanceCaller is an auto generated read-only Go binding around an Ethereum contract.
type IssuanceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IssuanceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IssuanceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IssuanceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IssuanceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IssuanceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IssuanceSession struct {
	Contract     *Issuance         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IssuanceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IssuanceCallerSession struct {
	Contract *IssuanceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// IssuanceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IssuanceTransactorSession struct {
	Contract     *IssuanceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IssuanceRaw is an auto generated low-level Go binding around an Ethereum contract.
type IssuanceRaw struct {
	Contract *Issuance // Generic contract binding to access the raw methods on
}

// IssuanceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IssuanceCallerRaw struct {
	Contract *IssuanceCaller // Generic read-only contract binding to access the raw methods on
}

// IssuanceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IssuanceTransactorRaw struct {
	Contract *IssuanceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIssuance creates a new instance of Issuance, bound to a specific deployed contract.
func NewIssuance(address common.Address, backend bind.ContractBackend) (*Issuance, error) {
	contract, err := bindIssuance(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Issuance{IssuanceCaller: IssuanceCaller{contract: contract}, IssuanceTransactor: IssuanceTransactor{contract: contract}, IssuanceFilterer: IssuanceFilterer{contract: contract}}, nil
}

// NewIssuanceCaller creates a new read-only instance of Issuance, bound to a specific deployed contract.
func NewIssuanceCaller(address common.Address, caller bind.ContractCaller) (*IssuanceCaller, error) {
	contract, err := bindIssuance(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IssuanceCaller{contract: contract}, nil
}

// NewIssuanceTransactor creates a new write-only instance of Issuance, bound to a specific deployed contract.
func NewIssuanceTransactor(address common.Address, transactor bind.ContractTransactor) (*IssuanceTransactor, error) {
	contract, err := bindIssuance(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IssuanceTransactor{contract: contract}, nil
}

// NewIssuanceFilterer creates a new log filterer instance of Issuance, bound to a specific deployed contract.
func NewIssuanceFilterer(address common.Address, filterer bind.ContractFilterer) (*IssuanceFilterer, error) {
	contract, err := bindIssuance(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IssuanceFilterer{contract: contract}, nil
}

// bindIssuance binds a generic wrapper to an already deployed contract.
func bindIssuance(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IssuanceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Issuance *IssuanceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Issuance.Contract.IssuanceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Issuance *IssuanceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Issuance.Contract.IssuanceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Issuance *IssuanceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Issuance.Contract.IssuanceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Issuance *IssuanceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Issuance.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Issuance *IssuanceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Issuance.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Issuance *IssuanceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Issuance.Contract.contract.Transact(opts, method, params...)
}

// Auth is a free data retrieval call binding the contract method 0xde9375f2.
//
// Solidity: function auth() view returns(address)
func (_Issuance *IssuanceCaller) Auth(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Issuance.contract.Call(opts, &out, "auth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Auth is a free data retrieval call binding the contract method 0xde9375f2.
//
// Solidity: function auth() view returns(address)
func (_Issuance *IssuanceSession) Auth() (common.Address, error) {
	return _Issuance.Contract.Auth(&_Issuance.CallOpts)
}

// Auth is a free data retrieval call binding the contract method 0xde9375f2.
//
// Solidity: function auth() view returns(address)
func (_Issuance *IssuanceCallerSession) Auth() (common.Address, error) {
	return _Issuance.Contract.Auth(&_Issuance.CallOpts)
}

// Expire is a free data retrieval call binding the contract method 0x22228560.
//
// Solidity: function expire(uint64 ) view returns(uint256)
func (_Issuance *IssuanceCaller) Expire(opts *bind.CallOpts, arg0 uint64) (*big.Int, error) {
	var out []interface{}
	err := _Issuance.contract.Call(opts, &out, "expire", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Expire is a free data retrieval call binding the contract method 0x22228560.
//
// Solidity: function expire(uint64 ) view returns(uint256)
func (_Issuance *IssuanceSession) Expire(arg0 uint64) (*big.Int, error) {
	return _Issuance.Contract.Expire(&_Issuance.CallOpts, arg0)
}

// Expire is a free data retrieval call binding the contract method 0x22228560.
//
// Solidity: function expire(uint64 ) view returns(uint256)
func (_Issuance *IssuanceCallerSession) Expire(arg0 uint64) (*big.Int, error) {
	return _Issuance.Contract.Expire(&_Issuance.CallOpts, arg0)
}

// HasRole is a free data retrieval call binding the contract method 0x9e97b8f6.
//
// Solidity: function hasRole(uint8 _role, address _a) view returns(bool)
func (_Issuance *IssuanceCaller) HasRole(opts *bind.CallOpts, _role uint8, _a common.Address) (bool, error) {
	var out []interface{}
	err := _Issuance.contract.Call(opts, &out, "hasRole", _role, _a)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x9e97b8f6.
//
// Solidity: function hasRole(uint8 _role, address _a) view returns(bool)
func (_Issuance *IssuanceSession) HasRole(_role uint8, _a common.Address) (bool, error) {
	return _Issuance.Contract.HasRole(&_Issuance.CallOpts, _role, _a)
}

// HasRole is a free data retrieval call binding the contract method 0x9e97b8f6.
//
// Solidity: function hasRole(uint8 _role, address _a) view returns(bool)
func (_Issuance *IssuanceCallerSession) HasRole(_role uint8, _a common.Address) (bool, error) {
	return _Issuance.Contract.HasRole(&_Issuance.CallOpts, _role, _a)
}

// Last is a free data retrieval call binding the contract method 0x47799da8.
//
// Solidity: function last() view returns(uint64)
func (_Issuance *IssuanceCaller) Last(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Issuance.contract.Call(opts, &out, "last")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// Last is a free data retrieval call binding the contract method 0x47799da8.
//
// Solidity: function last() view returns(uint64)
func (_Issuance *IssuanceSession) Last() (uint64, error) {
	return _Issuance.Contract.Last(&_Issuance.CallOpts)
}

// Last is a free data retrieval call binding the contract method 0x47799da8.
//
// Solidity: function last() view returns(uint64)
func (_Issuance *IssuanceCallerSession) Last() (uint64, error) {
	return _Issuance.Contract.Last(&_Issuance.CallOpts)
}

// RewardResid is a free data retrieval call binding the contract method 0x48ac60aa.
//
// Solidity: function rewardResid() view returns(uint256)
func (_Issuance *IssuanceCaller) RewardResid(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Issuance.contract.Call(opts, &out, "rewardResid")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardResid is a free data retrieval call binding the contract method 0x48ac60aa.
//
// Solidity: function rewardResid() view returns(uint256)
func (_Issuance *IssuanceSession) RewardResid() (*big.Int, error) {
	return _Issuance.Contract.RewardResid(&_Issuance.CallOpts)
}

// RewardResid is a free data retrieval call binding the contract method 0x48ac60aa.
//
// Solidity: function rewardResid() view returns(uint256)
func (_Issuance *IssuanceCallerSession) RewardResid() (*big.Int, error) {
	return _Issuance.Contract.RewardResid(&_Issuance.CallOpts)
}

// Size is a free data retrieval call binding the contract method 0x949d225d.
//
// Solidity: function size() view returns(uint64)
func (_Issuance *IssuanceCaller) Size(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Issuance.contract.Call(opts, &out, "size")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// Size is a free data retrieval call binding the contract method 0x949d225d.
//
// Solidity: function size() view returns(uint64)
func (_Issuance *IssuanceSession) Size() (uint64, error) {
	return _Issuance.Contract.Size(&_Issuance.CallOpts)
}

// Size is a free data retrieval call binding the contract method 0x949d225d.
//
// Solidity: function size() view returns(uint64)
func (_Issuance *IssuanceCallerSession) Size() (uint64, error) {
	return _Issuance.Contract.Size(&_Issuance.CallOpts)
}

// SpaceTime is a free data retrieval call binding the contract method 0x43a2755c.
//
// Solidity: function spaceTime() view returns(uint256)
func (_Issuance *IssuanceCaller) SpaceTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Issuance.contract.Call(opts, &out, "spaceTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SpaceTime is a free data retrieval call binding the contract method 0x43a2755c.
//
// Solidity: function spaceTime() view returns(uint256)
func (_Issuance *IssuanceSession) SpaceTime() (*big.Int, error) {
	return _Issuance.Contract.SpaceTime(&_Issuance.CallOpts)
}

// SpaceTime is a free data retrieval call binding the contract method 0x43a2755c.
//
// Solidity: function spaceTime() view returns(uint256)
func (_Issuance *IssuanceCallerSession) SpaceTime() (*big.Int, error) {
	return _Issuance.Contract.SpaceTime(&_Issuance.CallOpts)
}

// Sprice is a free data retrieval call binding the contract method 0x8937e7f0.
//
// Solidity: function sprice() view returns(uint256)
func (_Issuance *IssuanceCaller) Sprice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Issuance.contract.Call(opts, &out, "sprice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Sprice is a free data retrieval call binding the contract method 0x8937e7f0.
//
// Solidity: function sprice() view returns(uint256)
func (_Issuance *IssuanceSession) Sprice() (*big.Int, error) {
	return _Issuance.Contract.Sprice(&_Issuance.CallOpts)
}

// Sprice is a free data retrieval call binding the contract method 0x8937e7f0.
//
// Solidity: function sprice() view returns(uint256)
func (_Issuance *IssuanceCallerSession) Sprice() (*big.Int, error) {
	return _Issuance.Contract.Sprice(&_Issuance.CallOpts)
}

// StageRatio is a free data retrieval call binding the contract method 0xe72b5250.
//
// Solidity: function stageRatio() view returns(uint256)
func (_Issuance *IssuanceCaller) StageRatio(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Issuance.contract.Call(opts, &out, "stageRatio")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StageRatio is a free data retrieval call binding the contract method 0xe72b5250.
//
// Solidity: function stageRatio() view returns(uint256)
func (_Issuance *IssuanceSession) StageRatio() (*big.Int, error) {
	return _Issuance.Contract.StageRatio(&_Issuance.CallOpts)
}

// StageRatio is a free data retrieval call binding the contract method 0xe72b5250.
//
// Solidity: function stageRatio() view returns(uint256)
func (_Issuance *IssuanceCallerSession) StageRatio() (*big.Int, error) {
	return _Issuance.Contract.StageRatio(&_Issuance.CallOpts)
}

// StageReward is a free data retrieval call binding the contract method 0xe12e4b21.
//
// Solidity: function stageReward() view returns(uint256)
func (_Issuance *IssuanceCaller) StageReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Issuance.contract.Call(opts, &out, "stageReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StageReward is a free data retrieval call binding the contract method 0xe12e4b21.
//
// Solidity: function stageReward() view returns(uint256)
func (_Issuance *IssuanceSession) StageReward() (*big.Int, error) {
	return _Issuance.Contract.StageReward(&_Issuance.CallOpts)
}

// StageReward is a free data retrieval call binding the contract method 0xe12e4b21.
//
// Solidity: function stageReward() view returns(uint256)
func (_Issuance *IssuanceCallerSession) StageReward() (*big.Int, error) {
	return _Issuance.Contract.StageReward(&_Issuance.CallOpts)
}

// StageRewardResid is a free data retrieval call binding the contract method 0xbd6ba39a.
//
// Solidity: function stageRewardResid() view returns(uint256)
func (_Issuance *IssuanceCaller) StageRewardResid(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Issuance.contract.Call(opts, &out, "stageRewardResid")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StageRewardResid is a free data retrieval call binding the contract method 0xbd6ba39a.
//
// Solidity: function stageRewardResid() view returns(uint256)
func (_Issuance *IssuanceSession) StageRewardResid() (*big.Int, error) {
	return _Issuance.Contract.StageRewardResid(&_Issuance.CallOpts)
}

// StageRewardResid is a free data retrieval call binding the contract method 0xbd6ba39a.
//
// Solidity: function stageRewardResid() view returns(uint256)
func (_Issuance *IssuanceCallerSession) StageRewardResid() (*big.Int, error) {
	return _Issuance.Contract.StageRewardResid(&_Issuance.CallOpts)
}

// StageSize is a free data retrieval call binding the contract method 0xfb752b8e.
//
// Solidity: function stageSize() view returns(uint256)
func (_Issuance *IssuanceCaller) StageSize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Issuance.contract.Call(opts, &out, "stageSize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StageSize is a free data retrieval call binding the contract method 0xfb752b8e.
//
// Solidity: function stageSize() view returns(uint256)
func (_Issuance *IssuanceSession) StageSize() (*big.Int, error) {
	return _Issuance.Contract.StageSize(&_Issuance.CallOpts)
}

// StageSize is a free data retrieval call binding the contract method 0xfb752b8e.
//
// Solidity: function stageSize() view returns(uint256)
func (_Issuance *IssuanceCallerSession) StageSize() (*big.Int, error) {
	return _Issuance.Contract.StageSize(&_Issuance.CallOpts)
}

// TargetReward is a free data retrieval call binding the contract method 0xeda7461a.
//
// Solidity: function targetReward() view returns(uint256)
func (_Issuance *IssuanceCaller) TargetReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Issuance.contract.Call(opts, &out, "targetReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TargetReward is a free data retrieval call binding the contract method 0xeda7461a.
//
// Solidity: function targetReward() view returns(uint256)
func (_Issuance *IssuanceSession) TargetReward() (*big.Int, error) {
	return _Issuance.Contract.TargetReward(&_Issuance.CallOpts)
}

// TargetReward is a free data retrieval call binding the contract method 0xeda7461a.
//
// Solidity: function targetReward() view returns(uint256)
func (_Issuance *IssuanceCallerSession) TargetReward() (*big.Int, error) {
	return _Issuance.Contract.TargetReward(&_Issuance.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint16)
func (_Issuance *IssuanceCaller) Version(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Issuance.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint16)
func (_Issuance *IssuanceSession) Version() (uint16, error) {
	return _Issuance.Contract.Version(&_Issuance.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint16)
func (_Issuance *IssuanceCallerSession) Version() (uint16, error) {
	return _Issuance.Contract.Version(&_Issuance.CallOpts)
}

// Issu is a paid mutator transaction binding the contract method 0x13df3828.
//
// Solidity: function issu(uint64 _start, uint64 _end, uint64 _size, uint256 _sPrice) returns(uint256)
func (_Issuance *IssuanceTransactor) Issu(opts *bind.TransactOpts, _start uint64, _end uint64, _size uint64, _sPrice *big.Int) (*types.Transaction, error) {
	return _Issuance.contract.Transact(opts, "issu", _start, _end, _size, _sPrice)
}

// Issu is a paid mutator transaction binding the contract method 0x13df3828.
//
// Solidity: function issu(uint64 _start, uint64 _end, uint64 _size, uint256 _sPrice) returns(uint256)
func (_Issuance *IssuanceSession) Issu(_start uint64, _end uint64, _size uint64, _sPrice *big.Int) (*types.Transaction, error) {
	return _Issuance.Contract.Issu(&_Issuance.TransactOpts, _start, _end, _size, _sPrice)
}

// Issu is a paid mutator transaction binding the contract method 0x13df3828.
//
// Solidity: function issu(uint64 _start, uint64 _end, uint64 _size, uint256 _sPrice) returns(uint256)
func (_Issuance *IssuanceTransactorSession) Issu(_start uint64, _end uint64, _size uint64, _sPrice *big.Int) (*types.Transaction, error) {
	return _Issuance.Contract.Issu(&_Issuance.TransactOpts, _start, _end, _size, _sPrice)
}

// SetRole is a paid mutator transaction binding the contract method 0xcfbfe5f3.
//
// Solidity: function setRole(uint8 _role, address _a, bytes[5] signs) returns()
func (_Issuance *IssuanceTransactor) SetRole(opts *bind.TransactOpts, _role uint8, _a common.Address, signs [5][]byte) (*types.Transaction, error) {
	return _Issuance.contract.Transact(opts, "setRole", _role, _a, signs)
}

// SetRole is a paid mutator transaction binding the contract method 0xcfbfe5f3.
//
// Solidity: function setRole(uint8 _role, address _a, bytes[5] signs) returns()
func (_Issuance *IssuanceSession) SetRole(_role uint8, _a common.Address, signs [5][]byte) (*types.Transaction, error) {
	return _Issuance.Contract.SetRole(&_Issuance.TransactOpts, _role, _a, signs)
}

// SetRole is a paid mutator transaction binding the contract method 0xcfbfe5f3.
//
// Solidity: function setRole(uint8 _role, address _a, bytes[5] signs) returns()
func (_Issuance *IssuanceTransactorSession) SetRole(_role uint8, _a common.Address, signs [5][]byte) (*types.Transaction, error) {
	return _Issuance.Contract.SetRole(&_Issuance.TransactOpts, _role, _a, signs)
}

// IssuanceSetRoleIterator is returned from FilterSetRole and is used to iterate over the raw logs and unpacked data for SetRole events raised by the Issuance contract.
type IssuanceSetRoleIterator struct {
	Event *IssuanceSetRole // Event containing the contract specifics and raw log

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
func (it *IssuanceSetRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IssuanceSetRole)
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
		it.Event = new(IssuanceSetRole)
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
func (it *IssuanceSetRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IssuanceSetRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IssuanceSetRole represents a SetRole event raised by the Issuance contract.
type IssuanceSetRole struct {
	Role uint8
	A    common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSetRole is a free log retrieval operation binding the contract event 0x5c22095cd1123887402a709a9698bc2961c12cb22eeb16e5d76480931f63ede0.
//
// Solidity: event SetRole(uint8 indexed role, address indexed a)
func (_Issuance *IssuanceFilterer) FilterSetRole(opts *bind.FilterOpts, role []uint8, a []common.Address) (*IssuanceSetRoleIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var aRule []interface{}
	for _, aItem := range a {
		aRule = append(aRule, aItem)
	}

	logs, sub, err := _Issuance.contract.FilterLogs(opts, "SetRole", roleRule, aRule)
	if err != nil {
		return nil, err
	}
	return &IssuanceSetRoleIterator{contract: _Issuance.contract, event: "SetRole", logs: logs, sub: sub}, nil
}

// WatchSetRole is a free log subscription operation binding the contract event 0x5c22095cd1123887402a709a9698bc2961c12cb22eeb16e5d76480931f63ede0.
//
// Solidity: event SetRole(uint8 indexed role, address indexed a)
func (_Issuance *IssuanceFilterer) WatchSetRole(opts *bind.WatchOpts, sink chan<- *IssuanceSetRole, role []uint8, a []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var aRule []interface{}
	for _, aItem := range a {
		aRule = append(aRule, aItem)
	}

	logs, sub, err := _Issuance.contract.WatchLogs(opts, "SetRole", roleRule, aRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IssuanceSetRole)
				if err := _Issuance.contract.UnpackLog(event, "SetRole", log); err != nil {
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

// ParseSetRole is a log parse operation binding the contract event 0x5c22095cd1123887402a709a9698bc2961c12cb22eeb16e5d76480931f63ede0.
//
// Solidity: event SetRole(uint8 indexed role, address indexed a)
func (_Issuance *IssuanceFilterer) ParseSetRole(log types.Log) (*IssuanceSetRole, error) {
	event := new(IssuanceSetRole)
	if err := _Issuance.contract.UnpackLog(event, "SetRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
