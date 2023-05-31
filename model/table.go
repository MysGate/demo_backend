package model

const (
	DB_TABLE_ORDER    = "orders"
	DB_TABLE_CONTRACT = "contract_config"
	DB_TABLE_BLOCK    = "block"
)

func GetOrderTableName() string {
	return DB_TABLE_ORDER
}

func GetContractTableName() string {
	return DB_TABLE_CONTRACT
}
func GetBlockTableName() string {
	return DB_TABLE_BLOCK
}
