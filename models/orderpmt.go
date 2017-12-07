package models

type PmtGoods struct{
	Id int
	//OrderId int `orm:"column(order_id)"`
	Order *Order `orm:"rel(fk)"`
	Name string
	Price float32 `orm:"digits(10);decimals(2);default(0.00)"` //优惠的价格 单价
}

type PmtOrder struct{
	Id int
	//OrderId int `orm:"column(order_id)"`
	Order *Order `orm:"rel(fk)"`
	Name string
	Price float32 `orm:"digits(10);decimals(2);default(0.00)"`
}