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

func Main(f *os.File, part2 bool) {
	lines := readFile(f)
	var sum uint
	count := make(map[int]uint)
	for cardNo, line := range *lines {
		scratchCard := NewScratchCard(&line)
		if !part2 {
			sum += scratchCard.Points()
		} else {
			count[cardNo]++
			for i, m := 1, scratchCard.Matches(); i <= m; i++ {
				count[cardNo+i] += count[cardNo]
			}
			sum += count[cardNo]
		}
	}
	log.Println("Sum:", sum)
}

type ScratchCard struct {
	winningValues *mapset.Set[uint8]
	scratchValues *mapset.Set[uint8]
}

func (sc ScratchCard) Matches() int {
	winners := *sc.winningValues
	scratched := *sc.scratchValues
	matches := winners.Intersect(scratched)
	return matches.Cardinality()
}

func (sc ScratchCard) Points() uint {
	if matches := sc.Matches(); matches == 0 {
		return 0
	} else {
		return uint(math.Pow(2.0, float64(matches-1)))
	}
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

func readFile(f *os.File) *[]string {
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
