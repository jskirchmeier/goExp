package location

import "github.com/jskirchmeier/explore/text"

// Path is a one way connection between locations
// could be a trail, road, hallway
// or a door which could have state, open, closed, locked, etc.
type Pathway interface {
	// short and long descriptions are the same
	text.Describer

	// where does the path lead to
	LeadsTo() Location
}

// Location is a node in a network of connected nodes, could be a room, a cave, a spot in the road
// The only requirement is that it has a way to move to another Location
//
type Location interface {
	// moves in the direction given, if is an invalid direction then nil is returned
	Go(direction string) Location

	// Describes where you're at
	text.Describer
}
