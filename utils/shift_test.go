package utils

import "testing"

func TestShift(t *testing.T) {
	got := Shift([]int{1, 2, 3}, 100)
	expected := []int{100, 1, 2, 3}
	if !Equal(got, expected) {
		t.Errorf("Got: %v. Expected: %v\n", got, expected)
	}
}
