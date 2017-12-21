package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Goods struct {
	Id       	int 	 `orm:"column(goods_id)";pk:"auto"`
	Name     	string
	Desc     	string  `orm:"size(255);null"`
	Bn       	string  `orm:"unique"`
	Cost     	float32 `orm:"digits(10);decimals(2);default(0.00)"`
	Price    	float32 `orm:"digits(10);decimals(2);default(0.00)"`
	PicUrl   	string  `orm:"null"`
	TypeId   	int     //像这样的字段定义，会保存成 type_id
	CatId    	int
	Created  	time.Time `orm:"auto_now_add;type(datetime)"`
	Modified 	time.Time `orm:"auto_now;type(datetime)"`
	OrderItem 	[]*OrderItem `orm:"reverse(many)"`
	Store  		[]*Store `orm:"reverse(many)"`
	IO	     	[]*IO `orm:"reverse(many)"`
}

func (this *Goods) GetAll() ([]Goods, error) {
	o := orm.NewOrm()
	var data []Goods
	_, err := o.QueryTable("oms_goods").All(&data)
	return data, err
}

func (this *Goods) GetOneById(goods_id int) ( Goods, error) {
	o := orm.NewOrm()
	var data Goods
	err := o.QueryTable("oms_goods").Filter("id",goods_id).One(&data)
	return data, err
}

func (this *Goods) GetOne(filters map[string]interface{})(Goods){
	o := orm.NewOrm()
	var data Goods
	tmp := o.QueryTable("oms_goods")
	for key,value := range filters{
		tmp = tmp.Filter(key,value)
	}
	_,err := tmp.All(&data)
	if err!=nil{
		return Goods{}
	}else {
		return data
	}
}

func (this *Goods) Add(goods *Goods) (int, error) {
	o := orm.NewOrm()
	id, err := o.Insert(goods)
	return int(id), err
}

func (this *Goods) Delete(goods_id int) (int64, error){
	o := orm.NewOrm()
	return o.QueryTable("oms_goods").Filter("id", goods_id).Delete()
}

func (this *Goods) Update(data map[string]interface{}, goods_id int) (int64, error){
	o := orm.NewOrm()
	return o.QueryTable("oms_goods").Filter("id",goods_id).Update(data);
}

//返回表名
func (goods *Goods) TableName() string {
	return "goods"
}

// 多字段索引
func (goods *Goods) TableIndex() [][]string {
	return [][]string{
		[]string{"Bn", "Name"},
	}
}

// 多字段唯一键 唯一索引
func (goods *Goods) TableUnique() [][]string {
	return [][]string{
		[]string{"Bn", "Name", "CatId"},
	}
}

// 设置引擎为 INNODB
func (goods *Goods) TableEngine() string {
	return "INNODB"
}
