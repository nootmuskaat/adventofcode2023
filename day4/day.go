package day4

import (
	"bufio"
	mapset "github.com/deckarep/golang-set/v2"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

const FILENAME = "./static/day4.txt"

func Main(part2 bool) {
	lines := readFile(FILENAME)
	var sum uint
	for lineNo, line := range *lines {
		if lineNo < 10 {
			log.Printf("%d:%s", lineNo, line)
		}
		if len(line) == 0 {
			continue
		}
		scratchCard := NewScratchCard(&line)
		sum += scratchCard.Points()
	}
	log.Println("Sum:", sum)
}

type ScratchCard struct {
	winningValues *mapset.Set[uint8]
	scratchValues *mapset.Set[uint8]
}

func (sc ScratchCard) Points() uint {
	winners := *sc.winningValues
	scratched := *sc.scratchValues
	matches := winners.Intersect(scratched)
	log.Printf("%v & %v -> %v\n", winners, scratched, matches)
	if matches.IsEmpty() {
		return 0
	}

	return uint(math.Pow(2.0, float64(matches.Cardinality()-1)))
}

func NewScratchCard(line *string) *ScratchCard {
	splitOffCardNo := strings.Split(*line, ":")
	if len(splitOffCardNo) != 2 {
		log.Fatalf("Unable to parse '%s': %s\n", *line, splitOffCardNo)
	}
	splitWinningFromOthers := strings.Split(splitOffCardNo[1], "|")
	if len(splitWinningFromOthers) != 2 {
		log.Fatalf("Unable to parse '%s': %s\n", splitOffCardNo[1], splitWinningFromOthers)
	}
	winners := asUintSet(splitWinningFromOthers[0])
	others := asUintSet(splitWinningFromOthers[1])
	return &ScratchCard{winners, others}
}

func asUintSet(s string) *mapset.Set[uint8] {
	values := mapset.NewSet[uint8]()
	for _, i := range strings.Split(s, " ") {
		if len(i) == 0 {
			continue
		}
		u, err := strconv.ParseUint(i, 10, 8)
		if err != nil {
			log.Fatalln("Failed to parse segment", s)
		}
		values.Add(uint8(u))
	}
	return &values
}

func readFile(filename string) *[]string {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	lines := make([]string, 0, 16)
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
