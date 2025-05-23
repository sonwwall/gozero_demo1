package logic

import (
	"context"
	"errors"
	"gozero_demo1/rpc_study/user_api_rpc/rpc/types/user"

	"gozero_demo1/rpc_study/user_api_rpc/api/internal/svc"
	"gozero_demo1/rpc_study/user_api_rpc/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserCreateLogic {
	return &UserCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserCreateLogic) UserCreate(req *types.UserCreateRequest) (resp string, err error) {
	// todo: add your logic here and delete this line

	response, err := l.svcCtx.UserRpc.UserCreate(l.ctx, &user.UserCreateRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return "", err
	}
	if response.Err != "" {
		return "", errors.New(response.Err)
	}
	return
}
