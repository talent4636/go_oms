package routers

import (
	"oms/controllers"
	"oms/controllers/goods"

	"github.com/astaxie/beego"
	"oms/controllers/branch"
	"oms/controllers/user"
	"github.com/astaxie/beego/context"
	"oms/controllers/testcase"
	"oms/controllers/order"
	"oms/controllers/store"
)

func init() {

	//所有内容必须登陆才能显示
	beego.InsertFilter("/*", beego.BeforeRouter, func(context *context.Context) {
		var url string = context.Input.URL()
		if url != "/login" {
			objUser := &user.UserController{}
			objUser.Ctx = context
			if !objUser.CheckLogin() {
				context.Redirect(302, "/login")
			}
		}
	})

	beego.Router("/", &controllers.MainController{})

	//登陆注册模块
	beego.Router("/login", &user.UserController{}, "Get:Login;Post:DoLogin")
	beego.Router("/logout", &user.UserController{}, "Get:Logout")

	//设置模块
	beego.Router("/setting/user", &user.UserSettingController{}, "Get:List")
	beego.Router("/setting/user/edit/?:id", &user.UserSettingController{}, "Get:Edit")
	beego.Router("/setting/user/add", &user.UserSettingController{}, "Get:Edit")
	beego.Router("/setting/user/save", &user.UserSettingController{}, "Post:Save")
	beego.Router("/setting/user/delete/?:id", &user.UserSettingController{}, "Post:Delete")

	//商品模块
	beego.Router("/goods", &goods.GoodsController{})
	beego.Router("/goods/add", &goods.GoodsController{}, "Get:Edit")
	beego.Router("/goods/edit/?:id", &goods.GoodsController{}, "Get:Edit")
	beego.Router("/goods/delete/?:id", &goods.GoodsController{}, "Post:Delete")
	beego.Router("/goods/save", &goods.GoodsController{}, "Post:Save")

	//订单模块
	beego.Router("/order/list", &order.OrderListController{}, "Get:ShowList")
	beego.Router("/order/detail/:id", &order.OrderListController{}, "Get:Detail")
	beego.Router("/order/new", &order.OrderListController{}, "Get:New")

	//仓库模块
	beego.Router("/branch", &branch.BranchController{}, "Get:Get")
	beego.Router("/branch/add", &branch.BranchController{}, "Get:Edit")
	beego.Router("/branch/edit/?:id", &branch.BranchController{}, "Get:Edit")
	beego.Router("/branch/save", &branch.BranchController{}, "Post:Save")
	beego.Router("/branch/delete/?:id", &branch.BranchController{}, "Post:Delete")

	//库存
	beego.Router("/store/all", &store.StoreController{}, "Get:Get")
	beego.Router("/store/branch", &store.StoreController{}, "Get:BranchStore")
	beego.Router("/store/branch/?:id", &store.StoreController{}, "Post:BranchStoreList")
	beego.Router("/store/goods", &store.StoreController{}, "Get:GoodsStore")
	beego.Router("/store/io", &store.IOController{}, "Get:Get")
	beego.Router("/store/io/newin", &store.IOController{}, "Get:AddIn")
	beego.Router("/store/io/newout", &store.IOController{}, "Get:AddOut")
	beego.Router("/store/io/save", &store.IOController{}, "Post:SaveIO")
	beego.Router("/store/inventory", &store.InventoryController{}, "Get:Get")
	beego.Router("/store/io/cancel/?:id", &store.IOController{}, "Post:Cancel")
	beego.Router("/store/io/confirm/?:id", &store.IOController{}, "Post:Confirm")

	//testCase
	beego.Router("/test/order", &testcase.TestController{},"Get:OrderCreate")
}
