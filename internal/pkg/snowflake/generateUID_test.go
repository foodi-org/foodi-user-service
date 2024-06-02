package snowflake

import (
	"fmt"
	"testing"
)

func TestRuleUID(t *testing.T) {
	type args struct {
		machineLen int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "TestRuleUID",
			args: args{machineLen: 900},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RuleUID(tt.args.machineLen)
			fmt.Println(got)
			if got == tt.want {
				t.Errorf("RuleUID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSnowUID(t *testing.T) {
	tests := []struct {
		name    string
		want    int64
		wantErr bool
	}{
		{name: "TestSnowUID", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SnowUID()
			if (err != nil) != tt.wantErr {
				t.Errorf("SnowUID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == tt.want {
				t.Errorf("SnowUID() got = %v, want %v", got, tt.want)
			}
		})
	}
}
