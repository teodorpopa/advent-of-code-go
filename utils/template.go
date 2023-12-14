package utils

import (
	"fmt"
	"github.com/gookit/color"
	"os"
)

func pathExists(path string) bool {
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		return true
	}
	return false
}

func CreateFromTemplate(y int, d int) {
	if !pathExists(fmt.Sprintf("y%04d/day%02d", y, d)) {
		err := os.MkdirAll(fmt.Sprintf("y%04d/day%02d", y, d), os.ModePerm)
		Panic(err)
	} else {
		color.Printf("\n<err>Path y%04d/day%02d/ already exists</>\n", y, d)
		return
	}

	dayFile := fmt.Sprintf("y%04d/day%02d/day.go", y, d)
	testFile := fmt.Sprintf("y%04d/day%02d/day_test.go", y, d)
	inputFile := fmt.Sprintf("y%04d/day%02d/input.txt", y, d)
	inputTestFile := fmt.Sprintf("y%04d/day%02d/input_test.txt", y, d)

	f, err := os.Create(dayFile)
	Panic(err)
	defer f.Close()
	f.WriteString(fmt.Sprintf(`package day%02d

import (
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input string

func part1(input string) int {
	
	return 0
}

func part2(input string) int {
	
	return 0
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
`, d))

	color.Print("\n<suc>Wrote day file:</> ")
	color.Printf("<comment>%s</>\n", dayFile)

	f, err = os.Create(testFile)
	Panic(err)
	defer f.Close()
	f.WriteString(fmt.Sprintf(`package day%02d

import (
	"github.com/stretchr/testify/assert"
	"github.com/teodorpopa/advent-of-code-go/utils"
	"testing"
)

func TestDayPart1(t *testing.T) {
	r := part1(utils.ReadFile("input_test.txt"))
	assert.Equal(t, 0, r)
}

func TestDayPart2(t *testing.T) {
	r := part2(utils.ReadFile("input_test.txt"))
	assert.Equal(t, 0, r)
}

`, d))

	color.Print("\n<suc>Wrote test file:</> ")
	color.Printf("<comment>%s</>\n", testFile)

	f, err = os.Create(inputFile)
	color.Print("\n<suc>Wrote input file:</> ")
	color.Printf("<comment>%s</>\n", inputFile)
	Panic(err)

	f, err = os.Create(inputTestFile)
	color.Print("\n<suc>Wrote input test file:</> ")
	color.Printf("<comment>%s</>\n", inputTestFile)
	Panic(err)
}

func Panic(err error) {
	if err != nil {
		panic(err)
	}
}
