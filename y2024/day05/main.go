package main

import (
	_ "embed"
	"flag"
	"fmt"
	"golang.org/x/exp/slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func part1(input string) int {
	rules, pages := readInput(input)
	rRules := composeRules(rules)
	pageInd := pageIndices(rRules, pages)
	return resolvePart1(rRules, pageInd, pages)
}

func part2(input string) int {
	rules, pages := readInput(input)
	rRules := composeRules(rules)
	pageInd := pageIndices(rRules, pages)
	return resolvePart2(rRules, pageInd, pages)
}

func readInput(input string) ([][]int, [][]int) {
	lines := []string{}
	for _, line := range strings.Split(input, "\n\n") {
		if string(line) != "" {
			lines = append(lines, string(line))
		}
	}

	sections := [][][]int{}
	for i, section := range lines {
		nums := [][]int{}
		lineStrs := strings.Split(section, "\n")
		separator := ","
		if i == 0 {
			separator = "|"
		}
		for _, lineStr := range lineStrs {
			if lineStr == "" {
				continue
			}
			numL := []int{}
			for _, numStr := range strings.Split(lineStr, separator) {
				num, _ := strconv.Atoi(numStr)
				numL = append(numL, num)
			}
			nums = append(nums, numL)
		}
		sections = append(sections, nums)
	}

	return sections[0], sections[1]
}

func composeRules(rulesList [][]int) map[int][]int {
	rules := make(map[int][]int)
	for _, rule := range rulesList {
		rules[rule[0]] = append(rules[rule[0]], rule[1])
	}
	return rules
}

func pageIndices(rules map[int][]int, pages [][]int) map[int][]int {
	nums := make(map[int]bool)
	for num, list := range rules {
		nums[num] = true
		for _, elem := range list {
			if !nums[elem] {
				nums[elem] = true
			}
		}
	}

	numIndices := make(map[int][]int)
	for num, _ := range nums {
		for _, numLine := range pages {
			index := -1
			for i, n := range numLine {
				if n == num {
					index = i
				}
			}
			numIndices[num] = append(numIndices[num], index)
		}
	}
	return numIndices
}

func resolvePart1(rules, numIndices map[int][]int, pages [][]int) int {

	score := 0
	for index, pageLine := range pages {
		ordered := true
		for _, num1 := range pageLine {
			rule := rules[num1]
			index1 := numIndices[num1][index]
			for _, num2 := range rule {
				index2 := numIndices[num2][index]
				if index1 == -1 || index2 == -1 {
					continue
				}
				if index1 > index2 {
					//fmt.Println(pageLine, num1, num2, index1, index2)
					ordered = false
				}
			}
		}
		if ordered {
			score += pageLine[int(len(pageLine)/2)]
		}
	}
	return score
}

func correctPageOrder(line []int, rules map[int][]int) []int {
	newLine := []int{}
	for _, num := range line {
		index := make(map[int]int)
		for i, n := range newLine {
			index[n] = i
		}
		newInsertIndex := len(newLine)
		for _, rule := range rules[num] {
			if idx, ok := index[rule]; ok {
				if newInsertIndex > idx {
					newInsertIndex = idx
				}
			}
		}
		afterNum := slices.Clone(newLine[newInsertIndex:])
		newLine = append(newLine[:newInsertIndex], num)
		newLine = append(newLine, afterNum...)
	}
	return newLine
}

func resolvePart2(rules, numIndices map[int][]int, pages [][]int) int {
	score := 0
	for index, pageLine := range pages {
		ordered := true
		for _, num1 := range pageLine {
			rule := rules[num1]
			index1 := numIndices[num1][index]
			for _, num2 := range rule {
				index2 := numIndices[num2][index]
				if index1 == -1 || index2 == -1 {
					continue
				}
				if index1 > index2 {
					ordered = false
				}
			}
		}
		if !ordered {
			newLine := correctPageOrder(pageLine, rules)
			score += newLine[len(newLine)/2]
		}
	}
	return score
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
