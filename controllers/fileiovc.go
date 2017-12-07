package controllers

import (
	"fmt"
	"selan/utils"

	"github.com/astaxie/beego"
)

type FileController struct {
	BaseController
}

func (c *FileController) Get() {
	path := beego.AppConfig.String("domainpath")

	ng_port := utils.GetValue(path, "nginx", "ng_port").(string)
	ng_cname := utils.GetValue(path, "nginx", "ng_cname").(string)
	ng_path := utils.GetValue(path, "nginx", "ng_path").(string)
	ng_conf_path := utils.GetValue(path, "nginx", "ng_conf_path").(string)
	data := utils.CreateTemp(ng_port, ng_cname, ng_path)
	c.Data["path"] = data
	fmt.Println(data)
	c.TplName = "temp.tpl"
	result := utils.WriteFile(ng_conf_path, data)
	if result {
		fmt.Println("网站配置创建成功")
	} else {
		fmt.Println("网站配置创建失败")
	}

}
