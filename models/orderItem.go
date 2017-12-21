package models

import (
	"github.com/astaxie/beego/orm"
)

type OrderItem struct {
	Id int
	Order *Order `orm:"rel(fk);"`
	Goods *Goods `orm:"rel(fk)"`
	SalePrice float32 `orm:"digits(10);decimals(2);default(0.00);column(sale_price)"`
	Price float32 `orm:"digits(10);decimals(2);default(0.00)"`
	PmtPrice float32 `orm:"digits(10);decimals(2);default(0.00);column(pmt_price)"`
	Quantity int
	SendQuantity int `orm:"column(send_quantity);default(0)"`
}

func (item *OrderItem) Save(oItem *OrderItem) (int, error){
	o := orm.NewOrm()
	if id64,err := o.Insert(oItem);err!=nil{
		return 0,err
	}else{
		return int(id64),nil
	}
}

func (item *OrderItem) GetItemByOrderId(order_id int) []*OrderItem{
	o := orm.NewOrm()
	var items []*OrderItem
	if _,err := o.QueryTable("oms_order_item").Filter("order_id",order_id).All(&items);err!=nil{
		return nil
	}else{
		//把goods信息也放进去
		for key,itm := range items{
			goods,_ := new(Goods).GetOneById(itm.Goods.Id)
			items[key].Goods = &goods
		}
		return items
	}
}