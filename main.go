package main

import (
	"encoding/gob"
	"github.com/astaxie/beego"
	"myblog/models"
	_ "myblog/routers"
)

func main() {
	beego.BConfig.WebConfig.Session.SessionProvider = "file"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "./tmp"
	beego.BConfig.WebConfig.Session.SessionOn = true
	gob.Register(models.User{})
	beego.Run("localhost:8080")
}
