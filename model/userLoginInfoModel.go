package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UserLoginInfoModel = (*customUserLoginInfoModel)(nil)

type (
	// UserLoginInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserLoginInfoModel.
	UserLoginInfoModel interface {
		userLoginInfoModel
		withSession(session sqlx.Session) UserLoginInfoModel
	}

	customUserLoginInfoModel struct {
		*defaultUserLoginInfoModel
	}
)

// NewUserLoginInfoModel returns a model for the database table.
func NewUserLoginInfoModel(conn sqlx.SqlConn) UserLoginInfoModel {
	return &customUserLoginInfoModel{
		defaultUserLoginInfoModel: newUserLoginInfoModel(conn),
	}
}

func (m *customUserLoginInfoModel) withSession(session sqlx.Session) UserLoginInfoModel {
	return NewUserLoginInfoModel(sqlx.NewSqlConnFromSession(session))
}
