package fizzbuzz_test

import (
	"go_exercises/fizzbuzz"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	var tests = []struct {
		num  int
		want string
	}{
		{3, "Fizz"},
		{5, "Buzz"},
		{15, "FizzBuzz"},
		{4, "4"},
		{7, "7"},
		{10, "Buzz"},
	}

	for _, test := range tests {
		if got := fizzbuzz.FizzBuzz(test.num); got != test.want {
			t.Errorf("FizzBuzz(%d) = %q; want %q", test.num, got, test.want)
		}
	}
}
