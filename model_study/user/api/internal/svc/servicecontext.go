package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"gozero_demo1/model_study/user/api/internal/config"
	"gozero_demo1/model_study/user/model"
)

type ServiceContext struct {
	Config     config.Config
	UsersModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlConn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:     c,
		UsersModel: model.NewUserModel(mysqlConn),
	}
}
