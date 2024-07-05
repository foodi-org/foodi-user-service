package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ AccountInfoModel = (*customAccountInfoModel)(nil)

type (
	// AccountInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAccountInfoModel.
	AccountInfoModel interface {
		accountInfoModel
		withSession(session sqlx.Session) AccountInfoModel
	}

	customAccountInfoModel struct {
		*defaultAccountInfoModel
	}
)

// NewAccountInfoModel returns a model for the database table.
func NewAccountInfoModel(conn sqlx.SqlConn) AccountInfoModel {
	return &customAccountInfoModel{
		defaultAccountInfoModel: newAccountInfoModel(conn),
	}
}

func (m *customAccountInfoModel) withSession(session sqlx.Session) AccountInfoModel {
	return NewAccountInfoModel(sqlx.NewSqlConnFromSession(session))
}
