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

const stateOn = "#"
const stateOff = "."

func part1(input string, steps int) int {
	grid := initGrid(input)

	for i := 0; i < steps; i++ {
		grid = changeLights(grid)
	}

	return countLightsOn(grid)
}

func part2(input string, steps int) int {
	grid := initGrid(input)

	rows := len(grid) - 1
	cols := len(grid[0]) - 1

	grid[0][0] = stateOn
	grid[0][cols] = stateOn
	grid[rows][0] = stateOn
	grid[rows][cols] = stateOn

	for i := 0; i < steps; i++ {
		grid = changeLights(grid)

		grid[0][0] = stateOn
		grid[0][cols] = stateOn
		grid[rows][0] = stateOn
		grid[rows][cols] = stateOn
	}

	return countLightsOn(grid)
}

func changeLights(grid [][]string) [][]string {
	var changedGrid [][]string

	for r, row := range grid {
		changedGrid = append(changedGrid, make([]string, len(grid[0])))
		for c, cell := range row {
			neighbours := countOnNeighbours(grid, r, c)
			changedGrid[r][c] = setLightState(cell, neighbours)
		}
	}

	return changedGrid
}

func countOnNeighbours(grid [][]string, r int, c int) int {
	var neighbours int

	for rDiff := -1; rDiff <= 1; rDiff++ {
		for cDiff := -1; cDiff <= 1; cDiff++ {
			if !(rDiff == 0 && cDiff == 0) {
				nextRow := r + rDiff
				nextCol := c + cDiff

				if nextRow >= 0 &&
					nextRow < len(grid) &&
					nextCol >= 0 &&
					nextCol < len(grid[0]) &&
					grid[nextRow][nextCol] == stateOn {
					neighbours++
				}
			}
		}
	}

	return neighbours
}

func initGrid(input string) [][]string {
	lines := utils.ReadLines(input)

	grid := make([][]string, len(lines))
	for i, line := range lines {
		grid[i] = strings.Split(line, "")
	}

	return grid
}

func countLightsOn(grid [][]string) int {
	var cnt int

	for _, row := range grid {
		for _, cell := range row {
			if cell == stateOn {
				cnt++
			}
		}
	}

	return cnt
}

func setLightState(current string, neighbors int) string {
	lightState := stateOff

	if (current == stateOn && (neighbors == 2 || neighbors == 3)) || (current == stateOff && neighbors == 3) {
		lightState = stateOn
	}

	return lightState
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "which part tu run? 1 or 2")
	flag.Parse()

	res := 0
	if part == 1 {
		res = part1(input, 100)
	} else {
		res = part2(input, 100)
	}

	fmt.Println("Result: ", res)
}
