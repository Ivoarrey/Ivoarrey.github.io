package builtins

import (
	"errors"
	"os"
)

var (
	ErrExportInvalidArgCount = errors.New("invalid argument count")
	ErrExportInvalidFormat   = errors.New("invalid format")
)

func ExportCommand(args ...string) error {
	if len(args) != 1 {
		return ErrExportInvalidArgCount
	}

	exportArg := args[0]

	kv := strings.SplitN(exportArg, "=", 2)
	if len(kv) != 2 {
		return ErrExportInvalidFormat
	}

	key, value := kv[0], kv[1]
	return os.Setenv(key, value)
}
