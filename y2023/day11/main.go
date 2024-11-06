package main

import (
	_ "embed"
	"flag"
	"fmt"
	"github.com/teodorpopa/advent-of-code-go/utils"
	"math"
	"strings"
)

//go:embed input.txt
var input string

type Galaxy struct {
	X int
	Y int
}

func getLinesToExpand(galaxyMap [][]string) []int {
	linesToExpand := []int{}
	for y := 0; y < len(galaxyMap); y++ {
		found := false
		for x := 0; x < len(galaxyMap[y]); x++ {
			if galaxyMap[y][x] == "#" {
				found = true
				break
			}
		}
		if !found {
			linesToExpand = append(linesToExpand, y)
		}
	}
	return linesToExpand
}

func getColsToExpand(galaxyMap [][]string) []int {
	colsToExpand := []int{}
	for x := 0; x < len(galaxyMap[0]); x++ {
		found := false
		for y := 0; y < len(galaxyMap); y++ {
			if galaxyMap[y][x] == "#" {
				found = true
				break
			}
		}
		if !found {
			colsToExpand = append(colsToExpand, x)
		}
	}
	return colsToExpand
}

func getGalaxyMap(galaxyMap [][]string, linesToExpand, colsToExpand []int, universeExpansion int) []Galaxy {
	galaxies := []Galaxy{}
	for y := 0; y < len(galaxyMap); y++ {
		for x := 0; x < len(galaxyMap[y]); x++ {
			if galaxyMap[y][x] == "#" {
				expandX := 0
				expandY := 0

				for i := 0; i < len(linesToExpand); i++ {
					if y > linesToExpand[i] {
						expandY += universeExpansion - 1
					}
				}

				for i := 0; i < len(colsToExpand); i++ {
					if x > colsToExpand[i] {
						expandX += universeExpansion - 1
					}
				}

				galaxyPosition := Galaxy{
					X: x + expandX,
					Y: y + expandY,
				}
				galaxies = append(galaxies, galaxyPosition)
			}
		}
	}
	return galaxies
}

func distance(a Galaxy, b Galaxy) int {
	return int(math.Abs(float64(b.X)-float64(a.X)) + math.Abs(float64(b.Y)-float64(a.Y)))
}

func part1(input string) int {
	sum := 0
	lines := utils.ReadLines(input)

	galaxyMap := [][]string{}
	for _, line := range lines {
		galaxyMap = append(galaxyMap, strings.Split(line, ""))
	}

	linesToExpand := getLinesToExpand(galaxyMap)
	colsToExpand := getColsToExpand(galaxyMap)

	galaxies := getGalaxyMap(
		galaxyMap,
		linesToExpand,
		colsToExpand,
		2,
	)

	for i := 0; i < len(galaxies)-1; i++ {
		for j := i + 1; j < len(galaxies); j++ {
			sum += distance(galaxies[i], galaxies[j])
		}
	}

	return sum
}

func part2(input string) int {
	sum := 0
	lines := utils.ReadLines(input)

	galaxyMap := [][]string{}
	for _, line := range lines {
		galaxyMap = append(galaxyMap, strings.Split(line, ""))
	}

	linesToExpand := getLinesToExpand(galaxyMap)
	colsToExpand := getColsToExpand(galaxyMap)

	galaxies := getGalaxyMap(
		galaxyMap,
		linesToExpand,
		colsToExpand,
		1000000,
	)

	for i := 0; i < len(galaxies)-1; i++ {
		for j := i + 1; j < len(galaxies); j++ {
			sum += distance(galaxies[i], galaxies[j])
		}
	}

	return sum
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
