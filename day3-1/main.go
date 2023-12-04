package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

type Number struct {
	byteS                   []byte
	startX, endX, y, length int
}

var total = 0

// https://adventofcode.com/2023/day/3
func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scan := bufio.NewScanner(f)

	var grid [][]byte

	// loop through every line
	for scan.Scan() {
		line := scan.Text()

		grid = append(grid, []byte(line))
	}

	// loop through grid
	num := Number{}
	for i, row := range grid {
		for j, v := range row {
			if unicode.IsDigit(rune(v)) {
				// create number
				if num.length == 0 {
					num.startX = j
					num.y = i
				}

				num.byteS = append(num.byteS, v)
				num.endX = j
				num.length = num.endX - num.startX + 1
			} else {
				// found end of number
				checkNumber(&num, &grid)
			}
		}

		// check end of row for number
		checkNumber(&num, &grid)
	}

	fmt.Println(total)
}

// checks if number is surrounded by symbol and adds to total if it does
// clears num when done
func checkNumber(num *Number, grid *[][]byte) {
	if num.length != 0 {
		if checkSurroundings(num, grid) {
			res, err := strconv.Atoi(string(num.byteS))
			if err != nil {
				log.Fatal(err)
			}

			total += res
		}
	}

	*num = Number{}
}

// checks if there are any symbols around the number
func checkSurroundings(num *Number, grid *[][]byte) bool {
	maxX := len((*grid)[0])
	maxY := len(*grid)

	for i := num.y - 1; i <= num.y+1; i++ {
		for j := num.startX - 1; j <= num.endX+1; j++ {
			// avoid checking itself
			if i == num.y && j >= num.startX && j <= num.endX {
				continue
			}

			if checkBounds(j, i, maxX, maxY) {
				val := (*grid)[i][j]
				if val != '.' {
					return true
				}
			}
		}
	}

	return false
}

// checks if x & y are in bounds of the grid
func checkBounds(x, y, maxX, maxY int) bool {
	return x >= 0 && x < maxX && y >= 0 && y < maxY
}
