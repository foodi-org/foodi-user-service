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

	Content string `gorm:"type:text"`

	PostTime time.Time `gorm:"type:datetime"`

	//
	//  ParentId
	//  @Description: parent_id字段的值为-1时表示该评论为顶级评论，即没有父评论。当parent_id字段的值不为-1时，则表示该评论是对parent_id所对应的评论的回复。
	//
	ParentId sql.NullInt64 `gorm:"comment:父评论id"`

	UpCount sql.NullInt16 `gorm:"comment:点赞数"`

	ReplyCount sql.NullInt16 `gorm:"comment:回复此评论的数量"`

	IsAnonymous bool `gorm:"default:false;comment:是否匿名"`
}
