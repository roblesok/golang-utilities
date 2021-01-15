package utils

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		input []int
		want  []int
	}{
		{input: []int{}, want: []int{}},
		{input: []int{1, 2, 3}, want: []int{3, 2, 1}},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("Reverse(%v)", tc.input), func(t *testing.T) {
			got := Reverse(tc.input)
			if !Equal(got, tc.want) {
				t.Errorf("Got: %v. Expected: %v", got, tc.want)
			}
		})
	}
}
