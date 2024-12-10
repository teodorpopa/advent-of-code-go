package main

import (
	_ "embed"
	"flag"
	"fmt"
	"github.com/teodorpopa/advent-of-code-go/utils"
	"regexp"
	"strings"
)

//go:embed input.txt
var input string

func findAndCalculateMultiplications(input string) int {
	total := 0

	multiplicationsRegex := regexp.MustCompile(`mul\(\d+,\d+\)`)
	numRegex := regexp.MustCompile(`(\d+)`)

	matches := multiplicationsRegex.FindAllString(input, -1)
	for _, match := range matches {
		nums := numRegex.FindAllString(match, -1)
		total += utils.ToInt(nums[0]) * utils.ToInt(nums[1])
	}

	return total
}

func part1(input string) int {
	return findAndCalculateMultiplications(input)
}

func part2(input string) int {
	matches := ""

	parts := strings.Split(input, "don't()")
	for i, part := range parts {
		if i == 0 {
			matches += part
		} else {
			_, after, _ := strings.Cut(part, "do()")
			if len(after) > 1 {
				matches += after
			}
		}
	}

	return findAndCalculateMultiplications(matches)
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
