package articlelogic

import (
	"context"
	"database/sql"
	"github.com/foodi-org/foodi-user-service/model"

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

// CreateArticle 新增文章
func (l *CreateArticleLogic) CreateArticle(in *foodi_user_service.CreateArticleRequest) (*foodi_user_service.OKReply, error) {

	_, err := l.svcCtx.ArticleModel.Insert(l.ctx, &model.ArticleInfo{
		Uid:         in.GetUid(),
		Published:   sql.NullInt64{},
		PublishedAt: sql.NullTime{},
		ArticleType: in.GetType().String(),
		Title:       sql.NullString{String: in.GetTitle(), Valid: true},
		Video:       sql.NullString{},
		ImageUrl:    sql.NullString{},
		Content:     sql.NullString{},
		Region:      sql.NullString{String: in.GetRegion(), Valid: in.GetRegion() != ""},
		Location:    sql.NullString{},
		Merchant:    sql.NullInt64{},
		Cover:       sql.NullString{String: in.GetCover(), Valid: in.GetCover() != ""},
	})
	if err != nil {
		l.Logger.Error("创建文章执行sql失败", "error:", err)
		return nil, err
	}

	return &foodi_user_service.OKReply{Ok: err == nil}, nil
}
