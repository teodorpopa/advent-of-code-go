package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/teodorpopa/advent-of-code-go/utils"
)

const resultPart1 = 7
const resultPart2 = 33

func TestDayPart1(t *testing.T) {
	r := part1(utils.ReadFile("input_test.txt"))
	assert.Equal(t, resultPart1, r)
}

func TestDayPart2(t *testing.T) {
	r := part2(utils.ReadFile("input_test.txt"))
	assert.Equal(t, resultPart2, r)
}
