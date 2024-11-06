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

const (
	minArea = 200_000_000_000_000
	maxArea = 400_000_000_000_000
)

type Point struct {
	x, y float64
}

type Position struct {
	x, y, z float64
}

type Hailstone struct {
	position Position
	velocity Position
}

func parseHailstones(lines []string) []Hailstone {
	hailStones := make([]Hailstone, 0, len(lines))

	for _, line := range lines {
		parts := strings.Split(line, " @ ")
		coords := utils.ToIntSlice(strings.Split(parts[0], ","))
		velocities := utils.ToIntSlice(strings.Split(parts[1], ","))

		hailStone := Hailstone{
			Position{
				float64(coords[0]),
				float64(coords[1]),
				float64(coords[2]),
			},
			Position{
				float64(velocities[0]),
				float64(velocities[1]),
				float64(velocities[2]),
			},
		}
		hailStones = append(hailStones, hailStone)
	}

	return hailStones
}

func doInterest(a, b Hailstone) (Point, bool) {
	a2 := Point{a.velocity.x, a.velocity.y}
	b2 := Point{b.velocity.x, b.velocity.y}
	d2 := Point{b.position.x - a.position.x, b.position.y - a.position.y}

	inter := vecInt(a2, b2)
	if inter == 0 {
		return Point{0, 0}, false
	}

	u := vecInt(d2, b2) / inter
	return Point{a.position.x + a.velocity.x*u, a.position.y + a.velocity.y*u}, true
}

func vecInt(a, b Point) float64 {
	return (a.x * b.y) - (a.y * b.x)
}

func checkIntersection(point Point, hailstone Hailstone) bool {
	dx := point.x - hailstone.position.x
	dy := point.y - hailstone.position.y

	return point.x >= minArea && point.x <= maxArea &&
		point.y >= minArea && point.y <= maxArea &&
		(dx > 0) == (hailstone.velocity.x > 0) &&
		(dy > 0) == (hailstone.velocity.y > 0)
}

func findMatchingVel(vel, pv int) []int {
	match := []int{}
	for v := -1000; v < 1000; v++ {
		if v != pv && vel%(v-pv) == 0 {
			match = append(match, v)
		}
	}
	return match
}

func getIntersect(a, b []int) []int {
	result := []int{}
	for _, val := range a {
		if slices.Contains(b, val) {
			result = append(result, val)
		}
	}
	return result
}

func part1(input string) int {
	lines := utils.ReadLines(input)
	hailStones := parseHailstones(lines)

	intersectCount := 0
	for i := 0; i < len(hailStones)-1; i++ {
		for j := i + 1; j < len(hailStones); j++ {
			a, b := hailStones[i], hailStones[j]
			if point, does := doInterest(a, b); does && checkIntersection(point, a) && checkIntersection(point, b) {
				intersectCount++
			}
		}
	}

	return intersectCount
}

func part2(input string) int {
	lines := utils.ReadLines(input)
	hailStones := parseHailstones(lines)

	maybeX, maybeY, maybeZ := []int{}, []int{}, []int{}
	for i := 0; i < len(hailStones)-1; i++ {
		for j := i + 1; j < len(hailStones); j++ {
			a, b := hailStones[i], hailStones[j]
			if a.velocity.x == b.velocity.x {
				nextMaybe := findMatchingVel(int(b.position.x-a.position.x), int(a.velocity.x))
				if len(maybeX) == 0 {
					maybeX = nextMaybe
				} else {
					maybeX = getIntersect(maybeX, nextMaybe)
				}
			}
			if a.velocity.y == b.velocity.y {
				nextMaybe := findMatchingVel(int(b.position.y-a.position.y), int(a.velocity.y))
				if len(maybeY) == 0 {
					maybeY = nextMaybe
				} else {
					maybeY = getIntersect(maybeY, nextMaybe)
				}
			}
			if a.velocity.z == b.velocity.z {
				nextMaybe := findMatchingVel(int(b.position.z-a.position.z), int(a.velocity.z))
				if len(maybeZ) == 0 {
					maybeZ = nextMaybe
				} else {
					maybeZ = getIntersect(maybeZ, nextMaybe)
				}
			}
		}
	}

	var result = 0
	if len(maybeX) == len(maybeY) && len(maybeY) == len(maybeZ) && len(maybeZ) == 1 {
		rockVel := Position{float64(maybeX[0]), float64(maybeY[0]), float64(maybeZ[0])}
		hailStoneA, hailStoneB := hailStones[0], hailStones[1]
		mA := (hailStoneA.velocity.y - rockVel.y) / (hailStoneA.velocity.x - rockVel.x)
		mB := (hailStoneB.velocity.y - rockVel.y) / (hailStoneB.velocity.x - rockVel.x)
		cA := hailStoneA.position.y - (mA * hailStoneA.position.x)
		cB := hailStoneB.position.y - (mB * hailStoneB.position.x)
		xPos := (cB - cA) / (mA - mB)
		yPos := mA*xPos + cA
		time := (xPos - hailStoneA.position.x) / (hailStoneA.velocity.x - rockVel.x)
		zPos := hailStoneA.position.z + (hailStoneA.velocity.z-rockVel.z)*time
		result = int(xPos + yPos + zPos)
	}

	return result
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
