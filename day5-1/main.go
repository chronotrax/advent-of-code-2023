package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Converter struct {
	destStart, sourceStart, length int
}

type Map struct {
	converters []Converter
}

// https://adventofcode.com/2023/day/5
func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scan := bufio.NewScanner(f)

	// scan seeds
	scan.Scan()
	line := scan.Text()
	var seeds []int
	seedsStr := strings.Fields(strings.SplitAfter(line, ":")[1])
	for _, s := range seedsStr {
		i, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}

		seeds = append(seeds, i)
	}
	scan.Scan()

	maps := make([]Map, 7)

	// scan every map
	for i := 0; i < 7; i++ {
		// skip text line
		scan.Scan()

		maps[i] = *readMap(scan)
	}

	// apply maps
	locations := make([]int, len(seeds))
	for i := range seeds {
		for _, m := range maps {
			var converter *Converter

			// find converter in range
			for _, c := range m.converters {
				if seeds[i] >= c.sourceStart && seeds[i] < c.sourceStart+c.length {
					converter = &c
					break
				}
			}

			// skip converting if no converter
			if converter == nil {
				continue
			}

			// convert
			seeds[i] = converter.destStart + (seeds[i] - converter.sourceStart)
		}

		locations[i] = seeds[i]
	}

	lowest := locations[0]
	for _, l := range locations {
		if l < lowest {
			lowest = l
		}
	}

	fmt.Println(lowest)
}

func readMap(scan *bufio.Scanner) *Map {
	m := Map{converters: []Converter{}}

	for scan.Scan() {
		// check for blank line
		line := scan.Text()
		if line == "" {
			break
		}

		c := Converter{}

		var nums []int

		numsStr := strings.Fields(line)

		// convert number strings into ints
		for _, i := range numsStr {
			n, err := strconv.Atoi(i)
			if err != nil {
				log.Fatal(err)
			}

			nums = append(nums, n)
		}

		c.destStart = nums[0]
		c.sourceStart = nums[1]
		c.length = nums[2]

		m.converters = append(m.converters, c)
	}

	// sort map's converters based on source start
	slices.SortFunc[[]Converter, Converter](m.converters, func(a, b Converter) int {
		if a.sourceStart < b.sourceStart {
			return -1
		}

		if a.sourceStart > b.sourceStart {
			return 1
		}

		return 0
	})

	return &m
}
