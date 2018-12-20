package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	fmt.Println("hhhhhhhh")
	orm.RegisterDataBase("default", "mysql", "root:@tcp(127.0.0.1:3306)/shenhua?charset=utf8mb4&loc=Local", 30)
	orm.Debug = true
}
