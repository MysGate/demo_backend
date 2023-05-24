package model

import (
	"time"

	"github.com/MysGate/demo_backend/contracts"
	"github.com/MysGate/demo_backend/core"
	"github.com/MysGate/demo_backend/util"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/go-xorm/xorm"
)

func Keccak256EncodePackedContractOrder(co *contracts.CrossControllerOrder) (orderHash [32]byte) {
	result := util.EncodePacked(
		util.EncodeBigInt(co.OrderId),
		util.EncodeBigInt(co.SrcChainId),
		util.EncodeAddress(co.SrcAddress),
		util.EncodeAddress(co.SrcToken),
		util.EncodeBigInt(co.SrcAmount),
		util.EncodeBigInt(co.DestChainId),
		util.EncodeAddress(co.DestAddress),
		util.EncodeAddress(co.DestToken),
		util.EncodeAddress(co.PorterPool),
	)
	hash := crypto.Keccak256Hash(result)
	copy(orderHash[:], hash.Bytes())
	return
}

type Order struct {
	ID         int64   `xorm:"'id' pk autoincr" json:"order_id"`
	PoterId    string  `xorm:"poter_id" json:"poter_id"`
	SrcChainId uint64  `xorm:"src_chain_id" json:"src_chain_id"`
	SrcAddress string  `xorm:"src_address" json:"src_address"`
	SrcToken   string  `xorm:"src_token" json:"src_token"`
	SrcAmount  float64 `xorm:"src_amount" json:"src_amount"`
	SrcTxHash  string  `xorm:"src_tx_hash" json:"src_tx_hash"`

	FixedFee float64 `xorm:"fixed_fee" json:"fixed_fee"`
	FeeRate  float64 `xorm:"fee_rate" json:"fee_rate"`
	TotalFee float64 `xorm:"total_fee" json:"total_fee"`

	DestChainId uint64  `xorm:"dest_chain_id" json:"dest_chain_id"`
	DestAddress string  `xorm:"dest_address" json:"dest_address"`
	DestToken   string  `xorm:"dest_token" json:"dest_token"`
	DestAmount  float64 `xorm:"dest_amount" json:"dest_amount"`
	DestTxHash  string  `xorm:"dest_tx_hash" json:"dest_tx_hash"`

	Proof               string `xorm:"proof" json:"proof"`
	CommitReceiptTxHash string `xorm:"commit_receipt_tx_hash" json:"commit_receipt_tx_hash"`

	Created      time.Time `xorm:"created" json:"created"`
	FinishedTime time.Time `xorm:"finished_time" json:"finish_time"`
	Updated      time.Time `xorm:"updated" json:"updated"`

	Status int `xorm:"status" json:"status"`
}

func (o *Order) TableName() string {
	return GetOrderTableName()
}

func GetOrder(orderId int64, db *xorm.Engine) (bool, *Order) {
	order := &Order{}
	has, _ := db.Table(GetOrderTableName()).ID(orderId).Get(&order)
	return has, order
}

func GetOrderList(src_chain_id uint64, dest_chain_id uint64, db *xorm.Engine) ([]Order, error) {
	orders := make([]Order, 0)
	err := db.Table(GetOrderTableName()).
		Where("src_chain_id = ?", src_chain_id).
		And("dest_chain_id = ?", dest_chain_id).Find(&orders)
	return orders, err
}

func UpdateOrderStatus(id int64, status int, db *xorm.Engine) error {
	order := &Order{
		Status:  status,
		Updated: time.Now(),
	}

	if status == core.Success {
		order.FinishedTime = time.Now()
	}

	_, err := db.Table(GetOrderTableName()).ID(id).Update(order)
	return err
}

func UpdateOrderDestTxHash(order *Order, db *xorm.Engine) error {
	o := &Order{
		DestTxHash: order.DestTxHash,
		Updated:    time.Now(),
	}

	_, err := db.Table(GetOrderTableName()).ID(order.ID).Update(o)
	return err
}

func UpdateOrderProof(id int64, proof string, db *xorm.Engine) error {
	order := &Order{
		Proof:   proof,
		Updated: time.Now(),
	}
	_, err := db.Table(GetOrderTableName()).ID(id).Update(order)
	return err
}

func InsertOrder(order *Order, db *xorm.Engine) error {
	_, err := db.Table(GetOrderTableName()).Insert(order)
	return err
}
