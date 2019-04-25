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

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
const SafeMathBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a72305820a3968129a6b83f89d1789cb5e64d378309030066b0fdbe2b5e0b4a239f4247210029`

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

// WRKChainRootABI is the input ABI used to generate the binding from.
const WRKChainRootABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_chainId\",\"type\":\"uint256\"},{\"name\":\"_newAuthAddresses\",\"type\":\"address[]\"}],\"name\":\"setAuthAddresses\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_chainId\",\"type\":\"uint256\"},{\"name\":\"_authAddresses\",\"type\":\"address[]\"},{\"name\":\"_genesisHash\",\"type\":\"bytes32\"}],\"name\":\"registerWrkChain\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_chainId\",\"type\":\"uint256\"}],\"name\":\"getNumBlocksSubmitted\",\"outputs\":[{\"name\":\"numBlocksSubmitted_\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_chainId\",\"type\":\"uint256\"},{\"name\":\"_height\",\"type\":\"uint256\"},{\"name\":\"_hash\",\"type\":\"bytes32\"},{\"name\":\"_parentHash\",\"type\":\"bytes32\"},{\"name\":\"_receiptRoot\",\"type\":\"bytes32\"},{\"name\":\"_txRoot\",\"type\":\"bytes32\"},{\"name\":\"_stateRoot\",\"type\":\"bytes32\"},{\"name\":\"_blockSigner\",\"type\":\"address\"}],\"name\":\"recordHeader\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_chainId\",\"type\":\"uint256\"}],\"name\":\"getAuthAddresses\",\"outputs\":[{\"name\":\"authAddressesIdx_\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_chainId\",\"type\":\"uint256\"}],\"name\":\"getGenesis\",\"outputs\":[{\"name\":\"genesisHash_\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_chainId\",\"type\":\"uint256\"}],\"name\":\"getLastRecordedBlockNum\",\"outputs\":[{\"name\":\"lastBlock_\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_wrkchainRegDeposit\",\"type\":\"uint256\"},{\"name\":\"_minBlocksForDepositReturn\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_chainId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_height\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_parentHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_receiptRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_txRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_stateRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_blockSigner\",\"type\":\"address\"}],\"name\":\"RecordHeader\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_chainId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_genesisHash\",\"type\":\"bytes32\"}],\"name\":\"RegisterWrkChain\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"LogFallbackFunctionCalled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_chainId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"WRKChainDepositRefund\",\"type\":\"event\"}]"

// WRKChainRootBin is the compiled bytecode used for deploying new contracts.
const WRKChainRootBin = `0x608060405234801561001057600080fd5b506040516040806113838339810160405280516020909101516000821161009857604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f4465706f73697420616d6f756e742073686f756c64206265203e203000000000604482015290519081900360640190fd5b6000811161010757604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f4d696e20626c6f636b732073686f756c64206265203e20300000000000000000604482015290519081900360640190fd5b6000919091556001556112648061011f6000396000f3006080604052600436106100825763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663244658f781146100be5780635be4e2281461011a5780635eefa77f14610134578063a78c49881461015e578063aa0d14de14610194578063cfdf3e84146101fc578063f4da14f614610214575b6040805133815234602082015281517fb1db2589d999730e8faec7dccf62c4d28a7f6ef1095b46c9a744dcd46d87130e929181900390910190a1005b3480156100ca57600080fd5b506040805160206004602480358281013584810280870186019097528086526101189684359636966044959194909101929182918501908490808284375094975061022c9650505050505050565b005b610118600480359060248035908101910135604435610595565b34801561014057600080fd5b5061014c6004356109d0565b60408051918252519081900360200190f35b34801561016a57600080fd5b5061011860043560243560443560643560843560a43560c435600160a060020a0360e43516610a86565b3480156101a057600080fd5b506101ac600435610e15565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156101e85781810151838201526020016101d0565b505050509050019250505060405180910390f35b34801561020857600080fd5b5061014c600435610f23565b34801561022057600080fd5b5061014c600435610fd9565b60008281526004602052604081206002015481908190859060ff16151561029d576040805160e560020a62461bcd02815260206004820152601660248201527f436861696e494420646f6573206e6f7420657869737400000000000000000000604482015290519081900360640190fd5b600081815260046020908152604080832033845260030190915290205460ff161515600114806102e65750600081815260046020526040902060050154600160a060020a031633145b15156102f157600080fd5b60008611610337576040805160e560020a62461bcd02815260206004820152601160248201526000805160206111f9833981519152604482015290519081900360640190fd5b8451600010610390576040805160e560020a62461bcd02815260206004820152601760248201527f4175746820616464726573736573207265717569726564000000000000000000604482015290519081900360640190fd5b845160001015610410576040805160e560020a62461bcd02815260206004820152602360248201527f4d6178207375626d697373696f6e2031302061646472657373657320616c6c6f60448201527f7765640000000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b60008681526004602052604090206002015460ff161515610469576040805160e560020a62461bcd0281526020600482015260176024820152600080516020611219833981519152604482015290519081900360640190fd5b6000868152600460205260408120945092505b60048401548310156104f357600060046000888152602001908152602001600020600301600086600401868154811015156104b357fe5b600091825260208083209190910154600160a060020a031683528201929092526040019020805460ff19169115159190911790556001929092019161047c565b600091505b8451821015610567576001600460008881526020019081526020016000206003016000878581518110151561052957fe5b602090810291909101810151600160a060020a03168252810191909152604001600020805460ff1916911515919091179055600191909101906104f8565b6000868152600460208181526040909220875161058c93919092019190880190611152565b50505050505050565b60008054341480156105a75750600034115b15156105fd576040805160e560020a62461bcd02815260206004820152601060248201527f4465706f73697420726571756972656400000000000000000000000000000000604482015290519081900360640190fd5b60008511610643576040805160e560020a62461bcd02815260206004820152601160248201526000805160206111f9833981519152604482015290519081900360640190fd5b600083116106c1576040805160e560020a62461bcd02815260206004820152602560248201527f496e697469616c20417574686f7269736564206164647265737365732072657160448201527f7569726564000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b600a831115610740576040805160e560020a62461bcd02815260206004820152602f60248201527f4d6178696d756d20313020496e697469616c20417574686f726973656420616460448201527f6472657373657320616c6c6f7765640000000000000000000000000000000000606482015290519081900360840190fd5b60008581526004602052604090206002015460ff16156107aa576040805160e560020a62461bcd02815260206004820152601760248201527f436861696e20494420616c726561647920657869737473000000000000000000604482015290519081900360640190fd5b60c060405190810160405280600081526020018360001916815260200160011515815260200185858080602002602001604051908101604052809392919081815260200183836020028082843750505092845250503360208084019190915260006040938401819052898152600480835290849020855181558583015160018201559385015160028501805460ff19169115159190911790556060850151805161085a9450918501920190611152565b50608082015160058201805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0390921691909117905560a0909101516006909101555060005b828110156109025760008581526004602052604081206001916003909101908686858181106108c957fe5b60209081029290920135600160a060020a0316835250810191909152604001600020805460ff191691151591909117905560010161089e565b600085815260046020818152604080842033808652600382018452918520805460ff191660019081179091558484529301805493840181558452922001805473ffffffffffffffffffffffffffffffffffffffff19169091179055600254610970903463ffffffff61108c16565b600255336000908152600360209081526040808320888452825291829020349055815187815290810184905281517fd49920535e715987d1fbf16411df1f3f3347d62e741006463c13550661ce4b69929181900390910190a15050505050565b6000808211610a17576040805160e560020a62461bcd02815260206004820152601160248201526000805160206111f9833981519152604482015290519081900360640190fd5b60008281526004602052604090206002015460ff161515610a70576040805160e560020a62461bcd0281526020600482015260176024820152600080516020611219833981519152604482015290519081900360640190fd5b5060009081526004602052604090206006015490565b600088815260046020526040812060020154899060ff161515610af3576040805160e560020a62461bcd02815260206004820152601660248201527f436861696e494420646f6573206e6f7420657869737400000000000000000000604482015290519081900360640190fd5b600081815260046020908152604080832033845260030190915290205460ff16151560011480610b3c5750600081815260046020526040902060050154600160a060020a031633145b1515610b4757600080fd5b60008a11610b8d576040805160e560020a62461bcd02815260206004820152601160248201526000805160206111f9833981519152604482015290519081900360640190fd5b60008a81526004602052604090206002015460ff161515610be6576040805160e560020a62461bcd0281526020600482015260176024820152600080516020611219833981519152604482015290519081900360640190fd5b60008a815260046020526040902080549092508911610c4f576040805160e560020a62461bcd02815260206004820152601d60248201527f426c6f636b2068656164657220616c7265616479207265636f72646564000000604482015290519081900360640190fd5b60008a815260046020526040902089815560060154610c7590600163ffffffff61108c16565b60008b8152600460205260409020600601819055600154148015610cc95750600080548b825260046020908152604080842060050154600160a060020a03168452600382528084208e855290915290912054145b15610d9d576005820154600160a060020a031660009081526003602090815260408083208d8452909152812081905554600254610d0b9163ffffffff6110f016565b6002556005820154600054604080518d8152600160a060020a03909316602084015282810191909152517f5fdfd7834e3e713189c97ffe1bc13c31b4e5003b84f2b73cdf2055cfe49bf6479181900360600190a1600582015460008054604051600160a060020a039093169281156108fc0292818181858888f19350505050158015610d9b573d6000803e3d6000fd5b505b604080518b8152602081018b90528082018a9052606081018990526080810188905260a0810187905260c08101869052600160a060020a03851660e082015290517f943621e8fb786306cd7b16a4467604107bbf3b0b04b536d39a245ea68c1062ef918190036101000190a150505050505050505050565b606060008211610e5d576040805160e560020a62461bcd02815260206004820152601160248201526000805160206111f9833981519152604482015290519081900360640190fd5b60008281526004602052604090206002015460ff161515610eb6576040805160e560020a62461bcd0281526020600482015260176024820152600080516020611219833981519152604482015290519081900360640190fd5b60008281526004602081815260409283902090910180548351818402810184019094528084529091830182828015610f1757602002820191906000526020600020905b8154600160a060020a03168152600190910190602001808311610ef9575b50505050509050919050565b6000808211610f6a576040805160e560020a62461bcd02815260206004820152601160248201526000805160206111f9833981519152604482015290519081900360640190fd5b60008281526004602052604090206002015460ff161515610fc3576040805160e560020a62461bcd0281526020600482015260176024820152600080516020611219833981519152604482015290519081900360640190fd5b5060009081526004602052604090206001015490565b6000808211611020576040805160e560020a62461bcd02815260206004820152601160248201526000805160206111f9833981519152604482015290519081900360640190fd5b60008281526004602052604090206002015460ff161515611079576040805160e560020a62461bcd0281526020600482015260176024820152600080516020611219833981519152604482015290519081900360640190fd5b5060009081526004602052604090205490565b6000828201838110156110e9576040805160e560020a62461bcd02815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b6000808383111561114b576040805160e560020a62461bcd02815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015290519081900360640190fd5b5050900390565b8280548282559060005260206000209081019282156111b4579160200282015b828111156111b4578251825473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03909116178255602090920191600190910190611172565b506111c09291506111c4565b5090565b6111f591905b808211156111c057805473ffffffffffffffffffffffffffffffffffffffff191681556001016111ca565b905600436861696e204944207265717569726564000000000000000000000000000000436861696e20494420646f6573206e6f74206578697374000000000000000000a165627a7a723058208ac53078cd772e45ae12079aa7c966b424e762098739432fd89b1368c1acd07c0029`

// DeployWRKChainRoot deploys a new Ethereum contract, binding an instance of WRKChainRoot to it.
func DeployWRKChainRoot(auth *bind.TransactOpts, backend bind.ContractBackend, _wrkchainRegDeposit *big.Int, _minBlocksForDepositReturn *big.Int) (common.Address, *types.Transaction, *WRKChainRoot, error) {
	parsed, err := abi.JSON(strings.NewReader(WRKChainRootABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(WRKChainRootBin), backend, _wrkchainRegDeposit, _minBlocksForDepositReturn)
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

// GetLastRecordedBlockNum is a free data retrieval call binding the contract method 0xf4da14f6.
//
// Solidity: function getLastRecordedBlockNum(uint256 _chainId) constant returns(uint256 lastBlock_)
func (_WRKChainRoot *WRKChainRootCaller) GetLastRecordedBlockNum(opts *bind.CallOpts, _chainId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _WRKChainRoot.contract.Call(opts, out, "getLastRecordedBlockNum", _chainId)
	return *ret0, err
}

// GetLastRecordedBlockNum is a free data retrieval call binding the contract method 0xf4da14f6.
//
// Solidity: function getLastRecordedBlockNum(uint256 _chainId) constant returns(uint256 lastBlock_)
func (_WRKChainRoot *WRKChainRootSession) GetLastRecordedBlockNum(_chainId *big.Int) (*big.Int, error) {
	return _WRKChainRoot.Contract.GetLastRecordedBlockNum(&_WRKChainRoot.CallOpts, _chainId)
}

// GetLastRecordedBlockNum is a free data retrieval call binding the contract method 0xf4da14f6.
//
// Solidity: function getLastRecordedBlockNum(uint256 _chainId) constant returns(uint256 lastBlock_)
func (_WRKChainRoot *WRKChainRootCallerSession) GetLastRecordedBlockNum(_chainId *big.Int) (*big.Int, error) {
	return _WRKChainRoot.Contract.GetLastRecordedBlockNum(&_WRKChainRoot.CallOpts, _chainId)
}

// GetNumBlocksSubmitted is a free data retrieval call binding the contract method 0x5eefa77f.
//
// Solidity: function getNumBlocksSubmitted(uint256 _chainId) constant returns(uint256 numBlocksSubmitted_)
func (_WRKChainRoot *WRKChainRootCaller) GetNumBlocksSubmitted(opts *bind.CallOpts, _chainId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _WRKChainRoot.contract.Call(opts, out, "getNumBlocksSubmitted", _chainId)
	return *ret0, err
}

// GetNumBlocksSubmitted is a free data retrieval call binding the contract method 0x5eefa77f.
//
// Solidity: function getNumBlocksSubmitted(uint256 _chainId) constant returns(uint256 numBlocksSubmitted_)
func (_WRKChainRoot *WRKChainRootSession) GetNumBlocksSubmitted(_chainId *big.Int) (*big.Int, error) {
	return _WRKChainRoot.Contract.GetNumBlocksSubmitted(&_WRKChainRoot.CallOpts, _chainId)
}

// GetNumBlocksSubmitted is a free data retrieval call binding the contract method 0x5eefa77f.
//
// Solidity: function getNumBlocksSubmitted(uint256 _chainId) constant returns(uint256 numBlocksSubmitted_)
func (_WRKChainRoot *WRKChainRootCallerSession) GetNumBlocksSubmitted(_chainId *big.Int) (*big.Int, error) {
	return _WRKChainRoot.Contract.GetNumBlocksSubmitted(&_WRKChainRoot.CallOpts, _chainId)
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

// WRKChainRootLogFallbackFunctionCalledIterator is returned from FilterLogFallbackFunctionCalled and is used to iterate over the raw logs and unpacked data for LogFallbackFunctionCalled events raised by the WRKChainRoot contract.
type WRKChainRootLogFallbackFunctionCalledIterator struct {
	Event *WRKChainRootLogFallbackFunctionCalled // Event containing the contract specifics and raw log

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
func (it *WRKChainRootLogFallbackFunctionCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WRKChainRootLogFallbackFunctionCalled)
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
		it.Event = new(WRKChainRootLogFallbackFunctionCalled)
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
func (it *WRKChainRootLogFallbackFunctionCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WRKChainRootLogFallbackFunctionCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WRKChainRootLogFallbackFunctionCalled represents a LogFallbackFunctionCalled event raised by the WRKChainRoot contract.
type WRKChainRootLogFallbackFunctionCalled struct {
	From   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLogFallbackFunctionCalled is a free log retrieval operation binding the contract event 0xb1db2589d999730e8faec7dccf62c4d28a7f6ef1095b46c9a744dcd46d87130e.
//
// Solidity: event LogFallbackFunctionCalled(address _from, uint256 _amount)
func (_WRKChainRoot *WRKChainRootFilterer) FilterLogFallbackFunctionCalled(opts *bind.FilterOpts) (*WRKChainRootLogFallbackFunctionCalledIterator, error) {

	logs, sub, err := _WRKChainRoot.contract.FilterLogs(opts, "LogFallbackFunctionCalled")
	if err != nil {
		return nil, err
	}
	return &WRKChainRootLogFallbackFunctionCalledIterator{contract: _WRKChainRoot.contract, event: "LogFallbackFunctionCalled", logs: logs, sub: sub}, nil
}

// WatchLogFallbackFunctionCalled is a free log subscription operation binding the contract event 0xb1db2589d999730e8faec7dccf62c4d28a7f6ef1095b46c9a744dcd46d87130e.
//
// Solidity: event LogFallbackFunctionCalled(address _from, uint256 _amount)
func (_WRKChainRoot *WRKChainRootFilterer) WatchLogFallbackFunctionCalled(opts *bind.WatchOpts, sink chan<- *WRKChainRootLogFallbackFunctionCalled) (event.Subscription, error) {

	logs, sub, err := _WRKChainRoot.contract.WatchLogs(opts, "LogFallbackFunctionCalled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WRKChainRootLogFallbackFunctionCalled)
				if err := _WRKChainRoot.contract.UnpackLog(event, "LogFallbackFunctionCalled", log); err != nil {
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

// WRKChainRootWRKChainDepositRefundIterator is returned from FilterWRKChainDepositRefund and is used to iterate over the raw logs and unpacked data for WRKChainDepositRefund events raised by the WRKChainRoot contract.
type WRKChainRootWRKChainDepositRefundIterator struct {
	Event *WRKChainRootWRKChainDepositRefund // Event containing the contract specifics and raw log

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
func (it *WRKChainRootWRKChainDepositRefundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WRKChainRootWRKChainDepositRefund)
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
		it.Event = new(WRKChainRootWRKChainDepositRefund)
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
func (it *WRKChainRootWRKChainDepositRefundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WRKChainRootWRKChainDepositRefundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WRKChainRootWRKChainDepositRefund represents a WRKChainDepositRefund event raised by the WRKChainRoot contract.
type WRKChainRootWRKChainDepositRefund struct {
	ChainId *big.Int
	Owner   common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWRKChainDepositRefund is a free log retrieval operation binding the contract event 0x5fdfd7834e3e713189c97ffe1bc13c31b4e5003b84f2b73cdf2055cfe49bf647.
//
// Solidity: event WRKChainDepositRefund(uint256 _chainId, address _owner, uint256 _amount)
func (_WRKChainRoot *WRKChainRootFilterer) FilterWRKChainDepositRefund(opts *bind.FilterOpts) (*WRKChainRootWRKChainDepositRefundIterator, error) {

	logs, sub, err := _WRKChainRoot.contract.FilterLogs(opts, "WRKChainDepositRefund")
	if err != nil {
		return nil, err
	}
	return &WRKChainRootWRKChainDepositRefundIterator{contract: _WRKChainRoot.contract, event: "WRKChainDepositRefund", logs: logs, sub: sub}, nil
}

// WatchWRKChainDepositRefund is a free log subscription operation binding the contract event 0x5fdfd7834e3e713189c97ffe1bc13c31b4e5003b84f2b73cdf2055cfe49bf647.
//
// Solidity: event WRKChainDepositRefund(uint256 _chainId, address _owner, uint256 _amount)
func (_WRKChainRoot *WRKChainRootFilterer) WatchWRKChainDepositRefund(opts *bind.WatchOpts, sink chan<- *WRKChainRootWRKChainDepositRefund) (event.Subscription, error) {

	logs, sub, err := _WRKChainRoot.contract.WatchLogs(opts, "WRKChainDepositRefund")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WRKChainRootWRKChainDepositRefund)
				if err := _WRKChainRoot.contract.UnpackLog(event, "WRKChainDepositRefund", log); err != nil {
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
