package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

type Converter struct {
	destStart, sourceStart, length int
}

type Map struct {
	converters []Converter
}

type SeedRange struct {
	start, length int
}

// https://adventofcode.com/2023/day/5#part2
func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scan := bufio.NewScanner(f)

	// scan seeds
	scan.Scan()
	line := scan.Text()
	var seeds []SeedRange
	seedsStr := strings.Fields(strings.SplitAfter(line, ":")[1])

	for i := 0; i < len(seedsStr); i += 2 {
		start, err := strconv.Atoi(seedsStr[i])
		if err != nil {
			log.Fatal(err)
		}

		length, err := strconv.Atoi(seedsStr[i+1])
		if err != nil {
			log.Fatal(err)
		}

		seeds = append(seeds, SeedRange{start: start, length: length})
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
	lowest := -1
	wg := sync.WaitGroup{}
	mu := &sync.Mutex{}

	// use goroutines to check each seed range async
	for _, seedRange := range seeds {
		s := seedRange
		wg.Add(1)

		go func() {
			defer wg.Done()
			res := applyMaps(s, maps)

			mu.Lock()
			if lowest == -1 || res < lowest {
				lowest = res
			}
			mu.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println(lowest)
}

// calculates the lowest location in the range of seeds
func applyMaps(sr SeedRange, maps []Map) int {
	lowest := -1

	// for every seed in the range
	for s := sr.start; s < sr.start+sr.length; s++ {
		seed := s

		// apply each map
		for _, m := range maps {
			var converter *Converter

			// check each converter if seed needs to be converted
			for _, c := range m.converters {
				if seed >= c.sourceStart && seed < c.sourceStart+c.length {
					converter = &c
					break
				}
			}

			// skip converting if no converter
			if converter == nil {
				continue
			}

			// convert
			seed = converter.destStart + (seed - converter.sourceStart)
		}

		if lowest == -1 || seed < lowest {
			lowest = seed
		}
	}

	return lowest
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

	return &m
}
