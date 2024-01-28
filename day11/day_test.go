package day11

import (
	"fmt"
	"testing"
)

func TestParts(t *testing.T) {
	gmap := GalaxyMap{
		{false, false, false, true, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, true, false, false},
		{true, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, true, false, false, false},
		{false, true, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, true},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, true, false, false},
		{true, false, false, false, true, false, false, false, false, false},
	}

	testCases := []struct {
		expansionFactor, expected int
	}{
		{2, 374},
		{10, 1030},
		{100, 8410},
	}

	for _, tc := range testCases {
		testName := fmt.Sprintf("%d->%d", tc.expansionFactor, tc.expected)
		t.Run(testName, func(t *testing.T) {
			got := trackDistances(&gmap, tc.expansionFactor)

			if got != tc.expected {
				t.Errorf("Expected %d but got %d", tc.expected, got)
			}
		})
	}
}
