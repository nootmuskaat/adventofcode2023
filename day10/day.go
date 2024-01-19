package day10

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
	"sync"
	"time"
)

func Main(f *os.File, part2 bool) {
	terrain := readFile(f)
	furthest := 0
	distances := measureDistances(terrain)
	for _, row := range distances {
		furthest = max(furthest, slices.Max(row))
	}
	fmt.Println("Furthest point:", furthest)
}

type Point struct {
	x, y int
}

func (p Point) neighbor(x, y int) Point {
	return Point{p.x + x, p.y + y}
}

func measureDistances(lines *[][]rune) [][]int {
	distances := make([][]int, len(*lines))
	var start Point
	for i, line := range *lines {
		distances[i] = make([]int, len(line))
		if s := slices.Index(line, START); s != -1 {
			start = Point{s, i}
		}
	}
	todos := make(chan Point, 4)
	var wg sync.WaitGroup
	go connections(lines, &distances, start, todos)
	exit, done := false, make(chan bool)
	go func() {
		time.Sleep(100 * time.Millisecond)
		wg.Wait()
		done <- true
	}()

	for !exit {
		select {
		case p := <-todos:
			wg.Add(1)
			go func() {
				defer wg.Done()
				connections(lines, &distances, p, todos)
			}()
		case <-done:
			exit = true
		}
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

func connections(lines *[][]rune, dist *[][]int, start Point, todos chan Point) {
	c := At(lines, start)
	v := Val(dist, start) + 1
	directions := []struct {
		neighbor                  Point
		connectables, inDirection string
	}{
		{start.neighbor(0, -1), FROM_ABOVE, TO_ABOVE},
		{start.neighbor(1, 0), FROM_RIGHT, TO_RIGHT},
		{start.neighbor(0, 1), FROM_BELOW, TO_BELOW},
		{start.neighbor(-1, 0), FROM_LEFT, TO_LEFT},
	}

	for _, nd := range directions {
		if shouldLook(nd.inDirection, c) {
			// does the neighboring rune 'connect' to our space
			if In(nd.connectables, At(lines, nd.neighbor)) {

				if update(dist, nd.neighbor, v) {
					todos <- nd.neighbor
				}
			}
		}
	}
}

func shouldLook(s string, r rune) bool {
	return In(s, r) || r == START
}

func In(s string, r rune) bool {
	return strings.IndexRune(s, r) != -1
}

func At(lines *[][]rune, point Point) rune {
	if point.y >= len(*lines) || point.y < 0 {
		return '.'
	}
	line := (*lines)[point.y]
	if point.x > len(line) || point.x < 0 {
		return '.'
	}
	return line[point.x]
}

func Val(board *[][]int, point Point) int {
	if point.y >= len(*board) || point.y < 0 {
		return -1
	}
	line := (*board)[point.y]
	if point.x > len(line) || point.x < 0 {
		return -1
	}
	return line[point.x]
}

func update(board *[][]int, point Point, v int) bool {
	current := Val(board, point)
	if current == 0 || current > v {
		(*board)[point.y][point.x] = v
		return true
	}
	return false
}

func readFile(f *os.File) *[][]rune {
	output := make([][]rune, 0, 64)

	scan := bufio.NewScanner(f)
	for scan.Scan() {
		line := scan.Text()
		if err := scan.Err(); err != nil {
			panic(err)
		}
		output = append(output, []rune(line))
	}
	return &output
}
