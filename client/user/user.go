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
	AddCommentRequest    = foodi_user_service.AddCommentRequest
	CreateArticleRequest = foodi_user_service.CreateArticleRequest
	OKReply              = foodi_user_service.OKReply
	RegisterReply        = foodi_user_service.RegisterReply
	RegisterRequest      = foodi_user_service.RegisterRequest
	SaveArticleRequest   = foodi_user_service.SaveArticleRequest
	TokenReply           = foodi_user_service.TokenReply
	TokenRequest         = foodi_user_service.TokenRequest
	UpRequest            = foodi_user_service.UpRequest
	UserDetailReply      = foodi_user_service.UserDetailReply
	UserDetailRequest    = foodi_user_service.UserDetailRequest
	UserReply            = foodi_user_service.UserReply
	UserRequest          = foodi_user_service.UserRequest

	User interface {
		BaseInfo(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserReply, error)
		DetailInfo(ctx context.Context, in *UserDetailRequest, opts ...grpc.CallOption) (*UserDetailReply, error)
		CreateArticle(ctx context.Context, in *CreateArticleRequest, opts ...grpc.CallOption) (*OKReply, error)
		AddComment(ctx context.Context, in *AddCommentRequest, opts ...grpc.CallOption) (*OKReply, error)
		Up(ctx context.Context, in *UpRequest, opts ...grpc.CallOption) (*OKReply, error)
		SaveArticle(ctx context.Context, in *SaveArticleRequest, opts ...grpc.CallOption) (*OKReply, error)
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

func (m *defaultUser) BaseInfo(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserReply, error) {
	client := foodi_user_service.NewUserClient(m.cli.Conn())
	return client.BaseInfo(ctx, in, opts...)
}

func (m *defaultUser) DetailInfo(ctx context.Context, in *UserDetailRequest, opts ...grpc.CallOption) (*UserDetailReply, error) {
	client := foodi_user_service.NewUserClient(m.cli.Conn())
	return client.DetailInfo(ctx, in, opts...)
}

func (m *defaultUser) CreateArticle(ctx context.Context, in *CreateArticleRequest, opts ...grpc.CallOption) (*OKReply, error) {
	client := foodi_user_service.NewUserClient(m.cli.Conn())
	return client.CreateArticle(ctx, in, opts...)
}

func (m *defaultUser) AddComment(ctx context.Context, in *AddCommentRequest, opts ...grpc.CallOption) (*OKReply, error) {
	client := foodi_user_service.NewUserClient(m.cli.Conn())
	return client.AddComment(ctx, in, opts...)
}

func (m *defaultUser) Up(ctx context.Context, in *UpRequest, opts ...grpc.CallOption) (*OKReply, error) {
	client := foodi_user_service.NewUserClient(m.cli.Conn())
	return client.Up(ctx, in, opts...)
}

func (m *defaultUser) SaveArticle(ctx context.Context, in *SaveArticleRequest, opts ...grpc.CallOption) (*OKReply, error) {
	client := foodi_user_service.NewUserClient(m.cli.Conn())
	return client.SaveArticle(ctx, in, opts...)
}
