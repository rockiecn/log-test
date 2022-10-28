// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package fs

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

// FsOut is an auto generated low-level Go binding around an user-defined struct.
type FsOut struct {
	Nonce    uint64
	SubNonce uint64
}

// GroupOut is an auto generated low-level Go binding around an user-defined struct.
type GroupOut struct {
	Size   uint64
	Sprice *big.Int
	Lost   *big.Int
}

// OrderIn is an auto generated low-level Go binding around an user-defined struct.
type OrderIn struct {
	UIndex uint64
	PIndex uint64
	Start  uint64
	End    uint64
	Size   uint64
	Nonce  uint64
	TIndex uint8
	Sprice *big.Int
}

// PWIn is an auto generated low-level Go binding around an user-defined struct.
type PWIn struct {
	PIndex uint64
	TIndex uint8
	Pay    *big.Int
	Lost   *big.Int
}

// SettleOut is an auto generated low-level Go binding around an user-defined struct.
type SettleOut struct {
	Time    uint64
	Size    uint64
	Price   *big.Int
	MaxPay  *big.Int
	HasPaid *big.Int
	CanPay  *big.Int
	Lost    *big.Int
}

// StoreOut is an auto generated low-level Go binding around an user-defined struct.
type StoreOut struct {
	Start  uint64
	End    uint64
	Size   uint64
	Sprice *big.Int
}

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

// FileSysABI is the input ABI used to generate the binding from.
const FileSysABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"ti\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"ui\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"pi\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"start\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"end\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sprice\",\"type\":\"uint256\"}],\"name\":\"AddOrder\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"ti\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"ui\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"pi\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"start\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"end\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sprice\",\"type\":\"uint256\"}],\"name\":\"AddRepair\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"role\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"SetRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"ti\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"ui\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"pi\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"start\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"end\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sprice\",\"type\":\"uint256\"}],\"name\":\"SubOrder\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"uIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"start\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"end\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"tIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"sprice\",\"type\":\"uint256\"}],\"internalType\":\"structOrderIn\",\"name\":\"ps\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"mr\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"tr\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"}],\"name\":\"addOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_tc\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_rtype\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_kPB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_pPB\",\"type\":\"uint256\"}],\"name\":\"addPenalty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"uIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"start\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"end\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"tIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"sprice\",\"type\":\"uint256\"}],\"internalType\":\"structOrderIn\",\"name\":\"ps\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"}],\"name\":\"addRepair\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"auth\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_ui\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_pi\",\"type\":\"uint64\"}],\"name\":\"getFsInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"subNonce\",\"type\":\"uint64\"}],\"internalType\":\"structFsOut\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"}],\"name\":\"getGInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"sprice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lost\",\"type\":\"uint256\"}],\"internalType\":\"structGroupOut\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_pi\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"}],\"name\":\"getSettleInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"time\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"hasPaid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"canPay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lost\",\"type\":\"uint256\"}],\"internalType\":\"structSettleOut\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_ui\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_pi\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"}],\"name\":\"getStoreInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"start\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"end\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"sprice\",\"type\":\"uint256\"}],\"internalType\":\"structStoreOut\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_role\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"pIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"tIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"pay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lost\",\"type\":\"uint256\"}],\"internalType\":\"structPWIn\",\"name\":\"ps\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"}],\"name\":\"proWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isV\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"money\",\"type\":\"uint256\"}],\"name\":\"recharge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_role\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"bytes[5]\",\"name\":\"signs\",\"type\":\"bytes[5]\"}],\"name\":\"setRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"uIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"start\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"end\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"tIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"sprice\",\"type\":\"uint256\"}],\"internalType\":\"structOrderIn\",\"name\":\"ps\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"}],\"name\":\"subOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_lock\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// FileSysFuncSigs maps the 4-byte function signature to its string representation.
var FileSysFuncSigs = map[string]string{
	"9058d8e7": "addOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256),uint8,uint8,uint64)",
	"2449a168": "addPenalty(uint64,uint64,uint8,uint8,uint256,uint256)",
	"80faaf88": "addRepair((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256),uint64)",
	"de9375f2": "auth()",
	"fc3ba0ad": "balanceOf(uint64,uint8)",
	"59f7b572": "getFsInfo(uint64,uint64)",
	"5475cf8d": "getGInfo(uint64,uint8)",
	"64c64a4d": "getSettleInfo(uint64,uint8)",
	"8f328ecd": "getStoreInfo(uint64,uint64,uint8)",
	"9e97b8f6": "hasRole(uint8,address)",
	"a4703e16": "proWithdraw((uint64,uint8,uint256,uint256),uint64)",
	"24d11d40": "recharge(uint64,uint8,bool,uint256)",
	"cfbfe5f3": "setRole(uint8,address,bytes[5])",
	"248d02a0": "subOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256),uint64)",
	"54fd4d50": "version()",
	"5affa0f3": "withdraw(uint64,uint8,uint256,uint256)",
}

// FileSysBin is the compiled bytecode used for deploying new contracts.
var FileSysBin = "0x60806040526001805461ffff60a01b1916600160a11b1790553480156200002557600080fd5b5060405162002d0338038062002d0383398101604081905262000048916200006e565b600180546001600160a01b0319166001600160a01b0392909216919091179055620000a0565b6000602082840312156200008157600080fd5b81516001600160a01b03811681146200009957600080fd5b9392505050565b612c5380620000b06000396000f3fe608060405234801561001057600080fd5b50600436106101005760003560e01c806380faaf8811610097578063a4703e1611610066578063a4703e161461042c578063cfbfe5f31461043f578063de9375f214610452578063fc3ba0ad1461047d57600080fd5b806380faaf88146102e55780638f328ecd146102f85780639058d8e7146103f65780639e97b8f61461040957600080fd5b806354fd4d50116100d357806354fd4d501461018357806359f7b572146101ab5780635affa0f31461024e57806364c64a4d1461026f57600080fd5b80632449a16814610105578063248d02a01461011a57806324d11d401461012d5780635475cf8d14610140575b600080fd5b61011861011336600461253f565b6104ab565b005b6101186101283660046126d3565b61072b565b61011861013b366004612709565b610b04565b61015361014e36600461275b565b610b9a565b6040805182516001600160401b031681526020808401519082015291810151908201526060015b60405180910390f35b60015461019890600160a01b900461ffff1681565b60405161ffff909116815260200161017a565b6102276101b9366004612785565b60408051808201825260008082526020918201819052825180840184528181528201819052825180840184526001600160401b0395861682526003835283822094861680835285845293822054808716825293909152928152600160401b9091049092169181019190915290565b6040805182516001600160401b03908116825260209384015116928101929092520161017a565b61026161025c3660046127af565b610c4e565b60405190815260200161017a565b61028261027d36600461275b565b610e30565b60405161017a9190600060e0820190506001600160401b038084511683528060208501511660208401525060408301516040830152606083015160608301526080830151608083015260a083015160a083015260c083015160c083015292915050565b6101186102f33660046126d3565b610ec7565b6103b66103063660046127f1565b6040805160808082018352600080835260208084018290528385018290526060938401829052845180840186528281528082018390528086018390528401829052845192830185526001600160401b03978816825260038152848220968816825295865283812060ff9590951680825260019586018088528583208054808b168652600160401b81048b16868b0152600160801b9004909916958401959095529052919093529201549082015290565b6040805182516001600160401b0390811682526020808501518216908301528383015116918101919091526060918201519181019190915260800161017a565b610118610404366004612834565b6110eb565b61041c6104173660046128a4565b611951565b604051901515815260200161017a565b61026161043a3660046128ce565b6119f3565b61011861044d366004612956565b611aaf565b600154610465906001600160a01b031681565b6040516001600160a01b03909116815260200161017a565b61049061048b36600461275b565b611c21565b6040805193845260208401929092529082015260600161017a565b60046104b78133611d7f565b6000805b8660ff168160ff161015610720578560ff16600203610516576001600160401b03808916600090815260046020908152604080832060ff8616845290915290205461050f918691600160401b900416612a7a565b915061055a565b8560ff16600303610555576001600160401b03808a16600090815260026020908152604080832060ff8616845290915290205461050f91879116612a7a565b610720565b6001600160401b038816600090815260056020908152604080832060ff85168452909152902054821161060e576001600160401b038816600090815260056020908152604080832060ff85168452909152812080548492906105bd908490612a99565b909155505060ff811660009081527f05b8ccbb9d4d8fb16ea74ce3c29a41f1b461fbdaff4714a0d9a8eb05499746bc602052604081208054849290610603908490612aac565b9091555061070a9050565b6001600160401b038816600090815260056020908152604080832060ff8516845290915290205461063f9083612a99565b6001600160401b038916600090815260076020908152604080832060ff8616845290915281208054929450849290919061067a908490612aac565b90915550506001600160401b038816600090815260056020908152604080832060ff851684528252808320547f05b8ccbb9d4d8fb16ea74ce3c29a41f1b461fbdaff4714a0d9a8eb05499746bc90925282208054919290916106dd908490612aac565b90915550506001600160401b038816600090815260056020908152604080832060ff851684529091528120555b600091508061071881612abf565b9150506104bb565b505050505050505050565b60026107378133611d7f565b60a083015183516001600160401b0390811660009081526003602090815260408083208289015185168452909152902054600160401b900481169116146107c55760405162461bcd60e51b815260206004820152601b60248201527f6e6f6e6365206e6f7420657175616c2066732e7375626e6f6e6365000000000060448201526064015b60405180910390fd5b60a083015183516001600160401b03908116600090815260036020908152604080832082890151851684529091529020549181169116116108485760405162461bcd60e51b815260206004820152601760248201527f6e6f6e6365206e6f74206c6573732066732e6e6f6e636500000000000000000060448201526064016107bc565b60e083015183516001600160401b03908116600090815260036020908152604080832082890151909416835292815282822060c088015160ff168352600190810190915291812090910180549091906108a2908490612a99565b9091555050608083015183516001600160401b0390811660009081526003602090815260408083208289015185168452825280832060c089015160ff168452600101909152902080549091601091610903918591600160801b900416612ade565b92506101000a8154816001600160401b0302191690836001600160401b0316021790555061094883602001518460c00151856060015186608001518760e00151611e1f565b82516001600160401b039081166000908152600360209081526040808320828801518516845290915290208054600160401b900490911690600861098b83612afe565b82546101009290920a6001600160401b038181021990931691831602179091556080850151848216600090815260026020908152604080832060c08a015160ff168452909152812080549294509290916109e791859116612ade565b92506101000a8154816001600160401b0302191690836001600160401b031602179055508260e0015160026000846001600160401b03166001600160401b0316815260200190815260200160002060008560c0015160ff1660ff1681526020019081526020016000206001016000828254610a629190612a99565b9250508190555082602001516001600160401b031683600001516001600160401b03168460c0015160ff167fe42ef64dc7dfe9121f13026a70df328cf6687348fdf8f71745f8e40fa5e52b828660400151876060015188608001518960e00151604051610af794939291906001600160401b039485168152928416602084015292166040820152606081019190915260800190565b60405180910390a4505050565b6003610b108133611d7f565b8215610b57576001600160401b038516600090815260066020908152604080832060ff8816845290915281208054849290610b4c908490612aac565b90915550610b939050565b6001600160401b038516600090815260056020908152604080832060ff8816845290915281208054849290610b8d908490612aac565b90915550505b5050505050565b610bc7604051806060016040528060006001600160401b0316815260200160008152602001600081525090565b610bf4604051806060016040528060006001600160401b0316815260200160008152602001600081525090565b5050604080516060810182526001600160401b0380851660009081526002602081815285832060ff881680855281835287852080549096168752600186015483880152909352919091520154918101919091525b92915050565b60006004610c5c8133611d7f565b6001600160401b038616600081815260076020908152604080832060ff8a16808552908352818420549484526005835281842090845290915290205410610d19576001600160401b038616600081815260076020908152604080832060ff8a16808552908352818420549484526005835281842090845290915281208054909190610ce8908490612a99565b90915550506001600160401b038616600090815260076020908152604080832060ff89168452909152812055610d91565b6001600160401b038616600081815260056020908152604080832060ff8a16808552908352818420549484526007835281842090845290915281208054909190610d64908490612a99565b90915550506001600160401b038616600090815260056020908152604080832060ff891684529091528120555b6001600160401b038616600090815260056020908152604080832060ff8916845290915290205483811115610dd157610dca8482612a99565b9050610dd5565b5060005b80851180610de1575084155b15610dea578094505b6001600160401b038716600090815260056020908152604080832060ff8a16845290915281208054879290610e20908490612a99565b9091555094979650505050505050565b610e386124c3565b610e406124c3565b50506040805160e0810182526001600160401b0393841660009081526004602081815284832060ff969096168084528682528584208054808a168752600160401b90049098168286015260018801549585019590955260028701546060850152600387015460808501529086015460a084015292905291905260059091015460c082015290565b6002610ed38133611d7f565b60a083015183516001600160401b03908116600090815260036020908152604080832082890151851684529091529020548116911614610f555760405162461bcd60e51b815260206004820152601860248201527f6e6f6e6365206e6f7420657175616c2066732e6e6f6e6365000000000000000060448201526064016107bc565b60008360e0015184604001518560600151610f709190612ade565b6001600160401b0316610f839190612a7a565b6001600160401b038416600090815260026020818152604080842060c08a015160ff16855290915282200180549293508392909190610fc3908490612a99565b92505081905550610fec84602001518560c00151866060015187608001518860e001518661201d565b83516001600160401b03908116600090815260036020908152604080832082890151851684529091528120805460019391929161102b91859116612b24565b92506101000a8154816001600160401b0302191690836001600160401b0316021790555083602001516001600160401b031684600001516001600160401b03168560c0015160ff167f82f944cf0134663ff19eea526226f1cd8d5746a71cc837983659b8cbb2674e488760400151886060015189608001518a60e001516040516110dd94939291906001600160401b039485168152928416602084015292166040820152606081019190915260800190565b60405180910390a450505050565b60026110f78133611d7f565b60008560e00151866040015187606001516111129190612ade565b6001600160401b03166111259190612a7a565b9050600060ff8616611138606484612b44565b6111429190612a7a565b9050600060ff8616611155606485612b44565b61115f9190612a7a565b905060008161116e8486612aac565b6111789190612aac565b905080600660008b600001516001600160401b03166001600160401b0316815260200190815260200160002060008b60c0015160ff1660ff16815260200190815260200160002054600560008c600001516001600160401b03166001600160401b0316815260200190815260200160002060008c60c0015160ff1660ff1681526020019081526020016000205461120f9190612aac565b10156112535760405162461bcd60e51b81526020600482015260136024820152720ea40cce65ac4c2d840dcdee840cadcdeeaced606b1b60448201526064016107bc565b60a089015189516001600160401b039081166000908152600360209081526040808320828f01518516845290915290205481169116146112c15760405162461bcd60e51b81526020600482015260096024820152683737b731b29032b93960b91b60448201526064016107bc565b6040808a01518a516001600160401b03908116600090815260036020908152848220818f015184168352815284822060c08f015160ff16835260010190529290922054908216911611156113435760405162461bcd60e51b815260206004820152600960248201526839ba30b93a1032b93960b91b60448201526064016107bc565b606089015189516001600160401b039081166000908152600360209081526040808320828f015185168452825280832060c08f015160ff168452600101909152902054918116600160401b9092041611156113ca5760405162461bcd60e51b815260206004820152600760248201526632b7321032b93960c91b60448201526064016107bc565b60e089015189516001600160401b039081166000908152600360209081526040808320828f0151909416835292815282822060c08e015160ff16835260019081019091529181209091018054909190611424908490612aac565b9091555050608089015189516001600160401b039081166000908152600360209081526040808320828f015185168452825280832060c08f015160ff168452600101909152902080549091601091611485918591600160801b900416612b24565b92506101000a8154816001600160401b0302191690836001600160401b031602179055508860400151600360008b600001516001600160401b03166001600160401b0316815260200190815260200160002060008b602001516001600160401b03166001600160401b0316815260200190815260200160002060010160008b60c0015160ff1660ff16815260200190815260200160002060000160006101000a8154816001600160401b0302191690836001600160401b031602179055508860600151600360008b600001516001600160401b03166001600160401b0316815260200190815260200160002060008b602001516001600160401b03166001600160401b0316815260200190815260200160002060010160008b60c0015160ff1660ff16815260200190815260200160002060000160086101000a8154816001600160401b0302191690836001600160401b031602179055506115ff89602001518a60c001518b604001518c608001518d60e001518961201d565b88516001600160401b039081166000908152600360209081526040808320828e01518516845290915281208054909216919061163a83612afe565b91906101000a8154816001600160401b0302191690836001600160401b0316021790555050886080015160026000886001600160401b03166001600160401b0316815260200190815260200160002060008b60c0015160ff1660ff16815260200190815260200160002060000160008282829054906101000a90046001600160401b03166116c89190612b24565b92506101000a8154816001600160401b0302191690836001600160401b031602179055508860e0015160026000886001600160401b03166001600160401b0316815260200190815260200160002060008b60c0015160ff1660ff16815260200190815260200160002060010160008282546117439190612aac565b909155505060008080526005602090815260c08b015160ff1682527f05b8ccbb9d4d8fb16ea74ce3c29a41f1b461fbdaff4714a0d9a8eb05499746bc9052604081208054849290611795908490612aac565b909155505088516001600160401b0316600090815260066020908152604080832060c08d015160ff16845290915290205481116118125788516001600160401b0316600090815260066020908152604080832060c08d015160ff16845290915281208054839290611807908490612a99565b909155506118b09050565b88516001600160401b0316600090815260066020908152604080832060c08d015160ff1684529091529020546118489082612a99565b89516001600160401b03908116600090815260066020908152604080832060c08f01805160ff90811686529184528285208590558f519095168452600583528184209451168352929052908120805492935083929091906118aa908490612a99565b90915550505b88602001516001600160401b031689600001516001600160401b03168a60c0015160ff167f9ad187def591246eb8fff84534926005f62988636593745516c0ac08fa0863148c604001518d606001518e608001518f60e0015160405161193e94939291906001600160401b039485168152928416602084015292166040820152606081019190915260800190565b60405180910390a4505050505050505050565b60ff82166000908152602081905260408120546001600160a01b0380841691160361197e57506001610c48565b6000805b60088160ff1610156119e857806002188516915060008260ff161180156119c6575060ff82166000908152602081905260409020546001600160a01b038581169116145b156119d657600192505050610c48565b806119e081612abf565b915050611982565b506000949350505050565b60006002611a018133611d7f565b600080611a208660000151876020015188604001518960600151612280565b6001600160401b03871660009081526002602081815260408084208c83015160ff16855290915282200180549395509193508392611a5f908490612aac565b909155505085516001600160401b03166000908152600560209081526040808320828a015160ff16845290915281208054849290611a9e908490612aac565b9091555091935050505b5092915050565b813b6000819003611af75760405162461bcd60e51b81526020600482015260126024820152713732b2b21031b7b73a3930b1ba1030b2323960711b60448201526064016107bc565b6040516bffffffffffffffffffffffff1930606090811b8216602084015266736574526f6c6560c81b603484015285901b16603b8201526001600160f81b031960f886901b16604f82015260009060500160408051601f1981840301815290829052805160209091012060015463a96bba9d60e01b83529092506001600160a01b03169063a96bba9d90611b919084908790600401612bac565b600060405180830381600087803b158015611bab57600080fd5b505af1158015611bbf573d6000803e3d6000fd5b5050505060ff851660008181526020819052604080822080546001600160a01b0319166001600160a01b03891690811790915590519092917f5c22095cd1123887402a709a9698bc2961c12cb22eeb16e5d76480931f63ede091a35050505050565b6001600160401b03808316600081815260056020818152604080842060ff88168086529083528185205486865260068452828620828752845282862054878752600785528387208388528552838720549787526004808652848820938852928552838720845160e0810186528154808c16808352600160401b909104909b169681019690965260018101549486019490945260028401546060860152600384015460808601529183015460a08501819052929094015460c0840152939586958695939092909115611d705781516001600160401b0316421115611d2f5760408201518251611d18906001600160401b031642612a99565b611d229190612a7a565b611d2c9082612aac565b90505b60008260c001518360600151611d459190612a99565b905080821115611d53578091505b6080830151611d629083612a99565b611d6c9086612aac565b9450505b50929891975095509350505050565b611d898282611951565b611e1b5760405171024b739ba30b731b29d1030b1b1b7bab73a160751b60208201526bffffffffffffffffffffffff19606083901b1660328201527001034b99036b4b9b9b4b733903937b6329607d1b60468201526001600160f81b031960f884901b16605782015260580160408051601f198184030181529082905262461bcd60e51b82526107bc91600401612c03565b5050565b6001600160401b03808616600090815260046020908152604080832060ff8916845290915290205481169084168111611f17576001600160401b038616600090815260046020908152604080832060ff89168452909152902060010154611e868286612ade565b6001600160401b0316611e999190612a7a565b6001600160401b038716600090815260046020818152604080842060ff8b1685529091528220018054909190611ed0908490612aac565b90915550506001600160401b03868116600090815260046020908152604080832060ff8a1684529091529020805467ffffffffffffffff1916918616919091179055611f72565b81611f228583612ade565b6001600160401b0316611f359190612a7a565b6001600160401b038716600090815260046020818152604080842060ff8b1685529091528220018054909190611f6c908490612a99565b90915550505b6001600160401b038616600090815260046020908152604080832060ff8916845290915281206001018054849290611fab908490612a99565b90915550506001600160401b03808716600090815260046020908152604080832060ff8a168452909152902080548592600891611ff1918591600160401b900416612ade565b92506101000a8154816001600160401b0302191690836001600160401b03160217905550505050505050565b6001600160401b03808716600090815260046020908152604080832060ff8a16845290915290205481169085168111612125576001600160401b038116156120e3576001600160401b038716600090815260046020908152604080832060ff8a1684529091529020600101546120938287612ade565b6001600160401b03166120a69190612a7a565b6001600160401b038816600090815260046020818152604080842060ff8c16855290915282200180549091906120dd908490612aac565b90915550505b6001600160401b03878116600090815260046020908152604080832060ff8b1684529091529020805467ffffffffffffffff1916918716919091179055612180565b826121308683612ade565b6001600160401b03166121439190612a7a565b6001600160401b038816600090815260046020818152604080842060ff8c168552909152822001805490919061217a908490612aac565b90915550505b6001600160401b038716600090815260046020908152604080832060ff8a168452909152812060010180548592906121b9908490612aac565b90915550506001600160401b03808816600090815260046020908152604080832060ff8b1684529091529020805486926008916121ff918591600160401b900416612b24565b92506101000a8154816001600160401b0302191690836001600160401b031602179055508160046000896001600160401b03166001600160401b0316815260200190815260200160002060008860ff1660ff16815260200190815260200160002060020160008282546122729190612aac565b909155505050505050505050565b6001600160401b03808516600090815260046020818152604080842060ff891685528252808420815160e08101835281548088168252600160401b900490961692860192909252600182015490850152600281015460608501526003810154608085018190529181015460a08501526005015460c08401529091829185108061230c5750838160c00151115b1561231e5760008092509250506124ba565b8484826060015161232f9190612a99565b10156123425760008092509250506124ba565b6001600160401b038716600090815260046020908152604080832060ff8a168452909152902060050184905560c081015161237d9085612a99565b815190945042906001600160401b038083169116101561243557604082015182516123a89083612ade565b6001600160401b03166123bb9190612a7a565b6001600160401b038916600090815260046020818152604080842060ff8d16855290915282200180549091906123f2908490612aac565b90915550506001600160401b03888116600090815260046020908152604080832060ff8c1684529091529020805467ffffffffffffffff19169183169190911790555b6001600160401b038816600090815260046020818152604080842060ff8c1685529091529091200154861115612473576000809350935050506124ba565b60008260800151876124859190612a99565b6001600160401b038a16600090815260046020908152604080832060ff8d168452909152902060030188905594508593505050505b94509492505050565b6040518060e0016040528060006001600160401b0316815260200160006001600160401b0316815260200160008152602001600081526020016000815260200160008152602001600081525090565b80356001600160401b038116811461252957600080fd5b919050565b803560ff8116811461252957600080fd5b60008060008060008060c0878903121561255857600080fd5b61256187612512565b955061256f60208801612512565b945061257d6040880161252e565b935061258b6060880161252e565b92506080870135915060a087013590509295509295509295565b634e487b7160e01b600052604160045260246000fd5b60405160a081016001600160401b03811182821017156125dd576125dd6125a5565b60405290565b604051601f8201601f191681016001600160401b038111828210171561260b5761260b6125a5565b604052919050565b600061010080838503121561262757600080fd5b604051908101906001600160401b0382118183101715612649576126496125a5565b8160405280925061265984612512565b815261266760208501612512565b602082015261267860408501612512565b604082015261268960608501612512565b606082015261269a60808501612512565b60808201526126ab60a08501612512565b60a08201526126bc60c0850161252e565b60c082015260e084013560e0820152505092915050565b60008061012083850312156126e757600080fd5b6126f18484612613565b91506127006101008401612512565b90509250929050565b6000806000806080858703121561271f57600080fd5b61272885612512565b93506127366020860161252e565b92506040850135801515811461274b57600080fd5b9396929550929360600135925050565b6000806040838503121561276e57600080fd5b61277783612512565b91506127006020840161252e565b6000806040838503121561279857600080fd5b6127a183612512565b915061270060208401612512565b600080600080608085870312156127c557600080fd5b6127ce85612512565b93506127dc6020860161252e565b93969395505050506040820135916060013590565b60008060006060848603121561280657600080fd5b61280f84612512565b925061281d60208501612512565b915061282b6040850161252e565b90509250925092565b600080600080610160858703121561284b57600080fd5b6128558686612613565b9350612864610100860161252e565b9250612873610120860161252e565b91506128826101408601612512565b905092959194509250565b80356001600160a01b038116811461252957600080fd5b600080604083850312156128b757600080fd5b6128c08361252e565b91506127006020840161288d565b60008082840360a08112156128e257600080fd5b60808112156128f057600080fd5b50604051608081018181106001600160401b0382111715612913576129136125a5565b60405261291f84612512565b815261292d6020850161252e565b602082015260408401356040820152606084013560608201528092505061270060808401612512565b60008060006060848603121561296b57600080fd5b6129748461252e565b9250602061298381860161288d565b925060408501356001600160401b038082111561299f57600080fd5b8187019150601f88818401126129b457600080fd5b6129bc6125bb565b8060a085018b8111156129ce57600080fd5b855b81811015612a52578035868111156129e85760008081fd5b87018581018e136129f95760008081fd5b803587811115612a0b57612a0b6125a5565b612a1c818801601f19168b016125e3565b8181528f8b838501011115612a315760008081fd5b818b84018c83013760009181018b01919091528552509287019287016129d0565b50508096505050505050509250925092565b634e487b7160e01b600052601160045260246000fd5b6000816000190483118215151615612a9457612a94612a64565b500290565b81810381811115610c4857610c48612a64565b80820180821115610c4857610c48612a64565b600060ff821660ff8103612ad557612ad5612a64565b60010192915050565b6001600160401b03828116828216039080821115611aa857611aa8612a64565b60006001600160401b03808316818103612b1a57612b1a612a64565b6001019392505050565b6001600160401b03818116838216019080821115611aa857611aa8612a64565b600082612b6157634e487b7160e01b600052601260045260246000fd5b500490565b6000815180845260005b81811015612b8c57602081850181015186830182015201612b70565b506000602082860101526020601f19601f83011685010191505092915050565b8281526040602080830182905260009160e084019190840185845b6005811015612bf657603f19878603018352612be4858351612b66565b94509183019190830190600101612bc7565b5092979650505050505050565b602081526000612c166020830184612b66565b939250505056fea2646970667358221220567892897d08535d6ebd66311b7103edb954485cb5e6aed9a19372c9a54be5cb64736f6c63430008100033"

