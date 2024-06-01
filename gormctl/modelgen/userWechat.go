package modelgen

import "gorm.io/gorm"

// UserWechatInfo
// @Description: 存储用户绑定的微信信息
type UserWechatInfo struct {
	ID int64 `gorm:"primarykey"`

	DeletedAt gorm.DeletedAt `gorm:"-,index:idx_deleted_at"`

	Uid int64 `gorm:"index:idx_uid"`

	Aid int64 `gorm:"index:idx_aid"`
}
