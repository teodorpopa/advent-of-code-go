package main

import (
	_ "embed"
	"flag"
	"fmt"

	"github.com/teodorpopa/advent-of-code-go/utils"
)

//go:embed input.txt
var input string

func countRolls(mapping map[int]map[int]string) (int, map[int]map[int]bool) {
	rolls := 0
	toRemove := make(map[int]map[int]bool)

	for i := 0; i < len(mapping); i++ {
		toRemove[i] = make(map[int]bool)
		for j := 0; j < len(mapping[i]); j++ {
			if mapping[i][j] == "@" {
				adjacent := countAdjacent(mapping, i, j)
				if adjacent < 4 {
					rolls++
					toRemove[i][j] = true
				}
			}
		}
	}

	return rolls, toRemove
}

func countAdjacent(mapping map[int]map[int]string, i, j int) int {
	count := 0

	for _, dir := range [][]int{{0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}, {-1, -1}, {-1, 0}, {-1, 1}} {
		x, y := i+dir[0], j+dir[1]
		if x >= 0 && x < len(mapping) && y >= 0 && y < len(mapping[x]) && mapping[x][y] == "@" {
			count++
		}
	}

	return count
}

func part1(input string) int {

	lines := utils.ReadLines(input)
	rollsMap := make(map[int]map[int]string)

	for i, line := range lines {
		rollsMap[i] = make(map[int]string)
		for j, char := range line {
			rollsMap[i][j] = string(char)
		}
	}

	rolls, _ := countRolls(rollsMap)

	return rolls
}

func part2(input string) int {
	lines := utils.ReadLines(input)
	rollsMap := make(map[int]map[int]string)
	rolls := 0

	for i, line := range lines {
		rollsMap[i] = make(map[int]string)
		for j, char := range line {
			rollsMap[i][j] = string(char)
		}
	}

	for {
		cnt, toRemove := countRolls(rollsMap)

		if cnt == 0 {
			break
		}

		rolls += cnt

		for i := range toRemove {
			for j := range toRemove[i] {
				rollsMap[i][j] = "."
			}
		}
	}

	return rolls
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
