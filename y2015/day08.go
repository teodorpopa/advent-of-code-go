package y2015

import (
	"fmt"
	"github.com/teodorpopa/advent-of-code-go/utils"
	"strconv"
)

func unescape(str string) string {
	s, _ := strconv.Unquote(str)
	return s
}

func escape(str string) string {
	return strconv.Quote(str)
}

func Day08First(input string) int {
	lines := utils.ReadLines(input)
	total := 0

	for _, str := range lines {
		total += len(str) - len(unescape(str))
	}

	return total
}

func Day08Second(input string) int {
	lines := utils.ReadLines(input)
	total := 0

	for _, str := range lines {
		total += len(escape(str)) - len(str)
	}

	return total
}

func Day08() {
	fmt.Println(utils.DAY_PREFIX, "Day 08")
	fmt.Println(utils.PART_PREFIX, "Part 1:", Day08First(utils.ReadFile("y2015/input/day08.txt")))
	fmt.Println(utils.PART_PREFIX, "Part 2:", Day08Second(utils.ReadFile("y2015/input/day08.txt")))
}
