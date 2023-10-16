package main

import (
	"github.com/gookit/color"
	"github.com/teodorpopa/advent-of-code-go/utils"
	"github.com/teodorpopa/advent-of-code-go/y2015"
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
		default:
			color.Error.Println("Invalid day")
		}
	default:
		color.Error.Println("Invalid year")
	}

}