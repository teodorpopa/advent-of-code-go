package day14

import (
	_ "embed"
	"fmt"
	"github.com/teodorpopa/advent-of-code-go/utils"
	"sort"
	"strings"
)

//go:embed input.txt
var input string

type reindeer struct {
	name      string
	speed     int
	speedTime int
	restTime  int
}

type distance struct {
	reindeer       reindeer
	remainingSpeed int
	remainingRest  int
	dist           int
	score          int
}

func getReindeers(input string) []reindeer {
	lines := utils.ReadLines(input)
	var reindeers []reindeer

	for _, l := range lines {
		parts := strings.Split(l, " ")

		reindeers = append(reindeers, reindeer{
			name:      parts[0],
			speed:     utils.ToInt(parts[3]),
			speedTime: utils.ToInt(parts[6]),
			restTime:  utils.ToInt(parts[13]),
		})
	}

	return reindeers
}

func part1(input string, iterations int) int {
	reindeers := getReindeers(input)
	var distances []distance

	for _, r := range reindeers {
		distances = append(distances, distance{
			reindeer:       r,
			remainingSpeed: r.speedTime,
			remainingRest:  r.restTime,
			dist:           0,
		})
	}

	for i := 1; i <= iterations; i++ {
		for k, d := range distances {
			if d.remainingSpeed > 0 {
				d.dist += d.reindeer.speed
				d.remainingSpeed--

			} else if d.remainingSpeed == 0 && d.remainingRest > 0 {
				d.remainingRest--
			}

			if d.remainingSpeed == 0 && d.remainingRest == 0 {
				d.remainingSpeed = d.reindeer.speedTime
				d.remainingRest = d.reindeer.restTime
			}

			distances[k] = d
		}
	}

	sort.Slice(distances, func(i, j int) bool {
		return distances[i].dist > distances[j].dist
	})

	return distances[0].dist
}

func part2(input string, iterations int) int {
	reindeers := getReindeers(input)
	var distances []distance

	for _, r := range reindeers {
		distances = append(distances, distance{
			reindeer:       r,
			remainingSpeed: r.speedTime,
			remainingRest:  r.restTime,
			dist:           0,
		})
	}

	for i := 1; i <= iterations; i++ {
		for k, d := range distances {
			if d.remainingSpeed > 0 {
				d.dist += d.reindeer.speed
				d.remainingSpeed--

			} else if d.remainingSpeed == 0 && d.remainingRest > 0 {
				d.remainingRest--
			}

			if d.remainingSpeed == 0 && d.remainingRest == 0 {
				d.remainingSpeed = d.reindeer.speedTime
				d.remainingRest = d.reindeer.restTime
			}

			distances[k] = d
		}

		sort.Slice(distances, func(i, j int) bool {
			return distances[i].dist > distances[j].dist
		})
		maxDist := 0
		for m, di := range distances {
			if m == 0 {
				maxDist = di.dist
				di.score++
				distances[m] = di
			} else if maxDist == di.dist {
				di.score++
				distances[m] = di
			}
		}
	}

	sort.Slice(distances, func(i, j int) bool {
		return distances[i].score > distances[j].score
	})

	return distances[0].score
}

func Solve(part int) int {
	if part == 1 {
		return part1(utils.ReadFile(input), 2503)
	} else if part == 2 {
		return part2(utils.ReadFile(input), 2503)
	} else {
		fmt.Println("Invalid part to run")
		return -1
	}
}
