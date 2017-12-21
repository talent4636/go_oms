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
	Store  	 []*Store `orm:"reverse(many)"`
	IO	     []*IO `orm:"reverse(many)"`
}

func (mdlBranch *Branch) GetList(filters map[string]string, offset int, limit int) []Branch{
	o := orm.NewOrm()
	var data []Branch
	tmp := o.QueryTable("oms_branch")
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

func (mdlBranch *Branch) GetOne(filters map[string]interface{}) Branch{
	o := orm.NewOrm()
	var data Branch
	tmp := o.QueryTable("oms_branch")
	for key,value := range filters{
		tmp = tmp.Filter(key,value)
	}
	_,err := tmp.All(&data)
	if err!=nil{
		return Branch{}
	}else {
		return data
	}
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

func (mdlBranch *Branch) Update (id int, modifyData map[string]interface{})(int64, error){
	return int64(1),nil
	//o := orm.NewOrm()
	//branch := mdlBranch.GetOne(map[string]interface{}{"id":id})
	//for name,value := range modifyData{
	//	switch name {
	//	case "id":
	//		break
	//	case "password":
	//		if value==""{
	//			break
	//		}
	//		objMd := md5.New()
	//		objMd.Write([]byte(value.(string)))
	//		pass := hex.EncodeToString(objMd.Sum(nil))
	//		user.Password = pass
	//		break
	//	case "neckname":
	//		user.Neckname = value.(string)
	//		break
	//	case "mobile":
	//		user.Mobile = value.(string)
	//		break
	//	case "email":
	//		user.Email = value.(string)
	//		break
	//	default:
	//		break
	//	}
	//}
	//return o.Update(&user)
}
