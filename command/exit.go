package command

import (
	"github.com/jskirchmeier/explore/adventure"
)

// exit is the command to terminate the adventure
func exit(_ *adventure.Adventure, _ []string) Response {
	return Response{Exit: true}
	// no response needed, the runner will handle that
}

func init() {
	Register(HandlerFunc(exit), "exit", "bye", "end")
}
