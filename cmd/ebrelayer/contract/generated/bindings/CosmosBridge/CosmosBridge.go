// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package CosmosBridge

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
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// CosmosBridgeABI is the input ABI used to generate the binding from.
const CosmosBridgeABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"bridgeBank\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_prophecyID\",\"type\":\"uint256\"}],\"name\":\"isProphecyClaimValidatorActive\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"operator\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"hasBridgeBank\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_prophecyID\",\"type\":\"uint256\"}],\"name\":\"completeProphecyClaim\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_oracle\",\"type\":\"address\"}],\"name\":\"setOracle\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"oracle\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"valset\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_bridgeBank\",\"type\":\"address\"}],\"name\":\"setBridgeBank\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"prophecyClaimCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_claimType\",\"type\":\"uint8\"},{\"name\":\"_cosmosSender\",\"type\":\"bytes\"},{\"name\":\"_ethereumReceiver\",\"type\":\"address\"},{\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"name\":\"_symbol\",\"type\":\"string\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"newProphecyClaim\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_prophecyID\",\"type\":\"uint256\"}],\"name\":\"isProphecyClaimActive\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"prophecyClaims\",\"outputs\":[{\"name\":\"claimType\",\"type\":\"uint8\"},{\"name\":\"cosmosSender\",\"type\":\"bytes\"},{\"name\":\"ethereumReceiver\",\"type\":\"address\"},{\"name\":\"originalValidator\",\"type\":\"address\"},{\"name\":\"tokenAddress\",\"type\":\"address\"},{\"name\":\"symbol\",\"type\":\"string\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"hasOracle\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\"},{\"name\":\"_valset\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_oracle\",\"type\":\"address\"}],\"name\":\"LogOracleSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_bridgeBank\",\"type\":\"address\"}],\"name\":\"LogBridgeBankSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_prophecyID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_claimType\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"_cosmosSender\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"_ethereumReceiver\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_validatorAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_symbol\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"LogNewProphecyClaim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_prophecyID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_claimType\",\"type\":\"uint8\"}],\"name\":\"LogProphecyCompleted\",\"type\":\"event\"}]"

