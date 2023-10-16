package y2015

import (
	"fmt"
	"github.com/teodorpopa/advent-of-code-go/utils"
	"sort"
	"strings"
)

func Day02First(input string) int {
	lines := utils.ReadLines(input)
	totalSqFt := 0

	for _, v := range lines {
		dim := strings.Split(v, "x")
		l := utils.ToInt(dim[0])
		w := utils.ToInt(dim[1])
		h := utils.ToInt(dim[2])

		side1 := l * w
		side2 := w * h
		side3 := h * l

		slack := side1
		comp := []int{side1, side2, side3}
		for _, value := range comp {
			if value < slack {
				slack = value
			}
		}

		totalSqFt += (2 * side1) + (2 * side2) + (2 * side3) + slack
	}

	return totalSqFt
}

func Day02Second(input string) int {
	lines := utils.ReadLines(input)
	totalRibbon := 0

	for _, v := range lines {
		dim := strings.Split(v, "x")
		l := utils.ToInt(dim[0])
		w := utils.ToInt(dim[1])
		h := utils.ToInt(dim[2])

		volume := l * w * h

		comp := []int{l, w, h}
		sort.Ints(comp)

		ss1 := comp[0]
		ss2 := comp[1]

		totalRibbon += 2*ss1 + 2*ss2 + volume
	}

	return totalRibbon
}

func Day02() {
	fmt.Println(utils.DAY_PREFIX, "Day 2 - I Was Told There Would Be No Math")
	fmt.Println(utils.PART_PREFIX, "Part 1:", Day02First(utils.ReadFile("y2015/input/day02.txt")))
	fmt.Println(utils.PART_PREFIX, "Part 2:", Day02Second(utils.ReadFile("y2015/input/day02.txt")))
}
