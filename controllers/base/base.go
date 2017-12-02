package base

import (
    "github.com/astaxie/beego"
	_ "flag"
)

type BaseController struct {
	beego.Controller
}

//返回NavData map类型
func NavData(url string) map[string]interface{}{
	data := map[string]interface{}{
		"nav_list":map[int]interface{}{
			1:map[string]interface{}{
				"name":"首页",
				"select":true,
				"url":"/",
				"sub_nav":map[int]interface{}{
					1:map[string]interface{}{
						"name":"概览",
						"select":false,
						"url":"/",
					},
				},
			},
			2:map[string]interface{}{
				"name":"基础资料",
				"select":false,
				"url":"/goods",
				"sub_nav":map[int]interface{}{
					1:map[string]interface{}{
						"name":"商品资料管理",
						"select":false,
						"url":"/goods",
					},
					2:map[string]interface{}{
						"name":"仓库管理",
						"select":false,
						"url":"/branch",
					},
				},
			},
			3:map[string]interface{}{
				"name":"设置",
				"select":false,
				"url":"/setting",
				"sub_nav":map[int]interface{}{
					1:map[string]interface{}{
						"name":"用户管理",
						"select":false,
						"url":"/setting/user",
					},
					2:map[string]interface{}{
						"name":"TODO",
						"select":false,
						"url":"/setting/todo",
					},
				},
			},
		},
		"site_name":"OrderManageSystem",
	}

	//设置选中
	var mainSelectKey int;
	for key, value := range data["nav_list"].(map[int]interface{}){
		//if value.(map[string]interface{})["url"] == url {
		//	data["nav_list"].(map[int]interface{})[int(key)].(map[string]interface{})["select"] = true
		//}else{
		//	data["nav_list"].(map[int]interface{})[int(key)].(map[string]interface{})["select"] = false
		//}
		for k,v := range value.(map[string]interface{})["sub_nav"].(map[int]interface{}){
			if(v.(map[string]interface{})["url"] == url){
				mainSelectKey = int(key)
				data["nav_list"].(map[int]interface{})[int(key)].(map[string]interface{})["sub_nav"].(map[int]interface{})[int(k)].(map[string]interface{})["select"] = true
			}else{
				data["nav_list"].(map[int]interface{})[int(key)].(map[string]interface{})["sub_nav"].(map[int]interface{})[int(k)].(map[string]interface{})["select"] = false
			}
		}
		data["nav_list"].(map[int]interface{})[int(key)].(map[string]interface{})["select"] = false

	}
	data["nav_list"].(map[int]interface{})[mainSelectKey].(map[string]interface{})["select"] = true

	//获取当前登录的用户信息
	ctlbase := *new(BaseController)
	data["_CURRENT_USER"] = ctlbase.GetNeckName()

	return data;
}

//TODO
func (this *BaseController) GetNeckName() string {
	//var ctx = new(beego.Controller)
	return "管理员11"
}