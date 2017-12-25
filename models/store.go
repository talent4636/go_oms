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

func (mdl *Store) GetList(filters map[string]interface{}) ([]Store, error){
	o := orm.NewOrm()
	var data []Store
	tmp := o.QueryTable("oms_store");
	for key,value := range filters{
		tmp = tmp.Filter(key,value)
	}
	if _,err := tmp.All(&data);err!=nil{
		return nil,err
	}else{
		for index,value := range data{
			branch := new(Branch).GetOne(map[string]interface{}{"id":value.Branch.Id})
			goods := new(Goods).GetOne(map[string]interface{}{"id":value.Goods.Id})
			data[index].Branch = &branch
			data[index].Goods = &goods
		}
		return data, nil
	}
}

func (mdl *Store) Init(branch_id int, goods_id int, store int, store_feeze int)(int, error){
	o := orm.NewOrm()
	branchInfo := new(Branch).GetOne(map[string]interface{}{"id":branch_id})
	goodsInfo := new(Goods).GetOne(map[string]interface{}{"id":goods_id})
	storeData := Store{
		Branch:&branchInfo,
		Goods:&goodsInfo,
		Freeze:store_feeze,
		Store:store,
	}
	if id64, err := o.Insert(&storeData); err!=nil{
		return 0,err
	}else{
		return int(id64),nil
	}
}

//入库
func (mdl *Store) In (branch_id int, goods_id int, store int) (bool,error){
	return mdl.Change(branch_id, goods_id, store)
}

//出库
func (mdl *Store) Out (branch_id int, goods_id int, store int) (bool,error){
	return mdl.Change(branch_id, goods_id, (0-store))
}

func (mdl *Store) Change(branch_id int, goods_id int, store int) (bool,error){
	//如果没有数据，初始化
	var storeNew Store
	o := orm.NewOrm()
	list,err := mdl.GetList(map[string]interface{}{"branch_id":branch_id,"goods_id":goods_id});
	if err!=nil || list==nil || len(list)==0{
		mdl.Init(branch_id, goods_id,0,0)
		list,_ := mdl.GetList(map[string]interface{}{"branch_id":branch_id,"goods_id":goods_id});
		storeNew = list[0]
	}else{
		storeNew = list[0]
	}
	if _store,_storeFreeze,err := mdl.GetStore(branch_id, goods_id);err!=nil{
		return false, err
	}else{
		branchInfo := new(Branch).GetOne(map[string]interface{}{"id":branch_id})
		goodsInfo := new(Goods).GetOne(map[string]interface{}{"id":goods_id})
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
func (mdl *Store) GetStore (branch_id int, goods_id int) (int, int, error){
	o := orm.NewOrm()
	var store Store
	if err := o.QueryTable("oms_store").Filter("branch_id", branch_id).
		Filter("goods_id",goods_id).One(&store);err!=nil{
			return 0,0,err
	}else{
		return store.Store,store.Freeze,nil
	}
}