package day10

import (
	"github.com/stretchr/testify/assert"
	"github.com/teodorpopa/advent-of-code-go/utils"
	"testing"
)

func TestDay10First(t *testing.T) {
	r := part1(utils.ReadFile("input_test.txt"), 5)
	assert.Equal(t, r, 6)
}

func TestDay10Second(t *testing.T) {
	r := part1(utils.ReadFile("input_test.txt"), 5)
	assert.Equal(t, r, 6)
}
