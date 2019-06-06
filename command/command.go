package command

import (
	"fmt"
	"github.com/jskirchmeier/explore/adventure"
	"io"
	"log"
)

// Command performs some action using the cmds and adventure instance to do so.
// Writes any response to the user in out (including any error messages).
// Returns if this command changed the adventures state in some manner.
type Command interface {
	Execute(adventure *adventure.Adventure, cmds []string, out io.Writer) (stateChange bool)
}

// Func allows a function to act as a Command
type Func func(*adventure.Adventure, []string, io.Writer) bool

func (f Func) Execute(adventure *adventure.Adventure, cmds []string, out io.Writer) bool {
	return f(adventure, cmds, out)
}

var commands = make(map[string]Command)

func Initialize() {
	// calling this will make all of the commands init functions fire
	// because the package will be pulled in
}

// Register makes the command available for use in the adventure
// A one or more names of how the command can be fired must be supplied.  Names MUST be unique!!!!!
func Register(c Command, names ...string) {
	fmt.Println("Registering commands : ", names)
	if len(names) == 0 {
		log.Fatal("At least one name must be registered with a command")
	}
	for _, n := range names {
		if _, ok := commands[n]; ok {
			log.Fatal("Command names must be unique, " + n + " is already in use")
		}
		commands[n] = c
	}
}

// ProcessCommand finds the proper command and executes it.
// The first element in cmd must be the command.  It is assumed that all of cmd are in lower case.
// Only single work commands are supported, sub commands can be used if supported by the particular command
func Execute(a *adventure.Adventure, cmd []string, out io.Writer) bool {
	if len(cmd) == 0 {
		_, _ = out.Write([]byte("Invalid command.  Type help for assistance."))
		return false
	}

	c, ok := commands[cmd[0]]

	if !ok {
		_, _ = out.Write([]byte(cmd[0] + " is not a recognized command.  Type help for assistance."))
		return false
	}
	return c.Execute(a, cmd, out)

}
