package day05

import (
	_ "embed"
	"fmt"
	"golang.org/x/exp/slices"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func readMappings(file string) ([]int, map[int][][]int) {
	lines := strings.Split(file, "\n\n")
	seedsData := strings.TrimSpace((strings.Split(lines[0], ":"))[1])

	seeds := []int{}
	matches := regexp.MustCompile(`-?\d+`).FindAllString(seedsData, -1)
	for _, match := range matches {
		seed, _ := strconv.Atoi(match)
		seeds = append(seeds, seed)
	}

	maps := getMaps(lines)
	return seeds, maps
}

func getMaps(lines []string) map[int][][]int {
	maps := make(map[int][][]int)

	for i := 1; i < len(lines); i++ {
		mapData := [][]int{}
		for _, line := range strings.Split(lines[i], "\n") {
			matches := regexp.MustCompile(`-?\d+`).FindAllString(line, -1)
			mapV := []int{}

			for _, match := range matches {
				num, err := strconv.Atoi(match)
				if err == nil {
					mapV = append(mapV, num)
				}
			}
			if len(mapV) > 0 {
				mapData = append(mapData, mapV)
			}
		}

		maps[i] = mapData
	}

	return maps
}

func checkSeed(seed int, mapping [][]int) int {
	for _, m := range mapping {
		d := m[0]
		s := m[1]
		r := m[2]

		if s <= seed && seed < (s+r) {
			return d + (seed - s)
		}
	}
	return seed
}

func checkSeedInterval(seed int, mapping [][]int) int {
	for _, m := range mapping {
		d := m[0]
		s := m[1]
		r := m[2]

		if d <= seed && seed < (d+r) {
			return s + (seed - d)
		}
	}
	return seed
}

func part1(input string) int {
	seeds, maps := readMappings(input)
	results := []int{}

	for _, seed := range seeds {
		for i := 1; i <= len(maps); i++ {
			seed = checkSeed(seed, maps[i])
		}
		results = append(results, seed)
	}

	return slices.Min(results)
}

func part2(input string) int {
	seeds, maps := readMappings(input)

	for n := 0; ; n++ {
		seedValue := n

		for j := len(maps); j > 0; j-- {
			seedValue = checkSeedInterval(seedValue, maps[j])
		}

		for i := 0; i < len(seeds); i += 2 {
			x := seeds[i]
			y := seeds[i+1]

			if x <= seedValue && seedValue < x+y {
				return n
			}
		}
	}
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
