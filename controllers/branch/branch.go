package branch

import (
	"github.com/astaxie/beego"
	"oms/controllers/base"
)

type BranchController struct {
	beego.Controller
}

func (this *BranchController) Get(){
	this.Data["_BASE"] = base.NavData("/branch", this.Ctx)
	this.Data["cssPath"] = "./"
	this.Layout = "layout/main.html"
	this.TplName = "branch/list.html"
}