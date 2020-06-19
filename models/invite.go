package models

import (
	"offergo/connect"
	"offergo/log"
	"strconv"
)

type Invite struct {
	ID int32 `json:"id"`
	// InviteID id
	InviteID int32 `json:"invite_id"`
	// InviteName 名称
	InviteName string `json:"invite_name"`
	// InviteAddress 地址
	InviteAddress string `json:"invite_address"`
	// InviteArea 地区
	InviteArea int8 `json:"invite_area"`
	// APIDescrible 接口自取点描述\n\n
	APIDescrible string `json:"api_describle"`
	// Status 是否显示 0 否 1 是
	Status int8 `json:"status"`
	//地点id
	InviteLocation int8 `json:"invite_location"`
	// InviteDescrible 用户端自取点描述
	InviteDescrible InviteDescrible `json:"invite_describle,omitempty"`
}

func (*Invite) TableName() string {
	return "invite"
}

//获取自取点列表
func (*Invite) GetInviteList(result *[]Invite, where *Invite, sel []string, option *map[string]interface{}) (string, bool) {
	getMany := connect.Getdb().
		Table("invite").
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
		PageNum := PAGENUM
		page, _ := data.(map[string]interface{})
		pageToString, _ := page["page"].(string)
		currentPage, _ := strconv.Atoi(pageToString)
		//总条数
		var total float64
		getMany.Count(&total)
		var getInfo PageStruct
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
		new(PageStruct).getPage(total, currentPage, &getInfo, PageNum)
		(*option)["pageInfo"] = getInfo
	}
	getMany = getMany.Preload("InviteDescrible")
	getMany = getMany.Order("invite_area ASC").Order("invite_location ASC").Order("invite_id ASC")
	getMany.Find(result)
	if getMany.Error != nil {
		log.LogInfo.Error(getMany.Error.Error())
		return "查询失败", false
	}

	return "查询成功", true
}

//清空自取点列表
func (*Invite) DeleteInviteList() {
	connect.Getdb().Delete(Invite{})
}

//修改自取点信息
func (*Invite) UpdateInviteInfo(where map[string]interface{}, update *map[string]interface{}) (string, bool) {
	updates := connect.Getdb().Table("invite")
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

//insert many records
func(*Invite) InsertManyRecords(sql string) (string, bool) {
	result:= connect.Getdb().Exec(sql)
	if result.Error == nil {
		return "批量添加成功",true
	}
	return result.Error.Error(),false
}