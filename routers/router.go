package routers

import (
	"github.com/astaxie/beego"
	"myblog/controllers"
)

func init() {
	beego.Router("/login/", &controllers.LoginController{})
	beego.Router("/register/", &controllers.RegisterController{})
	beego.Router("/article/", &controllers.Article{}, "get:List")
	beego.Router("/article/:id([0-9]+", &controllers.Article{}, "get:Detail")
	beego.Router("/article/create/", &controllers.Article{}, "post:CreateArticle")
	beego.Router("/article/create/", &controllers.Article{}, "get:CreateView")
	beego.Router("/article/:id([0-9]+/comment/", &controllers.Article{}, "post:CreateComment")
}
