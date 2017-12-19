package models

import (
	"time"
	"github.com/astaxie/beego/logs"
	"strconv"
)

type Testcase struct {
	//
}

var timeLayoutOrderBn string = "20060102150405"

func (this *Testcase) FakeOrder() int{
	var goods1 *Goods = &Goods{
		Bn:"test01",
		Id:1,
	}
	var goods2 *Goods = &Goods{
		Bn:"YGDSB01",
		Id:2,
	}
	Orderitems := []*OrderItem{
		{
			SalePrice:10.00,
			Price:15.00,
			PmtPrice:5.00,
			Quantity:2,
			Goods:goods1,
		},
		{
			SalePrice:80.00,
			Price:100.00,
			PmtPrice:20.00,
			Quantity:1,
			Goods:goods2,
		},
	}
	pmtGoods := []*PmtGoods{
		{
			Name:"购买任意商品优惠30元",
			Price:30.00,
		},
	}
	pmtOrder := []*PmtOrder{
		{
			Name:"开业大酬宾，运费优惠5元",
			Price:5.00,
		},
	}
	shop := &Shop{
		Bn:"123123",
		Name:"测试旗舰店",
	}
	OrderData := Order{
		OrderBn:"TBN"+time.Now().Format(timeLayoutOrderBn),
		OrderItem:Orderitems,
		Payed:110.00,
		CustomRemark:"测试订单",
		ServiceRemark:"客服备注111",
		PayStatus:1,
		TotalPrice:110.00,
		GoodsPrice:100.00,
		ShipPrice:15.00,
		PmtGoodsPrice:30.00,
		PmtOrderPrice:5.00,
		PmtGoods:pmtGoods,
		PmtOrder:pmtOrder,
		Created:time.Now(),
		Modified:time.Now(),
		ReceiverName:"谢俊",
		ReceiverProvince:"上海",
		ReceiverCity:"上海市",
		ReceiverDistinct:"徐汇区",
		ReceiverAddress:"越虹广场H88B9",
		ReceiverMobile:"18712344321",
		ReceiverZip:"123456",
		Shop:shop,
		PayMethod:"支付宝支付",
	}
	mdlOrder := new(Order)
	if orderId,err := mdlOrder.CreateOrder(&OrderData); err!=nil{
		logs.Error("这里？")
		logs.Error(err)
		logs.Error("这里!")
		return 0
	}else{
		logs.Warn("创建成功"+strconv.Itoa(orderId))
		return orderId
	}
}