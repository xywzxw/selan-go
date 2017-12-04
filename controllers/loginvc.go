package controllers

type LoginController struct {
	BaseController
}

func (c *LoginController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["body"] = "这是拦截器进入的登录页面"
	c.TplName = "outpage.tpl"
}
