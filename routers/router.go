package routers

import (
	"selan/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/index/", &controllers.SeLanController{})
	beego.Router("/login", &controllers.LoginController{})

}
