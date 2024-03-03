package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickSelect(t *testing.T) {
	assert := assert.New(t)
	array := []int{5, 7, 1, 2, 9, 1, 2}
	cases := []struct {
		nthHighest int
		value      int
	}{
		{1, 9}, // first highest
		{2, 7}, // Second highest
		{3, 5}, // third highest
		{5, 2}, // 5th highest is 2 because we have duplicated two, that counts!
		{7, 1}, // 7th highest
	}

	for _, c := range cases {
		got := QuickSelectHighest(array, c.nthHighest)
		assert.Equal(c.value, got)
	}
}
