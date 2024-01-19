package day10

import (
	"fmt"
	"testing"
)

func TestBasicCase(t *testing.T) {
	input := [][]rune{
		[]rune("....."),
		[]rune(".S-7."),
		[]rune(".|.|."),
		[]rune(".L-J."),
		[]rune("....."),
	}

	expected := [][]int{
		{0, 0, 0, 0, 0},
		{0, 0, 1, 2, 0},
		{0, 1, 0, 3, 0},
		{0, 2, 3, 4, 0},
		{0, 0, 0, 0, 0},
	}

	output := measureDistances(&input)
	for _, row := range output {
		fmt.Println(row)
	}

	for i, row := range expected {
		for j, val := range row {
			// fmt.Println(i, j, val)

			if val != output[i][j] {
				t.Errorf("Point %d,%d: got %d, expected %d", i, j, val, output[i][j])
			}
		}
	}
}

func TestMoreComplicatedCase(t *testing.T) {
	input := [][]rune{
		[]rune("..F7."),
		[]rune(".FJ|."),
		[]rune("SJ.L7"),
		[]rune("|F--J"),
		[]rune("LJ..."),
	}

	expected := [][]int{
		{0, 0, 4, 5, 0},
		{0, 2, 3, 6, 0},
		{0, 1, 0, 7, 8},
		{1, 4, 5, 6, 7},
		{2, 3, 0, 0, 0},
	}

	output := measureDistances(&input)
	for _, row := range output {
		fmt.Println(row)
	}

	for i, row := range expected {
		for j, val := range row {
			// fmt.Println(i, j, val)

			if val != output[i][j] {
				t.Errorf("Point %d,%d: got %d, expected %d", i, j, val, output[i][j])
			}
		}
	}
}
