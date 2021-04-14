// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package adjudicator

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

// ChannelAllocation is an auto generated low-level Go binding around an user-defined struct.
type ChannelAllocation struct {
	Assets   []common.Address
	Balances [][]*big.Int
	Locked   []ChannelSubAlloc
}

// ChannelParams is an auto generated low-level Go binding around an user-defined struct.
type ChannelParams struct {
	ChallengeDuration *big.Int
	Nonce             *big.Int
	App               common.Address
	Participants      []common.Address
}

// ChannelState is an auto generated low-level Go binding around an user-defined struct.
type ChannelState struct {
	ChannelID [32]byte
	Version   uint64
	Outcome   ChannelAllocation
	AppData   []byte
	IsFinal   bool
}

// ChannelSubAlloc is an auto generated low-level Go binding around an user-defined struct.
type ChannelSubAlloc struct {
	ID       [32]byte
	Balances []*big.Int
}

// AdjudicatorABI is the input ABI used to generate the binding from.
const AdjudicatorABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"channelID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"phase\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"timeout\",\"type\":\"uint64\"}],\"name\":\"ChannelUpdate\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"challengeDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"app\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"participants\",\"type\":\"address[]\"}],\"internalType\":\"structChannel.Params\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"channelID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"challengeDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"app\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"participants\",\"type\":\"address[]\"}],\"internalType\":\"structChannel.Params\",\"name\":\"params\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"channelID\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[][]\",\"name\":\"balances\",\"type\":\"uint256[][]\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"ID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"}],\"internalType\":\"structChannel.SubAlloc[]\",\"name\":\"locked\",\"type\":\"tuple[]\"}],\"internalType\":\"structChannel.Allocation\",\"name\":\"outcome\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"appData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"isFinal\",\"type\":\"bool\"}],\"internalType\":\"structChannel.State\",\"name\":\"state\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"channelID\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[][]\",\"name\":\"balances\",\"type\":\"uint256[][]\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"ID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"}],\"internalType\":\"structChannel.SubAlloc[]\",\"name\":\"locked\",\"type\":\"tuple[]\"}],\"internalType\":\"structChannel.Allocation\",\"name\":\"outcome\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"appData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"isFinal\",\"type\":\"bool\"}],\"internalType\":\"structChannel.State[]\",\"name\":\"subStates\",\"type\":\"tuple[]\"}],\"name\":\"conclude\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"challengeDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"app\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"participants\",\"type\":\"address[]\"}],\"internalType\":\"structChannel.Params\",\"name\":\"params\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"channelID\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[][]\",\"name\":\"balances\",\"type\":\"uint256[][]\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"ID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"}],\"internalType\":\"structChannel.SubAlloc[]\",\"name\":\"locked\",\"type\":\"tuple[]\"}],\"internalType\":\"structChannel.Allocation\",\"name\":\"outcome\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"appData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"isFinal\",\"type\":\"bool\"}],\"internalType\":\"structChannel.State\",\"name\":\"state\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"sigs\",\"type\":\"bytes[]\"}],\"name\":\"concludeFinal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"disputes\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"timeout\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"challengeDuration\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"hasApp\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"phase\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"stateHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"channelID\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[][]\",\"name\":\"balances\",\"type\":\"uint256[][]\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"ID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"}],\"internalType\":\"structChannel.SubAlloc[]\",\"name\":\"locked\",\"type\":\"tuple[]\"}],\"internalType\":\"structChannel.Allocation\",\"name\":\"outcome\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"appData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"isFinal\",\"type\":\"bool\"}],\"internalType\":\"structChannel.State\",\"name\":\"state\",\"type\":\"tuple\"}],\"name\":\"hashState\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"challengeDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"app\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"participants\",\"type\":\"address[]\"}],\"internalType\":\"structChannel.Params\",\"name\":\"params\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"channelID\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[][]\",\"name\":\"balances\",\"type\":\"uint256[][]\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"ID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"}],\"internalType\":\"structChannel.SubAlloc[]\",\"name\":\"locked\",\"type\":\"tuple[]\"}],\"internalType\":\"structChannel.Allocation\",\"name\":\"outcome\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"appData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"isFinal\",\"type\":\"bool\"}],\"internalType\":\"structChannel.State\",\"name\":\"stateOld\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"channelID\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[][]\",\"name\":\"balances\",\"type\":\"uint256[][]\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"ID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"}],\"internalType\":\"structChannel.SubAlloc[]\",\"name\":\"locked\",\"type\":\"tuple[]\"}],\"internalType\":\"structChannel.Allocation\",\"name\":\"outcome\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"appData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"isFinal\",\"type\":\"bool\"}],\"internalType\":\"structChannel.State\",\"name\":\"state\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"actorIdx\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"progress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"challengeDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"app\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"participants\",\"type\":\"address[]\"}],\"internalType\":\"structChannel.Params\",\"name\":\"params\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"channelID\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[][]\",\"name\":\"balances\",\"type\":\"uint256[][]\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"ID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"}],\"internalType\":\"structChannel.SubAlloc[]\",\"name\":\"locked\",\"type\":\"tuple[]\"}],\"internalType\":\"structChannel.Allocation\",\"name\":\"outcome\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"appData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"isFinal\",\"type\":\"bool\"}],\"internalType\":\"structChannel.State\",\"name\":\"state\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"sigs\",\"type\":\"bytes[]\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// AdjudicatorBin is the compiled bytecode used for deploying new contracts.
var AdjudicatorBin = "0x608060405234801561001057600080fd5b506124d9806100206000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c80636bbf706a1161005b5780636bbf706a146100d85780638bba7507146100eb578063a1ee1592146100fe578063a83a5cc51461011e5761007d565b806311be199714610082578063170e6715146100b057806336995831146100c5575b600080fd5b610095610090366004611825565b610131565b6040516100a7969594939291906123bc565b60405180910390f35b6100c36100be366004611877565b610180565b005b6100c36100d3366004611a0a565b610259565b6100c36100e6366004611877565b6103e4565b6100c36100f9366004611949565b6104df565b61011161010c36600461183d565b610546565b6040516100a79190611d3f565b61011161012c366004611abb565b610560565b600060208190529081526040902080546001909101546001600160401b0380831692600160401b8104821692600160801b82049092169160ff600160c01b8304811692600160c81b9004169086565b61018a838361056b565b610195838383610597565b61019d611348565b60006101ac8460000151610621565b9150915080156102465783602001516001600160401b031682604001516001600160401b0316106101f85760405162461bcd60e51b81526004016101ef90611dfd565b60405180910390fd5b608082015160ff161561021d5760405162461bcd60e51b81526004016101ef90611f0d565b81516001600160401b031642106102465760405162461bcd60e51b81526004016101ef906120b9565b610252858560006106b2565b5050505050565b610261611348565b835161026c9061078b565b608081015190915060ff166102aa5780516001600160401b03164210156102a55760405162461bcd60e51b81526004016101ef9061208d565b6102fb565b608081015160ff16600114156102e35780516001600160401b031642106102a55760405162461bcd60e51b81526004016101ef90611d73565b60405162461bcd60e51b81526004016101ef90611e86565b60408601516001600160a01b03166103255760405162461bcd60e51b81526004016101ef90612324565b85606001515183106103495760405162461bcd60e51b81526004016101ef90611fc3565b610353868561056b565b61035c85610560565b8160a001511461037e5760405162461bcd60e51b81526004016101ef90611ead565b6103a861038a856107ce565b838860600151868151811061039b57fe5b60200260200101516107f7565b6103c45760405162461bcd60e51b81526004016101ef90611dd2565b6103d086868686610832565b6103dc868560016106b2565b505050505050565b6080820151151560011461040a5760405162461bcd60e51b81526004016101ef90611e26565b604080830151015151156104305760405162461bcd60e51b81526004016101ef90611ed6565b61043a838361056b565b610445838383610597565b61044d611348565b600061045c8460000151610621565b91509150801561048e57608082015160ff166002141561048e5760405162461bcd60e51b81526004016101ef90611ff2565b61049a858560026106b2565b604080516000808252602082019092526060916104cd565b6104ba61137d565b8152602001906001900390816104b25790505b5090506103dc85828860600151610916565b6104e7611348565b82516104f29061078b565b608081015190915060ff166002141561051d5760405162461bcd60e51b81526004016101ef90611ff2565b610527848461056b565b6105318383610b5a565b61054083838660600151610916565b50505050565b600061055182610b98565b8051906020012090505b919050565b6000610551826107ce565b61057482610546565b8151146105935760405162461bcd60e51b81526004016101ef90612262565b5050565b60606105a2836107ce565b90508151846060015151146105c95760405162461bcd60e51b81526004016101ef9061211f565b60005b8251811015610252576105fd828483815181106105e557fe5b60200260200101518760600151848151811061039b57fe5b6106195760405162461bcd60e51b81526004016101ef90611dd2565b6001016105cc565b610629611348565b6000610633611348565b50505060008181526020818152604091829020825160c08101845281546001600160401b038082168352600160401b8204811694830194909452600160801b81049093169381019390935260ff600160c01b8304811615156060850152600160c81b90920490911660808301526001015460a082018190521515915091565b6106ba611348565b60006106c98460000151610621565b86516001600160401b03908116602080850191909152870151166040808401919091528701516001600160a01b031615156060830152909250905082600281111561071057fe5b60ff16608083015261072184610560565b60a0830152608084015115610741576001600160401b034216825261077f565b8015806107555750608082015160ff166001145b1561077f576020820151610773906001600160401b03421690610bab565b6001600160401b031682525b83516102529083610c00565b610793611348565b61079b611348565b60006107a684610621565b91509150806107c75760405162461bcd60e51b81526004016101ef90611f36565b5092915050565b6060816040516020016107e191906123a9565b6040516020818303038152906040529050919050565b60008061080a8580519060200120610cfb565b905060006108188286610d4c565b6001600160a01b0390811690851614925050509392505050565b82602001516001016001600160401b031682602001516001600160401b03161461086e5760405162461bcd60e51b81526004016101ef9061222b565b6080830151156108905760405162461bcd60e51b81526004016101ef90611f5e565b6108a883604001518360400151866060015151610f37565b6040808501519051637614eebf60e11b81526001600160a01b0382169063ec29dd7e906108df90889088908890889060040161235e565b60006040518083038186803b1580156108f757600080fd5b505afa15801561090b573d6000803e3d6000fd5b505050505050505050565b60408301515160005b815181101561025257606083516001600160401b038111801561094157600080fd5b5060405190808252806020026020018201604052801561096b578160200160208202803683370190505b50905060005b8151811015610ad357866040015160200151838151811061098e57fe5b602002602001015181815181106109a157fe5b60200260200101518282815181106109b557fe5b60200260200101818152505060005b8651811015610aca576109d561137d565b8782815181106109e157fe5b602002602001015190508585815181106109f757fe5b60200260200101516001600160a01b03168160400151600001518681518110610a1c57fe5b60200260200101516001600160a01b031614610a4a5760405162461bcd60e51b81526004016101ef90612029565b6000848481518110610a5857fe5b6020026020010151905060008260400151602001518781518110610a7857fe5b60200260200101518581518110610a8b57fe5b60200260200101519050610aa8818361115f90919063ffffffff16565b868681518110610ab457fe5b60209081029190910101525050506001016109c4565b50600101610971565b50828281518110610ae057fe5b60200260200101516001600160a01b031663fc79a66d876000015186846040518463ffffffff1660e01b8152600401610b1b93929190611d48565b600060405180830381600087803b158015610b3557600080fd5b505af1158015610b49573d6000803e3d6000fd5b50506001909301925061091f915050565b610b63826111c0565b6000610b7183836000611299565b905081518114610b935760405162461bcd60e51b81526004016101ef90612056565b505050565b6060816040516020016107e1919061234b565b8082016001600160401b038084169082161015610bfa576040805162461bcd60e51b81526020600482015260086024820152676f766572666c6f7760c01b604482015290519081900360640190fd5b92915050565b600082815260208181526040918290208351815492850151858501516060870151608088015167ffffffffffffffff199096166001600160401b03808616919091176fffffffffffffffff00000000000000001916600160401b948216949094029390931767ffffffffffffffff60801b1916600160801b938316939093029290921760ff60c01b1916600160c01b921515929092029190911760ff60c81b1916600160c81b60ff86160217835560a0860151600190930192909255925185937f895ef5a5fc3efd313a300b006d6ce97ff0670dfe04f6eea90417edf924fa786b93610cef93929091906123f9565b60405180910390a25050565b604080517f19457468657265756d205369676e6564204d6573736167653a0a333200000000602080830191909152603c8083019490945282518083039094018452605c909101909152815191012090565b60008151604114610da4576040805162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e67746800604482015290519081900360640190fd5b60208201516040830151606084015160001a7f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0821115610e155760405162461bcd60e51b81526004018080602001828103825260228152602001806124606022913960400191505060405180910390fd5b8060ff16601b14158015610e2d57508060ff16601c14155b15610e695760405162461bcd60e51b81526004018080602001828103825260228152602001806124826022913960400191505060405180910390fd5b600060018783868660405160008152602001604052604051808581526020018460ff1681526020018381526020018281526020019450505050506020604051602081039080840390855afa158015610ec5573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b038116610f2d576040805162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e61747572650000000000000000604482015290519081900360640190fd5b9695505050505050565b81602001515183602001515114610f605760405162461bcd60e51b81526004016101ef906122b6565b81515183515114610f835760405162461bcd60e51b81526004016101ef90611f93565b60408301515115610fa65760405162461bcd60e51b81526004016101ef90612156565b60408201515115610fc95760405162461bcd60e51b81526004016101ef90611e4f565b60005b825151811015610540578251805182908110610fe457fe5b60200260200101516001600160a01b03168460000151828151811061100557fe5b60200260200101516001600160a01b0316146110335760405162461bcd60e51b81526004016101ef90611d9b565b600080838660200151848151811061104757fe5b6020026020010151511461106d5760405162461bcd60e51b81526004016101ef906122ed565b838560200151848151811061107e57fe5b602002602001015151146110a45760405162461bcd60e51b81526004016101ef906121f4565b60005b84811015611135576110ec876020015185815181106110c257fe5b602002602001015182815181106110d557fe5b60200260200101518461115f90919063ffffffff16565b925061112b8660200151858151811061110157fe5b6020026020010151828151811061111457fe5b60200260200101518361115f90919063ffffffff16565b91506001016110a7565b508082146111555760405162461bcd60e51b81526004016101ef906121bd565b5050600101610fcc565b6000828201838110156111b9576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b6111c8611348565b81516111d39061078b565b90506111de82610560565b8160a00151146112005760405162461bcd60e51b81526004016101ef906120f0565b608081015160ff16600214156112165750611296565b608081015160ff1615801561122c575080606001515b15611259576020810151815161124d916001600160401b0390911690610bab565b6001600160401b031681525b80516001600160401b03164210156112835760405162461bcd60e51b81526004016101ef9061218d565b6002608082015281516105939082610c00565b50565b60408084015101516000908290825b815181101561133d576112b961137d565b8684815181106112c557fe5b6020026020010151905080600001518383815181106112e057fe5b602002602001015160000151146113095760405162461bcd60e51b81526004016101ef9061228a565b611312816111c0565b6040808201510151516001909401931561133457611331818886611299565b93505b506001016112a8565b509095945050505050565b6040805160c081018252600080825260208201819052918101829052606081018290526080810182905260a081019190915290565b6040805160a0810182526000808252602082015290810161139c6113b0565b815260606020820152600060409091015290565b60405180606001604052806060815260200160608152602001606081525090565b80356001600160a01b038116811461055b57600080fd5b600082601f8301126113f8578081fd5b813561140b61140682612442565b61241f565b81815291506020808301908481018184028601820187101561142c57600080fd5b60005b8481101561145257611440826113d1565b8452928201929082019060010161142f565b505050505092915050565b600082601f83011261146d578081fd5b813561147b61140682612442565b818152915060208083019084810160005b84811015611452576114a3888484358a0101611563565b8452928201929082019060010161148c565b600082601f8301126114c5578081fd5b81356114d361140682612442565b818152915060208083019084810160005b848110156114525781358701604080601f19838c0301121561150557600080fd5b80518181016001600160401b03828210818311171561152057fe5b81845284880135835292840135928084111561153b57600080fd5b505061154b8b8784860101611563565b818701528652505092820192908201906001016114e4565b600082601f830112611573578081fd5b813561158161140682612442565b8181529150602080830190848101818402860182018710156115a257600080fd5b60005b84811015611452578135845292820192908201906001016115a5565b8035801515811461055b57600080fd5b600082601f8301126115e1578081fd5b81356001600160401b038111156115f457fe5b611607601f8201601f191660200161241f565b915080825283602082850101111561161e57600080fd5b8060208401602084013760009082016020015292915050565b600060608284031215611648578081fd5b604051606081016001600160401b03828210818311171561166557fe5b81604052829350843591508082111561167d57600080fd5b611689868387016113e8565b8352602085013591508082111561169f57600080fd5b6116ab8683870161145d565b602084015260408501359150808211156116c457600080fd5b506116d1858286016114b5565b6040830152505092915050565b6000608082840312156116ef578081fd5b604051608081016001600160401b03828210818311171561170c57fe5b81604052829350843583526020850135602084015261172d604086016113d1565b6040840152606085013591508082111561174657600080fd5b50611753858286016113e8565b6060830152505092915050565b600060a08284031215611771578081fd5b60405160a081016001600160401b03828210818311171561178e57fe5b81604052829350843583526117a56020860161180e565b602084015260408501359150808211156117be57600080fd5b6117ca86838701611637565b604084015260608501359150808211156117e357600080fd5b506117f0858286016115d1565b606083015250611802608084016115c1565b60808201525092915050565b80356001600160401b038116811461055b57600080fd5b600060208284031215611836578081fd5b5035919050565b60006020828403121561184e578081fd5b81356001600160401b03811115611863578182fd5b61186f848285016116de565b949350505050565b60008060006060848603121561188b578182fd5b83356001600160401b03808211156118a1578384fd5b6118ad878388016116de565b94506020915081860135818111156118c3578485fd5b6118cf88828901611760565b9450506040860135818111156118e3578384fd5b86019050601f810187136118f5578283fd5b803561190361140682612442565b81815283810190838501865b84811015611938576119268c8884358901016115d1565b8452928601929086019060010161190f565b505080955050505050509250925092565b60008060006060848603121561195d578081fd5b83356001600160401b0380821115611973578283fd5b61197f878388016116de565b9450602091508186013581811115611995578384fd5b6119a188828901611760565b9450506040860135818111156119b5578384fd5b86019050601f810187136119c7578283fd5b80356119d561140682612442565b81815283810190838501865b84811015611938576119f88c888435890101611760565b845292860192908601906001016119e1565b600080600080600060a08688031215611a21578283fd5b85356001600160401b0380821115611a37578485fd5b611a4389838a016116de565b96506020880135915080821115611a58578485fd5b611a6489838a01611760565b95506040880135915080821115611a79578485fd5b611a8589838a01611760565b9450606088013593506080880135915080821115611aa1578283fd5b50611aae888289016115d1565b9150509295509295909350565b600060208284031215611acc578081fd5b81356001600160401b03811115611ae1578182fd5b61186f84828501611760565b6000815180845260208085019450808401835b83811015611b255781516001600160a01b031687529582019590820190600101611b00565b509495945050505050565b6000815180845260208085018081965082840281019150828601855b85811015611b8a5782840389528151805185528501516040868601819052611b7681870183611b97565b9a87019a9550505090840190600101611b4c565b5091979650505050505050565b6000815180845260208085019450808401835b83811015611b2557815187529582019590820190600101611baa565b15159052565b60008151808452815b81811015611bf157602081850181015186830182015201611bd5565b81811115611c025782602083870101525b50601f01601f19169290920160200192915050565b6000815183526020820151602084015260018060a01b03604083015116604084015260608201516080606085015261186f6080850182611aed565b60008151835260206001600160401b03818401511681850152604083015160a060408601528051606060a0870152611c8e610100870182611aed565b83830151609f19888303810160c08a01528151808452929350908501918386019080870285018701885b82811015611ce657601f19878303018452611cd4828751611b97565b95890195938901939150600101611cb8565b5060408701519750838b82030160e08c0152611d028189611b30565b97505050505050505060608301518482036060860152611d228282611bcc565b9150506080830151611d376080860182611bc6565b509392505050565b90815260200190565b600084825260606020830152611d616060830185611aed565b8281036040840152610f2d8185611b97565b6020808252600e908201526d1d1a5b595bdd5d081c185cdcd95960921b604082015260600190565b6020808252601a908201527f6173736574735b695d2061646472657373206d69736d61746368000000000000604082015260600190565b602080825260119082015270696e76616c6964207369676e617475726560781b604082015260600190565b6020808252600f908201526e34b73b30b634b2103b32b939b4b7b760891b604082015260600190565b6020808252600f908201526e1cdd185d19481b9bdd08199a5b985b608a1b604082015260600190565b60208082526019908201527f66756e6473206c6f636b656420696e206e657720737461746500000000000000604082015260600190565b6020808252600d908201526c696e76616c696420706861736560981b604082015260600190565b6020808252600f908201526e77726f6e67206f6c6420737461746560881b604082015260600190565b60208082526018908201527f63616e6e6f742068617665207375622d6368616e6e656c730000000000000000604082015260600190565b6020808252600f908201526e696e636f727265637420706861736560881b604082015260600190565b6020808252600e908201526d1b9bdd081c9959da5cdd195c995960921b604082015260600190565b6020808252818101527f63616e6e6f742070726f67726573732066726f6d2066696e616c207374617465604082015260600190565b6020808252601690820152750c2e6e6cae8e640d8cadccee8d040dad2e6dac2e8c6d60531b604082015260600190565b6020808252601590820152746163746f72496478206f7574206f662072616e676560581b604082015260600190565b60208082526019908201527f6368616e6e656c20616c726561647920636f6e636c7564656400000000000000604082015260600190565b6020808252601390820152720c2e6e6cae8e640c8de40dcdee840dac2e8c6d606b1b604082015260600190565b60208082526019908201527f77726f6e67206e756d626572206f662073756273746174657300000000000000604082015260600190565b6020808252601290820152711d1a5b595bdd5d081b9bdd081c185cdcd95960721b604082015260600190565b60208082526019908201527f72656675746174696f6e2074696d656f75742070617373656400000000000000604082015260600190565b602080825260159082015274696e76616c6964206368616e6e656c20737461746560581b604082015260600190565b6020808252601a908201527f7369676e617475726573206c656e677468206d69736d61746368000000000000604082015260600190565b60208082526019908201527f66756e6473206c6f636b656420696e206f6c6420737461746500000000000000604082015260600190565b6020808252601690820152751d1a5b595bdd5d081b9bdd081c185cdcd959081e595d60521b604082015260600190565b60208082526018908201527f73756d206f662062616c616e636573206d69736d617463680000000000000000604082015260600190565b6020808252601c908201527f6e65772062616c616e636573206c656e677468206d69736d6174636800000000604082015260600190565b6020808252601d908201527f76657273696f6e206d75737420696e6372656d656e74206279206f6e65000000604082015260600190565b6020808252600e908201526d696e76616c696420706172616d7360901b604082015260600190565b6020808252601290820152711a5b9d985b1a590818da185b9b995b08125160721b604082015260600190565b60208082526018908201527f62616c616e636573206c656e677468206d69736d617463680000000000000000604082015260600190565b6020808252601c908201527f6f6c642062616c616e636573206c656e677468206d69736d6174636800000000604082015260600190565b6020808252600d908201526c06d75737420686176652061707609c1b604082015260600190565b6000602082526111b96020830184611c17565b6000608082526123716080830187611c17565b82810360208401526123838187611c52565b905082810360408401526123978186611c52565b91505082606083015295945050505050565b6000602082526111b96020830184611c52565b6001600160401b03968716815294861660208601529290941660408401521515606083015260ff909216608082015260a081019190915260c00190565b6001600160401b03938416815260ff929092166020830152909116604082015260600190565b6040518181016001600160401b038111828210171561243a57fe5b604052919050565b60006001600160401b0382111561245557fe5b506020908102019056fe45434453413a20696e76616c6964207369676e6174757265202773272076616c756545434453413a20696e76616c6964207369676e6174757265202776272076616c7565a26469706673582212200cc2d448b3a016ac8b4425e9ef6133ab28caa3194af8911895c36619ab98b21264736f6c63430007040033"

