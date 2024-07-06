package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ SaveArticleInfoModel = (*customSaveArticleInfoModel)(nil)

type (
	// SaveArticleInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSaveArticleInfoModel.
	SaveArticleInfoModel interface {
		saveArticleInfoModel
		withSession(session sqlx.Session) SaveArticleInfoModel
	}

	customSaveArticleInfoModel struct {
		*defaultSaveArticleInfoModel
	}
)

// NewSaveArticleInfoModel returns a model for the database table.
func NewSaveArticleInfoModel(conn sqlx.SqlConn) SaveArticleInfoModel {
	return &customSaveArticleInfoModel{
		defaultSaveArticleInfoModel: newSaveArticleInfoModel(conn),
	}
}

func (m *customSaveArticleInfoModel) withSession(session sqlx.Session) SaveArticleInfoModel {
	return NewSaveArticleInfoModel(sqlx.NewSqlConnFromSession(session))
}
