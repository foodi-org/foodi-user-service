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

// 新增文章
func (l *CreateArticleLogic) CreateArticle(in *foodi_user_service.CreateArticleRequest) (*foodi_user_service.OKReply, error) {
	// todo: add your logic here and delete this line

	_, err := l.svcCtx.ArticleModel.Insert(l.ctx, &model.ArticleInfo{
		Uid:         in.GetUid(),
		Published:   sql.NullInt64{},
		PublishedAt: sql.NullTime{},
		ArticleType: "",
		Title:       sql.NullString{in.GetTitle(), true},
		Video:       sql.NullString{},
		ImageUrl:    sql.NullString{},
		Content:     sql.NullString{},
		Region:      sql.NullString{},
		Location:    sql.NullString{},
		Merchant:    sql.NullInt64{},
		Cover:       sql.NullString{},
	})
	if err != nil {
		return nil, err
	}

	return &foodi_user_service.OKReply{Ok: err == nil}, nil
}
