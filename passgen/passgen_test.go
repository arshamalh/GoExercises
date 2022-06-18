package passgen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPassgen(t *testing.T) {
	SeedRandom()
	assert := assert.New(t)
	for i := 0; i < 20; i++ {
		password := GeneratePass(uint8(i))
		assert.GreaterOrEqual(len(password), 8)
		assert.GreaterOrEqual(len(password), i)

		lowercase := containsAny(password, "abcdefghijklmnopqrstuvwxyz", 2)
		uppercase := containsAny(password, "ABCDEFGHIJKLMNOPQRSTUVWXYZ", 2)
		number := containsAny(password, "0123456789", 2)
		symbol := containsAny(password, Symbols, 2)
		assert.Equalf(true, lowercase && uppercase && number && symbol, "should include at least two uppercase, lowercase, number and symbol, pass: %s", password)
	}
}

func containsAny(source string, character_set string, number_of_contains int) bool {
	element_count := 0
	for _, char := range character_set {
		for _, char_from_source := range source {
			if char_from_source == char {
				element_count += 1
			}
		}
	}
	return element_count >= number_of_contains
}
