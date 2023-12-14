package day07

import (
	"github.com/stretchr/testify/assert"
	"github.com/teodorpopa/advent-of-code-go/utils"
	"testing"
)

func TestDay07First(t *testing.T) {
	r := Day07First(utils.ReadFile("input_test.txt"))
	assert.Equal(t, 6440, r)
}

func TestDay07Second(t *testing.T) {
	r := Day07Second(utils.ReadFile("input_test.txt"))
	assert.Equal(t, 5905, r)
}
