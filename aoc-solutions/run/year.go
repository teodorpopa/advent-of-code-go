package run

import (
	"github.com/gookit/color"
	"github.com/teodorpopa/advent-of-code-go/y2015"
	"github.com/teodorpopa/advent-of-code-go/y2016"
	"github.com/teodorpopa/advent-of-code-go/y2017"
	"github.com/teodorpopa/advent-of-code-go/y2018"
	"github.com/teodorpopa/advent-of-code-go/y2019"
	"github.com/teodorpopa/advent-of-code-go/y2020"
	"github.com/teodorpopa/advent-of-code-go/y2021"
	"github.com/teodorpopa/advent-of-code-go/y2022"
	"github.com/teodorpopa/advent-of-code-go/y2023"
)

func Run(year int, day int, partsToRun []int) bool {
	color.Cyanf("\nYear: %d \n", year)

	switch year {
	case 2015:
		y2015.Run(day, partsToRun)
	case 2016:
		y2016.Run(day, partsToRun)
	case 2017:
		y2017.Run(day, partsToRun)
	case 2018:
		y2018.Run(day, partsToRun)
	case 2019:
		y2019.Run(day, partsToRun)
	case 2020:
		y2020.Run(day, partsToRun)
	case 2021:
		y2021.Run(day, partsToRun)
	case 2022:
		y2022.Run(day, partsToRun)
	case 2023:
		y2023.Run(day, partsToRun)
	}

	return true
}
