package userlogic

import (
	"context"
	"database/sql"
	"errors"
	"github.com/foodi-org/foodi-user-service/model"

	"github.com/foodi-org/foodi-user-service/internal/svc"
	"github.com/foodi-org/foodi-user-service/pb/github.com/foodi-org/foodi-user-service"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveArticleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSaveArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveArticleLogic {
	return &SaveArticleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SaveArticleLogic) SaveArticle(in *foodi_user_service.SaveArticleRequest) (*foodi_user_service.OKReply, error) {
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

	_, err = l.svcCtx.SaveArticleModel.Insert(l.ctx, &model.SaveArticleInfo{
		Uid:       sql.NullInt64{Int64: in.Uid, Valid: true},
		ArticleId: sql.NullInt64{Int64: in.ArticleID, Valid: true},
	})
	if err != nil {
		return nil, errors.New("save article fail")
	}
	return &foodi_user_service.OKReply{Ok: true}, nil
}
