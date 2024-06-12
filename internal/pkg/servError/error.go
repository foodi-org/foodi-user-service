package servError

import "fmt"

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
	e.Code = COMMON
	e.Message = msg[UNKONW]
}

func (e *GRPCError) UnKnowErr() {
	e.Code = UNKONW
	e.Message = msg[UNKONW]
}

func GRPCErr() error {
	return &GRPCError{}
}

func NewGRPCError(code int, msg string) error {
	return &GRPCError{code, msg}
}
