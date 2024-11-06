package main

import (
	_ "embed"
	"flag"
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

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "which part tu run? 1 or 2")
	flag.Parse()

	res := 0
	if part == 1 {
		res = part1(input)
	} else {
		res = part2(input)
	}

	fmt.Println("Result: ", res)
}
