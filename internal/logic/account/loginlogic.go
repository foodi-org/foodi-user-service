package accountlogic

import (
	"context"
	"errors"
	"fmt"
	"github.com/foodi-org/foodi-user-service/internal/pkg/servError"
	"github.com/foodi-org/foodi-user-service/internal/svc"
	"github.com/foodi-org/foodi-user-service/pb/github.com/foodi-org/foodi-user-service"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"golang.org/x/crypto/bcrypt"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Login 用户登录
func (l *LoginLogic) Login(in *foodi_user_service.LoginRequest) (*foodi_user_service.LoginReply, error) {
	if in.GetPhone() == 0 {
		return nil, servError.NewGRPCError(servError.MissPhoneNumber, servError.Msg(servError.MissPhoneNumber))
	}

	switch in.GetLoginType() {
	case foodi_user_service.RegisterCoup_Password:
		if in.GetPassword() == "" {
			return nil, servError.NewGRPCError(servError.MissPassword, servError.Msg(servError.MissPassword))
		}
		switch in.GetType() {
		case foodi_user_service.UserCoup_BUYER:
			account, err := l.svcCtx.AccountModel.FindWithPhone(l.ctx, in.GetPhone())
			switch {
			case errors.Is(err, sqlx.ErrNotFound):
				return nil, servError.NewGRPCError(servError.PhoneNumberNotFound, servError.Msg(servError.PhoneNumberNotFound))
			case err == nil:
				if !account.Password.Valid {
					return nil, servError.NewGRPCError(servError.AccountNotBindPassword, servError.Msg(servError.AccountNotBindPassword))
				}

				// 密码验证
				if err = bcrypt.CompareHashAndPassword([]byte(account.Password.String), []byte(in.GetPassword())); err != nil {
					return nil, servError.NewGRPCError(servError.AccountPasswordNotMatch, servError.Msg(servError.AccountPasswordNotMatch))
				}

				if user, err := l.svcCtx.UserInfoModel.TakeWithAid(l.ctx, account.Id); err != nil {
					return nil, err
				} else {
					return &foodi_user_service.LoginReply{Uid: user.Uid.Int64}, nil
				}
			default:
				return nil, err
			}
		case foodi_user_service.UserCoup_MERCHANT:
			fmt.Println("this is merchant login")
			return &foodi_user_service.LoginReply{}, nil
		case foodi_user_service.UserCoup_RIDER:
			fmt.Println("this is rider login")
			return &foodi_user_service.LoginReply{}, nil
		default:
			return nil, servError.NewGRPCError(servError.InvalidUserCoup, servError.Msg(servError.InvalidUserCoup))
		}
	case foodi_user_service.RegisterCoup_Code:
		fmt.Println("this is code login")
		return &foodi_user_service.LoginReply{}, nil
	case foodi_user_service.RegisterCoup_Phone:
		fmt.Println("this is phone login")
		return &foodi_user_service.LoginReply{}, nil
	default:
		return nil, servError.NewGRPCError(servError.InvalidRegisterCoup, servError.Msg(servError.InvalidRegisterCoup))
	}
}
