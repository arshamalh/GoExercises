package palindrome_test

import (
	"go_exercises/palindrome"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPalindrome(t *testing.T) {
	tests := []struct {
		word     string
		expected bool
	}{
		{"", true},
		{"a", true},
		{"ab", false},
		{"aba", true},
		{"abc", false},
		{"abba", true},
		{"abac", false},
		{"abacaba", true},
	}
	for _, test := range tests {
		got := palindrome.IsPalindrome(test.word)
		assert.Equal(t, test.expected, got)
	}
}
