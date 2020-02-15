package components

import (
	"go.7yes.com/j7f/components/errors"
)

const (
	OK      = 0
	SUCCESS = "success"
)

type Error struct {
	code int64
	err  string
	errors.Error
}
