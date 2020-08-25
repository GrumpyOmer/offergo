package lib

//statistics.go's
type GetUserStatisticalStruct struct {
	CurrentUser      int
	CurrentMonthUser int
	LastMonthUser    int
	Percentage       float64
	Text             string
}
type GetUserStatisticalResponseData struct {
	WxUser                       GetUserStatisticalStruct //微信用户信息
	AppUser                      GetUserStatisticalStruct //总APP用户信息
	IosUser                      GetUserStatisticalStruct //ios用户信息
	AndroidUser                  GetUserStatisticalStruct //android用户信息
	JobBookUser                  GetUserStatisticalStruct //工作板块用户信息
	ActiveBookUser               GetUserStatisticalStruct //活动板块用户信息
	MbUser                       GetUserStatisticalStruct //集运信息
	SecondHandUser               GetUserStatisticalStruct //二手用户信息
	OfflineCommissionPublishUser GetUserStatisticalStruct //待寄待取发布者用户信息
	OfflineCommissionPayUser     GetUserStatisticalStruct //待寄待取付费者用户信息
	TelecomCardActivationedUser  GetUserStatisticalStruct //大K卡已激活总人数信息
	TelecomUsingUser             GetUserStatisticalStruct //正在使用的大k卡用户信息
	TelecomNewUser               GetUserStatisticalStruct //新申请的大k卡用户信息
	KPlusUser                    GetUserStatisticalStruct //k_plus会员用户信息
	BlankCardActivationedInfo    GetUserStatisticalStruct //白卡激活用户信息
	BlankCardUsingInfo           GetUserStatisticalStruct //白卡开卡用户信息
}

//page's struct
type PageStruct struct {
	From     int     `json:"from"`      //当前页
	LastPage float64 `json:"last_page"` //最后页
	PerPage  int     `json:"per_page"`  //每页条数
	Total    float64 `json:"total"`     //总条数
}

type PageResult struct {
	Data interface{}
	Page PageStruct
}

//shenZhouInvite's api struct
type ShenZhouInviteApiResult struct {
	Code int
	Msg  string
	Data InviteStruct
}

type InviteStruct struct {
	HongKong  []HongKongStruct  `json:"hong_kong"`
	TakePoint []TakePointStruct `json:"take_point"`
}

type HongKongStruct struct {
	MethodId          int    `json:"method_id"`
	MethodName        string `json:"method_name"`
	MethodDescription string `json:"method_description"`
	FirstKgFee        int    `json:"first_kg_fee"`
	SecondKgFee       int    `json:"second_kg_fee"`
	AdditionalFee     int    `json:"additional_fee"`
	MethodType        int    `json:"method_type"`
}

type TakePointStruct struct {
	MethodId          int    `json:"method_id"`
	MethodName        string `json:"method_name"`
	MethodDescription string `json:"method_description"`
	FirstKgFee        int    `json:"first_kg_fee"`
	SecondKgFee       int    `json:"second_kg_fee"`
	AdditionalFee     int    `json:"additional_fee"`
	TakePointAddress  string `json:"take_point_address"`
	TakePointStoreDay int    `json:"take_point_store_day"`
	TakePointArea     int    `json:"take_point_area"`
	TakePointLocation int    `json:"take_point_location"`
	MethodType        int    `json:"method_type"`
}

type SearchDocumentStruct struct {
	Html string
	Img  []string
}

//telecom.go's
