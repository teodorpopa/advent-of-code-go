package main

import (
	_ "embed"
	"flag"
	"fmt"
	"github.com/teodorpopa/advent-of-code-go/utils"
	"sort"
)

//go:embed input.txt
var input string

func parseInput(input string) [][]string {
	lines := utils.ReadLines(input)

	grid := make([][]string, len(lines))
	for i, line := range lines {
		grid[i] = make([]string, len(line))
		for j, c := range line {
			grid[i][j] = string(c)
		}
	}

	return grid
}

func visitedGrid(grid [][]string) [][]bool {
	visited := make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[i]))
	}
	return visited
}

func sideCount(sides [][3]int) int {
	sideMap := make(map[[3]int]bool)

	sort.Slice(sides, func(i, j int) bool {
		if sides[i][0] == sides[j][0] {
			return sides[i][1] < sides[j][1]
		}
		return sides[i][0] < sides[j][0]
	})

	sideCount := 0
	for _, s := range sides {
		getCombinations := utils.GetNeighbors(s[0], s[1])
		combFound := false

		for _, c := range getCombinations {
			c[2] = s[2]
			if _, found := sideMap[c]; found {
				combFound = true
			}
		}
		if !combFound {
			sideCount++
		}

		sideMap[s] = true

	}

	return sideCount
}

func part1(input string) int {
	grid := parseInput(input)
	visited := visitedGrid(grid)

	price := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if !visited[i][j] {
				a, p := utils.DFS(grid, visited, i, j, nil)
				price += a * p
			}
		}
	}
	return price
}

func part2(input string) int {
	grid := parseInput(input)
	visited := visitedGrid(grid)

	price := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if !visited[i][j] {
				sides := make([][3]int, 0)
				a, _ := utils.DFS(grid, visited, i, j, &sides)
				price += a * sideCount(sides)
			}
		}
	}
	return price
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
