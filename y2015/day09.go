package y2015

import (
	"fmt"
	"github.com/teodorpopa/advent-of-code-go/utils"
	"golang.org/x/exp/slices"
	"strings"
)

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

func permutations(arr []city) [][]city {
	var helper func([]city, int)
	var res [][]city

	helper = func(arr []city, n int) {
		if n == 1 {
			tmp := make([]city, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
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

func Day09First(input string) int {
	minDistance := 0
	lines := utils.ReadLines(input)
	destinations, cities := getDestinations(lines)

	permutations := permutations(cities)

	for _, cities := range permutations {
		curDist := 0
		var from city

		for i := 1; i < len(cities); i++ {
			if i == 1 {
				from = cities[0]
			}
			to := cities[i]
			curDist += getDistance(destinations, from, to)
			from = to
		}

		if minDistance == 0 || curDist < minDistance {
			minDistance = curDist
		}
	}

	return minDistance
}

func Day09Second(input string) int {
	maxDistance := 0
	lines := utils.ReadLines(input)
	destinations, cities := getDestinations(lines)

	permutations := permutations(cities)

	for _, cities := range permutations {
		curDist := 0
		var from city

		for i := 1; i < len(cities); i++ {
			if i == 1 {
				from = cities[0]
			}
			to := cities[i]
			curDist += getDistance(destinations, from, to)
			from = to
		}

		if curDist > maxDistance {
			maxDistance = curDist
		}
	}

	return maxDistance
}

func Day09() {
	fmt.Println(utils.DAY_PREFIX, "Day 09")
	fmt.Println(utils.PART_PREFIX, "Part 1:", Day09First(utils.ReadFile("y2015/input/day09.txt")))
	fmt.Println(utils.PART_PREFIX, "Part 2:", Day09Second(utils.ReadFile("y2015/input/day09.txt")))
}