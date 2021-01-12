package utils

// Shift add a item front of slice. It returns a new slice.
func Shift(slice []int, val int) []int {
	return append([]int{val}, slice...)
}
