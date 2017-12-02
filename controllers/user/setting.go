package user

import (
	"github.com/astaxie/beego"
	"oms/models"
	"oms/controllers/base"
	"strconv"
	"github.com/astaxie/beego/logs"
)

type UserSettingController struct {
	beego.Controller
}

func (this *UserSettingController) List(){
	this.Data["_BASE"] = base.NavData("/setting/user")
	this.Data["cssPath"] = "../"
	mdlUser := new(models.User)
	users := mdlUser.GetList(nil,0,10)
	this.Data["_COUNT"] = len(users)
	this.Data["users"] = users
	this.Layout = "layout/main.html"
	this.TplName = "user/setting.html"
}

func (this *UserSettingController) Edit(){
	if id, err := strconv.Atoi(this.Ctx.Input.Param(":id")); err != nil {
		this.Data["cssPath"] = "./../../"
		this.Data["title"] = "新增用户"
		this.Data["pageTitle"] = "新增用户"
		//this.Redirect("/goods", 302) //没有ID的话，新增
	}else{
		mdlUser := new(models.User)
		filter := map[string]interface{}{
			"id":id,
		}
		user := mdlUser.GetOne(filter)
		this.Data["user"] = user
		this.Data["cssPath"] = "./../../../"
		this.Data["title"] = "编辑用户"
		this.Data["pageTitle"] = "编辑用户"
	}
	this.Layout = "layout/singlePage.html"
	this.TplName = "user/edit.html"
}

func (this *UserSettingController) Save(){
	post := this.Input()
	logs.Warn(post)
	mdlUser := new(models.User)
	if _, ok := post["id"]; ok {
		//modify
		id,_ := strconv.Atoi(post["id"][0])
		modefyData := map[string]interface{}{}
		for key,value := range post{
			modefyData[key] = value[0]
		}
		if _,err := mdlUser.Modify(id,modefyData);err!=nil{
			this.Data["json"] = map[string]interface{}{"result":0,"msg":"修改失败"}
		}else{
			this.Data["json"] = map[string]interface{}{"result":1,"msg":"修改成功"}
		}
	}else{
		//create
		var newUser models.User = models.User{
			Username:post["username"][0],
			Password:post["password"][0],
			Neckname:post["neckname"][0],
			Mobile:post["mobile"][0],
			Email:post["email"][0],
		}
		if _,err := mdlUser.New(newUser);err!=nil{
			this.Data["json"] = map[string]interface{}{"result":0,"msg":"创建用户失败"}
		}else{
			this.Data["json"] = map[string]interface{}{"result":1,"msg":"创建用户成功"}
		}
	}
	this.ServeJSON()
}