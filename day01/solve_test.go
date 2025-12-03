package day01_test

import (
	"ii7102/advent-of-code-2025/day01"
	"ii7102/advent-of-code-2025/utils"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSolvePuzzle1(t *testing.T) {
	input := utils.ReadInput(t)
	result := day01.SolvePuzzle1(input)
	assert.Equal(t, utils.Result1Day1, result)
}

func TestSolvePuzzle2(t *testing.T) {
	input := utils.ReadInput(t)
	result := day01.SolvePuzzle2(input)
	assert.Equal(t, utils.Result2Day1, result)
}
