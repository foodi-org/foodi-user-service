package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UserInfoModel = (*customUserInfoModel)(nil)

type (
	// UserInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserInfoModel.
	UserInfoModel interface {
		userInfoModel
		withSession(session sqlx.Session) UserInfoModel
	}

	customUserInfoModel struct {
		*defaultUserInfoModel
	}
)

// NewUserInfoModel returns a model for the database table.
func NewUserInfoModel(conn sqlx.SqlConn) UserInfoModel {
	return &customUserInfoModel{
		defaultUserInfoModel: newUserInfoModel(conn),
	}
}

func (m *customUserInfoModel) withSession(session sqlx.Session) UserInfoModel {
	return NewUserInfoModel(sqlx.NewSqlConnFromSession(session))
}
