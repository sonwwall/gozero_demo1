package useractionlogic

import (
	"context"
	"fmt"

	"gozero_demo1/rpc_study/user_group/rpc/internal/svc"
	"gozero_demo1/rpc_study/user_group/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserCreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserCreateLogic {
	return &UserCreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserCreateLogic) UserCreate(in *user.UserCreateRequest) (*user.UserCreateResponse, error) {
	// todo: add your logic here and delete this line
	fmt.Println("用户已创建")

	return &user.UserCreateResponse{}, nil
}
