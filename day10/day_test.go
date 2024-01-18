package day10

import (
	// "fmt"
	"testing"
)

func TestSomething(t *testing.T) {
	input := []string{
		".....",
		".S-7.",
		".|.|.",
		".L-J.",
		".....",
	}

	expected := [][]int{
		{0, 0, 0, 0, 0},
		{0, 0, 1, 2, 0},
		{0, 1, 0, 3, 0},
		{0, 2, 3, 4, 0},
		{0, 0, 0, 0, 0},
	}

	output := measureDistances(&input)

	for i, row := range expected {
		for j, val := range row {
			// fmt.Println(i, j, val)

			if val != output[i][j] {
				t.Errorf("Point %d,%d: got %d, expected %d", i, j, val, output[i][j])
			}
		}
	}
}
