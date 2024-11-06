package main

import (
	_ "embed"
	"flag"
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
