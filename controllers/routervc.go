package controllers

import (
	"selan/utils"
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

	path := beego.AppConfig.String("domainpath")
	area := c.GetString("area")

	cn := utils.GetValue(path, area, cname).(string)

	printinfo := "path:" + path + "-----------area:" + area + "-----------host:" + host + "--------cname:" + cname + "--------cn:" + cn + "--------port:" + port
	// fmt.Println(printinfo)
	c.Data["Message"] = printinfo

	var FilterUser = func(ctx *context.Context) {
		_, ok := ctx.Input.Session("uid").(int)
		if !ok {
			// ctx.Data["Message"] = "拦截器进入"
			ctx.Redirect(302, "/login")
		}
	}
	beego.InsertFilter("/router/:id([0-9]+)", beego.BeforeRouter, FilterUser)
	c.Data["path"] = path
	c.Data["cname"] = cname
	c.Data["port"] = port
	c.Data["page"] = cn

	c.TplName = "index.html"

}
