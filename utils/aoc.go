package utils

import (
	"bufio"
	"fmt"
	"github.com/gookit/color"
	"os"
	"strconv"
	"strings"
)

const DAY_PREFIX = "  -"
const PART_PREFIX = "      "

func ReadFile(path string) string {
	buf, _ := os.ReadFile(path)
	return string(buf)
}

func ReadLines(input string) []string {
	return strings.Split(input, "\n")
}

func ToInt(i string) int {
	v, _ := strconv.Atoi(i)
	return v
}

func ReadArgs() (int, int) {
	args := os.Args

	var y, d int
	var yearErr, dayErr error

	if len(args) == 4 {
		y, yearErr = strconv.Atoi(args[2])
		d, dayErr = strconv.Atoi(args[3])
	} else if len(args) == 3 {
		y, yearErr = strconv.Atoi(args[1])
		d, dayErr = strconv.Atoi(args[2])
	} else {
		y, yearErr = ReadInt("Year to run: ")
		d, dayErr = ReadInt("Day to run: ")
	}

	if yearErr != nil {
		fmt.Println("Please enter a valid year: 2015 - 2022")
	}
	if dayErr != nil {
		fmt.Println("Please enter a valid day: 1 - 25")
	}

	if len(args) > 1 && args[1] == "create" {
		CreateFromTemplate(y, d)
		os.Exit(0)
	}

	color.Printf("\n<suc>Running day</> <comment>%02d</> <suc>from year</> <comment>%02d</>\n", d, y)
	return y, d
}

func CreateFromTemplate(y int, d int) {
	f, err := os.Create(fmt.Sprintf("y%04d/day%02d.go", y, d))
	Panic(err)
	defer f.Close()
	f.WriteString(fmt.Sprintf(`package y%04d

import (
	"fmt"
	"github.com/teodorpopa/advent-of-code-go/utils"
)

func Day%02dFirst(input string) int {
	return 0
}

func Day%02dSecond(input string) int {
	return 0
}

func Day%02d() {
	fmt.Println(utils.DAY_PREFIX, "Day %02d")
	fmt.Println(utils.PART_PREFIX, "Part 1:", Day%02dFirst(utils.ReadFile("y%04d/input/day%02d.txt")))
	fmt.Println(utils.PART_PREFIX, "Part 2:", Day%02dSecond(utils.ReadFile("y%04d/input/day%02d.txt")))
}
`, y, d, d, d, d, d, y, d, d, y, d))
	color.Printf("\n<suc>Wrote day</> <comment>y%04d/day%02d.go</>\n", y, d)

	f, err = os.Create(fmt.Sprintf("y%04d/day%02d_test.go", y, d))
	Panic(err)
	defer f.Close()
	f.WriteString(fmt.Sprintf(`package y%04d

import (
	"github.com/stretchr/testify/assert"
	"github.com/teodorpopa/advent-of-code-go/utils"
	"testing"
)

func TestDay%02dFirst(t *testing.T) {
	r := Day%02dFirst(utils.ReadFile("test_input/day%02d.txt"))
	assert.Equal(t, r, 0)
}

func TestDay%02dSecond(t *testing.T) {
	r := Day%02dSecond(utils.ReadFile("test_input/day%02d.txt"))
	assert.Equal(t, r, 0)
}
`, y, d, d, d, d, d, d))
	color.Printf("\n<suc>Wrote day</> <comment>y%04d/day%02d_test.go</>\n", y, d)

	f, err = os.Create(fmt.Sprintf("y%04d/input/day%02d.txt", y, d))
	color.Printf("\n<suc>Wrote day</> <comment>y%04d/input/day%02d.txt</>\n", y, d)
	Panic(err)

	f, err = os.Create(fmt.Sprintf("y%04d/test_input/day%02d.txt", y, d))
	color.Printf("\n<suc>Wrote day</> <comment>y%04d/test_input/day%02d.txt</>\n", y, d)
	Panic(err)
}

func Panic(err error) {
	if err != nil {
		panic(err)
	}
}

func ReadInt(label string) (int, error) {
	var s string
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stderr, label+" ")
		s, _ = r.ReadString('\n')
		if s != "" {
			break
		}
	}
	return strconv.Atoi(strings.TrimSpace(s))
}

func DisplayTitle(title string) {
	color.New(color.FgGreen, color.BgCyan).Println(strings.Repeat("=", len(title)))
	color.New(color.FgBlack, color.BgCyan).Println(title)
	color.New(color.FgGreen, color.BgCyan).Println(strings.Repeat("=", len(title)))
	fmt.Println("")
}
