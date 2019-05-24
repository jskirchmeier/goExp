package simple

import "github.com/jskirchmeier/goExp"

// abstract location to save connection information
// does NOT fully implement the Location interface
// TODO: 1) do we need an Add function, won't we be always creating from config 2) if this is all there is then punt

type AbstractLocation struct {
	Connections map[string]int
}

func (loc *AbstractLocation) Add(d string, l goExp.Location) {
	if loc.Connections == nil {
		loc.Connections = make(map[string]int)
	}
	loc.Connections[d] = l.Id()
}
