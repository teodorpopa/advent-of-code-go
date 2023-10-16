package y2015

import (
	"github.com/stretchr/testify/assert"
	"github.com/teodorpopa/advent-of-code-go/utils"
	"testing"
)

const d1_1 = 2
const d1_2 = 11

func TestDay01First(t *testing.T) {
	r := Day01First(utils.ReadFile("test_input/day01.txt"))
	assert.Equal(t, r, d1_1)
}

func TestDay01Second(t *testing.T) {
	r := Day01Second(utils.ReadFile("test_input/day01.txt"))
	assert.Equal(t, r, d1_2)
}
