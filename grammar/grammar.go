package grammar

import "errors"

const (
	NoWord = -1

	UK = iota // unknown
	TW        // Throw Away
	VB        // verb
	NN        // noun
	JJ        // adjective
	PP        // preposition
)

var reserved = map[string]int{
	"the":  TW,
	"a":    TW,
	"an":   TW,
	"and":  TW,
	"on":   PP,
	"with": PP,
	"in":   PP,
}

type Identifier interface {
	// is there enough information to identify this object
	// The noun describing it (including synonyms if any) MUST match
	// If adjectives are passed they MUST match an adjective on the object
	// Not all, or any, adjectives need to be present to identify an object
	// order of the adjectives do not matter
	// example:  noun: door, adj: shiny red metal
	// true:
	//	door
	//  red door
	//  shiny door
	// false:
	//  key
	//  blue door
	//  red shiny wooden door
	//  red shiny metal short door
	AreYou(noun string, adj ...string) bool
}

type Subject struct {
	Noun string
	Adj  map[string]bool
}

type Command struct {
	Verb     string
	Verb2    string
	Subjects []*Subject
}

// a simple grammar,
// verb [verb] subject [prep subject]
// some verbs support a secondary verb
// where a subject is : [adj...] noun
// Only one verb which must the first one
// At lease one subject
// and optional second subject with a connecting preposition
// a subject can have zero or more adjective
// the combination of noun and adjectives MUST be unique within the context the item is contained in
// NOTE: if the item is moveable it must be unique system wide
// all prepositions are treated the same (even if it sounds weird)
// examples:
// go north
// open door
// unlock red door with green key
//

// verb [verb] [adj...] noun [prep] [adj...] [noun]
// items need to register adjectives and nouns
// command need to register verbs
type Grammar map[string]int

func NewGrammar() Grammar {
	g := Grammar{}
	for w, t := range reserved {
		g[w] = t
	}
	return g
}

func (g Grammar) AddVerb(verbs ...string) {
	for _, s := range verbs {
		g[s] = VB
	}
}

func (g Grammar) AddNoun(nouns ...string) {
	for _, n := range nouns {
		g[n] = NN
	}
}

func (g Grammar) AddAdjective(adjs ...string) {
	for _, a := range adjs {
		g[a] = JJ
	}
}

func (g Grammar) parseSubject(q *Queue) (*Subject, error) {
	s := &Subject{Adj: make(map[string]bool)}
	for q.More() {
		v, t := q.Peek()
		switch t {
		case PP:
			return s, nil

		case NN:
			if len(s.Noun) > 0 {
				return nil, errors.New("more than one noun in subject " + s.Noun + ", amd " + v)
			}
			s.Noun = v
			q.Next()

		case JJ:
			has := s.Adj[v]
			if has {
				return nil, errors.New("adjective used more than once : " + v)
			}
			s.Adj[v] = true
			q.Next()
		default:
			return nil, errors.New(v + " not recognized as p[art of subject")
		}
	}
	return s, nil
}

func (g Grammar) Parse(line string) (*Command, error) {
	q := NewQueue(line, g)

	// first word needs to be a verb
	v, t := q.Next()
	if t != VB {
		return nil, errors.New("needs to start with a verb, found " + v)
	}
	c := &Command{Verb: v}
	v, t = q.Peek()
	if t == VB {
		// sub verb
		c.Verb2 = v
		q.Next()
	}

	for q.More() {
		_, t := q.Peek()
		if t == PP {
			// a new subject
			q.Next()
		}
		s, err := g.parseSubject(q)
		if err != nil {
			return nil, err
		}
		c.Subjects = append(c.Subjects, s)
	}
	return c, nil
}
