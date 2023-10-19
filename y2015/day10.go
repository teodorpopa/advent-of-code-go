package y2015

import (
	"fmt"
	"github.com/teodorpopa/advent-of-code-go/utils"
	"strconv"
	"strings"
)

func Day10First(input string, iterations int) int {
	start := input

	for i := 0; i < iterations; i++ {
		digits := strings.Split(start, "")
		newString := ""
		digit := ""
		repeat := 0

		for _, v := range digits {
			if digit == "" {
				digit = v
				repeat = 1
			} else if v == digit {
				repeat++
			} else {
				newString = newString + strconv.Itoa(repeat) + digit

				digit = v
				repeat = 1
			}
		}
		newString = newString + strconv.Itoa(repeat) + digit
		start = newString
	}

	return len(start)
}

func Day10() {
	fmt.Println(utils.DAY_PREFIX, "Day 10")
	fmt.Println(utils.PART_PREFIX, "Part 1:", Day10First(utils.ReadFile("y2015/input/day10.txt"), 40))
	fmt.Println(utils.PART_PREFIX, "Part 2:", Day10First(utils.ReadFile("y2015/input/day10.txt"), 50))
}
