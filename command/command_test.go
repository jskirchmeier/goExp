package command

import (
	"fmt"
	"strconv"
	"testing"
)

func TestImplCommander_Execute(t *testing.T) {
	var one, two, three int
	c := &implCommander{make(map[string]Command)}

	c.commands["one"] = CommandFunc(func(ctx *Context, params ...string) error {
		one++
		return nil
	})

	c.commands["two"] = CommandFunc(func(ctx *Context, params ...string) error {
		two++
		return nil
	})

	c.commands["three"] = CommandFunc(func(ctx *Context, params ...string) error {
		if len(params) != 1 {
			return fmt.Errorf("required one paramater, recieved %d", len(params))
		}
		cnt, err := strconv.Atoi(params[0])
		if err != nil {
			return err
		}
		three += cnt
		return nil
	})

	err := c.Execute(nil, "unknowncommand")
	if err == nil {
		t.Errorf("Expected error for unknown command")
	}

	err = c.Execute(nil, "one")
	if err != nil {
		t.Errorf("Unexpected error for command one")
	}
	if one != 1 {
		t.Errorf("Command one not executed properly expexted value of 1 got %d", one)
	}

}
