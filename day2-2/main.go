package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type round struct {
	red, green, blue int
}

type game struct {
	id                        int
	rounds                    []round
	minRed, minGreen, minBlue int
}

// https://adventofcode.com/2023/day/2#part2
func main() {
	var games []game

	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scan := bufio.NewScanner(f)

	// loop through every game (line)
	for scan.Scan() {
		gameStr := scan.Text()

		game := game{}

		// split into "game # : rounds"
		gameSplit := strings.Split(gameStr, ":")

		// scan game number
		n, err := fmt.Sscanf(gameSplit[0], "Game %d", &game.id)
		if err != nil {
			log.Fatal(err)
		}
		if n != 1 {
			log.Fatal("did not scan game # properly")
		}

		// split into rounds
		roundSplit := strings.Split(gameSplit[1], ";")
		for _, roundStr := range roundSplit {
			round := round{}

			// split into colors
			colors := strings.Split(roundStr, ",")
			for _, color := range colors {
				switch {
				case strings.Contains(color, "red"):
					n, err := fmt.Sscanf(color, "%d red", &round.red)
					if err != nil {
						log.Fatal(err)
					}
					if n != 1 {
						log.Fatal("did not scan red properly")
					}

					// update game minimum
					if game.minRed < round.red {
						game.minRed = round.red
					}

				case strings.Contains(color, "green"):
					n, err := fmt.Sscanf(color, "%d green", &round.green)
					if err != nil {
						log.Fatal(err)
					}
					if n != 1 {
						log.Fatal("did not scan green properly")
					}

					// update game minimum
					if game.minGreen < round.green {
						game.minGreen = round.green
					}

				case strings.Contains(color, "blue"):
					n, err := fmt.Sscanf(color, "%d blue", &round.blue)
					if err != nil {
						log.Fatal(err)
					}
					if n != 1 {
						log.Fatal("did not scan blue properly")
					}

					// update game minimum
					if game.minBlue < round.blue {
						game.minBlue = round.blue
					}
				}
			}

			game.rounds = append(game.rounds, round)
		}

		games = append(games, game)
	}

	total := 0

	for _, game := range games {
		total += game.minRed * game.minGreen * game.minBlue
	}

	fmt.Println(total)
}
