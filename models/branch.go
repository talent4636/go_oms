package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type Branch struct {
	Id       int
	Name     string
	Desc     string  `orm:"size(255);null"`
	Bn       string  `orm:"unique"`
	Created  time.Time `orm:"auto_now_add;type(datetime)"`
	Modified time.Time `orm:"auto_now;type(datetime)"`
	IsSelf	 bool `orm:"column(is_self);default(true)"`
}

func (mdlBranch *Branch) New(branch *Branch)(int, error){
	o := orm.NewOrm()
	if id64, err := o.Insert(branch); err!=nil{
		return 0,err
	}else{
		return int(id64),nil
	}
}

func (mdlBranch *Branch) DeleteById(branch_id int)(bool, error){
	o := orm.NewOrm()
	var branch *Branch
	if err := o.QueryTable("oms_branch").Filter("id", branch_id).One(&branch); err!=nil{
		return false, err
	}else{
		if _,err := o.Delete(&branch);err!=nil{
			return false, err
		}else{
			return true, nil
		}
	}
}
