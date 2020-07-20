package lib

import (
	"encoding/json"
)

//结构体转字典
func StructToMap(data interface{}, result *map[string]interface{}) {
	bytes, _ := json.Marshal(data)
	json.Unmarshal(bytes, result)
}

//字典转结构体
func MapToStruct(data interface{}, result interface{}) {
	bytes, _ := json.Marshal(data)
	json.Unmarshal(bytes, result)
}
