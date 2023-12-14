package day08

import (
	_ "embed"
	"fmt"
	"github.com/teodorpopa/advent-of-code-go/utils"
	"strconv"
)

//go:embed input.txt
var input string

func unescape(str string) string {
	s, _ := strconv.Unquote(str)
	return s
}

func escape(str string) string {
	return strconv.Quote(str)
}

func part1(input string) int {
	lines := utils.ReadLines(input)
	total := 0

	for _, str := range lines {
		total += len(str) - len(unescape(str))
	}

	return total
}

func part2(input string) int {
	lines := utils.ReadLines(input)
	total := 0

	for _, str := range lines {
		total += len(escape(str)) - len(str)
	}

	return total
}

func Solve(part int) int {
	if part == 1 {
		return part1(input)
	} else if part == 2 {
		return part2(input)
	} else {
		fmt.Println("Invalid part to run")
		return -1
	}
}
