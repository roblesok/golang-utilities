package utils

import "testing"

func TestInsert(t *testing.T) {
	var got, expected []int
	t.Run("Insert items", func(t *testing.T) {
		got = Insert([]int{1, 2, 3, 4}, 1, []int{100, 200}...)
		expected = []int{1, 100, 200, 2, 3, 4}
		if !Equal(got, expected) {
			t.Errorf("Got: %v. Expected: %v\n", got, expected)
		}
	})
	t.Run("Insert a item", func(t *testing.T) {
		got = Insert([]int{1, 2, 3, 4}, 1, 100)
		expected = []int{1, 100, 2, 3, 4}
		if !Equal(got, expected) {
			t.Errorf("Got: %v. Expected: %v\n", got, expected)
		}
	})
}
