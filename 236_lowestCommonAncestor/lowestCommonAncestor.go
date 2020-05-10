package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	pd := make(map[int]int)
	qd := make(map[int]int)

	pfind := false
	qfind := false

	check(root, pd, qd, 0, p.Val, q.Val, &pfind, &qfind)

	find := false
	dep := 0
	for find == false {

		pt, pOk := pd[dep]
		qt, qOk := qd[dep]

		if !pOk || !qOk || pt != qt {
			return &TreeNode{pd[dep-1], nil, nil}
		} else {
			dep++
		}
	}

	return root
}

func check(node *TreeNode, pd, qd map[int]int, dep, p, q int, pfind, qfind *bool) {

	if !*pfind {
		pd[dep] = node.Val

		if node.Val == p {
			*pfind = true
		}
	}

	if !*qfind {
		qd[dep] = node.Val

		if node.Val == q {
			*qfind = true
		}
	}

	if *pfind && *qfind {
		return
	}

	if node.Left != nil {

		check(node.Left, pd, qd, dep+1, p, q, pfind, qfind)
	}

	if node.Right != nil {

		check(node.Right, pd, qd, dep+1, p, q, pfind, qfind)
	}
}
