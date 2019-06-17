package command

import (
	"errors"
	"github.com/jskirchmeier/explore/adventure"
)

func move(a *adventure.Adventure, cmds []string) Response {

	// cmd[0] is the move command
	// cmd[1] is the direction
	if len(cmds) < 2 {
		return Response{Err: errors.New("direction to move must be specified")}

	}
	newLoc := a.CurrentLocation.Go(cmds[1])
	if newLoc == nil {
		return Response{Err: errors.New("you cannot go " + cmds[1])}
	}
	a.CurrentLocation = newLoc
	return Response{StateChange: true}
}

func init() {
	Register(HandlerFunc(move), "move", "go", "mv")
}
