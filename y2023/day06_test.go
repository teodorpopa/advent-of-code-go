package y2023

import (
	"github.com/stretchr/testify/assert"
	"github.com/teodorpopa/advent-of-code-go/utils"
	"testing"
)

func TestDay06First(t *testing.T) {
	r := Day06First(utils.ReadFile("test_input/day06.txt"))
	assert.Equal(t, 288, r)
}

func TestDay06Second(t *testing.T) {
	r := Day06Second(utils.ReadFile("test_input/day06.txt"))
	assert.Equal(t, 71503, r)
}
