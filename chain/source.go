package chain

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"strings"

	"github.com/MysGate/demo_backend/contracts"
	"github.com/MysGate/demo_backend/core"
	"github.com/MysGate/demo_backend/model"
	"github.com/MysGate/demo_backend/pubsub"
	"github.com/MysGate/demo_backend/util"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-xorm/xorm"
)

type SrcChainHandler struct {
	Db              *xorm.Engine
	PrivKey         *ecdsa.PrivateKey
	WssClient       *ethclient.Client
	HttpClient      *ethclient.Client
	QuitListen      chan bool
	ContractAddress common.Address
	Caller          common.Address
	disp            IDispatcher
	keys            []string
}

func NewSrcChainHandler(wssClient *ethclient.Client, httpClient *ethclient.Client, addr common.Address, key *ecdsa.PrivateKey, db *xorm.Engine, disp IDispatcher, keys []string) *SrcChainHandler {
	cch := &SrcChainHandler{
		WssClient:       wssClient,
		HttpClient:      httpClient,
		PrivKey:         key,
		ContractAddress: addr,
		QuitListen:      make(chan bool, 10),
		Db:              db,
		disp:            disp,
		keys:            keys,
	}

	publicKey := key.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		util.Logger().Error("NewDestChainHandler key err")
		return nil
	}

	caller := crypto.PubkeyToAddress(*publicKeyECDSA)
	cch.Caller = caller
	return cch
}

func (sch *SrcChainHandler) close() {
	sch.QuitListen <- true
}

var orderCrossToTopic string = util.GetCrossToTopic()

func (sch *SrcChainHandler) runListenEvent() {
	logs := make(chan interface{}, 10000000)
	m := pubsub.GetSubscribeManager()
	for _, k := range sch.keys {
		m.AddSubscribe(k, logs)
	}

	for {
		select {
		case vLog := <-logs:
			sch.DispatchEvent(vLog)
		case <-sch.QuitListen:
			return
		}
	}
}
func (sch *SrcChainHandler) DispatchEvent(v interface{}) {
	vLog, ok := v.(*types.Log)
	if !ok {
		util.Logger().Error("DispatchEvent input type mismatch")
		return
	}

	if vLog.Topics[0].Hex() != orderCrossToTopic {
		return
	}
	order, succeed := sch.parseCrossToEvent(vLog)
	if !succeed || order == nil {
		util.Logger().Error(fmt.Sprintf("DispatchEvent parseEvent failed: %+v", vLog))
		return
	}

	sch.disp.PayForDest(order)
}

func (sch *SrcChainHandler) parseCrossToEvent(vLog *types.Log) (*model.Order, bool) {
	contractAbi, err := abi.JSON(strings.NewReader(string(contracts.CrossABI)))
	if err != nil {
		util.Logger().Error(fmt.Sprintf("Not found abi json, err:%+v", err))
		return nil, false
	}

	orderEvent := &contracts.CrossCrossTo{}
	err = contractAbi.UnpackIntoInterface(orderEvent, "CrossTo", vLog.Data)
	if err != nil {
		util.Logger().Error(fmt.Sprintf("[Order] failed to UnpackIntoInterface: %+v", err))
		return nil, false
	}

	order := &model.Order{
		ID:          orderEvent.Order.OrderId.String(),
		SrcAddress:  orderEvent.Order.SrcAddress.Hex(),
		SrcAmount:   util.ConvertTokenAmountToFloat64(orderEvent.Order.SrcAmount.String(), 18),
		SrcToken:    orderEvent.Order.SrcToken.Hex(),
		SrcChainId:  orderEvent.Order.SrcChainId.Uint64(),
		DestAddress: orderEvent.Order.DestAddress.Hex(),
		DestChainId: orderEvent.Order.DestChainId.Uint64(),
		DestToken:   orderEvent.Order.DestAddress.Hex(),
		DestAmount:  util.ConvertTokenAmountToFloat64(orderEvent.CrossAmount.String(), 18),
		PoterId:     orderEvent.Order.PorterPool.Hex(),
		FixedFee:    util.ConvertTokenAmountToFloat64(orderEvent.FixedFeeAmount.String(), 18),
		FloatFee:    util.ConvertTokenAmountToFloat64(orderEvent.FloatFeeAmount.String(), 18),
		Status:      core.CrossTo,
	}

	srcChainId, _ := sch.HttpClient.NetworkID(context.Background())
	order.SrcChainId = srcChainId.Uint64()
	model.InsertOrder(order, sch.Db)
	return order, true
}

func (sch *SrcChainHandler) commitReceipt(order *model.Order) error {
	nonce, err := sch.HttpClient.PendingNonceAt(context.Background(), sch.Caller)
	if err != nil {
		errMsg := fmt.Sprintf("commitReceipt:sch.HttpClient.PendingNonceAt err: %+v", err)
		util.Logger().Error(errMsg)
		return err
	}

	gasPrice, err := sch.HttpClient.SuggestGasPrice(context.Background())
	if err != nil {
		errMsg := fmt.Sprintf("commitReceipt:sch.HttpClient.SuggestGasPrice err: %+v", err)
		util.Logger().Error(errMsg)
		return err
	}

	srcChainID := big.NewInt(int64(order.SrcChainId))
	opts, err := bind.NewKeyedTransactorWithChainID(sch.PrivKey, srcChainID)
	if err != nil {
		errMsg := fmt.Sprintf("commitReceipt:NewKeyedTransactorWithChainID err: %+v", err)
		util.Logger().Error(errMsg)
		return err
	}

	opts.Nonce = big.NewInt(int64(nonce))
	opts.Value = big.NewInt(0)     // in wei
	opts.GasLimit = uint64(300000) // in units
	opts.GasPrice = gasPrice

	instance, err := contracts.NewCrossTransactor(sch.ContractAddress, sch.HttpClient)
	if err != nil {
		util.Logger().Error(fmt.Sprintf("commitReceipt: create instance err:%+v", err))
		return err
	}
	orderId, _ := new(big.Int).SetString(order.ID, 10)
	contractOrder := &contracts.CrossControllerOrder{
		OrderId:     orderId,
		SrcChainId:  new(big.Int).SetUint64(order.SrcChainId),
		SrcAddress:  common.HexToAddress(order.SrcAddress),
		SrcToken:    common.HexToAddress(order.SrcToken),
		SrcAmount:   util.ConvertFloat64ToTokenAmount(order.SrcAmount, 18),
		DestChainId: new(big.Int).SetUint64(order.DestChainId),
		DestAddress: common.HexToAddress(order.DestAddress),
		DestToken:   common.HexToAddress(order.DestToken),
		PorterPool:  common.HexToAddress(order.PoterId),
	}
	orderHash := model.Keccak256EncodePackedContractOrder(contractOrder)

	hash := common.HexToHash(order.DestTxHash)
	receipt := &contracts.CrossControllerReceipt{}
	copy(receipt.DestTxHash[:], hash.Bytes())
	// TODO add proof
	// receipt.Proof
	tx, err := instance.CommitReceipt(opts, orderHash, *receipt)
	if err != nil {
		errMsg := fmt.Sprintf("commitReceipt:instance.CommitReceipt err: %+v", err)
		util.Logger().Error(errMsg)
		return err
	}

	order.ReceiptTxHash = tx.Hash().Hex()

	return nil
}
