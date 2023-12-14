package day04

import (
	_ "embed"
	"fmt"
	"github.com/teodorpopa/advent-of-code-go/utils"
	"math"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type Card struct {
	Name             string
	WinningNumbers   []int
	ScratchedNumbers []int
}

func part1(input string) int {
	lines := utils.ReadLines(input)
	total := 0
	cards := []Card{}

	for _, line := range lines {
		parts := strings.Split(line, " | ")

		cardInfo := strings.Split(parts[0], ": ")
		cardName := cardInfo[0]

		winningNo := strings.Split(cardInfo[1], " ")
		winningNoI := make([]int, len(winningNo))

		for i, s := range winningNo {
			winningNoI[i], _ = strconv.Atoi(s)
		}

		scratchedNo := strings.Split(parts[1], " ")
		scratchedNoI := make([]int, len(scratchedNo))

		for i, s := range scratchedNo {
			scratchedNoI[i], _ = strconv.Atoi(s)
		}

		card := Card{
			cardName,
			winningNoI,
			scratchedNoI,
		}
		cards = append(cards, card)
	}

	for _, c := range cards {
		matchedNo := 0

		for _, sNum := range c.ScratchedNumbers {
			for _, wNum := range c.WinningNumbers {
				if sNum > 0 && wNum > 0 && sNum == wNum {
					matchedNo++
				}
			}
		}

		if matchedNo > 1 {
			total += int(math.Pow(2.0, float64(matchedNo-1)))
		} else if matchedNo == 1 {
			total += 1
		}
	}

	return total
}

func part2(input string) int {
	lines := utils.ReadLines(input)
	total := 0
	cards := []Card{}

	for _, line := range lines {
		parts := strings.Split(line, " | ")

		cardInfo := strings.Split(parts[0], ": ")
		cardName := cardInfo[0]

		winningNo := strings.Split(cardInfo[1], " ")
		winningNoI := make([]int, len(winningNo))

		for i, s := range winningNo {
			winningNoI[i], _ = strconv.Atoi(s)
		}

		scratchedNo := strings.Split(parts[1], " ")
		scratchedNoI := make([]int, len(scratchedNo))

		for i, s := range scratchedNo {
			scratchedNoI[i], _ = strconv.Atoi(s)
		}

		card := Card{
			cardName,
			winningNoI,
			scratchedNoI,
		}
		cards = append(cards, card)
	}

	total += getTotalCards(cards)
	return total
}

func getTotalCards(cards []Card) int {
	total := 0

	for i, c := range cards {
		matchedNo := 0

		for _, sNum := range c.ScratchedNumbers {
			for _, wNum := range c.WinningNumbers {
				if sNum > 0 && wNum > 0 && sNum == wNum {
					matchedNo++
				}
			}
		}

		if matchedNo > 0 {
			total += getTotalCards(cards[i+1 : i+1+matchedNo])
		}
		total++
	}

	return total
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
