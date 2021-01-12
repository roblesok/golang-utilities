package utils

// Equal returns true if both slices contains the same elements.
// A nil argument is equivalent to empty slice.
func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for idx, val := range a {
		if val != b[idx] {
			return false
		}
	}
	return true
}
