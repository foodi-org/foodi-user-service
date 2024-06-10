package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ ArticleReadInfoModel = (*customArticleReadInfoModel)(nil)

type (
	// ArticleReadInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customArticleReadInfoModel.
	ArticleReadInfoModel interface {
		articleReadInfoModel
		withSession(session sqlx.Session) ArticleReadInfoModel
	}

	customArticleReadInfoModel struct {
		*defaultArticleReadInfoModel
	}
)

// NewArticleReadInfoModel returns a model for the database table.
func NewArticleReadInfoModel(conn sqlx.SqlConn) ArticleReadInfoModel {
	return &customArticleReadInfoModel{
		defaultArticleReadInfoModel: newArticleReadInfoModel(conn),
	}
}

func (m *customArticleReadInfoModel) withSession(session sqlx.Session) ArticleReadInfoModel {
	return NewArticleReadInfoModel(sqlx.NewSqlConnFromSession(session))
}
