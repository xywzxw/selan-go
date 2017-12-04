package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

// BaseController 基控制器
type BaseController struct {
	beego.Controller
}

// Prepare 基控制器所有请求先走的方法
func (this *BaseController) Prepare() {
	filter(this)
	method := this.Ctx.Request.Method
	if method == "GET" {
		fmt.Println("GET处理的")
		dealGET(this)
	} else if method == "POST" {
		fmt.Println("POST处理的")
		dealPOST(this)
	}
}

//处理GET请求
func dealGET(this *BaseController) {
	path := this.Ctx.Request.URL.Path
	host := this.Ctx.Request.Host
	method := this.Ctx.Request.Method
	req := this.Ctx.Request
	req.ParseForm()

	if len(req.Form) > 0 {
		for k, v := range req.Form {
			fmt.Println("k:" + k + "----v:" + v[0])
		}
	}
	fmt.Println(method + "---" + path + "---" + host)
}

//处理POST请求
func dealPOST(this *BaseController) {
	dealGET(this)
}

//拦截器
func filter(this *BaseController) {
	fmt.Println("即将进入拦截器")
	var FilterUser = func(ctx *context.Context) {
		fmt.Println("进入拦截器")
		uid := this.GetString("uid")
		if len(uid) == 0 && ctx.Request.RequestURI != "/login" {
			ctx.Redirect(302, "/login")
		}
	}
	beego.InsertFilter("/index", beego.BeforeRouter, FilterUser)
}
