package y2015

import (
	"fmt"
	"github.com/teodorpopa/advent-of-code-go/utils"
	"golang.org/x/exp/slices"
	"strings"
)

type Point struct {
	X int
	Y int
}

func getNewCoords(direction string, x *int, y *int) Point {
	switch direction {
	case "^":
		*y++
		break
	case "v":
		*y--
		break
	case "<":
		*x--
		break
	case ">":
		*x++
		break
	}

	return Point{
		X: *x,
		Y: *y,
	}
}

func Day03First(input string) int {
	houses := []Point{
		{
			X: 0,
			Y: 0,
		},
	}

	lastX := 0
	lastY := 0
	for _, v := range strings.Split(input, "") {
		newPoint := getNewCoords(v, &lastX, &lastY)

		if !slices.Contains(houses, newPoint) {
			houses = append(houses, newPoint)
		}
	}

	return len(houses)
}

func Day03Second(input string) int {
	houses := []Point{
		{
			X: 0,
			Y: 0,
		},
	}

	lastXSanta := 0
	lastYSanta := 0
	lastXRobo := 0
	lastYRobo := 0

	for k, v := range strings.Split(input, "") {
		if k%2 == 0 {
			newPoint := getNewCoords(v, &lastXSanta, &lastYSanta)
			if !slices.Contains(houses, newPoint) {
				houses = append(houses, newPoint)
			}
		} else {
			newPoint := getNewCoords(v, &lastXRobo, &lastYRobo)
			if !slices.Contains(houses, newPoint) {
				houses = append(houses, newPoint)
			}
		}
	}

	return len(houses)
}

func Day03() {
	fmt.Println(utils.DAY_PREFIX, "Day 03")
	fmt.Println(utils.PART_PREFIX, "Part 1:", Day03First(utils.ReadFile("y2015/input/day03.txt")))
	fmt.Println(utils.PART_PREFIX, "Part 2:", Day03Second(utils.ReadFile("y2015/input/day03.txt")))
}
