package routers

import (
	"selan/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/router/", &controllers.RouterController{})
	beego.Router("/index/", &controllers.SeLanController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/file", &controllers.FileController{})

}
