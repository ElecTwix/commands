package command_test

import (
	"testing"

	"github.com/ElecTwix/commands/pkg/command"
)

func HelperCommandFunction(i *int) (o *int, err error) {
	return nil, nil
}

func TestCreateCommand(t *testing.T) {
	cmd := command.CreateCommand[string, int, int, int]("test", HelperCommandFunction, 1, "test", "test")
	if cmd.Key != "test" {
		t.Error("Key is not test")
	}
	if cmd.CustomField != 1 {
		t.Error("CustomField is not 1")
	}
	if cmd.Helper.Disc != "test" {
		t.Error("Disc is not test")
	}
	if cmd.Helper.Usage != "test" {
		t.Error("Usage is not test")
	}
}
