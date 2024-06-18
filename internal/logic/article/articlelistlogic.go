package articlelogic

import (
	"context"

	"github.com/foodi-org/foodi-user-service/internal/svc"
	"github.com/foodi-org/foodi-user-service/pb/github.com/foodi-org/foodi-user-service"

	"github.com/zeromicro/go-zero/core/logx"
)

type ArticleListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewArticleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ArticleListLogic {
	return &ArticleListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 用户下滑继续获取文章列表
func (l *ArticleListLogic) ArticleList(in *foodi_user_service.ArticleListRequest) (*foodi_user_service.ArticleListReply, error) {
	// todo: add your logic here and delete this line

	return &foodi_user_service.ArticleListReply{}, nil
}
