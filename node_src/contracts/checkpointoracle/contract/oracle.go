// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"errors"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
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
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
	_ = bind.Bind
)

// CheckpointOracleABI is the input ABI used to generate the binding from.
const CheckpointOracleABI = `[{"inputs":[{"internalType":"address[]","name":"_adminlist","type":"address[]"},{"internalType":"uint256","name":"_sectionSize","type":"uint256"},{"internalType":"uint256","name":"_processConfirms","type":"uint256"},{"internalType":"uint256","name":"_threshold","type":"uint256"},{"internalType":"uint256","name":"_chainId","type":"uint256"}],"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"uint64","name":"index","type":"uint64"},{"indexed":false,"internalType":"bytes32","name":"checkpointHash","type":"bytes32"},{"indexed":false,"internalType":"uint8","name":"v","type":"uint8"},{"indexed":false,"internalType":"bytes32","name":"r","type":"bytes32"},{"indexed":false,"internalType":"bytes32","name":"s","type":"bytes32"}],"name":"NewCheckpointVote","type":"event"},{"inputs":[],"name":"CHECKPOINT_TYPEHASH","outputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"DOMAIN_NAME","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"DOMAIN_SEPARATOR_TYPEHASH","outputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"DOMAIN_VERSION","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"GetAllAdmin","outputs":[{"internalType":"address[]","name":"","type":"address[]"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"GetLatestCheckpoint","outputs":[{"internalType":"uint64","name":"","type":"uint64"},{"internalType":"bytes32","name":"","type":"bytes32"},{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"bytes32","name":"_hash","type":"bytes32"},{"internalType":"uint64","name":"_sectionIndex","type":"uint64"},{"internalType":"uint256[]","name":"_nonces","type":"uint256[]"},{"internalType":"uint8[]","name":"v","type":"uint8[]"},{"internalType":"bytes32[]","name":"r","type":"bytes32[]"},{"internalType":"bytes32[]","name":"s","type":"bytes32[]"}],"name":"SetCheckpoint","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"","type":"address"}],"name":"adminNonces","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"chainId","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"admin","type":"address"}],"name":"getAdminNonce","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"}]`

