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
	byteS                        []byte
	startX, endX, y, length, val int
}

type Gear struct {
	nums []int
	x, y int
}

var total = 0
var gears []Gear

// https://adventofcode.com/2023/day/3#part2
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

	// check gears
	for _, g := range gears {
		if len(g.nums) == 2 {
			total += g.nums[0] * g.nums[1]
		}
	}

	fmt.Println(total)
}

// checks if number is surrounded by a gear and adds itself to the gear
// clears num when done
func checkNumber(num *Number, grid *[][]byte) {
	if num.length != 0 {

		res, err := strconv.Atoi(string(num.byteS))
		if err != nil {
			log.Fatal(err)
		}

		num.val = res

		checkSurroundings(num, grid)

		*num = Number{}
	}
}

// checks if there are any gears surrounding the number
func checkSurroundings(num *Number, grid *[][]byte) {
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
				if val == '*' {
					found := false
					for g := 0; g < len(gears); g++ {
						if gears[g].x == j && gears[g].y == i {
							gears[g].nums = append(gears[g].nums, num.val)
							return
						}
					}

					if !found {
						gears = append(gears, Gear{x: j, y: i, nums: []int{num.val}})
					}
				}
			}
		}
	}
}

// checks if x & y are in bounds of the grid
func checkBounds(x, y, maxX, maxY int) bool {
	return x >= 0 && x < maxX && y >= 0 && y < maxY
}
