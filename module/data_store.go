package module

import (
	"fmt"

	"github.com/MysGate/demo_backend/util"
	_ "github.com/go-sql-driver/mysql" // mysql
	"github.com/go-xorm/xorm"
)

var mysgate_enging *xorm.Engine

func InitMySQLXorm(addr string, showSQL bool) *xorm.Engine {
	mysgate_enging, err := xorm.NewEngine("mysql", addr)
	if err != nil {
		util.Logger().Fatal(fmt.Sprintf("create mysql connection failed. err = %v", err))
	}
	mysgate_enging.ShowSQL(showSQL)
	return mysgate_enging
}

func getMySql() *xorm.Engine {
	return mysgate_enging
}
