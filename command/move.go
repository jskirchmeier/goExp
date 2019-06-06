package command

import (
	"github.com/jskirchmeier/explore/adventure"
	"io"
)

func move(a *adventure.Adventure, cmd []string, out io.Writer) (changeState bool) {

	// cmd[0] is the move command
	// cmd[1] is the direction
	if len(cmd) < 2 {
		_, _ = out.Write([]byte("You must specify what direction to wish to go."))
		return false

	}

	newLoc := a.CurrentLocation.Go(cmd[1])
	if newLoc == nil {
		_, _ = out.Write([]byte("You cannot go " + cmd[1]))
		return false
	}
	a.CurrentLocation = newLoc
	return true
}

func init() {
	Register(Func(move), "move", "go", "mv")
}
