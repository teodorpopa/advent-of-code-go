package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/teodorpopa/advent-of-code-go/utils"
	"testing"
)

const resultPart1 = 55312
const resultPart2 = 0

func TestDayPart1(t *testing.T) {
	r := part1(utils.ReadFile("input_test.txt"), 25)
	assert.Equal(t, resultPart1, r)
}

func TestDayPart2(t *testing.T) {
	r := part1(utils.ReadFile("input_test.txt"), 75)
	assert.Equal(t, resultPart2, r)
}
