package main

import (
	_ "embed"
	"flag"
	"fmt"
	"math"
	"strings"
)

//go:embed input.txt
var input string

var directions = [][]int{
	{1, 0},
	{0, 1},
	{-1, 0},
	{0, -1},
}

func parseInput(input string) (garden []string, x, y int) {
	garden = strings.Split(input, "\n")
	for y, l := range garden {
		for x, c := range l {
			if c == 'S' {
				return garden, x, y
			}
		}
	}
	return garden, 0, 0
}

func bfs(gardenMap []string, startX, startY, steps int, infinite bool) map[string]int {
	dst := map[string]int{
		fmt.Sprintf("%d,%d", startX, startY): 0,
	}
	seen := [][]int{{startX, startY, steps}}
	for len(seen) > 0 {
		v := seen[0]
		seen = seen[1:]
		for _, d := range directions {
			nx, ny := v[0]+d[0], v[1]+d[1]
			if !infinite {
				if ny >= 0 && ny < len(gardenMap) && nx >= 0 && nx < len(gardenMap[ny]) && gardenMap[ny][nx] != '#' {
					seen = addSteps(nx, ny, dst, v, seen)
				}
			} else {
				checkNx, checkNy := nx, ny
				if ny >= len(gardenMap) {
					checkNy = ny % len(gardenMap)
				} else if ny < 0 {
					checkNy = (ny%len(gardenMap) + len(gardenMap)) % len(gardenMap)
				}
				if nx >= len(gardenMap[checkNy]) {
					checkNx = nx % len(gardenMap[checkNy])
				} else if nx < 0 {
					checkNx = (nx%len(gardenMap[checkNy]) + len(gardenMap[checkNy])) % len(gardenMap[checkNy])
				}
				if gardenMap[checkNy][checkNx] != '#' {
					seen = addSteps(nx, ny, dst, v, seen)
				}
			}
		}
	}
	return dst
}

func addSteps(nx int, ny int, dst map[string]int, v []int, seen [][]int) [][]int {
	sw := fmt.Sprintf("%d,%d", nx, ny)
	if _, p := dst[sw]; !p && v[2]-1 >= 0 {
		seen = append(seen, []int{nx, ny, v[2] - 1})
		dst[sw] = dst[fmt.Sprintf("%d,%d", v[0], v[1])] + 1
	}
	return seen
}

func calculate(n int, p ...int) int {
	return p[0] + n*(p[1]-p[0]) + n*(n-1)/2*((p[2]-p[1])-(p[1]-p[0]))
}

func part1(input string) int {
	stepLimit := 64
	steps := 0
	gardenMap, startX, startY := parseInput(input)

	for _, d := range bfs(gardenMap, startX, startY, stepLimit, false) {
		if (d+stepLimit%2)%2 == 0 {
			steps++
		}
	}
	return steps
}

func part2(input string) int {
	params := []int{}

	stepLimit := 26501365
	gardenMap, startX, startY := parseInput(input)

	for i := 0; i < len(gardenMap)*3; i++ {
		if i%len(gardenMap) == int(math.Floor(float64(len(gardenMap))/2)) {
			r := 0
			for _, d := range bfs(gardenMap, startX, startY, i, true) {
				if (d+i%2)%2 == 0 {
					r++
				}
			}
			params = append(params, r)
		}
	}

	return calculate(int(math.Floor(float64(stepLimit)/float64(len(gardenMap)))), params...)
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
