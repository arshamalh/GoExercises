package fibosq

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFiboSqLoop(t *testing.T) {
	tests := []struct {
		limit int
		want  []int
	}{
		{0, []int(nil)},
		{1, []int{0}},
		{2, []int{0, 1, 1}},
		{3, []int{0, 1, 1, 2}},
		{10, []int{0, 1, 1, 2, 3, 5, 8}},
		{13, []int{0, 1, 1, 2, 3, 5, 8}},
	}

	for _, test := range tests {
		got := FiboSqLoop(test.limit)
		assert.Equal(t, test.want, got)
	}
}
