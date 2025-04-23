package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"gozero_demo1/rpc_study/user_api_rpc/api/internal/config"
	"gozero_demo1/rpc_study/user_api_rpc/rpc/userclient"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
