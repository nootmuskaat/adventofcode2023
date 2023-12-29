package day1

import (
	"bufio"
	"log"
	"os"
	"unicode"
)

const DAY1_FILE = "static/day1.txt"

func Main() {
	f, err := os.Open(DAY1_FILE)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	var line_values []int
	var line_value int

	for scanner.Scan() {
		line := scanner.Text()
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		line_value = 10*firstInt(line) + lastInt(line)
		// log.Printf("Line %s yielded value %d", line, line_value)
		line_values = append(line_values, line_value)
	}

	sum := sumValues(&line_values)
	log.Printf("Sum: %d\n", sum)
}

func firstInt(line string) int {
	zero := '0'
	for _, chr := range line {
		if unicode.IsDigit(chr) {
			return int(chr - zero)
		}
	}
	return 0
}

func lastInt(line string) int {
	zero := '0'
	var chr rune
	for i := len(line) - 1; i >= 0; i-- {
		chr = rune(line[i])
		if unicode.IsDigit(chr) {
			return int(chr - zero)
		}
	}
	return 0
}

func sumValues(values *[]int) (sum int) {
	for _, i := range *values {
		sum += i
	}
	return
}
