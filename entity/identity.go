package entity

import (
	"strings"
)

// Identity implements the Nameable interface and can be used
// to compose more complex entities without needing to reimplement the functionality
//
// Name is the main identifier for the item
// optional Adjectives further describe the item
// an Identity must be unique in the context that contains it,
// if it identifies an item than can be picked up then it must be unique system wide
type Identity struct {
	Name       string
	Adjectives []string
}

// Does the Identity match you?
func (i *Identity) AreYou(words ...string) bool {
	foundName := false

	for _, w := range words {
		if w == i.Name {
			foundName = true
		} else {
			if !stringInSlice(i.Adjectives, w) {
				return false
			}
		}
	}
	return foundName

}

func stringInSlice(slc []string, str string) bool {
	for _, s := range slc {
		if s == str {
			return true
		}
	}
	return false
}

func (i Identity) String() string {
	buf := strings.Builder{}
	for _, a := range i.Adjectives {
		buf.WriteString(a)
		buf.WriteString(" ")
	}
	buf.WriteString(i.Name)

	return buf.String()
}