// CheckpointOracleBin is the compiled bytecode used for deploying new contracts.
var CheckpointOracleBin = `608060405234801561001057600080fd5b50604051610d53380380610d53833981810160405260a081101561003357600080fd5b810190808051604051939291908464010000000082111561005357600080fd5b90830190602082018581111561006857600080fd5b825186602082028301116401000000008211171561008557600080fd5b82525081516020918201928201910280838360005b838110156100b257818101518382015260200161009a565b5050505091909101604090815260208301519083015160608401516080909401519195509350505060005b85518110156101895760016000808884815181106100f757fe5b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002060006101000a81548160ff021916908315150217905550600186828151811061014457fe5b60209081029190910181015182546001808201855560009485529290932090920180546001600160a01b0319166001600160a01b0390931692909217909155016100dd565b5060059390935560069190915560075560085550610ba7806101ac6000396000f3fe608060405234801561001057600080fd5b506004361061009e5760003560e01c806398b58e341161006657806398b58e34146101e95780639a8a05921461020f578063acb8cc4914610217578063e1193e4c1461021f578063ed4590b2146102275761009e565b806302f2987e146100a35780631db61b54146100db57806345848dfc146100e35780634d6a304c1461013b578063796f077b1461016c575b600080fd5b6100c9600480360360208110156100b957600080fd5b50356001600160a01b0316610478565b60408051918252519081900360200190f35b6100c9610493565b6100eb6104ae565b60408051602080825283518183015283519192839290830191858101910280838360005b8381101561012757818101518382015260200161010f565b505050509050019250505060405180910390f35b61014361054d565b6040805167ffffffffffffffff9094168452602084019290925282820152519081900360600190f35b610174610568565b6040805160208082528351818301528351919283929083019185019080838360005b838110156101ae578181015183820152602001610196565b50505050905090810190601f1680156101db5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6100c9600480360360208110156101ff57600080fd5b50356001600160a01b0316610594565b6100c96105a6565b6101746105ac565b6100c96105cb565b610464600480360360c081101561023d57600080fd5b81359167ffffffffffffffff60208201351691810190606081016040820135600160201b81111561026d57600080fd5b82018360208201111561027f57600080fd5b803590602001918460208302840111600160201b831117156102a057600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156102ef57600080fd5b82018360208201111561030157600080fd5b803590602001918460208302840111600160201b8311171561032257600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561037157600080fd5b82018360208201111561038357600080fd5b803590602001918460208302840111600160201b831117156103a457600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156103f357600080fd5b82018360208201111561040557600080fd5b803590602001918460208302840111600160201b8311171561042657600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295506105e6945050505050565b604080519115158252519081900360200190f35b6001600160a01b031660009081526009602052604090205490565b604051806052610b2082396052019050604051809103902081565b6060806001805490506040519080825280602002602001820160405280156104e0578160200160208202803883390190505b50905060005b60015481101561054757600181815481106104fd57fe5b9060005260206000200160009054906101000a90046001600160a01b031682828151811061052757fe5b6001600160a01b03909216602092830291909101909101526001016104e6565b50905090565b60025460045460035467ffffffffffffffff90921691909192565b6040518060400160405280601081526020016f436865636b706f696e744f7261636c6560801b81525081565b60096020526000908152604090205481565b60085481565b604051806040016040528060038152602001620312e360ec1b81525081565b60405180603a610ae68239603a019050604051809103902081565b3360009081526020819052604081205460ff1661060257600080fd5b825184511461061057600080fd5b815184511461061e57600080fd5b845184511461062c57600080fd5b6006546005548760010167ffffffffffffffff16020143101561065157506000610adb565b60025467ffffffffffffffff908116908716101561067157506000610adb565b60025467ffffffffffffffff87811691161480156106a3575067ffffffffffffffff86161515806106a3575060035415155b156106b057506000610adb565b866106bd57506000610adb565b60006040518080610b2060529139604080519182900360520182208282018252601083526f436865636b706f696e744f7261636c6560801b6020938401528151808301835260038152620312e360ec1b908401526008548251808501929092527f82c2167c8128d46b36868ca82fd80baa0fcb2e1a8bfd00518a0325c9d15c426a828401527fe6bbd6277e1bf288eed5e8d1780f9a50b239e86b153736bceebccf4ea79d90b3606083015260808201523060a0808301919091528251808303909101815260c09091019091528051910120915060009050805b8651811015610ad55760006040518080610ae6603a9139603a01905060405180910390208a8c8b85815181106107c857fe5b6020026020010151604051602001808581526020018467ffffffffffffffff1667ffffffffffffffff16815260200183815260200182815260200194505050505060405160208183030381529060405280519060200120905060008482604051602001808061190160f01b8152506002018381526020018281526020019250505060405160208183030381529060405280519060200120905060006001828b868151811061087257fe5b60200260200101518b878151811061088657fe5b60200260200101518b888151811061089a57fe5b602002602001015160405160008152602001604052604051808581526020018460ff1660ff1681526020018381526020018281526020019450505050506020604051602081039080840390855afa1580156108f9573d6000803e3d6000fd5b505060408051601f1901516001600160a01b03811660009081526020819052919091205490925060ff16905061092e57600080fd5b846001600160a01b0316816001600160a01b03161161094c57600080fd5b8a848151811061095857fe5b602002602001015160096000836001600160a01b03166001600160a01b0316815260200190815260200160002054146109c8576040805162461bcd60e51b815260206004820152600d60248201526c496e76616c6964206e6f6e636560981b604482015290519081900360640190fd5b6001600160a01b0381166000908152600960205260409020805460010190558951909450849067ffffffffffffffff8d16907fce51ffa16246bcaf0899f6504f473cd0114f430f566cef71ab7e03d3dde42a41908f908d9088908110610a2a57fe5b60200260200101518c8881518110610a3e57fe5b60200260200101518c8981518110610a5257fe5b6020026020010151604051808581526020018460ff1660ff16815260200183815260200182815260200194505050505060405180910390a26007548460010110610aca5750505060048a905550504360035550506002805467ffffffffffffffff191667ffffffffffffffff87161790556001610adb565b505050600101610796565b50600080fd5b969550505050505056fe436865636b706f696e742875696e7436342073656374696f6e496e6465782c6279746573333220686173682c75696e74323536206e6f6e636529454950373132446f6d61696e28737472696e67206e616d652c737472696e672076657273696f6e2c75696e7432353620636861696e49642c6164647265737320766572696679696e67436f6e747261637429a2646970667358221220a5062a3febd2a0f627cd304a80bbb3d7a411a335f2d6b05299f7840cec5fc3b564736f6c63430006000033`

