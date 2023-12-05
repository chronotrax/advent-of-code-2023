package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// https://adventofcode.com/2023/day/4#part2
func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Println(err)
	}

	scan := bufio.NewScanner(f)

	total := 0

	var cardQueue []int

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

		// find how many copies
		copies := 0
		if len(cardQueue) > 0 {
			copies = cardQueue[0]
		}

		total += copies + 1

		// remove self from queue
		if len(cardQueue) > 0 {
			cardQueue = cardQueue[1:]
		}

		for i := 0; i < numOfMatches; i++ {
			// add to queue
			if i >= len(cardQueue) {
				cardQueue = append(cardQueue, copies+1)
			} else { // increment copies of cards
				cardQueue[i] += copies + 1
			}
		}
	}

	fmt.Println(total)
}
