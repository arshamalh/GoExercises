package golomb

// Golomb sequence, get the index and returns the value of that index in golomb sequence
func Golomb(n int) int {
	if n == 1 {
		return 1
	}
	return 1 + Golomb(n-Golomb(Golomb(n-1)))
}

func GolombOptimized(n int, data map[int]int) int {
	if n == 1 {
		return 1
	}
	if val, ok := data[n]; ok {
		return val
	}
	statement1 := GolombOptimized(n-1, data)
	data[n-1] = statement1

	statement2 := GolombOptimized(statement1, data)
	data[statement1] = statement2

	statement3 := GolombOptimized(n-statement2, data)
	data[n-statement2] = statement3

	result := 1 + statement3
	data[n] = result

	return result
}
