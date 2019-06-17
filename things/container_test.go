package things

import (
	"fmt"
	"testing"
)

type testi1 string

func (i testi1) AreYou(names ...string) bool {
	if len(names) == 1 {
		return names[0] == "door"
	}

	for _, n := range names {
		switch n {
		case "door":
		// nothing
		default:
			return n == string(i)
		}
	}
	return false
}

func Test(t *testing.T) {
	red := testi1("red")
	blue := testi1("blue")
	green := testi1("green")

	c := Container{}
	c.Add(red, blue, green)
	var doors []Nameable

	doors = c.Get([]string{"red", "door"})
	if len(doors) != 1 {
		t.Errorf("Did not return red door")
	}

	doors = c.Get([]string{"door", "red"})
	if len(doors) != 1 {
		t.Errorf("Did not return red door")
	}

	if doors[0] != red {
		t.Errorf("Did not return red door")
	}
	_, ok := doors[0].(testi1)
	fmt.Println(ok)

	doors = c.Get([]string{"door"})
	if len(doors) != 3 {
		t.Errorf("Did not return doors")
	}

}
