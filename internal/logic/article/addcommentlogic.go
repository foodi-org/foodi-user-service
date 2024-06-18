package articlelogic

import (
	"context"
	"database/sql"
	"github.com/foodi-org/foodi-user-service/internal/pkg/servError"
	"github.com/foodi-org/foodi-user-service/internal/svc"
	"github.com/foodi-org/foodi-user-service/model"
	"github.com/foodi-org/foodi-user-service/pb/github.com/foodi-org/foodi-user-service"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCommentLogic {
	return &AddCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddComment 评论文章或者回复评论
func (l *AddCommentLogic) AddComment(in *foodi_user_service.AddCommentRequest) (*foodi_user_service.OKReply, error) {
	if _, err := l.svcCtx.ArticleModel.FindOne(l.ctx, in.GetArticleID()); err != nil {
		return nil, servError.NewGRPCError(servError.ArticleNotFound, servError.Msg(servError.ArticleNotFound))
	}

	// 回复评论, 检查 parent id
	if in.GetParentID() > 0 {
		_, err := l.svcCtx.ArticleCommentModel.FindOne(l.ctx, in.GetParentID())
		if err != nil {
			return nil, err
		}
	}

	// todo 加锁更新 reply count

	_, err := l.svcCtx.ArticleCommentModel.Insert(l.ctx, &model.ArticleCommentInfo{
		ArticleId:      sql.NullInt64{Int64: in.GetArticleID(), Valid: true},
		Uid:            sql.NullInt64{Int64: in.GetUid(), Valid: true},
		Content:        sql.NullString{String: in.GetContent(), Valid: true},
		PostTime:       sql.NullTime{Time: time.Now(), Valid: true},
		ParentId:       sql.NullInt64{Int64: in.GetParentID(), Valid: in.GetParentID() > 0},
		FirstCommentId: sql.NullInt64{},
		UpCount:        sql.NullInt64{},
		ReplyCount:     sql.NullInt64{},
		IsAnonymous:    in.GetIsAnonymous(),
	})
	return &foodi_user_service.OKReply{Ok: err == nil}, err
}
