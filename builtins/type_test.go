package builtins_test

import (
	"github.com/jh125486/CSCE4600/Project2/builtins"
	"testing"
)

func TestTypeCommand(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "error too many args",
			args: args{
				args: []string{"ls", "cd"},
			},
			wantErr: true,
		},
		{
			name: "builtin command",
			args: args{
				args: []string{"cd"},
			},
			want: "cd is a shell builtin",
		},
		{
			name: "external command",
			args: args{
				args: []string{"ls"},
			},
			want: "ls is /bin/ls",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := builtins.TypeCommand(tt.args.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("TypeCommand() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("TypeCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}
