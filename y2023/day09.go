package y2023

import (
	"fmt"
	"github.com/teodorpopa/advent-of-code-go/utils"
	"strings"
)

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

func Day09First(input string) int {
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

func Day09Second(input string) int {
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

func Day09() {
	fmt.Println(utils.DAY_PREFIX, "Day 09")
	fmt.Println(utils.PART_PREFIX, "Part 1:", Day09First(utils.ReadFile("y2023/input/day09.txt")))
	fmt.Println(utils.PART_PREFIX, "Part 2:", Day09Second(utils.ReadFile("y2023/input/day09.txt")))
}
