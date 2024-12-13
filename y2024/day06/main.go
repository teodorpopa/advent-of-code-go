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

var grid [][]string
var startIndex point

type point struct {
	y int
	x int
}

var directions = []point{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

func part1(input string) int {
	readInput(input)
	_, res := getPath(grid, startIndex)
	return res
}

func part2(input string) int {
	readInput(input)
	path, _ := getPath(grid, startIndex)
	return getObstacles(grid, startIndex, path)
}

func readInput(input string) {
	lines := utils.ReadLines(input)
	for y, l := range lines {
		line := strings.Split(l, "")
		for x, c := range line {
			if c == "^" {
				startIndex = point{y, x}
			}
		}
		grid = append(grid, line)
	}
}

func getPath(grid [][]string, startIndex point) ([]point, int) {
	maxRow := len(grid)
	maxCol := len(grid[0])
	cache := make(map[point]bool)
	at := startIndex
	way := 0

	for validPoint(at, maxRow, maxCol) {
		if grid[at.y][at.x] == "#" {
			at.y -= directions[way].y
			at.x -= directions[way].x
			way = (way + 1) % len(directions)
			continue
		}
		cache[at] = true
		grid[at.y][at.x] = "X"
		at.y += directions[way].y
		at.x += directions[way].x
	}

	visited := make([]point, 0, len(cache))
	for idx := range cache {
		visited = append(visited, idx)
	}
	return visited, len(visited)
}

func getObstacles(grid [][]string, startIndex point, path []point) int {
	bad := 0
	for _, index := range path {
		if (index == startIndex) || grid[index.y][index.x] == "#" {
			continue
		}
		grid[index.y][index.x] = "#"
		if hasCycle(grid, startIndex) {
			bad++
		}
		grid[index.y][index.x] = "."
	}

	return bad
}

func hasCycle(grid [][]string, startIndex point) bool {
	maxRow := len(grid)
	maxCol := len(grid[0])
	cache := make(map[point]point)
	at := startIndex
	way := 0

	for validPoint(at, maxRow, maxCol) {
		if cache[at] == directions[way] {
			return true
		}

		cache[at] = directions[way]

		if grid[at.y][at.x] == "#" {
			at.y -= directions[way].y
			at.x -= directions[way].x
			way = (way + 1) % len(directions)
			continue
		}

		grid[at.y][at.x] = "X"
		at.y += directions[way].y
		at.x += directions[way].x
	}

	return false
}

func validPoint(p point, maxRow, maxCol int) bool {
	return p.x >= 0 && p.y >= 0 && p.y < maxRow && p.x < maxCol
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