// CheckpointOracleFuncSigs maps the 4-byte function signature to its string representation.
var CheckpointOracleFuncSigs = map[string]string{
	"45848dfc": "GetAllAdmin()",
	"4d6a304c": "GetLatestCheckpoint()",
	"d459fc46": "SetCheckpoint(bytes32,uint64,uint256[],uint8[],bytes32[],bytes32[])",
	"8d1fdf2f": "getAdminNonce(address)",
	"3850c7bd": "chainId()",
	"a0e67e2b": "DOMAIN_NAME()",
	"54fd4d50": "DOMAIN_VERSION()",
}

// CheckpointOracle is an auto generated Go binding around an Ethereum contract.
type CheckpointOracle struct {
	CheckpointOracleCaller     // Read-only binding to the contract
	CheckpointOracleTransactor // Write-only binding to the contract
	CheckpointOracleFilterer   // Log filterer for contract events
}

// CheckpointOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type CheckpointOracleCaller struct {
	contract *bind.BoundContract // Generic wrapper to the ABI methods
}

// CheckpointOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CheckpointOracleTransactor struct {
	contract *bind.BoundContract // Generic wrapper to the ABI methods
}

// CheckpointOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CheckpointOracleFilterer struct {
	contract *bind.BoundContract // Generic wrapper to the ABI methods
}

// CheckpointOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CheckpointOracleSession struct {
	Contract     *CheckpointOracle  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// CheckpointOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CheckpointOracleCallerSession struct {
	Contract *CheckpointOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// CheckpointOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CheckpointOracleTransactorSession struct {
	Contract     *CheckpointOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// CheckpointOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type CheckpointOracleRaw struct {
	Contract *CheckpointOracle // Generic contract binding to access the raw methods on
}

// CheckpointOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CheckpointOracleCallerRaw struct {
	Contract *CheckpointOracleCaller // Generic read-only contract binding to access the raw methods on
}

// CheckpointOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CheckpointOracleTransactorRaw struct {
	Contract *CheckpointOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCheckpointOracle is an auto generated write-only Go binding around an Ethereum contract.
func NewCheckpointOracle(address common.Address, backend bind.ContractBackend) (*CheckpointOracle, error) {
	contract, err := bindCheckpointOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CheckpointOracle{CheckpointOracleCaller: CheckpointOracleCaller{contract: contract}, CheckpointOracleTransactor: CheckpointOracleTransactor{contract: contract}, CheckpointOracleFilterer: CheckpointOracleFilterer{contract: contract}}, nil
}

