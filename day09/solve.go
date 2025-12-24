package day09

import "ii7102/advent-of-code-2025/utils"

type coordinate struct {
	x, y int
}

func (c coordinate) area(c2 coordinate) int {
	return (1 + abs(c.x-c2.x)) * (1 + abs(c.y-c2.y))
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func SolvePuzzle1(input string) (res int) {
	coords := parseInput(input)

	for i, c1 := range coords {
		for _, c2 := range coords[i+1:] {
			res = max(res, c1.area(c2))
		}
	}

	return
}

func SolvePuzzle2(input string) (res int) {
	coords := parseInput(input)

	for i, c1 := range coords {
		for _, c3 := range coords[i+1:] {
			c2 := coordinate{c1.x, c3.y}
			c4 := coordinate{c3.x, c1.y}

			if isInside(c2, coords) && isInside(c4, coords) && !edgeCutsRect(coords, c1, c3) {
				res = max(res, c1.area(c3))
			}
		}
	}

	return
}

func parseInput(input string) (coords []coordinate) {
	for line := range utils.SplitByNewLine(input) {
		x, y := utils.Coordinates2D(line)
		coords = append(coords, coordinate{x, y})
	}

	return coords
}

func isInside(p coordinate, poly []coordinate) bool {
	n := len(poly)

	for i := range n {
		a, b := poly[i], poly[(i+1)%n]

		if a.x == b.x && p.x == a.x && p.y >= min(a.y, b.y) && p.y <= max(a.y, b.y) {
			return true
		}

		if a.y == b.y && p.y == a.y && p.x >= min(a.x, b.x) && p.x <= max(a.x, b.x) {
			return true
		}

		if (a.y > p.y) != (b.y > p.y) {
			xHit := a.x + (p.y-a.y)*(b.x-a.x)/(b.y-a.y)
			if p.x < xHit {
				return true
			}
		}
	}

	return false
}

func edgeCutsRect(poly []coordinate, c1, c3 coordinate) bool {
	minX, maxX := min(c1.x, c3.x), max(c1.x, c3.x)
	minY, maxY := min(c1.y, c3.y), max(c1.y, c3.y)

	for i := range len(poly) {
		a, b := poly[i], poly[(i+1)%len(poly)]

		if a.y == b.y {
			if a.y > minY && a.y < maxY && min(a.x, b.x) < maxX && max(a.x, b.x) > minX {
				return true
			}

			continue
		}

		if a.x > minX && a.x < maxX && min(a.y, b.y) < maxY && max(a.y, b.y) > minY {
			return true
		}
	}

	return false
}