// DeployAdjudicator deploys a new Ethereum contract, binding an instance of Adjudicator to it.
func DeployAdjudicator(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Adjudicator, error) {
	parsed, err := abi.JSON(strings.NewReader(AdjudicatorABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AdjudicatorBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Adjudicator{AdjudicatorCaller: AdjudicatorCaller{contract: contract}, AdjudicatorTransactor: AdjudicatorTransactor{contract: contract}, AdjudicatorFilterer: AdjudicatorFilterer{contract: contract}}, nil
}

// Adjudicator is an auto generated Go binding around an Ethereum contract.
type Adjudicator struct {
	AdjudicatorCaller     // Read-only binding to the contract
	AdjudicatorTransactor // Write-only binding to the contract
	AdjudicatorFilterer   // Log filterer for contract events
}

// AdjudicatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type AdjudicatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdjudicatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AdjudicatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdjudicatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AdjudicatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdjudicatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AdjudicatorSession struct {
	Contract     *Adjudicator      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AdjudicatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AdjudicatorCallerSession struct {
	Contract *AdjudicatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// AdjudicatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AdjudicatorTransactorSession struct {
	Contract     *AdjudicatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// AdjudicatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type AdjudicatorRaw struct {
	Contract *Adjudicator // Generic contract binding to access the raw methods on
}

// AdjudicatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AdjudicatorCallerRaw struct {
	Contract *AdjudicatorCaller // Generic read-only contract binding to access the raw methods on
}

// AdjudicatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AdjudicatorTransactorRaw struct {
	Contract *AdjudicatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAdjudicator creates a new instance of Adjudicator, bound to a specific deployed contract.
func NewAdjudicator(address common.Address, backend bind.ContractBackend) (*Adjudicator, error) {
	contract, err := bindAdjudicator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Adjudicator{AdjudicatorCaller: AdjudicatorCaller{contract: contract}, AdjudicatorTransactor: AdjudicatorTransactor{contract: contract}, AdjudicatorFilterer: AdjudicatorFilterer{contract: contract}}, nil
}

// NewAdjudicatorCaller creates a new read-only instance of Adjudicator, bound to a specific deployed contract.
func NewAdjudicatorCaller(address common.Address, caller bind.ContractCaller) (*AdjudicatorCaller, error) {
	contract, err := bindAdjudicator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AdjudicatorCaller{contract: contract}, nil
}

// NewAdjudicatorTransactor creates a new write-only instance of Adjudicator, bound to a specific deployed contract.
func NewAdjudicatorTransactor(address common.Address, transactor bind.ContractTransactor) (*AdjudicatorTransactor, error) {
	contract, err := bindAdjudicator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AdjudicatorTransactor{contract: contract}, nil
}

// NewAdjudicatorFilterer creates a new log filterer instance of Adjudicator, bound to a specific deployed contract.
func NewAdjudicatorFilterer(address common.Address, filterer bind.ContractFilterer) (*AdjudicatorFilterer, error) {
	contract, err := bindAdjudicator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AdjudicatorFilterer{contract: contract}, nil
}

// bindAdjudicator binds a generic wrapper to an already deployed contract.
func bindAdjudicator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AdjudicatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Adjudicator *AdjudicatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Adjudicator.Contract.AdjudicatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Adjudicator *AdjudicatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Adjudicator.Contract.AdjudicatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Adjudicator *AdjudicatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Adjudicator.Contract.AdjudicatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Adjudicator *AdjudicatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Adjudicator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Adjudicator *AdjudicatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Adjudicator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Adjudicator *AdjudicatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Adjudicator.Contract.contract.Transact(opts, method, params...)
}

