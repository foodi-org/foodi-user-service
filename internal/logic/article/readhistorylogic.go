package articlelogic

import (
	"context"
	"database/sql"
	"github.com/foodi-org/foodi-user-service/internal/pkg/servError"
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

// ReadHistory 获取用户最近一周的文章访问历史记录
func (l *ReadHistoryLogic) ReadHistory(in *foodi_user_service.ReadHistoryRequest) (*foodi_user_service.ArticleListReply, error) {
	var datetime string
	switch in.Action {
	case foodi_user_service.ActionCoup_ADD: // 将文章加入用户近期阅读列表
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
		return nil, servError.NewGRPCError(servError.ArticleNotFound, servError.Msg(servError.ArticleNotFound))
	case foodi_user_service.ActionCoup_GET: // 获取用户近期阅读列表
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
		if len(ids) == 0 {
			return &foodi_user_service.ArticleListReply{
				Size: 0,
				End:  0,
				List: make([]*foodi_user_service.ArticleInfo, 0),
			}, nil
		}

		res, err := l.svcCtx.ArticleModel.ListByIDS(l.ctx, ids)
		if err != nil {
			return nil, err
		}

		list := make([]*foodi_user_service.ArticleInfo, len(res))
		for k, item := range res {
			user, err := l.svcCtx.UserInfoModel.FindOne(l.ctx, item.Uid)
			if err != nil {
				continue
			}
			list[k] = &foodi_user_service.ArticleInfo{
				Id: item.Id,
				Author: &foodi_user_service.Author{
					Uid:   user.Uid.Int64,
					Name:  user.Name.String,
					Image: user.Image.String,
				},
				UpNum:     item.Up,
				Title:     item.Title.String,
				Cover:     item.Cover.String,
				CoverType: foodi_user_service.ResourceCoup_IMAGE, // todo 修改model field type
			}
		}
		return &foodi_user_service.ArticleListReply{
			End:  ids[len(ids)-1],
			Size: int64(len(list)),
			List: list,
		}, nil
	default:
		return &foodi_user_service.ArticleListReply{}, servError.NewGRPCError(servError.InvalidAction, servError.Msg(servError.InvalidAction))
	}
}
