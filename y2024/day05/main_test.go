package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/teodorpopa/advent-of-code-go/utils"
	"testing"
)

const resultPart1 = 143
const resultPart2 = 123

func TestDayPart1(t *testing.T) {
	r := resolvePart1(utils.ReadFile("input_test.txt"))
	assert.Equal(t, resultPart1, r)
}

func TestDayPart2(t *testing.T) {
	r := resolvePart2(utils.ReadFile("input_test.txt"))
	assert.Equal(t, resultPart2, r)
}
