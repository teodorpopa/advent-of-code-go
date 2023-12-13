package y2023

import (
	"github.com/stretchr/testify/assert"
	"github.com/teodorpopa/advent-of-code-go/utils"
	"testing"
)

func TestDay13First(t *testing.T) {
	r := Day13First(utils.ReadFile("test_input/day13.txt"))
	assert.Equal(t, 405, r)
}

func TestDay13Second(t *testing.T) {
	r := Day13Second(utils.ReadFile("test_input/day13.txt"))
	assert.Equal(t, 400, r)
}
