package sort

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

var test_cases = []struct {
	name  string
	given []int
	want  []int
}{
	{
		name:  "empty",
		given: []int{},
		want:  []int{},
	},
	{
		name:  "negative included",
		given: []int{5, -1, 2, 1, 0},
		want:  []int{-1, 0, 1, 2, 5},
	},
	{
		name:  "long list",
		given: []int{2, 4, 8, 5, 2, 1, 9, 3, 7},
		want:  []int{1, 2, 2, 3, 4, 5, 7, 8, 9},
	},
	{
		name:  "almost sorted",
		given: []int{1, 2, 5, 4, 6, 7},
		want:  []int{1, 2, 4, 5, 6, 7},
	},
}

func TestIsSorted(t *testing.T) {
	assert := assert.New(t)

	array := []int{}
	for i := 0; i < 100; i++ {
		array = append(array, i)
	}

	assert.Equal(true, isSorted(array))
}

func TestSort(t *testing.T) {
	assert := assert.New(t)

	for i, test := range test_cases {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			got := BubbleSortExtended(test.given)
			assert.Equal(test.want, got)
		})
	}
}

func TestSortedInsert(t *testing.T) {
	assert := assert.New(t)
	array := Array{}

	for i := 0; i < 30; i++ {
		random := rand.Intn(150)
		array.SortedInsert(random)
	}
	assert.NotEmpty(array)
	got := isSorted([]int(array))
	assert.Equal(true, got)
}
