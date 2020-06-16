package models

import (
	"offergo/connect"
	"offergo/log"
)

type User struct {
	// UserID 用户唯一ID
	UserID string `gorm:"column:user_id" json:"user_id"`
	// Username 账号登录名（邮箱）
	Username string `gorm:"column:username" json:"username"`
	// Nickname 昵称
	Nickname string `gorm:"column:nickname" json:"nickname"`
	// Password 密码
	Password string `gorm:"column:password" json:"password"`
	// Sex 性别
	Sex string `gorm:"column:sex" json:"sex"`
	// Age 年龄
	Age int32 `gorm:"column:age" json:"age"`
	// City 城市
	City string `gorm:"column:city" json:"city"`
	// Home app家乡个人资料那边的展示
	Home         string `gorm:"column:home" json:"home"`
	Province     string `gorm:"column:province" json:"province"`
	Country      string `gorm:"column:country" json:"country"`
	School       string `gorm:"column:school" json:"school"`
	EnterYear    int32  `gorm:"column:enter_year" json:"enter_year"`
	Major        string `gorm:"column:major" json:"major"`
	CnSchool     string `gorm:"column:cn_school" json:"cn_school"`
	IsSearchable int32  `gorm:"column:is_searchable" json:"is_searchable"`
	// IsVerified offer认证:0未认证 1已认证 2待认证 3未通过
	IsVerified           int32  `gorm:"column:is_verified" json:"is_verified"`
	ContactWx            string `gorm:"column:contact_wx" json:"contact_wx"`
	VerFileName          string `gorm:"column:ver_file_name" json:"ver_file_name"`
	Headimgurl           string `gorm:"column:headimgurl" json:"headimgurl"`
	Openid               string `gorm:"column:openid" json:"openid"`
	MemberCode           string `gorm:"column:member_code" json:"member_code"`
	SelfDesc             string `gorm:"column:self_desc" json:"self_desc"`
	LiveCity             string `gorm:"column:live_city" json:"live_city"`
	MatchSex             string `gorm:"column:match_sex" json:"match_sex"`
	MatchCity            string `gorm:"column:match_city" json:"match_city"`
	PsHint               string `gorm:"column:ps_hint" json:"ps_hint"`
	HtAnswer             string `gorm:"column:ht_answer" json:"ht_answer"`
	GenDate              string `gorm:"column:gen_date" json:"gen_date"`
	LastLoginDate        string `gorm:"column:last_login_date" json:"last_login_date"`
	LoginCount           int32  `gorm:"column:login_count" json:"login_count"`
	Email                string `gorm:"column:email" json:"email"`
	UserAuth             string `gorm:"column:user_auth" json:"user_auth"`
	UserStatus           int32  `gorm:"column:user_status" json:"user_status"`
	SearchCount          int32  `gorm:"column:search_count" json:"search_count"`
	FindCount            int32  `gorm:"column:find_count" json:"find_count"`
	Vote                 string `gorm:"column:vote" json:"vote"`
	TmStatus             string `gorm:"column:tm_status" json:"tm_status"`
	RememberToken        string `gorm:"column:remember_token" json:"remember_token"`
	SubscriptionStartsAt string `gorm:"column:subscription_starts_at" json:"subscription_starts_at"`
	SubscriptionEndsAt   string `gorm:"column:subscription_ends_at" json:"subscription_ends_at"`
	RewardPoint          int32  `gorm:"column:reward_point" json:"reward_point"`
	CreatedAt            string `gorm:"column:created_at" json:"created_at"`
	UpdatedAt            string `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt            string `gorm:"column:deleted_at" json:"deleted_at"`
	CheckinAt            string `gorm:"column:checkin_at" json:"checkin_at"`
	// IsCheckIn 是否开启签到消息提醒（每天八点）
	IsCheckIn          int32  `gorm:"column:is_check_in" json:"is_check_in"`
	JobSubscriber      int32  `gorm:"column:job_subscriber" json:"job_subscriber"`
	ActivitySubscriber int32  `gorm:"column:activity_subscriber" json:"activity_subscriber"`
	RepostStatus       int32  `gorm:"column:repost_status" json:"repost_status"`
	VerCode            string `gorm:"column:ver_code" json:"ver_code"`
	VerCodeAt          string `gorm:"column:ver_code_at" json:"ver_code_at"`
	FirstInHouse       int32  `gorm:"column:first_in_house" json:"first_in_house"`
	LastNotice         int32  `gorm:"column:last_notice" json:"last_notice"`
	HasSaveChatRecord  int32  `gorm:"column:has_save_chat_record" json:"has_save_chat_record"`
	// Unionid 用户在公众号和公众平台关联的app应用唯一id
	Unionid string `gorm:"column:unionid" json:"unionid"`
	// Appopenid 开放平台关联公众号的一个用户识别id（唯一）
	Appopenid string `gorm:"column:appopenid" json:"appopenid"`
	// Sign 当月签到的日期记录
	Sign string `gorm:"column:sign" json:"sign"`
	// Collect 收藏文章记录
	Collect            string `gorm:"column:collect" json:"collect"`
	AndroiddeviceToken string `gorm:"column:androiddeviceToken" json:"androiddeviceToken"`
	AndroiddeviceType  string `gorm:"column:androiddeviceType" json:"androiddeviceType"`
	IosdeviceToken     string `gorm:"column:iosdeviceToken" json:"iosdeviceToken"`
	IosdeviceType      string `gorm:"column:iosdeviceType" json:"iosdeviceType"`
	First              int32  `gorm:"column:first" json:"first"`
	// SendAt 模板消息發送時間
	SendAt int32 `gorm:"column:send_at" json:"send_at"`
	// IsNewStudent 是否新生； 0未认证1已认证2待认证3未通过
	IsNewStudent uint32 `gorm:"column:is_new_student" json:"is_new_student"`
	// NewStudentImage 新生认证图片
	NewStudentImage string `gorm:"column:new_student_image" json:"new_student_image"`
	// NewStudentUpdatedAt 新生验证更新时间
	NewStudentUpdatedAt string `gorm:"column:new_student_updated_at" json:"new_student_updated_at"`
	// LastLogin app最后登录时间
	LastLogin      int32  `gorm:"column:last_login" json:"last_login"`
	Phone          string `gorm:"column:phone" json:"phone"`
	WechatIcon     string `gorm:"column:wechat_icon" json:"wechat_icon"`
	WechatNickname string `gorm:"column:wechat_nickname" json:"wechat_nickname"`
	// JobSubscribeType 1应届2实习4全职6宣讲会
	JobSubscribeType string `gorm:"column:job_subscribe_type" json:"job_subscribe_type"`
	// JobSubscribeArea 1香港2深圳3广州4上海5北京
	JobSubscribeArea string `gorm:"column:job_subscribe_area" json:"job_subscribe_area"`
	// LikeSubscriber 是否接收点赞通知 1 ：接收 ； 0 ： 不接收
	LikeSubscriber uint8 `gorm:"column:like_subscriber" json:"like_subscriber"`
	// EmploymentStatus 1在读2待业3在职
	EmploymentStatus int8 `gorm:"column:employment_status" json:"employment_status"`
	// JobSubscribeES 订阅工作的就业状态类型
	JobSubscribeES string `gorm:"column:job_subscribe_e_s" json:"job_subscribe_e_s"`
	// Realname 真实姓名
	Realname string `gorm:"column:realname" json:"realname"`
	// OfferRejectReason 室友验证不通过原因
	OfferRejectReason string `gorm:"column:offer_reject_reason" json:"offer_reject_reason"`
	OfferApplyTime    int32  `gorm:"column:offer_apply_time" json:"offer_apply_time"`
	// RoommateAccess 用户进入找室友功能的时间
	RoommateAccess int32 `gorm:"column:roommate_access" json:"roommate_access"`
	// Wechat offer验证添加微信号
	Wechat       string `gorm:"column:wechat" json:"wechat"`
	FreightShare int32  `gorm:"column:freight_share" json:"freight_share"`
}

//用户信息表
func (*User) TableName() string {
	return "register_info"
}

//获取用户信息(多个)
func (*User) GetUser(user *[]User, where *User, sel []string, option *map[string]interface{}) (string, bool) {
	getMany := connect.Getdb().
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

	//Or
	if data, ok := (*option)["or"]; ok {
		for k, v := range data.(map[string]interface{}) {
			if v == nil {
				getMany = getMany.Or(k)
			} else {
				getMany = getMany.Or(k, v)
			}
		}
	}

	getMany.Find(user)
	if getMany.Error != nil {
		log.LogInfo.Error(getMany.Error.Error())
		return "查询失败", false
	}
	return "查询成功", true
}
