package main

import (
	_ "embed"
	"flag"
	"fmt"
	"github.com/teodorpopa/advent-of-code-go/utils"
	"math"
)

//go:embed input.txt
var input string

func part1(input string, qty int) int {
	return len(inputCombinations(input, qty))
}

func part2(input string, qty int) int {
	combs := inputCombinations(input, qty)

	minLen := math.MaxInt32
	for _, comb := range combs {
		minLen = utils.Min(minLen, len(comb))
	}

	var count int
	for _, comb := range combs {
		if len(comb) == minLen {
			count++
		}
	}

	return count
}

func toIntInput(input string) []int {
	var nums []int
	lines := utils.ReadLines(input)

	for _, line := range lines {
		nums = append(nums, utils.ToInt(line))
	}
	return nums
}

func tryCombinations(nums []int, sIndex, total int, used []int) [][]int {
	if total == 0 {
		return [][]int{append([]int{}, used...)}
	}
	if total < 0 {
		return nil
	}

	var validReturns [][]int
	for i := sIndex; i < len(nums); i++ {
		used = append(used, i)
		validReturns = append(validReturns, tryCombinations(nums, i+1, total-nums[i], used)...)
		used = used[:len(used)-1]
	}
	return validReturns
}

func inputCombinations(input string, total int) [][]int {
	var nums = toIntInput(input)
	return tryCombinations(nums, 0, total, []int{})
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "which part tu run? 1 or 2")
	flag.Parse()

	res := 0
	if part == 1 {
		res = part1(input, 150)
	} else {
		res = part2(input, 150)
	}

	fmt.Println("Result: ", res)
}
