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

type Point struct {
	x, y int
}

type Line struct {
	p1 Point
	p2 Point
}

func (line *Line) len() int {
	if line.p1.x == line.p2.x {
		return utils.Abs(line.p1.y - line.p2.y)
	}
	return utils.Abs(line.p1.x - line.p2.x)
}

func part1(input string) int {
	lines := utils.ReadLines(input)

	x1, y1 := 0, 0
	x2, y2 := 0, 0

	lagoon := []Line{}
	for _, line := range lines {
		parts := strings.Split(line, " ")
		value := utils.ToInt(parts[1])

		switch parts[0] {
		case "R":
			x2 = x1 + value
		case "L":
			x2 = x1 - value
		case "D":
			y2 = y1 + value
		case "U":
			y2 = y1 - value
		}

		lagoon = append(lagoon, Line{
			p1: Point{
				x1,
				y1,
			},
			p2: Point{
				x2,
				y2,
			},
		})

		x1, y1 = x2, y2
	}

	sum := 0

	for i := 0; i < len(lagoon); i++ {
		sum += (lagoon[i].p1.y + lagoon[i].p2.y) * (lagoon[i].p1.x - lagoon[i].p2.x)
	}

	for _, l := range lagoon {
		sum += l.len()
	}

	return sum/2 + 1
}

func part2(input string) int {
	lines := utils.ReadLines(input)

	x1, y1 := 0, 0
	x2, y2 := 0, 0

	lagoon := []Line{}
	for _, line := range lines {
		parts := strings.Split(line, " ")
		value, _ := strconv.ParseInt(parts[2][2:7], 16, 64)

		switch parts[2][7] {
		case '0':
			x2 = x1 + int(value)
		case '2':
			x2 = x1 - int(value)
		case '1':
			y2 = y1 + int(value)
		case '3':
			y2 = y1 - int(value)
		}

		lagoon = append(lagoon, Line{
			p1: Point{
				x1,
				y1,
			},
			p2: Point{
				x2,
				y2,
			},
		})

		x1, y1 = x2, y2
	}

	sum := 0
	for i := 0; i < len(lagoon); i++ {
		sum += (lagoon[i].p1.y + lagoon[i].p2.y) * (lagoon[i].p1.x - lagoon[i].p2.x)
	}

	for _, l := range lagoon {
		sum += l.len()
	}

	return sum/2 + 1
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
