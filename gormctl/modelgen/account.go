package modelgen

import (
	"database/sql"
	"github.com/foodi-org/foodi-user-service/gormctl/enumgen"
	"gorm.io/gorm"
	"time"
)

// AccountInfo
// @Description: 账号信息表
type AccountInfo struct {
	ID int64 `gorm:"primarykey"`

	CreatedAt time.Time

	UpdatedAt time.Time

	DeletedAt gorm.DeletedAt `gorm:"-,index:idx_deleted_at"`

	Type enumgen.UserType `gorm:"type:enum('consumer','merchant');comment:账号类型"`

	NikeName string `gorm:"type:varchar(100);comment:昵称"`

	Image string `gorm:"comment:用户头像"`

	Phone sql.NullInt16 `gorm:"index:idx_phone"`

	Password sql.NullString `gorm:"type:varchar(64)"`

	FirstRegister sql.NullString `gorm:"type:varchar(32);comment:首次注册时间"`

	LastLogin sql.NullString `gorm:"type:varchar(32);comment:上次登录时间"`

	Verified sql.NullBool `gorm:"comment:是否实名认证"`

	BindWX sql.NullBool `gorm:"comment:是否绑定了微信"`
}
