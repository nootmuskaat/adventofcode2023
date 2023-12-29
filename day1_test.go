// I have not yet figured out how to set up testing so this will
// eventually be setup correctly
package main

import "testing"

func testFirstInt(t *testing.T) {

	input, first, last := "abc123", 1, 3
	if got := firstInt(input); got != first {
		t.Errorf("Expected %d, got %d", first, got)
	}
	if got := lastInt(input); got != last {
		t.Errorf("Expected %d, got %d", last, got)
	}

	input, first, last = "987", 9, 7
	if got := firstInt(input); got != first {
		t.Errorf("Expected %d, got %d", first, got)
	}
	if got := lastInt(input); got != last {
		t.Errorf("Expected %d, got %d", last, got)
	}

	input, first, last = "asdf", 0, 0
	if got := firstInt(input); got != first {
		t.Errorf("Expected %d, got %d", first, got)
	}
	if got := lastInt(input); got != last {
		t.Errorf("Expected %d, got %d", last, got)
	}
}
