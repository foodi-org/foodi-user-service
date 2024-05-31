package modelgen

import (
	"database/sql"
	"github.com/foodi-org/foodi-user-service/gormctl/enumgen"
	"gorm.io/gorm"
	"time"
)

type UserInfo struct {
	ID        int64 `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"-,index:idx_deleted_at"`

	Uid int64

	Name string

	CardType enumgen.CardType `gorm:"default:1;comment:'证件类型'"`

	CardID sql.NullString `gorm:"column:card_id;size:64"`

	Gender enumgen.GenderType `gorm:"default:null"`
}
