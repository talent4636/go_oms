package goods

import (
	"oms/models"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type GoodsController struct {
	beego.Controller
}

func (this *GoodsController) Get() {
	this.Data["cssPath"] = "./"
	mdlGoods := new(models.Goods)
	if goods, _ := mdlGoods.GetAll(); len(goods) > 0 {
		//logs.Warn(goods)
		this.Data["goods"] = goods
	}
	this.Layout = "layout/main.html"
	this.TplName = "goods/list.html"
}

func (this *GoodsController) Edit() {
	if id, err := strconv.Atoi(this.Ctx.Input.Param(":id")); err != nil {
		logs.Warn("id wrong!")
		this.Data["cssPath"] = "./../"
		this.Data["title"] = "新增商品"
		this.Data["pageTitle"] = "新增商品"
		//this.Redirect("/goods", 302) //没有ID的话，新增
	}else{
		mdlGoods := new(models.Goods)
		data, _ := mdlGoods.GetOne(id)
		this.Data["goods"] = data
		this.Data["cssPath"] = "./../../"
		this.Data["title"] = "编辑商品"
		this.Data["pageTitle"] = "编辑商品"
	}
	this.Layout = "layout/singlePage.html"
	this.TplName = "goods/edit.html"
}

func (this *GoodsController) Delete(){
	if id, err := strconv.Atoi(this.Ctx.Input.Param(":id")); err != nil {
		this.Data["json"] = map[string]interface{}{"result":0,"msg":"服务器繁忙，请重试"}
	}else{
		mdlGoods := new(models.Goods)
		if _, err := mdlGoods.Delete(id);err!=nil{
			this.Data["json"] = map[string]interface{}{"result":0,"msg":"删除失败，请重试"}
		}else{
			this.Data["json"] = map[string]interface{}{"result":1,"msg":"删除成功"}
		}
	}
	this.ServeJSON()
}

func (this *GoodsController) Save() {
	post := this.Input()

	if _, ok := post["id"]; ok {
		//TODO modify goods
		id, _ := strconv.Atoi(post["id"][0])
		this.Data["json"] = map[string]interface{}{"result":0,"msg":"有ID:["+strconv.Itoa(id)+"],comming soon!"}
		mdlGoods := new(models.Goods)
		data := map[string]interface{}{
			"name":post["name"][0],
			"bn":post["bn"][0],
			"price":post["price"][0],
			"desc":post["desc"][0],
			"type_id":post["type_id"][0],
			"cost":post["cost"][0],
			"pic_url":post["pic_url"][0],
			"cat_id":post["cat_id"][0],
		}
		if _,err := mdlGoods.Update(data, id);err!=nil{
			this.Data["json"] = map[string]interface{}{"result":0,"msg":"更新失败，请重试"}
		}else{
			this.Data["json"] = map[string]interface{}{"result":1,"msg":"更新成功"}
		}
	} else {
		//add goods
		mdlGoods := new(models.Goods)
		goods := new(models.Goods)
		//		var goods *models.Goods
		goods.Name = post["name"][0]
		goods.Desc = post["desc"][0]
		goods.Bn = post["bn"][0]
		cost, _ := strconv.ParseFloat(post["cost"][0], 32)
		goods.Cost = float32(cost)
		price, _ := strconv.ParseFloat(post["price"][0], 32)
		goods.Price = float32(price)
		//		goods.Price = post["price"][0]
		goods.PicUrl = post["pic_url"][0]
		type_id, _ := strconv.Atoi(post["type_id"][0])
		goods.TypeId = type_id
		cat_id, _ := strconv.Atoi(post["cat_id"][0])
		goods.CatId = cat_id
		if _, err := mdlGoods.Add(goods); err != nil {
			//save error TODO how to display a message-box
			//can we just do post with ajax
			//if error response , just show the error with js
			//if succ response, we show message and jump to new page ?
			//we'll see
			//logs.Warn("error comes")
			//logs.Warn(err)
			this.Data["json"] = map[string]interface{}{"result":0,"msg":"创建商品失败"}
		} else {
			//save succ
			//logs.Warn(id)
			//this.Redirect("/goods", 302)
			this.Data["json"] = map[string]interface{}{"result":1,"msg":"创建商品成功"}
		}
	}
	this.ServeJSON()

}
