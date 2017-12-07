package testcase

import (
	"github.com/astaxie/beego"
	"oms/models"
)

type TestController struct {
	beego.Controller
}

func (this *TestController)OrderCreate(){
	mdlTest := new(models.Testcase)
	order_id := mdlTest.FakeOrder()
	if order_id == 0{
		this.Data["json"] = `{"创建失败"}`
	}else{
		mdlOrder := new(models.Order)
		this.Data["json"] = mdlOrder.GetDetail(order_id)
	}
	this.ServeJSON()
}