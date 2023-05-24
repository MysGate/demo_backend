package chain

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/MysGate/demo_backend/contracts"
	"github.com/MysGate/demo_backend/model"
	"github.com/MysGate/demo_backend/util"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
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

func (dest *DestChainHandler) crossFrom(order *model.Order) error {
	nonce, err := dest.HttpClient.PendingNonceAt(context.Background(), dest.Caller)
	if err != nil {
		errMsg := fmt.Sprintf("crossFrom:dest.HttpClient.PendingNonceAt err: %+v", err)
		util.Logger().Error(errMsg)
		return err
	}

	gasPrice, err := dest.HttpClient.SuggestGasPrice(context.Background())
	if err != nil {
		errMsg := fmt.Sprintf("crossFrom:dest.HttpClient.SuggestGasPrice err: %+v", err)
		util.Logger().Error(errMsg)
		return err
	}

	destChainID := big.NewInt(int64(order.DestChainId))
	opts, err := bind.NewKeyedTransactorWithChainID(dest.PrivKey, destChainID)
	if err != nil {
		errMsg := fmt.Sprintf("crossFrom:NewKeyedTransactorWithChainID err: %+v", err)
		util.Logger().Error(errMsg)
		return err
	}

	opts.Nonce = big.NewInt(int64(nonce))
	opts.Value = big.NewInt(0)     // in wei
	opts.GasLimit = uint64(300000) // in units
	opts.GasPrice = gasPrice

	instance, err := contracts.NewCrossTransactor(dest.ContractAddress, dest.HttpClient)
	if err != nil {
		util.Logger().Error(fmt.Sprintf("crossFrom: create instance err:%+v", err))
		return err
	}
	o := &contracts.CrossControllerOrder{
		OrderId:     big.NewInt(order.ID),
		SrcChainId:  big.NewInt(int64(order.SrcChainId)),
		SrcAddress:  common.HexToAddress(order.SrcAddress),
		SrcToken:    common.HexToAddress(order.SrcToken),
		SrcAmount:   util.ConvertFloat64ToTokenAmount(order.SrcAmount, 18),
		DestChainId: destChainID,
		DestAddress: common.HexToAddress(order.DestAddress),
		DestToken:   common.HexToAddress(order.DestToken),
		PorterPool:  common.HexToAddress(order.PoterId),
	}
	destAmount := util.ConvertFloat64ToTokenAmount(order.DestAmount, 18)
	tx, err := instance.CrossFrom(opts, *o, 18, destAmount)
	if err != nil {
		errMsg := fmt.Sprintf("crossFrom:instance.CrossFrom err: %+v", err)
		util.Logger().Error(errMsg)
		return err
	}

	order.DestTxHash = tx.Hash().Hex()
	return nil
}
