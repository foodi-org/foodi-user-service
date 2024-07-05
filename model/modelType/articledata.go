package modelType

import (
	"database/sql"
)

type (
	DraftArticle struct {
		ID        int64
		CreatedAt sql.NullTime // 创建时间
		UpdatedAt sql.NullTime // 最近一次的更新时间
		Title     string
	}
)