// CosmosBridgeBin is the compiled bytecode used for deploying new contracts.
var CosmosBridgeBin = "0x608060405234801561001057600080fd5b5060405160408061269f8339810180604052604081101561003057600080fd5b8101908080519060200190929190805190602001909291905050506000600481905550816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000600260146101000a81548160ff0219169083151502179055506000600360146101000a81548160ff02191690831515021790555050506125848061011b6000396000f3fe6080604052600436106100d0576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680630e41f373146100d5578063529f3dd21461012c578063570ca7351461017f57806369294a4e146101d65780636b3ce98c146102055780637adbf973146102405780637dc0d1d0146102915780637f54af0c146102e8578063814c92c31461033f5780638ea5352d146103905780639408ee9c146103bb578063d8da69ea14610571578063db4237af146105c4578063fb7831f2146107ae575b600080fd5b3480156100e157600080fd5b506100ea6107dd565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561013857600080fd5b506101656004803603602081101561014f57600080fd5b8101908080359060200190929190505050610803565b604051808215151515815260200191505060405180910390f35b34801561018b57600080fd5b50610194610938565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156101e257600080fd5b506101eb61095d565b604051808215151515815260200191505060405180910390f35b34801561021157600080fd5b5061023e6004803603602081101561022857600080fd5b8101908080359060200190929190505050610970565b005b34801561024c57600080fd5b5061028f6004803603602081101561026357600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610bc2565b005b34801561029d57600080fd5b506102a6610e15565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156102f457600080fd5b506102fd610e3b565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561034b57600080fd5b5061038e6004803603602081101561036257600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610e61565b005b34801561039c57600080fd5b506103a56110b4565b6040518082815260200191505060405180910390f35b3480156103c757600080fd5b5061056f600480360360c08110156103de57600080fd5b81019080803560ff1690602001909291908035906020019064010000000081111561040857600080fd5b82018360208201111561041a57600080fd5b8035906020019184600183028401116401000000008311171561043c57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001906401000000008111156104df57600080fd5b8201836020820111156104f157600080fd5b8035906020019184600183028401116401000000008311171561051357600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290803590602001909291905050506110ba565b005b34801561057d57600080fd5b506105aa6004803603602081101561059457600080fd5b8101908080359060200190929190505050611783565b604051808215151515815260200191505060405180910390f35b3480156105d057600080fd5b506105fd600480360360208110156105e757600080fd5b81019080803590602001909291905050506117c9565b6040518089600281111561060d57fe5b60ff168152602001806020018873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001806020018581526020018460038111156106c557fe5b60ff16815260200183810383528a818151815260200191508051906020019080838360005b838110156107055780820151818401526020810190506106ea565b50505050905090810190601f1680156107325780820380516001836020036101000a031916815260200191505b50838103825286818151815260200191508051906020019080838360005b8381101561076b578082015181840152602081019050610750565b50505050905090810190601f1680156107985780820380516001836020036101000a031916815260200191505b509a505050505050505050505060405180910390f35b3480156107ba57600080fd5b506107c36119bb565b604051808215151515815260200191505060405180910390f35b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166340550a1c6005600085815260200190815260200160002060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b1580156108f657600080fd5b505afa15801561090a573d6000803e3d6000fd5b505050506040513d602081101561092057600080fd5b81019080805190602001909291905050509050919050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600360149054906101000a900460ff1681565b8061097a81611783565b15156109ee576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601c8152602001807f50726f706865637920636c61696d206973206e6f74206163746976650000000081525060200191505060405180910390fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610ad9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260278152602001807f4f6e6c7920746865204f7261636c65206d617920636f6d706c6574652070726f81526020017f706865636965730000000000000000000000000000000000000000000000000081525060400191505060405180910390fd5b60026005600084815260200190815260200160002060070160006101000a81548160ff02191690836003811115610b0c57fe5b021790555060006005600084815260200190815260200160002060000160009054906101000a900460ff16905060016002811115610b4657fe5b816002811115610b5257fe5b1415610b6657610b61836119ce565b610b70565b610b6f83611e34565b5b7f79e7c1c0bd54f11809c3bf6023c242783602d61ceff272c6bba6f8559c24ad0d838260405180838152602001826002811115610ba957fe5b60ff1681526020019250505060405180910390a1505050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610c86576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260158152602001807f4d75737420626520746865206f70657261746f722e000000000000000000000081525060200191505060405180910390fd5b600260149054906101000a900460ff16151515610d31576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260318152602001807f546865204f7261636c652063616e6e6f742062652075706461746564206f6e6381526020017f6520697420686173206265656e2073657400000000000000000000000000000081525060400191505060405180910390fd5b6001600260146101000a81548160ff02191690831515021790555080600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f6efb0434342713e2e9b1501dbebf76b4ed18406ea77ab5d56535cc26dec3adc0600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a150565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610f25576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260158152602001807f4d75737420626520746865206f70657261746f722e000000000000000000000081525060200191505060405180910390fd5b600360149054906101000a900460ff16151515610fd0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260368152602001807f546865204272696467652042616e6b2063616e6e6f742062652075706461746581526020017f64206f6e636520697420686173206265656e207365740000000000000000000081525060400191505060405180910390fd5b6001600360146101000a81548160ff02191690831515021790555080600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507fc8b65043fb196ac032b79a435397d1d14a96b4e9d12e366c3b1f550cb01d2dfa600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a150565b60045481565b60011515600260149054906101000a900460ff1615151480156110f0575060011515600360149054906101000a900460ff161515145b15156111b0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260468152602001807f546865204f70657261746f72206d7573742073657420746865206f7261636c6581526020017f20616e64206272696467652062616e6b20666f7220627269646765206163746981526020017f766174696f6e000000000000000000000000000000000000000000000000000081525060600191505060405180910390fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166340550a1c336040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b15801561126b57600080fd5b505afa15801561127f573d6000803e3d6000fd5b505050506040513d602081101561129557600080fd5b8101908080519060200190929190505050151561131a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f4d75737420626520616e206163746976652076616c696461746f72000000000081525060200191505060405180910390fd5b611330600160045461230b90919063ffffffff16565b600481905550600033905060006001600281111561134a57fe5b88600281111561135657fe5b14156113655760019050611389565b60028081111561137157fe5b88600281111561137d57fe5b141561138857600290505b5b611391612395565b610100604051908101604052808360028111156113aa57fe5b81526020018981526020018873ffffffffffffffffffffffffffffffffffffffff1681526020018473ffffffffffffffffffffffffffffffffffffffff1681526020018773ffffffffffffffffffffffffffffffffffffffff1681526020018681526020018581526020016001600381111561142257fe5b81525090508060056000600454815260200190815260200160002060008201518160000160006101000a81548160ff0219169083600281111561146157fe5b02179055506020820151816001019080519060200190611482929190612433565b5060408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060808201518160040160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060a08201518160050190805190602001906115749291906124b3565b5060c0820151816006015560e08201518160070160006101000a81548160ff021916908360038111156115a357fe5b02179055509050507f4c4b04a2b190e6bb01b6243f150fc76174861acd19cf98841801baaff5262dd8600454838a8a878b8b8b604051808981526020018860028111156115ec57fe5b60ff168152602001806020018773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200180602001848152602001838103835289818151815260200191508051906020019080838360005b838110156116d05780820151818401526020810190506116b5565b50505050905090810190601f1680156116fd5780820380516001836020036101000a031916815260200191505b50838103825285818151815260200191508051906020019080838360005b8381101561173657808201518184015260208101905061171b565b50505050905090810190601f1680156117635780820380516001836020036101000a031916815260200191505b509a505050505050505050505060405180910390a1505050505050505050565b60006001600381111561179257fe5b6005600084815260200190815260200160002060070160009054906101000a900460ff1660038111156117c157fe5b149050919050565b60056020528060005260406000206000915090508060000160009054906101000a900460ff1690806001018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156118885780601f1061185d57610100808354040283529160200191611888565b820191906000526020600020905b81548152906001019060200180831161186b57829003601f168201915b5050505050908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806005018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156119985780601f1061196d57610100808354040283529160200191611998565b820191906000526020600020905b81548152906001019060200180831161197b57829003601f168201915b5050505050908060060154908060070160009054906101000a900460ff16905088565b600260149054906101000a900460ff1681565b6119d6612395565b6005600083815260200190815260200160002061010060405190810160405290816000820160009054906101000a900460ff166002811115611a1457fe5b6002811115611a1f57fe5b8152602001600182018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015611abc5780601f10611a9157610100808354040283529160200191611abc565b820191906000526020600020905b815481529060010190602001808311611a9f57829003601f168201915b505050505081526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016004820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600582018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015611c605780601f10611c3557610100808354040283529160200191611c60565b820191906000526020600020905b815481529060010190602001808311611c4357829003601f168201915b50505050508152602001600682015481526020016007820160009054906101000a900460ff166003811115611c9157fe5b6003811115611c9c57fe5b815250509050600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632f89c91c826040015183608001518460a001518560c001516040518563ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200180602001838152602001828103825284818151815260200191508051906020019080838360005b83811015611dc9578082015181840152602081019050611dae565b50505050905090810190601f168015611df65780820380516001836020036101000a031916815260200191505b5095505050505050600060405180830381600087803b158015611e1857600080fd5b505af1158015611e2c573d6000803e3d6000fd5b505050505050565b611e3c612395565b6005600083815260200190815260200160002061010060405190810160405290816000820160009054906101000a900460ff166002811115611e7a57fe5b6002811115611e8557fe5b8152602001600182018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015611f225780601f10611ef757610100808354040283529160200191611f22565b820191906000526020600020905b815481529060010190602001808311611f0557829003601f168201915b505050505081526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016004820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600582018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156120c65780601f1061209b576101008083540402835291602001916120c6565b820191906000526020600020905b8154815290600101906020018083116120a957829003601f168201915b50505050508152602001600682015481526020016007820160009054906101000a900460ff1660038111156120f757fe5b600381111561210257fe5b815250509050600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663cdf68c418260200151836040015184608001518560a001518660c001516040518663ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180806020018673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200180602001848152602001838103835288818151815260200191508051906020019080838360005b8381101561223857808201518184015260208101905061221d565b50505050905090810190601f1680156122655780820380516001836020036101000a031916815260200191505b50838103825285818151815260200191508051906020019080838360005b8381101561229e578082015181840152602081019050612283565b50505050905090810190601f1680156122cb5780820380516001836020036101000a031916815260200191505b50975050505050505050600060405180830381600087803b1580156122ef57600080fd5b505af1158015612303573d6000803e3d6000fd5b505050505050565b600080828401905083811015151561238b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f536166654d6174683a206164646974696f6e206f766572666c6f77000000000081525060200191505060405180910390fd5b8091505092915050565b61010060405190810160405280600060028111156123af57fe5b815260200160608152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff16815260200160608152602001600081526020016000600381111561242d57fe5b81525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061247457805160ff19168380011785556124a2565b828001600101855582156124a2579182015b828111156124a1578251825591602001919060010190612486565b5b5090506124af9190612533565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106124f457805160ff1916838001178555612522565b82800160010185558215612522579182015b82811115612521578251825591602001919060010190612506565b5b50905061252f9190612533565b5090565b61255591905b80821115612551576000816000905550600101612539565b5090565b9056fea165627a7a72305820145d982a8aa79543ba6ef54e331edc61265bb801ca9f8d6d0be6ce7b03504cef0029"

