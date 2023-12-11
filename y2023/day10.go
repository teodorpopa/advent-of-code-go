package y2023

import (
	"fmt"
	"github.com/gookit/color"
	"github.com/teodorpopa/advent-of-code-go/utils"
	"golang.org/x/exp/slices"
	"strings"
)

var charMap = map[string]string{
	"|": "║",
	"-": "═",
	"L": "╚",
	"J": "╝",
	"7": "╗",
	"F": "╔",
	".": " ",
	"S": "▪",
	"X": "█",
}

type mazePoint struct {
	X     int
	Y     int
	Value string
}

var goUp = []string{"-", "J", "7"}
var goRight = []string{"|", "J", "L"}
var goDown = []string{"F", "-", "L"}
var goLeft = []string{"|", "L", "7"}

func displayMainMaze(input string, mainMazeCoords []mazePoint) {

	lines := utils.ReadLines(input)
	specialChars := []string{"|", "-", "L", "J", "7", "F", "S", "X"}

	newDesign := make([][]string, len(lines))
	for i, line := range lines {
		newDesign[i] = make([]string, len(line))
		for j, char := range line {
			newDesign[i][j] = charMap[string(char)]
		}
	}

	c := color.New(color.FgCyan)
	for _, mCoord := range mainMazeCoords {
		if slices.Contains(specialChars, mCoord.Value) {
			newDesign[mCoord.X][mCoord.Y] = c.Sprintf(charMap[mCoord.Value])
			continue
		}
	}

	for _, line := range newDesign {
		fmt.Println()
		for _, char := range line {
			if char != "" {
				fmt.Print(char)
			} else {
				fmt.Print(" ")
			}
		}
	}
	fmt.Println()
	fmt.Println()
}

func displaySecondaryMaze(input [][]string) {
	specialChars := []string{"X"}

	c := color.New(color.Blue)
	for _, line := range input {
		for _, char := range line {
			if slices.Contains(specialChars, char) {
				c.Print(charMap[char])
			} else {
				fmt.Print(char)
			}
		}
		fmt.Println()
	}

}

func sPosition(input string) mazePoint {
	lines := utils.ReadLines(input)

	for i, line := range lines {
		for j, char := range line {
			if string(char) == "S" {
				return mazePoint{
					X:     i,
					Y:     j,
					Value: "S",
				}
			}
		}
	}

	panic("Start point not found")
}

func findMainMaze(input string) []mazePoint {
	lines := utils.ReadLines(input)
	pastMoves := []mazePoint{}
	s := sPosition(input)
	pastMoves = append(pastMoves, s)

	startMove := getStartMove(s.X, s.Y, lines)

	for startMove.Value != "S" {
		lastMove := pastMoves[len(pastMoves)-1]
		newMove := findNextMove(startMove, lastMove, lines)
		pastMoves = append(pastMoves, startMove)
		startMove = newMove
	}

	return pastMoves
}

func getStartMove(x, y int, input []string) mazePoint {
	newX := x
	newY := y
	found := false

	if y != (len(input[x])-1) && slices.Contains(goUp, string(input[x][y+1])) {
		newY = y + 1
		found = true
	}

	if !found && x != (len(input)-1) && slices.Contains(goRight, string(input[x+1][y])) {
		newX = x + 1
		found = true
	}

	if !found && y != 0 && slices.Contains(goDown, string(input[x][y-1])) {
		newY = y - 1
		found = true
	}

	if !found && x != 0 && slices.Contains(goLeft, string(input[x-1][y])) {
		newX = x - 1
		found = true
	}

	return mazePoint{
		X:     newX,
		Y:     newY,
		Value: string(input[newX][newY]),
	}
}

func findNextMove(current, prev mazePoint, input []string) mazePoint {
	x := current.X
	y := current.Y
	prevX := prev.X
	prevY := prev.Y

	newX := current.X
	newY := current.Y

	if current.Value == "-" {
		if y != 0 && prevY == (y-1) {
			newY = y + 1
		} else if y != (len(input[x])-1) && prevY == (y+1) {
			newY = y - 1
		}
	} else if current.Value == "J" {
		if x != 0 && y != 0 && prevX == (x-1) {
			newY = y - 1
		} else if x != 0 && y != 0 && prevY == (y-1) {
			newX = x - 1
		}
	} else if current.Value == "|" {
		if x != 0 && x != len(input)-1 && prevX == (x-1) {
			newX = x + 1
		} else if x != 0 && x != len(input)-1 && prevX == (x+1) {
			newX = x - 1
		}
	} else if current.Value == "L" {
		if x != 0 && y != len(input[x])-1 && prevX == (x-1) {
			newY = y + 1
		} else if x != 0 && y != len(input[x])-1 && prevY == (y+1) {
			newX = x - 1
		}
	} else if current.Value == "7" {
		if y != 0 && x != (len(input)-1) && prevY == (y-1) {
			newX = x + 1
		} else if x != (len(input)-1) && y != 0 && prevX == (x+1) {
			newY = y - 1
		}
	} else if current.Value == "F" {
		if x != (len(input)-1) && y != (len(input[x])-1) && prevY == (y+1) {
			newX = x + 1
		} else if x != (len(input)-1) && y != (len(input[x])-1) && prevX == (x+1) {
			newY = y + 1
		}
	}

	return mazePoint{
		X:     newX,
		Y:     newY,
		Value: string(input[newX][newY]),
	}
}

