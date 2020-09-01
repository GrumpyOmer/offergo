package models

import (
	"offergo/connect"
	"offergo/log"
)

type KPlusPrepairdOrder struct {
	CreatedAt      int64   `gorm:"column:created_at" json:"created_at"`
	DeletedAt      int64   `gorm:"column:deleted_at" json:"deleted_at"`
	ID             int64   `gorm:"column:id;primary_key" json:"id"`
	OrderNum       string  `gorm:"column:order_num" json:"order_num"`
	PayStatus      int64   `gorm:"column:pay_status" json:"pay_status"`
	PayTime        int64   `gorm:"column:pay_time" json:"pay_time"`
	Payment        int64   `gorm:"column:payment" json:"payment"`
	PrepaidEndTime int64   `gorm:"column:prepaid_end_time" json:"prepaid_end_time"`
	Price          float64 `gorm:"column:price" json:"price"`
	TradeNo        string  `gorm:"column:trade_no" json:"trade_no"`
	Type           int64   `gorm:"column:type" json:"type"`
	UpdatedAt      int64   `gorm:"column:updated_at" json:"updated_at"`
	UserID         string  `gorm:"column:user_id" json:"user_id"`
}

// TableName sets the insert table name for this struct type
func (k *KPlusPrepairdOrder) TableName() string {
	return "k_plus_prepaid_order"
}

//获取k_plus分期用户 (多个)
func (*KPlusPrepairdOrder) GetKPlusUser(user *[]KPlusPrepairdOrder, where *KPlusPrepairdOrder, sel []string, option *map[string]interface{}) (string, bool) {
	getMany := connect.GetHkokDb().
		Table("k_plus_prepaid_order").
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
