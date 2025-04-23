package svc

import (
	"gorm.io/gorm"
	"gozero_demo1/common/init_gorm"
	"gozero_demo1/rpc_study/user_api_rpc/rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := init_gorm.InitGorm(c.Mysql.DataSource)
	return &ServiceContext{
		Config: c,
		DB:     db,
	}
}
