package main

import (
	_ "embed"
	"flag"
	"fmt"
	"github.com/teodorpopa/advent-of-code-go/utils"
	"strings"
)

//go:embed input.txt
var input string

func readInput(input string) map[int][]int {
	lines := utils.ReadLines(input)

	rows := make(map[int][]int)

	for _, l := range lines {
		parts := strings.Split(l, ": ")

		total := utils.ToInt(parts[0])
		nums := strings.Split(parts[1], " ")

		var ns []int
		for _, n := range nums {
			ni := utils.ToInt(n)
			ns = append(ns, ni)
		}

		rows[total] = ns
	}

	return rows
}

func part1(input string) int {
	rows := readInput(input)

	var sums int
	for total, nums := range rows {
		if canTotal(nums, total, 0, 0) {
			sums += total
		}
	}

	return sums
}

func part2(input string) int {
	rows := readInput(input)

	var sums int
	for total, nums := range rows {
		if canTotalConcat(nums, total, 0, 0) {
			sums += total
		}

	}

	return sums
}

func canTotal(nums []int, target int, current int, index int) bool {
	if index == 0 {
		return canTotal(nums, target, nums[0], 1)
	}

	i := nums[index]
	nextSum := current + i
	nextMult := current * i

	if index == len(nums)-1 {
		return (nextSum == target) || (nextMult == target)
	}

	return canTotal(nums, target, nextSum, index+1) ||
		canTotal(nums, target, nextMult, index+1)
}

func canTotalConcat(nums []int, target int, current int, index int) bool {
	if index == 0 {
		return canTotalConcat(nums, target, nums[0], 1)
	}

	i := nums[index]
	nextSum := current + i
	nextMult := current * i

	if index == len(nums)-1 {
		return (nextSum == target) || (nextMult == target) ||
			(utils.ConcatTwoInt(current, i) == target)
	}

	nextIndex := index + 1

	return canTotalConcat(nums, target, nextSum, nextIndex) ||
		canTotalConcat(nums, target, nextMult, nextIndex) ||
		canTotalConcat(nums, target, utils.ConcatTwoInt(current, i), nextIndex)
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
