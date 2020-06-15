package lib

import (
	"encoding/json"
	"offergo/models"
)

//结构体去重操作
func DuplicateRemoval(data interface{}, str interface{}) interface{} {
	//字典存放切片
	resultMap := make(map[string]interface{})
	switch str.(type) {
	case models.SecondHandInfo:
		temp := data.([]models.SecondHandInfo)
		for _, v := range temp {
			data, _ := json.Marshal(v)
			resultMap[string(data)] = true
		}
		result := []models.SecondHandInfo{}
		for k := range resultMap {
			var t models.SecondHandInfo
			json.Unmarshal([]byte(k), &t)
			result = append(result, t)
		}
		return result
	}
	return nil
}

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
