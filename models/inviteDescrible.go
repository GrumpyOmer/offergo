package models

import (
	"github.com/astaxie/beego/logs"
	"offergo/connect"
)

type InviteDescrible struct {
	ID            int32  `json:"-"`
	UserDescrible string `json:"user_describle"`
	InviteID      int32  `json:"-"`
}

func (*InviteDescrible) TableName() string {
	return "invite_describle"
}

//修改自定义用户描述
func (*InviteDescrible) UpdateInviteDescribleInfo(where map[string]interface{}, update *map[string]interface{}) (string, bool) {
	updates := connect.Getdb().Table("invite_describle")
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
		logs.Error(updates.Error.Error())
		return "更新失败", false
	}
	return "更新成功", true
}

//查看自定义用户描述
func (*InviteDescrible) GetInviteDescribleInfo(where map[string]interface{}, describle *InviteDescrible) (string, bool) {
	find := connect.Getdb().Table("invite_describle")
	for k, v := range where {
		if v == nil {
			find = find.Where(k)
		} else {
			find = find.Where(k, v)
		}
	}
	row := find.First(describle).RowsAffected
	if row == 0 {
		return "无数据", false
	}
	if find.Error != nil {
		logs.Error(find.Error.Error())
		return "查找失败", false
	}
	return "查找成功", true
}

//新增用户自定义描述
func (*InviteDescrible) AddInviteDescribleInfo(describle *InviteDescrible) (string, bool) {
	add := connect.Getdb().Table("invite_describle").Create(describle)
	row := add.RowsAffected
	if row == 0 {
		return "新增失败", false
	}
	if add.Error != nil {
		logs.Error(add.Error.Error())
		return "新增失败", false
	}
	return "新增成功", true
}
