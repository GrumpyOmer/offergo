package models

import (
	"offergo/connect"
	"offergo/log"
)

type SecondHandInfo struct {
	Shareid          string `gorm:"column:shareid" json:"shareid"`
	UserID           string `gorm:"column:user_id" json:"user_id"`
	Openid           string `gorm:"column:openid" json:"openid"`
	MemberCode       string `gorm:"column:member_code" json:"member_code"`
	Type             string `gorm:"column:type" json:"type"`
	Title            string `gorm:"column:title" json:"title"`
	ReviewCount      int32  `gorm:"column:review_count" json:"review_count"`
	PostTimelineTime int32  `gorm:"column:post_timeline_time" json:"post_timeline_time"`
	PostAppTime      int32  `gorm:"column:post_app_time" json:"post_app_time"`
	Price            string `gorm:"column:price" json:"price"`
	PriceTag         string `gorm:"column:price_tag" json:"price_tag"`
	Detail           string `gorm:"column:detail" json:"detail"`
	Contact          string `gorm:"column:contact" json:"contact"`
	ContactTag       string `gorm:"column:contact_tag" json:"contact_tag"`
	Addr             string `gorm:"column:addr" json:"addr"`
	AddrTag          string `gorm:"column:addr_tag" json:"addr_tag"`
	Imgurl1          string `gorm:"column:imgurl_1" json:"imgurl_1"`
	Imgurl2          string `gorm:"column:imgurl_2" json:"imgurl_2"`
	PostAt           string `gorm:"column:post_at" json:"post_at"`
	UpdatedAt        string `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt        string `gorm:"column:deleted_at" json:"deleted_at"`
	ExpiredAt        string `gorm:"column:expired_at" json:"expired_at"`
	ExpiredStatus    string `gorm:"column:expired_status" json:"expired_status"`
	Status           string `gorm:"column:status" json:"status"`
	// IsChecking  是否审核:0默认待审核，1为审核通过，2为审核不通过
	IsChecking int8 `gorm:"column:is_checking" json:"is_checking"`
	// Like 点赞数
	Like int32 `gorm:"column:like" json:"like"`
}

//二手信息表
func (*SecondHandInfo) TableName() string {
	return "second_hand_info"
}

//获取二手用户信息(多个)
func (*SecondHandInfo) GetSecondHandUser(user *[]SecondHandInfo, where *SecondHandInfo, sel []string, option *map[string]interface{}) (string, bool) {
	getMany := connect.Getdb().
		Select(sel).
		Where(where)

	//wheres
	if data, ok := (*option)["wheres"]; ok {
		for k, v := range data.(map[string]interface{}) {
			getMany = getMany.Where(k, v)
		}
	}
	getMany.Find(user)
	if getMany.Error != nil {
		log.LogInfo.Error(getMany.Error.Error())
		return "查询失败", false
	}
	return "查询成功", true
}
