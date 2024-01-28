package day11

import (
	"bufio"
	"fmt"
	"os"
)

type GalaxyMap [][]bool

const Galaxy = '#'

func Main(f *os.File, part2 bool) {
	expansionFactor := 2
	if part2 {
		expansionFactor = 1_000_000
	}
	gmap := readFile(f)
	totalDistance := trackDistances(gmap, expansionFactor)
	fmt.Println("Total distance:", totalDistance)
}

type Point struct {
	row, col int
}

func trackDistances(gmap *GalaxyMap, expansionFactor int) (totalDistance int) {
	expansionFactor--
	nonEmptyCols := make([]bool, len((*gmap)[0]))
	nonEmptyRows := make([]bool, len(*gmap))
	galaxies := make([]Point, 0, 16)
	for rowNo, row := range *gmap {
		for colNo, val := range row {
			if val {
				nonEmptyCols[colNo] = true
				nonEmptyRows[rowNo] = true
				galaxies = append(galaxies, Point{rowNo, colNo})
			}
		}
	}

	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			ga, gb := galaxies[i], galaxies[j]
			totalDistance += absDiff(ga.row, gb.row) + expansionFactor*blanks(ga.row, gb.row, nonEmptyRows)
			totalDistance += absDiff(ga.col, gb.col) + expansionFactor*blanks(ga.col, gb.col, nonEmptyCols)
		}
	}

	return
}

func absDiff(a, b int) int {
	if b > a {
		return b - a
	}
	return a - b
}

func blanks(pa, pb int, occupiedSpots []bool) (blankPoints int) {
	from, to := min(pa, pb), max(pa, pb)
	for _, isOccupied := range occupiedSpots[from:to] {
		if !isOccupied {
			blankPoints++
		}
	}
	return
}

func readFile(f *os.File) *GalaxyMap {
	gmap := make(GalaxyMap, 0, 32)

	scan := bufio.NewScanner(f)
	for scan.Scan() {
		line := scan.Text()
		row := make([]bool, 0, 32)
		if err := scan.Err(); err != nil {
			panic(err)
		}
		for _, c := range line {
			row = append(row, c == Galaxy)
		}
		gmap = append(gmap, row)
	}
	return &gmap
}
