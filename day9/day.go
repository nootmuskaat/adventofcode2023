package day9

import (
	"bufio"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Main(f *os.File, reverse bool) {
	total := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if err := scanner.Err(); err != nil {
			panic(err)
		}
		values := make([]int, 0, 8)
		for _, p := range strings.Fields(line) {
			if v, e := strconv.Atoi(p); e != nil {
				panic(e)
			} else {
				values = append(values, v)
			}
		}
		if reverse {
			slices.Reverse(values)
		}
		total += DrillDown(values)
	}
	log.Println("Total", total)
}

func DrillDown(values []int) int {
	diffs := make([]int, 0, len(values))
	allZeroes := true
	for i, v := range values[1:] {
		diff := v - values[i]
		allZeroes = allZeroes && (diff == 0)
		diffs = append(diffs, diff)
	}
	if allZeroes {
		return values[len(values)-1]
	} else {
		return values[len(values)-1] + DrillDown(diffs)
	}
}
