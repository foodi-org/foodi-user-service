package userlogic

import (
	"context"
	"database/sql"
	"github.com/foodi-org/foodi-user-service/github.com/foodi-org/foodi-user-service"
	"github.com/foodi-org/foodi-user-service/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserInfoLogic {
	return &UpdateUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserInfoLogic) UpdateUserInfo(in *foodi_user_service.UpdateUserInfoRequest) (*foodi_user_service.UserOKReply, error) {

	user, err := l.svcCtx.UserInfoModel.FindByUID(l.ctx, in.GetUuid())
	if err != nil {
		return nil, err
	}

	if in.GetName() != "" {
		user.Name = sql.NullString{String: in.GetName(), Valid: true}
	}
	if in.GetNikeName() != "" {
		user.NikeName = sql.NullString{String: in.GetNikeName(), Valid: true}
	}
	if in.GetAge() != 0 {
		user.Age = sql.NullInt64{Int64: in.GetAge(), Valid: true}
	}
	if in.GetCardID() != "" {
		user.CardId = sql.NullString{String: in.GetCardID(), Valid: true}
	}
	if in.GetBirthday() != "" {
		user.Birthday = sql.NullString{String: in.GetBirthday(), Valid: true}
	}
	if in.GetGender() != "" {
		user.Gender = sql.NullString{String: in.GetGender(), Valid: true}
	}
	if in.GetRegion() != "" {
		user.Region = sql.NullString{String: in.GetRegion(), Valid: true}
	}

	err = l.svcCtx.UserInfoModel.Update(l.ctx, user)

	return &foodi_user_service.UserOKReply{Ok: err == nil}, err
}
