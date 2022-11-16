package sort

import (
	"math"
)

// This only works on integers, it's not based on comparison.
func RadixSort(array []float64) []float64 {
	max_digits := MaxDigits(array)
	for k := 0; k < max_digits; k++ {
		bucks := Buckets{}
		for _, item := range array {
			digit := GetDigit(item, k)
			bucks.Add(digit, item)
		}

		new_array := []float64{}
		for j := 0; j < 10; j++ {
			if buck, ok := bucks.Get(j); ok {
				new_array = append(new_array, buck...)
			}
		}
		array = new_array
	}
	return array
}

// Bucks type specifically made for radix sort, radix sort will keep 10 bucks from 0 to 10.
// It can also be a 2d matrix
type Buckets map[int][]float64

func (b *Buckets) Add(digit int, number float64) {
	(*b)[digit] = append((*b)[digit], number)
}

func (b *Buckets) Get(digit int) (number []float64, ok bool) {
	number, ok = (*b)[digit]
	return
}

// Get digit at k
func GetDigit(number float64, k int) int {
	return int(math.Floor(
		math.Abs(number)/math.Pow10(k),
	)) % 10
}

// Get Number of digits in a number
func NumberOfDigits(number float64) int {
	return int(math.Floor(math.Log10(number))) + 1
}

// Get the number of digits that biggest number in an array has.
func MaxDigits(array []float64) int {
	max_digits := 0
	for _, item := range array {
		if digits := NumberOfDigits(item); digits > max_digits {
			max_digits = digits
		}
	}
	return max_digits
}
