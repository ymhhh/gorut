package builtin

import (
	"errors"
)

var (
	ErrGoroot                   = errors.New("goroot is not set")
	ErrDepsOnlyFlag             = errors.New("you can use only update or delete flag")
	ErrDepsDeleteValueNotExists = errors.New("delete flag not exists")
	ErrCommandNotExists         = errors.New("command not exists")
	ErrTemplateNameNotExists    = errors.New("template name not exists")
)
