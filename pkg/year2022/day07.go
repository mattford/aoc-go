package year2022

import (
	"math"
	"strconv"
	"strings"
)

type Day07 struct{}

type treeNode struct {
	nodeType string
	size     int
	children map[string]*treeNode
	parent   *treeNode
}

func (p Day07) PartA(lines []string) any {
	tree := getDirTree(lines)
	return getDirsUnder(tree, 100000)
}

func (p Day07) PartB(lines []string) any {
	tree := getDirTree(lines)
	requiredSize := 30000000 - (70000000 - tree.size)
	return findSmallestToDelete(tree, requiredSize)
}

func findSmallestToDelete(tree *treeNode, sizeRequired int) int {
	smallest := math.MaxInt
	if tree.nodeType != "dir" {
		return smallest
	}
	if tree.size >= sizeRequired {
		smallest = tree.size
	}
	for _, node := range tree.children {
		if node.nodeType == "dir" && node.size >= sizeRequired && node.size < smallest {
			smallest = node.size
		}
		subSmallest := findSmallestToDelete(node, sizeRequired)
		if subSmallest < smallest {
			smallest = subSmallest
		}
	}
	return smallest
}

func getDirsUnder(tree *treeNode, sizeThreshold int) int {
	total := 0
	for _, node := range tree.children {
		if node.nodeType == "dir" && node.size <= sizeThreshold {
			total += node.size
		}
		total += getDirsUnder(node, sizeThreshold)
	}
	return total
}

func getDirTree(lines []string) *treeNode {
	var rootNode = &treeNode{
		nodeType: "dir",
		size:     0,
		children: make(map[string]*treeNode),
	}
	var currentNode = rootNode
	for _, line := range lines {
		parts := strings.Split(strings.TrimLeft(line, "$ "), " ")
		switch parts[0] {
		case "cd":
			if parts[1] == "/" {
				currentNode = rootNode
			} else if parts[1] == ".." {
				currentNode = currentNode.parent
			} else {
				nextNode, ok := currentNode.children[parts[1]]
				if !ok {
					nextNode = &treeNode{
						nodeType: "dir",
						size:     0,
						children: make(map[string]*treeNode),
						parent:   currentNode,
					}
					currentNode.children[parts[1]] = nextNode
				}
				currentNode = nextNode
			}
		case "dir":
			currentNode.children[parts[1]] = &treeNode{
				nodeType: "dir",
				size:     0,
				children: make(map[string]*treeNode),
				parent:   currentNode,
			}
		case "ls":
			continue
		default:
			size, _ := strconv.Atoi(parts[0])
			fileName := parts[1]
			currentNode.children[fileName] = &treeNode{
				nodeType: "file",
				size:     size,
			}
			currentNode.size += size
			parent := currentNode.parent
			for parent != nil {
				parent.size += size
				parent = parent.parent
			}

		}
	}
	return rootNode
}
