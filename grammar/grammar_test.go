package grammar

import (
	"fmt"
	"testing"
)

func printC(c *Command) {
	fmt.Print("Command : ")

	if c == nil {
		return
	}
	fmt.Print(c.Verb, "   ", c.Verb2, "  ")
	for _, s := range c.Subjects {
		fmt.Print("  (")
		for a := range s.Adj {
			fmt.Print(a, " ")
		}
		fmt.Print(s.Noun)
		fmt.Print(" )  ")
	}
	fmt.Println()

}

func TestGrammar1(t *testing.T) {
	g := NewGrammar()
	g.AddVerb("move", "go", "take", "open", "unlock")
	g.AddNoun("door", "north", "south", "key")
	g.AddAdjective("red", "shinny", "blue", "rusty")

	c, err := g.Parse("Open red door with rusty key")
	if err != nil {
		t.Error("ORDWRK : ", err)
	}
	printC(c)

	c, err = g.Parse("Move   ")
	//fmt.Print(err)
	printC(c)

	c, err = g.Parse("key the door")
	fmt.Println(err)

	c, err = g.Parse("open the dor")
	fmt.Println(err)

}
