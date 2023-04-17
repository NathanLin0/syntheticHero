package utils

func IndexOf[T string | ~int | bool](arr []T, val T) int {
	for i, v := range arr {
		if v == val {
			return i
		}
	}
	return -1
}
