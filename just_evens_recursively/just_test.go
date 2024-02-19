package justevensrecursively

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJustEvens(t *testing.T) {
	assert := assert.New(t)
	input := []int{0, 1, 2, 3, 4, 5}
	expectedOutput := []int{0, 2, 4}

	got := JustEvens(input)

	assert.ElementsMatch(got, expectedOutput)
}
