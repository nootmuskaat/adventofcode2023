// I have not yet figured out how to set up testing so this will
// eventually be setup correctly
package day1

import (
	"fmt"
	"testing"
)

func TestFirstInt(t *testing.T) {

	testCases := []struct {
		input       string
		first, last int
	}{
		{"abc123", 1, 3},
		{"987", 9, 7},
		{"asdf", 0, 0},
		{"onetwo", 1, 2},
		{"one55two", 1, 2},
		{"one5two5", 1, 5},
		{"6one5two", 6, 2},
	}

	for i, tc := range testCases {

		testname := fmt.Sprintf("('%s')->%d%d)", tc.input, tc.first, tc.last)
		t.Run(testname, func(t *testing.T) {
			if got := firstInt(&tc.input); got != tc.first {
				t.Errorf("Test case %d, Expected %d, got %d", i, tc.first, got)
			}
			if got := lastInt(&tc.input); got != tc.last {
				t.Errorf("Test case %d, Expected %d, got %d", i, tc.last, got)
			}
		})
	}
}
