package main

import (
	_ "embed"
	"flag"
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
