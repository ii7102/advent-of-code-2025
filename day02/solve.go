package day02

import (
	"strconv"
	"strings"
)

func SolvePuzzle1(input string) (res int) {
	return solvePuzzle(isInvalidIDPart1, input)
}

func SolvePuzzle2(input string) (res int) {
	return solvePuzzle(isInvalidIDPart2, input)
}

func solvePuzzle(invalidIDFunc func(string) bool, input string) (res int) {
	for line := range strings.SplitSeq(input, ",") {
		leftStr, rightStr, found := strings.Cut(line, "-")
		if !found {
			panic("invalid line: " + line)
		}

		left, err := strconv.Atoi(leftStr)
		if err != nil {
			panic("invalid left: " + leftStr)
		}

		right, err := strconv.Atoi(rightStr)
		if err != nil {
			panic("invalid right: " + rightStr)
		}

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

	partLen := len(s) / parts

	pattern := s[:partLen]
	for i := range parts {
		if s[i*partLen:(i+1)*partLen] != pattern {
			return false
		}
	}

	return true
}