// ChannelID is a free data retrieval call binding the contract method 0xa1ee1592.
//
// Solidity: function channelID((uint256,uint256,address,address[]) params) pure returns(bytes32)
func (_Adjudicator *AdjudicatorCaller) ChannelID(opts *bind.CallOpts, params ChannelParams) ([32]byte, error) {
	var out []interface{}
	err := _Adjudicator.contract.Call(opts, &out, "channelID", params)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ChannelID is a free data retrieval call binding the contract method 0xa1ee1592.
//
// Solidity: function channelID((uint256,uint256,address,address[]) params) pure returns(bytes32)
func (_Adjudicator *AdjudicatorSession) ChannelID(params ChannelParams) ([32]byte, error) {
	return _Adjudicator.Contract.ChannelID(&_Adjudicator.CallOpts, params)
}

// ChannelID is a free data retrieval call binding the contract method 0xa1ee1592.
//
// Solidity: function channelID((uint256,uint256,address,address[]) params) pure returns(bytes32)
func (_Adjudicator *AdjudicatorCallerSession) ChannelID(params ChannelParams) ([32]byte, error) {
	return _Adjudicator.Contract.ChannelID(&_Adjudicator.CallOpts, params)
}

// Disputes is a free data retrieval call binding the contract method 0x11be1997.
//
// Solidity: function disputes(bytes32 ) view returns(uint64 timeout, uint64 challengeDuration, uint64 version, bool hasApp, uint8 phase, bytes32 stateHash)
func (_Adjudicator *AdjudicatorCaller) Disputes(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Timeout           uint64
	ChallengeDuration uint64
	Version           uint64
	HasApp            bool
	Phase             uint8
	StateHash         [32]byte
}, error) {
	var out []interface{}
	err := _Adjudicator.contract.Call(opts, &out, "disputes", arg0)

	outstruct := new(struct {
		Timeout           uint64
		ChallengeDuration uint64
		Version           uint64
		HasApp            bool
		Phase             uint8
		StateHash         [32]byte
	})

	outstruct.Timeout = out[0].(uint64)
	outstruct.ChallengeDuration = out[1].(uint64)
	outstruct.Version = out[2].(uint64)
	outstruct.HasApp = out[3].(bool)
	outstruct.Phase = out[4].(uint8)
	outstruct.StateHash = out[5].([32]byte)

	return *outstruct, err

}

