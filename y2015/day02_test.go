package y2015

import (
	"github.com/stretchr/testify/assert"
	"github.com/teodorpopa/advent-of-code-go/utils"
	"testing"
)

const d2_1 = 101
const d2_2 = 48

func TestDay02First(t *testing.T) {
	r := Day02First(utils.ReadFile("test_input/day02.txt"))
	assert.Equal(t, r, d2_1)
}

func TestDay02Second(t *testing.T) {
	r := Day02Second(utils.ReadFile("test_input/day02.txt"))
	assert.Equal(t, r, d2_2)
}
