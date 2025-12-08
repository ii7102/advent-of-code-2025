package day07

import (
	"ii7102/advent-of-code-2025/utils"
)

func SolvePuzzle1(input string) (res int) {
	grid := make([][]byte, 0)
	for line := range utils.SplitByNewLine(input) {
		grid = append(grid, []byte(line))
	}

	for i := range grid[:len(grid)-1] {
		for j, ch := range grid[i] {
			switch ch {
			case 'S', '|':
				switch grid[i+1][j] {
				case '.':
					grid[i+1][j] = '|'
				case '^':
					grid[i+1][j-1] = '|'
					grid[i+1][j+1] = '|'
					res++
				}
			}
		}
	}

	return
}

func SolvePuzzle2(input string) (res int) {
	var (
		grid = make([][]byte, 0)
		nums = make([][]int, 0)
	)

	for line := range utils.SplitByNewLine(input) {
		grid = append(grid, []byte(line))
		nums = append(nums, make([]int, len(line)))
	}

	for i := range nums[:len(nums)-1] {
		for j, num := range nums[i] {
			switch grid[i][j] {
			case 'S':
				grid[i+1][j] = '|'
				nums[i+1][j] = 1
			case '|':
				switch grid[i+1][j] {
				case '.', '|':
					grid[i+1][j] = '|'
					nums[i+1][j] += num
				case '^':
					grid[i+1][j-1] = '|'
					grid[i+1][j+1] = '|'
					nums[i+1][j-1] += num
					nums[i+1][j+1] += num
				}
			}
		}
	}

	for _, num := range nums[len(nums)-1] {
		res += num
	}

	return
}
