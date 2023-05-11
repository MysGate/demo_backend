package chain

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type DestChainHandler struct {
	DestClient      *ethclient.Client
	ContractAddress common.Address
}
