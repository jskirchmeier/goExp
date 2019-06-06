package command

import (
	"github.com/jskirchmeier/explore/adventure"
	"io"
	"log"
	"text/template"
)

const (
	welcomeText = `
Welcome {{ if .Turns }}back {{ end }}to {{.Name}}
{{.Description}}

Type help for additional instructions.
Ready...
`
)

var templateWelcome *template.Template

// welcome prints out the starting message for the user
func welcome(a *adventure.Adventure, _ []string, out io.Writer) (changeState bool) {

	err := templateWelcome.Execute(out, a)
	if err != nil {
		// TODO : clean this up
		log.Fatal(err)
	}
	return false
}

func init() {
	Register(Func(welcome), "start", "welcome")
	var err error
	templateWelcome, err = template.New("welcome").Parse(welcomeText)
	if err != nil {
		// TODO : clean this up
		log.Fatal(err)
	}
}
