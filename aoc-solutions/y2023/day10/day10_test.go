package day10

import (
	"github.com/stretchr/testify/assert"
	"github.com/teodorpopa/advent-of-code-go/utils"
	"testing"
)

func TestDay10First(t *testing.T) {
	r := Day10First(utils.ReadFile("input_test.txt"))
	assert.Equal(t, 8, r)
}

func TestDay10Second(t *testing.T) {
	r := Day10Second(utils.ReadFile("input_test_2.txt"))
	assert.Equal(t, 10, r)
}
