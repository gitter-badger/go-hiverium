// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"math/big"
	"strings"

	ethereum "github.com/unification-com/mainchain"
	"github.com/unification-com/mainchain/accounts/abi"
	"github.com/unification-com/mainchain/accounts/abi/bind"
	"github.com/unification-com/mainchain/common"
	"github.com/unification-com/mainchain/core/types"
	"github.com/unification-com/mainchain/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// WRKChainRootABI is the input ABI used to generate the binding from.
const WRKChainRootABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getMasterRegistrar\",\"outputs\":[{\"name\":\"masterRegistrar_\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_chainId\",\"type\":\"uint256\"},{\"name\":\"_newAuthAddresses\",\"type\":\"address[]\"}],\"name\":\"setAuthAddresses\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_chainId\",\"type\":\"uint256\"},{\"name\":\"_authAddresses\",\"type\":\"address[]\"},{\"name\":\"_genesisHash\",\"type\":\"bytes32\"}],\"name\":\"registerWrkChain\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_chainId\",\"type\":\"uint256\"},{\"name\":\"_height\",\"type\":\"uint256\"},{\"name\":\"_hash\",\"type\":\"bytes32\"},{\"name\":\"_parentHash\",\"type\":\"bytes32\"},{\"name\":\"_receiptRoot\",\"type\":\"bytes32\"},{\"name\":\"_txRoot\",\"type\":\"bytes32\"},{\"name\":\"_stateRoot\",\"type\":\"bytes32\"},{\"name\":\"_blockSigner\",\"type\":\"address\"}],\"name\":\"recordHeader\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_chainId\",\"type\":\"uint256\"}],\"name\":\"getAuthAddresses\",\"outputs\":[{\"name\":\"authAddressesIdx_\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_registrar\",\"type\":\"address\"}],\"name\":\"addRegistrar\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_registrar\",\"type\":\"address\"}],\"name\":\"removeRegistrar\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_chainId\",\"type\":\"uint256\"}],\"name\":\"getGenesis\",\"outputs\":[{\"name\":\"genesisHash_\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_masterRegistrar\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_chainId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_height\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_parentHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_receiptRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_txRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_stateRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_blockSigner\",\"type\":\"address\"}],\"name\":\"RecordHeader\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_chainId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_genesisHash\",\"type\":\"bytes32\"}],\"name\":\"RegisterWrkChain\",\"type\":\"event\"}]"

