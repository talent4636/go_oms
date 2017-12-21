package branch

import (
	"github.com/astaxie/beego"
	"oms/controllers/base"
	"oms/models"
	"strconv"
)

type BranchController struct {
	beego.Controller
}

func (this *BranchController) Get(){
	this.Data["_BASE"] = base.NavData("/branch", this.Ctx)
	this.Data["cssPath"] = "./"
	mdlBranch := new(models.Branch)
	branchs := mdlBranch.GetList(nil,0,10)
	this.Data["branchs"] = branchs
	this.Data["_COUNT"] = len(branchs)
	//this.Data["json"] = branchs
	//this.ServeJSON()
	this.Layout = "layout/main.html"
	this.TplName = "branch/list.html"
}

func (this *BranchController) Edit(){
	if id, err := strconv.Atoi(this.Ctx.Input.Param(":id")); err != nil {
		this.Data["cssPath"] = "./../../"
		this.Data["title"] = "新增仓库"
		this.Data["pageTitle"] = "新增仓库"
	}else{
		mdlBranch := new(models.Branch)
		filter := map[string]interface{}{
			"id":id,
		}
		branchInfo := mdlBranch.GetOne(filter)
		this.Data["branch"] = branchInfo
		this.Data["cssPath"] = "./../../../"
		this.Data["title"] = "编辑仓库"
		this.Data["pageTitle"] = "编辑仓库"
	}
	this.Layout = "layout/singlePage.html"
	this.TplName = "branch/edit.html"
}

func (this *BranchController) Save(){
	post := this.Input()
	mdlBranch := new(models.Branch)
	if _, ok := post["id"]; ok {
		id,_ := strconv.Atoi(post["id"][0])
		Data := map[string]interface{}{}
		for key,value := range post{
			Data[key] = value[0]
		}
		//mdlBranch.Update();
		this.Data["json"] = id;
	}else{
		var newBranch models.Branch = models.Branch{
			Name:post["name"][0],
			Desc:post["desc"][0],
			Bn:post["bn"][0],
			IsSelf:post["is_self"][0]=="true",
		}
		if _,err := mdlBranch.New(&newBranch); err!=nil{
			this.Data["json"] = map[string]interface{}{"result":0,"msg":"创建仓库失败"}
		}else{
			this.Data["json"] = map[string]interface{}{"result":1,"msg":"创建仓库成功"}
		}
	}
	this.ServeJSON()
}

func (this *BranchController) Delete(){
	//
}