package servError

import "fmt"

const (
	COMMONCODE = 111000
	UNKONWCODE = 111099
)

const (
	COMMONMSG = "user service error"
	UNKNOWMSG = "unknow service error"
)

type (
	GRPCError struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}
)

func (e *GRPCError) Error() string {
	return fmt.Sprintf("grpc code: %d, message: %s", e.Code, e.Message)
}

func (e *GRPCError) CommonErr() {
	e.Code = COMMONCODE
	e.Message = COMMONMSG
}

func (e *GRPCError) UnKnowErr() {
	e.Code = UNKONWCODE
	e.Message = UNKNOWMSG
}

func NewGRPCError(code int, msg string) error {
	return &GRPCError{code, msg}
}
