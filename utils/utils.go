package utils

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func CreateTree(treeStr string) *TreeNode {
	var root *TreeNode

	nodeStrs := strings.Split(treeStr, ",")

	nodes := make([]*TreeNode, len(nodeStrs))
	index := 0

	for i, v := range nodeStrs {

		var node *TreeNode

		if v == "null" {
			node = nil
		} else {
			nodeVal, _ := strconv.Atoi(v)
			node = &TreeNode{nodeVal, nil, nil}
		}

		nodes[i] = node

		if i == 0 {
			root = node
			continue
		}

		if i%2 == 0 {
			nodes[index].Right = node
			index++
		} else {
			nodes[index].Left = node
		}
	}

	return root
}
