package day7

import (
	"bufio"
	"cmp"
	"log/slog"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Main(f *os.File, jokers bool) {
	hands := readFile(f)

	if jokers {
		slices.SortFunc(*hands, compareHandsWithJokers)
	} else {
		slices.SortFunc(*hands, compareHands)
	}

	scores := 0
	for rank, hand := range *hands {
		scores += (rank + 1) * hand.wager
	}
	slog.Info("Total", "winnings", scores)
}

type Hand struct {
	cards string
	wager int
}

func compareHandsWithJokers(a, b Hand) int {
	byType := cmp.Compare(identifyHandWithJokers(&a.cards), identifyHandWithJokers(&b.cards))
	if byType != 0 {
		return byType
	}
	return cmp.Compare(cardsToIntWithJokers(&a.cards), cardsToIntWithJokers(&b.cards))
}

func compareHands(a, b Hand) int {
	byType := cmp.Compare(identifyHand(&a.cards), identifyHand(&b.cards))
	if byType != 0 {
		return byType
	}
	return cmp.Compare(cardsToInt(&a.cards), cardsToInt(&b.cards))
}

const cardOrder = "23456789TJQKA"
const jokerOrder = "J23456789TQKA"

func cardsToInt(cards *string) int {
	val := 0
	for _, card := range *cards {
		val = (val << 4) + strings.IndexRune(cardOrder, card)
	}
	return val
}

func cardsToIntWithJokers(cards *string) int {
	val := 0
	for _, card := range *cards {
		val = (val << 4) + strings.IndexRune(jokerOrder, card)
	}
	return val
}

type HandType int8

const (
	HIGH_CARD HandType = iota
	ONE_PAIR
	TWO_PAIR
	THREE_OF_A_KIND
	FULL_HOUSE
	FOUR_OF_A_KIND
	FIVE_OF_A_KIND
)

func identifyHandWithJokers(cards *string) HandType {
	count := make(map[rune]int8)
	for _, card := range *cards {
		count[card]++
	}
	if jokers, ok := count['J']; ok {
		delete(count, 'J')
		actAs := keyWithMost(count)
		count[actAs] += jokers

	}
	return identifyHandFromCount(count)
}

func keyWithMost(count map[rune]int8) rune {
	mostK, mostV := 'J', int8(0)
	for k, v := range count {
		if v > mostV {
			mostK, mostV = k, v
		}
	}
	return mostK
}

func identifyHand(cards *string) HandType {
	count := make(map[rune]int8)
	for _, card := range *cards {
		count[card]++
	}
	return identifyHandFromCount(count)
}

func identifyHandFromCount(count map[rune]int8) HandType {
	switch len(count) {
	case 1:
		return FIVE_OF_A_KIND
	case 2:
		for _, v := range count {
			if v == 4 || v == 1 {
				return FOUR_OF_A_KIND
			} else {
				return FULL_HOUSE
			}
		}
	case 3:
		for _, v := range count {
			if v == 3 {
				return THREE_OF_A_KIND
			}
		}
		return TWO_PAIR
	case 4:
		return ONE_PAIR
	case 5:
		return HIGH_CARD
	}
	return 0
}

func readFile(f *os.File) *[]Hand {
	scanner := bufio.NewScanner(f)

	hands := make([]Hand, 0, 64)
	for scanner.Scan() {
		line := scanner.Text()
		if err := scanner.Err(); err != nil {
			slog.Error("Failure reading line", "error", err)
		}

		parts := strings.Split(line, " ")
		for _, p := range parts[1:] {
			if len(p) > 0 {
				i, err := strconv.Atoi(p)
				if err != nil {
					slog.Error("Failure decoding int", "error", err)
				}
				hands = append(hands, Hand{parts[0], i})
			}
		}
	}
	return &hands
}
