package utils

func Insert(slice []int, idx int, values ...int) []int {
	// If slice has capacity
	if n := len(slice) + len(values); n <= cap(slice) {
		res := slice[:n]                         // add items before index
		copy(res[idx+len(values):], slice[idx:]) // add values
		copy(res[idx:], values)                  // add items after new values
		return res
	}
	// create a slice with just the right capacity
	res := make([]int, len(slice)+len(values))
	copy(res, slice[:idx])
	copy(res[idx:], values)
	copy(res[idx+len(values):], slice[idx:])
	return res
}
