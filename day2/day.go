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

func Main() {
	allGames := readGames()
	for gameNo, game := range *allGames {
		fmt.Printf("Game %d: ", gameNo)
		for drawNo, draw := range game {
			fmt.Printf("draw %d: %v |", drawNo, draw)
		}
		fmt.Println()
	}
}

func readGames() *[][]Draw {
	games := make([][]Draw, 0, 100)

	f, err := os.Open(DAY2_FILE)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	tmp_counter := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		game := readLine(&line)
		games = append(games, *game)

		tmp_counter++
		if tmp_counter >= 10 {
			break
		}

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
