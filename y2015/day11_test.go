package y2015

import (
	"github.com/stretchr/testify/assert"
	"github.com/teodorpopa/advent-of-code-go/utils"
	"testing"
)

func TestDay11First(t *testing.T) {
	r := Day11First(utils.ReadFile("test_input/day11.txt"))
	assert.Equal(t, r, "abcdffaa")
}

func TestDay11Second(t *testing.T) {
	r := Day11Second(utils.ReadFile("test_input/day11.txt"))
	assert.Equal(t, r, "abcdffbb")
}
