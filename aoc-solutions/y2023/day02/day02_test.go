package day02

import (
	"github.com/stretchr/testify/assert"
	"github.com/teodorpopa/advent-of-code-go/utils"
	"testing"
)

func TestDay02First(t *testing.T) {
	r := part1(utils.ReadFile("input_test.txt"))
	assert.Equal(t, 8, r)
}

func TestDay02Second(t *testing.T) {
	r := part2(utils.ReadFile("input_test.txt"))
	assert.Equal(t, 2286, r)
}
