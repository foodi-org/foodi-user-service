// Code generated by goctl. DO NOT EDIT.
// Source: foodiUserService.proto

package user

import (
	"context"

	"github.com/foodi-org/foodi-user-service/pb/github.com/foodi-org/foodi-user-service"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddCommentRequest     = foodi_user_service.AddCommentRequest
	ArticleInfo           = foodi_user_service.ArticleInfo
	ArticleListReply      = foodi_user_service.ArticleListReply
	ArticleListRequest    = foodi_user_service.ArticleListRequest
	ArticleRefreshRequest = foodi_user_service.ArticleRefreshRequest
	Author                = foodi_user_service.Author
	CollectArticleRequest = foodi_user_service.CollectArticleRequest
	CreateArticleRequest  = foodi_user_service.CreateArticleRequest
	DelArticleRequest     = foodi_user_service.DelArticleRequest
	Draft                 = foodi_user_service.Draft
	DraftListReply        = foodi_user_service.DraftListReply
	DraftListRequest      = foodi_user_service.DraftListRequest
	LoginReply            = foodi_user_service.LoginReply
	LoginRequest          = foodi_user_service.LoginRequest
	OKReply               = foodi_user_service.OKReply
	ReadHistoryRequest    = foodi_user_service.ReadHistoryRequest
	RegisterReply         = foodi_user_service.RegisterReply
	RegisterRequest       = foodi_user_service.RegisterRequest
	UpRequest             = foodi_user_service.UpRequest
	UserDetailReply       = foodi_user_service.UserDetailReply
	UserDetailRequest     = foodi_user_service.UserDetailRequest
	UserImageRequest      = foodi_user_service.UserImageRequest
	UserReply             = foodi_user_service.UserReply
	UserRequest           = foodi_user_service.UserRequest

	User interface {
		// 获取用户基础信息
		BaseInfo(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserReply, error)
		// 获取用户详情
		DetailInfo(ctx context.Context, in *UserDetailRequest, opts ...grpc.CallOption) (*UserDetailReply, error)
		// 更新用户头像
		UserImage(ctx context.Context, in *UserImageRequest, opts ...grpc.CallOption) (*OKReply, error)
	}

	defaultUser struct {
		cli zrpc.Client
	}
)

func NewUser(cli zrpc.Client) User {
	return &defaultUser{
		cli: cli,
	}
}

// 获取用户基础信息
func (m *defaultUser) BaseInfo(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserReply, error) {
	client := foodi_user_service.NewUserClient(m.cli.Conn())
	return client.BaseInfo(ctx, in, opts...)
}

// 获取用户详情
func (m *defaultUser) DetailInfo(ctx context.Context, in *UserDetailRequest, opts ...grpc.CallOption) (*UserDetailReply, error) {
	client := foodi_user_service.NewUserClient(m.cli.Conn())
	return client.DetailInfo(ctx, in, opts...)
}

// 更新用户头像
func (m *defaultUser) UserImage(ctx context.Context, in *UserImageRequest, opts ...grpc.CallOption) (*OKReply, error) {
	client := foodi_user_service.NewUserClient(m.cli.Conn())
	return client.UserImage(ctx, in, opts...)
}
