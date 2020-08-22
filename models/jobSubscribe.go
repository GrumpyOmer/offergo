package models

import (
	"offergo/connect"
	"offergo/log"
)

type JobSubscribe struct {
	CreatedAt        int64  `gorm:"column:created_at" json:"created_at"`
	DeletedAt        int64  `gorm:"column:deleted_at" json:"deleted_at"`
	ID               int    `gorm:"column:id;primary_key" json:"id"`
	Industry         string `gorm:"column:industry" json:"industry"`
	JobSubscribeArea string `gorm:"column:job_subscribe_area" json:"job_subscribe_area"`
	JobSubscribeType string `gorm:"column:job_subscribe_type" json:"job_subscribe_type"`
	UpdatedAt        string `gorm:"column:updated_at" json:"updated_at"`
	UserID           string `gorm:"column:user_id" json:"user_id"`
}

// TableName sets the insert table name for this struct type
func (j *JobSubscribe) TableName() string {
	return "job_subscribe"
}

func (j *JobSubscribe) GetJobSubscribe(user *[]JobSubscribe, where *JobSubscribe, sel []string, option *map[string]interface{}) (string, bool) {
	getMany := connect.GetHkokJobDb().
		Table("job_subscribe").
		Select(sel).
		Where(where).
		Unscoped()
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

	//Or
	if data, ok := (*option)["or"]; ok && data != "" {
		for k, v := range data.(map[string]interface{}) {
			if v == nil {
				getMany = getMany.Or(k)
			} else {
				getMany = getMany.Or(k, v)
			}
		}
	}

	getMany.Find(user)

	//Count
	if data, ok := (*option)["count"]; ok {
		getMany.Count(data)
	}

	if getMany.Error != nil {
		log.LogInfo.Error(getMany.Error.Error())
		return "查询失败", false
	}
	return "查询成功", true
}
