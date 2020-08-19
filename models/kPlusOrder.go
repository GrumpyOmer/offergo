package models

import (
	"offergo/connect"
	"offergo/log"
)

type KPlusOrder struct {
	CreatedAt    int64   `gorm:"column:created_at" json:"created_at"`
	DeletedAt    int64   `gorm:"column:deleted_at" json:"deleted_at"`
	ID           int     `gorm:"column:id;primary_key" json:"id"`
	OrderNum     string  `gorm:"column:order_num" json:"order_num"`
	PayStatus    int64   `gorm:"column:pay_status" json:"pay_status"`
	PayTime      int64   `gorm:"column:pay_time" json:"pay_time"`
	Payment      int64   `gorm:"column:payment" json:"payment"`
	PeriodNum    int64   `gorm:"column:period_num" json:"period_num"`
	PeriodPrice  float64 `gorm:"column:period_price" json:"period_price"`
	TotalPrice   float64 `gorm:"column:total_price" json:"total_price"`
	TradeNo      string  `gorm:"column:trade_no" json:"trade_no"`
	UpdatedAt    int64   `gorm:"column:updated_at" json:"updated_at"`
	UserID       string  `gorm:"column:user_id" json:"user_id"`
	VipStartTime int64   `gorm:"column:vip_start_time" json:"vip_start_time"`
}

// TableName sets the insert table name for this struct type
func (k *KPlusOrder) TableName() string {
	return "k_plus_order"
}

//获取k_plus用户 (多个)
func (*KPlusOrder) GetKPlusUser(user *[]KPlusOrder, where *KPlusOrder, sel []string, option *map[string]interface{}) (string, bool) {
	getMany := connect.GetHkokDb().
		Table("k_plus_order").
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
	//group
	if data, ok := (*option)["group"]; ok {
		getMany = getMany.Group(data.(string))
	}
	getMany = getMany.Find(user)

	//count
	if data, ok := (*option)["count"]; ok {
		getMany = getMany.Count(data)
	}

	if getMany.Error != nil {
		log.LogInfo.Error(getMany.Error.Error())
		return "查询失败", false
	}
	return "查询成功", true
}
