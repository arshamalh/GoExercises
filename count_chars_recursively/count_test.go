package countcharsrecursively

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCount(t *testing.T) {
	assert := assert.New(t)
	testInput := []string{"ab", "c", "def", "ghij"}
	expectedOutput := 10

	got := Count(testInput)
	assert.Equal(expectedOutput, got)
}
