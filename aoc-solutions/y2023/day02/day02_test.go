package day02

import (
	"github.com/stretchr/testify/assert"
	"github.com/teodorpopa/advent-of-code-go/utils"
	"testing"
)

func TestDay02First(t *testing.T) {
	r := Day02First(utils.ReadFile("input_test.txt"))
	assert.Equal(t, 8, r)
}

func TestDay02Second(t *testing.T) {
	r := Day02Second(utils.ReadFile("input_test.txt"))
	assert.Equal(t, 2286, r)
}
