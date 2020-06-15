package models

import (
	"encoding/json"
	"math"
	"offergo/connect"
	"time"
)

type SzUser struct {
	ID uint32 `gorm:"column:id" json:"id"`
	// UserID 用户id(hkok)
	UserID string `gorm:"column:user_id" json:"user_id"`
	// SzUsername sz预设账号名（hkok生成）
	SzUsername string `gorm:"column:sz_username" json:"sz_username"`
	// SzToken sz账户令牌
	SzToken string `gorm:"column:sz_token" json:"sz_token"`
	// SzID sz账户id
	SzID string `gorm:"column:sz_id" json:"sz_id"`
	// Passcode sz预设密码（hkok生成，统一）
	Passcode string `gorm:"column:passcode" json:"passcode"`
	// Name sz预设昵称（可统一，非必填）
	Name   string `gorm:"column:name"  json:"name"`
	Status int32  `gorm:"column:status" json:"status"`
	// CreatedAt 创建时间
	CreatedAt int32 `gorm:"column:created_at" json:"created_at"`
	// UpdatedAt 更新时间
	UpdatedAt int32 `gorm:"column:updated_at" json:"updated_at"`
}

//神州集运用户信息表
func (*SzUser) TableName() string {
	return "sz_user"
}

//获取神州信息(多个)
func (s *SzUser) GetStatisticsSzUser(monthOneDay string, LastMonthOneDay string) map[string]interface{} {
	type res struct {
		CurrentNum      int     `json:"CurrentUser, omitempty"`      //当前用户数量
		CurrentMonthNum int     `json:"CurrentMonthUser, omitempty"` //当月1号用户数量
		LastMonthNum    int     `json:"LastMonthUser, omitempty"`    //上月1号用户数量
		RateOfIncrease  float64 `json:"Percentage, omitempty"`       //增长率
		Text            string
	}
	var result res
	responseResult := make(map[string]interface{})
	db := connect.Getdb()
	//当前用户数量
	db.Raw("select COUNT(member_code) as current_num from (select member_code,COUNT(DISTINCT member_code) from (select member_code from mb_parcel_info union all select user_id as member_code from sz_user) as d GROUP BY member_code) as c").
		First(&result)
	//本月初用户
	CurrentMonth, _ := time.ParseInLocation("2006-01-02 15:04:05", monthOneDay, time.Local)
	db.Raw("select COUNT(member_code) as current_month_num from (select member_code,COUNT(DISTINCT member_code) from (select member_code from mb_parcel_info WHERE created_at < ? union all select user_id as member_code from sz_user WHERE created_at < ?) as d GROUP BY member_code) as c", monthOneDay, CurrentMonth.Unix()).
		Find(&result)
	//上月初用户
	LastMonth, _ := time.ParseInLocation("2006-01-02 15:04:05", LastMonthOneDay, time.Local)
	db.Raw("select COUNT(member_code) as last_month_num from (select member_code,COUNT(DISTINCT member_code) from (select member_code from mb_parcel_info WHERE created_at < ? union all select user_id as member_code from sz_user WHERE created_at < ?) as d GROUP BY member_code) as c", LastMonthOneDay, LastMonth.Unix()).
		Find(&result)
	//获取增长率
	differ := result.CurrentMonthNum - result.LastMonthNum
	percentage := math.Round(float64(differ)/float64(result.LastMonthNum)*10000) / 100
	result.RateOfIncrease = percentage
	result.Text = "集运服务用户数"
	//转换为字典返回
	bytes, _ := json.Marshal(result)
	json.Unmarshal(bytes, &responseResult)
	return responseResult
}
