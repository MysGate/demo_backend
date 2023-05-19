package module

import (
	"time"

	"github.com/go-xorm/xorm"
)

type Order struct {
	ID         int    `xorm:"'id' pk autoincr"`
	PoterId    string `xorm:"poter_id" json:"poter_id"`
	SrcChainId int    `xorm:"src_chain_id" json:"src_chain_id"`
	SrcAddress string `xorm:"src_address" json:"src_address"`
	SrcToken   string `xorm:"src_token" json:"src_token"`
	SrcAmount  string `xorm:"src_amount" json:"src_amount"`
	SrcTxHash  string `xorm:"src_tx_hash" json:"src_tx_hash"`

	FixedFee    string `xorm:"fixed_fee" json:"fixed_fee"`
	FeeRate     string `xorm:"fee_rate" json:"fee_rate"`
	TransferFee string `xorm:"transfer_fee" json:"transfer_fee"`
	TotalFee    string `xorm:"total_fee" json:"total_fee"`

	DestChainId int    `xorm:"dest_chain_id" json:"dest_chain_id"`
	DestAddress string `xorm:"dest_address" json:"dest_address"`
	DestAmount  string `xorm:"dest_amount" json:"dest_amount"`
	DestTxHash  string `xorm:"dest_tx_hash" json:"dest_tx_hash"`

	Proof string `xorm:"proof" json:"proof"`

	Created      time.Time `xorm:"created" json:"created"`
	FinishedTime time.Time `xorm:"finished_time" json:"finish_time"`
	Updated      time.Time `xorm:"updated" json:"updated"`
}

func (o *Order) TableName() string {
	return GetOrderTableName()
}

func (o *Order) GetOrder(orderId int, db *xorm.Engine) (bool, Order) {
	order := Order{}
	has, _ := db.Table(GetOrderTableName()).ID(orderId).Get(&order)
	return has, order
}

func (o *Order) GetOrderList(src_chain_id int, dest_chain_id int, db *xorm.Engine) (error, []Order) {
	orders := make([]Order, 0)
	err := db.Table(GetOrderTableName()).Where("src_chain_id = ? and dest_chain_id = ?", src_chain_id, dest_chain_id).Find(&orders)
	return err, orders
}
