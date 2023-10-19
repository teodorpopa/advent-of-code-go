package y2015

import (
	"github.com/stretchr/testify/assert"
	"github.com/teodorpopa/advent-of-code-go/utils"
	"testing"
)

func TestDay09First(t *testing.T) {
	r := Day09First(utils.ReadFile("test_input/day09.txt"))
	assert.Equal(t, r, 605)
}

func TestDay09Second(t *testing.T) {
	r := Day09Second(utils.ReadFile("test_input/day09.txt"))
	assert.Equal(t, r, 982)
}
