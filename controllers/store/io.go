package store

import (
	"github.com/astaxie/beego"
	"oms/controllers/base"
	"oms/models"
	"strconv"
	"time"
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

func (this *IOController) Confirm(){
	if id, err := strconv.Atoi(this.Ctx.Input.Param(":id")); err != nil {
		this.Data["json"] = map[string]interface{}{"result":0,"msg":"服务器繁忙，请重试"}
	}else{
		mdlIO := new(models.IO)
		if result := mdlIO.Confirm(id);result!=true{
			this.Data["json"] = map[string]interface{}{"result":0,"msg":"审核失败"}
		}else{
			this.Data["json"] = map[string]interface{}{"result":1,"msg":"审核成功"}
		}
	}
	this.ServeJSON()
}

func (this *IOController) Cancel(){
	if id, err := strconv.Atoi(this.Ctx.Input.Param(":id")); err != nil {
		this.Data["json"] = map[string]interface{}{"result":0,"msg":"服务器繁忙，请重试"}
	}else{
		mdlIO := new(models.IO)
		if result := mdlIO.Cancel(id);result!=true{
			this.Data["json"] = map[string]interface{}{"result":0,"msg":"作废失败，请重试"}
		}else{
			this.Data["json"] = map[string]interface{}{"result":1,"msg":"作废成功"}
		}
	}
	this.ServeJSON()
}

func (this *IOController) AddIn(){
	this.GetInOut("in")
}

func (this *IOController) AddOut(){
	this.GetInOut("out")
}

func (this *IOController) SaveIO(){
	this.Data["json"] = this.Input()
	branch_id,_ := strconv.Atoi(this.Input()["branch_id"][0])
	goods_id,_ := strconv.Atoi(this.Input()["goods_id"][0])
	number,_ := strconv.Atoi(this.Input()["number"][0])
	_type := this.Input()["type"][0]
	var ioType int
	if _type=="in"{
		ioType=1
	}else{
		ioType=2
	}
	var ioData *models.IO = &models.IO{
		Type:ioType,
		Branch:&models.Branch{
			Id:branch_id,
		},
		Goods:&models.Goods{
			Id:goods_id,
		},
		Number:number,
		Created:time.Now(),
		CreateBy:base.GetNeckName(this.Ctx),
	}
	mdlIO := new(models.IO)
	if mdlIO.New(ioData){
		this.Data["json"] = map[string]interface{}{"result":1,"msg":"创建成功!"}
	}else{
		this.Data["json"] = map[string]interface{}{"result":0,"msg":"创建失败，请重试!"}
	}

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