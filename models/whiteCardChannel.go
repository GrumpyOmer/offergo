package models

import (
	"github.com/jinzhu/gorm"
	"offergo/connect"
	"offergo/lib"
	"offergo/log"
	"strconv"
	"time"
)

type WhiteCardChannel struct {
	ID uint32 `json:"id"`
	// Name 渠道名
	Name string `json:"name"`
	// Balance 初始余额
	Balance int32 `json:"balance"`
	// Status 0过期1可用
	Status int8 `json:"status"`
	//创建时间
	CreatedAt int32 `json:"created_at"`
}

//白卡渠道表
func (*WhiteCardChannel) TableName() string {
	return "whiteCardChannel"
}

//获取白卡渠道列表
func (*WhiteCardChannel) GetWhiteCardChannel(result *[]WhiteCardChannel, where *WhiteCardChannel, sel []string, option *map[string]interface{}) (string, bool) {
	getMany := connect.Gettdb().
		Table("whiteCardChannel").
		Select(sel).
		Where(where)
	//wheres
	if data, ok := (*option)["wheres"]; ok {
		for k, v := range data.(map[string]interface{}) {
			if v == nil {
				getMany = getMany.Where(k)
			} else {
				getMany = getMany.Where(k, v)
			}
		}
	}

	//or
	if data, ok := (*option)["or"]; ok {
		for k, v := range data.(map[string]interface{}) {
			if v == nil {
				getMany = getMany.Or(k)
			} else {
				getMany = getMany.Or(k, v)
			}
		}
	}

	//order by
	if data, ok := (*option)["orderBy"]; ok {
		getMany = getMany.Order(data)
	}

	//join
	if data, ok := (*option)["join"]; ok {
		for k, _ := range data.(map[string]interface{}) {
			getMany = getMany.Joins(k)
		}
	}

	//page
	if data, ok := (*option)["page"]; ok {
		PageNum := lib.PAGENUM
		page, _ := data.(map[string]interface{})
		pageToString, _ := page["page"].(string)
		currentPage, _ := strconv.Atoi(pageToString)
		//总条数
		var total float64
		getMany.Count(&total)
		var getInfo lib.PageStruct
		if pageNum, ok := page["pageNum"]; ok && pageNum != "" {
			pageNumToString, _ := pageNum.(string)
			pageNumToInt, _ := strconv.Atoi(pageNumToString)
			//有传参页面展示数据数量则使用传参值 否则使用默认的每页展示条数
			getMany = getMany.Offset((currentPage - 1) * pageNumToInt).Limit(pageNumToInt)
			PageNum = pageNumToInt
		} else {
			getMany = getMany.Offset((currentPage - 1) * PageNum).Limit(PageNum)
		}
		//获取分页信息
		new(lib.PageStruct).GetPage(total, currentPage, &getInfo, PageNum)
		(*option)["pageInfo"] = getInfo
	}

	getMany.Find(result)
	if getMany.Error != nil {
		log.LogInfo.Error(getMany.Error.Error())
		return "查询失败", false
	}

	return "查询成功", true
}

//新增白卡渠道
func (*WhiteCardChannel) CreateWhiteCardChannel(data *WhiteCardChannel) (string, bool) {
	create := connect.Gettdb().
		Omit("status").
		Create(data)
	if create.Error != nil {
		log.LogInfo.Error(create.Error.Error())
		return "添加失败", false
	}
	return "添加成功", true
}

//修改白卡渠道状态
func (*WhiteCardChannel) UpdateWhiteCardChannel(where map[string]interface{}, update *map[string]interface{}) (string, bool) {
	updates := connect.Gettdb().Table("whiteCardChannel")
	for k, v := range where {
		if v == nil {
			updates = updates.Where(k)
		} else {
			updates = updates.Where(k, v)
		}
	}
	row := updates.Updates(*update).RowsAffected
	if row == 0 {
		return "无任何更新", false
	}
	if updates.Error != nil {
		log.LogInfo.Error(updates.Error.Error())
		return "更新失败", false
	}
	return "更新成功", true
}

//钩子时间，设置created_at字段
func (w *WhiteCardChannel) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("created_at", time.Now().Unix())
	return nil
}
