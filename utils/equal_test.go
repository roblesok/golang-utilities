package utils

import "testing"

func TestEqual(t *testing.T) {
	var got, expected bool
	t.Run("empty slices", func(t *testing.T) {
		got = Equal(nil, nil)
		expected = true
		if got != expected {
			t.Errorf("Got: %v. Expected: %v", got, expected)
		}
	})
	t.Run("different slices", func(t *testing.T) {
		got = Equal([]int{1, 2}, nil)
		expected = false
		if got != expected {
			t.Errorf("Got: %v. Expected: %v", got, expected)
		}
	})
	t.Run("equal slices", func(t *testing.T) {
		got = Equal([]int{1, 2}, []int{1, 2})
		expected = true
		if got != expected {
			t.Errorf("Got: %v. Expected: %v", got, expected)
		}
	})
}
