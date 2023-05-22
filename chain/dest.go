package chain

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

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
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
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

	auth := bind.NewKeyedTransactor(dest.PrivKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	//TODO call contract crossController.sol::crossFrom
	return nil
}
