package y2023

import (
	"github.com/stretchr/testify/assert"
	"github.com/teodorpopa/advent-of-code-go/utils"
	"testing"
)

func TestDay08First(t *testing.T) {
	r := Day08First(utils.ReadFile("test_input/day08.txt"))
	assert.Equal(t, 6, r)
}

func TestDay08Second(t *testing.T) {
	r := Day08Second(utils.ReadFile("test_input/day08_part2.txt"))
	assert.Equal(t, 6, r)
}
