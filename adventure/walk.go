package adventure

import (
	"github.com/jskirchmeier/explore/location"
)

func createWalk() *Adventure {
	a := &Adventure{}
	a.Name = "A walk in the woods"
	a.Description = "Our first (simple) adventure, only movement used"

	start := &location.Base{Name: "at the end of a road, on the edge of the woods"}
	path := &location.Base{Name: "on a tree lined path"}
	clearing := &location.Base{Name: "in a clearing, surrounded by trees"}

	start.AddNorth("wooded path", path)
	path.AddSouth("wooded path", start).AddWest("wooded path", clearing)
	clearing.AddEast("wooded path", path)

	a.CurrentLocation = start
	return a
}

func init() {
	Register("walk", createWalk)
}
