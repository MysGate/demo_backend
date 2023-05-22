package model

const (
	PARAM_ERROR    string = "Params error"
	AMOUNT_ERROR   string = "Amount error"
	PONG           string = "Pong"
	INTERNAL_ERROR string = "Internal error"
)

type Message struct {
	Message string `json:"message"`
}

type Chain struct {
	ChainName    string   `json:"chain_name"`
	ChainID      uint64   `json:"chain_id"`
	SupportCoins []string `json:"support_coins"`
}

type CrossChainPair struct {
	SrcChain   *Chain   `json:"src_chain"`
	DestChains []*Chain `json:"dest_chains"`
}

type SupportCrossChains struct {
	Pairs []*CrossChainPair `json:"chain_pairs"`
}

type Fee struct {
	Fixed     float64 `json:"fixed"`
	FloatRate float64 `json:"float_rate"`
}

type CostReq struct {
	Coin   string  `json:"coin"`
	Amount float64 `json:"amount"`
}

type CostResp struct {
	Coin      string  `json:"coin"`
	Amount    float64 `json:"amount"`
	RealSend  float64 `json:"real_send"`
	Received  float64 `json:"received"`
	Fixed     float64 `json:"fixed_fee"`
	FloatRate float64 `json:"float_fee_rate"`
	FloatFee  float64 `json:"float_fee"`
}

type NextIDResp struct {
	NextOrderID int64 `json:"next_order_id"`
}

func GetMessage(m string) *Message {
	msg := &Message{
		Message: m,
	}
	return msg
}
