package sort

func MergeSort(array []int) []int {
	if len(array) <= 1 {
		return array
	}
	middle := len(array) / 2
	left_side := MergeSort(array[:middle])
	right_side := MergeSort(array[middle:])
	return Merge(left_side, right_side)
}

func Merge(arleft []int, aright []int) (result []int) {
	i, j := 0, 0
	ll, lr := len(arleft), len(aright)

	// Comparing Values one by one of both arrays
	for i < ll && j < lr {
		if arleft[i] < aright[j] {
			result = append(result, arleft[i])
			i++
		} else {
			result = append(result, aright[j])
			j++
		}
	}

	// Finding remaining included array and appending the rest to the result.
	if i < ll {
		return append(result, arleft[i:]...)
	} else if j < lr {
		return append(result, aright[j:]...)
	}

	return result
}
