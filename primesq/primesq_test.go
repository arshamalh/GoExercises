package primesq_test

import (
	"go_exercises/primesq"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrimeSq(t *testing.T) {
	tests := []struct {
		limit int
		want  []int
	}{
		{2, []int(nil)},
		{4, []int{2, 3}},
		{6, []int{2, 3, 5}},
		{7, []int{2, 3, 5}},
		{10, []int{2, 3, 5, 7}},
		{13, []int{2, 3, 5, 7, 11}},
	}
	for _, test := range tests {
		got := primesq.PrimeSq(test.limit)
		assert.Equal(t, test.want, got)
	}
}
