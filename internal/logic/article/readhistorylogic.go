package articlelogic

import (
	"context"
	"database/sql"
	"errors"
	"github.com/foodi-org/foodi-user-service/internal/svc"
	"github.com/foodi-org/foodi-user-service/model"
	"github.com/foodi-org/foodi-user-service/pb/github.com/foodi-org/foodi-user-service"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

type ReadHistoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewReadHistoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReadHistoryLogic {
	return &ReadHistoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取用户最近一周的文章访问历史记录
func (l *ReadHistoryLogic) ReadHistory(in *foodi_user_service.ReadHistoryRequest) (*foodi_user_service.ArticleListReply, error) {
	var datetime string
	switch in.Action {
	case foodi_user_service.ActionCoup_ADD:
		info, err := l.svcCtx.ArticleModel.FindOne(l.ctx, in.GetAid())
		if err != nil {
			return nil, err
		}
		if info == nil {
			return nil, err
		}
		if !info.DeletedAt.Valid && info.PublishedAt.Valid {
			_, err = l.svcCtx.ArticleReadModel.Insert(l.ctx, &model.ArticleReadInfo{
				Uid: sql.NullInt64{Int64: in.GetUid(), Valid: in.GetUid() != 0},
				Aid: sql.NullInt64{Int64: in.GetAid(), Valid: in.GetAid() != 0},
			})
			return nil, nil
		}
		return nil, errors.New("unValid aid")
	case foodi_user_service.ActionCoup_GET:
		if in.GetDatetime() == "" {
			now := time.Now().AddDate(0, 0, -7)
			datetime = now.Format("2006-01-02")
		} else {
			now, err := time.Parse(time.DateTime, in.GetDatetime())
			if err != nil {
				return nil, err
			}
			datetime = now.Format("2006-01-02")
		}

		ids, err := l.svcCtx.ArticleReadModel.AidsAfterDate(l.ctx, in.GetUid(), datetime)
		if err != nil {
			return nil, err
		}
	}

	return &foodi_user_service.ArticleListReply{}, nil
}
