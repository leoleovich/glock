package main

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

const base = 10
const margin = 2
const numProbable = 3
const loglevel = log.InfoLevel

type probable struct {
	number int
	// How often we see this number
	frequency int
}

// get number +- margin
func around(number, margin int) []int {
	around := []int{}
	around = append(around, number)

	for i := 1; i <= margin; i++ {
		before := number - i
		if before < 0 {
			before += base
		}
		around = append(around, before)

		after := number + i
		if after > 9 {
			after -= base
		}
		around = append(around, after)
	}
	return around
}

// get top N frequent numbers in all "around" combinations
func topN(combinations map[int]int) []probable {
	options := make([]probable, numProbable)

	for num, freq := range combinations {
		pushed := false
		for i, opt := range options {
			if freq >= opt.frequency {
				// push down
				for c := numProbable - 1; c > i; c-- {
					options[c] = options[c-1]
				}
				options[i] = probable{number: num, frequency: freq}
				pushed = true
			}
			if pushed {
				break
			}
		}
	}

	return options
}

func main() {
	log.SetLevel(loglevel)

	if len(os.Args) < 3 {
		log.Fatal("Must be at least 2 wrong combinations")
	}

	wLen := len(os.Args[1])
	cLen := len(os.Args[1:])

	// map [combination Number][combination Options][freq]
	uniq := make([]map[int]int, wLen)
	for i := 0; i < wLen; i++ {
		uniq[i] = make(map[int]int)
	}

	// From input combinations build the map with possible options and frequency
	for _, wrong := range os.Args[1:] {
		if wLen != len(wrong) {
			log.Fatal("Wrong combinations must be the same lenght")
		}

		for c, num := range wrong {
			combinations := uniq[c]

			for _, a := range around(int(num-'0'), margin) {
				if _, ok := combinations[a]; !ok {
					combinations[a] = 1
				} else {
					combinations[a]++
				}
			}
		}
	}

	for n, combinations := range uniq {
		log.Debug(combinations)

		options := topN(combinations)

		fmt.Printf("Number %d is probably %d (%d/%d). If not try: ", n+1, options[0].number, options[0].frequency, cLen)
		for i := 1; i < numProbable; i++ {
			fmt.Printf("%d (%d/%d)", options[i].number, options[i].frequency, cLen)
			if i != numProbable-1 {
				fmt.Printf(" or ")
			} else {
				fmt.Printf("\n")
			}
		}
	}
}
