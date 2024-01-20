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
	if part2 {
		setStartTo1(terrain, distances)
		typeMap := convertToTypeMap(distances)
		for changes := 1; changes > 0; {
			changes = identifyOutside(terrain, typeMap) + findLeakage(terrain, typeMap)
		}
	} else {
		for _, row := range *distances {
			furthest = max(furthest, slices.Max(row))
		}
	}
	fmt.Println("Furthest point:", furthest)
}

type PType uint8

const (
	UNKN PType = iota
	PIPE
	SEAP
	OUTS
)

func convertToTypeMap(distances *[][]int) *[][]PType {
	typeMap := make([][]PType, len(*distances))
	for i, row := range *distances {
		typeMap[i] = make([]PType, len(row))
		for j, dist := range row {
			if dist > 0 {
				typeMap[i][j] = PIPE
			} else {
				typeMap[i][j] = UNKN
			}
		}
	}
	return &typeMap
}

type Point struct {
	x, y int
}

func (p Point) neighbor(x, y int) Point {
	return Point{p.x + x, p.y + y}
}

func findLeakage(terrain *[][]rune, typeMap *[][]PType) (changed int) {
	// TODO!
	return
}

func identifyOutside(terrain *[][]rune, typeMap *[][]PType) (changed int) {
	// candidates := countCandidates(distances)

	// Moving outside in, mark as external any point that's currently 0 (undefined)
	// and whose outside neighbor is -1 (exteranl)
	iters := max(len((*typeMap)[0]), len(*typeMap)) - 1
	width, height := len((*typeMap)[0]), len(*typeMap)
	for offset := 0; offset < iters; offset++ {
		bottom := height - offset - 1
		right := width - offset - 1
		for x, v := range (*typeMap)[offset] {
			if neigh := valueAt(typeMap, Point{x, offset - 1}, OUTS); v == UNKN && neigh == OUTS {
				set(typeMap, Point{x, offset}, OUTS)
				changed++
			}
		}
		for x, v := range (*typeMap)[bottom] {
			if neigh := valueAt(typeMap, Point{x, height + 1}, OUTS); v == UNKN && neigh == OUTS {
				set(typeMap, Point{x, height}, OUTS)
				changed++
			}
		}
		for y, row := range *typeMap {
			if neigh := valueAt(typeMap, Point{offset - 1, y}, OUTS); row[offset] == UNKN && neigh == OUTS {
				set(typeMap, Point{offset, y}, OUTS)
				changed++
			}
			if neigh := valueAt(typeMap, Point{right + 1, y}, OUTS); row[right] == UNKN && neigh == OUTS {
				set(typeMap, Point{right, y}, OUTS)
				changed++
			}
		}
	}
	return
}

func countCandidates(distances *[][]int) (candidates int) {
	for _, row := range *distances {
		for _, v := range row {
			if v == 0 {
				candidates++
			}
		}
	}
	return
}

func setStartTo1(terrain *[][]rune, distances *[][]int) {
	for y, row := range *terrain {
		for x, r := range row {
			if r == START {
				update(distances, Point{x, y}, 1)
				return
			}
		}
	}
}

func measureDistances(lines *[][]rune) *[][]int {
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
	return &distances
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
	c := valueAt(lines, start, '.')
	v := valueAt(dist, start, -1) + 1
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
			if In(nd.connectables, valueAt(lines, nd.neighbor, '.')) {

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

func valueAt[T any](board *[][]T, point Point, defaultValue T) T {
	if point.y >= len(*board) || point.y < 0 {
		return defaultValue
	}
	line := (*board)[point.y]
	if point.x >= len(line) || point.x < 0 {
		return defaultValue
	}
	return line[point.x]
}

func set[T any](board *[][]T, point Point, val T) {
	(*board)[point.y][point.x] = val
}

func update(board *[][]int, point Point, v int) bool {
	current := valueAt(board, point, -1)
	if current == 0 || current > v {
		set(board, point, v)
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
