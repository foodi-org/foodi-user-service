package articlelogic

import (
	"context"
	"github.com/foodi-org/foodi-user-service/pkg/ternary"
	"time"

	"github.com/foodi-org/foodi-user-service/internal/svc"
	"github.com/foodi-org/foodi-user-service/pb/github.com/foodi-org/foodi-user-service"

	"github.com/zeromicro/go-zero/core/logx"
)

type DraftListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDraftListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DraftListLogic {
	return &DraftListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DraftList 获取用户草稿文章列表
// @param uid 用户id
func (l *DraftListLogic) DraftList(in *foodi_user_service.DraftListRequest) (*foodi_user_service.DraftListReply, error) {
	query, err := l.svcCtx.ArticleModel.DraftList(l.ctx, in.GetUid())
	if err != nil {
		return nil, err
	}
	res := make([]*foodi_user_service.Draft, len(query))
	for k, v := range query {
		res[k] = &foodi_user_service.Draft{
			Id:    v.ID,
			Title: v.Title,
			Time:  ternary.Ternary(v.UpdatedAt.Valid, v.UpdatedAt, v.CreatedAt).Time.Format(time.DateTime),
		}
	}

	return &foodi_user_service.DraftListReply{List: res}, nil
}