// WRKChainRootBin is the compiled bytecode used for deploying new contracts.
const WRKChainRootBin = `0x608060405234801561001057600080fd5b50604051602080610ed9833981016040525160028054600160a060020a031916600160a060020a03909216919091179055610e89806100506000396000f30060806040526004361061008d5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416630bd2cb6e8114610092578063244658f7146100c35780635be4e2281461011f578063a78c49881461017b578063aa0d14de146101b1578063af92a69314610219578063ba904eed1461023a578063cfdf3e841461025b575b600080fd5b34801561009e57600080fd5b506100a7610285565b60408051600160a060020a039092168252519081900360200190f35b3480156100cf57600080fd5b5060408051602060046024803582810135848102808701860190975280865261011d968435963696604495919490910192918291850190849080828437509497506102949650505050505050565b005b34801561012b57600080fd5b5060408051602060046024803582810135848102808701860190975280865261011d9684359636966044959194909101929182918501908490808284375094975050933594506105d19350505050565b34801561018757600080fd5b5061011d60043560243560443560643560843560a43560c435600160a060020a0360e435166108d2565b3480156101bd57600080fd5b506101c9600435610afb565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156102055781810151838201526020016101ed565b505050509050019250505060405180910390f35b34801561022557600080fd5b5061011d600160a060020a0360043516610c07565b34801561024657600080fd5b5061011d600160a060020a0360043516610c67565b34801561026757600080fd5b50610273600435610cc1565b60408051918252519081900360200190f35b600254600160a060020a031690565b60008281526020819052604081206002015481908190859060ff161515610305576040805160e560020a62461bcd02815260206004820152601660248201527f436861696e494420646f6573206e6f7420657869737400000000000000000000604482015290519081900360640190fd5b60008181526020818152604080832033845260030190915290205460ff16151560011461033157600080fd5b60008611610377576040805160e560020a62461bcd0281526020600482015260116024820152600080516020610e1e833981519152604482015290519081900360640190fd5b84516000106103d0576040805160e560020a62461bcd02815260206004820152601760248201527f4175746820616464726573736573207265717569726564000000000000000000604482015290519081900360640190fd5b845160001015610450576040805160e560020a62461bcd02815260206004820152602360248201527f4d6178207375626d697373696f6e2031302061646472657373657320616c6c6f60448201527f7765640000000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b60008681526020819052604090206002015460ff1615156104a9576040805160e560020a62461bcd0281526020600482015260176024820152600080516020610e3e833981519152604482015290519081900360640190fd5b6000868152602081905260408120945092505b6004840154831015610532576000806000888152602001908152602001600020600301600086600401868154811015156104f257fe5b600091825260208083209190910154600160a060020a031683528201929092526040019020805460ff1916911515919091179055600192909201916104bc565b600091505b84518210156105a55760016000808881526020019081526020016000206003016000878581518110151561056757fe5b602090810291909101810151600160a060020a03168252810191909152604001600020805460ff191691151591909117905560019190910190610537565b60008681526020818152604090912086516105c892600490920191880190610d77565b50505050505050565b33600090815260016020819052604082205460ff16151514806105fe5750600254600160a060020a031633145b151561060957600080fd5b6000841161064f576040805160e560020a62461bcd0281526020600482015260116024820152600080516020610e1e833981519152604482015290519081900360640190fd5b82516000106106ce576040805160e560020a62461bcd02815260206004820152602560248201527f496e697469616c20417574686f7269736564206164647265737365732072657160448201527f7569726564000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b8251600a101561074e576040805160e560020a62461bcd02815260206004820152602f60248201527f4d6178696d756d20313020496e697469616c20417574686f726973656420616460448201527f6472657373657320616c6c6f7765640000000000000000000000000000000000606482015290519081900360840190fd5b60008481526020819052604090206002015460ff16156107b8576040805160e560020a62461bcd02815260206004820152601760248201527f436861696e20494420616c726561647920657869737473000000000000000000604482015290519081900360640190fd5b60408051608081018252600080825260208083018681526001848601818152606086018a81528b86528585529690942085518155915190820155915160028301805460ff1916911515919091179055925180519293919261081f9260048501920190610d77565b506000925050505b82518110156108915760016000808681526020019081526020016000206003016000858481518110151561085757fe5b602090810291909101810151600160a060020a03168252810191909152604001600020805460ff1916911515919091179055600101610827565b604080518581526020810184905281517fd49920535e715987d1fbf16411df1f3f3347d62e741006463c13550661ce4b69929181900390910190a150505050565b600088815260208190526040812060020154899060ff16151561093f576040805160e560020a62461bcd02815260206004820152601660248201527f436861696e494420646f6573206e6f7420657869737400000000000000000000604482015290519081900360640190fd5b60008181526020818152604080832033845260030190915290205460ff16151560011461096b57600080fd5b60008a116109b1576040805160e560020a62461bcd0281526020600482015260116024820152600080516020610e1e833981519152604482015290519081900360640190fd5b60008a81526020819052604090206002015460ff161515610a0a576040805160e560020a62461bcd0281526020600482015260176024820152600080516020610e3e833981519152604482015290519081900360640190fd5b60008a815260208190526040902080549092508911610a73576040805160e560020a62461bcd02815260206004820152601d60248201527f426c6f636b2068656164657220616c7265616479207265636f72646564000000604482015290519081900360640190fd5b60008a815260208181526040918290208b905581518c81529081018b90528082018a9052606081018990526080810188905260a0810187905260c08101869052600160a060020a03851660e082015290517f943621e8fb786306cd7b16a4467604107bbf3b0b04b536d39a245ea68c1062ef918190036101000190a150505050505050505050565b606060008211610b43576040805160e560020a62461bcd0281526020600482015260116024820152600080516020610e1e833981519152604482015290519081900360640190fd5b60008281526020819052604090206002015460ff161515610b9c576040805160e560020a62461bcd0281526020600482015260176024820152600080516020610e3e833981519152604482015290519081900360640190fd5b6000828152602081815260409182902060040180548351818402810184019094528084529091830182828015610bfb57602002820191906000526020600020905b8154600160a060020a03168152600190910190602001808311610bdd575b50505050509050919050565b3360009081526001602081905260409091205460ff1615151480610c355750600254600160a060020a031633145b1515610c4057600080fd5b600160a060020a03166000908152600160208190526040909120805460ff19169091179055565b3360009081526001602081905260409091205460ff1615151480610c955750600254600160a060020a031633145b1515610ca057600080fd5b600160a060020a03166000908152600160205260409020805460ff19169055565b6000808211610d08576040805160e560020a62461bcd0281526020600482015260116024820152600080516020610e1e833981519152604482015290519081900360640190fd5b60008281526020819052604090206002015460ff161515610d61576040805160e560020a62461bcd0281526020600482015260176024820152600080516020610e3e833981519152604482015290519081900360640190fd5b5060009081526020819052604090206001015490565b828054828255906000526020600020908101928215610dd9579160200282015b82811115610dd9578251825473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03909116178255602090920191600190910190610d97565b50610de5929150610de9565b5090565b610e1a91905b80821115610de557805473ffffffffffffffffffffffffffffffffffffffff19168155600101610def565b905600436861696e204944207265717569726564000000000000000000000000000000436861696e20494420646f6573206e6f74206578697374000000000000000000a165627a7a723058206686536df06442c895b30612f01b25ba0bb088beb1bc69e6766dbfb95e0ee76f0029`