// DeployFileSys deploys a new Ethereum contract, binding an instance of FileSys to it.
func DeployFileSys(auth *bind.TransactOpts, backend bind.ContractBackend, _a common.Address) (common.Address, *types.Transaction, *FileSys, error) {
	parsed, err := abi.JSON(strings.NewReader(FileSysABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(FileSysBin), backend, _a)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FileSys{FileSysCaller: FileSysCaller{contract: contract}, FileSysTransactor: FileSysTransactor{contract: contract}, FileSysFilterer: FileSysFilterer{contract: contract}}, nil
}

// FileSys is an auto generated Go binding around an Ethereum contract.
type FileSys struct {
	FileSysCaller     // Read-only binding to the contract
	FileSysTransactor // Write-only binding to the contract
	FileSysFilterer   // Log filterer for contract events
}

// FileSysCaller is an auto generated read-only Go binding around an Ethereum contract.
type FileSysCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FileSysTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FileSysTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FileSysFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FileSysFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FileSysSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FileSysSession struct {
	Contract     *FileSys          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FileSysCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FileSysCallerSession struct {
	Contract *FileSysCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// FileSysTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FileSysTransactorSession struct {
	Contract     *FileSysTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// FileSysRaw is an auto generated low-level Go binding around an Ethereum contract.
type FileSysRaw struct {
	Contract *FileSys // Generic contract binding to access the raw methods on
}

// FileSysCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FileSysCallerRaw struct {
	Contract *FileSysCaller // Generic read-only contract binding to access the raw methods on
}

// FileSysTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FileSysTransactorRaw struct {
	Contract *FileSysTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFileSys creates a new instance of FileSys, bound to a specific deployed contract.
func NewFileSys(address common.Address, backend bind.ContractBackend) (*FileSys, error) {
	contract, err := bindFileSys(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FileSys{FileSysCaller: FileSysCaller{contract: contract}, FileSysTransactor: FileSysTransactor{contract: contract}, FileSysFilterer: FileSysFilterer{contract: contract}}, nil
}

// NewFileSysCaller creates a new read-only instance of FileSys, bound to a specific deployed contract.
func NewFileSysCaller(address common.Address, caller bind.ContractCaller) (*FileSysCaller, error) {
	contract, err := bindFileSys(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FileSysCaller{contract: contract}, nil
}

// NewFileSysTransactor creates a new write-only instance of FileSys, bound to a specific deployed contract.
func NewFileSysTransactor(address common.Address, transactor bind.ContractTransactor) (*FileSysTransactor, error) {
	contract, err := bindFileSys(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FileSysTransactor{contract: contract}, nil
}

// NewFileSysFilterer creates a new log filterer instance of FileSys, bound to a specific deployed contract.
func NewFileSysFilterer(address common.Address, filterer bind.ContractFilterer) (*FileSysFilterer, error) {
	contract, err := bindFileSys(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FileSysFilterer{contract: contract}, nil
}

// bindFileSys binds a generic wrapper to an already deployed contract.
func bindFileSys(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FileSysABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FileSys *FileSysRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FileSys.Contract.FileSysCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FileSys *FileSysRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FileSys.Contract.FileSysTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FileSys *FileSysRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FileSys.Contract.FileSysTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FileSys *FileSysCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FileSys.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FileSys *FileSysTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FileSys.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FileSys *FileSysTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FileSys.Contract.contract.Transact(opts, method, params...)
}

// Auth is a free data retrieval call binding the contract method 0xde9375f2.
//
// Solidity: function auth() view returns(address)
func (_FileSys *FileSysCaller) Auth(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FileSys.contract.Call(opts, &out, "auth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Auth is a free data retrieval call binding the contract method 0xde9375f2.
//
// Solidity: function auth() view returns(address)
func (_FileSys *FileSysSession) Auth() (common.Address, error) {
	return _FileSys.Contract.Auth(&_FileSys.CallOpts)
}

// Auth is a free data retrieval call binding the contract method 0xde9375f2.
//
// Solidity: function auth() view returns(address)
func (_FileSys *FileSysCallerSession) Auth() (common.Address, error) {
	return _FileSys.Contract.Auth(&_FileSys.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0xfc3ba0ad.
//
// Solidity: function balanceOf(uint64 _i, uint8 _ti) view returns(uint256, uint256, uint256)
func (_FileSys *FileSysCaller) BalanceOf(opts *bind.CallOpts, _i uint64, _ti uint8) (*big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _FileSys.contract.Call(opts, &out, "balanceOf", _i, _ti)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// BalanceOf is a free data retrieval call binding the contract method 0xfc3ba0ad.
//
// Solidity: function balanceOf(uint64 _i, uint8 _ti) view returns(uint256, uint256, uint256)
func (_FileSys *FileSysSession) BalanceOf(_i uint64, _ti uint8) (*big.Int, *big.Int, *big.Int, error) {
	return _FileSys.Contract.BalanceOf(&_FileSys.CallOpts, _i, _ti)
}

// BalanceOf is a free data retrieval call binding the contract method 0xfc3ba0ad.
//
// Solidity: function balanceOf(uint64 _i, uint8 _ti) view returns(uint256, uint256, uint256)
func (_FileSys *FileSysCallerSession) BalanceOf(_i uint64, _ti uint8) (*big.Int, *big.Int, *big.Int, error) {
	return _FileSys.Contract.BalanceOf(&_FileSys.CallOpts, _i, _ti)
}

// GetFsInfo is a free data retrieval call binding the contract method 0x59f7b572.
//
// Solidity: function getFsInfo(uint64 _ui, uint64 _pi) view returns((uint64,uint64))
func (_FileSys *FileSysCaller) GetFsInfo(opts *bind.CallOpts, _ui uint64, _pi uint64) (FsOut, error) {
	var out []interface{}
	err := _FileSys.contract.Call(opts, &out, "getFsInfo", _ui, _pi)

	if err != nil {
		return *new(FsOut), err
	}

	out0 := *abi.ConvertType(out[0], new(FsOut)).(*FsOut)

	return out0, err

}

// GetFsInfo is a free data retrieval call binding the contract method 0x59f7b572.
//
// Solidity: function getFsInfo(uint64 _ui, uint64 _pi) view returns((uint64,uint64))
func (_FileSys *FileSysSession) GetFsInfo(_ui uint64, _pi uint64) (FsOut, error) {
	return _FileSys.Contract.GetFsInfo(&_FileSys.CallOpts, _ui, _pi)
}

// GetFsInfo is a free data retrieval call binding the contract method 0x59f7b572.
//
// Solidity: function getFsInfo(uint64 _ui, uint64 _pi) view returns((uint64,uint64))
func (_FileSys *FileSysCallerSession) GetFsInfo(_ui uint64, _pi uint64) (FsOut, error) {
	return _FileSys.Contract.GetFsInfo(&_FileSys.CallOpts, _ui, _pi)
}

// GetGInfo is a free data retrieval call binding the contract method 0x5475cf8d.
//
// Solidity: function getGInfo(uint64 _gi, uint8 _ti) view returns((uint64,uint256,uint256))
func (_FileSys *FileSysCaller) GetGInfo(opts *bind.CallOpts, _gi uint64, _ti uint8) (GroupOut, error) {
	var out []interface{}
	err := _FileSys.contract.Call(opts, &out, "getGInfo", _gi, _ti)

	if err != nil {
		return *new(GroupOut), err
	}

	out0 := *abi.ConvertType(out[0], new(GroupOut)).(*GroupOut)

	return out0, err

}

// GetGInfo is a free data retrieval call binding the contract method 0x5475cf8d.
//
// Solidity: function getGInfo(uint64 _gi, uint8 _ti) view returns((uint64,uint256,uint256))
func (_FileSys *FileSysSession) GetGInfo(_gi uint64, _ti uint8) (GroupOut, error) {
	return _FileSys.Contract.GetGInfo(&_FileSys.CallOpts, _gi, _ti)
}

// GetGInfo is a free data retrieval call binding the contract method 0x5475cf8d.
//
// Solidity: function getGInfo(uint64 _gi, uint8 _ti) view returns((uint64,uint256,uint256))
func (_FileSys *FileSysCallerSession) GetGInfo(_gi uint64, _ti uint8) (GroupOut, error) {
	return _FileSys.Contract.GetGInfo(&_FileSys.CallOpts, _gi, _ti)
}

// GetSettleInfo is a free data retrieval call binding the contract method 0x64c64a4d.
//
// Solidity: function getSettleInfo(uint64 _pi, uint8 _ti) view returns((uint64,uint64,uint256,uint256,uint256,uint256,uint256))
func (_FileSys *FileSysCaller) GetSettleInfo(opts *bind.CallOpts, _pi uint64, _ti uint8) (SettleOut, error) {
	var out []interface{}
	err := _FileSys.contract.Call(opts, &out, "getSettleInfo", _pi, _ti)

	if err != nil {
		return *new(SettleOut), err
	}

	out0 := *abi.ConvertType(out[0], new(SettleOut)).(*SettleOut)

	return out0, err

}

// GetSettleInfo is a free data retrieval call binding the contract method 0x64c64a4d.
//
// Solidity: function getSettleInfo(uint64 _pi, uint8 _ti) view returns((uint64,uint64,uint256,uint256,uint256,uint256,uint256))
func (_FileSys *FileSysSession) GetSettleInfo(_pi uint64, _ti uint8) (SettleOut, error) {
	return _FileSys.Contract.GetSettleInfo(&_FileSys.CallOpts, _pi, _ti)
}

// GetSettleInfo is a free data retrieval call binding the contract method 0x64c64a4d.
//
// Solidity: function getSettleInfo(uint64 _pi, uint8 _ti) view returns((uint64,uint64,uint256,uint256,uint256,uint256,uint256))
func (_FileSys *FileSysCallerSession) GetSettleInfo(_pi uint64, _ti uint8) (SettleOut, error) {
	return _FileSys.Contract.GetSettleInfo(&_FileSys.CallOpts, _pi, _ti)
}

// GetStoreInfo is a free data retrieval call binding the contract method 0x8f328ecd.
//
// Solidity: function getStoreInfo(uint64 _ui, uint64 _pi, uint8 _ti) view returns((uint64,uint64,uint64,uint256))
func (_FileSys *FileSysCaller) GetStoreInfo(opts *bind.CallOpts, _ui uint64, _pi uint64, _ti uint8) (StoreOut, error) {
	var out []interface{}
	err := _FileSys.contract.Call(opts, &out, "getStoreInfo", _ui, _pi, _ti)

	if err != nil {
		return *new(StoreOut), err
	}

	out0 := *abi.ConvertType(out[0], new(StoreOut)).(*StoreOut)

	return out0, err

}

// GetStoreInfo is a free data retrieval call binding the contract method 0x8f328ecd.
//
// Solidity: function getStoreInfo(uint64 _ui, uint64 _pi, uint8 _ti) view returns((uint64,uint64,uint64,uint256))
func (_FileSys *FileSysSession) GetStoreInfo(_ui uint64, _pi uint64, _ti uint8) (StoreOut, error) {
	return _FileSys.Contract.GetStoreInfo(&_FileSys.CallOpts, _ui, _pi, _ti)
}

// GetStoreInfo is a free data retrieval call binding the contract method 0x8f328ecd.
//
// Solidity: function getStoreInfo(uint64 _ui, uint64 _pi, uint8 _ti) view returns((uint64,uint64,uint64,uint256))
func (_FileSys *FileSysCallerSession) GetStoreInfo(_ui uint64, _pi uint64, _ti uint8) (StoreOut, error) {
	return _FileSys.Contract.GetStoreInfo(&_FileSys.CallOpts, _ui, _pi, _ti)
}

// HasRole is a free data retrieval call binding the contract method 0x9e97b8f6.
//
// Solidity: function hasRole(uint8 _role, address _a) view returns(bool)
func (_FileSys *FileSysCaller) HasRole(opts *bind.CallOpts, _role uint8, _a common.Address) (bool, error) {
	var out []interface{}
	err := _FileSys.contract.Call(opts, &out, "hasRole", _role, _a)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x9e97b8f6.
//
// Solidity: function hasRole(uint8 _role, address _a) view returns(bool)
func (_FileSys *FileSysSession) HasRole(_role uint8, _a common.Address) (bool, error) {
	return _FileSys.Contract.HasRole(&_FileSys.CallOpts, _role, _a)
}

// HasRole is a free data retrieval call binding the contract method 0x9e97b8f6.
//
// Solidity: function hasRole(uint8 _role, address _a) view returns(bool)
func (_FileSys *FileSysCallerSession) HasRole(_role uint8, _a common.Address) (bool, error) {
	return _FileSys.Contract.HasRole(&_FileSys.CallOpts, _role, _a)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint16)
func (_FileSys *FileSysCaller) Version(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _FileSys.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint16)
func (_FileSys *FileSysSession) Version() (uint16, error) {
	return _FileSys.Contract.Version(&_FileSys.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint16)
func (_FileSys *FileSysCallerSession) Version() (uint16, error) {
	return _FileSys.Contract.Version(&_FileSys.CallOpts)
}

// AddOrder is a paid mutator transaction binding the contract method 0x9058d8e7.
//
// Solidity: function addOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) ps, uint8 mr, uint8 tr, uint64 _gi) returns()
func (_FileSys *FileSysTransactor) AddOrder(opts *bind.TransactOpts, ps OrderIn, mr uint8, tr uint8, _gi uint64) (*types.Transaction, error) {
	return _FileSys.contract.Transact(opts, "addOrder", ps, mr, tr, _gi)
}

// AddOrder is a paid mutator transaction binding the contract method 0x9058d8e7.
//
// Solidity: function addOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) ps, uint8 mr, uint8 tr, uint64 _gi) returns()
func (_FileSys *FileSysSession) AddOrder(ps OrderIn, mr uint8, tr uint8, _gi uint64) (*types.Transaction, error) {
	return _FileSys.Contract.AddOrder(&_FileSys.TransactOpts, ps, mr, tr, _gi)
}

// AddOrder is a paid mutator transaction binding the contract method 0x9058d8e7.
//
// Solidity: function addOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) ps, uint8 mr, uint8 tr, uint64 _gi) returns()
func (_FileSys *FileSysTransactorSession) AddOrder(ps OrderIn, mr uint8, tr uint8, _gi uint64) (*types.Transaction, error) {
	return _FileSys.Contract.AddOrder(&_FileSys.TransactOpts, ps, mr, tr, _gi)
}

// AddPenalty is a paid mutator transaction binding the contract method 0x2449a168.
//
// Solidity: function addPenalty(uint64 _gi, uint64 _i, uint8 _tc, uint8 _rtype, uint256 _kPB, uint256 _pPB) returns()
func (_FileSys *FileSysTransactor) AddPenalty(opts *bind.TransactOpts, _gi uint64, _i uint64, _tc uint8, _rtype uint8, _kPB *big.Int, _pPB *big.Int) (*types.Transaction, error) {
	return _FileSys.contract.Transact(opts, "addPenalty", _gi, _i, _tc, _rtype, _kPB, _pPB)
}

// AddPenalty is a paid mutator transaction binding the contract method 0x2449a168.
//
// Solidity: function addPenalty(uint64 _gi, uint64 _i, uint8 _tc, uint8 _rtype, uint256 _kPB, uint256 _pPB) returns()
func (_FileSys *FileSysSession) AddPenalty(_gi uint64, _i uint64, _tc uint8, _rtype uint8, _kPB *big.Int, _pPB *big.Int) (*types.Transaction, error) {
	return _FileSys.Contract.AddPenalty(&_FileSys.TransactOpts, _gi, _i, _tc, _rtype, _kPB, _pPB)
}

// AddPenalty is a paid mutator transaction binding the contract method 0x2449a168.
//
// Solidity: function addPenalty(uint64 _gi, uint64 _i, uint8 _tc, uint8 _rtype, uint256 _kPB, uint256 _pPB) returns()
func (_FileSys *FileSysTransactorSession) AddPenalty(_gi uint64, _i uint64, _tc uint8, _rtype uint8, _kPB *big.Int, _pPB *big.Int) (*types.Transaction, error) {
	return _FileSys.Contract.AddPenalty(&_FileSys.TransactOpts, _gi, _i, _tc, _rtype, _kPB, _pPB)
}

// AddRepair is a paid mutator transaction binding the contract method 0x80faaf88.
//
// Solidity: function addRepair((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) ps, uint64 _gi) returns()
func (_FileSys *FileSysTransactor) AddRepair(opts *bind.TransactOpts, ps OrderIn, _gi uint64) (*types.Transaction, error) {
	return _FileSys.contract.Transact(opts, "addRepair", ps, _gi)
}

// AddRepair is a paid mutator transaction binding the contract method 0x80faaf88.
//
// Solidity: function addRepair((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) ps, uint64 _gi) returns()
func (_FileSys *FileSysSession) AddRepair(ps OrderIn, _gi uint64) (*types.Transaction, error) {
	return _FileSys.Contract.AddRepair(&_FileSys.TransactOpts, ps, _gi)
}

// AddRepair is a paid mutator transaction binding the contract method 0x80faaf88.
//
// Solidity: function addRepair((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) ps, uint64 _gi) returns()
func (_FileSys *FileSysTransactorSession) AddRepair(ps OrderIn, _gi uint64) (*types.Transaction, error) {
	return _FileSys.Contract.AddRepair(&_FileSys.TransactOpts, ps, _gi)
}

// ProWithdraw is a paid mutator transaction binding the contract method 0xa4703e16.
//
// Solidity: function proWithdraw((uint64,uint8,uint256,uint256) ps, uint64 _gi) returns(uint256)
func (_FileSys *FileSysTransactor) ProWithdraw(opts *bind.TransactOpts, ps PWIn, _gi uint64) (*types.Transaction, error) {
	return _FileSys.contract.Transact(opts, "proWithdraw", ps, _gi)
}

// ProWithdraw is a paid mutator transaction binding the contract method 0xa4703e16.
//
// Solidity: function proWithdraw((uint64,uint8,uint256,uint256) ps, uint64 _gi) returns(uint256)
func (_FileSys *FileSysSession) ProWithdraw(ps PWIn, _gi uint64) (*types.Transaction, error) {
	return _FileSys.Contract.ProWithdraw(&_FileSys.TransactOpts, ps, _gi)
}

// ProWithdraw is a paid mutator transaction binding the contract method 0xa4703e16.
//
// Solidity: function proWithdraw((uint64,uint8,uint256,uint256) ps, uint64 _gi) returns(uint256)
func (_FileSys *FileSysTransactorSession) ProWithdraw(ps PWIn, _gi uint64) (*types.Transaction, error) {
	return _FileSys.Contract.ProWithdraw(&_FileSys.TransactOpts, ps, _gi)
}

// Recharge is a paid mutator transaction binding the contract method 0x24d11d40.
//
// Solidity: function recharge(uint64 _i, uint8 _ti, bool isV, uint256 money) returns()
func (_FileSys *FileSysTransactor) Recharge(opts *bind.TransactOpts, _i uint64, _ti uint8, isV bool, money *big.Int) (*types.Transaction, error) {
	return _FileSys.contract.Transact(opts, "recharge", _i, _ti, isV, money)
}

// Recharge is a paid mutator transaction binding the contract method 0x24d11d40.
//
// Solidity: function recharge(uint64 _i, uint8 _ti, bool isV, uint256 money) returns()
func (_FileSys *FileSysSession) Recharge(_i uint64, _ti uint8, isV bool, money *big.Int) (*types.Transaction, error) {
	return _FileSys.Contract.Recharge(&_FileSys.TransactOpts, _i, _ti, isV, money)
}

// Recharge is a paid mutator transaction binding the contract method 0x24d11d40.
//
// Solidity: function recharge(uint64 _i, uint8 _ti, bool isV, uint256 money) returns()
func (_FileSys *FileSysTransactorSession) Recharge(_i uint64, _ti uint8, isV bool, money *big.Int) (*types.Transaction, error) {
	return _FileSys.Contract.Recharge(&_FileSys.TransactOpts, _i, _ti, isV, money)
}

// SetRole is a paid mutator transaction binding the contract method 0xcfbfe5f3.
//
// Solidity: function setRole(uint8 _role, address _a, bytes[5] signs) returns()
func (_FileSys *FileSysTransactor) SetRole(opts *bind.TransactOpts, _role uint8, _a common.Address, signs [5][]byte) (*types.Transaction, error) {
	return _FileSys.contract.Transact(opts, "setRole", _role, _a, signs)
}

// SetRole is a paid mutator transaction binding the contract method 0xcfbfe5f3.
//
// Solidity: function setRole(uint8 _role, address _a, bytes[5] signs) returns()
func (_FileSys *FileSysSession) SetRole(_role uint8, _a common.Address, signs [5][]byte) (*types.Transaction, error) {
	return _FileSys.Contract.SetRole(&_FileSys.TransactOpts, _role, _a, signs)
}

// SetRole is a paid mutator transaction binding the contract method 0xcfbfe5f3.
//
// Solidity: function setRole(uint8 _role, address _a, bytes[5] signs) returns()
func (_FileSys *FileSysTransactorSession) SetRole(_role uint8, _a common.Address, signs [5][]byte) (*types.Transaction, error) {
	return _FileSys.Contract.SetRole(&_FileSys.TransactOpts, _role, _a, signs)
}

// SubOrder is a paid mutator transaction binding the contract method 0x248d02a0.
//
// Solidity: function subOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) ps, uint64 _gi) returns()
func (_FileSys *FileSysTransactor) SubOrder(opts *bind.TransactOpts, ps OrderIn, _gi uint64) (*types.Transaction, error) {
	return _FileSys.contract.Transact(opts, "subOrder", ps, _gi)
}

// SubOrder is a paid mutator transaction binding the contract method 0x248d02a0.
//
// Solidity: function subOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) ps, uint64 _gi) returns()
func (_FileSys *FileSysSession) SubOrder(ps OrderIn, _gi uint64) (*types.Transaction, error) {
	return _FileSys.Contract.SubOrder(&_FileSys.TransactOpts, ps, _gi)
}

// SubOrder is a paid mutator transaction binding the contract method 0x248d02a0.
//
// Solidity: function subOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) ps, uint64 _gi) returns()
func (_FileSys *FileSysTransactorSession) SubOrder(ps OrderIn, _gi uint64) (*types.Transaction, error) {
	return _FileSys.Contract.SubOrder(&_FileSys.TransactOpts, ps, _gi)
}

// Withdraw is a paid mutator transaction binding the contract method 0x5affa0f3.
//
// Solidity: function withdraw(uint64 _i, uint8 _ti, uint256 amount, uint256 _lock) returns(uint256)
func (_FileSys *FileSysTransactor) Withdraw(opts *bind.TransactOpts, _i uint64, _ti uint8, amount *big.Int, _lock *big.Int) (*types.Transaction, error) {
	return _FileSys.contract.Transact(opts, "withdraw", _i, _ti, amount, _lock)
}

// Withdraw is a paid mutator transaction binding the contract method 0x5affa0f3.
//
// Solidity: function withdraw(uint64 _i, uint8 _ti, uint256 amount, uint256 _lock) returns(uint256)
func (_FileSys *FileSysSession) Withdraw(_i uint64, _ti uint8, amount *big.Int, _lock *big.Int) (*types.Transaction, error) {
	return _FileSys.Contract.Withdraw(&_FileSys.TransactOpts, _i, _ti, amount, _lock)
}

// Withdraw is a paid mutator transaction binding the contract method 0x5affa0f3.
//
// Solidity: function withdraw(uint64 _i, uint8 _ti, uint256 amount, uint256 _lock) returns(uint256)
func (_FileSys *FileSysTransactorSession) Withdraw(_i uint64, _ti uint8, amount *big.Int, _lock *big.Int) (*types.Transaction, error) {
	return _FileSys.Contract.Withdraw(&_FileSys.TransactOpts, _i, _ti, amount, _lock)
}

// FileSysAddOrderIterator is returned from FilterAddOrder and is used to iterate over the raw logs and unpacked data for AddOrder events raised by the FileSys contract.
type FileSysAddOrderIterator struct {
	Event *FileSysAddOrder // Event containing the contract specifics and raw log

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
func (it *FileSysAddOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FileSysAddOrder)
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
		it.Event = new(FileSysAddOrder)
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
func (it *FileSysAddOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FileSysAddOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FileSysAddOrder represents a AddOrder event raised by the FileSys contract.
type FileSysAddOrder struct {
	Ti     uint8
	Ui     uint64
	Pi     uint64
	Start  uint64
	End    uint64
	Size   uint64
	Sprice *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAddOrder is a free log retrieval operation binding the contract event 0x9ad187def591246eb8fff84534926005f62988636593745516c0ac08fa086314.
//
// Solidity: event AddOrder(uint8 indexed ti, uint64 indexed ui, uint64 indexed pi, uint64 start, uint64 end, uint64 size, uint256 sprice)
func (_FileSys *FileSysFilterer) FilterAddOrder(opts *bind.FilterOpts, ti []uint8, ui []uint64, pi []uint64) (*FileSysAddOrderIterator, error) {

	var tiRule []interface{}
	for _, tiItem := range ti {
		tiRule = append(tiRule, tiItem)
	}
	var uiRule []interface{}
	for _, uiItem := range ui {
		uiRule = append(uiRule, uiItem)
	}
	var piRule []interface{}
	for _, piItem := range pi {
		piRule = append(piRule, piItem)
	}

	logs, sub, err := _FileSys.contract.FilterLogs(opts, "AddOrder", tiRule, uiRule, piRule)
	if err != nil {
		return nil, err
	}
	return &FileSysAddOrderIterator{contract: _FileSys.contract, event: "AddOrder", logs: logs, sub: sub}, nil
}

// WatchAddOrder is a free log subscription operation binding the contract event 0x9ad187def591246eb8fff84534926005f62988636593745516c0ac08fa086314.
//
// Solidity: event AddOrder(uint8 indexed ti, uint64 indexed ui, uint64 indexed pi, uint64 start, uint64 end, uint64 size, uint256 sprice)
func (_FileSys *FileSysFilterer) WatchAddOrder(opts *bind.WatchOpts, sink chan<- *FileSysAddOrder, ti []uint8, ui []uint64, pi []uint64) (event.Subscription, error) {

	var tiRule []interface{}
	for _, tiItem := range ti {
		tiRule = append(tiRule, tiItem)
	}
	var uiRule []interface{}
	for _, uiItem := range ui {
		uiRule = append(uiRule, uiItem)
	}
	var piRule []interface{}
	for _, piItem := range pi {
		piRule = append(piRule, piItem)
	}

	logs, sub, err := _FileSys.contract.WatchLogs(opts, "AddOrder", tiRule, uiRule, piRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FileSysAddOrder)
				if err := _FileSys.contract.UnpackLog(event, "AddOrder", log); err != nil {
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

// ParseAddOrder is a log parse operation binding the contract event 0x9ad187def591246eb8fff84534926005f62988636593745516c0ac08fa086314.
//
// Solidity: event AddOrder(uint8 indexed ti, uint64 indexed ui, uint64 indexed pi, uint64 start, uint64 end, uint64 size, uint256 sprice)
func (_FileSys *FileSysFilterer) ParseAddOrder(log types.Log) (*FileSysAddOrder, error) {
	event := new(FileSysAddOrder)
	if err := _FileSys.contract.UnpackLog(event, "AddOrder", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FileSysAddRepairIterator is returned from FilterAddRepair and is used to iterate over the raw logs and unpacked data for AddRepair events raised by the FileSys contract.
type FileSysAddRepairIterator struct {
	Event *FileSysAddRepair // Event containing the contract specifics and raw log

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
func (it *FileSysAddRepairIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FileSysAddRepair)
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
		it.Event = new(FileSysAddRepair)
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
func (it *FileSysAddRepairIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FileSysAddRepairIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FileSysAddRepair represents a AddRepair event raised by the FileSys contract.
type FileSysAddRepair struct {
	Ti     uint8
	Ui     uint64
	Pi     uint64
	Start  uint64
	End    uint64
	Size   uint64
	Sprice *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAddRepair is a free log retrieval operation binding the contract event 0x82f944cf0134663ff19eea526226f1cd8d5746a71cc837983659b8cbb2674e48.
//
// Solidity: event AddRepair(uint8 indexed ti, uint64 indexed ui, uint64 indexed pi, uint64 start, uint64 end, uint64 size, uint256 sprice)
func (_FileSys *FileSysFilterer) FilterAddRepair(opts *bind.FilterOpts, ti []uint8, ui []uint64, pi []uint64) (*FileSysAddRepairIterator, error) {

	var tiRule []interface{}
	for _, tiItem := range ti {
		tiRule = append(tiRule, tiItem)
	}
	var uiRule []interface{}
	for _, uiItem := range ui {
		uiRule = append(uiRule, uiItem)
	}
	var piRule []interface{}
	for _, piItem := range pi {
		piRule = append(piRule, piItem)
	}

	logs, sub, err := _FileSys.contract.FilterLogs(opts, "AddRepair", tiRule, uiRule, piRule)
	if err != nil {
		return nil, err
	}
	return &FileSysAddRepairIterator{contract: _FileSys.contract, event: "AddRepair", logs: logs, sub: sub}, nil
}

// WatchAddRepair is a free log subscription operation binding the contract event 0x82f944cf0134663ff19eea526226f1cd8d5746a71cc837983659b8cbb2674e48.
//
// Solidity: event AddRepair(uint8 indexed ti, uint64 indexed ui, uint64 indexed pi, uint64 start, uint64 end, uint64 size, uint256 sprice)
func (_FileSys *FileSysFilterer) WatchAddRepair(opts *bind.WatchOpts, sink chan<- *FileSysAddRepair, ti []uint8, ui []uint64, pi []uint64) (event.Subscription, error) {

	var tiRule []interface{}
	for _, tiItem := range ti {
		tiRule = append(tiRule, tiItem)
	}
	var uiRule []interface{}
	for _, uiItem := range ui {
		uiRule = append(uiRule, uiItem)
	}
	var piRule []interface{}
	for _, piItem := range pi {
		piRule = append(piRule, piItem)
	}

	logs, sub, err := _FileSys.contract.WatchLogs(opts, "AddRepair", tiRule, uiRule, piRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FileSysAddRepair)
				if err := _FileSys.contract.UnpackLog(event, "AddRepair", log); err != nil {
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

// ParseAddRepair is a log parse operation binding the contract event 0x82f944cf0134663ff19eea526226f1cd8d5746a71cc837983659b8cbb2674e48.
//
// Solidity: event AddRepair(uint8 indexed ti, uint64 indexed ui, uint64 indexed pi, uint64 start, uint64 end, uint64 size, uint256 sprice)
func (_FileSys *FileSysFilterer) ParseAddRepair(log types.Log) (*FileSysAddRepair, error) {
	event := new(FileSysAddRepair)
	if err := _FileSys.contract.UnpackLog(event, "AddRepair", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FileSysSetRoleIterator is returned from FilterSetRole and is used to iterate over the raw logs and unpacked data for SetRole events raised by the FileSys contract.
type FileSysSetRoleIterator struct {
	Event *FileSysSetRole // Event containing the contract specifics and raw log

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
func (it *FileSysSetRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FileSysSetRole)
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
		it.Event = new(FileSysSetRole)
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
func (it *FileSysSetRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FileSysSetRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FileSysSetRole represents a SetRole event raised by the FileSys contract.
type FileSysSetRole struct {
	Role uint8
	A    common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSetRole is a free log retrieval operation binding the contract event 0x5c22095cd1123887402a709a9698bc2961c12cb22eeb16e5d76480931f63ede0.
//
// Solidity: event SetRole(uint8 indexed role, address indexed a)
func (_FileSys *FileSysFilterer) FilterSetRole(opts *bind.FilterOpts, role []uint8, a []common.Address) (*FileSysSetRoleIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var aRule []interface{}
	for _, aItem := range a {
		aRule = append(aRule, aItem)
	}

	logs, sub, err := _FileSys.contract.FilterLogs(opts, "SetRole", roleRule, aRule)
	if err != nil {
		return nil, err
	}
	return &FileSysSetRoleIterator{contract: _FileSys.contract, event: "SetRole", logs: logs, sub: sub}, nil
}

// WatchSetRole is a free log subscription operation binding the contract event 0x5c22095cd1123887402a709a9698bc2961c12cb22eeb16e5d76480931f63ede0.
//
// Solidity: event SetRole(uint8 indexed role, address indexed a)
func (_FileSys *FileSysFilterer) WatchSetRole(opts *bind.WatchOpts, sink chan<- *FileSysSetRole, role []uint8, a []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var aRule []interface{}
	for _, aItem := range a {
		aRule = append(aRule, aItem)
	}

	logs, sub, err := _FileSys.contract.WatchLogs(opts, "SetRole", roleRule, aRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FileSysSetRole)
				if err := _FileSys.contract.UnpackLog(event, "SetRole", log); err != nil {
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
func (_FileSys *FileSysFilterer) ParseSetRole(log types.Log) (*FileSysSetRole, error) {
	event := new(FileSysSetRole)
	if err := _FileSys.contract.UnpackLog(event, "SetRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FileSysSubOrderIterator is returned from FilterSubOrder and is used to iterate over the raw logs and unpacked data for SubOrder events raised by the FileSys contract.
type FileSysSubOrderIterator struct {
	Event *FileSysSubOrder // Event containing the contract specifics and raw log

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
func (it *FileSysSubOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FileSysSubOrder)
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
		it.Event = new(FileSysSubOrder)
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
func (it *FileSysSubOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FileSysSubOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FileSysSubOrder represents a SubOrder event raised by the FileSys contract.
type FileSysSubOrder struct {
	Ti     uint8
	Ui     uint64
	Pi     uint64
	Start  uint64
	End    uint64
	Size   uint64
	Sprice *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSubOrder is a free log retrieval operation binding the contract event 0xe42ef64dc7dfe9121f13026a70df328cf6687348fdf8f71745f8e40fa5e52b82.
//
// Solidity: event SubOrder(uint8 indexed ti, uint64 indexed ui, uint64 indexed pi, uint64 start, uint64 end, uint64 size, uint256 sprice)
func (_FileSys *FileSysFilterer) FilterSubOrder(opts *bind.FilterOpts, ti []uint8, ui []uint64, pi []uint64) (*FileSysSubOrderIterator, error) {

	var tiRule []interface{}
	for _, tiItem := range ti {
		tiRule = append(tiRule, tiItem)
	}
	var uiRule []interface{}
	for _, uiItem := range ui {
		uiRule = append(uiRule, uiItem)
	}
	var piRule []interface{}
	for _, piItem := range pi {
		piRule = append(piRule, piItem)
	}

	logs, sub, err := _FileSys.contract.FilterLogs(opts, "SubOrder", tiRule, uiRule, piRule)
	if err != nil {
		return nil, err
	}
	return &FileSysSubOrderIterator{contract: _FileSys.contract, event: "SubOrder", logs: logs, sub: sub}, nil
}

// WatchSubOrder is a free log subscription operation binding the contract event 0xe42ef64dc7dfe9121f13026a70df328cf6687348fdf8f71745f8e40fa5e52b82.
//
// Solidity: event SubOrder(uint8 indexed ti, uint64 indexed ui, uint64 indexed pi, uint64 start, uint64 end, uint64 size, uint256 sprice)
func (_FileSys *FileSysFilterer) WatchSubOrder(opts *bind.WatchOpts, sink chan<- *FileSysSubOrder, ti []uint8, ui []uint64, pi []uint64) (event.Subscription, error) {

	var tiRule []interface{}
	for _, tiItem := range ti {
		tiRule = append(tiRule, tiItem)
	}
	var uiRule []interface{}
	for _, uiItem := range ui {
		uiRule = append(uiRule, uiItem)
	}
	var piRule []interface{}
	for _, piItem := range pi {
		piRule = append(piRule, piItem)
	}

	logs, sub, err := _FileSys.contract.WatchLogs(opts, "SubOrder", tiRule, uiRule, piRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FileSysSubOrder)
				if err := _FileSys.contract.UnpackLog(event, "SubOrder", log); err != nil {
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

// ParseSubOrder is a log parse operation binding the contract event 0xe42ef64dc7dfe9121f13026a70df328cf6687348fdf8f71745f8e40fa5e52b82.
//
// Solidity: event SubOrder(uint8 indexed ti, uint64 indexed ui, uint64 indexed pi, uint64 start, uint64 end, uint64 size, uint256 sprice)
func (_FileSys *FileSysFilterer) ParseSubOrder(log types.Log) (*FileSysSubOrder, error) {
	event := new(FileSysSubOrder)
	if err := _FileSys.contract.UnpackLog(event, "SubOrder", log); err != nil {
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

// IFileSysABI is the input ABI used to generate the binding from.
const IFileSysABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"ti\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"ui\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"pi\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"start\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"end\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sprice\",\"type\":\"uint256\"}],\"name\":\"AddOrder\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"ti\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"ui\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"pi\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"start\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"end\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sprice\",\"type\":\"uint256\"}],\"name\":\"AddRepair\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"ti\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"ui\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"pi\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"start\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"end\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sprice\",\"type\":\"uint256\"}],\"name\":\"SubOrder\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"uIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"start\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"end\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"tIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"sprice\",\"type\":\"uint256\"}],\"internalType\":\"structOrderIn\",\"name\":\"ps\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"_mr\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_tr\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"}],\"name\":\"addOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_rtype\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_kPB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_pPB\",\"type\":\"uint256\"}],\"name\":\"addPenalty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"uIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"start\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"end\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"tIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"sprice\",\"type\":\"uint256\"}],\"internalType\":\"structOrderIn\",\"name\":\"ps\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"}],\"name\":\"addRepair\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_ui\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_pi\",\"type\":\"uint64\"}],\"name\":\"getFsInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"subNonce\",\"type\":\"uint64\"}],\"internalType\":\"structFsOut\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"}],\"name\":\"getGInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"sprice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lost\",\"type\":\"uint256\"}],\"internalType\":\"structGroupOut\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_pi\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"}],\"name\":\"getSettleInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"time\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"hasPaid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"canPay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lost\",\"type\":\"uint256\"}],\"internalType\":\"structSettleOut\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_ui\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_pi\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"}],\"name\":\"getStoreInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"start\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"end\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"sprice\",\"type\":\"uint256\"}],\"internalType\":\"structStoreOut\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"pIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"tIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"pay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lost\",\"type\":\"uint256\"}],\"internalType\":\"structPWIn\",\"name\":\"ps\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"}],\"name\":\"proWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isLock\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"money\",\"type\":\"uint256\"}],\"name\":\"recharge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"uIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"start\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"end\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"tIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"sprice\",\"type\":\"uint256\"}],\"internalType\":\"structOrderIn\",\"name\":\"ps\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"}],\"name\":\"subOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"money\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_lock\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IFileSysFuncSigs maps the 4-byte function signature to its string representation.
var IFileSysFuncSigs = map[string]string{
	"9058d8e7": "addOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256),uint8,uint8,uint64)",
	"2449a168": "addPenalty(uint64,uint64,uint8,uint8,uint256,uint256)",
	"80faaf88": "addRepair((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256),uint64)",
	"fc3ba0ad": "balanceOf(uint64,uint8)",
	"59f7b572": "getFsInfo(uint64,uint64)",
	"5475cf8d": "getGInfo(uint64,uint8)",
	"64c64a4d": "getSettleInfo(uint64,uint8)",
	"8f328ecd": "getStoreInfo(uint64,uint64,uint8)",
	"a4703e16": "proWithdraw((uint64,uint8,uint256,uint256),uint64)",
	"24d11d40": "recharge(uint64,uint8,bool,uint256)",
	"248d02a0": "subOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256),uint64)",
	"5affa0f3": "withdraw(uint64,uint8,uint256,uint256)",
}

// IFileSys is an auto generated Go binding around an Ethereum contract.
type IFileSys struct {
	IFileSysCaller     // Read-only binding to the contract
	IFileSysTransactor // Write-only binding to the contract
	IFileSysFilterer   // Log filterer for contract events
}

// IFileSysCaller is an auto generated read-only Go binding around an Ethereum contract.
type IFileSysCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFileSysTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IFileSysTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFileSysFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IFileSysFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFileSysSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IFileSysSession struct {
	Contract     *IFileSys         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IFileSysCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IFileSysCallerSession struct {
	Contract *IFileSysCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// IFileSysTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IFileSysTransactorSession struct {
	Contract     *IFileSysTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IFileSysRaw is an auto generated low-level Go binding around an Ethereum contract.
type IFileSysRaw struct {
	Contract *IFileSys // Generic contract binding to access the raw methods on
}

// IFileSysCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IFileSysCallerRaw struct {
	Contract *IFileSysCaller // Generic read-only contract binding to access the raw methods on
}

// IFileSysTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IFileSysTransactorRaw struct {
	Contract *IFileSysTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIFileSys creates a new instance of IFileSys, bound to a specific deployed contract.
func NewIFileSys(address common.Address, backend bind.ContractBackend) (*IFileSys, error) {
	contract, err := bindIFileSys(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IFileSys{IFileSysCaller: IFileSysCaller{contract: contract}, IFileSysTransactor: IFileSysTransactor{contract: contract}, IFileSysFilterer: IFileSysFilterer{contract: contract}}, nil
}

// NewIFileSysCaller creates a new read-only instance of IFileSys, bound to a specific deployed contract.
func NewIFileSysCaller(address common.Address, caller bind.ContractCaller) (*IFileSysCaller, error) {
	contract, err := bindIFileSys(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IFileSysCaller{contract: contract}, nil
}

// NewIFileSysTransactor creates a new write-only instance of IFileSys, bound to a specific deployed contract.
func NewIFileSysTransactor(address common.Address, transactor bind.ContractTransactor) (*IFileSysTransactor, error) {
	contract, err := bindIFileSys(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IFileSysTransactor{contract: contract}, nil
}

// NewIFileSysFilterer creates a new log filterer instance of IFileSys, bound to a specific deployed contract.
func NewIFileSysFilterer(address common.Address, filterer bind.ContractFilterer) (*IFileSysFilterer, error) {
	contract, err := bindIFileSys(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IFileSysFilterer{contract: contract}, nil
}

// bindIFileSys binds a generic wrapper to an already deployed contract.
func bindIFileSys(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IFileSysABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IFileSys *IFileSysRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IFileSys.Contract.IFileSysCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IFileSys *IFileSysRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IFileSys.Contract.IFileSysTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IFileSys *IFileSysRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IFileSys.Contract.IFileSysTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IFileSys *IFileSysCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IFileSys.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IFileSys *IFileSysTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IFileSys.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IFileSys *IFileSysTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IFileSys.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0xfc3ba0ad.
//
// Solidity: function balanceOf(uint64 _i, uint8 _ti) view returns(uint256, uint256, uint256)
func (_IFileSys *IFileSysCaller) BalanceOf(opts *bind.CallOpts, _i uint64, _ti uint8) (*big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _IFileSys.contract.Call(opts, &out, "balanceOf", _i, _ti)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// BalanceOf is a free data retrieval call binding the contract method 0xfc3ba0ad.
//
// Solidity: function balanceOf(uint64 _i, uint8 _ti) view returns(uint256, uint256, uint256)
func (_IFileSys *IFileSysSession) BalanceOf(_i uint64, _ti uint8) (*big.Int, *big.Int, *big.Int, error) {
	return _IFileSys.Contract.BalanceOf(&_IFileSys.CallOpts, _i, _ti)
}

// BalanceOf is a free data retrieval call binding the contract method 0xfc3ba0ad.
//
// Solidity: function balanceOf(uint64 _i, uint8 _ti) view returns(uint256, uint256, uint256)
func (_IFileSys *IFileSysCallerSession) BalanceOf(_i uint64, _ti uint8) (*big.Int, *big.Int, *big.Int, error) {
	return _IFileSys.Contract.BalanceOf(&_IFileSys.CallOpts, _i, _ti)
}

// GetFsInfo is a free data retrieval call binding the contract method 0x59f7b572.
//
// Solidity: function getFsInfo(uint64 _ui, uint64 _pi) view returns((uint64,uint64))
func (_IFileSys *IFileSysCaller) GetFsInfo(opts *bind.CallOpts, _ui uint64, _pi uint64) (FsOut, error) {
	var out []interface{}
	err := _IFileSys.contract.Call(opts, &out, "getFsInfo", _ui, _pi)

	if err != nil {
		return *new(FsOut), err
	}

	out0 := *abi.ConvertType(out[0], new(FsOut)).(*FsOut)

	return out0, err

}

// GetFsInfo is a free data retrieval call binding the contract method 0x59f7b572.
//
// Solidity: function getFsInfo(uint64 _ui, uint64 _pi) view returns((uint64,uint64))
func (_IFileSys *IFileSysSession) GetFsInfo(_ui uint64, _pi uint64) (FsOut, error) {
	return _IFileSys.Contract.GetFsInfo(&_IFileSys.CallOpts, _ui, _pi)
}

// GetFsInfo is a free data retrieval call binding the contract method 0x59f7b572.
//
// Solidity: function getFsInfo(uint64 _ui, uint64 _pi) view returns((uint64,uint64))
func (_IFileSys *IFileSysCallerSession) GetFsInfo(_ui uint64, _pi uint64) (FsOut, error) {
	return _IFileSys.Contract.GetFsInfo(&_IFileSys.CallOpts, _ui, _pi)
}

// GetGInfo is a free data retrieval call binding the contract method 0x5475cf8d.
//
// Solidity: function getGInfo(uint64 _gi, uint8 _ti) view returns((uint64,uint256,uint256))
func (_IFileSys *IFileSysCaller) GetGInfo(opts *bind.CallOpts, _gi uint64, _ti uint8) (GroupOut, error) {
	var out []interface{}
	err := _IFileSys.contract.Call(opts, &out, "getGInfo", _gi, _ti)

	if err != nil {
		return *new(GroupOut), err
	}

	out0 := *abi.ConvertType(out[0], new(GroupOut)).(*GroupOut)

	return out0, err

}

// GetGInfo is a free data retrieval call binding the contract method 0x5475cf8d.
//
// Solidity: function getGInfo(uint64 _gi, uint8 _ti) view returns((uint64,uint256,uint256))
func (_IFileSys *IFileSysSession) GetGInfo(_gi uint64, _ti uint8) (GroupOut, error) {
	return _IFileSys.Contract.GetGInfo(&_IFileSys.CallOpts, _gi, _ti)
}

// GetGInfo is a free data retrieval call binding the contract method 0x5475cf8d.
//
// Solidity: function getGInfo(uint64 _gi, uint8 _ti) view returns((uint64,uint256,uint256))
func (_IFileSys *IFileSysCallerSession) GetGInfo(_gi uint64, _ti uint8) (GroupOut, error) {
	return _IFileSys.Contract.GetGInfo(&_IFileSys.CallOpts, _gi, _ti)
}

// GetSettleInfo is a free data retrieval call binding the contract method 0x64c64a4d.
//
// Solidity: function getSettleInfo(uint64 _pi, uint8 _ti) view returns((uint64,uint64,uint256,uint256,uint256,uint256,uint256))
func (_IFileSys *IFileSysCaller) GetSettleInfo(opts *bind.CallOpts, _pi uint64, _ti uint8) (SettleOut, error) {
	var out []interface{}
	err := _IFileSys.contract.Call(opts, &out, "getSettleInfo", _pi, _ti)

	if err != nil {
		return *new(SettleOut), err
	}

	out0 := *abi.ConvertType(out[0], new(SettleOut)).(*SettleOut)

	return out0, err

}

// GetSettleInfo is a free data retrieval call binding the contract method 0x64c64a4d.
//
// Solidity: function getSettleInfo(uint64 _pi, uint8 _ti) view returns((uint64,uint64,uint256,uint256,uint256,uint256,uint256))
func (_IFileSys *IFileSysSession) GetSettleInfo(_pi uint64, _ti uint8) (SettleOut, error) {
	return _IFileSys.Contract.GetSettleInfo(&_IFileSys.CallOpts, _pi, _ti)
}

// GetSettleInfo is a free data retrieval call binding the contract method 0x64c64a4d.
//
// Solidity: function getSettleInfo(uint64 _pi, uint8 _ti) view returns((uint64,uint64,uint256,uint256,uint256,uint256,uint256))
func (_IFileSys *IFileSysCallerSession) GetSettleInfo(_pi uint64, _ti uint8) (SettleOut, error) {
	return _IFileSys.Contract.GetSettleInfo(&_IFileSys.CallOpts, _pi, _ti)
}

// GetStoreInfo is a free data retrieval call binding the contract method 0x8f328ecd.
//
// Solidity: function getStoreInfo(uint64 _ui, uint64 _pi, uint8 _ti) view returns((uint64,uint64,uint64,uint256))
func (_IFileSys *IFileSysCaller) GetStoreInfo(opts *bind.CallOpts, _ui uint64, _pi uint64, _ti uint8) (StoreOut, error) {
	var out []interface{}
	err := _IFileSys.contract.Call(opts, &out, "getStoreInfo", _ui, _pi, _ti)

	if err != nil {
		return *new(StoreOut), err
	}

	out0 := *abi.ConvertType(out[0], new(StoreOut)).(*StoreOut)

	return out0, err

}

// GetStoreInfo is a free data retrieval call binding the contract method 0x8f328ecd.
//
// Solidity: function getStoreInfo(uint64 _ui, uint64 _pi, uint8 _ti) view returns((uint64,uint64,uint64,uint256))
func (_IFileSys *IFileSysSession) GetStoreInfo(_ui uint64, _pi uint64, _ti uint8) (StoreOut, error) {
	return _IFileSys.Contract.GetStoreInfo(&_IFileSys.CallOpts, _ui, _pi, _ti)
}

// GetStoreInfo is a free data retrieval call binding the contract method 0x8f328ecd.
//
// Solidity: function getStoreInfo(uint64 _ui, uint64 _pi, uint8 _ti) view returns((uint64,uint64,uint64,uint256))
func (_IFileSys *IFileSysCallerSession) GetStoreInfo(_ui uint64, _pi uint64, _ti uint8) (StoreOut, error) {
	return _IFileSys.Contract.GetStoreInfo(&_IFileSys.CallOpts, _ui, _pi, _ti)
}

// AddOrder is a paid mutator transaction binding the contract method 0x9058d8e7.
//
// Solidity: function addOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) ps, uint8 _mr, uint8 _tr, uint64 _gi) returns()
func (_IFileSys *IFileSysTransactor) AddOrder(opts *bind.TransactOpts, ps OrderIn, _mr uint8, _tr uint8, _gi uint64) (*types.Transaction, error) {
	return _IFileSys.contract.Transact(opts, "addOrder", ps, _mr, _tr, _gi)
}

// AddOrder is a paid mutator transaction binding the contract method 0x9058d8e7.
//
// Solidity: function addOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) ps, uint8 _mr, uint8 _tr, uint64 _gi) returns()
func (_IFileSys *IFileSysSession) AddOrder(ps OrderIn, _mr uint8, _tr uint8, _gi uint64) (*types.Transaction, error) {
	return _IFileSys.Contract.AddOrder(&_IFileSys.TransactOpts, ps, _mr, _tr, _gi)
}

// AddOrder is a paid mutator transaction binding the contract method 0x9058d8e7.
//
// Solidity: function addOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) ps, uint8 _mr, uint8 _tr, uint64 _gi) returns()
func (_IFileSys *IFileSysTransactorSession) AddOrder(ps OrderIn, _mr uint8, _tr uint8, _gi uint64) (*types.Transaction, error) {
	return _IFileSys.Contract.AddOrder(&_IFileSys.TransactOpts, ps, _mr, _tr, _gi)
}

// AddPenalty is a paid mutator transaction binding the contract method 0x2449a168.
//
// Solidity: function addPenalty(uint64 _gi, uint64 _i, uint8 _ti, uint8 _rtype, uint256 _kPB, uint256 _pPB) returns()
func (_IFileSys *IFileSysTransactor) AddPenalty(opts *bind.TransactOpts, _gi uint64, _i uint64, _ti uint8, _rtype uint8, _kPB *big.Int, _pPB *big.Int) (*types.Transaction, error) {
	return _IFileSys.contract.Transact(opts, "addPenalty", _gi, _i, _ti, _rtype, _kPB, _pPB)
}

// AddPenalty is a paid mutator transaction binding the contract method 0x2449a168.
//
// Solidity: function addPenalty(uint64 _gi, uint64 _i, uint8 _ti, uint8 _rtype, uint256 _kPB, uint256 _pPB) returns()
func (_IFileSys *IFileSysSession) AddPenalty(_gi uint64, _i uint64, _ti uint8, _rtype uint8, _kPB *big.Int, _pPB *big.Int) (*types.Transaction, error) {
	return _IFileSys.Contract.AddPenalty(&_IFileSys.TransactOpts, _gi, _i, _ti, _rtype, _kPB, _pPB)
}

// AddPenalty is a paid mutator transaction binding the contract method 0x2449a168.
//
// Solidity: function addPenalty(uint64 _gi, uint64 _i, uint8 _ti, uint8 _rtype, uint256 _kPB, uint256 _pPB) returns()
func (_IFileSys *IFileSysTransactorSession) AddPenalty(_gi uint64, _i uint64, _ti uint8, _rtype uint8, _kPB *big.Int, _pPB *big.Int) (*types.Transaction, error) {
	return _IFileSys.Contract.AddPenalty(&_IFileSys.TransactOpts, _gi, _i, _ti, _rtype, _kPB, _pPB)
}

// AddRepair is a paid mutator transaction binding the contract method 0x80faaf88.
//
// Solidity: function addRepair((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) ps, uint64 _gi) returns()
func (_IFileSys *IFileSysTransactor) AddRepair(opts *bind.TransactOpts, ps OrderIn, _gi uint64) (*types.Transaction, error) {
	return _IFileSys.contract.Transact(opts, "addRepair", ps, _gi)
}

// AddRepair is a paid mutator transaction binding the contract method 0x80faaf88.
//
// Solidity: function addRepair((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) ps, uint64 _gi) returns()
func (_IFileSys *IFileSysSession) AddRepair(ps OrderIn, _gi uint64) (*types.Transaction, error) {
	return _IFileSys.Contract.AddRepair(&_IFileSys.TransactOpts, ps, _gi)
}

// AddRepair is a paid mutator transaction binding the contract method 0x80faaf88.
//
// Solidity: function addRepair((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) ps, uint64 _gi) returns()
func (_IFileSys *IFileSysTransactorSession) AddRepair(ps OrderIn, _gi uint64) (*types.Transaction, error) {
	return _IFileSys.Contract.AddRepair(&_IFileSys.TransactOpts, ps, _gi)
}

// ProWithdraw is a paid mutator transaction binding the contract method 0xa4703e16.
//
// Solidity: function proWithdraw((uint64,uint8,uint256,uint256) ps, uint64 _gi) returns(uint256)
func (_IFileSys *IFileSysTransactor) ProWithdraw(opts *bind.TransactOpts, ps PWIn, _gi uint64) (*types.Transaction, error) {
	return _IFileSys.contract.Transact(opts, "proWithdraw", ps, _gi)
}

// ProWithdraw is a paid mutator transaction binding the contract method 0xa4703e16.
//
// Solidity: function proWithdraw((uint64,uint8,uint256,uint256) ps, uint64 _gi) returns(uint256)
func (_IFileSys *IFileSysSession) ProWithdraw(ps PWIn, _gi uint64) (*types.Transaction, error) {
	return _IFileSys.Contract.ProWithdraw(&_IFileSys.TransactOpts, ps, _gi)
}

// ProWithdraw is a paid mutator transaction binding the contract method 0xa4703e16.
//
// Solidity: function proWithdraw((uint64,uint8,uint256,uint256) ps, uint64 _gi) returns(uint256)
func (_IFileSys *IFileSysTransactorSession) ProWithdraw(ps PWIn, _gi uint64) (*types.Transaction, error) {
	return _IFileSys.Contract.ProWithdraw(&_IFileSys.TransactOpts, ps, _gi)
}

// Recharge is a paid mutator transaction binding the contract method 0x24d11d40.
//
// Solidity: function recharge(uint64 _i, uint8 _ti, bool isLock, uint256 money) returns()
func (_IFileSys *IFileSysTransactor) Recharge(opts *bind.TransactOpts, _i uint64, _ti uint8, isLock bool, money *big.Int) (*types.Transaction, error) {
	return _IFileSys.contract.Transact(opts, "recharge", _i, _ti, isLock, money)
}

// Recharge is a paid mutator transaction binding the contract method 0x24d11d40.
//
// Solidity: function recharge(uint64 _i, uint8 _ti, bool isLock, uint256 money) returns()
func (_IFileSys *IFileSysSession) Recharge(_i uint64, _ti uint8, isLock bool, money *big.Int) (*types.Transaction, error) {
	return _IFileSys.Contract.Recharge(&_IFileSys.TransactOpts, _i, _ti, isLock, money)
}

// Recharge is a paid mutator transaction binding the contract method 0x24d11d40.
//
// Solidity: function recharge(uint64 _i, uint8 _ti, bool isLock, uint256 money) returns()
func (_IFileSys *IFileSysTransactorSession) Recharge(_i uint64, _ti uint8, isLock bool, money *big.Int) (*types.Transaction, error) {
	return _IFileSys.Contract.Recharge(&_IFileSys.TransactOpts, _i, _ti, isLock, money)
}

// SubOrder is a paid mutator transaction binding the contract method 0x248d02a0.
//
// Solidity: function subOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) ps, uint64 _gi) returns()
func (_IFileSys *IFileSysTransactor) SubOrder(opts *bind.TransactOpts, ps OrderIn, _gi uint64) (*types.Transaction, error) {
	return _IFileSys.contract.Transact(opts, "subOrder", ps, _gi)
}

// SubOrder is a paid mutator transaction binding the contract method 0x248d02a0.
//
// Solidity: function subOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) ps, uint64 _gi) returns()
func (_IFileSys *IFileSysSession) SubOrder(ps OrderIn, _gi uint64) (*types.Transaction, error) {
	return _IFileSys.Contract.SubOrder(&_IFileSys.TransactOpts, ps, _gi)
}

// SubOrder is a paid mutator transaction binding the contract method 0x248d02a0.
//
// Solidity: function subOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) ps, uint64 _gi) returns()
func (_IFileSys *IFileSysTransactorSession) SubOrder(ps OrderIn, _gi uint64) (*types.Transaction, error) {
	return _IFileSys.Contract.SubOrder(&_IFileSys.TransactOpts, ps, _gi)
}

// Withdraw is a paid mutator transaction binding the contract method 0x5affa0f3.
//
// Solidity: function withdraw(uint64 _i, uint8 _ti, uint256 money, uint256 _lock) returns(uint256)
func (_IFileSys *IFileSysTransactor) Withdraw(opts *bind.TransactOpts, _i uint64, _ti uint8, money *big.Int, _lock *big.Int) (*types.Transaction, error) {
	return _IFileSys.contract.Transact(opts, "withdraw", _i, _ti, money, _lock)
}

// Withdraw is a paid mutator transaction binding the contract method 0x5affa0f3.
//
// Solidity: function withdraw(uint64 _i, uint8 _ti, uint256 money, uint256 _lock) returns(uint256)
func (_IFileSys *IFileSysSession) Withdraw(_i uint64, _ti uint8, money *big.Int, _lock *big.Int) (*types.Transaction, error) {
	return _IFileSys.Contract.Withdraw(&_IFileSys.TransactOpts, _i, _ti, money, _lock)
}

// Withdraw is a paid mutator transaction binding the contract method 0x5affa0f3.
//
// Solidity: function withdraw(uint64 _i, uint8 _ti, uint256 money, uint256 _lock) returns(uint256)
func (_IFileSys *IFileSysTransactorSession) Withdraw(_i uint64, _ti uint8, money *big.Int, _lock *big.Int) (*types.Transaction, error) {
	return _IFileSys.Contract.Withdraw(&_IFileSys.TransactOpts, _i, _ti, money, _lock)
}

// IFileSysAddOrderIterator is returned from FilterAddOrder and is used to iterate over the raw logs and unpacked data for AddOrder events raised by the IFileSys contract.
type IFileSysAddOrderIterator struct {
	Event *IFileSysAddOrder // Event containing the contract specifics and raw log

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
func (it *IFileSysAddOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IFileSysAddOrder)
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
		it.Event = new(IFileSysAddOrder)
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
func (it *IFileSysAddOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IFileSysAddOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IFileSysAddOrder represents a AddOrder event raised by the IFileSys contract.
type IFileSysAddOrder struct {
	Ti     uint8
	Ui     uint64
	Pi     uint64
	Start  uint64
	End    uint64
	Size   uint64
	Sprice *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAddOrder is a free log retrieval operation binding the contract event 0x9ad187def591246eb8fff84534926005f62988636593745516c0ac08fa086314.
//
// Solidity: event AddOrder(uint8 indexed ti, uint64 indexed ui, uint64 indexed pi, uint64 start, uint64 end, uint64 size, uint256 sprice)
func (_IFileSys *IFileSysFilterer) FilterAddOrder(opts *bind.FilterOpts, ti []uint8, ui []uint64, pi []uint64) (*IFileSysAddOrderIterator, error) {

	var tiRule []interface{}
	for _, tiItem := range ti {
		tiRule = append(tiRule, tiItem)
	}
	var uiRule []interface{}
	for _, uiItem := range ui {
		uiRule = append(uiRule, uiItem)
	}
	var piRule []interface{}
	for _, piItem := range pi {
		piRule = append(piRule, piItem)
	}

	logs, sub, err := _IFileSys.contract.FilterLogs(opts, "AddOrder", tiRule, uiRule, piRule)
	if err != nil {
		return nil, err
	}
	return &IFileSysAddOrderIterator{contract: _IFileSys.contract, event: "AddOrder", logs: logs, sub: sub}, nil
}

// WatchAddOrder is a free log subscription operation binding the contract event 0x9ad187def591246eb8fff84534926005f62988636593745516c0ac08fa086314.
//
// Solidity: event AddOrder(uint8 indexed ti, uint64 indexed ui, uint64 indexed pi, uint64 start, uint64 end, uint64 size, uint256 sprice)
func (_IFileSys *IFileSysFilterer) WatchAddOrder(opts *bind.WatchOpts, sink chan<- *IFileSysAddOrder, ti []uint8, ui []uint64, pi []uint64) (event.Subscription, error) {

	var tiRule []interface{}
	for _, tiItem := range ti {
		tiRule = append(tiRule, tiItem)
	}
	var uiRule []interface{}
	for _, uiItem := range ui {
		uiRule = append(uiRule, uiItem)
	}
	var piRule []interface{}
	for _, piItem := range pi {
		piRule = append(piRule, piItem)
	}

	logs, sub, err := _IFileSys.contract.WatchLogs(opts, "AddOrder", tiRule, uiRule, piRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IFileSysAddOrder)
				if err := _IFileSys.contract.UnpackLog(event, "AddOrder", log); err != nil {
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

// ParseAddOrder is a log parse operation binding the contract event 0x9ad187def591246eb8fff84534926005f62988636593745516c0ac08fa086314.
//
// Solidity: event AddOrder(uint8 indexed ti, uint64 indexed ui, uint64 indexed pi, uint64 start, uint64 end, uint64 size, uint256 sprice)
func (_IFileSys *IFileSysFilterer) ParseAddOrder(log types.Log) (*IFileSysAddOrder, error) {
	event := new(IFileSysAddOrder)
	if err := _IFileSys.contract.UnpackLog(event, "AddOrder", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IFileSysAddRepairIterator is returned from FilterAddRepair and is used to iterate over the raw logs and unpacked data for AddRepair events raised by the IFileSys contract.
type IFileSysAddRepairIterator struct {
	Event *IFileSysAddRepair // Event containing the contract specifics and raw log

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
func (it *IFileSysAddRepairIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IFileSysAddRepair)
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
		it.Event = new(IFileSysAddRepair)
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
func (it *IFileSysAddRepairIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IFileSysAddRepairIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IFileSysAddRepair represents a AddRepair event raised by the IFileSys contract.
type IFileSysAddRepair struct {
	Ti     uint8
	Ui     uint64
	Pi     uint64
	Start  uint64
	End    uint64
	Size   uint64
	Sprice *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAddRepair is a free log retrieval operation binding the contract event 0x82f944cf0134663ff19eea526226f1cd8d5746a71cc837983659b8cbb2674e48.
//
// Solidity: event AddRepair(uint8 indexed ti, uint64 indexed ui, uint64 indexed pi, uint64 start, uint64 end, uint64 size, uint256 sprice)
func (_IFileSys *IFileSysFilterer) FilterAddRepair(opts *bind.FilterOpts, ti []uint8, ui []uint64, pi []uint64) (*IFileSysAddRepairIterator, error) {

	var tiRule []interface{}
	for _, tiItem := range ti {
		tiRule = append(tiRule, tiItem)
	}
	var uiRule []interface{}
	for _, uiItem := range ui {
		uiRule = append(uiRule, uiItem)
	}
	var piRule []interface{}
	for _, piItem := range pi {
		piRule = append(piRule, piItem)
	}

	logs, sub, err := _IFileSys.contract.FilterLogs(opts, "AddRepair", tiRule, uiRule, piRule)
	if err != nil {
		return nil, err
	}
	return &IFileSysAddRepairIterator{contract: _IFileSys.contract, event: "AddRepair", logs: logs, sub: sub}, nil
}

// WatchAddRepair is a free log subscription operation binding the contract event 0x82f944cf0134663ff19eea526226f1cd8d5746a71cc837983659b8cbb2674e48.
//
// Solidity: event AddRepair(uint8 indexed ti, uint64 indexed ui, uint64 indexed pi, uint64 start, uint64 end, uint64 size, uint256 sprice)
func (_IFileSys *IFileSysFilterer) WatchAddRepair(opts *bind.WatchOpts, sink chan<- *IFileSysAddRepair, ti []uint8, ui []uint64, pi []uint64) (event.Subscription, error) {

	var tiRule []interface{}
	for _, tiItem := range ti {
		tiRule = append(tiRule, tiItem)
	}
	var uiRule []interface{}
	for _, uiItem := range ui {
		uiRule = append(uiRule, uiItem)
	}
	var piRule []interface{}
	for _, piItem := range pi {
		piRule = append(piRule, piItem)
	}

	logs, sub, err := _IFileSys.contract.WatchLogs(opts, "AddRepair", tiRule, uiRule, piRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IFileSysAddRepair)
				if err := _IFileSys.contract.UnpackLog(event, "AddRepair", log); err != nil {
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

// ParseAddRepair is a log parse operation binding the contract event 0x82f944cf0134663ff19eea526226f1cd8d5746a71cc837983659b8cbb2674e48.
//
// Solidity: event AddRepair(uint8 indexed ti, uint64 indexed ui, uint64 indexed pi, uint64 start, uint64 end, uint64 size, uint256 sprice)
func (_IFileSys *IFileSysFilterer) ParseAddRepair(log types.Log) (*IFileSysAddRepair, error) {
	event := new(IFileSysAddRepair)
	if err := _IFileSys.contract.UnpackLog(event, "AddRepair", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IFileSysSubOrderIterator is returned from FilterSubOrder and is used to iterate over the raw logs and unpacked data for SubOrder events raised by the IFileSys contract.
type IFileSysSubOrderIterator struct {
	Event *IFileSysSubOrder // Event containing the contract specifics and raw log

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
func (it *IFileSysSubOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IFileSysSubOrder)
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
		it.Event = new(IFileSysSubOrder)
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
func (it *IFileSysSubOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IFileSysSubOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IFileSysSubOrder represents a SubOrder event raised by the IFileSys contract.
type IFileSysSubOrder struct {
	Ti     uint8
	Ui     uint64
	Pi     uint64
	Start  uint64
	End    uint64
	Size   uint64
	Sprice *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSubOrder is a free log retrieval operation binding the contract event 0xe42ef64dc7dfe9121f13026a70df328cf6687348fdf8f71745f8e40fa5e52b82.
//
// Solidity: event SubOrder(uint8 indexed ti, uint64 indexed ui, uint64 indexed pi, uint64 start, uint64 end, uint64 size, uint256 sprice)
func (_IFileSys *IFileSysFilterer) FilterSubOrder(opts *bind.FilterOpts, ti []uint8, ui []uint64, pi []uint64) (*IFileSysSubOrderIterator, error) {

	var tiRule []interface{}
	for _, tiItem := range ti {
		tiRule = append(tiRule, tiItem)
	}
	var uiRule []interface{}
	for _, uiItem := range ui {
		uiRule = append(uiRule, uiItem)
	}
	var piRule []interface{}
	for _, piItem := range pi {
		piRule = append(piRule, piItem)
	}

	logs, sub, err := _IFileSys.contract.FilterLogs(opts, "SubOrder", tiRule, uiRule, piRule)
	if err != nil {
		return nil, err
	}
	return &IFileSysSubOrderIterator{contract: _IFileSys.contract, event: "SubOrder", logs: logs, sub: sub}, nil
}

// WatchSubOrder is a free log subscription operation binding the contract event 0xe42ef64dc7dfe9121f13026a70df328cf6687348fdf8f71745f8e40fa5e52b82.
//
// Solidity: event SubOrder(uint8 indexed ti, uint64 indexed ui, uint64 indexed pi, uint64 start, uint64 end, uint64 size, uint256 sprice)
func (_IFileSys *IFileSysFilterer) WatchSubOrder(opts *bind.WatchOpts, sink chan<- *IFileSysSubOrder, ti []uint8, ui []uint64, pi []uint64) (event.Subscription, error) {

	var tiRule []interface{}
	for _, tiItem := range ti {
		tiRule = append(tiRule, tiItem)
	}
	var uiRule []interface{}
	for _, uiItem := range ui {
		uiRule = append(uiRule, uiItem)
	}
	var piRule []interface{}
	for _, piItem := range pi {
		piRule = append(piRule, piItem)
	}

	logs, sub, err := _IFileSys.contract.WatchLogs(opts, "SubOrder", tiRule, uiRule, piRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IFileSysSubOrder)
				if err := _IFileSys.contract.UnpackLog(event, "SubOrder", log); err != nil {
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

// ParseSubOrder is a log parse operation binding the contract event 0xe42ef64dc7dfe9121f13026a70df328cf6687348fdf8f71745f8e40fa5e52b82.
//
// Solidity: event SubOrder(uint8 indexed ti, uint64 indexed ui, uint64 indexed pi, uint64 start, uint64 end, uint64 size, uint256 sprice)
func (_IFileSys *IFileSysFilterer) ParseSubOrder(log types.Log) (*IFileSysSubOrder, error) {
	event := new(IFileSysSubOrder)
	if err := _IFileSys.contract.UnpackLog(event, "SubOrder", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IFileSysGetterABI is the input ABI used to generate the binding from.
const IFileSysGetterABI = "[{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_ui\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_pi\",\"type\":\"uint64\"}],\"name\":\"getFsInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"subNonce\",\"type\":\"uint64\"}],\"internalType\":\"structFsOut\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"}],\"name\":\"getGInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"sprice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lost\",\"type\":\"uint256\"}],\"internalType\":\"structGroupOut\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_pi\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"}],\"name\":\"getSettleInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"time\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"hasPaid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"canPay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lost\",\"type\":\"uint256\"}],\"internalType\":\"structSettleOut\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_ui\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_pi\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"}],\"name\":\"getStoreInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"start\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"end\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"sprice\",\"type\":\"uint256\"}],\"internalType\":\"structStoreOut\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IFileSysGetterFuncSigs maps the 4-byte function signature to its string representation.
var IFileSysGetterFuncSigs = map[string]string{
	"fc3ba0ad": "balanceOf(uint64,uint8)",
	"59f7b572": "getFsInfo(uint64,uint64)",
	"5475cf8d": "getGInfo(uint64,uint8)",
	"64c64a4d": "getSettleInfo(uint64,uint8)",
	"8f328ecd": "getStoreInfo(uint64,uint64,uint8)",
}

// IFileSysGetter is an auto generated Go binding around an Ethereum contract.
type IFileSysGetter struct {
	IFileSysGetterCaller     // Read-only binding to the contract
	IFileSysGetterTransactor // Write-only binding to the contract
	IFileSysGetterFilterer   // Log filterer for contract events
}

// IFileSysGetterCaller is an auto generated read-only Go binding around an Ethereum contract.
type IFileSysGetterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFileSysGetterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IFileSysGetterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFileSysGetterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IFileSysGetterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFileSysGetterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IFileSysGetterSession struct {
	Contract     *IFileSysGetter   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IFileSysGetterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IFileSysGetterCallerSession struct {
	Contract *IFileSysGetterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IFileSysGetterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IFileSysGetterTransactorSession struct {
	Contract     *IFileSysGetterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IFileSysGetterRaw is an auto generated low-level Go binding around an Ethereum contract.
type IFileSysGetterRaw struct {
	Contract *IFileSysGetter // Generic contract binding to access the raw methods on
}

// IFileSysGetterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IFileSysGetterCallerRaw struct {
	Contract *IFileSysGetterCaller // Generic read-only contract binding to access the raw methods on
}

// IFileSysGetterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IFileSysGetterTransactorRaw struct {
	Contract *IFileSysGetterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIFileSysGetter creates a new instance of IFileSysGetter, bound to a specific deployed contract.
func NewIFileSysGetter(address common.Address, backend bind.ContractBackend) (*IFileSysGetter, error) {
	contract, err := bindIFileSysGetter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IFileSysGetter{IFileSysGetterCaller: IFileSysGetterCaller{contract: contract}, IFileSysGetterTransactor: IFileSysGetterTransactor{contract: contract}, IFileSysGetterFilterer: IFileSysGetterFilterer{contract: contract}}, nil
}

// NewIFileSysGetterCaller creates a new read-only instance of IFileSysGetter, bound to a specific deployed contract.
func NewIFileSysGetterCaller(address common.Address, caller bind.ContractCaller) (*IFileSysGetterCaller, error) {
	contract, err := bindIFileSysGetter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IFileSysGetterCaller{contract: contract}, nil
}

// NewIFileSysGetterTransactor creates a new write-only instance of IFileSysGetter, bound to a specific deployed contract.
func NewIFileSysGetterTransactor(address common.Address, transactor bind.ContractTransactor) (*IFileSysGetterTransactor, error) {
	contract, err := bindIFileSysGetter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IFileSysGetterTransactor{contract: contract}, nil
}

// NewIFileSysGetterFilterer creates a new log filterer instance of IFileSysGetter, bound to a specific deployed contract.
func NewIFileSysGetterFilterer(address common.Address, filterer bind.ContractFilterer) (*IFileSysGetterFilterer, error) {
	contract, err := bindIFileSysGetter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IFileSysGetterFilterer{contract: contract}, nil
}

// bindIFileSysGetter binds a generic wrapper to an already deployed contract.
func bindIFileSysGetter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IFileSysGetterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IFileSysGetter *IFileSysGetterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IFileSysGetter.Contract.IFileSysGetterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IFileSysGetter *IFileSysGetterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IFileSysGetter.Contract.IFileSysGetterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IFileSysGetter *IFileSysGetterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IFileSysGetter.Contract.IFileSysGetterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IFileSysGetter *IFileSysGetterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IFileSysGetter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IFileSysGetter *IFileSysGetterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IFileSysGetter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IFileSysGetter *IFileSysGetterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IFileSysGetter.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0xfc3ba0ad.
//
// Solidity: function balanceOf(uint64 _i, uint8 _ti) view returns(uint256, uint256, uint256)
func (_IFileSysGetter *IFileSysGetterCaller) BalanceOf(opts *bind.CallOpts, _i uint64, _ti uint8) (*big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _IFileSysGetter.contract.Call(opts, &out, "balanceOf", _i, _ti)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// BalanceOf is a free data retrieval call binding the contract method 0xfc3ba0ad.
//
// Solidity: function balanceOf(uint64 _i, uint8 _ti) view returns(uint256, uint256, uint256)
func (_IFileSysGetter *IFileSysGetterSession) BalanceOf(_i uint64, _ti uint8) (*big.Int, *big.Int, *big.Int, error) {
	return _IFileSysGetter.Contract.BalanceOf(&_IFileSysGetter.CallOpts, _i, _ti)
}

// BalanceOf is a free data retrieval call binding the contract method 0xfc3ba0ad.
//
// Solidity: function balanceOf(uint64 _i, uint8 _ti) view returns(uint256, uint256, uint256)
func (_IFileSysGetter *IFileSysGetterCallerSession) BalanceOf(_i uint64, _ti uint8) (*big.Int, *big.Int, *big.Int, error) {
	return _IFileSysGetter.Contract.BalanceOf(&_IFileSysGetter.CallOpts, _i, _ti)
}

// GetFsInfo is a free data retrieval call binding the contract method 0x59f7b572.
//
// Solidity: function getFsInfo(uint64 _ui, uint64 _pi) view returns((uint64,uint64))
func (_IFileSysGetter *IFileSysGetterCaller) GetFsInfo(opts *bind.CallOpts, _ui uint64, _pi uint64) (FsOut, error) {
	var out []interface{}
	err := _IFileSysGetter.contract.Call(opts, &out, "getFsInfo", _ui, _pi)

	if err != nil {
		return *new(FsOut), err
	}

	out0 := *abi.ConvertType(out[0], new(FsOut)).(*FsOut)

	return out0, err

}

// GetFsInfo is a free data retrieval call binding the contract method 0x59f7b572.
//
// Solidity: function getFsInfo(uint64 _ui, uint64 _pi) view returns((uint64,uint64))
func (_IFileSysGetter *IFileSysGetterSession) GetFsInfo(_ui uint64, _pi uint64) (FsOut, error) {
	return _IFileSysGetter.Contract.GetFsInfo(&_IFileSysGetter.CallOpts, _ui, _pi)
}

// GetFsInfo is a free data retrieval call binding the contract method 0x59f7b572.
//
// Solidity: function getFsInfo(uint64 _ui, uint64 _pi) view returns((uint64,uint64))
func (_IFileSysGetter *IFileSysGetterCallerSession) GetFsInfo(_ui uint64, _pi uint64) (FsOut, error) {
	return _IFileSysGetter.Contract.GetFsInfo(&_IFileSysGetter.CallOpts, _ui, _pi)
}

// GetGInfo is a free data retrieval call binding the contract method 0x5475cf8d.
//
// Solidity: function getGInfo(uint64 _gi, uint8 _ti) view returns((uint64,uint256,uint256))
func (_IFileSysGetter *IFileSysGetterCaller) GetGInfo(opts *bind.CallOpts, _gi uint64, _ti uint8) (GroupOut, error) {
	var out []interface{}
	err := _IFileSysGetter.contract.Call(opts, &out, "getGInfo", _gi, _ti)

	if err != nil {
		return *new(GroupOut), err
	}

	out0 := *abi.ConvertType(out[0], new(GroupOut)).(*GroupOut)

	return out0, err

}

// GetGInfo is a free data retrieval call binding the contract method 0x5475cf8d.
//
// Solidity: function getGInfo(uint64 _gi, uint8 _ti) view returns((uint64,uint256,uint256))
func (_IFileSysGetter *IFileSysGetterSession) GetGInfo(_gi uint64, _ti uint8) (GroupOut, error) {
	return _IFileSysGetter.Contract.GetGInfo(&_IFileSysGetter.CallOpts, _gi, _ti)
}

// GetGInfo is a free data retrieval call binding the contract method 0x5475cf8d.
//
// Solidity: function getGInfo(uint64 _gi, uint8 _ti) view returns((uint64,uint256,uint256))
func (_IFileSysGetter *IFileSysGetterCallerSession) GetGInfo(_gi uint64, _ti uint8) (GroupOut, error) {
	return _IFileSysGetter.Contract.GetGInfo(&_IFileSysGetter.CallOpts, _gi, _ti)
}

// GetSettleInfo is a free data retrieval call binding the contract method 0x64c64a4d.
//
// Solidity: function getSettleInfo(uint64 _pi, uint8 _ti) view returns((uint64,uint64,uint256,uint256,uint256,uint256,uint256))
func (_IFileSysGetter *IFileSysGetterCaller) GetSettleInfo(opts *bind.CallOpts, _pi uint64, _ti uint8) (SettleOut, error) {
	var out []interface{}
	err := _IFileSysGetter.contract.Call(opts, &out, "getSettleInfo", _pi, _ti)

	if err != nil {
		return *new(SettleOut), err
	}

	out0 := *abi.ConvertType(out[0], new(SettleOut)).(*SettleOut)

	return out0, err

}

// GetSettleInfo is a free data retrieval call binding the contract method 0x64c64a4d.
//
// Solidity: function getSettleInfo(uint64 _pi, uint8 _ti) view returns((uint64,uint64,uint256,uint256,uint256,uint256,uint256))
func (_IFileSysGetter *IFileSysGetterSession) GetSettleInfo(_pi uint64, _ti uint8) (SettleOut, error) {
	return _IFileSysGetter.Contract.GetSettleInfo(&_IFileSysGetter.CallOpts, _pi, _ti)
}

// GetSettleInfo is a free data retrieval call binding the contract method 0x64c64a4d.
//
// Solidity: function getSettleInfo(uint64 _pi, uint8 _ti) view returns((uint64,uint64,uint256,uint256,uint256,uint256,uint256))
func (_IFileSysGetter *IFileSysGetterCallerSession) GetSettleInfo(_pi uint64, _ti uint8) (SettleOut, error) {
	return _IFileSysGetter.Contract.GetSettleInfo(&_IFileSysGetter.CallOpts, _pi, _ti)
}

// GetStoreInfo is a free data retrieval call binding the contract method 0x8f328ecd.
//
// Solidity: function getStoreInfo(uint64 _ui, uint64 _pi, uint8 _ti) view returns((uint64,uint64,uint64,uint256))
func (_IFileSysGetter *IFileSysGetterCaller) GetStoreInfo(opts *bind.CallOpts, _ui uint64, _pi uint64, _ti uint8) (StoreOut, error) {
	var out []interface{}
	err := _IFileSysGetter.contract.Call(opts, &out, "getStoreInfo", _ui, _pi, _ti)

	if err != nil {
		return *new(StoreOut), err
	}

	out0 := *abi.ConvertType(out[0], new(StoreOut)).(*StoreOut)

	return out0, err

}

// GetStoreInfo is a free data retrieval call binding the contract method 0x8f328ecd.
//
// Solidity: function getStoreInfo(uint64 _ui, uint64 _pi, uint8 _ti) view returns((uint64,uint64,uint64,uint256))
func (_IFileSysGetter *IFileSysGetterSession) GetStoreInfo(_ui uint64, _pi uint64, _ti uint8) (StoreOut, error) {
	return _IFileSysGetter.Contract.GetStoreInfo(&_IFileSysGetter.CallOpts, _ui, _pi, _ti)
}

// GetStoreInfo is a free data retrieval call binding the contract method 0x8f328ecd.
//
// Solidity: function getStoreInfo(uint64 _ui, uint64 _pi, uint8 _ti) view returns((uint64,uint64,uint64,uint256))
func (_IFileSysGetter *IFileSysGetterCallerSession) GetStoreInfo(_ui uint64, _pi uint64, _ti uint8) (StoreOut, error) {
	return _IFileSysGetter.Contract.GetStoreInfo(&_IFileSysGetter.CallOpts, _ui, _pi, _ti)
}

// IFileSysSetterABI is the input ABI used to generate the binding from.
const IFileSysSetterABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"ti\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"ui\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"pi\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"start\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"end\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sprice\",\"type\":\"uint256\"}],\"name\":\"AddOrder\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"ti\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"ui\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"pi\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"start\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"end\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sprice\",\"type\":\"uint256\"}],\"name\":\"AddRepair\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"ti\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"ui\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"pi\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"start\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"end\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sprice\",\"type\":\"uint256\"}],\"name\":\"SubOrder\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"uIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"start\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"end\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"tIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"sprice\",\"type\":\"uint256\"}],\"internalType\":\"structOrderIn\",\"name\":\"ps\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"_mr\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_tr\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"}],\"name\":\"addOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_rtype\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_kPB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_pPB\",\"type\":\"uint256\"}],\"name\":\"addPenalty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"uIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"start\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"end\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"tIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"sprice\",\"type\":\"uint256\"}],\"internalType\":\"structOrderIn\",\"name\":\"ps\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"}],\"name\":\"addRepair\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"pIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"tIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"pay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lost\",\"type\":\"uint256\"}],\"internalType\":\"structPWIn\",\"name\":\"ps\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"}],\"name\":\"proWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isLock\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"money\",\"type\":\"uint256\"}],\"name\":\"recharge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"uIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"start\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"end\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"tIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"sprice\",\"type\":\"uint256\"}],\"internalType\":\"structOrderIn\",\"name\":\"ps\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"}],\"name\":\"subOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"money\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_lock\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IFileSysSetterFuncSigs maps the 4-byte function signature to its string representation.
var IFileSysSetterFuncSigs = map[string]string{
	"9058d8e7": "addOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256),uint8,uint8,uint64)",
	"2449a168": "addPenalty(uint64,uint64,uint8,uint8,uint256,uint256)",
	"80faaf88": "addRepair((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256),uint64)",
	"a4703e16": "proWithdraw((uint64,uint8,uint256,uint256),uint64)",
	"24d11d40": "recharge(uint64,uint8,bool,uint256)",
	"248d02a0": "subOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256),uint64)",
	"5affa0f3": "withdraw(uint64,uint8,uint256,uint256)",
}

// IFileSysSetter is an auto generated Go binding around an Ethereum contract.
type IFileSysSetter struct {
	IFileSysSetterCaller     // Read-only binding to the contract
	IFileSysSetterTransactor // Write-only binding to the contract
	IFileSysSetterFilterer   // Log filterer for contract events
}

// IFileSysSetterCaller is an auto generated read-only Go binding around an Ethereum contract.
type IFileSysSetterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFileSysSetterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IFileSysSetterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFileSysSetterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IFileSysSetterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFileSysSetterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IFileSysSetterSession struct {
	Contract     *IFileSysSetter   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IFileSysSetterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IFileSysSetterCallerSession struct {
	Contract *IFileSysSetterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IFileSysSetterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IFileSysSetterTransactorSession struct {
	Contract     *IFileSysSetterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IFileSysSetterRaw is an auto generated low-level Go binding around an Ethereum contract.
type IFileSysSetterRaw struct {
	Contract *IFileSysSetter // Generic contract binding to access the raw methods on
}

// IFileSysSetterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IFileSysSetterCallerRaw struct {
	Contract *IFileSysSetterCaller // Generic read-only contract binding to access the raw methods on
}

// IFileSysSetterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IFileSysSetterTransactorRaw struct {
	Contract *IFileSysSetterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIFileSysSetter creates a new instance of IFileSysSetter, bound to a specific deployed contract.
func NewIFileSysSetter(address common.Address, backend bind.ContractBackend) (*IFileSysSetter, error) {
	contract, err := bindIFileSysSetter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IFileSysSetter{IFileSysSetterCaller: IFileSysSetterCaller{contract: contract}, IFileSysSetterTransactor: IFileSysSetterTransactor{contract: contract}, IFileSysSetterFilterer: IFileSysSetterFilterer{contract: contract}}, nil
}

// NewIFileSysSetterCaller creates a new read-only instance of IFileSysSetter, bound to a specific deployed contract.
func NewIFileSysSetterCaller(address common.Address, caller bind.ContractCaller) (*IFileSysSetterCaller, error) {
	contract, err := bindIFileSysSetter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IFileSysSetterCaller{contract: contract}, nil
}

// NewIFileSysSetterTransactor creates a new write-only instance of IFileSysSetter, bound to a specific deployed contract.
func NewIFileSysSetterTransactor(address common.Address, transactor bind.ContractTransactor) (*IFileSysSetterTransactor, error) {
	contract, err := bindIFileSysSetter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IFileSysSetterTransactor{contract: contract}, nil
}

// NewIFileSysSetterFilterer creates a new log filterer instance of IFileSysSetter, bound to a specific deployed contract.
func NewIFileSysSetterFilterer(address common.Address, filterer bind.ContractFilterer) (*IFileSysSetterFilterer, error) {
	contract, err := bindIFileSysSetter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IFileSysSetterFilterer{contract: contract}, nil
}

// bindIFileSysSetter binds a generic wrapper to an already deployed contract.
func bindIFileSysSetter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IFileSysSetterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IFileSysSetter *IFileSysSetterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IFileSysSetter.Contract.IFileSysSetterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IFileSysSetter *IFileSysSetterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IFileSysSetter.Contract.IFileSysSetterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IFileSysSetter *IFileSysSetterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IFileSysSetter.Contract.IFileSysSetterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IFileSysSetter *IFileSysSetterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IFileSysSetter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IFileSysSetter *IFileSysSetterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IFileSysSetter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IFileSysSetter *IFileSysSetterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IFileSysSetter.Contract.contract.Transact(opts, method, params...)
}

// AddOrder is a paid mutator transaction binding the contract method 0x9058d8e7.
//
// Solidity: function addOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) ps, uint8 _mr, uint8 _tr, uint64 _gi) returns()
func (_IFileSysSetter *IFileSysSetterTransactor) AddOrder(opts *bind.TransactOpts, ps OrderIn, _mr uint8, _tr uint8, _gi uint64) (*types.Transaction, error) {
	return _IFileSysSetter.contract.Transact(opts, "addOrder", ps, _mr, _tr, _gi)
}

// AddOrder is a paid mutator transaction binding the contract method 0x9058d8e7.
//
// Solidity: function addOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) ps, uint8 _mr, uint8 _tr, uint64 _gi) returns()
func (_IFileSysSetter *IFileSysSetterSession) AddOrder(ps OrderIn, _mr uint8, _tr uint8, _gi uint64) (*types.Transaction, error) {
	return _IFileSysSetter.Contract.AddOrder(&_IFileSysSetter.TransactOpts, ps, _mr, _tr, _gi)
}

// AddOrder is a paid mutator transaction binding the contract method 0x9058d8e7.
//
// Solidity: function addOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) ps, uint8 _mr, uint8 _tr, uint64 _gi) returns()
func (_IFileSysSetter *IFileSysSetterTransactorSession) AddOrder(ps OrderIn, _mr uint8, _tr uint8, _gi uint64) (*types.Transaction, error) {
	return _IFileSysSetter.Contract.AddOrder(&_IFileSysSetter.TransactOpts, ps, _mr, _tr, _gi)
}

// AddPenalty is a paid mutator transaction binding the contract method 0x2449a168.
//
// Solidity: function addPenalty(uint64 _gi, uint64 _i, uint8 _ti, uint8 _rtype, uint256 _kPB, uint256 _pPB) returns()
func (_IFileSysSetter *IFileSysSetterTransactor) AddPenalty(opts *bind.TransactOpts, _gi uint64, _i uint64, _ti uint8, _rtype uint8, _kPB *big.Int, _pPB *big.Int) (*types.Transaction, error) {
	return _IFileSysSetter.contract.Transact(opts, "addPenalty", _gi, _i, _ti, _rtype, _kPB, _pPB)
}

// AddPenalty is a paid mutator transaction binding the contract method 0x2449a168.
//
// Solidity: function addPenalty(uint64 _gi, uint64 _i, uint8 _ti, uint8 _rtype, uint256 _kPB, uint256 _pPB) returns()
func (_IFileSysSetter *IFileSysSetterSession) AddPenalty(_gi uint64, _i uint64, _ti uint8, _rtype uint8, _kPB *big.Int, _pPB *big.Int) (*types.Transaction, error) {
	return _IFileSysSetter.Contract.AddPenalty(&_IFileSysSetter.TransactOpts, _gi, _i, _ti, _rtype, _kPB, _pPB)
}

// AddPenalty is a paid mutator transaction binding the contract method 0x2449a168.
//
// Solidity: function addPenalty(uint64 _gi, uint64 _i, uint8 _ti, uint8 _rtype, uint256 _kPB, uint256 _pPB) returns()
func (_IFileSysSetter *IFileSysSetterTransactorSession) AddPenalty(_gi uint64, _i uint64, _ti uint8, _rtype uint8, _kPB *big.Int, _pPB *big.Int) (*types.Transaction, error) {
	return _IFileSysSetter.Contract.AddPenalty(&_IFileSysSetter.TransactOpts, _gi, _i, _ti, _rtype, _kPB, _pPB)
}

// AddRepair is a paid mutator transaction binding the contract method 0x80faaf88.
//
// Solidity: function addRepair((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) ps, uint64 _gi) returns()
func (_IFileSysSetter *IFileSysSetterTransactor) AddRepair(opts *bind.TransactOpts, ps OrderIn, _gi uint64) (*types.Transaction, error) {
	return _IFileSysSetter.contract.Transact(opts, "addRepair", ps, _gi)
}

// AddRepair is a paid mutator transaction binding the contract method 0x80faaf88.
//
// Solidity: function addRepair((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) ps, uint64 _gi) returns()
func (_IFileSysSetter *IFileSysSetterSession) AddRepair(ps OrderIn, _gi uint64) (*types.Transaction, error) {
	return _IFileSysSetter.Contract.AddRepair(&_IFileSysSetter.TransactOpts, ps, _gi)
}

// AddRepair is a paid mutator transaction binding the contract method 0x80faaf88.
//
// Solidity: function addRepair((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) ps, uint64 _gi) returns()
func (_IFileSysSetter *IFileSysSetterTransactorSession) AddRepair(ps OrderIn, _gi uint64) (*types.Transaction, error) {
	return _IFileSysSetter.Contract.AddRepair(&_IFileSysSetter.TransactOpts, ps, _gi)
}

// ProWithdraw is a paid mutator transaction binding the contract method 0xa4703e16.
//
// Solidity: function proWithdraw((uint64,uint8,uint256,uint256) ps, uint64 _gi) returns(uint256)
func (_IFileSysSetter *IFileSysSetterTransactor) ProWithdraw(opts *bind.TransactOpts, ps PWIn, _gi uint64) (*types.Transaction, error) {
	return _IFileSysSetter.contract.Transact(opts, "proWithdraw", ps, _gi)
}

// ProWithdraw is a paid mutator transaction binding the contract method 0xa4703e16.
//
// Solidity: function proWithdraw((uint64,uint8,uint256,uint256) ps, uint64 _gi) returns(uint256)
func (_IFileSysSetter *IFileSysSetterSession) ProWithdraw(ps PWIn, _gi uint64) (*types.Transaction, error) {
	return _IFileSysSetter.Contract.ProWithdraw(&_IFileSysSetter.TransactOpts, ps, _gi)
}

// ProWithdraw is a paid mutator transaction binding the contract method 0xa4703e16.
//
// Solidity: function proWithdraw((uint64,uint8,uint256,uint256) ps, uint64 _gi) returns(uint256)
func (_IFileSysSetter *IFileSysSetterTransactorSession) ProWithdraw(ps PWIn, _gi uint64) (*types.Transaction, error) {
	return _IFileSysSetter.Contract.ProWithdraw(&_IFileSysSetter.TransactOpts, ps, _gi)
}

// Recharge is a paid mutator transaction binding the contract method 0x24d11d40.
//
// Solidity: function recharge(uint64 _i, uint8 _ti, bool isLock, uint256 money) returns()
func (_IFileSysSetter *IFileSysSetterTransactor) Recharge(opts *bind.TransactOpts, _i uint64, _ti uint8, isLock bool, money *big.Int) (*types.Transaction, error) {
	return _IFileSysSetter.contract.Transact(opts, "recharge", _i, _ti, isLock, money)
}

// Recharge is a paid mutator transaction binding the contract method 0x24d11d40.
//
// Solidity: function recharge(uint64 _i, uint8 _ti, bool isLock, uint256 money) returns()
func (_IFileSysSetter *IFileSysSetterSession) Recharge(_i uint64, _ti uint8, isLock bool, money *big.Int) (*types.Transaction, error) {
	return _IFileSysSetter.Contract.Recharge(&_IFileSysSetter.TransactOpts, _i, _ti, isLock, money)
}

// Recharge is a paid mutator transaction binding the contract method 0x24d11d40.
//
// Solidity: function recharge(uint64 _i, uint8 _ti, bool isLock, uint256 money) returns()
func (_IFileSysSetter *IFileSysSetterTransactorSession) Recharge(_i uint64, _ti uint8, isLock bool, money *big.Int) (*types.Transaction, error) {
	return _IFileSysSetter.Contract.Recharge(&_IFileSysSetter.TransactOpts, _i, _ti, isLock, money)
}

// SubOrder is a paid mutator transaction binding the contract method 0x248d02a0.
//
// Solidity: function subOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) ps, uint64 _gi) returns()
func (_IFileSysSetter *IFileSysSetterTransactor) SubOrder(opts *bind.TransactOpts, ps OrderIn, _gi uint64) (*types.Transaction, error) {
	return _IFileSysSetter.contract.Transact(opts, "subOrder", ps, _gi)
}

// SubOrder is a paid mutator transaction binding the contract method 0x248d02a0.
//
// Solidity: function subOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) ps, uint64 _gi) returns()
func (_IFileSysSetter *IFileSysSetterSession) SubOrder(ps OrderIn, _gi uint64) (*types.Transaction, error) {
	return _IFileSysSetter.Contract.SubOrder(&_IFileSysSetter.TransactOpts, ps, _gi)
}

// SubOrder is a paid mutator transaction binding the contract method 0x248d02a0.
//
// Solidity: function subOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) ps, uint64 _gi) returns()
func (_IFileSysSetter *IFileSysSetterTransactorSession) SubOrder(ps OrderIn, _gi uint64) (*types.Transaction, error) {
	return _IFileSysSetter.Contract.SubOrder(&_IFileSysSetter.TransactOpts, ps, _gi)
}

// Withdraw is a paid mutator transaction binding the contract method 0x5affa0f3.
//
// Solidity: function withdraw(uint64 _i, uint8 _ti, uint256 money, uint256 _lock) returns(uint256)
func (_IFileSysSetter *IFileSysSetterTransactor) Withdraw(opts *bind.TransactOpts, _i uint64, _ti uint8, money *big.Int, _lock *big.Int) (*types.Transaction, error) {
	return _IFileSysSetter.contract.Transact(opts, "withdraw", _i, _ti, money, _lock)
}

// Withdraw is a paid mutator transaction binding the contract method 0x5affa0f3.
//
// Solidity: function withdraw(uint64 _i, uint8 _ti, uint256 money, uint256 _lock) returns(uint256)
func (_IFileSysSetter *IFileSysSetterSession) Withdraw(_i uint64, _ti uint8, money *big.Int, _lock *big.Int) (*types.Transaction, error) {
	return _IFileSysSetter.Contract.Withdraw(&_IFileSysSetter.TransactOpts, _i, _ti, money, _lock)
}

// Withdraw is a paid mutator transaction binding the contract method 0x5affa0f3.
//
// Solidity: function withdraw(uint64 _i, uint8 _ti, uint256 money, uint256 _lock) returns(uint256)
func (_IFileSysSetter *IFileSysSetterTransactorSession) Withdraw(_i uint64, _ti uint8, money *big.Int, _lock *big.Int) (*types.Transaction, error) {
	return _IFileSysSetter.Contract.Withdraw(&_IFileSysSetter.TransactOpts, _i, _ti, money, _lock)
}

// IFileSysSetterAddOrderIterator is returned from FilterAddOrder and is used to iterate over the raw logs and unpacked data for AddOrder events raised by the IFileSysSetter contract.
type IFileSysSetterAddOrderIterator struct {
	Event *IFileSysSetterAddOrder // Event containing the contract specifics and raw log

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
func (it *IFileSysSetterAddOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IFileSysSetterAddOrder)
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
		it.Event = new(IFileSysSetterAddOrder)
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
func (it *IFileSysSetterAddOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IFileSysSetterAddOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IFileSysSetterAddOrder represents a AddOrder event raised by the IFileSysSetter contract.
type IFileSysSetterAddOrder struct {
	Ti     uint8
	Ui     uint64
	Pi     uint64
	Start  uint64
	End    uint64
	Size   uint64
	Sprice *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAddOrder is a free log retrieval operation binding the contract event 0x9ad187def591246eb8fff84534926005f62988636593745516c0ac08fa086314.
//
// Solidity: event AddOrder(uint8 indexed ti, uint64 indexed ui, uint64 indexed pi, uint64 start, uint64 end, uint64 size, uint256 sprice)
func (_IFileSysSetter *IFileSysSetterFilterer) FilterAddOrder(opts *bind.FilterOpts, ti []uint8, ui []uint64, pi []uint64) (*IFileSysSetterAddOrderIterator, error) {

	var tiRule []interface{}
	for _, tiItem := range ti {
		tiRule = append(tiRule, tiItem)
	}
	var uiRule []interface{}
	for _, uiItem := range ui {
		uiRule = append(uiRule, uiItem)
	}
	var piRule []interface{}
	for _, piItem := range pi {
		piRule = append(piRule, piItem)
	}

	logs, sub, err := _IFileSysSetter.contract.FilterLogs(opts, "AddOrder", tiRule, uiRule, piRule)
	if err != nil {
		return nil, err
	}
	return &IFileSysSetterAddOrderIterator{contract: _IFileSysSetter.contract, event: "AddOrder", logs: logs, sub: sub}, nil
}

// WatchAddOrder is a free log subscription operation binding the contract event 0x9ad187def591246eb8fff84534926005f62988636593745516c0ac08fa086314.
//
// Solidity: event AddOrder(uint8 indexed ti, uint64 indexed ui, uint64 indexed pi, uint64 start, uint64 end, uint64 size, uint256 sprice)
func (_IFileSysSetter *IFileSysSetterFilterer) WatchAddOrder(opts *bind.WatchOpts, sink chan<- *IFileSysSetterAddOrder, ti []uint8, ui []uint64, pi []uint64) (event.Subscription, error) {

	var tiRule []interface{}
	for _, tiItem := range ti {
		tiRule = append(tiRule, tiItem)
	}
	var uiRule []interface{}
	for _, uiItem := range ui {
		uiRule = append(uiRule, uiItem)
	}
	var piRule []interface{}
	for _, piItem := range pi {
		piRule = append(piRule, piItem)
	}

	logs, sub, err := _IFileSysSetter.contract.WatchLogs(opts, "AddOrder", tiRule, uiRule, piRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IFileSysSetterAddOrder)
				if err := _IFileSysSetter.contract.UnpackLog(event, "AddOrder", log); err != nil {
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

// ParseAddOrder is a log parse operation binding the contract event 0x9ad187def591246eb8fff84534926005f62988636593745516c0ac08fa086314.
//
// Solidity: event AddOrder(uint8 indexed ti, uint64 indexed ui, uint64 indexed pi, uint64 start, uint64 end, uint64 size, uint256 sprice)
func (_IFileSysSetter *IFileSysSetterFilterer) ParseAddOrder(log types.Log) (*IFileSysSetterAddOrder, error) {
	event := new(IFileSysSetterAddOrder)
	if err := _IFileSysSetter.contract.UnpackLog(event, "AddOrder", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IFileSysSetterAddRepairIterator is returned from FilterAddRepair and is used to iterate over the raw logs and unpacked data for AddRepair events raised by the IFileSysSetter contract.
type IFileSysSetterAddRepairIterator struct {
	Event *IFileSysSetterAddRepair // Event containing the contract specifics and raw log

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
func (it *IFileSysSetterAddRepairIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IFileSysSetterAddRepair)
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
		it.Event = new(IFileSysSetterAddRepair)
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
func (it *IFileSysSetterAddRepairIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IFileSysSetterAddRepairIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IFileSysSetterAddRepair represents a AddRepair event raised by the IFileSysSetter contract.
type IFileSysSetterAddRepair struct {
	Ti     uint8
	Ui     uint64
	Pi     uint64
	Start  uint64
	End    uint64
	Size   uint64
	Sprice *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAddRepair is a free log retrieval operation binding the contract event 0x82f944cf0134663ff19eea526226f1cd8d5746a71cc837983659b8cbb2674e48.
//
// Solidity: event AddRepair(uint8 indexed ti, uint64 indexed ui, uint64 indexed pi, uint64 start, uint64 end, uint64 size, uint256 sprice)
func (_IFileSysSetter *IFileSysSetterFilterer) FilterAddRepair(opts *bind.FilterOpts, ti []uint8, ui []uint64, pi []uint64) (*IFileSysSetterAddRepairIterator, error) {

	var tiRule []interface{}
	for _, tiItem := range ti {
		tiRule = append(tiRule, tiItem)
	}
	var uiRule []interface{}
	for _, uiItem := range ui {
		uiRule = append(uiRule, uiItem)
	}
	var piRule []interface{}
	for _, piItem := range pi {
		piRule = append(piRule, piItem)
	}

	logs, sub, err := _IFileSysSetter.contract.FilterLogs(opts, "AddRepair", tiRule, uiRule, piRule)
	if err != nil {
		return nil, err
	}
	return &IFileSysSetterAddRepairIterator{contract: _IFileSysSetter.contract, event: "AddRepair", logs: logs, sub: sub}, nil
}

// WatchAddRepair is a free log subscription operation binding the contract event 0x82f944cf0134663ff19eea526226f1cd8d5746a71cc837983659b8cbb2674e48.
//
// Solidity: event AddRepair(uint8 indexed ti, uint64 indexed ui, uint64 indexed pi, uint64 start, uint64 end, uint64 size, uint256 sprice)
func (_IFileSysSetter *IFileSysSetterFilterer) WatchAddRepair(opts *bind.WatchOpts, sink chan<- *IFileSysSetterAddRepair, ti []uint8, ui []uint64, pi []uint64) (event.Subscription, error) {

	var tiRule []interface{}
	for _, tiItem := range ti {
		tiRule = append(tiRule, tiItem)
	}
	var uiRule []interface{}
	for _, uiItem := range ui {
		uiRule = append(uiRule, uiItem)
	}
	var piRule []interface{}
	for _, piItem := range pi {
		piRule = append(piRule, piItem)
	}

	logs, sub, err := _IFileSysSetter.contract.WatchLogs(opts, "AddRepair", tiRule, uiRule, piRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IFileSysSetterAddRepair)
				if err := _IFileSysSetter.contract.UnpackLog(event, "AddRepair", log); err != nil {
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

// ParseAddRepair is a log parse operation binding the contract event 0x82f944cf0134663ff19eea526226f1cd8d5746a71cc837983659b8cbb2674e48.
//
// Solidity: event AddRepair(uint8 indexed ti, uint64 indexed ui, uint64 indexed pi, uint64 start, uint64 end, uint64 size, uint256 sprice)
func (_IFileSysSetter *IFileSysSetterFilterer) ParseAddRepair(log types.Log) (*IFileSysSetterAddRepair, error) {
	event := new(IFileSysSetterAddRepair)
	if err := _IFileSysSetter.contract.UnpackLog(event, "AddRepair", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IFileSysSetterSubOrderIterator is returned from FilterSubOrder and is used to iterate over the raw logs and unpacked data for SubOrder events raised by the IFileSysSetter contract.
type IFileSysSetterSubOrderIterator struct {
	Event *IFileSysSetterSubOrder // Event containing the contract specifics and raw log

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
func (it *IFileSysSetterSubOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IFileSysSetterSubOrder)
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
		it.Event = new(IFileSysSetterSubOrder)
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
func (it *IFileSysSetterSubOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IFileSysSetterSubOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IFileSysSetterSubOrder represents a SubOrder event raised by the IFileSysSetter contract.
type IFileSysSetterSubOrder struct {
	Ti     uint8
	Ui     uint64
	Pi     uint64
	Start  uint64
	End    uint64
	Size   uint64
	Sprice *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSubOrder is a free log retrieval operation binding the contract event 0xe42ef64dc7dfe9121f13026a70df328cf6687348fdf8f71745f8e40fa5e52b82.
//
// Solidity: event SubOrder(uint8 indexed ti, uint64 indexed ui, uint64 indexed pi, uint64 start, uint64 end, uint64 size, uint256 sprice)
func (_IFileSysSetter *IFileSysSetterFilterer) FilterSubOrder(opts *bind.FilterOpts, ti []uint8, ui []uint64, pi []uint64) (*IFileSysSetterSubOrderIterator, error) {

	var tiRule []interface{}
	for _, tiItem := range ti {
		tiRule = append(tiRule, tiItem)
	}
	var uiRule []interface{}
	for _, uiItem := range ui {
		uiRule = append(uiRule, uiItem)
	}
	var piRule []interface{}
	for _, piItem := range pi {
		piRule = append(piRule, piItem)
	}

	logs, sub, err := _IFileSysSetter.contract.FilterLogs(opts, "SubOrder", tiRule, uiRule, piRule)
	if err != nil {
		return nil, err
	}
	return &IFileSysSetterSubOrderIterator{contract: _IFileSysSetter.contract, event: "SubOrder", logs: logs, sub: sub}, nil
}

// WatchSubOrder is a free log subscription operation binding the contract event 0xe42ef64dc7dfe9121f13026a70df328cf6687348fdf8f71745f8e40fa5e52b82.
//
// Solidity: event SubOrder(uint8 indexed ti, uint64 indexed ui, uint64 indexed pi, uint64 start, uint64 end, uint64 size, uint256 sprice)
func (_IFileSysSetter *IFileSysSetterFilterer) WatchSubOrder(opts *bind.WatchOpts, sink chan<- *IFileSysSetterSubOrder, ti []uint8, ui []uint64, pi []uint64) (event.Subscription, error) {

	var tiRule []interface{}
	for _, tiItem := range ti {
		tiRule = append(tiRule, tiItem)
	}
	var uiRule []interface{}
	for _, uiItem := range ui {
		uiRule = append(uiRule, uiItem)
	}
	var piRule []interface{}
	for _, piItem := range pi {
		piRule = append(piRule, piItem)
	}

	logs, sub, err := _IFileSysSetter.contract.WatchLogs(opts, "SubOrder", tiRule, uiRule, piRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IFileSysSetterSubOrder)
				if err := _IFileSysSetter.contract.UnpackLog(event, "SubOrder", log); err != nil {
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

// ParseSubOrder is a log parse operation binding the contract event 0xe42ef64dc7dfe9121f13026a70df328cf6687348fdf8f71745f8e40fa5e52b82.
//
// Solidity: event SubOrder(uint8 indexed ti, uint64 indexed ui, uint64 indexed pi, uint64 start, uint64 end, uint64 size, uint256 sprice)
func (_IFileSysSetter *IFileSysSetterFilterer) ParseSubOrder(log types.Log) (*IFileSysSetterSubOrder, error) {
	event := new(IFileSysSetterSubOrder)
	if err := _IFileSysSetter.contract.UnpackLog(event, "SubOrder", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
