package articlelogic

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/foodi-org/foodi-user-service/internal/pkg/servError"
	"github.com/foodi-org/foodi-user-service/internal/svc"
	"github.com/foodi-org/foodi-user-service/model"
	"github.com/foodi-org/foodi-user-service/pb/github.com/foodi-org/foodi-user-service"
	"math/rand"
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
	var parentID int64
	if _, err := l.svcCtx.ArticleModel.FindOne(l.ctx, in.GetArticleID()); err != nil {
		return nil, servError.NewGRPCError(servError.ArticleNotFound, servError.Msg(servError.ArticleNotFound))
	}

	// 没有父节点id，当前为给文章添加评论
	if in.GetParentID() == 0 {
		if _, err := l.svcCtx.ArticleCommentModel.Insert(l.ctx, &model.ArticleCommentInfo{
			ArticleId:   sql.NullInt64{Int64: in.GetArticleID(), Valid: true},
			Uid:         sql.NullInt64{Int64: in.GetUid(), Valid: true},
			Content:     sql.NullString{String: in.GetContent(), Valid: true},
			PostTime:    sql.NullTime{Time: time.Now(), Valid: true},
			IsAnonymous: in.GetIsAnonymous(),
		}); err != nil {
			return &foodi_user_service.OKReply{Ok: false}, err
		} else {
			return &foodi_user_service.OKReply{Ok: true}, nil
		}
	}

	// 若为回复评论, 检查 parent id 有有效性
	if in.GetParentID() > 0 {
		parent, err := l.svcCtx.ArticleCommentModel.FindOne(l.ctx, in.GetParentID())
		if err != nil {
			return nil, servError.NewGRPCError(servError.CommentNotFound, servError.Msg(servError.CommentNotFound))
		}
		parentID = parent.Id
	}

	_, err := l.svcCtx.ArticleCommentModel.Insert(l.ctx, &model.ArticleCommentInfo{
		ArticleId:      sql.NullInt64{Int64: in.GetArticleID(), Valid: true},
		Uid:            sql.NullInt64{Int64: in.GetUid(), Valid: true},
		Content:        sql.NullString{String: in.GetContent(), Valid: true},
		PostTime:       sql.NullTime{Time: time.Now(), Valid: true},
		ParentId:       sql.NullInt64{Int64: in.GetParentID(), Valid: in.GetParentID() > 0},
		FirstCommentId: sql.NullInt64{Int64: parentID, Valid: parentID > 0},
		IsAnonymous:    in.GetIsAnonymous(),
	})
	if err != nil {
		return nil, servError.NewGRPCError(servError.ArticleCreateFail, servError.Msg(servError.ArticleCreateFail))
	}

	// 异步加锁更新 reply count
	go func() {
		var getMtx bool
		value := rand.Intn(100)
		now := time.Now().String()
		expireTime := rand.Intn(5) + 5 // 随机范围过期时间，避免雪崩
		defer func() {
			if getMtx {
				_, _ = l.svcCtx.Redis.DelCtx(l.ctx, "reply_count_mtx") // 释放锁
			}
		}()
		for {
			select {
			case <-time.After(500 * time.Millisecond):
				getMtx, err = l.svcCtx.Redis.SetnxEx("reply_count_mtx", fmt.Sprintf("%d-%s", value, now), expireTime)
				if err != nil {
					l.Logger.Errorf("redis set ex err: %v", err)
					return
				}
				// 抢锁成功,更新回复顶节点评论的回复数量
				if getMtx {
					_ = l.svcCtx.ArticleCommentModel.ReplyAddOne(l.ctx, parentID)
					return
				}
			case <-time.After(time.Duration(5*expireTime) * time.Second):
				l.Logger.Errorf("up article comment reply count fail, redis ex expired")
			}
		}
	}()

	return &foodi_user_service.OKReply{Ok: err == nil}, err
}
