package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/astaxie/beego"
	"github.com/luckykris/go-utilbox/Conf/ReadConf"
	"github.com/luckykris/go-utilbox/Env"
)

//ChangeMapToJSON 将map转换为json数据
func ChangeMapToJSON(map1 map[string]interface{}) string {
	data, err := json.Marshal(map1)
	if err != nil {
		return "转换json错误"
	}
	return string(data)
}

//GetValue 从配置文件中读取数据
//path 配置文件的路径
//area 读取的区域("[]"里的)
//key  键值
//return 读取到的数据
func GetValue(path string, area string, key string) interface{} {
	config := ReadConf.CONFIG{
		AREA: area,
		CONF: map[string]ReadConf.CONFIGROW{
			key: {TYPE: "string", EXIT: false, DEFAULT: ""},
		},
	}
	val := area + "/" + key
	ReadConf.LoadConf(path, config)
	return Env.ENV[val]
}

//ReadFile 读取文件
//@param path 读取文件的路径
//return 读出的字符串
func ReadFile(path string) string {
	dat, _ := ioutil.ReadFile(path)
	return string(dat)
}

//WriteFile 写入文件
//@param path 写入的文件
//@param body 写入的内容
//return 写入的结果
func WriteFile(path string, body string) bool {
	fout, err := os.Create(path) //根据路径创建File的内存地址
	defer fout.Close()           //延迟关闭资源
	if err != nil {
		if os.IsNotExist(err) {
			h := strings.Split(path, "/") //拆串
			m := h[len(h)-1]
			v := path[0 : len(path)-len(m)]
			err1 := os.MkdirAll(v, os.ModePerm)
			var result bool
			if err1 != nil {
				fmt.Println("项目文件目录创建失败")
				result = false
			} else {
				result = WriteFile(path, body)
			}
			return result
		}
		fmt.Println(path, err)
		return false
	}
	fout.WriteString(body)
	return true
}

//CreateTemp 创建配置文件
//port 端口号
//name 域名
//path 网站存放路径
//return 生成配置文件字符串
func CreateTemp(port string, name string, path string) string {
	nginxpath := beego.AppConfig.String("domainpath")
	v := "server{\n    listen " + port + "; \n    server_name " + name + "; \n    root " + path + ";\n    access_log  /www/wwwlogs/" + name + ".log;\n}"
	substr := v[0 : len(v)-1]
	substr += ReadFile(GetValue(nginxpath, "nginx", "nginxConf").(string))
	substr += "}"
	return substr
}
