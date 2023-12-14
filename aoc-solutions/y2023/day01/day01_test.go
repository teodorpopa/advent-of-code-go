package day01

import (
	"github.com/stretchr/testify/assert"
	"github.com/teodorpopa/advent-of-code-go/utils"
	"testing"
)

func TestDay01First(t *testing.T) {
	r := Day01First(utils.ReadFile("input_test.txt"))
	assert.Equal(t, r, 142)
}

func TestDay01Second(t *testing.T) {
	r := Day01Second(utils.ReadFile("input_test_2.txt"))
	assert.Equal(t, r, 281)
}
