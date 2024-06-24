package accountlogic

import (
	"context"
	"database/sql"
	_const "github.com/foodi-org/foodi-user-service/internal/const"
	"github.com/foodi-org/foodi-user-service/internal/pkg/snowflake"
	"github.com/foodi-org/foodi-user-service/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"time"

	"github.com/foodi-org/foodi-user-service/internal/svc"
	"github.com/foodi-org/foodi-user-service/pb/github.com/foodi-org/foodi-user-service"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *foodi_user_service.RegisterRequest) (*foodi_user_service.RegisterReply, error) {
	if err := l.svcCtx.DB.Transact(func(session sqlx.Session) error {
		res, err := l.svcCtx.AccountModel.TransInsert(l.ctx, session, &model.AccountInfo{
			Type:          sql.NullString{String: in.GetType().String(), Valid: true},
			NikeName:      sql.NullString{String: "自动昵称", Valid: true},
			Image:         sql.NullString{String: "这是基础用户头像oss地址", Valid: true},
			Phone:         sql.NullInt64{Int64: in.GetPhone(), Valid: true},
			Password:      sql.NullString{String: in.GetPassword(), Valid: len(in.GetPassword()) > 0},
			FirstRegister: sql.NullString{String: time.Now().Format(time.DateTime), Valid: true},
			LastLogin:     sql.NullString{String: time.Now().Format(time.DateTime), Valid: true},
		})
		if err != nil {
			return err
		}

		aid, err := res.LastInsertId()
		if err != nil {
			return err
		}

		if _, err = l.svcCtx.UserInfoModel.TransInsert(l.ctx, session, &model.UserInfo{
			Uid:      sql.NullInt64{Int64: snowflake.RuleUID(_const.MachineLen), Valid: true},
			Aid:      sql.NullInt64{Int64: aid, Valid: true},
			Name:     sql.NullString{},
			CardType: 0,
			CardId:   sql.NullString{},
			Gender:   sql.NullString{},
			Age:      sql.NullInt64{},
			Birthday: sql.NullString{},
			Region:   sql.NullString{},
			WeChatId: sql.NullString{},
			Lv:       sql.NullInt64{},
			Vip:      sql.NullInt64{},
			NikeName: sql.NullString{},
			Image:    sql.NullString{},
		}); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}

	if reply, err := NewLoginLogic(l.ctx, l.svcCtx).Login(&foodi_user_service.LoginRequest{
		Type:      in.GetType(),
		LoginType: in.GetRegisterType(),
		Phone:     in.GetPhone(),
		Code:      in.GetCode(),
		Password:  in.GetPassword(),
		Length:    in.GetLength(),
	}); err != nil {
		return nil, err
	} else {
		return &foodi_user_service.RegisterReply{
			Ok:    true,
			Uid:   reply.GetUid(),
			Token: reply.GetToken(),
		}, nil
	}
}
