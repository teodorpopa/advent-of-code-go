package main

import (
	_ "embed"
	"flag"
	"fmt"
	"github.com/teodorpopa/advent-of-code-go/utils"
	"sort"
	"strings"
)

//go:embed input.txt
var input string

func part1(input string) int {
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

func part2(input string) int {
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
