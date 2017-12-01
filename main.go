package main

import (
	_ "oms/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionAutoSetCookie = true
	beego.Run()
}

