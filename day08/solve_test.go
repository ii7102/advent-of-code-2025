package day08_test

import (
	"ii7102/advent-of-code-2025/day08"
	"ii7102/advent-of-code-2025/utils"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSolvePuzzle1(t *testing.T) {
	input := utils.ReadInput(t)
	result := day08.SolvePuzzle1(input)
	assert.Equal(t, utils.Result1Day8, result)
}

func TestSolvePuzzle2(t *testing.T) {
	input := utils.ReadInput(t)
	result := day08.SolvePuzzle2(input)
	assert.Equal(t, utils.Result2Day8, result)
}
