package sort

func BubbleSort(array []int) []int {
	for range array {
		for i := 0; i < len(array)-1; i++ {
			if array[i] > array[i+1] {
				array[i], array[i+1] = array[i+1], array[i]
			}
		}
	}
	return array
}

// In this version, if we don't have any swap in the inner loop, we're done.
func BubbleSortExtended(array []int) []int {
	didSwap := false
	for range array {
		didSwap = false
		for i := 0; i < len(array)-1; i++ {
			if array[i] > array[i+1] {
				array[i], array[i+1] = array[i+1], array[i]
				didSwap = true
			}
		}
		if !didSwap {
			return array
		}
	}
	return array
}
