package userlogic

import (
	"context"
	"github.com/foodi-org/foodi-user-service/model/modelType/bo"

	"github.com/foodi-org/foodi-user-service/internal/svc"
	"github.com/foodi-org/foodi-user-service/pb/github.com/foodi-org/foodi-user-service"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserImageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserImageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserImageLogic {
	return &UserImageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新用户头像
func (l *UserImageLogic) UserImage(in *foodi_user_service.UserImageRequest) (*foodi_user_service.OKReply, error) {

	return &foodi_user_service.OKReply{
		Ok: nil == l.svcCtx.UserInfoModel.UpdateImage(l.ctx, &bo.UpdateImageBO{
			Uid:   in.GetUid(),
			Image: in.GetImage(),
		}),
	}, nil
}