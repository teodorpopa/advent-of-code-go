package main

import (
	_ "embed"
	"flag"
	"fmt"
	"github.com/teodorpopa/advent-of-code-go/utils"
	"golang.org/x/exp/slices"
	"sync"
)

//go:embed input.txt
var input string

type beam struct {
	direction string
	x         int
	y         int
}

func getEnergizedMap(lines []string) [][]string {
	energized := make([][]string, len(lines))
	for i, line := range lines {
		row := make([]string, len(line))
		for j, _ := range line {
			row[j] = "."
		}
		energized[i] = row
	}

	return energized
}

func getContraption(lines []string) [][]string {
	contraption := make([][]string, len(lines))
	for i, _ := range contraption {
		contraption[i] = make([]string, len(lines[0]))
	}

	for i, line := range lines {
		for j, char := range line {
			contraption[i][j] = string(char)
		}
	}

	return contraption
}

func cloneContraption(contraption [][]string) [][]string {
	newContraption := make([][]string, len(contraption))
	for i := range contraption {
		newContraption[i] = slices.Clone(contraption[i])
	}
	return newContraption
}

func followBeam(lines []string, contraption [][]string, x int, y int, direction string) int {
	energized := getEnergizedMap(lines)

	beams := []beam{}
	defaultBeam := beam{
		direction: direction,
		x:         x,
		y:         y,
	}
	beams = append(beams, defaultBeam)
	energized[0][0] = "#"

	history := []string{}
	key := fmt.Sprintf("%s-%d-%d", defaultBeam.direction, defaultBeam.x, defaultBeam.y)
	history = append(history, key)

	countBeams := true
	for countBeams {

		for i, b := range beams {

			if contraption[b.x][b.y] == "/" {
				if b.direction == "right" {
					b.direction = "up"
				} else if b.direction == "down" {
					b.direction = "left"
				} else if b.direction == "left" {
					b.direction = "down"
				} else if b.direction == "up" {
					b.direction = "right"
				}
			} else if contraption[b.x][b.y] == "\\" {
				if b.direction == "right" {
					b.direction = "down"
				} else if b.direction == "down" {
					b.direction = "right"
				} else if b.direction == "left" {
					b.direction = "up"
				} else if b.direction == "up" {
					b.direction = "left"
				}
			} else if contraption[b.x][b.y] == "|" {
				if b.direction == "right" {
					b.direction = "up"

					beam := beam{
						direction: "down",
						x:         b.x,
						y:         b.y,
					}
					beams = append(beams, beam)
				} else if b.direction == "left" {
					b.direction = "up"

					beam := beam{
						direction: "down",
						x:         b.x,
						y:         b.y,
					}
					beams = append(beams, beam)
				}
			} else if contraption[b.x][b.y] == "-" {
				if b.direction == "down" {
					b.direction = "left"

					beam := beam{
						direction: "right",
						x:         b.x,
						y:         b.y,
					}
					beams = append(beams, beam)

				} else if b.direction == "up" {
					b.direction = "left"

					beam := beam{
						direction: "right",
						x:         b.x,
						y:         b.y,
					}
					beams = append(beams, beam)
				}
			}

			if b.direction == "right" {
				b.y++
			} else if b.direction == "down" {
				b.x++
			} else if b.direction == "left" {
				b.y--
			} else if b.direction == "up" {
				b.x--
			}

			key := fmt.Sprintf("%s-%d-%d", b.direction, b.x, b.y)
			if b.x < 0 || b.y < 0 || b.x >= len(contraption) || b.y >= len(contraption[b.x]) || slices.Contains(history, key) {
				beams = append(beams[:i], beams[i+1:]...)
				break
			}

			beams[i] = b
			energized[b.x][b.y] = "#"
			history = append(history, key)
		}

		if len(beams) == 0 {
			countBeams = false
		}

	}

	sum := 0
	for i, _ := range energized {
		for j, _ := range energized[i] {
			if energized[i][j] == "#" {
				sum++
			}
		}
	}

	return sum
}

func part1(input string) int {
	lines := utils.ReadLines(input)
	contraption := getContraption(lines)

	sum := followBeam(lines, contraption, 0, 0, "right")
	return sum
}

func part2(input string) int {
	lines := utils.ReadLines(input)
	contraption := getContraption(lines)

	out := make(chan int)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := range contraption {
			out <- followBeam(lines, cloneContraption(contraption), i, 0, "right")
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := range contraption {
			out <- followBeam(lines, cloneContraption(contraption), i, len(contraption)-1, "left")
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := range contraption[0] {
			out <- followBeam(lines, cloneContraption(contraption), 0, i, "down")
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := range contraption[0] {
			out <- followBeam(lines, cloneContraption(contraption), len(contraption)-1, i, "up")
		}
	}()

	go func() {
		wg.Wait()
		close(out)
	}()

	var maximum int
	for i := range out {
		maximum = utils.Max(maximum, i)
	}
	return maximum
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
