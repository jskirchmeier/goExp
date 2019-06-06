package location

// Base is a helper implementation of the Location interface
// it contains a map of paths to other locations.
// Uses common direction, and their alias north (n), south (s), east (e), west (w)
// Additionally if no direction is given and there is only one direction to go in
// then that one will be used
type Base struct {
	directions map[string]Path
	Name       string
}

func (l *Base) Go(direction string) Location {
	if len(direction) == 0 {
		if len(l.directions) == 1 {
			for _, p := range l.directions {
				return p.Location
			}
		}
	}
	p, ok := l.directions[direction]
	if ok {
		return p.Location
	}
	return nil
}

func (l *Base) ShortDescription() string {
	return l.Name
}

func (l *Base) addDirection(dir, desc string, loc Location) {
	if l.directions == nil {
		l.directions = make(map[string]Path)
	}
	l.directions[dir] = Path{desc, loc}
}

func (l *Base) AddNorth(description string, loc Location) *Base {
	l.addDirection("north", description, loc)
	l.addDirection("n", description, loc)
	return l
}

func (l *Base) AddSouth(description string, loc Location) *Base {
	l.addDirection("south", description, loc)
	l.addDirection("s", description, loc)
	return l
}

func (l *Base) AddEast(description string, loc Location) *Base {
	l.addDirection("east", description, loc)
	l.addDirection("e", description, loc)
	return l
}

func (l *Base) AddWest(description string, loc Location) *Base {
	l.addDirection("west", description, loc)
	l.addDirection("w", description, loc)
	return l
}

// Describes where a path leads to
type Path struct {
	Description string
	Location    Location
}
