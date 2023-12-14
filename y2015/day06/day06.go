package day06

import (
	_ "embed"
	"fmt"
	"github.com/teodorpopa/advent-of-code-go/utils"
	"regexp"
)

//go:embed input.txt
var input string

type Action struct {
	action string
	fromX  int
	fromY  int
	toX    int
	toY    int
}

func createGrid(x int, y int) [][]int {
	grid := [][]int{}

	for i := 0; i <= x; i++ {
		grid = append(grid, make([]int, y+1))
	}

	return grid
}

func countLights(grid [][]int) int {
	lights := 0

	for _, v := range grid {
		for _, y := range v {
			lights = lights + y
		}
	}

	return lights
}

func transAction(s string) Action {
	r, _ := regexp.Compile("^(turn on|turn off|toggle)\\s(.*?),(.*?)\\s\\w+\\s(.*?),(.*?)$")
	parts := r.FindStringSubmatch(s)

	return Action{
		action: parts[1],
		fromX:  utils.ToInt(parts[2]),
		fromY:  utils.ToInt(parts[3]),
		toX:    utils.ToInt(parts[4]),
		toY:    utils.ToInt(parts[5]),
	}
}

func lights(action Action, grid *[][]int) {
	for i, v := range *grid {
		for j, w := range v {

			if i >= action.fromX && i <= action.toX && j >= action.fromY && j <= action.toY {

				curV := w
				switch action.action {
				case "turn on":
					curV = 1
					break
				case "turn off":
					curV = 0
					break
				case "toggle":
					if curV == 1 {
						curV = 0
					} else {
						curV = 1
					}
				}
				(*grid)[i][j] = curV
			}
		}
	}
}

func lightsBright(action Action, grid *[][]int) {
	for i, v := range *grid {
		for j, w := range v {

			if i >= action.fromX && i <= action.toX && j >= action.fromY && j <= action.toY {

				curV := w
				switch action.action {
				case "turn on":
					curV++
					break
				case "turn off":
					curV--
					if curV < 0 {
						curV = 0
					}
					break
				case "toggle":
					curV = curV + 2
				}
				(*grid)[i][j] = curV
			}
		}
	}
}

func part1(input string) int {
	grid := createGrid(999, 999)
	lines := utils.ReadLines(input)

	for _, line := range lines {
		action := transAction(line)
		lights(action, &grid)
	}

	return countLights(grid)
}

func part2(input string) int {
	grid := createGrid(999, 999)
	lines := utils.ReadLines(input)

	for _, line := range lines {
		action := transAction(line)
		lightsBright(action, &grid)
	}

	return countLights(grid)
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
