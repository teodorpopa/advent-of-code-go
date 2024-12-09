package main

import (
	_ "embed"
	"flag"
	"fmt"
	"github.com/teodorpopa/advent-of-code-go/utils"
	"math"
	"slices"
	"strings"
)

//go:embed input.txt
var input string

func ReadLists(input string) ([]int, []int) {
	lines := utils.ReadLines(input)

	left := []int{}
	right := []int{}

	for _, line := range lines {
		parts := strings.Split(line, "   ")

		left = append(left, utils.ToInt(parts[0]))
		right = append(right, utils.ToInt(parts[1]))
	}

	slices.Sort(left)
	slices.Sort(right)

	return left, right
}

func part1(input string) int {

	left, right := ReadLists(input)
	diff := 0

	for i := range left {
		difference := int(math.Abs(float64(right[i] - left[i])))
		diff += difference
	}

	return diff
}

func part2(input string) int {

	left, right := ReadLists(input)
	similarity := 0

	for _, l := range left {
		instances := 0
		for _, r := range right {
			if l == r {
				instances++
			}
		}
		similarity += l * instances
	}

	return similarity
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
