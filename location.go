package explore

import "github.com/jskirchmeier/explore/simple"

// A place and its contents, could be a room, a cave, a spot in the road
// the name of the location will be used as : You are in a ........
type Location struct {
	simple.Identity `yaml:",inline"`

	// map of direction (up, down, north, south, etc.) to the path
	// the grammar parser may allow multiple terms for the sane direction (n, north) but
	// will have to translate to a single 'main' term
	Paths map[string]*Path `yaml:"paths,inline"`
}

// Describes where a path leads to,
// When two locations (A,B) are connected there will be two paths (A => B and B => A)
type Path struct {
	Description string `yaml:"desc"`
	LocationID  int    `yaml:"locID"`
}
