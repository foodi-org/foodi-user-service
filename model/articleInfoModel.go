package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ ArticleInfoModel = (*customArticleInfoModel)(nil)

type (
	// ArticleInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customArticleInfoModel.
	ArticleInfoModel interface {
		articleInfoModel
		withSession(session sqlx.Session) ArticleInfoModel
	}

	customArticleInfoModel struct {
		*defaultArticleInfoModel
	}
)

// NewArticleInfoModel returns a model for the database table.
func NewArticleInfoModel(conn sqlx.SqlConn) ArticleInfoModel {
	return &customArticleInfoModel{
		defaultArticleInfoModel: newArticleInfoModel(conn),
	}
}

func (m *customArticleInfoModel) withSession(session sqlx.Session) ArticleInfoModel {
	return NewArticleInfoModel(sqlx.NewSqlConnFromSession(session))
}
