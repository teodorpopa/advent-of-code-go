package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/teodorpopa/advent-of-code-go/utils"
	"testing"
)

const resultPart1 = 8
const resultPart2 = 10

func TestDayPart1(t *testing.T) {
	r := part1(utils.ReadFile("input_test.txt"))
	assert.Equal(t, resultPart1, r)
}

func TestDayPart2(t *testing.T) {
	r := part2(utils.ReadFile("input_test_2.txt"))
	assert.Equal(t, resultPart2, r)
}
