package main

import (
	_ "embed"
	"flag"
	"fmt"

	"github.com/teodorpopa/advent-of-code-go/utils"
)

//go:embed input.txt
var input string

const START = 50
const MIN = 0
const MAX = 99

func part1(input string) int {
	lines := utils.ReadLines(input)

	var dial = START
	var count = 0

	for _, line := range lines {
		direction := string(line[0])
		amount := utils.ToInt(line[1:len(line)])

		clicks := amount % (MAX + 1)
		if direction == "L" {
			for i := 0; i < clicks; i++ {
				dial--

				if dial < MIN {
					dial = MAX
				}

			}
		} else {
			for i := 0; i < clicks; i++ {
				dial++

				if dial == MAX+1 {
					dial = MIN
				}
			}
		}

		if dial == MIN {
			count++
		}
	}

	return count
}

func part2(input string) int {
	lines := utils.ReadLines(input)

	var dial = START
	var count = 0

	for _, line := range lines {
		direction := string(line[0])
		amount := utils.ToInt(line[1:len(line)])

		for range amount {
			if direction == "L" {
				dial++
			} else if direction == "R" {
				dial--
			}

			if dial%100 == 0 {
				count++
			}
		}
	}

	return count
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
