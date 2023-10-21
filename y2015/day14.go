package y2015

import (
	"fmt"
	"github.com/teodorpopa/advent-of-code-go/utils"
	"sort"
	"strings"
)

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

func Day14First(input string, iterations int) int {
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

func Day14Second(input string, iterations int) int {
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

func Day14() {
	fmt.Println(utils.DAY_PREFIX, "Day 14")
	fmt.Println(utils.PART_PREFIX, "Part 1:", Day14First(utils.ReadFile("y2015/input/day14.txt"), 2503))
	fmt.Println(utils.PART_PREFIX, "Part 2:", Day14Second(utils.ReadFile("y2015/input/day14.txt"), 2503))
}
