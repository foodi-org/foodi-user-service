// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/foodi-org/foodi-user-service/model/modelType/bo"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	saveArticleInfoFieldNames          = builder.RawFieldNames(&SaveArticleInfo{})
	saveArticleInfoRows                = strings.Join(saveArticleInfoFieldNames, ",")
	saveArticleInfoRowsExpectAutoSet   = strings.Join(stringx.Remove(saveArticleInfoFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	saveArticleInfoRowsWithPlaceHolder = strings.Join(stringx.Remove(saveArticleInfoFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	saveArticleInfoModel interface {
		Insert(ctx context.Context, data *SaveArticleInfo) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*SaveArticleInfo, error)
		Update(ctx context.Context, data *SaveArticleInfo) error
		Delete(ctx context.Context, id int64) error

		/*DelSaveArticle
		@Description: 移除文章收藏
		*/
		DelSaveArticle(ctx context.Context, bo bo.ArticleBo) error
	}

	defaultSaveArticleInfoModel struct {
		conn  sqlx.SqlConn
		table string
	}

	SaveArticleInfo struct {
		Id        int64         `db:"id"`
		Uid       sql.NullInt64 `db:"uid"`
		ArticleId sql.NullInt64 `db:"article_id"`
	}
)

func newSaveArticleInfoModel(conn sqlx.SqlConn) *defaultSaveArticleInfoModel {
	return &defaultSaveArticleInfoModel{
		conn:  conn,
		table: "`save_article_info`",
	}
}

func (m *defaultSaveArticleInfoModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultSaveArticleInfoModel) FindOne(ctx context.Context, id int64) (*SaveArticleInfo, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", saveArticleInfoRows, m.table)
	var resp SaveArticleInfo
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultSaveArticleInfoModel) Insert(ctx context.Context, data *SaveArticleInfo) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, saveArticleInfoRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Uid, data.ArticleId)
	return ret, err
}

func (m *defaultSaveArticleInfoModel) Update(ctx context.Context, data *SaveArticleInfo) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, saveArticleInfoRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Uid, data.ArticleId, data.Id)
	return err
}

// DelSaveArticle
//
//	@Description: 移除文章收藏
//	@param ctx
//	@param bo
//	@return error
func (m *defaultSaveArticleInfoModel) DelSaveArticle(ctx context.Context, bo bo.ArticleBo) error {
	query := fmt.Sprintf("delete from %s where `uid` = ? and `article_id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, bo.Uid, bo.ArticleID)
	return err
}

func (m *defaultSaveArticleInfoModel) tableName() string {
	return m.table
}
