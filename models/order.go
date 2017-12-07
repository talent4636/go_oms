package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type Order struct {
	Id       int `orm:"column(order_id)";pk:"auto"`
	OrderBn  string  `orm:"column(order_bn)"`
	OrderItem	[]*OrderItem	`orm:"reverse(many)"` //设置一对多的反向关系
	Status int `orm:"default(0)"` //0活动订单   1完成订单finish    2作废订单
	ShipStatus int `orm:"default(0);column(ship_status)"` //0未发货  1已发货   2部分发货   3已退货
	Payed float32 `orm:"digits(10);decimals(2);default(0.00)"`
	CustomRemark string `orm:"null;column(custom_remark)"` //客户备注
	ServiceRemark string `orm:"null;column(service_remark)"` //客服备注
	PayStatus int `orm:"default(0);column(pay_status)"`  //0未支付  1已支付   2部分支付  3已部分退款  4全额退款
	TotalPrice float32 `orm:"digits(10);decimals(2);default(0.00);column(total_price)"`
	GoodsPrice float32 `orm:"digits(10);decimals(2);default(0.00);column(goods_price)"`
	ShipPrice float32 `orm:"digits(10);decimals(2);default(0.00);column(ship_price)"`
	PmtGoodsPrice float32 `orm:"digits(10);decimals(2);default(0.00);column(pmt_goods_price)"`
	PmtOrderPrice float32 `orm:"digits(10);decimals(2);default(0.00);column(pmt_order_price)"`
	PmtGoods []*PmtGoods `orm:"reverse(many);null"`
	PmtOrder []*PmtOrder `orm:"reverse(many);null"`
	Created  time.Time `orm:"auto_now_add;type(datetime)"`
	Modified time.Time `orm:"auto_now;type(datetime)"`
	//收货人信息
	ReceiverName string `orm:"size(255)"`
	ReceiverZip string `orm:"size(10);null"`
	ReceiverProvince string `orm:"size(32);null;column(receiver_province)"`
	ReceiverCity string `orm:"size(32);null;column(receiver_city)"`
	ReceiverDistinct string `orm:"size(32);null;column(receiver_distinct)"`
	ReceiverAddress string `orm:"size(255);null;column(receiver_address)"`
	ReceiverMobile string `orm:"size(20);column(receiver_mobile);null"`
	//店铺信息
	Shop *Shop `orm:"reverse(one)"`
	//支付信息
	PayMethod string `orm:"column(pay_method);null;size(100)"`
}

func (this *Order) CreateOrder(order *Order) (int, error){
	o := orm.NewOrm()
	//创建订单加入事物
	o.Begin()
	if orderid64, err := o.Insert(order);err!=nil{
		o.Rollback()
		return 0,err
	}else{
		order_id := int(orderid64)
		mdlItem := new(OrderItem)
		for _,item := range order.OrderItem{
			item.Order = order
			//item.Goods = item.Goods
			if _,err := mdlItem.Save(item); err!=nil{
				o.Rollback()
				return 0,err
			}
		}
		o.Commit()
		return order_id,nil
	}
}

func (this *Order) GetDetail(order_id int) Order{
	o := orm.NewOrm()
	var order Order
	if err:=o.QueryTable("oms_order").
		Filter("order_id", order_id).One(&order); err!=nil{
		return Order{}
	}else{
		mdlItem := new(OrderItem)
		order.OrderItem = mdlItem.GetItemByOrderId(order_id)
		return order
	}
}

func (this *Order) GetList(filters map[string]string, offset int, limit int) []Order{
	o := orm.NewOrm()
	var orders []Order
	tmp := o.QueryTable("oms_order")
	for key,value := range filters{
		tmp = tmp.Filter(key,value)
	}
	_,err := tmp.Offset(offset).Limit(limit).RelatedSel().All(&orders)
	if err!=nil{
		return nil
	}else{
		return orders
	}
}
