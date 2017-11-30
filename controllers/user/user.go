package user

import (
	"github.com/astaxie/beego"
	"oms/models"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) Login(){

	//如果没有admin，就创建一个  初始 admin  admin123
	mdlUser := new(models.User)
	if !mdlUser.AdminInit() {
		models.CreateAdmin()
	}

	this.Layout = "layout/onlyBody.html"
	this.TplName = "user/login.html"
}

func (this *UserController) DoLogin(){
	username := this.GetString("username")
	password := this.GetString("password")
	if len(username)<4 || len(password)<6{
		this.Data["json"] = map[string]interface{}{"result":0, "msg":"请输入正确的用户名和密码"}
	}else{
		objUser := new(models.User)
		if objUser.Check(username, password){
			this.Data["json"] = map[string]interface{}{"result":1,"msg":"登陆成功"}
		}else{
			this.Data["json"] = map[string]interface{}{"result":0,"msg":"用户名或密码错误"}
		}
	}
	this.ServeJSON()
}