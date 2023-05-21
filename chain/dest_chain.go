package chain

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type DestChainHandler struct {
	Client          *ethclient.Client
	ContractAddress common.Address
	PrivKey         *ecdsa.PrivateKey
	Caller          common.Address
}

func NewDestChainHandler(client *ethclient.Client, addr common.Address, key *ecdsa.PrivateKey) *DestChainHandler {
	dch := &DestChainHandler{
		Client:          client,
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

func (dest *DestChainHandler) crossFrom(receipt common.Address, amount float64) {
	nonce, err := dest.Client.PendingNonceAt(context.Background(), dest.Caller)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := dest.Client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(dest.PrivKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	//TODO contract call

}
