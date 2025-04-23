package logic

import (
	"context"
	"errors"
	"gozero_demo1/rpc_study/user_api_rpc/models"

	"gozero_demo1/rpc_study/user_api_rpc/rpc/internal/svc"
	"gozero_demo1/rpc_study/user_api_rpc/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserInfoLogic) UserInfo(in *user.UserInfoRequest) (*user.UserInfoResponse, error) {
	var model models.UserModel
	err := l.svcCtx.DB.Take(&model, in.UserId).Error
	if err != nil {
		return nil, errors.New("用户不存在")
	}

	return &user.UserInfoResponse{
		UserId:   uint32(model.ID),
		Username: model.Username,
	}, nil
}
