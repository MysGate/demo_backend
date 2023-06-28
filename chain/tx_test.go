package chain

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/MysGate/demo_backend/contracts"
	"github.com/MysGate/demo_backend/core"
	"github.com/MysGate/demo_backend/model"
	"github.com/MysGate/demo_backend/util"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/crypto/sha3"
)

var crossContractOwnerKey = "283398d001e28892198fecb98ae86c38c10c6e64de883e6ffb86359d3f5e7a99"
var crossContractOwnerAddr = common.HexToAddress("0x24b4F2A3A30fe6e7bB060573D2C7C538BB242C8C")

var bobAddr = common.HexToAddress("0x327E046E0799b704517dF415d1AfE03bA11021cc")
var aliceKey = ""
var aliceAddr = common.HexToAddress("0x61249E8d708d5A8b0d46673242fBD03D96D94b4F")

var arbContract = common.HexToAddress("0xfcFe5d2e0842f14290Df274Bfc26FddA76027f2e")
var arbUsdtAddr = common.HexToAddress("0x1C056c622395cA50c5aC10001C6fFa91B316DfD8")
var arbPorter = crossContractOwnerAddr

var scrollContract = common.HexToAddress("0x89Cb7B97593a9D48F28CF5ff88905fA32faa0755")
var scrollUsdtAddr = common.HexToAddress("0xc6cE82a8670be7c78D56aC46Dc9f557251Cd5465")
var scrollPorter = crossContractOwnerAddr

var arbChainID = big.NewInt(421613)
var scrollCHainID = big.NewInt(534353)

var arbRpc = "https://arbitrum-goerli.publicnode.com"
var scrollRpc = "https://scroll-alphanet.blastapi.io/57226f34-917f-4e5d-84ed-76f527dac6ce"

func ScrollToPay() {
	opts, client, err := CreateTrxOpts(scrollRpc, crossContractOwnerKey, scrollCHainID, crossContractOwnerAddr)
	if err != nil {
		return
	}
	orderId, err := util.GenerateIncreaseID()
	if err != nil {
		return
	}
	fmt.Println("scroll orderId: ", orderId)
	o := &contracts.CrossControllerOrder{
		OrderId:     big.NewInt(orderId),
		SrcChainId:  arbChainID,
		SrcAddress:  crossContractOwnerAddr,
		SrcToken:    arbUsdtAddr,
		SrcAmount:   util.ConvertFloat64ToTokenAmount(30, 18),
		DestChainId: scrollCHainID,
		DestAddress: bobAddr,
		DestToken:   scrollUsdtAddr,
		Porter:      scrollPorter,
	}
	destAmount := util.ConvertFloat64ToTokenAmount(10, 18)
	fmt.Println("scroll destAmount: ", destAmount)
	instance, err := contracts.NewCrossTransactor(scrollContract, client)
	if err != nil {
		return
	}
	tx, err := instance.CrossFrom(opts, *o, 18, destAmount)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(tx.Hash().Hex())
}

