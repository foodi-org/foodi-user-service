package modelgen

import "time"

// ArticleReadInfo
// @Description: 文章阅读扩展记录
type ArticleReadInfo struct {
	ID int64 `gorm:"primarykey"`

	CreatedAt time.Time

	Uid int64 `gorm:"index: idx_uid"`

	Aid int64 `gorm:"index: idx_aid"`
}
