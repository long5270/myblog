package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

func init(){
	fmt.Println("hhhhhhhh")
	orm.RegisterDataBase("default", "mysql", "root:@tcp(127.0.0.1:3306)/shenhua?charset=utf8mb4", 30)
	orm.Debug = true
}