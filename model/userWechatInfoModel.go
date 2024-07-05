package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UserWechatInfoModel = (*customUserWechatInfoModel)(nil)

type (
	// UserWechatInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserWechatInfoModel.
	UserWechatInfoModel interface {
		userWechatInfoModel
		withSession(session sqlx.Session) UserWechatInfoModel
	}

	customUserWechatInfoModel struct {
		*defaultUserWechatInfoModel
	}
)

// NewUserWechatInfoModel returns a model for the database table.
func NewUserWechatInfoModel(conn sqlx.SqlConn) UserWechatInfoModel {
	return &customUserWechatInfoModel{
		defaultUserWechatInfoModel: newUserWechatInfoModel(conn),
	}
}

func (m *customUserWechatInfoModel) withSession(session sqlx.Session) UserWechatInfoModel {
	return NewUserWechatInfoModel(sqlx.NewSqlConnFromSession(session))
}
