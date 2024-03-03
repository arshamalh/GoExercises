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

func Merge(arrLeft []int, arrRight []int) (result []int) {
	i, j := 0, 0
	lenLeftArr, lenRightArr := len(arrLeft), len(arrRight)

	// Comparing Values one by one of both arrays
	for i < lenLeftArr && j < lenRightArr {
		if arrLeft[i] < arrRight[j] {
			result = append(result, arrLeft[i])
			i++
		} else {
			result = append(result, arrRight[j])
			j++
		}
	}

	// Finding remaining included array and appending the rest to the result.
	if i < lenLeftArr {
		return append(result, arrLeft[i:]...)
	} else if j < lenRightArr {
		return append(result, arrRight[j:]...)
	}

	return result
}
