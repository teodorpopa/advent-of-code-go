package day08

import (
	_ "embed"
	"fmt"
	"github.com/teodorpopa/advent-of-code-go/utils"
	"strings"
)

//go:embed input.txt
var input string

type Node struct {
	Left  string
	Right string
}

var Nodes = make(map[string]Node)

func getNodes(lines []string) map[string]Node {
	for _, line := range lines {
		parts := strings.Split(line, " = ")
		nodeName := parts[0]

		dirs := strings.Split(parts[1], ", ")

		node := Node{
			Left:  strings.Trim(dirs[0], "("),
			Right: strings.Trim(dirs[1], ")"),
		}

		Nodes[nodeName] = node
	}

	return Nodes
}

func getLCM(numbers []int) int {
	//Least Common Multiple
	lcm := numbers[0]
	for i := 0; i < len(numbers); i++ {
		num1 := lcm
		num2 := numbers[i]
		gcd := 1
		for num2 != 0 {
			temp := num2
			num2 = num1 % num2
			num1 = temp
		}
		gcd = num1
		lcm = (lcm * numbers[i]) / gcd
	}

	return lcm
}

func part1(input string) int {
	lines := utils.ReadLines(input)

	directions := strings.Split(lines[0], "")
	nodes := getNodes(lines[2:len(lines)])

	totalSteps := 0
	startNode := "AAA"
	endNode := "ZZZ"
	currentNode := startNode

	for {
		for i := 0; i < len(directions); i++ {

			if directions[i] == "L" {
				currentNode = nodes[currentNode].Left
			} else if directions[i] == "R" {
				currentNode = nodes[currentNode].Right
			}

			totalSteps++
		}

		if currentNode == endNode {
			break
		}
	}

	return totalSteps
}

func part2(input string) int {
	lines := utils.ReadLines(input)
	directions := strings.Split(lines[0], "")
	nodes := getNodes(lines[2:len(lines)])

	startNodes := []string{}
	for name, _ := range nodes {
		nodeName := strings.Split(name, "")
		if nodeName[len(nodeName)-1] == "A" {
			startNodes = append(startNodes, name)
		}
	}

	iterPerKeys := make([]int, len(startNodes))
	dirIndex := 0

	for index := range startNodes {
		totalSteps := 0
		for {

			if dirIndex == len(directions) {
				dirIndex = 0
			}

			if string(startNodes[index][2]) == "Z" {
				break
			}

			if directions[dirIndex] == "L" {
				startNodes[index] = nodes[startNodes[index]].Left
			} else if directions[dirIndex] == "R" {
				startNodes[index] = nodes[startNodes[index]].Right
			}

			totalSteps++
			dirIndex++
		}

		iterPerKeys[index] = totalSteps
	}

	return getLCM(iterPerKeys)
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
