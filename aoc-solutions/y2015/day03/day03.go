package day03

import (
	_ "embed"
	"fmt"
	"golang.org/x/exp/slices"
	"strings"
)

//go:embed input.txt
var input string

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

func part1(input string) int {
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

func part2(input string) int {
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
