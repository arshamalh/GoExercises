package findxrecursively

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFind(t *testing.T) {
	assert := assert.New(t)
	input := "abcdefghijklmnopqrstuvwxyz"
	expectedOutput := 23

	got := Find(input)
	assert.Equal(expectedOutput, got)
}
