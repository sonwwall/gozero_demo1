package main

import (
	"flag"
	"fmt"

	"gozero_demo1/rpc_study/user_group/rpc/internal/config"
	useractionServer "gozero_demo1/rpc_study/user_group/rpc/internal/server/useraction"
	userinfoServer "gozero_demo1/rpc_study/user_group/rpc/internal/server/userinfo"
	"gozero_demo1/rpc_study/user_group/rpc/internal/svc"
	"gozero_demo1/rpc_study/user_group/rpc/types/user"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "rpc_study/user_group/rpc/etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		user.RegisterUserInfoServer(grpcServer, userinfoServer.NewUserInfoServer(ctx))
		user.RegisterUserActionServer(grpcServer, useractionServer.NewUserActionServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
