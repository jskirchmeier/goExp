package adventure

import (
	"fmt"
	"log"

	"github.com/jskirchmeier/explore/location"
)

type Creator func() *Adventure

var creators = make(map[string]Creator)

func Register(name string, c Creator) {
	if _, ok := creators[name]; ok {
		log.Fatal("Adventure names need to be unique, " + name + " is already in use")
	}
	creators[name] = c
}

func New(name string) *Adventure {
	fmt.Printf("New exploration : %s, %d availiable : %+v\n", name, len(creators), creators)
	c, ok := creators[name]
	if !ok {
		return nil
	}
	return c()
}

type Adventure struct {
	Name        string
	Description string

	//	Locations       map[int]*Location
	CurrentLocation location.Location
	Turns           int
	Exit            bool
	History         []string
}
