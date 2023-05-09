package builtins

import (
	"errors"
	"strconv"
)

var (
	ErrShiftInvalidArgCount = errors.New("invalid argument count")
)

func ShiftCommand(args *[]string, shiftBy ...string) error {
	if len(shiftBy) > 1 {
		return ErrShiftInvalidArgCount
	}

	shift := 1
	if len(shiftBy) == 1 {
		var err error
		shift, err = strconv.Atoi(shiftBy[0])
		if err != nil {
			return ErrShiftInvalidArgCount
		}
	}

	if len(*args) < shift {
		return ErrShiftInvalidArgCount
	}

	*args = (*args)[shift:]
	return nil
}
