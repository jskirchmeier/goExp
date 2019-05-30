package cli

import (
	"bufio"
	"github.com/jskirchmeier/explore"
	"os"
)

// runs from the command line
type Runner struct {
}

func (r *Runner) Run(a *explore.Adventure) {
	reader := bufio.NewReader(os.Stdin)
	a.Start(os.Stdout)
	for {
		text, _ := reader.ReadString('\n')
		if a.Do(text, os.Stdout) {
			break
		}
	}
}
