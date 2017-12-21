package store

import (
	"github.com/astaxie/beego"
	"oms/controllers/base"
	"oms/models"
)

type IOController struct {
	beego.Controller
}

func (this *IOController) Get(){
	this.Data["_BASE"] = base.NavData("/store/io", this.Ctx)
	this.Data["cssPath"] = "../"
	mdlIO := new(models.IO)
	ios := mdlIO.GetList(nil,0,20)
	this.Data["_COUNT"] = len(ios)
	this.Data["ios"] = ios
	this.Layout = "layout/main.html"
	this.TplName = "store/io/list.html"
}

func (this *IOController) AddIn(){
	this.GetInOut("in")
}

func (this *IOController) AddOut(){
	this.GetInOut("out")
}

func (this *IOController) SaveIO(){
	this.Data["json"] = map[string]interface{}{"result":0,"msg":"功能还没有 TODO"}
	this.ServeJSON()
}

func (this *IOController) GetInOut(ioType string){
	this.Data["cssPath"] = "./../../"
	if ioType == "in"{
		this.Data["title"] = "新增入库单"
		this.Data["pageTitle"] = "新增入库单"
	}else{
		this.Data["title"] = "新增出库单"
		this.Data["pageTitle"] = "新增出库单"
	}
	this.Data["type"] = ioType
	this.Data["branchs"],this.Data["goods"] = this.GetDefaultInfo()
	this.Layout = "layout/singlePage.html"
	this.TplName = "store/io/edit.html"
}

func (this *IOController) GetDefaultInfo() ([]models.Branch, []models.Goods){
	mdlBranch := new(models.Branch)
	mdlGoods := new(models.Goods)
	branchs := mdlBranch.GetList(nil,0,100)
	goods,_ := mdlGoods.GetAll()
	return branchs,goods
}