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

// swagger:response successOperate
type SuccessOperateRsp struct {

	// success operate data
	//
	// in: body
	Body OperateResult
}

type RevValueBase struct {
	Status bool   `json:"state"`
	Err    string `json:"err"`
}

type RevError struct {
	ErrorCode string `json:"errorCode"`
	ErrorDesc string `json:"errorDesc"`
}

type OperateResult struct {
	// swagger:allOf
	RevValueBase
	Error RevError `json:"error"`
}
