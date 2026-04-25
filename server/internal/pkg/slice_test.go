package pkg

import (
	"testing"
)

func TestIsContain(t *testing.T) {
	tests := []struct {
		name   string
		list   []int
		item   int
		expect bool
	}{
		{"item exists", []int{1, 2, 3}, 2, true},
		{"item not exists", []int{1, 2, 3}, 4, false},
		{"empty list", []int{}, 1, false},
		{"single item match", []int{5}, 5, true},
		{"single item no match", []int{5}, 3, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsContain(tt.list, tt.item)
			if result != tt.expect {
				t.Errorf("IsContain(%v, %d) = %v, want %v", tt.list, tt.item, result, tt.expect)
			}
		})
	}
}

func TestIsContainString(t *testing.T) {
	tests := []struct {
		name   string
		list   []string
		item   string
		expect bool
	}{
		{"string exists", []string{"a", "b", "c"}, "b", true},
		{"string not exists", []string{"a", "b", "c"}, "d", false},
		{"empty string list", []string{}, "a", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsContain(tt.list, tt.item)
			if result != tt.expect {
				t.Errorf("IsContain(%v, %q) = %v, want %v", tt.list, tt.item, result, tt.expect)
			}
		})
	}
}
