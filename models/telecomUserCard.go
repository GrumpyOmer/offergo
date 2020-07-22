package models

import (
	"offergo/connect"
	"offergo/log"
)

type TelecomUserCard struct {
	ID int32 `gorm:"column:id" json:"id"`
	// UserID 用户ID
	UserID string `gorm:"column:user_id" json:"user_id"`
	// Mobile 手机号码
	Mobile string `gorm:"column:mobile" json:"mobile"`
	// RealName 真实姓名
	RealName string `gorm:"column:real_name" json:"real_name"`
	// Balance 余额
	Balance float32 `gorm:"column:balance" json:"balance"`
	// ServerPlan 1.小K，2.大K
	ServerPlan int8   `gorm:"column:server_plan" json:"server_plan"`
	Iccid      string `gorm:"column:iccid" json:"iccid"`
	Imsi       string `gorm:"column:imsi" json:"imsi"`
	// CardStatus -1.停机！0.待开通1.已开通2.欠费！3.冻结4.注销5.解绑6.过期
	CardStatus int8 `gorm:"column:card_status" json:"card_status"`
	// ActivateDate 激活日期
	ActivateDate string `gorm:"column:activate_date" json:"activate_date"`
	// ExpiredAt 激活的截止日期
	ExpiredAt string `gorm:"column:expired_at" json:"expired_at"`
	CreatedAt string `gorm:"column:created_at" json:"created_at"`
	UpdatedAt string `gorm:"column:updated_at" json:"updated_at"`
	// CancellationStatus 是否销户用户
	CancellationStatus uint32 `gorm:"column:cancellation_status" json:"cancellation_status"`
	// Type 卡获取类型:BUY：购买，GIFT：新生礼包
	Type string `gorm:"column:type" json:"type"`
	// DeductionStatus 是否扣费
	DeductionStatus uint32 `gorm:"column:deduction_status" json:"deduction_status"`
	// FreezeTime 最近冻结时间
	FreezeTime string `gorm:"column:freeze_time" json:"freeze_time"`
	// RestoreTime 最近恢复时间
	RestoreTime string `gorm:"column:restore_time" json:"restore_time"`
	// Deduction 扣费记录
	Deduction int32 `gorm:"column:deduction" json:"deduction"`
}

//大K卡用户信息表
func (*TelecomUserCard) TableName() string {
	return "user_card"
}

//获取大K卡用户信息(多条)
func (*TelecomUserCard) GetTelecomUser(user *[]TelecomUserCard, where *TelecomUserCard, sel []string, option *map[string]interface{}) (string, bool) {
	getMany := connect.GetTelecomDb().
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
	getMany.Find(user)
	if getMany.Error != nil {
		log.LogInfo.Error(getMany.Error.Error())
		return "查询失败", false
	}
	return "查询成功", true
}
