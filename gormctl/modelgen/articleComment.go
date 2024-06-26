package modelgen

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

// ArticleCommentInfo
// @Description: 评论表
type ArticleCommentInfo struct {
	ID int64 `gorm:"primarykey"`

	CreatedAt time.Time

	UpdatedAt time.Time

	DeletedAt gorm.DeletedAt `gorm:"-,index:idx_deleted_at"`

	ArticleId int64 `gorm:"index:idx_article_id"`

	Uid int64 `gorm:"index:idx_uid;comment:发布评论用户id"`

	Content string `gorm:"type:text;comment:评论内容"`

	PostTime time.Time `gorm:"type:datetime;comment:发布时间"`

	//
	//  ParentId
	//  @Description: parent_id字段的值为-1时表示该评论为顶级评论，即没有父评论。当parent_id字段的值不为-1时，则表示该评论是对parent_id所对应的评论的回复。
	//
	ParentId sql.NullInt64 `gorm:"comment:父评论id"`

	//
	//  FirstCommentId
	//  @Description: 回复链的第一个回复id
	//
	FirstCommentId sql.NullInt64 `gorm:"comment:回复的起始评论"`

	UpCount sql.NullInt16 `gorm:"comment:点赞数"`

	ReplyCount sql.NullInt16 `gorm:"comment:回复此评论的数量"`

	IsAnonymous bool `gorm:"default:false;comment:是否匿名"`
}
