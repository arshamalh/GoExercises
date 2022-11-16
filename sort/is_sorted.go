package sort

func isSorted(array []int) bool {
	for i := 0; i < len(array)-1; i++ {
		if array[i+1] < array[i] {
			return false
		}
	}
	return true
}
