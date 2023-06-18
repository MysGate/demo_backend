package util

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/MysGate/demo_backend/contracts"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/crypto/sha3"
)

var crossContractOwnerKey = ""
var crossContractOwnerAddr = common.HexToAddress("0x24b4F2A3A30fe6e7bB060573D2C7C538BB242C8C")

var bobKey = ""
var bobAddr = common.HexToAddress("0x327E046E0799b704517dF415d1AfE03bA11021cc")
var aliceAddr = common.HexToAddress("0x61249E8d708d5A8b0d46673242fBD03D96D94b4F")
var aliceKey = ""

var arbContract = common.HexToAddress("0xF15D89a32E62A1e69C5DEa052df5511b298eb404")
var arbUsdtAddr = common.HexToAddress("0x1C056c622395cA50c5aC10001C6fFa91B316DfD8")
var arbPorter = crossContractOwnerAddr

// var arbPorter = crossContractOwnerAddr

var scrollContract = common.HexToAddress("0xEbc1e4Df0fE0790fccD2a6615D9510C913F684B1")
var scrollUsdtAddr = common.HexToAddress("0xc6cE82a8670be7c78D56aC46Dc9f557251Cd5465")
var scrollPorter = crossContractOwnerAddr

// var scrollPorter = crossContractOwnerAddr

var arbChainID = big.NewInt(421613)
var scrollCHainID = big.NewInt(534353)

var arbRpc = "https://endpoints.omniatech.io/v1/arbitrum/goerli/public"
var scrollRpc = "https://scroll-alphanet.blastapi.io/57226f34-917f-4e5d-84ed-76f527dac6ce"

func ScrollToPay() {

	// aliceAddr := common.HexToAddress("0x61249E8d708d5A8b0d46673242fBD03D96D94b4F")
	// bobKey := ""
	//"0xb0184F8c19a31e819e74c0F61b1E3ba9b1ab4634"
	// scrollPorterPoolAddrKey := "71c0cdcdf4df917b715252066c3e1cca2f81722d367a46eeac4a59ae003e48d4"
	// arbRpc := "https://arb-goerli.g.alchemy.com/v2/SoAdZnLAhCJgYXNrCprnca8pddcPNqAM"
	//wss://arb-goerli.g.alchemy.com/v2/SoAdZnLAhCJgYXNrCprnca8pddcPNqAM

	client, err := ethclient.Dial(scrollRpc)
	if err != nil {
		return
	}

	nonce, err := client.PendingNonceAt(context.Background(), crossContractOwnerAddr)
	if err != nil {
		return
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return
	}

	fmt.Println(gasPrice.String())

	key, err := crypto.HexToECDSA(crossContractOwnerKey)
	if err != nil {
		return
	}

	opts, err := bind.NewKeyedTransactorWithChainID(key, scrollCHainID)
	if err != nil {
		return
	}

	opts.Nonce = big.NewInt(int64(nonce))
	opts.Value = big.NewInt(0) // in wei
	// opts.GasLimit = 0          // in units
	opts.GasLimit = uint64(0)
	opts.GasPrice = gasPrice

	instance, err := contracts.NewCrossTransactor(scrollContract, client)
	if err != nil {
		return
	}
	orderId := int64(1685690886)
	fmt.Println("scroll orderId: ", orderId)
	o := &contracts.CrossControllerOrder{
		OrderId:     big.NewInt(orderId),
		SrcChainId:  arbChainID,
		SrcAddress:  crossContractOwnerAddr,
		SrcToken:    arbUsdtAddr,
		SrcAmount:   ConvertFloat64ToTokenAmount(30, 18),
		DestChainId: scrollCHainID,
		DestAddress: bobAddr,
		DestToken:   scrollUsdtAddr,
		Porter:      scrollPorter,
	}
	destAmount := ConvertFloat64ToTokenAmount(10, 18)
	// destAmount := big.NewInt(1)
	fmt.Println("scroll destAmount: ", destAmount)
	tx, err := instance.CrossFrom(opts, *o, 18, destAmount)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(tx.Hash().Hex())
}

