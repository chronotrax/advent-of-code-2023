package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// https://adventofcode.com/2023/day/1#part2
func main() {
	numberStrings := [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	// converts int 0 to ascii 1 (49)
	// so numberStrings[0] will translate to 1
	converter := 49

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

		// find first number string
		firstIndex := len(line)
		for i, numStr := range numberStrings {
			if res := strings.Index(line, numStr); res != -1 && res < firstIndex {
				firstIndex = res
				first = byte(i + converter)
			}
		}

		// search for earlier digits
		for i := 0; i < firstIndex; i++ {
			// check for digit
			if unicode.IsDigit(rune(line[i])) {
				first = line[i]
				break
			}
		}

		// find last number
		lastIndex := 0
		for i, numStr := range numberStrings {
			if res := strings.LastIndex(line, numStr); res != -1 && res > lastIndex {
				lastIndex = res
				last = byte(i + converter)
			}
		}

		// search for later digits
		for i := len(line) - 1; i >= lastIndex; i-- {
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
