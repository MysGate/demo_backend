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
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-xorm/xorm"
)

type SrcChainHandler struct {
	Db              *xorm.Engine
	PrivKey         *ecdsa.PrivateKey
	HttpClient      *ethclient.Client
	QuitListen      chan bool
	ContractAddress common.Address
	BridgeAddress   common.Address
	Caller          common.Address
	disp            IDispatcher
	keys            []string
}

func NewSrcChainHandler(httpClient *ethclient.Client, addr common.Address, key *ecdsa.PrivateKey, db *xorm.Engine, disp IDispatcher, keys []string) *SrcChainHandler {
	cch := &SrcChainHandler{
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

	if vLog.Topics[0].Hex() == util.GetCrossToTopic() {
		order, succeed := sch.parseCrossToEvent(vLog)
		if !succeed || order == nil {
			util.Logger().Error(fmt.Sprintf("DispatchEvent parseEvent failed: %+v", vLog))
			return
		}
		// crossfrom
		sch.disp.PayForDest(order)
	}
}

// addcommitment
func (sch *SrcChainHandler) AddCommitment(order *model.Order) (bool, error) {
	opts, err := util.CreateTransactionOpts(sch.HttpClient, sch.PrivKey, order.SrcChainId, sch.Caller)
	if err != nil {
		util.Logger().Error(fmt.Sprintf("AddCommitment:CreateTransactionOpts err:%+v", err))
		return false, err
	}

	instance, err := contracts.NewBridge(sch.BridgeAddress, sch.HttpClient)
	if err != nil {
		util.Logger().Error(fmt.Sprintf("AddCommitment: create instance err:%+v", err))
		return false, err
	}

	tx, err := instance.AddCommitment(opts, big.NewInt(order.ID))
	if err != nil {
		errMsg := fmt.Sprintf("AddCommitment:instance.AddCommitment err: %+v", err)
		util.Logger().Error(errMsg)
		return false, err
	}

	order.AddCommitmentTxHash = tx.Hash().Hex()
	receipt, ret, err := util.TxWaitToSync(context.Background(), sch.HttpClient, tx)
	if err != nil {
		errMsg := fmt.Sprintf("AddCommitment:TxWaitToSync err: %+v", err)
		util.Logger().Error(errMsg)
		return false, err
	}

	if !ret {
		errMsg := fmt.Sprintf("AddCommitment:TxWaitToSync failed, tx: %+v", tx)
		util.Logger().Error(errMsg)
		return false, err
	}
	order.AddCommitmentTxStatus = 1

	for _, log := range receipt.Logs {
		sch.parseCommitmentAddedEvent(order, log)
	}

	return true, nil
}

func (sch *SrcChainHandler) commitReceipt(order *model.Order) (bool, error) {
	opts, err := util.CreateTransactionOpts(sch.HttpClient, sch.PrivKey, order.SrcChainId, sch.Caller)
	if err != nil {
		util.Logger().Error(fmt.Sprintf("AddCommitment:CreateTransactionOpts err:%+v", err))
		return false, err
	}

	instance, err := contracts.NewCrossTransactor(sch.ContractAddress, sch.HttpClient)
	if err != nil {
		util.Logger().Error(fmt.Sprintf("commitReceipt: create instance err:%+v", err))
		return false, err
	}
	contractOrder := &contracts.CrossControllerOrder{
		OrderId:     big.NewInt(order.ID),
		SrcChainId:  new(big.Int).SetUint64(order.SrcChainId),
		SrcAddress:  common.HexToAddress(order.SrcAddress),
		SrcToken:    common.HexToAddress(order.SrcToken),
		SrcAmount:   util.ConvertFloat64ToTokenAmount(order.SrcAmount, 18),
		DestChainId: new(big.Int).SetUint64(order.DestChainId),
		DestAddress: common.HexToAddress(order.DestAddress),
		DestToken:   common.HexToAddress(order.DestToken),
		Porter:      common.HexToAddress(order.PoterId),
	}
	orderHash := model.Keccak256EncodePackedContractOrder(contractOrder)

	hash := common.HexToHash(order.DestTxHash)
	receipt := &contracts.CrossControllerReceipt{}
	copy(receipt.DestTxHash[:], hash.Bytes())
	tx, err := instance.CommitReceipt(opts, orderHash, *receipt)
	if err != nil {
		errMsg := fmt.Sprintf("commitReceipt:instance.CommitReceipt err: %+v", err)
		util.Logger().Error(errMsg)
		return false, err
	}
	order.ReceiptTxHash = tx.Hash().Hex()
	_, ret, err := util.TxWaitToSync(context.Background(), sch.HttpClient, tx)
	return ret, err
}

func (sch *SrcChainHandler) commitReceiptWithZk(order *model.Order) (bool, error) {
	opts, err := util.CreateTransactionOpts(sch.HttpClient, sch.PrivKey, order.SrcChainId, sch.Caller)
	if err != nil {
		util.Logger().Error(fmt.Sprintf("commitReceiptWithZk:CreateTransactionOpts err:%+v", err))
		return false, err
	}

	instance, err := contracts.NewCrossTransactor(sch.ContractAddress, sch.HttpClient)
	if err != nil {
		util.Logger().Error(fmt.Sprintf("commitReceiptWithZk: create instance err:%+v", err))
		return false, err
	}
	contractOrder := &contracts.CrossControllerOrder{
		OrderId:     big.NewInt(order.ID),
		SrcChainId:  new(big.Int).SetUint64(order.SrcChainId),
		SrcAddress:  common.HexToAddress(order.SrcAddress),
		SrcToken:    common.HexToAddress(order.SrcToken),
		SrcAmount:   util.ConvertFloat64ToTokenAmount(order.SrcAmount, 18),
		DestChainId: new(big.Int).SetUint64(order.DestChainId),
		DestAddress: common.HexToAddress(order.DestAddress),
		DestToken:   common.HexToAddress(order.DestToken),
		Porter:      common.HexToAddress(order.PoterId),
	}
	orderHash := model.Keccak256EncodePackedContractOrder(contractOrder)

	hash := common.HexToHash(order.DestTxHash)
	receipt := &contracts.CrossControllerReceipt{}
	copy(receipt.DestTxHash[:], hash.Bytes())

	proof := &contracts.CrossControllerProof{
		A: order.ZKP.Proof.A,
		B: order.ZKP.Proof.B,
		C: order.ZKP.Proof.C,
	}
	input := order.ZKP.PublicInfo
	tx, err := instance.CommitReceiptWithZK(opts, *proof, input, orderHash, receipt.DestTxHash)
	if err != nil {
		errMsg := fmt.Sprintf("commitReceipt:instance.CommitReceipt err: %+v", err)
		util.Logger().Error(errMsg)
		return false, err
	}

	order.ReceiptTxHash = tx.Hash().Hex()
	_, ret, err := util.TxWaitToSync(context.Background(), sch.HttpClient, tx)
	return ret, err
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
		ID:          orderEvent.Order.OrderId.Int64(),
		SrcAddress:  orderEvent.Order.SrcAddress.Hex(),
		SrcAmount:   util.ConvertTokenAmountToFloat64(orderEvent.Order.SrcAmount.String(), 18),
		SrcToken:    orderEvent.Order.SrcToken.Hex(),
		SrcChainId:  orderEvent.Order.SrcChainId.Uint64(),
		DestAddress: orderEvent.Order.DestAddress.Hex(),
		DestChainId: orderEvent.Order.DestChainId.Uint64(),
		DestToken:   orderEvent.Order.DestToken.Hex(),
		DestAmount:  util.ConvertTokenAmountToFloat64(orderEvent.CrossAmount.String(), 18),
		PoterId:     orderEvent.Order.Porter.Hex(),
		FixedFee:    util.ConvertTokenAmountToFloat64(orderEvent.FixedFeeAmount.String(), 18),
		FloatFee:    util.ConvertTokenAmountToFloat64(orderEvent.FloatFeeAmount.String(), 18),
		Status:      core.CrossTo,
	}

	srcChainId, _ := sch.HttpClient.NetworkID(context.Background())
	order.SrcChainId = srcChainId.Uint64()
	order.SrcTxHash = vLog.TxHash.Hex()
	model.InsertOrder(order, sch.Db)
	return order, true
}

func (sch *SrcChainHandler) parseCommitmentAddedEvent(order *model.Order, vLog *types.Log) bool {
	contractAbi, err := abi.JSON(strings.NewReader(string(contracts.BridgeABI)))
	if err != nil {
		util.Logger().Error(fmt.Sprintf("Not found abi json, err:%+v", err))
		return false
	}

	commitmentAddedEvent := contracts.BridgeCommitmentAdded{}
	err = contractAbi.UnpackIntoInterface(commitmentAddedEvent, "CommitmentAdded", vLog.Data)
	if err != nil {
		util.Logger().Error(fmt.Sprintf("[Order] failed to UnpackIntoInterface: %+v", err))
		return false
	}

	order.CommitmentIdx = commitmentAddedEvent.LeafIndex.Uint64()
	return true
}
