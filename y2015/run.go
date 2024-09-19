package y2015

import (
	"github.com/teodorpopa/advent-of-code-go/utils"
	"github.com/teodorpopa/advent-of-code-go/y2015/day01"
	"github.com/teodorpopa/advent-of-code-go/y2015/day02"
	"github.com/teodorpopa/advent-of-code-go/y2015/day03"
	"github.com/teodorpopa/advent-of-code-go/y2015/day04"
	"github.com/teodorpopa/advent-of-code-go/y2015/day05"
	"github.com/teodorpopa/advent-of-code-go/y2015/day06"
	"github.com/teodorpopa/advent-of-code-go/y2015/day07"
	"github.com/teodorpopa/advent-of-code-go/y2015/day08"
	"github.com/teodorpopa/advent-of-code-go/y2015/day09"
	"github.com/teodorpopa/advent-of-code-go/y2015/day10"
	"github.com/teodorpopa/advent-of-code-go/y2015/day11"
	"github.com/teodorpopa/advent-of-code-go/y2015/day12"
	"github.com/teodorpopa/advent-of-code-go/y2015/day13"
	"github.com/teodorpopa/advent-of-code-go/y2015/day14"
	"github.com/teodorpopa/advent-of-code-go/y2015/day15"
	"github.com/teodorpopa/advent-of-code-go/y2015/day16"
)

var days = map[int]interface{}{
	1:  day01.Solve,
	2:  day02.Solve,
	3:  day03.Solve,
	4:  day04.Solve,
	5:  day05.Solve,
	6:  day06.Solve,
	7:  day07.Solve,
	8:  day08.Solve,
	9:  day09.Solve,
	10: day10.Solve,
	11: day11.Solve,
	12: day12.Solve,
	13: day13.Solve,
	14: day14.Solve,
	15: day15.Solve,
	16: day16.Solve,
}

func Run(day int, parts []int) {
	utils.RunDay(days, day, parts)
}
