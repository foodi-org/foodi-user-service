package articlelogic

import (
	"context"
	"github.com/foodi-org/foodi-user-service/github.com/foodi-org/foodi-user-service"
	"github.com/foodi-org/foodi-user-service/internal/config"
	"github.com/foodi-org/foodi-user-service/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
	"reflect"
	"testing"
)

func TestAddCommentLogic_AddComment(t *testing.T) {
	type fields struct {
		ctx    context.Context
		svcCtx *svc.ServiceContext
		Logger logx.Logger
	}
	type args struct {
		in *foodi_user_service.AddCommentRequest
	}
	if err := svc.NewServiceContext(config.ServConf(), "/Users/edy/Documents/foodi-org/foodi-user-service", "dev.yaml"); err != nil {
		t.Fatal(err)
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *foodi_user_service.OKReply
		wantErr bool
	}{
		{
			name: "test",
			fields: fields{
				ctx:    context.TODO(),
				svcCtx: svc.Svc(),
				Logger: nil,
			},
			args: args{in: &foodi_user_service.AddCommentRequest{
				Uid:         10011,
				ArticleID:   1021912,
				Content:     "this is a test content",
				IsAnonymous: true,
			}},
			wantErr: false,
			want:    &foodi_user_service.OKReply{Ok: true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &AddCommentLogic{
				ctx:    tt.fields.ctx,
				svcCtx: tt.fields.svcCtx,
				Logger: tt.fields.Logger,
			}
			got, err := l.AddComment(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddComment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddComment() got = %v, want %v", got, tt.want)
			}
		})
	}
}
