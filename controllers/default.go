package controllers

import (
	"github.com/astaxie/beego"
	"myblog/models"
	"fmt"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}


type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	this.TplName = "login.html"
}

func (this *LoginController) Post() {
	username := this.GetString("UserName")
	password := this.GetString("PassWord")
	fmt.Println(username, password)
	user, err := models.UserLogin(username, password)
	fmt.Println(user)
	if err != nil{
		this.Ctx.WriteString("登录失败")
		return
	}
	this.Ctx.WriteString("登录成功")
}

type RegisterController struct {
	beego.Controller
}

func (this *RegisterController) Get(){
	this.TplName = "register.html"
}

func (this *RegisterController) Post(){
	username := this.GetString("UserName")
	password := this.GetString("PassWord")
	name := this.GetString("Name")
	user, err := models.CreateUser(username, password, name)
	this.TplName = "register_result.html"
	if err != nil{
		this.Data["Result"] = "注册失败"
		return
	}
	this.Data["Result"] = "注册成功 您的用户名是 " + user.UserName
}