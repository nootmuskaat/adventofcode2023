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

type Point struct {
	x, y int
}

func (p Point) neighbor(x, y int) Point {
	return Point{p.x + x, p.y + y}
}

func Main(f *os.File, part2 bool) {
	terrain := readFile(f)

	if part2 {
		typeMap := part2Main(terrain)
		inside := 0
		for _, row := range *typeMap {
			for _, t := range row {
				if t == UNKN {
					inside++
				}
			}
		}
		fmt.Println("Number of inside points:", inside)
	} else {
		distances := measureDistances(terrain)
		furthest := 0
		for _, row := range *distances {
			furthest = max(furthest, slices.Max(row))
		}
		fmt.Println("Furthest point:", furthest)
	}
}

func part2Main(terrain *[][]rune) *[][]Ptype {
	typeMap := createTypeMap(terrain, measureDistances(terrain))
	intersections := createIntersections(terrain, typeMap)
	changes := identifyExternalPoints(intersections, typeMap)
	for i := 1; changes > 0; i++ {
		changes = findExternalIntersections(intersections, typeMap) + identifyExternalPoints(intersections, typeMap)
	}
	return typeMap
}

type Ptype uint8

const (
	UNKN Ptype = iota
	PIPE
	OUTS
)

func createTypeMap(terrain *[][]rune, distances *[][]int) *[][]Ptype {
	typeMap := make([][]Ptype, len(*distances))
	for i, row := range *distances {
		typeMap[i] = make([]Ptype, len(row))
		for j, dist := range row {
			if dist > 0 || (*terrain)[i][j] == START {
				typeMap[i][j] = PIPE
			} else {
				typeMap[i][j] = UNKN
			}
		}
	}
	return &typeMap
}

type Itype uint8

const (
	OPEN_NORTH Itype = 1 << 0
	OPEN_SOUTH       = 1 << 1
	OPEN_EAST        = 1 << 2
	OPEN_WEST        = 1 << 3
	EXTERNAL         = 1 << 4
)

type Isect struct {
	val uint8
}

func (i Isect) String() string {
	chars := make([]rune, 0, 5)
	if i.Is(EXTERNAL) {
		chars = append(chars, 'X')
	}
	if i.Is(OPEN_NORTH) {
		chars = append(chars, 'N')
	}
	if i.Is(OPEN_SOUTH) {
		chars = append(chars, 'S')
	}
	if i.Is(OPEN_EAST) {
		chars = append(chars, 'E')
	}
	if i.Is(OPEN_WEST) {
		chars = append(chars, 'W')
	}
	return string(chars)
}

func (i Isect) Is(t Itype) bool {
	return i.val&uint8(t) == uint8(t)
}

func (i *Isect) Set(t Itype) {
	i.val += uint8(t)
}

const (
	CONNECT_EAST  string = "LF-S"
	CONNECT_WEST         = "J7-S"
	CONNECT_NORTH        = "LJ|S"
	CONNECT_SOUTH        = "F7|S"
)

// given the initial map (terrain) and an indication of what is and is not
// part of the actual pipe (types), determine in what directions one can move
// from that points, e.g.
//
//	JL
//	--
//
// This 2x2 block creates an intersection where one can move North, East, and West
// assuming all 4 blocks are themselves part of the pipe structure.
func createIntersections(terrain *[][]rune, types *[][]Ptype) *[][]Isect {

	ter := func(p Point) rune {
		return valueAt(terrain, p, '.')
	}
	typ := func(p Point) Ptype {
		return valueAt(types, p, OUTS)
	}

	intersections := make([][]Isect, len(*terrain)-1)
	for y := range intersections {
		intersections[y] = make([]Isect, len((*terrain)[y])-1)
		for x := range intersections[y] {
			isect := Isect{0}
			p1, p2, p3, p4 := Point{x, y}, Point{x + 1, y}, Point{x, y + 1}, Point{x + 1, y + 1}
			t1, t2, t3, t4 := typ(p1), typ(p2), typ(p3), typ(p4)
			c1, c4 := ter(p1), ter(p4)
			if t1 == UNKN || t2 == UNKN || !In(CONNECT_EAST, c1) {
				isect.Set(OPEN_NORTH)
			}
			if t1 == UNKN || t3 == UNKN || !In(CONNECT_SOUTH, c1) {
				isect.Set(OPEN_WEST)
			}
			if t3 == UNKN || t4 == UNKN || !In(CONNECT_WEST, c4) {
				isect.Set(OPEN_SOUTH)
			}
			if t2 == UNKN || t4 == UNKN || !In(CONNECT_NORTH, c4) {
				isect.Set(OPEN_EAST)
			}
			intersections[y][x] = isect
		}
	}
	return &intersections
}

