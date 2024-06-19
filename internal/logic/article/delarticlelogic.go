package articlelogic

import (
	"context"

	"github.com/foodi-org/foodi-user-service/internal/svc"
	"github.com/foodi-org/foodi-user-service/pb/github.com/foodi-org/foodi-user-service"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelArticleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelArticleLogic {
	return &DelArticleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DelArticle 删除文章
func (l *DelArticleLogic) DelArticle(in *foodi_user_service.DelArticleRequest) (*foodi_user_service.OKReply, error) {
	if err := l.svcCtx.ArticleModel.Disable(l.ctx, in.GetId(), in.GetUid()); err != nil {
		return &foodi_user_service.OKReply{}, err
	}
	return &foodi_user_service.OKReply{Ok: true}, nil
}
