package userlogic

import (
	"context"

	"github.com/foodi-org/foodi-user-service/github.com/foodi-org/foodi-user-service"
	"github.com/foodi-org/foodi-user-service/internal/svc"

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

// DetailInfo 获取用户详情
func (l *DetailInfoLogic) DetailInfo(in *foodi_user_service.UserDetailRequest) (*foodi_user_service.UserDetailReply, error) {
	if user, err := l.svcCtx.UserInfoModel.FindByUID(l.ctx, in.GetUid()); err != nil {
		return nil, err
	} else {
		return &foodi_user_service.UserDetailReply{
			Name:     user.Name.String,
			NikeName: user.NikeName.String,
			Image:    user.Image.String,
			Age:      int32(user.Age.Int64),
			Gender:   user.Gender.String,
			Birthday: user.Birthday.String,
			Region:   user.Region.String,
			LV:       int32(user.Lv.Int64),
			VIP:      int32(user.Vip.Int64),
		}, nil
	}
}
