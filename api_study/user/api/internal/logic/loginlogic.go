package logic

import (
	"context"
	"fmt"

	"gozero_demo1/api_study/user/api/internal/svc"
	"gozero_demo1/api_study/user/api/internal/types"

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

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	// todo: add your logic here and delete this line
	fmt.Println(req.Username, req.Password)
	return &types.LoginResponse{Code: 0, Data: "hello", Msg: "你好啊"}, nil
}
