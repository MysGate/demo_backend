package module

import (
	"fmt"

	"github.com/MysGate/demo_backend/util"
	_ "github.com/go-sql-driver/mysql" // mysql
	"github.com/go-xorm/xorm"
)

func InitMySQLXorm(addr string, showSQL bool) *xorm.Engine {
	e, err := xorm.NewEngine("mysql", addr)
	if err != nil {
		util.Logger().Fatal(fmt.Sprintf("create mysql connection failed. err = %v", err))
	}
	e.ShowSQL(showSQL)
	return e
}
