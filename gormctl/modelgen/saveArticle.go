package modelgen

// SaveArticleInfo
// @Description: 用户收藏文章信息表
type SaveArticleInfo struct {
	ID int64 `gorm:"primarykey"`

	Uid int64 `gorm:"index:idx_uid"`

	ArticleID int64 `gorm:"index:idx_article_id"`
}
