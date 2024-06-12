package servError

// 用户中心 - 文章分层服务 error code
// 码段区间 113000 ~ 113999

func (e *GRPCError) ParamInvalid() {
	e.Code = InvalidParam
	e.Message = msg[InvalidParam]
}

func (e *GRPCError) ActionInvalid() {
	e.Code = InvalidAction
	e.Message = msg[InvalidAction]
}
