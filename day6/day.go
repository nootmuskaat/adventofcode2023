package day6

import (
	"log"
	"os"
	"strconv"
	"strings"
)

const FILENAME = "./static/day6.txt"

func Main(part2 bool) {
	races := readFile(FILENAME, part2)
	allWinningWays := uint(1)
	for _, race := range races {
		winningWays := uint(0)
		for i := uint(1); i < race.time; i++ {
			if i*(race.time-i) > race.distance {
				winningWays++
			}
		}
		allWinningWays *= winningWays
	}
	log.Println("Result", allWinningWays)
}

type Race struct {
	time, distance uint
}

func readFile(filename string, joinInts bool) []Race {
	races := make([]Race, 0, 4)
	times := make([]uint, 0, 4)
	distances := make([]uint, 0, 4)

	contents, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalln(err)
	}
	lines := strings.Split(string(contents), "\n")
	for _, i := range strings.Split(lines[0], " ")[1:] {
		if len(i) == 0 {
			continue
		}
		ii, e := strconv.ParseUint(i, 10, 64)
		if e != nil {
			log.Fatalln(e)
		}
		times = append(times, uint(ii))
	}
	for _, i := range strings.Split(lines[1], " ")[1:] {
		if len(i) == 0 {
			continue
		}
		ii, e := strconv.ParseUint(i, 10, 64)
		if e != nil {
			log.Fatalln(e)
		}
		distances = append(distances, uint(ii))
	}
	if joinInts {
		timeStrs := make([]string, 0, len(times))
		distStrs := make([]string, 0, len(distances))
		for i := 0; i < len(times); i++ {
			timeStrs = append(timeStrs, strconv.Itoa(int(times[i])))
			distStrs = append(distStrs, strconv.Itoa(int(distances[i])))
		}
		singleTime, et := strconv.ParseUint(strings.Join(timeStrs, ""), 10, 64)
		singleDist, ed := strconv.ParseUint(strings.Join(distStrs, ""), 10, 64)
		if et != nil || ed != nil {
			log.Fatalln(et, ed)
		}
		times = []uint{uint(singleTime)}
		distances = []uint{uint(singleDist)}
	}
	for i := 0; i < len(times); i++ {
		races = append(races, Race{times[i], distances[i]})
	}
	return races
}
