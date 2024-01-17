package main

import (
	"flag"
	"fmt"
	"os"

	day1 "nootmuskaat/adventofcode2023/day1"
	day2 "nootmuskaat/adventofcode2023/day2"
	day3 "nootmuskaat/adventofcode2023/day3"
	day4 "nootmuskaat/adventofcode2023/day4"
	day5 "nootmuskaat/adventofcode2023/day5"
	day6 "nootmuskaat/adventofcode2023/day6"
	day7 "nootmuskaat/adventofcode2023/day7"
	day8 "nootmuskaat/adventofcode2023/day8"
)

func main() {
	days := map[uint]func(*os.File, bool){
		1: day1.Main,
		2: day2.Main,
		3: day3.Main,
		4: day4.Main,
		5: day5.Main,
		6: day6.Main,
		7: day7.Main,
		8: day8.Main,
	}

	day := flag.Uint("day", uint(len(days)), "The day to run")
	isPartTwo := flag.Bool("part2", false, "If true, run the second part of the day's task")
	filename := flag.String("f", "", "File to use, if not './static/day${dayNumber}.txt'")
	flag.Parse()

	if len(*filename) == 0 {
		*filename = fmt.Sprintf("./static/day%d.txt", *day)
	}

	f, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fmt.Printf("Running day %d (%s) - part two %v\n", *day, *filename, *isPartTwo)

	days[*day](f, *isPartTwo)
}
