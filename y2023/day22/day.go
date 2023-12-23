package day22

import (
	_ "embed"
	"fmt"
	"github.com/teodorpopa/advent-of-code-go/utils"
	"sort"
	"strings"
)

//go:embed input.txt
var input string

type Brick struct {
	X1, Y1, Z1, X2, Y2, Z2 int
}

func isSubset(set1, set2 map[int]struct{}) bool {
	for key := range set1 {
		if _, ok := set2[key]; !ok {
			return false
		}
	}
	return true
}

func parseBricks(lines []string) []Brick {
	var bricks []Brick

	for _, line := range lines {
		var brick Brick
		for i, s := range rs(line, "~", ",") {
			switch i {
			case 0:
				brick.X1 = utils.ToInt(s)
			case 1:
				brick.Y1 = utils.ToInt(s)
			case 2:
				brick.Z1 = utils.ToInt(s)
			case 3:
				brick.X2 = utils.ToInt(s)
			case 4:
				brick.Y2 = utils.ToInt(s)
			case 5:
				brick.Z2 = utils.ToInt(s)
			}
		}
		bricks = append(bricks, brick)
	}

	sort.SliceStable(bricks, func(i, j int) bool {
		return bricks[i].Z1 < bricks[j].Z1
	})

	return bricks
}

func rs(str, s1, s2 string) []string {
	return strings.Split(strings.ReplaceAll(str, s1, s2), s2)
}

func overlaps(b1, b2 Brick) bool {
	return utils.Max(b1.X1, b2.X1) <= utils.Min(b1.X2, b2.X2) &&
		utils.Max(b1.Y1, b2.Y1) <= utils.Min(b1.Y2, b2.Y2)
}

func updateBrickZ(bricks []Brick, index int, maxZ int) {
	bricks[index].Z2 -= bricks[index].Z1 - maxZ
	bricks[index].Z1 = maxZ
}

func buildSupportGraph(bricks []Brick) (map[int]map[int]struct{}, map[int]map[int]struct{}) {
	supR := make(map[int]map[int]struct{})
	supC := make(map[int]map[int]struct{})

	for i := range bricks {
		supR[i] = make(map[int]struct{})
		supC[i] = make(map[int]struct{})
	}

	for j, upper := range bricks {
		for i, lower := range bricks[:j] {
			if overlaps(lower, upper) && upper.Z1 == lower.Z2+1 {
				supR[i][j] = struct{}{}
				supC[j][i] = struct{}{}
			}
		}
	}

	return supR, supC
}

func countValidStacks(bricks []Brick, supR, supC map[int]map[int]struct{}) int {
	total := 0

	for i := range bricks {
		ok := true
		for j := range supR[i] {
			if len(supC[j]) < 2 {
				ok = false
				break
			}
		}
		if ok {
			total++
		}
	}

	return total
}

func countFallingStacks(bricks []Brick, supR, supC map[int]map[int]struct{}) int {
	total := 0

	for i := range bricks {
		q := make([]int, 0)
		for j := range supR[i] {
			if len(supC[j]) == 1 {
				q = append(q, j)
			}
		}

		fall := make(map[int]struct{})
		for _, j := range q {
			fall[j] = struct{}{}
		}
		fall[i] = struct{}{}

		for len(q) > 0 {
			j := q[0]
			q = q[1:]
			for k := range supR[j] {
				if _, ok := fall[k]; !ok {
					if isSubset(supC[k], fall) {
						q = append(q, k)
						fall[k] = struct{}{}
					}
				}
			}
		}

		total += len(fall) - 1
	}

	return total
}

func part1(input string) int {
	lines := utils.ReadLines(input)
	bricks := parseBricks(lines)

	for index, brick := range bricks {
		maxZ := 1
		for _, check := range bricks[:index] {
			if overlaps(brick, check) {
				maxZ = utils.Max(maxZ, check.Z2+1)
			}
		}
		updateBrickZ(bricks, index, maxZ)
	}

	sort.Slice(bricks, func(i, j int) bool {
		return bricks[i].Z1 < bricks[j].Z1
	})

	supR, supC := buildSupportGraph(bricks)

	return countValidStacks(bricks, supR, supC)
}

func part2(input string) int {
	lines := utils.ReadLines(input)
	bricks := parseBricks(lines)

	for index, brick := range bricks {
		maxZ := 1
		for _, check := range bricks[:index] {
			if overlaps(brick, check) {
				maxZ = utils.Max(maxZ, check.Z2+1)
			}
		}
		updateBrickZ(bricks, index, maxZ)
	}

	sort.SliceStable(bricks, func(i, j int) bool {
		return bricks[i].Z1 < bricks[j].Z1
	})

	supR, supC := buildSupportGraph(bricks)

	return countFallingStacks(bricks, supR, supC)
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
