package utils

import (
	"encoding/json"

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
