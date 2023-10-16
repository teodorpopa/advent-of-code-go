package y2015

import (
	"fmt"
	"github.com/teodorpopa/advent-of-code-go/utils"
	"strings"
)

func Day01First(input string) int {
	floor := 0

	for _, v := range strings.Split(input, "") {
		switch v {
		case "(":
			floor++
			break
		case ")":
			floor--
			break
		}
	}

	return floor
}

func Day01Second(input string) int {
	floor := 0

	for p, v := range strings.Split(input, "") {
		switch v {
		case "(":
			floor++
			break
		case ")":
			floor--
			break
		}

		if floor == -1 {
			return p + 1
		}
	}

	return 0
}

func Day01() {
	fmt.Println(utils.DAY_PREFIX, "Day 01 - Not Quite Lisp")
	fmt.Println(utils.PART_PREFIX, "Part 1:", Day01First(utils.ReadFile("y2015/input/day01.txt")))
	fmt.Println(utils.PART_PREFIX, "Part 2:", Day01Second(utils.ReadFile("y2015/input/day01.txt")))
}
