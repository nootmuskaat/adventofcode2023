package day9

import (
	"fmt"
	"testing"
)

func TestDrillDown(t *testing.T) {
	testCases := []struct {
		values   []int
		expected int
	}{
		{[]int{0, 3, 6, 9, 12, 15}, 18},
		{[]int{1, 3, 6, 10, 15, 21}, 28},
		{[]int{10, 13, 16, 21, 30, 45}, 68},
	}

	for _, tc := range testCases {
		name := fmt.Sprintf("%v->%d", tc.values, tc.expected)
		t.Run(name, func(t *testing.T) {
			got := DrillDown(tc.values)

			if got != tc.expected {
				t.Errorf("Expected %d, got %d", tc.expected, got)
			}
		})
	}
}
