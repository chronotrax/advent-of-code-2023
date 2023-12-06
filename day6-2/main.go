package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Race struct {
	time, distance int
}

// https://adventofcode.com/2023/day/6#part2
func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scan := bufio.NewScanner(f)

	scan.Scan()
	line := scan.Text()
	line = strings.ReplaceAll(line, " ", "")
	timesStr := strings.Split(line, ":")[1]
	scan.Scan()
	line = scan.Text()
	line = strings.ReplaceAll(line, " ", "")
	distancesStr := strings.Split(line, ":")[1]

	time, err := strconv.Atoi(timesStr)
	if err != nil {
		log.Fatal(err)
	}
	distance, err := strconv.Atoi(distancesStr)
	if err != nil {
		log.Fatal(err)
	}

	race := Race{time: time, distance: distance}

	total := 0

	for prepTime := 0; prepTime < race.time; prepTime++ {
		if willFinishRace(prepTime, race) {
			total++
		}
	}

	fmt.Println(total)
}

func willFinishRace(prepTime int, race Race) bool {
	timeRemaining := race.time - prepTime

	// note prepTime == speed of boat
	return prepTime*timeRemaining > race.distance
}
