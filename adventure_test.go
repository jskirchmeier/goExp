package explore

import (
	"fmt"
	"github.com/jskirchmeier/explore/simple"
	"gopkg.in/yaml.v3"
	"testing"
)

const a1 = `
locations:
    101:
        id: 101
        name: clearing
        n:
            desc: path to the north
            locID: 102
    102:
        id: 102
        name: fork in the path
        s:
            desc: path to the south
            locID: 101

`

func TestUnmarshal(t *testing.T) {
	a := &Adventure{}
	err := yaml.Unmarshal([]byte(a1), a)
	if err != nil {
		t.Error("Unable to marshall", err)
	}
}

func TestMarshal(t *testing.T) {
	a := Adventure{}
	a.Name = "A Sunday walk"
	a.Description = "our first simple adventure"
	a.Locations = make(map[int]*Location)
	l := &Location{Identity: simple.Identity{ID: 101, Nombre: "clearing"}}
	l.Paths = make(map[string]*Path)
	l.Paths["n"] = &Path{Description: "path to the north", LocationID: 102}
	a.Locations[l.ID] = l

	l = &Location{Identity: simple.Identity{ID: 102, Nombre: "fork in the path"}}
	l.Paths = make(map[string]*Path)
	l.Paths["s"] = &Path{Description: "path to the south", LocationID: 101}
	a.Locations[l.ID] = l

	buf, err := yaml.Marshal(a)

	if err != nil {
		t.Error("Unable to marshall", err)
	}

	fmt.Println(string(buf))

}
