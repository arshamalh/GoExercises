package pymod

func PyMod(d int, m int) int {
	if d < 0 && m < 0 {
		return d % m
	}
	d %= m
	if d < 0 {
		d += m
	}
	return d
}
