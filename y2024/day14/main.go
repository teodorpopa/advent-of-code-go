package main

import (
	_ "embed"
	"flag"
	"fmt"
	"github.com/teodorpopa/advent-of-code-go/utils"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var input string

type robot struct {
	coords point
	velX   int
	velY   int
}

type point struct {
	x int
	y int
}

func parseInput(input string) []robot {
	robots := []robot{}
	lines := utils.ReadLines(input)

	for _, line := range lines {
		parts := strings.Split(line, " ")

		p0 := strings.Split(parts[0], "=")
		p1 := strings.Split(parts[1], "=")

		coords := strings.Split(p0[1], ",")
		vel := strings.Split(p1[1], ",")

		robots = append(robots, robot{
			coords: point{
				x: utils.ToInt(coords[0]),
				y: utils.ToInt(coords[1]),
			},
			velX: utils.ToInt(vel[0]),
			velY: utils.ToInt(vel[1]),
		})
	}

	return robots
}

func createGrid(width, height int) [][]string {
	grid := make([][]string, height)
	val := "."

	for i := 0; i < height; i++ {
		grid[i] = make([]string, width)
		for j := 0; j < width; j++ {
			grid[i][j] = val
		}
	}

	return grid
}

func moveRobots(robots *[]robot, grid [][]string) {
	for i := 0; i < len(*robots); i++ {
		robot := (*robots)[i]
		curX := robot.coords.x
		curY := robot.coords.y

		newX := curX + robot.velX
		newY := curY + robot.velY

		if !utils.ValidIndex(newX, newY, len(grid[0]), len(grid)) {
			if newX < 0 {
				newX = len(grid[0]) + newX
			} else if newX >= len(grid[0]) {
				newX = newX - len(grid[0])
			}

			if newY < 0 {
				newY = len(grid) + newY
			} else if newY >= len(grid) {
				newY = newY - len(grid)
			}
		}

		robot.coords.x = newX
		robot.coords.y = newY
		(*robots)[i] = robot
	}
}

func displayGrid(grid [][]string, robots *[]robot) {

	grid = createGrid(len(grid[0]), len(grid))

	for _, robot := range *robots {
		if grid[robot.coords.y][robot.coords.x] == "." {
			grid[robot.coords.y][robot.coords.x] = "1"
		} else {
			curValue := utils.ToInt(grid[robot.coords.y][robot.coords.x])
			curValue++
			grid[robot.coords.y][robot.coords.x] = strconv.Itoa(curValue)
		}
	}

	for i := 0; i < len(grid); i++ {
		fmt.Println(grid[i])
	}
}

func part1(input string, width int, height int) int {
	robots := parseInput(input)
	grid := createGrid(width, height)

	seconds := 100

	for i := 0; i < seconds; i++ {
		fmt.Print("\033[H\033[2J")
		fmt.Println("\nSecond: ", i)

		moveRobots(&robots, grid)
		displayGrid(grid, &robots)

		time.Sleep(90 * time.Millisecond)
	}

	midX := (len(grid[0]) - 1) / 2
	midY := (len(grid) - 1) / 2

	robotCount := make(map[string]int)
	robotCount["NW"] = 0
	robotCount["NE"] = 0
	robotCount["SE"] = 0
	robotCount["SW"] = 0

	for _, robot := range robots {
		if robot.coords.x <= midX-1 && robot.coords.y >= midY+1 {
			robotCount["NW"]++
		} else if robot.coords.x >= midX+1 && robot.coords.y >= midY+1 {
			robotCount["NE"]++
		} else if robot.coords.x >= midX+1 && robot.coords.y <= midY-1 {
			robotCount["SE"]++
		} else if robot.coords.x <= midX-1 && robot.coords.y <= midY-1 {
			robotCount["SW"]++
		}
	}

	return robotCount["NW"] * robotCount["NE"] * robotCount["SE"] * robotCount["SW"]
}

func part2(input string, width int, height int) int {
	robots := parseInput(input)
	grid := createGrid(width, height)
	visited := make(map[point]string)

	var i int
	for ; len(visited) != len(robots); i++ {
		clear(visited)
		moveRobots(&robots, grid)
		for j := range robots {
			visited[point{robots[j].coords.x, robots[j].coords.y}] = ""
		}
	}

	displayGrid(grid, &robots)

	return i
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "which part tu run? 1 or 2")
	flag.Parse()

	res := 0
	if part == 1 {
		res = part1(input, 101, 103)
	} else {
		res = part2(input, 101, 103)
	}

	fmt.Println("Result: ", res)
}
