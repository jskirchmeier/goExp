package goExp

// everything in the place has to be a thing
// used to look up items

type Thing interface {
	// ID must be unique within the system
	Id() int
	Name() string
	// flowery phrase to be displayed, can be multiple lines
	Description() string
}

// a location is where you are or want to be
// is a Thing
type Location interface {
	Thing

	// travel in the indicated direction
	// returns the id of the location that is in that direction
	// returns 0 if nothing there
	Go(direction string) int
}
