package day04

import (
	"ii7102/advent-of-code-2025/utils"
)

func SolvePuzzle1(input string) (res int) {
	return solvePuzzle(input, false)
}

func SolvePuzzle2(input string) (res int) {
	return solvePuzzle(input, true)
}

func solvePuzzle(input string, cascade bool) (res int) {
	rolls := initRolls(input)
	adjacentRolls := initAdjacentRolls(rolls)

	for i := 0; i < len(rolls); i++ {
		for j := 0; j < len(rolls[i]); j++ {
			if !rolls[i][j] || adjacentRolls[i][j] > 4 {
				continue
			}

			res++

			if cascade {
				rolls[i][j] = false
				updateAdjacentRolls(adjacentRolls, i, j)
				i, j = updateIndexes(i, j)
			}
		}
	}

	return
}


func initRolls(input string) (rolls [][]bool) {
	rolls = make([][]bool, 0)
	for line := range utils.SplitByNewLine(input) {
		row := make([]bool, 0)
		for _, cell := range line {
			row = append(row, cell == '@')
		}

		rolls = append(rolls, row)
	}

	return
}

func initAdjacentRolls(rolls [][]bool) (adjacentRolls [][]int) {
	adjacentRolls = make([][]int, len(rolls))
	for i, row := range rolls {
		adjacentRolls[i] = make([]int, len(row))
		for j := range row {
			adjacentRolls[i][j] = countRolls(rolls, i, j)
		}
	}

	return
}

func countRolls(rolls [][]bool, i, j int) (count int) {
	for ni := i - 1; ni <= i+1; ni++ {
		for nj := j - 1; nj <= j+1; nj++ {
			if ni < 0 || ni >= len(rolls) || nj < 0 || nj >= len(rolls[ni]) {
				continue
			}

			if rolls[ni][nj] {
				count++
			}
		}
	}

	return
}

func updateAdjacentRolls(adjacentRolls [][]int, i, j int) {
	for ni := i - 1; ni <= i+1; ni++ {
		for nj := j - 1; nj <= j+1; nj++ {
			if ni < 0 || ni >= len(adjacentRolls) || nj < 0 || nj >= len(adjacentRolls[ni]) {
				continue
			}

			adjacentRolls[ni][nj]--
		}
	}
}

func updateIndexes(i, j int) (int, int) {
	return max(i-1, 0), max(j-2, -1)
}
