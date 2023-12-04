package y2023

import (
	"fmt"
	"github.com/teodorpopa/advent-of-code-go/utils"
	"regexp"
	"strconv"
	"unicode"
)

type Point struct {
	X int
	Y int
}

func (pnt Point) addPoint(dest Point) Point {
	return Point{
		pnt.X + dest.X,
		pnt.Y + dest.Y,
	}
}

func getSymbols(lines []string) map[Point]string {
	sym := map[Point]string{}

	for y, line := range lines {
		for x, c := range line {
			if c != '.' && !unicode.IsDigit(c) {
				sym[Point{x, y}] = string(c)
			}
		}
	}

	return sym
}

func getEngineSchema(lines []string, symbols map[Point]string) map[Point][]int {
	parts := map[Point][]int{}
	re := regexp.MustCompile(`\d+`)

	dirs := []Point{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}

	for y, s := range lines {
		for _, m := range re.FindAllStringIndex(s, -1) {
			keys := map[Point]bool{}
			for x := m[0]; x < m[1]; x++ {
				for _, d := range dirs {
					keys[Point{x, y}.addPoint(d)] = true
				}
			}

			n, _ := strconv.Atoi(s[m[0]:m[1]])
			for p := range keys {
				if _, exists := symbols[p]; exists {
					parts[p] = append(parts[p], n)
				}
			}
		}
	}
	return parts

}

func Day03First(input string) int {
	sum := 0
	lines := utils.ReadLines(input)
	symbols := getSymbols(lines)
	engineParts := getEngineSchema(lines, symbols)

	for _, values := range engineParts {
		for _, value := range values {
			sum += value
		}
	}
	return sum
}

func Day03Second(input string) int {
	ratio := 0
	lines := utils.ReadLines(input)
	symbols := getSymbols(lines)
	engineParts := getEngineSchema(lines, symbols)

	for index, neighbors := range engineParts {
		if symbols[index] == "*" && len(neighbors) == 2 {
			ratio += neighbors[0] * neighbors[1]
		}
	}
	return ratio
}

func Day03() {
	fmt.Println(utils.DAY_PREFIX, "Day 03")
	fmt.Println(utils.PART_PREFIX, "Part 1:", Day03First(utils.ReadFile("y2023/input/day03.txt")))
	fmt.Println(utils.PART_PREFIX, "Part 2:", Day03Second(utils.ReadFile("y2023/input/day03.txt")))
}
