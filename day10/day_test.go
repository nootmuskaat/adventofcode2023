package day10

import (
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
		t.Log(row)
	} */

	for i, row := range expected {
		for j, val := range row {
			// t.Log(i, j, val)

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
		t.Log(row)
	} */

	for i, row := range expected {
		for j, val := range row {
			if val != (*output)[i][j] {
				t.Errorf("Point %d,%d: got %d, expected %d", i, j, val, (*output)[i][j])
			}
		}
	}
}

func TestInsideOutSimple(t *testing.T) {
	input := [][]rune{
		//      0123456789
		[]rune(".........."), // 0
		[]rune(".S------7."),
		[]rune(".|F----7|."),
		[]rune(".||....||."), // 3
		[]rune(".||....||."),
		[]rune(".|L-7F-J|."),
		[]rune(".|..||..|."), // 6
		[]rune(".L--JL--J."),
		[]rune(".........."),
	}

	expected := [][]Ptype{
		{OUTS, OUTS, OUTS, OUTS, OUTS, OUTS, OUTS, OUTS, OUTS, OUTS},
		{OUTS, PIPE, PIPE, PIPE, PIPE, PIPE, PIPE, PIPE, PIPE, OUTS},
		{OUTS, PIPE, PIPE, PIPE, PIPE, PIPE, PIPE, PIPE, PIPE, OUTS},
		{OUTS, PIPE, PIPE, OUTS, OUTS, OUTS, OUTS, PIPE, PIPE, OUTS},
		{OUTS, PIPE, PIPE, OUTS, OUTS, OUTS, OUTS, PIPE, PIPE, OUTS},
		{OUTS, PIPE, PIPE, PIPE, PIPE, PIPE, PIPE, PIPE, PIPE, OUTS},
		{OUTS, PIPE, UNKN, UNKN, PIPE, PIPE, UNKN, UNKN, PIPE, OUTS},
		{OUTS, PIPE, PIPE, PIPE, PIPE, PIPE, PIPE, PIPE, PIPE, OUTS},
		{OUTS, OUTS, OUTS, OUTS, OUTS, OUTS, OUTS, OUTS, OUTS, OUTS},
	}

	typeMap := part2Main(&input)

	for i, row := range *typeMap {
		t.Log(row, expected[i])
	}

	for i, row := range expected {
		for j, val := range row {
			if val != (*typeMap)[i][j] {
				t.Errorf("Point %d,%d: got %d, expected %d", i, j, val, (*typeMap)[i][j])
			}
		}
	}
}

func TestInsideNext(t *testing.T) {
	input := [][]rune{
		[]rune(".F----7F7F7F7F-7...."),
		[]rune(".|F--7||||||||FJ...."),
		[]rune(".||.FJ||||||||L7...."),
		[]rune("FJL7L7LJLJ||LJ.L-7.."),
		[]rune("L--J.L7...LJS7F-7L7."),
		[]rune("....F-J..F7FJ|L7L7L7"),
		[]rune("....L7.F7||L7|.L7L7|"),
		[]rune(".....|FJLJ|FJ|F7|.LJ"),
		[]rune("....FJL-7.||.||||..."),
		[]rune("....L---J.LJ.LJLJ..."),
	}

	expected := [][]Ptype{
		{OUTS, PIPE, PIPE, PIPE, PIPE, PIPE, PIPE, PIPE, PIPE, PIPE, PIPE, PIPE, PIPE, PIPE, PIPE, PIPE, OUTS, OUTS, OUTS, OUTS},
		{OUTS, PIPE, PIPE, PIPE, PIPE, PIPE, PIPE, PIPE, PIPE, PIPE, PIPE, PIPE, PIPE, PIPE, PIPE, PIPE, OUTS, OUTS, OUTS, OUTS},
		{OUTS, PIPE, PIPE, OUTS, PIPE, PIPE, PIPE, PIPE, PIPE, PIPE, PIPE, PIPE, PIPE, PIPE, PIPE, PIPE, OUTS, OUTS, OUTS, OUTS},
		{PIPE, PIPE, PIPE, PIPE, PIPE, PIPE, PIPE, PIPE, PIPE, PIPE, PIPE, PIPE, PIPE, PIPE, UNKN, PIPE, PIPE, PIPE, OUTS, OUTS},
		{PIPE, PIPE, PIPE, PIPE, OUTS, PIPE, PIPE, UNKN, UNKN, UNKN, PIPE, PIPE, PIPE, PIPE, PIPE, PIPE, PIPE, PIPE, PIPE, OUTS},
		{OUTS, OUTS, OUTS, OUTS, PIPE, PIPE, PIPE, UNKN, UNKN, PIPE, PIPE, PIPE, PIPE, PIPE, PIPE, PIPE, PIPE, PIPE, PIPE, PIPE},
		{OUTS, OUTS, OUTS, OUTS, PIPE, PIPE, UNKN, PIPE, PIPE, PIPE, PIPE, PIPE, PIPE, PIPE, UNKN, PIPE, PIPE, PIPE, PIPE, PIPE},
		{OUTS, OUTS, OUTS, OUTS, OUTS, PIPE, PIPE, PIPE, PIPE, PIPE, PIPE, PIPE, PIPE, PIPE, PIPE, PIPE, PIPE, OUTS, PIPE, PIPE},
		{OUTS, OUTS, OUTS, OUTS, PIPE, PIPE, PIPE, PIPE, PIPE, OUTS, PIPE, PIPE, OUTS, PIPE, PIPE, PIPE, PIPE, OUTS, OUTS, OUTS},
		{OUTS, OUTS, OUTS, OUTS, PIPE, PIPE, PIPE, PIPE, PIPE, OUTS, PIPE, PIPE, OUTS, PIPE, PIPE, PIPE, PIPE, OUTS, OUTS, OUTS},
	}

	typeMap := part2Main(&input)

	for i, row := range *typeMap {
		t.Log(row, expected[i])
	}

	for i, row := range expected {
		for j, val := range row {
			if val != (*typeMap)[i][j] {
				t.Errorf("Point %d,%d: got %d, expected %d", i, j, (*typeMap)[i][j], val)
			}
		}
	}
}

