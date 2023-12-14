package day06

import (
	_ "embed"
	"fmt"
	"github.com/teodorpopa/advent-of-code-go/utils"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func getLineTimes(line string) []int {
	parts := strings.Split(line, ": ")
	times := strings.Split(strings.Trim(parts[1], " "), "  ")
	data := []int{}

	for _, t := range times {
		if t != "" {
			nt, _ := strconv.Atoi(strings.Trim(t, " "))
			data = append(data, nt)
		}
	}

	return data
}

func getLineTimesS(line string) []string {
	parts := strings.Split(line, ": ")
	times := strings.Split(strings.Trim(parts[1], " "), "  ")
	data := []string{}

	for _, t := range times {
		if t != "" {
			data = append(data, strings.Trim(t, " "))
		}
	}

	return data
}

func mixValues(times []int, records []int) map[int]int {
	data := make(map[int]int)

	for k, v := range times {
		data[v] = records[k]
	}

	return data
}

func part1(input string) int {
	mult := 1
	games := []int{}

	lines := utils.ReadLines(input)

	times := getLineTimes(lines[0])
	records := getLineTimes(lines[1])

	inputData := mixValues(times, records)

	for time, record := range inputData {
		wins := 0

		for press := 0; press < time; press++ {
			remTime := time - press
			dist := remTime * press

			if dist > record {
				wins++
			}
		}
		games = append(games, wins)

	}

	for _, gv := range games {
		mult *= gv
	}

	return mult
}

func part2(input string) int {
	lines := utils.ReadLines(input)

	times := getLineTimesS(lines[0])
	records := getLineTimesS(lines[1])

	time := strings.Join(times, "")
	record := strings.Join(records, "")

	timeI, _ := strconv.Atoi(time)
	recordI, _ := strconv.Atoi(record)

	wins := 0

	for press := 0; press < timeI; press++ {
		remTime := timeI - press
		dist := remTime * press

		if dist > recordI {
			wins++
		}
	}

	return wins
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
