package arrayrot

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotate(t *testing.T) {
	tests := []struct {
		nums   []int
		wanted []int
		k      int
	}{
		{[]int{1, 2, 3, 4, 5, 6}, []int{1, 2, 3, 4, 5, 6}, 0},
		{[]int{1, 2, 3, 4, 5, 6}, []int{6, 1, 2, 3, 4, 5}, 1},
		{[]int{1, 2, 3, 4, 5, 6}, []int{4, 5, 6, 1, 2, 3}, 3},
		{[]int{1, 2, 3, 4, 5, 6}, []int{5, 6, 1, 2, 3, 4}, 8},
		{[]int{1, 2, 3, 4, 5, 6}, []int{2, 3, 4, 5, 6, 1}, -1},
		{[]int{1, 2, 3, 4, 5, 6}, []int{3, 4, 5, 6, 1, 2}, -8},
		{[]int{-1}, []int{-1}, 4},
	}
	for _, test := range tests {
		Rotate(test.nums, test.k)
		assert.Equal(t, test.nums, test.wanted)
	}
}