// based on neighboring intersections and adjacent points, mark intersections as external
func findExternalIntersections(intersections *[][]Isect, typeMap *[][]Ptype) (changed int) {
	tp := func(p Point) Ptype {
		return valueAt(typeMap, p, OUTS)
	}
	isExt := func(p Point) bool {
		return valueAt(intersections, p, Isect{0b11111}).Is(EXTERNAL)
	}
	// are any of the four of points that make up an intersection external
	containsExternal := func(thisPoint Point) bool {
		p2, p3, p4 := thisPoint.neighbor(1, 0), thisPoint.neighbor(0, 1), thisPoint.neighbor(1, 1)
		return tp(thisPoint) == OUTS || tp(p2) == OUTS || tp(p3) == OUTS || tp(p4) == OUTS
	}

	width, height := len((*intersections)[0]), len(*intersections)
	for row := 0; row < height; row++ {
		for x, thisIsect := range (*intersections)[row] {
			if thisIsect.Is(EXTERNAL) {
				continue
			}
			thisPoint := Point{x, row}
			if containsExternal(thisPoint) || thisIsect.Is(OPEN_WEST) && isExt(thisPoint.neighbor(-1, 0)) {
				thisIsect.Set(EXTERNAL)
				set(intersections, thisPoint, thisIsect)
				changed++
			}
		}
		for x := width - 1; x >= 0; x-- {
			thisPoint := Point{x, row}
			thisIsect := (*intersections)[row][x]
			if thisIsect.Is(EXTERNAL) {
				continue
			}
			if containsExternal(thisPoint) || thisIsect.Is(OPEN_EAST) && isExt(thisPoint.neighbor(1, 0)) {
				thisIsect.Set(EXTERNAL)
				set(intersections, thisPoint, thisIsect)
				changed++
			}
		}
	}
	for col := 0; col < width; col++ {
		for y := range *intersections {
			thisPoint := Point{col, y}
			thisIsect := (*intersections)[y][col]
			if thisIsect.Is(EXTERNAL) {
				continue
			}
			if containsExternal(thisPoint) || thisIsect.Is(OPEN_NORTH) && isExt(thisPoint.neighbor(0, -1)) {
				thisIsect.Set(EXTERNAL)
				set(intersections, thisPoint, thisIsect)
				changed++
			}
		}
		for y := height - 1; y >= 0; y-- {
			thisPoint := Point{col, y}
			thisIsect := (*intersections)[y][col]
			if thisIsect.Is(EXTERNAL) {
				continue
			}
			if containsExternal(thisPoint) || thisIsect.Is(OPEN_SOUTH) && isExt(thisPoint.neighbor(0, 1)) {
				thisIsect.Set(EXTERNAL)
				set(intersections, thisPoint, thisIsect)
				changed++
			}
		}
	}
	return
}

// based on neighboring items and adjacent intersections, mark points as external
func identifyExternalPoints(intersections *[][]Isect, typeMap *[][]Ptype) (changed int) {

	tp := func(p Point) Ptype {
		return valueAt(typeMap, p, OUTS)
	}
	isect := func(p Point) Isect {
		return valueAt(intersections, p, Isect{0b11111})
	}
	update := func(this, l, r, c, li, ri Point) {
		if tp(l) == OUTS || tp(r) == OUTS || tp(c) == OUTS || isect(li).Is(EXTERNAL) || isect(ri).Is(EXTERNAL) {
			set(typeMap, this, OUTS)
			changed++
		}
	}

	width, height := len((*typeMap)[0]), len(*typeMap)
	for row := 0; row < height; row++ {
		for x, thisType := range (*typeMap)[row] {
			if thisType != UNKN {
				continue
			}
			thisPoint := Point{x, row}
			l, r, c := thisPoint.neighbor(-1, -1), thisPoint.neighbor(-1, 1), thisPoint.neighbor(-1, 0)
			update(thisPoint, l, r, c, l, c)
		}
		for x := width - 1; x >= 0; x-- {
			thisPoint := Point{x, row}
			if tp(thisPoint) != UNKN {
				continue
			}
			l, r, c := thisPoint.neighbor(1, 1), thisPoint.neighbor(1, -1), thisPoint.neighbor(1, 0)
			update(thisPoint, l, r, c, thisPoint, thisPoint.neighbor(0, -1))
		}
	}

	for col := 0; col < width; col++ {
		for y := range *typeMap {
			thisPoint := Point{col, y}
			if tp(thisPoint) != UNKN {
				continue
			}
			l, r, c := thisPoint.neighbor(1, -1), thisPoint.neighbor(-1, -1), thisPoint.neighbor(0, -1)
			update(thisPoint, l, r, c, c, r)
		}
		for y := height - 1; y >= 0; y-- {
			thisPoint := Point{col, y}
			if tp(thisPoint) != UNKN {
				continue
			}
			l, r, c := thisPoint.neighbor(-1, 1), thisPoint.neighbor(1, 1), thisPoint.neighbor(0, 1)
			update(thisPoint, l, r, c, thisPoint.neighbor(-1, 0), thisPoint)
		}
	}
	return
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
