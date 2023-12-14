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

var years = map[int]interface{}{
	2015: y2015.Run,
	2016: y2016.Run,
	2017: y2017.Run,
	2018: y2018.Run,
	2019: y2019.Run,
	2020: y2020.Run,
	2021: y2021.Run,
	2022: y2022.Run,
	2023: y2023.Run,
}

func Run(year int, day int, partsToRun []int) {
	color.Cyanf("\nYear: %d \n", year)

	v, ext := years[year]
	if !ext {
		color.Errorf("Year %d has not been implemented yet\n\n", year)
		return
	}

	v.(func(int, []int))(day, partsToRun)

}
