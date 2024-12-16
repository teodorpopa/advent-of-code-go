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

type point struct {
	y int
	x int
}

func readInput(input string) (map[string][]point, [][]string) {
	antennas := make(map[string][]point)
	grid := make([][]string, 0)

	lines := utils.ReadLines(input)
	for y, l := range lines {
		line := strings.Split(l, "")
		grid = append(grid, line)

		for x, c := range line {
			if c == "." {
				continue
			}

			ant, ok := antennas[c]

			p := point{
				y,
				x,
			}

			if ok {
				points := ant
				points = append(points, p)
				antennas[c] = points
			} else {
				antennas[c] = []point{p}
			}
		}
	}

	return antennas, grid
}

func part1(input string) int {
	antennas, grid := readInput(input)
	antinodeMap := make(map[string]bool)

	for _, points := range antennas {
		for i, p1 := range points {
			for j := i + 1; j < len(points); j++ {
				p2 := points[j]

				pointA := point{2*p1.y - p2.y, 2*p1.x - p2.x}

				if utils.ValidIndex(pointA.x, pointA.y, len(grid), len(grid[0])) {
					if grid[pointA.y][pointA.x] == "." {
						grid[pointA.y][pointA.x] = "#"
					}
					antinodeMap[fmt.Sprint(pointA.x, "x", pointA.y)] = true
				}

				pointB := point{2*p2.y - p1.y, 2*p2.x - p1.x}

				if utils.ValidIndex(pointB.x, pointB.y, len(grid), len(grid[0])) {
					if grid[pointB.y][pointB.x] == "." {
						grid[pointB.y][pointB.x] = "#"
					}
					antinodeMap[fmt.Sprint(pointB.x, "x", pointB.y)] = true
				}
			}
		}
	}

	return len(antinodeMap)
}

func part2(input string) int {
	antennas, grid := readInput(input)

	for _, points := range antennas {
		for _, p1 := range points {
			for _, p2 := range points {
				if p1 == p2 {
					continue
				}

				diff := point{p1.y - p2.y, p1.x - p2.x}
				tryLoc := point{p1.y + diff.y, p1.x + diff.x}

				for utils.ValidIndex(tryLoc.x, tryLoc.y, len(grid[0]), len(grid)) {
					if grid[tryLoc.y][tryLoc.x] == "." {
						grid[tryLoc.y][tryLoc.x] = "#"
					}
					tryLoc = point{tryLoc.y + diff.y, tryLoc.x + diff.x}
				}
			}
		}
	}

	count := 0
	for _, row := range grid {
		fmt.Println(row)
		for _, col := range row {
			if col != "." {
				count += 1
			}
		}
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
