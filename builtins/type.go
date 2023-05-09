package builtins

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

var (
	ErrTypeInvalidArgCount = errors.New("invalid argument count")
)

func TypeCommand(args ...string) (string, error) {
	if len(args) != 1 {
		return "", fmt.Errorf("%w: expected one argument (command name)", ErrTypeInvalidArgCount)
	}

	cmd := args[0]

	if IsBuiltin(cmd) {
		return fmt.Sprintf("%s is a shell builtin", cmd), nil
	}

	path, err := exec.LookPath(cmd)
	if err != nil {
		return "", fmt.Errorf("%s: command not found", cmd)
	}

	return fmt.Sprintf("%s is %s", cmd, path), nil
}

func IsBuiltin(cmd string) bool {
	builtins := []string{"cd", "type", "shift", "export","exit","env"}
	for _, b := range builtins {
		if b == cmd {
			return true
		}
	}
	return false
}
