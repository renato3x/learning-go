package strings

import (
	"strings"
	"testing"
)

const defaultError = "Expected %v, got %v"

func TestIndex(t *testing.T) {
	tests := []struct {
		text     string
		search   string
		expected int
	}{
		{"Hello, world!", "world", 7},
		{"Hello, world!", "Hello", 0},
		{"Hello, world!", "!", 12},
		{"Hello, world!", "notfound", -1},
	}

	for _, test := range tests {
		t.Logf("Cool: %v", test.text)
		actual := strings.Index(test.text, test.search)

		if actual != test.expected {
			t.Errorf(defaultError, test.expected, actual)
		}
	}
}
