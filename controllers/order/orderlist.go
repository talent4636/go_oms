package order

import (
	"github.com/astaxie/beego"
	"oms/controllers/base"
	"strconv"
	"oms/models"
)

type OrderListController struct {
	beego.Controller
}

func (this *OrderListController) ShowList(){
	this.Data["_BASE"] = base.NavData("/order/list", this.Ctx)
	this.Data["cssPath"] = "../"
	mdlOrder := new(models.Order)
	if orders := mdlOrder.GetList(nil,0,0); len(orders) > 0 {
		//this.Data["json"] = orders
		//this.ServeJSON()
		//logs.Warn(goods)
		this.Data["orders"] = orders
		this.Data["_COUNT"] = strconv.Itoa(len(orders))
	}
	this.Layout = "layout/main.html"
	this.TplName = "order/list.html"
}