package main

import (
	_ "embed"
	"flag"
	"fmt"
	"slices"
	"strings"

	"github.com/teodorpopa/advent-of-code-go/utils"
)

//go:embed input.txt
var input string

func part1(input string) int {
	lines := utils.ReadLines(input)
	worksheet := [][]int{}
	sum := 0
	for _, line := range lines[0:1] {
		nums := strings.Fields(strings.Trim(line, " "))
		for range nums {
			worksheet = append(worksheet, make([]int, len(lines)-1))
		}
	}

	for i, line := range lines[:len(lines)-1] {
		nums := strings.Fields(strings.Trim(line, " "))
		for j, num := range nums {
			worksheet[j][i] = utils.ToInt(num)
		}
	}

	signs := strings.Fields(strings.Trim(lines[len(lines)-1], " "))

	for i, sheet := range worksheet {
		sign := signs[i]
		total := 0
		if sign == "*" {
			total = 1
		}

		for _, num := range sheet {
			if sign == "+" {
				total += num
			} else {
				total *= num
			}
		}
		sum += total
	}

	return sum
}

func part2(input string) int {

	lines := utils.ReadLines(input)
	sum := 0
	dataRows := lines[:len(lines)-1]

	cols := 0
	for _, r := range lines {
		if len(r) > cols {
			cols = len(r)
		}
	}

	finalNums := [][]int{}
	signs := strings.Fields(strings.Trim(lines[len(lines)-1], " "))

	finalN := []int{}

	for c := cols - 1; c >= 0; c-- {
		nums := []string{}
		for _, dr := range dataRows {
			if c < len(dr) {
				tok := dr[c]
				if string(tok) == "" {
					continue
				}
				nums = append(nums, string(tok))
			}
		}

		num := utils.ToInt(strings.Trim(strings.Join(nums, ""), " "))
		if num == 0 {
			finalNums = append(finalNums, finalN)
			finalN = []int{}
			continue
		} else {
			finalN = append(finalN, num)
		}
	}
	finalNums = append(finalNums, finalN)
	slices.Reverse(signs)

	for i, nums := range finalNums {
		total := 0
		if signs[i] == "*" {
			total = 1
		}

		for _, num := range nums {
			if signs[i] == "+" {
				total += num
			} else {
				total *= num
			}
		}
		sum += total
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
