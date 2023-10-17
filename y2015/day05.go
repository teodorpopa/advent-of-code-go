package y2015

import (
	"fmt"
	"github.com/teodorpopa/advent-of-code-go/utils"
	"golang.org/x/exp/slices"
	"regexp"
	"strings"
)

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

func Day05First(input string) int {
	lines := utils.ReadLines(input)

	i := 0
	for _, l := range lines {
		if containsVowels(l) && containsDoubleLetters(l) && !containsNaughty(l) {
			i++
		}
	}

	return i
}

func Day05Second(input string) int {
	lines := utils.ReadLines(input)

	i := 0
	for _, l := range lines {
		if hasTwicePair(l) && repeatedLetter(l) {
			i++
		}
	}

	return i
}

func Day05() {
	fmt.Println(utils.DAY_PREFIX, "Day 05")
	fmt.Println(utils.PART_PREFIX, "Part 1:", Day05First(utils.ReadFile("y2015/input/day05.txt")))
	fmt.Println(utils.PART_PREFIX, "Part 2:", Day05Second(utils.ReadFile("y2015/input/day05.txt")))
}
