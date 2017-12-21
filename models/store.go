package models

import (
	"github.com/astaxie/beego/orm"
)

type Store struct {
	Id			int
	//BranchId	int		`orm:"column(branch_id)"`
	Branch		*Branch `orm:"column(branch_id);rel(fk)"`
	Goods		*Goods  `orm:"column(goods_id);rel(fk)"`
	//GoodsBn		string  `orm:"size(50);column(goods_bn)"`
	Freeze		int		`orm:"default(0)"`
	Store		int		`orm:"default(0)"`
}

func (mdl *Store) GetList(filters map[string]interface{}) ([]*Store, error){
	o := orm.NewOrm()
	var data []*Store
	tmp := o.QueryTable("oms_store");
	for key,value := range filters{
		tmp = tmp.Filter(key,value)
	}
	if _,err := tmp.All(&data);err!=nil{
		return nil,err
	}else{
		return data, nil
	}
}

func (mdl *Store) Init(branch_id int, goods_bn string, store int, store_feeze int)(int, error){
	o := orm.NewOrm()
	branchInfo := new(Branch).GetOne(map[string]interface{}{"branch_id":branch_id})
	goodsInfo := new(Goods).GetOne(map[string]interface{}{"bn":goods_bn})
	storeData := Store{
		Branch:&branchInfo,
		Goods:&goodsInfo,
		Freeze:store_feeze,
		Store:store,
	}
	if id64, err := o.Insert(storeData); err!=nil{
		return 0,err
	}else{
		return int(id64),nil
	}
}

//入库
func (mdl *Store) In (branch_id int, goods_bn string, store int) (bool,error){
	return mdl.Change(branch_id, goods_bn, store)
}

//出库
func (mdl *Store) Out (branch_id int, goods_bn string, store int) (bool,error){
	return mdl.Change(branch_id, goods_bn, (0-store))
}

func (mdl *Store) Change(branch_id int, goods_bn string, store int) (bool,error){
	o := orm.NewOrm()
	var storeNew *Store
	if _store,_storeFreeze,err := mdl.GetStore(branch_id, goods_bn);err!=nil{
		return false, err
	}else{
		branchInfo := new(Branch).GetOne(map[string]interface{}{"branch_id":branch_id})
		goodsInfo := new(Goods).GetOne(map[string]interface{}{"bn":goods_bn})
		storeNew.Store = _store+store
		storeNew.Branch = &branchInfo
		storeNew.Goods = &goodsInfo
		storeNew.Freeze = _storeFreeze
	}
	if _,err := o.Update(&storeNew);err!=nil{
		return false,err
	}else{
		return true,nil
	}
}

//return 库存 冻结 error
func (mdl *Store) GetStore (branch_id int, goods_bn string) (int, int, error){
	o := orm.NewOrm()
	var store *Store
	if err := o.QueryTable("oms_store").Filter("branch_id", branch_id).
		Filter("goods_bn", goods_bn).One(&store);err!=nil{
			return 0,0,err
	}else{
		return store.Store,store.Freeze,nil
	}
}