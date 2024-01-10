package day5

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const FILENAME = "./static/day5.txt"
const MaxUint = ^uint(0)

var fieldOrder []string = []string{
	"seed", "soil", "fertilizer", "water", "light", "temperature", "humidity", "location",
}

func Main(part2 bool) {
	seeds, maps := readFile(FILENAME)
	location := MaxUint

	if !part2 {
		// Part 1
		// What is the lowest location number that corresponds to any of the initial seed numbers?
		// you'll need to convert each seed number through other categories until you can find its
		// corresponding location number. In this example, the corresponding types are:
		// Seed 79->soil 81->fertilizer 81->water 81->light 74->temperature 78->humidity 78->location 82.
		for _, val := range seeds {
			val = RunForwards(val, &maps)
			if val < location {
				location = val
			}
		}
	} else {
		seedRanges := FullRange{}
		for i, j := 0, 1; i < len(seeds); {
			seedRanges.Append(Range{seeds[i], seeds[i], seeds[j]})
			i, j = i+2, j+2
		}
		location = 0
		factor := uint(10000)
		// hone in on target value one order of magnitude at a time
		for {
			if seedVal := RunBackwards(location, &maps); seedRanges.Contains(seedVal) {
				if factor == 1 {
					break
				} else {
					location -= factor
					factor /= 10
				}
			}
			location += factor
		}
	}
	log.Println("Lowest:", location)
}

func RunForwards(num uint, maps *map[string]FullRange) uint {
	for idx, factor := range fieldOrder[1:] {
		key := fmt.Sprintf("%s-to-%s", fieldOrder[idx], factor)
		num = (*maps)[key].Map(num)
	}
	return num
}

func RunBackwards(num uint, maps *map[string]FullRange) uint {
	for i := len(fieldOrder) - 1; i > 0; i-- {
		// fmt.Printf("%d\t->\t", val)
		key := fmt.Sprintf("%s-to-%s", fieldOrder[i-1], fieldOrder[i])
		num = (*maps)[key].ReverseMap(num)
		// fmt.Printf("%s\t->\t%d\n", key, val)
	}
	return num
}

type Range struct {
	dest, source, length uint
}

func (r Range) Contains(num uint) bool {
	return num >= r.source && num < (r.source+r.length)
}

func (r Range) ReverseContains(num uint) bool {
	return num >= r.dest && num < (r.dest+r.length)
}

func (r Range) Map(num uint) (uint, error) {
	if !r.Contains(num) {
		return 0, errors.New("Num not in range")
	}
	return r.dest + (num - r.source), nil
}

func (r Range) ReverseMap(num uint) (uint, error) {
	if !r.ReverseContains(num) {
		return 0, errors.New("Num not in range")
	}
	return r.source + (num - r.dest), nil
}

func RangeFromLine(s *string) (Range, error) {
	parts := strings.Split(*s, " ")
	ints := make([]uint, 3)
	for idx, part := range parts {
		if i, err := strconv.ParseUint(part, 10, 64); err != nil {
			return Range{}, err
		} else {
			ints[idx] = uint(i)
		}
	}
	if len(ints) == 3 {
		return Range{ints[0], ints[1], ints[2]}, nil
	} else {
		return Range{}, errors.New(fmt.Sprintf("Line contained %d ints", len(ints)))
	}
}

type FullRange struct {
	ranges []Range
}

func (fr *FullRange) Append(r Range) {
	fr.ranges = append(fr.ranges, r)
}

func (fr FullRange) Contains(num uint) bool {
	for _, r := range fr.ranges {
		if r.Contains(num) {
			return true
		}
	}
	return false
}

func (fr FullRange) Map(num uint) uint {
	for _, r := range fr.ranges {
		if n, err := r.Map(num); err == nil {
			return n
		}
	}
	return num
}

func (fr FullRange) ReverseMap(num uint) uint {
	for _, r := range fr.ranges {
		if n, err := r.ReverseMap(num); err == nil {
			return n
		}
	}
	return num
}

func readFile(filename string) ([]uint, map[string]FullRange) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()
	seeds := make([]uint, 0, 16)
	ranges := make(map[string]FullRange)
	key := ""
	fr := FullRange{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if err := scanner.Err(); err != nil {
			log.Fatalln(err)
		}
		if strings.HasPrefix(line, "seeds:") {
			readSeeds(&line, &seeds)
		} else if strings.Contains(line, "map") {
			if len(key) > 0 {
				ranges[key] = fr
			}
			key = strings.Split(line, " ")[0]
			fr = FullRange{}
		} else if len(line) > 0 {
			if r, e := RangeFromLine(&line); e != nil {
				log.Fatalln(e)
			} else {
				fr.Append(r)
			}
		}
	}
	ranges[key] = fr
	return seeds, ranges
}

func readSeeds(line *string, seeds *[]uint) {
	ints := strings.TrimSpace(strings.Split(*line, ":")[1])
	for _, i := range strings.Split(ints, " ") {
		if len(i) == 0 {
			continue
		}
		u, e := strconv.ParseUint(i, 10, 64)
		if e != nil {
			log.Fatalf("Could not parse %s", i)
		}
		*seeds = append(*seeds, uint(u))
	}
}
