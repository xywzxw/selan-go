package controllers

type MainController struct {
	BaseController
}

func (c *MainController) Get() {
	c.Data["Message"] = "色兰网数据接口-go语言"
	c.TplName = "index.tpl"
}
