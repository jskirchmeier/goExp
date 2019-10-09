package entity

import (
	"testing"
)

func TestIdentity_AreYou(t *testing.T) {
	tests := []struct {
		name       string
		adjectives []string
		query      []string
		output     bool
	}{
		{"item", []string{"a", "b", "c"}, []string{"item"}, true},
		{"item", []string{"a", "b", "c"}, []string{"a", "item"}, true},
		{"item", []string{"a", "b", "c"}, []string{"b", "a", "item"}, true},
		{"item", []string{"a", "b", "c"}, []string{"a", "b", "item"}, true},
		{"item", []string{"a", "b", "c"}, []string{"a", "b", "c", "item"}, true},
		{"item", []string{"a", "b", "c"}, []string{"a", "c", "b", "item"}, true},
		{"item", []string{"a", "b", "c"}, []string{"a", "c", "item", "b", "item"}, true},
		{"item", []string{"a", "b", "c"}, []string{"q", "item"}, false},
		{"item", []string{}, []string{"item"}, true},
		{"item", []string{}, []string{"g", "item"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := Identity{tt.name, tt.adjectives}

			if i.AreYou(tt.query...) != tt.output {
				t.Errorf("failed query expected %v,  got %v    %s  : %v", tt.output, i.AreYou(tt.query...), i.String(), tt.query)
			}
		})
	}
}

func TestIdentity_Set(t *testing.T) {
	tests := []struct {
		name       string
		adjectives []string
		output     string
	}{
		{"item1", []string{"a", "b", "c"}, "a b c item1"},
		{"item2", []string{}, "item2"},
		{"item3", []string{"a", "b"}, "a b item3"},
		{"item4", []string{"a"}, "a item4"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := Identity{tt.name, tt.adjectives}
			if i.String() != tt.output {
				t.Errorf("expected %s,  got %s", tt.output, i.String())
			}
		})
	}
}
