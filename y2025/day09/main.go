package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"
	"github.com/teodorpopa/advent-of-code-go/utils"
)

//go:embed input.txt
var input string

func parseGrid(input string) (grid [][2]int) {
	for _, line := range utils.ReadLines(input) {
		parts := strings.Split(line, ",")
		grid = append(grid, [2]int{
			utils.ToInt(parts[0]),
			utils.ToInt(parts[1]),
		})
	}
	return grid
}

func checkHorizontal(edges [][][2]int, x1, x2, y1, y2 int) bool {
	for _, edge := range edges {
		row := edge[0][1]

		if row <= y1 || row >= y2 {
			continue
		}

		edgeLeft := min(edge[0][0], edge[1][0])
		edgeRight := max(edge[0][0], edge[1][0])

		if edgeLeft <= x1 && edgeRight > x1 {
			return false
		} else if edgeLeft < x2 && edgeRight >= x2 {
			return false
		}
	}
	return true
}

func checkVertical(edges [][][2]int, x1, x2, y1, y2 int) bool {
	for _, edge := range edges {
		col := edge[0][0]

		if col <= x1 || col >= x2 {
			continue
		}

		edgeTop := min(edge[0][1], edge[1][1])
		edgeBottom := max(edge[0][1], edge[1][1])

		if edgeTop <= y1 && edgeBottom > y1 {
			return false
		} else if edgeTop < y2 && edgeBottom >= y2 {
			return false
		}
	}
	return true
}

func readEdges(grid [][2]int) (hEdges, vEdges [][][2]int) {
	for i := range len(grid) {
		p1, p2 := grid[i], grid[(i+1) % len(grid)]
		if p1[0] == p2[0] {
			vEdges = append(vEdges, [][2]int{
				p1, p2,
			})
		} else {
			hEdges = append(hEdges, [][2]int{
				p1, p2,
			})
		}
	}

	return hEdges, vEdges
}

func part1(input string) int {
	grid := parseGrid(input)

	area := 0
	for i, p1 := range grid {
		for _, p2 := range grid[i+1:] {

			x1 := min(p1[0], p2[0])
			x2 := max(p1[0], p2[0])
			y1 := min(p1[1], p2[1])
			y2 := max(p1[1], p2[1])
			size := (y2 - y1 + 1) * (x2 - x1 + 1)

			area = max(area, size)
		}
	}

	return area
}

func part2(input string) int {
	grid := parseGrid(input)
	hEdges, vEdges := readEdges(grid)

	area := 0
	for i, p1 := range grid {
		for _, p2 := range grid[i+1:] {

			x1 := min(p1[0], p2[0])
			x2 := max(p1[0], p2[0])
			y1 := min(p1[1], p2[1])
			y2 := max(p1[1], p2[1])
			size := (y2 - y1 + 1) * (x2 - x1 + 1)

			if size <= area {
				continue
			}

			checkH := checkHorizontal(hEdges, x1, x2, y1, y2)
			if !checkH {
				continue
			}

			checkV := checkVertical(vEdges, x1, x2, y1, y2)
			if !checkV {
				continue
			}

			area = size
		}
	}

	return area
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
