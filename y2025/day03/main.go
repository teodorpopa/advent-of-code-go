package main

import (
	_ "embed"
	"flag"
	"fmt"
	"math"

	"github.com/teodorpopa/advent-of-code-go/utils"
)

//go:embed input.txt
var input string

func getJoltage(line string, n int) int64 {
	result := int64(0)
	start := 0

	for i := n; i > 0; i-- {
		max := byte('0')
		end := len(line) - i

		for j := start; j <= end; j++ {
			if line[j] > max {
				max = line[j]
				start = j + 1
			}
		}

		result += int64(max-'0') * int64(math.Pow(10.0, float64(i-1)))
	}

	return result
}

func part1(input string) int {
	lines := utils.ReadLines(input)
	sum := 0

	for _, line := range lines {
		sum += int(getJoltage(line, 2))
	}

	return sum
}

func part2(input string) int {
	lines := utils.ReadLines(input)
	sum := 0

	for _, line := range lines {
		sum += int(getJoltage(line, 12))
	}

	return sum
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
