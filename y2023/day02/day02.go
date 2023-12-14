package day02

import (
	_ "embed"
	"fmt"
	"github.com/teodorpopa/advent-of-code-go/utils"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

var MaxBlue = 14
var MaxGreen = 13
var MaxRed = 12

func part1(input string) int {
	lines := utils.ReadLines(input)
	total := 0

	for _, line := range lines {

		parts := strings.Split(line, ": ")
		gameId := strings.ReplaceAll(parts[0], "Game ", "")

		games := strings.Split(parts[1], "; ")

		validLine := true

		for _, g := range games {
			opts := strings.Split(g, ", ")
			validGame := true

			for _, o := range opts {
				oBlue := 0
				oGreen := 0
				oRed := 0

				cubes := strings.Split(o, " ")

				cubeNo, _ := strconv.Atoi(cubes[0])
				switch cubes[1] {
				case "blue":
					oBlue = cubeNo
					break
				case "green":
					oGreen = cubeNo
					break
				case "red":
					oRed = cubeNo
					break
				}

				if oBlue > MaxBlue || oRed > MaxRed || oGreen > MaxGreen {
					validGame = false
					break
				}

			}

			if validGame == false {
				validLine = false
				break

			}
		}

		if validLine {
			gameIdI, _ := strconv.Atoi(gameId)
			total += gameIdI
		}

	}

	return total
}

func part2(input string) int {
	lines := utils.ReadLines(input)
	total := 0

	for _, line := range lines {

		parts := strings.Split(line, ": ")
		games := strings.Split(parts[1], "; ")

		minBlue := 1
		minGreen := 1
		minRed := 1

		for _, g := range games {
			opts := strings.Split(g, ", ")

			for _, o := range opts {
				oBlue := 0
				oGreen := 0
				oRed := 0

				cubes := strings.Split(o, " ")

				cubeNo, _ := strconv.Atoi(cubes[0])
				switch cubes[1] {
				case "blue":
					oBlue = cubeNo
					if minBlue <= oBlue {
						minBlue = oBlue
					}
					break
				case "green":
					oGreen = cubeNo
					if minGreen <= oGreen {
						minGreen = oGreen
					}
					break
				case "red":
					oRed = cubeNo
					if minRed <= oRed {
						minRed = oRed
					}
					break
				}

			}

		}
		gameScore := minBlue * minGreen * minRed
		total += gameScore
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
