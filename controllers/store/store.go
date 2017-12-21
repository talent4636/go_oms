package store

import (
	"github.com/astaxie/beego"
	"oms/models"
	"oms/controllers/base"
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
	this.Data["json"] = "store branch store"
	this.ServeJSON()
}

func (this *StoreController) GoodsStore(){
	this.Data["json"] = "store goods store"
	this.ServeJSON()
}