package builtins_test

import (
	"github.com/jh125486/CSCE4600/Project2/builtins"
	"reflect"
	"testing"
)

func TestShiftCommand(t *testing.T) {
	type args struct {
		args    *[]string
		shiftBy []string
	}
	tests := []struct {
		name    string
		args    args
		want    *[]string
		wantErr bool
	}{
		{
			name: "shift by one",
			args: args{
				args:    &[]string{"arg1", "arg2", "arg3"},
				shiftBy: []string{},
			},
			want:    &[]string{"arg2", "arg3"},
			wantErr: false,
		},
		{
			name: "shift by two",
			args: args{
				args:    &[]string{"arg1", "arg2", "arg3"},
				shiftBy: []string{"2"},
			},
			want:    &[]string{"arg3"},
			wantErr: false,
		},
		{
			name: "error too many args",
			args: args{
				args:    &[]string{"arg1", "arg2", "arg3"},
				shiftBy: []string{"1", "2"},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := builtins.ShiftCommand(tt.args.args, tt.args.shiftBy...)
			if (err != nil) != tt.wantErr {
				t.Errorf("ShiftCommand() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(tt.args.args, tt.want) {
				t.Errorf("ShiftCommand() = %v, want %v", tt.args.args, tt.want)
			}
		})
	}
}
