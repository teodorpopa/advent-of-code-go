package main

import (
	_ "embed"
	"flag"
	"fmt"
	"github.com/teodorpopa/advent-of-code-go/utils"
	"slices"
	"strings"
)

//go:embed input.txt
var input string

var grid [][]int

type point struct {
	y int
	x int
}

var trailheads = make(map[point][]point)

func readInput(input string) {
	lines := utils.ReadLines(input)
	for _, l := range lines {
		line := strings.Split(l, "")
		grid = append(grid, utils.ToIntSlice(line))
	}
}

func countTrails(grid [][]int, y, x, next int, start point, part int) {
	if y < 0 || y >= len(grid[0]) || x < 0 || x >= len(grid) {
		return
	}

	curr := grid[y][x]
	if curr != next {
		return
	}

	if curr == 9 && next == 9 {
		if part == 2 {
			trailheads[start] = append(trailheads[start], point{y, x})
		}
		positions := trailheads[start]
		if !slices.Contains(positions, point{y, x}) {
			trailheads[start] = append(trailheads[start], point{y, x})
		}
		return
	}
	countTrails(grid, y+1, x, next+1, start, part)
	countTrails(grid, y-1, x, next+1, start, part)
	countTrails(grid, y, x-1, next+1, start, part)
	countTrails(grid, y, x+1, next+1, start, part)
}

func part1(input string) int {
	readInput(input)

	count := 0
	for y, row := range grid {
		for x, item := range row {
			if item == 0 {
				trailheads[point{y, x}] = make([]point, 0)
				countTrails(grid, y, x, 0, point{y, x}, 1)
			}
		}
	}

	for _, item := range trailheads {
		count += len(item)
	}

	return count
}

func part2(input string) int {
	readInput(input)

	count := 0
	for y, row := range grid {
		for x, item := range row {
			if item == 0 {
				trailheads[point{y, x}] = make([]point, 0)
				countTrails(grid, y, x, 0, point{y, x}, 2)
			}
		}
	}

	for _, item := range trailheads {
		count += len(item)
	}

	return count
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
