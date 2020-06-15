package lib

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