func closeLoop(start, first, last mazePoint, input []string) mazePoint {
	for _, char := range "F-|JL7" {
		newInput := input
		a := strings.Split(newInput[start.X], "")

		a[start.Y] = string(char)
		newInput[start.X] = strings.Join(a, "")

		start.Value = string(char)

		mP := findNextMove(start, last, newInput)
		if mP.X == first.X && mP.Y == first.Y {
			return start
		}
	}

	return start
}

func izolateMaze(input string, mazePoints []mazePoint) [][]string {
	lines := utils.ReadLines(input)
	newDesign := make([][]string, len(lines))
	for i, line := range lines {
		newDesign[i] = make([]string, len(line))
		for j, _ := range line {
			newDesign[i][j] = " "
		}
	}

	for _, mCoord := range mazePoints {
		if charMap[mCoord.Value] != "" {
			newDesign[mCoord.X][mCoord.Y] = charMap[mCoord.Value]
			continue
		}
	}

	for i, line := range newDesign {
		for j, char := range line {
			if char != "" {
				newDesign[i][j] = char
			} else {
				newDesign[i][j] = " "
			}
		}
	}

	return newDesign
}

func mapMainPositions(mazePoints []mazePoint) map[int][]int {
	mapPosition := make(map[int][]int)
	for _, p := range mazePoints {
		mapPosition[p.X] = append(mapPosition[p.X], p.Y)
	}

	return mapPosition
}

func findElementsInside(mainMaze []mazePoint, input string) int {
	lines := utils.ReadLines(input)

	totalElements := 0

	for k, _ := range lines {
		points := strings.Split(lines[k], "")

		for i, _ := range points {
			if points[i] == "S" {
				replacement := closeLoop(
					mainMaze[0],
					mainMaze[1],
					mainMaze[len(mainMaze)-1],
					lines,
				)
				mainMaze = append(
					[]mazePoint{replacement},
					mainMaze[1:len(mainMaze)]...,
				)
				points[i] = charMap[replacement.Value]
			} else {
				points[i] = charMap[points[i]]
			}
		}
		lines[k] = strings.Join(points, "")
	}

	newDesign := izolateMaze(input, mainMaze)
	mapPositions := mapMainPositions(mainMaze)

	for k, _ := range mapPositions {
		chars := newDesign[k]

		insideBlocks := 0
		lastChar := "#"

		rowCount := 0
		for i, c := range chars {
			if c == "║" {
				insideBlocks++
				continue
			} else if c == "═" {
				continue
			}

			if lastChar != "#" && slices.Contains([]string{"╝", "╚", "╗", "╔"}, c) {
				if (lastChar == "╚" && c == "╗") || (lastChar == "╔" && c == "╝") {
					insideBlocks++
				}
				lastChar = "#"
				continue
			} else if lastChar == "#" && slices.Contains([]string{"╝", "╚", "╗", "╔"}, c) {
				if slices.Contains([]string{"╝", "╚", "╗", "╔"}, c) {
					lastChar = c
					continue
				}
			}

			if insideBlocks%2 == 0 {
				chars[i] = " "
			} else if insideBlocks%2 != 0 {
				chars[i] = "X"
				rowCount++
			}
		}
		newDesign[k] = chars
		totalElements += rowCount
	}

	displaySecondaryMaze(newDesign)

	return totalElements
}

func Day10First(input string) int {
	mainMaze := findMainMaze(input)
	displayMainMaze(input, mainMaze)

	return len(mainMaze) / 2
}

func Day10Second(input string) int {
	mainMaze := findMainMaze(input)
	elementsInside := findElementsInside(mainMaze, input)

	return elementsInside
}

func Day10() {
	fmt.Println(utils.DAY_PREFIX, "Day 10")
	fmt.Println(utils.PART_PREFIX, "Part 1:", Day10First(utils.ReadFile("y2023/input/day10.txt")))
	fmt.Println(utils.PART_PREFIX, "Part 2:", Day10Second(utils.ReadFile("y2023/input/day10.txt")))
}
