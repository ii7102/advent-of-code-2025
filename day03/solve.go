package day03

import (
	"ii7102/advent-of-code-2025/utils"
	"strings"
)

func SolvePuzzle1(input string) (res int) {
	return solvePuzzle(input, 2)
}

func SolvePuzzle2(input string) (res int) {
	return solvePuzzle(input, 12)
}

func solvePuzzle(input string, length int) (res int) {
	for line := range utils.SplitByNewLine(input) {
		if strings.ContainsFunc(line, isNotDigit) {
			panic("invalid line: " + line)
		}

		res += biggestNumber(line, length)
	}

	return
}

func biggestNumber(line string, length int) (res int) {
	var index int

	for currDigit := length; currDigit > 0; currDigit-- {
		var maxDigit byte

		for i := index; i <= len(line)-currDigit; i++ {
			if line[i] > maxDigit {
				maxDigit = line[i]
				index = i + 1
			}
		}

		res = res*10 + int(maxDigit-'0')
	}

	return
}

func isNotDigit(r rune) bool {
	return r < '0' || r > '9'
}
