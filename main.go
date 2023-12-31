package main

import (
	"flag"
	"fmt"

	day1 "nootmuskaat/adventofcode2023/day1"
	day2 "nootmuskaat/adventofcode2023/day2"
	day3 "nootmuskaat/adventofcode2023/day3"
	day4 "nootmuskaat/adventofcode2023/day4"
)

func main() {
	days := map[uint]func(bool){
		1: day1.Main,
		2: day2.Main,
		3: day3.Main,
		4: day4.Main,
	}

	day := flag.Uint("day", uint(len(days)), "The day to run")
	isPartTwo := flag.Bool("part2", false, "If true, run the second part of the day's task")
	flag.Parse()

	fmt.Printf("Running day %d - part two %v\n", *day, *isPartTwo)

	days[*day](*isPartTwo)
}
