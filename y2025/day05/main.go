package main

import (
	_ "embed"
	"flag"
	"fmt"
	"sort"
	"strings"

	"github.com/teodorpopa/advent-of-code-go/utils"
)

//go:embed input.txt
var input string

type Interval struct {
	start int
	end   int
}

func gatherIntervals(input string) []Interval {
	intervals := make([]Interval, 0)
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, "-")
		start, end := utils.ToInt(parts[0]), utils.ToInt(parts[1])
		intervals = append(intervals, Interval{start, end})
	}
	return intervals
}

func isFresh(intervals []Interval, ingredient int) bool {
	for _, interval := range intervals {
		if ingredient >= interval.start && ingredient <= interval.end {
			return true
		}
	}
	return false
}

func part1(input string) int {
	parts := strings.Split(input, "\n\n")
	intervals := gatherIntervals(parts[0])
	lines := utils.ReadLines(parts[1])
	fresh := 0

	for _, line := range lines {
		if isFresh(intervals, utils.ToInt(line)) {
			fresh++
		}
	}

	return fresh
}

func part2(input string) int {
	parts := strings.Split(input, "\n\n")
	intervals := gatherIntervals(parts[0])
	fresh := 0

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].start < intervals[j].start
	})

	start := intervals[0].start
	end := intervals[0].end

	for _, interval := range intervals[1:] {
		newStart := interval.start
		newEnd := interval.end

		if newStart > end {
			fresh += (end - start) + 1

			start = newStart
			end = newEnd
		} else if newEnd > end {
			end = newEnd
		}
	}

	fresh += (end - start) + 1

	return fresh
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
