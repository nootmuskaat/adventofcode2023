package day7

import (
	"fmt"
	"slices"
	"testing"
)

func TestOrderHands(t *testing.T) {
	hands := []Hand{
		Hand{"T55J5", 684},
		Hand{"KK677", 28},
		Hand{"KTJJT", 220},
		Hand{"QQQJA", 483},
		Hand{"32T3K", 765},
	}

	slices.SortFunc(hands, compareHands)

	expected := []Hand{
		Hand{"32T3K", 765},
		Hand{"KTJJT", 220},
		Hand{"KK677", 28},
		Hand{"T55J5", 684},
		Hand{"QQQJA", 483},
	}

	for i := 0; i < len(expected); i++ {
		if hands[i].cards != expected[i].cards {
			t.Fatal("Got", hands, "expected", expected)
		}
	}
}

func TestIdentifyHand(t *testing.T) {
	testCases := []struct {
		cards    string
		expected HandType
	}{
		{"AAAAA", FIVE_OF_A_KIND},
		{"AA8AA", FOUR_OF_A_KIND},
		{"23332", FULL_HOUSE},
		{"TTT98", THREE_OF_A_KIND},
		{"23432", TWO_PAIR},
		{"A23A4", ONE_PAIR},
		{"23456", HIGH_CARD},
	}

	for _, tc := range testCases {
		testName := fmt.Sprintf("identifyHand(%s)->%d", tc.cards, tc.expected)
		t.Run(testName, func(t *testing.T) {
			got := identifyHand(&tc.cards)
			if got != tc.expected {
				t.Errorf("Got %d expected %d", got, tc.expected)
			}
		})
	}
}

func TestIdentifyHandWithJokers(t *testing.T) {
	testCases := []struct {
		cards    string
		expected HandType
	}{
		{"32T3K", ONE_PAIR},
		{"T55J5", FOUR_OF_A_KIND},
		{"KK677", TWO_PAIR},
		{"KTJJT", FOUR_OF_A_KIND},
		{"QQQJA", FOUR_OF_A_KIND},
		{"3245J", ONE_PAIR},
		{"324JJ", THREE_OF_A_KIND},
		{"3JJJJ", FIVE_OF_A_KIND},
	}

	for _, tc := range testCases {
		testName := fmt.Sprintf("identifyHand(%s)->%d", tc.cards, tc.expected)
		t.Run(testName, func(t *testing.T) {
			got := identifyHandWithJokers(&tc.cards)
			if got != tc.expected {
				t.Errorf("Got %d expected %d", got, tc.expected)
			}
		})
	}
}

func TestCardsToInt(t *testing.T) {
	testCases := []struct {
		cards    string
		expected int
	}{
		{"AAAAA", 0b11001100110011001100},
		{"22222", 0b0},
		{"K2222", 0b10110000000000000000},
	}

	for _, tc := range testCases {
		testName := fmt.Sprintf("cardsToInt(%s)->%d", tc.cards, tc.expected)
		t.Run(testName, func(t *testing.T) {
			got := cardsToInt(&tc.cards)
			if got != tc.expected {
				t.Errorf("Got %d expected %d", got, tc.expected)
			}
		})
	}
}
