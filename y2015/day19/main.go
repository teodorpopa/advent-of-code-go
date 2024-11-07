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

func part1(input string) int {
	replaceMap, molecule := parseInput(input)
	var combinations []string

	for k, v := range replaceMap {
		indexes := allIndicesForString(k, molecule)
		for _, i := range indexes {
			start := molecule[:i]
			end := molecule[i+len(k):]

			for _, com := range v {
				newMol := start + com + end
				if !slices.Contains(combinations, newMol) {
					combinations = append(combinations, newMol)
				}
			}
		}
	}

	return len(combinations)
}

func part2(input string) int {
	replaceMap, molecule := parseInputPart2(input)
	originalMolecule := molecule
	replaceSlice := replaceSlice(replaceMap)

	var steps int
	for molecule != "e" {
		var changeMade bool

		for _, rep := range replaceSlice {
			count := strings.Count(molecule, rep)
			if count <= 0 {
				continue
			}
			changeMade = true
			steps += count
			molecule = strings.ReplaceAll(molecule, rep, replaceMap[rep])
			break
		}

		if !changeMade {
			replaceSlice = utils.ShuffleSlice(replaceSlice)
			molecule = originalMolecule
			steps = 0
		}
	}

	return steps
}

func replaceSlice(replaceMap map[string]string) []string {
	var slice []string

	for value := range replaceMap {
		slice = append(slice, value)
	}

	return slice
}

func allIndicesForString(s, in string) []int {
	var indices []int

	index := strings.Index(in, s)
	offset := 0
	for index > -1 {
		indices = append(indices, index+offset)
		offset += len(in[:index+len(s)])
		in = in[index+len(s):]
		index = strings.Index(in, s)
	}

	return indices
}

func parseInput(input string) (map[string][]string, string) {
	inputs := strings.Split(input, "\n\n")

	replaces := utils.ReadLines(inputs[0])
	molecule := inputs[1]

	replaceMap := map[string][]string{}
	for _, l := range replaces {
		parts := strings.Split(l, " => ")
		replaceMap[parts[0]] = append(replaceMap[parts[0]], parts[1])
	}

	return replaceMap, molecule
}

func parseInputPart2(input string) (map[string]string, string) {
	inputs := strings.Split(input, "\n\n")

	replaces := utils.ReadLines(inputs[0])
	molecule := inputs[1]

	replaceMap := map[string]string{}
	for _, l := range replaces {
		parts := strings.Split(l, " => ")
		replaceMap[parts[1]] = parts[0]
	}

	return replaceMap, molecule
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
