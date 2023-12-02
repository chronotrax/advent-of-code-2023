package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

// https://adventofcode.com/2023/day/1
func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scan := bufio.NewScanner(f)

	total := 0

	// loop through every line
	for scan.Scan() {
		first, last := byte(0), byte(0)
		line := scan.Text()

		// find first number
		for i := 0; i < len(line); i++ {
			if unicode.IsDigit(rune(line[i])) {
				first = line[i]
				break
			}
		}

		// find last number
		for i := len(line) - 1; i >= 0; i-- {
			if unicode.IsDigit(rune(line[i])) {
				last = line[i]
				break
			}
		}

		// combine the 2 numbers together and convert to an int
		num, err := strconv.Atoi(fmt.Sprintf("%c%c", first, last))
		if err != nil {
			log.Fatal(err)
		}

		total += num
	}

	fmt.Println(total)
}
