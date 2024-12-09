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

const (
	Increasing = iota
	Decreasing
)

func Reports(input string) [][]int {
	lines := utils.ReadLines(input)

	reports := [][]int{}
	for _, l := range lines {
		parts := strings.Split(l, " ")

		report := []int{}
		for _, p := range parts {
			report = append(report, utils.ToInt(p))
		}
		reports = append(reports, report)
	}

	return reports
}

func compareReport(r []int) bool {
	safe := true

	levelsType := Decreasing
	if (r[0] - r[1]) < 0 {
		levelsType = Increasing
	}

	for i := 0; i < len(r)-1; i++ {
		diff := r[i] - r[i+1]

		if (levelsType == Increasing && (diff >= 0 || math.Abs(float64(diff)) > 3)) ||
			(levelsType == Decreasing && (diff <= 0 || diff > 3)) {
			safe = false
			break
		}
	}

	return safe
}

func compareWithDampener(r []int) bool {
	for i := range r {
		reportClone := utils.CopyIntSlice(r)
		rClone := append(reportClone[:i], reportClone[i+1:]...)
		ok := compareReport(rClone)
		if ok {
			return true
		}
	}

	return false
}

func part1(input string) int {
	safeReports := 0
	for _, r := range Reports(input) {
		safe := compareReport(r)

		if safe {
			safeReports++
			continue
		}
	}

	return safeReports
}

func part2(input string) int {
	safeReports := 0
	for _, r := range Reports(input) {
		safe := compareReport(r)
		if safe {
			safeReports++
			continue
		}

		safe = compareWithDampener(r)
		if safe {
			safeReports++
			continue
		}
	}

	return safeReports
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
