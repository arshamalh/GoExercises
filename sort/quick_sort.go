package sort

func QuickSort(array []int) []int {
	if len(array) <= 1 {
		return array
	}

	pivot_point, left_array, right_array := Divider(array)
	left := QuickSort(left_array)
	right := QuickSort(right_array)
	left = append(left, array[pivot_point])
	return append(left, right...)
}

// Divider will select a pivot number and order fewer and higher numbers on the left and right side of it.
func Divider(array []int) (pp int, left []int, right []int) {
	for i := 1; i < len(array); i++ {
		if array[i] <= array[0] {
			pp++
			array[pp], array[i] = array[i], array[pp]
		}
	}
	array[0], array[pp] = array[pp], array[0]
	return pp, array[:pp], array[pp+1:]
}
