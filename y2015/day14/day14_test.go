package day14

import (
	"github.com/stretchr/testify/assert"
	"github.com/teodorpopa/advent-of-code-go/utils"
	"testing"
)

func TestDayPart1(t *testing.T) {
	r := part1(utils.ReadFile("input_test.txt"), 1000)
	assert.Equal(t, r, 1120)
}

func TestDayPart2(t *testing.T) {
	r := part2(utils.ReadFile("input_test.txt"), 1000)
	assert.Equal(t, r, 689)
}
