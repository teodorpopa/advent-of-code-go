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

	if len(args) == 3 {
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

	color.Printf("\n<suc>Running day</> <comment>%02d</> <suc>from year</> <comment>%02d</>\n", d, y)
	return y, d
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
