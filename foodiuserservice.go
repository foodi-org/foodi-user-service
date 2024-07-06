package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/foodi-org/foodi-user-service/github.com/foodi-org/foodi-user-service"
	"github.com/foodi-org/foodi-user-service/internal/config"
	accountServer "github.com/foodi-org/foodi-user-service/internal/server/account"
	articleServer "github.com/foodi-org/foodi-user-service/internal/server/article"
	userServer "github.com/foodi-org/foodi-user-service/internal/server/user"
	"github.com/foodi-org/foodi-user-service/internal/svc"

	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	zrpcconsul "github.com/zeromicro/zero-contrib/zrpc/registry/consul"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "dev.yaml", "the config file")

func main() {
	var c = config.ServConf()
	dir, _ := os.Getwd()
	flag.Parse()

	if err := svc.NewServiceContext(c, dir, *configFile); err != nil {
		panic(err)
	}

	// 注册服务
	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {

		// 注册账号服务
		foodi_user_service.RegisterAccountServer(grpcServer, accountServer.NewAccountServer(svc.Svc()))

		// 注册用户服务
		foodi_user_service.RegisterUserServer(grpcServer, userServer.NewUserServer(svc.Svc()))

		// 注册文章服务
		foodi_user_service.RegisterArticleServer(grpcServer, articleServer.NewArticleServer(svc.Svc()))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})

	// 注册中心使用 consul
	if err := zrpcconsul.RegisterService(c.ListenOn, c.Consul); err != nil {
		panic(err)
	}

	defer s.Stop()

	fmt.Printf("Starting foodi-user rpc server at %s...\n", c.ListenOn)
	s.Start()
}
