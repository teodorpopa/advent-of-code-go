package main

import (
	_ "embed"
	"flag"
	"fmt"
	"github.com/teodorpopa/advent-of-code-go/utils"
	"strings"
)

//go:embed input.txt
var input string

type info struct {
	x, y float64
}

type machineData struct {
	btnA, btnB, prize info
}

func parseInput(part int) []machineData {
	machines := strings.Split(input, "\n\n")
	m := make([]machineData, 0)

	for _, machine := range machines {
		lines := strings.Split(machine, "\n")

		btnAInfo := strings.Split(strings.Split(lines[0], ": ")[1], ", ")
		btnA := info{
			x: utils.ToFloat(btnAInfo[0][2:]),
			y: utils.ToFloat(btnAInfo[1][2:]),
		}

		btnBInfo := strings.Split(strings.Split(lines[1], ": ")[1], ", ")
		btnB := info{
			x: utils.ToFloat(btnBInfo[0][2:]),
			y: utils.ToFloat(btnBInfo[1][2:]),
		}

		prizeInfo := strings.Split(strings.Split(lines[2], ": ")[1], ", ")

		var add float64
		if part == 2 {
			add = 10_000_000_000_000
		}

		prize := info{
			x: utils.ToFloat(prizeInfo[0][2:]) + add,
			y: utils.ToFloat(prizeInfo[1][2:]) + add,
		}

		m = append(m, machineData{
			btnA,
			btnB,
			prize,
		})
	}

	return m
}

func validTry(num float64) bool {
	return num <= 100 && num > 0
}

func part1(input string) int {
	machines := parseInput(1)

	coins := 0
	for _, machine := range machines {
		timesB := (machine.prize.y*machine.btnA.x - machine.prize.x*machine.btnA.y) / (machine.btnB.y*machine.btnA.x - machine.btnB.x*machine.btnA.y)
		timesA := (machine.prize.x - machine.btnB.x*timesB) / machine.btnA.x

		if utils.IsInt(timesA) && utils.IsInt(timesB) && validTry(timesA) && validTry(timesB) {
			coins += int(timesA)*3 + int(timesB)
		}
	}

	return coins
}

func part2(input string) int {
	machines := parseInput(2)

	coins := 0
	for _, machine := range machines {
		timesB := (machine.prize.y*machine.btnA.x - machine.prize.x*machine.btnA.y) / (machine.btnB.y*machine.btnA.x - machine.btnB.x*machine.btnA.y)
		timesA := (machine.prize.x - machine.btnB.x*timesB) / machine.btnA.x

		if utils.IsInt(timesA) && utils.IsInt(timesB) {
			coins += int(timesA)*3 + int(timesB)
		}
	}

	return coins
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
