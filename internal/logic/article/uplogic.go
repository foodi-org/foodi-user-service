package articlelogic

import (
	"context"
	"database/sql"
	"errors"
	"github.com/foodi-org/foodi-user-service/github.com/foodi-org/foodi-user-service"
	"github.com/foodi-org/foodi-user-service/internal/pkg/servError"
	"github.com/foodi-org/foodi-user-service/model"
	"github.com/foodi-org/foodi-user-service/model/modelType/bo"

	"github.com/foodi-org/foodi-user-service/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpLogic {
	return &UpLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Up
//
//	@Description:
//	@param in
//	@return *foodi_user_service.OKReply
//	@return error
func (l *UpLogic) Up(in *foodi_user_service.UpRequest) (*foodi_user_service.OKReply, error) {
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
		if _, err = l.svcCtx.UpModel.Insert(l.ctx, &model.UpInfo{
			Uid:       sql.NullInt64{Int64: in.Uid, Valid: true},
			ArticleId: sql.NullInt64{Int64: in.GetArticleID(), Valid: true},
			CommentId: sql.NullInt64{Int64: in.GetCommentID(), Valid: in.GetCommentID() > 0},
		}); err != nil {
			return nil, err
		} else {
			// 更新文章或者评论的点赞数
			go func() {

			}()
			return &foodi_user_service.OKReply{Ok: true}, nil
		}
	case foodi_user_service.ActionCoup_Cancel:
		if err = l.svcCtx.UpModel.DelArticleUP(l.ctx, bo.DelBO{
			Uid:       in.GetUid(),
			ArticleID: in.GetArticleID(),
		}); err != nil {
			return nil, err
		}
		return &foodi_user_service.OKReply{Ok: true}, nil
	default:
		return nil, servError.NewGRPCError(servError.InvalidAction, servError.Msg(servError.InvalidAction))
	}
}
