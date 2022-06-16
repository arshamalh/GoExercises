package even_odd_test

import (
	"go_exercises/even_odd"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvenOdd(t *testing.T) {
	var tests = []struct {
		num  int
		want []string
	}{
		{1, []string{"1 odd"}},
		{2, []string{"1 odd", "2 even"}},
		{4, []string{"1 odd", "2 even", "3 odd", "4 even"}},
		{6, []string{"1 odd", "2 even", "3 odd", "4 even", "5 odd", "6 even"}},
	}
	for _, test := range tests {
		got := even_odd.EvenOdd(test.num)
		assert.Equal(t, test.want, got, "should be equal")

	}
}
