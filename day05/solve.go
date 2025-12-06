package day05

import (
	"ii7102/advent-of-code-2025/utils"
	"slices"
	"strings"
)

type Range struct {
	left, right int
}

func SolvePuzzle1(input string) (res int) {
	ingredientRanges := make([]Range, 0)
	for line := range utils.SplitByNewLine(input) {
		if line == "" {
			continue
		}

		if strings.Contains(line, "-") {
			left, right := utils.Range(line)
			ingredientRanges = append(ingredientRanges, Range{left: left, right: right})
			continue
		}

		ingredient := utils.ToInt(line)
		for _, ingredientRange := range ingredientRanges {
			if ingredient >= ingredientRange.left && ingredient <= ingredientRange.right {
				res++
				break
			}
		}
	}

	return
}

func SolvePuzzle2(input string) (res int) {
	ingredientRanges := make([]Range, 0)
	for line := range utils.SplitByNewLine(input) {
		if line == "" {
			break
		}

		left, right := utils.Range(line)
		ingredientRanges = append(ingredientRanges, Range{left: left, right: right})
	}

	slices.SortFunc(ingredientRanges, func(a, b Range) int {
		if a.left == b.left {
			return b.right - a.right
		}

		return a.left - b.left
	})

	var maxRight int
	for _, r := range ingredientRanges {
		if r.right <= maxRight {
			continue
		}

		if r.left > maxRight {
			res += r.right - r.left + 1
		} else {
			res += r.right - maxRight
		}

		maxRight = r.right
	}

	return
}
