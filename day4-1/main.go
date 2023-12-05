package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

// https://adventofcode.com/2023/day/4
func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Println(err)
	}

	scan := bufio.NewScanner(f)

	total := 0

	for scan.Scan() {
		line := strings.Split(scan.Text(), ":")

		cards := strings.Split(line[1], "|")

		myNums := strings.Split(cards[0], " ")
		winNums := strings.Split(cards[1], " ")

		numOfMatches := 0

		// match cards
		for _, my := range myNums {
			for _, win := range winNums {
				if my == "" || win == "" {
					continue
				}
				if my == win {
					numOfMatches++
				}
			}
		}

		switch numOfMatches {
		case 0:

		case 1:
			total++
		default:
			total += int(math.Pow(2, float64(numOfMatches)-1))
		}
	}

	fmt.Println(total)
}
