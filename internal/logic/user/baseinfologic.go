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

// BaseInfo 获取用户基础信息
func (l *BaseInfoLogic) BaseInfo(in *foodi_user_service.UserRequest) (*foodi_user_service.UserReply, error) {
	if user, err := l.svcCtx.UserInfoModel.FindByUID(l.ctx, in.GetUid()); err != nil {
		return nil, err
	} else {
		return &foodi_user_service.UserReply{
			Name:     user.Name.String,
			NikeName: user.NikeName.String,
			Image:    user.Image.String,
			LV:       int32(user.Lv.Int64),
		}, nil
	}
}
