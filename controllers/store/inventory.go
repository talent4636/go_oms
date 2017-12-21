package store

import "github.com/astaxie/beego"

type InventoryController struct{
	beego.Controller
}

func (this *InventoryController) Get(){
	this.Data["json"] = "inventory get"
	this.ServeJSON()
}
