package grammar

import "strings"

type Queue struct {
	words   []string
	idx     int
	grammar Grammar
}

func NewQueue(line string, grammar Grammar) *Queue {
	words := strings.Fields(strings.ToLower(line))
	q := &Queue{words: words, grammar: grammar}
	q.Peek() // eat any leading throw aways
	return q
}

func (q *Queue) More() bool {
	_, t := q.Peek()
	return t != NoWord
}

// Peek returns the next string available and the type
// will eat any TW words,
// returns "",NoWord if no more are left
func (q *Queue) Peek() (string, int) {
	// eat any TW words
	for q.idx < len(q.words) && reserved[q.words[q.idx]] == TW {
		q.idx++
	}
	if q.idx < len(q.words) {
		w := q.words[q.idx]
		return w, q.grammar[w]
	}
	return "", NoWord
}

func (q *Queue) Next() (string, int) {
	w, t := q.Peek()
	if t != NoWord {
		q.idx++
	}
	return w, t
}
