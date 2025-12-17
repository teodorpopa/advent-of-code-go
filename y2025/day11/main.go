package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"

	"github.com/teodorpopa/advent-of-code-go/utils"
)

//go:embed input.txt
var input string

type Node struct {
	Code     string
	OutCodes []string
}

func readNodes(input string) map[string]Node {
	nodes := map[string]Node{}
	for _, line := range utils.ReadLines(input) {
		split := strings.Split(line, ": ")
		code := split[0]

		outCodes := strings.Split(split[1], " ")
		nodes[code] = Node{code, outCodes}
	}

	return nodes
}

func countPaths(code, destNode string, nodes map[string]Node, memo map[string]int) int {
	if count, exists := memo[code]; exists {
		return count
	}

	if code == destNode {
		return 1
	}

	node, exists := nodes[code]
	if !exists {
		return 0
	}

	totalPaths := 0
	for _, nextCode := range node.OutCodes {
		totalPaths += countPaths(nextCode, destNode, nodes, memo)
	}

	memo[code] = totalPaths
	return totalPaths
}

func countAllPaths(currentCode, destinationNode string, nodes map[string]Node) int {
	memo := make(map[string]int)
	return countPaths(currentCode, destinationNode, nodes, memo)
}

func part1(input string) int {
	nodes := readNodes(input)
	count := countAllPaths("you", "out", nodes)

	return count
}

func part2(input string) int {
	nodes := readNodes(input)

	case1 := countAllPaths("svr", "dac", nodes) * countAllPaths("dac", "fft", nodes) * countAllPaths("fft", "out", nodes)
	case2 := countAllPaths("svr", "fft", nodes) * countAllPaths("fft", "dac", nodes) * countAllPaths("dac", "out", nodes)
	count := case1 + case2

	return count
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
