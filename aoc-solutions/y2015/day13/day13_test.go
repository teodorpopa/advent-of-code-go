package day13

import (
	"github.com/stretchr/testify/assert"
	"github.com/teodorpopa/advent-of-code-go/utils"
	"testing"
)

func TestDay13First(t *testing.T) {
	r := Day13First(utils.ReadFile("input_test.txt"))
	assert.Equal(t, r, 330)
}

func TestDay13Second(t *testing.T) {
	r := Day13Second(utils.ReadFile("input_test.txt"))
	assert.Equal(t, r, 286)
}
