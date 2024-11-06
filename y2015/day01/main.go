package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func part1(input string) int {
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

func part2(input string) int {
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
