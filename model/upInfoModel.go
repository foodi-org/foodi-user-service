package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UpInfoModel = (*customUpInfoModel)(nil)

type (
	// UpInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUpInfoModel.
	UpInfoModel interface {
		upInfoModel
		withSession(session sqlx.Session) UpInfoModel
	}

	customUpInfoModel struct {
		*defaultUpInfoModel
	}
)

// NewUpInfoModel returns a model for the database table.
func NewUpInfoModel(conn sqlx.SqlConn) UpInfoModel {
	return &customUpInfoModel{
		defaultUpInfoModel: newUpInfoModel(conn),
	}
}

func (m *customUpInfoModel) withSession(session sqlx.Session) UpInfoModel {
	return NewUpInfoModel(sqlx.NewSqlConnFromSession(session))
}
