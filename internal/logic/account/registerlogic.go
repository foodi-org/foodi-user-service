package accountlogic

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	_const "github.com/foodi-org/foodi-user-service/internal/const"
	"github.com/foodi-org/foodi-user-service/internal/pkg/snowflake"
	"github.com/foodi-org/foodi-user-service/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"math/rand"
	"time"

	"github.com/foodi-org/foodi-user-service/github.com/foodi-org/foodi-user-service"
	"github.com/foodi-org/foodi-user-service/internal/svc"

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
	var (
		user *model.UserInfo
		res  sql.Result
		aid  int64
		err  error
	)

	// phone exist
	_, err = l.svcCtx.AccountModel.FindWithPhone(l.ctx, in.GetPhone())
	switch {
	case errors.Is(err, sqlx.ErrNotFound):
		break
	case err == nil:
		return &foodi_user_service.RegisterReply{}, errors.New("该手机号已注册")
	default:
		return &foodi_user_service.RegisterReply{}, err
	}

	err = l.svcCtx.DB.Transact(func(session sqlx.Session) error {
		res, err = l.svcCtx.AccountModel.TransInsert(l.ctx, session, &model.AccountInfo{
			Type:          sql.NullString{String: in.GetType().String(), Valid: true},
			NikeName:      l.generateNikeName(),
			Image:         sql.NullString{String: l.defaultImage(), Valid: true},
			Phone:         sql.NullInt64{Int64: in.GetPhone(), Valid: true},
			Password:      sql.NullString{String: in.GetPassword(), Valid: len(in.GetPassword()) > 0},
			FirstRegister: sql.NullString{String: time.Now().Format(time.DateTime), Valid: true},
			LastLogin:     sql.NullString{String: time.Now().Format(time.DateTime), Valid: true},
		})
		if err != nil {
			return err
		}

		aid, err = res.LastInsertId()
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
	})
	if err != nil {
		return nil, err
	}
	user, err = l.svcCtx.UserInfoModel.TakeWithAid(l.ctx, aid)
	return &foodi_user_service.RegisterReply{
		Ok:  true,
		Uid: user.Uid.Int64,
	}, nil
}
func (l *RegisterLogic) generateNikeName() sql.NullString {
	head := "袋熊"
	now := time.Now().Unix()
	return sql.NullString{String: head + fmt.Sprintf("%d", now), Valid: true}
}

func (l *RegisterLogic) defaultImage() string {
	idx := rand.Intn(5)
	switch idx {
	case 0:
		return "default 0 image address of oss"
	case 1:
		return "default 1 image address of oss"
	case 2:
		return "default 2 image address of oss"
	case 3:
		return "default 3 image address of oss"
	case 4:
		return "default 4 image address of oss"
	}
	return "final image address of oss"
}
