package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/teodorpopa/advent-of-code-go/utils"
)

//go:embed input.txt
var input string

type Present struct {
	ID     int
	Area   int
	Layout []string
}

type Region struct {
	ID            int
	Width         int
	Height        int
	Area          int
	PresentCounts []int
}

func splitByEmptyLines(text string) [][]string {
	normalized := strings.ReplaceAll(text, "\r\n", "\n")

	re := regexp.MustCompile(`\n\s*\n`)
	groups := re.Split(normalized, -1)

	result := make([][]string, 0, len(groups))
	for _, group := range groups {
		lines := strings.Split(strings.TrimSpace(group), "\n")
		result = append(result, lines)
	}

	return result
}

// parseInts extracts integers from a space-separated string
func parseInts(text string) []int {
	fields := strings.Fields(text)
	result := make([]int, 0, len(fields))

	for _, field := range fields {
		if num, err := strconv.Atoi(field); err == nil {
			result = append(result, num)
		}
	}

	return result
}

func countChars(char rune, lines []string) int {
	count := 0
	for _, line := range lines {
		count += strings.Count(line, string(char))
	}
	return count
}

func parsePresent(lines []string) Present {
	idLine := strings.TrimRight(lines[0], ":")
	id := utils.ToInt(idLine)

	layout := lines[1:]
	area := countChars('#', layout)

	return Present{
		ID:     id,
		Area:   area,
		Layout: layout,
	}
}

func parseRegion(line string, id int) Region {
	parts := strings.SplitN(line, ":", 2)

	r := Region{ID: id}
	fmt.Sscanf(parts[0], "%dx%d", &r.Width, &r.Height)
	r.Area = r.Width * r.Height

	if len(parts) > 1 {
		r.PresentCounts = parseInts(parts[1])
	}

	return r
}

func parseInput(input string) (map[int]Present, []Region) {
	segments := splitByEmptyLines(input)

	presents := make(map[int]Present, 6)
	for _, presentLines := range segments[:6] {
		p := parsePresent(presentLines)
		presents[p.ID] = p
	}

	regionLines := segments[6]
	regions := make([]Region, 0, len(regionLines))
	for i, line := range regionLines {
		r := parseRegion(line, i)
		regions = append(regions, r)
	}

	return presents, regions
}

func calculateTotals(presents map[int]Present, presentCounts []int) (totalCount, totalArea int) {
	for id, count := range presentCounts {
		if present, exists := presents[id]; exists {
			totalArea += count * present.Area
			totalCount += count
		}
	}
	return totalCount, totalArea
}

func canFitPresents(presents map[int]Present, region Region) bool {
	totalCount, totalArea := calculateTotals(presents, region.PresentCounts)

	if totalArea > region.Area {
		return false
	}

	minRequiredArea := totalCount * 9
	if minRequiredArea <= region.Area {
		return true
	}

	return false
}

func part1(input string) int {
	presents, regions := parseInput(input)

	validRegions := 0
	for _, region := range regions {
		if canFitPresents(presents, region) {
			validRegions++
		}
	}

	return validRegions
}

func part2(input string) int {
	return 0
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "which part to run? 1 or 2")
	flag.Parse()

	var result int
	switch part {
	case 1:
		result = part1(input)
	case 2:
		result = part2(input)
	default:
		log.Fatalf("Invalid part: %d (must be 1 or 2)", part)
	}

	fmt.Println("Result:", result)
}
