package userlogic

import (
	"context"

	"github.com/foodi-org/foodi-user-service/internal/svc"
	"github.com/foodi-org/foodi-user-service/pb/github.com/foodi-org/foodi-user-service"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateArticleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateArticleLogic {
	return &CreateArticleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateArticleLogic) CreateArticle(in *foodi_user_service.CreateArticleRequest) (*foodi_user_service.OKReply, error) {
	// todo: add your logic here and delete this line

	return &foodi_user_service.OKReply{}, nil
}
