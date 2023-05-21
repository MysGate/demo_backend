package model

import "github.com/go-xorm/xorm"

type RouterReq struct {
	SrcChain  string  `json:"src_chain"`
	DestChain string  `json:"dest_chain"`
	From      string  `json:"from"`
	To        string  `json:"to"`
	Amount    float64 `json:"amount"`
	Token     string  `json:"token"`
}

type Porter struct {
	Address    string  `json:"address"`
	Transfered int64   `json:"transfered"`
	Completion float32 `json:"completion"`
}

type Router struct {
	From string `json:"from"`
	To   string `json:"to"`
}

type RouterResponse struct {
	Porters []*Porter `json:"porters"`
	Routers []*Router `json:"routers"`
}

func GetPorterTransferedAndCompletion(addr string, e *xorm.Engine) (transfered int64, completion float32, err error) {
	return
}
