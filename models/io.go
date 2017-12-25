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

func (io *IO) GetList(filters map[string]interface{}, offset int, limit int) []IO{
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
		mdlGoods := new(Goods)
		mdlBranch := new(Branch)
		for k,v := range data{
			goods,_ := mdlGoods.GetOneById(v.Goods.Id)
			branch := mdlBranch.GetOne(map[string]interface{}{"id":v.Branch.Id})
			data[k].Goods = &goods
			data[k].Branch = &branch
		}
		return data
	}
}

func (io *IO) GetOne(io_id int) *IO {
	ioInfo := io.GetList(map[string]interface{}{"id":io_id},0,1)[0]
	return &ioInfo
}

func (io *IO) Update(io_id int, modifyData map[string]interface{}) (bool,error){
	o := orm.NewOrm()
	if _,err := o.QueryTable("oms_io").Filter("id",io_id).Update(modifyData);err!=nil{
		return false,err
	}else{
		return true,nil
	}
}

//返回表名
func (io *IO) TableName() string {
	return "io"
}

func (io *IO) New(ioPnt *IO) bool{
	o:=orm.NewOrm()
	if _,err := o.Insert(ioPnt);err!=nil{
		return false;
	}else{
		return true;
	}
}

func (io *IO) Cancel(io_id int) bool{
	if _,err := io.Update(io_id,map[string]interface{}{"status":2,"confirm_at":time.Now(),});err!=nil{
		return false;
	}else{
		return true;
	}
}

func (io *IO) Confirm(io_id int) bool {
	o := orm.NewOrm()
	o.Begin()
	if _,err := io.Update(io_id,map[string]interface{}{"status":1,"confirm_at":time.Now(),});err!=nil{
		o.Rollback()
		return false;
	}else{
		//审核通过了，就开始改库存
		ioInfo := io.GetOne(io_id)
		mdlStore := new(Store)
		var res bool
		var err error
		if ioInfo.Type == 1{
			res,err = mdlStore.In(ioInfo.Branch.Id, ioInfo.Goods.Id, ioInfo.Number)
		}else{
			res,err = mdlStore.Out(ioInfo.Branch.Id, ioInfo.Goods.Id, ioInfo.Number)
		}

		if err!=nil || res==false{
			o.Rollback()
			return false;
		}else{
			o.Commit()
			return true;
		}
	}
}