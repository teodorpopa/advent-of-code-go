package y2023

import (
	"github.com/stretchr/testify/assert"
	"github.com/teodorpopa/advent-of-code-go/utils"
	"testing"
)

func TestDay05First(t *testing.T) {
	r := Day05First(utils.ReadFile("test_input/day05.txt"))
	assert.Equal(t, 35, r)
}

func TestDay05Second(t *testing.T) {
	r := Day05Second(utils.ReadFile("test_input/day05.txt"))
	assert.Equal(t, 46, r)
}
