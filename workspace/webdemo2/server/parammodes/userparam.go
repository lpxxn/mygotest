package parammodes

//
// swagger:parameters upUserParam
type UserParam struct {
	// IDD
	//
	// required: true
	// in: body
	Param UserParamInfo
}

// swagger:model param
type UserParamInfo struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Age  int32  `json:"age"`
}
