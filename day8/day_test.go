package day8

import (
	// "fmt"
	"testing"
)

func TestIterateInstructions(t *testing.T) {
	expected := "LLRLLRLLR"

	vals, done := Instructor("LLR")

	got := make([]rune, 9)
	for i := 0; i < 9; i++ {
		got[i] = <-vals
	}
	done <- true
	close(done)

	if s := string(got); s != expected {
		t.Errorf("Got %s, expected %s", s, expected)
	}

}

func TestReadLine(t *testing.T) {
	line := "CDQ = (MSF, BKM)"
	expectedL := Location("CDQ")
	expectedSP := SignPost{Location("MSF"), Location("BKM")}

	gotL, gotSP := readLine(line)
	if string(gotL) != string(expectedL) {
		t.Errorf("Location: Got %s, expected %s", gotL, expectedL)
	}
	if string(gotSP.left) != string(expectedSP.left) {
		t.Errorf("SingPost.left: Got %s, expected %s", gotSP.left, expectedSP.left)
	}
	if string(gotSP.right) != string(expectedSP.right) {
		t.Errorf("SignPost.right: Got %s, expected %s", gotSP.right, expectedSP.right)
	}

}
