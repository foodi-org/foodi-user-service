package servError

import "fmt"

type (
	GRPCError struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}
)

// Error 接口实现
func (e *GRPCError) Error() string {
	return fmt.Sprintf("grpc code: %d, message: %s", e.Code, e.Message)
}

func (e *GRPCError) CommonErr() {
	e.Code = Common
	e.Message = msg[UNKNOWN]
}

func (e *GRPCError) UnKnowErr() {
	e.Code = UNKNOWN
	e.Message = msg[UNKNOWN]
}

func GRPCErr() error {
	return &GRPCError{}
}

func NewGRPCError(code int, msg string) error {
	return &GRPCError{code, msg}
}
