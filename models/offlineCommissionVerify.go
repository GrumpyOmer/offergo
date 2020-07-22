package models

import (
	"offergo/connect"
	"offergo/log"
)

type OfflineCommissionVerify struct {
	ID     int32  `gorm:"column:id" json:"id"`
	UserID string `gorm:"column:user_id" json:"user_id"`
	// VerifyMaterial 审核材料
	VerifyMaterial string `gorm:"column:verify_material" json:"verify_material"`
	// VerifyStatus 1.提交审核，2.未通过审核，3.审核通过
	VerifyStatus int8 `gorm:"column:verify_status" json:"verify_status"`
	// DenyReason 拒绝通过原因
	DenyReason      string `gorm:"column:deny_reason" json:"deny_reason"`
	CompletionTimes int32  `gorm:"column:completion_times" json:"completion_times"`
	CreatedAt       int32  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt       int32  `gorm:"column:updated_at" json:"updated_at"`
	// Source 来源\r\n\r\nagent: ''代寄代取'',\r\nbibubi: ''币不币任务'',\r\ntelecom: ''电话卡
	Source string `gorm:"column:source" json:"source"`
}

//待寄待取信息表
func (*OfflineCommissionVerify) TableName() string {
	return "offline_commission_verify"
}

//获取待寄待取发布者信息(多条)
func (*OfflineCommissionVerify) GetOfflineCommission(user *[]OfflineCommissionVerify, where *OfflineCommissionVerify, sel []string, option *map[string]interface{}) (string, bool) {
	getOne := connect.GetHkokDb().
		Select(sel).
		Where(where)
	//wheres
	if data, ok := (*option)["wheres"]; ok && data != "" {
		for k, v := range data.(map[string]interface{}) {
			if v == nil {
				getOne = getOne.Where(k)
			} else {
				getOne = getOne.Where(k, v)
			}
		}
	}

	//Or
	if data, ok := (*option)["or"]; ok && data != "" {
		for k, v := range data.(map[string]interface{}) {
			if v == nil {
				getOne = getOne.Or(k)
			} else {
				getOne = getOne.Or(k, v)
			}
		}
	}

	getOne.Find(user)
	if getOne.Error != nil {
		log.LogInfo.Error(getOne.Error.Error())
		return "查询失败", false
	}
	return "查询成功", true
}
