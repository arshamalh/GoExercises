package sort

// One of the famous sorting algorithms
func InsertionSort(array []int) []int {
	for i := 1; i < len(array); i++ {
		currentVal := array[i]
		last_node := 0
		for j := i - 1; j >= 0; j-- {
			if array[j] > currentVal {
				array[j+1] = array[j]
			} else {
				last_node = j + 1
				break
			}
		}
		array[last_node] = currentVal
	}
	return array
}

// Array type for sorted insert
type Array []int

// Sorted Insert algorithm which inserts new items in order.
func (arr *Array) SortedInsert(item int) {
	*arr = append(*arr, item)
	for i := len(*arr) - 2; 0 <= i; i-- {
		current := (*arr)[i]
		if item < current {
			(*arr)[i+1] = current
			(*arr)[i] = item
		}
	}
}
