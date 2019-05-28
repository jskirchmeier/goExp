package explore

import (
	"fmt"
)

func check(value, desc string) bool {
	if len(value) == 0 {
		fmt.Println(desc, "not set")
		return false
	}
	return true
}

// helper for writing a new exploration configuration file
func (a *Adventure) Validate() {

	// make sure we have required elements
	ok :=
		check(a.Name, "name of exploration adventure") ||
			check(a.Description, "description of exploration adventure")

	if ok {
		fmt.Println("Required elements exist")
	}

	// Locations
	// IDs are unique
	// required values are there
	// paths lead to valid locations
	// TODO: all locations are reachable

	if len(a.Locations) == 0 {
		fmt.Println("No locations are defined")
	}

	// make sure all ids are unique
	m := make(map[int]Identifier)
	for _, l := range a.Locations {
		el, ok := m[l.ID]
		if ok {
			fmt.Println("Duplicate IDs", l.ID, l.Name(), el.Name())
		}
		m[l.ID] = l
		check(l.Name(), fmt.Sprintf("Loc %d : No Name defined\n", l.ID))
		if len(l.Paths) == 0 {
			fmt.Printf("Loc %d : No Paths defined, dead end\n", l.ID)
		}
	}
	// make sure app paths are good
	for _, l := range a.Locations {

		for n, p := range l.Paths {
			if len(p.Description) == 0 {
				fmt.Printf("Loc %d  Path %s : No decription defined\n", l.ID, n)
			}
			_, ok := m[p.LocationID]
			if !ok {
				fmt.Printf("Loc %d  Path %s : Not a valid destination\n", l.ID, n)
			}
		}
	}

}
