package day15

import (
	_ "embed"
	"fmt"
	"github.com/teodorpopa/advent-of-code-go/utils"
	"golang.org/x/exp/slices"
	"strings"
)

//go:embed input.txt
var input string

type Lens struct {
	Name  string
	Value int
}

func hash(input string) int {
	start := 0

	for i := 0; i < len(input); i++ {
		curVal := (start + int(input[i])) * 17
		curVal = curVal % 256
		start = curVal

		if i == len(input)-1 {
			return curVal
		}
	}

	return 0
}

func part1(input string) int {
	sum := 0
	parts := strings.Split(input, ",")

	for _, part := range parts {
		sum += hash(part)
	}

	return sum
}

func part2(input string) int {
	sum := 0
	parts := strings.Split(input, ",")
	boxes := [256][]Lens{}

	for _, part := range parts {
		action := ""
		lp := []string{}

		if strings.Contains(part, "=") {
			action = "add"
			lp = strings.Split(part, "=")
		} else {
			action = "remove"
			lp = strings.Split(part, "-")
		}

		curBox := &boxes[hash(lp[0])]

		lensIndex := slices.IndexFunc(*curBox, func(l Lens) bool {
			return l.Name == lp[0]
		})

		if action == "remove" {
			if lensIndex > -1 {
				*curBox = slices.Delete(*curBox, lensIndex, lensIndex+1)
			}
		} else {
			lens := Lens{
				lp[0],
				utils.ToInt(lp[1]),
			}

			if lensIndex != -1 {
				(*curBox)[lensIndex] = lens
			} else {
				*curBox = append(*curBox, lens)
			}
		}
	}

	for i, box := range boxes {
		for j, lens := range box {
			sum += (i + 1) * (j + 1) * lens.Value
		}
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
