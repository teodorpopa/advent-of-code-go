package y2023

import (
	"github.com/teodorpopa/advent-of-code-go/utils"
	"github.com/teodorpopa/advent-of-code-go/y2023/day01"
	"github.com/teodorpopa/advent-of-code-go/y2023/day02"
	"github.com/teodorpopa/advent-of-code-go/y2023/day03"
	"github.com/teodorpopa/advent-of-code-go/y2023/day04"
	"github.com/teodorpopa/advent-of-code-go/y2023/day05"
	"github.com/teodorpopa/advent-of-code-go/y2023/day06"
	"github.com/teodorpopa/advent-of-code-go/y2023/day07"
	"github.com/teodorpopa/advent-of-code-go/y2023/day08"
	"github.com/teodorpopa/advent-of-code-go/y2023/day09"
	"github.com/teodorpopa/advent-of-code-go/y2023/day10"
	"github.com/teodorpopa/advent-of-code-go/y2023/day11"
	"github.com/teodorpopa/advent-of-code-go/y2023/day12"
	"github.com/teodorpopa/advent-of-code-go/y2023/day13"
	"github.com/teodorpopa/advent-of-code-go/y2023/day14"
	"github.com/teodorpopa/advent-of-code-go/y2023/day15"
	"github.com/teodorpopa/advent-of-code-go/y2023/day16"
	"github.com/teodorpopa/advent-of-code-go/y2023/day17"
	"github.com/teodorpopa/advent-of-code-go/y2023/day18"
	"github.com/teodorpopa/advent-of-code-go/y2023/day19"
	"github.com/teodorpopa/advent-of-code-go/y2023/day20"
	"github.com/teodorpopa/advent-of-code-go/y2023/day21"
	"github.com/teodorpopa/advent-of-code-go/y2023/day22"
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
	17: day17.Solve,
	18: day18.Solve,
	19: day19.Solve,
	20: day20.Solve,
	21: day21.Solve,
	22: day22.Solve,
}

func Run(day int, parts []int) {
	utils.RunDay(days, day, parts)
}
