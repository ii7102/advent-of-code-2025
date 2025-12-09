package day08

import (
	"ii7102/advent-of-code-2025/utils"
	"math"
	"slices"
)

type junctionBox struct {
	x, y, z int
}

type boxDistance struct {
	box1, box2 junctionBox
	distSq     int64
}

func (jb junctionBox) distance(jb2 junctionBox) boxDistance {
	var (
		xDist = jb.x - jb2.x
		yDist = jb.y - jb2.y
		zDist = jb.z - jb2.z
	)

	return boxDistance{
		box1:   jb,
		box2:   jb2,
		distSq: pow2(xDist) + pow2(yDist) + pow2(zDist),
	}
}

//go:inline
func pow2(x int) int64 {
	return int64(x * x)
}

func SolvePuzzle1(input string) int {
	boxes := make([]junctionBox, 0)
	for line := range utils.SplitByNewLine(input) {
		x, y, z := utils.Coordinates(line)
		boxes = append(boxes, junctionBox{x: x, y: y, z: z})
	}

	distances := findDistances(boxes, 1000)
	circuits := findCircuits(distances)

	res := 1
	for range 3 {
		var maxIdx int
		for idx, circuit := range circuits {
			if len(circuits[maxIdx]) < len(circuit) {
				maxIdx = idx
			}
		}

		res *= len(circuits[maxIdx])
		circuits = slices.Delete(circuits, maxIdx, maxIdx+1)
	}

	return res
}

func findDistances(boxes []junctionBox, numOfPairs int) []boxDistance {
	distances := make([]boxDistance, 0, pow2(len(boxes)))

	for i := range boxes {
		for j := i + 1; j < len(boxes); j++ {
			distances = append(distances, boxes[i].distance(boxes[j]))
		}
	}

	for i := range numOfPairs {
		minIdx := i
		for j := i + 1; j < len(distances); j++ {
			if distances[minIdx].distSq > distances[j].distSq {
				minIdx = j
			}
		}
		distances[i], distances[minIdx] = distances[minIdx], distances[i]
	}

	return distances[:numOfPairs]
}

func findCircuits(distances []boxDistance) [][]junctionBox {
	circuitByBox := make(map[junctionBox]int, len(distances)*2)
	circuits := make([][]junctionBox, 0)

	for _, distance := range distances {
		c1, ok1 := circuitByBox[distance.box1]
		c2, ok2 := circuitByBox[distance.box2]

		switch {
		case ok1 && ok2:
			if c1 == c2 {
				continue
			}

			for _, jb := range circuits[c2] {
				circuitByBox[jb] = c1
			}

			circuits[c1] = append(circuits[c1], circuits[c2]...)
			circuits[c2] = nil
		case ok1:
			circuits[c1] = append(circuits[c1], distance.box2)
			circuitByBox[distance.box2] = c1
		case ok2:
			circuits[c2] = append(circuits[c2], distance.box1)
			circuitByBox[distance.box1] = c2
		default:
			circuits = append(circuits, []junctionBox{distance.box1, distance.box2})
			idx := len(circuits) - 1
			circuitByBox[distance.box1] = idx
			circuitByBox[distance.box2] = idx
		}
	}

	return circuits
}

func SolvePuzzle2(input string) int {
	boxes := make([]junctionBox, 0)
	for line := range utils.SplitByNewLine(input) {
		x, y, z := utils.Coordinates(line)
		boxes = append(boxes, junctionBox{x: x, y: y, z: z})
	}

	var maxMinDist boxDistance
	for i, jb1 := range boxes {
		minDist := boxDistance{distSq: math.MaxInt}
		for j, jb2 := range boxes {
			if i == j {
				continue
			}

			if distance := jb1.distance(jb2); distance.distSq < minDist.distSq {
				minDist = distance
			}
		}

		if minDist.distSq > maxMinDist.distSq {
			maxMinDist = minDist
		}
	}

	return maxMinDist.box1.x * maxMinDist.box2.x
}
