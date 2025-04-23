package logic

import (
	"context"
	"fmt"
	"gozero_demo1/model_study/user_gorm/models"

	"gozero_demo1/model_study/user_gorm/api/internal/svc"
	"gozero_demo1/model_study/user_gorm/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp string, err error) {
	// todo: add your logic here and delete this line
	err = l.svcCtx.DB.Create(&models.UserModel{
		Username: "龙哥",
		Password: "1234560",
	}).Error
	if err != nil {
		fmt.Println(err)
	}
	return "ok", nil
}
