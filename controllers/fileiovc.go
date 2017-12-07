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

	ngPort := utils.GetValue(path, "nginx", "ngPort").(string)
	ngCname := utils.GetValue(path, "nginx", "ngCname").(string)
	ngPath := utils.GetValue(path, "nginx", "ngPath").(string)
	ngConfPath := utils.GetValue(path, "nginx", "ngConfPath").(string)
	data := utils.CreateTemp(ngPort, ngCname, ngPath)
	// c.Data["path"] = data
	// fmt.Println(data)
	c.TplName = "index.tpl"
	result := utils.WriteFile(ngConfPath, data)
	if result {
		fmt.Println("网站配置创建成功")
	} else {
		fmt.Println("网站配置创建失败")
	}

}
