package day3

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

const DAY3_FILE = "./static/day3.txt"

func Main() {
	values, symbols := whatIsWhere(readFile())

	var sum uint

	// Part 1
	// Any number adjacent to a symbol, even diagonally, is a "part number"
	// (Periods (.) do not count as a symbol.)
	// What is the sum of all of the part numbers in the engine schematic?

	// for _, symbol := range *symbols {
	// 	for _, value := range *values {
	// 		if symbol.Neighbors(&value) && !value.relevant {
	// 			value.relevant = true
	// 			sum += value.value
	// 		}
	// 	}
	// }

	// Part 2
	// A gear is any * symbol that is adjacent to exactly two part numbers.
	// Its gear ratio is the result of multiplying those two numbers together.
	// What is the sum of all of the gear ratios in your engine schematic?
	neighbors := make(map[Point][]Value)

	for _, gear := range *symbols {
		for _, value := range *values {
			if gear.Neighbors(&value) {
				neighbors[gear] = append(neighbors[gear], value)
			}
		}
	}

	for _, values := range neighbors {
		if len(values) == 2 {
			sum += values[0].value * values[1].value
		}
	}

	fmt.Println("Sum", sum)
}

type Point struct {
	x, y int16
}

type Value struct {
	value    uint     // the actual numerical value
	points   *[]Point // the points in space occupied by the value
	relevant bool     // exists to prevent double counting
}

func NewValue(value uint, starts Point) Value {
	width := int16(math.Log10(float64(value))) + 1 // e.g. 999 -> 3, 1000 -> 4
	points := make([]Point, 0, width)
	var i int16 = 0
	for ; i < width; i++ {
		points = append(points, Point{y: starts.y, x: starts.x + i})
	}
	return Value{value, &points, false}
}

func (p Point) Neighbors(value *Value) bool {
	abs := func(i int16) int16 {
		return int16(math.Abs(float64(i)))
	}

	for _, point := range *value.points {
		if abs(p.x-point.x) <= 1 && abs(p.y-point.y) <= 1 {
			return true
		}
	}
	return false
}

// Build a mapping of what numerical values are where
// and what non-. symbols are where
func whatIsWhere(lines *[]string) (*[]Value, *[]Point) {
	values := make([]Value, 0, 16)
	symbols := make([]Point, 0, 16)

	intValue := func(c rune) uint {
		return uint(c - '0')
	}

	var starting Point
	var value uint = 0

	for y, line := range *lines {
		for x, chr := range line {
			if uint('0') <= uint(chr) && uint(chr) <= uint('9') {
				if value > 0 {
					value = 10*value + intValue(chr)
				} else {
					value = intValue(chr)
					starting = Point{int16(x), int16(y)}
				}
			} else {
				if value > 0 {
					values = append(values, NewValue(value, starting))
					value = 0
				}
				// if chr != '.' {  // Part 1 version
				if chr == '*' { // Part 2 version
					symbols = append(symbols, Point{int16(x), int16(y)})
				}
			}
		}
	}
	return &values, &symbols

}

func readFile() *[]string {
	f, err := os.Open(DAY3_FILE)
	if err != nil {
		log.Fatal(err)
	}
	lines := make([]string, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		lines = append(lines, line)
	}
	return &lines
}
