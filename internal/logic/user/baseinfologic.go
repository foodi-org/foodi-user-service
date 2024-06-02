package userlogic

import (
	"context"

	"github.com/foodi-org/foodi-user-service/internal/svc"
	"github.com/foodi-org/foodi-user-service/pb/github.com/foodi-org/foodi-user-service"

	"github.com/zeromicro/go-zero/core/logx"
)

type BaseInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBaseInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BaseInfoLogic {
	return &BaseInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BaseInfoLogic) BaseInfo(in *foodi_user_service.UserRequest) (*foodi_user_service.UserReply, error) {
	// todo: add your logic here and delete this line

	return &foodi_user_service.UserReply{}, nil
}
