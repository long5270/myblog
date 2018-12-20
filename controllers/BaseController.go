package controllers

import (
	"github.com/astaxie/beego"
	"myblog/models"
)

type BaseController struct {
	beego.Controller
	controllerName string      //当前控制名称
	actionName     string      //当前action名称
	curUser        models.User //当前用户信息
}

func (this *BaseController) Prepare() {
	//附值
	this.controllerName, this.actionName = this.GetControllerAndAction()

	this.Data["siteApp"] = beego.AppConfig.String("site.app")
	this.Data["siteName"] = beego.AppConfig.String("site.name")
	this.Data["siteVersion"] = beego.AppConfig.String("site.version")

	//从Session里获取数据 设置用户信息
	this.adapterUserInfo()
}

//从session里取用户信息
func (this *BaseController) adapterUserInfo() {
	a := this.GetSession("backenduser")
	if a != nil {
		this.curUser = a.(models.User)
		this.Data["backenduser"] = a
	}
}

// checkLogin判断用户是否登录，未登录则跳转至登录页面
// 一定要在BaseController.Prepare()后执行
func (this *BaseController) checkLogin() {
	if this.curUser.Id == 0 {
		//登录页面地址
		urlstr := this.URLFor("HomeController.Login") + "?url="
		//登录成功后返回的址为当前
		returnURL := this.Ctx.Request.URL.Path
		this.Redirect(urlstr+returnURL, 302)
		this.StopRun()
	}
}

//SetBackendUser2Session 获取用户信息（包括资源UrlFor）保存至Session
//被 HomeController.DoLogin 调用
func (this *BaseController) setUser2Session(userId int) error {
	m, err := models.UserOne(userId)
	if err != nil {
		return err
	}
	this.SetSession("backenduser", *m)
	return nil
}

// 重定向
func (this *BaseController) redirect(url string) {
	this.Redirect(url, 302)
	this.StopRun()
}

// 重定向 去错误页
func (this *BaseController) pageError(msg string) {
	errorurl := this.URLFor("HomeController.Error") + "/" + msg
	this.Redirect(errorurl, 302)
	this.StopRun()
}

// 重定向 去登录页
func (this *BaseController) pageLogin() {
	url := this.URLFor("HomeController.Login")
	this.Redirect(url, 302)
	this.StopRun()
}
