package explore

// Describer returns a descriptive string with details
// of the item.  It will be displayed to the user to
// enable them to make decisions on their next actions.
// The description can be multiple lines
// TODO: do we need a simple and detailed descriptions, also any formatting
type Describer interface {
	// flowery phrase to be displayed, can be multiple lines
	Description() string
}

// Identifier allows the object to be stored in a
// centralized item store.  The identity has to be unique across the entire
// system.
type Identifier interface {
	// A system unique identifier, will be used to store and retrieve the item
	Identify() int

	// A short name useful for display.  Does not have to be unique but the user
	// will use this when referring to the item, (ex. open red door) so it
	// should be unique within a particular context.
	// TODO: need to supply rules, will we allow spaces?  If so how many words.  Will depend on grammar parser
	Name() string
}
