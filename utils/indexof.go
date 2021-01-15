package utils

// IndexOf return the index of first ocurrence of val
// returns -1 if the item not found
func IndexOf(slice []int, val int) int {
	for i := 0; i < len(slice); i++ {
		if slice[i] == val {
			return i
		}
	}
	return -1
}
