package svc

import (
	"github.com/foodi-org/foodi-user-service/internal/config"
	"github.com/foodi-org/foodi-user-service/model"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"time"
)

var svc = ServiceContext{}

type ServiceContext struct {
	Config *config.Config
	Redis  *redis.Redis

	AccountModel        model.AccountInfoModel
	ArticleCommentModel model.ArticleCommentInfoModel
	SaveArticleModel    model.SaveArticleInfoModel
	UpModel             model.UpInfoModel
	UserInfoModel       model.UserInfoModel
	UserLocationModel   model.UserLocationInfoModel
	UserLoginModel      model.UserLoginInfoModel
	ArticleModel        model.ArticleInfoModel
	UserWechatModel     model.UserWechatInfoModel
	ArticleReadModel    model.ArticleReadInfoModel
}

// NewServiceContext
//
//	@Description: 加载 svc
//	@param c
//	@param path 项目路径
//	@param file 配置文件名称
//	@return *ServiceContext
func NewServiceContext(c *config.Config, path string, file string) error {
	// 加载配置
	if err := config.InitServConf(path, file); err != nil {
		return err
	} else {
		svc.Config = c
	}

	// 日志配置
	logx.MustSetup(logx.LogConf{
		ServiceName: c.Name,
		Mode:        "file",
		//TimeFormat:  "",
		Path:       path,
		Level:      "info",
		KeepDays:   10,
		MaxBackups: 10,
		Rotation:   "daily",
	})
	logx.DisableStat() // disables the stat logs.

	// redis client
	svc.Redis = redis.MustNewRedis(redis.RedisConf{
		Host:        c.Redis.Host,
		Type:        c.Redis.Type,
		Pass:        c.Redis.Pass,
		Tls:         c.Redis.Tls,
		PingTimeout: 3 * time.Second,
	})

	// mysql
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	svc.AccountModel = model.NewAccountInfoModel(conn)
	svc.ArticleCommentModel = model.NewArticleCommentInfoModel(conn)
	svc.SaveArticleModel = model.NewSaveArticleInfoModel(conn)
	svc.UpModel = model.NewUpInfoModel(conn)
	svc.UserInfoModel = model.NewUserInfoModel(conn)
	svc.UserLocationModel = model.NewUserLocationInfoModel(conn)
	svc.UserLoginModel = model.NewUserLoginInfoModel(conn)
	svc.ArticleModel = model.NewArticleInfoModel(conn)
	svc.UserWechatModel = model.NewUserWechatInfoModel(conn)
	svc.ArticleReadModel = model.NewArticleReadInfoModel(conn)

	return nil
}

func RedisCli() *redis.Redis {
	return svc.Redis
}

func Svc() *ServiceContext {
	return &svc
}
