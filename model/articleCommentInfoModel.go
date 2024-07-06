package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ ArticleCommentInfoModel = (*customArticleCommentInfoModel)(nil)

type (
	// ArticleCommentInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customArticleCommentInfoModel.
	ArticleCommentInfoModel interface {
		articleCommentInfoModel
		withSession(session sqlx.Session) ArticleCommentInfoModel
	}

	customArticleCommentInfoModel struct {
		*defaultArticleCommentInfoModel
	}
)

// NewArticleCommentInfoModel returns a model for the database table.
func NewArticleCommentInfoModel(conn sqlx.SqlConn) ArticleCommentInfoModel {
	return &customArticleCommentInfoModel{
		defaultArticleCommentInfoModel: newArticleCommentInfoModel(conn),
	}
}

func (m *customArticleCommentInfoModel) withSession(session sqlx.Session) ArticleCommentInfoModel {
	return NewArticleCommentInfoModel(sqlx.NewSqlConnFromSession(session))
}
