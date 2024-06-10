package servError

// 用户中心 - 文章分层服务 error code
// 码段区间 113000 ~ 113999

const (
	InvalidParamCode = 113000 // 无效的参数
	ActionErrCode    = 113001 // 操作类型错误
)

const (
	InvalidParamMsg = "invalid parameter"
	ActionErrMsg    = "action type valid"
)

func (e *GRPCError) ParamInvalid() {
	e.Code = InvalidParamCode
	e.Message = InvalidParamMsg
}

func (e *GRPCError) ActionInvalid() {
	e.Code = ActionErrCode
	e.Message = ActionErrMsg
}
