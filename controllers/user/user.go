package user

import (
	"github.com/astaxie/beego"
	"oms/models"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) Login(){

	if this.CheckLogin(){
		this.Redirect("/", 302)
	}

	//如果没有admin，就创建一个  初始 admin  admin123
	mdlUser := new(models.User)
	if !mdlUser.AdminInit() {
		models.CreateAdmin()
	}

	//x := this.GetSession("aa")
	//logs.Warn(x)
	//sess := this.StartSession()
	//logs.Warn(sess.Get("uname"))
	//sess.Set("uname", "123123")
	//this.SetSession("aa","aaaaa1")

	this.Layout = "layout/onlyBody.html"
	this.TplName = "user/login.html"
}

func (this *UserController) DoLogin(){
	sess := this.StartSession()
	username := this.GetString("username")
	password := this.GetString("password")
	if len(username)<4 || len(password)<6{
		this.Data["json"] = map[string]interface{}{"result":0, "msg":"请输入正确的用户名和密码"}
	}else{
		objUser := new(models.User)
		if objUser.Check(username, password){
			if(username == "admin"){
				sess.Set("neckname", "超级管理员")
			}else{
				userInfo := objUser.GetOne(map[string]interface{}{"username":username})
				sess.Set("neckname", userInfo.Neckname)
			}
			sess.Set("uname", username)
			this.Data["json"] = map[string]interface{}{"result":1,"msg":"登陆成功"}
		}else{
			this.Data["json"] = map[string]interface{}{"result":0,"msg":"用户名或密码错误"}
		}
	}
	this.ServeJSON()
}

func (this *UserController) Logout (){
	sess := this.StartSession()
	sess.Delete("uname")
	sess.Delete("neckname")
	this.Redirect("/login", 302)
}

func (this *UserController) CheckLogin() bool {
	sess := this.StartSession()
	uname := sess.Get("uname")
	if nil==uname {
		return false
	}else{
		return true
	}
}

func (this *UserController) GetNeckName() string {
	if this.CheckLogin(){
		sess := this.StartSession()
		neckname := sess.Get("neckname")
		return neckname.(string)
	}else{
		return ""
	}
}