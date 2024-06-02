package modelgen

import (
	"database/sql"
	"github.com/foodi-org/foodi-user-service/gormctl/enumgen"
	"gorm.io/gorm"
	"time"
)

// UserInfo
// @Description: 用户信息表
type UserInfo struct {
	ID int64 `gorm:"primarykey"`

	CreatedAt time.Time

	UpdatedAt time.Time

	DeletedAt gorm.DeletedAt `gorm:"-,index:idx_deleted_at"`

	Uid int64 `gorm:"column:uid;index:idx_uid;comment:用户id"`

	Aid int64 `gorm:"column:aid;index:idx_uid"`

	Name string `gorm:"size:32;comment:用户姓名"`

	NikeName string `gorm:"type:varchar(100);comment:昵称"`

	Image string `gorm:"type:varchar(255);comment:用户头像"`

	CardType enumgen.CardType `gorm:"default:1;comment:'证件类型'"`

	CardID sql.NullString `gorm:"column:card_id;size:64;comment:证件号码"`

	Gender enumgen.GenderType `gorm:"type:enum('male','female');default:null;comment:性别"`

	Age sql.NullInt16 `gorm:"comment:用户年龄"`

	Birthday sql.NullString `gorm:"size:32;comment:用户生日"`

	Region sql.NullString `gorm:"size:32;comment:用户所在地区"`

	WeChatID sql.NullString `gorm:"size:256;comment:关联微信id"`

	LV int8 `gorm:"column:lv;comment:用户平台等级"`

	VIP sql.NullInt16 `gorm:"column:vip;comment:用户会员等级"`
}
