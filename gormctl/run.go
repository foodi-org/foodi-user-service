package main

import (
	foodipkg "github.com/foodi-org/foodi-pkg/mysql"
	"github.com/foodi-org/foodi-user-service/gormctl/modelgen"
)

func main() {
	foodipkg.InitConn("root:123456@tcp(8.134.206.241:3306)/foodi_lbs_service?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai")
	db := foodipkg.GetDBConn()
	db.AutoMigrate(&modelgen.UserInfo{})
}
