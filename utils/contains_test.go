package utils

import (
	"fmt"
	"testing"
)

func TestContains(t *testing.T) {
	type Input struct {
		slice []int
		value int
	}
	tests := []struct {
		input Input
		want  bool
	}{
		{Input{[]int{1, 2, 3}, 100}, false},
		{Input{[]int{1, 2, 3}, 1}, true},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("Contains(%v, %d)", tc.input.slice, tc.input.value), func(t *testing.T) {
			got := Contains(tc.input.slice, tc.input.value)
			if got != tc.want {
				t.Errorf("Got: %v. Expected: %v.", got, tc.want)
			}
		})
	}
}
