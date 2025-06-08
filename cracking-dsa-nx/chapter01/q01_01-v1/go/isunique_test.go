package isunique

import (
	"testing"
)

func TestIsUnique(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{"ascii unique", "abc", true},
		{"ascii duplicate", "aabc", false},
		{"unicode unique", "🙉💩", true},
		{"unicode unique", "🙉🙉", false},
		{"empty string", "", true},
	}

	// List of implementations to test
	implementations := []struct {
		name string
		fn   isUniqueImpl
	}{
		{"Brute Force (nested loops)", BruteForce},
		{"Single-loop implementation with map-as-set", WithSet},
	}

	for _, impl := range implementations {
		for _, tt := range tests {
			t.Run(impl.name+"_"+tt.name, func(t *testing.T) {
				got := impl.fn(tt.input)
				if got != tt.want {
					t.Errorf("[%s] got %v, want %v for input %q", impl.name, got, tt.want, tt.input)
				}
			})
		}
	}
}
