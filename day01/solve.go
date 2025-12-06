package day01

import (
	"ii7102/advent-of-code-2025/utils"
)

const (
	startDial   = 50
	left, right = 'L', 'R'
)

func SolvePuzzle1(input string) (res int) {
	dial := startDial

	for line := range utils.SplitByNewLine(input) {
		direction, rotation := directionAndRotation(line)

		dial = applyRotation(direction, dial, rotation)
		dial = normalizeDial(dial)

		if dial == 0 {
			res++
		}
	}

	return
}

func SolvePuzzle2(input string) (res int) {
	dial := startDial

	for line := range utils.SplitByNewLine(input) {
		direction, rotation := directionAndRotation(line)

		if direction == left && dial > 0 && rotation >= dial { // special case for left direction
			res++
		}

		dial = applyRotation(direction, dial, rotation)

		res += abs(dial) / 100

		dial = normalizeDial(dial)
	}

	return
}

func normalizeDial(dial int) int {
	return (dial%100 + 100) % 100
}

func directionAndRotation(line string) (direction byte, rotation int) {
	direction = line[0]
	rotation = utils.ToInt(line[1:])

	return
}

func applyRotation(direction byte, dial, rotation int) int {
	switch direction {
	case right:
		return dial + rotation
	case left:
		return dial - rotation
	default:
		panic("invalid direction")
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
