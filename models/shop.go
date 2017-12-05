package models

type Shop struct {
	Id int `pk:"auto";orm:"column(shop_id)"`
	Bn string `orm:"column(shop_bn)"`
	Name string
	Status int `orm:"default(0)"` //0未绑定  1已绑定
	Info string `orm:"null"`
	Orders *Orders `orm:"rel(one)"`
}