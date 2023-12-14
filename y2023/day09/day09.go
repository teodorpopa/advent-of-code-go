package day09

import (
	_ "embed"
	"fmt"
	"github.com/teodorpopa/advent-of-code-go/utils"
	"strings"
)

//go:embed input.txt
var input string

func calcNextItem(diffs [][]int) int {

	subTotal := 0
	for _, diff := range diffs {
		subTotal += diff[len(diff)-1]
	}
	return subTotal
}

func calcFirstItem(diffs [][]int) int {
	nValues := make(map[int][]int)
	for i, diff := range diffs {
		nValues[i] = diff
	}

	newValues := nValues[len(nValues)]
	newValues = append([]int{0}, newValues...)
	nValues[len(nValues)] = newValues

	for i := len(nValues) - 1; i >= 0; i-- {
		nVal := []int{}
		newFirstVal := 0

		if i == 0 {
			continue
		} else {
			newFirstVal = nValues[i-1][0] - nValues[i][0]
			nVal = append([]int{newFirstVal}, nValues[i-1]...)
			nValues[i-1] = nVal
		}
	}

	return nValues[0][0]
}

func part1(input string) int {
	lines := utils.ReadLines(input)
	sum := 0

	for _, line := range lines {
		values := strings.Split(line, " ")
		valuesInt := []int{}
		for _, value := range values {
			valuesInt = append(valuesInt, utils.ToInt(value))
		}

		diffs := [][]int{
			valuesInt,
		}
		prev := 0
		allZero := false

		for !allZero {
			newVal := []int{}
			for i := 1; i < len(valuesInt); i++ {

				prev = valuesInt[i-1]

				diff := valuesInt[i] - prev
				newVal = append(newVal, diff)

			}

			diffs = append(diffs, newVal)
			valuesInt = newVal

			for _, nv := range newVal {
				if nv != 0 {
					allZero = false
					break
				}
				allZero = true
			}
		}

		nextItem := calcNextItem(diffs)
		sum += nextItem

	}

	return sum
}

func part2(input string) int {
	lines := utils.ReadLines(input)
	sum := 0

	for _, line := range lines {
		values := strings.Split(line, " ")
		valuesInt := []int{}
		for _, value := range values {
			valuesInt = append(valuesInt, utils.ToInt(value))
		}

		diffs := [][]int{
			valuesInt,
		}
		prev := 0
		allZero := false

		for !allZero {
			newVal := []int{}
			for i := 1; i < len(valuesInt); i++ {

				prev = valuesInt[i-1]

				diff := valuesInt[i] - prev
				newVal = append(newVal, diff)

			}

			diffs = append(diffs, newVal)
			valuesInt = newVal

			for _, nv := range newVal {
				if nv != 0 {
					allZero = false
					break
				}
				allZero = true
			}
		}

		nextItem := calcFirstItem(diffs)
		sum += nextItem

	}

	return sum
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
