package day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const DAY2_FILE = "./static/day2.txt"

type Draw struct {
	red, green, blue uint64
}

func (d Draw) contains(other *Draw) bool {
	return other.red <= d.red && other.green <= d.green && other.blue <= d.blue
}

func (d Draw) Power() uint64 {
	return d.red * d.green * d.blue
}

func maxDraw(draws ...Draw) *Draw {
	combined := Draw{}
	for _, draw := range draws {
		if draw.red > combined.red {
			combined.red = draw.red
		}
		if draw.green > combined.green {
			combined.green = draw.green
		}
		if draw.blue > combined.blue {
			combined.blue = draw.blue
		}
	}
	return &combined
}

func Main() {
	allGames := readGames()

	// Part 1
	// Determine which games would have been possible if the bag had been loaded with only
	// 12 red cubes, 13 green cubes, and 14 blue cubes. What is the sum of the IDs of those games?
	// limit := Draw{12, 13, 14}
	var sum uint64

	// Part 2
	// The power of a set of cubes is equal to the numbers of red, green, and blue cubes multiplied together.
	// For each game, find the minimum set of cubes that must have been present. What is the sum of the power of these sets?

	for _, draws := range *allGames {
		maxPossibleDraw := maxDraw(draws...)
		sum += maxPossibleDraw.Power()

		// if limit.contains(maxPossibleDraw) {
		// 	sum += gameNo + 1
		// }
	}
	fmt.Println("Sum", sum)
}

func readGames() *[][]Draw {
	games := make([][]Draw, 0, 100)

	f, err := os.Open(DAY2_FILE)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		game := readLine(&line)
		games = append(games, *game)
	}
	return &games
}

func readLine(line *string) *[]Draw {
	draws := make([]Draw, 0, 4)

	parts := strings.Split(*line, ": ")
	if len(parts) != 2 {
		log.Fatal("Too many parts:", parts)
	}
	for _, drawStr := range strings.Split(parts[1], "; ") {
		var red, green, blue uint64
		for _, partStr := range strings.Split(drawStr, ", ") {
			part := strings.Split(partStr, " ")
			if len(part) != 2 {
				log.Fatal("Too many parts in amount declaration:", part)
			}
			amount, err := strconv.ParseUint(part[0], 10, 0)
			if err != nil {
				log.Fatalf("Failed to decode amount: %s\n", part)
			}
			switch part[1] {
			case "red":
				red = amount
			case "green":
				green = amount
			case "blue":
				blue = amount
			}
		}
		draws = append(draws, Draw{red, green, blue})
	}
	return &draws
}
