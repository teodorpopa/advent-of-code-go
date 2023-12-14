package day14

import (
	"github.com/stretchr/testify/assert"
	"github.com/teodorpopa/advent-of-code-go/utils"
	"testing"
)

func TestDay14First(t *testing.T) {
	r := Day14First(utils.ReadFile("input_test.txt"), 1000)
	assert.Equal(t, r, 1120)
}

func TestDay14Second(t *testing.T) {
	r := Day14Second(utils.ReadFile("input_test.txt"), 1000)
	assert.Equal(t, r, 689)
}