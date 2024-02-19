package triangular_test

import (
	triangular "go_exercises/triangular_numbers"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTriangular(t *testing.T) {
	assert := assert.New(t)
	input := 7
	expectedOutput := 28

	got := triangular.Triangular(input)
	assert.Equal(expectedOutput, got)
}
