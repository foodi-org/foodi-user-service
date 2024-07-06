package main

import (
	"flag"
	"fmt"

	"github.com/foodi-org/foodi-user-service/github.com/foodi-org/foodi-user-service"
	"github.com/foodi-org/foodi-user-service/internal/config"
	accountServer "github.com/foodi-org/foodi-user-service/internal/server/account"
	"github.com/foodi-org/foodi-user-service/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/foodiaccountservice.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		foodi_user_service.RegisterAccountServer(grpcServer, accountServer.NewAccountServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