func ArbsToScroll() {
	// arbRpc := "https://arb-goerli.g.alchemy.com/v2/SoAdZnLAhCJgYXNrCprnca8pddcPNqAM"
	//wss://arb-goerli.g.alchemy.com/v2/SoAdZnLAhCJgYXNrCprnca8pddcPNqAM

	client, err := ethclient.Dial(arbRpc)
	if err != nil {
		return
	}

	nonce, err := client.PendingNonceAt(context.Background(), aliceAddr)
	if err != nil {
		return
	}
	fmt.Println("alic nonce: ", nonce)

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return
	}
	//101680000
	//3235990000
	// gasPrice := big.NewInt(3235990000)
	fmt.Println("gasPrice: ", gasPrice)

	key, err := crypto.HexToECDSA(aliceKey)
	if err != nil {
		return
	}

	opts, err := bind.NewKeyedTransactorWithChainID(key, arbChainID)
	if err != nil {
		return
	}

	opts.Nonce = big.NewInt(int64(nonce))
	opts.Value = big.NewInt(0) // in wei
	opts.GasLimit = uint64(0)  // in units
	opts.GasPrice = gasPrice

	orderId := time.Now().Unix()
	fmt.Println("arb orderId: ", orderId)
	o := &contracts.CrossControllerOrder{
		OrderId:    big.NewInt(int64(orderId)),
		SrcChainId: arbChainID,
		SrcAddress: aliceAddr,
		SrcToken:   arbUsdtAddr,
		SrcAmount:  ConvertFloat64ToTokenAmount(30, 18),

		DestChainId: scrollCHainID,
		DestAddress: bobAddr,
		DestToken:   scrollUsdtAddr,
		Porter:      arbPorter,
	}
	fmt.Println(o)

	GetGasLimit(arbRpc, aliceAddr, arbContract, o)

	instance, err := contracts.NewCrossTransactor(arbContract, client)
	if err != nil {
		return
	}
	tx, err := instance.CrossTo(opts, *o)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("arb tx: ", tx.Hash().Hex())

}

func ArbToCommit() {

	// aliceKey := ""
	// arbRpc := "https://arb-goerli.g.alchemy.com/v2/SoAdZnLAhCJgYXNrCprnca8pddcPNqAM"
	//wss://arb-goerli.g.alchemy.com/v2/SoAdZnLAhCJgYXNrCprnca8pddcPNqAM

	client, err := ethclient.Dial(arbRpc)
	if err != nil {
		return
	}

	nonce, err := client.PendingNonceAt(context.Background(), crossContractOwnerAddr)
	if err != nil {
		return
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return
	}
	// gasPrice := big.NewInt(3235990000)

	key, err := crypto.HexToECDSA(crossContractOwnerKey)
	if err != nil {
		return
	}

	opts, err := bind.NewKeyedTransactorWithChainID(key, arbChainID)
	if err != nil {
		return
	}

	opts.Nonce = big.NewInt(int64(nonce))
	opts.Value = big.NewInt(0) // in wei
	opts.GasLimit = uint64(0)  // in units
	opts.GasPrice = gasPrice

	instance, err := contracts.NewCross(arbContract, client)
	if err != nil {
		return
	}

	c := &contracts.CrossControllerReceipt{
		// Proof:      common.HexToHash("success"),
		DestTxHash: common.HexToHash("0x8803fbc06b9b6ae484a4a29ba579f81fb47f5d878a59aa89721391fc2d7bdc25"),
	}
	tx, err := instance.CommitReceipt(opts, common.HexToHash("0xac09027dbe785a0a5f35bc4974fb79e811e7ada552260bf0f7164f330a88e408"), *c)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(tx.Hash().Hex())
}

