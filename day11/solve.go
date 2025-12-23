package day11

import (
	"ii7102/advent-of-code-2025/utils"
	"slices"
	"strings"
)

type Node struct {
	prevNodes []string
	nextNodes []string

	pathCount int
}

func SolvePuzzle1(input string) int {
	nodes := initNodes(input)

	return pathCount(nodes, "you", "out")
}

func SolvePuzzle2(input string) int {
	nodes := initNodes(input)
	fromSvrToFft := pathCount(nodes, "svr", "fft")

	fromDacToOut := pathCount(nodes, "dac", "out")

	fromDacToFft := pathCount(nodes, "dac", "fft")
	fromFftToDac := pathCount(nodes, "fft", "dac")

	return (fromSvrToFft - fromDacToFft) * fromFftToDac * fromDacToOut
}

func pathCount(nodesOriginal map[string]*Node, from, to string) int {
	nodes := make(map[string]*Node)

	for name, node := range nodesOriginal {
		nodes[name] = &Node{prevNodes: slices.Clone(node.prevNodes), nextNodes: slices.Clone(node.nextNodes), pathCount: node.pathCount}
	}

	nodes[from].pathCount = 1

	for len(nodes[to].prevNodes) != 0 {
		for name, node := range nodes {
			if len(node.prevNodes) != 0 || name == to {
				continue
			}

			for _, nextNode := range node.nextNodes {
				nodes[nextNode].pathCount += node.pathCount

				idx := slices.Index(nodes[nextNode].prevNodes, name)
				nodes[nextNode].prevNodes = slices.Delete(nodes[nextNode].prevNodes, idx, idx+1)
			}

			delete(nodes, name)
		}
	}

	return nodes[to].pathCount
}

func initNodes(input string) map[string]*Node {
	nodes := make(map[string]*Node)

	for line := range utils.SplitByNewLine(input) {
		strSplit := strings.Split(line, ":")

		from := strSplit[0]
		nodes[from] = &Node{}

		for to := range utils.SplitByEmptySpace(strSplit[1]) {
			if to == "" {
				continue
			}

			nodes[from].nextNodes = append(nodes[from].nextNodes, to)
		}
	}

	nodes["out"] = &Node{}

	for name, node := range nodes {
		for _, to := range node.nextNodes {
			nodes[to].prevNodes = append(nodes[to].prevNodes, name)
		}
	}

	return nodes
}