// DeployCosmosBridge deploys a new Ethereum contract, binding an instance of CosmosBridge to it.
func DeployCosmosBridge(auth *bind.TransactOpts, backend bind.ContractBackend, _operator common.Address, _valset common.Address) (common.Address, *types.Transaction, *CosmosBridge, error) {
	parsed, err := abi.JSON(strings.NewReader(CosmosBridgeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CosmosBridgeBin), backend, _operator, _valset)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CosmosBridge{CosmosBridgeCaller: CosmosBridgeCaller{contract: contract}, CosmosBridgeTransactor: CosmosBridgeTransactor{contract: contract}, CosmosBridgeFilterer: CosmosBridgeFilterer{contract: contract}}, nil
}

// CosmosBridge is an auto generated Go binding around an Ethereum contract.
type CosmosBridge struct {
	CosmosBridgeCaller     // Read-only binding to the contract
	CosmosBridgeTransactor // Write-only binding to the contract
	CosmosBridgeFilterer   // Log filterer for contract events
}

// CosmosBridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type CosmosBridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CosmosBridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CosmosBridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CosmosBridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CosmosBridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CosmosBridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CosmosBridgeSession struct {
	Contract     *CosmosBridge     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CosmosBridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CosmosBridgeCallerSession struct {
	Contract *CosmosBridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// CosmosBridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CosmosBridgeTransactorSession struct {
	Contract     *CosmosBridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// CosmosBridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type CosmosBridgeRaw struct {
	Contract *CosmosBridge // Generic contract binding to access the raw methods on
}

// CosmosBridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CosmosBridgeCallerRaw struct {
	Contract *CosmosBridgeCaller // Generic read-only contract binding to access the raw methods on
}

// CosmosBridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CosmosBridgeTransactorRaw struct {
	Contract *CosmosBridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCosmosBridge creates a new instance of CosmosBridge, bound to a specific deployed contract.
func NewCosmosBridge(address common.Address, backend bind.ContractBackend) (*CosmosBridge, error) {
	contract, err := bindCosmosBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CosmosBridge{CosmosBridgeCaller: CosmosBridgeCaller{contract: contract}, CosmosBridgeTransactor: CosmosBridgeTransactor{contract: contract}, CosmosBridgeFilterer: CosmosBridgeFilterer{contract: contract}}, nil
}

// NewCosmosBridgeCaller creates a new read-only instance of CosmosBridge, bound to a specific deployed contract.
func NewCosmosBridgeCaller(address common.Address, caller bind.ContractCaller) (*CosmosBridgeCaller, error) {
	contract, err := bindCosmosBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CosmosBridgeCaller{contract: contract}, nil
}

// NewCosmosBridgeTransactor creates a new write-only instance of CosmosBridge, bound to a specific deployed contract.
func NewCosmosBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*CosmosBridgeTransactor, error) {
	contract, err := bindCosmosBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CosmosBridgeTransactor{contract: contract}, nil
}

// NewCosmosBridgeFilterer creates a new log filterer instance of CosmosBridge, bound to a specific deployed contract.
func NewCosmosBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*CosmosBridgeFilterer, error) {
	contract, err := bindCosmosBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CosmosBridgeFilterer{contract: contract}, nil
}

// bindCosmosBridge binds a generic wrapper to an already deployed contract.
func bindCosmosBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CosmosBridgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CosmosBridge *CosmosBridgeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CosmosBridge.Contract.CosmosBridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CosmosBridge *CosmosBridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CosmosBridge.Contract.CosmosBridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CosmosBridge *CosmosBridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CosmosBridge.Contract.CosmosBridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CosmosBridge *CosmosBridgeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CosmosBridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CosmosBridge *CosmosBridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CosmosBridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CosmosBridge *CosmosBridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CosmosBridge.Contract.contract.Transact(opts, method, params...)
}

// BridgeBank is a free data retrieval call binding the contract method 0x0e41f373.
//
// Solidity: function bridgeBank() view returns(address)
func (_CosmosBridge *CosmosBridgeCaller) BridgeBank(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CosmosBridge.contract.Call(opts, out, "bridgeBank")
	return *ret0, err
}