func TestInsideLast(t *testing.T) {
	input := [][]rune{
		[]rune("FF7FSF7F7F7F7F7F---7"),
		[]rune("L|LJ||||||||||||F--J"),
		[]rune("FL-7LJLJ||||||LJL-77"),
		[]rune("F--JF--7||LJLJ7F7FJ-"),
		[]rune("L---JF-JLJ.||-FJLJJ7"),
		[]rune("|F|F-JF---7F7-L7L|7|"),
		[]rune("|FFJF7L7F-JF7|JL---7"),
		[]rune("7-L-JL7||F7|L7F-7F7|"),
		[]rune("L.L7LFJ|||||FJL7||LJ"),
		[]rune("L7JLJL-JLJLJL--JLJ.L"),
	}

	typeMap := part2Main(&input)
	got, expected := 0, 10

	for _, row := range *typeMap {
		for _, val := range row {
			if val == UNKN {
				got++
			}
		}
	}
	if got != expected {
		t.Errorf("Expected %d internal points, got %d", expected, got)
	}
}

func TestFromActual(t *testing.T) {
	input := [][]rune{
		[]rune("..F--S-7"),
		[]rune("..|F-7.|"),
		[]rune("..LJ.L7|"),
		[]rune("...F--J|"),
		[]rune("...L---J"),
	}

	expected := [][]Ptype{
		{OUTS, OUTS, PIPE, PIPE, PIPE, PIPE, PIPE, PIPE},
		{OUTS, OUTS, PIPE, PIPE, PIPE, PIPE, UNKN, PIPE},
		{OUTS, OUTS, PIPE, PIPE, OUTS, PIPE, PIPE, PIPE},
		{OUTS, OUTS, OUTS, PIPE, PIPE, PIPE, PIPE, PIPE},
		{OUTS, OUTS, OUTS, PIPE, PIPE, PIPE, PIPE, PIPE},
	}

	typeMap := part2Main(&input)

	for i, row := range *typeMap {
		t.Log(row, expected[i])
	}

	for i, row := range expected {
		for j, val := range row {
			if val != (*typeMap)[i][j] {
				t.Errorf("Point %d,%d: got %d, expected %d", i, j, (*typeMap)[i][j], val)
			}
		}
	}
}

func TestBendInTheRoad(t *testing.T) {
	input := [][]rune{
		[]rune("..F---7"),
		[]rune(".FJF-7|"),
		[]rune(".L-JFJ|"),
		[]rune("..F7L7|"),
		[]rune(".FJ|.||"),
		[]rune("FJ.L-J|"),
		[]rune("L-----S"),
	}
	expected := [][]Ptype{
		{OUTS, OUTS, PIPE, PIPE, PIPE, PIPE, PIPE},
		{OUTS, PIPE, PIPE, PIPE, PIPE, PIPE, PIPE},
		{OUTS, PIPE, PIPE, PIPE, PIPE, PIPE, PIPE},
		{OUTS, OUTS, PIPE, PIPE, PIPE, PIPE, PIPE},
		{OUTS, PIPE, PIPE, PIPE, OUTS, PIPE, PIPE},
		{PIPE, PIPE, UNKN, PIPE, PIPE, PIPE, PIPE},
		{PIPE, PIPE, PIPE, PIPE, PIPE, PIPE, PIPE},
	}

	typeMap := part2Main(&input)

	for i, row := range *typeMap {
		t.Log(row, expected[i])
	}

	for i, row := range expected {
		for j, val := range row {
			if val != (*typeMap)[i][j] {
				t.Errorf("Point %d,%d: got %d, expected %d", i, j, (*typeMap)[i][j], val)
			}
		}
	}
}
