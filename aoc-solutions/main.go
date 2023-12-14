package main

import (
	"flag"
	"github.com/gookit/color"
	"github.com/teodorpopa/advent-of-code-go/runner"
	"github.com/teodorpopa/advent-of-code-go/utils"
)

const title = "   GO Advent of Code   "

var year int
var day int
var part int
var create bool

func main() {

	utils.DisplayTitle(title)

	flag.IntVar(&year, "year", -1, "Year to run the challenge for")
	flag.IntVar(&year, "y", -1, "Year to run the challenge for")

	flag.IntVar(&day, "day", -1, "Day number for the specified year")
	flag.IntVar(&day, "d", -1, "Day number for the specified year")

	flag.IntVar(&part, "part", 0, "0 For both, 1 for first part, 2 for second part")
	flag.IntVar(&part, "p", 0, "0 For both, 1 for first part, 2 for second part")

	flag.BoolVar(&create, "create", false, "create a day from template")

	flag.Parse()

	if !validate(year, day, part) {
		color.Error.Println("Invalid flags")
		return
	}

	partsToRun := []int{1, 2}
	if part != 0 {
		partsToRun = []int{part}
	}

	runner.Run(year, day, partsToRun)
}

func validate(y int, d int, p int) bool {
	if y == -1 || d == -1 {
		return false
	}
	if y < 2015 || y > 2023 {
		return false
	}
	if d < 1 || d > 25 {
		return false
	}
	if p < 0 || p > 2 {
		return false
	}

	return true
}
