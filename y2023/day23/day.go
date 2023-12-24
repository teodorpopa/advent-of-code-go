package day23

import (
	_ "embed"
	"fmt"
	"github.com/teodorpopa/advent-of-code-go/utils"
)

//go:embed input.txt
var input string

type Point struct {
	x, y int
}

type PathTo struct {
	end           Point
	length, index int
}

func getNeighbours(grid []string) map[Point]bool {
	nb := map[Point]bool{}
	for i, line := range grid {
		for j, char := range line {
			if char == '#' {
				continue
			}
			point := Point{j, i}
			neighbours := 0

			for _, dir := range [4]Point{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
				next := Point{point.x + dir.x, point.y + dir.y}
				if inLimits(grid, next) && grid[next.y][next.x] != '#' {
					neighbours++
				}
			}

			if neighbours > 2 {
				nb[point] = true
			}
		}
	}
	return nb
}

func getPaths(grid []string, junctions map[Point]bool) map[Point][]PathTo {
	paths := map[Point][]PathTo{}
	junctionIndex := 0
	for junctionPoint := range junctions {
		for _, startDir := range [4]Point{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
			currentPoint := Point{junctionPoint.x + startDir.x, junctionPoint.y + startDir.y}
			if inLimits(grid, currentPoint) && grid[currentPoint.y][currentPoint.x] != '#' {
				path := getPath(grid, junctionPoint, currentPoint, startDir, 1, junctions)
				path.index = junctionIndex
				paths[junctionPoint] = append(paths[junctionPoint], path)
			}
		}
		junctionIndex++
	}
	return paths
}

func getPath(grid []string, pathStart, currentPoint, currentDir Point, pathLength int, junctions map[Point]bool) PathTo {
	for _, dir := range [3]Point{currentDir, left(currentDir), right(currentDir)} {
		next := Point{currentPoint.x + dir.x, currentPoint.y + dir.y}
		if grid[next.y][next.x] != '#' {
			if _, found := junctions[next]; found {
				return PathTo{next, pathLength + 1, 0}
			} else {
				return getPath(grid, pathStart, next, dir, pathLength+1, junctions)
			}
		}
	}
	return PathTo{Point{-1, -1}, 0, 0}
}

func findLongestPath(grid []string, paths map[Point][]PathTo, start, end Point, step int, visited []bool) int {
	maxStep := 0
	for _, path := range paths[start] {
		index := paths[path.end][0].index
		if !visited[index] {
			if path.end == end {
				return step + path.length
			}
			visited[index] = true
			maxStep = utils.Max(maxStep, findLongestPath(grid, paths, path.end, end, step+path.length, visited))
			visited[index] = false
		}
	}
	return maxStep
}

func findPath(grid []string, start, currentDir Point, visited map[Point]int) {
	current := start
	currentStep := visited[current]

	for _, dir := range [3]Point{currentDir, left(currentDir), right(currentDir)} {
		next := Point{current.x + dir.x, current.y + dir.y}
		if inLimits(grid, next) && grid[next.y][next.x] != '#' {
			char := grid[next.y][next.x]
			oppositeChar := map[Point]byte{{1, 0}: '<', {-1, 0}: '>', {0, 1}: '^', {0, -1}: 'v'}
			if oppositeChar[dir] == char {
				continue
			}

			if val, found := visited[next]; !found || val < currentStep+1 {
				visited[next] = currentStep + 1
				findPath(grid, next, dir, visited)
			}
		}
	}
}

func inLimits(grid []string, pos Point) bool {
	return pos.x >= 0 && pos.x < len(grid[0]) && pos.y >= 0 && pos.y < len(grid)
}

func left(p Point) Point {
	return Point{p.y, -p.x}
}

func right(p Point) Point {
	return Point{-p.y, p.x}
}

func getStartEndPoints(grid []string) (Point, Point) {
	return Point{1, 0}, Point{len(grid[0]) - 2, len(grid) - 1}
}

func part1(input string) int {
	grid := utils.ReadLines(input)
	start, end := getStartEndPoints(grid)

	visited := map[Point]int{start: 0}
	currentDir := Point{0, 1}
	findPath(grid, start, currentDir, visited)

	return visited[end]
}

func part2(input string) int {
	grid := utils.ReadLines(input)
	start, end := getStartEndPoints(grid)

	neighbours := getNeighbours(grid)
	neighbours[start] = true
	neighbours[end] = true

	paths := getPaths(grid, neighbours)
	visited := make([]bool, len(neighbours))
	visited[paths[start][0].index] = true
	return findLongestPath(grid, paths, start, end, 0, visited)
}

func Solve(part int) int {
	if part == 1 {
		return part1(input)
	} else if part == 2 {
		return part2(input)
	} else {
		fmt.Println("Invalid part to run")
		return -1
	}
}
