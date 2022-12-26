package year2022

import (
	"aocgen/pkg/common"
	"fmt"
	"math"
)

type Day20 struct{}

type encryptedNode struct {
	mixed         int
	value         int
	originalValue int
	order         int
}

func (node encryptedNode) String() string {
	return fmt.Sprintf("%v/%v", node.value, node.order)
}

func (p Day20) PartA(lines []string) any {
	nodes := getNodesFromInput(lines, 1)
	for idx := 0; idx >= 0; idx = firstUnmixed(nodes, 1) {
		nodes = moveNode(nodes, idx)

	}
	zeroIdx := findZero(nodes)
	return nodes[(zeroIdx+1000)%len(nodes)].value + nodes[(zeroIdx+2000)%len(nodes)].value + nodes[(zeroIdx+3000)%len(nodes)].value
}

func (p Day20) PartB(lines []string) any {
	nodes := getNodesFromInput(lines, 811589153)
	for round := 1; round <= 10; round++ {
		for idx := firstUnmixed(nodes, round); idx >= 0; idx = firstUnmixed(nodes, round) {
			nodes = moveNode(nodes, idx)
		}
	}
	zeroIdx := findZero(nodes)
	length := len(nodes)
	return nodes[(zeroIdx+1000)%length].value +
		nodes[(zeroIdx+2000)%length].value +
		nodes[(zeroIdx+3000)%length].value
}

func findZero(nodes []encryptedNode) int {
	for k, n := range nodes {
		if n.originalValue == 0 {
			return k
		}
	}
	return -1
}

func moveNode(nodes []encryptedNode, idx int) []encryptedNode {
	nodeToMove := nodes[idx]
	nodeToMove.mixed++
	length := len(nodes) - 1
	steps := nodeToMove.value
	if steps == 0 {
		nodes[idx].mixed++
		return nodes
	}
	newIdx := (idx + steps) % length
	if newIdx < 0 {
		newIdx += length
	}
	newNodes := common.RemoveIndex(nodes, idx)
	return common.InsertAtIndex(newNodes, newIdx, nodeToMove)
}

func firstUnmixed(nodes []encryptedNode, round int) int {
	min := math.MaxInt
	minIdx := -1
	for i, n := range nodes {
		if n.mixed < round && n.order < min {
			min = n.order
			minIdx = i
		}
	}
	return minIdx
}

func getNodesFromInput(lines []string, key int) []encryptedNode {
	ints := common.GetInts(lines)
	out := make([]encryptedNode, 0)
	for k, i := range ints {
		out = append(out, encryptedNode{mixed: 0, value: i * key, originalValue: i, order: k})
	}
	return out
}
