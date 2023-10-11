package commands

import (
	"fmt"
	"io"
	"os"

	"github.com/ElecTwix/commands/pkg/errcmd"
)

// K for key
// C for custom field on command
type CommandHandler[K comparable, C any] struct {
	cmdMap map[K]Command[K, C]
	writer io.StringWriter
}

type HelperHandler struct {
	Disc  string
	Usage string
}

type Command[K comparable, C any] struct {
	Key K

	Fn     CommandFunc
	Helper HelperHandler

	CustomField C
}

type CommandFunc func(interface{}) (interface{}, error)

func New[K comparable, C any](w io.StringWriter) *CommandHandler[K, C] {
	cmdMap := make(map[K]Command[K, C], 0)
	if w == nil {
		w = os.Stdout
	}
	return &CommandHandler[K, C]{cmdMap: cmdMap, writer: w}
}

func (cmdHandler *CommandHandler[K, C]) HandleCmd(key K, args interface{}) (interface{}, error) {
	cmd, ok := cmdHandler.cmdMap[key]
	if !ok {
		return nil, errcmd.ErrCmdNotExists
	}

	return cmd.Fn(args)
}

func (cmdHandler *CommandHandler[K, C]) AddCmd(cmd Command[K, C]) error {
	_, ok := cmdHandler.cmdMap[cmd.Key]
	if ok {
		return errcmd.ErrCmdAlreadyExists
	}

	cmdHandler.cmdMap[cmd.Key] = cmd
	return nil
}

func (cmdHandler *CommandHandler[K, C]) PrintAllCommands() error {
	for key := range cmdHandler.cmdMap {
		_, err := cmdHandler.writer.WriteString(fmt.Sprintf("Key: %v", key))
		if err != nil {
			return err
		}
	}
	return nil
}
