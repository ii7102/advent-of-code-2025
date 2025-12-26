package day12

import (
	"fmt"
	"ii7102/advent-of-code-2025/utils"
	"strings"
)

type cords struct {
	x, y int
}

type present struct {
	areaSize          int
	shape             []cords
	shapeRotatedLeft  []cords
	shapeRotatedRight []cords
	shapeFlipped      []cords
}

func SolvePuzzle1(input string) (res int) {
	presents := make(map[int]present)

	var (
		gridIdx    int
		presentIdx int
	)

	for line := range utils.SplitByNewLine(input) {
		if line == "" {
			continue
		}

		switch {
		case strings.Contains(line, "x"):
			if canPresentsFit(line, presents) {
				res++
			}
		case strings.Contains(line, ":"):
			presentIdx = utils.ToInt(string(line[0]))
			gridIdx = 0
		default:
			presents[presentIdx] = handlePresentShape(line, presents[presentIdx], gridIdx)
			gridIdx++
		}
	}

	return
}

func handlePresentShape(line string, present present, gridIdx int) present {
	for i, ch := range line {
		if ch == '.' {
			continue
		}

		present.shape = append(present.shape, cords{x: gridIdx, y: i})
		present.shapeRotatedLeft = append(present.shapeRotatedLeft, cords{x: 2 - i, y: gridIdx})
		present.shapeRotatedRight = append(present.shapeRotatedRight, cords{x: i, y: 2 - gridIdx})
		present.shapeFlipped = append(present.shapeFlipped, cords{x: 2 - gridIdx, y: i})
		present.areaSize++
	}

	return present
}

func canPresentsFit(line string, presents map[int]present) bool {
	strSplit := strings.Split(line, ":")

	grid := strings.Split(strSplit[0], "x")
	gridX, gridY := utils.ToInt(grid[0]), utils.ToInt(grid[1])

	quantityOfPresents := strings.Split(strSplit[1], " ")[1:]
	if len(quantityOfPresents) != 6 {
		panic(fmt.Sprintf("invalid quantity %d of presents", len(quantityOfPresents)))
	}

	emptyArea := gridX * gridY
	for i, presentQuantity := range quantityOfPresents {
		emptyArea -= utils.ToInt(presentQuantity) * presents[i].areaSize
	}

	return emptyArea >= 0
}

func SolvePuzzle2(input string) (res int) {
	return
}
