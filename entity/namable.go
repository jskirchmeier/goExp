package entity

// thr various interfaces needed by the entities

// Nameable is something that can be selected by name
// the name has to be unique by context (only one red door in a room)
// if the item is  moveable then it has to be unique system wide
type Nameable interface {
	// Does the name match you?
	// There should be a 'primary' name for the entity ('door') and
	// zero or more optional adjectives associated with it ('red shiny metal door')
	// For the method to return true at least the main name has to be mentioned.
	// none of the attributes are required although specifying one that does not belong if a fail
	// examples : using the identity 'red shiny metal door'
	// returning True : 'door', 'red door', 'shiny door', 'red metal door'
	// order does not matter : metal red door
	// returning False: 'cat', 'blue door', 'red shiny metal dented door', 'red shiny metal cat'
	AreYou(name ...string) bool
}
