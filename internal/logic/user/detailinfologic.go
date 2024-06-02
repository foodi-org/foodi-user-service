package userlogic

import (
	"context"

	"github.com/foodi-org/foodi-user-service/internal/svc"
	"github.com/foodi-org/foodi-user-service/pb/github.com/foodi-org/foodi-user-service"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDetailInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailInfoLogic {
	return &DetailInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DetailInfoLogic) DetailInfo(in *foodi_user_service.UserDetailRequest) (*foodi_user_service.UserDetailReply, error) {
	// todo: add your logic here and delete this line

	return &foodi_user_service.UserDetailReply{}, nil
}
