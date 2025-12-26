package day10

import (
	"ii7102/advent-of-code-2025/utils"
	"math"
	"slices"
	"strings"
)

type indicatorLights struct {
	onLights []bool
}

func (il indicatorLights) isOff() bool {
	for _, on := range il.onLights {
		if on {
			return false
		}
	}

	return true
}

type button struct {
	toggle []int
}

type joltageRequirement struct {
	counters []int
}

func (jr joltageRequirement) requirementMet() bool {
	for _, counter := range jr.counters {
		if counter > 0 {
			return false
		}
	}

	return true
}

func parseInput(input string) (indicatorLights, []button, joltageRequirement) {
	strSplit := strings.Split(input, " ")
	for i := range len(strSplit) {
		strSplit[i] = strSplit[i][1 : len(strSplit[i])-1]
	}

	var (
		indicatorLights    = indicatorLights{}
		buttons            = make([]button, len(strSplit)-2)
		joltageRequirement = joltageRequirement{}
	)

	for _, ch := range strSplit[0] {
		indicatorLights.onLights = append(indicatorLights.onLights, ch == '#')
	}

	for i := range len(strSplit) - 2 {
		for ch := range utils.SplitByComma(strSplit[i+1]) {
			buttons[i].toggle = append(buttons[i].toggle, utils.ToInt(ch))
		}
	}

	for ch := range utils.SplitByComma(strSplit[len(strSplit)-1]) {
		joltageRequirement.counters = append(joltageRequirement.counters, utils.ToInt(ch))
	}

	return indicatorLights, buttons, joltageRequirement
}

func SolvePuzzle1(input string) (res int) {
	for line := range utils.SplitByNewLine(input) {
		indicatorLights, buttons, _ := parseInput(line)
		res += solvePart1(indicatorLights, buttons, 0, 0)
	}

	return
}

func solvePart1(finalIL indicatorLights, buttons []button, idx, pressedCount int) int {
	if idx == len(buttons) {
		if finalIL.isOff() {
			return pressedCount
		}

		return math.MaxInt
	}

	for _, light := range buttons[idx].toggle {
		finalIL.onLights[light] = !finalIL.onLights[light]
	}

	res := solvePart1(finalIL, buttons, idx+1, pressedCount+1)

	for _, light := range buttons[idx].toggle {
		finalIL.onLights[light] = !finalIL.onLights[light]
	}

	res = min(res, solvePart1(finalIL, buttons, idx+1, pressedCount))

	return res
}

func SolvePuzzle2(input string) (res int) {
	for line := range utils.SplitByNewLine(input) {
		_, buttons, jr := parseInput(line)
		res += solvePart2(jr, buttons)
	}

	return
}

func solvePart2(targetJR joltageRequirement, buttons []button) (res int) {
	if targetJR.requirementMet() {
		return 0
	}

	targetParity := make([]bool, len(targetJR.counters))
	for i := range targetParity {
		targetParity[i] = targetJR.counters[i]%2 == 1
	}

	allSolutions := findAllParitySolutions(targetParity, buttons)

	res = math.MaxInt
	if len(allSolutions) == 0 {
		return
	}

solutionsLoop:
	for _, presses := range allSolutions {
		newTargetJR := joltageRequirement{counters: slices.Clone(targetJR.counters)}

		var parityPresses int
		for j, count := range presses {
			if count == 0 {
				continue
			}

			parityPresses += count
			for _, idx := range buttons[j].toggle {
				newTargetJR.counters[idx] -= count
			}
		}

		if parityPresses >= res {
			continue
		}

		for i, counter := range newTargetJR.counters {
			if counter < 0 {
				continue solutionsLoop
			}

			newTargetJR.counters[i] /= 2
		}

		sub := solvePart2(newTargetJR, buttons)
		if sub == math.MaxInt {
			continue
		}

		res = min(res, parityPresses+2*sub)
	}

	return
}

func findAllParitySolutions(target []bool, buttons []button) (solutions [][]int) {
	var dfs func(int, []bool, []int)
	dfs = func(idx int, cur []bool, counts []int) {
		if idx == len(buttons) {
			match := true
			for i := range len(target) {
				match = match && cur[i] == target[i]
			}

			if match {
				solutions = append(solutions, slices.Clone(counts))
			}

			return
		}

		dfs(idx+1, cur, counts)

		for _, light := range buttons[idx].toggle {
			cur[light] = !cur[light]
		}

		counts[idx] = 1
		dfs(idx+1, cur, counts)
		counts[idx] = 0

		for _, light := range buttons[idx].toggle {
			cur[light] = !cur[light]
		}
	}

	dfs(0, make([]bool, len(target)), make([]int, len(buttons)))

	return
}
