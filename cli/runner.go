package cli

import (
	"bufio"
	"os"
	"strings"

	"github.com/jskirchmeier/explore/adventure"
	"github.com/jskirchmeier/explore/command"
)

// runs from the command line
type Runner struct {
}

func (r *Runner) Run(a *adventure.Adventure) {
	reader := bufio.NewReader(os.Stdin)
	// lets get started
	command.Execute(a, []string{"welcome"}, os.Stdout)
	for !a.Exit {
		_, _ = os.Stdout.WriteString("You are standing ")
		_, _ = os.Stdout.WriteString(a.CurrentLocation.ShortDescription())
		_, _ = os.Stdout.WriteString("\n==> ")

		text, _ := reader.ReadString('\n')
		a.Turns++
		parts := strings.Fields(strings.ToLower(text))
		if command.Execute(a, parts, os.Stdout) {
			// this command made a state change, so save the cause
			a.History = append(a.History, text)
		}
	}
}
