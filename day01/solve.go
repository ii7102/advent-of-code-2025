package day01

import (
	"strconv"
	"strings"
)

const (
	startDial   = 50
	left, right = 'L', 'R'
)

func SolvePuzzle1(input string) (res int) {
	dial := startDial

	for line := range strings.SplitSeq(input, "\n") {
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

	for line := range strings.SplitSeq(input, "\n") {
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

func directionAndRotation(line string) (byte, int) {
	direction := line[0]
	rotation, err := strconv.Atoi(strings.TrimSpace(line[1:]))
	if err != nil {
		panic(err)
	}

	return direction, rotation
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