// DeployCheckpointOracle deploys a new Ethereum contract, binding an instance of CheckpointOracle to it.
func DeployCheckpointOracle(auth *bind.TransactOpts, backend bind.ContractBackend, _adminlist []common.Address, _sectionSize *big.Int, _processConfirms *big.Int, _threshold *big.Int, _chainId *big.Int) (common.Address, *types.Transaction, *CheckpointOracle, error) {
	parsed, err := abi.JSON(strings.NewReader(CheckpointOracleABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	
	// Deploy the contract using bind.DeployContract
	address, tx, _, err := bind.DeployContract(auth, parsed, common.FromHex(CheckpointOracleBin), backend, _adminlist, _sectionSize, _processConfirms, _threshold, _chainId)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	
	// Create the CheckpointOracle instance
	contract, err := NewCheckpointOracle(address, backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	
	return address, tx, contract, nil
}

// NewCheckpointOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
func NewCheckpointOracleCaller(address common.Address, caller bind.ContractCaller) (*CheckpointOracleCaller, error) {
	contract, err := bindCheckpointOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CheckpointOracleCaller{contract: contract}, nil
}

// NewCheckpointOracleTransactor creates a new write-only instance of CheckpointOracle, bound to a specific deployed contract.
func NewCheckpointOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*CheckpointOracleTransactor, error) {
	contract, err := bindCheckpointOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CheckpointOracleTransactor{contract: contract}, nil
}

// NewCheckpointOracleFilterer creates a new log filterer instance of CheckpointOracle, bound to a specific deployed contract.
func NewCheckpointOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*CheckpointOracleFilterer, error) {
	contract, err := bindCheckpointOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CheckpointOracleFilterer{contract: contract}, nil
}

// bindCheckpointOracle binds a generic wrapper to an already deployed contract.
func bindCheckpointOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CheckpointOracleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CheckpointOracle *CheckpointOracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CheckpointOracle.Contract.CheckpointOracleTransactor.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one exists.
func (_CheckpointOracle *CheckpointOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CheckpointOracle.Contract.CheckpointOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CheckpointOracle *CheckpointOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CheckpointOracle.Contract.CheckpointOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CheckpointOracle *CheckpointOracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CheckpointOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one exists.
func (_CheckpointOracle *CheckpointOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CheckpointOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CheckpointOracle *CheckpointOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CheckpointOracle.Contract.contract.Transact(opts, method, params...)
}

// GetAllAdmin is a free data retrieval call binding the contract method 0x45848dfc.
func (_CheckpointOracle *CheckpointOracleCaller) GetAllAdmin(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _CheckpointOracle.contract.Call(opts, &out, "GetAllAdmin")
	if err != nil {
		return *new([]common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	return out0, err
}

// GetAllAdmin is a free data retrieval call binding the contract method 0x45848dfc.
func (_CheckpointOracle *CheckpointOracleSession) GetAllAdmin() ([]common.Address, error) {
	return _CheckpointOracle.Contract.GetAllAdmin(&_CheckpointOracle.CallOpts)
}

// GetAllAdmin is a free data retrieval call binding the contract method 0x45848dfc.
func (_CheckpointOracle *CheckpointOracleCallerSession) GetAllAdmin() ([]common.Address, error) {
	return _CheckpointOracle.Contract.GetAllAdmin(&_CheckpointOracle.CallOpts)
}

// GetLatestCheckpoint is a free data retrieval call binding the contract method 0x4d6a304c.
func (_CheckpointOracle *CheckpointOracleCaller) GetLatestCheckpoint(opts *bind.CallOpts) (uint64, [32]byte, *big.Int, error) {
	var out []interface{}
	err := _CheckpointOracle.contract.Call(opts, &out, "GetLatestCheckpoint")
	if err != nil {
		return *new(uint64), *new([32]byte), *new(*big.Int), err
	}
	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	out1 := *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	return out0, out1, out2, err
}

// GetLatestCheckpoint is a free data retrieval call binding the contract method 0x4d6a304c.
func (_CheckpointOracle *CheckpointOracleSession) GetLatestCheckpoint() (uint64, [32]byte, *big.Int, error) {
	return _CheckpointOracle.Contract.GetLatestCheckpoint(&_CheckpointOracle.CallOpts)
}

// GetLatestCheckpoint is a free data retrieval call binding the contract method 0x4d6a304c.
func (_CheckpointOracle *CheckpointOracleCallerSession) GetLatestCheckpoint() (uint64, [32]byte, *big.Int, error) {
	return _CheckpointOracle.Contract.GetLatestCheckpoint(&_CheckpointOracle.CallOpts)
}

// SetCheckpoint is a paid mutator transaction binding the contract method 0xd459fc46.
func (_CheckpointOracle *CheckpointOracleTransactor) SetCheckpoint(opts *bind.TransactOpts, _hash [32]byte, _sectionIndex uint64, _nonces []*big.Int, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _CheckpointOracle.contract.Transact(opts, "SetCheckpoint", _hash, _sectionIndex, _nonces, v, r, s)
}

// SetCheckpoint is a paid mutator transaction binding the contract method 0xd459fc46.
func (_CheckpointOracle *CheckpointOracleSession) SetCheckpoint(_hash [32]byte, _sectionIndex uint64, _nonces []*big.Int, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _CheckpointOracle.Contract.SetCheckpoint(&_CheckpointOracle.TransactOpts, _hash, _sectionIndex, _nonces, v, r, s)
}

// SetCheckpoint is a paid mutator transaction binding the contract method 0xd459fc46.
func (_CheckpointOracle *CheckpointOracleTransactorSession) SetCheckpoint(_hash [32]byte, _sectionIndex uint64, _nonces []*big.Int, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _CheckpointOracle.Contract.SetCheckpoint(&_CheckpointOracle.TransactOpts, _hash, _sectionIndex, _nonces, v, r, s)
}

// GetAdminNonce is a free data retrieval call binding the contract method 0x8d1fdf2f.
func (_CheckpointOracle *CheckpointOracleCaller) GetAdminNonce(opts *bind.CallOpts, admin common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CheckpointOracle.contract.Call(opts, &out, "getAdminNonce", admin)
	if err != nil {
		return *new(*big.Int), err
	}
	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	return out0, err
}

// GetAdminNonce is a free data retrieval call binding the contract method 0x8d1fdf2f.
func (_CheckpointOracle *CheckpointOracleSession) GetAdminNonce(admin common.Address) (*big.Int, error) {
	return _CheckpointOracle.Contract.GetAdminNonce(&_CheckpointOracle.CallOpts, admin)
}

// GetAdminNonce is a free data retrieval call binding the contract method 0x8d1fdf2f.
func (_CheckpointOracle *CheckpointOracleCallerSession) GetAdminNonce(admin common.Address) (*big.Int, error) {
	return _CheckpointOracle.Contract.GetAdminNonce(&_CheckpointOracle.CallOpts, admin)
}

// ChainId is a free data retrieval call binding the contract method 0x3850c7bd.
func (_CheckpointOracle *CheckpointOracleCaller) ChainId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CheckpointOracle.contract.Call(opts, &out, "chainId")
	if err != nil {
		return *new(*big.Int), err
	}
	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	return out0, err
}

// ChainId is a free data retrieval call binding the contract method 0x3850c7bd.
func (_CheckpointOracle *CheckpointOracleSession) ChainId() (*big.Int, error) {
	return _CheckpointOracle.Contract.ChainId(&_CheckpointOracle.CallOpts)
}

// ChainId is a free data retrieval call binding the contract method 0x3850c7bd.
func (_CheckpointOracle *CheckpointOracleCallerSession) ChainId() (*big.Int, error) {
	return _CheckpointOracle.Contract.ChainId(&_CheckpointOracle.CallOpts)
}

// AdminNonces is a free data retrieval call binding the contract method 0x8d1fdf2f.
func (_CheckpointOracle *CheckpointOracleCaller) AdminNonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CheckpointOracle.contract.Call(opts, &out, "adminNonces", arg0)
	if err != nil {
		return *new(*big.Int), err
	}
	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	return out0, err
}

// AdminNonces is a free data retrieval call binding the contract method 0x8d1fdf2f.
func (_CheckpointOracle *CheckpointOracleSession) AdminNonces(arg0 common.Address) (*big.Int, error) {
	return _CheckpointOracle.Contract.AdminNonces(&_CheckpointOracle.CallOpts, arg0)
}

// AdminNonces is a free data retrieval call binding the contract method 0x8d1fdf2f.
func (_CheckpointOracle *CheckpointOracleCallerSession) AdminNonces(arg0 common.Address) (*big.Int, error) {
	return _CheckpointOracle.Contract.AdminNonces(&_CheckpointOracle.CallOpts, arg0)
}

// CheckpointOracleNewCheckpointVoteIterator is returned from FilterNewCheckpointVote and is used to iterate over the raw logs and unpacked data for NewCheckpointVote events raised by the CheckpointOracle contract.
type CheckpointOracleNewCheckpointVoteIterator struct {
	Event *CheckpointOracleNewCheckpointVote // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for event logs to catch and remove them as they're received
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CheckpointOracleNewCheckpointVoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CheckpointOracleNewCheckpointVote)
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
		it.Event = new(CheckpointOracleNewCheckpointVote)
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
func (it *CheckpointOracleNewCheckpointVoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CheckpointOracleNewCheckpointVoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CheckpointOracleNewCheckpointVote represents a NewCheckpointVote event raised by the CheckpointOracle contract.
type CheckpointOracleNewCheckpointVote struct {
	Index          uint64
	CheckpointHash [32]byte
	V              uint8
	R              [32]byte
	S              [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNewCheckpointVote is a free log retrieval operation binding the contract event 0xce51ffa16246bcaf0899f6504f473cd0114f430f566cef71ab7e03d3dde42a4.
func (_CheckpointOracle *CheckpointOracleFilterer) FilterNewCheckpointVote(opts *bind.FilterOpts, index []uint64) (*CheckpointOracleNewCheckpointVoteIterator, error) {

	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _CheckpointOracle.contract.FilterLogs(opts, "NewCheckpointVote", indexRule)
	if err != nil {
		return nil, err
	}
	return &CheckpointOracleNewCheckpointVoteIterator{contract: _CheckpointOracle.contract, event: "NewCheckpointVote", logs: logs, sub: sub}, nil
}

// WatchNewCheckpointVote is a free log subscription operation binding the contract event 0xce51ffa16246bcaf0899f6504f473cd0114f430f566cef71ab7e03d3dde42a4.
func (_CheckpointOracle *CheckpointOracleFilterer) WatchNewCheckpointVote(opts *bind.WatchOpts, sink chan<- *CheckpointOracleNewCheckpointVote, index []uint64) (event.Subscription, error) {

	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _CheckpointOracle.contract.WatchLogs(opts, "NewCheckpointVote", indexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// Dump chain reading and rescaling the big.Int parameters
				val := new(CheckpointOracleNewCheckpointVote)
				if err := _CheckpointOracle.contract.UnpackLog(val, "NewCheckpointVote", log); err != nil {
					return err
				}
				val.Raw = log
				select {
				case sink <- val:
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

// ParseNewCheckpointVote is a log parse operation binding the contract event 0xce51ffa16246bcaf0899f6504f473cd0114f430f566cef71ab7e03d3dde42a4.
func (_CheckpointOracle *CheckpointOracleFilterer) ParseNewCheckpointVote(log types.Log) (*CheckpointOracleNewCheckpointVote, error) {
	event := new(CheckpointOracleNewCheckpointVote)
	if err := _CheckpointOracle.contract.UnpackLog(event, "NewCheckpointVote", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ParseNewCheckpointVote is a log parse operation binding the contract event 0xce51ffa16246bcaf0899f6504f473cd0114f430f566cef71ab7e03d3dde42a4.
func ParseNewCheckpointVote(log types.Log) (*CheckpointOracleNewCheckpointVote, error) {
	parsed, err := abi.JSON(strings.NewReader(CheckpointOracleABI))
	if err != nil {
		return nil, err
	}
	event := new(CheckpointOracleNewCheckpointVote)
	if err := parsed.UnpackIntoInterface(event, "NewCheckpointVote", log.Data); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
