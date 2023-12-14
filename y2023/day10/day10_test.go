package day10

import (
	"github.com/stretchr/testify/assert"
	"github.com/teodorpopa/advent-of-code-go/utils"
	"testing"
)

func TestDayPart1(t *testing.T) {
	r := part1(utils.ReadFile("input_test.txt"))
	assert.Equal(t, 8, r)
}

func TestDayPart2(t *testing.T) {
	r := part2(utils.ReadFile("input_test_2.txt"))
	assert.Equal(t, 10, r)
}
