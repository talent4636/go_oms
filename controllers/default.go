package controllers

import (
	"github.com/astaxie/beego"
	"oms/controllers/user"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	//判断登陆，如果没登录就登陆
	u := new(user.UserController)
	u.Ctx = c.Ctx
	if !u.CheckLogin() {
		c.Redirect("/login", 302)
	}
	c.Redirect("/goods", 302)
	//	return
	//	c.Data["Website"] = "beego.me"
	//	c.Data["Email"] = "astaxie@gmail.com"
	//	c.TplName = "index.tpl"
}
