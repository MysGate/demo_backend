package util

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type PairCreatedEvent struct {
	Pair common.Address
	Arg3 *big.Int
}

func GetPairCreatedTopic() {
	topic := []byte("PairCreated(address,address,address,uint256)")
	topicHash := crypto.Keccak256Hash(topic)
	fmt.Println(topicHash.Hex()) //0x0d3648bd0f6ba80134a33ba9275ac585d9d315f0ad8355cddefde31afa28d0e9
}

type OrderEvent struct {
	SrcAddress  common.Address
	SrcToken    common.Address
	SrcAmount   uint32
	DestChain   uint32
	DestAddress common.Address
	DestToken   common.Address
	Porters     []string
	Proof       string
	DestTxHash  string
}

func GetOrderTopic() string {
	//SrcAddress,SrcToken,SrcAmount,DestChain,DestAddress,DestToken,Porters,Proof,DestTxHash
	topic := []byte("Order(address,address,uint256,uint256,address,address,address,string,string)")
	topicHash := crypto.Keccak256Hash(topic)
	return topicHash.Hex()
}

func GetTradeTopic() string {
	topic := []byte("Trade(address,address,address,uint256)")
	topicHash := crypto.Keccak256Hash(topic)
	return topicHash.Hex()
}
