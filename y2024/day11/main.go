package main

import (
	_ "embed"
	"flag"
	"fmt"
	"github.com/teodorpopa/advent-of-code-go/utils"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string
var cache map[string]int

func part1(input string, steps int) int {
	cache = make(map[string]int)
	nums := []int{}
	parts := strings.Split(input, " ")

	for _, p := range parts {
		nums = append(nums, utils.ToInt(p))
	}

	count := 0
	for i := 0; i < len(nums); i++ {
		count += countStones(nums[i], 0, steps)
	}

	return count
}

func countStones(num int, depth int, max int) int {
	stones := 0
	if depth == max {
		return 1
	}

	str := strconv.Itoa(num)
	key := str + "_" + strconv.Itoa(depth)

	if cache[key] != 0 {
		return cache[key]
	} else if num == 0 {
		stones = countStones(1, depth+1, max)
		cache[key] = stones
		return stones
	}
	if len(str)%2 == 0 {
		mid := len(str) / 2
		left := utils.ToInt(str[0:mid])
		right := utils.ToInt(str[mid:])
		stones = countStones(left, depth+1, max) + countStones(right, depth+1, max)
		cache[key] = stones
		return stones
	}
	stones = countStones(num*2024, depth+1, max)
	cache[key] = stones
	return stones
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "which part tu run? 1 or 2")
	flag.Parse()

	res := 0
	if part == 1 {
		res = part1(input, 25)
	} else {
		res = part1(input, 75)
	}

	fmt.Println("Result: ", res)
}
