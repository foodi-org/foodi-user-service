// Code generated by goctl. DO NOT EDIT.
// Source: foodiUserService.proto

package server

import (
	"context"

	"github.com/foodi-org/foodi-user-service/internal/logic/article"
	"github.com/foodi-org/foodi-user-service/internal/svc"
	"github.com/foodi-org/foodi-user-service/pb/github.com/foodi-org/foodi-user-service"
)

type ArticleServer struct {
	svcCtx *svc.ServiceContext
	foodi_user_service.UnimplementedArticleServer
}

func NewArticleServer(svcCtx *svc.ServiceContext) *ArticleServer {
	return &ArticleServer{
		svcCtx: svcCtx,
	}
}

// 新增文章
func (s *ArticleServer) CreateArticle(ctx context.Context, in *foodi_user_service.CreateArticleRequest) (*foodi_user_service.OKReply, error) {
	l := articlelogic.NewCreateArticleLogic(ctx, s.svcCtx)
	return l.CreateArticle(in)
}

// 评论文章或者回复评论
func (s *ArticleServer) AddComment(ctx context.Context, in *foodi_user_service.AddCommentRequest) (*foodi_user_service.OKReply, error) {
	l := articlelogic.NewAddCommentLogic(ctx, s.svcCtx)
	return l.AddComment(in)
}

// 点赞文章或者评论
func (s *ArticleServer) Up(ctx context.Context, in *foodi_user_service.UpRequest) (*foodi_user_service.OKReply, error) {
	l := articlelogic.NewUpLogic(ctx, s.svcCtx)
	return l.Up(in)
}

// 收藏文章列表 & 取消收藏
func (s *ArticleServer) CollectArticle(ctx context.Context, in *foodi_user_service.CollectArticleRequest) (*foodi_user_service.OKReply, error) {
	l := articlelogic.NewCollectArticleLogic(ctx, s.svcCtx)
	return l.CollectArticle(in)
}

// 删除文章
func (s *ArticleServer) DelArticle(ctx context.Context, in *foodi_user_service.DelArticleRequest) (*foodi_user_service.OKReply, error) {
	l := articlelogic.NewDelArticleLogic(ctx, s.svcCtx)
	return l.DelArticle(in)
}

// 获取用户草稿文章列表
func (s *ArticleServer) DraftList(ctx context.Context, in *foodi_user_service.DraftListRequest) (*foodi_user_service.DraftListReply, error) {
	l := articlelogic.NewDraftListLogic(ctx, s.svcCtx)
	return l.DraftList(in)
}

// 获取用户最近一周的文章访问历史记录
func (s *ArticleServer) ReadHistory(ctx context.Context, in *foodi_user_service.ReadHistoryRequest) (*foodi_user_service.ArticleListReply, error) {
	l := articlelogic.NewReadHistoryLogic(ctx, s.svcCtx)
	return l.ReadHistory(in)
}