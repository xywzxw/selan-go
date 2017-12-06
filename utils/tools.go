package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/luckykris/go-utilbox/Conf/ReadConf"
	"github.com/luckykris/go-utilbox/Env"
)

func ChangeMapToJSON(map1 map[string]interface{}) string {
	data, err := json.Marshal(map1)
	if err != nil {
		return "转换json错误"
	}
	return string(data)
}

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
func ReadFile(path string) string {
	dat, _ := ioutil.ReadFile(path)
	return string(dat)
}
func WriteFile(path string, body string) bool {
	fout, err := os.Create(path) //根据路径创建File的内存地址
	defer fout.Close()           //延迟关闭资源
	if err != nil {
		fmt.Println(path, err)
		return false
	}
	fout.WriteString(body)
	return true
}
