package even_odd

import (
	"strconv"
)

func EvenOdd(num int) (result_sequence []string) {
	for i := 1; i <= num; i++ {
		str_num := strconv.Itoa(i)
		if i%2 == 0 {
			result_sequence = append(result_sequence, str_num+" even")
		} else {
			result_sequence = append(result_sequence, str_num+" odd")
		}
	}
	return result_sequence
}
