package utils

import (
	"fmt"
	"testing"
)

func TestIndexOf(t *testing.T) {
	type Input struct {
		slice []int
		val   int
	}
	tests := []struct {
		input Input
		want  int
	}{
		{Input{slice: []int{}, val: 2}, -1},
		{Input{slice: []int{1, 2}, val: 2}, 1},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("IndexOf(%v, %d)", tc.input.slice, tc.input.val), func(t *testing.T) {
			got := IndexOf(tc.input.slice, tc.input.val)
			if got != tc.want {
				t.Errorf("Got: %v. Expected: %v", got, tc.want)
			}
		})
	}
}
