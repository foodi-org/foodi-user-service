package accountlogic

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"github.com/foodi-org/foodi-user-service/github.com/foodi-org/foodi-user-service"
	"github.com/foodi-org/foodi-user-service/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type PhoneExistLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPhoneExistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PhoneExistLogic {
	return &PhoneExistLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PhoneExistLogic) PhoneExist(in *foodi_user_service.PhoneExistRequest) (*foodi_user_service.PhoneExistReply, error) {
	_, err := l.svcCtx.AccountModel.FindWithPhone(l.ctx, in.GetPhone())
	switch {
	case errors.Is(err, sqlx.ErrNotFound):
		return &foodi_user_service.PhoneExistReply{}, nil
	case err == nil:
		return &foodi_user_service.PhoneExistReply{Exist: true}, nil
	default:
		return nil, err
	}
}
