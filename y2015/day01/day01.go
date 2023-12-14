package day01

import (
	_ "embed"
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
