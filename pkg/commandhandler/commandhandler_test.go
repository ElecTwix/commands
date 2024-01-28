package commandhandler_test

import (
	"os"
	"testing"

	"github.com/ElecTwix/commands/pkg/command"
	"github.com/ElecTwix/commands/pkg/commandhandler"
)

func CreateSimpleHandler(t *testing.T) *commandhandler.CommandHandler[string, int, int, int] {
	handler := commandhandler.NewHandler[string, int, int, int](os.Stdout)
	if handler == nil {
		t.Error("Handler is nil")
	}
	return handler
}

func CreateCustomKeyedHandler[T comparable](t *testing.T) *commandhandler.CommandHandler[T, int, int, int] {
	handler := commandhandler.NewHandler[T, int, int, int](os.Stdout)
	if handler == nil {
		t.Error("Handler is nil")
	}
	return handler
}

func HelperCommandFunction(i *int) (o *int, err error) {
	return nil, nil
}

func TestNewHandler(t *testing.T) {
	handler := CreateSimpleHandler(t)
	if handler == nil {
		t.Error("Handler is nil")
	}
}

func TestCommandHandler_AddCmd(t *testing.T) {
	cmd := command.CreateCommand[string, int, int, int]("test", HelperCommandFunction, 1, "test", "test")

	handler := CreateSimpleHandler(t)
	err := handler.AddCmd(cmd)
	if err != nil {
		t.Error(err)
	}
}

func TestCommandHandler_HandleCmd(t *testing.T) {
	cmd := command.CreateCommand[string, int, int, int]("test", HelperCommandFunction, 1, "test", "test")
	handler := CreateSimpleHandler(t)
	err := handler.AddCmd(cmd)
	if err != nil {
		t.Error(err)
	}
	_, err = handler.HandleCmd("test", nil)
	if err != nil {
		t.Error(err)
	}
}

func TestCommandHandler_PrintAllCommands(t *testing.T) {
	cmd := command.CreateCommand[string, int, int, int]("test", HelperCommandFunction, 1, "test", "test")
	handler := CreateSimpleHandler(t)
	err := handler.AddCmd(cmd)
	if err != nil {
		t.Error(err)
	}
	err = handler.PrintAllCommands()
	if err != nil {
		t.Error(err)
	}
}

func TestCommandHandler_GetAllCommands(t *testing.T) {
	totalCmd := 10
	cmdArr := make([]command.Command[int, int, int, int], 10)

	for i := 0; i < totalCmd; i++ {
		cmdArr[i] = command.CreateCommand[int, int, int, int](i, HelperCommandFunction, 1, "test", "test")
	}

	handler := CreateCustomKeyedHandler[int](t)

	for i := 0; i < totalCmd; i++ {
		err := handler.AddCmd(cmdArr[i])
		if err != nil {
			t.Error(err)
		}
	}

	commands := handler.GetAllCommands()
	if commands == nil {
		t.Error("Commands is nil")
	}

	if len(commands) != totalCmd {
		t.Error("Commands length is not equal to totalCmd")
	}
}
