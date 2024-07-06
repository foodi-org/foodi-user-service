package modelgen

// UpInfo
// @Description: 用户点赞信息表
type UpInfo struct {
	ID int64 `gorm:"primarykey"`

	Uid int64 `gorm:"index:idx_uid"`

	ArticleID int64 `gorm:"index:idx_article_id"`

	CommentID int64 `gorm:"index:idx_comment_id"`
}
