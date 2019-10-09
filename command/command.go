package command

import (
	"fmt"
	"strings"

	"github.com/jskirchmeier/explore/entity"
)

// Context gives the command all the information it needs to perform its job
// TODO move into runner package when ready, has to add methods to allow commands to modify state
// TODO how do we return response to user
type Context struct {
	Turn         int               // what turn are we on
	User         string            // wio is playing
	CurLocation  string            // where are we : TODO change to location object once that is completed
	CurContainer *entity.Container // what are we looking at (a location is also a container so it could be both)

}

type Command interface {
	Execute(ctx *Context, params ...string) error
}

type CommandFunc func(ctx *Context, params ...string) error

func (cf CommandFunc) Execute(ctx *Context, params ...string) error {
	return cf(ctx, params...)
}

// Commander is the entry point to the commands
type Commander interface {
	Execute(ctx *Context, cmd string) error
}

func NewCommander() Commander {
	c := &implCommander{make(map[string]Command)}

	return c
}

type implCommander struct {
	commands map[string]Command
}

func (c *implCommander) Execute(ctx *Context, cmd string) error {
	words := strings.Fields(cmd)
	if len(words) == 0 {
		return fmt.Errorf("command line is blank")
	}
	command, ok := c.commands[words[0]]
	if !ok {
		return fmt.Errorf("unknown command (%s)", words[0])
	}
	return command.Execute(ctx, words[1:]...)
}
