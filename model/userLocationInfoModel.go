package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UserLocationInfoModel = (*customUserLocationInfoModel)(nil)

type (
	// UserLocationInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserLocationInfoModel.
	UserLocationInfoModel interface {
		userLocationInfoModel
		withSession(session sqlx.Session) UserLocationInfoModel
	}

	customUserLocationInfoModel struct {
		*defaultUserLocationInfoModel
	}
)

// NewUserLocationInfoModel returns a model for the database table.
func NewUserLocationInfoModel(conn sqlx.SqlConn) UserLocationInfoModel {
	return &customUserLocationInfoModel{
		defaultUserLocationInfoModel: newUserLocationInfoModel(conn),
	}
}

func (m *customUserLocationInfoModel) withSession(session sqlx.Session) UserLocationInfoModel {
	return NewUserLocationInfoModel(sqlx.NewSqlConnFromSession(session))
}
