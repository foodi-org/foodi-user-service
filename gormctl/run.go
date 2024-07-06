package gormctl

import (
	foodipkg "github.com/foodi-org/foodi-pkg/mysql"
	"github.com/foodi-org/foodi-user-service/gormctl/modelgen"
)

func RunGormCTL() {
	foodipkg.InitConn("root:123456@tcp(8.134.206.241:3306)/foodi_user_service?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai")
	db := foodipkg.GetDBConn()
	_ = db.AutoMigrate(&modelgen.UserInfo{})
	_ = db.AutoMigrate(&modelgen.AccountInfo{})
	_ = db.AutoMigrate(&modelgen.ArticleInfo{})
	_ = db.AutoMigrate(&modelgen.ArticleCommentInfo{})
	_ = db.AutoMigrate(&modelgen.UserLocationInfo{})
	_ = db.AutoMigrate(&modelgen.UserLoginInfo{})
	_ = db.AutoMigrate(&modelgen.UserWechatInfo{})
	_ = db.AutoMigrate(&modelgen.UpInfo{})
	_ = db.AutoMigrate(&modelgen.SaveArticleInfo{})
	_ = db.AutoMigrate(&modelgen.ArticleReadInfo{})
}
