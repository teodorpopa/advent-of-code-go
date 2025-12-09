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

var splits int

func part1(input string) int {
	lines := utils.ReadLines(input)
	design := make([][]string, len(lines))
	for i, line := range lines {
		design[i] = strings.Split(line, "")
	}

	loop := true
	for loop {
		loop = splitBeam(&design)
	}

	return splits
}

func splitBeam(design *[][]string) bool {
Outer:
	for i, row := range *design {
		for j, col := range row {
			if col == "S" {
				(*design)[i+1][j] = "|"
				continue Outer
			} else if col == "^" {
				if (*design)[i-1][j] == "|" {
					splits++
				}
				(*design)[i][j-1] = "|"
				(*design)[i][j+1] = "|"
			} else if col == "|" && (*design)[i][j-1] != "^" {
				if (*design)[i+1][j] == "." {
					(*design)[i+1][j] = "|"
				}
			} else if col == "." && i > 0 {
				if j >= 1 && (*design)[i-1][j] == "|" {
					(*design)[i][j] = "|"
				}
				if (*design)[i-1][j] == "|" {
					(*design)[i][j] = "|"
				}
			}
		}
		//display(i+1, design)

		if i == len(*design)-1 {
			return false
		}

		continue Outer
	}

	return true
}

// func display(iteration int, design *[][]string) {
// 	fmt.Print("\033[H\033[2J")
// 	fmt.Println("Iteration:", iteration)
// 	fmt.Println("Splits:", splits)
// 	for _, row := range *design {
// 		for _, col := range row {
// 			fmt.Print(col)
// 		}
// 		fmt.Println()
// 	}
// 	time.Sleep(10 * time.Millisecond)
// }

func part2(input string) int {
	lines := utils.ReadLines(input)
	design := make([][]string, len(lines))
	for i, line := range lines {
		design[i] = strings.Split(line, "")
	}

	cols := map[int]int{}
	for c := range len(design[0]) {
		if design[0][c] == "S" {
			cols[c] = 1
		}
	}

	for _, line := range design {
		nextCols := map[int]int{}
		for i, val := range line {
			if val == "^" && cols[i] > 0 {
				nextCols[i-1] += cols[i]
				nextCols[i+1] += cols[i]
				delete(cols, i)
			}
		}

		for c, val := range nextCols {
			cols[c] += val
		}
	}

	timelines := 0
	for _, val := range cols {
		timelines += val
	}
	return timelines
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
