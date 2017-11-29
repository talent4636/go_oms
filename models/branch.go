package models

import (
	_ "github.com/astaxie/beego/orm"
	"time"
)

type Branch struct {
	Id       int
	Name     string
	Desc     string  `orm:"size(255);null"`
	Bn       string  `orm:"unique"`
	Created  time.Time `orm:"auto_now_add;type(datetime)"`
	Modified time.Time `orm:"auto_now;type(datetime)"`
}

//type BranchStore struct{
//	BranchId 	int	`orm:"pk:column(branch_id)"`
//	GoodsId 	int
//	Store    	int
//	StoreFreez 	int
//}