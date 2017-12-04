package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
)

// BaseController 基控制器
type BaseController struct {
	beego.Controller
}

// Prepare 基控制器所有请求先走的方法
func (this *BaseController) Prepare() {
	method := this.Ctx.Request.Method
	if method == "GET" {
		fmt.Println("GET处理的")
		dealGET(this)
	} else if method == "POST" {
		fmt.Println("POST处理的")
		dealPOST(this)
	}
}
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
func dealPOST(this *BaseController) {
	dealGET(this)
}
