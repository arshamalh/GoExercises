package sort

/*
Let's say you have an array in random order and you do not need to sort it,
but you do want ot know the tenth lowest value in the array,
or the fifth highest.

This can useful if we had a lot of test grades and want to know what the 25th percentile was,
or if we want to find the median grade.
*/
func QuickSelectHighest(array []int, nthHighest int) int {
	pp, leftArr, rightArr := Divider(array)
	index := len(array) - nthHighest
	if index < pp {
		return QuickSelectHighest(leftArr, nthHighest-len(rightArr)-1)
	}
	if index > pp {
		return QuickSelectHighest(rightArr, nthHighest)
	}
	return array[pp]
}
