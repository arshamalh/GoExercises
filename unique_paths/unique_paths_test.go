package uniquepaths_test

import (
	uniquepaths "go_exercises/unique_paths"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniquePaths(t *testing.T) {
	assert := assert.New(t)

	cases := []struct {
		rows     int
		columns  int
		expected int
	}{
		{1, 5, 1},
		{5, 1, 1},
		{2, 3, 3},
		{3, 3, 6},
	}

	for _, test := range cases {
		t.Run("", func(t *testing.T) {
			got := uniquepaths.UniquePaths(test.rows, test.columns)
			assert.Equal(test.expected, got)
		})
	}
}
