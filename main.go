package main

import (
	"github.com/gookit/color"
	"github.com/teodorpopa/advent-of-code-go/utils"
	"github.com/teodorpopa/advent-of-code-go/y2015"
	"github.com/teodorpopa/advent-of-code-go/y2023"
)

const title = "   GO Advent of Code   "

func main() {

	utils.DisplayTitle(title)
	y, d := utils.ReadArgs()

	switch y {
	case 2015:
		switch d {
		case 1:
			y2015.Day01()
		case 2:
			y2015.Day02()
		case 3:
			y2015.Day03()
		case 4:
			y2015.Day04()
		case 5:
			y2015.Day05()
		case 6:
			y2015.Day06()
		case 7:
			y2015.Day07()
		case 8:
			y2015.Day08()
		case 9:
			y2015.Day09()
		case 10:
			y2015.Day10()
		case 11:
			y2015.Day11()
		case 12:
			y2015.Day12()
		case 13:
			y2015.Day13()
		case 14:
			y2015.Day14()
		default:
			color.Error.Println("Invalid day")
		}
	case 2023:
		switch d {
		case 1:
			y2023.Day01()
		case 2:
			y2023.Day02()
		case 4:
			y2023.Day04()
		}
	default:
		color.Error.Println("Invalid year")
	}

}
