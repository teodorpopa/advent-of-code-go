package day08

import (
	"github.com/stretchr/testify/assert"
	"github.com/teodorpopa/advent-of-code-go/utils"
	"testing"
)

func TestDay08First(t *testing.T) {
	r := part1(utils.ReadFile("input_test.txt"))
	assert.Equal(t, 6, r)
}

func TestDay08Second(t *testing.T) {
	r := part2(utils.ReadFile("input_test_2.txt"))
	assert.Equal(t, 6, r)
}
