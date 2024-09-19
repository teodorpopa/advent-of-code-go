package day15

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/teodorpopa/advent-of-code-go/utils"
)

func TestDayPart1(t *testing.T) {
	r := part1(utils.ReadFile("input_test.txt"))
	assert.Equal(t, 62842880, r)
}

func TestDayPart2(t *testing.T) {
	r := part2(utils.ReadFile("input_test.txt"))
	assert.Equal(t, 57600000, r)
}
