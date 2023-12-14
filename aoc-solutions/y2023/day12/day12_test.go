package day12

import (
	"github.com/stretchr/testify/assert"
	"github.com/teodorpopa/advent-of-code-go/utils"
	"testing"
)

func TestDay12First(t *testing.T) {
	r := Day12First(utils.ReadFile("input_test.txt"))
	assert.Equal(t, 21, r)
}

func TestDay12Second(t *testing.T) {
	r := Day12Second(utils.ReadFile("input_test.txt"))
	assert.Equal(t, 525152, r)
}