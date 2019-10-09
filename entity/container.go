package entity

// Container holds some number of Nameable entities
// supports the addition, removal, and query of those entities
// TODO: handle non-removal of entities
// TODO: will have to be reworked to handle other attributes of an entity but for now ...
type Container struct {
	items []Nameable
}

func (c *Container) len() int {
	return len(c.items)
}

func (c *Container) Query(name ...string) (resp []Nameable) {
	for _, n := range c.items {
		if n.AreYou(name...) {
			resp = append(resp, n)
		}
	}
	return
}

func (c *Container) Add(n Nameable) {
	c.items = append(c.items, n)
}

func (c *Container) Remove(n Nameable) {
	for i, nn := range c.items {
		if n == nn {
			copy(c.items[i:], c.items[i+1:])
			c.items[len(c.items)-1] = nil // or the zero value of T
			c.items = c.items[:len(c.items)-1]
			return
		}
	}

}

// Container contains items, these items have behaviours that can be searched for

// type DContainer struct {
// 	items []Nameable
// }
//
// func (c *DContainer) Add(items ...Nameable) {
// 	c.items = append(c.items, items...)
// }
//
// func (c DContainer) Get(names []string) (items []Nameable) {
// 	for _, i := range c.items {
// 		if i.AreYou(names...) {
// 			items = append(items, i)
// 		}
// 	}
// 	return
// }
