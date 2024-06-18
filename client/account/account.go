// Code generated by goctl. DO NOT EDIT.
// Source: foodiUserService.proto

package account

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
	OKReply               = foodi_user_service.OKReply
	ReadHistoryRequest    = foodi_user_service.ReadHistoryRequest
	RegisterReply         = foodi_user_service.RegisterReply
	RegisterRequest       = foodi_user_service.RegisterRequest
	TokenReply            = foodi_user_service.TokenReply
	TokenRequest          = foodi_user_service.TokenRequest
	UpRequest             = foodi_user_service.UpRequest
	UserDetailReply       = foodi_user_service.UserDetailReply
	UserDetailRequest     = foodi_user_service.UserDetailRequest
	UserImageRequest      = foodi_user_service.UserImageRequest
	UserReply             = foodi_user_service.UserReply
	UserRequest           = foodi_user_service.UserRequest

	Account interface {
		// 获取token
		Token(ctx context.Context, in *TokenRequest, opts ...grpc.CallOption) (*TokenReply, error)
		// 普通用户注册
		Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterReply, error)
	}

	defaultAccount struct {
		cli zrpc.Client
	}
)

func NewAccount(cli zrpc.Client) Account {
	return &defaultAccount{
		cli: cli,
	}
}

// 获取token
func (m *defaultAccount) Token(ctx context.Context, in *TokenRequest, opts ...grpc.CallOption) (*TokenReply, error) {
	client := foodi_user_service.NewAccountClient(m.cli.Conn())
	return client.Token(ctx, in, opts...)
}

// 普通用户注册
func (m *defaultAccount) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterReply, error) {
	client := foodi_user_service.NewAccountClient(m.cli.Conn())
	return client.Register(ctx, in, opts...)
}
