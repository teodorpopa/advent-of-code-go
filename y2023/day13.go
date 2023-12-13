package y2023

import (
	"fmt"
	"github.com/teodorpopa/advent-of-code-go/utils"
	"golang.org/x/exp/slices"
	"strings"
)

func readBlocks(input string) []string {
	return strings.Split(input, "\n\n")
}

func stringDiff(a string, b string) int {
	if len(a) != len(b) {
		return -1
	}

	diff := 0
	for i, _ := range a {
		if a[i] != b[i] {
			diff++
		}
	}

	return diff
}

func copyAndReverse(s []string) []string {
	str := make([]string, len(s))
	copy(str, s)

	slices.Reverse(str)

	return str
}

func check(block []string, smudge int) int {
	for i := 1; i < len(block); i++ {
		lBlock := block[:i]
		rBlock := block[i:]

		min := utils.Min(len(lBlock), len(rBlock))
		left := lBlock[len(lBlock)-min:]
		right := rBlock[:min]

		leftStr := strings.Join(left, "")
		rightStr := strings.Join(copyAndReverse(right), "")

		if (smudge == 0 && leftStr == rightStr) || (smudge == 1 && stringDiff(leftStr, rightStr) == 1) {
			return i
		}
	}

	return 0
}

func Day13First(input string) int {
	total := 0

	blocks := readBlocks(input)

	for _, block := range blocks {

		blockLines := utils.ReadLines(block)

		blockRows := []string{}
		blockCols := []string{}

		for _, line := range blockLines {
			blockRows = append(blockRows, line)
		}

		for i, _ := range blockLines[0] {
			blockCol := ""
			for _, line := range blockLines {
				blockCol += string(line[i])
			}
			blockCols = append(blockCols, blockCol)
		}

		mirrorRow := check(blockRows, 0)
		mirrorCol := check(blockCols, 0)

		total += mirrorCol + mirrorRow*100
	}

	return total
}

func Day13Second(input string) int {
	total := 0

	blocks := readBlocks(input)

	for _, block := range blocks {
		blockLines := utils.ReadLines(block)

		blockRows := []string{}
		blockCols := []string{}

		for _, line := range blockLines {
			blockRows = append(blockRows, line)
		}

		for i, _ := range blockLines[0] {
			blockCol := ""
			for _, line := range blockLines {
				blockCol += string(line[i])
			}
			blockCols = append(blockCols, blockCol)
		}

		mirrorRow := check(blockRows, 1)
		mirrorCol := check(blockCols, 1)

		total += mirrorCol + mirrorRow*100
	}

	return total
}

func Day13() {
	fmt.Println(utils.DAY_PREFIX, "Day 13")
	fmt.Println(utils.PART_PREFIX, "Part 1:", Day13First(utils.ReadFile("y2023/input/day13.txt")))
	fmt.Println(utils.PART_PREFIX, "Part 2:", Day13Second(utils.ReadFile("y2023/input/day13.txt")))
}
