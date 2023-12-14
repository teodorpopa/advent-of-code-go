package day08

import (
	"github.com/stretchr/testify/assert"
	"github.com/teodorpopa/advent-of-code-go/utils"
	"testing"
)

func TestDay08First(t *testing.T) {
	r := Day08First(utils.ReadFile("input_test.txt"))
	assert.Equal(t, r, 12)
}

func TestDay08Second(t *testing.T) {
	r := Day08Second(utils.ReadFile("input_test.txt"))
	assert.Equal(t, r, 19)
}
