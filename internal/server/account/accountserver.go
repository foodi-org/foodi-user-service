// Code generated by goctl. DO NOT EDIT.
// Source: foodiUserService.proto

package server

import (
	"context"

	"github.com/foodi-org/foodi-user-service/internal/logic/account"
	"github.com/foodi-org/foodi-user-service/internal/svc"
	"github.com/foodi-org/foodi-user-service/pb/github.com/foodi-org/foodi-user-service"
)

type AccountServer struct {
	svcCtx *svc.ServiceContext
	foodi_user_service.UnimplementedAccountServer
}

func NewAccountServer(svcCtx *svc.ServiceContext) *AccountServer {
	return &AccountServer{
		svcCtx: svcCtx,
	}
}

// 普通用户注册
func (s *AccountServer) Register(ctx context.Context, in *foodi_user_service.RegisterRequest) (*foodi_user_service.RegisterReply, error) {
	l := accountlogic.NewRegisterLogic(ctx, s.svcCtx)
	return l.Register(in)
}

// 用户登录
func (s *AccountServer) Login(ctx context.Context, in *foodi_user_service.LoginRequest) (*foodi_user_service.LoginReply, error) {
	l := accountlogic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}
