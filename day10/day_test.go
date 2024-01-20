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
	/* for _, row := range *output {
		fmt.Println(row)
	} */

	for i, row := range expected {
		for j, val := range row {
			// fmt.Println(i, j, val)

			if val != (*output)[i][j] {
				t.Errorf("Point %d,%d: got %d, expected %d", i, j, val, (*output)[i][j])
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
	/* for _, row := range *output {
		fmt.Println(row)
	} */

	for i, row := range expected {
		for j, val := range row {
			// fmt.Println(i, j, val)

			if val != (*output)[i][j] {
				t.Errorf("Point %d,%d: got %d, expected %d", i, j, val, (*output)[i][j])
			}
		}
	}
}

func TestInsideOut(t *testing.T) {
	input := [][]rune{
		[]rune(".........."),
		[]rune(".S------7."),
		[]rune(".|F----7|."),
		[]rune(".||....||."),
		[]rune(".||....||."),
		[]rune(".|L-7F-J|."),
		[]rune(".|..||..|."),
		[]rune(".L--JL--J."),
		[]rune(".........."),
	}

	expected := [][]int{
		{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
		{-1, 1, 1, 2, 3, 4, 5, 6, 7, -1},
		{-1, 1, 16, 17, 18, 19, 20, 21, 8, -1},
		{-1, 2, 15, -1, -1, -1, -1, 22, 9, -1},
		{-1, 3, 14, -1, -1, -1, -1, 21, 10, -1},
		{-1, 4, 13, 12, 11, 18, 19, 20, 11, -1},
		{-1, 5, -2, -2, 10, 17, -2, -2, 12, -1},
		{-1, 6, 7, 8, 9, 16, 15, 14, 13, -1},
		{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
	}

	output := measureDistances(&input)
	identifyInsideOut(&input, output)

	for _, row := range *output {
		fmt.Println(row)
	}

	for i, row := range expected {
		for j, val := range row {
			// fmt.Println(i, j, val)

			if val != (*output)[i][j] {
				t.Errorf("Point %d,%d: got %d, expected %d", i, j, val, (*output)[i][j])
			}
		}
	}
}
