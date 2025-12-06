package day02

import (
	"ii7102/advent-of-code-2025/utils"
	"strconv"
)

func SolvePuzzle1(input string) (res int) {
	return solvePuzzle(isInvalidIDPart1, input)
}

func SolvePuzzle2(input string) (res int) {
	return solvePuzzle(isInvalidIDPart2, input)
}

func solvePuzzle(invalidIDFunc func(string) bool, input string) (res int) {
	for line := range utils.SplitByComma(input) {
		left, right := utils.Range(line)

		for id := left; id <= right; id++ {
			if invalidIDFunc(strconv.Itoa(id)) {
				res += id
			}
		}
	}

	return
}

func isInvalidIDPart1(s string) bool {
	return hasRepeatingPattern(s, 2)
}

func isInvalidIDPart2(s string) bool {
	for i := 2; i*i <= len(s); i++ {
		if hasRepeatingPattern(s, i) || hasRepeatingPattern(s, len(s)/i) {
			return true
		}
	}

	return len(s) > 1 && hasRepeatingPattern(s, len(s))
}

func hasRepeatingPattern(s string, parts int) bool {
	if len(s)%parts != 0 {
		return false
	}

	var (
		partLen = len(s) / parts
		pattern = s[:partLen]
	)

	for i := range parts {
		if s[i*partLen:(i+1)*partLen] != pattern {
			return false
		}
	}

	return true
}
