package main

import (
	_ "embed"
	"flag"
	"fmt"
	"github.com/teodorpopa/advent-of-code-go/utils"
	"golang.org/x/exp/slices"
	"strings"
)

//go:embed input.txt
var input string

type destination struct {
	from     city
	to       city
	distance int
}

type city string

func getDestinations(dest []string) ([]destination, []city) {
	var destinations []destination
	var cities []city
	for _, l := range dest {
		parts := strings.Split(l, " = ")
		lParts := strings.Split(parts[0], " to ")

		destinations = append(destinations, destination{
			from:     city(lParts[0]),
			to:       city(lParts[1]),
			distance: utils.ToInt(parts[1]),
		})

		if !slices.Contains(cities, city(lParts[0])) {
			cities = append(cities, city(lParts[0]))
		}
		if !slices.Contains(cities, city(lParts[1])) {
			cities = append(cities, city(lParts[1]))
		}

	}

	return destinations, cities
}

func getDistance(destinations []destination, from city, to city) int {
	for _, v := range destinations {
		if v.from == from && v.to == to {
			return v.distance
		} else if v.to == from && v.from == to {
			return v.distance
		}
	}

	return 0
}

func part1(input string) int {
	minDistance := 0
	lines := utils.ReadLines(input)
	destinations, cities := getDestinations(lines)

	s := make([]interface{}, len(cities))
	for i, v := range cities {
		s[i] = v
	}
	permutations := utils.Permutations(s)

	for _, cities := range permutations {
		curDist := 0
		var from city

		for i := 1; i < len(cities); i++ {
			if i == 1 {
				from = cities[0].(city)
			}
			to := cities[i].(city)
			curDist += getDistance(destinations, from, to)
			from = to
		}

		if minDistance == 0 || curDist < minDistance {
			minDistance = curDist
		}
	}

	return minDistance
}

func part2(input string) int {
	maxDistance := 0
	lines := utils.ReadLines(input)
	destinations, cities := getDestinations(lines)

	s := make([]interface{}, len(cities))
	for i, v := range cities {
		s[i] = v
	}
	permutations := utils.Permutations(s)

	for _, cities := range permutations {
		curDist := 0
		var from city

		for i := 1; i < len(cities); i++ {
			if i == 1 {
				from = cities[0].(city)
			}
			to := cities[i].(city)
			curDist += getDistance(destinations, from, to)
			from = to
		}

		if curDist > maxDistance {
			maxDistance = curDist
		}
	}

	return maxDistance
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
