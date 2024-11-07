package main

import (
	_ "embed"
	"flag"
	"fmt"
	"regexp"
	"slices"
	"strings"

	"github.com/teodorpopa/advent-of-code-go/utils"
)

//go:embed input.txt
var input string

var toMatch = map[string]int{
	"children":    3,
	"cats":        7,
	"samoyeds":    2,
	"pomeranians": 3,
	"akitas":      0,
	"vizslas":     0,
	"goldfish":    5,
	"trees":       3,
	"cars":        2,
	"perfumes":    1,
}

func part1(input string) int {
	lines := utils.ReadLines(input)
	var id int
	for _, line := range lines {
		r, _ := regexp.Compile(`Sue (\d+): (.*)`)
		founds := r.FindStringSubmatch(line)
		var notGood bool

		id = utils.ToInt(founds[1])
		for _, s := range strings.Split(founds[2], ", ") {
			r2, _ := regexp.Compile(`(\w+): (\d+)`)
			stats := r2.FindStringSubmatch(s)

			item := stats[1]
			count := utils.ToInt(stats[2])

			if count != toMatch[item] {
				notGood = true
			}
		}

		if !notGood {
			return id
		}
	}
	return 0
}

func part2(input string) int {
	lines := utils.ReadLines(input)

	more := []string{"cats", "trees"}
	less := []string{"pomeranians", "goldfish"}

	var id int
	for _, line := range lines {
		r, _ := regexp.Compile(`Sue (\d+): (.*)`)
		founds := r.FindStringSubmatch(line)
		var notGood bool

		id = utils.ToInt(founds[1])
		for _, s := range strings.Split(founds[2], ", ") {
			r2, _ := regexp.Compile(`(\w+): (\d+)`)
			stats := r2.FindStringSubmatch(s)

			item := stats[1]
			count := utils.ToInt(stats[2])

			if slices.Contains(more, item) {
				if count <= toMatch[item] {
					notGood = true
				}
			} else if slices.Contains(less, item) {
				if count >= toMatch[item] {
					notGood = true
				}
			} else {
				if count != toMatch[item] {
					notGood = true
				}
			}
		}

		if !notGood {
			return id
		}
	}
	return 0
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
