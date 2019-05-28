package explore

import (
	"io"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

// Adventure contains everything that is needed for your adventure.
// The initial read from config will contain the start state.
// The same structure will also be used to store state.
// No assumptions on the interface (cli or web) can be made
type Adventure struct {
	Name        string
	Description string
	Locations   map[int]*Location `yaml:"locations"`
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
