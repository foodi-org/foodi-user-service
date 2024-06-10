package articlelogic

import (
	"context"

	"github.com/foodi-org/foodi-user-service/internal/svc"
	"github.com/foodi-org/foodi-user-service/pb/github.com/foodi-org/foodi-user-service"

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

// 评论文章或者回复评论
func (l *AddCommentLogic) AddComment(in *foodi_user_service.AddCommentRequest) (*foodi_user_service.OKReply, error) {
	// todo: add your logic here and delete this line

	return &foodi_user_service.OKReply{}, nil
}
