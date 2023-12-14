package day05

import (
	_ "embed"
	"fmt"
	"github.com/teodorpopa/advent-of-code-go/utils"
	"golang.org/x/exp/slices"
	"regexp"
	"strings"
)

//go:embed input.txt
var input string

var vowels = []string{"a", "e", "i", "o", "u"}

func containsVowels(s string) bool {
	c := 0

	for _, l := range strings.Split(s, "") {
		if slices.Contains(vowels, l) {
			c++
		}
	}

	return c >= 3
}

func containsDoubleLetters(s string) bool {
	prevC := ""

	for _, l := range strings.Split(s, "") {
		c := l

		if c == prevC {
			return true
		} else {
			prevC = c
		}
	}

	return false
}

func containsNaughty(s string) bool {
	match, _ := regexp.MatchString("(ab|cd|pq|xy)", s)
	return match
}

func hasTwicePair(s string) bool {
	match, _ := regexp.MatchString("(..).*\\1", s)
	return match
}

func repeatedLetter(s string) bool {
	match, _ := regexp.MatchString("(.).\\1", s)
	return match
}

func part1(input string) int {
	lines := utils.ReadLines(input)

	i := 0
	for _, l := range lines {
		if containsVowels(l) && containsDoubleLetters(l) && !containsNaughty(l) {
			i++
		}
	}

	return i
}

func part2(input string) int {
	lines := utils.ReadLines(input)

	i := 0
	for _, l := range lines {
		if hasTwicePair(l) && repeatedLetter(l) {
			i++
		}
	}

	return i
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
