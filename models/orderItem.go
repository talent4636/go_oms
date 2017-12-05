package models

type OrderItem struct {
	Id int
	Orders *Orders `orm:"rel(fk)"`
	Goods *Goods `orm:"rel(one)"`
	SalePrice float32 `orm:"digits(10);decimals(2);default(0.00);column(sale_price)"`
	Price float32 `orm:"digits(10);decimals(2);default(0.00)"`
	PmtPrice float32 `orm:"digits(10);decimals(2);default(0.00);column(pmt_price)"`
	Quantity int
	SendQuantity int `orm:"column(send_quantity);default(0)"`
}
