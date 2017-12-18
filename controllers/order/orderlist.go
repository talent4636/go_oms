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

/**
{
  "Id": 1,
  "OrderBn": "TBN20171207152512",
  "OrderItem": [
    {
      "Id": 1,
      "Order": {
        "Id": 1,
        "OrderBn": "",
        "OrderItem": null,
        "Status": 0,
        "ShipStatus": 0,
        "Payed": 0,
        "CustomRemark": "",
        "ServiceRemark": "",
        "PayStatus": 0,
        "TotalPrice": 0,
        "GoodsPrice": 0,
        "ShipPrice": 0,
        "PmtGoodsPrice": 0,
        "PmtOrderPrice": 0,
        "PmtGoods": null,
        "PmtOrder": null,
        "Created": "0001-01-01T00:00:00Z",
        "Modified": "0001-01-01T00:00:00Z",
        "ReceiverName": "",
        "ReceiverZip": "",
        "ReceiverProvince": "",
        "ReceiverCity": "",
        "ReceiverDistinct": "",
        "ReceiverAddress": "",
        "ReceiverMobile": "",
        "Shop": null,
        "PayMethod": ""
      },
      "Goods": {
        "Id": 1,
        "Name": "test01",
        "Desc": "test01",
        "Bn": "test01",
        "Cost": 11,
        "Price": 11,
        "PicUrl": "/static/img/goods/7988df89ecc57fa34ca47f402ef98e8c.JPG",
        "TypeId": 11,
        "CatId": 111,
        "Created": "2017-12-07T07:25:00+08:00",
        "Modified": "2017-12-07T07:25:00+08:00",
        "OrderItem": null
      },
      "SalePrice": 10,
      "Price": 15,
      "PmtPrice": 5,
      "Quantity": 2,
      "SendQuantity": 0
    }
  ],
  "Status": 0,
  "ShipStatus": 0,
  "Payed": 30,
  "CustomRemark": "测试订单",
  "ServiceRemark": "客服备注111",
  "PayStatus": 1,
  "TotalPrice": 30,
  "GoodsPrice": 30,
  "ShipPrice": 15,
  "PmtGoodsPrice": 10,
  "PmtOrderPrice": 5,
  "PmtGoods": null,
  "PmtOrder": null,
  "Created": "2017-12-07T07:25:12+08:00",
  "Modified": "2017-12-07T07:25:12+08:00",
  "ReceiverName": "谢俊",
  "ReceiverZip": "123456",
  "ReceiverProvince": "上海",
  "ReceiverCity": "上海市",
  "ReceiverDistinct": "徐汇区",
  "ReceiverAddress": "越虹广场H88B9",
  "ReceiverMobile": "18712344321",
  "Shop": null,
  "PayMethod": "支付宝支付"
}
 */
func (this *OrderListController) Detail(){
	if id, err := strconv.Atoi(this.Ctx.Input.Param(":id")); err != nil {
		//
	}else{
		mdlOrder := new(models.Order)
		order := mdlOrder.GetDetail(id)
		//this.Data["json"] = order
		//this.ServeJSON()
		this.Data["order"] = order
		this.Data["cssPath"] = "./../../"
		this.Data["title"] = "订单详情"
		this.Data["pageTitle"] = "订单详情"
	}
	this.Layout = "layout/singlePage.html"
	this.TplName = "order/detail.html"
}