package command

import (
	"errors"
	"fmt"
	"log"

	"github.com/jskirchmeier/explore/adventure"
)

// Response contains all of the possible elements
// that a command handler can return
// StateChange - the command changed the state of the adventure
// Exit - we are done, say goodbye and terminate
// Text - Explain to the user that happened (optional)
// Err - Something happened, unable to execute command
// TODO : do we need a separate Text and Err fields?
type Response struct {
	StateChange bool
	Exit        bool
	Text        string
	Err         error
}

// Handler performs some action using the cmds and adventure instance to do so.
// the output type is what kind of text we want, the options are html, ansi (colored console), and plain
type Handler interface {
	Execute(adventure *adventure.Adventure, cmds []string) Response
}

// HandlerFunc allows a function to act as a Handler
type HandlerFunc func(adventure *adventure.Adventure, cmds []string) Response

func (f HandlerFunc) Execute(adventure *adventure.Adventure, cmds []string) Response {
	return f(adventure, cmds)
}

var commands = make(map[string]Handler)

// Register makes the command available for use in the adventure
// A one or more names of how the command can be fired must be supplied.  Names MUST be unique!!!!!
func Register(h Handler, names ...string) {
	fmt.Println("Registering commands : ", names)
	if len(names) == 0 {
		log.Fatal("At least one name must be registered with a command")
	}
	for _, n := range names {
		if _, ok := commands[n]; ok {
			log.Fatal("Handler names must be unique, " + n + " is already in use")
		}
		commands[n] = h
	}
}

// Execute finds the proper command handler and executes it.
// The first element in cmd must be the command.  It is assumed that all of cmd are in lower case.
// Only single work commands are supported, sub commands can be used if supported by the particular command
func Execute(adventure *adventure.Adventure, cmds []string) Response {
	if len(cmds) == 0 {
		return Response{Err: errors.New("no command found")}
	}

	c, ok := commands[cmds[0]]

	if !ok {
		return Response{Err: errors.New("not a recognized command : " + cmds[0])}
	}
	return c.Execute(adventure, cmds)
}
