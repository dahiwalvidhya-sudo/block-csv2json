package validate

import (
	"errors"
	"strconv"
)

func Int(value string, field string) (int, error) {
	i, err := strconv.Atoi(value)
	if err != nil {
		return 0, errors.New(field + " must be an integer")
	}
	return i, nil
}
