package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/teodorpopa/advent-of-code-go/utils"
	"testing"
)

const resultPart1 = 1000000
const resultPart2 = 1000000

func TestDayPart1(t *testing.T) {
	r := part1(utils.ReadFile("input_test.txt"))
	assert.Equal(t, r, resultPart1)
}

func TestDayPart2(t *testing.T) {
	r := part2(utils.ReadFile("input_test.txt"))
	assert.Equal(t, r, resultPart2)
}