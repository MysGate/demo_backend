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

const (
	orderToTopic            string = "0xeb354ff2ff6b3d6392f3c14565a5e0c60fc642b456cd2538e94968fbc54467e8"
	orderFromTopic          string = "0x104f0c1d6ebbba9acf834bd5f27d78481d562d83159d076b974d16bca9c66c21"
	orderCommitReceiptTopic string = "0x581db44feed8ab7f2b0e591fd633c1326a4ba3ea20a5c346ab38fd1f42208e81"
)

func FindOrderEventTopic(topic string) string {
	var orderEvent string
	switch topic {
	case orderToTopic:
		orderEvent = "CrossTo"
	case orderFromTopic:
		orderEvent = "CrossFrom"
	case orderCommitReceiptTopic:
		orderEvent = "CommitReceipt"
	default:
		orderEvent = ""
	}
	return orderEvent
}

// event CrossTo 0xeb354ff2ff6b3d6392f3c14565a5e0c60fc642b456cd2538e94968fbc54467e8
func GetCrossToTopic() string {
	return orderToTopic
	// topic := []byte("CrossTo(address indexed account, (uint256,uint256,address,address,uint256,uint256,address,address,address) order, uint256 fixedFeeAmount, uint256 floatFeeAmount, uint256 crossAmount)")
	// topicHash := crypto.Keccak256Hash(topic)
	// return topicHash.Hex()
}

// event CrossFrom 0x104f0c1d6ebbba9acf834bd5f27d78481d562d83159d076b974d16bca9c66c21
func GetCrossFromTopic() string {
	topic := []byte("CrossFrom(address indexed validator, (uint256,uint256,address,address,uint256,uint256,address,address,address) order, uint8 srcTokenDecimals, uint256 crossAmount, uint256 paidAmount)")
	topicHash := crypto.Keccak256Hash(topic)
	return topicHash.Hex()
}

// event CommitReceipt 0x581db44feed8ab7f2b0e591fd633c1326a4ba3ea20a5c346ab38fd1f42208e81
func GetCommitReceiptTopic() string {
	topic := []byte("CommitReceipt(address indexed validator, bytes32 indexed orderHash, (bytes32,bytes32) receipt)")
	topicHash := crypto.Keccak256Hash(topic)
	return topicHash.Hex()
}
