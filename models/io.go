package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type IO struct{
	Id		int
	Type    int	`orm:"column(type);default(1)"` //1入库  2出库
	Status  int `orm:"default(0)"` //0未审核  1已审核  2已作废
	Branch	*Branch `orm:"column(branch_id);rel(fk)"`
	Goods	*Goods	`orm:"column(goods_id);rel(fk)"`
	Number 	int	`orm:"column(number)"` //出入库数量
	Created	time.Time `orm:"auto_now_add;type(datetime)"`
	ConfirmAt	time.Time `orm:"auto_now_add;type(datetime);column(confirm_at);null"`
	CreateBy	string	`orm:"create_by;null"` //创建人
	ConfirmBy	string	`orm:"confirm_by;null"` //审核人
}

func (io *IO) GetList(filters map[string]string, offset int, limit int) []IO{
	o := orm.NewOrm()
	var data []IO
	tmp := o.QueryTable("oms_io")
	for key,value := range filters{
		tmp = tmp.Filter(key,value)
	}
	_,err := tmp.Offset(offset).Limit(limit).All(&data)
	if err!=nil{
		return nil
	}else{
		return data
	}
}

//返回表名
func (io *IO) TableName() string {
	return "io"
}