// Disputes is a free data retrieval call binding the contract method 0x11be1997.
//
// Solidity: function disputes(bytes32 ) view returns(uint64 timeout, uint64 challengeDuration, uint64 version, bool hasApp, uint8 phase, bytes32 stateHash)
func (_Adjudicator *AdjudicatorSession) Disputes(arg0 [32]byte) (struct {
	Timeout           uint64
	ChallengeDuration uint64
	Version           uint64
	HasApp            bool
	Phase             uint8
	StateHash         [32]byte
}, error) {
	return _Adjudicator.Contract.Disputes(&_Adjudicator.CallOpts, arg0)
}

// Disputes is a free data retrieval call binding the contract method 0x11be1997.
//
// Solidity: function disputes(bytes32 ) view returns(uint64 timeout, uint64 challengeDuration, uint64 version, bool hasApp, uint8 phase, bytes32 stateHash)
func (_Adjudicator *AdjudicatorCallerSession) Disputes(arg0 [32]byte) (struct {
	Timeout           uint64
	ChallengeDuration uint64
	Version           uint64
	HasApp            bool
	Phase             uint8
	StateHash         [32]byte
}, error) {
	return _Adjudicator.Contract.Disputes(&_Adjudicator.CallOpts, arg0)
}

// HashState is a free data retrieval call binding the contract method 0xa83a5cc5.
//
// Solidity: function hashState((bytes32,uint64,(address[],uint256[][],(bytes32,uint256[])[]),bytes,bool) state) pure returns(bytes32)
func (_Adjudicator *AdjudicatorCaller) HashState(opts *bind.CallOpts, state ChannelState) ([32]byte, error) {
	var out []interface{}
	err := _Adjudicator.contract.Call(opts, &out, "hashState", state)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashState is a free data retrieval call binding the contract method 0xa83a5cc5.
//
// Solidity: function hashState((bytes32,uint64,(address[],uint256[][],(bytes32,uint256[])[]),bytes,bool) state) pure returns(bytes32)
func (_Adjudicator *AdjudicatorSession) HashState(state ChannelState) ([32]byte, error) {
	return _Adjudicator.Contract.HashState(&_Adjudicator.CallOpts, state)
}

// HashState is a free data retrieval call binding the contract method 0xa83a5cc5.
//
// Solidity: function hashState((bytes32,uint64,(address[],uint256[][],(bytes32,uint256[])[]),bytes,bool) state) pure returns(bytes32)
func (_Adjudicator *AdjudicatorCallerSession) HashState(state ChannelState) ([32]byte, error) {
	return _Adjudicator.Contract.HashState(&_Adjudicator.CallOpts, state)
}

// Conclude is a paid mutator transaction binding the contract method 0x8bba7507.
//
// Solidity: function conclude((uint256,uint256,address,address[]) params, (bytes32,uint64,(address[],uint256[][],(bytes32,uint256[])[]),bytes,bool) state, (bytes32,uint64,(address[],uint256[][],(bytes32,uint256[])[]),bytes,bool)[] subStates) returns()
func (_Adjudicator *AdjudicatorTransactor) Conclude(opts *bind.TransactOpts, params ChannelParams, state ChannelState, subStates []ChannelState) (*types.Transaction, error) {
	return _Adjudicator.contract.Transact(opts, "conclude", params, state, subStates)
}

// Conclude is a paid mutator transaction binding the contract method 0x8bba7507.
//
// Solidity: function conclude((uint256,uint256,address,address[]) params, (bytes32,uint64,(address[],uint256[][],(bytes32,uint256[])[]),bytes,bool) state, (bytes32,uint64,(address[],uint256[][],(bytes32,uint256[])[]),bytes,bool)[] subStates) returns()
func (_Adjudicator *AdjudicatorSession) Conclude(params ChannelParams, state ChannelState, subStates []ChannelState) (*types.Transaction, error) {
	return _Adjudicator.Contract.Conclude(&_Adjudicator.TransactOpts, params, state, subStates)
}

// Conclude is a paid mutator transaction binding the contract method 0x8bba7507.
//
// Solidity: function conclude((uint256,uint256,address,address[]) params, (bytes32,uint64,(address[],uint256[][],(bytes32,uint256[])[]),bytes,bool) state, (bytes32,uint64,(address[],uint256[][],(bytes32,uint256[])[]),bytes,bool)[] subStates) returns()
func (_Adjudicator *AdjudicatorTransactorSession) Conclude(params ChannelParams, state ChannelState, subStates []ChannelState) (*types.Transaction, error) {
	return _Adjudicator.Contract.Conclude(&_Adjudicator.TransactOpts, params, state, subStates)
}

// ConcludeFinal is a paid mutator transaction binding the contract method 0x6bbf706a.
//
// Solidity: function concludeFinal((uint256,uint256,address,address[]) params, (bytes32,uint64,(address[],uint256[][],(bytes32,uint256[])[]),bytes,bool) state, bytes[] sigs) returns()
func (_Adjudicator *AdjudicatorTransactor) ConcludeFinal(opts *bind.TransactOpts, params ChannelParams, state ChannelState, sigs [][]byte) (*types.Transaction, error) {
	return _Adjudicator.contract.Transact(opts, "concludeFinal", params, state, sigs)
}

// ConcludeFinal is a paid mutator transaction binding the contract method 0x6bbf706a.
//
// Solidity: function concludeFinal((uint256,uint256,address,address[]) params, (bytes32,uint64,(address[],uint256[][],(bytes32,uint256[])[]),bytes,bool) state, bytes[] sigs) returns()
func (_Adjudicator *AdjudicatorSession) ConcludeFinal(params ChannelParams, state ChannelState, sigs [][]byte) (*types.Transaction, error) {
	return _Adjudicator.Contract.ConcludeFinal(&_Adjudicator.TransactOpts, params, state, sigs)
}

// ConcludeFinal is a paid mutator transaction binding the contract method 0x6bbf706a.
//
// Solidity: function concludeFinal((uint256,uint256,address,address[]) params, (bytes32,uint64,(address[],uint256[][],(bytes32,uint256[])[]),bytes,bool) state, bytes[] sigs) returns()
func (_Adjudicator *AdjudicatorTransactorSession) ConcludeFinal(params ChannelParams, state ChannelState, sigs [][]byte) (*types.Transaction, error) {
	return _Adjudicator.Contract.ConcludeFinal(&_Adjudicator.TransactOpts, params, state, sigs)
}

// Progress is a paid mutator transaction binding the contract method 0x36995831.
//
// Solidity: function progress((uint256,uint256,address,address[]) params, (bytes32,uint64,(address[],uint256[][],(bytes32,uint256[])[]),bytes,bool) stateOld, (bytes32,uint64,(address[],uint256[][],(bytes32,uint256[])[]),bytes,bool) state, uint256 actorIdx, bytes sig) returns()
func (_Adjudicator *AdjudicatorTransactor) Progress(opts *bind.TransactOpts, params ChannelParams, stateOld ChannelState, state ChannelState, actorIdx *big.Int, sig []byte) (*types.Transaction, error) {
	return _Adjudicator.contract.Transact(opts, "progress", params, stateOld, state, actorIdx, sig)
}

// Progress is a paid mutator transaction binding the contract method 0x36995831.
//
// Solidity: function progress((uint256,uint256,address,address[]) params, (bytes32,uint64,(address[],uint256[][],(bytes32,uint256[])[]),bytes,bool) stateOld, (bytes32,uint64,(address[],uint256[][],(bytes32,uint256[])[]),bytes,bool) state, uint256 actorIdx, bytes sig) returns()
func (_Adjudicator *AdjudicatorSession) Progress(params ChannelParams, stateOld ChannelState, state ChannelState, actorIdx *big.Int, sig []byte) (*types.Transaction, error) {
	return _Adjudicator.Contract.Progress(&_Adjudicator.TransactOpts, params, stateOld, state, actorIdx, sig)
}

// Progress is a paid mutator transaction binding the contract method 0x36995831.
//
// Solidity: function progress((uint256,uint256,address,address[]) params, (bytes32,uint64,(address[],uint256[][],(bytes32,uint256[])[]),bytes,bool) stateOld, (bytes32,uint64,(address[],uint256[][],(bytes32,uint256[])[]),bytes,bool) state, uint256 actorIdx, bytes sig) returns()
func (_Adjudicator *AdjudicatorTransactorSession) Progress(params ChannelParams, stateOld ChannelState, state ChannelState, actorIdx *big.Int, sig []byte) (*types.Transaction, error) {
	return _Adjudicator.Contract.Progress(&_Adjudicator.TransactOpts, params, stateOld, state, actorIdx, sig)
}

// Register is a paid mutator transaction binding the contract method 0x170e6715.
//
// Solidity: function register((uint256,uint256,address,address[]) params, (bytes32,uint64,(address[],uint256[][],(bytes32,uint256[])[]),bytes,bool) state, bytes[] sigs) returns()
func (_Adjudicator *AdjudicatorTransactor) Register(opts *bind.TransactOpts, params ChannelParams, state ChannelState, sigs [][]byte) (*types.Transaction, error) {
	return _Adjudicator.contract.Transact(opts, "register", params, state, sigs)
}

// Register is a paid mutator transaction binding the contract method 0x170e6715.
//
// Solidity: function register((uint256,uint256,address,address[]) params, (bytes32,uint64,(address[],uint256[][],(bytes32,uint256[])[]),bytes,bool) state, bytes[] sigs) returns()
func (_Adjudicator *AdjudicatorSession) Register(params ChannelParams, state ChannelState, sigs [][]byte) (*types.Transaction, error) {
	return _Adjudicator.Contract.Register(&_Adjudicator.TransactOpts, params, state, sigs)
}

// Register is a paid mutator transaction binding the contract method 0x170e6715.
//
// Solidity: function register((uint256,uint256,address,address[]) params, (bytes32,uint64,(address[],uint256[][],(bytes32,uint256[])[]),bytes,bool) state, bytes[] sigs) returns()
func (_Adjudicator *AdjudicatorTransactorSession) Register(params ChannelParams, state ChannelState, sigs [][]byte) (*types.Transaction, error) {
	return _Adjudicator.Contract.Register(&_Adjudicator.TransactOpts, params, state, sigs)
}

// AdjudicatorChannelUpdateIterator is returned from FilterChannelUpdate and is used to iterate over the raw logs and unpacked data for ChannelUpdate events raised by the Adjudicator contract.
type AdjudicatorChannelUpdateIterator struct {
	Event *AdjudicatorChannelUpdate // Event containing the contract specifics and raw log

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
func (it *AdjudicatorChannelUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdjudicatorChannelUpdate)
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
		it.Event = new(AdjudicatorChannelUpdate)
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
func (it *AdjudicatorChannelUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdjudicatorChannelUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdjudicatorChannelUpdate represents a ChannelUpdate event raised by the Adjudicator contract.
type AdjudicatorChannelUpdate struct {
	ChannelID [32]byte
	Version   uint64
	Phase     uint8
	Timeout   uint64
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterChannelUpdate is a free log retrieval operation binding the contract event 0x895ef5a5fc3efd313a300b006d6ce97ff0670dfe04f6eea90417edf924fa786b.
//
// Solidity: event ChannelUpdate(bytes32 indexed channelID, uint64 version, uint8 phase, uint64 timeout)
func (_Adjudicator *AdjudicatorFilterer) FilterChannelUpdate(opts *bind.FilterOpts, channelID [][32]byte) (*AdjudicatorChannelUpdateIterator, error) {

	var channelIDRule []interface{}
	for _, channelIDItem := range channelID {
		channelIDRule = append(channelIDRule, channelIDItem)
	}

	logs, sub, err := _Adjudicator.contract.FilterLogs(opts, "ChannelUpdate", channelIDRule)
	if err != nil {
		return nil, err
	}
	return &AdjudicatorChannelUpdateIterator{contract: _Adjudicator.contract, event: "ChannelUpdate", logs: logs, sub: sub}, nil
}

// WatchChannelUpdate is a free log subscription operation binding the contract event 0x895ef5a5fc3efd313a300b006d6ce97ff0670dfe04f6eea90417edf924fa786b.
//
// Solidity: event ChannelUpdate(bytes32 indexed channelID, uint64 version, uint8 phase, uint64 timeout)
func (_Adjudicator *AdjudicatorFilterer) WatchChannelUpdate(opts *bind.WatchOpts, sink chan<- *AdjudicatorChannelUpdate, channelID [][32]byte) (event.Subscription, error) {

	var channelIDRule []interface{}
	for _, channelIDItem := range channelID {
		channelIDRule = append(channelIDRule, channelIDItem)
	}

	logs, sub, err := _Adjudicator.contract.WatchLogs(opts, "ChannelUpdate", channelIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdjudicatorChannelUpdate)
				if err := _Adjudicator.contract.UnpackLog(event, "ChannelUpdate", log); err != nil {
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

// ParseChannelUpdate is a log parse operation binding the contract event 0x895ef5a5fc3efd313a300b006d6ce97ff0670dfe04f6eea90417edf924fa786b.
//
// Solidity: event ChannelUpdate(bytes32 indexed channelID, uint64 version, uint8 phase, uint64 timeout)
func (_Adjudicator *AdjudicatorFilterer) ParseChannelUpdate(log types.Log) (*AdjudicatorChannelUpdate, error) {
	event := new(AdjudicatorChannelUpdate)
	if err := _Adjudicator.contract.UnpackLog(event, "ChannelUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
