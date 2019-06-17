package adventure

import (
	"github.com/jskirchmeier/explore/location"
)

type Adventure struct {
	Name        string
	Description string

	//	Locations       map[int]*Location
	CurrentLocation location.Location
	Turns           int
	Exit            bool
	History         []string
}