func Approve(rpc string, tokenAddres common.Address, spenderAddress common.Address) {
	client, err := ethclient.Dial(rpc)
	if err != nil {
		return
	}

	privateKey, err := crypto.HexToECDSA(aliceKey)
	if err != nil {
		return
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Println("fromAddress: ", fromAddress.Hex())
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return
	}
	fmt.Println("nonce: ", nonce)

	value := big.NewInt(0) // in wei (0 eth)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return
	}
	fmt.Println("gasPrice: ", gasPrice)

	methodSignature := []byte("approve(address,uint256)")
	crypto.NewKeccakState()
	hash := sha3.NewLegacyKeccak256()

	hash.Write(methodSignature)
	methodID := hash.Sum(nil)[:4]
	fmt.Println("methodID: ", hexutil.Encode(methodID))

	amount := new(big.Int)
	amount.SetString("1000000000000000000000", 10) // 1000 tokens
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)

	var data []byte
	data = append(data, methodID...)
	data = append(data, common.LeftPadBytes(spenderAddress.Bytes(), 32)...)
	data = append(data, paddedAmount...)

	fmt.Println("data: ", common.Bytes2Hex(data))

	// gasLimit := uint64(0)
	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		From:  fromAddress,
		To:    &tokenAddres,
		Data:  data,
		Value: value,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("gasLimit: ", gasLimit)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("chainID: ", hexutil.Encode(methodID))

	tx := types.NewTransaction(nonce, tokenAddres, value, gasLimit, gasPrice, data)
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		fmt.Println("signedTx error")
		return
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return
	}
	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())

	for {
		receipt, err := client.TransactionReceipt(context.Background(), signedTx.Hash())
		if err != nil {
			fmt.Println("not found")
			time.Sleep(1 * time.Second)
			continue
		}
		if receipt.Status == 1 {
			break
		} else {
			fmt.Sprintln("receipt.Status:%d ", receipt.Status)
			break
		}
	}

}

func GetGasLimit(rpc string, fromAddress common.Address, toAddress common.Address, o *contracts.CrossControllerOrder) uint64 {
	// fromAddress := common.HexToAddress("0x61249E8d708d5A8b0d46673242fBD03D96D94b4F")
	// toAddress := common.HexToAddress("0xAc29bB4862e7865cbF228Ece1c519D4b34452fbd")
	client, err := ethclient.Dial(rpc)
	if err != nil {
		return 0
	}
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return 0
	}
	fmt.Println(nonce)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return 0
	}
	fmt.Println(gasPrice)

	methodSignature := []byte("crossTo((uint256,uint256,address,address,uint256,uint256,address,address,address))")
	crypto.NewKeccakState()
	hash := sha3.NewLegacyKeccak256()

	hash.Write(methodSignature)
	methodID := hash.Sum(nil)[:4]
	fmt.Println(hexutil.Encode(methodID)) // 0xdef66322

	// orderId := time.Now().Unix()
	// o := &contracts.CrossControllerOrder{
	// 	OrderId:    big.NewInt(int64(orderId)),
	// 	SrcChainId: arbChainID,
	// 	SrcAddress: fromAddress,
	// 	SrcToken:   arbUsdtAddr,
	// 	SrcAmount:  ConvertFloat64ToTokenAmount(30, 18),

	// 	DestChainId: scrollCHainID,
	// 	DestAddress: bobAddr,
	// 	DestToken:   scrollUsdtAddr,
	// 	PorterPool:  arbPorterPoolAddr,
	// }

	var data []byte
	data = append(data, methodID...)
	data = append(data, common.LeftPadBytes(o.OrderId.Bytes(), 32)...)
	data = append(data, common.LeftPadBytes(o.SrcChainId.Bytes(), 32)...)
	data = append(data, common.LeftPadBytes(o.SrcAddress.Bytes(), 32)...)
	data = append(data, common.LeftPadBytes(o.SrcToken.Bytes(), 32)...)
	data = append(data, common.LeftPadBytes(o.SrcAmount.Bytes(), 32)...)
	data = append(data, common.LeftPadBytes(o.DestChainId.Bytes(), 32)...)
	data = append(data, common.LeftPadBytes(o.DestAddress.Bytes(), 32)...)
	data = append(data, common.LeftPadBytes(o.DestToken.Bytes(), 32)...)
	data = append(data, common.LeftPadBytes(o.Porter.Bytes(), 32)...)
	fmt.Println(common.Bytes2Hex(data))

	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		From:  fromAddress,
		To:    &toAddress,
		Data:  data,
		Value: big.NewInt(0),
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(gasLimit) // 23256
	return gasLimit
}

func cross_test(t *testing.T) {
	// Approve(arbRpc, arbUsdtAddr, arbContract)
	ArbsToScroll()
	// ScrollToPay()
	// ArbToCommit()
	// GetGasLimit()
}
