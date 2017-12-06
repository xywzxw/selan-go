package controllers

type MainController struct {
	BaseController
}

func (c *MainController) Get() {

	c.Data["Message"] = "色兰网接口-Go语言编写"
	c.TplName = "index.tpl"
}
