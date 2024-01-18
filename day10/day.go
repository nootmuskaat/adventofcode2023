package day10

import (
	"os"
	"strings"
	"fmt"
)

func Main(f *os.File, part2 bool) {
}

func measureDistances(lines *[]string) [][]int {
	distances := make([][]int, len(*lines))
	var startX, startY int
	for i, line := range *lines {
		distances[i] = make([]int, len(line))
		if s := strings.IndexRune(line, START); s != -1 {
			startX, startY = s, i
		}
	}
	connections(lines, &distances, startX, startY)
	for _, row := range distances {
		fmt.Println(row)
	}
	return distances
}

const START = 'S'
const (
	FROM_ABOVE = "F7|"
	FROM_RIGHT = "J7-"
	FROM_BELOW = "JL|"
	FROM_LEFT  = "FL-"
)

const (
	TO_ABOVE = FROM_BELOW
	TO_RIGHT = FROM_LEFT
	TO_BELOW = FROM_ABOVE
	TO_LEFT  = FROM_RIGHT
)


func connections(lines *[]string, dist *[][]int, startX, startY int) {
	c := rune((*lines)[startY][startX])
	v := (*dist)[startY][startX] + 1
	if look(TO_ABOVE, c) {
		if above := At(lines, startX, startY-1); In(FROM_ABOVE, above) {
			(*dist)[startY-1][startX] = update(v, (*dist)[startY-1][startX])
		}
	}
	if look(TO_RIGHT, c) {
		if right := At(lines, startX+1, startY); In(FROM_RIGHT, right) {
			(*dist)[startY][startX+1] = update(v, (*dist)[startY][startX+1])
		}
	}
	if look(TO_BELOW, c) {
		if below := At(lines, startX, startY+1); In(FROM_BELOW, below) {
			(*dist)[startY+1][startX] = update(v, (*dist)[startY+1][startX])
		}
	}
	if look(TO_LEFT, c) {
		if left := At(lines, startX-1, startY); In(FROM_LEFT, left) {
			(*dist)[startY][startX-1] = update(v, (*dist)[startY][startX-1])
		}
	}
}

func look(s string, r rune) bool {
	return In(s, r) || r == START
}

func In(s string, r rune) bool {
	return strings.IndexRune(s, r) != -1
}

func At(lines *[]string, x, y int) rune {
	if y >= len(*lines) || y < 0 {
		return '.'
	}
	line := (*lines)[y]
	if x > len(line) || x < 0 {
		return '.'
	}
	r := rune(line[x])
	fmt.Printf("[%d, %d] rune %c\n", x, y, r)
	return r
}

func update(a, b int) int {
	if a < b || b == 0 {
		fmt.Println("Set to", a)
		return a
	}
	return b
}
