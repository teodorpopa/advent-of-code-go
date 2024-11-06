package main

import (
	_ "embed"
	"flag"
	"fmt"
	"github.com/teodorpopa/advent-of-code-go/utils"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

var nums = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func getFirstAndLastDigitsAsNumber(s *string) int {
	firstDigit := 0
	lastDigit := 0

	chars := strings.Split(*s, "")
	for i, _ := range chars {
		intValue, err := strconv.Atoi(chars[i])
		if err == nil && firstDigit == 0 {
			firstDigit = intValue
		}
	}

	revLine := utils.ReverseString(*s)
	revChars := strings.Split(revLine, "")
	for j, _ := range revChars {
		intValue, err := strconv.Atoi(revChars[j])
		if err == nil && lastDigit == 0 {
			lastDigit = intValue
		}
	}

	add, err := strconv.Atoi(fmt.Sprintf("%d%d", firstDigit, lastDigit))
	if err != nil {
		panic(err)
	}

	return add
}

func part1(input string) int {
	lines := utils.ReadLines(input)
	total := 0

	for _, line := range lines {
		total += getFirstAndLastDigitsAsNumber(&line)
	}
	return total
}

func part2(input string) int {
	lines := utils.ReadLines(input)
	total := 0

	for _, line := range lines {
		oldLine := line
		newLine := ""

		replaceFirst := false
		replaceLast := false

		ch := strings.Split(oldLine, "")
		for i := 0; i < len(ch); i++ {
			strToComp := oldLine[i:len(ch)]

			_, err := strconv.Atoi(string(ch[i]))
			if err == nil && replaceFirst == false {
				newLine = line
				break
			}

			if !replaceFirst {
				for j, n := range nums {
					if strings.HasPrefix(strToComp, n) {
						replaceFirst = true
						newLine = oldLine[0:i] + strconv.Itoa(j+1) + oldLine[i+len(n):len(ch)]
						i += len(nums[j]) - 1
						break
					}
				}
			} else {
				break
			}
		}

		ch = strings.Split(newLine, "")

		for i := 0; i < len(ch); i++ {

			if !replaceLast {
				strToComp := newLine[len(ch)-i : len(ch)]

				for j, n := range nums {
					if strings.HasPrefix(strToComp, n) {
						replaceLast = true
						newLine = newLine[0:len(newLine)-len(strToComp)] + strconv.Itoa(j+1) + newLine[len(newLine)-len(strToComp)+len(n):len(newLine)]
						break
					}
				}
			} else {
				break
			}
		}

		if newLine == "" {
			newLine = line
		}

		add := getFirstAndLastDigitsAsNumber(&newLine)
		total += add

	}
	return total
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
