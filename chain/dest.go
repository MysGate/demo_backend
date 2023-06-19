package chain

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/MysGate/demo_backend/contracts"
	"github.com/MysGate/demo_backend/model"
	"github.com/MysGate/demo_backend/util"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type DestChainHandler struct {
	HttpClient      *ethclient.Client
	ContractAddress common.Address
	PrivKey         *ecdsa.PrivateKey
	Caller          common.Address
}

func NewDestChainHandler(client *ethclient.Client, addr common.Address, key *ecdsa.PrivateKey) *DestChainHandler {
	dch := &DestChainHandler{
		HttpClient:      client,
		ContractAddress: addr,
		PrivKey:         key,
	}

	publicKey := key.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		util.Logger().Error("NewDestChainHandler key err")
		return nil
	}

	caller := crypto.PubkeyToAddress(*publicKeyECDSA)
	dch.Caller = caller

	return dch
}

func (dest *DestChainHandler) crossFrom(order *model.Order) (bool, error) {
	opts, err := util.CreateTransactionOpts(dest.HttpClient, dest.PrivKey, order.DestChainId, dest.Caller)
	if err != nil {
		util.Logger().Error(fmt.Sprintf("crossFrom:CreateTransactionOpts err:%+v", err))
		return false, err
	}
	instance, err := contracts.NewCrossTransactor(dest.ContractAddress, dest.HttpClient)
	if err != nil {
		util.Logger().Error(fmt.Sprintf("crossFrom: create instance err:%+v", err))
		return false, err
	}

	o := &contracts.CrossControllerOrder{
		OrderId:     big.NewInt(order.ID),
		SrcChainId:  big.NewInt(int64(order.SrcChainId)),
		SrcAddress:  common.HexToAddress(order.SrcAddress),
		SrcToken:    common.HexToAddress(order.SrcToken),
		SrcAmount:   util.ConvertFloat64ToTokenAmount(order.SrcAmount, 18),
		DestChainId: big.NewInt(int64(order.DestChainId)),
		DestAddress: common.HexToAddress(order.DestAddress),
		DestToken:   common.HexToAddress(order.DestToken),
		Porter:      common.HexToAddress(order.PoterId),
	}
	destAmount := util.ConvertFloat64ToTokenAmount(order.DestAmount, 18)
	tx, err := instance.CrossFrom(opts, *o, 18, destAmount)
	if err != nil {
		errMsg := fmt.Sprintf("crossFrom:instance.CrossFrom err: %+v", err)
		util.Logger().Error(errMsg)
		return false, err
	}

	order.DestTxHash = tx.Hash().Hex()

	_, ret, err := util.TxWaitToSync(context.Background(), dest.HttpClient, tx)
	return ret, err
}
