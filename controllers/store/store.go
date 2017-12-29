package store

import (
	"github.com/astaxie/beego"
	"oms/models"
	"oms/controllers/base"
	"strconv"
)

type StoreController struct {
	beego.Controller
}

func (this *StoreController) Get(){
	this.Data["_BASE"] = base.NavData("/store/all", this.Ctx)
	this.Data["cssPath"] = "../"
	mdlStore := new(models.Store)
	stores,_ := mdlStore.GetList(nil)
	//this.Data["json"] = stores
	//this.ServeJSON()
	this.Data["_COUNT"] = len(stores)
	this.Data["stores"] = stores
	this.Layout = "layout/main.html"
	this.TplName = "store/all.html"
}

func (this *StoreController) BranchStore(){
	this.Data["_BASE"] = base.NavData("/store/branch", this.Ctx)
	this.Data["cssPath"] = "../"
	mdlBranch := new(models.Branch)
	branchList := mdlBranch.GetList(nil,0,20)
	//this.Data["json"] = branchList
	//this.ServeJSON()
	this.Data["branchs"] = branchList
	this.Layout = "layout/main.html"
	this.TplName = "store/select.html"
}

func (this *StoreController) BranchStoreList(){
	if id, err := strconv.Atoi(this.Ctx.Input.Param(":id")); err != nil {
		this.Data["json"] = map[string]interface{}{"result":0,"msg":"仓库ID为空，请刷新重试"}
	}else{
		mdlStore := new(models.Store)
		if stores,err := mdlStore.GetList(map[string]interface{}{"branch_id":id});err!=nil{
			this.Data["json"] = map[string]interface{}{"result":0,"msg":"查询失败，请重试"}
		}else{
			this.Data["json"] = map[string]interface{}{"result":1,"msg":"","data":stores}
		}
	}
	this.ServeJSON()
}

func (this *StoreController) GoodsStore(){
	this.Data["json"] = "store goods store"
	this.ServeJSON()
}