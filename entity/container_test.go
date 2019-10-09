package entity

import (
	"testing"
)

func TestContainer(t *testing.T) {
	c := Container{}
	d1 := &Identity{"door", []string{"red", "shiny", "metal"}}
	d2 := &Identity{"door", []string{"blue"}}

	c.Add(d1)
	if c.len() != 1 {
		t.Errorf("Entity not added correctly expected size of 1 got %d", c.len())
	}
	c.Add(d2)
	if c.len() != 2 {
		t.Errorf("Entity not added correctly expected size of 2 got %d", c.len())
	}

	e := c.Query("red", "door")
	if len(e) != 1 {
		t.Errorf("Query not returning expected number of items (1) got %d", c.len())
	}
	if e[0] != d1 {
		t.Errorf("Query not returning expected item %+v", e[0])
	}

	e = c.Query("door")
	if len(e) != 2 {
		t.Errorf("Query not returning expected number of items (2) got %d", c.len())
	}

	e = c.Query("desk")
	if len(e) != 0 {
		t.Errorf("Query not returning expected number of items (0) got %d", c.len())
	}

}
