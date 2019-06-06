package location

//import "io"
//
// Location is a node in a network of connected nodes, could be a room, a cave, a spot in the road
// The only requirement is that it has a way to move to another Location
//
type Location interface {
	// moves in the direction given, if is an invalid direction then nil is returned
	Go(direction string) Location

	// ShortDescription tells the user where they are standing:
	// follows "you are standing "
	// examples : "in a clearing", "at the end of a road"
	// does not describe the things that are at that location, that would be LongDescription
	ShortDescription() string
}
