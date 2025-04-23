package logic

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"gozero_demo1/model_study/user/api/internal/svc"
	"gozero_demo1/model_study/user/api/internal/types"
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
	//res, err := l.svcCtx.UsersModel.Insert(l.ctx, &model.User{
	//	Username: "外城",
	//	Password: "123456",
	//})
	//if err != nil {
	//	return "", err
	//}
	//fmt.Println(res)
	//return "ok", nil

	user, err := l.svcCtx.UsersModel.FindOneByUsernameAndPassword(l.ctx, req.Username, req.Password)
	if err != nil {
		return "", errors.New("登录失败")
	}

	return user.Username, nil
}
