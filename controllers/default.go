package controllers

import (
	"selan/utils"

	"github.com/astaxie/beego"
)

type MainController struct {
	BaseController
}

func (c *MainController) Get() {

	path := beego.AppConfig.String("dmainpath")
	area := c.GetString("area")
	key := c.GetString("key")
	values := utils.GetValue(path, area, key)
	// fmt.Println(values)
	str := "读取配置文件[area:" + area + "--------key:" + key + "-----value:" + values.(string) + "]"
	c.Data["Message"] = str
	c.TplName = "index.tpl"
}
