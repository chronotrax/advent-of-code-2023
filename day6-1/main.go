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

// https://adventofcode.com/2023/day/6
func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scan := bufio.NewScanner(f)

	scan.Scan()
	line := scan.Text()
	timesStr := strings.Fields(strings.Split(line, ":")[1])
	scan.Scan()
	line = scan.Text()
	distancesStr := strings.Fields(strings.Split(line, ":")[1])

	var races []Race

	for i := range timesStr {
		time, err := strconv.Atoi(timesStr[i])
		if err != nil {
			log.Fatal(err)
		}

		distance, err := strconv.Atoi(distancesStr[i])
		if err != nil {
			log.Fatal(err)
		}

		races = append(races, Race{time: time, distance: distance})
	}

	totals := make([]int, len(races))

	for i := range races {
		for prepTime := 0; prepTime < races[i].time; prepTime++ {
			if willFinishRace(prepTime, races[i]) {
				totals[i]++
			}
		}
	}

	total := totals[0]
	for i := 1; i < len(totals); i++ {
		total *= totals[i]
	}

	fmt.Println(total)
}

func willFinishRace(prepTime int, race Race) bool {
	timeRemaining := race.time - prepTime

	// note prepTime == speed of boat
	return prepTime*timeRemaining > race.distance
}
