package command

import (
	"log"
	"strings"
	"text/template"

	"github.com/jskirchmeier/explore/adventure"
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
func welcome(a *adventure.Adventure, _ []string) Response {

	buf := &strings.Builder{}
	err := templateWelcome.Execute(buf, a)
	if err != nil {
		// TODO : clean this up
		log.Fatal(err)
	}
	return Response{}
}

func init() {
	Register(HandlerFunc(welcome), "start", "welcome")
	var err error
	templateWelcome, err = template.New("welcome").Parse(welcomeText)
	if err != nil {
		// TODO : clean this up
		log.Fatal(err)
	}
}
