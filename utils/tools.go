package utils

import "encoding/json"

func ChangeMapToJSON(map1 map[string]interface{}) string {
	data, err := json.Marshal(map1)
	if err != nil {
		return "转换json错误"
	}
	return string(data)
}
