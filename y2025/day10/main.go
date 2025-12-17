package main

import (
	_ "embed"
	"flag"
	"fmt"
	"math"
	"strings"

	"github.com/teodorpopa/advent-of-code-go/utils"
)

//go:embed input.txt
var input string

const NoSolution = math.MaxInt

func parseLine(
	textline string,
	display *int,
	buttons *[]int,
	other *[]int,
) {
	r := strings.NewReader(textline)

	read := func() (rune, bool) {
		ch, _, err := r.ReadRune()
		return ch, err == nil
	}

	readNonSpace := func() (rune, bool) {
		for {
			ch, ok := read()
			if !ok {
				return 0, false
			}
			if ch != ' ' && ch != '\n' && ch != '\t' {
				return ch, true
			}
		}
	}

	*display = 0
	displaySize := 0

	read()

	for {
		ch, ok := read()
		if !ok {
			break
		}

		if ch == '.' {
			*display <<= 1
		} else if ch == '#' {
			*display = (*display << 1) | 1
		} else {
			break
		}
		displaySize++
	}

	for {
		ch, ok := readNonSpace()
		if !ok || ch != '(' {
			break
		}

		button := 0
		for {
			temp := 0
			fmt.Fscan(r, &temp)
			ch, _ = readNonSpace()
			button |= (1 << (displaySize - temp - 1))
			if ch == ')' {
				break
			}
		}
		*buttons = append(*buttons, button)
	}

	for {
		temp := 0
		_, err := fmt.Fscan(r, &temp)
		if err != nil {
			break
		}
		ch, _ := readNonSpace()
		*other = append(*other, temp)
		if ch == '}' {
			break
		}
	}
}

func permutation(
	display int,
	buttons []int,
	buttonIdx int,
	depth int,
	minDepth *int,
	stack []int,
) {
	if depth > *minDepth {
		return
	}

	if buttonIdx != -1 {
		display ^= buttons[buttonIdx]

		if display == 0 {
			if depth < *minDepth {
				*minDepth = depth
			}
			return
		}
	}

	for i := buttonIdx + 1; i < len(buttons); i++ {
		stack = append(stack, i)
		permutation(display, buttons, i, depth+1, minDepth, stack)
		stack = stack[:len(stack)-1]
	}
}

func processLine(textline string, sum *int) {
	if strings.TrimSpace(textline) == "" {
		return
	}

	var display int
	var buttons []int
	var other []int

	parseLine(textline, &display, &buttons, &other)

	if display == 0 {
		*sum += 0
		return
	}

	minDepth := NoSolution

	permutation(display, buttons, -1, 0, &minDepth, []int{})

	if minDepth == NoSolution {
		return
	}

	*sum += minDepth
}

func part1(input string) int {

	var sum int = 0
	for _, line := range utils.ReadLines(input) {
		processLine(line, &sum)
	}
	return sum
}

func part2(input string) int {
	sum := 0

	return sum
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
