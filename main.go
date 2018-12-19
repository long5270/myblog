package main

import (
	_ "myblog/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run("localhost:8080")
}

