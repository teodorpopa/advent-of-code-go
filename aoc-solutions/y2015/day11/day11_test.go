package day11

import (
	"github.com/stretchr/testify/assert"
	"github.com/teodorpopa/advent-of-code-go/utils"
	"testing"
)

func TestDay11First(t *testing.T) {
	r := part1(utils.ReadFile("input_test.txt"))
	assert.Equal(t, r, "abcdffaa")
}

func TestDay11Second(t *testing.T) {
	r := part2(utils.ReadFile("input_test.txt"))
	assert.Equal(t, r, "abcdffbb")
}
