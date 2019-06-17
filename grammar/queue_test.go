package grammar

import (
	"testing"
)

func TestQueue1(t *testing.T) {
	q := NewQueue("One Two Three Four Five Six Seven Eight", reserved)

	tests := []string{"one", "two", "three", "four", "five", "six", "seven", "eight"}

	for _, tt := range tests {
		t.Run(tt, func(t *testing.T) {
			if !q.More() {
				t.Error("Expected more")
			}

			w, _ := q.Peek()
			if w != tt {
				t.Errorf("Peek - expected %s got %s", tt, w)
			}
			w, _ = q.Next()
			if w != tt {
				t.Errorf("Pop - expected %s got %s", tt, w)
			}

		})
	}
	if q.More() {
		t.Error("Expected no more")
	}
}

func TestQueue2(t *testing.T) {
	q := NewQueue("AND One a Two the the and Three an Four Five the Six Seven Eight and and an a the", reserved)

	tests := []string{"one", "two", "three", "four", "five", "six", "seven", "eight"}

	for _, tt := range tests {
		t.Run(tt, func(t *testing.T) {
			if !q.More() {
				t.Error("Expected more")
			}

			w, _ := q.Peek()
			if w != tt {
				t.Errorf("Peek - expected %s got %s", tt, w)
			}
			w, _ = q.Next()
			if w != tt {
				t.Errorf("Pop - expected %s got %s", tt, w)
			}
		})
	}
	if q.More() {
		t.Error("Expected no more")
	}
}