// BridgeBank is a free data retrieval call binding the contract method 0x0e41f373.
//
// Solidity: function bridgeBank() view returns(address)
func (_CosmosBridge *CosmosBridgeSession) BridgeBank() (common.Address, error) {
	return _CosmosBridge.Contract.BridgeBank(&_CosmosBridge.CallOpts)
}

// BridgeBank is a free data retrieval call binding the contract method 0x0e41f373.
//
// Solidity: function bridgeBank() view returns(address)
func (_CosmosBridge *CosmosBridgeCallerSession) BridgeBank() (common.Address, error) {
	return _CosmosBridge.Contract.BridgeBank(&_CosmosBridge.CallOpts)
}

// HasBridgeBank is a free data retrieval call binding the contract method 0x69294a4e.
//
// Solidity: function hasBridgeBank() view returns(bool)
func (_CosmosBridge *CosmosBridgeCaller) HasBridgeBank(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _CosmosBridge.contract.Call(opts, out, "hasBridgeBank")
	return *ret0, err
}

// HasBridgeBank is a free data retrieval call binding the contract method 0x69294a4e.
//
// Solidity: function hasBridgeBank() view returns(bool)
func (_CosmosBridge *CosmosBridgeSession) HasBridgeBank() (bool, error) {
	return _CosmosBridge.Contract.HasBridgeBank(&_CosmosBridge.CallOpts)
}

// HasBridgeBank is a free data retrieval call binding the contract method 0x69294a4e.
//
// Solidity: function hasBridgeBank() view returns(bool)
func (_CosmosBridge *CosmosBridgeCallerSession) HasBridgeBank() (bool, error) {
	return _CosmosBridge.Contract.HasBridgeBank(&_CosmosBridge.CallOpts)
}

// HasOracle is a free data retrieval call binding the contract method 0xfb7831f2.
//
// Solidity: function hasOracle() view returns(bool)
func (_CosmosBridge *CosmosBridgeCaller) HasOracle(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _CosmosBridge.contract.Call(opts, out, "hasOracle")
	return *ret0, err
}

// HasOracle is a free data retrieval call binding the contract method 0xfb7831f2.
//
// Solidity: function hasOracle() view returns(bool)
func (_CosmosBridge *CosmosBridgeSession) HasOracle() (bool, error) {
	return _CosmosBridge.Contract.HasOracle(&_CosmosBridge.CallOpts)
}

// HasOracle is a free data retrieval call binding the contract method 0xfb7831f2.
//
// Solidity: function hasOracle() view returns(bool)
func (_CosmosBridge *CosmosBridgeCallerSession) HasOracle() (bool, error) {
	return _CosmosBridge.Contract.HasOracle(&_CosmosBridge.CallOpts)
}

// IsProphecyClaimActive is a free data retrieval call binding the contract method 0xd8da69ea.
//
// Solidity: function isProphecyClaimActive(uint256 _prophecyID) view returns(bool)
func (_CosmosBridge *CosmosBridgeCaller) IsProphecyClaimActive(opts *bind.CallOpts, _prophecyID *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _CosmosBridge.contract.Call(opts, out, "isProphecyClaimActive", _prophecyID)
	return *ret0, err
}

// IsProphecyClaimActive is a free data retrieval call binding the contract method 0xd8da69ea.
//
// Solidity: function isProphecyClaimActive(uint256 _prophecyID) view returns(bool)
func (_CosmosBridge *CosmosBridgeSession) IsProphecyClaimActive(_prophecyID *big.Int) (bool, error) {
	return _CosmosBridge.Contract.IsProphecyClaimActive(&_CosmosBridge.CallOpts, _prophecyID)
}

// IsProphecyClaimActive is a free data retrieval call binding the contract method 0xd8da69ea.
//
// Solidity: function isProphecyClaimActive(uint256 _prophecyID) view returns(bool)
func (_CosmosBridge *CosmosBridgeCallerSession) IsProphecyClaimActive(_prophecyID *big.Int) (bool, error) {
	return _CosmosBridge.Contract.IsProphecyClaimActive(&_CosmosBridge.CallOpts, _prophecyID)
}

// IsProphecyClaimValidatorActive is a free data retrieval call binding the contract method 0x529f3dd2.
//
// Solidity: function isProphecyClaimValidatorActive(uint256 _prophecyID) view returns(bool)
func (_CosmosBridge *CosmosBridgeCaller) IsProphecyClaimValidatorActive(opts *bind.CallOpts, _prophecyID *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _CosmosBridge.contract.Call(opts, out, "isProphecyClaimValidatorActive", _prophecyID)
	return *ret0, err
}

// IsProphecyClaimValidatorActive is a free data retrieval call binding the contract method 0x529f3dd2.
//
// Solidity: function isProphecyClaimValidatorActive(uint256 _prophecyID) view returns(bool)
func (_CosmosBridge *CosmosBridgeSession) IsProphecyClaimValidatorActive(_prophecyID *big.Int) (bool, error) {
	return _CosmosBridge.Contract.IsProphecyClaimValidatorActive(&_CosmosBridge.CallOpts, _prophecyID)
}

