package day10

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func part1(input string, iterations int) int {
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

func Solve(part int) int {
	if part == 1 {
		return part1(input, 40)
	} else if part == 2 {
		return part1(input, 50)
	} else {
		fmt.Println("Invalid part to run")
		return -1
	}
}
