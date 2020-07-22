package models

import (
	"offergo/connect"
	"offergo/lib"
	"offergo/log"
	"strconv"
)

type TelecomCard struct {
	//字段改成interface{}的字段,是因为字段为0时是有意义的不应该被过滤掉，同时不想要为不同的select返回数量去创造不一样的结构体来返回，太麻烦了！！！解决参考案例：https://blog.csdn.net/m0_37422289/article/details/103569013
	ID int32 `json:"id" form:"-"`
	// Status 0.可用，1.已使用，2.已停用
	Status    int8   `json:"status" form:"status"`
	Imsi      int8   `json:"imsi" form:"imsi"`
	Iccid     string `json:"iccid" form:"iccid"`
	CreatedAt string `json:"created_at,omitempty" form:"created_at"`
	// WhiteCard 是否白卡：1白卡
	WhiteCard int8 `json:"white_card" form:"white_card"`
	// IsProvide 白卡是否发放：0 未发放 1已发放
	IsProvide int8 `json:"is_provide" form:"is_provide"`
	//渠道关联id
	ChannelId int8 `json:"-" form:"channel_id"`
	// 关联渠道
	WhiteCardChannel WhiteCardChannel `json:"white_card_channel" gorm:"foreignkey:id;association_foreignkey:channel_id"`
}

//大K卡表（iccid）
func (*TelecomCard) TableName() string {
	return "cards"
}

//获取cards信息(多条)
func (*TelecomCard) GetTelecomCard(result *[]TelecomCard, where *TelecomCard, sel []string, option *map[string]interface{}) (string, bool) {
	getMany := connect.GetTelecomDb().
		Table("cards").
		Select(sel).
		Where(where)
	//wheres
	if data, ok := (*option)["wheres"]; ok && data != "" {
		for k, v := range data.(map[string]interface{}) {
			if v == nil {
				getMany = getMany.Where(k)
			} else {
				getMany = getMany.Where(k, v)
			}
		}
	}

	//or
	if data, ok := (*option)["or"]; ok && data != "" {
		for k, v := range data.(map[string]interface{}) {
			if v == nil {
				getMany = getMany.Or(k)
			} else {
				getMany = getMany.Or(k, v)
			}
		}
	}

	//join
	if data, ok := (*option)["join"]; ok && data != "" {
		for k, _ := range data.(map[string]interface{}) {
			getMany = getMany.Joins(k)
		}
	}

	//page
	if data, ok := (*option)["page"]; ok && data != "" {
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

	getMany = getMany.Preload("WhiteCardChannel")
	getMany.Find(result)
	if getMany.Error != nil {
		log.LogInfo.Error(getMany.Error.Error())
		return "查询失败", false
	}

	return "查询成功", true
}

//修改cards
func (*TelecomCard) UpdateTelecomCard(where map[string]interface{}, update *map[string]interface{}) (string, bool) {
	updates := connect.GetTelecomDb().
		Table("cards")
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
