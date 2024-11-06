package main

import (
	_ "embed"
	"flag"
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

	fmt.Println("Result:", res)
}
