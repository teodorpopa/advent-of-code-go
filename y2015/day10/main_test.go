package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/teodorpopa/advent-of-code-go/utils"
	"testing"
)

const resultPart1 = 6
const resultPart2 = 6

func TestDayPart1(t *testing.T) {
	r := part1(utils.ReadFile("input_test.txt"), 5)
	assert.Equal(t, r, resultPart1)
}

func TestDayPart2(t *testing.T) {
	r := part1(utils.ReadFile("input_test.txt"), 5)
	assert.Equal(t, r, resultPart2)
}