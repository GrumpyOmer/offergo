package models

import (
	"offergo/connect"
	"offergo/log"
)

type JobList struct {
	ApplyLink    string `gorm:"column:apply_link" json:"apply_link"`
	BrowseNum    int    `gorm:"column:browse_num" json:"browse_num"`
	ClickNum     int    `gorm:"column:click_num" json:"click_num"`
	Company      string `gorm:"column:company" json:"company"`
	CreatedAt    int    `gorm:"column:created_at" json:"created_at"`
	Details      string `gorm:"column:details" json:"details"`
	EndTime      int64  `gorm:"column:end_time" json:"end_time"`
	Experience   int    `gorm:"column:experience" json:"experience"`
	Highlighted  int    `gorm:"column:highlighted" json:"highlighted"`
	ID           int    `gorm:"column:id;primary_key" json:"id"`
	Industry     int    `gorm:"column:industry" json:"industry"`
	Labels       string `gorm:"column:labels" json:"labels"`
	LogoURL      string `gorm:"column:logo_url" json:"logo_url"`
	Place        string `gorm:"column:place" json:"place"`
	PlusPosition int    `gorm:"column:plus_position" json:"plus_position"`
	Position     string `gorm:"column:position" json:"position"`
	PushStatus   int    `gorm:"column:push_status" json:"push_status"`
	PushTime     int64  `gorm:"column:push_time" json:"push_time"`
	ShareNum     int    `gorm:"column:share_num" json:"share_num"`
	ShowTime     int64  `gorm:"column:show_time" json:"show_time"`
	SortMethod   int    `gorm:"column:sort_method" json:"sort_method"`
	Status       int    `gorm:"column:status" json:"status"`
	UpdatedAt    int    `gorm:"column:updated_at" json:"updated_at"`
}

// TableName sets the insert table name for this struct type
func (j *JobList) TableName() string {
	return "job_list"
}

func (j *JobList) GetJobList(user *[]JobList, where *JobList, sel []string, option *map[string]interface{}) (string, bool) {
	getMany := connect.GetHkokJobDb().
		Table("job_list").
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
