package main

import (
	_ "embed"
	"flag"
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
