package day12

import (
	_ "embed"
	"fmt"
	"github.com/teodorpopa/advent-of-code-go/utils"
	"strings"
)

//go:embed input.txt
var input string

func findSprings(combinations []int, springMap string, sIndex int, cIndex int, history [][]int) int {
	total := 0

	if sIndex >= len(springMap) {
		if cIndex < len(combinations) {
			return 0
		}
		return 1
	}

	if history[sIndex][cIndex] != -1 {
		return history[sIndex][cIndex]
	}

	if springMap[sIndex] == '.' {
		total = findSprings(
			combinations,
			springMap,
			sIndex+1,
			cIndex,
			history,
		)
	} else {
		if springMap[sIndex] == '?' {
			total += findSprings(
				combinations,
				springMap,
				sIndex+1,
				cIndex,
				history,
			)
		}
		if cIndex < len(combinations) {
			count := 0
			for k := sIndex; k < len(springMap); k++ {
				if count > combinations[cIndex] || springMap[k] == '.' || count == combinations[cIndex] && springMap[k] == '?' {
					break
				}
				count += 1
			}

			if count == combinations[cIndex] {
				if sIndex+count < len(springMap) && springMap[sIndex+count] != '#' {
					total += findSprings(
						combinations,
						springMap,
						sIndex+count+1,
						cIndex+1,
						history,
					)
				} else {
					total += findSprings(
						combinations,
						springMap,
						sIndex+count,
						cIndex+1,
						history,
					)
				}
			}
		}
	}

	history[sIndex][cIndex] = total
	return total
}

func createHistoryMap(springMap string, combinations []int) [][]int {
	var history [][]int
	for i := 0; i < len(springMap); i++ {
		history = append(history, make([]int, len(combinations)+1))
		for j := 0; j < len(combinations)+1; j++ {
			history[i][j] = -1
		}
	}
	return history
}

func part1(input string) int {
	total := 0
	lines := utils.ReadLines(input)

	for _, line := range lines {
		parts := strings.Split(line, " ")

		combinations := utils.ToIntSlice(strings.Split(parts[1], ","))
		springMap := parts[0]
		history := createHistoryMap(springMap, combinations)

		total += findSprings(combinations, springMap, 0, 0, history)
	}

	return total
}

func part2(input string) int {
	total := 0
	lines := utils.ReadLines(input)

	for _, line := range lines {
		parts := strings.Split(line, " ")

		combinations := utils.ToIntSlice(strings.Split(parts[1], ","))
		springMap := parts[0]

		expandCombinations := make([]int, len(combinations)*5)
		for i := 0; i < len(combinations)*5; i++ {
			expandCombinations[i] = combinations[i%len(combinations)]
		}

		var expandSpringMap strings.Builder
		for i := 0; i < len(springMap)*5; i++ {
			if i != 0 && i%len(springMap) == 0 {
				expandSpringMap.WriteByte('?')
			}
			expandSpringMap.WriteByte(springMap[i%len(springMap)])
		}

		history := createHistoryMap(expandSpringMap.String(), expandCombinations)

		total += findSprings(expandCombinations, expandSpringMap.String(), 0, 0, history)
	}

	return total
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
