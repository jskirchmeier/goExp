package explore

import (
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"strings"

	"gopkg.in/yaml.v3"
)

// Adventure contains everything that is needed for your adventure.
// The initial read from config will contain the start state.
// The same structure will also be used to store state.
// No assumptions on the interface (cli or web) can be made
type Adventure struct {
	Name            string
	Description     string
	Locations       map[int]*Location `yaml:"locations"`
	CurrentLocation int               `yaml:"curLocation"`
	Turns           int               `yaml:"turns"`
}

// NewAdventure loads an adventure definition from the specified (yaml or JSON) file
// return error if there is something wrong with the file
func NewAdventure(fn string) (*Adventure, error) {

	buf, err := ioutil.ReadFile(fn)
	if err != nil {
		return nil, err
	}
	a := &Adventure{}

	err = yaml.Unmarshal(buf, a)
	return a, err
}

func ReadAdventure(in io.Reader) (*Adventure, error) {
	buf, err := ioutil.ReadAll(in)

	if err != nil {
		return nil, err
	}
	a := &Adventure{}

	err = yaml.Unmarshal(buf, a)
	return a, err
}

// starts or welcomes back the player
func (a *Adventure) Start(out io.Writer) {
	t, err := template.New("welcome").Parse(welcome)
	if err != nil {
		// TODO : clean this up
		log.Fatal(err)
	}
	err = t.Execute(out, a)
	if err != nil {
		// TODO : clean this up
		log.Fatal(err)
	}

}

// do the action(s) defined in the command, modify state if needed
// returns a string to display, if end is true then we're done here
func (a *Adventure) Do(command string, out io.Writer) bool {
	a.Turns++
	parts := strings.Fields(strings.ToLower(command))
	return ProcessCommand(a, parts, out)
}
