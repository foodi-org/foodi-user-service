package articlelogic

import (
	"context"

	"github.com/foodi-org/foodi-user-service/internal/svc"
	"github.com/foodi-org/foodi-user-service/pb/github.com/foodi-org/foodi-user-service"

	"github.com/zeromicro/go-zero/core/logx"
)

type ArticleRefreshLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewArticleRefreshLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ArticleRefreshLogic {
	return &ArticleRefreshLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 用户下拉刷新文章列表
func (l *ArticleRefreshLogic) ArticleRefresh(in *foodi_user_service.ArticleRefreshRequest) (*foodi_user_service.ArticleListReply, error) {
	// todo: add your logic here and delete this line

	return &foodi_user_service.ArticleListReply{}, nil
}
