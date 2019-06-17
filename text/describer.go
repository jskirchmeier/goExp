package text

import (
	"log"
	"strings"
	"text/template"
)

// Describer is the basic mechanism to relate messages to the user
// It has one option; short, true => a one sentence description, false => a longer one including contents
type Describer interface {
	Description(short bool) string
}

type DescriberFunc func(short bool) string

func (h DescriberFunc) Description(short bool) string {
	return h(short)
}

// SimpleDescriber implements a Describer using a single string
type SimpleDescriber string

func (s SimpleDescriber) Description(_ bool) string {
	return string(s)
}

// DescriberHelper implements the Describer interface for
// those that have only simple strings for the two options.
type DescriberHelper struct {
	Short, Long string
}

func (h DescriberHelper) Description(short bool) string {
	if !short && len(h.Long) > 0 {
		return h.Long
	}
	return h.Short
}

// TemplateFunc implements the Describer interface for
// those objects that creates the description from a template
func TemplateFunc(shortTemplate, longTemplate string, o interface{}) Describer {

	return DescriberFunc(func(short bool) string {
		t := shortTemplate
		if !short && len(longTemplate) > 0 {
			t = longTemplate
		}
		tmpl, err := template.New("t").Parse(t)
		if err != nil {
			// TODO : clean this up
			log.Fatal(err)
		}
		buf := &strings.Builder{}
		err = tmpl.Execute(buf, o)
		if err != nil {
			// TODO : clean this up
			log.Fatal(err)
		}
		return buf.String()
	})
}
