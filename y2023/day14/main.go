package main

import (
	"bytes"
	_ "embed"
	"flag"
	"fmt"
	"github.com/teodorpopa/advent-of-code-go/utils"
)

//go:embed input.txt
var input string

const billion = 1_000_000_000

var byteInput [][]byte

func parseInput(input string) {
	lines := utils.ReadLines(input)
	byteInput = make([][]byte, len(lines))
	for i, line := range lines {
		byteInput[i] = []byte(line)
	}
}

func cycleMap1(cols int, rows int) {
	for j := 0; j < cols; j++ {
		partIndex := -1
		check := 0
		for i := 0; i < rows; i++ {
			switch byteInput[i][j] {
			case '.':
				continue
			case '#':
				for n := 0; n < check; n++ {
					byteInput[partIndex+n+1][j] = 'O'
				}
				check = 0
				partIndex = i
			case 'O':
				byteInput[i][j] = '.'
				check++
			}
		}
		for n := 0; n < check; n++ {
			byteInput[partIndex+n+1][j] = 'O'
		}
	}
}

func cycleMap2(cols int, rows int) {
	for i := 0; i < rows; i++ {
		partIndex := -1
		check := 0
		for j := 0; j < cols; j++ {
			switch byteInput[i][j] {
			case '.':
				continue
			case '#':
				for n := 0; n < check; n++ {
					byteInput[i][partIndex+n+1] = 'O'
				}
				check = 0
				partIndex = j
			case 'O':
				byteInput[i][j] = '.'
				check++
			}
		}
		for n := 0; n < check; n++ {
			byteInput[i][partIndex+n+1] = 'O'
		}
	}
}

func cycleMap3(cols int, rows int) {
	for j := 0; j < cols; j++ {
		partIndex := rows
		check := 0
		for i := rows - 1; i >= 0; i-- {
			switch byteInput[i][j] {
			case '.':
				continue
			case '#':
				for n := 0; n < check; n++ {
					byteInput[partIndex-n-1][j] = 'O'
				}
				check = 0
				partIndex = i
			case 'O':
				byteInput[i][j] = '.'
				check++
			}
		}
		for n := 0; n < check; n++ {
			byteInput[partIndex-n-1][j] = 'O'
		}
	}
}

func cycleMap4(cols int, rows int) {
	for i := 0; i < rows; i++ {
		partIndex := cols
		check := 0
		for j := cols - 1; j >= 0; j-- {
			switch byteInput[i][j] {
			case '.':
				continue
			case '#':
				for n := 0; n < check; n++ {
					byteInput[i][partIndex-n-1] = 'O'
				}
				check = 0
				partIndex = j
			case 'O':
				byteInput[i][j] = '.'
				check++
			}
		}
		for n := 0; n < check; n++ {
			byteInput[i][partIndex-n-1] = 'O'
		}
	}
}

func cycleMap(cols int, rows int) {
	cycleMap1(cols, rows)
	cycleMap2(cols, rows)
	cycleMap3(cols, rows)
	cycleMap4(cols, rows)
}

func part1(input string) int {
	parseInput(input)
	sum := 0

	rows := len(byteInput)
	cols := len(byteInput[0])

	for j := 0; j < rows; j++ {
		partIndex := -1
		check := 0

		for i := 0; i < cols; i++ {
			if byteInput[i][j] == '.' {
				continue
			} else if byteInput[i][j] == '#' {
				sum += check*(cols-partIndex-1) - check*(check-1)/2
				check = 0
				partIndex = i
			} else if byteInput[i][j] == 'O' {
				check++
			}

			continue
		}

		if check > 0 {
			sum += check*(cols-partIndex-1) - check*(check-1)/2
		}
	}

	return sum
}

func part2(input string) int {
	parseInput(input)
	sum := 0

	rows := len(byteInput)
	cols := len(byteInput[0])

	breakpoint := billion
	history := map[string]int{}

	cycleMap(cols, rows)
	for k := 0; k < billion; k++ {
		cycleMap(cols, rows)

		if breakpoint == billion {
			h := string(bytes.Join(byteInput, []byte{}))
			if i, ok := history[h]; ok {
				breakpoint = k + (billion-i)%(k-i) - 1
			} else {
				history[h] = k
			}
		}

		if k == breakpoint {
			break
		}

	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if byteInput[i][j] == 'O' {
				sum += rows - i
			}
		}
	}

	return sum
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
