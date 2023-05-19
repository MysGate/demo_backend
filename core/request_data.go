package core

type TradeRequest struct {
	SrcChainId uint64 `form:"src_chain_id" json:"src_chain_id" binding:"required"`
	SrcAddress string `form:"src_address" json:"src_address" binding:"required"`
	SrcToken   string `form:"src_token" json:"src_token" binding:"required"`
	SrcAmount  uint64 `form:"src_amount" json:"src_amount" binding:"required"`

	DestChainId uint64 `form:"dest_chain_id" json:"dest_chain_id" binding:"required"`
	DestAddress string `form:"dest_address" json:"dest_address" binding:"required"`
	DestToken   string `form:"dest_token" json:"dest_token" binding:"required"`
	DestAmount  uint64 `form:"dest_amount" json:"dest_amount" binding:"required"`
}

type OrderSearchRequest struct {
	OrderId int `form:"order_id" json:"order_id" binding:"required"`
}

type OrderListRequest struct {
	SrcChainId  int `form:"src_chain_id" json:"src_chain_id" binding:"required"`
	DestChainId int `form:"dest_chain_id" json:"dest_chain_id" binding:"required"`
}
