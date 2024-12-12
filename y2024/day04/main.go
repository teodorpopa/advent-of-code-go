package main

import (
	_ "embed"
	"flag"
	"fmt"
	"github.com/teodorpopa/advent-of-code-go/utils"
)

//go:embed input.txt
var input string

func parseInput(input string) int {
	lines := utils.ReadLines(input)
	word := "XMAS"
	grid := [][]string{}

	for _, l := range lines {
		line := []string{}
		for _, c := range l {
			line = append(line, string(c))
		}
		grid = append(grid, line)
	}

	found := 0
	for y, l := range lines {
		for x, c := range l {
			if c == 'X' {
				found = found + searchWord(word, grid, x, y)
			}
		}
	}

	return found
}

func searchWord(word string, grid [][]string, x int, y int) int {
	wordLen := len(word)
	searchRange := make([]int, wordLen)
	for i := 1; i < wordLen; i++ {
		searchRange[i] = i
	}
	found := 0

	ways := []string{"N", "NE", "E", "SE", "S", "SW", "W", "NW"}

	for _, way := range ways {
		found += checkWord(word, way, x, y, grid, searchRange)
	}

	return found
}

func checkWord(word string, way string, x int, y int, grid [][]string, searchRange []int) int {
	for _, sr := range searchRange[1:] {

		switch way {
		case "N":
			y--

			if y < 0 {
				return 0
			}

			if string(word[sr]) != grid[y][x] {
				return 0
			}
		case "NE":
			y--
			x++

			if y < 0 || y > len(grid[0]) || x >= (len(grid[y])) {
				return 0
			}

			if string(word[sr]) != grid[y][x] {
				return 0
			}
		case "E":
			x++

			if x >= (len(grid[y])) {
				return 0
			}

			if string(word[sr]) != grid[y][x] {
				return 0
			}
		case "SE":
			x++
			y++

			if y >= len(grid) || x >= len(grid[y]) {
				return 0
			}

			if string(word[sr]) != grid[y][x] {
				return 0
			}
		case "S":
			y++

			if y >= len(grid) {
				return 0
			}

			if string(word[sr]) != grid[y][x] {
				return 0
			}
		case "SW":
			y++
			x--

			if y >= len(grid) || x < 0 {
				return 0
			}

			if string(word[sr]) != grid[y][x] {
				return 0
			}
		case "W":
			x--

			if x < 0 {
				return 0
			}

			if string(word[sr]) != grid[y][x] {
				return 0
			}
		case "NW":
			x--
			y--

			if x < 0 || y < 0 {
				return 0
			}

			if string(word[sr]) != grid[y][x] {
				return 0
			}
		default:
			return 0
		}

	}

	return 1
}

func part1(input string) int {
	return parseInput(input)
}

func part2(input string) int {
	lines := utils.ReadLines(input)
	grid := [][]string{}

	for _, l := range lines {
		line := []string{}
		for _, c := range l {
			line = append(line, string(c))
		}
		grid = append(grid, line)
	}

	found := 0
	for y, l := range lines {
		for x, c := range l {
			if c == 'A' {
				if x > 0 && y > 0 && x < len(grid[0])-1 && y < len(grid)-1 {
					if (grid[y+1][x-1] == "M" && grid[y+1][x+1] == "M" && grid[y-1][x-1] == "S" && grid[y-1][x+1] == "S") ||
						(grid[y+1][x-1] == "S" && grid[y+1][x+1] == "S" && grid[y-1][x-1] == "M" && grid[y-1][x+1] == "M") ||
						(grid[y+1][x-1] == "S" && grid[y+1][x+1] == "M" && grid[y-1][x-1] == "S" && grid[y-1][x+1] == "M") ||
						(grid[y+1][x-1] == "M" && grid[y+1][x+1] == "S" && grid[y-1][x-1] == "M" && grid[y-1][x+1] == "S") {
						found++
					}
				}
			}
		}
	}

	return found
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
