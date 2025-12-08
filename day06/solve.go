package day06

import (
	"ii7102/advent-of-code-2025/utils"
	"strings"
)

func SolvePuzzle1(input string) int {
	return solvePuzzle(input, getNumbersHorizontally)
}

func getNumbersHorizontally(lines []string) (numbers [][]int) {
	for _, line := range lines {
		row := make([]int, 0)
		for strNum := range utils.SplitByEmptySpace(line) {
			if strNum != "" {
				row = append(row, utils.ToInt(strNum))
			}
		}

		numbers = append(numbers, row)
	}

	return utils.Transpose(numbers)
}

func SolvePuzzle2(input string) int {
	return solvePuzzle(input, getNumbersVertically)
}

func getNumbersVertically(lines []string) (numbers [][]int) {
	lines = utils.TransposeStrings(lines)
	var idx int

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			idx++
			continue
		}

		if len(numbers) == idx {
			numbers = append(numbers, make([]int, 0, 1))
		}

		numbers[idx] = append(numbers[idx], utils.ToInt(line))
	}

	return numbers
}

func solvePuzzle(input string, getNumbers func([]string) [][]int) (res int) {
	var (
		lines = strings.Split(input, "\r\n")
		opIdx = len(lines) - 1
		nums  = getNumbers(lines[:opIdx])
	)

	lines[opIdx] = strings.ReplaceAll(lines[opIdx], " ", "")
	for i, ch := range lines[opIdx] {
		switch ch {
		case '+':
			for _, num := range nums[i] {
				res += num
			}
		case '*':
			prod := 1
			for _, num := range nums[i] {
				prod *= num
			}
			res += prod
		}
	}

	return
}
