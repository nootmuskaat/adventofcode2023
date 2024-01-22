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
	display := map[PType]rune {
		UNKN: ' ',
		OUTS: 'X',
		LEKL: '<',
		LEKR: '>',
	}


	terrain := readFile(f)
	distances := measureDistances(terrain)
	if part2 {
		setStartTo1(terrain, distances)
		typeMap := convertToTypeMap(distances)
		for changes := 1; changes > 0; {
			changes = identifyOutside(terrain, typeMap) + findLeakage(terrain, typeMap)
		}
		inside := 0
		for y, row := range *typeMap {
			for x, t := range row {
				if t == PIPE {
					fmt.Printf("%c", (*terrain)[y][x])
				} else {
					fmt.Printf("%c", display[t])
				}
				if t == UNKN {
					inside++
				}
			}
			fmt.Println()
		}
		fmt.Println("Number of inside points:", inside)
	} else {
		furthest := 0
		for _, row := range *distances {
			furthest = max(furthest, slices.Max(row))
		}
		fmt.Println("Furthest point:", furthest)
	}
}

type PType uint8

const (
	UNKN PType = iota
	PIPE
	LEKL  // left side of leak
	LEKR  // right side
	LEKB  // bend if leak, i.e. a fork or elbow
	OUTS
)

type Isect uint8
const (
	OPEN_NORTH Isect = 1 << 0
	OPEN_SOUTH       = 1 << 1
	OPEN_EAST        = 1 << 2
	OPEN_WEST        = 1 << 3
	EXTERNAL         = 1 << 4
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


const (
	CONNECT_EAST string = "LF-S"
	CONNECT_WEST        = "J7-S"
	CONNECT_NORTH       = "LJ|S"
	CONNECT_SOUTH       = "F7|S"
)

// Moving outside in, mark as LEKL or LEKR any two points where outside can
// leak beyond the initial parimeter
// TODO - possibly: implement elbows
func findLeakage(terrain *[][]rune, typeMap *[][]PType) (changed int) {
	width, height := len((*typeMap)[0]), len(*typeMap)
	for offset := 0; offset < height; offset++ {
		bottom := height - offset - 1
		// top moving down
		for rx, lx := 0, 1; lx < len((*typeMap)[offset]); rx, lx = lx, lx+1 {
			rp, lp := Point{rx, offset}, Point{lx, offset}
			rn, ln := Point{rx, offset-1}, Point{lx, offset-1}
			if possibleLeak := leakSite(typeMap, lp, rp, ln, rn); !possibleLeak {
				continue
			}
			rChar, lChar := valueAt(terrain, rp, '.'), valueAt(terrain, lp, '.')
			if In(CONNECT_WEST, lChar) && In(CONNECT_EAST, rChar) {
				continue
			}
			// fmt.Printf("T! LEAK %v %c %v %c\n", lp, lChar, rp, rChar)
			set(typeMap, rp, LEKR)
			set(typeMap, lp, LEKL)
			changed++
		}
		// bottom moving up
		for lx, rx := 0, 1; rx < len((*typeMap)[bottom]); lx, rx = rx, rx+1 {
			rp, lp := Point{rx, bottom}, Point{lx, bottom}
			rn, ln := Point{rx, bottom+1}, Point{lx, bottom+1}
			if possibleLeak := leakSite(typeMap, lp, rp, ln, rn); !possibleLeak {
				continue
			}
			rChar, lChar := valueAt(terrain, rp, '.'), valueAt(terrain, lp, '.')
			if In(CONNECT_EAST, lChar) && In(CONNECT_WEST, rChar) {
				continue
			}
			// fmt.Printf("B! LEAK %v %c %v %c\n", lp, lChar, rp, rChar)
			set(typeMap, rp, LEKR)
			set(typeMap, lp, LEKL)
			changed++
		}
	}
	for offset := 0; offset < width; offset++ {
		right := width - offset - 1
		// left moving right
		for ly, ry := 0, 1; ry < len(*typeMap); ly, ry = ry, ry+1 {
			rp, lp := Point{offset, ry}, Point{offset, ly}
			rn, ln := Point{offset-1, ry}, Point{offset-1, ly}
			if possibleLeak := leakSite(typeMap, lp, rp, ln, rn); !possibleLeak {
				continue
			}
			rChar, lChar := valueAt(terrain, rp, '.'), valueAt(terrain, lp, '.')
			if In(CONNECT_SOUTH, lChar) && In(CONNECT_NORTH, rChar) {
				continue
			}
			// fmt.Printf("L! LEAK %v %c %v %c\n", lp, lChar, rp, rChar)
			set(typeMap, rp, LEKR)
			set(typeMap, lp, LEKL)
			changed++
		}
		// right moving left
		for ry, ly := 0, 1; ly < len(*typeMap); ry, ly = ly, ly+1 {
			rp, lp := Point{right, ry}, Point{right, ly}
			rn, ln := Point{right+1, ry}, Point{right+1, ly}
			if possibleLeak := leakSite(typeMap, lp, rp, ln, rn); !possibleLeak {
				continue
			}
			rChar, lChar := valueAt(terrain, rp, '.'), valueAt(terrain, lp, '.')
			if In(CONNECT_NORTH, lChar) && In(CONNECT_SOUTH, rChar) {
				continue
			}
			// fmt.Printf("R! LEAK %v %c %v %c\n", lp, lChar, rp, rChar)
			set(typeMap, rp, LEKR)
			set(typeMap, lp, LEKL)
			changed++
		}
	}
	return
}


// Examine a 2x2 block of map to determine if a leak is even worth assessing
// For example, the below set of blocks, evaluated bottom up
//
// row A:  F7
// row B:  ||
// row C:  J|
// row D:  .L
// row E:  ..
//
// Assuming the points on row E are determined as external,
// row C would be determined as a leak potential, due to the left neighbor
// being an external point. Similarly rows B and A would need to be passed for
// further evaluation as the points below them would be assessed ultimately as leak
// points. Row A will ultimately close the loop, but this is evaluated outside this func
func leakSite(typeMap *[][]PType, leftPoint, rightPoint, leftNeighbor, rightNeighbor Point) bool {
	// confirm we are looking at pipe elements on both sides
	rType := valueAt(typeMap, rightPoint, OUTS)
	lType := valueAt(typeMap, leftPoint, OUTS)
	if rType != PIPE || lType != PIPE {
		return false
	}
	// confirm that we are at a possible leak point to begin with
	rNeighType := valueAt(typeMap, rightNeighbor, OUTS)
	lNeighType := valueAt(typeMap, leftNeighbor, OUTS)
	return lNeighType == OUTS || rNeighType == OUTS || (lNeighType == LEKL && rNeighType == LEKR)
	// if p {
	// 	fmt.Println("Potential leak @", leftPoint, rightPoint)
	// }
	// return p
}

// Moving outside in, mark as external any point that's currently 0 (undefined)
// and whose outside neighbor is -1 (exteranl)
func identifyOutside(terrain *[][]rune, typeMap *[][]PType) (changed int) {

	val := func(p Point) PType {
		return valueAt(typeMap, p, OUTS)
	}

	width, height := len((*typeMap)[0]), len(*typeMap)
	for offset := 0; offset < height; offset++ {
		bottom := height - offset - 1
		for x, v := range (*typeMap)[offset] {
			// check neighbors left, right and center
			l, r, c := val(Point{x+1, offset-1}), val(Point{x-1, offset-1}), val(Point{x, offset-1})
			if v == UNKN && external(l, r, c) {
				set(typeMap, Point{x, offset}, OUTS)
				changed++
			}
		}
		for x, v := range (*typeMap)[bottom] {
			l, r, c := val(Point{x-1, bottom+1}), val(Point{x+1, bottom+1}), val(Point{x, bottom+1})
			if v == UNKN && external(l, r, c) {
				set(typeMap, Point{x, bottom}, OUTS)
				changed++
			}
		}
	}
	for offset := 0; offset < width; offset++ {
		right := width - offset - 1
		for y, row := range *typeMap {
			l, r, c := val(Point{offset-1, y-1}), val(Point{offset-1, y+1}), val(Point{offset-1, y})
			if row[offset] == UNKN && external(l, r, c) {
				set(typeMap, Point{offset, y}, OUTS)
				changed++
			}
			l, r, c = val(Point{right+1, y+1}), val(Point{right+1, y-1}), val(Point{right+1, y})
			if row[right] == UNKN && external(l, r, c) {
				set(typeMap, Point{right, y}, OUTS)
				changed++
			}
		}
	}
	return
}

func external(left, right, center PType) bool {
	if center == OUTS || left == OUTS || right == OUTS{
		return true
	}
	return (left == LEKL && center == LEKR) || (center == LEKL && right == LEKR)
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
