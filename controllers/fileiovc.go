package controllers

import (
	"fmt"
	"selan/utils"
)

type FileController struct {
	BaseController
}

func (c *FileController) Get() {

	// c.Data["cname"] = "www.zxw.com"
	data := createTemp("999我去玩", "baidu.com", "www/sd/dd")
	c.Data["path"] = data
	fmt.Println(data)
	c.TplName = "temp.tpl"
	result := utils.WriteFile("/Users/zxw/www/aa.conf", data)
	if result {
		fmt.Println("网站配置创建成功")
	} else {
		fmt.Println("网站配置创建失败")
	}
}
func createTemp(port string, name string, path string) string {
	v := "server{\n    listen " + port + "; \n    server_name " + name + "; \n    root " + path + ";\n}"
	substr := v[0 : len(v)-1]
	substr += utils.ReadFile("views/conf.tpl")
	substr += "}"
	return substr
}
