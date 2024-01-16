package day8

import (
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

func TestLoopFinding(t *testing.T) {
	inst := "LR"
	whereTo := map[Location]SignPost{
		Location("11A"): {Location("11B"), Location("XXX")},
		Location("11B"): {Location("XXX"), Location("11Z")},
		Location("11Z"): {Location("11B"), Location("XXX")},
		Location("22A"): {Location("22B"), Location("XXX")},
		Location("22B"): {Location("22C"), Location("22C")},
		Location("22C"): {Location("22Z"), Location("22Z")},
		Location("22Z"): {Location("22B"), Location("22B")},
		Location("XXX"): {Location("XXX"), Location("XXX")},
	}
	done := make(chan bool)

	testCases := []struct {
		start       Location
		expectedLen int
		expectedEP  []int
	}{
		{Location("11A"), 2, []int{2}},
		{Location("22A"), 6, []int{3, 6}},
	}

	for _, tc := range testCases {
		t.Run(string(tc.start), func(t *testing.T) {
			loop := Loop{0, []int{}}

			go findLoop(inst, tc.start, &whereTo, &loop, done)
			<-done

			if tc.expectedLen != loop.length {
				t.Errorf("expected len %d, got %d", tc.expectedLen, loop.length)
			}

			if len(tc.expectedEP) != len(loop.endpoints) {
				t.Errorf("expected len %v, got %v", tc.expectedEP, loop.endpoints)
			}
			for i, v := range tc.expectedEP {
				if v != loop.endpoints[i] {
					t.Errorf("expected len %v, got %v", tc.expectedEP, loop.endpoints)
				}
			}
		})
	}
}
