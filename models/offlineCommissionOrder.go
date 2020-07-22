package models

import (
	"offergo/connect"
	"offergo/log"
)

type OfflineCommissionOrder struct {
	ID          int32  `gorm:"column:id" json:"id"`
	UserID      string `gorm:"column:user_id" json:"user_id"`
	PublisherID string `gorm:"column:publisher_id" json:"publisher_id"`
	OrderNum    string `gorm:"column:order_num" json:"order_num"`
	// PaymentMethod 1.微信，2.支付宝
	PaymentMethod int8    `gorm:"column:payment_method" json:"payment_method"`
	Paid          float32 `gorm:"column:paid" json:"paid"`
	PayTime       int32   `gorm:"column:pay_time" json:"pay_time"`
	// Status 1.已支付，2.已完成，3.已退款，4.取消订单
	Status       int8   `gorm:"column:status" json:"status"`
	PayerWechat  string `gorm:"column:payer_wechat" json:"payer_wechat"`
	PayerPhone   string `gorm:"column:payer_phone" json:"payer_phone"`
	CommissionID int32  `gorm:"column:commission_id" json:"commission_id"`
	CreatedAt    int32  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt    int32  `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt    int32  `gorm:"column:deleted_at" json:"deleted_at"`
	// RefundReason 退款理由
	RefundReason string `gorm:"column:refund_reason" json:"refund_reason"`
	// MchID 商户ID
	MchID string `gorm:"column:mch_id" json:"mch_id"`
	// TradeType 支付方式
	TradeType string `gorm:"column:trade_type" json:"trade_type"`
	// TransactionID 微信支付订单号
	TransactionID string `gorm:"column:transaction_id" json:"transaction_id"`
	// Appid 公众账号ID
	Appid string `gorm:"column:appid" json:"appid"`
	// PayType 支付途径
	PayType string `gorm:"column:pay_type" json:"pay_type"`
}

func (*OfflineCommissionOrder) TableName() string {
	return "offline_commission_order"
}

//获取待寄待取付费者信息(多条)
func (*OfflineCommissionOrder) GetOfflineCommission(user *[]OfflineCommissionOrder, where *OfflineCommissionOrder, sel []string, option *map[string]interface{}) (string, bool) {
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
