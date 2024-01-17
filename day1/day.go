package day1

import (
	"bufio"
	"log"
	"os"
	"strings"
	"unicode"
)

const ZERO = '0'

// golang does not allow this to be const
var NUMBERS_AS_TEXT = map[string]int{
	"zero":  0,
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func Main(f *os.File, includeStrings bool) {
	scanner := bufio.NewScanner(f)

	sum := 0
	var line_value int

	for scanner.Scan() {
		line := scanner.Text()
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		line_value = 10*firstInt(&line, includeStrings) + lastInt(&line, includeStrings)
		// log.Printf("Line %s yielded value %d", line, line_value)
		sum += line_value
	}

	log.Printf("Sum: %d\n", sum)
}

// locate the first spelled-out number in the string
func firstString(line string) int {
	numeral, numeral_idx := -1, len(line)
	for k, v := range NUMBERS_AS_TEXT {
		if idx := strings.Index(line, k); idx != -1 && idx < numeral_idx {
			numeral, numeral_idx = v, idx
		}
	}
	return numeral
}

// locate the last spelled-out number in the string
func lastString(line string) int {
	numeral, numeral_idx := -1, -1
	for k, v := range NUMBERS_AS_TEXT {
		if idx := strings.LastIndex(line, k); idx != -1 && idx > numeral_idx {
			numeral, numeral_idx = v, idx
		}
	}
	return numeral
}

// Locate the first integer in the string
func firstInt(line *string, includeStrings bool) int {
	numeral, numeral_idx := 0, -1
	for idx, chr := range *line {
		if unicode.IsDigit(chr) {
			numeral, numeral_idx = int(chr-ZERO), idx
			break
		}
	}
	var prefix string
	if !includeStrings {
		return numeral
	}
	if numeral_idx == -1 {
		prefix = *line
	} else {
		prefix = (*line)[:numeral_idx]
	}
	if first_string := firstString(prefix); first_string != -1 {
		numeral = first_string
	}
	return numeral
}

// Locate the last integer in the string
func lastInt(line *string, includeStrings bool) int {
	numeral, numeral_idx := 0, 0
	var chr rune
	for idx := len(*line) - 1; idx >= 0; idx-- {
		chr = rune((*line)[idx])
		if unicode.IsDigit(chr) {
			numeral, numeral_idx = int(chr-ZERO), idx
			break
		}
	}
	if !includeStrings {
		return numeral
	}
	suffix := (*line)[numeral_idx:]
	if last_string := lastString(suffix); last_string != -1 {
		numeral = last_string
	}
	return numeral
}
