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
//
// Important Note: As we're using 0 for pp, we also used it for keeping count, but if pp initial value is not 0,
// we should use two separate variables.
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

// Not completely implemented, it will get custom pivot point
func DividerPivot(array []int, pp int) (left []int, right []int) {
	count := 0
	for i := 0; i < len(array); i++ {
		if i == pp {
			continue
		}
		if array[i] <= array[pp] {
			count++
			array[count], array[i] = array[i], array[count]
		}
	}
	// TODO: Not working, needs more work around it.
	array[pp], array[count] = array[count], array[pp]
	return array[:count], array[count+1:]
}
