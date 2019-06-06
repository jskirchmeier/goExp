package command

import (
	"github.com/jskirchmeier/explore/adventure"
	"io"
)

// Exit is the command to terminate the adventure

func exit(adventure *adventure.Adventure, _ []string, out io.Writer) (changeState bool) {

	adventure.Exit = true
	_, _ = out.Write([]byte("Thank you for exploring with us.\nGoodbye...\n"))
	// not really a change state, won't be saved anyway
	return false
}

func init() {
	Register(Func(exit), "exit", "bye", "end")
}
