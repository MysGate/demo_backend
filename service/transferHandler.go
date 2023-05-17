package service

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/MysGate/demo_backend/conf"
	"github.com/MysGate/demo_backend/core"
	"github.com/MysGate/demo_backend/core/errno"
	"github.com/MysGate/demo_backend/log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/sha3"
)

func (s *Server) trade(c *gin.Context) {
	var requestData core.TradeRequest
	if err := c.BindJSON(&requestData); err != nil {
		errMsg := fmt.Sprintf("Failed to bind request params, reason=[%s]", err)
		log.Logger().Error(errMsg)
		core.SendResponse(c, errno.BindErr, nil)
		return
	}
	src_chain := s.cfg.FindCrossChain(requestData.SrcChainId)
	saveResult, txHash := SaveTradeToSourceChain(src_chain, requestData.SrcAddress, requestData.SrcToken, requestData.SrcAmount)
	log.Logger().Info(fmt.Sprintf("src txhash: %s", txHash))
	if saveResult {
		core.SendResponse(c, errno.OK, nil)
	} else {
		core.SendResponse(c, errno.InternalServerErr, nil)
	}
}

func SaveTradeToSourceChain(src_chain *conf.Chain, src_address string, src_token string, src_amount uint64) (bool, string) {

	log.Logger().Info(src_chain.SrcRpcUrl)

	// //用私钥来对交易事务进行签名
	privateKey, err := crypto.HexToECDSA(src_chain.PrivateKey)
	if err != nil {
		log.Logger().Fatal(err.Error())
		return false, ""
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Logger().Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		return false, ""
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	log.Logger().Info(fmt.Sprintf("fromAddress:%s ", fromAddress))

	//查到nonce和燃气价格。
	client, err := ethclient.Dial(src_chain.SrcRpcUrl)
	if err != nil {
		log.Logger().Fatal(err.Error())
		return false, ""
	}
	defer client.Close()
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Logger().Fatal(err.Error())
		return false, ""
	}
	value := big.NewInt(0) // in wei (0 eth)
	toAddress := common.HexToAddress(src_chain.SrcContractAddress)

	//Alice
	srcAddress := common.HexToAddress(src_chain.SrcContractAddress)

	srcTokenAddress := common.HexToAddress(src_token)

	tradeFnSignature := []byte("trade(address,address,uint256)")
	crypto.NewKeccakState()
	hash := sha3.NewLegacyKeccak256()

	hash.Write(tradeFnSignature)
	methodID := hash.Sum(nil)[:4]

	paddedSrcAddress := common.LeftPadBytes((srcAddress).Bytes(), 32)
	paddedSrcTokenAddress := common.LeftPadBytes((srcTokenAddress).Bytes(), 32)

	amount := big.NewInt(int64(src_amount))
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)

	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedSrcAddress...)
	data = append(data, paddedSrcTokenAddress...)
	data = append(data, paddedAmount...)

	tx := types.NewTransaction(nonce, toAddress, value, src_chain.GasLimit, src_chain.GasSuggest, data)

	chainID := big.NewInt(int64(src_chain.ChainID))

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Logger().Fatal(err.Error())
		return false, ""
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Logger().Fatal(err.Error())
		return false, ""
	}
	return true, signedTx.Hash().Hex()
}
