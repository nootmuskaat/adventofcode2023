package day5

import (
	"fmt"
	"testing"
)

func TestRangePositive(t *testing.T) {
	testRange := Range{100, 50, 10}

	tests := []struct {
		num, expected uint
	}{
		{50, 100},
		{51, 101},
		{59, 109},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("Range.Map(%d)->%d", tt.num, tt.expected)
		t.Run(testname, func(t *testing.T) {
			got, err := testRange.Map(tt.num)
			if err != nil {
				t.Errorf("Got unexpected error: %v", err)
			}
			if got != tt.expected {
				t.Errorf("got %d, expected %d", got, tt.expected)
			}
		})
	}
}

func TestRangeNegative(t *testing.T) {
	testRange := Range{100, 50, 10}

	tests := []struct {
		num, expected uint
	}{
		{60, 0},
		{49, 0},
		{101, 0},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("Range.Map(%d)->%d", tt.num, tt.expected)
		t.Run(testname, func(t *testing.T) {
			got, err := testRange.Map(tt.num)
			if got != tt.expected {
				t.Errorf("got %d, expected %d", got, tt.expected)
			}
			if err == nil {
				t.Errorf("Expected error, got nil")
			}
		})
	}
}

func TestFullRange(t *testing.T) {
	fr := FullRange{}

	fr.ranges = append(fr.ranges, Range{50, 98, 2})
	fr.ranges = append(fr.ranges, Range{52, 50, 48})

	tests := []struct {
		num, expected uint
	}{
		{79, 81},
		{14, 14},
		{55, 57},
		{13, 13},
		{98, 50},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("Range.Map(%d)->%d", tt.num, tt.expected)
		t.Run(testname, func(t *testing.T) {
			got := fr.Map(tt.num)
			if got != tt.expected {
				t.Errorf("got %d, expected %d", got, tt.expected)
			}
		})
	}
}
