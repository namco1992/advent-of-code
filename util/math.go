package util

func Max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func Abs(i int) int {
	if i >= 0 {
		return i
	}
	return -i
}