// IsProphecyClaimValidatorActive is a free data retrieval call binding the contract method 0x529f3dd2.
//
// Solidity: function isProphecyClaimValidatorActive(uint256 _prophecyID) view returns(bool)
func (_CosmosBridge *CosmosBridgeCallerSession) IsProphecyClaimValidatorActive(_prophecyID *big.Int) (bool, error) {
	return _CosmosBridge.Contract.IsProphecyClaimValidatorActive(&_CosmosBridge.CallOpts, _prophecyID)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_CosmosBridge *CosmosBridgeCaller) Operator(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CosmosBridge.contract.Call(opts, out, "operator")
	return *ret0, err
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_CosmosBridge *CosmosBridgeSession) Operator() (common.Address, error) {
	return _CosmosBridge.Contract.Operator(&_CosmosBridge.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_CosmosBridge *CosmosBridgeCallerSession) Operator() (common.Address, error) {
	return _CosmosBridge.Contract.Operator(&_CosmosBridge.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_CosmosBridge *CosmosBridgeCaller) Oracle(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CosmosBridge.contract.Call(opts, out, "oracle")
	return *ret0, err
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_CosmosBridge *CosmosBridgeSession) Oracle() (common.Address, error) {
	return _CosmosBridge.Contract.Oracle(&_CosmosBridge.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_CosmosBridge *CosmosBridgeCallerSession) Oracle() (common.Address, error) {
	return _CosmosBridge.Contract.Oracle(&_CosmosBridge.CallOpts)
}

// ProphecyClaimCount is a free data retrieval call binding the contract method 0x8ea5352d.
//
// Solidity: function prophecyClaimCount() view returns(uint256)
func (_CosmosBridge *CosmosBridgeCaller) ProphecyClaimCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CosmosBridge.contract.Call(opts, out, "prophecyClaimCount")
	return *ret0, err
}

// ProphecyClaimCount is a free data retrieval call binding the contract method 0x8ea5352d.
//
// Solidity: function prophecyClaimCount() view returns(uint256)
func (_CosmosBridge *CosmosBridgeSession) ProphecyClaimCount() (*big.Int, error) {
	return _CosmosBridge.Contract.ProphecyClaimCount(&_CosmosBridge.CallOpts)
}

// ProphecyClaimCount is a free data retrieval call binding the contract method 0x8ea5352d.
//
// Solidity: function prophecyClaimCount() view returns(uint256)
func (_CosmosBridge *CosmosBridgeCallerSession) ProphecyClaimCount() (*big.Int, error) {
	return _CosmosBridge.Contract.ProphecyClaimCount(&_CosmosBridge.CallOpts)
}

// ProphecyClaims is a free data retrieval call binding the contract method 0xdb4237af.
//
// Solidity: function prophecyClaims(uint256 ) view returns(uint8 claimType, bytes cosmosSender, address ethereumReceiver, address originalValidator, address tokenAddress, string symbol, uint256 amount, uint8 status)
func (_CosmosBridge *CosmosBridgeCaller) ProphecyClaims(opts *bind.CallOpts, arg0 *big.Int) (struct {
	ClaimType         uint8
	CosmosSender      []byte
	EthereumReceiver  common.Address
	OriginalValidator common.Address
	TokenAddress      common.Address
	Symbol            string
	Amount            *big.Int
	Status            uint8
}, error) {
	ret := new(struct {
		ClaimType         uint8
		CosmosSender      []byte
		EthereumReceiver  common.Address
		OriginalValidator common.Address
		TokenAddress      common.Address
		Symbol            string
		Amount            *big.Int
		Status            uint8
	})
	out := ret
	err := _CosmosBridge.contract.Call(opts, out, "prophecyClaims", arg0)
	return *ret, err
}

// ProphecyClaims is a free data retrieval call binding the contract method 0xdb4237af.
//
// Solidity: function prophecyClaims(uint256 ) view returns(uint8 claimType, bytes cosmosSender, address ethereumReceiver, address originalValidator, address tokenAddress, string symbol, uint256 amount, uint8 status)
func (_CosmosBridge *CosmosBridgeSession) ProphecyClaims(arg0 *big.Int) (struct {
	ClaimType         uint8
	CosmosSender      []byte
	EthereumReceiver  common.Address
	OriginalValidator common.Address
	TokenAddress      common.Address
	Symbol            string
	Amount            *big.Int
	Status            uint8
}, error) {
	return _CosmosBridge.Contract.ProphecyClaims(&_CosmosBridge.CallOpts, arg0)
}

// ProphecyClaims is a free data retrieval call binding the contract method 0xdb4237af.
//
// Solidity: function prophecyClaims(uint256 ) view returns(uint8 claimType, bytes cosmosSender, address ethereumReceiver, address originalValidator, address tokenAddress, string symbol, uint256 amount, uint8 status)
func (_CosmosBridge *CosmosBridgeCallerSession) ProphecyClaims(arg0 *big.Int) (struct {
	ClaimType         uint8
	CosmosSender      []byte
	EthereumReceiver  common.Address
	OriginalValidator common.Address
	TokenAddress      common.Address
	Symbol            string
	Amount            *big.Int
	Status            uint8
}, error) {
	return _CosmosBridge.Contract.ProphecyClaims(&_CosmosBridge.CallOpts, arg0)
}

// Valset is a free data retrieval call binding the contract method 0x7f54af0c.
//
// Solidity: function valset() view returns(address)
func (_CosmosBridge *CosmosBridgeCaller) Valset(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CosmosBridge.contract.Call(opts, out, "valset")
	return *ret0, err
}

// Valset is a free data retrieval call binding the contract method 0x7f54af0c.
//
// Solidity: function valset() view returns(address)
func (_CosmosBridge *CosmosBridgeSession) Valset() (common.Address, error) {
	return _CosmosBridge.Contract.Valset(&_CosmosBridge.CallOpts)
}

// Valset is a free data retrieval call binding the contract method 0x7f54af0c.
//
// Solidity: function valset() view returns(address)
func (_CosmosBridge *CosmosBridgeCallerSession) Valset() (common.Address, error) {
	return _CosmosBridge.Contract.Valset(&_CosmosBridge.CallOpts)
}

// CompleteProphecyClaim is a paid mutator transaction binding the contract method 0x6b3ce98c.
//
// Solidity: function completeProphecyClaim(uint256 _prophecyID) returns()
func (_CosmosBridge *CosmosBridgeTransactor) CompleteProphecyClaim(opts *bind.TransactOpts, _prophecyID *big.Int) (*types.Transaction, error) {
	return _CosmosBridge.contract.Transact(opts, "completeProphecyClaim", _prophecyID)
}

// CompleteProphecyClaim is a paid mutator transaction binding the contract method 0x6b3ce98c.
//
// Solidity: function completeProphecyClaim(uint256 _prophecyID) returns()
func (_CosmosBridge *CosmosBridgeSession) CompleteProphecyClaim(_prophecyID *big.Int) (*types.Transaction, error) {
	return _CosmosBridge.Contract.CompleteProphecyClaim(&_CosmosBridge.TransactOpts, _prophecyID)
}

// CompleteProphecyClaim is a paid mutator transaction binding the contract method 0x6b3ce98c.
//
// Solidity: function completeProphecyClaim(uint256 _prophecyID) returns()
func (_CosmosBridge *CosmosBridgeTransactorSession) CompleteProphecyClaim(_prophecyID *big.Int) (*types.Transaction, error) {
	return _CosmosBridge.Contract.CompleteProphecyClaim(&_CosmosBridge.TransactOpts, _prophecyID)
}

// NewProphecyClaim is a paid mutator transaction binding the contract method 0x9408ee9c.
//
// Solidity: function newProphecyClaim(uint8 _claimType, bytes _cosmosSender, address _ethereumReceiver, address _tokenAddress, string _symbol, uint256 _amount) returns()
func (_CosmosBridge *CosmosBridgeTransactor) NewProphecyClaim(opts *bind.TransactOpts, _claimType uint8, _cosmosSender []byte, _ethereumReceiver common.Address, _tokenAddress common.Address, _symbol string, _amount *big.Int) (*types.Transaction, error) {
	return _CosmosBridge.contract.Transact(opts, "newProphecyClaim", _claimType, _cosmosSender, _ethereumReceiver, _tokenAddress, _symbol, _amount)
}

// NewProphecyClaim is a paid mutator transaction binding the contract method 0x9408ee9c.
//
// Solidity: function newProphecyClaim(uint8 _claimType, bytes _cosmosSender, address _ethereumReceiver, address _tokenAddress, string _symbol, uint256 _amount) returns()
func (_CosmosBridge *CosmosBridgeSession) NewProphecyClaim(_claimType uint8, _cosmosSender []byte, _ethereumReceiver common.Address, _tokenAddress common.Address, _symbol string, _amount *big.Int) (*types.Transaction, error) {
	return _CosmosBridge.Contract.NewProphecyClaim(&_CosmosBridge.TransactOpts, _claimType, _cosmosSender, _ethereumReceiver, _tokenAddress, _symbol, _amount)
}

// NewProphecyClaim is a paid mutator transaction binding the contract method 0x9408ee9c.
//
// Solidity: function newProphecyClaim(uint8 _claimType, bytes _cosmosSender, address _ethereumReceiver, address _tokenAddress, string _symbol, uint256 _amount) returns()
func (_CosmosBridge *CosmosBridgeTransactorSession) NewProphecyClaim(_claimType uint8, _cosmosSender []byte, _ethereumReceiver common.Address, _tokenAddress common.Address, _symbol string, _amount *big.Int) (*types.Transaction, error) {
	return _CosmosBridge.Contract.NewProphecyClaim(&_CosmosBridge.TransactOpts, _claimType, _cosmosSender, _ethereumReceiver, _tokenAddress, _symbol, _amount)
}

// SetBridgeBank is a paid mutator transaction binding the contract method 0x814c92c3.
//
// Solidity: function setBridgeBank(address _bridgeBank) returns()
func (_CosmosBridge *CosmosBridgeTransactor) SetBridgeBank(opts *bind.TransactOpts, _bridgeBank common.Address) (*types.Transaction, error) {
	return _CosmosBridge.contract.Transact(opts, "setBridgeBank", _bridgeBank)
}

// SetBridgeBank is a paid mutator transaction binding the contract method 0x814c92c3.
//
// Solidity: function setBridgeBank(address _bridgeBank) returns()
func (_CosmosBridge *CosmosBridgeSession) SetBridgeBank(_bridgeBank common.Address) (*types.Transaction, error) {
	return _CosmosBridge.Contract.SetBridgeBank(&_CosmosBridge.TransactOpts, _bridgeBank)
}

// SetBridgeBank is a paid mutator transaction binding the contract method 0x814c92c3.
//
// Solidity: function setBridgeBank(address _bridgeBank) returns()
func (_CosmosBridge *CosmosBridgeTransactorSession) SetBridgeBank(_bridgeBank common.Address) (*types.Transaction, error) {
	return _CosmosBridge.Contract.SetBridgeBank(&_CosmosBridge.TransactOpts, _bridgeBank)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address _oracle) returns()
func (_CosmosBridge *CosmosBridgeTransactor) SetOracle(opts *bind.TransactOpts, _oracle common.Address) (*types.Transaction, error) {
	return _CosmosBridge.contract.Transact(opts, "setOracle", _oracle)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address _oracle) returns()
func (_CosmosBridge *CosmosBridgeSession) SetOracle(_oracle common.Address) (*types.Transaction, error) {
	return _CosmosBridge.Contract.SetOracle(&_CosmosBridge.TransactOpts, _oracle)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address _oracle) returns()
func (_CosmosBridge *CosmosBridgeTransactorSession) SetOracle(_oracle common.Address) (*types.Transaction, error) {
	return _CosmosBridge.Contract.SetOracle(&_CosmosBridge.TransactOpts, _oracle)
}

// CosmosBridgeLogBridgeBankSetIterator is returned from FilterLogBridgeBankSet and is used to iterate over the raw logs and unpacked data for LogBridgeBankSet events raised by the CosmosBridge contract.
type CosmosBridgeLogBridgeBankSetIterator struct {
	Event *CosmosBridgeLogBridgeBankSet // Event containing the contract specifics and raw log

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
func (it *CosmosBridgeLogBridgeBankSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmosBridgeLogBridgeBankSet)
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
		it.Event = new(CosmosBridgeLogBridgeBankSet)
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
func (it *CosmosBridgeLogBridgeBankSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmosBridgeLogBridgeBankSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmosBridgeLogBridgeBankSet represents a LogBridgeBankSet event raised by the CosmosBridge contract.
type CosmosBridgeLogBridgeBankSet struct {
	BridgeBank common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLogBridgeBankSet is a free log retrieval operation binding the contract event 0xc8b65043fb196ac032b79a435397d1d14a96b4e9d12e366c3b1f550cb01d2dfa.
//
// Solidity: event LogBridgeBankSet(address _bridgeBank)
func (_CosmosBridge *CosmosBridgeFilterer) FilterLogBridgeBankSet(opts *bind.FilterOpts) (*CosmosBridgeLogBridgeBankSetIterator, error) {

	logs, sub, err := _CosmosBridge.contract.FilterLogs(opts, "LogBridgeBankSet")
	if err != nil {
		return nil, err
	}
	return &CosmosBridgeLogBridgeBankSetIterator{contract: _CosmosBridge.contract, event: "LogBridgeBankSet", logs: logs, sub: sub}, nil
}

// WatchLogBridgeBankSet is a free log subscription operation binding the contract event 0xc8b65043fb196ac032b79a435397d1d14a96b4e9d12e366c3b1f550cb01d2dfa.
//
// Solidity: event LogBridgeBankSet(address _bridgeBank)
func (_CosmosBridge *CosmosBridgeFilterer) WatchLogBridgeBankSet(opts *bind.WatchOpts, sink chan<- *CosmosBridgeLogBridgeBankSet) (event.Subscription, error) {

	logs, sub, err := _CosmosBridge.contract.WatchLogs(opts, "LogBridgeBankSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmosBridgeLogBridgeBankSet)
				if err := _CosmosBridge.contract.UnpackLog(event, "LogBridgeBankSet", log); err != nil {
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

// ParseLogBridgeBankSet is a log parse operation binding the contract event 0xc8b65043fb196ac032b79a435397d1d14a96b4e9d12e366c3b1f550cb01d2dfa.
//
// Solidity: event LogBridgeBankSet(address _bridgeBank)
func (_CosmosBridge *CosmosBridgeFilterer) ParseLogBridgeBankSet(log types.Log) (*CosmosBridgeLogBridgeBankSet, error) {
	event := new(CosmosBridgeLogBridgeBankSet)
	if err := _CosmosBridge.contract.UnpackLog(event, "LogBridgeBankSet", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CosmosBridgeLogNewProphecyClaimIterator is returned from FilterLogNewProphecyClaim and is used to iterate over the raw logs and unpacked data for LogNewProphecyClaim events raised by the CosmosBridge contract.
type CosmosBridgeLogNewProphecyClaimIterator struct {
	Event *CosmosBridgeLogNewProphecyClaim // Event containing the contract specifics and raw log

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
func (it *CosmosBridgeLogNewProphecyClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmosBridgeLogNewProphecyClaim)
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
		it.Event = new(CosmosBridgeLogNewProphecyClaim)
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
func (it *CosmosBridgeLogNewProphecyClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmosBridgeLogNewProphecyClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmosBridgeLogNewProphecyClaim represents a LogNewProphecyClaim event raised by the CosmosBridge contract.
type CosmosBridgeLogNewProphecyClaim struct {
	ProphecyID       *big.Int
	ClaimType        uint8
	CosmosSender     []byte
	EthereumReceiver common.Address
	ValidatorAddress common.Address
	TokenAddress     common.Address
	Symbol           string
	Amount           *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterLogNewProphecyClaim is a free log retrieval operation binding the contract event 0x4c4b04a2b190e6bb01b6243f150fc76174861acd19cf98841801baaff5262dd8.
//
// Solidity: event LogNewProphecyClaim(uint256 _prophecyID, uint8 _claimType, bytes _cosmosSender, address _ethereumReceiver, address _validatorAddress, address _tokenAddress, string _symbol, uint256 _amount)
func (_CosmosBridge *CosmosBridgeFilterer) FilterLogNewProphecyClaim(opts *bind.FilterOpts) (*CosmosBridgeLogNewProphecyClaimIterator, error) {

	logs, sub, err := _CosmosBridge.contract.FilterLogs(opts, "LogNewProphecyClaim")
	if err != nil {
		return nil, err
	}
	return &CosmosBridgeLogNewProphecyClaimIterator{contract: _CosmosBridge.contract, event: "LogNewProphecyClaim", logs: logs, sub: sub}, nil
}

// WatchLogNewProphecyClaim is a free log subscription operation binding the contract event 0x4c4b04a2b190e6bb01b6243f150fc76174861acd19cf98841801baaff5262dd8.
//
// Solidity: event LogNewProphecyClaim(uint256 _prophecyID, uint8 _claimType, bytes _cosmosSender, address _ethereumReceiver, address _validatorAddress, address _tokenAddress, string _symbol, uint256 _amount)
func (_CosmosBridge *CosmosBridgeFilterer) WatchLogNewProphecyClaim(opts *bind.WatchOpts, sink chan<- *CosmosBridgeLogNewProphecyClaim) (event.Subscription, error) {

	logs, sub, err := _CosmosBridge.contract.WatchLogs(opts, "LogNewProphecyClaim")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmosBridgeLogNewProphecyClaim)
				if err := _CosmosBridge.contract.UnpackLog(event, "LogNewProphecyClaim", log); err != nil {
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

// ParseLogNewProphecyClaim is a log parse operation binding the contract event 0x4c4b04a2b190e6bb01b6243f150fc76174861acd19cf98841801baaff5262dd8.
//
// Solidity: event LogNewProphecyClaim(uint256 _prophecyID, uint8 _claimType, bytes _cosmosSender, address _ethereumReceiver, address _validatorAddress, address _tokenAddress, string _symbol, uint256 _amount)
func (_CosmosBridge *CosmosBridgeFilterer) ParseLogNewProphecyClaim(log types.Log) (*CosmosBridgeLogNewProphecyClaim, error) {
	event := new(CosmosBridgeLogNewProphecyClaim)
	if err := _CosmosBridge.contract.UnpackLog(event, "LogNewProphecyClaim", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CosmosBridgeLogOracleSetIterator is returned from FilterLogOracleSet and is used to iterate over the raw logs and unpacked data for LogOracleSet events raised by the CosmosBridge contract.
type CosmosBridgeLogOracleSetIterator struct {
	Event *CosmosBridgeLogOracleSet // Event containing the contract specifics and raw log

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
func (it *CosmosBridgeLogOracleSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmosBridgeLogOracleSet)
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
		it.Event = new(CosmosBridgeLogOracleSet)
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
func (it *CosmosBridgeLogOracleSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmosBridgeLogOracleSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmosBridgeLogOracleSet represents a LogOracleSet event raised by the CosmosBridge contract.
type CosmosBridgeLogOracleSet struct {
	Oracle common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLogOracleSet is a free log retrieval operation binding the contract event 0x6efb0434342713e2e9b1501dbebf76b4ed18406ea77ab5d56535cc26dec3adc0.
//
// Solidity: event LogOracleSet(address _oracle)
func (_CosmosBridge *CosmosBridgeFilterer) FilterLogOracleSet(opts *bind.FilterOpts) (*CosmosBridgeLogOracleSetIterator, error) {

	logs, sub, err := _CosmosBridge.contract.FilterLogs(opts, "LogOracleSet")
	if err != nil {
		return nil, err
	}
	return &CosmosBridgeLogOracleSetIterator{contract: _CosmosBridge.contract, event: "LogOracleSet", logs: logs, sub: sub}, nil
}

// WatchLogOracleSet is a free log subscription operation binding the contract event 0x6efb0434342713e2e9b1501dbebf76b4ed18406ea77ab5d56535cc26dec3adc0.
//
// Solidity: event LogOracleSet(address _oracle)
func (_CosmosBridge *CosmosBridgeFilterer) WatchLogOracleSet(opts *bind.WatchOpts, sink chan<- *CosmosBridgeLogOracleSet) (event.Subscription, error) {

	logs, sub, err := _CosmosBridge.contract.WatchLogs(opts, "LogOracleSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmosBridgeLogOracleSet)
				if err := _CosmosBridge.contract.UnpackLog(event, "LogOracleSet", log); err != nil {
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

// ParseLogOracleSet is a log parse operation binding the contract event 0x6efb0434342713e2e9b1501dbebf76b4ed18406ea77ab5d56535cc26dec3adc0.
//
// Solidity: event LogOracleSet(address _oracle)
func (_CosmosBridge *CosmosBridgeFilterer) ParseLogOracleSet(log types.Log) (*CosmosBridgeLogOracleSet, error) {
	event := new(CosmosBridgeLogOracleSet)
	if err := _CosmosBridge.contract.UnpackLog(event, "LogOracleSet", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CosmosBridgeLogProphecyCompletedIterator is returned from FilterLogProphecyCompleted and is used to iterate over the raw logs and unpacked data for LogProphecyCompleted events raised by the CosmosBridge contract.
type CosmosBridgeLogProphecyCompletedIterator struct {
	Event *CosmosBridgeLogProphecyCompleted // Event containing the contract specifics and raw log

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
func (it *CosmosBridgeLogProphecyCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmosBridgeLogProphecyCompleted)
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
		it.Event = new(CosmosBridgeLogProphecyCompleted)
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
func (it *CosmosBridgeLogProphecyCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmosBridgeLogProphecyCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmosBridgeLogProphecyCompleted represents a LogProphecyCompleted event raised by the CosmosBridge contract.
type CosmosBridgeLogProphecyCompleted struct {
	ProphecyID *big.Int
	ClaimType  uint8
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLogProphecyCompleted is a free log retrieval operation binding the contract event 0x79e7c1c0bd54f11809c3bf6023c242783602d61ceff272c6bba6f8559c24ad0d.
//
// Solidity: event LogProphecyCompleted(uint256 _prophecyID, uint8 _claimType)
func (_CosmosBridge *CosmosBridgeFilterer) FilterLogProphecyCompleted(opts *bind.FilterOpts) (*CosmosBridgeLogProphecyCompletedIterator, error) {

	logs, sub, err := _CosmosBridge.contract.FilterLogs(opts, "LogProphecyCompleted")
	if err != nil {
		return nil, err
	}
	return &CosmosBridgeLogProphecyCompletedIterator{contract: _CosmosBridge.contract, event: "LogProphecyCompleted", logs: logs, sub: sub}, nil
}

// WatchLogProphecyCompleted is a free log subscription operation binding the contract event 0x79e7c1c0bd54f11809c3bf6023c242783602d61ceff272c6bba6f8559c24ad0d.
//
// Solidity: event LogProphecyCompleted(uint256 _prophecyID, uint8 _claimType)
func (_CosmosBridge *CosmosBridgeFilterer) WatchLogProphecyCompleted(opts *bind.WatchOpts, sink chan<- *CosmosBridgeLogProphecyCompleted) (event.Subscription, error) {

	logs, sub, err := _CosmosBridge.contract.WatchLogs(opts, "LogProphecyCompleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmosBridgeLogProphecyCompleted)
				if err := _CosmosBridge.contract.UnpackLog(event, "LogProphecyCompleted", log); err != nil {
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

// ParseLogProphecyCompleted is a log parse operation binding the contract event 0x79e7c1c0bd54f11809c3bf6023c242783602d61ceff272c6bba6f8559c24ad0d.
//
// Solidity: event LogProphecyCompleted(uint256 _prophecyID, uint8 _claimType)
func (_CosmosBridge *CosmosBridgeFilterer) ParseLogProphecyCompleted(log types.Log) (*CosmosBridgeLogProphecyCompleted, error) {
	event := new(CosmosBridgeLogProphecyCompleted)
	if err := _CosmosBridge.contract.UnpackLog(event, "LogProphecyCompleted", log); err != nil {
		return nil, err
	}
	return event, nil
}
