package utils

import (
	"testing"
)

func TestValidateVersion(t *testing.T) {
	tests := []struct {
		minVersion     string
		currentVersion string
		expected       bool
	}{
		{"1.0.0", "1.0.0", true},
		{"1.0.0", "1.0.1", true},
		{"1.0.1", "1.0.0", false},
		{"2.0.0", "3.0.0", true},
	}

	for _, tt := range tests {
		t.Run(tt.currentVersion, func(t *testing.T) {
			result := ValidateVersion(tt.minVersion, tt.currentVersion)
			if result != tt.expected {
				t.Errorf("versionCheck(%s, %s) = %v, want %v", tt.minVersion, tt.currentVersion, result, tt.expected)
			}
		})
	}
}
