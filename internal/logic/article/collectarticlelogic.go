package articlelogic

import (
	"context"
	"database/sql"
	"errors"
	"github.com/foodi-org/foodi-user-service/internal/pkg/servError"
	"github.com/foodi-org/foodi-user-service/model"
	"github.com/foodi-org/foodi-user-service/model/modelType/bo"

	"github.com/foodi-org/foodi-user-service/internal/svc"
	"github.com/foodi-org/foodi-user-service/pb/github.com/foodi-org/foodi-user-service"

	"github.com/zeromicro/go-zero/core/logx"
)

type CollectArticleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCollectArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CollectArticleLogic {
	return &CollectArticleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 收藏文章列表 & 取消收藏
func (l *CollectArticleLogic) CollectArticle(in *foodi_user_service.CollectArticleRequest) (*foodi_user_service.OKReply, error) {
	var err = servError.GRPCErrorDO()
	// 参数校验
	if in.GetUid() == 0 || in.GetArticleID() == 0 {
		return nil, errors.New("invalid parameter")
	}

	// 校验文章
	record, err := l.svcCtx.ArticleModel.FindOne(l.ctx, in.ArticleID)
	if err != nil {
		return nil, errors.New("found article fail")
	}
	if record == nil {
		return nil, errors.New("article not found")
	}

	switch in.GetAction() {
	case foodi_user_service.ActionCoup_ADD:
		_, err = l.svcCtx.SaveArticleModel.Insert(l.ctx, &model.SaveArticleInfo{
			Uid:       sql.NullInt64{Int64: in.Uid, Valid: true},
			ArticleId: sql.NullInt64{Int64: in.ArticleID, Valid: true},
		})
		if err != nil {
			return nil, errors.New("save article fail")
		}
		return &foodi_user_service.OKReply{Ok: true}, nil
	case foodi_user_service.ActionCoup_Cancel:
		if err = l.svcCtx.SaveArticleModel.DelSaveArticle(l.ctx, bo.ArticleBo{
			Uid:       in.GetUid(),
			ArticleID: in.GetArticleID(),
			Type:      foodi_user_service.ActionCoup_Cancel.String(),
		}); err != nil {
			return nil, errors.New("save article fail")
		}
		return &foodi_user_service.OKReply{Ok: true}, nil
	default:
		return nil, servError.NewGRPCError(servError.ActionErrCode, servError.ActionErrMsg)
	}
}
