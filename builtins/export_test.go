package builtins_test

import (
	"github.com/jh125486/CSCE4600/Project2/builtins"
	"os"
	"testing"
)

func TestExportCommand(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "valid export",
			args: args{
				args: []string{"MY_VAR=value"},
			},
			wantErr: false,
		},
		{
			name: "error too many args",
			args: args{
				args: []string{"MY_VAR=value", "OTHER_VAR=value"},
			},
			wantErr: true,
		},
		{
			name: "error invalid format",
			args: args{
				args: []string{"MY_VAR value"},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := builtins.ExportCommand(tt.args.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExportCommand() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				key, value := tt.args.args[0][:strings.Index(tt.args.args[0], "=")], tt.args.args[0][strings.Index(tt.args.args[0], "=")+1:]
				gotValue, exists := os.LookupEnv(key)
				if !exists {
					t.Errorf("ExportCommand() failed to set environment variable: %v", key)
				}
				if value != gotValue {
					t.Errorf("ExportCommand() set environment variable to %v, want %v", gotValue, value)
				}
			}
		})
	}
}
