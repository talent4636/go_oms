package routers

import (
	"oms/controllers"
	"oms/controllers/goods"

	"github.com/astaxie/beego"
	"oms/controllers/branch"
	"oms/controllers/user"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	//	beego.Router("/", &goods.GoodsController{})

	beego.Router("/login", &user.UserController{}, "Get:Login")
	beego.Router("/login", &user.UserController{}, "Post:DoLogin")

	beego.Router("/goods", &goods.GoodsController{})
	beego.Router("/goods/add", &goods.GoodsController{}, "Get:Edit")
	beego.Router("/goods/edit/?:id", &goods.GoodsController{}, "Get:Edit")
	beego.Router("/goods/delete/?:id", &goods.GoodsController{}, "Post:Delete")
	beego.Router("/goods/save", &goods.GoodsController{}, "Post:Save")

	beego.Router("/branch", &branch.BranchController{}, "Get:Get")
}
