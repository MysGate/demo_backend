package model

import (
	"time"

	"github.com/go-xorm/xorm"
)

type Block struct {
	ID          int       `xorm:"'id' pk autoincr"`
	ChainId     int       `xorm:"chain_id" json:"chain_id"`
	Contract    string    `xorm:"contract" json:"contract"`
	BlockNumber int64     `xorm:"block_number" json:"block_number"`
	Created     time.Time `xorm:"created" json:"created"`
	Updated     time.Time `xorm:"updated" json:"updated"`
}

func (o *Block) TableName() string {
	return GetBlockTableName()
}

func InsertBlock(block *Block, db *xorm.Engine) error {
	_, err := db.Table(GetBlockTableName()).Insert(block)
	return err
}

func GetBlock(chainId int64, contract string, db *xorm.Engine) (bool, *Block) {
	var block Block
	has, _ := db.Table(GetBlockTableName()).Where(" chain_id= ? and contract= ?", chainId, contract).Get(&block)
	return has, &block
}
func UpdateBlock(id int, blockNumber int64, db *xorm.Engine) error {
	_, err := db.Table(GetBlockTableName()).ID(id).Update(&Block{BlockNumber: blockNumber})
	return err
}
