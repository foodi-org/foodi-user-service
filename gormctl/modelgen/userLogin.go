package modelgen

import (
	"database/sql"
	"time"
)

// UserLoginInfo
// @Description: 预留，统计用户登录时间和下线时间等埋点信息，用于月活等统计
type UserLoginInfo struct {
	ID int64 `gorm:"primarykey"`

	Aid int64 `gorm:"comment:账号id"`

	LastLogin time.Time `gorm:"上次登录时间"`

	// Count 登录次数
	Count int

	// BuyActionCount 购买次数
	BuyActionCount sql.NullInt64

	// NoteActionCount 发布笔记次数
	NoteActionCount sql.NullInt64

	// ActivityDuration 每次活跃时长
	ActivityDuration sql.NullInt16

	// ConsecutiveActiveDays 连续活跃天数
	ConsecutiveActiveDays int8
}
