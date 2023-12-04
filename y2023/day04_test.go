package y2023

import (
	"github.com/stretchr/testify/assert"
	"github.com/teodorpopa/advent-of-code-go/utils"
	"testing"
)

func TestDay04First(t *testing.T) {
	r := Day04First(utils.ReadFile("test_input/day04.txt"))
	assert.Equal(t, 13, r)
}

func TestDay04Second(t *testing.T) {
	r := Day04Second(utils.ReadFile("test_input/day04.txt"))
	assert.Equal(t, 30, r)
}
