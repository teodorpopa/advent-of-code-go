package y2023

import (
	"github.com/stretchr/testify/assert"
	"github.com/teodorpopa/advent-of-code-go/utils"
	"testing"
)

func TestDay03First(t *testing.T) {
	r := Day03First(utils.ReadFile("test_input/day03.txt"))
	assert.Equal(t, 4361, r)
}

func TestDay03Second(t *testing.T) {
	r := Day03Second(utils.ReadFile("test_input/day03.txt"))
	assert.Equal(t, 467835, r)
}
