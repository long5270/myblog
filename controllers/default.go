package controllers

import (
	"fmt"
	"myblog/models"
	"strconv"
)

type MainController struct {
	BaseController
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

type LoginController struct {
	BaseController
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
	if err != nil {
		this.Ctx.WriteString("登录失败")
		return
	}
	this.setUser2Session(user.Id)
	this.Redirect("/article/", 302)
}

type RegisterController struct {
	BaseController
}

func (this *RegisterController) Get() {
	this.TplName = "register.html"
}

func (this *RegisterController) Post() {
	username := this.GetString("UserName")
	password := this.GetString("PassWord")
	name := this.GetString("Name")
	user, err := models.CreateUser(username, password, name)
	this.TplName = "register_result.html"
	if err != nil {
		this.Data["Result"] = "注册失败"
		return
	}
	this.Data["Result"] = "注册成功 您的用户名是 " + user.UserName
}

type Article struct {
	BaseController
}

func (this *Article) List() {
	pageSize, err := this.GetInt("pageSize")
	if err != nil {
		pageSize = 5
	}
	pageNumber, err := this.GetInt("pageNumber")
	if err != nil {
		pageNumber = 1
	}
	articles, total := models.ListArticle(pageNumber, pageSize)
	this.TplName = "article_list.html"
	this.Data["Articles"] = articles
	this.Data["Total"] = total
	this.Data["PageNumber"] = pageNumber
	AllPage := total / int64(pageSize)
	if total%int64(pageSize) > 0 {
		AllPage += 1
	}
	this.Data["AllPage"] = AllPage
}

func (this *Article) CreateView() {
	this.checkLogin()
	this.TplName = "article_create.html"
}

func (this *Article) CreateArticle() {
	this.checkLogin()
	title := this.GetString("title")
	content := this.GetString("content")
	tagSting := this.GetString("tagSting")
	models.CreateArticle(&this.curUser, title, content, tagSting)
	this.Redirect("/article/", 302)
}

func (this *Article) Detail() {
	this.checkLogin()
	Id, _ := this.GetInt(":id")
	article := models.ArticleOne(Id)
	this.TplName = "article_detail.html"
	this.Data["Article"] = article
}

func (this *Article) CreateComment() {
	Id, _ := this.GetInt(":id")
	parentId, _ := this.GetInt(":parentId")
	content := this.GetString("content")
	models.CreateComment(&this.curUser, content, Id, parentId)
	this.Redirect("/article/"+strconv.Itoa(Id), 302)
}
