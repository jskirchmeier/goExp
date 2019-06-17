package things

// Nameable is something that can be selected by name
// the name has to be unique by context (only one red door in a room)
// if the item is  moveable (sword) then it has to be unique system wide
type Nameable interface {
	// Does the name match you?
	// if there are multiple parts the item is responsible to determine if
	// there is enough there to identify itself
	// For example the item is a red door, it might except {"red", "door"}, or just {"door"}
	// in the second case the location may return more than one item that accepts "door"
	// and would return an error at that level
	AreYou(name ...string) bool
}

// Container contains items, these items have behaviours that can be searched for

type Container struct {
	items []Nameable
}

func (c *Container) Add(items ...Nameable) {
	c.items = append(c.items, items...)
}

func (c Container) Get(names []string) (items []Nameable) {
	for _, i := range c.items {
		if i.AreYou(names...) {
			items = append(items, i)
		}
	}
	return
}
