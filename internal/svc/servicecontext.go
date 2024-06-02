package svc

import (
	foodipkg "github.com/foodi-org/foodi-pkg/mysql"
	"github.com/foodi-org/foodi-user-service/internal/config"
	"github.com/foodi-org/foodi-user-service/model"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"gorm.io/gorm"
	"time"
)

var svc ServiceContext

type ServiceContext struct {
	Config config.Config
	Redis  *redis.Redis
	conn   sqlx.SqlConn
	orm    *gorm.DB

	AccountModel        model.AccountInfoModel
	ArticleCommentModel model.ArticleCommentInfoModel
	SaveArticleModel    model.SaveArticleInfoModel
	UpModel             model.UpInfoModel
	UserInfoModel       model.UserInfoModel
	UserLocationModel   model.UserLocationInfoModel
	UserLoginModel      model.UserLoginInfoModel
	ArticleModel        model.ArticleInfoModel
	UserWechatModel     model.UserWechatInfoModel
}

// NewServiceContext
//
//	@Description: 加载 svc
//	@param c
//	@param dir 项目路径
//	@param file 配置文件名称
//	@return *ServiceContext
func NewServiceContext(c *config.Config, dir string, file string) error {
	// 加载配置
	if err := config.InitServConf(dir, file); err != nil {
		return err
	}

	// mysql client
	foodipkg.InitConn(c.Mysql.DataSource)

	// gorm mysql conn
	svc.orm = foodipkg.GetDBConn()

	// sqlx mysql
	svc.conn = sqlx.NewMysql(c.Mysql.DataSource)

	// redis client
	red := redis.MustNewRedis(redis.RedisConf{
		Host:        c.Redis.Host,
		Type:        c.Redis.Type,
		Pass:        c.Redis.Pass,
		Tls:         c.Redis.Tls,
		PingTimeout: 3 * time.Second,
	})
	svc = ServiceContext{
		Config: *c,
		Redis:  red,
	}

	svc.AccountModel = model.NewAccountInfoModel(svc.conn)
	svc.ArticleCommentModel = model.NewArticleCommentInfoModel(svc.conn)
	svc.SaveArticleModel = model.NewSaveArticleInfoModel(svc.conn)
	svc.UpModel = model.NewUpInfoModel(svc.conn)
	svc.UserInfoModel = model.NewUserInfoModel(svc.conn)
	svc.UserLocationModel = model.NewUserLocationInfoModel(svc.conn)
	svc.UserLoginModel = model.NewUserLoginInfoModel(svc.conn)
	svc.ArticleModel = model.NewArticleInfoModel(svc.conn)
	svc.UserWechatModel = model.NewUserWechatInfoModel(svc.conn)

	return nil
}

func RedisCli() *redis.Redis {
	return svc.Redis
}

func Svc() *ServiceContext {
	return &svc
}
