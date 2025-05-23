package main

import (
	"flag"
	"fmt"

	"gozero_demo1/model_study/user_gorm/api/internal/config"
	"gozero_demo1/model_study/user_gorm/api/internal/handler"
	"gozero_demo1/model_study/user_gorm/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "./model_study/user_gorm/api/etc/users.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
