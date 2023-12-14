package day09

import (
	"github.com/stretchr/testify/assert"
	"github.com/teodorpopa/advent-of-code-go/utils"
	"testing"
)

func TestDay09First(t *testing.T) {
	r := part1(utils.ReadFile("input_test.txt"))
	assert.Equal(t, r, 605)
}

func TestDay09Second(t *testing.T) {
	r := part2(utils.ReadFile("input_test.txt"))
	assert.Equal(t, r, 982)
}
