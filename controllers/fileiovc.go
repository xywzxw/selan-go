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
	data := createTemp("999", "baidu.com", "www/sd/dd")
	c.Data["path"] = data
	fmt.Println(data)
	c.TplName = "temp.tpl"
	utils.WriteFile("/Users/zxw/www/aa.conf", data)

}
func createTemp(port string, name string, path string) string {
	v := "server{\n    listen " + port + "; \n    server_name " + name + "; \n    root " + path + ";\n}"
	substr := v[0 : len(v)-1]
	substr += utils.ReadFile("views/conf.tpl")
	substr += "}"
	return substr
}
