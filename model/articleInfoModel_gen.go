// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/foodi-org/foodi-user-service/model/modelType"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	articleInfoFieldNames          = builder.RawFieldNames(&ArticleInfo{})
	articleInfoRows                = strings.Join(articleInfoFieldNames, ",")
	articleInfoRowsExpectAutoSet   = strings.Join(stringx.Remove(articleInfoFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	articleInfoRowsWithPlaceHolder = strings.Join(stringx.Remove(articleInfoFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	articleInfoModel interface {
		Insert(ctx context.Context, data *ArticleInfo) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*ArticleInfo, error)
		Update(ctx context.Context, data *ArticleInfo) error
		Delete(ctx context.Context, id int64) error

		/*Disable
		@Description: 删除文章
		@param id 文章id
		@param uid 用户id
		@return error
		*/
		Disable(ctx context.Context, id int64, uid int64) error

		/*DraftList
		@Description: 获取用户保存的文稿
		@param uid 用户id
		*/
		DraftList(ctx context.Context, uid int64) ([]*modelType.DraftArticle, error)
	}

	defaultArticleInfoModel struct {
		conn  sqlx.SqlConn
		table string
	}

	ArticleInfo struct {
		Id          int64          `db:"id"`
		CreatedAt   sql.NullTime   `db:"created_at"`
		UpdatedAt   sql.NullTime   `db:"updated_at"`
		DeletedAt   sql.NullTime   `db:"deleted_at"`
		Uid         int64          `db:"uid"`
		Published   sql.NullInt64  `db:"published"`    // 是否发布
		PublishedAt sql.NullTime   `db:"published_at"` // 发布时间
		ArticleType string         `db:"article_type"` // 文章类型
		Title       sql.NullString `db:"title"`
		Video       sql.NullString `db:"video"`
		ImageUrl    sql.NullString `db:"image_url"` // 图片地址
		Content     sql.NullString `db:"content"`
		Up          int64          `db:"up"`       // 点赞数
		Save        int64          `db:"save"`     // 收藏数
		Region      sql.NullString `db:"region"`   // 定位地区
		Location    sql.NullString `db:"location"` // 精准定位信息
		Merchant    sql.NullInt64  `db:"merchant"` // 关联的商家
		Cover       sql.NullString `db:"cover"`    // 封面url
	}
)

func newArticleInfoModel(conn sqlx.SqlConn) *defaultArticleInfoModel {
	return &defaultArticleInfoModel{
		conn:  conn,
		table: "`article_info`",
	}
}

func (m *defaultArticleInfoModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultArticleInfoModel) FindOne(ctx context.Context, id int64) (*ArticleInfo, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", articleInfoRows, m.table)
	var resp ArticleInfo
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

func (m *defaultArticleInfoModel) Insert(ctx context.Context, data *ArticleInfo) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, articleInfoRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.DeletedAt, data.Uid, data.Published, data.PublishedAt, data.ArticleType, data.Title, data.Video, data.ImageUrl, data.Content, data.Up, data.Save, data.Region, data.Location, data.Merchant, data.Cover)
	return ret, err
}

func (m *defaultArticleInfoModel) Update(ctx context.Context, data *ArticleInfo) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, articleInfoRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.DeletedAt, data.Uid, data.Published, data.PublishedAt, data.ArticleType, data.Title, data.Video, data.ImageUrl, data.Content, data.Up, data.Save, data.Region, data.Location, data.Merchant, data.Cover, data.Id)
	return err
}

// Disable
//
//	@Description: 删除文章
//	@param id 文章id
//	@param uid 用户id
//	@return error
func (m *defaultArticleInfoModel) Disable(ctx context.Context, id int64, uid int64) error {
	query := fmt.Sprintf("update %s set `deleted_at` = ? where `id` = ? and `uid` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, time.Now(), id, uid)
	return err
}

// 获取用户保存的文稿
func (m *defaultArticleInfoModel) DraftList(ctx context.Context, uid int64) ([]*modelType.DraftArticle, error) {
	var res []*modelType.DraftArticle
	query := fmt.Sprintf("select `id`, `created_at`, `updated_at`, `title` from %s where `uid` = ? and `published_at` is null and `deleted_at` is null", m.table)
	if err := m.conn.QueryRowsCtx(ctx, &res, query, uid); err != nil {
		return nil, err
	}
	return res, nil
}

func (m *defaultArticleInfoModel) tableName() string {
	return m.table
}
