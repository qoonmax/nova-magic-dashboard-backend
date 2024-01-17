package utils

import (
	"testing"
)

func TestGetSingularName(t *testing.T) {
	tests := []struct {
		pluralWord   string
		singularWord string
	}{
		{"items", "Item"},
		{"users", "User"},
		{"lists", "List"},
		{"people", "Person"},
	}

	for _, tt := range tests {
		t.Run(tt.pluralWord, func(t *testing.T) {
			result := GetModelName(tt.pluralWord)
			if result != tt.singularWord {
				t.Errorf("GetSingularName(%s) = %v, want %v", tt.pluralWord, result, tt.singularWord)
			}
		})
	}
}
