package commandhandler

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/ElecTwix/commands/pkg/command"
)

var (
	ErrCmdNotExists     error = errors.New("command not exists")
	ErrCmdAlreadyExists error = errors.New("command not exists")
)

// K for key
// C for custom field on command
// ID for input data
// OD for output data
type CommandHandler[K comparable, C, I, O any] struct { // nolint: golint
	cmdMap map[K]command.Command[K, C, I, O]
	writer io.StringWriter
}

func NewHandler[K comparable, C, I, O any](w io.StringWriter) *CommandHandler[K, C, I, O] { // nolint: golint
	cmdMap := make(map[K]command.Command[K, C, I, O], 0)
	if w == nil {
		w = os.Stdout
	}
	return &CommandHandler[K, C, I, O]{cmdMap: cmdMap, writer: w}
}

func (cmdHandler *CommandHandler[K, C, I, O]) HandleCmd(key K, input *I) (*O, error) {
	cmd, ok := cmdHandler.cmdMap[key]
	if !ok {
		return nil, ErrCmdNotExists
	}

	return cmd.Fn(input)
}

func (cmdHandler *CommandHandler[K, C, I, O]) AddCmd(cmd command.Command[K, C, I, O]) error {
	_, ok := cmdHandler.cmdMap[cmd.Key]
	if ok {
		return ErrCmdAlreadyExists
	}

	cmdHandler.cmdMap[cmd.Key] = cmd
	return nil
}

func (cmdHandler *CommandHandler[K, C, I, O]) PrintAllCommands() error {
	for key := range cmdHandler.cmdMap {
		_, err := cmdHandler.writer.WriteString(fmt.Sprintf("Key: %v", key))
		if err != nil {
			return err
		}
	}
	return nil
}

func (cmdHandler *CommandHandler[K, C, I, O]) GetAllCommands() []command.Command[K, C, I, O] {
	cmdList := make([]command.Command[K, C, I, O], 0, len(cmdHandler.cmdMap))
	for _, cmd := range cmdHandler.cmdMap {
		cmdList = append(cmdList, cmd)
	}
	return cmdList
}

func (cmdHandler *CommandHandler[K, C, I, O]) PrintCommand(key K) error {
	cmd, ok := cmdHandler.cmdMap[key]
	if !ok {
		return ErrCmdAlreadyExists
	}
	_, err := cmdHandler.writer.WriteString(fmt.Sprintf("Key: %v, Disc: %v, Usage: %v", key, cmd.Helper.Disc, cmd.Helper.Usage))
	if err != nil {
		return err
	}
	return nil
}
