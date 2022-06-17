package sumdiv_test

import (
	"go_exercises/sumdiv"
	"testing"
)

func TestSumDiv(t *testing.T) {
	tests := []struct {
		limit int
		sum   int
	}{
		{10, 25},
		{100, 2208},
		{300, 18966},
	}

	for _, test := range tests {
		if sumdiv.SumDiv(test.limit) != test.sum {
			t.Errorf("SumDiv(%d) = %d, want %d", test.limit, sumdiv.SumDiv(test.limit), test.sum)
		}
	}
}
