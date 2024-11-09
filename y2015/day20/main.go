package main

import (
	_ "embed"
	"flag"
	"fmt"
	"math"

	"github.com/teodorpopa/advent-of-code-go/utils"
)

//go:embed input.txt
var input string

func part1(input string) int {
	inputValue := utils.ToInt(input)

	for house := 1; house < math.MaxInt32; house++ {
		gifts := 0

		for _, inc := range getHouseIncrements(house) {
			gifts += inc * 10

			if gifts >= inputValue {
				return house
			}
		}
	}

	return 0
}

func part2(input string) int {
	inputValue := utils.ToInt(input)

	for house := 1; house < math.MaxInt32; house++ {
		gifts := 0

		for _, inc := range getHouseIncrements(house) {
			if house/inc <= 50 {
				gifts += inc * 11
			}

			if gifts >= inputValue {
				return house
			}
		}
	}

	return 0
}

func getHouseIncrements(house int) []int {
	var increments []int
	sqrt := int(math.Sqrt(float64(house)))

	for i := 1; i <= sqrt; i++ {
		if house%i == 0 {
			increments = append(increments, i, house/i)
		}
	}

	return increments
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
