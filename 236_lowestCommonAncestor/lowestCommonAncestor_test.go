package leetcode

import (
	"strconv"
	"strings"
	"testing"
)

func CreateTree(rootStr string, pVal int, qVal int) (*TreeNode, *TreeNode, *TreeNode) {

	var root *TreeNode
	var p *TreeNode
	var q *TreeNode

	nodeStrs := strings.Split(rootStr, ",")

	nodes := make([]*TreeNode, len(nodeStrs))
	index := 0

	for i, v := range nodeStrs {

		var node *TreeNode

		if v == "null" {
			node = nil
		} else {
			nodeVal, _ := strconv.Atoi(v)
			node = &TreeNode{nodeVal, nil, nil}

			if node.Val == pVal {
				p = node
			}
			if node.Val == qVal {
				q = node
			}
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

	return root, p, q
}

func TestExample_1(t *testing.T) {
	rootStr := "3,5,1,6,2,0,8,null,null,7,4"
	pVal := 5
	qVal := 1

	ancestorVal := 3

	root, p, q := CreateTree(rootStr, pVal, qVal)

	ancestor := lowestCommonAncestor(root, p, q)

	if ancestor.Val != ancestorVal {
		t.Errorf("ancestor.val is %d , not %d", ancestor.Val, ancestorVal)
	}
}

func TestExample_2(t *testing.T) {
	rootStr := "3,5,1,6,2,0,8,null,null,7,4"
	pVal := 5
	qVal := 4

	ancestorVal := 5

	root, p, q := CreateTree(rootStr, pVal, qVal)

	ancestor := lowestCommonAncestor(root, p, q)

	if ancestor.Val != ancestorVal {
		t.Errorf("ancestor.val is %d , not %d", ancestor.Val, ancestorVal)
	}
}

func TestExample_4(t *testing.T) {
	rootStr := "2,null,1"
	pVal := 2
	qVal := 1

	ancestorVal := 2

	root, p, q := CreateTree(rootStr, pVal, qVal)

	ancestor := lowestCommonAncestor(root, p, q)

	if ancestor.Val != ancestorVal {
		t.Errorf("ancestor.val is %d , not %d", ancestor.Val, ancestorVal)
	}
}

func TestExample_29(t *testing.T) {
	rootStr := "-1,0,3,-2,4,null,null,8"
	pVal := 8
	qVal := -1

	ancestorVal := -1

	root, p, q := CreateTree(rootStr, pVal, qVal)

	ancestor := lowestCommonAncestor(root, p, q)

	if ancestor.Val != ancestorVal {
		t.Errorf("ancestor.val is %d , not %d", ancestor.Val, ancestorVal)
	}
}
