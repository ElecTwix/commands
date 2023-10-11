package errcmd

import "errors"

var (
	ErrCmdNotExists     error = errors.New("Command not exists")
	ErrCmdAlreadyExists error = errors.New("Command not exists")
)
