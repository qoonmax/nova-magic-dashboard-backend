package utils

import (
	"testing"
)

func TestGetSingularName(t *testing.T) {
	tests := []struct {
		pluralWord   string
		singularWord string
	}{
		{"items", "item"},
		{"users", "user"},
		{"lists", "list"},
		{"people", "person"},
	}

	for _, tt := range tests {
		t.Run(tt.pluralWord, func(t *testing.T) {
			result := GetSingularName(tt.pluralWord)
			if result != tt.singularWord {
				t.Errorf("GetSingularName(%s) = %v, want %v", tt.pluralWord, result, tt.singularWord)
			}
		})
	}
}
