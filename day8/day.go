package day8

import (
	"bufio"
	"log"
	"os"
	"regexp"
)

const FILENAME = "./static/day8.txt"

func Main(part2 bool) {
	inst, whereTo := readFile(FILENAME)
	steps := 0
	if !part2 {
		vals, done := Instructor(inst)
		where := Location("AAA")
		destination := Location("ZZZ")
		for {
			steps += 1
			switch <-vals {
			case 'L':
				where = whereTo[where].left
			case 'R':
				where = whereTo[where].right
			}
			if where == destination {
				done <- true
				close(done)
				break
			}
		}
	} else {
		starts := startingPoints(&whereTo)
		stepChans := make([]chan bool, 0, len(starts))
		arrChans := make([]chan bool, 0, len(starts))

		for _, startFrom := range starts {
			takeStep := make(chan bool)
			hasArrived := make(chan bool)

			go func(where Location, step <-chan bool, arrived chan<- bool) {
				log.Printf("Starting gorouting starting at %s", where)

				vals, done := Instructor(inst)
				for <- step {
					switch <-vals {
					case 'L':
						where = whereTo[where].left
					case 'R':
						where = whereTo[where].right
					}
					arrived <- where[2] == 'Z'
				}
				done <- true
				close(done)
				close(arrived)
			}(startFrom, takeStep, hasArrived)

			stepChans = append(stepChans, takeStep)
			arrChans = append(arrChans, hasArrived)
		}

		for {
			steps += 1
			allArrived := true

			for i, takeStep := range stepChans {
				takeStep <- true
				hm := <- arrChans[i]
				// if hm {
				// 	log.Println("Goroutine", i, "has arrived after", steps)
				// }
				allArrived = allArrived && hm
			}
			if allArrived {
				for _, takeStep := range stepChans {
					takeStep <- false
					close(takeStep)
				}
				break
			}
		}
	}
	log.Println("Arrived after", steps)
}

func startingPoints(m *map[Location]SignPost) []Location {
	starts := make([]Location, 0, 8)
	for l := range *m {
		if l[2] == 'A' {
			starts = append(starts, l)
		}
	}
	return starts
}

func Instructor(s string) (chan rune, chan bool) {
	vals := make(chan rune, 1)
	done := make(chan bool)

	go func() {
		for {
			for i := 0; i < len(s); i++ {
				select {
				case vals <- rune(s[i]):
					continue
				case <-done:
					close(vals)
					return
				}
			}
		}
	}()
	return vals, done
}

type Location string
type SignPost struct {
	left, right Location
}

var LineRegex = regexp.MustCompile(`(.{3})\s*=\s*\((.{3}),\s*(.{3})\)`)

func readLine(s string) (Location, SignPost) {
	if !LineRegex.MatchString(s) {
		return Location(""), SignPost{Location(""), Location("")}
	}
	m := LineRegex.FindStringSubmatch(s)
	return Location(m[1]), SignPost{Location(m[2]), Location(m[3])}
}

func readFile(filename string) (string, map[Location]SignPost) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	var inst string
	whereTo := make(map[Location]SignPost)

	for i, scan := 0, bufio.NewScanner(f); scan.Scan(); i++ {
		line := scan.Text()
		if err := scan.Err(); err != nil {
			log.Fatalln(err)
		}
		if i == 0 {
			inst = line
		} else if len(line) > 0 {
			loc, sp := readLine(line)
			whereTo[loc] = sp
		}
	}
	return inst, whereTo
}
