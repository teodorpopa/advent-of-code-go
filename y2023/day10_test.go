package y2023

import (
	"github.com/stretchr/testify/assert"
	"github.com/teodorpopa/advent-of-code-go/utils"
	"testing"
)

func TestDay10First(t *testing.T) {
	r := Day10First(utils.ReadFile("test_input/day10.txt"))
	assert.Equal(t, 8, r)
}

func TestDay10Second(t *testing.T) {
	r := Day10Second(utils.ReadFile("test_input/day10_part2.txt"))
	assert.Equal(t, 10, r)
}
