package models

type Articlelist struct {
	// ID 文章列表自增id
	ID int32 `json:"id"`
	// ClickNum 文章的点击量
	ClickNum int32 `json:"click_num"`
	// Title 文章的标题
	Title string `json:"title"`
	// Classify 文章的分类 1:生活2:毕业3:出行
	Classify string `json:"classify"`
	// Digest 文章的摘要
	Digest string `json:"digest"`
	// ThumbMediaID 图文消息封面图片素材id
	ThumbMediaID string `json:"thumb_media_id"`
	// URL 图文地址
	URL string `json:"url"`
	// ThumbURL 图片地址
	ThumbURL string `json:"thumb_url"`
	// ReadMoreURL 阅读原文的链接
	ReadMoreURL string `json:"read_more_url"`
	// IsPush 是否进行文章推送
	IsPush int32 `json:"is_push"`
	// CreateTime 创建时间
	CreateTime int32 `json:"create_time"`
	// UpdateTime 更新时间（时间戳）
	UpdateTime int32 `json:"update_time"`
	// StartTime 文章展示起始时间
	StartTime int32 `json:"start_time"`
	// EndTime 文章展示结束时间
	EndTime int32 `json:"end_time"`
	// Sticky 置顶优先度
	Sticky int32 `json:"sticky"`
	// StickyStartTime 文章置顶起始时间
	StickyStartTime int32 `json:"sticky_start_time"`
	// StickyEndTime 文章置顶结束时间
	StickyEndTime int32  `json:"sticky_end_time"`
	DeletedAt     int32  `json:"deleted_at"`
	ShowToUser    uint32 `json:"show_to_user"`
}

//文章表
func (*Articlelist) TableName() string {
	return "articlelist"
}
