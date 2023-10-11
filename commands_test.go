package commands_test

import (
	"bytes"
	"errors"
	"os"
	"strings"
	"testing"

	"github.com/ElecTwix/commands"
)

type SomeCustomFields struct {
	a int
	b string
}

func TestCommmandCreate(t *testing.T) {
	cmdHandler := commands.New[string, SomeCustomFields](os.Stdout)

	cmd := commands.Command[string, SomeCustomFields]{
		Key: "test123",
		Fn: func(i interface{}) (interface{}, error) {
			return nil, nil
		},
		CustomField: SomeCustomFields{
			a: 15,
			b: "test123",
		},
	}

	err := cmdHandler.AddCmd(cmd)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCommmandFuncExection(t *testing.T) {
	cmdHandler := commands.New[string, SomeCustomFields](os.Stdout)

	testValue, testErr := "testValue123", errors.New("testErr")

	cmd := commands.Command[string, SomeCustomFields]{
		Key: "test123",
		Fn: func(i interface{}) (interface{}, error) {
			return testValue, testErr
		},
		CustomField: SomeCustomFields{
			a: 15,
			b: "test123",
		},
	}

	err := cmdHandler.AddCmd(cmd)
	if err != nil {
		t.Fatal(err)
	}

	returnVal, returnErr := cmdHandler.HandleCmd(cmd.Key, nil)
	if returnVal != testValue {
		t.Fail()
	}

	if returnErr != testErr {
		t.Fail()
	}
}

func TestCommmandPrint(t *testing.T) {
	buff := bytes.NewBuffer([]byte{})
	cmdHandler := commands.New[string, SomeCustomFields](buff)

	testValue, testErr := "testValue123", errors.New("testErr")

	cmd := commands.Command[string, SomeCustomFields]{
		Key: "test123",
		Fn: func(i interface{}) (interface{}, error) {
			return testValue, testErr
		},
		CustomField: SomeCustomFields{
			a: 15,
			b: "test123",
		},
	}

	err := cmdHandler.AddCmd(cmd)
	if err != nil {
		t.Fatal(err)
	}

	err = cmdHandler.PrintAllCommands()
	if err != nil {
		t.Fatal(err)
	}

	printOut := buff.String()
	if !strings.Contains(printOut, cmd.Key) {
		t.Fatal(errors.New("stdout not contains key that created"))
	}

}
