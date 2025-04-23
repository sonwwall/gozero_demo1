package logic

import (
	"context"
	"encoding/json"

	"gozero_demo1/api_study/user/api_jwt/internal/svc"
	"gozero_demo1/api_study/user/api_jwt/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo() (resp *types.UserInfoResponse, err error) {
	// todo: add your logic here and delete this line
	Userid := l.ctx.Value("user_id").(json.Number)
	UserId, _ := Userid.Int64()

	Username := l.ctx.Value("username").(string)

	return &types.UserInfoResponse{
		UserId:   uint(UserId),
		Username: Username,
	}, nil
}