// DeployWRKChainRoot deploys a new Ethereum contract, binding an instance of WRKChainRoot to it.
func DeployWRKChainRoot(auth *bind.TransactOpts, backend bind.ContractBackend, _masterRegistrar common.Address) (common.Address, *types.Transaction, *WRKChainRoot, error) {
	parsed, err := abi.JSON(strings.NewReader(WRKChainRootABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(WRKChainRootBin), backend, _masterRegistrar)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &WRKChainRoot{WRKChainRootCaller: WRKChainRootCaller{contract: contract}, WRKChainRootTransactor: WRKChainRootTransactor{contract: contract}, WRKChainRootFilterer: WRKChainRootFilterer{contract: contract}}, nil
}

// WRKChainRoot is an auto generated Go binding around an Ethereum contract.
type WRKChainRoot struct {
	WRKChainRootCaller     // Read-only binding to the contract
	WRKChainRootTransactor // Write-only binding to the contract
	WRKChainRootFilterer   // Log filterer for contract events
}

// WRKChainRootCaller is an auto generated read-only Go binding around an Ethereum contract.
type WRKChainRootCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WRKChainRootTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WRKChainRootTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WRKChainRootFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WRKChainRootFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WRKChainRootSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WRKChainRootSession struct {
	Contract     *WRKChainRoot     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WRKChainRootCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WRKChainRootCallerSession struct {
	Contract *WRKChainRootCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// WRKChainRootTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WRKChainRootTransactorSession struct {
	Contract     *WRKChainRootTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// WRKChainRootRaw is an auto generated low-level Go binding around an Ethereum contract.
type WRKChainRootRaw struct {
	Contract *WRKChainRoot // Generic contract binding to access the raw methods on
}

// WRKChainRootCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WRKChainRootCallerRaw struct {
	Contract *WRKChainRootCaller // Generic read-only contract binding to access the raw methods on
}

// WRKChainRootTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WRKChainRootTransactorRaw struct {
	Contract *WRKChainRootTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWRKChainRoot creates a new instance of WRKChainRoot, bound to a specific deployed contract.
func NewWRKChainRoot(address common.Address, backend bind.ContractBackend) (*WRKChainRoot, error) {
	contract, err := bindWRKChainRoot(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WRKChainRoot{WRKChainRootCaller: WRKChainRootCaller{contract: contract}, WRKChainRootTransactor: WRKChainRootTransactor{contract: contract}, WRKChainRootFilterer: WRKChainRootFilterer{contract: contract}}, nil
}

// NewWRKChainRootCaller creates a new read-only instance of WRKChainRoot, bound to a specific deployed contract.
func NewWRKChainRootCaller(address common.Address, caller bind.ContractCaller) (*WRKChainRootCaller, error) {
	contract, err := bindWRKChainRoot(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WRKChainRootCaller{contract: contract}, nil
}

// NewWRKChainRootTransactor creates a new write-only instance of WRKChainRoot, bound to a specific deployed contract.
func NewWRKChainRootTransactor(address common.Address, transactor bind.ContractTransactor) (*WRKChainRootTransactor, error) {
	contract, err := bindWRKChainRoot(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WRKChainRootTransactor{contract: contract}, nil
}

// NewWRKChainRootFilterer creates a new log filterer instance of WRKChainRoot, bound to a specific deployed contract.
func NewWRKChainRootFilterer(address common.Address, filterer bind.ContractFilterer) (*WRKChainRootFilterer, error) {
	contract, err := bindWRKChainRoot(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WRKChainRootFilterer{contract: contract}, nil
}

// bindWRKChainRoot binds a generic wrapper to an already deployed contract.
func bindWRKChainRoot(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WRKChainRootABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WRKChainRoot *WRKChainRootRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _WRKChainRoot.Contract.WRKChainRootCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WRKChainRoot *WRKChainRootRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WRKChainRoot.Contract.WRKChainRootTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WRKChainRoot *WRKChainRootRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WRKChainRoot.Contract.WRKChainRootTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WRKChainRoot *WRKChainRootCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _WRKChainRoot.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WRKChainRoot *WRKChainRootTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WRKChainRoot.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WRKChainRoot *WRKChainRootTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WRKChainRoot.Contract.contract.Transact(opts, method, params...)
}

// GetAuthAddresses is a free data retrieval call binding the contract method 0xaa0d14de.
//
// Solidity: function getAuthAddresses(uint256 _chainId) constant returns(address[] authAddressesIdx_)
func (_WRKChainRoot *WRKChainRootCaller) GetAuthAddresses(opts *bind.CallOpts, _chainId *big.Int) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _WRKChainRoot.contract.Call(opts, out, "getAuthAddresses", _chainId)
	return *ret0, err
}

// GetAuthAddresses is a free data retrieval call binding the contract method 0xaa0d14de.
//
// Solidity: function getAuthAddresses(uint256 _chainId) constant returns(address[] authAddressesIdx_)
func (_WRKChainRoot *WRKChainRootSession) GetAuthAddresses(_chainId *big.Int) ([]common.Address, error) {
	return _WRKChainRoot.Contract.GetAuthAddresses(&_WRKChainRoot.CallOpts, _chainId)
}

// GetAuthAddresses is a free data retrieval call binding the contract method 0xaa0d14de.
//
// Solidity: function getAuthAddresses(uint256 _chainId) constant returns(address[] authAddressesIdx_)
func (_WRKChainRoot *WRKChainRootCallerSession) GetAuthAddresses(_chainId *big.Int) ([]common.Address, error) {
	return _WRKChainRoot.Contract.GetAuthAddresses(&_WRKChainRoot.CallOpts, _chainId)
}

// GetGenesis is a free data retrieval call binding the contract method 0xcfdf3e84.
//
// Solidity: function getGenesis(uint256 _chainId) constant returns(bytes32 genesisHash_)
func (_WRKChainRoot *WRKChainRootCaller) GetGenesis(opts *bind.CallOpts, _chainId *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _WRKChainRoot.contract.Call(opts, out, "getGenesis", _chainId)
	return *ret0, err
}

// GetGenesis is a free data retrieval call binding the contract method 0xcfdf3e84.
//
// Solidity: function getGenesis(uint256 _chainId) constant returns(bytes32 genesisHash_)
func (_WRKChainRoot *WRKChainRootSession) GetGenesis(_chainId *big.Int) ([32]byte, error) {
	return _WRKChainRoot.Contract.GetGenesis(&_WRKChainRoot.CallOpts, _chainId)
}

// GetGenesis is a free data retrieval call binding the contract method 0xcfdf3e84.
//
// Solidity: function getGenesis(uint256 _chainId) constant returns(bytes32 genesisHash_)
func (_WRKChainRoot *WRKChainRootCallerSession) GetGenesis(_chainId *big.Int) ([32]byte, error) {
	return _WRKChainRoot.Contract.GetGenesis(&_WRKChainRoot.CallOpts, _chainId)
}

// GetMasterRegistrar is a free data retrieval call binding the contract method 0x0bd2cb6e.
//
// Solidity: function getMasterRegistrar() constant returns(address masterRegistrar_)
func (_WRKChainRoot *WRKChainRootCaller) GetMasterRegistrar(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _WRKChainRoot.contract.Call(opts, out, "getMasterRegistrar")
	return *ret0, err
}

// GetMasterRegistrar is a free data retrieval call binding the contract method 0x0bd2cb6e.
//
// Solidity: function getMasterRegistrar() constant returns(address masterRegistrar_)
func (_WRKChainRoot *WRKChainRootSession) GetMasterRegistrar() (common.Address, error) {
	return _WRKChainRoot.Contract.GetMasterRegistrar(&_WRKChainRoot.CallOpts)
}

// GetMasterRegistrar is a free data retrieval call binding the contract method 0x0bd2cb6e.
//
// Solidity: function getMasterRegistrar() constant returns(address masterRegistrar_)
func (_WRKChainRoot *WRKChainRootCallerSession) GetMasterRegistrar() (common.Address, error) {
	return _WRKChainRoot.Contract.GetMasterRegistrar(&_WRKChainRoot.CallOpts)
}

// AddRegistrar is a paid mutator transaction binding the contract method 0xaf92a693.
//
// Solidity: function addRegistrar(address _registrar) returns()
func (_WRKChainRoot *WRKChainRootTransactor) AddRegistrar(opts *bind.TransactOpts, _registrar common.Address) (*types.Transaction, error) {
	return _WRKChainRoot.contract.Transact(opts, "addRegistrar", _registrar)
}

// AddRegistrar is a paid mutator transaction binding the contract method 0xaf92a693.
//
// Solidity: function addRegistrar(address _registrar) returns()
func (_WRKChainRoot *WRKChainRootSession) AddRegistrar(_registrar common.Address) (*types.Transaction, error) {
	return _WRKChainRoot.Contract.AddRegistrar(&_WRKChainRoot.TransactOpts, _registrar)
}

// AddRegistrar is a paid mutator transaction binding the contract method 0xaf92a693.
//
// Solidity: function addRegistrar(address _registrar) returns()
func (_WRKChainRoot *WRKChainRootTransactorSession) AddRegistrar(_registrar common.Address) (*types.Transaction, error) {
	return _WRKChainRoot.Contract.AddRegistrar(&_WRKChainRoot.TransactOpts, _registrar)
}

// RecordHeader is a paid mutator transaction binding the contract method 0xa78c4988.
//
// Solidity: function recordHeader(uint256 _chainId, uint256 _height, bytes32 _hash, bytes32 _parentHash, bytes32 _receiptRoot, bytes32 _txRoot, bytes32 _stateRoot, address _blockSigner) returns()
func (_WRKChainRoot *WRKChainRootTransactor) RecordHeader(opts *bind.TransactOpts, _chainId *big.Int, _height *big.Int, _hash [32]byte, _parentHash [32]byte, _receiptRoot [32]byte, _txRoot [32]byte, _stateRoot [32]byte, _blockSigner common.Address) (*types.Transaction, error) {
	return _WRKChainRoot.contract.Transact(opts, "recordHeader", _chainId, _height, _hash, _parentHash, _receiptRoot, _txRoot, _stateRoot, _blockSigner)
}

// RecordHeader is a paid mutator transaction binding the contract method 0xa78c4988.
//
// Solidity: function recordHeader(uint256 _chainId, uint256 _height, bytes32 _hash, bytes32 _parentHash, bytes32 _receiptRoot, bytes32 _txRoot, bytes32 _stateRoot, address _blockSigner) returns()
func (_WRKChainRoot *WRKChainRootSession) RecordHeader(_chainId *big.Int, _height *big.Int, _hash [32]byte, _parentHash [32]byte, _receiptRoot [32]byte, _txRoot [32]byte, _stateRoot [32]byte, _blockSigner common.Address) (*types.Transaction, error) {
	return _WRKChainRoot.Contract.RecordHeader(&_WRKChainRoot.TransactOpts, _chainId, _height, _hash, _parentHash, _receiptRoot, _txRoot, _stateRoot, _blockSigner)
}

// RecordHeader is a paid mutator transaction binding the contract method 0xa78c4988.
//
// Solidity: function recordHeader(uint256 _chainId, uint256 _height, bytes32 _hash, bytes32 _parentHash, bytes32 _receiptRoot, bytes32 _txRoot, bytes32 _stateRoot, address _blockSigner) returns()
func (_WRKChainRoot *WRKChainRootTransactorSession) RecordHeader(_chainId *big.Int, _height *big.Int, _hash [32]byte, _parentHash [32]byte, _receiptRoot [32]byte, _txRoot [32]byte, _stateRoot [32]byte, _blockSigner common.Address) (*types.Transaction, error) {
	return _WRKChainRoot.Contract.RecordHeader(&_WRKChainRoot.TransactOpts, _chainId, _height, _hash, _parentHash, _receiptRoot, _txRoot, _stateRoot, _blockSigner)
}

// RegisterWrkChain is a paid mutator transaction binding the contract method 0x5be4e228.
//
// Solidity: function registerWrkChain(uint256 _chainId, address[] _authAddresses, bytes32 _genesisHash) returns()
func (_WRKChainRoot *WRKChainRootTransactor) RegisterWrkChain(opts *bind.TransactOpts, _chainId *big.Int, _authAddresses []common.Address, _genesisHash [32]byte) (*types.Transaction, error) {
	return _WRKChainRoot.contract.Transact(opts, "registerWrkChain", _chainId, _authAddresses, _genesisHash)
}

// RegisterWrkChain is a paid mutator transaction binding the contract method 0x5be4e228.
//
// Solidity: function registerWrkChain(uint256 _chainId, address[] _authAddresses, bytes32 _genesisHash) returns()
func (_WRKChainRoot *WRKChainRootSession) RegisterWrkChain(_chainId *big.Int, _authAddresses []common.Address, _genesisHash [32]byte) (*types.Transaction, error) {
	return _WRKChainRoot.Contract.RegisterWrkChain(&_WRKChainRoot.TransactOpts, _chainId, _authAddresses, _genesisHash)
}

// RegisterWrkChain is a paid mutator transaction binding the contract method 0x5be4e228.
//
// Solidity: function registerWrkChain(uint256 _chainId, address[] _authAddresses, bytes32 _genesisHash) returns()
func (_WRKChainRoot *WRKChainRootTransactorSession) RegisterWrkChain(_chainId *big.Int, _authAddresses []common.Address, _genesisHash [32]byte) (*types.Transaction, error) {
	return _WRKChainRoot.Contract.RegisterWrkChain(&_WRKChainRoot.TransactOpts, _chainId, _authAddresses, _genesisHash)
}

// RemoveRegistrar is a paid mutator transaction binding the contract method 0xba904eed.
//
// Solidity: function removeRegistrar(address _registrar) returns()
func (_WRKChainRoot *WRKChainRootTransactor) RemoveRegistrar(opts *bind.TransactOpts, _registrar common.Address) (*types.Transaction, error) {
	return _WRKChainRoot.contract.Transact(opts, "removeRegistrar", _registrar)
}

// RemoveRegistrar is a paid mutator transaction binding the contract method 0xba904eed.
//
// Solidity: function removeRegistrar(address _registrar) returns()
func (_WRKChainRoot *WRKChainRootSession) RemoveRegistrar(_registrar common.Address) (*types.Transaction, error) {
	return _WRKChainRoot.Contract.RemoveRegistrar(&_WRKChainRoot.TransactOpts, _registrar)
}

// RemoveRegistrar is a paid mutator transaction binding the contract method 0xba904eed.
//
// Solidity: function removeRegistrar(address _registrar) returns()
func (_WRKChainRoot *WRKChainRootTransactorSession) RemoveRegistrar(_registrar common.Address) (*types.Transaction, error) {
	return _WRKChainRoot.Contract.RemoveRegistrar(&_WRKChainRoot.TransactOpts, _registrar)
}

// SetAuthAddresses is a paid mutator transaction binding the contract method 0x244658f7.
//
// Solidity: function setAuthAddresses(uint256 _chainId, address[] _newAuthAddresses) returns()
func (_WRKChainRoot *WRKChainRootTransactor) SetAuthAddresses(opts *bind.TransactOpts, _chainId *big.Int, _newAuthAddresses []common.Address) (*types.Transaction, error) {
	return _WRKChainRoot.contract.Transact(opts, "setAuthAddresses", _chainId, _newAuthAddresses)
}

// SetAuthAddresses is a paid mutator transaction binding the contract method 0x244658f7.
//
// Solidity: function setAuthAddresses(uint256 _chainId, address[] _newAuthAddresses) returns()
func (_WRKChainRoot *WRKChainRootSession) SetAuthAddresses(_chainId *big.Int, _newAuthAddresses []common.Address) (*types.Transaction, error) {
	return _WRKChainRoot.Contract.SetAuthAddresses(&_WRKChainRoot.TransactOpts, _chainId, _newAuthAddresses)
}

// SetAuthAddresses is a paid mutator transaction binding the contract method 0x244658f7.
//
// Solidity: function setAuthAddresses(uint256 _chainId, address[] _newAuthAddresses) returns()
func (_WRKChainRoot *WRKChainRootTransactorSession) SetAuthAddresses(_chainId *big.Int, _newAuthAddresses []common.Address) (*types.Transaction, error) {
	return _WRKChainRoot.Contract.SetAuthAddresses(&_WRKChainRoot.TransactOpts, _chainId, _newAuthAddresses)
}

// WRKChainRootRecordHeaderIterator is returned from FilterRecordHeader and is used to iterate over the raw logs and unpacked data for RecordHeader events raised by the WRKChainRoot contract.
type WRKChainRootRecordHeaderIterator struct {
	Event *WRKChainRootRecordHeader // Event containing the contract specifics and raw log

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
func (it *WRKChainRootRecordHeaderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WRKChainRootRecordHeader)
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
		it.Event = new(WRKChainRootRecordHeader)
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
func (it *WRKChainRootRecordHeaderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WRKChainRootRecordHeaderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WRKChainRootRecordHeader represents a RecordHeader event raised by the WRKChainRoot contract.
type WRKChainRootRecordHeader struct {
	ChainId     *big.Int
	Height      *big.Int
	Hash        [32]byte
	ParentHash  [32]byte
	ReceiptRoot [32]byte
	TxRoot      [32]byte
	StateRoot   [32]byte
	BlockSigner common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRecordHeader is a free log retrieval operation binding the contract event 0x943621e8fb786306cd7b16a4467604107bbf3b0b04b536d39a245ea68c1062ef.
//
// Solidity: event RecordHeader(uint256 _chainId, uint256 _height, bytes32 _hash, bytes32 _parentHash, bytes32 _receiptRoot, bytes32 _txRoot, bytes32 _stateRoot, address _blockSigner)
func (_WRKChainRoot *WRKChainRootFilterer) FilterRecordHeader(opts *bind.FilterOpts) (*WRKChainRootRecordHeaderIterator, error) {

	logs, sub, err := _WRKChainRoot.contract.FilterLogs(opts, "RecordHeader")
	if err != nil {
		return nil, err
	}
	return &WRKChainRootRecordHeaderIterator{contract: _WRKChainRoot.contract, event: "RecordHeader", logs: logs, sub: sub}, nil
}

// WatchRecordHeader is a free log subscription operation binding the contract event 0x943621e8fb786306cd7b16a4467604107bbf3b0b04b536d39a245ea68c1062ef.
//
// Solidity: event RecordHeader(uint256 _chainId, uint256 _height, bytes32 _hash, bytes32 _parentHash, bytes32 _receiptRoot, bytes32 _txRoot, bytes32 _stateRoot, address _blockSigner)
func (_WRKChainRoot *WRKChainRootFilterer) WatchRecordHeader(opts *bind.WatchOpts, sink chan<- *WRKChainRootRecordHeader) (event.Subscription, error) {

	logs, sub, err := _WRKChainRoot.contract.WatchLogs(opts, "RecordHeader")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WRKChainRootRecordHeader)
				if err := _WRKChainRoot.contract.UnpackLog(event, "RecordHeader", log); err != nil {
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

// WRKChainRootRegisterWrkChainIterator is returned from FilterRegisterWrkChain and is used to iterate over the raw logs and unpacked data for RegisterWrkChain events raised by the WRKChainRoot contract.
type WRKChainRootRegisterWrkChainIterator struct {
	Event *WRKChainRootRegisterWrkChain // Event containing the contract specifics and raw log

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
func (it *WRKChainRootRegisterWrkChainIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WRKChainRootRegisterWrkChain)
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
		it.Event = new(WRKChainRootRegisterWrkChain)
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
func (it *WRKChainRootRegisterWrkChainIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WRKChainRootRegisterWrkChainIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WRKChainRootRegisterWrkChain represents a RegisterWrkChain event raised by the WRKChainRoot contract.
type WRKChainRootRegisterWrkChain struct {
	ChainId     *big.Int
	GenesisHash [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRegisterWrkChain is a free log retrieval operation binding the contract event 0xd49920535e715987d1fbf16411df1f3f3347d62e741006463c13550661ce4b69.
//
// Solidity: event RegisterWrkChain(uint256 _chainId, bytes32 _genesisHash)
func (_WRKChainRoot *WRKChainRootFilterer) FilterRegisterWrkChain(opts *bind.FilterOpts) (*WRKChainRootRegisterWrkChainIterator, error) {

	logs, sub, err := _WRKChainRoot.contract.FilterLogs(opts, "RegisterWrkChain")
	if err != nil {
		return nil, err
	}
	return &WRKChainRootRegisterWrkChainIterator{contract: _WRKChainRoot.contract, event: "RegisterWrkChain", logs: logs, sub: sub}, nil
}

// WatchRegisterWrkChain is a free log subscription operation binding the contract event 0xd49920535e715987d1fbf16411df1f3f3347d62e741006463c13550661ce4b69.
//
// Solidity: event RegisterWrkChain(uint256 _chainId, bytes32 _genesisHash)
func (_WRKChainRoot *WRKChainRootFilterer) WatchRegisterWrkChain(opts *bind.WatchOpts, sink chan<- *WRKChainRootRegisterWrkChain) (event.Subscription, error) {

	logs, sub, err := _WRKChainRoot.contract.WatchLogs(opts, "RegisterWrkChain")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WRKChainRootRegisterWrkChain)
				if err := _WRKChainRoot.contract.UnpackLog(event, "RegisterWrkChain", log); err != nil {
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
