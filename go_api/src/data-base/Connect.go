package data_base

import (
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"github.com/go-xorm/xorm"
	"common"
)



func init() {
	fmt.Println("初始化数据库连接")
	common.Engine, _ = xorm.NewEngine("mysql", "root:@tcp(localhost:3306)/todo?charset=utf8")
	common.Engine.ShowSQL(true, true, true, true)
	common.Engine.SetMaxIdleConns(50)
	common.Engine.SetMaxOpenConns(10)
	fmt.Println("数据库连接成功")
}
