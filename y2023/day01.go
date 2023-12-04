package y2023

import (
	"fmt"
	"github.com/teodorpopa/advent-of-code-go/utils"
	"strconv"
	"strings"
)

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

func isNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func Day01First(input string) int {
	lines := utils.ReadLines(input)
	total := 0

	for _, line := range lines {
		total += getFirstAndLastDigitsAsNumber(&line)
	}
	return total
}

func Day01Second(input string) int {
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

func Day01() {
	fmt.Println(utils.DAY_PREFIX, "Day 01")
	fmt.Println(utils.PART_PREFIX, "Part 1:", Day01First(utils.ReadFile("y2023/input/day01.txt")))
	fmt.Println(utils.PART_PREFIX, "Part 2:", Day01Second(utils.ReadFile("y2023/input/day01.txt")))
}
