package day8

import (
	"bufio"
	"log"
	"os"
	"regexp"
)

func Main(f *os.File, part2 bool) {
	inst, whereTo := readFile(f)
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
		isDone := make(chan bool)
		loops := make([]Loop, len(starts))

		for i, startFrom := range starts {
			loops[i] = Loop{0, []int{}}
			go findLoop(inst, startFrom, &whereTo, &loops[i], isDone)
		}
		for range starts {
			<-isDone
		}
		lens := make([]int, len(starts))
		// from earlier testing, I know that each loop has one endpoint which is
		// equal to the length of the loop
		for i, l := range loops {
			lens[i] = l.length
		}
		steps = lcm(lens[0], lens[1], lens[2:]...)
	}
	log.Println("Arrived after", steps)
}

func findLoop(inst string, startPoint Location, whereTo *map[Location]SignPost, loop *Loop, done chan bool) {
	vals, stopIter := Instructor(inst)
	secondaryCheck := make(map[LoopPoint]int)
	arrivals := make([]int, 0, 8)
	current := startPoint

	steps := 0
	for ; ; steps++ {
		if (steps%len(inst)) == 0 && current == startPoint && steps > 0 {
			break
		}
		switch <-vals {
		case 'L':
			current = (*whereTo)[current].left
		case 'R':
			current = (*whereTo)[current].right
		}
		if current[2] == 'Z' {
			lp := LoopPoint{current, steps % len(inst)}
			if secondaryCheck[lp] > 0 {
				steps -= secondaryCheck[lp]
				break
			}
			secondaryCheck[lp] = steps
			arrivals = append(arrivals, steps+1)
		}
	}
	stopIter <- true
	close(stopIter)
	*loop = Loop{steps, arrivals}
	done <- true
}

func gcd(a, b int, o ...int) int {
	if len(o) > 0 {
		return gcd(a, gcd(b, o[0], o[1:]...))
	}
	for a != b {
		if a > b {
			a -= b
		} else {
			b -= a
		}
	}
	return a
}

func lcm(a, b int, o ...int) int {
	if len(o) > 0 {
		return lcm(a, lcm(b, o[0], o[1:]...))
	}
	if a > b {
		return (a / gcd(a, b)) * b
	} else {
		return (b / gcd(a, b)) * a
	}
}

type LoopPoint struct {
	location Location
	when     int
}

type Loop struct {
	length    int
	endpoints []int
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

func readFile(f *os.File) (string, map[Location]SignPost) {
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
