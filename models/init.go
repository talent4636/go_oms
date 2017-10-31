package models

import (
	_ "time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	//	orm.DefaultTimeLoc = time.UTC //这里有问题，时区不对
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:passw0rd@/beego_oms?charset=utf8")
	//	orm.RegisterModel(new(User))
	orm.SetMaxIdleConns("default", 30)
	orm.SetMaxOpenConns("default", 30)
	//需要注册的表都写到这里
	orm.RegisterModelWithPrefix("oms_", new(Goods))
	//END
	orm.RunSyncdb("default", false, true) //第二个参数，true就强制更新表了
}