func ArbToCommit() {
	opts, client, err := CreateTrxOpts(arbRpc, crossContractOwnerKey, arbChainID, crossContractOwnerAddr)
	if err != nil {
		return
	}
	c := &contracts.CrossControllerReceipt{
		DestTxHash: common.HexToHash("0x8803fbc06b9b6ae484a4a29ba579f81fb47f5d878a59aa89721391fc2d7bdc25"),
	}
	instance, err := contracts.NewCross(arbContract, client)
	if err != nil {
		return
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
			fmt.Println("receipt.Status:", receipt.Status)
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

func CreateTrxOpts(rpc string, keyStr string, chainId *big.Int, caller common.Address) (opts *bind.TransactOpts, client *ethclient.Client, err error) {
	client, err = ethclient.Dial(rpc)
	if err != nil {
		return nil, nil, err
	}
	key, err := crypto.HexToECDSA(keyStr)
	if err != nil {
		return nil, nil, err
	}
	nonce, err := client.PendingNonceAt(context.Background(), caller)
	if err != nil {
		errMsg := fmt.Sprintf("CreateTransactionOpts:client.PendingNonceAt err: %+v", err)
		util.Logger().Error(errMsg)
		return nil, nil, err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		errMsg := fmt.Sprintf("CreateTransactionOpts:client.SuggestGasPrice err: %+v", err)
		util.Logger().Error(errMsg)
		return nil, nil, err
	}

	opts, err = bind.NewKeyedTransactorWithChainID(key, chainId)
	if err != nil {
		errMsg := fmt.Sprintf("CreateTransactionOpts:NewKeyedTransactorWithChainID err: %+v", err)
		util.Logger().Error(errMsg)
		return nil, nil, err
	}

	opts.Nonce = big.NewInt(int64(nonce))
	opts.Value = big.NewInt(0) // in wei
	opts.GasLimit = uint64(0)  // in units
	opts.GasPrice = new(big.Int).Mul(gasPrice, big.NewInt(2))

	return opts, client, nil
}

type JSONData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Order   model.Order `json:"data"`
}

func getOrder(orderid string) *model.Order {
	d := &JSONData{}
	url := fmt.Sprintf("http://demoapi.mysgate.xyz/order/search?orderid=%s", orderid)
	headers := make(map[string]string)
	headers["Content-Type"] = " application/json"
	content := []byte("")
	hc := util.GetHTTPClient()
	body, err := util.HTTPReq("GET", url, hc, content, headers)
	if err != nil {
		return &d.Order
	}

	err = json.Unmarshal(body, &d)
	if err != nil {
		return &d.Order
	}

	return &d.Order
}

func printOrder(o *model.Order) {
	fmt.Println("OrderSucceed, order detail:#################")
	fmt.Println("OrderID: ", o.ID)
	fmt.Println("SrcChainId: ", o.SrcChainId)
	fmt.Println("SrcAddress: ", o.SrcAddress)
	fmt.Println("SrcToken: ", o.SrcToken)
	fmt.Println("SrcAmount: ", o.SrcAmount)
	fmt.Println("SrcTxHash: ", o.SrcTxHash)
	fmt.Println("DestAmount: ", o.DestAmount)
	fmt.Println("DestChain: ", o.DestChainId)
	fmt.Println("DestAddress: ", o.DestAddress)
	fmt.Println("DestToken: ", o.DestToken)
	fmt.Println("DestTxHash: ", o.DestTxHash)
	fmt.Println("ZkProof: ", o.Proof)
	fmt.Println("Stats: ", o.Status)
	fmt.Println("FinishTime: ", o.FinishedTime)
	fmt.Println("#############################################")
}

func ArbsToScroll(amount float64) {
	opts, client, err := CreateTrxOpts(arbRpc, aliceKey, arbChainID, aliceAddr)
	if err != nil {
		return
	}
	orderId, err := util.GenerateIncreaseID()
	if err != nil {
		return
	}
	fmt.Println("arb orderId: ", orderId)
	o := &contracts.CrossControllerOrder{
		OrderId:     big.NewInt(orderId),
		SrcChainId:  arbChainID,
		SrcAddress:  aliceAddr,
		SrcToken:    arbUsdtAddr,
		SrcAmount:   util.ConvertFloat64ToTokenAmount(amount, 18),
		DestChainId: scrollCHainID,
		DestAddress: bobAddr,
		DestToken:   scrollUsdtAddr,
		Porter:      arbPorter,
	}
	instance, err := contracts.NewCrossTransactor(arbContract, client)
	if err != nil {
		return
	}
	tx, err := instance.CrossTo(opts, *o)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("arb crossTo tx: ", tx.Hash().Hex())
	_, _, err = util.TxWaitToSync(context.Background(), client, tx)
	if err != nil {
		fmt.Println(err)
		return
	}

	ticker := time.NewTicker(1 * time.Second)
	for range ticker.C {
		order := getOrder(o.OrderId.String())
		switch order.Status {
		case int(core.CrossTo):
			fmt.Println("Source chain transaction!")
		case int(core.CrossFrom):
			fmt.Println("Destination chain transaction!")
		case int(core.AddCommitment):
			fmt.Println("Add Commitment for gen Merkel tree!")
		case int(core.Generate):
			fmt.Println("Generate ZK Proof!")
		case int(core.CommitReceipt):
			fmt.Println("CommitReceipt!")
		case int(core.Success):
			fmt.Println("Order Success!")
			printOrder(order)
			ticker.Stop()
			return
		default:
			fmt.Println("err status")
		}
	}
}

func Test_cross_bridge(t *testing.T) {
	Approve(arbRpc, arbUsdtAddr, arbContract)
	ArbsToScroll(101)
	// ScrollToPay()
	// ArbToCommit()
	// GetGasLimit()
}
