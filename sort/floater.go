package sort

// Converting an array of ints to an array of floats in a BAD way, just for testing purposes.
func Floater(array []int) (result []float64) {
	for _, item := range array {
		result = append(result, float64(item))
	}
	return result
}

// Converting an array of floats to an array of ints in a BAD way, just for testing purposes.
func UnFloater(array []float64) (result []int) {
	for _, item := range array {
		result = append(result, int(item))
	}
	return result
}
