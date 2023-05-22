package util

import (
	"fmt"
	"math/big"
	"regexp"

	"github.com/MysGate/demo_backend/core"
	"github.com/shopspring/decimal"
	"github.com/bwmarrin/snowflake"
)

var IsAlphanumeric = regexp.MustCompile(`^[0-9a-zA-Z]+$`).MatchString

func ConvertHexToDecimalInStringFormat(hexString string) string {
	i := new(big.Int)
	// if hexString with '0x' prefix, using fmt.Sscan()
	fmt.Sscan(hexString, i)
	// if hexString without '0x' prefix, using i.SetString()
	//i.SetString(hexString, 16)

	return fmt.Sprintf("%v", i)
}

func PadLeft(str, pad string, length int) string {
	for {
		str = pad + str
		if len(str) >= length {
			return str[0:length]
		}
	}
}

func IsAnAddress(address string) bool {
	return len(address) == core.AddressFixedLength+2 && address[:2] == "0x" && IsAlphanumeric(address)
}

func IsValidTxHash(txHash string) bool {
	return len(txHash) == core.TxHashFixedLength && txHash[:2] == "0x" && IsAlphanumeric(txHash)
}

func ConvertTokenAmountToFloat64(amt string, tokenDecimal int32) float64 {
	amount, _ := decimal.NewFromString(amt)
	amount_converted := amount.Div(decimal.New(1, tokenDecimal))
	amountFloat, _ := amount_converted.Float64()
	return amountFloat
}

func GenerateIncreaseID() (int64, error) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		Logger().Error(fmt.Sprintf("GenerateIncreaseID err:%+v", err))
		return 0, err
	}
	// Generate a snowflake ID.
	id := node.Generate()

	return id.Int64(), nil
}
