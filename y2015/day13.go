package y2015

import (
	"fmt"
	"github.com/teodorpopa/advent-of-code-go/utils"
	"golang.org/x/exp/slices"
	"strings"
)

type tableSeat struct {
	left  person
	right person
	gain  int
}

type person string

func getPersons(data []string, includeMe bool) ([]tableSeat, []person) {
	var persons []person
	var seats []tableSeat
	for _, l := range data {
		parts := strings.Split(l, " ")

		gain := 0
		if parts[2] == "gain" {
			gain = utils.ToInt(parts[3])
		} else {
			gain = -(utils.ToInt(parts[3]))
		}

		seats = append(seats, tableSeat{
			left:  person(parts[0]),
			right: person(parts[10][:len(parts[10])-1]),
			gain:  gain,
		})

		if !slices.Contains(persons, person(parts[0])) {
			persons = append(persons, person(parts[0]))
		}
		if !slices.Contains(persons, person(parts[10][:len(parts[10])-1])) {
			persons = append(persons, person(parts[10][:len(parts[10])-1]))
		}
	}

	if includeMe {
		for _, p := range persons {
			seats = append(seats, tableSeat{
				left:  p,
				right: "Me",
				gain:  0,
			})
			seats = append(seats, tableSeat{
				left:  "Me",
				right: p,
				gain:  0,
			})
		}

		persons = append(persons, "Me")
	}

	return seats, persons
}

func getHappiness(seats []tableSeat, left person, right person) int {
	happiness := 0
	for _, v := range seats {
		if v.left == left && v.right == right {
			happiness += v.gain
		} else if v.right == left && v.left == right {
			happiness += v.gain
		}
	}

	return happiness
}

func Day13First(input string) int {
	maxHappiness := 0
	lines := utils.ReadLines(input)
	seats, persons := getPersons(lines, false)

	s := make([]interface{}, len(persons))
	for i, v := range persons {
		s[i] = v
	}
	permutations := utils.Permutations(s)

	for _, persons := range permutations {
		curHap := 0
		var left person
		var right person

		for i := 0; i < len(persons); i++ {

			if i == 0 {
				left = persons[0].(person)
				right = persons[i+1].(person)
			} else if i+1 == len(persons) {
				right = persons[0].(person)
			} else {
				right = persons[i+1].(person)
			}

			curHap += getHappiness(seats, left, right)
			left = right
		}

		if maxHappiness == 0 || curHap > maxHappiness {
			maxHappiness = curHap
		}
	}

	return maxHappiness
}

func Day13Second(input string) int {
	maxHappiness := 0
	lines := utils.ReadLines(input)
	seats, persons := getPersons(lines, true)

	s := make([]interface{}, len(persons))
	for i, v := range persons {
		s[i] = v
	}
	permutations := utils.Permutations(s)

	for _, persons := range permutations {
		curHap := 0
		var left person
		var right person

		for i := 0; i < len(persons); i++ {

			if i == 0 {
				left = persons[0].(person)
				right = persons[i+1].(person)
			} else if i+1 == len(persons) {
				right = persons[0].(person)
			} else {
				right = persons[i+1].(person)
			}

			curHap += getHappiness(seats, left, right)
			left = right
		}

		if maxHappiness == 0 || curHap > maxHappiness {
			maxHappiness = curHap
		}
	}

	return maxHappiness
}

func Day13() {
	fmt.Println(utils.DAY_PREFIX, "Day 13")
	fmt.Println(utils.PART_PREFIX, "Part 1:", Day13First(utils.ReadFile("y2015/input/day13.txt")))
	fmt.Println(utils.PART_PREFIX, "Part 2:", Day13Second(utils.ReadFile("y2015/input/day13.txt")))
}
