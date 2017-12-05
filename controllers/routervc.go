package controllers

import (
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type RouterController struct {
	BaseController
}

func (c *RouterController) Get() {

	host := c.Ctx.Request.Host

	h := strings.Split(host, ":") //拆串

	cname := h[0] //域名
	port := h[1]  //端口号

	cn := beego.AppConfig.String(cname) //从配置文件中获取域名对应的页面

	printinfo := "host:" + host + "--------" + "cname:" + cname + "--------" + "cn:" + cn + "--------" + "port:" + port

	c.Data["Message"] = printinfo

	var FilterUser = func(ctx *context.Context) {
		_, ok := ctx.Input.Session("uid").(int)
		if !ok {
			// ctx.Data["Message"] = "拦截器进入"
			ctx.Redirect(302, "/index")
		}
	}
	beego.InsertFilter("/router/:id([0-9]+)", beego.BeforeRouter, FilterUser)

	c.TplName = cn
}
