package modelgen

import (
	"database/sql"
	"github.com/foodi-org/foodi-user-service/gormctl/enumgen"
	"gorm.io/gorm"
	"time"
)

// ArticleInfo
// @Description: 用户发布的文章、笔记
type ArticleInfo struct {
	ID int64 `gorm:"primarykey"`

	CreatedAt time.Time

	UpdatedAt time.Time

	DeletedAt gorm.DeletedAt `gorm:"-,index:idx_deleted_at"`

	Uid int64 `gorm:"not null;index:idx_note_uid"`

	Published sql.NullBool `gorm:"comment:是否发布"`

	PublishedAt sql.NullTime `gorm:"comment:发布时间"`

	ArticleType enumgen.ArticleType `gorm:"type:enum('video','text');default:text;comment:文章类型"`

	Title sql.NullString `gorm:"type:varchar(255)"`

	Video sql.NullString `gorm:"type:varchar(255)"`

	ImageURL sql.NullString `gorm:"type:varchar(255);column:image_url;comment:图片地址"`

	Content sql.NullString `gorm:"type:text"`

	Up int `gorm:"default:0;comment:点赞数"`

	Save int `gorm:"default:0;comment:收藏数"`

	Region sql.NullString `gorm:"type:varchar(255);comment:定位地区"`

	Location sql.NullString `gorm:"type:varchar(255);comment:精准定位信息"`

	Merchant sql.NullInt64 `gorm:"comment:关联的商家"`

	Cover sql.NullString `gorm:"type:varchar(255);comment:封面url"`
}
