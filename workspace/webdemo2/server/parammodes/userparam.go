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

// swagger:response operateRev
type OperateRevRsp struct {

	// success operate data
	//
	// in: body
	Body OperateResult
}

/*
/// RevValueBase是嵌入的类型用 swagger all
/// RevValueBase 必须定义为 swagger model swagger 才能生成引用的model
*/

// Operate status
// swagger:model
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
