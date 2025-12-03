package day05_test

import (
	"ii7102/advent-of-code-2025/day05"
	"ii7102/advent-of-code-2025/utils"
	"testing"
)

func TestSolvePuzzle1(t *testing.T) {
	input := utils.ReadInput(t)
	result := day05.SolvePuzzle1(input)
	utils.LogResult(t, result)
}

func TestSolvePuzzle2(t *testing.T) {
	input := utils.ReadInput(t)
	result := day05.SolvePuzzle2(input)
	utils.LogResult(t, result)
}
