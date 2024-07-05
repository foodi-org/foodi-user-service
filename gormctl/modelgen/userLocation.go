package modelgen

// UserLocationInfo
// @Description: 用户收件地址信息表
type UserLocationInfo struct {
	ID int64 `gorm:"primarykey"`

	Uid int64 `gorm:"index"`

	LocationLabel string `gorm:"type:varchar(255)"`

	Tag string `gorm:"type:varchar(255)"`

	Description string `gorm:"type:varchar(255)"`

	Receiver string `gorm:"type:varchar(128)"`

	Phone string `gorm:"type:varchar(255)"`

	Longitude float64 `gorm:"comment:经度"`

	Latitude float64 `gorm:"comment:纬度"`
}
