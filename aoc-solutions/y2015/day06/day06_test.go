package day06

import (
	"github.com/stretchr/testify/assert"
	"github.com/teodorpopa/advent-of-code-go/utils"
	"testing"
)

func TestDay06First(t *testing.T) {
	r := Day06First(utils.ReadFile("input_test.txt"))
	assert.Equal(t, r, 1000000)
}

func TestDay06Second(t *testing.T) {
	r := Day06Second(utils.ReadFile("input_test.txt"))
	assert.Equal(t, r, 1000000)
}
