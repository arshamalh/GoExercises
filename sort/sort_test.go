package sort

import (
	"fmt"
	"math/rand"
	"slices"
	"testing"

	"github.com/go-faker/faker/v4"
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

func TestRadixSort(t *testing.T) {
	assert := assert.New(t)

	for i, test := range test_cases {
		t.Run(fmt.Sprintf("%d-%s", i, test.name), func(t *testing.T) {
			got := RadixSort(Floater(test.given))
			assert.Equal(test.want, UnFloater(got))
		})
	}
}

func TestQuickSort(t *testing.T) {
	assert := assert.New(t)

	for i, test := range test_cases {
		t.Run(fmt.Sprintf("%d-%s", i, test.name), func(t *testing.T) {
			got := QuickSort(test.given)
			assert.Equal(test.want, got)
		})
	}
}

func TestInsertionSort(t *testing.T) {
	assert := assert.New(t)

	for i, test := range test_cases {
		t.Run(fmt.Sprintf("%d-%s", i, test.name), func(t *testing.T) {
			got := InsertionSort(test.given)
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

func TestMerge(t *testing.T) {
	assert := assert.New(t)
	merge_test_cases := []struct {
		left_side  []int
		right_side []int
		expected   []int
	}{
		{
			left_side:  []int{1, 2, 4},
			right_side: []int{5, 9, 11, 22},
			expected:   []int{1, 2, 4, 5, 9, 11, 22},
		},
		{
			left_side:  []int{3},
			right_side: []int{1},
			expected:   []int{1, 3},
		},
		{
			// Merge function should work, although, this scenario should not happen in Merge sort itself.
			left_side:  []int{5, -1},
			right_side: []int{2, 1, 0},
			expected:   []int{2, 1, 0, 5, -1},
		},
	}
	for i, test := range merge_test_cases {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			got := Merge(test.left_side, test.right_side)
			assert.Equal(test.expected, got)
		})
	}
}

func TestDivider(t *testing.T) {
	assert := assert.New(t)
	merge_test_cases := []struct {
		given       []int
		expected_pp int
		expected_ls []int
		expected_rs []int
	}{
		{
			given:       []int{0, 1},
			expected_pp: 0,
			expected_ls: []int{},
			expected_rs: []int{1},
		},
	}

	for i, test := range merge_test_cases {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			got_pp, got_ls, got_rs := Divider(test.given)
			assert.Equal(test.expected_pp, got_pp)
			assert.Equal(test.expected_ls, got_ls)
			assert.Equal(test.expected_rs, got_rs)
		})
	}
}

var llArr, _ = faker.RandomInt(1, 100, 1000)

// *** Benchmarks *** //
func BenchmarkInsertionSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsertionSort(llArr)
	}
}

func BenchmarkQuickSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QuickSort(llArr)
	}
}

func BenchmarkNativeSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slices.SortStableFunc(llArr, func(a, b int) int {
			return a - b
		})
	}
}
