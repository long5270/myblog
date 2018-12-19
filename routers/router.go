package routers

import (
	"myblog/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/login/", &controllers.LoginController{})
    beego.Router("/register/", &controllers.RegisterController{})
}
